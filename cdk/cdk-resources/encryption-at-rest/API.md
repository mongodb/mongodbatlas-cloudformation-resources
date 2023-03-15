# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnEncryptionAtRest <a name="CfnEncryptionAtRest" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest"></a>

A CloudFormation `MongoDB::Atlas::EncryptionAtRest`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.Initializer"></a>

```typescript
import { CfnEncryptionAtRest } from '@mongodbatlas-awscdk/encryption-at-rest'

new CfnEncryptionAtRest(scope: Construct, id: string, props: CfnEncryptionAtRestProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps">CfnEncryptionAtRestProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps">CfnEncryptionAtRestProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isConstruct"></a>

```typescript
import { CfnEncryptionAtRest } from '@mongodbatlas-awscdk/encryption-at-rest'

CfnEncryptionAtRest.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isCfnElement"></a>

```typescript
import { CfnEncryptionAtRest } from '@mongodbatlas-awscdk/encryption-at-rest'

CfnEncryptionAtRest.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isCfnResource"></a>

```typescript
import { CfnEncryptionAtRest } from '@mongodbatlas-awscdk/encryption-at-rest'

CfnEncryptionAtRest.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::EncryptionAtRest.Id`. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps">CfnEncryptionAtRestProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::EncryptionAtRest.Id`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.props"></a>

```typescript
public readonly props: CfnEncryptionAtRestProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps">CfnEncryptionAtRestProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRest.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### AwsKmsConfiguration <a name="AwsKmsConfiguration" id="@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration"></a>

Specifies AWS KMS configuration details and whether Encryption at Rest is enabled for an Atlas project.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.Initializer"></a>

```typescript
import { AwsKmsConfiguration } from '@mongodbatlas-awscdk/encryption-at-rest'

const awsKmsConfiguration: AwsKmsConfiguration = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.property.customerMasterKeyId">customerMasterKeyId</a></code> | <code>string</code> | The AWS customer master key used to encrypt and decrypt the MongoDB master keys. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.property.enabled">enabled</a></code> | <code>boolean</code> | Specifies whether Encryption at Rest is enabled for an Atlas project. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.property.region">region</a></code> | <code>string</code> | The AWS region in which the AWS customer master key exists. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.property.roleId">roleId</a></code> | <code>string</code> | ID of an AWS IAM role authorized to manage an AWS customer master key. |

---

##### `customerMasterKeyId`<sup>Optional</sup> <a name="customerMasterKeyId" id="@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.property.customerMasterKeyId"></a>

```typescript
public readonly customerMasterKeyId: string;
```

- *Type:* string

The AWS customer master key used to encrypt and decrypt the MongoDB master keys.

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

Specifies whether Encryption at Rest is enabled for an Atlas project.

To disable Encryption at Rest, pass only this parameter with a value of false. When you disable Encryption at Rest, Atlas also removes the configuration details.

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

The AWS region in which the AWS customer master key exists.

---

##### `roleId`<sup>Optional</sup> <a name="roleId" id="@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration.property.roleId"></a>

```typescript
public readonly roleId: string;
```

- *Type:* string

ID of an AWS IAM role authorized to manage an AWS customer master key.

---

### CfnEncryptionAtRestProps <a name="CfnEncryptionAtRestProps" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps"></a>

Returns and edits the Encryption at Rest using Customer Key Management configuration.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps.Initializer"></a>

```typescript
import { CfnEncryptionAtRestProps } from '@mongodbatlas-awscdk/encryption-at-rest'

const cfnEncryptionAtRestProps: CfnEncryptionAtRestProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps.property.awsKms">awsKms</a></code> | <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration">AwsKmsConfiguration</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps.property.projectId">projectId</a></code> | <code>string</code> | Unique identifier of the Atlas project to which the user belongs. |
| <code><a href="#@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps.property.profile">profile</a></code> | <code>string</code> | The profile is defined in AWS Secret manager. |

---

##### `awsKms`<sup>Required</sup> <a name="awsKms" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps.property.awsKms"></a>

```typescript
public readonly awsKms: AwsKmsConfiguration;
```

- *Type:* <a href="#@mongodbatlas-awscdk/encryption-at-rest.AwsKmsConfiguration">AwsKmsConfiguration</a>

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique identifier of the Atlas project to which the user belongs.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/encryption-at-rest.CfnEncryptionAtRestProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

The profile is defined in AWS Secret manager.

See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

---



