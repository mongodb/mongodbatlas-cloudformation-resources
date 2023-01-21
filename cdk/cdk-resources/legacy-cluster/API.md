# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnCluster <a name="CfnCluster" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster"></a>

A CloudFormation `MongoDB::Atlas::Cluster`.

#### Initializers <a name="Initializers" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.Initializer"></a>

```typescript
import { CfnCluster } from '@mongodb-cdk/atlas-legacy-cluster'

new CfnCluster(scope: Construct, id: string, props: CfnClusterProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps">CfnClusterProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps">CfnClusterProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isConstruct"></a>

```typescript
import { CfnCluster } from '@mongodb-cdk/atlas-legacy-cluster'

CfnCluster.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isCfnElement"></a>

```typescript
import { CfnCluster } from '@mongodb-cdk/atlas-legacy-cluster'

CfnCluster.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isCfnResource"></a>

```typescript
import { CfnCluster } from '@mongodb-cdk/atlas-legacy-cluster'

CfnCluster.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.Id`. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrMongoDBVersion">attrMongoDBVersion</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.MongoDBVersion`. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrMongoURI">attrMongoURI</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.MongoURI`. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrMongoURIUpdated">attrMongoURIUpdated</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.MongoURIUpdated`. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrMongoURIWithOptions">attrMongoURIWithOptions</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.MongoURIWithOptions`. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrSrvAddress">attrSrvAddress</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.SrvAddress`. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrStateName">attrStateName</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Cluster.StateName`. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps">CfnClusterProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.Id`.

---

##### `attrMongoDBVersion`<sup>Required</sup> <a name="attrMongoDBVersion" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrMongoDBVersion"></a>

```typescript
public readonly attrMongoDBVersion: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.MongoDBVersion`.

---

##### `attrMongoURI`<sup>Required</sup> <a name="attrMongoURI" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrMongoURI"></a>

```typescript
public readonly attrMongoURI: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.MongoURI`.

---

##### `attrMongoURIUpdated`<sup>Required</sup> <a name="attrMongoURIUpdated" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrMongoURIUpdated"></a>

```typescript
public readonly attrMongoURIUpdated: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.MongoURIUpdated`.

---

##### `attrMongoURIWithOptions`<sup>Required</sup> <a name="attrMongoURIWithOptions" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrMongoURIWithOptions"></a>

```typescript
public readonly attrMongoURIWithOptions: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.MongoURIWithOptions`.

---

##### `attrSrvAddress`<sup>Required</sup> <a name="attrSrvAddress" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrSrvAddress"></a>

```typescript
public readonly attrSrvAddress: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.SrvAddress`.

---

##### `attrStateName`<sup>Required</sup> <a name="attrStateName" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.attrStateName"></a>

```typescript
public readonly attrStateName: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Cluster.StateName`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.props"></a>

```typescript
public readonly props: CfnClusterProps;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps">CfnClusterProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodb-cdk/atlas-legacy-cluster.CfnCluster.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodb-cdk/atlas-legacy-cluster.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodb-cdk/atlas-legacy-cluster'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodb-cdk/atlas-legacy-cluster.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodb-cdk/atlas-legacy-cluster.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### AutoScaling <a name="AutoScaling" id="@mongodb-cdk/atlas-legacy-cluster.AutoScaling"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.AutoScaling.Initializer"></a>

```typescript
import { AutoScaling } from '@mongodb-cdk/atlas-legacy-cluster'

const autoScaling: AutoScaling = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.AutoScaling.property.compute">compute</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.Compute">Compute</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.AutoScaling.property.diskGbEnabled">diskGbEnabled</a></code> | <code>boolean</code> | *No description.* |

---

##### `compute`<sup>Optional</sup> <a name="compute" id="@mongodb-cdk/atlas-legacy-cluster.AutoScaling.property.compute"></a>

```typescript
public readonly compute: Compute;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.Compute">Compute</a>

---

##### `diskGbEnabled`<sup>Optional</sup> <a name="diskGbEnabled" id="@mongodb-cdk/atlas-legacy-cluster.AutoScaling.property.diskGbEnabled"></a>

```typescript
public readonly diskGbEnabled: boolean;
```

- *Type:* boolean

---

### CfnClusterProps <a name="CfnClusterProps" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps"></a>

The cluster resource provides access to your cluster configurations.

The resource lets you create, edit and delete clusters. The resource requires your Project ID.

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.Initializer"></a>

```typescript
import { CfnClusterProps } from '@mongodb-cdk/atlas-legacy-cluster'

