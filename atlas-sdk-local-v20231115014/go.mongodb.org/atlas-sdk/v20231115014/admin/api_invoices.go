// Copyright 2023 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type InvoicesApi interface {

	/*
		CreateCostExplorerQueryProcess Create Cost Explorer query process

		[experimental] Creates a query process within the Cost Explorer for the given parameters. A token is returned that can be used to poll the status of the query and eventually retrieve the results.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return CreateCostExplorerQueryProcessApiRequest
	*/
	CreateCostExplorerQueryProcess(ctx context.Context, orgId string, costExplorerFilterRequestBody *CostExplorerFilterRequestBody) CreateCostExplorerQueryProcessApiRequest
	/*
		CreateCostExplorerQueryProcess Create Cost Explorer query process


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateCostExplorerQueryProcessApiParams - Parameters for the request
		@return CreateCostExplorerQueryProcessApiRequest
	*/
	CreateCostExplorerQueryProcessWithParams(ctx context.Context, args *CreateCostExplorerQueryProcessApiParams) CreateCostExplorerQueryProcessApiRequest

	// Method available only for mocking purposes
	CreateCostExplorerQueryProcessExecute(r CreateCostExplorerQueryProcessApiRequest) (*CostExplorerFilterResponse, *http.Response, error)

	/*
		CreateCostExplorerQueryProcess1 Return results from a given Cost Explorer query, or notify that the results are not ready yet.

		[experimental] Returns the usage details for a Cost Explorer query, if the query is finished and the data is ready to be viewed. If the data is not ready, a 'processing' response willindicate that another request should be sent later to view the data.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@param token Unique 64 digit string that identifies the Cost Explorer query.
		@return CreateCostExplorerQueryProcess1ApiRequest
	*/
	CreateCostExplorerQueryProcess1(ctx context.Context, orgId string, token string) CreateCostExplorerQueryProcess1ApiRequest
	/*
		CreateCostExplorerQueryProcess1 Return results from a given Cost Explorer query, or notify that the results are not ready yet.


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateCostExplorerQueryProcess1ApiParams - Parameters for the request
		@return CreateCostExplorerQueryProcess1ApiRequest
	*/
	CreateCostExplorerQueryProcess1WithParams(ctx context.Context, args *CreateCostExplorerQueryProcess1ApiParams) CreateCostExplorerQueryProcess1ApiRequest

	// Method available only for mocking purposes
	CreateCostExplorerQueryProcess1Execute(r CreateCostExplorerQueryProcess1ApiRequest) (string, *http.Response, error)

	/*
			DownloadInvoiceCSV Return One Organization Invoice as CSV

			[experimental] Returns one invoice that MongoDB issued to the specified organization in CSV format. A unique 24-hexadecimal digit string identifies the invoice. To use this resource, the requesting API Key have at least the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can query for a linked invoice if you have the Organization Billing Admin or Organization Owner Role.
		 To compute the total owed amount of the invoice - sum up total owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
			@return DownloadInvoiceCSVApiRequest
	*/
	DownloadInvoiceCSV(ctx context.Context, orgId string, invoiceId string) DownloadInvoiceCSVApiRequest
	/*
		DownloadInvoiceCSV Return One Organization Invoice as CSV


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DownloadInvoiceCSVApiParams - Parameters for the request
		@return DownloadInvoiceCSVApiRequest
	*/
	DownloadInvoiceCSVWithParams(ctx context.Context, args *DownloadInvoiceCSVApiParams) DownloadInvoiceCSVApiRequest

	// Method available only for mocking purposes
	DownloadInvoiceCSVExecute(r DownloadInvoiceCSVApiRequest) (string, *http.Response, error)

	/*
			GetInvoice Return One Organization Invoice

			[experimental] Returns one invoice that MongoDB issued to the specified organization. A unique 24-hexadecimal digit string identifies the invoice. You can choose to receive this invoice in JSON or CSV format. To use this resource, the requesting API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can query for a linked invoice if you have the Organization Billing Admin or Organization Owner role.
		To compute the total owed amount of the invoice - sum up total owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
			@return GetInvoiceApiRequest
	*/
	GetInvoice(ctx context.Context, orgId string, invoiceId string) GetInvoiceApiRequest
	/*
		GetInvoice Return One Organization Invoice


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetInvoiceApiParams - Parameters for the request
		@return GetInvoiceApiRequest
	*/
	GetInvoiceWithParams(ctx context.Context, args *GetInvoiceApiParams) GetInvoiceApiRequest

	// Method available only for mocking purposes
	GetInvoiceExecute(r GetInvoiceApiRequest) (string, *http.Response, error)

	/*
			ListInvoices Return All Invoices for One Organization

			[experimental] Returns all invoices that MongoDB issued to the specified organization. This list includes all invoices regardless of invoice status. To use this resource, the requesting API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can view linked invoices if you have the Organization Billing Admin or Organization Owner role.
		To compute the total owed amount of the invoices - sum up total owed of each invoice. It could be computed as a sum of owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
			@return ListInvoicesApiRequest
	*/
	ListInvoices(ctx context.Context, orgId string) ListInvoicesApiRequest
	/*
		ListInvoices Return All Invoices for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListInvoicesApiParams - Parameters for the request
		@return ListInvoicesApiRequest
	*/
	ListInvoicesWithParams(ctx context.Context, args *ListInvoicesApiParams) ListInvoicesApiRequest

	// Method available only for mocking purposes
	ListInvoicesExecute(r ListInvoicesApiRequest) (*PaginatedApiInvoice, *http.Response, error)

	/*
		ListPendingInvoices Return All Pending Invoices for One Organization

		[experimental] Returns all invoices accruing charges for the current billing cycle for the specified organization. To use this resource, the requesting API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can view linked invoices if you have the Organization Billing Admin or Organization Owner Role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
		@return ListPendingInvoicesApiRequest
	*/
	ListPendingInvoices(ctx context.Context, orgId string) ListPendingInvoicesApiRequest
	/*
		ListPendingInvoices Return All Pending Invoices for One Organization


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListPendingInvoicesApiParams - Parameters for the request
		@return ListPendingInvoicesApiRequest
	*/
	ListPendingInvoicesWithParams(ctx context.Context, args *ListPendingInvoicesApiParams) ListPendingInvoicesApiRequest

	// Method available only for mocking purposes
	ListPendingInvoicesExecute(r ListPendingInvoicesApiRequest) (*PaginatedApiInvoice, *http.Response, error)
}

