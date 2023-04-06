# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnFederatedSettingOrgConfigs <a name="CfnFederatedSettingOrgConfigs" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs"></a>

A CloudFormation `MongoDB::Atlas::FederatedSettingOrgConfigs`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.Initializer"></a>

```typescript
import { CfnFederatedSettingOrgConfigs } from '@mongodbatlas-awscdk/federated-setting-org-configs'

new CfnFederatedSettingOrgConfigs(scope: Construct, id: string, props: CfnFederatedSettingOrgConfigsProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps">CfnFederatedSettingOrgConfigsProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps">CfnFederatedSettingOrgConfigsProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isConstruct"></a>

```typescript
import { CfnFederatedSettingOrgConfigs } from '@mongodbatlas-awscdk/federated-setting-org-configs'

CfnFederatedSettingOrgConfigs.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isCfnElement"></a>

```typescript
import { CfnFederatedSettingOrgConfigs } from '@mongodbatlas-awscdk/federated-setting-org-configs'

CfnFederatedSettingOrgConfigs.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isCfnResource"></a>

```typescript
import { CfnFederatedSettingOrgConfigs } from '@mongodbatlas-awscdk/federated-setting-org-configs'

CfnFederatedSettingOrgConfigs.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.attrFederationSettingsId">attrFederationSettingsId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::FederatedSettingOrgConfigs.FederationSettingsId`. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.attrOrgId">attrOrgId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::FederatedSettingOrgConfigs.OrgId`. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.attrTestMode">attrTestMode</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::FederatedSettingOrgConfigs.TestMode`. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps">CfnFederatedSettingOrgConfigsProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrFederationSettingsId`<sup>Required</sup> <a name="attrFederationSettingsId" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.attrFederationSettingsId"></a>

```typescript
public readonly attrFederationSettingsId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::FederatedSettingOrgConfigs.FederationSettingsId`.

---

##### `attrOrgId`<sup>Required</sup> <a name="attrOrgId" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.attrOrgId"></a>

```typescript
public readonly attrOrgId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::FederatedSettingOrgConfigs.OrgId`.

---

##### `attrTestMode`<sup>Required</sup> <a name="attrTestMode" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.attrTestMode"></a>

```typescript
public readonly attrTestMode: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::FederatedSettingOrgConfigs.TestMode`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.props"></a>

```typescript
public readonly props: CfnFederatedSettingOrgConfigsProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps">CfnFederatedSettingOrgConfigsProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigs.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/federated-setting-org-configs'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnFederatedSettingOrgConfigsProps <a name="CfnFederatedSettingOrgConfigsProps" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps"></a>

Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.Initializer"></a>

```typescript
import { CfnFederatedSettingOrgConfigsProps } from '@mongodbatlas-awscdk/federated-setting-org-configs'

const cfnFederatedSettingOrgConfigsProps: CfnFederatedSettingOrgConfigsProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.domainAllowList">domainAllowList</a></code> | <code>string[]</code> | Approved domains that restrict users who can join the organization based on their email address. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.domainRestrictionEnabled">domainRestrictionEnabled</a></code> | <code>boolean</code> | Value that indicates whether domain restriction is enabled for this connected org. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.identityProviderId">identityProviderId</a></code> | <code>string</code> | Unique 20-hexadecimal digit string that identifies the identity provider that this connected org config is associated with. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.postAuthRoleGrants">postAuthRoleGrants</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants">CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants</a>[]</code> | Atlas roles that are granted to a user in this organization after authenticating. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.roleMappings">roleMappings</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView">RoleMappingView</a>[]</code> | Role mappings that are configured in this organization. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.userConflicts">userConflicts</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView">FederatedUserView</a>[]</code> | List that contains the users who have an email address that doesn't match any domain on the allowed list. |

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `domainAllowList`<sup>Optional</sup> <a name="domainAllowList" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.domainAllowList"></a>

```typescript
public readonly domainAllowList: string[];
```

- *Type:* string[]

Approved domains that restrict users who can join the organization based on their email address.

---

##### `domainRestrictionEnabled`<sup>Optional</sup> <a name="domainRestrictionEnabled" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.domainRestrictionEnabled"></a>

```typescript
public readonly domainRestrictionEnabled: boolean;
```

- *Type:* boolean

Value that indicates whether domain restriction is enabled for this connected org.

---

##### `identityProviderId`<sup>Optional</sup> <a name="identityProviderId" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.identityProviderId"></a>

```typescript
public readonly identityProviderId: string;
```

- *Type:* string

Unique 20-hexadecimal digit string that identifies the identity provider that this connected org config is associated with.

---

##### `postAuthRoleGrants`<sup>Optional</sup> <a name="postAuthRoleGrants" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.postAuthRoleGrants"></a>

```typescript
public readonly postAuthRoleGrants: CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants">CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants</a>[]

Atlas roles that are granted to a user in this organization after authenticating.

---

##### `roleMappings`<sup>Optional</sup> <a name="roleMappings" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.roleMappings"></a>

```typescript
public readonly roleMappings: RoleMappingView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView">RoleMappingView</a>[]

