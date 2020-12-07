package resource

import (
    "bytes"
	"context"
    "errors"
	"fmt"
	"log"
    "os"
    "strings"
    "text/template"
    "time"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
    "go.mongodb.org/atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/spf13/cast"
    "encoding/json"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/iam"
    "github.com/aws/aws-sdk-go/service/secretsmanager"
)

const DefaultMongoDBVersion = "4.4"

func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}
func cast64(i int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}

type PolicyDocument struct {
    Version   string
    Statement []StatementEntry
}

type StatementEntry struct {
    Effect          string
    Action          []string
    Resource        string
    Principal       map[string]interface{}
    Condition       map[string]interface{}
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
    DBUsername          string  `json:"DBUsername"` 
}


//(&req, iamRole.AtlasAWSAccountARN, iamRole.AtlasAssumedRoleExternalID, username
// TODO - this should be much more graceful, we should be able to just
// inject the trust policy we need not overwrite it all.
const roleTrustPolicyTemplate = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "{{ .AtlasAWSAccountARN }}"
      },
      "Action": "sts:AssumeRole",
      "Condition": {
        "StringEquals": {
          "sts:ExternalId": "{{ .AtlasAssumedRoleExternalID }}"
        }
      }
    },
    {
      "Effect": "Allow",
      "Principal": {
        "AWS": "{{ .AWSUserPrincipalARN }}"
      },
      "Action": "sts:AssumeRole"
    },
    {
      "Effect": "Allow",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

func addRoleWithTrustPolicy(req *handler.Request, atlasAWSAccountARN string, atlasAssumedRoleExternalID string, targetAWSIAMRoleARN string) error {
    log.Printf("addRoleWithTrustPolicy -- atlasAWSAccountARN:%s, atlasAssumedRoleExternalID:%s, targetAWSIAMRoleARN:%s",atlasAWSAccountARN, atlasAssumedRoleExternalID, targetAWSIAMRoleARN)
    iamService := iam.New(req.Session)


    t := strings.Split(targetAWSIAMRoleARN,"/")
    roleName := t[1]
    ap := strings.Split(t[0],":")
    awsUserPrincipalARN := fmt.Sprintf("%s:root",strings.Join(ap[0:len(ap)-1],":"))
    log.Printf("addROleWithTrustPolicy roleName:%s awsUserPrincipalARN:%s",roleName, awsUserPrincipalARN)

    trustTemplate, err := template.New("todos").Parse( roleTrustPolicyTemplate )
    if err != nil {
        log.Printf("Error parsing role trust policy %v",err)
        return err
    }
    trustPolicyContext := struct {
        AtlasAWSAccountARN,
        AtlasAssumedRoleExternalID,
        AWSUserPrincipalARN string
    }{
        atlasAWSAccountARN,
        atlasAssumedRoleExternalID,
        awsUserPrincipalARN,
    }

    log.Printf("addRoleWithTrustPolicy trustPolicyContext:%v",trustPolicyContext)

    var rawTrustPolicy bytes.Buffer
    err = trustTemplate.Execute(&rawTrustPolicy, trustPolicyContext)
    if err != nil {
        log.Printf("Error executing role trust policy template %v",err)
        return err
    }


    log.Printf("trustPolicyDoc:   string(rawTrustPolicy):%s",rawTrustPolicy.String())
    input := &iam.UpdateAssumeRolePolicyInput{
        PolicyDocument: aws.String(rawTrustPolicy.String()),
        RoleName:       aws.String(roleName),
    }
    log.Printf("UpdateAssumeRolePolicyInput input:%v",input)
    result, err := iamService.UpdateAssumeRolePolicy(input)


    if err != nil {
      log.Printf("createRole error:%v",err)
      //fmt.Println(err.Error())
      return err
    }
    log.Printf("UpdateAssumeRolePolicy result:%v",result)
    return nil
}


func createDeploymentSecret(req *handler.Request, tid *TableCFNIdentifier, publicKey string, privateKey string, dbUsername string) (*string, error) {
    deploySecret := &DeploymentSecret{
        PublicKey:      publicKey,
        PrivateKey:     privateKey,
        ResourceID:     fmt.Sprintf("%v",tid),
        DBUsername:     dbUsername,
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
       log.Printf("Error --- %v", err.Error())
       return key, err
   }
   fmt.Println("%v",key)
   return key, nil
}


// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    log.Printf("Create - currentModel: %v, prevModel: %v", currentModel, prevModel)
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

    // If this is first call, create a deployment secret to cache the api keys
    if _, ok := thisCallbackContext["TableCFNIdentifier"]; !ok {
        secretName, err := createDeploymentSecret(&req, tid, *currentModel.PublicApiKey, *currentModel.PrivateApiKey,*currentModel.Username)
        if err != nil {
            return handler.ProgressEvent{}, err
        }
        log.Printf("secretName: %s",secretName)
    }

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // Cancel ctx as soon as this returns.

    // Check if this is the FINAL callback after Cluster, IAM, DBUser is all set.
    // If so - issue the driver interaction to setup the actual Database/Collection - 
    // Write some simple document (TODO/ inject sample data from our sets here)
    if status, ok := thisCallbackContext["status"]; ok {
        log.Printf("Had a status for callback:%s",status)
        if status == "table-setup" {
            log.Printf("TABLE_SETUP!!!!!!!! os.Environ():%v",os.Environ())
            log.Printf("HERE -----> call util/mongodb with ......")
            cluster, _, err := client.Clusters.Get(ctx, projectID, clusterName)
            if err != nil {
                return handler.ProgressEvent{}, err
            }
            roleArn := thisCallbackContext["table-setup-role-arn"].(string)

            currentModel.ConnectionStringsStandard = &cluster.ConnectionStrings.Standard
            currentModel.ConnectionStringsStandardSrv = &cluster.ConnectionStrings.StandardSrv
            log.Printf("try list db names roleArn: %s, srv:%v", roleArn, *currentModel.ConnectionStringsStandardSrv)

            dbs, err := util.ListDatabaseNames(&req, *currentModel.ConnectionStringsStandardSrv,roleArn)
            if err != nil {
                log.Printf("ListDatabaseNames err:%v",err)
                return handler.ProgressEvent{}, err
            }
            log.Printf("ListDatabaseNames - did it work????? dbs:%v",dbs) 
            return handler.ProgressEvent{
                OperationStatus: handler.Success,
                Message:         "Create Complete, Go Data.",
                ResourceModel:   currentModel,
            }, nil
        }
    }

	cluster, _, err := client.Clusters.Get(ctx, projectID, clusterName)
	if err != nil {
        log.Printf("Cluster was not found, creating it now... clusterName:(%s) err:%s", clusterName, err)
        clusterRequest := &mongodbatlas.Cluster{
            Name:                     clusterName,
            ClusterType:              "REPLICASET",
            MongoDBMajorVersion:      DefaultMongoDBVersion,
            ProviderSettings:         &mongodbatlas.ProviderSettings{
                ProviderName:         "AWS",
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
            "TableCFNIdentifier": fmt.Sprintf("%v",tid),
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
    username := *currentModel.Username
    log.Printf("++++++++++++ IAM setup now +++++++++++++++ username:%s",username)
    if !strings.HasPrefix(username,"arn:aws:iam:") {
        return handler.ProgressEvent{}, fmt.Errorf("error CloudProviderAccess username must be AWS IAM Role or User: %s", username)
    }
    up := strings.Split(username,":")
    log.Printf(" ~~~~~~~~~~~~~ up:%v",up)
    iamType := strings.Split(up[len(up)-1], "/")[0]
    awsIAMType := strings.ToUpper(iamType)
    dbUserDBName := "$external"
    log.Printf("username:%s",username)
    cpaReq := mongodbatlas.CloudProviderAccessRoleRequest{ProviderName: "AWS"}
    iamRole, _, err := client.CloudProviderAccess.CreateRole(ctx, projectID, &cpaReq)
    if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error CloudProviderAccess.CreateRole AWS: %s", err)
    }

    log.Printf("create - CloudProviderAccess role db user iamRole:%+v",iamRole)

    /* In order to authorize this AWS intergration we need to add a trust policy so
    that the new IAM Role Atlas just generated is allowed to connect. this is linked through
    the projectID setting, so we really only need one per project (think)

    */
    err = addRoleWithTrustPolicy(&req, iamRole.AtlasAWSAccountARN, iamRole.AtlasAssumedRoleExternalID, username)
    if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error addRoleWithTrustPolicy: %s", err)
    }

    // pause a moment to allow the new trust policy to "settle in" to AWS
    time.Sleep(10 * time.Second)
    /**/
    authReq := mongodbatlas.CloudProviderAuthorizationRequest{ProviderName: "AWS",IAMAssumedRoleARN: username}
    log.Printf("create - CloudProviderAccess - attempt authorize role RoleID:%s, authReq:+%V",iamRole.RoleID,authReq)
    iamRole2, res, err := client.CloudProviderAccess.AuthorizeRole(ctx, projectID, iamRole.RoleID, &authReq)
    log.Printf("AuthorizeRole res.StatusCode: %v", res.StatusCode)

    if err != nil {
        if res.StatusCode == 409 {
            log.Printf("IAM Role already authorized, provisioning continuing. err: %v", err)
            iamRoles, _, err2 := client.CloudProviderAccess.ListRoles(ctx, projectID)
            if err2 != nil {
                log.Printf("Error ListRoles err:%v",err)
                return handler.ProgressEvent{}, fmt.Errorf("error CloudProviderAccess.ListRoles err:%v", err)
            }
            log.Printf("iamRoles:%v",iamRoles)
            for _,role := range iamRoles.AWSIAMRoles {
                log.Printf("role:%v",role)
                if role.IAMAssumedRoleARN == username {
                    iamRole2 = &role
                    log.Printf("Found existing role!  iamRole2:%v", iamRole2)
                }
            }

        } else {
            log.Printf("was NOT 409 Conflict, err: %v", err)
            return handler.ProgressEvent{}, fmt.Errorf("error CloudProviderAccess.AuthorizeRole AWS: %s", err)
        }
    }

    log.Printf("authorize - iam role db user iamRole2:%+v",iamRole2)
    /**/

    // By default, open up access for a temporary duration.
    // TODO - Review this with the team!
    // ? Should we just generate an AWS Security Group on the fly (or allow passed in),
    // then users could just add that group to their app's VPCs?

    tomorrow := time.Now().Add(time.Hour*24)
	accessListRequest := []*mongodbatlas.ProjectIPAccessList{
		{
			GroupID:   projectID,
		    IPAddress: "0.0.0.0/1",
            DeleteAfterDate: tomorrow.Format(time.RFC3339),
		},
	}

	projectIPAccessList, _, err := client.ProjectIPAccessList.Create(ctx, projectID, accessListRequest)
	if err != nil {
		log.Printf("ProjectIPAccessList.Create returned error: %v", err)
        return handler.ProgressEvent{}, fmt.Errorf("ProjectIPAccessList.Create returned error: %v", err)
	}

    log.Printf("projectIDAccessList: %v",projectIPAccessList)

    var labels []mongodbatlas.Label
    labels = append(labels, mongodbatlas.Label{Key:"Comment",Value:"Created from AWS Quickstart"})
    log.Printf("labels: %v", labels)
    var scopes []mongodbatlas.Scope
    scopes = append(scopes, mongodbatlas.Scope{clusterName,"CLUSTER"})
    log.Printf("scopes: %v", scopes)

    var roles []mongodbatlas.Role
    newDBRole := mongodbatlas.Role{
        RoleName:       "readWrite",
        DatabaseName:   databaseName,
        CollectionName: tableName,
    }
    roles = append(roles, newDBRole)
    log.Printf("roles: %v", roles)
    user := &mongodbatlas.DatabaseUser{
        Roles:        roles,
        GroupID:      projectID,
        Username:     username,
        DatabaseName: dbUserDBName,
        Labels:       labels,
        Scopes:       scopes,
        AWSIAMType:   awsIAMType,
    }


    log.Printf("user: %v", user)

    log.Printf("Arguments: Project ID: %s, Request %v", projectID, user)

    newUser, res2, err := client.DatabaseUsers.Create(ctx, projectID, user)
    if err != nil {
        if res2.StatusCode == 409 {
            log.Printf("DatabaseUser already exists, provisioning continuing. err: %v", err)
            // Now we need check we have a role to access this new "Table"
            user, _, err2 := client.DatabaseUsers.Get(ctx,dbUserDBName,projectID,user.Username)
            if err2 != nil {
                log.Printf("DatabaseUser.Get returned error: %v", err)
                return handler.ProgressEvent{}, fmt.Errorf("DatabaseUser.Get returned error: %v", err)
            }
            log.Printf("Found existing dbuser user:%v, attempt add new role",user)

            isSameRole := func(r1 mongodbatlas.Role, r2 mongodbatlas.Role) bool {
                log.Printf("isSameRole r1:%v r2:%v",r1,r2)
                if r1.RoleName != r2.RoleName { return false }
                if r1.DatabaseName != r2.DatabaseName { return false }
                if r1.CollectionName != r2.CollectionName { return false }
                log.Println("isSameRole - passed all checks")
                return true

            }
            foundRole := false
            for _, role := range user.Roles {
                foundRole = isSameRole(role, newDBRole)
            }
            if !foundRole {   // didn't find it, so add it

                user.Roles = append(user.Roles, newDBRole)
                fixedUser, _, err3 := client.DatabaseUsers.Update(ctx,dbUserDBName,projectID, user)
                if err2 != nil {
                    log.Printf("DatabaseUser.Update returned error: %v", err3)
                    return handler.ProgressEvent{}, fmt.Errorf("DatabaseUser.Update returned error: %v", err3)
                }
                log.Printf("Found existing dbuser user:%v, attempt add new role",fixedUser)

            }


        } else {
            return handler.ProgressEvent{}, fmt.Errorf("error creating database user: %s", err)
        }
    }
    log.Printf("newUser: %s", newUser)

    // Once here everything should be provisioned,
    
    // TODO// Issue ONE more callback to create the collection using driver connection.
    // This step can fail but the whole stack still be OK.
    // return properties
    // todo - util/mongodb
    // this means the "user" running this function (the lambda iam execution role, needs to also be a dbuser)
    // or we need something else here.


    /*
    log.Printf("AtlasTable provisioning penultimate step complete. Issuing final callback 'status':'table-setup' ")
    cc := map[string]interface{}{
        "status": "table-setup",
        "table-setup-role-arn": newUser.Username,
        "cluster": cluster,
        "counter": callbackCount,
        "publicKey": *currentModel.PublicApiKey,
        "privateKey": *currentModel.PublicApiKey,
        "TableCFNIdentifier": fmt.Sprintf("%v",tid),
    }
    return handler.ProgressEvent{
        OperationStatus:      handler.InProgress,
        Message:              fmt.Sprintf("Final step, connecting to cluster and creating Database.Collection %s.%s",databaseName, tableName),
        CallbackDelaySeconds: 30,
        CallbackContext:      cc,
        ResourceModel:        currentModel,
    }, nil
    */
    
    /**/
	currentModel.ConnectionStringsStandard = &cluster.ConnectionStrings.Standard
	currentModel.ConnectionStringsStandardSrv = &cluster.ConnectionStrings.StandardSrv
    currentModel.Username = &newUser.Username
    log.Printf("read-------> cluster:%v",cluster)
    log.Printf("about to return currentModel: %v", currentModel)
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
    /**/
}


// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    // * Check/set any callback context (
    callback := map[string]interface{}(req.CallbackContext)
    log.Printf("Read -  callback: %v",callback)

    if currentModel != nil {
        log.Printf("Read - currentModel: %v", currentModel)
    }
    if prevModel != nil {
        log.Printf("Read - prevModel: %v", prevModel)
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
    currentModel.Username = &key.DBUsername
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


/*




    policy := PolicyDocument{
        Version: "2012-10-17",
        Statement: []StatementEntry{
            StatementEntry{
                Effect: "Allow",
                Action: []string{
                    "sts:AssumeRole", // Allow for creating log groups
                },
            },
        },
    }
    log.Printf("addRoleTrustPolicy policy:%+v",policy)
    b, err := json.Marshal(&policy)
    if err != nil {
        log.Println(fmt.Println("Error marshaling policy", err))
        return err
    }

    createPolicyResult, err := iamService.CreatePolicy(&iam.CreatePolicyInput{
        PolicyDocument: aws.String(string(b)),
        PolicyName:     aws.String(newPolicyName),
    })

    if err != nil {
        fmt.Println("Error", err)
        return err
    }

    log.Println("createPolicyResult: %+v", createPolicyResult)
 
    params := &iam.CreateRoleInput{
        AssumeRolePolicyDocument: aws.String(fmt.Sprintf("{\"Version\": \"2012-10-17\",\"Statement\": [{\"Effect\": \"Allow\",\"Principal\": {\"AWS\": \"%s\"},\"Action\": \"sts:AssumeRole\",\"Condition\": {\"StringEquals\" : { \"sts:ExternalId\" : \"%s\" } }}]}",atlasIAMARN,projectID)),
      Description:              aws.String("Role description"),
      PermissionsBoundary:      aws.String(*createPolicyResult.Policy.Arn),
      RoleName:                 aws.String("rolename"),
    }

    resp, err := iamService.CreateRole(params)

    /*
    attachResult, err := iamService.AttachRolePolicy(&iam.AttachRolePolicyInput{
        PolicyArn: aws.String(*createPolicyResult.Policy.Arn),
        RoleName: aws.String(*targetAWSIAMRoleARN),
    })
    if err != nil {
        log.Printf("attachRolePolicy error:%v",err)
        return err
    }
    log.Printf("attachRolePolicyResult: %+v", attachResult)
*/
