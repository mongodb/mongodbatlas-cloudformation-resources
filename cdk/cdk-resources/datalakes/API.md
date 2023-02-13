# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnDataLakes <a name="CfnDataLakes" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes"></a>

A CloudFormation `MongoDB::Atlas::DataLakes`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.Initializer"></a>

```typescript
import { CfnDataLakes } from '@mongodbatlas-awscdk/datalakes'

new CfnDataLakes(scope: Construct, id: string, props: CfnDataLakesProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps">CfnDataLakesProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps">CfnDataLakesProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.isConstruct"></a>

```typescript
import { CfnDataLakes } from '@mongodbatlas-awscdk/datalakes'

CfnDataLakes.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.isCfnElement"></a>

```typescript
import { CfnDataLakes } from '@mongodbatlas-awscdk/datalakes'

CfnDataLakes.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.isCfnResource"></a>

```typescript
import { CfnDataLakes } from '@mongodbatlas-awscdk/datalakes'

CfnDataLakes.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrEndDate">attrEndDate</a></code> | <code>number</code> | Attribute `MongoDB::Atlas::DataLakes.EndDate`. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrGroupId">attrGroupId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::DataLakes.GroupId`. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrHostnames">attrHostnames</a></code> | <code>string[]</code> | Attribute `MongoDB::Atlas::DataLakes.Hostnames`. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrStartDate">attrStartDate</a></code> | <code>number</code> | Attribute `MongoDB::Atlas::DataLakes.StartDate`. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrState">attrState</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::DataLakes.State`. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrTenantName">attrTenantName</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::DataLakes.TenantName`. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps">CfnDataLakesProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrEndDate`<sup>Required</sup> <a name="attrEndDate" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrEndDate"></a>

```typescript
public readonly attrEndDate: number;
```

- *Type:* number

Attribute `MongoDB::Atlas::DataLakes.EndDate`.

---

##### `attrGroupId`<sup>Required</sup> <a name="attrGroupId" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrGroupId"></a>

```typescript
public readonly attrGroupId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::DataLakes.GroupId`.

---

##### `attrHostnames`<sup>Required</sup> <a name="attrHostnames" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrHostnames"></a>

```typescript
public readonly attrHostnames: string[];
```

- *Type:* string[]

Attribute `MongoDB::Atlas::DataLakes.Hostnames`.

---

##### `attrStartDate`<sup>Required</sup> <a name="attrStartDate" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrStartDate"></a>

```typescript
public readonly attrStartDate: number;
```

- *Type:* number

Attribute `MongoDB::Atlas::DataLakes.StartDate`.

---

##### `attrState`<sup>Required</sup> <a name="attrState" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrState"></a>

```typescript
public readonly attrState: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::DataLakes.State`.

---

##### `attrTenantName`<sup>Required</sup> <a name="attrTenantName" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.attrTenantName"></a>

```typescript
public readonly attrTenantName: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::DataLakes.TenantName`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.props"></a>

```typescript
public readonly props: CfnDataLakesProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps">CfnDataLakesProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/datalakes.CfnDataLakes.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/datalakes.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/datalakes'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/datalakes.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/datalakes.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/datalakes.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnDataLakesProps <a name="CfnDataLakesProps" id="@mongodbatlas-awscdk/datalakes.CfnDataLakesProps"></a>

Returns, adds, edits, and removes Federated Database Instances.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.Initializer"></a>

```typescript
import { CfnDataLakesProps } from '@mongodbatlas-awscdk/datalakes'

const cfnDataLakesProps: CfnDataLakesProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.cloudProviderConfig">cloudProviderConfig</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeCloudProviderConfigView">DataLakeCloudProviderConfigView</a></code> | Cloud provider linked to this data lake. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.dataProcessRegion">dataProcessRegion</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionView">DataLakeDataProcessRegionView</a></code> | Information about the cloud provider region to which the data lake routes client connections. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.skipRoleValidation">skipRoleValidation</a></code> | <code>boolean</code> | Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.storage">storage</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeStorageView">DataLakeStorageView</a></code> | Configuration information for each data store and its mapping to MongoDB Cloud databases. |

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `cloudProviderConfig`<sup>Optional</sup> <a name="cloudProviderConfig" id="@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.cloudProviderConfig"></a>

```typescript
public readonly cloudProviderConfig: DataLakeCloudProviderConfigView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeCloudProviderConfigView">DataLakeCloudProviderConfigView</a>