Role mappings that are configured in this organization.

---

##### `userConflicts`<sup>Optional</sup> <a name="userConflicts" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsProps.property.userConflicts"></a>

```typescript
public readonly userConflicts: FederatedUserView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView">FederatedUserView</a>[]

List that contains the users who have an email address that doesn't match any domain on the allowed list.

---

### FederatedUserView <a name="FederatedUserView" id="@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.Initializer"></a>

```typescript
import { FederatedUserView } from '@mongodbatlas-awscdk/federated-setting-org-configs'

const federatedUserView: FederatedUserView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.emailAddress">emailAddress</a></code> | <code>string</code> | Email address of the MongoDB Cloud user linked to the federated organization. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.federationSettingsId">federationSettingsId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the federation to which this MongoDB Cloud user belongs. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.firstName">firstName</a></code> | <code>string</code> | First or given name that belongs to the MongoDB Cloud user. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.lastName">lastName</a></code> | <code>string</code> | Last name, family name, or surname that belongs to the MongoDB Cloud user. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.userId">userId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies this user. |

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `emailAddress`<sup>Optional</sup> <a name="emailAddress" id="@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.emailAddress"></a>

```typescript
public readonly emailAddress: string;
```

- *Type:* string

Email address of the MongoDB Cloud user linked to the federated organization.

---

##### `federationSettingsId`<sup>Optional</sup> <a name="federationSettingsId" id="@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.federationSettingsId"></a>

```typescript
public readonly federationSettingsId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the federation to which this MongoDB Cloud user belongs.

---

##### `firstName`<sup>Optional</sup> <a name="firstName" id="@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.firstName"></a>

```typescript
public readonly firstName: string;
```

- *Type:* string

First or given name that belongs to the MongoDB Cloud user.

---

##### `lastName`<sup>Optional</sup> <a name="lastName" id="@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.lastName"></a>

```typescript
public readonly lastName: string;
```

- *Type:* string

Last name, family name, or surname that belongs to the MongoDB Cloud user.

---

##### `userId`<sup>Optional</sup> <a name="userId" id="@mongodbatlas-awscdk/federated-setting-org-configs.FederatedUserView.property.userId"></a>

```typescript
public readonly userId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies this user.

---

### RoleAssignment <a name="RoleAssignment" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment.Initializer"></a>

```typescript
import { RoleAssignment } from '@mongodbatlas-awscdk/federated-setting-org-configs'

const roleAssignment: RoleAssignment = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment.property.groupId">groupId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment.property.orgId">orgId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment.property.role">role</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole">RoleAssignmentRole</a></code> | *No description.* |

---

##### `groupId`<sup>Optional</sup> <a name="groupId" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment.property.groupId"></a>

```typescript
public readonly groupId: string;
```

- *Type:* string

---

##### `orgId`<sup>Optional</sup> <a name="orgId" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment.property.orgId"></a>

```typescript
public readonly orgId: string;
```

- *Type:* string

---

##### `role`<sup>Optional</sup> <a name="role" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment.property.role"></a>

```typescript
public readonly role: RoleAssignmentRole;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole">RoleAssignmentRole</a>

---

### RoleMappingView <a name="RoleMappingView" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView.Initializer"></a>

