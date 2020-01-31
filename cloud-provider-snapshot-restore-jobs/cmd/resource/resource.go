package resource

import (
	"context"
	"fmt"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/encoding"
	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	matlasClient "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
)

// Create handles the Create event from the Cloudformation service.
func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	deliveryType := currentModel.DeliveryType.Value()
	if deliveryType == nil || (*deliveryType != "automated" && *deliveryType != "download") {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "error creating cloud provider snapshot restore job: you need to implement only one, `automated` and `download` delivery types",
			ResourceModel:   currentModel,
		}, nil
	}
	targetClusterName := currentModel.TargetClusterName.Value()
	targetProjectId := currentModel.TargetProjectId.Value()
	if *deliveryType == "automated" {
		if targetClusterName == nil || *targetClusterName == "" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "error creating cloud provider snapshot restore job: `target_cluster_name` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
		if targetProjectId == nil || *targetProjectId == "" {
			return handler.ProgressEvent{
				OperationStatus: handler.Failed,
				Message:         "error creating cloud provider snapshot restore job: `target_project_id` must be set if delivery type is `automated`",
				ResourceModel:   currentModel,
			}, nil
		}
	}
	requestParameters := &matlasClient.SnapshotReqPathParameters{
		GroupID:     *currentModel.ProjectId.Value(),
		ClusterName: *currentModel.ClusterName.Value(),
	}
	snapshotRequest := &matlasClient.CloudProviderSnapshotRestoreJob{
		SnapshotID:   *currentModel.SnapshotId.Value(),
		DeliveryType: *deliveryType,
	}
	if *deliveryType == "automated" {
		snapshotRequest.TargetClusterName = *targetClusterName
		snapshotRequest.TargetGroupID = *targetProjectId
	}

	restoreJob, _, err := client.CloudProviderSnapshotRestoreJobs.Create(context.Background(), requestParameters, snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error creating cloud provider snapshot restore job: %s", err)
	}

	currentModel.Id = encoding.NewString(restoreJob.ID)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Create Complete",
		ResourceModel:   currentModel,
	}, nil
}

// Read handles the Read event from the Cloudformation service.
func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId.Value()
	jobId := *currentModel.Id.Value()
	requestParameters := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: *currentModel.ClusterName.Value(),
		JobID:       jobId,
	}

	restoreJob, _, err := client.CloudProviderSnapshotRestoreJobs.Get(context.Background(), requestParameters)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot restore job with id(project: %s, job: %s): %s", projectId, jobId, err)
	}

	currentModel.TargetClusterName = encoding.NewString(restoreJob.TargetClusterName)
	currentModel.DeliveryType = encoding.NewString(restoreJob.DeliveryType)
	currentModel.ExpiresAt = encoding.NewString(restoreJob.ExpiresAt)
	currentModel.CreatedAt = encoding.NewString(restoreJob.CreatedAt)
	currentModel.Id = encoding.NewString(restoreJob.ID)
	currentModel.FinishedAt = encoding.NewString(restoreJob.FinishedAt)
	currentModel.SnapshotId = encoding.NewString(restoreJob.SnapshotID)
	currentModel.TargetProjectId = encoding.NewString(restoreJob.TargetGroupID)
	currentModel.Timestamp = encoding.NewString(restoreJob.Timestamp)
	currentModel.Cancelled = encoding.NewBool(restoreJob.Cancelled)
	currentModel.Expired = encoding.NewBool(restoreJob.Expired)
	currentModel.DeliveryUrl = flattenDeliveryUrl(restoreJob.DeliveryURL)
	currentModel.Links = flattenLinks(restoreJob.Links)

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read complete",
		ResourceModel:   currentModel,
	}, nil
}

// Update handles the Update event from the Cloudformation service.
func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	// operation not available :(
	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Update complete",
		ResourceModel:   currentModel,
	}, nil
}

// Delete handles the Delete event from the Cloudformation service.
func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId.Value()
	jobId := *currentModel.Id.Value()
	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: *currentModel.ClusterName.Value(),
		JobID:       jobId,
	}
	if *currentModel.DeliveryType.Value() == "automated" {
		return handler.ProgressEvent{
			OperationStatus: handler.Failed,
			Message:         "Automated restore cannot be cancelled",
			ResourceModel:   currentModel,
		}, nil
	}
	_, err = client.CloudProviderSnapshotRestoreJobs.Delete(context.Background(), snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error deleting cloud provider snapshot restore job with id(project: %s, job: %s): %s", projectId, jobId, err)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete complete",
		ResourceModel:   currentModel,
	}, nil
}

// List handles the List event from the Cloudformation service.
func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	client, err := util.CreateMongoDBClient(*currentModel.ApiKeys.PublicKey.Value(), *currentModel.ApiKeys.PrivateKey.Value())
	if err != nil {
		return handler.ProgressEvent{}, err
	}

	projectId := *currentModel.ProjectId.Value()
	snapshotRequest := &matlasClient.SnapshotReqPathParameters{
		GroupID:     projectId,
		ClusterName: *currentModel.ClusterName.Value(),
	}
	restoreJobs, _, err := client.CloudProviderSnapshotRestoreJobs.List(context.Background(), snapshotRequest)
	if err != nil {
		return handler.ProgressEvent{}, fmt.Errorf("error reading cloud provider snapshot restore job list with id(project: %s): %s", projectId, err)
	}

	var models []Model
	for _, restoreJob := range restoreJobs.Results {
		var model Model
		model.TargetClusterName = encoding.NewString(restoreJob.TargetClusterName)
		model.DeliveryType = encoding.NewString(restoreJob.DeliveryType)
		model.ExpiresAt = encoding.NewString(restoreJob.ExpiresAt)
		model.CreatedAt = encoding.NewString(restoreJob.CreatedAt)
		model.Id = encoding.NewString(restoreJob.ID)
		model.FinishedAt = encoding.NewString(restoreJob.FinishedAt)
		model.SnapshotId = encoding.NewString(restoreJob.SnapshotID)
		model.TargetProjectId = encoding.NewString(restoreJob.TargetGroupID)
		model.Timestamp = encoding.NewString(restoreJob.Timestamp)
		model.Cancelled = encoding.NewBool(restoreJob.Cancelled)
		model.Expired = encoding.NewBool(restoreJob.Expired)
		model.DeliveryUrl = flattenDeliveryUrl(restoreJob.DeliveryURL)
		model.Links = flattenLinks(restoreJob.Links)
		models = append(models, model)
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "List complete",
		ResourceModel:   models,
	}, nil
}

func flattenDeliveryUrl(deliveryUrlResult []string) []*encoding.String {
	deliveryUrls := make([]*encoding.String, 0)
	for _, deliveryUrl := range deliveryUrlResult {
		deliveryUrls = append(deliveryUrls, encoding.NewString(deliveryUrl))
	}
	return deliveryUrls
}

func flattenLinks(linksResult []*matlasClient.Link) []Links {
	links := make([]Links, 0)
	for _, link := range linksResult {
		var lin Links
		lin.Href = encoding.NewString(link.Href)
		lin.Rel = encoding.NewString(link.Rel)
		links = append(links, lin)
	}
	return links
}