Cloud provider linked to this data lake.

---

##### `dataProcessRegion`<sup>Optional</sup> <a name="dataProcessRegion" id="@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.dataProcessRegion"></a>

```typescript
public readonly dataProcessRegion: DataLakeDataProcessRegionView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionView">DataLakeDataProcessRegionView</a>

Information about the cloud provider region to which the data lake routes client connections.

MongoDB Cloud supports AWS only.

---

##### `skipRoleValidation`<sup>Optional</sup> <a name="skipRoleValidation" id="@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.skipRoleValidation"></a>

```typescript
public readonly skipRoleValidation: boolean;
```

- *Type:* boolean

Flag that indicates whether this request should check if the requesting IAM role can read from the S3 bucket.

AWS checks if the role can list the objects in the bucket before writing to it. Some IAM roles only need write permissions. This flag allows you to skip that check.

---

##### `storage`<sup>Optional</sup> <a name="storage" id="@mongodbatlas-awscdk/datalakes.CfnDataLakesProps.property.storage"></a>

```typescript
public readonly storage: DataLakeStorageView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeStorageView">DataLakeStorageView</a>

Configuration information for each data store and its mapping to MongoDB Cloud databases.

---

### DataLakeAwsCloudProviderConfigView <a name="DataLakeAwsCloudProviderConfigView" id="@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.Initializer"></a>

```typescript
import { DataLakeAwsCloudProviderConfigView } from '@mongodbatlas-awscdk/datalakes'

const dataLakeAwsCloudProviderConfigView: DataLakeAwsCloudProviderConfigView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.externalId">externalId</a></code> | <code>string</code> | Unique identifier associated with the Identity and Access Management (IAM) role that the data lake assumes when accessing the data stores. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.iamAssumedRoleArn">iamAssumedRoleArn</a></code> | <code>string</code> | Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the data lake assumes when accessing data stores. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.iamUserArn">iamUserArn</a></code> | <code>string</code> | Amazon Resource Name (ARN) of the user that the data lake assumes when accessing data stores. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.roleId">roleId</a></code> | <code>string</code> | Unique identifier of the role that the data lake can use to access the data stores.Required if specifying cloudProviderConfig. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.testS3Bucket">testS3Bucket</a></code> | <code>string</code> | Name of the S3 data bucket that the provided role ID is authorized to access.Required if specifying cloudProviderConfig. |

---

##### `externalId`<sup>Optional</sup> <a name="externalId" id="@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.externalId"></a>

```typescript
public readonly externalId: string;
```

- *Type:* string

Unique identifier associated with the Identity and Access Management (IAM) role that the data lake assumes when accessing the data stores.

---

##### `iamAssumedRoleArn`<sup>Optional</sup> <a name="iamAssumedRoleArn" id="@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.iamAssumedRoleArn"></a>

```typescript
public readonly iamAssumedRoleArn: string;
```

- *Type:* string

Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role that the data lake assumes when accessing data stores.

---

##### `iamUserArn`<sup>Optional</sup> <a name="iamUserArn" id="@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.iamUserArn"></a>

```typescript
public readonly iamUserArn: string;
```

- *Type:* string

Amazon Resource Name (ARN) of the user that the data lake assumes when accessing data stores.

---

##### `roleId`<sup>Optional</sup> <a name="roleId" id="@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.roleId"></a>

```typescript
public readonly roleId: string;
```

- *Type:* string

Unique identifier of the role that the data lake can use to access the data stores.Required if specifying cloudProviderConfig.

---

##### `testS3Bucket`<sup>Optional</sup> <a name="testS3Bucket" id="@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView.property.testS3Bucket"></a>

```typescript
public readonly testS3Bucket: string;
```

- *Type:* string

Name of the S3 data bucket that the provided role ID is authorized to access.Required if specifying cloudProviderConfig.

---

### DataLakeCloudProviderConfigView <a name="DataLakeCloudProviderConfigView" id="@mongodbatlas-awscdk/datalakes.DataLakeCloudProviderConfigView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.DataLakeCloudProviderConfigView.Initializer"></a>

```typescript
import { DataLakeCloudProviderConfigView } from '@mongodbatlas-awscdk/datalakes'

