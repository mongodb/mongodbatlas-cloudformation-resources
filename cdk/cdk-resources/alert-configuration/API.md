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
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrCreated">attrCreated</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::AlertConfiguration.Created`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrEnabled">attrEnabled</a></code> | <code>aws-cdk-lib.IResolvable</code> | Attribute `MongoDB::Atlas::AlertConfiguration.Enabled`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrGroupId">attrGroupId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::AlertConfiguration.GroupId`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::AlertConfiguration.Id`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrTotalCount">attrTotalCount</a></code> | <code>number</code> | Attribute `MongoDB::Atlas::AlertConfiguration.TotalCount`. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrTypeName">attrTypeName</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::AlertConfiguration.TypeName`. |
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

##### `attrCreated`<sup>Required</sup> <a name="attrCreated" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrCreated"></a>

```typescript
public readonly attrCreated: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::AlertConfiguration.Created`.

---

##### `attrEnabled`<sup>Required</sup> <a name="attrEnabled" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrEnabled"></a>

```typescript
public readonly attrEnabled: IResolvable;
```

- *Type:* aws-cdk-lib.IResolvable

Attribute `MongoDB::Atlas::AlertConfiguration.Enabled`.

---

##### `attrGroupId`<sup>Required</sup> <a name="attrGroupId" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrGroupId"></a>

```typescript
public readonly attrGroupId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::AlertConfiguration.GroupId`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::AlertConfiguration.Id`.

---

##### `attrTotalCount`<sup>Required</sup> <a name="attrTotalCount" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrTotalCount"></a>

```typescript
public readonly attrTotalCount: number;
```

- *Type:* number

Attribute `MongoDB::Atlas::AlertConfiguration.TotalCount`.

---

##### `attrTypeName`<sup>Required</sup> <a name="attrTypeName" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfiguration.property.attrTypeName"></a>

```typescript
public readonly attrTypeName: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::AlertConfiguration.TypeName`.

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

### AlertView <a name="AlertView" id="@mongodbatlas-awscdk/alert-configuration.AlertView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.AlertView.Initializer"></a>

```typescript
import { AlertView } from '@mongodbatlas-awscdk/alert-configuration'

const alertView: AlertView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.acknowledgedUntil">acknowledgedUntil</a></code> | <code>string</code> | Date and time until which this alert has been acknowledged. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.acknowledgementComment">acknowledgementComment</a></code> | <code>string</code> | Comment that a MongoDB Cloud user submitted when acknowledging the alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.acknowledgingUsername">acknowledgingUsername</a></code> | <code>string</code> | MongoDB Cloud username of the person who acknowledged the alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.alertConfigId">alertConfigId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.clusterName">clusterName</a></code> | <code>string</code> | Human-readable label that identifies the cluster to which this alert applies. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.created">created</a></code> | <code>string</code> | Date and time when MongoDB Cloud created this alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.currentValue">currentValue</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValue">CurrentValue</a></code> | Value of the metric that triggered the alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.eventTypeName">eventTypeName</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName">AlertViewEventTypeName</a></code> | Incident that triggered this alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.groupId">groupId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the project that owns this alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.hostnameAndPort">hostnameAndPort</a></code> | <code>string</code> | Hostname and port of the host to which this alert applies. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.id">id</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies this alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.lastNotified">lastNotified</a></code> | <code>string</code> | Date and time that any notifications were last sent for this alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.links">links</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.Link">Link</a>[]</code> | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.metricName">metricName</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName">AlertViewMetricName</a></code> | Human-readable label that identifies the metric against which MongoDB Cloud checks the alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.replicaSetName">replicaSetName</a></code> | <code>string</code> | Name of the replica set to which this alert applies. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.resolved">resolved</a></code> | <code>string</code> | Date and time that this alert changed to '"status" : "CLOSED"'. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.status">status</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewStatus">AlertViewStatus</a></code> | State of this alert at the time you requested its details. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.typeName">typeName</a></code> | <code>string</code> | Category in which MongoDB Cloud classifies this alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView.property.updated">updated</a></code> | <code>string</code> | Date and time when someone last updated this alert. |

---

##### `acknowledgedUntil`<sup>Optional</sup> <a name="acknowledgedUntil" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.acknowledgedUntil"></a>

```typescript
public readonly acknowledgedUntil: string;
```

- *Type:* string

Date and time until which this alert has been acknowledged.

This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if a MongoDB User previously acknowledged this alert.

- To acknowledge this alert forever, set the parameter value to 100 years in the future.

- To unacknowledge a previously acknowledged alert, set the parameter value to a date in the past.

---

##### `acknowledgementComment`<sup>Optional</sup> <a name="acknowledgementComment" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.acknowledgementComment"></a>

```typescript
public readonly acknowledgementComment: string;
```

- *Type:* string

Comment that a MongoDB Cloud user submitted when acknowledging the alert.

---

##### `acknowledgingUsername`<sup>Optional</sup> <a name="acknowledgingUsername" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.acknowledgingUsername"></a>

```typescript
public readonly acknowledgingUsername: string;
```

- *Type:* string

MongoDB Cloud username of the person who acknowledged the alert.

The response returns this parameter if a MongoDB Cloud user previously acknowledged this alert.

---

##### `alertConfigId`<sup>Optional</sup> <a name="alertConfigId" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.alertConfigId"></a>

```typescript
public readonly alertConfigId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the alert configuration that sets this alert.

---

##### `clusterName`<sup>Optional</sup> <a name="clusterName" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

Human-readable label that identifies the cluster to which this alert applies.

This resource returns this parameter for alerts of events impacting backups, replica sets, or sharded clusters.

---

##### `created`<sup>Optional</sup> <a name="created" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.created"></a>

```typescript
public readonly created: string;
```

- *Type:* string

Date and time when MongoDB Cloud created this alert.

This parameter expresses its value in the ISO 8601 timestamp format in UTC.

---

##### `currentValue`<sup>Optional</sup> <a name="currentValue" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.currentValue"></a>

```typescript
public readonly currentValue: CurrentValue;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValue">CurrentValue</a>

Value of the metric that triggered the alert.

The resource returns this parameter for alerts of events impacting hosts.

---

##### `eventTypeName`<sup>Optional</sup> <a name="eventTypeName" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.eventTypeName"></a>

```typescript
public readonly eventTypeName: AlertViewEventTypeName;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName">AlertViewEventTypeName</a>

Incident that triggered this alert.

---

##### `groupId`<sup>Optional</sup> <a name="groupId" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.groupId"></a>

```typescript
public readonly groupId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the project that owns this alert.

---

##### `hostnameAndPort`<sup>Optional</sup> <a name="hostnameAndPort" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.hostnameAndPort"></a>

```typescript
public readonly hostnameAndPort: string;
```

- *Type:* string

Hostname and port of the host to which this alert applies.

The resource returns this parameter for alerts of events impacting hosts or replica sets.

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies this alert.

---

##### `lastNotified`<sup>Optional</sup> <a name="lastNotified" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.lastNotified"></a>

```typescript
public readonly lastNotified: string;
```

- *Type:* string

Date and time that any notifications were last sent for this alert.

This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter if MongoDB Cloud has sent notifications for this alert.

---

##### `links`<sup>Optional</sup> <a name="links" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.links"></a>

```typescript
public readonly links: Link[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.Link">Link</a>[]

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.

RFC 5988 outlines these relationships.

---

##### `metricName`<sup>Optional</sup> <a name="metricName" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.metricName"></a>

```typescript
public readonly metricName: AlertViewMetricName;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName">AlertViewMetricName</a>

Human-readable label that identifies the metric against which MongoDB Cloud checks the alert.

---

##### `replicaSetName`<sup>Optional</sup> <a name="replicaSetName" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.replicaSetName"></a>

```typescript
public readonly replicaSetName: string;
```

- *Type:* string

Name of the replica set to which this alert applies.

The response returns this parameter for alerts of events impacting backups, hosts, or replica sets.

---

##### `resolved`<sup>Optional</sup> <a name="resolved" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.resolved"></a>

```typescript
public readonly resolved: string;
```

- *Type:* string

Date and time that this alert changed to '"status" : "CLOSED"'.

This parameter expresses its value in the ISO 8601 timestamp format in UTC. The resource returns this parameter once '"status" : "CLOSED"'.

---

##### `status`<sup>Optional</sup> <a name="status" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.status"></a>

```typescript
public readonly status: AlertViewStatus;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewStatus">AlertViewStatus</a>

State of this alert at the time you requested its details.

---

##### `typeName`<sup>Optional</sup> <a name="typeName" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.typeName"></a>

```typescript
public readonly typeName: string;
```

- *Type:* string

Category in which MongoDB Cloud classifies this alert.

---

##### `updated`<sup>Optional</sup> <a name="updated" id="@mongodbatlas-awscdk/alert-configuration.AlertView.property.updated"></a>

```typescript
public readonly updated: string;
```

- *Type:* string

Date and time when someone last updated this alert.

This parameter expresses its value in the ISO 8601 timestamp format in UTC.

---

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/alert-configuration.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/alert-configuration'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/alert-configuration.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/alert-configuration.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

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
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.eventTypeName">eventTypeName</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName">CfnAlertConfigurationPropsEventTypeName</a></code> | Event type that triggers an alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.links">links</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.Link">Link</a>[]</code> | List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.matchers">matchers</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.Matcher">Matcher</a>[]</code> | List of rules that determine whether MongoDB Cloud checks an object for the alert configuration. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.metricThreshold">metricThreshold</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView">MetricThresholdView</a></code> | Threshold for the metric that, when exceeded, triggers an alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.notifications">notifications</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.NotificationView">NotificationView</a>[]</code> | List that contains the targets that MongoDB Cloud sends notifications. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.profile">profile</a></code> | <code>string</code> | Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.results">results</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertView">AlertView</a>[]</code> | List of returned documents that MongoDB Cloud provides when completing this request. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.threshold">threshold</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView">IntegerThresholdView</a></code> | Limit that triggers an alert when exceeded. |

---

##### `eventTypeName`<sup>Optional</sup> <a name="eventTypeName" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.eventTypeName"></a>

```typescript
public readonly eventTypeName: CfnAlertConfigurationPropsEventTypeName;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName">CfnAlertConfigurationPropsEventTypeName</a>

Event type that triggers an alert.

---

##### `links`<sup>Optional</sup> <a name="links" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.links"></a>

```typescript
public readonly links: Link[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.Link">Link</a>[]

List of one or more Uniform Resource Locators (URLs) that point to API sub-resources, related API resources, or both.

RFC 5988 outlines these relationships.

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

##### `results`<sup>Optional</sup> <a name="results" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.results"></a>

```typescript
public readonly results: AlertView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.AlertView">AlertView</a>[]

List of returned documents that MongoDB Cloud provides when completing this request.

---

##### `threshold`<sup>Optional</sup> <a name="threshold" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationProps.property.threshold"></a>

```typescript
public readonly threshold: IntegerThresholdView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView">IntegerThresholdView</a>

Limit that triggers an alert when exceeded.

The resource returns this parameter when **eventTypeName** has not been set to 'OUTSIDE_METRIC_THRESHOLD'.

---

### CurrentValue <a name="CurrentValue" id="@mongodbatlas-awscdk/alert-configuration.CurrentValue"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.CurrentValue.Initializer"></a>

```typescript
import { CurrentValue } from '@mongodbatlas-awscdk/alert-configuration'

const currentValue: CurrentValue = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValue.property.number">number</a></code> | <code>number</code> | Amount of the **metricName** recorded at the time of the event. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValue.property.units">units</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits">CurrentValueUnits</a></code> | Element used to express the quantity in **currentValue.number**. This can be an element of time, storage capacity, and the like. This metric triggered the alert. |

---

##### `number`<sup>Optional</sup> <a name="number" id="@mongodbatlas-awscdk/alert-configuration.CurrentValue.property.number"></a>

```typescript
public readonly number: number;
```

- *Type:* number

Amount of the **metricName** recorded at the time of the event.

This value triggered the alert.

---

##### `units`<sup>Optional</sup> <a name="units" id="@mongodbatlas-awscdk/alert-configuration.CurrentValue.property.units"></a>

```typescript
public readonly units: CurrentValueUnits;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits">CurrentValueUnits</a>

Element used to express the quantity in **currentValue.number**. This can be an element of time, storage capacity, and the like. This metric triggered the alert.

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
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdView.property.units">units</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits">IntegerThresholdViewUnits</a></code> | Element used to express the quantity. |

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
public readonly units: IntegerThresholdViewUnits;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits">IntegerThresholdViewUnits</a>

Element used to express the quantity.

This can be an element of time, storage capacity, and the like.

---

### Link <a name="Link" id="@mongodbatlas-awscdk/alert-configuration.Link"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/alert-configuration.Link.Initializer"></a>

```typescript
import { Link } from '@mongodbatlas-awscdk/alert-configuration'

const link: Link = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.Link.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.Link.property.href">href</a></code> | <code>string</code> | Uniform Resource Locator (URL) that points another API resource to which this response has some relationship. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.Link.property.rel">rel</a></code> | <code>string</code> | Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource. |

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/alert-configuration.Link.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `href`<sup>Optional</sup> <a name="href" id="@mongodbatlas-awscdk/alert-configuration.Link.property.href"></a>

```typescript
public readonly href: string;
```

- *Type:* string

Uniform Resource Locator (URL) that points another API resource to which this response has some relationship.

This URL often begins with 'https://mms.mongodb.com'.

---

##### `rel`<sup>Optional</sup> <a name="rel" id="@mongodbatlas-awscdk/alert-configuration.Link.property.rel"></a>

```typescript
public readonly rel: string;
```

- *Type:* string

Uniform Resource Locator (URL) that defines the semantic relationship between this resource and another API resource.

This URL often begins with 'https://mms.mongodb.com'.

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
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.metricName">metricName</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName">MetricThresholdViewMetricName</a></code> | Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.mode">mode</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMode">MetricThresholdViewMode</a></code> | MongoDB Cloud computes the current metric value as an average. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.operator">operator</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewOperator">MetricThresholdViewOperator</a></code> | Comparison operator to apply when checking the current metric value. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.threshold">threshold</a></code> | <code>number</code> | Value of metric that, when exceeded, triggers an alert. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.units">units</a></code> | <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits">MetricThresholdViewUnits</a></code> | Element used to express the quantity. |

