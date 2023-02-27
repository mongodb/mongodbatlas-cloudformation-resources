package profile

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

const (
	defaultProfile    = "default"
	profileNamePrefix = "cfn/atlas/profile"
)

type Profile struct {
	PublicKey  string `json:"PublicKey"`
	PrivateKey string `json:"PrivateKey"`
	BaseURL    string `json:"BaseUrl"`
}

func NewProfile(req *handler.Request, profileName *string) (*Profile, error) {
	if profileName == nil || *profileName == "" {
		profileName = aws.String(defaultProfile)
	}

	secretsManagerClient := secretsmanager.New(req.Session)
	resp, err := secretsManagerClient.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: aws.String(fmt.Sprintf("%s/%s", profileNamePrefix, *profileName))})
	if err != nil {
		return nil, err
	}

	profile := new(Profile)
	err = json.Unmarshal([]byte(*resp.SecretString), &profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

func (p *Profile) NewBaseURL() string {
	if baseURL := os.Getenv("MONGODB_ATLAS_BASE_URL"); baseURL != "" {
		return baseURL
	}

	return p.BaseURL
}

func (p *Profile) NewPublicKey() string {
	if k := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY"); k != "" {
		return k
	}

	return p.PublicKey
}

func (p *Profile) NewPrivateKey() string {
	if k := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY"); k != "" {
		return k
	}

	return p.PrivateKey
}

func (p *Profile) AreKeysAvailable() bool {
	return p.NewPublicKey() == "" || p.PrivateKey == ""
}
