// Generated by cdk-import
import * as cdk from 'aws-cdk-lib';
import * as constructs from 'constructs';

/**
 * An example resource schema demonstrating some basic constructs and validation rules.
 *
 * @schema CfnCloudBackupScheduleProps
 */
export interface CfnCloudBackupScheduleProps {
  /**
   * Unique identifier of the snapshot.
   *
   * @schema CfnCloudBackupScheduleProps#Id
   */
  readonly id?: string;

  /**
   * Unique 24-hexadecimal digit string that identifies your project.
   *
   * @schema CfnCloudBackupScheduleProps#ProjectId
   */
  readonly projectId: string;

  /**
   * The name of the Atlas cluster that contains the snapshots you want to retrieve.
   *
   * @schema CfnCloudBackupScheduleProps#ClusterName
   */
  readonly clusterName: string;

  /**
   * Flag that indicates whether automatic export of cloud backup snapshots to the AWS bucket is enabled.
   *
   * @schema CfnCloudBackupScheduleProps#AutoExportEnabled
   */
  readonly autoExportEnabled?: boolean;

  /**
   * Specify true to use organization and project names instead of organization and project UUIDs in the path for the metadata files that Atlas uploads to your S3 bucket after it finishes exporting the snapshots.
   *
   * @schema CfnCloudBackupScheduleProps#UseOrgAndGroupNamesInExportPrefix
   */
  readonly useOrgAndGroupNamesInExportPrefix?: boolean;

  /**
   * Policy for automatically exporting cloud backup snapshots.
   *
   * @schema CfnCloudBackupScheduleProps#Export
   */
  readonly export?: Export;

  /**
   * List that contains a document for each copy setting item in the desired backup policy.
   *
   * @schema CfnCloudBackupScheduleProps#CopySettings
   */
  readonly copySettings?: ApiAtlasDiskBackupCopySettingView[];

  /**
   * List that contains a document for each deleted copy setting whose backup copies you want to delete.
   *
   * @schema CfnCloudBackupScheduleProps#DeleteCopiedBackups
   */
  readonly deleteCopiedBackups?: ApiDeleteCopiedBackupsView[];

  /**
   * Rules set for this backup schedule.
   *
   * @schema CfnCloudBackupScheduleProps#Policies
   */
  readonly policies?: ApiPolicyView[];

  /**
   * UTC Hour of day between 0 and 23 representing which hour of the day that Atlas takes a snapshot
   *
   * @schema CfnCloudBackupScheduleProps#ReferenceHourOfDay
   */
  readonly referenceHourOfDay?: number;

  /**
   * UTC Minute of day between 0 and 59 representing which minute of the referenceHourOfDay that Atlas takes the snapshot.
   *
   * @schema CfnCloudBackupScheduleProps#ReferenceMinuteOfHour
   */
  readonly referenceMinuteOfHour?: number;

  /**
   * Number of days back in time you can restore to with Continuous Cloud Backup accuracy. Must be a positive, non-zero integer.
   *
   * @schema CfnCloudBackupScheduleProps#RestoreWindowDays
   */
  readonly restoreWindowDays?: number;

  /**
   * Flag indicating if updates to retention in the backup policy were applied to snapshots that Atlas took earlier.
   *
   * @schema CfnCloudBackupScheduleProps#UpdateSnapshots
   */
  readonly updateSnapshots?: boolean;

  /**
   * Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used
   *
   * @schema CfnCloudBackupScheduleProps#Profile
   */
  readonly profile: string;

  /**
   * List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. RFC 5988 outlines these relationships.
   *
   * @schema CfnCloudBackupScheduleProps#Links
   */
  readonly links?: Link[];

}