---

##### `metricName`<sup>Optional</sup> <a name="metricName" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdView.property.metricName"></a>

```typescript
public readonly metricName: MetricThresholdViewMetricName;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName">MetricThresholdViewMetricName</a>

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
public readonly units: MetricThresholdViewUnits;
```

- *Type:* <a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits">MetricThresholdViewUnits</a>

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

### AlertViewEventTypeName <a name="AlertViewEventTypeName" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName"></a>

Incident that triggered this alert.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.AWS_ENCRYPTION_KEY_NEEDS_ROTATION">AWS_ENCRYPTION_KEY_NEEDS_ROTATION</a></code> | AWS_ENCRYPTION_KEY_NEEDS_ROTATION. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.AZURE_ENCRYPTION_KEY_NEEDS_ROTATION">AZURE_ENCRYPTION_KEY_NEEDS_ROTATION</a></code> | AZURE_ENCRYPTION_KEY_NEEDS_ROTATION. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CLUSTER_MONGOS_IS_MISSING">CLUSTER_MONGOS_IS_MISSING</a></code> | CLUSTER_MONGOS_IS_MISSING. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_RESTORE_FAILED">CPS_RESTORE_FAILED</a></code> | CPS_RESTORE_FAILED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_RESTORE_SUCCESSFUL">CPS_RESTORE_SUCCESSFUL</a></code> | CPS_RESTORE_SUCCESSFUL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_BEHIND">CPS_SNAPSHOT_BEHIND</a></code> | CPS_SNAPSHOT_BEHIND. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED">CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED</a></code> | CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_FALLBACK_FAILED">CPS_SNAPSHOT_FALLBACK_FAILED</a></code> | CPS_SNAPSHOT_FALLBACK_FAILED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_FALLBACK_SUCCESSFUL">CPS_SNAPSHOT_FALLBACK_SUCCESSFUL</a></code> | CPS_SNAPSHOT_FALLBACK_SUCCESSFUL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_SUCCESSFUL">CPS_SNAPSHOT_SUCCESSFUL</a></code> | CPS_SNAPSHOT_SUCCESSFUL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CREDIT_CARD_ABOUT_TO_EXPIRE">CREDIT_CARD_ABOUT_TO_EXPIRE</a></code> | CREDIT_CARD_ABOUT_TO_EXPIRE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.DAILY_BILL_OVER_THRESHOLD">DAILY_BILL_OVER_THRESHOLD</a></code> | DAILY_BILL_OVER_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.GCP_ENCRYPTION_KEY_NEEDS_ROTATION">GCP_ENCRYPTION_KEY_NEEDS_ROTATION</a></code> | GCP_ENCRYPTION_KEY_NEEDS_ROTATION. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.HOST_DOWN">HOST_DOWN</a></code> | HOST_DOWN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.JOINED_GROUP">JOINED_GROUP</a></code> | JOINED_GROUP. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK">NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK</a></code> | NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK">NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK</a></code> | NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.NO_PRIMARY">NO_PRIMARY</a></code> | NO_PRIMARY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.OUTSIDE_METRIC_THRESHOLD">OUTSIDE_METRIC_THRESHOLD</a></code> | OUTSIDE_METRIC_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.OUTSIDE_SERVERLESS_METRIC_THRESHOLD">OUTSIDE_SERVERLESS_METRIC_THRESHOLD</a></code> | OUTSIDE_SERVERLESS_METRIC_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.PENDING_INVOICE_OVER_THRESHOLD">PENDING_INVOICE_OVER_THRESHOLD</a></code> | PENDING_INVOICE_OVER_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.PRIMARY_ELECTED">PRIMARY_ELECTED</a></code> | PRIMARY_ELECTED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.REMOVED_FROM_GROUP">REMOVED_FROM_GROUP</a></code> | REMOVED_FROM_GROUP. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.REPLICATION_OPLOG_WINDOW_RUNNING_OUT">REPLICATION_OPLOG_WINDOW_RUNNING_OUT</a></code> | REPLICATION_OPLOG_WINDOW_RUNNING_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.TOO_MANY_ELECTIONS">TOO_MANY_ELECTIONS</a></code> | TOO_MANY_ELECTIONS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.USERS_WITHOUT_MULTIFACTOR_AUTH">USERS_WITHOUT_MULTIFACTOR_AUTH</a></code> | USERS_WITHOUT_MULTIFACTOR_AUTH. |

---

##### `AWS_ENCRYPTION_KEY_NEEDS_ROTATION` <a name="AWS_ENCRYPTION_KEY_NEEDS_ROTATION" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.AWS_ENCRYPTION_KEY_NEEDS_ROTATION"></a>

AWS_ENCRYPTION_KEY_NEEDS_ROTATION.

---


##### `AZURE_ENCRYPTION_KEY_NEEDS_ROTATION` <a name="AZURE_ENCRYPTION_KEY_NEEDS_ROTATION" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.AZURE_ENCRYPTION_KEY_NEEDS_ROTATION"></a>

AZURE_ENCRYPTION_KEY_NEEDS_ROTATION.

---


##### `CLUSTER_MONGOS_IS_MISSING` <a name="CLUSTER_MONGOS_IS_MISSING" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CLUSTER_MONGOS_IS_MISSING"></a>

CLUSTER_MONGOS_IS_MISSING.

---


##### `CPS_RESTORE_FAILED` <a name="CPS_RESTORE_FAILED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_RESTORE_FAILED"></a>

CPS_RESTORE_FAILED.

---


##### `CPS_RESTORE_SUCCESSFUL` <a name="CPS_RESTORE_SUCCESSFUL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_RESTORE_SUCCESSFUL"></a>

CPS_RESTORE_SUCCESSFUL.

---


##### `CPS_SNAPSHOT_BEHIND` <a name="CPS_SNAPSHOT_BEHIND" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_BEHIND"></a>

CPS_SNAPSHOT_BEHIND.

---


##### `CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED` <a name="CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED"></a>

CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED.

---


##### `CPS_SNAPSHOT_FALLBACK_FAILED` <a name="CPS_SNAPSHOT_FALLBACK_FAILED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_FALLBACK_FAILED"></a>

CPS_SNAPSHOT_FALLBACK_FAILED.

---


##### `CPS_SNAPSHOT_FALLBACK_SUCCESSFUL` <a name="CPS_SNAPSHOT_FALLBACK_SUCCESSFUL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_FALLBACK_SUCCESSFUL"></a>

CPS_SNAPSHOT_FALLBACK_SUCCESSFUL.

---


##### `CPS_SNAPSHOT_SUCCESSFUL` <a name="CPS_SNAPSHOT_SUCCESSFUL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CPS_SNAPSHOT_SUCCESSFUL"></a>

CPS_SNAPSHOT_SUCCESSFUL.

---


##### `CREDIT_CARD_ABOUT_TO_EXPIRE` <a name="CREDIT_CARD_ABOUT_TO_EXPIRE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.CREDIT_CARD_ABOUT_TO_EXPIRE"></a>

CREDIT_CARD_ABOUT_TO_EXPIRE.

---


##### `DAILY_BILL_OVER_THRESHOLD` <a name="DAILY_BILL_OVER_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.DAILY_BILL_OVER_THRESHOLD"></a>

DAILY_BILL_OVER_THRESHOLD.

---


##### `GCP_ENCRYPTION_KEY_NEEDS_ROTATION` <a name="GCP_ENCRYPTION_KEY_NEEDS_ROTATION" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.GCP_ENCRYPTION_KEY_NEEDS_ROTATION"></a>

GCP_ENCRYPTION_KEY_NEEDS_ROTATION.

---


##### `HOST_DOWN` <a name="HOST_DOWN" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.HOST_DOWN"></a>

HOST_DOWN.

---


##### `JOINED_GROUP` <a name="JOINED_GROUP" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.JOINED_GROUP"></a>

JOINED_GROUP.

---


##### `NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK` <a name="NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK"></a>

NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK.

---


##### `NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK` <a name="NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK"></a>

NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK.

---


##### `NO_PRIMARY` <a name="NO_PRIMARY" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.NO_PRIMARY"></a>

NO_PRIMARY.

---


##### `OUTSIDE_METRIC_THRESHOLD` <a name="OUTSIDE_METRIC_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.OUTSIDE_METRIC_THRESHOLD"></a>

OUTSIDE_METRIC_THRESHOLD.

---


##### `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` <a name="OUTSIDE_SERVERLESS_METRIC_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.OUTSIDE_SERVERLESS_METRIC_THRESHOLD"></a>

OUTSIDE_SERVERLESS_METRIC_THRESHOLD.

---


##### `PENDING_INVOICE_OVER_THRESHOLD` <a name="PENDING_INVOICE_OVER_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.PENDING_INVOICE_OVER_THRESHOLD"></a>

PENDING_INVOICE_OVER_THRESHOLD.

---


##### `PRIMARY_ELECTED` <a name="PRIMARY_ELECTED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.PRIMARY_ELECTED"></a>

PRIMARY_ELECTED.

---


##### `REMOVED_FROM_GROUP` <a name="REMOVED_FROM_GROUP" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.REMOVED_FROM_GROUP"></a>

REMOVED_FROM_GROUP.

---


##### `REPLICATION_OPLOG_WINDOW_RUNNING_OUT` <a name="REPLICATION_OPLOG_WINDOW_RUNNING_OUT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.REPLICATION_OPLOG_WINDOW_RUNNING_OUT"></a>

REPLICATION_OPLOG_WINDOW_RUNNING_OUT.

---


##### `TOO_MANY_ELECTIONS` <a name="TOO_MANY_ELECTIONS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.TOO_MANY_ELECTIONS"></a>

TOO_MANY_ELECTIONS.

---


##### `USERS_WITHOUT_MULTIFACTOR_AUTH` <a name="USERS_WITHOUT_MULTIFACTOR_AUTH" id="@mongodbatlas-awscdk/alert-configuration.AlertViewEventTypeName.USERS_WITHOUT_MULTIFACTOR_AUTH"></a>

USERS_WITHOUT_MULTIFACTOR_AUTH.

---


### AlertViewMetricName <a name="AlertViewMetricName" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName"></a>