// InvoicesApiService InvoicesApi service
type InvoicesApiService service

type CreateCostExplorerQueryProcessApiRequest struct {
	ctx                           context.Context
	ApiService                    InvoicesApi
	orgId                         string
	costExplorerFilterRequestBody *CostExplorerFilterRequestBody
}

type CreateCostExplorerQueryProcessApiParams struct {
	OrgId                         string
	CostExplorerFilterRequestBody *CostExplorerFilterRequestBody
}

func (a *InvoicesApiService) CreateCostExplorerQueryProcessWithParams(ctx context.Context, args *CreateCostExplorerQueryProcessApiParams) CreateCostExplorerQueryProcessApiRequest {
	return CreateCostExplorerQueryProcessApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         args.OrgId,
		costExplorerFilterRequestBody: args.CostExplorerFilterRequestBody,
	}
}

func (r CreateCostExplorerQueryProcessApiRequest) Execute() (*CostExplorerFilterResponse, *http.Response, error) {
	return r.ApiService.CreateCostExplorerQueryProcessExecute(r)
}

/*
CreateCostExplorerQueryProcess Create Cost Explorer query process

[experimental] Creates a query process within the Cost Explorer for the given parameters. A token is returned that can be used to poll the status of the query and eventually retrieve the results.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return CreateCostExplorerQueryProcessApiRequest
*/
func (a *InvoicesApiService) CreateCostExplorerQueryProcess(ctx context.Context, orgId string, costExplorerFilterRequestBody *CostExplorerFilterRequestBody) CreateCostExplorerQueryProcessApiRequest {
	return CreateCostExplorerQueryProcessApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		orgId:                         orgId,
		costExplorerFilterRequestBody: costExplorerFilterRequestBody,
	}
}

