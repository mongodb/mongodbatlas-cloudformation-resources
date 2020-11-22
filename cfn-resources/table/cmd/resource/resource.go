package resource

import (
	"context"
	"fmt"
	"log"
    "strings"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
    "go.mongodb.org/atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/spf13/cast"
)

func castNO64(i *int64) *int {
	x := cast.ToInt(&i)
	return &x
}
func cast64(i int) *int64 {
	x := cast.ToInt64(&i)
	return &x
}

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    log.Printf("Create - currentModel: %#+v, prevModel: %#+v", currentModel, prevModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
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

    username := tableName
    if currentModel.Username != nil {
	    username = cast.ToString(currentModel.Username)
    }
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

    cfnid := fmt.Sprintf("%s-%s-%s",cast.ToString(currentModel.ProjectId),tableName,username)
    currentModel.TableCNFIdentifier = &cfnid
    log.Printf("TableCFNIdentifier: %s",cfnid)

	cluster, _, err := client.Clusters.Get(context.Background(), projectID, clusterName)
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
        cluster, resp, err := client.Clusters.Create(context.Background(), projectID, clusterRequest)
        if err != nil {
            return handler.ProgressEvent{}, fmt.Errorf("error creating cluster: %w %v", err, &resp)
        }
        cc := map[string]interface{}{
            "status": "cluster-create",
            "cluster": cluster,
            "counter": callbackCount,
        }
        log.Printf("Created cluster- request callback in 30 seconds cluster:%+v",cluster)
        return handler.ProgressEvent{
            OperationStatus:      handler.InProgress,
            Message:              "In Progress, provisioning cluster",
            CallbackDelaySeconds: 30,
            CallbackContext:      cc,
            ResourceModel:        &currentModel,
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
        }
        return handler.ProgressEvent{
            OperationStatus:      handler.InProgress,
            Message:              fmt.Sprintf("In Progress, cluster state %s",cluster.StateName),
            CallbackDelaySeconds: 30,
            CallbackContext:      cc,
            ResourceModel:        &currentModel,
        }, nil

    }

    // cluster must be ready. see if we have the db-user
    dbUserDBName := "admin" // Or will allow "$external" TODO
    // For example, "arn:aws:iam::466197078724:role/puffin-123-AtlasIAMRole-FO9UEDNJ9MZL"
    awsIAMType := "NONE"
    password := username
    if strings.HasPrefix(username,"arn:aws:iam:") {
        dbUserDBName = "$external" 
        iamType := strings.Split(strings.Split(username,":")[5], "/")[0]
        awsIAMType = strings.ToUpper(iamType)
        password = ""

        // TODO - need add call to CloudProviderAccess to enable IAM
    }
	databaseUser, _, err := client.DatabaseUsers.Get(context.Background(), dbUserDBName, projectID, username)
	if err != nil {
		//return handler.ProgressEvent{}, 
        log.Printf("error fetching database user (%s): %s", username, err)
        log.Printf("create - got error, user does not exist, try create them, %s", err)

        var labels []mongodbatlas.Label
        for _, l := range currentModel.Labels {

            label := mongodbatlas.Label{
                Key: *l.Key,
                Value: *l.Value,
            }
            labels = append(labels, label)
        }
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
            Password:     password,
            DatabaseName: dbUserDBName,
            Labels:       labels,
            Scopes:       scopes,
            AWSIAMType:   awsIAMType,
        }
        log.Printf("user: %#+v", user)

        log.Printf("Arguments: Project ID: %s, Request %#+v", projectID, user)

        newUser, _, err := client.DatabaseUsers.Create(context.Background(), projectID, user)
        if err != nil {
            return handler.ProgressEvent{}, fmt.Errorf("error creating database user: %s", err)
        }
        log.Printf("newUser: %s", newUser)
	} else {
        log.Printf("Found existing user: %+v",databaseUser)
    }

    // Once here everything should be provisioned, setup the
    // return properties

	currentModel.ConnectionStrings = &ConnectionStringsDefinition{
        Standard:               &cluster.ConnectionStrings.Standard,
        StandardSrv:            &cluster.ConnectionStrings.StandardSrv,
	    Private:                &cluster.ConnectionStrings.Private,
        PrivateSrv:             &cluster.ConnectionStrings.PrivateSrv,
	    //AwsPrivateLink:         &cluster.ConnectionStrings.AwsPrivateLink,
	    //AwsPrivateLinkSrv:      &cluster.ConnectionStrings.AwsPrivateLinkSrv,
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil

}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
    // Add your code here:
    // * Make API calls (use req.Session)
    // * Mutate the model
    // * Check/set any callback context (req.CallbackContext / response.CallbackContext)

    /*
        // Construct a new handler.ProgressEvent and return it
        response := handler.ProgressEvent{
            OperationStatus: handler.Success,
            Message: "Read complete",
            ResourceModel: currentModel,
        }

        return response, nil
    */

    // Not implemented, return an empty handler.ProgressEvent
    // and an error
    log.Printf("Read - currentModel: %#+v, prevModel: %#+v", currentModel, prevModel)
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey, *currentModel.ApiKeys.PrivateKey)
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

    username := tableName
    if currentModel.Username != nil {
	    username = cast.ToString(currentModel.Username)
    }
    databaseName := tableName
    if currentModel.DatabaseName != nil {
	    databaseName = *currentModel.DatabaseName
    }

    clusterName := tableName
    if currentModel.ClusterName != nil {
	    clusterName = *currentModel.ClusterName
    }
    cfnid := fmt.Sprintf("%s-%s-%s",cast.ToString(currentModel.ProjectId),tableName,username)
    currentModel.TableCNFIdentifier = &cfnid
    log.Printf("TableCFNIdentifier: %s",cfnid)
    log.Printf("Read - Get clusterName:%s databaseName:%s",clusterName,databaseName)
	cluster, resp, err := client.Clusters.Get(context.Background(), projectID, clusterName)
	if err != nil {
        return handler.ProgressEvent{}, fmt.Errorf("error reading cluster: %w %v", err, &resp)
    }
	currentModel.ConnectionStrings = &ConnectionStringsDefinition{
        Standard:               &cluster.ConnectionStrings.Standard,
        StandardSrv:            &cluster.ConnectionStrings.StandardSrv,
	    Private:                &cluster.ConnectionStrings.Private,
        PrivateSrv:             &cluster.ConnectionStrings.PrivateSrv,
	    //AwsPrivateLink:         &cluster.ConnectionStrings.AwsPrivateLink,
	    //AwsPrivateLinkSrv:      &cluster.ConnectionStrings.AwsPrivateLinkSrv,
	}
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
    response := handler.ProgressEvent{
        OperationStatus: handler.Success,
        Message: "List Complete",
        ResourceModel: currentModel,
    }

    return response, nil
}
