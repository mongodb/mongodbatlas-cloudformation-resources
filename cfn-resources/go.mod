module github.com/mongodb/mongodbatlas-cloudformation-resources

go 1.26.2

// Replacing with local copy of Atlas SDK v20231115014 to support new AdvancedConfiguration in *admin.AdvancedClusterDescription
replace go.mongodb.org/atlas-sdk/v20231115014 => ../vendor/go.mongodb.org/atlas-sdk/v20231115014

require (
	github.com/aws-cloudformation/cloudformation-cli-go-plugin v1.2.0
	github.com/aws/aws-sdk-go-v2 v1.41.7
	github.com/aws/aws-sdk-go-v2/config v1.32.17
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.71.11
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.299.1
	github.com/aws/aws-sdk-go-v2/service/secretsmanager v1.41.7
	github.com/aws/smithy-go v1.25.1
	github.com/dave/jennifer v1.7.1
	github.com/getkin/kin-openapi v0.137.0
	github.com/ghodss/yaml v1.0.0
	github.com/mongodb-forks/digest v1.1.0
	github.com/mongodb-labs/go-client-mongodb-atlas-app-services v1.0.0
	github.com/rs/xid v1.6.0
	github.com/spf13/cast v1.10.0
	github.com/stretchr/testify v1.11.1
	github.com/tidwall/pretty v1.2.1
	go.mongodb.org/atlas-sdk/v20231115002 v20231115002.1.0
	go.mongodb.org/atlas-sdk/v20231115014 v20231115014.0.0
	go.mongodb.org/atlas-sdk/v20250312013 v20250312013.2.0
)

require (
	github.com/aws/aws-lambda-go v1.54.0 // indirect
	github.com/aws/aws-sdk-go v1.55.8 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.19.16 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.23 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.24 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.9 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.0.11 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.30.17 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.35.21 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.42.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-openapi/jsonpointer v0.23.1 // indirect
	github.com/go-openapi/swag/jsonname v0.26.0 // indirect
	github.com/google/go-querystring v1.2.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.9.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/oasdiff/yaml v0.0.9 // indirect
	github.com/oasdiff/yaml3 v0.0.12 // indirect
	github.com/perimeterx/marshmallow v1.1.5 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/santhosh-tekuri/jsonschema/v6 v6.0.2 // indirect
	github.com/segmentio/ksuid v1.0.4 // indirect
	github.com/stretchr/objx v0.5.3 // indirect
	github.com/woodsbury/decimal128 v1.4.0 // indirect
	go.mongodb.org/atlas v0.38.0 // indirect
	golang.org/x/oauth2 v0.36.0 // indirect
	golang.org/x/text v0.36.0 // indirect
	gopkg.in/validator.v2 v2.0.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
