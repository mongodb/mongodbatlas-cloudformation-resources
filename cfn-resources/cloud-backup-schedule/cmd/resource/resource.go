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

package resource

import (
	"context"
	"errors"

	"github.com/aws-cloudformation/cloudformation-cli-go-plugin/cfn/handler"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/constants"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/progressevent"
	"github.com/mongodb/mongodbatlas-cloudformation-resources/util/validator"
	admin20231115014 "go.mongodb.org/atlas-sdk/v20231115014/admin"
)

var RequiredFields = []string{constants.ProjectID, constants.ClusterName}

func setup() {
	util.SetupLogger("mongodb-atlas-cloud-backup-schedule")
}

func Create(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(RequiredFields, currentModel); err != nil {
		return *err, nil
	}

	return cloudBackupScheduleCreateOrUpdate(req, prevModel, currentModel)
}

func Read(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(RequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	backupPolicy, resp, err := client.Atlas20231115014.CloudBackupsApi.GetBackupSchedule(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if pe := validateExist(backupPolicy); pe != nil {
		return *pe, nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Read Complete",
		ResourceModel:   currentModel.newModel(backupPolicy),
	}, nil
}

func Update(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(RequiredFields, currentModel); err != nil {
		return *err, nil
	}

	events, _ := Read(req, prevModel, currentModel)
	if events.HandlerErrorCode == cloudformation.HandlerErrorCodeNotFound {
		return handler.ProgressEvent{
			Message:          "Not Found",
			OperationStatus:  handler.Failed,
			HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}, nil
	}

	return cloudBackupScheduleCreateOrUpdate(req, prevModel, currentModel)
}

func Delete(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	setup()
	util.SetDefaultProfileIfNotDefined(&currentModel.Profile)
	if err := validator.ValidateModel(RequiredFields, currentModel); err != nil {
		return *err, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}

	backupPolicy, resp, err := client.Atlas20231115014.CloudBackupsApi.GetBackupSchedule(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	if pe := validateExist(backupPolicy); pe != nil {
		return *pe, nil
	}

	_, resp, err = client.Atlas20231115014.CloudBackupsApi.DeleteAllBackupSchedules(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "Delete Complete",
		ResourceModel:   nil,
	}, nil
}

func List(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	return handler.ProgressEvent{}, errors.New("not implemented: List")
}

func cloudBackupScheduleCreateOrUpdate(req handler.Request, prevModel *Model, currentModel *Model) (handler.ProgressEvent, error) {
	if event, err := validateExportDetails(currentModel); err != nil {
		return event, nil
	}

	if event, err := validatePolicies(currentModel); err != nil {
		return event, nil
	}

	client, pe := util.NewAtlasClient(&req, currentModel.Profile)
	if pe != nil {
		return *pe, nil
	}
	_, resp, err := client.Atlas20231115014.CloudBackupsApi.DeleteAllBackupSchedules(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	params := currentModel.getParams()

	// From https://jira.mongodb.org/browse/HELP-55421
	// Even after deleting and recreating the schedules, it is required to fetch the ID of the policy in order to be passed to the update request.
	if len(params.GetPolicies()) == 1 && params.GetPolicies()[0].GetId() == "" {
		backupSchedule, resp, err := client.Atlas20231115014.CloudBackupsApi.GetBackupSchedule(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName).Execute()
		if err != nil {
			return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
		} else if len(backupSchedule.GetPolicies()) == 1 {
			params.GetPolicies()[0].Id = backupSchedule.GetPolicies()[0].Id
		}
	}

	backupPolicy, resp, err := client.Atlas20231115014.CloudBackupsApi.UpdateBackupSchedule(context.Background(), *currentModel.ProjectId, *currentModel.ClusterName, params).Execute()
	if err != nil {
		return progressevent.GetFailedEventByResponse(err.Error(), resp), nil
	}

	return handler.ProgressEvent{
		OperationStatus: handler.Success,
		Message:         "PATCH Complete",
		ResourceModel:   currentModel.newModel(backupPolicy),
	}, nil
}

func validatePolicies(currentModel *Model) (pe handler.ProgressEvent, err error) {
	if len(currentModel.Policies) == 0 {
		msg := "validation error: policies cannot be empty"
		return progressevent.GetFailedEventByCode(msg, cloudformation.HandlerErrorCodeInvalidRequest), errors.New(msg)
	}
	for _, policy := range currentModel.Policies {
		if len(policy.PolicyItems) == 0 {
			msg := "validation error: policy items cannot be empty"
			return progressevent.GetFailedEventByCode(msg, cloudformation.HandlerErrorCodeInvalidRequest), errors.New(msg)
		}
		for _, policyItem := range policy.PolicyItems {
			if policyItem.FrequencyInterval == nil || policyItem.FrequencyType == nil ||
				policyItem.RetentionUnit == nil || policyItem.RetentionValue == nil {
				err := errors.New("validation error: All values from PolicyItem should be set when `PolicyItems` is set")
				return handler.ProgressEvent{
					OperationStatus:  handler.Failed,
					Message:          err.Error(),
					HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, err
			}
		}
	}
	return handler.ProgressEvent{}, nil
}

func validateExportDetails(currentModel *Model) (pe handler.ProgressEvent, err error) {
	if *currentModel.AutoExportEnabled && currentModel.Export != nil {
		if (currentModel.Export.FrequencyType) == nil {
			err := errors.New("error updating cloud backup schedule: FrequencyType should be set when `Export` is set")
			return handler.ProgressEvent{
				OperationStatus:  handler.Failed,
				Message:          err.Error(),
				HandlerErrorCode: cloudformation.HandlerErrorCodeInvalidRequest}, err
		}
	}
	return handler.ProgressEvent{}, nil
}

func validateExist(policy *admin20231115014.DiskBackupSnapshotSchedule) *handler.ProgressEvent {
	if policy.Policies != nil && len(policy.GetPolicies()) > 0 && len(policy.GetPolicies()[0].GetPolicyItems()) > 0 {
		return nil
	}
	return &handler.ProgressEvent{
		OperationStatus:  handler.Failed,
		Message:          "Not Found",
		HandlerErrorCode: cloudformation.HandlerErrorCodeNotFound}
}

func (m *Model) getParams() *admin20231115014.DiskBackupSnapshotSchedule {
	return &admin20231115014.DiskBackupSnapshotSchedule{
		AutoExportEnabled:                 m.AutoExportEnabled,
		ReferenceHourOfDay:                m.ReferenceHourOfDay,
		ReferenceMinuteOfHour:             m.ReferenceMinuteOfHour,
		RestoreWindowDays:                 m.RestoreWindowDays,
		NextSnapshot:                      util.StringPtrToTimePtr(m.NextSnapshot),
		UseOrgAndGroupNamesInExportPrefix: m.UseOrgAndGroupNamesInExportPrefix,
		Policies:                          expandPolicies(m.Policies),
		CopySettings:                      expandCopySettings(m.CopySettings),
		Export:                            expandExport(m.Export, aws.BoolValue(m.AutoExportEnabled)),
		UpdateSnapshots:                   m.UpdateSnapshots,
		DeleteCopiedBackups:               expandDeleteCopiedBackups(m.DeleteCopiedBackups),
	}
}

func (m *Model) newModel(policy *admin20231115014.DiskBackupSnapshotSchedule) *Model {
	return &Model{
		Profile:                           m.Profile,
		ProjectId:                         m.ProjectId,
		ClusterName:                       m.ClusterName,
		ClusterId:                         policy.ClusterId,
		AutoExportEnabled:                 policy.AutoExportEnabled,
		ReferenceHourOfDay:                policy.ReferenceHourOfDay,
		ReferenceMinuteOfHour:             policy.ReferenceMinuteOfHour,
		RestoreWindowDays:                 policy.RestoreWindowDays,
		NextSnapshot:                      util.TimePtrToStringPtr(policy.NextSnapshot),
		UseOrgAndGroupNamesInExportPrefix: policy.UseOrgAndGroupNamesInExportPrefix,
		Policies:                          flattenPolicies(policy.Policies),
		Links:                             flattenLinks(policy.Links),
		CopySettings:                      flattenCopySettings(policy.CopySettings),
		Export:                            flattenExport(policy.Export, aws.BoolValue(policy.AutoExportEnabled)),
	}
}

func expandExport(export *Export, enabled bool) *admin20231115014.AutoExportPolicy {
	if export == nil || !enabled {
		return nil
	}

	return &admin20231115014.AutoExportPolicy{
		ExportBucketId: export.ExportBucketId,
		FrequencyType:  export.FrequencyType,
	}
}

func flattenExport(export *admin20231115014.AutoExportPolicy, enabled bool) *Export {
	if export == nil || !enabled {
		return nil
	}
	return &Export{
		ExportBucketId: export.ExportBucketId,
		FrequencyType:  export.FrequencyType,
	}
}

func expandPolicies(policies []ApiPolicyView) *[]admin20231115014.AdvancedDiskBackupSnapshotSchedulePolicy {
	schedulePolicies := make([]admin20231115014.AdvancedDiskBackupSnapshotSchedulePolicy, 0)
	for _, s := range policies {
		policy := admin20231115014.AdvancedDiskBackupSnapshotSchedulePolicy{
			Id:          s.ID,
			PolicyItems: expandPolicyItems(s.PolicyItems),
		}
		schedulePolicies = append(schedulePolicies, policy)
	}
	return &schedulePolicies
}

func flattenPolicies(policies *[]admin20231115014.AdvancedDiskBackupSnapshotSchedulePolicy) []ApiPolicyView {
	snapPolicies := make([]ApiPolicyView, 0)
	for _, policy := range *policies {
		snapPolicy := ApiPolicyView{
			ID:          policy.Id,
			PolicyItems: flattenPolicyItems(*policy.PolicyItems),
		}
		snapPolicies = append(snapPolicies, snapPolicy)
	}
	return snapPolicies
}

func expandPolicyItems(cloudPolicyItems []ApiPolicyItemView) *[]admin20231115014.DiskBackupApiPolicyItem {
	policyItems := make([]admin20231115014.DiskBackupApiPolicyItem, 0)
	for _, policyItem := range cloudPolicyItems {
		cPolicyItem := admin20231115014.DiskBackupApiPolicyItem{
			Id:                policyItem.ID,
			FrequencyInterval: aws.IntValue(policyItem.FrequencyInterval),
			FrequencyType:     util.SafeString(policyItem.FrequencyType),
			RetentionUnit:     util.SafeString(policyItem.RetentionUnit),
			RetentionValue:    aws.IntValue(policyItem.RetentionValue),
		}
		policyItems = append(policyItems, cPolicyItem)
	}
	return &policyItems
}

func flattenPolicyItems(policyItems []admin20231115014.DiskBackupApiPolicyItem) []ApiPolicyItemView {
	cloudPolicyItems := make([]ApiPolicyItemView, 0)
	for _, policyItem := range policyItems {
		snapPolicy := ApiPolicyItemView{
			ID:                policyItem.Id,
			FrequencyInterval: &policyItem.FrequencyInterval,
			FrequencyType:     &policyItem.FrequencyType,
			RetentionUnit:     &policyItem.RetentionUnit,
			RetentionValue:    &policyItem.RetentionValue,
		}
		cloudPolicyItems = append(cloudPolicyItems, snapPolicy)
	}
	return cloudPolicyItems
}

func expandCopySettings(copySettings []ApiAtlasDiskBackupCopySettingView) *[]admin20231115014.DiskBackupCopySetting {
	cloudCopySettings := make([]admin20231115014.DiskBackupCopySetting, 0)
	for _, copySetting := range copySettings {
		backupSetting := admin20231115014.DiskBackupCopySetting{
			CloudProvider:     copySetting.CloudProvider,
			RegionName:        copySetting.RegionName,
			ReplicationSpecId: copySetting.ReplicationSpecId,
			ShouldCopyOplogs:  copySetting.ShouldCopyOplogs,
			Frequencies:       &copySetting.Frequencies,
		}
		cloudCopySettings = append(cloudCopySettings, backupSetting)
	}
	return &cloudCopySettings
}

func flattenCopySettings(copySettings *[]admin20231115014.DiskBackupCopySetting) []ApiAtlasDiskBackupCopySettingView {
	cloudCopySettings := make([]ApiAtlasDiskBackupCopySettingView, 0)
	for _, copySetting := range *copySettings {
		cloudCopySettings = append(cloudCopySettings, ApiAtlasDiskBackupCopySettingView{
			CloudProvider:     copySetting.CloudProvider,
			RegionName:        copySetting.RegionName,
			ReplicationSpecId: copySetting.ReplicationSpecId,
			ShouldCopyOplogs:  copySetting.ShouldCopyOplogs,
			Frequencies:       copySetting.GetFrequencies(),
		})
	}
	return cloudCopySettings
}

func expandDeleteCopiedBackups(deleteCopiedBackups []ApiDeleteCopiedBackupsView) *[]admin20231115014.DeleteCopiedBackups {
	cloudDeleteCopiedBackups := make([]admin20231115014.DeleteCopiedBackups, 0)
	for _, deleteCopiedBackup := range deleteCopiedBackups {
		copiedSetting := admin20231115014.DeleteCopiedBackups{
			CloudProvider:     deleteCopiedBackup.CloudProvider,
			RegionName:        deleteCopiedBackup.RegionName,
			ReplicationSpecId: deleteCopiedBackup.ReplicationSpecId,
		}
		cloudDeleteCopiedBackups = append(cloudDeleteCopiedBackups, copiedSetting)
	}
	return &cloudDeleteCopiedBackups
}

func flattenLinks(linksResult *[]admin20231115014.Link) []Link {
	links := make([]Link, 0)
	for _, link := range *linksResult {
		var lin Link
		lin.Href = link.Href
		lin.Rel = link.Rel
		links = append(links, lin)
	}
	return links
}