Human-readable label that identifies the metric against which MongoDB Cloud checks the alert.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.ASSERT_MSG">ASSERT_MSG</a></code> | ASSERT_MSG. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.ASSERT_REGULAR">ASSERT_REGULAR</a></code> | ASSERT_REGULAR. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.ASSERT_USER">ASSERT_USER</a></code> | ASSERT_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.ASSERT_WARNING">ASSERT_WARNING</a></code> | ASSERT_WARNING. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.AVG_COMMAND_EXECUTION_TIME">AVG_COMMAND_EXECUTION_TIME</a></code> | AVG_COMMAND_EXECUTION_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.AVG_READ_EXECUTION_TIME">AVG_READ_EXECUTION_TIME</a></code> | AVG_READ_EXECUTION_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.AVG_WRITE_EXECUTION_TIME">AVG_WRITE_EXECUTION_TIME</a></code> | AVG_WRITE_EXECUTION_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CACHE_BYTES_READ_INTO">CACHE_BYTES_READ_INTO</a></code> | CACHE_BYTES_READ_INTO. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CACHE_BYTES_WRITTEN_FROM">CACHE_BYTES_WRITTEN_FROM</a></code> | CACHE_BYTES_WRITTEN_FROM. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CACHE_DIRTY_BYTES">CACHE_DIRTY_BYTES</a></code> | CACHE_DIRTY_BYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CACHE_USED_BYTES">CACHE_USED_BYTES</a></code> | CACHE_USED_BYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.COMPUTED_MEMORY">COMPUTED_MEMORY</a></code> | COMPUTED_MEMORY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CONNECTIONS">CONNECTIONS</a></code> | CONNECTIONS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CONNECTIONS_PERCENT">CONNECTIONS_PERCENT</a></code> | CONNECTIONS_PERCENT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CURSORS_TOTAL_OPEN">CURSORS_TOTAL_OPEN</a></code> | CURSORS_TOTAL_OPEN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CURSORS_TOTAL_TIMED_OUT">CURSORS_TOTAL_TIMED_OUT</a></code> | CURSORS_TOTAL_TIMED_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DB_DATA_SIZE_TOTAL">DB_DATA_SIZE_TOTAL</a></code> | DB_DATA_SIZE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DB_INDEX_SIZE_TOTAL">DB_INDEX_SIZE_TOTAL</a></code> | DB_INDEX_SIZE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DB_STORAGE_TOTAL">DB_STORAGE_TOTAL</a></code> | DB_STORAGE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_SPACE_USED_DATA">DISK_PARTITION_SPACE_USED_DATA</a></code> | DISK_PARTITION_SPACE_USED_DATA. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_SPACE_USED_INDEX">DISK_PARTITION_SPACE_USED_INDEX</a></code> | DISK_PARTITION_SPACE_USED_INDEX. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_SPACE_USED_JOURNAL">DISK_PARTITION_SPACE_USED_JOURNAL</a></code> | DISK_PARTITION_SPACE_USED_JOURNAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_UTILIZATION_DATA">DISK_PARTITION_UTILIZATION_DATA</a></code> | DISK_PARTITION_UTILIZATION_DATA. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_UTILIZATION_INDEX">DISK_PARTITION_UTILIZATION_INDEX</a></code> | DISK_PARTITION_UTILIZATION_INDEX. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_UTILIZATION_JOURNAL">DISK_PARTITION_UTILIZATION_JOURNAL</a></code> | DISK_PARTITION_UTILIZATION_JOURNAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DOCUMENT_DELETED">DOCUMENT_DELETED</a></code> | DOCUMENT_DELETED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DOCUMENT_INSERTED">DOCUMENT_INSERTED</a></code> | DOCUMENT_INSERTED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DOCUMENT_RETURNED">DOCUMENT_RETURNED</a></code> | DOCUMENT_RETURNED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DOCUMENT_UPDATED">DOCUMENT_UPDATED</a></code> | DOCUMENT_UPDATED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.EXTRA_INFO_PAGE_FAULTS">EXTRA_INFO_PAGE_FAULTS</a></code> | EXTRA_INFO_PAGE_FAULTS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_MEMORY_RESIDENT">FTS_MEMORY_RESIDENT</a></code> | FTS_MEMORY_RESIDENT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_MEMORY_SHARED">FTS_MEMORY_SHARED</a></code> | FTS_MEMORY_SHARED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_MEMORY_VIRTUAL">FTS_MEMORY_VIRTUAL</a></code> | FTS_MEMORY_VIRTUAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_PROCESS_CPU_KERNEL">FTS_PROCESS_CPU_KERNEL</a></code> | FTS_PROCESS_CPU_KERNEL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_PROCESS_CPU_USER">FTS_PROCESS_CPU_USER</a></code> | FTS_PROCESS_CPU_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_PROCESS_DISK">FTS_PROCESS_DISK</a></code> | FTS_PROCESS_DISK. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_READERS">GLOBAL_LOCK_CURRENT_QUEUE_READERS</a></code> | GLOBAL_LOCK_CURRENT_QUEUE_READERS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_TOTAL">GLOBAL_LOCK_CURRENT_QUEUE_TOTAL</a></code> | GLOBAL_LOCK_CURRENT_QUEUE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_WRITERS">GLOBAL_LOCK_CURRENT_QUEUE_WRITERS</a></code> | GLOBAL_LOCK_CURRENT_QUEUE_WRITERS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.LOGICAL_SIZE">LOGICAL_SIZE</a></code> | LOGICAL_SIZE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.MEMORY_RESIDENT">MEMORY_RESIDENT</a></code> | MEMORY_RESIDENT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.MEMORY_VIRTUAL">MEMORY_VIRTUAL</a></code> | MEMORY_VIRTUAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NETWORK_BYTES_IN">NETWORK_BYTES_IN</a></code> | NETWORK_BYTES_IN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NETWORK_BYTES_OUT">NETWORK_BYTES_OUT</a></code> | NETWORK_BYTES_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NETWORK_NUM_REQUESTS">NETWORK_NUM_REQUESTS</a></code> | NETWORK_NUM_REQUESTS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NORMALIZED_FTS_PROCESS_CPU_KERNEL">NORMALIZED_FTS_PROCESS_CPU_KERNEL</a></code> | NORMALIZED_FTS_PROCESS_CPU_KERNEL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NORMALIZED_FTS_PROCESS_CPU_USER">NORMALIZED_FTS_PROCESS_CPU_USER</a></code> | NORMALIZED_FTS_PROCESS_CPU_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NORMALIZED_SYSTEM_CPU_STEAL">NORMALIZED_SYSTEM_CPU_STEAL</a></code> | NORMALIZED_SYSTEM_CPU_STEAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NORMALIZED_SYSTEM_CPU_USER">NORMALIZED_SYSTEM_CPU_USER</a></code> | NORMALIZED_SYSTEM_CPU_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_CMD">OPCOUNTER_CMD</a></code> | OPCOUNTER_CMD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_DELETE">OPCOUNTER_DELETE</a></code> | OPCOUNTER_DELETE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_GETMORE">OPCOUNTER_GETMORE</a></code> | OPCOUNTER_GETMORE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_INSERT">OPCOUNTER_INSERT</a></code> | OPCOUNTER_INSERT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_QUERY">OPCOUNTER_QUERY</a></code> | OPCOUNTER_QUERY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_REPL_CMD">OPCOUNTER_REPL_CMD</a></code> | OPCOUNTER_REPL_CMD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_REPL_DELETE">OPCOUNTER_REPL_DELETE</a></code> | OPCOUNTER_REPL_DELETE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_REPL_INSERT">OPCOUNTER_REPL_INSERT</a></code> | OPCOUNTER_REPL_INSERT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_REPL_UPDATE">OPCOUNTER_REPL_UPDATE</a></code> | OPCOUNTER_REPL_UPDATE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_UPDATE">OPCOUNTER_UPDATE</a></code> | OPCOUNTER_UPDATE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPERATIONS_SCAN_AND_ORDER">OPERATIONS_SCAN_AND_ORDER</a></code> | OPERATIONS_SCAN_AND_ORDER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPLOG_MASTER_LAG_TIME_DIFF">OPLOG_MASTER_LAG_TIME_DIFF</a></code> | OPLOG_MASTER_LAG_TIME_DIFF. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPLOG_MASTER_TIME">OPLOG_MASTER_TIME</a></code> | OPLOG_MASTER_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPLOG_RATE_GB_PER_HOUR">OPLOG_RATE_GB_PER_HOUR</a></code> | OPLOG_RATE_GB_PER_HOUR. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPLOG_SLAVE_LAG_MASTER_TIME">OPLOG_SLAVE_LAG_MASTER_TIME</a></code> | OPLOG_SLAVE_LAG_MASTER_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.QUERY_EXECUTOR_SCANNED">QUERY_EXECUTOR_SCANNED</a></code> | QUERY_EXECUTOR_SCANNED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.QUERY_EXECUTOR_SCANNED_OBJECTS">QUERY_EXECUTOR_SCANNED_OBJECTS</a></code> | QUERY_EXECUTOR_SCANNED_OBJECTS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED">QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED</a></code> | QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.QUERY_TARGETING_SCANNED_PER_RETURNED">QUERY_TARGETING_SCANNED_PER_RETURNED</a></code> | QUERY_TARGETING_SCANNED_PER_RETURNED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.RESTARTS_IN_LAST_HOUR">RESTARTS_IN_LAST_HOUR</a></code> | RESTARTS_IN_LAST_HOUR. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_CONNECTIONS">SERVERLESS_CONNECTIONS</a></code> | SERVERLESS_CONNECTIONS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_CONNECTIONS_PERCENT">SERVERLESS_CONNECTIONS_PERCENT</a></code> | SERVERLESS_CONNECTIONS_PERCENT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_DATA_SIZE_TOTAL">SERVERLESS_DATA_SIZE_TOTAL</a></code> | SERVERLESS_DATA_SIZE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_NETWORK_BYTES_IN">SERVERLESS_NETWORK_BYTES_IN</a></code> | SERVERLESS_NETWORK_BYTES_IN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_NETWORK_BYTES_OUT">SERVERLESS_NETWORK_BYTES_OUT</a></code> | SERVERLESS_NETWORK_BYTES_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_NETWORK_NUM_REQUESTS">SERVERLESS_NETWORK_NUM_REQUESTS</a></code> | SERVERLESS_NETWORK_NUM_REQUESTS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_CMD">SERVERLESS_OPCOUNTER_CMD</a></code> | SERVERLESS_OPCOUNTER_CMD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_DELETE">SERVERLESS_OPCOUNTER_DELETE</a></code> | SERVERLESS_OPCOUNTER_DELETE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_GETMORE">SERVERLESS_OPCOUNTER_GETMORE</a></code> | SERVERLESS_OPCOUNTER_GETMORE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_INSERT">SERVERLESS_OPCOUNTER_INSERT</a></code> | SERVERLESS_OPCOUNTER_INSERT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_QUERY">SERVERLESS_OPCOUNTER_QUERY</a></code> | SERVERLESS_OPCOUNTER_QUERY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_UPDATE">SERVERLESS_OPCOUNTER_UPDATE</a></code> | SERVERLESS_OPCOUNTER_UPDATE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_TOTAL_READ_UNITS">SERVERLESS_TOTAL_READ_UNITS</a></code> | SERVERLESS_TOTAL_READ_UNITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_TOTAL_WRITE_UNITS">SERVERLESS_TOTAL_WRITE_UNITS</a></code> | SERVERLESS_TOTAL_WRITE_UNITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.TICKETS_AVAILABLE_READS">TICKETS_AVAILABLE_READS</a></code> | TICKETS_AVAILABLE_READS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.TICKETS_AVAILABLE_WRITES">TICKETS_AVAILABLE_WRITES</a></code> | TICKETS_AVAILABLE_WRITES. |

---

##### `ASSERT_MSG` <a name="ASSERT_MSG" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.ASSERT_MSG"></a>

ASSERT_MSG.

---


##### `ASSERT_REGULAR` <a name="ASSERT_REGULAR" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.ASSERT_REGULAR"></a>

ASSERT_REGULAR.

---


##### `ASSERT_USER` <a name="ASSERT_USER" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.ASSERT_USER"></a>

ASSERT_USER.

---


##### `ASSERT_WARNING` <a name="ASSERT_WARNING" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.ASSERT_WARNING"></a>

ASSERT_WARNING.

---


##### `AVG_COMMAND_EXECUTION_TIME` <a name="AVG_COMMAND_EXECUTION_TIME" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.AVG_COMMAND_EXECUTION_TIME"></a>

AVG_COMMAND_EXECUTION_TIME.

---


##### `AVG_READ_EXECUTION_TIME` <a name="AVG_READ_EXECUTION_TIME" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.AVG_READ_EXECUTION_TIME"></a>

AVG_READ_EXECUTION_TIME.

---


##### `AVG_WRITE_EXECUTION_TIME` <a name="AVG_WRITE_EXECUTION_TIME" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.AVG_WRITE_EXECUTION_TIME"></a>

AVG_WRITE_EXECUTION_TIME.

---


##### `CACHE_BYTES_READ_INTO` <a name="CACHE_BYTES_READ_INTO" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CACHE_BYTES_READ_INTO"></a>

CACHE_BYTES_READ_INTO.

---


##### `CACHE_BYTES_WRITTEN_FROM` <a name="CACHE_BYTES_WRITTEN_FROM" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CACHE_BYTES_WRITTEN_FROM"></a>

CACHE_BYTES_WRITTEN_FROM.

---


##### `CACHE_DIRTY_BYTES` <a name="CACHE_DIRTY_BYTES" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CACHE_DIRTY_BYTES"></a>

CACHE_DIRTY_BYTES.

---


##### `CACHE_USED_BYTES` <a name="CACHE_USED_BYTES" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CACHE_USED_BYTES"></a>

CACHE_USED_BYTES.

---


##### `COMPUTED_MEMORY` <a name="COMPUTED_MEMORY" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.COMPUTED_MEMORY"></a>

COMPUTED_MEMORY.

---


##### `CONNECTIONS` <a name="CONNECTIONS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CONNECTIONS"></a>

CONNECTIONS.

---


##### `CONNECTIONS_PERCENT` <a name="CONNECTIONS_PERCENT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CONNECTIONS_PERCENT"></a>

CONNECTIONS_PERCENT.

---


##### `CURSORS_TOTAL_OPEN` <a name="CURSORS_TOTAL_OPEN" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CURSORS_TOTAL_OPEN"></a>

CURSORS_TOTAL_OPEN.

---


##### `CURSORS_TOTAL_TIMED_OUT` <a name="CURSORS_TOTAL_TIMED_OUT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.CURSORS_TOTAL_TIMED_OUT"></a>

CURSORS_TOTAL_TIMED_OUT.

---


##### `DB_DATA_SIZE_TOTAL` <a name="DB_DATA_SIZE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DB_DATA_SIZE_TOTAL"></a>

DB_DATA_SIZE_TOTAL.

---


##### `DB_INDEX_SIZE_TOTAL` <a name="DB_INDEX_SIZE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DB_INDEX_SIZE_TOTAL"></a>

DB_INDEX_SIZE_TOTAL.

---


##### `DB_STORAGE_TOTAL` <a name="DB_STORAGE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DB_STORAGE_TOTAL"></a>

DB_STORAGE_TOTAL.

---


##### `DISK_PARTITION_SPACE_USED_DATA` <a name="DISK_PARTITION_SPACE_USED_DATA" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_SPACE_USED_DATA"></a>

DISK_PARTITION_SPACE_USED_DATA.

---


##### `DISK_PARTITION_SPACE_USED_INDEX` <a name="DISK_PARTITION_SPACE_USED_INDEX" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_SPACE_USED_INDEX"></a>

DISK_PARTITION_SPACE_USED_INDEX.

---


##### `DISK_PARTITION_SPACE_USED_JOURNAL` <a name="DISK_PARTITION_SPACE_USED_JOURNAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_SPACE_USED_JOURNAL"></a>

DISK_PARTITION_SPACE_USED_JOURNAL.

---


##### `DISK_PARTITION_UTILIZATION_DATA` <a name="DISK_PARTITION_UTILIZATION_DATA" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_UTILIZATION_DATA"></a>

DISK_PARTITION_UTILIZATION_DATA.

---


##### `DISK_PARTITION_UTILIZATION_INDEX` <a name="DISK_PARTITION_UTILIZATION_INDEX" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_UTILIZATION_INDEX"></a>

DISK_PARTITION_UTILIZATION_INDEX.

---


##### `DISK_PARTITION_UTILIZATION_JOURNAL` <a name="DISK_PARTITION_UTILIZATION_JOURNAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DISK_PARTITION_UTILIZATION_JOURNAL"></a>

DISK_PARTITION_UTILIZATION_JOURNAL.

---


##### `DOCUMENT_DELETED` <a name="DOCUMENT_DELETED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DOCUMENT_DELETED"></a>

DOCUMENT_DELETED.

---


##### `DOCUMENT_INSERTED` <a name="DOCUMENT_INSERTED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DOCUMENT_INSERTED"></a>

DOCUMENT_INSERTED.

---


##### `DOCUMENT_RETURNED` <a name="DOCUMENT_RETURNED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DOCUMENT_RETURNED"></a>

DOCUMENT_RETURNED.

---


##### `DOCUMENT_UPDATED` <a name="DOCUMENT_UPDATED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.DOCUMENT_UPDATED"></a>

DOCUMENT_UPDATED.

---


##### `EXTRA_INFO_PAGE_FAULTS` <a name="EXTRA_INFO_PAGE_FAULTS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.EXTRA_INFO_PAGE_FAULTS"></a>

EXTRA_INFO_PAGE_FAULTS.

---


##### `FTS_MEMORY_RESIDENT` <a name="FTS_MEMORY_RESIDENT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_MEMORY_RESIDENT"></a>

FTS_MEMORY_RESIDENT.

---


##### `FTS_MEMORY_SHARED` <a name="FTS_MEMORY_SHARED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_MEMORY_SHARED"></a>

FTS_MEMORY_SHARED.

---


##### `FTS_MEMORY_VIRTUAL` <a name="FTS_MEMORY_VIRTUAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_MEMORY_VIRTUAL"></a>

FTS_MEMORY_VIRTUAL.

---


##### `FTS_PROCESS_CPU_KERNEL` <a name="FTS_PROCESS_CPU_KERNEL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_PROCESS_CPU_KERNEL"></a>

FTS_PROCESS_CPU_KERNEL.

---