const cfnClusterProps: CfnClusterProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.name">name</a></code> | <code>string</code> | Name of the cluster. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.autoScaling">autoScaling</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.AutoScaling">AutoScaling</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.backupEnabled">backupEnabled</a></code> | <code>boolean</code> | Applicable only for M10+ clusters. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.biConnector">biConnector</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsBiConnector">CfnClusterPropsBiConnector</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.clusterType">clusterType</a></code> | <code>string</code> | Type of the cluster that you want to create. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.connectionStrings">connectionStrings</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings">ConnectionStrings</a></code> | Set of connection strings that your applications use to connect to this cluster. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.diskSizeGb">diskSizeGb</a></code> | <code>number</code> | Capacity, in gigabytes, of the hosts root volume. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.encryptionAtRestProvider">encryptionAtRestProvider</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider">CfnClusterPropsEncryptionAtRestProvider</a></code> | Set the Encryption at Rest parameter. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.labels">labels</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsLabels">CfnClusterPropsLabels</a>[]</code> | Array containing key-value pairs that tag and categorize the cluster. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.mongoDbMajorVersion">mongoDbMajorVersion</a></code> | <code>string</code> | Major version of the cluster to deploy. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.numShards">numShards</a></code> | <code>number</code> | Positive integer that specifies the number of shards to deploy for a sharded cluster. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.paused">paused</a></code> | <code>boolean</code> | Flag that indicates whether the cluster is paused or not. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.pitEnabled">pitEnabled</a></code> | <code>boolean</code> | Flag that indicates if the cluster uses Point-in-Time backups. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.projectId">projectId</a></code> | <code>string</code> | Unique identifier of the project the cluster belongs to. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.providerBackupEnabled">providerBackupEnabled</a></code> | <code>boolean</code> | Applicable only for M10+ clusters. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.providerSettings">providerSettings</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings">CfnClusterPropsProviderSettings</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.replicationFactor">replicationFactor</a></code> | <code>number</code> | ReplicationFactor is deprecated. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.replicationSpecs">replicationSpecs</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec">ReplicationSpec</a>[]</code> | Configuration for cluster regions. |

---

##### `name`<sup>Required</sup> <a name="name" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Name of the cluster.

Once the cluster is created, its name cannot be changed.

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `autoScaling`<sup>Optional</sup> <a name="autoScaling" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.autoScaling"></a>

```typescript
public readonly autoScaling: AutoScaling;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.AutoScaling">AutoScaling</a>

---

##### `backupEnabled`<sup>Optional</sup> <a name="backupEnabled" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.backupEnabled"></a>

```typescript
public readonly backupEnabled: boolean;
```

- *Type:* boolean

Applicable only for M10+ clusters.

Set to true to enable Atlas continuous backups for the cluster. Set to false to disable continuous backups for the cluster. Atlas deletes any stored snapshots. See the continuous backup Snapshot Schedule for more information. You cannot enable continuous backups if you have an existing cluster in the project with Cloud Provider Snapshots enabled. The default value is false.

---

##### `biConnector`<sup>Optional</sup> <a name="biConnector" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.biConnector"></a>

```typescript
public readonly biConnector: CfnClusterPropsBiConnector;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsBiConnector">CfnClusterPropsBiConnector</a>

---

##### `clusterType`<sup>Optional</sup> <a name="clusterType" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.clusterType"></a>

```typescript
public readonly clusterType: string;
```

- *Type:* string

Type of the cluster that you want to create.

---

##### `connectionStrings`<sup>Optional</sup> <a name="connectionStrings" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.connectionStrings"></a>

```typescript
public readonly connectionStrings: ConnectionStrings;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings">ConnectionStrings</a>

Set of connection strings that your applications use to connect to this cluster.

Use the parameters in this object to connect your applications to this cluster. See the MongoDB [Connection String URI Format](https://docs.mongodb.com/manual/reference/connection-string/) reference for further details.

---

##### `diskSizeGb`<sup>Optional</sup> <a name="diskSizeGb" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.diskSizeGb"></a>

```typescript
public readonly diskSizeGb: number;
```

- *Type:* number

Capacity, in gigabytes, of the hosts root volume.

Increase this number to add capacity, up to a maximum possible value of 4096 (i.e., 4 TB). This value must be a positive integer.

---

##### `encryptionAtRestProvider`<sup>Optional</sup> <a name="encryptionAtRestProvider" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.encryptionAtRestProvider"></a>

```typescript
public readonly encryptionAtRestProvider: CfnClusterPropsEncryptionAtRestProvider;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider">CfnClusterPropsEncryptionAtRestProvider</a>

