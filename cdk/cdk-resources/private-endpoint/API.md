# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnPrivateEndpoint <a name="CfnPrivateEndpoint" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint"></a>

A CloudFormation `MongoDB::Atlas::PrivateEndpoint`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.Initializer"></a>

```typescript
import { CfnPrivateEndpoint } from '@mongodbatlas-awscdk/private-endpoint'

new CfnPrivateEndpoint(scope: Construct, id: string, props: CfnPrivateEndpointProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps">CfnPrivateEndpointProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps">CfnPrivateEndpointProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isConstruct"></a>

```typescript
import { CfnPrivateEndpoint } from '@mongodbatlas-awscdk/private-endpoint'

CfnPrivateEndpoint.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isCfnElement"></a>

```typescript
import { CfnPrivateEndpoint } from '@mongodbatlas-awscdk/private-endpoint'

CfnPrivateEndpoint.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isCfnResource"></a>

```typescript
import { CfnPrivateEndpoint } from '@mongodbatlas-awscdk/private-endpoint'

CfnPrivateEndpoint.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::PrivateEndpoint.Id`. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps">CfnPrivateEndpointProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::PrivateEndpoint.Id`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.props"></a>

```typescript
public readonly props: CfnPrivateEndpointProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps">CfnPrivateEndpointProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpoint.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKey <a name="ApiKey" id="@mongodbatlas-awscdk/private-endpoint.ApiKey"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/private-endpoint.ApiKey.Initializer"></a>

```typescript
import { ApiKey } from '@mongodbatlas-awscdk/private-endpoint'

const apiKey: ApiKey = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.ApiKey.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.ApiKey.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/private-endpoint.ApiKey.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/private-endpoint.ApiKey.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnPrivateEndpointProps <a name="CfnPrivateEndpointProps" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps"></a>

The Private Endpoint creation flow consists of the creation of three related resources in the next order: 1.

Atlas Private Endpoint Service 2. Aws VPC private Endpoint 3. Atlas Private Endpoint

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.Initializer"></a>

```typescript
import { CfnPrivateEndpointProps } from '@mongodbatlas-awscdk/private-endpoint'

const cfnPrivateEndpointProps: CfnPrivateEndpointProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/private-endpoint.ApiKey">ApiKey</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.groupId">groupId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.region">region</a></code> | <code>string</code> | Aws Region. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.endpointServiceName">endpointServiceName</a></code> | <code>string</code> | Name of the AWS PrivateLink endpoint service. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.errorMessage">errorMessage</a></code> | <code>string</code> | Error message pertaining to the AWS PrivateLink connection. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.privateEndpoints">privateEndpoints</a></code> | <code><a href="#@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint">PrivateEndpoint</a>[]</code> | List of private endpoint associated to the service. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.status">status</a></code> | <code>string</code> | Status of the Atlas PrivateEndpoint service connection. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKey;
```

- *Type:* <a href="#@mongodbatlas-awscdk/private-endpoint.ApiKey">ApiKey</a>

---

##### `groupId`<sup>Required</sup> <a name="groupId" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.groupId"></a>

```typescript
public readonly groupId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `region`<sup>Required</sup> <a name="region" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

Aws Region.

---

##### `endpointServiceName`<sup>Optional</sup> <a name="endpointServiceName" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.endpointServiceName"></a>

```typescript
public readonly endpointServiceName: string;
```

- *Type:* string

Name of the AWS PrivateLink endpoint service.

Atlas returns null while it is creating the endpoint service.

---

##### `errorMessage`<sup>Optional</sup> <a name="errorMessage" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.errorMessage"></a>

```typescript
public readonly errorMessage: string;
```

- *Type:* string

Error message pertaining to the AWS PrivateLink connection.

Returns null if there are no errors.

---

##### `privateEndpoints`<sup>Optional</sup> <a name="privateEndpoints" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.privateEndpoints"></a>

```typescript
public readonly privateEndpoints: PrivateEndpoint[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint">PrivateEndpoint</a>[]

List of private endpoint associated to the service.

---

##### `status`<sup>Optional</sup> <a name="status" id="@mongodbatlas-awscdk/private-endpoint.CfnPrivateEndpointProps.property.status"></a>

```typescript
public readonly status: string;
```

- *Type:* string

Status of the Atlas PrivateEndpoint service connection.

---

### PrivateEndpoint <a name="PrivateEndpoint" id="@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.Initializer"></a>

```typescript
import { PrivateEndpoint } from '@mongodbatlas-awscdk/private-endpoint'

const privateEndpoint: PrivateEndpoint = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.atlasPrivateEndpointStatus">atlasPrivateEndpointStatus</a></code> | <code>string</code> | Status of the Atlas PrivateEndpoint connection. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.awsPrivateEndpointStatus">awsPrivateEndpointStatus</a></code> | <code>string</code> | Status of the AWS PrivateEndpoint connection. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.interfaceEndpointId">interfaceEndpointId</a></code> | <code>string</code> | Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection. |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.subnetId">subnetId</a></code> | <code>string</code> | String Representing the AWS VPC Subnet ID (like: subnet-xxxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint). |
| <code><a href="#@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.vpcId">vpcId</a></code> | <code>string</code> | String Representing the AWS VPC ID (like: vpc-xxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint). |

---

##### `atlasPrivateEndpointStatus`<sup>Optional</sup> <a name="atlasPrivateEndpointStatus" id="@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.atlasPrivateEndpointStatus"></a>

```typescript
public readonly atlasPrivateEndpointStatus: string;
```

- *Type:* string

Status of the Atlas PrivateEndpoint connection.

---

##### `awsPrivateEndpointStatus`<sup>Optional</sup> <a name="awsPrivateEndpointStatus" id="@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.awsPrivateEndpointStatus"></a>

```typescript
public readonly awsPrivateEndpointStatus: string;
```

- *Type:* string

Status of the AWS PrivateEndpoint connection.

---

##### `interfaceEndpointId`<sup>Optional</sup> <a name="interfaceEndpointId" id="@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.interfaceEndpointId"></a>

```typescript
public readonly interfaceEndpointId: string;
```

- *Type:* string

Unique identifiers of the interface endpoints in your VPC that you added to the AWS PrivateLink connection.

---

##### `subnetId`<sup>Optional</sup> <a name="subnetId" id="@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.subnetId"></a>

```typescript
public readonly subnetId: string;
```

- *Type:* string

String Representing the AWS VPC Subnet ID (like: subnet-xxxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint).

---

##### `vpcId`<sup>Optional</sup> <a name="vpcId" id="@mongodbatlas-awscdk/private-endpoint.PrivateEndpoint.property.vpcId"></a>

```typescript
public readonly vpcId: string;
```

- *Type:* string

String Representing the AWS VPC ID (like: vpc-xxxxxxxxxxxxxxxx) (Used For Creating the AWS VPC Endpoint).

---



