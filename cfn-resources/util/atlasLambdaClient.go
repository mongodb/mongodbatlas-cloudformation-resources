package util

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/mongodb-forks/digest"

	"github.com/mongodb/mongodbatlas-cloudformation-resources/profile"
)

// lambdaForwardingTransport implements http.RoundTripper
type lambdaForwardingTransport struct {
	lambdaArn    string
	lambdaClient *lambda.Lambda
}

func newLambdaForwardingTransport(req *handler.Request, lambdaArn string) *lambdaForwardingTransport {
	log.Printf("Initializing lambdaForwardingTransport with Lambda ARN: %s", lambdaArn)
	// Extract region from Lambda ARN - TODO: probably not required, remove
	arnParts := strings.Split(lambdaArn, ":")
	region := "us-east-1" // Default
	if len(arnParts) >= 4 {
		region = arnParts[3]
	}
	svc := lambda.New(req.Session, aws.NewConfig().WithRegion(region))

	return &lambdaForwardingTransport{
		lambdaArn:    lambdaArn,
		lambdaClient: svc,
	}
}

// LambdaRequestPayload is sent to the Lambda function
type LambdaRequestPayload struct {
	Method  string            `json:"method"`
	URL     string            `json:"url"`
	Headers map[string]string `json:"headers"`
	Body    string            `json:"body"`
}

// LambdaResponsePayload is returned by the Lambda function
type LambdaResponsePayload struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
}

// This method currently uses extensive logging for POC purpose which should be reduced
func (t *lambdaForwardingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("Entering lambdaForwardingTransport.RoundTrip for URL: %s, Method: %s", req.URL.String(), req.Method)

	headers := make(map[string]string)
	for key, values := range req.Header {
		if len(values) > 0 {
			headers[key] = values[0]
		}
	}
	log.Printf("Captured request headers: %+v", headers)

	var bodyBytes []byte
	if req.Body != nil {
		var err error
		bodyBytes, err = ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			return nil, err
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	} else {
		log.Printf("No request body to read (req.Body is nil)")
	}

	payloadStruct := LambdaRequestPayload{
		Method:  req.Method,
		URL:     req.URL.String(),
		Headers: headers,
		Body:    string(bodyBytes),
	}
	payloadBytes, err := json.Marshal(payloadStruct)
	if err != nil {
		log.Printf("Error marshaling payload: %v", err)
		return nil, err
	}
	log.Printf("Payload to be sent to Lambda: %s", string(payloadBytes))

	input := &lambda.InvokeInput{
		FunctionName: aws.String(t.lambdaArn),
		Payload:      payloadBytes,
	}
	log.Printf("Invoking Lambda with input: %+v", input)
	result, err := t.lambdaClient.Invoke(input)
	if err != nil {
		log.Printf("Error invoking Lambda: %v", err)
		return nil, err
	}
	if result.FunctionError != nil {
		errMsg := "Lambda function error: " + *result.FunctionError
		log.Printf(errMsg)
		return nil, errors.New(errMsg)
	}
	log.Printf("Lambda invocation result: %+v", result)

	var respPayload LambdaResponsePayload
	err = json.Unmarshal(result.Payload, &respPayload)
	if err != nil {
		log.Printf("Error unmarshaling Lambda response payload: %v", err)
		return nil, err
	}
	log.Printf("Parsed Lambda response payload: %+v", respPayload)

	resp := &http.Response{
		StatusCode: respPayload.StatusCode,
		Status:     http.StatusText(respPayload.StatusCode),
		Header:     make(http.Header),
		Body:       ioutil.NopCloser(bytes.NewBufferString(respPayload.Body)),
		Request:    req,
	}
	for key, value := range respPayload.Headers {
		resp.Header.Set(key, value)
	}
	log.Printf("Returning HTTP response from RoundTrip: %+v", resp)
	return resp, nil
}

// This method currently uses extensive logging for POC purpose which should be reduced
func newAtlasV2ClientWithLambdaProxySupport(req *handler.Request, profileName *string, profileNamePrefixRequired bool, lambdaArn *string) (*MongoDBClient, *handler.ProgressEvent) {
	log.Printf("Initializing newAtlasV2Client with profileName: %v", profileName)
	prof, err := profile.NewProfile(req, profileName, profileNamePrefixRequired)
	if err != nil {
		log.Printf("Error creating profile: %v", err)
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound,
		}
	}

	var client *http.Client
	if lambdaArn != nil && *lambdaArn != "" {
		// Chain the digest transport with the lambda forwarding transport
		log.Printf("Using chained digest transport with Lambda forwarding. Lambda ARN: %s", *lambdaArn)
		// Create the Lambda forwarding transport
		lambdaTransport := newLambdaForwardingTransport(req, *lambdaArn)
		// Create the digest transport
		digestTransport := digest.NewTransport(prof.PublicKey, prof.PrivateKey)
		// IMPORTANT: Set the underlying transport to our Lambda transport
		digestTransport.Transport = lambdaTransport
		// Use the digest transport as the client transport
		client = &http.Client{Transport: digestTransport}
	} else {
		log.Printf("Using default digest transport with PublicKey: %s", prof.PublicKey)
		transport := digest.NewTransport(prof.PublicKey, prof.PrivateKey)
		client, err = transport.Client()
		if err != nil {
			log.Printf("Error creating digest transport client: %v", err)
			return nil, &handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
			}
		}
	}

	c := Config{BaseURL: prof.BaseURL, DebugClient: prof.UseDebug()}
	log.Printf("Config initialized: %+v", c)

	sdk20231115002Client, err := c.NewSDKv20231115002Client(client)
	if err != nil {
		log.Printf("Error creating SDKv20231115002Client: %v", err)
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
		}
	}

	sdk20231115014Client, err := c.NewSDKv20231115014Client(client)
	if err != nil {
		log.Printf("Error creating SDKv20231115014Client: %v", err)
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
		}
	}

	sdkV2LatestClient, err := c.NewSDKV2LatestClient(client)
	if err != nil {
		log.Printf("Error creating SDKV2LatestClient: %v", err)
		return nil, &handler.ProgressEvent{
			OperationStatus:  handler.Failed,
			Message:          err.Error(),
			HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest,
		}
	}

	clients := &MongoDBClient{
		Atlas20231115002: sdk20231115002Client,
		Atlas20231115014: sdk20231115014Client,
		AtlasSDK:         sdkV2LatestClient,
		Config:           &c,
	}
	log.Printf("newAtlasV2Client successfully created clients: %+v", clients)
	return clients, nil
}

func NewAtlasClientWithLambdaProxySupport(req *handler.Request, profileName, lambdaARN *string) (*MongoDBClient, *handler.ProgressEvent) {
	return newAtlasV2ClientWithLambdaProxySupport(req, profileName, true, lambdaARN)
}