##### `FTS_PROCESS_CPU_USER` <a name="FTS_PROCESS_CPU_USER" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_PROCESS_CPU_USER"></a>

FTS_PROCESS_CPU_USER.

---


##### `FTS_PROCESS_DISK` <a name="FTS_PROCESS_DISK" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.FTS_PROCESS_DISK"></a>

FTS_PROCESS_DISK.

---


##### `GLOBAL_LOCK_CURRENT_QUEUE_READERS` <a name="GLOBAL_LOCK_CURRENT_QUEUE_READERS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_READERS"></a>

GLOBAL_LOCK_CURRENT_QUEUE_READERS.

---


##### `GLOBAL_LOCK_CURRENT_QUEUE_TOTAL` <a name="GLOBAL_LOCK_CURRENT_QUEUE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_TOTAL"></a>

GLOBAL_LOCK_CURRENT_QUEUE_TOTAL.

---


##### `GLOBAL_LOCK_CURRENT_QUEUE_WRITERS` <a name="GLOBAL_LOCK_CURRENT_QUEUE_WRITERS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_WRITERS"></a>

GLOBAL_LOCK_CURRENT_QUEUE_WRITERS.

---


##### `LOGICAL_SIZE` <a name="LOGICAL_SIZE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.LOGICAL_SIZE"></a>

LOGICAL_SIZE.

---


##### `MEMORY_RESIDENT` <a name="MEMORY_RESIDENT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.MEMORY_RESIDENT"></a>

MEMORY_RESIDENT.

---


##### `MEMORY_VIRTUAL` <a name="MEMORY_VIRTUAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.MEMORY_VIRTUAL"></a>

MEMORY_VIRTUAL.

---


##### `NETWORK_BYTES_IN` <a name="NETWORK_BYTES_IN" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NETWORK_BYTES_IN"></a>

NETWORK_BYTES_IN.

---


##### `NETWORK_BYTES_OUT` <a name="NETWORK_BYTES_OUT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NETWORK_BYTES_OUT"></a>

NETWORK_BYTES_OUT.

---


##### `NETWORK_NUM_REQUESTS` <a name="NETWORK_NUM_REQUESTS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NETWORK_NUM_REQUESTS"></a>

NETWORK_NUM_REQUESTS.

---


##### `NORMALIZED_FTS_PROCESS_CPU_KERNEL` <a name="NORMALIZED_FTS_PROCESS_CPU_KERNEL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NORMALIZED_FTS_PROCESS_CPU_KERNEL"></a>

NORMALIZED_FTS_PROCESS_CPU_KERNEL.

---


##### `NORMALIZED_FTS_PROCESS_CPU_USER` <a name="NORMALIZED_FTS_PROCESS_CPU_USER" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NORMALIZED_FTS_PROCESS_CPU_USER"></a>

NORMALIZED_FTS_PROCESS_CPU_USER.

---


##### `NORMALIZED_SYSTEM_CPU_STEAL` <a name="NORMALIZED_SYSTEM_CPU_STEAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NORMALIZED_SYSTEM_CPU_STEAL"></a>

NORMALIZED_SYSTEM_CPU_STEAL.

---


##### `NORMALIZED_SYSTEM_CPU_USER` <a name="NORMALIZED_SYSTEM_CPU_USER" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.NORMALIZED_SYSTEM_CPU_USER"></a>

NORMALIZED_SYSTEM_CPU_USER.

---


##### `OPCOUNTER_CMD` <a name="OPCOUNTER_CMD" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_CMD"></a>

OPCOUNTER_CMD.

---


##### `OPCOUNTER_DELETE` <a name="OPCOUNTER_DELETE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_DELETE"></a>

OPCOUNTER_DELETE.

---


##### `OPCOUNTER_GETMORE` <a name="OPCOUNTER_GETMORE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_GETMORE"></a>

OPCOUNTER_GETMORE.

---


##### `OPCOUNTER_INSERT` <a name="OPCOUNTER_INSERT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_INSERT"></a>

OPCOUNTER_INSERT.

---


##### `OPCOUNTER_QUERY` <a name="OPCOUNTER_QUERY" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_QUERY"></a>

OPCOUNTER_QUERY.

---


##### `OPCOUNTER_REPL_CMD` <a name="OPCOUNTER_REPL_CMD" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_REPL_CMD"></a>

OPCOUNTER_REPL_CMD.

---


##### `OPCOUNTER_REPL_DELETE` <a name="OPCOUNTER_REPL_DELETE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_REPL_DELETE"></a>

OPCOUNTER_REPL_DELETE.

---


##### `OPCOUNTER_REPL_INSERT` <a name="OPCOUNTER_REPL_INSERT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_REPL_INSERT"></a>

OPCOUNTER_REPL_INSERT.

---


##### `OPCOUNTER_REPL_UPDATE` <a name="OPCOUNTER_REPL_UPDATE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_REPL_UPDATE"></a>

OPCOUNTER_REPL_UPDATE.

---


##### `OPCOUNTER_UPDATE` <a name="OPCOUNTER_UPDATE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPCOUNTER_UPDATE"></a>

OPCOUNTER_UPDATE.

---


##### `OPERATIONS_SCAN_AND_ORDER` <a name="OPERATIONS_SCAN_AND_ORDER" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPERATIONS_SCAN_AND_ORDER"></a>

OPERATIONS_SCAN_AND_ORDER.

---


##### `OPLOG_MASTER_LAG_TIME_DIFF` <a name="OPLOG_MASTER_LAG_TIME_DIFF" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPLOG_MASTER_LAG_TIME_DIFF"></a>

OPLOG_MASTER_LAG_TIME_DIFF.

---


##### `OPLOG_MASTER_TIME` <a name="OPLOG_MASTER_TIME" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPLOG_MASTER_TIME"></a>

OPLOG_MASTER_TIME.

---


##### `OPLOG_RATE_GB_PER_HOUR` <a name="OPLOG_RATE_GB_PER_HOUR" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPLOG_RATE_GB_PER_HOUR"></a>

OPLOG_RATE_GB_PER_HOUR.

---


##### `OPLOG_SLAVE_LAG_MASTER_TIME` <a name="OPLOG_SLAVE_LAG_MASTER_TIME" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.OPLOG_SLAVE_LAG_MASTER_TIME"></a>

OPLOG_SLAVE_LAG_MASTER_TIME.

---


##### `QUERY_EXECUTOR_SCANNED` <a name="QUERY_EXECUTOR_SCANNED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.QUERY_EXECUTOR_SCANNED"></a>

QUERY_EXECUTOR_SCANNED.

---


##### `QUERY_EXECUTOR_SCANNED_OBJECTS` <a name="QUERY_EXECUTOR_SCANNED_OBJECTS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.QUERY_EXECUTOR_SCANNED_OBJECTS"></a>

QUERY_EXECUTOR_SCANNED_OBJECTS.

---


##### `QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED` <a name="QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED"></a>

QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED.

---


##### `QUERY_TARGETING_SCANNED_PER_RETURNED` <a name="QUERY_TARGETING_SCANNED_PER_RETURNED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.QUERY_TARGETING_SCANNED_PER_RETURNED"></a>

QUERY_TARGETING_SCANNED_PER_RETURNED.

---


##### `RESTARTS_IN_LAST_HOUR` <a name="RESTARTS_IN_LAST_HOUR" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.RESTARTS_IN_LAST_HOUR"></a>

RESTARTS_IN_LAST_HOUR.

---


##### `SERVERLESS_CONNECTIONS` <a name="SERVERLESS_CONNECTIONS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_CONNECTIONS"></a>

SERVERLESS_CONNECTIONS.

---


##### `SERVERLESS_CONNECTIONS_PERCENT` <a name="SERVERLESS_CONNECTIONS_PERCENT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_CONNECTIONS_PERCENT"></a>

SERVERLESS_CONNECTIONS_PERCENT.

---


##### `SERVERLESS_DATA_SIZE_TOTAL` <a name="SERVERLESS_DATA_SIZE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_DATA_SIZE_TOTAL"></a>

SERVERLESS_DATA_SIZE_TOTAL.

---


##### `SERVERLESS_NETWORK_BYTES_IN` <a name="SERVERLESS_NETWORK_BYTES_IN" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_NETWORK_BYTES_IN"></a>

SERVERLESS_NETWORK_BYTES_IN.

---


##### `SERVERLESS_NETWORK_BYTES_OUT` <a name="SERVERLESS_NETWORK_BYTES_OUT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_NETWORK_BYTES_OUT"></a>

SERVERLESS_NETWORK_BYTES_OUT.

---


##### `SERVERLESS_NETWORK_NUM_REQUESTS` <a name="SERVERLESS_NETWORK_NUM_REQUESTS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_NETWORK_NUM_REQUESTS"></a>

SERVERLESS_NETWORK_NUM_REQUESTS.

---


##### `SERVERLESS_OPCOUNTER_CMD` <a name="SERVERLESS_OPCOUNTER_CMD" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_CMD"></a>

SERVERLESS_OPCOUNTER_CMD.

---


##### `SERVERLESS_OPCOUNTER_DELETE` <a name="SERVERLESS_OPCOUNTER_DELETE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_DELETE"></a>

SERVERLESS_OPCOUNTER_DELETE.

---


##### `SERVERLESS_OPCOUNTER_GETMORE` <a name="SERVERLESS_OPCOUNTER_GETMORE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_GETMORE"></a>

SERVERLESS_OPCOUNTER_GETMORE.

---


##### `SERVERLESS_OPCOUNTER_INSERT` <a name="SERVERLESS_OPCOUNTER_INSERT" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_INSERT"></a>

SERVERLESS_OPCOUNTER_INSERT.

---


##### `SERVERLESS_OPCOUNTER_QUERY` <a name="SERVERLESS_OPCOUNTER_QUERY" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_QUERY"></a>

SERVERLESS_OPCOUNTER_QUERY.

---


##### `SERVERLESS_OPCOUNTER_UPDATE` <a name="SERVERLESS_OPCOUNTER_UPDATE" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_OPCOUNTER_UPDATE"></a>

SERVERLESS_OPCOUNTER_UPDATE.

---


##### `SERVERLESS_TOTAL_READ_UNITS` <a name="SERVERLESS_TOTAL_READ_UNITS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_TOTAL_READ_UNITS"></a>

SERVERLESS_TOTAL_READ_UNITS.

---


##### `SERVERLESS_TOTAL_WRITE_UNITS` <a name="SERVERLESS_TOTAL_WRITE_UNITS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.SERVERLESS_TOTAL_WRITE_UNITS"></a>

SERVERLESS_TOTAL_WRITE_UNITS.

---


##### `TICKETS_AVAILABLE_READS` <a name="TICKETS_AVAILABLE_READS" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.TICKETS_AVAILABLE_READS"></a>

TICKETS_AVAILABLE_READS.

---


##### `TICKETS_AVAILABLE_WRITES` <a name="TICKETS_AVAILABLE_WRITES" id="@mongodbatlas-awscdk/alert-configuration.AlertViewMetricName.TICKETS_AVAILABLE_WRITES"></a>

TICKETS_AVAILABLE_WRITES.

---


### AlertViewStatus <a name="AlertViewStatus" id="@mongodbatlas-awscdk/alert-configuration.AlertViewStatus"></a>

State of this alert at the time you requested its details.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewStatus.CANCELLED">CANCELLED</a></code> | CANCELLED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewStatus.CLOSED">CLOSED</a></code> | CLOSED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewStatus.OPEN">OPEN</a></code> | OPEN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.AlertViewStatus.TRACKING">TRACKING</a></code> | TRACKING. |

---

##### `CANCELLED` <a name="CANCELLED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewStatus.CANCELLED"></a>

CANCELLED.

---


##### `CLOSED` <a name="CLOSED" id="@mongodbatlas-awscdk/alert-configuration.AlertViewStatus.CLOSED"></a>

CLOSED.

---


##### `OPEN` <a name="OPEN" id="@mongodbatlas-awscdk/alert-configuration.AlertViewStatus.OPEN"></a>

OPEN.

---


##### `TRACKING` <a name="TRACKING" id="@mongodbatlas-awscdk/alert-configuration.AlertViewStatus.TRACKING"></a>

TRACKING.

---


### CfnAlertConfigurationPropsEventTypeName <a name="CfnAlertConfigurationPropsEventTypeName" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName"></a>

