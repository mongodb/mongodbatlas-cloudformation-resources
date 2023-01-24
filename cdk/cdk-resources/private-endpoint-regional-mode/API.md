# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnPrivateEndPointRegionalMode <a name="CfnPrivateEndPointRegionalMode" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode"></a>

A CloudFormation `MongoDB::Atlas::PrivateEndPointRegionalMode`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.Initializer"></a>

```typescript
import { CfnPrivateEndPointRegionalMode } from '@mongodbatlas-awscdk/private-endpoint-regional-mode'

new CfnPrivateEndPointRegionalMode(scope: Construct, id: string, props: CfnPrivateEndPointRegionalModeProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps">CfnPrivateEndPointRegionalModeProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps">CfnPrivateEndPointRegionalModeProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isConstruct"></a>

```typescript
import { CfnPrivateEndPointRegionalMode } from '@mongodbatlas-awscdk/private-endpoint-regional-mode'

CfnPrivateEndPointRegionalMode.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isCfnElement"></a>

```typescript
import { CfnPrivateEndPointRegionalMode } from '@mongodbatlas-awscdk/private-endpoint-regional-mode'

CfnPrivateEndPointRegionalMode.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isCfnResource"></a>

```typescript
import { CfnPrivateEndPointRegionalMode } from '@mongodbatlas-awscdk/private-endpoint-regional-mode'

CfnPrivateEndPointRegionalMode.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps">CfnPrivateEndPointRegionalModeProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.props"></a>

```typescript
public readonly props: CfnPrivateEndPointRegionalModeProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps">CfnPrivateEndPointRegionalModeProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalMode.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKey <a name="ApiKey" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.ApiKey"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.ApiKey.Initializer"></a>

```typescript
import { ApiKey } from '@mongodbatlas-awscdk/private-endpoint-regional-mode'

const apiKey: ApiKey = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.ApiKey.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.ApiKey.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.ApiKey.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.ApiKey.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnPrivateEndPointRegionalModeProps <a name="CfnPrivateEndPointRegionalModeProps" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps"></a>

An example resource schema demonstrating some basic constructs and validation rules.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps.Initializer"></a>

```typescript
import { CfnPrivateEndPointRegionalModeProps } from '@mongodbatlas-awscdk/private-endpoint-regional-mode'

const cfnPrivateEndPointRegionalModeProps: CfnPrivateEndPointRegionalModeProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.ApiKey">ApiKey</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps.property.enabled">enabled</a></code> | <code>boolean</code> | Flag that indicates whether someone enabled the regionalized private endpoint setting for the specified project. |

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKey;
```

- *Type:* <a href="#@mongodbatlas-awscdk/private-endpoint-regional-mode.ApiKey">ApiKey</a>

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodbatlas-awscdk/private-endpoint-regional-mode.CfnPrivateEndPointRegionalModeProps.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

Flag that indicates whether someone enabled the regionalized private endpoint setting for the specified project.

---