// Execute executes the request
//
//	@return CostExplorerFilterResponse
func (a *InvoicesApiService) CreateCostExplorerQueryProcessExecute(r CreateCostExplorerQueryProcessApiRequest) (*CostExplorerFilterResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CostExplorerFilterResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.CreateCostExplorerQueryProcess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.costExplorerFilterRequestBody == nil {
		return localVarReturnValue, nil, reportError("costExplorerFilterRequestBody is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.costExplorerFilterRequestBody
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateCostExplorerQueryProcess1ApiRequest struct {
	ctx        context.Context
	ApiService InvoicesApi
	orgId      string
	token      string
}

type CreateCostExplorerQueryProcess1ApiParams struct {
	OrgId string
	Token string
}

func (a *InvoicesApiService) CreateCostExplorerQueryProcess1WithParams(ctx context.Context, args *CreateCostExplorerQueryProcess1ApiParams) CreateCostExplorerQueryProcess1ApiRequest {
	return CreateCostExplorerQueryProcess1ApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		token:      args.Token,
	}
}

func (r CreateCostExplorerQueryProcess1ApiRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.CreateCostExplorerQueryProcess1Execute(r)
}

/*
CreateCostExplorerQueryProcess1 Return results from a given Cost Explorer query, or notify that the results are not ready yet.

[experimental] Returns the usage details for a Cost Explorer query, if the query is finished and the data is ready to be viewed. If the data is not ready, a 'processing' response willindicate that another request should be sent later to view the data.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param token Unique 64 digit string that identifies the Cost Explorer query.
	@return CreateCostExplorerQueryProcess1ApiRequest
*/
func (a *InvoicesApiService) CreateCostExplorerQueryProcess1(ctx context.Context, orgId string, token string) CreateCostExplorerQueryProcess1ApiRequest {
	return CreateCostExplorerQueryProcess1ApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		token:      token,
	}
}

// Execute executes the request
//
//	@return string
func (a *InvoicesApiService) CreateCostExplorerQueryProcess1Execute(r CreateCostExplorerQueryProcess1ApiRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.CreateCostExplorerQueryProcess1")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/billing/costExplorer/usage/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+csv", "application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DownloadInvoiceCSVApiRequest struct {
	ctx        context.Context
	ApiService InvoicesApi
	orgId      string
	invoiceId  string
}

type DownloadInvoiceCSVApiParams struct {
	OrgId     string
	InvoiceId string
}

func (a *InvoicesApiService) DownloadInvoiceCSVWithParams(ctx context.Context, args *DownloadInvoiceCSVApiParams) DownloadInvoiceCSVApiRequest {
	return DownloadInvoiceCSVApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		invoiceId:  args.InvoiceId,
	}
}

func (r DownloadInvoiceCSVApiRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.DownloadInvoiceCSVExecute(r)
}

/*
DownloadInvoiceCSV Return One Organization Invoice as CSV

[experimental] Returns one invoice that MongoDB issued to the specified organization in CSV format. A unique 24-hexadecimal digit string identifies the invoice. To use this resource, the requesting API Key have at least the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can query for a linked invoice if you have the Organization Billing Admin or Organization Owner Role.

	To compute the total owed amount of the invoice - sum up total owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
	@return DownloadInvoiceCSVApiRequest
*/
func (a *InvoicesApiService) DownloadInvoiceCSV(ctx context.Context, orgId string, invoiceId string) DownloadInvoiceCSVApiRequest {
	return DownloadInvoiceCSVApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		invoiceId:  invoiceId,
	}
}

// Execute executes the request
//
//	@return string
func (a *InvoicesApiService) DownloadInvoiceCSVExecute(r DownloadInvoiceCSVApiRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.DownloadInvoiceCSV")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}/csv"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+csv", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetInvoiceApiRequest struct {
	ctx        context.Context
	ApiService InvoicesApi
	orgId      string
	invoiceId  string
}

type GetInvoiceApiParams struct {
	OrgId     string
	InvoiceId string
}

func (a *InvoicesApiService) GetInvoiceWithParams(ctx context.Context, args *GetInvoiceApiParams) GetInvoiceApiRequest {
	return GetInvoiceApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
		invoiceId:  args.InvoiceId,
	}
}

func (r GetInvoiceApiRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.GetInvoiceExecute(r)
}

/*
GetInvoice Return One Organization Invoice

[experimental] Returns one invoice that MongoDB issued to the specified organization. A unique 24-hexadecimal digit string identifies the invoice. You can choose to receive this invoice in JSON or CSV format. To use this resource, the requesting API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can query for a linked invoice if you have the Organization Billing Admin or Organization Owner role.
To compute the total owed amount of the invoice - sum up total owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@param invoiceId Unique 24-hexadecimal digit string that identifies the invoice submitted to the specified organization. Charges typically post the next day.
	@return GetInvoiceApiRequest
*/
func (a *InvoicesApiService) GetInvoice(ctx context.Context, orgId string, invoiceId string) GetInvoiceApiRequest {
	return GetInvoiceApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
		invoiceId:  invoiceId,
	}
}