Event type that triggers an alert.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.AWS_ENCRYPTION_KEY_NEEDS_ROTATION">AWS_ENCRYPTION_KEY_NEEDS_ROTATION</a></code> | AWS_ENCRYPTION_KEY_NEEDS_ROTATION. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.AZURE_ENCRYPTION_KEY_NEEDS_ROTATION">AZURE_ENCRYPTION_KEY_NEEDS_ROTATION</a></code> | AZURE_ENCRYPTION_KEY_NEEDS_ROTATION. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CLUSTER_MONGOS_IS_MISSING">CLUSTER_MONGOS_IS_MISSING</a></code> | CLUSTER_MONGOS_IS_MISSING. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_RESTORE_FAILED">CPS_RESTORE_FAILED</a></code> | CPS_RESTORE_FAILED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_RESTORE_SUCCESSFUL">CPS_RESTORE_SUCCESSFUL</a></code> | CPS_RESTORE_SUCCESSFUL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_BEHIND">CPS_SNAPSHOT_BEHIND</a></code> | CPS_SNAPSHOT_BEHIND. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED">CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED</a></code> | CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_FALLBACK_FAILED">CPS_SNAPSHOT_FALLBACK_FAILED</a></code> | CPS_SNAPSHOT_FALLBACK_FAILED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_FALLBACK_SUCCESSFUL">CPS_SNAPSHOT_FALLBACK_SUCCESSFUL</a></code> | CPS_SNAPSHOT_FALLBACK_SUCCESSFUL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_SUCCESSFUL">CPS_SNAPSHOT_SUCCESSFUL</a></code> | CPS_SNAPSHOT_SUCCESSFUL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CREDIT_CARD_ABOUT_TO_EXPIRE">CREDIT_CARD_ABOUT_TO_EXPIRE</a></code> | CREDIT_CARD_ABOUT_TO_EXPIRE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.DAILY_BILL_OVER_THRESHOLD">DAILY_BILL_OVER_THRESHOLD</a></code> | DAILY_BILL_OVER_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.GCP_ENCRYPTION_KEY_NEEDS_ROTATION">GCP_ENCRYPTION_KEY_NEEDS_ROTATION</a></code> | GCP_ENCRYPTION_KEY_NEEDS_ROTATION. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.HOST_DOWN">HOST_DOWN</a></code> | HOST_DOWN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.JOINED_GROUP">JOINED_GROUP</a></code> | JOINED_GROUP. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK">NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK</a></code> | NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK">NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK</a></code> | NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.NO_PRIMARY">NO_PRIMARY</a></code> | NO_PRIMARY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.OUTSIDE_METRIC_THRESHOLD">OUTSIDE_METRIC_THRESHOLD</a></code> | OUTSIDE_METRIC_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.OUTSIDE_SERVERLESS_METRIC_THRESHOLD">OUTSIDE_SERVERLESS_METRIC_THRESHOLD</a></code> | OUTSIDE_SERVERLESS_METRIC_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.OUTSIDE_REALM_METRIC_THRESHOLD">OUTSIDE_REALM_METRIC_THRESHOLD</a></code> | OUTSIDE_REALM_METRIC_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.PENDING_INVOICE_OVER_THRESHOLD">PENDING_INVOICE_OVER_THRESHOLD</a></code> | PENDING_INVOICE_OVER_THRESHOLD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.PRIMARY_ELECTED">PRIMARY_ELECTED</a></code> | PRIMARY_ELECTED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.REMOVED_FROM_GROUP">REMOVED_FROM_GROUP</a></code> | REMOVED_FROM_GROUP. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.REPLICATION_OPLOG_WINDOW_RUNNING_OUT">REPLICATION_OPLOG_WINDOW_RUNNING_OUT</a></code> | REPLICATION_OPLOG_WINDOW_RUNNING_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.TOO_MANY_ELECTIONS">TOO_MANY_ELECTIONS</a></code> | TOO_MANY_ELECTIONS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.USER_ROLES_CHANGED_AUDIT">USER_ROLES_CHANGED_AUDIT</a></code> | USER_ROLES_CHANGED_AUDIT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.USERS_WITHOUT_MULTIFACTOR_AUTH">USERS_WITHOUT_MULTIFACTOR_AUTH</a></code> | USERS_WITHOUT_MULTIFACTOR_AUTH. |

---

##### `AWS_ENCRYPTION_KEY_NEEDS_ROTATION` <a name="AWS_ENCRYPTION_KEY_NEEDS_ROTATION" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.AWS_ENCRYPTION_KEY_NEEDS_ROTATION"></a>

AWS_ENCRYPTION_KEY_NEEDS_ROTATION.

---


##### `AZURE_ENCRYPTION_KEY_NEEDS_ROTATION` <a name="AZURE_ENCRYPTION_KEY_NEEDS_ROTATION" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.AZURE_ENCRYPTION_KEY_NEEDS_ROTATION"></a>

AZURE_ENCRYPTION_KEY_NEEDS_ROTATION.

---


##### `CLUSTER_MONGOS_IS_MISSING` <a name="CLUSTER_MONGOS_IS_MISSING" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CLUSTER_MONGOS_IS_MISSING"></a>

CLUSTER_MONGOS_IS_MISSING.

---


##### `CPS_RESTORE_FAILED` <a name="CPS_RESTORE_FAILED" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_RESTORE_FAILED"></a>

CPS_RESTORE_FAILED.

---


##### `CPS_RESTORE_SUCCESSFUL` <a name="CPS_RESTORE_SUCCESSFUL" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_RESTORE_SUCCESSFUL"></a>

CPS_RESTORE_SUCCESSFUL.

---


##### `CPS_SNAPSHOT_BEHIND` <a name="CPS_SNAPSHOT_BEHIND" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_BEHIND"></a>

CPS_SNAPSHOT_BEHIND.

---


##### `CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED` <a name="CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED"></a>

CPS_SNAPSHOT_DOWNLOAD_REQUEST_FAILED.

---


##### `CPS_SNAPSHOT_FALLBACK_FAILED` <a name="CPS_SNAPSHOT_FALLBACK_FAILED" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_FALLBACK_FAILED"></a>

CPS_SNAPSHOT_FALLBACK_FAILED.

---


##### `CPS_SNAPSHOT_FALLBACK_SUCCESSFUL` <a name="CPS_SNAPSHOT_FALLBACK_SUCCESSFUL" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_FALLBACK_SUCCESSFUL"></a>

CPS_SNAPSHOT_FALLBACK_SUCCESSFUL.

---


##### `CPS_SNAPSHOT_SUCCESSFUL` <a name="CPS_SNAPSHOT_SUCCESSFUL" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CPS_SNAPSHOT_SUCCESSFUL"></a>

CPS_SNAPSHOT_SUCCESSFUL.

---


##### `CREDIT_CARD_ABOUT_TO_EXPIRE` <a name="CREDIT_CARD_ABOUT_TO_EXPIRE" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.CREDIT_CARD_ABOUT_TO_EXPIRE"></a>

CREDIT_CARD_ABOUT_TO_EXPIRE.

---


##### `DAILY_BILL_OVER_THRESHOLD` <a name="DAILY_BILL_OVER_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.DAILY_BILL_OVER_THRESHOLD"></a>

DAILY_BILL_OVER_THRESHOLD.

---


##### `GCP_ENCRYPTION_KEY_NEEDS_ROTATION` <a name="GCP_ENCRYPTION_KEY_NEEDS_ROTATION" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.GCP_ENCRYPTION_KEY_NEEDS_ROTATION"></a>

GCP_ENCRYPTION_KEY_NEEDS_ROTATION.

---


##### `HOST_DOWN` <a name="HOST_DOWN" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.HOST_DOWN"></a>

HOST_DOWN.

---


##### `JOINED_GROUP` <a name="JOINED_GROUP" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.JOINED_GROUP"></a>

JOINED_GROUP.

---


##### `NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK` <a name="NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK"></a>

NDS_X509_USER_AUTHENTICATION_CUSTOMER_CA_EXPIRATION_CHECK.

---


##### `NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK` <a name="NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK"></a>

NDS_X509_USER_AUTHENTICATION_MANAGED_USER_CERTS_EXPIRATION_CHECK.

---


##### `NO_PRIMARY` <a name="NO_PRIMARY" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.NO_PRIMARY"></a>

NO_PRIMARY.

---


##### `OUTSIDE_METRIC_THRESHOLD` <a name="OUTSIDE_METRIC_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.OUTSIDE_METRIC_THRESHOLD"></a>

OUTSIDE_METRIC_THRESHOLD.

---


##### `OUTSIDE_SERVERLESS_METRIC_THRESHOLD` <a name="OUTSIDE_SERVERLESS_METRIC_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.OUTSIDE_SERVERLESS_METRIC_THRESHOLD"></a>

OUTSIDE_SERVERLESS_METRIC_THRESHOLD.

---


##### `OUTSIDE_REALM_METRIC_THRESHOLD` <a name="OUTSIDE_REALM_METRIC_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.OUTSIDE_REALM_METRIC_THRESHOLD"></a>

OUTSIDE_REALM_METRIC_THRESHOLD.

---


##### `PENDING_INVOICE_OVER_THRESHOLD` <a name="PENDING_INVOICE_OVER_THRESHOLD" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.PENDING_INVOICE_OVER_THRESHOLD"></a>

PENDING_INVOICE_OVER_THRESHOLD.

---


##### `PRIMARY_ELECTED` <a name="PRIMARY_ELECTED" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.PRIMARY_ELECTED"></a>

PRIMARY_ELECTED.

---


##### `REMOVED_FROM_GROUP` <a name="REMOVED_FROM_GROUP" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.REMOVED_FROM_GROUP"></a>

REMOVED_FROM_GROUP.

---


##### `REPLICATION_OPLOG_WINDOW_RUNNING_OUT` <a name="REPLICATION_OPLOG_WINDOW_RUNNING_OUT" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.REPLICATION_OPLOG_WINDOW_RUNNING_OUT"></a>

REPLICATION_OPLOG_WINDOW_RUNNING_OUT.

---


##### `TOO_MANY_ELECTIONS` <a name="TOO_MANY_ELECTIONS" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.TOO_MANY_ELECTIONS"></a>

TOO_MANY_ELECTIONS.

---


##### `USER_ROLES_CHANGED_AUDIT` <a name="USER_ROLES_CHANGED_AUDIT" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.USER_ROLES_CHANGED_AUDIT"></a>

USER_ROLES_CHANGED_AUDIT.

---


##### `USERS_WITHOUT_MULTIFACTOR_AUTH` <a name="USERS_WITHOUT_MULTIFACTOR_AUTH" id="@mongodbatlas-awscdk/alert-configuration.CfnAlertConfigurationPropsEventTypeName.USERS_WITHOUT_MULTIFACTOR_AUTH"></a>

USERS_WITHOUT_MULTIFACTOR_AUTH.

---


### CurrentValueUnits <a name="CurrentValueUnits" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits"></a>

Element used to express the quantity in **currentValue.number**. This can be an element of time, storage capacity, and the like. This metric triggered the alert.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.BITS">BITS</a></code> | BITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.BYTES">BYTES</a></code> | BYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.DAYS">DAYS</a></code> | DAYS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.GIGABITS">GIGABITS</a></code> | GIGABITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.GIGABYTES">GIGABYTES</a></code> | GIGABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.HOURS">HOURS</a></code> | HOURS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.KILOBITS">KILOBITS</a></code> | KILOBITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.KILOBYTES">KILOBYTES</a></code> | KILOBYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.MEGABITS">MEGABITS</a></code> | MEGABITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.MEGABYTES">MEGABYTES</a></code> | MEGABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.MILLISECONDS">MILLISECONDS</a></code> | MILLISECONDS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.MINUTES">MINUTES</a></code> | MINUTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.PETABYTES">PETABYTES</a></code> | PETABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.RAW">RAW</a></code> | RAW. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.SECONDS">SECONDS</a></code> | SECONDS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.TERABYTES">TERABYTES</a></code> | TERABYTES. |

---

##### `BITS` <a name="BITS" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.BITS"></a>

BITS.

---


##### `BYTES` <a name="BYTES" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.BYTES"></a>

BYTES.

---


##### `DAYS` <a name="DAYS" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.DAYS"></a>

DAYS.

---


##### `GIGABITS` <a name="GIGABITS" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.GIGABITS"></a>

GIGABITS.

---


##### `GIGABYTES` <a name="GIGABYTES" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.GIGABYTES"></a>

GIGABYTES.

---


##### `HOURS` <a name="HOURS" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.HOURS"></a>

HOURS.

---


##### `KILOBITS` <a name="KILOBITS" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.KILOBITS"></a>

KILOBITS.

---


##### `KILOBYTES` <a name="KILOBYTES" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.KILOBYTES"></a>

KILOBYTES.

---


##### `MEGABITS` <a name="MEGABITS" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.MEGABITS"></a>

MEGABITS.

---


##### `MEGABYTES` <a name="MEGABYTES" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.MEGABYTES"></a>

MEGABYTES.

---


##### `MILLISECONDS` <a name="MILLISECONDS" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.MILLISECONDS"></a>

MILLISECONDS.

---


##### `MINUTES` <a name="MINUTES" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.MINUTES"></a>

MINUTES.

---


##### `PETABYTES` <a name="PETABYTES" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.PETABYTES"></a>

PETABYTES.

---


##### `RAW` <a name="RAW" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.RAW"></a>

RAW.

---


##### `SECONDS` <a name="SECONDS" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.SECONDS"></a>

SECONDS.

---


##### `TERABYTES` <a name="TERABYTES" id="@mongodbatlas-awscdk/alert-configuration.CurrentValueUnits.TERABYTES"></a>

TERABYTES.

---


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


### IntegerThresholdViewUnits <a name="IntegerThresholdViewUnits" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits"></a>

Element used to express the quantity.

This can be an element of time, storage capacity, and the like.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.BITS">BITS</a></code> | BITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.BYTES">BYTES</a></code> | BYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.DAYS">DAYS</a></code> | DAYS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.GIGABITS">GIGABITS</a></code> | GIGABITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.GIGABYTES">GIGABYTES</a></code> | GIGABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.HOURS">HOURS</a></code> | HOURS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.KILOBITS">KILOBITS</a></code> | KILOBITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.KILOBYTES">KILOBYTES</a></code> | KILOBYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.MEGABITS">MEGABITS</a></code> | MEGABITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.MEGABYTES">MEGABYTES</a></code> | MEGABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.MILLISECONDS">MILLISECONDS</a></code> | MILLISECONDS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.MINUTES">MINUTES</a></code> | MINUTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.PETABYTES">PETABYTES</a></code> | PETABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.RAW">RAW</a></code> | RAW. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.SECONDS">SECONDS</a></code> | SECONDS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.TERABYTES">TERABYTES</a></code> | TERABYTES. |

---

##### `BITS` <a name="BITS" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.BITS"></a>

BITS.

---


##### `BYTES` <a name="BYTES" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.BYTES"></a>

BYTES.

---


##### `DAYS` <a name="DAYS" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.DAYS"></a>

DAYS.

---


##### `GIGABITS` <a name="GIGABITS" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.GIGABITS"></a>

GIGABITS.

---


##### `GIGABYTES` <a name="GIGABYTES" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.GIGABYTES"></a>

GIGABYTES.

---


##### `HOURS` <a name="HOURS" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.HOURS"></a>

HOURS.

---


##### `KILOBITS` <a name="KILOBITS" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.KILOBITS"></a>

KILOBITS.

---


##### `KILOBYTES` <a name="KILOBYTES" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.KILOBYTES"></a>

KILOBYTES.

---


##### `MEGABITS` <a name="MEGABITS" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.MEGABITS"></a>

MEGABITS.

---


##### `MEGABYTES` <a name="MEGABYTES" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.MEGABYTES"></a>

MEGABYTES.

---


##### `MILLISECONDS` <a name="MILLISECONDS" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.MILLISECONDS"></a>

MILLISECONDS.

---


##### `MINUTES` <a name="MINUTES" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.MINUTES"></a>

MINUTES.

---


##### `PETABYTES` <a name="PETABYTES" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.PETABYTES"></a>

PETABYTES.

---


##### `RAW` <a name="RAW" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.RAW"></a>

RAW.

---


##### `SECONDS` <a name="SECONDS" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.SECONDS"></a>

