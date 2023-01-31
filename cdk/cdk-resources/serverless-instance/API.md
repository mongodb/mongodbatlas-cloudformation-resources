# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnServerlessInstance <a name="CfnServerlessInstance" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance"></a>

A CloudFormation `MongoDB::Atlas::ServerlessInstance`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.Initializer"></a>

```typescript
import { CfnServerlessInstance } from '@mongodbatlas-awscdk/serverless-instance'

new CfnServerlessInstance(scope: Construct, id: string, props: CfnServerlessInstanceProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps">CfnServerlessInstanceProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps">CfnServerlessInstanceProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isConstruct"></a>

```typescript
import { CfnServerlessInstance } from '@mongodbatlas-awscdk/serverless-instance'

CfnServerlessInstance.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isCfnElement"></a>

```typescript
import { CfnServerlessInstance } from '@mongodbatlas-awscdk/serverless-instance'

CfnServerlessInstance.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isCfnResource"></a>

```typescript
import { CfnServerlessInstance } from '@mongodbatlas-awscdk/serverless-instance'

CfnServerlessInstance.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrCreateDate">attrCreateDate</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::ServerlessInstance.CreateDate`. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::ServerlessInstance.Id`. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrLinks">attrLinks</a></code> | <code>any[]</code> | Attribute `MongoDB::Atlas::ServerlessInstance.Links`. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrMongoDBVersion">attrMongoDBVersion</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::ServerlessInstance.MongoDBVersion`. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrStateName">attrStateName</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::ServerlessInstance.StateName`. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrTotalCount">attrTotalCount</a></code> | <code>number</code> | Attribute `MongoDB::Atlas::ServerlessInstance.TotalCount`. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps">CfnServerlessInstanceProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrCreateDate`<sup>Required</sup> <a name="attrCreateDate" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrCreateDate"></a>

```typescript
public readonly attrCreateDate: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::ServerlessInstance.CreateDate`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::ServerlessInstance.Id`.

---

##### `attrLinks`<sup>Required</sup> <a name="attrLinks" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrLinks"></a>

```typescript
public readonly attrLinks: any[];
```

- *Type:* any[]

Attribute `MongoDB::Atlas::ServerlessInstance.Links`.

---

##### `attrMongoDBVersion`<sup>Required</sup> <a name="attrMongoDBVersion" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrMongoDBVersion"></a>

```typescript
public readonly attrMongoDBVersion: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::ServerlessInstance.MongoDBVersion`.

---

##### `attrStateName`<sup>Required</sup> <a name="attrStateName" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrStateName"></a>

```typescript
public readonly attrStateName: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::ServerlessInstance.StateName`.

---

##### `attrTotalCount`<sup>Required</sup> <a name="attrTotalCount" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.attrTotalCount"></a>

```typescript
public readonly attrTotalCount: number;
```

- *Type:* number

Attribute `MongoDB::Atlas::ServerlessInstance.TotalCount`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.props"></a>

```typescript
public readonly props: CfnServerlessInstanceProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps">CfnServerlessInstanceProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstance.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/serverless-instance.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/serverless-instance.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/serverless-instance'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/serverless-instance.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/serverless-instance.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnServerlessInstanceProps <a name="CfnServerlessInstanceProps" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps"></a>

Returns, adds, edits, and removes serverless instances.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.Initializer"></a>

```typescript
import { CfnServerlessInstanceProps } from '@mongodbatlas-awscdk/serverless-instance'

const cfnServerlessInstanceProps: CfnServerlessInstanceProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.connectionStrings">connectionStrings</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceConnectionStrings">ServerlessInstanceConnectionStrings</a></code> | Collection of Uniform Resource Locators that point to the MongoDB database. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.continuousBackupEnabled">continuousBackupEnabled</a></code> | <code>boolean</code> | Flag that indicates whether the serverless instances uses Serverless Continuous Backup. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.includeCount">includeCount</a></code> | <code>boolean</code> | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.itemsPerPage">itemsPerPage</a></code> | <code>number</code> | Number of items that the response returns per page. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.name">name</a></code> | <code>string</code> | Human-readable label that identifies the serverless instance. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.pageNum">pageNum</a></code> | <code>number</code> | Number of the page that displays the current set of the total objects that the response returns. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.providerSettings">providerSettings</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettings">ServerlessInstanceProviderSettings</a></code> | Group of settings that configure the provisioned MongoDB serverless instance. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.terminationProtectionEnabled">terminationProtectionEnabled</a></code> | <code>boolean</code> | Flag that indicates whether termination protection is enabled on the serverless instance. |

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `connectionStrings`<sup>Optional</sup> <a name="connectionStrings" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.connectionStrings"></a>

```typescript
public readonly connectionStrings: ServerlessInstanceConnectionStrings;
```

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceConnectionStrings">ServerlessInstanceConnectionStrings</a>

Collection of Uniform Resource Locators that point to the MongoDB database.

---

##### `continuousBackupEnabled`<sup>Optional</sup> <a name="continuousBackupEnabled" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.continuousBackupEnabled"></a>

```typescript
public readonly continuousBackupEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether the serverless instances uses Serverless Continuous Backup.