/**
 * Converts an object of type 'CfnCloudBackupScheduleProps' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_CfnCloudBackupScheduleProps(obj: CfnCloudBackupScheduleProps | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'Id': obj.id,
    'ProjectId': obj.projectId,
    'ClusterName': obj.clusterName,
    'AutoExportEnabled': obj.autoExportEnabled,
    'UseOrgAndGroupNamesInExportPrefix': obj.useOrgAndGroupNamesInExportPrefix,
    'Export': toJson_Export(obj.export),
    'CopySettings': obj.copySettings?.map(y => toJson_ApiAtlasDiskBackupCopySettingView(y)),
    'DeleteCopiedBackups': obj.deleteCopiedBackups?.map(y => toJson_ApiDeleteCopiedBackupsView(y)),
    'Policies': obj.policies?.map(y => toJson_ApiPolicyView(y)),
    'ReferenceHourOfDay': obj.referenceHourOfDay,
    'ReferenceMinuteOfHour': obj.referenceMinuteOfHour,
    'RestoreWindowDays': obj.restoreWindowDays,
    'UpdateSnapshots': obj.updateSnapshots,
    'Profile': obj.profile,
    'Links': obj.links?.map(y => toJson_Link(y)),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema Export
 */
export interface Export {
  /**
   * Unique identifier of the AWS bucket to export the cloud backup snapshot to
   *
   * @schema Export#ExportBucketId
   */
  readonly exportBucketId?: string;

  /**
   * Frequency associated with the export policy. Value can be daily, weekly, or monthly.
   *
   * @schema Export#FrequencyType
   */
  readonly frequencyType?: string;

}

/**
 * Converts an object of type 'Export' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_Export(obj: Export | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'ExportBucketId': obj.exportBucketId,
    'FrequencyType': obj.frequencyType,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema ApiAtlasDiskBackupCopySettingView
 */
export interface ApiAtlasDiskBackupCopySettingView {
  /**
   * A label that identifies the cloud provider that stores the snapshot copy.
   *
   * @schema ApiAtlasDiskBackupCopySettingView#CloudProvider
   */
  readonly cloudProvider?: string;

  /**
   * Target region to copy snapshots belonging to replicationSpecId to.
   *
   * @schema ApiAtlasDiskBackupCopySettingView#RegionName
   */
  readonly regionName?: string;

  /**
   * Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster.
   *
   * @schema ApiAtlasDiskBackupCopySettingView#ReplicationSpecId
   */
  readonly replicationSpecId?: string;

  /**
   * Flag that indicates whether to copy the oplogs to the target region.
   *
   * @schema ApiAtlasDiskBackupCopySettingView#ShouldCopyOplogs
   */
  readonly shouldCopyOplogs?: boolean;

  /**
   * List that describes which types of snapshots to copy.
   *
   * @schema ApiAtlasDiskBackupCopySettingView#Frequencies
   */
  readonly frequencies?: string[];

}

/**
 * Converts an object of type 'ApiAtlasDiskBackupCopySettingView' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_ApiAtlasDiskBackupCopySettingView(obj: ApiAtlasDiskBackupCopySettingView | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'CloudProvider': obj.cloudProvider,
    'RegionName': obj.regionName,
    'ReplicationSpecId': obj.replicationSpecId,
    'ShouldCopyOplogs': obj.shouldCopyOplogs,
    'Frequencies': obj.frequencies?.map(y => y),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema ApiDeleteCopiedBackupsView
 */
export interface ApiDeleteCopiedBackupsView {
  /**
   * A label that identifies the cloud provider for the deleted copy setting whose backup copies you want to delete
   *
   * @schema ApiDeleteCopiedBackupsView#CloudProvider
   */
  readonly cloudProvider?: string;

  /**
   * Target region for the deleted copy setting whose backup copies you want to delete.
   *
   * @schema ApiDeleteCopiedBackupsView#RegionName
   */
  readonly regionName?: string;

  /**
   * Unique 24-hexadecimal digit string that identifies the replication object for a zone in a cluster.
   *
   * @schema ApiDeleteCopiedBackupsView#ReplicationSpecId
   */
  readonly replicationSpecId?: string;

}

/**
 * Converts an object of type 'ApiDeleteCopiedBackupsView' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_ApiDeleteCopiedBackupsView(obj: ApiDeleteCopiedBackupsView | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'CloudProvider': obj.cloudProvider,
    'RegionName': obj.regionName,
    'ReplicationSpecId': obj.replicationSpecId,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema ApiPolicyView
 */
export interface ApiPolicyView {
  /**
   * @schema ApiPolicyView#ID
   */
  readonly id?: string;

  /**
   * @schema ApiPolicyView#PolicyItems
   */
  readonly policyItems?: ApiPolicyItemView[];

}