SECONDS.

---


##### `TERABYTES` <a name="TERABYTES" id="@mongodbatlas-awscdk/alert-configuration.IntegerThresholdViewUnits.TERABYTES"></a>

TERABYTES.

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


### MetricThresholdViewMetricName <a name="MetricThresholdViewMetricName" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName"></a>

Human-readable label that identifies the metric against which MongoDB Cloud checks the configured **metricThreshold.threshold**.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.ASSERT_MSG">ASSERT_MSG</a></code> | ASSERT_MSG. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.ASSERT_REGULAR">ASSERT_REGULAR</a></code> | ASSERT_REGULAR. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.ASSERT_USER">ASSERT_USER</a></code> | ASSERT_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.ASSERT_WARNING">ASSERT_WARNING</a></code> | ASSERT_WARNING. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.AVG_COMMAND_EXECUTION_TIME">AVG_COMMAND_EXECUTION_TIME</a></code> | AVG_COMMAND_EXECUTION_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.AVG_READ_EXECUTION_TIME">AVG_READ_EXECUTION_TIME</a></code> | AVG_READ_EXECUTION_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.AVG_WRITE_EXECUTION_TIME">AVG_WRITE_EXECUTION_TIME</a></code> | AVG_WRITE_EXECUTION_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.BACKGROUND_FLUSH_AVG">BACKGROUND_FLUSH_AVG</a></code> | BACKGROUND_FLUSH_AVG. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CACHE_BYTES_READ_INTO">CACHE_BYTES_READ_INTO</a></code> | CACHE_BYTES_READ_INTO. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CACHE_BYTES_WRITTEN_FROM">CACHE_BYTES_WRITTEN_FROM</a></code> | CACHE_BYTES_WRITTEN_FROM. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CACHE_USAGE_DIRTY">CACHE_USAGE_DIRTY</a></code> | CACHE_USAGE_DIRTY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CACHE_USAGE_USED">CACHE_USAGE_USED</a></code> | CACHE_USAGE_USED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.COMPUTED_MEMORY">COMPUTED_MEMORY</a></code> | COMPUTED_MEMORY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CONNECTIONS">CONNECTIONS</a></code> | CONNECTIONS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CONNECTIONS_MAX">CONNECTIONS_MAX</a></code> | CONNECTIONS_MAX. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CONNECTIONS_PERCENT">CONNECTIONS_PERCENT</a></code> | CONNECTIONS_PERCENT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CURSORS_TOTAL_CLIENT_CURSORS_SIZE">CURSORS_TOTAL_CLIENT_CURSORS_SIZE</a></code> | CURSORS_TOTAL_CLIENT_CURSORS_SIZE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CURSORS_TOTAL_OPEN">CURSORS_TOTAL_OPEN</a></code> | CURSORS_TOTAL_OPEN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CURSORS_TOTAL_TIMED_OUT">CURSORS_TOTAL_TIMED_OUT</a></code> | CURSORS_TOTAL_TIMED_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DB_DATA_SIZE_TOTAL">DB_DATA_SIZE_TOTAL</a></code> | DB_DATA_SIZE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DB_INDEX_SIZE_TOTAL">DB_INDEX_SIZE_TOTAL</a></code> | DB_INDEX_SIZE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DB_STORAGE_TOTAL">DB_STORAGE_TOTAL</a></code> | DB_STORAGE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_SPACE_USED_DATA">DISK_PARTITION_SPACE_USED_DATA</a></code> | DISK_PARTITION_SPACE_USED_DATA. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_SPACE_USED_INDEX">DISK_PARTITION_SPACE_USED_INDEX</a></code> | DISK_PARTITION_SPACE_USED_INDEX. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_SPACE_USED_JOURNAL">DISK_PARTITION_SPACE_USED_JOURNAL</a></code> | DISK_PARTITION_SPACE_USED_JOURNAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_UTILIZATION_DATA">DISK_PARTITION_UTILIZATION_DATA</a></code> | DISK_PARTITION_UTILIZATION_DATA. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_UTILIZATION_INDEX">DISK_PARTITION_UTILIZATION_INDEX</a></code> | DISK_PARTITION_UTILIZATION_INDEX. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_UTILIZATION_JOURNAL">DISK_PARTITION_UTILIZATION_JOURNAL</a></code> | DISK_PARTITION_UTILIZATION_JOURNAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DOCUMENT_DELETED">DOCUMENT_DELETED</a></code> | DOCUMENT_DELETED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DOCUMENT_INSERTED">DOCUMENT_INSERTED</a></code> | DOCUMENT_INSERTED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DOCUMENT_RETURNED">DOCUMENT_RETURNED</a></code> | DOCUMENT_RETURNED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DOCUMENT_UPDATED">DOCUMENT_UPDATED</a></code> | DOCUMENT_UPDATED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.EXTRA_INFO_PAGE_FAULTS">EXTRA_INFO_PAGE_FAULTS</a></code> | EXTRA_INFO_PAGE_FAULTS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_DISK_UTILIZATION">FTS_DISK_UTILIZATION</a></code> | FTS_DISK_UTILIZATION. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_MEMORY_MAPPED">FTS_MEMORY_MAPPED</a></code> | FTS_MEMORY_MAPPED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_MEMORY_RESIDENT">FTS_MEMORY_RESIDENT</a></code> | FTS_MEMORY_RESIDENT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_MEMORY_VIRTUAL">FTS_MEMORY_VIRTUAL</a></code> | FTS_MEMORY_VIRTUAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_PROCESS_CPU_KERNEL">FTS_PROCESS_CPU_KERNEL</a></code> | FTS_PROCESS_CPU_KERNEL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_PROCESS_CPU_USER">FTS_PROCESS_CPU_USER</a></code> | FTS_PROCESS_CPU_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_ACCESSES_NOT_IN_MEMORY">GLOBAL_ACCESSES_NOT_IN_MEMORY</a></code> | GLOBAL_ACCESSES_NOT_IN_MEMORY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_READERS">GLOBAL_LOCK_CURRENT_QUEUE_READERS</a></code> | GLOBAL_LOCK_CURRENT_QUEUE_READERS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_TOTAL">GLOBAL_LOCK_CURRENT_QUEUE_TOTAL</a></code> | GLOBAL_LOCK_CURRENT_QUEUE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_WRITERS">GLOBAL_LOCK_CURRENT_QUEUE_WRITERS</a></code> | GLOBAL_LOCK_CURRENT_QUEUE_WRITERS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_LOCK_PERCENTAGE">GLOBAL_LOCK_PERCENTAGE</a></code> | GLOBAL_LOCK_PERCENTAGE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN">GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN</a></code> | GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.INDEX_COUNTERS_BTREE_ACCESSES">INDEX_COUNTERS_BTREE_ACCESSES</a></code> | INDEX_COUNTERS_BTREE_ACCESSES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.INDEX_COUNTERS_BTREE_HITS">INDEX_COUNTERS_BTREE_HITS</a></code> | INDEX_COUNTERS_BTREE_HITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.INDEX_COUNTERS_BTREE_MISS_RATIO">INDEX_COUNTERS_BTREE_MISS_RATIO</a></code> | INDEX_COUNTERS_BTREE_MISS_RATIO. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.INDEX_COUNTERS_BTREE_MISSES">INDEX_COUNTERS_BTREE_MISSES</a></code> | INDEX_COUNTERS_BTREE_MISSES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.JOURNALING_COMMITS_IN_WRITE_LOCK">JOURNALING_COMMITS_IN_WRITE_LOCK</a></code> | JOURNALING_COMMITS_IN_WRITE_LOCK. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.JOURNALING_MB">JOURNALING_MB</a></code> | JOURNALING_MB. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.JOURNALING_WRITE_DATA_FILES_MB">JOURNALING_WRITE_DATA_FILES_MB</a></code> | JOURNALING_WRITE_DATA_FILES_MB. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.LOGICAL_SIZE">LOGICAL_SIZE</a></code> | LOGICAL_SIZE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MEMORY_MAPPED">MEMORY_MAPPED</a></code> | MEMORY_MAPPED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MEMORY_RESIDENT">MEMORY_RESIDENT</a></code> | MEMORY_RESIDENT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MEMORY_VIRTUAL">MEMORY_VIRTUAL</a></code> | MEMORY_VIRTUAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_IOWAIT">MUNIN_CPU_IOWAIT</a></code> | MUNIN_CPU_IOWAIT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_IRQ">MUNIN_CPU_IRQ</a></code> | MUNIN_CPU_IRQ. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_NICE">MUNIN_CPU_NICE</a></code> | MUNIN_CPU_NICE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_SOFTIRQ">MUNIN_CPU_SOFTIRQ</a></code> | MUNIN_CPU_SOFTIRQ. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_STEAL">MUNIN_CPU_STEAL</a></code> | MUNIN_CPU_STEAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_SYSTEM">MUNIN_CPU_SYSTEM</a></code> | MUNIN_CPU_SYSTEM. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_USER">MUNIN_CPU_USER</a></code> | MUNIN_CPU_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NETWORK_BYTES_IN">NETWORK_BYTES_IN</a></code> | NETWORK_BYTES_IN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NETWORK_BYTES_OUT">NETWORK_BYTES_OUT</a></code> | NETWORK_BYTES_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NETWORK_NUM_REQUESTS">NETWORK_NUM_REQUESTS</a></code> | NETWORK_NUM_REQUESTS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NORMALIZED_FTS_PROCESS_CPU_KERNEL">NORMALIZED_FTS_PROCESS_CPU_KERNEL</a></code> | NORMALIZED_FTS_PROCESS_CPU_KERNEL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NORMALIZED_FTS_PROCESS_CPU_USER">NORMALIZED_FTS_PROCESS_CPU_USER</a></code> | NORMALIZED_FTS_PROCESS_CPU_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NORMALIZED_SYSTEM_CPU_STEAL">NORMALIZED_SYSTEM_CPU_STEAL</a></code> | NORMALIZED_SYSTEM_CPU_STEAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NORMALIZED_SYSTEM_CPU_USER">NORMALIZED_SYSTEM_CPU_USER</a></code> | NORMALIZED_SYSTEM_CPU_USER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_CMD">OPCOUNTER_CMD</a></code> | OPCOUNTER_CMD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_DELETE">OPCOUNTER_DELETE</a></code> | OPCOUNTER_DELETE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_GETMORE">OPCOUNTER_GETMORE</a></code> | OPCOUNTER_GETMORE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_INSERT">OPCOUNTER_INSERT</a></code> | OPCOUNTER_INSERT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_QUERY">OPCOUNTER_QUERY</a></code> | OPCOUNTER_QUERY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_REPL_CMD">OPCOUNTER_REPL_CMD</a></code> | OPCOUNTER_REPL_CMD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_REPL_DELETE">OPCOUNTER_REPL_DELETE</a></code> | OPCOUNTER_REPL_DELETE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_REPL_INSERT">OPCOUNTER_REPL_INSERT</a></code> | OPCOUNTER_REPL_INSERT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_REPL_UPDATE">OPCOUNTER_REPL_UPDATE</a></code> | OPCOUNTER_REPL_UPDATE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_UPDATE">OPCOUNTER_UPDATE</a></code> | OPCOUNTER_UPDATE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPERATIONS_SCAN_AND_ORDER">OPERATIONS_SCAN_AND_ORDER</a></code> | OPERATIONS_SCAN_AND_ORDER. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_MASTER_LAG_TIME_DIFF">OPLOG_MASTER_LAG_TIME_DIFF</a></code> | OPLOG_MASTER_LAG_TIME_DIFF. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_MASTER_TIME">OPLOG_MASTER_TIME</a></code> | OPLOG_MASTER_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_MASTER_TIME_ESTIMATED_TTL">OPLOG_MASTER_TIME_ESTIMATED_TTL</a></code> | OPLOG_MASTER_TIME_ESTIMATED_TTL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_RATE_GB_PER_HOUR">OPLOG_RATE_GB_PER_HOUR</a></code> | OPLOG_RATE_GB_PER_HOUR. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_SLAVE_LAG_MASTER_TIME">OPLOG_SLAVE_LAG_MASTER_TIME</a></code> | OPLOG_SLAVE_LAG_MASTER_TIME. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.QUERY_EXECUTOR_SCANNED">QUERY_EXECUTOR_SCANNED</a></code> | QUERY_EXECUTOR_SCANNED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.QUERY_EXECUTOR_SCANNED_OBJECTS">QUERY_EXECUTOR_SCANNED_OBJECTS</a></code> | QUERY_EXECUTOR_SCANNED_OBJECTS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED">QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED</a></code> | QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.QUERY_TARGETING_SCANNED_PER_RETURNED">QUERY_TARGETING_SCANNED_PER_RETURNED</a></code> | QUERY_TARGETING_SCANNED_PER_RETURNED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.RESTARTS_IN_LAST_HOUR">RESTARTS_IN_LAST_HOUR</a></code> | RESTARTS_IN_LAST_HOUR. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_CONNECTIONS">SERVERLESS_CONNECTIONS</a></code> | SERVERLESS_CONNECTIONS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_CONNECTIONS_PERCENT">SERVERLESS_CONNECTIONS_PERCENT</a></code> | SERVERLESS_CONNECTIONS_PERCENT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_DATA_SIZE_TOTAL">SERVERLESS_DATA_SIZE_TOTAL</a></code> | SERVERLESS_DATA_SIZE_TOTAL. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_NETWORK_BYTES_IN">SERVERLESS_NETWORK_BYTES_IN</a></code> | SERVERLESS_NETWORK_BYTES_IN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_NETWORK_BYTES_OUT">SERVERLESS_NETWORK_BYTES_OUT</a></code> | SERVERLESS_NETWORK_BYTES_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_NETWORK_NUM_REQUESTS">SERVERLESS_NETWORK_NUM_REQUESTS</a></code> | SERVERLESS_NETWORK_NUM_REQUESTS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_CMD">SERVERLESS_OPCOUNTER_CMD</a></code> | SERVERLESS_OPCOUNTER_CMD. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_DELETE">SERVERLESS_OPCOUNTER_DELETE</a></code> | SERVERLESS_OPCOUNTER_DELETE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_GETMORE">SERVERLESS_OPCOUNTER_GETMORE</a></code> | SERVERLESS_OPCOUNTER_GETMORE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_INSERT">SERVERLESS_OPCOUNTER_INSERT</a></code> | SERVERLESS_OPCOUNTER_INSERT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_QUERY">SERVERLESS_OPCOUNTER_QUERY</a></code> | SERVERLESS_OPCOUNTER_QUERY. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_UPDATE">SERVERLESS_OPCOUNTER_UPDATE</a></code> | SERVERLESS_OPCOUNTER_UPDATE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_TOTAL_READ_UNITS">SERVERLESS_TOTAL_READ_UNITS</a></code> | SERVERLESS_TOTAL_READ_UNITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_TOTAL_WRITE_UNITS">SERVERLESS_TOTAL_WRITE_UNITS</a></code> | SERVERLESS_TOTAL_WRITE_UNITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SWAP_USAGE_FREE">SWAP_USAGE_FREE</a></code> | SWAP_USAGE_FREE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SWAP_USAGE_USED">SWAP_USAGE_USED</a></code> | SWAP_USAGE_USED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_MEMORY_AVAILABLE">SYSTEM_MEMORY_AVAILABLE</a></code> | SYSTEM_MEMORY_AVAILABLE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_MEMORY_FREE">SYSTEM_MEMORY_FREE</a></code> | SYSTEM_MEMORY_FREE. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_MEMORY_USED">SYSTEM_MEMORY_USED</a></code> | SYSTEM_MEMORY_USED. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_NETWORK_IN">SYSTEM_NETWORK_IN</a></code> | SYSTEM_NETWORK_IN. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_NETWORK_OUT">SYSTEM_NETWORK_OUT</a></code> | SYSTEM_NETWORK_OUT. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.TICKETS_AVAILABLE_READS">TICKETS_AVAILABLE_READS</a></code> | TICKETS_AVAILABLE_READS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.TICKETS_AVAILABLE_WRITES">TICKETS_AVAILABLE_WRITES</a></code> | TICKETS_AVAILABLE_WRITES. |

