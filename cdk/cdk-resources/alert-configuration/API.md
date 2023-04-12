# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnAlertConfiguration <a name="CfnAlertConfiguration" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration"></a>

A CloudFormation `MongoDB::Atlas::AlertConfiguration`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.Initializer"></a>

```typescript
import { CfnAlertConfiguration } from '@mongodbatlas-awscdk/alert-configuration'

new CfnAlertConfiguration(scope: Construct, id: string, props: CfnAlertConfigurationProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps">CfnAlertConfigurationProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps">CfnAlertConfigurationProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addMetadata"></a>

```typescript
public addMetadata(key: string, value: any): void
```

Add a value to the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addOverride"></a>

```typescript
public addOverride(path: string, value: any): void
```

Adds an override to the synthesized CloudFormation resource.

To add a
property override, either use `addPropertyOverride` or prefix `path` with
"Properties." (i.e. `Properties.TopicName`).

If the override is nested, separate each nested level using a dot (.) in the path parameter.
If there is an array as part of the nesting, specify the index in the path.

To include a literal `.` in the property name, prefix with a `\`. In most
programming languages you will need to write this as `"\\."` because the
`\` itself will need to be escaped.

For example,
```typescript
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.0.Projection.NonKeyAttributes', ['myattribute']);
cfnResource.addOverride('Properties.GlobalSecondaryIndexes.1.ProjectionType', 'INCLUDE');
```
would add the overrides
```json
"Properties": {
   "GlobalSecondaryIndexes": [
     {
       "Projection": {
         "NonKeyAttributes": [ "myattribute" ]
         ...
       }
       ...
     },
     {
       "ProjectionType": "INCLUDE"
       ...
     },
   ]
   ...
}
```

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.applyRemovalPolicy"></a>

```typescript
public applyRemovalPolicy(policy?: RemovalPolicy, options?: RemovalPolicyOptions): void
```

Sets the deletion policy of the resource based on the removal policy specified.

The Removal Policy controls what happens to this resource when it stops
being managed by CloudFormation, either because you've removed it from the
CDK application or because you've made a change that requires the resource
to be replaced.

The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.getMetadata"></a>

```typescript
public getMetadata(key: string): any
```

Retrieve a value value from the CloudFormation Resource Metadata.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/metadata-section-structure.html

Note that this is a different set of metadata from CDK node metadata; this
metadata ends up in the stack template under the resource, whereas CDK
node metadata ends up in the Cloud Assembly.)

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isConstruct"></a>

```typescript
import { CfnAlertConfiguration } from '@mongodbatlas-awscdk/alert-configuration'

CfnAlertConfiguration.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isCfnElement"></a>

```typescript
import { CfnAlertConfiguration } from '@mongodbatlas-awscdk/alert-configuration'

CfnAlertConfiguration.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isCfnResource"></a>

```typescript
import { CfnAlertConfiguration } from '@mongodbatlas-awscdk/alert-configuration'

CfnAlertConfiguration.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrEnabled">attrEnabled</a></code> | <code>aws-cdk-lib.IResolvable</code> | Attribute `MongoDB::Atlas::AlertConfiguration.Enabled`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::AlertConfiguration.Id`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrUpdated">attrUpdated</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::AlertConfiguration.Updated`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps">CfnAlertConfigurationProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrEnabled`<sup>Required</sup> <a name="attrEnabled" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrEnabled"></a>

```typescript
public readonly attrEnabled: IResolvable;
```

- *Type:* aws-cdk-lib.IResolvable

Attribute `MongoDB::Atlas::AlertConfiguration.Enabled`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::AlertConfiguration.Id`.

---

##### `attrUpdated`<sup>Required</sup> <a name="attrUpdated" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrUpdated"></a>

```typescript
public readonly attrUpdated: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::AlertConfiguration.Updated`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.props"></a>

```typescript
public readonly props: CfnAlertConfigurationProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps">CfnAlertConfigurationProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnAlertConfigurationProps <a name="CfnAlertConfigurationProps" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps"></a>

Returns and edits the conditions that trigger alerts and how MongoDB Cloud notifies users.

This collection remains under revision and may change. Refer to the legacy documentation for this collection in the following link.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.Initializer"></a>

