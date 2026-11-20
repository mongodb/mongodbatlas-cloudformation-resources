module github.com/mongodb/mongodbatlas-cloudformation-resources

go 1.26.2

// Replacing with local copy of Atlas SDK v20231115014 to support new AdvancedConfiguration in *admin.AdvancedClusterDescription
replace go.mongodb.org/atlas-sdk/v20231115014 => ../vendor/go.mongodb.org/atlas-sdk/v20231115014

require (
	github.com/aws-cloudformation/cloudformation-cli-go-plugin v1.2.0
	github.com/aws/aws-sdk-go v1.55.8
	github.com/aws/aws-sdk-go-v2 v1.48.0
	github.com/aws/aws-sdk-go-v2/config v1.33.18
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.83.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.345.0
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.50.12
	github.com/aws/smithy-go v1.28.9
	github.com/dave/jennifer v1.7.1
	github.com/getkin/kin-openapi v0.153.0
	github.com/ghodss/yaml v1.0.0
	github.com/mongodb-forks/digest v1.2.0
	github.com/mongodb-labs/go-client-mongodb-atlas-app-services v1.0.0
	github.com/rs/xid v1.6.0
	github.com/spf13/cast v1.10.0
	github.com/stretchr/testify v1.12.1
	github.com/tidwall/pretty v1.3.0
	go.mongodb.org/atlas-sdk/v20231115002 v20231115002.1.0
	go.mongodb.org/atlas-sdk/v20231115014 v20231115014.0.0
	go.mongodb.org/atlas-sdk/v20250312028 v20250312028.1.0
)

require (
	github.com/aws/aws-lambda-go v1.55.1 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.21.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.20.9 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.5.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.8.12 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.5.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.14.12 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.13.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.39.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.44.5 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.52.0 // indirect
	github.com/go-openapi/jsonpointer v1.0.2 // indirect
	github.com/google/go-querystring v1.2.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.3 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	github.com/stretchr/objx v0.8.0 // indirect
	go.mongodb.org/atlas v0.38.0 // indirect
	go.yaml.in/yaml/v3 v3.0.6 // indirect
	go.yaml.in/yaml/v4 v4.0.0 // indirect
	golang.org/x/oauth2 v0.38.0 // indirect
	golang.org/x/text v0.44.0 // indirect
	gopkg.in/validator.v2 v2.0.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
