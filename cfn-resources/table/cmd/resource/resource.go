package resource

import (
	"context"
    "errors"
	"fmt"
	"log"
    "os"
    "strings"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
    "go.mongodb.org/atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/spf13/cast"
    "encoding/json"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
)

func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}
func cast64(i int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}

type TableCFNIdentifier struct {
    ProjectId           string
    ClusterName         string
    DatabaseName        string
    TableName           string
}
/* Note string version is "+" delimited string of the fields, in proper heirachry
*/
func (t TableCFNIdentifier) String() string {
    fields := []string{t.ProjectId,t.ClusterName,t.DatabaseName,t.TableName}
    return strings.Join(fields, "+")
}

func parseTableId(tableId string) (TableCFNIdentifier, error) {
    var t TableCFNIdentifier
    log.Printf("parseTableId tableId:%v",tableId)
    if ! strings.Contains(tableId,"+") {
        return t, errors.New("Invalid format TableCFNIdentifier")
    }
    parts := strings.Split(tableId,"+")
    log.Printf("parseTableId parts:%v",parts)
    //if len(parts)==3 {
    //    return t, errors.New("Invalid format TableCFNIdentifier")
    //}
    t = TableCFNIdentifier{
        ProjectId:          parts[0],
        ClusterName:        parts[1],
        DatabaseName:       parts[2],
        TableName:          parts[3],
    }
    return t, nil
}

type DeploymentSecret struct {
    PublicKey           string  `json:"PublicKey"`
    PrivateKey          string  `json:"PrivateKey"`
    ResourceID          string  `json:"ResourceID"`
}

func createDeploymentSecret(req *handler.Request, tid *TableCFNIdentifier, publicKey string, privateKey string) (*string, error) {
    deploySecret := &DeploymentSecret{
        PublicKey:      publicKey,
        PrivateKey:     privateKey,
        ResourceID:     fmt.Sprintf("%v",tid),
    }
    log.Printf("deploySecret: %v", deploySecret)
    deploySecretString, err := json.Marshal(deploySecret)
    log.Printf("deploySecretString: %s", deploySecretString)

    log.Println("===============================================")
    log.Printf("%+v",os.Environ())
    log.Println("===============================================")

    //sess := credentials.SessionFromCredentialsProvider(creds)
    // create a new secret from this struct with the json string


    // Create service client value configured for credentials
    // from assumed role.
    svc := secretsmanager.New(req.Session)

    //config := &aws.Config{
    //    Region: aws.String("us-east-1"),
    //}
    //svc := secretsmanager.New(session.New(), config)

    input := &secretsmanager.CreateSecretInput{
        Description:        aws.String("MongoDB Atlas Quickstart Deployment Secret"),
        Name:               aws.String(fmt.Sprintf("%v",tid)),
        SecretString:       aws.String(string(deploySecretString)),
    }

    result, err := svc.CreateSecret(input)
    if err != nil {
        // Print the error, cast err to awserr.Error to get the Code and
        // Message from an error.
        log.Printf("error create secret: %+v", err.Error())
		return nil, err
        //fmt.Println(err.Error())

    }
    log.Printf("Created secret result:%+v",result)
    return result.Name, nil

}


func getApiKeyFromDeploymentSecret(req *handler.Request, secretName string) (DeploymentSecret, error) {
   fmt.Printf("secretName=%s\n",secretName)
   sm := secretsmanager.New(req.Session)
   output, err := sm.GetSecretValue(&secretsmanager.GetSecretValueInput{SecretId: &secretName})
   if err != nil {
      panic(err.Error())
   }
   fmt.Println(*output.SecretString)
   var key DeploymentSecret
   err = json.Unmarshal( []byte(*output.SecretString), &key )
   if err != nil {
       log.Printf("Error --- %#+v", err.Error())
       return key, err
   }
   fmt.Println("%v",key)
   return key, nil
}


// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    log.Printf("Create - currentModel: %#+v, prevModel: %#+v", currentModel, prevModel)
	log.Printf("APIKEYS=======>%v,%v",*currentModel.PublicApiKey, *currentModel.PrivateApiKey)
	client, err := util.CreateMongoDBClient(*currentModel.PublicApiKey, *currentModel.PrivateApiKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}

    /* Create - this will look for a table and create it if it doesn't exist
    */

    thisCallbackContext := req.CallbackContext
    log.Printf("thisCallbackContext:%+v",thisCallbackContext)
    callbackCount, gotCount := thisCallbackContext["count"].(int)
    if ! gotCount {
        callbackCount = 1
    } else {
        callbackCount += 1
    }

    log.Printf("callbackCount: %i",callbackCount)

	projectID := *currentModel.ProjectId
	tableName := *currentModel.TableName
    log.Printf("projectID=%v, tableName=%v",projectID,tableName)

    databaseName := tableName
    if currentModel.DatabaseName != nil {
	    databaseName = *currentModel.DatabaseName
    }

    clusterName := tableName
    if currentModel.ClusterName != nil {
	    clusterName = *currentModel.ClusterName
    }


    // convert AWS- regions to MDB regions
    regionName := strings.ToUpper(strings.Replace(string(*currentModel.RegionName),"-","_",-1))
    log.Printf("regionName:%s",regionName)

    tid := &TableCFNIdentifier{
        ProjectId:          projectID,
        ClusterName:        clusterName,
        DatabaseName:       databaseName,
        TableName:          tableName,
    }
    tidString := fmt.Sprintf("%v",tid) 
    currentModel.TableCNFIdentifier = &tidString
    log.Printf("tid: %v, TableCFNIdentifier: %s",tid,currentModel.TableCNFIdentifier)

    secretName, err := createDeploymentSecret(&req, tid, *currentModel.PublicApiKey, *currentModel.PrivateApiKey)
	if err != nil {
		return handler.ProgressEvent{}, err
	}
    log.Printf("secretName: %s",secretName)

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // Cancel ctx as soon as this returns.
	cluster, _, err := client.Clusters.Get(ctx, projectID, clusterName)
	if err != nil {
        log.Printf("Cluster was not found, creating it now... clusterName:(%s) err:%s", clusterName, err)
        clusterRequest := &mongodbatlas.Cluster{
            Name:                     clusterName,
            ClusterType:              "REPLICASET",
            ProviderSettings:         &mongodbatlas.ProviderSettings{
                ProviderName:           "AWS",
                RegionName:           regionName,
                InstanceSizeName:     "M10",    //TODO: Wish could be M0!
            },
            ReplicationFactor:        cast64(3),
            NumShards:                cast64(1),
        }
        log.Printf("clusterRequest: %+v",clusterRequest)
        cluster, resp, err := client.Clusters.Create(ctx, projectID, clusterRequest)
        if err != nil {
            return handler.ProgressEvent{}, fmt.Errorf("error creating cluster: %w %v", err, &resp)
        }
        cc := map[string]interface{}{
            "status": "cluster-create",
            "cluster": cluster,
            "counter": callbackCount,
            "publicKey": *currentModel.PublicApiKey,
            "privateKey": *currentModel.PublicApiKey,
            "TableCFNIdentifier": fmt.Sprintf("%v",tid),
        }
        log.Printf("Created cluster- request callback in 30 seconds cluster:%+v",cluster)
        return handler.ProgressEvent{
            OperationStatus:      handler.InProgress,
            Message:              "In Progress, provisioning cluster",
            CallbackDelaySeconds: 30,
            CallbackContext:      cc,
            ResourceModel:        currentModel,
        }, nil
	}
    log.Printf("cluster:%+v",cluster)
    // Have a cluster, is it ready yet?
    if cluster.StateName != "IDLE" {
        log.Printf("Cluster not ready yet StateName:%s request callback in 30 seconds",cluster.StateName)
        cc := map[string]interface{}{
            "status": "cluster-create-wait",
            "cluster": cluster,
            "counter": callbackCount,
            "publicKey": *currentModel.PublicApiKey,
            "privateKey": *currentModel.PublicApiKey,
        }
        return handler.ProgressEvent{
            OperationStatus:      handler.InProgress,
            Message:              fmt.Sprintf("In Progress, cluster state %s",cluster.StateName),
            CallbackDelaySeconds: 30,
            CallbackContext:      cc,
            ResourceModel:        currentModel,
        }, nil

    }

    // cluster must be ready. 
    // Enable AWS cloud provider access and add a DB user for the new IAM Role

    // For example, "arn:aws:iam::466197078724:role/puffin-123-AtlasIAMRole-FO9UEDNJ9MZL"
    log.Println("++++++++++++ IAM setup now +++++++++++++++")
    username := *currentModel.Username
    if !strings.HasPrefix(username,"arn:aws:iam:") {
        return handler.ProgressEvent{}, fmt.Errorf("error CloudProviderAccess username must be AWS IAM Role or User: %s", username)
    }
    iamType := strings.Split(strings.Split(username,":")[5], "/")[0]
    awsIAMType := strings.ToUpper(iamType)
    dbUserDBName := "$external"
    log.Printf("username:%s",username)
    cpaReq := mongodbatlas.CloudProviderAccessRoleRequest{ProviderName: "AWS"}
    iamRole, _, err := client.CloudProviderAccess.CreateRole(ctx, projectID, &cpaReq)
    if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error CloudProviderAccess.CreateRole AWS: %s", err)
    }

    log.Printf("create - iam role db user iamRole:%+v",iamRole)

    /*
    authReq := mongodbatlas.CloudProviderAuthorizationRequest{ProviderName: "AWS",IAMAssumedRoleARN: username}
    iamRole2, _, err := client.CloudProviderAccess.AuthorizeRole(ctx, projectID, iamRole.RoleID, &authReq)
    if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error CloudProviderAccess.CreateRole AWS: %s", err)
    }

    log.Printf("authorize - iam role db user iamRole2:%+v",iamRole2)
    */

    var labels []mongodbatlas.Label
    labels = append(labels, mongodbatlas.Label{Key:"Comment",Value:"Created from AWS Quickstart"})
    log.Printf("labels: %#+v", labels)
    var scopes []mongodbatlas.Scope
    scopes = append(scopes, mongodbatlas.Scope{clusterName,"CLUSTER"})
    log.Printf("scopes: %#+v", scopes)

    var roles []mongodbatlas.Role
    roles = append(roles, mongodbatlas.Role{
                RoleName:       "readWrite",
                DatabaseName:   databaseName,
                CollectionName: tableName,
        })
    log.Printf("roles: %#+v", roles)
    user := &mongodbatlas.DatabaseUser{
        Roles:        roles,
        GroupID:      projectID,
        Username:     username,
        DatabaseName: dbUserDBName,
        Labels:       labels,
        Scopes:       scopes,
        AWSIAMType:   awsIAMType,
    }


    log.Printf("user: %#+v", user)

    log.Printf("Arguments: Project ID: %s, Request %#+v", projectID, user)

    newUser, _, err := client.DatabaseUsers.Create(ctx, projectID, user)
    if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error creating database user: %s", err)
    }
    log.Printf("newUser: %s", newUser)

    // Once here everything should be provisioned,
    
    // TODO// Issue ONE more callback to create the collection using driver connection.
    // This step can fail but the whole stack still be OK.
    // return properties
    // todo - util/mongodb

	currentModel.ConnectionStringsStandard = &cluster.ConnectionStrings.Standard
	currentModel.ConnectionStringsStandardSrv = &cluster.ConnectionStrings.StandardSrv
    log.Printf("read-------> cluster:%#+v",cluster)
    log.Printf("about to return currentModel: %#+v", currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil

}


// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    // * Check/set any callback context (
    callback := map[string]interface{}(req.CallbackContext)
    log.Printf("Read -  callback: %#+v",callback)

    if currentModel != nil {
        log.Printf("Read - currentModel: %#+v", currentModel)
    }
    if prevModel != nil {
        log.Printf("Read - prevModel: %#+v", prevModel)
    }

    key, err := getApiKeyFromDeploymentSecret(&req, *currentModel.TableCNFIdentifier)
    if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error lookupSecret: %w", err)
    }
    log.Printf("key:%+v",key)

    cfnid, err := parseTableId(*currentModel.TableCNFIdentifier)
	if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error parsing TableId: %w %v", err, *currentModel.TableCNFIdentifier)
    }
    log.Printf("cfnid: %v",cfnid)

    log.Printf("Read - Get clusterName:%s databaseName:%s",cfnid.ClusterName,cfnid.DatabaseName)

    client, err := util.CreateMongoDBClient(key.PublicKey, key.PrivateKey)
    if err != nil {
        return handler.ProgressEvent{}, err
    }
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // Cancel ctx as soon as this returns.
    cluster, resp, err := client.Clusters.Get(ctx, cfnid.ProjectId, cfnid.ClusterName)
    if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error reading cluster: %w %v", err, &resp)
    }
    currentModel.ConnectionStringsStandard = &cluster.ConnectionStrings.Standard
    currentModel.ConnectionStringsStandardSrv = &cluster.ConnectionStrings.StandardSrv

    log.Printf("Read - currentModel:+%v",currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    response := handler.ProgressEvent{
        OperationStatus: handler.Success,
        Message: "Update Complete - no changes applied, please use mongocli or the Atlas UI",
        ResourceModel: currentModel,
    }

    return response, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    response := handler.ProgressEvent{
        OperationStatus: handler.Success,
        Message: "Delete Complete - no changes applied, please use mongocli or the Atlas UI to delete clusters",
        ResourceModel: currentModel,
    }

    return response, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    /*
    response := handler.ProgressEvent{
        OperationStatus: handler.Success,
        Message: "List Complete",
        ResourceModel: currentModel,
    }
    */
    log.Printf("List called - alias for Read")
    return Read(req, prevModel, currentModel)
    //return response, nil
}