```typescript
import { CfnAlertConfigurationProps } from '@mongodbatlas-awscdk/alert-configuration'

const cfnAlertConfigurationProps: CfnAlertConfigurationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.created">created</a></code> | <code>string</code> | Date and time when MongoDB Cloud created the alert configuration. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.eventTypeName">eventTypeName</a></code> | <code>string</code> | Event type that triggers an alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.matchers">matchers</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.Matcher">Matcher</a>[]</code> | List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.metricThreshold">metricThreshold</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView">MetricThresholdView</a></code> | Threshold for the metric that, when exceeded, triggers an alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.notifications">notifications</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView">NotificationView</a>[]</code> | List that contains the targets that MongoDB Cloud sends notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.profile">profile</a></code> | <code>string</code> | Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.threshold">threshold</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView">IntegerThresholdView</a></code> | Limit that triggers an alert when exceeded. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.typeName">typeName</a></code> | <code>string</code> | Human-readable label that displays the alert type. |

---

##### `created`<sup>Optional</sup> <a name="created" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.created"></a>

```typescript
public readonly created: string;
```

- *Type:* string

Date and time when MongoDB Cloud created the alert configuration.

This parameter expresses its value in the ISO 8601 timestamp format in UTC.

---

##### `eventTypeName`<sup>Optional</sup> <a name="eventTypeName" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.eventTypeName"></a>

```typescript
public readonly eventTypeName: string;
```

- *Type:* string

Event type that triggers an alert.

---

##### `matchers`<sup>Optional</sup> <a name="matchers" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.matchers"></a>

```typescript
public readonly matchers: Matcher[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.Matcher">Matcher</a>[]

List of rules that determine whether MongoDB Cloud checks an object for the alert configuration.

You can filter using the matchers array if the **eventTypeName** specifies an event for a host, replica set, or sharded cluster.

---

##### `metricThreshold`<sup>Optional</sup> <a name="metricThreshold" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.metricThreshold"></a>

```typescript
public readonly metricThreshold: MetricThresholdView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView">MetricThresholdView</a>

Threshold for the metric that, when exceeded, triggers an alert.

The resource returns this parameter when '"eventTypeName" : "OUTSIDE_METRIC_THRESHOLD"'.

---

##### `notifications`<sup>Optional</sup> <a name="notifications" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.notifications"></a>

```typescript
public readonly notifications: NotificationView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView">NotificationView</a>[]

List that contains the targets that MongoDB Cloud sends notifications.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `threshold`<sup>Optional</sup> <a name="threshold" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.threshold"></a>

```typescript
public readonly threshold: IntegerThresholdView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView">IntegerThresholdView</a>

Limit that triggers an alert when exceeded.

The resource returns this parameter when **eventTypeName** has not been set to 'OUTSIDE_METRIC_THRESHOLD'.

---

##### `typeName`<sup>Optional</sup> <a name="typeName" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.typeName"></a>

```typescript
public readonly typeName: string;
```

- *Type:* string

Human-readable label that displays the alert type.

---

### IntegerThresholdView <a name="IntegerThresholdView" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView.Initializer"></a>

```typescript
import { IntegerThresholdView } from '@mongodbatlas-awscdk/alert-configuration'

const integerThresholdView: IntegerThresholdView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView.property.operator">operator</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewOperator">IntegerThresholdViewOperator</a></code> | Comparison operator to apply when checking the current metric value. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView.property.threshold">threshold</a></code> | <code>number</code> | Value of metric that, when exceeded, triggers an alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView.property.units">units</a></code> | <code>string</code> | Element used to express the quantity. |

---

##### `operator`<sup>Optional</sup> <a name="operator" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView.property.operator"></a>

```typescript
public readonly operator: IntegerThresholdViewOperator;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewOperator">IntegerThresholdViewOperator</a>

Comparison operator to apply when checking the current metric value.

---

##### `threshold`<sup>Optional</sup> <a name="threshold" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView.property.threshold"></a>

```typescript
public readonly threshold: number;
```

- *Type:* number

Value of metric that, when exceeded, triggers an alert.

---

##### `units`<sup>Optional</sup> <a name="units" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView.property.units"></a>

```typescript
public readonly units: string;
```

- *Type:* string

Element used to express the quantity.

This can be an element of time, storage capacity, and the like.

---

### Matcher <a name="Matcher" id="@mongodbatlas-awscdk/alert-configuration.Matcher"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.Matcher.Initializer"></a>

```typescript
import { Matcher } from '@mongodbatlas-awscdk/alert-configuration'

const matcher: Matcher = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.Matcher.property.fieldName">fieldName</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName">MatcherFieldName</a></code> | Name of the parameter in the target object that MongoDB Cloud checks. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.Matcher.property.operator">operator</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator">MatcherOperator</a></code> | Comparison operator to apply when checking the current metric value against **matcher[n].value**. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.Matcher.property.value">value</a></code> | <code>string</code> | Value to match or exceed using the specified **matchers.operator**. |

---

##### `fieldName`<sup>Optional</sup> <a name="fieldName" id="@mongodbatlas-awscdk/alert-configuration.Matcher.property.fieldName"></a>

```typescript
public readonly fieldName: MatcherFieldName;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName">MatcherFieldName</a>

Name of the parameter in the target object that MongoDB Cloud checks.

The parameter must match all rules for MongoDB Cloud to check for alert configurations.

---

##### `operator`<sup>Optional</sup> <a name="operator" id="@mongodbatlas-awscdk/alert-configuration.Matcher.property.operator"></a>

```typescript
public readonly operator: MatcherOperator;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator">MatcherOperator</a>

Comparison operator to apply when checking the current metric value against **matcher[n].value**.

---

##### `value`<sup>Optional</sup> <a name="value" id="@mongodbatlas-awscdk/alert-configuration.Matcher.property.value"></a>

```typescript
public readonly value: string;
```

- *Type:* string

Value to match or exceed using the specified **matchers.operator**.

---

### MetricThresholdView <a name="MetricThresholdView" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.Initializer"></a>

```typescript
import { MetricThresholdView } from '@mongodbatlas-awscdk/alert-configuration'

const metricThresholdView: MetricThresholdView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.metricName">metricName</a></code> | <code>string</code> | Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.mode">mode</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMode">MetricThresholdViewMode</a></code> | MongoDB Cloud computes the current metric value as an average. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.operator">operator</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewOperator">MetricThresholdViewOperator</a></code> | Comparison operator to apply when checking the current metric value. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.threshold">threshold</a></code> | <code>number</code> | Value of metric that, when exceeded, triggers an alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.units">units</a></code> | <code>string</code> | Element used to express the quantity. |

---

##### `metricName`<sup>Optional</sup> <a name="metricName" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.metricName"></a>

```typescript
public readonly metricName: string;
```

- *Type:* string

Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**.

---

##### `mode`<sup>Optional</sup> <a name="mode" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.mode"></a>

```typescript
public readonly mode: MetricThresholdViewMode;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMode">MetricThresholdViewMode</a>

MongoDB Cloud computes the current metric value as an average.

---

##### `operator`<sup>Optional</sup> <a name="operator" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.operator"></a>

```typescript
public readonly operator: MetricThresholdViewOperator;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewOperator">MetricThresholdViewOperator</a>

Comparison operator to apply when checking the current metric value.

---

##### `threshold`<sup>Optional</sup> <a name="threshold" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.threshold"></a>

```typescript
public readonly threshold: number;
```

- *Type:* number

Value of metric that, when exceeded, triggers an alert.

---

##### `units`<sup>Optional</sup> <a name="units" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.units"></a>

```typescript
public readonly units: string;
```

- *Type:* string

Element used to express the quantity.

This can be an element of time, storage capacity, and the like.

---

### NotificationView <a name="NotificationView" id="@mongodbatlas-awscdk/alert-configuration.NotificationView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.Initializer"></a>

```typescript
import { NotificationView } from '@mongodbatlas-awscdk/alert-configuration'

const notificationView: NotificationView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.apiToken">apiToken</a></code> | <code>string</code> | Slack API token or Bot token that MongoDB Cloud needs to send alert notifications via Slack. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.channelName">channelName</a></code> | <code>string</code> | Name of the Slack channel to which MongoDB Cloud sends alert notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.datadogApiKey">datadogApiKey</a></code> | <code>string</code> | Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.datadogRegion">datadogRegion</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewDatadogRegion">NotificationViewDatadogRegion</a></code> | Datadog region that indicates which API Uniform Resource Locator (URL) to use. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.delayMin">delayMin</a></code> | <code>number</code> | Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.emailAddress">emailAddress</a></code> | <code>string</code> | Email address to which MongoDB Cloud sends alert notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.emailEnabled">emailEnabled</a></code> | <code>boolean</code> | Flag that indicates whether MongoDB Cloud should send email notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.flowdockApiToken">flowdockApiToken</a></code> | <code>string</code> | Flowdock API token that MongoDB Cloud needs to send alert notifications to Flowdock. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.flowName">flowName</a></code> | <code>string</code> | Flowdock flow name to which MongoDB Cloud sends alert notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.intervalMin">intervalMin</a></code> | <code>number</code> | Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.microsoftTeamsWebhookUrl">microsoftTeamsWebhookUrl</a></code> | <code>string</code> | Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.mobileNumber">mobileNumber</a></code> | <code>string</code> | Mobile phone number to which MongoDB Cloud sends alert notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.notificationToken">notificationToken</a></code> | <code>string</code> | HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.opsGenieApiKey">opsGenieApiKey</a></code> | <code>string</code> | API Key that MongoDB Cloud needs to send this notification via Opsgenie. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.opsGenieRegion">opsGenieRegion</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewOpsGenieRegion">NotificationViewOpsGenieRegion</a></code> | Opsgenie region that indicates which API Uniform Resource Locator (URL) to use. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.orgName">orgName</a></code> | <code>string</code> | Flowdock organization name to which MongoDB Cloud sends alert notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.roles">roles</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles">NotificationViewRoles</a>[]</code> | List that contains the one or more organization or project roles that receive the configured alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.roomName">roomName</a></code> | <code>string</code> | HipChat API room name to which MongoDB Cloud sends alert notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.serviceKey">serviceKey</a></code> | <code>string</code> | PagerDuty service key that MongoDB Cloud needs to send notifications via PagerDuty. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.severity">severity</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity">NotificationViewSeverity</a></code> | Degree of seriousness given to this notification. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.smsEnabled">smsEnabled</a></code> | <code>boolean</code> | Flag that indicates whether MongoDB Cloud should send text message notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.teamId">teamId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies one MongoDB Cloud team. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.teamName">teamName</a></code> | <code>string</code> | Name of the MongoDB Cloud team that receives this notification. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.typeName">typeName</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName">NotificationViewTypeName</a></code> | Human-readable label that displays the alert notification type. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.username">username</a></code> | <code>string</code> | MongoDB Cloud username of the person to whom MongoDB Cloud sends notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.victorOpsApiKey">victorOpsApiKey</a></code> | <code>string</code> | API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.victorOpsRoutingKey">victorOpsRoutingKey</a></code> | <code>string</code> | Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.webhookSecret">webhookSecret</a></code> | <code>string</code> | An optional field for your webhook secret. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView.property.webhookUrl">webhookUrl</a></code> | <code>string</code> | Your webhook URL. |

---

##### `apiToken`<sup>Optional</sup> <a name="apiToken" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.apiToken"></a>

```typescript
public readonly apiToken: string;
```

- *Type:* string

Slack API token or Bot token that MongoDB Cloud needs to send alert notifications via Slack.

The resource requires this parameter when '"notifications.typeName" : "SLACK"'. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes the token.

---

##### `channelName`<sup>Optional</sup> <a name="channelName" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.channelName"></a>

```typescript
public readonly channelName: string;
```

- *Type:* string

Name of the Slack channel to which MongoDB Cloud sends alert notifications.

The resource requires this parameter when '"notifications.typeName" : "SLACK"'.

---

##### `datadogApiKey`<sup>Optional</sup> <a name="datadogApiKey" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.datadogApiKey"></a>

```typescript
public readonly datadogApiKey: string;
```

- *Type:* string

Datadog API Key that MongoDB Cloud needs to send alert notifications to Datadog.

You can find this API key in the Datadog dashboard. The resource requires this parameter when '"notifications.typeName" : "DATADOG"'.

---

##### `datadogRegion`<sup>Optional</sup> <a name="datadogRegion" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.datadogRegion"></a>

```typescript
public readonly datadogRegion: NotificationViewDatadogRegion;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewDatadogRegion">NotificationViewDatadogRegion</a>

Datadog region that indicates which API Uniform Resource Locator (URL) to use.

The resource requires this parameter when '"notifications.typeName" : "DATADOG"'.

---

##### `delayMin`<sup>Optional</sup> <a name="delayMin" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.delayMin"></a>

```typescript
public readonly delayMin: number;
```

- *Type:* number

Number of minutes that MongoDB Cloud waits after detecting an alert condition before it sends out the first notification.

---

##### `emailAddress`<sup>Optional</sup> <a name="emailAddress" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.emailAddress"></a>

```typescript
public readonly emailAddress: string;
```

- *Type:* string

Email address to which MongoDB Cloud sends alert notifications.

The resource requires this parameter when '"notifications.typeName" : "EMAIL"'. You don't need to set this value to send emails to individual or groups of MongoDB Cloud users including:

- specific MongoDB Cloud users ('"notifications.typeName" : "USER"')
- MongoDB Cloud users with specific project roles ('"notifications.typeName" : "GROUP"')
- MongoDB Cloud users with specific organization roles ('"notifications.typeName" : "ORG"')
- MongoDB Cloud teams ('"notifications.typeName" : "TEAM"')

To send emails to one MongoDB Cloud user or grouping of users, set the **notifications.emailEnabled** parameter.

---

##### `emailEnabled`<sup>Optional</sup> <a name="emailEnabled" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.emailEnabled"></a>

```typescript
public readonly emailEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether MongoDB Cloud should send email notifications.

The resource requires this parameter when one of the following values have been set:

- '"notifications.typeName" : "ORG"'
- '"notifications.typeName" : "GROUP"'
- '"notifications.typeName" : "USER"'

---

##### `flowdockApiToken`<sup>Optional</sup> <a name="flowdockApiToken" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.flowdockApiToken"></a>

```typescript
public readonly flowdockApiToken: string;
```

- *Type:* string

Flowdock API token that MongoDB Cloud needs to send alert notifications to Flowdock.

The resource requires this parameter when '"notifications.typeName" : "FLOWDOCK"'. If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes the token.

---

##### `flowName`<sup>Optional</sup> <a name="flowName" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.flowName"></a>

```typescript
public readonly flowName: string;
```

- *Type:* string

Flowdock flow name to which MongoDB Cloud sends alert notifications.

This name appears after the organization name in the Uniform Resource Locator (URL) path: 'www.flowdock.com/app/<organization-name>/<flow-name>'. The resource requires this parameter when '"notifications.typeName" : "FLOWDOCK"'.

---

##### `intervalMin`<sup>Optional</sup> <a name="intervalMin" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.intervalMin"></a>

```typescript
public readonly intervalMin: number;
```

- *Type:* number

Number of minutes to wait between successive notifications. MongoDB Cloud sends notifications until someone acknowledges the unacknowledged alert.

PagerDuty, VictorOps, and OpsGenie notifications don't return this element. Configure and manage the notification interval within each of those services.

---

##### `microsoftTeamsWebhookUrl`<sup>Optional</sup> <a name="microsoftTeamsWebhookUrl" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.microsoftTeamsWebhookUrl"></a>

```typescript
public readonly microsoftTeamsWebhookUrl: string;
```

- *Type:* string

Microsoft Teams Webhook Uniform Resource Locator (URL) that MongoDB Cloud needs to send this notification via Microsoft Teams.

The resource requires this parameter when '"notifications.typeName" : "MICROSOFT_TEAMS"'. If the URL later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

---

##### `mobileNumber`<sup>Optional</sup> <a name="mobileNumber" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.mobileNumber"></a>

```typescript
public readonly mobileNumber: string;
```

- *Type:* string

Mobile phone number to which MongoDB Cloud sends alert notifications.

The resource requires this parameter when '"notifications.typeName" : "SMS"'.

---

##### `notificationToken`<sup>Optional</sup> <a name="notificationToken" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.notificationToken"></a>

```typescript
public readonly notificationToken: string;
```

- *Type:* string

HipChat API token that MongoDB Cloud needs to send alert notifications to HipChat.

The resource requires this parameter when '"notifications.typeName" : "HIP_CHAT"'". If the token later becomes invalid, MongoDB Cloud sends an email to the project owners. If the token remains invalid, MongoDB Cloud removes it.

---

##### `opsGenieApiKey`<sup>Optional</sup> <a name="opsGenieApiKey" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.opsGenieApiKey"></a>

```typescript
public readonly opsGenieApiKey: string;
```

- *Type:* string

API Key that MongoDB Cloud needs to send this notification via Opsgenie.

The resource requires this parameter when '"notifications.typeName" : "OPS_GENIE"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

---

##### `opsGenieRegion`<sup>Optional</sup> <a name="opsGenieRegion" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.opsGenieRegion"></a>

```typescript
public readonly opsGenieRegion: NotificationViewOpsGenieRegion;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewOpsGenieRegion">NotificationViewOpsGenieRegion</a>

Opsgenie region that indicates which API Uniform Resource Locator (URL) to use.

---

##### `orgName`<sup>Optional</sup> <a name="orgName" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.orgName"></a>

```typescript
public readonly orgName: string;
```

- *Type:* string

Flowdock organization name to which MongoDB Cloud sends alert notifications.

This name appears after 'www.flowdock.com/app/' in the Uniform Resource Locator (URL) path. The resource requires this parameter when '"notifications.typeName" : "FLOWDOCK"'.

---

##### `roles`<sup>Optional</sup> <a name="roles" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.roles"></a>

```typescript
public readonly roles: NotificationViewRoles[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles">NotificationViewRoles</a>[]

List that contains the one or more organization or project roles that receive the configured alert.

The resource requires this parameter when '"notifications.typeName" : "GROUP"' or '"notifications.typeName" : "ORG"'. If you include this parameter, MongoDB Cloud sends alerts only to users assigned the roles you specify in the array. If you omit this parameter, MongoDB Cloud sends alerts to users assigned any role.

---

##### `roomName`<sup>Optional</sup> <a name="roomName" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.roomName"></a>

```typescript
public readonly roomName: string;
```

- *Type:* string

HipChat API room name to which MongoDB Cloud sends alert notifications.

The resource requires this parameter when '"notifications.typeName" : "HIP_CHAT"'".

---

##### `serviceKey`<sup>Optional</sup> <a name="serviceKey" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.serviceKey"></a>

```typescript
public readonly serviceKey: string;
```

- *Type:* string

PagerDuty service key that MongoDB Cloud needs to send notifications via PagerDuty.

The resource requires this parameter when '"notifications.typeName" : "PAGER_DUTY"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

---

##### `severity`<sup>Optional</sup> <a name="severity" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.severity"></a>

```typescript
public readonly severity: NotificationViewSeverity;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity">NotificationViewSeverity</a>

Degree of seriousness given to this notification.

---

##### `smsEnabled`<sup>Optional</sup> <a name="smsEnabled" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.smsEnabled"></a>

```typescript
public readonly smsEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether MongoDB Cloud should send text message notifications.

The resource requires this parameter when one of the following values have been set:

- '"notifications.typeName" : "ORG"'
- '"notifications.typeName" : "GROUP"'
- '"notifications.typeName" : "USER"'

---

##### `teamId`<sup>Optional</sup> <a name="teamId" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.teamId"></a>

```typescript
public readonly teamId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies one MongoDB Cloud team.

The resource requires this parameter when '"notifications.typeName" : "TEAM"'.

---

##### `teamName`<sup>Optional</sup> <a name="teamName" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.teamName"></a>

```typescript
public readonly teamName: string;
```

- *Type:* string

Name of the MongoDB Cloud team that receives this notification.

The resource requires this parameter when '"notifications.typeName" : "TEAM"'.

---

##### `typeName`<sup>Optional</sup> <a name="typeName" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.typeName"></a>

```typescript
public readonly typeName: NotificationViewTypeName;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName">NotificationViewTypeName</a>

Human-readable label that displays the alert notification type.

---

##### `username`<sup>Optional</sup> <a name="username" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.username"></a>

```typescript
public readonly username: string;
```

- *Type:* string

MongoDB Cloud username of the person to whom MongoDB Cloud sends notifications.

Specify only MongoDB Cloud users who belong to the project that owns the alert configuration. The resource requires this parameter when '"notifications.typeName" : "USER"'.

---

##### `victorOpsApiKey`<sup>Optional</sup> <a name="victorOpsApiKey" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.victorOpsApiKey"></a>

```typescript
public readonly victorOpsApiKey: string;
```

- *Type:* string

API key that MongoDB Cloud needs to send alert notifications to Splunk On-Call.

The resource requires this parameter when '"notifications.typeName" : "VICTOR_OPS"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

---

##### `victorOpsRoutingKey`<sup>Optional</sup> <a name="victorOpsRoutingKey" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.victorOpsRoutingKey"></a>

```typescript
public readonly victorOpsRoutingKey: string;
```

- *Type:* string

Routing key that MongoDB Cloud needs to send alert notifications to Splunk On-Call.

The resource requires this parameter when '"notifications.typeName" : "VICTOR_OPS"'. If the key later becomes invalid, MongoDB Cloud sends an email to the project owners. If the key remains invalid, MongoDB Cloud removes it.

---

##### `webhookSecret`<sup>Optional</sup> <a name="webhookSecret" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.webhookSecret"></a>

```typescript
public readonly webhookSecret: string;
```

- *Type:* string

An optional field for your webhook secret.

---

##### `webhookUrl`<sup>Optional</sup> <a name="webhookUrl" id="@mongodbatlas-awscdk/alert-configuration.NotificationView.property.webhookUrl"></a>

```typescript
public readonly webhookUrl: string;
```

- *Type:* string

Your webhook URL.

---



## Enums <a name="Enums" id="Enums"></a>

### IntegerThresholdViewOperator <a name="IntegerThresholdViewOperator" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewOperator"></a>

Comparison operator to apply when checking the current metric value.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewOperator.GREATER_THAN">GREATER_THAN</a></code> | GREATER_THAN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewOperator.LESS_THAN">LESS_THAN</a></code> | LESS_THAN. |

---

##### `GREATER_THAN` <a name="GREATER_THAN" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewOperator.GREATER_THAN"></a>

GREATER_THAN.

---


##### `LESS_THAN` <a name="LESS_THAN" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewOperator.LESS_THAN"></a>

LESS_THAN.

---


### MatcherFieldName <a name="MatcherFieldName" id="@mongodbatlas-awscdk/alert-configuration.MatcherFieldName"></a>

Name of the parameter in the target object that MongoDB Cloud checks.

The parameter must match all rules for MongoDB Cloud to check for alert configurations.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.CLUSTER_NAME">CLUSTER_NAME</a></code> | CLUSTER_NAME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.HOSTNAME">HOSTNAME</a></code> | HOSTNAME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.HOSTNAME_AND_PORT">HOSTNAME_AND_PORT</a></code> | HOSTNAME_AND_PORT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.PORT">PORT</a></code> | PORT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.REPLICA_SET_NAME">REPLICA_SET_NAME</a></code> | REPLICA_SET_NAME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.SHARD_NAME">SHARD_NAME</a></code> | SHARD_NAME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.TYPE_NAME">TYPE_NAME</a></code> | TYPE_NAME. |

---

##### `CLUSTER_NAME` <a name="CLUSTER_NAME" id="@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.CLUSTER_NAME"></a>

CLUSTER_NAME.

---


##### `HOSTNAME` <a name="HOSTNAME" id="@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.HOSTNAME"></a>

HOSTNAME.

---


##### `HOSTNAME_AND_PORT` <a name="HOSTNAME_AND_PORT" id="@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.HOSTNAME_AND_PORT"></a>

HOSTNAME_AND_PORT.

---


##### `PORT` <a name="PORT" id="@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.PORT"></a>

PORT.

---


##### `REPLICA_SET_NAME` <a name="REPLICA_SET_NAME" id="@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.REPLICA_SET_NAME"></a>

REPLICA_SET_NAME.

---


##### `SHARD_NAME` <a name="SHARD_NAME" id="@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.SHARD_NAME"></a>

SHARD_NAME.

---


##### `TYPE_NAME` <a name="TYPE_NAME" id="@mongodbatlas-awscdk/alert-configuration.MatcherFieldName.TYPE_NAME"></a>

TYPE_NAME.

---


### MatcherOperator <a name="MatcherOperator" id="@mongodbatlas-awscdk/alert-configuration.MatcherOperator"></a>

Comparison operator to apply when checking the current metric value against **matcher[n].value**.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator.EQUALS">EQUALS</a></code> | EQUALS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator.CONTAINS">CONTAINS</a></code> | CONTAINS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator.STARTS_WITH">STARTS_WITH</a></code> | STARTS_WITH. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator.ENDS_WITH">ENDS_WITH</a></code> | ENDS_WITH. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator.NOT_EQUALS">NOT_EQUALS</a></code> | NOT_EQUALS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator.NOT_CONTAINS">NOT_CONTAINS</a></code> | NOT_CONTAINS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MatcherOperator.REGEX">REGEX</a></code> | REGEX. |

---

##### `EQUALS` <a name="EQUALS" id="@mongodbatlas-awscdk/alert-configuration.MatcherOperator.EQUALS"></a>

EQUALS.

---


##### `CONTAINS` <a name="CONTAINS" id="@mongodbatlas-awscdk/alert-configuration.MatcherOperator.CONTAINS"></a>

CONTAINS.

---


##### `STARTS_WITH` <a name="STARTS_WITH" id="@mongodbatlas-awscdk/alert-configuration.MatcherOperator.STARTS_WITH"></a>

STARTS_WITH.

---


##### `ENDS_WITH` <a name="ENDS_WITH" id="@mongodbatlas-awscdk/alert-configuration.MatcherOperator.ENDS_WITH"></a>

ENDS_WITH.

---


##### `NOT_EQUALS` <a name="NOT_EQUALS" id="@mongodbatlas-awscdk/alert-configuration.MatcherOperator.NOT_EQUALS"></a>

NOT_EQUALS.

---


##### `NOT_CONTAINS` <a name="NOT_CONTAINS" id="@mongodbatlas-awscdk/alert-configuration.MatcherOperator.NOT_CONTAINS"></a>

NOT_CONTAINS.

---


##### `REGEX` <a name="REGEX" id="@mongodbatlas-awscdk/alert-configuration.MatcherOperator.REGEX"></a>

REGEX.

---


### MetricThresholdViewMode <a name="MetricThresholdViewMode" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMode"></a>

MongoDB Cloud computes the current metric value as an average.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMode.AVERAGE">AVERAGE</a></code> | AVERAGE. |

---

##### `AVERAGE` <a name="AVERAGE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMode.AVERAGE"></a>

AVERAGE.

---


### MetricThresholdViewOperator <a name="MetricThresholdViewOperator" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewOperator"></a>

Comparison operator to apply when checking the current metric value.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewOperator.GREATER_THAN">GREATER_THAN</a></code> | GREATER_THAN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewOperator.LESS_THAN">LESS_THAN</a></code> | LESS_THAN. |

---

##### `GREATER_THAN` <a name="GREATER_THAN" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewOperator.GREATER_THAN"></a>

GREATER_THAN.

---


##### `LESS_THAN` <a name="LESS_THAN" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewOperator.LESS_THAN"></a>

LESS_THAN.

---


### NotificationViewDatadogRegion <a name="NotificationViewDatadogRegion" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewDatadogRegion"></a>

Datadog region that indicates which API Uniform Resource Locator (URL) to use.

The resource requires this parameter when '"notifications.typeName" : "DATADOG"'.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewDatadogRegion.EU">EU</a></code> | EU. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewDatadogRegion.US">US</a></code> | US. |

---

##### `EU` <a name="EU" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewDatadogRegion.EU"></a>

EU.

---


##### `US` <a name="US" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewDatadogRegion.US"></a>

US.

---


### NotificationViewOpsGenieRegion <a name="NotificationViewOpsGenieRegion" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewOpsGenieRegion"></a>

Opsgenie region that indicates which API Uniform Resource Locator (URL) to use.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewOpsGenieRegion.EU">EU</a></code> | EU. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewOpsGenieRegion.US">US</a></code> | US. |

---

##### `EU` <a name="EU" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewOpsGenieRegion.EU"></a>

EU.

---


##### `US` <a name="US" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewOpsGenieRegion.US"></a>

US.

---


### NotificationViewRoles <a name="NotificationViewRoles" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_CLUSTER_MANAGER">GROUP_CLUSTER_MANAGER</a></code> | GROUP_CLUSTER_MANAGER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_DATA_ACCESS_ADMIN">GROUP_DATA_ACCESS_ADMIN</a></code> | GROUP_DATA_ACCESS_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_DATA_ACCESS_READ_ONLY">GROUP_DATA_ACCESS_READ_ONLY</a></code> | GROUP_DATA_ACCESS_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_DATA_ACCESS_READ_WRITE">GROUP_DATA_ACCESS_READ_WRITE</a></code> | GROUP_DATA_ACCESS_READ_WRITE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_OWNER">GROUP_OWNER</a></code> | GROUP_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_READ_WRITE">GROUP_READ_WRITE</a></code> | GROUP_READ_WRITE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_OWNER">ORG_OWNER</a></code> | ORG_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_MEMBER">ORG_MEMBER</a></code> | ORG_MEMBER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_GROUP_CREATOR">ORG_GROUP_CREATOR</a></code> | ORG_GROUP_CREATOR. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_BILLING_ADMIN">ORG_BILLING_ADMIN</a></code> | ORG_BILLING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_READ_ONLY">ORG_READ_ONLY</a></code> | ORG_READ_ONLY. |

---

##### `GROUP_CLUSTER_MANAGER` <a name="GROUP_CLUSTER_MANAGER" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_CLUSTER_MANAGER"></a>

GROUP_CLUSTER_MANAGER.

---


##### `GROUP_DATA_ACCESS_ADMIN` <a name="GROUP_DATA_ACCESS_ADMIN" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_DATA_ACCESS_ADMIN"></a>

GROUP_DATA_ACCESS_ADMIN.

---


##### `GROUP_DATA_ACCESS_READ_ONLY` <a name="GROUP_DATA_ACCESS_READ_ONLY" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_DATA_ACCESS_READ_ONLY"></a>

GROUP_DATA_ACCESS_READ_ONLY.

---


##### `GROUP_DATA_ACCESS_READ_WRITE` <a name="GROUP_DATA_ACCESS_READ_WRITE" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_DATA_ACCESS_READ_WRITE"></a>

GROUP_DATA_ACCESS_READ_WRITE.

---


##### `GROUP_OWNER` <a name="GROUP_OWNER" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_OWNER"></a>

GROUP_OWNER.

---


##### `GROUP_READ_WRITE` <a name="GROUP_READ_WRITE" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.GROUP_READ_WRITE"></a>

GROUP_READ_WRITE.

---


##### `ORG_OWNER` <a name="ORG_OWNER" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_OWNER"></a>

ORG_OWNER.

---


##### `ORG_MEMBER` <a name="ORG_MEMBER" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_MEMBER"></a>

ORG_MEMBER.

---


##### `ORG_GROUP_CREATOR` <a name="ORG_GROUP_CREATOR" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_GROUP_CREATOR"></a>

ORG_GROUP_CREATOR.

---


##### `ORG_BILLING_ADMIN` <a name="ORG_BILLING_ADMIN" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_BILLING_ADMIN"></a>

ORG_BILLING_ADMIN.

---


##### `ORG_READ_ONLY` <a name="ORG_READ_ONLY" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewRoles.ORG_READ_ONLY"></a>

ORG_READ_ONLY.

---


### NotificationViewSeverity <a name="NotificationViewSeverity" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity"></a>

Degree of seriousness given to this notification.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity.CRITICAL">CRITICAL</a></code> | CRITICAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity.ERROR">ERROR</a></code> | ERROR. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity.WARNING">WARNING</a></code> | WARNING. |

---

##### `CRITICAL` <a name="CRITICAL" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity.CRITICAL"></a>

CRITICAL.

---


##### `ERROR` <a name="ERROR" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity.ERROR"></a>

ERROR.

---


##### `WARNING` <a name="WARNING" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewSeverity.WARNING"></a>

WARNING.

---


### NotificationViewTypeName <a name="NotificationViewTypeName" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName"></a>

Human-readable label that displays the alert notification type.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.DATADOG">DATADOG</a></code> | DATADOG. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.EMAIL">EMAIL</a></code> | EMAIL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.FLOWDOCK">FLOWDOCK</a></code> | FLOWDOCK. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.GROUP">GROUP</a></code> | GROUP. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.MICROSOFT_TEAMS">MICROSOFT_TEAMS</a></code> | MICROSOFT_TEAMS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.OPS_GENIE">OPS_GENIE</a></code> | OPS_GENIE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.ORG">ORG</a></code> | ORG. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.PAGER_DUTY">PAGER_DUTY</a></code> | PAGER_DUTY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.PROMETHEUS">PROMETHEUS</a></code> | PROMETHEUS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.SLACK">SLACK</a></code> | SLACK. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.SMS">SMS</a></code> | SMS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.TEAM">TEAM</a></code> | TEAM. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.USER">USER</a></code> | USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.VICTOR_OPS">VICTOR_OPS</a></code> | VICTOR_OPS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.WEBHOOK">WEBHOOK</a></code> | WEBHOOK. |

---

##### `DATADOG` <a name="DATADOG" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.DATADOG"></a>

DATADOG.

---


##### `EMAIL` <a name="EMAIL" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.EMAIL"></a>

EMAIL.

---


##### `FLOWDOCK` <a name="FLOWDOCK" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.FLOWDOCK"></a>

FLOWDOCK.

---


##### `GROUP` <a name="GROUP" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.GROUP"></a>

GROUP.

---


##### `MICROSOFT_TEAMS` <a name="MICROSOFT_TEAMS" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.MICROSOFT_TEAMS"></a>

MICROSOFT_TEAMS.

---


##### `OPS_GENIE` <a name="OPS_GENIE" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.OPS_GENIE"></a>

OPS_GENIE.

---


##### `ORG` <a name="ORG" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.ORG"></a>

ORG.

---


##### `PAGER_DUTY` <a name="PAGER_DUTY" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.PAGER_DUTY"></a>

PAGER_DUTY.

---


##### `PROMETHEUS` <a name="PROMETHEUS" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.PROMETHEUS"></a>

PROMETHEUS.

---


##### `SLACK` <a name="SLACK" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.SLACK"></a>

SLACK.

---


##### `SMS` <a name="SMS" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.SMS"></a>

SMS.

---


##### `TEAM` <a name="TEAM" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.TEAM"></a>

TEAM.

---


##### `USER` <a name="USER" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.USER"></a>

USER.

---


##### `VICTOR_OPS` <a name="VICTOR_OPS" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.VICTOR_OPS"></a>

VICTOR_OPS.

---


##### `WEBHOOK` <a name="WEBHOOK" id="@mongodbatlas-awscdk/alert-configuration.NotificationViewTypeName.WEBHOOK"></a>

WEBHOOK.

---