```typescript
import { RoleMappingView } from '@mongodbatlas-awscdk/federated-setting-org-configs'

const roleMappingView: RoleMappingView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView.property.externalGroupName">externalGroupName</a></code> | <code>string</code> | Unique human-readable label that identifies the identity provider group to whichthis role mapping applies. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView.property.id">id</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies this role mapping. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView.property.roleAssignments">roleAssignments</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment">RoleAssignment</a>[]</code> | Atlas roles and the unique identifiers of the groups and organizations associated with each role. |

---

##### `externalGroupName`<sup>Optional</sup> <a name="externalGroupName" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView.property.externalGroupName"></a>

```typescript
public readonly externalGroupName: string;
```

- *Type:* string

Unique human-readable label that identifies the identity provider group to whichthis role mapping applies.

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies this role mapping.

---

##### `roleAssignments`<sup>Optional</sup> <a name="roleAssignments" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleMappingView.property.roleAssignments"></a>

```typescript
public readonly roleAssignments: RoleAssignment[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignment">RoleAssignment</a>[]

Atlas roles and the unique identifiers of the groups and organizations associated with each role.

---



## Enums <a name="Enums" id="Enums"></a>

### CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants <a name="CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_AUTOMATION_ADMIN">GLOBAL_AUTOMATION_ADMIN</a></code> | GLOBAL_AUTOMATION_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_BACKUP_ADMIN">GLOBAL_BACKUP_ADMIN</a></code> | GLOBAL_BACKUP_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_METERING_USER">GLOBAL_METERING_USER</a></code> | GLOBAL_METERING_USER. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_METRICS_INTERNAL_USER">GLOBAL_METRICS_INTERNAL_USER</a></code> | GLOBAL_METRICS_INTERNAL_USER. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_MONITORING_ADMIN">GLOBAL_MONITORING_ADMIN</a></code> | GLOBAL_MONITORING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_OWNER">GLOBAL_OWNER</a></code> | GLOBAL_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.ORG_OWNER">ORG_OWNER</a></code> | ORG_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.ORG_MEMBER">ORG_MEMBER</a></code> | ORG_MEMBER. |

---

##### `GLOBAL_AUTOMATION_ADMIN` <a name="GLOBAL_AUTOMATION_ADMIN" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_AUTOMATION_ADMIN"></a>

GLOBAL_AUTOMATION_ADMIN.

---


##### `GLOBAL_BACKUP_ADMIN` <a name="GLOBAL_BACKUP_ADMIN" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_BACKUP_ADMIN"></a>

GLOBAL_BACKUP_ADMIN.

---


##### `GLOBAL_METERING_USER` <a name="GLOBAL_METERING_USER" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_METERING_USER"></a>

GLOBAL_METERING_USER.

---


##### `GLOBAL_METRICS_INTERNAL_USER` <a name="GLOBAL_METRICS_INTERNAL_USER" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_METRICS_INTERNAL_USER"></a>

GLOBAL_METRICS_INTERNAL_USER.

---


##### `GLOBAL_MONITORING_ADMIN` <a name="GLOBAL_MONITORING_ADMIN" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_MONITORING_ADMIN"></a>

GLOBAL_MONITORING_ADMIN.

---


##### `GLOBAL_OWNER` <a name="GLOBAL_OWNER" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.GLOBAL_OWNER"></a>

GLOBAL_OWNER.

---


##### `ORG_OWNER` <a name="ORG_OWNER" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.ORG_OWNER"></a>

ORG_OWNER.

---


##### `ORG_MEMBER` <a name="ORG_MEMBER" id="@mongodbatlas-awscdk/federated-setting-org-configs.CfnFederatedSettingOrgConfigsPropsPostAuthRoleGrants.ORG_MEMBER"></a>

ORG_MEMBER.

---


### RoleAssignmentRole <a name="RoleAssignmentRole" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_AUTOMATION_ADMIN">GLOBAL_AUTOMATION_ADMIN</a></code> | GLOBAL_AUTOMATION_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_BACKUP_ADMIN">GLOBAL_BACKUP_ADMIN</a></code> | GLOBAL_BACKUP_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_METERING_USER">GLOBAL_METERING_USER</a></code> | GLOBAL_METERING_USER. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_METRICS_INTERNAL_USER">GLOBAL_METRICS_INTERNAL_USER</a></code> | GLOBAL_METRICS_INTERNAL_USER. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_MONITORING_ADMIN">GLOBAL_MONITORING_ADMIN</a></code> | GLOBAL_MONITORING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_OWNER">GLOBAL_OWNER</a></code> | GLOBAL_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_READ_ONLY">GLOBAL_READ_ONLY</a></code> | GLOBAL_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.ORG_OWNER">ORG_OWNER</a></code> | ORG_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.ORG_MEMBER">ORG_MEMBER</a></code> | ORG_MEMBER. |

---

##### `GLOBAL_AUTOMATION_ADMIN` <a name="GLOBAL_AUTOMATION_ADMIN" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_AUTOMATION_ADMIN"></a>

GLOBAL_AUTOMATION_ADMIN.

---


##### `GLOBAL_BACKUP_ADMIN` <a name="GLOBAL_BACKUP_ADMIN" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_BACKUP_ADMIN"></a>

GLOBAL_BACKUP_ADMIN.

---


##### `GLOBAL_METERING_USER` <a name="GLOBAL_METERING_USER" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_METERING_USER"></a>

GLOBAL_METERING_USER.

---


##### `GLOBAL_METRICS_INTERNAL_USER` <a name="GLOBAL_METRICS_INTERNAL_USER" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_METRICS_INTERNAL_USER"></a>

GLOBAL_METRICS_INTERNAL_USER.

---


##### `GLOBAL_MONITORING_ADMIN` <a name="GLOBAL_MONITORING_ADMIN" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_MONITORING_ADMIN"></a>

GLOBAL_MONITORING_ADMIN.

---


##### `GLOBAL_OWNER` <a name="GLOBAL_OWNER" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_OWNER"></a>

GLOBAL_OWNER.

---


##### `GLOBAL_READ_ONLY` <a name="GLOBAL_READ_ONLY" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.GLOBAL_READ_ONLY"></a>

GLOBAL_READ_ONLY.

---


##### `ORG_OWNER` <a name="ORG_OWNER" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.ORG_OWNER"></a>

ORG_OWNER.

---


##### `ORG_MEMBER` <a name="ORG_MEMBER" id="@mongodbatlas-awscdk/federated-setting-org-configs.RoleAssignmentRole.ORG_MEMBER"></a>

ORG_MEMBER.

---