// Execute executes the request
//
//	@return string
func (a *InvoicesApiService) GetInvoiceExecute(r GetInvoiceApiRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.GetInvoice")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/{invoiceId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"invoiceId"+"}", url.PathEscape(parameterValueToString(r.invoiceId, "invoiceId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+csv", "application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListInvoicesApiRequest struct {
	ctx          context.Context
	ApiService   InvoicesApi
	orgId        string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListInvoicesApiParams struct {
	OrgId        string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *InvoicesApiService) ListInvoicesWithParams(ctx context.Context, args *ListInvoicesApiParams) ListInvoicesApiRequest {
	return ListInvoicesApiRequest{
		ApiService:   a,
		ctx:          ctx,
		orgId:        args.OrgId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListInvoicesApiRequest) IncludeCount(includeCount bool) ListInvoicesApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListInvoicesApiRequest) ItemsPerPage(itemsPerPage int) ListInvoicesApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListInvoicesApiRequest) PageNum(pageNum int) ListInvoicesApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListInvoicesApiRequest) Execute() (*PaginatedApiInvoice, *http.Response, error) {
	return r.ApiService.ListInvoicesExecute(r)
}

/*
ListInvoices Return All Invoices for One Organization

[experimental] Returns all invoices that MongoDB issued to the specified organization. This list includes all invoices regardless of invoice status. To use this resource, the requesting API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can view linked invoices if you have the Organization Billing Admin or Organization Owner role.
To compute the total owed amount of the invoices - sum up total owed of each invoice. It could be computed as a sum of owed amount of each payment included into the invoice. To compute payment's owed amount - use formula *totalBilledCents* * *unitPrice* + *salesTax* - *startingBalanceCents*.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListInvoicesApiRequest
*/
func (a *InvoicesApiService) ListInvoices(ctx context.Context, orgId string) ListInvoicesApiRequest {
	return ListInvoicesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return PaginatedApiInvoice
func (a *InvoicesApiService) ListInvoicesExecute(r ListInvoicesApiRequest) (*PaginatedApiInvoice, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiInvoice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.ListInvoices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListPendingInvoicesApiRequest struct {
	ctx        context.Context
	ApiService InvoicesApi
	orgId      string
}

type ListPendingInvoicesApiParams struct {
	OrgId string
}

func (a *InvoicesApiService) ListPendingInvoicesWithParams(ctx context.Context, args *ListPendingInvoicesApiParams) ListPendingInvoicesApiRequest {
	return ListPendingInvoicesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      args.OrgId,
	}
}

func (r ListPendingInvoicesApiRequest) Execute() (*PaginatedApiInvoice, *http.Response, error) {
	return r.ApiService.ListPendingInvoicesExecute(r)
}

/*
ListPendingInvoices Return All Pending Invoices for One Organization

[experimental] Returns all invoices accruing charges for the current billing cycle for the specified organization. To use this resource, the requesting API Key must have the Organization Billing Viewer, Organization Billing Admin, or Organization Owner role. If you have a cross-organization setup, you can view linked invoices if you have the Organization Billing Admin or Organization Owner Role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param orgId Unique 24-hexadecimal digit string that identifies the organization that contains your projects. Use the [/orgs](#tag/Organizations/operation/listOrganizations) endpoint to retrieve all organizations to which the authenticated user has access.
	@return ListPendingInvoicesApiRequest
*/
func (a *InvoicesApiService) ListPendingInvoices(ctx context.Context, orgId string) ListPendingInvoicesApiRequest {
	return ListPendingInvoicesApiRequest{
		ApiService: a,
		ctx:        ctx,
		orgId:      orgId,
	}
}

// Execute executes the request
//
//	@return PaginatedApiInvoice
func (a *InvoicesApiService) ListPendingInvoicesExecute(r ListPendingInvoicesApiRequest) (*PaginatedApiInvoice, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PaginatedApiInvoice
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "InvoicesApiService.ListPendingInvoices")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/orgs/{orgId}/invoices/pending"
	localVarPath = strings.Replace(localVarPath, "{"+"orgId"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v ApiError
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