Set the Encryption at Rest parameter.

---

##### `labels`<sup>Optional</sup> <a name="labels" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.labels"></a>

```typescript
public readonly labels: CfnClusterPropsLabels[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsLabels">CfnClusterPropsLabels</a>[]

Array containing key-value pairs that tag and categorize the cluster.

---

##### `mongoDbMajorVersion`<sup>Optional</sup> <a name="mongoDbMajorVersion" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.mongoDbMajorVersion"></a>

```typescript
public readonly mongoDbMajorVersion: string;
```

- *Type:* string

Major version of the cluster to deploy.

---

##### `numShards`<sup>Optional</sup> <a name="numShards" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.numShards"></a>

```typescript
public readonly numShards: number;
```

- *Type:* number

Positive integer that specifies the number of shards to deploy for a sharded cluster.

---

##### `paused`<sup>Optional</sup> <a name="paused" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.paused"></a>

```typescript
public readonly paused: boolean;
```

- *Type:* boolean

Flag that indicates whether the cluster is paused or not.

---

##### `pitEnabled`<sup>Optional</sup> <a name="pitEnabled" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.pitEnabled"></a>

```typescript
public readonly pitEnabled: boolean;
```

- *Type:* boolean

Flag that indicates if the cluster uses Point-in-Time backups.

If set to true, providerBackupEnabled must also be set to true.

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique identifier of the project the cluster belongs to.

---

##### `providerBackupEnabled`<sup>Optional</sup> <a name="providerBackupEnabled" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.providerBackupEnabled"></a>

```typescript
public readonly providerBackupEnabled: boolean;
```

- *Type:* boolean

Applicable only for M10+ clusters.

Set to true to enable Atlas Cloud Provider Snapshots backups for the cluster. Set to false to disable Cloud Provider Snapshots backups for the cluster. You cannot enable Cloud Provider Snapshots if you have an existing cluster in the project with continuous backups enabled. Note that you must set this value to true for NVMe clusters. The default value is false.

---

##### `providerSettings`<sup>Optional</sup> <a name="providerSettings" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.providerSettings"></a>

```typescript
public readonly providerSettings: CfnClusterPropsProviderSettings;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings">CfnClusterPropsProviderSettings</a>

---

##### `replicationFactor`<sup>Optional</sup> <a name="replicationFactor" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.replicationFactor"></a>

```typescript
public readonly replicationFactor: number;
```

- *Type:* number

ReplicationFactor is deprecated.

Use replicationSpecs.

---

##### `replicationSpecs`<sup>Optional</sup> <a name="replicationSpecs" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterProps.property.replicationSpecs"></a>

```typescript
public readonly replicationSpecs: ReplicationSpec[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec">ReplicationSpec</a>[]

Configuration for cluster regions.

---

### CfnClusterPropsBiConnector <a name="CfnClusterPropsBiConnector" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsBiConnector"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsBiConnector.Initializer"></a>

```typescript
import { CfnClusterPropsBiConnector } from '@mongodb-cdk/atlas-legacy-cluster'

const cfnClusterPropsBiConnector: CfnClusterPropsBiConnector = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsBiConnector.property.enabled">enabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsBiConnector.property.readPreference">readPreference</a></code> | <code>string</code> | *No description.* |

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsBiConnector.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

---

##### `readPreference`<sup>Optional</sup> <a name="readPreference" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsBiConnector.property.readPreference"></a>

```typescript
public readonly readPreference: string;
```

- *Type:* string

---

### CfnClusterPropsLabels <a name="CfnClusterPropsLabels" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsLabels"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsLabels.Initializer"></a>

```typescript
import { CfnClusterPropsLabels } from '@mongodb-cdk/atlas-legacy-cluster'

const cfnClusterPropsLabels: CfnClusterPropsLabels = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsLabels.property.key">key</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsLabels.property.value">value</a></code> | <code>string</code> | *No description.* |

---

##### `key`<sup>Optional</sup> <a name="key" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsLabels.property.key"></a>

```typescript
public readonly key: string;
```

- *Type:* string

---

##### `value`<sup>Optional</sup> <a name="value" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsLabels.property.value"></a>

```typescript
public readonly value: string;
```

- *Type:* string

---

### CfnClusterPropsProviderSettings <a name="CfnClusterPropsProviderSettings" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.Initializer"></a>

```typescript
import { CfnClusterPropsProviderSettings } from '@mongodb-cdk/atlas-legacy-cluster'