If this parameter is false, the serverless instance uses Basic Backup. | Option | Description | |---|---| | Serverless Continuous Backup | Atlas takes incremental snapshots of the data in your serverless instance every six hours and lets you restore the data from a selected point in time within the last 72 hours. Atlas also takes daily snapshots and retains these daily snapshots for 35 days. To learn more, see Serverless Instance Costs. | | Basic Backup | Atlas takes incremental snapshots of the data in your serverless instance every six hours and retains only the two most recent snapshots. You can use this option for free.

---

##### `includeCount`<sup>Optional</sup> <a name="includeCount" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.includeCount"></a>

```typescript
public readonly includeCount: boolean;
```

- *Type:* boolean

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

---

##### `itemsPerPage`<sup>Optional</sup> <a name="itemsPerPage" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.itemsPerPage"></a>

```typescript
public readonly itemsPerPage: number;
```

- *Type:* number

Number of items that the response returns per page.

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable label that identifies the serverless instance.

---

##### `pageNum`<sup>Optional</sup> <a name="pageNum" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.pageNum"></a>

```typescript
public readonly pageNum: number;
```

- *Type:* number

Number of the page that displays the current set of the total objects that the response returns.

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `providerSettings`<sup>Optional</sup> <a name="providerSettings" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.providerSettings"></a>

```typescript
public readonly providerSettings: ServerlessInstanceProviderSettings;
```

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettings">ServerlessInstanceProviderSettings</a>

Group of settings that configure the provisioned MongoDB serverless instance.

The options available relate to the cloud service provider.

---

##### `terminationProtectionEnabled`<sup>Optional</sup> <a name="terminationProtectionEnabled" id="@mongodbatlas-awscdk/serverless-instance.CfnServerlessInstanceProps.property.terminationProtectionEnabled"></a>

```typescript
public readonly terminationProtectionEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether termination protection is enabled on the serverless instance.

If set to true, MongoDB Cloud won't delete the serverless instance. If set to false, MongoDB cloud will delete the serverless instance."

---

### ServerlessInstanceConnectionStrings <a name="ServerlessInstanceConnectionStrings" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceConnectionStrings"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceConnectionStrings.Initializer"></a>

```typescript
import { ServerlessInstanceConnectionStrings } from '@mongodbatlas-awscdk/serverless-instance'

const serverlessInstanceConnectionStrings: ServerlessInstanceConnectionStrings = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceConnectionStrings.property.privateEndpoint">privateEndpoint</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint">ServerlessInstancePrivateEndpoint</a>[]</code> | List of private endpoint connection strings that you can use to connect to this serverless instance through a private endpoint. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceConnectionStrings.property.standardSrv">standardSrv</a></code> | <code>string</code> | Public connection string that you can use to connect to this serverless instance. |

---

##### `privateEndpoint`<sup>Optional</sup> <a name="privateEndpoint" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceConnectionStrings.property.privateEndpoint"></a>

```typescript
public readonly privateEndpoint: ServerlessInstancePrivateEndpoint[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint">ServerlessInstancePrivateEndpoint</a>[]

List of private endpoint connection strings that you can use to connect to this serverless instance through a private endpoint.

This parameter returns only if you created a private endpoint for this serverless instance and it is AVAILABLE.

---

##### `standardSrv`<sup>Optional</sup> <a name="standardSrv" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceConnectionStrings.property.standardSrv"></a>

```typescript
public readonly standardSrv: string;
```

- *Type:* string

Public connection string that you can use to connect to this serverless instance.

This connection string uses the `mongodb+srv://` protocol.

---

### ServerlessInstancePrivateEndpoint <a name="ServerlessInstancePrivateEndpoint" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint.Initializer"></a>

```typescript
import { ServerlessInstancePrivateEndpoint } from '@mongodbatlas-awscdk/serverless-instance'

const serverlessInstancePrivateEndpoint: ServerlessInstancePrivateEndpoint = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint.property.endpoints">endpoints</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint">ServerlessInstancePrivateEndpointEndpoint</a>[]</code> | List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].srvConnectionString**. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint.property.srvConnectionString">srvConnectionString</a></code> | <code>string</code> | Private endpoint-aware connection string that uses the `mongodb+srv://` protocol to connect to MongoDB Cloud through a private endpoint. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint.property.type">type</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointType">ServerlessInstancePrivateEndpointType</a></code> | MongoDB process type to which your application connects. |

---