---

##### `ASSERT_MSG` <a name="ASSERT_MSG" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.ASSERT_MSG"></a>

ASSERT_MSG.

---


##### `ASSERT_REGULAR` <a name="ASSERT_REGULAR" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.ASSERT_REGULAR"></a>

ASSERT_REGULAR.

---


##### `ASSERT_USER` <a name="ASSERT_USER" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.ASSERT_USER"></a>

ASSERT_USER.

---


##### `ASSERT_WARNING` <a name="ASSERT_WARNING" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.ASSERT_WARNING"></a>

ASSERT_WARNING.

---


##### `AVG_COMMAND_EXECUTION_TIME` <a name="AVG_COMMAND_EXECUTION_TIME" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.AVG_COMMAND_EXECUTION_TIME"></a>

AVG_COMMAND_EXECUTION_TIME.

---


##### `AVG_READ_EXECUTION_TIME` <a name="AVG_READ_EXECUTION_TIME" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.AVG_READ_EXECUTION_TIME"></a>

AVG_READ_EXECUTION_TIME.

---


##### `AVG_WRITE_EXECUTION_TIME` <a name="AVG_WRITE_EXECUTION_TIME" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.AVG_WRITE_EXECUTION_TIME"></a>

AVG_WRITE_EXECUTION_TIME.

---


##### `BACKGROUND_FLUSH_AVG` <a name="BACKGROUND_FLUSH_AVG" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.BACKGROUND_FLUSH_AVG"></a>

BACKGROUND_FLUSH_AVG.

---


##### `CACHE_BYTES_READ_INTO` <a name="CACHE_BYTES_READ_INTO" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CACHE_BYTES_READ_INTO"></a>

CACHE_BYTES_READ_INTO.

---


##### `CACHE_BYTES_WRITTEN_FROM` <a name="CACHE_BYTES_WRITTEN_FROM" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CACHE_BYTES_WRITTEN_FROM"></a>

CACHE_BYTES_WRITTEN_FROM.

---


##### `CACHE_USAGE_DIRTY` <a name="CACHE_USAGE_DIRTY" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CACHE_USAGE_DIRTY"></a>

CACHE_USAGE_DIRTY.

---


##### `CACHE_USAGE_USED` <a name="CACHE_USAGE_USED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CACHE_USAGE_USED"></a>

CACHE_USAGE_USED.

---


##### `COMPUTED_MEMORY` <a name="COMPUTED_MEMORY" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.COMPUTED_MEMORY"></a>

COMPUTED_MEMORY.

---


##### `CONNECTIONS` <a name="CONNECTIONS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CONNECTIONS"></a>

CONNECTIONS.

---


##### `CONNECTIONS_MAX` <a name="CONNECTIONS_MAX" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CONNECTIONS_MAX"></a>

CONNECTIONS_MAX.

---


##### `CONNECTIONS_PERCENT` <a name="CONNECTIONS_PERCENT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CONNECTIONS_PERCENT"></a>

CONNECTIONS_PERCENT.

---


##### `CURSORS_TOTAL_CLIENT_CURSORS_SIZE` <a name="CURSORS_TOTAL_CLIENT_CURSORS_SIZE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CURSORS_TOTAL_CLIENT_CURSORS_SIZE"></a>

CURSORS_TOTAL_CLIENT_CURSORS_SIZE.

---


##### `CURSORS_TOTAL_OPEN` <a name="CURSORS_TOTAL_OPEN" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CURSORS_TOTAL_OPEN"></a>

CURSORS_TOTAL_OPEN.

---


##### `CURSORS_TOTAL_TIMED_OUT` <a name="CURSORS_TOTAL_TIMED_OUT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.CURSORS_TOTAL_TIMED_OUT"></a>

CURSORS_TOTAL_TIMED_OUT.

---


##### `DB_DATA_SIZE_TOTAL` <a name="DB_DATA_SIZE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DB_DATA_SIZE_TOTAL"></a>

DB_DATA_SIZE_TOTAL.

---


##### `DB_INDEX_SIZE_TOTAL` <a name="DB_INDEX_SIZE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DB_INDEX_SIZE_TOTAL"></a>

DB_INDEX_SIZE_TOTAL.

---


##### `DB_STORAGE_TOTAL` <a name="DB_STORAGE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DB_STORAGE_TOTAL"></a>

DB_STORAGE_TOTAL.

---


##### `DISK_PARTITION_SPACE_USED_DATA` <a name="DISK_PARTITION_SPACE_USED_DATA" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_SPACE_USED_DATA"></a>

DISK_PARTITION_SPACE_USED_DATA.

---


##### `DISK_PARTITION_SPACE_USED_INDEX` <a name="DISK_PARTITION_SPACE_USED_INDEX" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_SPACE_USED_INDEX"></a>

DISK_PARTITION_SPACE_USED_INDEX.

---


##### `DISK_PARTITION_SPACE_USED_JOURNAL` <a name="DISK_PARTITION_SPACE_USED_JOURNAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_SPACE_USED_JOURNAL"></a>

DISK_PARTITION_SPACE_USED_JOURNAL.

---


##### `DISK_PARTITION_UTILIZATION_DATA` <a name="DISK_PARTITION_UTILIZATION_DATA" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_UTILIZATION_DATA"></a>

DISK_PARTITION_UTILIZATION_DATA.

---


##### `DISK_PARTITION_UTILIZATION_INDEX` <a name="DISK_PARTITION_UTILIZATION_INDEX" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_UTILIZATION_INDEX"></a>

DISK_PARTITION_UTILIZATION_INDEX.

---


##### `DISK_PARTITION_UTILIZATION_JOURNAL` <a name="DISK_PARTITION_UTILIZATION_JOURNAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DISK_PARTITION_UTILIZATION_JOURNAL"></a>

DISK_PARTITION_UTILIZATION_JOURNAL.

---


##### `DOCUMENT_DELETED` <a name="DOCUMENT_DELETED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DOCUMENT_DELETED"></a>

DOCUMENT_DELETED.

---


##### `DOCUMENT_INSERTED` <a name="DOCUMENT_INSERTED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DOCUMENT_INSERTED"></a>

DOCUMENT_INSERTED.

---


##### `DOCUMENT_RETURNED` <a name="DOCUMENT_RETURNED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DOCUMENT_RETURNED"></a>

DOCUMENT_RETURNED.

---


##### `DOCUMENT_UPDATED` <a name="DOCUMENT_UPDATED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.DOCUMENT_UPDATED"></a>

DOCUMENT_UPDATED.

---


##### `EXTRA_INFO_PAGE_FAULTS` <a name="EXTRA_INFO_PAGE_FAULTS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.EXTRA_INFO_PAGE_FAULTS"></a>

EXTRA_INFO_PAGE_FAULTS.

---


##### `FTS_DISK_UTILIZATION` <a name="FTS_DISK_UTILIZATION" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_DISK_UTILIZATION"></a>

FTS_DISK_UTILIZATION.

---


##### `FTS_MEMORY_MAPPED` <a name="FTS_MEMORY_MAPPED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_MEMORY_MAPPED"></a>

FTS_MEMORY_MAPPED.

---


##### `FTS_MEMORY_RESIDENT` <a name="FTS_MEMORY_RESIDENT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_MEMORY_RESIDENT"></a>

FTS_MEMORY_RESIDENT.

---


##### `FTS_MEMORY_VIRTUAL` <a name="FTS_MEMORY_VIRTUAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_MEMORY_VIRTUAL"></a>

FTS_MEMORY_VIRTUAL.

---


##### `FTS_PROCESS_CPU_KERNEL` <a name="FTS_PROCESS_CPU_KERNEL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_PROCESS_CPU_KERNEL"></a>

FTS_PROCESS_CPU_KERNEL.

---


##### `FTS_PROCESS_CPU_USER` <a name="FTS_PROCESS_CPU_USER" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.FTS_PROCESS_CPU_USER"></a>

FTS_PROCESS_CPU_USER.

---


##### `GLOBAL_ACCESSES_NOT_IN_MEMORY` <a name="GLOBAL_ACCESSES_NOT_IN_MEMORY" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_ACCESSES_NOT_IN_MEMORY"></a>

GLOBAL_ACCESSES_NOT_IN_MEMORY.

---


##### `GLOBAL_LOCK_CURRENT_QUEUE_READERS` <a name="GLOBAL_LOCK_CURRENT_QUEUE_READERS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_READERS"></a>

GLOBAL_LOCK_CURRENT_QUEUE_READERS.

---


##### `GLOBAL_LOCK_CURRENT_QUEUE_TOTAL` <a name="GLOBAL_LOCK_CURRENT_QUEUE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_TOTAL"></a>

GLOBAL_LOCK_CURRENT_QUEUE_TOTAL.

---


##### `GLOBAL_LOCK_CURRENT_QUEUE_WRITERS` <a name="GLOBAL_LOCK_CURRENT_QUEUE_WRITERS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_LOCK_CURRENT_QUEUE_WRITERS"></a>

GLOBAL_LOCK_CURRENT_QUEUE_WRITERS.

---


##### `GLOBAL_LOCK_PERCENTAGE` <a name="GLOBAL_LOCK_PERCENTAGE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_LOCK_PERCENTAGE"></a>

GLOBAL_LOCK_PERCENTAGE.

---


##### `GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN` <a name="GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN"></a>

GLOBAL_PAGE_FAULT_EXCEPTIONS_THROWN.

---


##### `INDEX_COUNTERS_BTREE_ACCESSES` <a name="INDEX_COUNTERS_BTREE_ACCESSES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.INDEX_COUNTERS_BTREE_ACCESSES"></a>

INDEX_COUNTERS_BTREE_ACCESSES.

---


##### `INDEX_COUNTERS_BTREE_HITS` <a name="INDEX_COUNTERS_BTREE_HITS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.INDEX_COUNTERS_BTREE_HITS"></a>

INDEX_COUNTERS_BTREE_HITS.

---


##### `INDEX_COUNTERS_BTREE_MISS_RATIO` <a name="INDEX_COUNTERS_BTREE_MISS_RATIO" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.INDEX_COUNTERS_BTREE_MISS_RATIO"></a>

INDEX_COUNTERS_BTREE_MISS_RATIO.

---


##### `INDEX_COUNTERS_BTREE_MISSES` <a name="INDEX_COUNTERS_BTREE_MISSES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.INDEX_COUNTERS_BTREE_MISSES"></a>

INDEX_COUNTERS_BTREE_MISSES.

---


##### `JOURNALING_COMMITS_IN_WRITE_LOCK` <a name="JOURNALING_COMMITS_IN_WRITE_LOCK" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.JOURNALING_COMMITS_IN_WRITE_LOCK"></a>

JOURNALING_COMMITS_IN_WRITE_LOCK.

---


##### `JOURNALING_MB` <a name="JOURNALING_MB" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.JOURNALING_MB"></a>

JOURNALING_MB.

---


##### `JOURNALING_WRITE_DATA_FILES_MB` <a name="JOURNALING_WRITE_DATA_FILES_MB" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.JOURNALING_WRITE_DATA_FILES_MB"></a>

JOURNALING_WRITE_DATA_FILES_MB.

---


##### `LOGICAL_SIZE` <a name="LOGICAL_SIZE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.LOGICAL_SIZE"></a>

LOGICAL_SIZE.

---


##### `MEMORY_MAPPED` <a name="MEMORY_MAPPED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MEMORY_MAPPED"></a>

MEMORY_MAPPED.

---


##### `MEMORY_RESIDENT` <a name="MEMORY_RESIDENT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MEMORY_RESIDENT"></a>

MEMORY_RESIDENT.

---


##### `MEMORY_VIRTUAL` <a name="MEMORY_VIRTUAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MEMORY_VIRTUAL"></a>

MEMORY_VIRTUAL.

---


##### `MUNIN_CPU_IOWAIT` <a name="MUNIN_CPU_IOWAIT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_IOWAIT"></a>

MUNIN_CPU_IOWAIT.

---


##### `MUNIN_CPU_IRQ` <a name="MUNIN_CPU_IRQ" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_IRQ"></a>

MUNIN_CPU_IRQ.

---


##### `MUNIN_CPU_NICE` <a name="MUNIN_CPU_NICE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_NICE"></a>

MUNIN_CPU_NICE.

---


##### `MUNIN_CPU_SOFTIRQ` <a name="MUNIN_CPU_SOFTIRQ" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_SOFTIRQ"></a>

MUNIN_CPU_SOFTIRQ.

---