const dataLakeCloudProviderConfigView: DataLakeCloudProviderConfigView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeCloudProviderConfigView.property.aws">aws</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView">DataLakeAwsCloudProviderConfigView</a></code> | Name of the cloud service that hosts the data lake's data stores. |

---

##### `aws`<sup>Optional</sup> <a name="aws" id="@mongodbatlas-awscdk/datalakes.DataLakeCloudProviderConfigView.property.aws"></a>

```typescript
public readonly aws: DataLakeAwsCloudProviderConfigView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeAwsCloudProviderConfigView">DataLakeAwsCloudProviderConfigView</a>

Name of the cloud service that hosts the data lake's data stores.

---

### DataLakeDatabaseCollectionView <a name="DataLakeDatabaseCollectionView" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseCollectionView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseCollectionView.Initializer"></a>

```typescript
import { DataLakeDatabaseCollectionView } from '@mongodbatlas-awscdk/datalakes'

const dataLakeDatabaseCollectionView: DataLakeDatabaseCollectionView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseCollectionView.property.dataSources">dataSources</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView">DataLakeDatabaseDataSourceView</a>[]</code> | Array that contains the data stores that map to a collection for this data lake. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseCollectionView.property.name">name</a></code> | <code>string</code> | Human-readable label that identifies the collection to which MongoDB Cloud maps the data in the data stores. |

---

##### `dataSources`<sup>Optional</sup> <a name="dataSources" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseCollectionView.property.dataSources"></a>

```typescript
public readonly dataSources: DataLakeDatabaseDataSourceView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView">DataLakeDatabaseDataSourceView</a>[]

Array that contains the data stores that map to a collection for this data lake.

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseCollectionView.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable label that identifies the collection to which MongoDB Cloud maps the data in the data stores.

---

### DataLakeDatabaseDataSourceView <a name="DataLakeDatabaseDataSourceView" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.Initializer"></a>

```typescript
import { DataLakeDatabaseDataSourceView } from '@mongodbatlas-awscdk/datalakes'

const dataLakeDatabaseDataSourceView: DataLakeDatabaseDataSourceView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.allowInsecure">allowInsecure</a></code> | <code>boolean</code> | Flag that validates the scheme in the specified URLs. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.collection">collection</a></code> | <code>string</code> | Human-readable label that identifies the collection in the database. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.collectionRegex">collectionRegex</a></code> | <code>string</code> | Regex pattern to use for creating the wildcard (*) collection. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.database">database</a></code> | <code>string</code> | Human-readable label that identifies the database, which contains the collection in the cluster. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.defaultFormat">defaultFormat</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat">DataLakeDatabaseDataSourceViewDefaultFormat</a></code> | File format that MongoDB Cloud uses if it encounters a file without a file extension while searching **storeName**. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.path">path</a></code> | <code>string</code> | File path that controls how MongoDB Cloud searches for and parses files in the **storeName** before mapping them to a collection.Specify ``/`` to capture all files and folders from the ``prefix`` path. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.storeName">storeName</a></code> | <code>string</code> | Human-readable label that identifies the data store that MongoDB Cloud maps to the collection. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.urls">urls</a></code> | <code>string[]</code> | URLs of the publicly accessible data files. |

---

##### `allowInsecure`<sup>Optional</sup> <a name="allowInsecure" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.allowInsecure"></a>

```typescript
public readonly allowInsecure: boolean;
```

- *Type:* boolean

Flag that validates the scheme in the specified URLs.