const cfnClusterPropsProviderSettings: CfnClusterPropsProviderSettings = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.autoScaling">autoScaling</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.AutoScaling">AutoScaling</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.backingProviderName">backingProviderName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.diskIops">diskIops</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.encryptEbsVolume">encryptEbsVolume</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.instanceSizeName">instanceSizeName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.providerName">providerName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.regionName">regionName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.volumeType">volumeType</a></code> | <code>string</code> | *No description.* |

---

##### `autoScaling`<sup>Optional</sup> <a name="autoScaling" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.autoScaling"></a>

```typescript
public readonly autoScaling: AutoScaling;
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.AutoScaling">AutoScaling</a>

---

##### `backingProviderName`<sup>Optional</sup> <a name="backingProviderName" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.backingProviderName"></a>

```typescript
public readonly backingProviderName: string;
```

- *Type:* string

---

##### `diskIops`<sup>Optional</sup> <a name="diskIops" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.diskIops"></a>

```typescript
public readonly diskIops: number;
```

- *Type:* number

---

##### `encryptEbsVolume`<sup>Optional</sup> <a name="encryptEbsVolume" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.encryptEbsVolume"></a>

```typescript
public readonly encryptEbsVolume: boolean;
```

- *Type:* boolean

---

##### `instanceSizeName`<sup>Optional</sup> <a name="instanceSizeName" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.instanceSizeName"></a>

```typescript
public readonly instanceSizeName: string;
```

- *Type:* string

---

##### `providerName`<sup>Optional</sup> <a name="providerName" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.providerName"></a>

```typescript
public readonly providerName: string;
```

- *Type:* string

---

##### `regionName`<sup>Optional</sup> <a name="regionName" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.regionName"></a>

```typescript
public readonly regionName: string;
```

- *Type:* string

---

##### `volumeType`<sup>Optional</sup> <a name="volumeType" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsProviderSettings.property.volumeType"></a>

```typescript
public readonly volumeType: string;
```

- *Type:* string

---

### Compute <a name="Compute" id="@mongodb-cdk/atlas-legacy-cluster.Compute"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.Compute.Initializer"></a>

```typescript
import { Compute } from '@mongodb-cdk/atlas-legacy-cluster'

const compute: Compute = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.Compute.property.enabled">enabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.Compute.property.maxInstanceSize">maxInstanceSize</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.Compute.property.minInstanceSize">minInstanceSize</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.Compute.property.scaleDownEnabled">scaleDownEnabled</a></code> | <code>boolean</code> | *No description.* |

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodb-cdk/atlas-legacy-cluster.Compute.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

---

##### `maxInstanceSize`<sup>Optional</sup> <a name="maxInstanceSize" id="@mongodb-cdk/atlas-legacy-cluster.Compute.property.maxInstanceSize"></a>

```typescript
public readonly maxInstanceSize: string;
```

- *Type:* string

---

##### `minInstanceSize`<sup>Optional</sup> <a name="minInstanceSize" id="@mongodb-cdk/atlas-legacy-cluster.Compute.property.minInstanceSize"></a>

```typescript
public readonly minInstanceSize: string;
```

- *Type:* string

---

##### `scaleDownEnabled`<sup>Optional</sup> <a name="scaleDownEnabled" id="@mongodb-cdk/atlas-legacy-cluster.Compute.property.scaleDownEnabled"></a>

```typescript
public readonly scaleDownEnabled: boolean;
```

- *Type:* boolean

---

### ConnectionStrings <a name="ConnectionStrings" id="@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.Initializer"></a>