##### `MUNIN_CPU_STEAL` <a name="MUNIN_CPU_STEAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_STEAL"></a>

MUNIN_CPU_STEAL.

---


##### `MUNIN_CPU_SYSTEM` <a name="MUNIN_CPU_SYSTEM" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_SYSTEM"></a>

MUNIN_CPU_SYSTEM.

---


##### `MUNIN_CPU_USER` <a name="MUNIN_CPU_USER" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.MUNIN_CPU_USER"></a>

MUNIN_CPU_USER.

---


##### `NETWORK_BYTES_IN` <a name="NETWORK_BYTES_IN" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NETWORK_BYTES_IN"></a>

NETWORK_BYTES_IN.

---


##### `NETWORK_BYTES_OUT` <a name="NETWORK_BYTES_OUT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NETWORK_BYTES_OUT"></a>

NETWORK_BYTES_OUT.

---


##### `NETWORK_NUM_REQUESTS` <a name="NETWORK_NUM_REQUESTS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NETWORK_NUM_REQUESTS"></a>

NETWORK_NUM_REQUESTS.

---


##### `NORMALIZED_FTS_PROCESS_CPU_KERNEL` <a name="NORMALIZED_FTS_PROCESS_CPU_KERNEL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NORMALIZED_FTS_PROCESS_CPU_KERNEL"></a>

NORMALIZED_FTS_PROCESS_CPU_KERNEL.

---


##### `NORMALIZED_FTS_PROCESS_CPU_USER` <a name="NORMALIZED_FTS_PROCESS_CPU_USER" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NORMALIZED_FTS_PROCESS_CPU_USER"></a>

NORMALIZED_FTS_PROCESS_CPU_USER.

---


##### `NORMALIZED_SYSTEM_CPU_STEAL` <a name="NORMALIZED_SYSTEM_CPU_STEAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NORMALIZED_SYSTEM_CPU_STEAL"></a>

NORMALIZED_SYSTEM_CPU_STEAL.

---


##### `NORMALIZED_SYSTEM_CPU_USER` <a name="NORMALIZED_SYSTEM_CPU_USER" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.NORMALIZED_SYSTEM_CPU_USER"></a>

NORMALIZED_SYSTEM_CPU_USER.

---


##### `OPCOUNTER_CMD` <a name="OPCOUNTER_CMD" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_CMD"></a>

OPCOUNTER_CMD.

---


##### `OPCOUNTER_DELETE` <a name="OPCOUNTER_DELETE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_DELETE"></a>

OPCOUNTER_DELETE.

---


##### `OPCOUNTER_GETMORE` <a name="OPCOUNTER_GETMORE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_GETMORE"></a>

OPCOUNTER_GETMORE.

---


##### `OPCOUNTER_INSERT` <a name="OPCOUNTER_INSERT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_INSERT"></a>

OPCOUNTER_INSERT.

---


##### `OPCOUNTER_QUERY` <a name="OPCOUNTER_QUERY" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_QUERY"></a>

OPCOUNTER_QUERY.

---


##### `OPCOUNTER_REPL_CMD` <a name="OPCOUNTER_REPL_CMD" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_REPL_CMD"></a>

OPCOUNTER_REPL_CMD.

---


##### `OPCOUNTER_REPL_DELETE` <a name="OPCOUNTER_REPL_DELETE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_REPL_DELETE"></a>

OPCOUNTER_REPL_DELETE.

---


##### `OPCOUNTER_REPL_INSERT` <a name="OPCOUNTER_REPL_INSERT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_REPL_INSERT"></a>

OPCOUNTER_REPL_INSERT.

---


##### `OPCOUNTER_REPL_UPDATE` <a name="OPCOUNTER_REPL_UPDATE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_REPL_UPDATE"></a>

OPCOUNTER_REPL_UPDATE.

---


##### `OPCOUNTER_UPDATE` <a name="OPCOUNTER_UPDATE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPCOUNTER_UPDATE"></a>

OPCOUNTER_UPDATE.

---


##### `OPERATIONS_SCAN_AND_ORDER` <a name="OPERATIONS_SCAN_AND_ORDER" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPERATIONS_SCAN_AND_ORDER"></a>

OPERATIONS_SCAN_AND_ORDER.

---


##### `OPLOG_MASTER_LAG_TIME_DIFF` <a name="OPLOG_MASTER_LAG_TIME_DIFF" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_MASTER_LAG_TIME_DIFF"></a>

OPLOG_MASTER_LAG_TIME_DIFF.

---


##### `OPLOG_MASTER_TIME` <a name="OPLOG_MASTER_TIME" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_MASTER_TIME"></a>

OPLOG_MASTER_TIME.

---


##### `OPLOG_MASTER_TIME_ESTIMATED_TTL` <a name="OPLOG_MASTER_TIME_ESTIMATED_TTL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_MASTER_TIME_ESTIMATED_TTL"></a>

OPLOG_MASTER_TIME_ESTIMATED_TTL.

---


##### `OPLOG_RATE_GB_PER_HOUR` <a name="OPLOG_RATE_GB_PER_HOUR" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_RATE_GB_PER_HOUR"></a>

OPLOG_RATE_GB_PER_HOUR.

---


##### `OPLOG_SLAVE_LAG_MASTER_TIME` <a name="OPLOG_SLAVE_LAG_MASTER_TIME" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.OPLOG_SLAVE_LAG_MASTER_TIME"></a>

OPLOG_SLAVE_LAG_MASTER_TIME.

---


##### `QUERY_EXECUTOR_SCANNED` <a name="QUERY_EXECUTOR_SCANNED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.QUERY_EXECUTOR_SCANNED"></a>

QUERY_EXECUTOR_SCANNED.

---


##### `QUERY_EXECUTOR_SCANNED_OBJECTS` <a name="QUERY_EXECUTOR_SCANNED_OBJECTS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.QUERY_EXECUTOR_SCANNED_OBJECTS"></a>

QUERY_EXECUTOR_SCANNED_OBJECTS.

---


##### `QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED` <a name="QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED"></a>

QUERY_TARGETING_SCANNED_OBJECTS_PER_RETURNED.

---


##### `QUERY_TARGETING_SCANNED_PER_RETURNED` <a name="QUERY_TARGETING_SCANNED_PER_RETURNED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.QUERY_TARGETING_SCANNED_PER_RETURNED"></a>

QUERY_TARGETING_SCANNED_PER_RETURNED.

---


##### `RESTARTS_IN_LAST_HOUR` <a name="RESTARTS_IN_LAST_HOUR" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.RESTARTS_IN_LAST_HOUR"></a>

RESTARTS_IN_LAST_HOUR.

---


##### `SERVERLESS_CONNECTIONS` <a name="SERVERLESS_CONNECTIONS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_CONNECTIONS"></a>

SERVERLESS_CONNECTIONS.

---


##### `SERVERLESS_CONNECTIONS_PERCENT` <a name="SERVERLESS_CONNECTIONS_PERCENT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_CONNECTIONS_PERCENT"></a>

SERVERLESS_CONNECTIONS_PERCENT.

---


##### `SERVERLESS_DATA_SIZE_TOTAL` <a name="SERVERLESS_DATA_SIZE_TOTAL" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_DATA_SIZE_TOTAL"></a>

SERVERLESS_DATA_SIZE_TOTAL.

---


##### `SERVERLESS_NETWORK_BYTES_IN` <a name="SERVERLESS_NETWORK_BYTES_IN" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_NETWORK_BYTES_IN"></a>

SERVERLESS_NETWORK_BYTES_IN.

---


##### `SERVERLESS_NETWORK_BYTES_OUT` <a name="SERVERLESS_NETWORK_BYTES_OUT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_NETWORK_BYTES_OUT"></a>

SERVERLESS_NETWORK_BYTES_OUT.

---


##### `SERVERLESS_NETWORK_NUM_REQUESTS` <a name="SERVERLESS_NETWORK_NUM_REQUESTS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_NETWORK_NUM_REQUESTS"></a>

SERVERLESS_NETWORK_NUM_REQUESTS.

---


##### `SERVERLESS_OPCOUNTER_CMD` <a name="SERVERLESS_OPCOUNTER_CMD" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_CMD"></a>

SERVERLESS_OPCOUNTER_CMD.

---


##### `SERVERLESS_OPCOUNTER_DELETE` <a name="SERVERLESS_OPCOUNTER_DELETE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_DELETE"></a>

SERVERLESS_OPCOUNTER_DELETE.

---


##### `SERVERLESS_OPCOUNTER_GETMORE` <a name="SERVERLESS_OPCOUNTER_GETMORE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_GETMORE"></a>

SERVERLESS_OPCOUNTER_GETMORE.

---


##### `SERVERLESS_OPCOUNTER_INSERT` <a name="SERVERLESS_OPCOUNTER_INSERT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_INSERT"></a>

SERVERLESS_OPCOUNTER_INSERT.

---


##### `SERVERLESS_OPCOUNTER_QUERY` <a name="SERVERLESS_OPCOUNTER_QUERY" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_QUERY"></a>

SERVERLESS_OPCOUNTER_QUERY.

---


##### `SERVERLESS_OPCOUNTER_UPDATE` <a name="SERVERLESS_OPCOUNTER_UPDATE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_OPCOUNTER_UPDATE"></a>

SERVERLESS_OPCOUNTER_UPDATE.

---


##### `SERVERLESS_TOTAL_READ_UNITS` <a name="SERVERLESS_TOTAL_READ_UNITS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_TOTAL_READ_UNITS"></a>

SERVERLESS_TOTAL_READ_UNITS.

---


##### `SERVERLESS_TOTAL_WRITE_UNITS` <a name="SERVERLESS_TOTAL_WRITE_UNITS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SERVERLESS_TOTAL_WRITE_UNITS"></a>

SERVERLESS_TOTAL_WRITE_UNITS.

---


##### `SWAP_USAGE_FREE` <a name="SWAP_USAGE_FREE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SWAP_USAGE_FREE"></a>

SWAP_USAGE_FREE.

---


##### `SWAP_USAGE_USED` <a name="SWAP_USAGE_USED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SWAP_USAGE_USED"></a>

SWAP_USAGE_USED.

---


##### `SYSTEM_MEMORY_AVAILABLE` <a name="SYSTEM_MEMORY_AVAILABLE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_MEMORY_AVAILABLE"></a>

SYSTEM_MEMORY_AVAILABLE.

---


##### `SYSTEM_MEMORY_FREE` <a name="SYSTEM_MEMORY_FREE" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_MEMORY_FREE"></a>

SYSTEM_MEMORY_FREE.

---


##### `SYSTEM_MEMORY_USED` <a name="SYSTEM_MEMORY_USED" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_MEMORY_USED"></a>

SYSTEM_MEMORY_USED.

---


##### `SYSTEM_NETWORK_IN` <a name="SYSTEM_NETWORK_IN" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_NETWORK_IN"></a>

SYSTEM_NETWORK_IN.

---


##### `SYSTEM_NETWORK_OUT` <a name="SYSTEM_NETWORK_OUT" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.SYSTEM_NETWORK_OUT"></a>

SYSTEM_NETWORK_OUT.

---


##### `TICKETS_AVAILABLE_READS` <a name="TICKETS_AVAILABLE_READS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.TICKETS_AVAILABLE_READS"></a>

TICKETS_AVAILABLE_READS.

---


##### `TICKETS_AVAILABLE_WRITES` <a name="TICKETS_AVAILABLE_WRITES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewMetricName.TICKETS_AVAILABLE_WRITES"></a>

TICKETS_AVAILABLE_WRITES.

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


### MetricThresholdViewUnits <a name="MetricThresholdViewUnits" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits"></a>

Element used to express the quantity.

This can be an element of time, storage capacity, and the like.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.BITS">BITS</a></code> | BITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.BYTES">BYTES</a></code> | BYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.DAYS">DAYS</a></code> | DAYS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.GIGABITS">GIGABITS</a></code> | GIGABITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.GIGABYTES">GIGABYTES</a></code> | GIGABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.HOURS">HOURS</a></code> | HOURS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.KILOBITS">KILOBITS</a></code> | KILOBITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.KILOBYTES">KILOBYTES</a></code> | KILOBYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.MEGABITS">MEGABITS</a></code> | MEGABITS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.MEGABYTES">MEGABYTES</a></code> | MEGABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.MILLISECONDS">MILLISECONDS</a></code> | MILLISECONDS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.MINUTES">MINUTES</a></code> | MINUTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.PETABYTES">PETABYTES</a></code> | PETABYTES. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.RAW">RAW</a></code> | RAW. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.SECONDS">SECONDS</a></code> | SECONDS. |
| <code><a href="#@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.TERABYTES">TERABYTES</a></code> | TERABYTES. |

---

##### `BITS` <a name="BITS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.BITS"></a>

BITS.

---


##### `BYTES` <a name="BYTES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.BYTES"></a>

BYTES.

---


##### `DAYS` <a name="DAYS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.DAYS"></a>

DAYS.

---


##### `GIGABITS` <a name="GIGABITS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.GIGABITS"></a>

GIGABITS.

---


##### `GIGABYTES` <a name="GIGABYTES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.GIGABYTES"></a>

GIGABYTES.

---


##### `HOURS` <a name="HOURS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.HOURS"></a>

HOURS.

---


##### `KILOBITS` <a name="KILOBITS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.KILOBITS"></a>

KILOBITS.

---


##### `KILOBYTES` <a name="KILOBYTES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.KILOBYTES"></a>

KILOBYTES.

---


##### `MEGABITS` <a name="MEGABITS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.MEGABITS"></a>

MEGABITS.

---


##### `MEGABYTES` <a name="MEGABYTES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.MEGABYTES"></a>

MEGABYTES.

---


##### `MILLISECONDS` <a name="MILLISECONDS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.MILLISECONDS"></a>

MILLISECONDS.

---


##### `MINUTES` <a name="MINUTES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.MINUTES"></a>

MINUTES.

---


##### `PETABYTES` <a name="PETABYTES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.PETABYTES"></a>

PETABYTES.

---


##### `RAW` <a name="RAW" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.RAW"></a>

RAW.

---


##### `SECONDS` <a name="SECONDS" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.SECONDS"></a>

SECONDS.

---


##### `TERABYTES` <a name="TERABYTES" id="@mongodbatlas-awscdk/alert-configuration.MetricThresholdViewUnits.TERABYTES"></a>

TERABYTES.

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