If `true`, allows insecure `HTTP` scheme, doesn't verify the server's certificate chain and hostname, and accepts any certificate with any hostname presented by the server. If `false`, allows secure `HTTPS` scheme only.

---

##### `collection`<sup>Optional</sup> <a name="collection" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.collection"></a>

```typescript
public readonly collection: string;
```

- *Type:* string

Human-readable label that identifies the collection in the database.

For creating a wildcard (`*`) collection, you must omit this parameter.

---

##### `collectionRegex`<sup>Optional</sup> <a name="collectionRegex" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.collectionRegex"></a>

```typescript
public readonly collectionRegex: string;
```

- *Type:* string

Regex pattern to use for creating the wildcard (*) collection.

To learn more about the regex syntax, see [Go programming language](https://pkg.go.dev/regexp).

---

##### `database`<sup>Optional</sup> <a name="database" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.database"></a>

```typescript
public readonly database: string;
```

- *Type:* string

Human-readable label that identifies the database, which contains the collection in the cluster.

You must omit this parameter to generate wildcard (`*`) collections for dynamically generated databases.

---

##### `defaultFormat`<sup>Optional</sup> <a name="defaultFormat" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.defaultFormat"></a>

```typescript
public readonly defaultFormat: DataLakeDatabaseDataSourceViewDefaultFormat;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat">DataLakeDatabaseDataSourceViewDefaultFormat</a>

File format that MongoDB Cloud uses if it encounters a file without a file extension while searching **storeName**.

---

##### `path`<sup>Optional</sup> <a name="path" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.path"></a>

```typescript
public readonly path: string;
```

- *Type:* string

File path that controls how MongoDB Cloud searches for and parses files in the **storeName** before mapping them to a collection.Specify ``/`` to capture all files and folders from the ``prefix`` path.

---

##### `storeName`<sup>Optional</sup> <a name="storeName" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.storeName"></a>

```typescript
public readonly storeName: string;
```

- *Type:* string

Human-readable label that identifies the data store that MongoDB Cloud maps to the collection.

---

##### `urls`<sup>Optional</sup> <a name="urls" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceView.property.urls"></a>

```typescript
public readonly urls: string[];
```

- *Type:* string[]

URLs of the publicly accessible data files.

You can't specify URLs that require authentication. Atlas Data Lake creates a partition for each URL. If empty or omitted, Data Lake uses the URLs from the store specified in the **dataSources.storeName** parameter.

---

### DataLakeDatabaseView <a name="DataLakeDatabaseView" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.Initializer"></a>

```typescript
import { DataLakeDatabaseView } from '@mongodbatlas-awscdk/datalakes'

const dataLakeDatabaseView: DataLakeDatabaseView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.property.collections">collections</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseCollectionView">DataLakeDatabaseCollectionView</a>[]</code> | Array of collections and data sources that map to a ``stores`` data store. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.property.maxWildcardCollections">maxWildcardCollections</a></code> | <code>number</code> | Maximum number of wildcard collections in the database. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.property.name">name</a></code> | <code>string</code> | Human-readable label that identifies the database to which the data lake maps data. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.property.views">views</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeViewView">DataLakeViewView</a>[]</code> | Array of aggregation pipelines that apply to the collection. |

---

##### `collections`<sup>Optional</sup> <a name="collections" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.property.collections"></a>

```typescript
public readonly collections: DataLakeDatabaseCollectionView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseCollectionView">DataLakeDatabaseCollectionView</a>[]

Array of collections and data sources that map to a ``stores`` data store.

---

##### `maxWildcardCollections`<sup>Optional</sup> <a name="maxWildcardCollections" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.property.maxWildcardCollections"></a>

```typescript
public readonly maxWildcardCollections: number;
```

- *Type:* number

Maximum number of wildcard collections in the database.

This only applies to S3 data sources.

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable label that identifies the database to which the data lake maps data.

---

##### `views`<sup>Optional</sup> <a name="views" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView.property.views"></a>

```typescript
public readonly views: DataLakeViewView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeViewView">DataLakeViewView</a>[]

Array of aggregation pipelines that apply to the collection.

This only applies to S3 data sources.

---

### DataLakeDataProcessRegionView <a name="DataLakeDataProcessRegionView" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionView.Initializer"></a>

```typescript
import { DataLakeDataProcessRegionView } from '@mongodbatlas-awscdk/datalakes'

const dataLakeDataProcessRegionView: DataLakeDataProcessRegionView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionView.property.cloudProvider">cloudProvider</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider">DataLakeDataProcessRegionViewCloudProvider</a></code> | Name of the cloud service that hosts the data lake's data stores. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionView.property.region">region</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion">DataLakeDataProcessRegionViewRegion</a></code> | Name of the region to which the data lake routes client connections. |

---

##### `cloudProvider`<sup>Optional</sup> <a name="cloudProvider" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionView.property.cloudProvider"></a>

```typescript
public readonly cloudProvider: DataLakeDataProcessRegionViewCloudProvider;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider">DataLakeDataProcessRegionViewCloudProvider</a>

Name of the cloud service that hosts the data lake's data stores.

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionView.property.region"></a>

```typescript
public readonly region: DataLakeDataProcessRegionViewRegion;
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion">DataLakeDataProcessRegionViewRegion</a>

Name of the region to which the data lake routes client connections.

---

### DataLakeStorageView <a name="DataLakeStorageView" id="@mongodbatlas-awscdk/datalakes.DataLakeStorageView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.DataLakeStorageView.Initializer"></a>

```typescript
import { DataLakeStorageView } from '@mongodbatlas-awscdk/datalakes'

const dataLakeStorageView: DataLakeStorageView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeStorageView.property.databases">databases</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView">DataLakeDatabaseView</a>[]</code> | Array that contains the queryable databases and collections for this data lake. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeStorageView.property.stores">stores</a></code> | <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail">StoreDetail</a>[]</code> | Array that contains the data stores for the data lake. |

---

##### `databases`<sup>Optional</sup> <a name="databases" id="@mongodbatlas-awscdk/datalakes.DataLakeStorageView.property.databases"></a>

```typescript
public readonly databases: DataLakeDatabaseView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseView">DataLakeDatabaseView</a>[]

Array that contains the queryable databases and collections for this data lake.

---

##### `stores`<sup>Optional</sup> <a name="stores" id="@mongodbatlas-awscdk/datalakes.DataLakeStorageView.property.stores"></a>

```typescript
public readonly stores: StoreDetail[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/datalakes.StoreDetail">StoreDetail</a>[]

Array that contains the data stores for the data lake.

---

### DataLakeViewView <a name="DataLakeViewView" id="@mongodbatlas-awscdk/datalakes.DataLakeViewView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.DataLakeViewView.Initializer"></a>

```typescript
import { DataLakeViewView } from '@mongodbatlas-awscdk/datalakes'

const dataLakeViewView: DataLakeViewView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeViewView.property.name">name</a></code> | <code>string</code> | Human-readable label that identifies the view, which corresponds to an aggregation pipeline on a collection. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeViewView.property.pipeline">pipeline</a></code> | <code>string</code> | Aggregation pipeline stages to apply to the source collection. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeViewView.property.source">source</a></code> | <code>string</code> | Human-readable label that identifies the source collection for the view. |

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/datalakes.DataLakeViewView.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable label that identifies the view, which corresponds to an aggregation pipeline on a collection.

---

##### `pipeline`<sup>Optional</sup> <a name="pipeline" id="@mongodbatlas-awscdk/datalakes.DataLakeViewView.property.pipeline"></a>

```typescript
public readonly pipeline: string;
```

- *Type:* string

Aggregation pipeline stages to apply to the source collection.

---

##### `source`<sup>Optional</sup> <a name="source" id="@mongodbatlas-awscdk/datalakes.DataLakeViewView.property.source"></a>

```typescript
public readonly source: string;
```

- *Type:* string

Human-readable label that identifies the source collection for the view.

---

### StoreDetail <a name="StoreDetail" id="@mongodbatlas-awscdk/datalakes.StoreDetail"></a>

Configuration information for each data store and its mapping to MongoDB Cloud databases.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/datalakes.StoreDetail.Initializer"></a>

```typescript
import { StoreDetail } from '@mongodbatlas-awscdk/datalakes'

const storeDetail: StoreDetail = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail.property.additionalStorageClasses">additionalStorageClasses</a></code> | <code>string[]</code> | Human-readable label that identifies the Federated Database to update. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail.property.bucket">bucket</a></code> | <code>string</code> | Human-readable label that identifies the Federated Database to update. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail.property.delimiter">delimiter</a></code> | <code>string</code> | Human-readable label that identifies the Federated Database to update. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail.property.includeTags">includeTags</a></code> | <code>boolean</code> | Human-readable label that identifies the Federated Database to update. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail.property.name">name</a></code> | <code>string</code> | Human-readable label that identifies the data store. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail.property.prefix">prefix</a></code> | <code>string</code> | Human-readable label that identifies the Federated Database to update. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail.property.provider">provider</a></code> | <code>string</code> | Human-readable label that identifies the Federated Database to update. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.StoreDetail.property.region">region</a></code> | <code>string</code> | Human-readable label that identifies the Federated Database to update. |

---

##### `additionalStorageClasses`<sup>Optional</sup> <a name="additionalStorageClasses" id="@mongodbatlas-awscdk/datalakes.StoreDetail.property.additionalStorageClasses"></a>

```typescript
public readonly additionalStorageClasses: string[];
```

- *Type:* string[]

Human-readable label that identifies the Federated Database to update.

---

##### `bucket`<sup>Optional</sup> <a name="bucket" id="@mongodbatlas-awscdk/datalakes.StoreDetail.property.bucket"></a>

```typescript
public readonly bucket: string;
```

- *Type:* string

Human-readable label that identifies the Federated Database to update.

---

##### `delimiter`<sup>Optional</sup> <a name="delimiter" id="@mongodbatlas-awscdk/datalakes.StoreDetail.property.delimiter"></a>

```typescript
public readonly delimiter: string;
```

- *Type:* string

Human-readable label that identifies the Federated Database to update.

---

##### `includeTags`<sup>Optional</sup> <a name="includeTags" id="@mongodbatlas-awscdk/datalakes.StoreDetail.property.includeTags"></a>

```typescript
public readonly includeTags: boolean;
```

- *Type:* boolean

Human-readable label that identifies the Federated Database to update.

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/datalakes.StoreDetail.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable label that identifies the data store.

---

##### `prefix`<sup>Optional</sup> <a name="prefix" id="@mongodbatlas-awscdk/datalakes.StoreDetail.property.prefix"></a>

```typescript
public readonly prefix: string;
```

- *Type:* string

Human-readable label that identifies the Federated Database to update.

---

##### `provider`<sup>Optional</sup> <a name="provider" id="@mongodbatlas-awscdk/datalakes.StoreDetail.property.provider"></a>

```typescript
public readonly provider: string;
```

- *Type:* string

Human-readable label that identifies the Federated Database to update.

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/datalakes.StoreDetail.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

Human-readable label that identifies the Federated Database to update.

---



## Enums <a name="Enums" id="Enums"></a>

### DataLakeDatabaseDataSourceViewDefaultFormat <a name="DataLakeDatabaseDataSourceViewDefaultFormat" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat"></a>

File format that MongoDB Cloud uses if it encounters a file without a file extension while searching **storeName**.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.AVRO">AVRO</a></code> | .avro. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.AVRO_GZ">AVRO_GZ</a></code> | .avro.gz. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.BSON">BSON</a></code> | .bson. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.BSON_GZ">BSON_GZ</a></code> | .bson.gz. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.CSV">CSV</a></code> | .csv. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.JSON">JSON</a></code> | .json. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.JSON_GZ">JSON_GZ</a></code> | .json.gz. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.ORC">ORC</a></code> | .orc. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.TSV">TSV</a></code> | .tsv. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.TSV_GZ">TSV_GZ</a></code> | .tsv.gz. |

---

##### `AVRO` <a name="AVRO" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.AVRO"></a>

.avro.

---


##### `AVRO_GZ` <a name="AVRO_GZ" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.AVRO_GZ"></a>

.avro.gz.

---


##### `BSON` <a name="BSON" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.BSON"></a>

.bson.

---


##### `BSON_GZ` <a name="BSON_GZ" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.BSON_GZ"></a>

.bson.gz.

---


##### `CSV` <a name="CSV" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.CSV"></a>

.csv.

---


##### `JSON` <a name="JSON" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.JSON"></a>

.json.

---


##### `JSON_GZ` <a name="JSON_GZ" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.JSON_GZ"></a>

.json.gz.

---


##### `ORC` <a name="ORC" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.ORC"></a>

.orc.

---


##### `TSV` <a name="TSV" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.TSV"></a>

.tsv.

---


##### `TSV_GZ` <a name="TSV_GZ" id="@mongodbatlas-awscdk/datalakes.DataLakeDatabaseDataSourceViewDefaultFormat.TSV_GZ"></a>

.tsv.gz.

---


### DataLakeDataProcessRegionViewCloudProvider <a name="DataLakeDataProcessRegionViewCloudProvider" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider"></a>

Name of the cloud service that hosts the data lake's data stores.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.AWS">AWS</a></code> | AWS. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.GCP">GCP</a></code> | GCP. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.AZURE">AZURE</a></code> | AZURE. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.TENANT">TENANT</a></code> | TENANT. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.SERVERLESS">SERVERLESS</a></code> | SERVERLESS. |

---

##### `AWS` <a name="AWS" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.AWS"></a>

AWS.

---


##### `GCP` <a name="GCP" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.GCP"></a>

GCP.

---


##### `AZURE` <a name="AZURE" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.AZURE"></a>

AZURE.

---


##### `TENANT` <a name="TENANT" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.TENANT"></a>

TENANT.

---


##### `SERVERLESS` <a name="SERVERLESS" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewCloudProvider.SERVERLESS"></a>

SERVERLESS.

---


### DataLakeDataProcessRegionViewRegion <a name="DataLakeDataProcessRegionViewRegion" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion"></a>

Name of the region to which the data lake routes client connections.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.DUBLIN_IRL">DUBLIN_IRL</a></code> | DUBLIN_IRL. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.FRANKFURT_DEU">FRANKFURT_DEU</a></code> | FRANKFURT_DEU. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.LONDON_GBR">LONDON_GBR</a></code> | LONDON_GBR. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.MUMBAI_IND">MUMBAI_IND</a></code> | MUMBAI_IND. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.OREGON_USA">OREGON_USA</a></code> | OREGON_USA. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.SYDNEY_AUS">SYDNEY_AUS</a></code> | SYDNEY_AUS. |
| <code><a href="#@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.VIRGINIA_USA">VIRGINIA_USA</a></code> | VIRGINIA_USA. |

---

##### `DUBLIN_IRL` <a name="DUBLIN_IRL" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.DUBLIN_IRL"></a>

DUBLIN_IRL.

---


##### `FRANKFURT_DEU` <a name="FRANKFURT_DEU" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.FRANKFURT_DEU"></a>

FRANKFURT_DEU.

---


##### `LONDON_GBR` <a name="LONDON_GBR" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.LONDON_GBR"></a>

LONDON_GBR.

---


##### `MUMBAI_IND` <a name="MUMBAI_IND" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.MUMBAI_IND"></a>

MUMBAI_IND.

---


##### `OREGON_USA` <a name="OREGON_USA" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.OREGON_USA"></a>

OREGON_USA.

---


##### `SYDNEY_AUS` <a name="SYDNEY_AUS" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.SYDNEY_AUS"></a>

SYDNEY_AUS.

---


##### `VIRGINIA_USA` <a name="VIRGINIA_USA" id="@mongodbatlas-awscdk/datalakes.DataLakeDataProcessRegionViewRegion.VIRGINIA_USA"></a>

VIRGINIA_USA.

---