##### `endpoints`<sup>Optional</sup> <a name="endpoints" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint.property.endpoints"></a>

```typescript
public readonly endpoints: ServerlessInstancePrivateEndpointEndpoint[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint">ServerlessInstancePrivateEndpointEndpoint</a>[]

List that contains the private endpoints through which you connect to MongoDB Cloud when you use **connectionStrings.privateEndpoint[n].srvConnectionString**.

---

##### `srvConnectionString`<sup>Optional</sup> <a name="srvConnectionString" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint.property.srvConnectionString"></a>

```typescript
public readonly srvConnectionString: string;
```

- *Type:* string

Private endpoint-aware connection string that uses the `mongodb+srv://` protocol to connect to MongoDB Cloud through a private endpoint.

The `mongodb+srv` protocol tells the driver to look up the seed list of hosts in the Domain Name System (DNS).

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpoint.property.type"></a>

```typescript
public readonly type: ServerlessInstancePrivateEndpointType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointType">ServerlessInstancePrivateEndpointType</a>

MongoDB process type to which your application connects.

---

### ServerlessInstancePrivateEndpointEndpoint <a name="ServerlessInstancePrivateEndpointEndpoint" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint.Initializer"></a>

```typescript
import { ServerlessInstancePrivateEndpointEndpoint } from '@mongodbatlas-awscdk/serverless-instance'

const serverlessInstancePrivateEndpointEndpoint: ServerlessInstancePrivateEndpointEndpoint = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint.property.endpointId">endpointId</a></code> | <code>string</code> | Unique provider identifier of the private endpoint. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint.property.providerName">providerName</a></code> | <code>string</code> | Cloud provider where the private endpoint is deployed. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint.property.region">region</a></code> | <code>string</code> | Region where the private endpoint is deployed. |

---

##### `endpointId`<sup>Optional</sup> <a name="endpointId" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint.property.endpointId"></a>

```typescript
public readonly endpointId: string;
```

- *Type:* string

Unique provider identifier of the private endpoint.

---

##### `providerName`<sup>Optional</sup> <a name="providerName" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint.property.providerName"></a>

```typescript
public readonly providerName: string;
```

- *Type:* string

Cloud provider where the private endpoint is deployed.

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointEndpoint.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

Region where the private endpoint is deployed.

---

### ServerlessInstanceProviderSettings <a name="ServerlessInstanceProviderSettings" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettings"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettings.Initializer"></a>

```typescript
import { ServerlessInstanceProviderSettings } from '@mongodbatlas-awscdk/serverless-instance'

const serverlessInstanceProviderSettings: ServerlessInstanceProviderSettings = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettings.property.providerName">providerName</a></code> | <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettingsProviderName">ServerlessInstanceProviderSettingsProviderName</a></code> | Human-readable label that identifies the cloud service provider. |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettings.property.regionName">regionName</a></code> | <code>string</code> | Human-readable label that identifies the geographic location of your MongoDB serverless instance. |

---

##### `providerName`<sup>Optional</sup> <a name="providerName" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettings.property.providerName"></a>

```typescript
public readonly providerName: ServerlessInstanceProviderSettingsProviderName;
```

- *Type:* <a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettingsProviderName">ServerlessInstanceProviderSettingsProviderName</a>

Human-readable label that identifies the cloud service provider.

---

##### `regionName`<sup>Optional</sup> <a name="regionName" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettings.property.regionName"></a>

```typescript
public readonly regionName: string;
```

- *Type:* string

Human-readable label that identifies the geographic location of your MongoDB serverless instance.

The region you choose can affect network latency for clients accessing your databases. For a complete list of region names, see [AWS](https://docs.atlas.mongodb.com/reference/amazon-aws/#std-label-amazon-aws), [GCP](https://docs.atlas.mongodb.com/reference/google-gcp/), and [Azure](https://docs.atlas.mongodb.com/reference/microsoft-azure/).

---



## Enums <a name="Enums" id="Enums"></a>

### ServerlessInstancePrivateEndpointType <a name="ServerlessInstancePrivateEndpointType" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointType"></a>

MongoDB process type to which your application connects.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointType.MONGOS">MONGOS</a></code> | MONGOS. |

---

##### `MONGOS` <a name="MONGOS" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstancePrivateEndpointType.MONGOS"></a>

MONGOS.

---


### ServerlessInstanceProviderSettingsProviderName <a name="ServerlessInstanceProviderSettingsProviderName" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettingsProviderName"></a>

Human-readable label that identifies the cloud service provider.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettingsProviderName.SERVERLESS">SERVERLESS</a></code> | SERVERLESS. |

---

##### `SERVERLESS` <a name="SERVERLESS" id="@mongodbatlas-awscdk/serverless-instance.ServerlessInstanceProviderSettingsProviderName.SERVERLESS"></a>

SERVERLESS.

---