/**
 * Converts an object of type 'ApiPolicyView' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_ApiPolicyView(obj: ApiPolicyView | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'ID': obj.id,
    'PolicyItems': obj.policyItems?.map(y => toJson_ApiPolicyItemView(y)),
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema Link
 */
export interface Link {
  /**
   * Uniform Resource Locator (URL) that points another API resource to which this response has some relationship. This URL often begins with `https://mms.mongodb.com`.
   *
   * @schema Link#Href
   */
  readonly href?: string;

  /**
   * Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource. This URL often begins with `https://mms.mongodb.com`.
   *
   * @schema Link#Rel
   */
  readonly rel?: string;

}

/**
 * Converts an object of type 'Link' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_Link(obj: Link | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'Href': obj.href,
    'Rel': obj.rel,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */

/**
 * @schema ApiPolicyItemView
 */
export interface ApiPolicyItemView {
  /**
   * Unique identifier of the backup policy item.
   *
   * @schema ApiPolicyItemView#ID
   */
  readonly id?: string;

  /**
   * Frequency associated with the backup policy item. One of the following values: hourly, daily, weekly or monthly.
   *
   * @schema ApiPolicyItemView#FrequencyType
   */
  readonly frequencyType?: string;

  /**
   * Desired frequency of the new backup policy item specified by frequencyType.
   *
   * @schema ApiPolicyItemView#FrequencyInterval
   */
  readonly frequencyInterval?: number;

  /**
   * Duration for which the backup is kept. Associated with retentionUnit.
   *
   * @schema ApiPolicyItemView#RetentionValue
   */
  readonly retentionValue?: number;

  /**
   * Metric of duration of the backup policy item: days, weeks, or months.
   *
   * @schema ApiPolicyItemView#RetentionUnit
   */
  readonly retentionUnit?: string;

}

/**
 * Converts an object of type 'ApiPolicyItemView' to JSON representation.
 */
/* eslint-disable max-len, quote-props */
export function toJson_ApiPolicyItemView(obj: ApiPolicyItemView | undefined): Record<string, any> | undefined {
  if (obj === undefined) { return undefined; }
  const result = {
    'ID': obj.id,
    'FrequencyType': obj.frequencyType,
    'FrequencyInterval': obj.frequencyInterval,
    'RetentionValue': obj.retentionValue,
    'RetentionUnit': obj.retentionUnit,
  };
  // filter undefined values
  return Object.entries(result).reduce((r, i) => (i[1] === undefined) ? r : ({ ...r, [i[0]]: i[1] }), {});
}
/* eslint-enable max-len, quote-props */


/**
 * A CloudFormation `MongoDB::Atlas::CloudBackupSchedule`
 *
 * @cloudformationResource MongoDB::Atlas::CloudBackupSchedule
 * @stability external
 */
export class CfnCloudBackupSchedule extends cdk.CfnResource {
  /**
  * The CloudFormation resource type name for this resource class.
  */
  public static readonly CFN_RESOURCE_TYPE_NAME = 'MongoDB::Atlas::CloudBackupSchedule';

  /**
   * Resource props.
   */
  public readonly props: CfnCloudBackupScheduleProps;

  /**
   * Attribute `MongoDB::Atlas::CloudBackupSchedule.ClusterId`
   */
  public readonly attrClusterId: string;
  /**
   * Attribute `MongoDB::Atlas::CloudBackupSchedule.NextSnapshot`
   */
  public readonly attrNextSnapshot: string;

  /**
   * Create a new `MongoDB::Atlas::CloudBackupSchedule`.
   *
   * @param scope - scope in which this resource is defined
   * @param id    - scoped id of the resource
   * @param props - resource properties
   */
  constructor(scope: constructs.Construct, id: string, props: CfnCloudBackupScheduleProps) {
    super(scope, id, { type: CfnCloudBackupSchedule.CFN_RESOURCE_TYPE_NAME, properties: toJson_CfnCloudBackupScheduleProps(props)! });

    this.props = props;

    this.attrClusterId = cdk.Token.asString(this.getAtt('ClusterId'));
    this.attrNextSnapshot = cdk.Token.asString(this.getAtt('NextSnapshot'));
  }
}