```typescript
import { ConnectionStrings } from '@mongodb-cdk/atlas-legacy-cluster'

const connectionStrings: ConnectionStrings = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.property.private">private</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.property.privateSrv">privateSrv</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.property.standard">standard</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.property.standardSrv">standardSrv</a></code> | <code>string</code> | *No description.* |

---

##### `private`<sup>Optional</sup> <a name="private" id="@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.property.private"></a>

```typescript
public readonly private: string;
```

- *Type:* string

---

##### `privateSrv`<sup>Optional</sup> <a name="privateSrv" id="@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.property.privateSrv"></a>

```typescript
public readonly privateSrv: string;
```

- *Type:* string

---

##### `standard`<sup>Optional</sup> <a name="standard" id="@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.property.standard"></a>

```typescript
public readonly standard: string;
```

- *Type:* string

---

##### `standardSrv`<sup>Optional</sup> <a name="standardSrv" id="@mongodb-cdk/atlas-legacy-cluster.ConnectionStrings.property.standardSrv"></a>

```typescript
public readonly standardSrv: string;
```

- *Type:* string

---

### RegionsConfig <a name="RegionsConfig" id="@mongodb-cdk/atlas-legacy-cluster.RegionsConfig"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.Initializer"></a>

```typescript
import { RegionsConfig } from '@mongodb-cdk/atlas-legacy-cluster'

const regionsConfig: RegionsConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.analyticsNodes">analyticsNodes</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.electableNodes">electableNodes</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.priority">priority</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.readOnlyNodes">readOnlyNodes</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.regionName">regionName</a></code> | <code>string</code> | *No description.* |

---

##### `analyticsNodes`<sup>Optional</sup> <a name="analyticsNodes" id="@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.analyticsNodes"></a>

```typescript
public readonly analyticsNodes: number;
```

- *Type:* number

---

##### `electableNodes`<sup>Optional</sup> <a name="electableNodes" id="@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.electableNodes"></a>

```typescript
public readonly electableNodes: number;
```

- *Type:* number

---

##### `priority`<sup>Optional</sup> <a name="priority" id="@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.priority"></a>

```typescript
public readonly priority: number;
```

- *Type:* number

---

##### `readOnlyNodes`<sup>Optional</sup> <a name="readOnlyNodes" id="@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.readOnlyNodes"></a>

```typescript
public readonly readOnlyNodes: number;
```

- *Type:* number

---

##### `regionName`<sup>Optional</sup> <a name="regionName" id="@mongodb-cdk/atlas-legacy-cluster.RegionsConfig.property.regionName"></a>

```typescript
public readonly regionName: string;
```

- *Type:* string

---

### ReplicationSpec <a name="ReplicationSpec" id="@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.Initializer"></a>

```typescript
import { ReplicationSpec } from '@mongodb-cdk/atlas-legacy-cluster'

const replicationSpec: ReplicationSpec = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.property.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.property.numShards">numShards</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.property.regionsConfig">regionsConfig</a></code> | <code><a href="#@mongodb-cdk/atlas-legacy-cluster.RegionsConfig">RegionsConfig</a>[]</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.property.zoneName">zoneName</a></code> | <code>string</code> | *No description.* |

---

##### `id`<sup>Optional</sup> <a name="id" id="@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.property.id"></a>

```typescript
public readonly id: string;
```

- *Type:* string

---

##### `numShards`<sup>Optional</sup> <a name="numShards" id="@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.property.numShards"></a>

```typescript
public readonly numShards: number;
```

- *Type:* number

---

##### `regionsConfig`<sup>Optional</sup> <a name="regionsConfig" id="@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.property.regionsConfig"></a>

```typescript
public readonly regionsConfig: RegionsConfig[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-legacy-cluster.RegionsConfig">RegionsConfig</a>[]

---

##### `zoneName`<sup>Optional</sup> <a name="zoneName" id="@mongodb-cdk/atlas-legacy-cluster.ReplicationSpec.property.zoneName"></a>

```typescript
public readonly zoneName: string;
```

- *Type:* string

---



## Enums <a name="Enums" id="Enums"></a>

### CfnClusterPropsEncryptionAtRestProvider <a name="CfnClusterPropsEncryptionAtRestProvider" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider"></a>

Set the Encryption at Rest parameter.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider.AWS">AWS</a></code> | AWS. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider.GCP">GCP</a></code> | GCP. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider.AZURE">AZURE</a></code> | AZURE. |
| <code><a href="#@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider.NONE">NONE</a></code> | NONE. |

---

##### `AWS` <a name="AWS" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider.AWS"></a>

AWS.

---


##### `GCP` <a name="GCP" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider.GCP"></a>

GCP.

---


##### `AZURE` <a name="AZURE" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider.AZURE"></a>

AZURE.

---


##### `NONE` <a name="NONE" id="@mongodb-cdk/atlas-legacy-cluster.CfnClusterPropsEncryptionAtRestProvider.NONE"></a>

NONE.

---

