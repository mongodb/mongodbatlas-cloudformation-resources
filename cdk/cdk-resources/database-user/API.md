# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnDatabaseUser <a name="CfnDatabaseUser" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser"></a>

A CloudFormation `MongoDB::Atlas::DatabaseUser`.

#### Initializers <a name="Initializers" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.Initializer"></a>

```typescript
import { CfnDatabaseUser } from '@mongodb-cdk/atlas-database-user'

new CfnDatabaseUser(scope: Construct, id: string, props: CfnDatabaseUserProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps">CfnDatabaseUserProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps">CfnDatabaseUserProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isConstruct"></a>

```typescript
import { CfnDatabaseUser } from '@mongodb-cdk/atlas-database-user'

CfnDatabaseUser.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isCfnElement"></a>

```typescript
import { CfnDatabaseUser } from '@mongodb-cdk/atlas-database-user'

CfnDatabaseUser.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isCfnResource"></a>

```typescript
import { CfnDatabaseUser } from '@mongodb-cdk/atlas-database-user'

CfnDatabaseUser.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.attrUserCFNIdentifier">attrUserCFNIdentifier</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::DatabaseUser.UserCFNIdentifier`. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps">CfnDatabaseUserProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrUserCFNIdentifier`<sup>Required</sup> <a name="attrUserCFNIdentifier" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.attrUserCFNIdentifier"></a>

```typescript
public readonly attrUserCFNIdentifier: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::DatabaseUser.UserCFNIdentifier`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.props"></a>

```typescript
public readonly props: CfnDatabaseUserProps;
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps">CfnDatabaseUserProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUser.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodb-cdk/atlas-database-user.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-database-user.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodb-cdk/atlas-database-user'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-database-user.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodb-cdk/atlas-database-user.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodb-cdk/atlas-database-user.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnDatabaseUserProps <a name="CfnDatabaseUserProps" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps"></a>

Returns, adds, edits, and removes database users.

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.Initializer"></a>

```typescript
import { CfnDatabaseUserProps } from '@mongodb-cdk/atlas-database-user'

const cfnDatabaseUserProps: CfnDatabaseUserProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.databaseName">databaseName</a></code> | <code>string</code> | MongoDB database against which the MongoDB database user authenticates. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.projectId">projectId</a></code> | <code>string</code> | Unique identifier of the Atlas project to which the user belongs. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.roles">roles</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.RoleDefinition">RoleDefinition</a>[]</code> | List that provides the pairings of one role with one applicable database. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.username">username</a></code> | <code>string</code> | Human-readable label that represents the user that authenticates to MongoDB. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.awsiamType">awsiamType</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType">CfnDatabaseUserPropsAwsiamType</a></code> | Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user's role. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.deleteAfterDate">deleteAfterDate</a></code> | <code>string</code> | Date and time when MongoDB Cloud deletes the user. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.labels">labels</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.LabelDefinition">LabelDefinition</a>[]</code> | List that contains the key-value pairs for tagging and categorizing the MongoDB database user. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.ldapAuthType">ldapAuthType</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType">CfnDatabaseUserPropsLdapAuthType</a></code> | Method by which the provided username is authenticated. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.password">password</a></code> | <code>string</code> | The user’s password. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.scopes">scopes</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.ScopeDefinition">ScopeDefinition</a>[]</code> | List that contains clusters and MongoDB Atlas Data Lakes that this database user can access. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.x509Type">x509Type</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type">CfnDatabaseUserPropsX509Type</a></code> | Method that briefs who owns the certificate provided. |

---

##### `databaseName`<sup>Required</sup> <a name="databaseName" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.databaseName"></a>

```typescript
public readonly databaseName: string;
```

- *Type:* string

MongoDB database against which the MongoDB database user authenticates.

MongoDB database users must provide both a username and authentication database to log into MongoDB.

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique identifier of the Atlas project to which the user belongs.

---

##### `roles`<sup>Required</sup> <a name="roles" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.roles"></a>

```typescript
public readonly roles: RoleDefinition[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.RoleDefinition">RoleDefinition</a>[]

List that provides the pairings of one role with one applicable database.

---

##### `username`<sup>Required</sup> <a name="username" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.username"></a>

```typescript
public readonly username: string;
```

- *Type:* string

Human-readable label that represents the user that authenticates to MongoDB.

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `awsiamType`<sup>Optional</sup> <a name="awsiamType" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.awsiamType"></a>

```typescript
public readonly awsiamType: CfnDatabaseUserPropsAwsiamType;
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType">CfnDatabaseUserPropsAwsiamType</a>

Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user's role.

---

##### `deleteAfterDate`<sup>Optional</sup> <a name="deleteAfterDate" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.deleteAfterDate"></a>

```typescript
public readonly deleteAfterDate: string;
```

- *Type:* string

Date and time when MongoDB Cloud deletes the user.

This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. You must specify a future date that falls within one week of making the Application Programming Interface (API) request.

---

##### `labels`<sup>Optional</sup> <a name="labels" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.labels"></a>

```typescript
public readonly labels: LabelDefinition[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.LabelDefinition">LabelDefinition</a>[]

List that contains the key-value pairs for tagging and categorizing the MongoDB database user.

The labels that you define do not appear in the console.

---

##### `ldapAuthType`<sup>Optional</sup> <a name="ldapAuthType" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.ldapAuthType"></a>

```typescript
public readonly ldapAuthType: CfnDatabaseUserPropsLdapAuthType;
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType">CfnDatabaseUserPropsLdapAuthType</a>

Method by which the provided username is authenticated.

If no value is given, Atlas uses the default value of NONE.

---

##### `password`<sup>Optional</sup> <a name="password" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.password"></a>

```typescript
public readonly password: string;
```

- *Type:* string

The user’s password.

This field is not included in the entity returned from the server.

---

##### `scopes`<sup>Optional</sup> <a name="scopes" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.scopes"></a>

```typescript
public readonly scopes: ScopeDefinition[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.ScopeDefinition">ScopeDefinition</a>[]

List that contains clusters and MongoDB Atlas Data Lakes that this database user can access.

If omitted, MongoDB Cloud grants the database user access to all the clusters and MongoDB Atlas Data Lakes in the project.

---

##### `x509Type`<sup>Optional</sup> <a name="x509Type" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserProps.property.x509Type"></a>

```typescript
public readonly x509Type: CfnDatabaseUserPropsX509Type;
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type">CfnDatabaseUserPropsX509Type</a>

Method that briefs who owns the certificate provided.

If no value is given while using X509Type, Atlas uses the default value of MANAGED.

---

### LabelDefinition <a name="LabelDefinition" id="@mongodb-cdk/atlas-database-user.LabelDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-database-user.LabelDefinition.Initializer"></a>

```typescript
import { LabelDefinition } from '@mongodb-cdk/atlas-database-user'

const labelDefinition: LabelDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.LabelDefinition.property.key">key</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-database-user.LabelDefinition.property.value">value</a></code> | <code>string</code> | *No description.* |

---

##### `key`<sup>Optional</sup> <a name="key" id="@mongodb-cdk/atlas-database-user.LabelDefinition.property.key"></a>

```typescript
public readonly key: string;
```

- *Type:* string

---

##### `value`<sup>Optional</sup> <a name="value" id="@mongodb-cdk/atlas-database-user.LabelDefinition.property.value"></a>

```typescript
public readonly value: string;
```

- *Type:* string

---

### RoleDefinition <a name="RoleDefinition" id="@mongodb-cdk/atlas-database-user.RoleDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-database-user.RoleDefinition.Initializer"></a>

```typescript
import { RoleDefinition } from '@mongodb-cdk/atlas-database-user'

const roleDefinition: RoleDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.RoleDefinition.property.collectionName">collectionName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-database-user.RoleDefinition.property.databaseName">databaseName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-database-user.RoleDefinition.property.roleName">roleName</a></code> | <code>string</code> | *No description.* |

---

##### `collectionName`<sup>Optional</sup> <a name="collectionName" id="@mongodb-cdk/atlas-database-user.RoleDefinition.property.collectionName"></a>

```typescript
public readonly collectionName: string;
```

- *Type:* string

---

##### `databaseName`<sup>Optional</sup> <a name="databaseName" id="@mongodb-cdk/atlas-database-user.RoleDefinition.property.databaseName"></a>

```typescript
public readonly databaseName: string;
```

- *Type:* string

---

##### `roleName`<sup>Optional</sup> <a name="roleName" id="@mongodb-cdk/atlas-database-user.RoleDefinition.property.roleName"></a>

```typescript
public readonly roleName: string;
```

- *Type:* string

---

### ScopeDefinition <a name="ScopeDefinition" id="@mongodb-cdk/atlas-database-user.ScopeDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-database-user.ScopeDefinition.Initializer"></a>

```typescript
import { ScopeDefinition } from '@mongodb-cdk/atlas-database-user'

const scopeDefinition: ScopeDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.ScopeDefinition.property.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-database-user.ScopeDefinition.property.type">type</a></code> | <code><a href="#@mongodb-cdk/atlas-database-user.ScopeDefinitionType">ScopeDefinitionType</a></code> | *No description.* |

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodb-cdk/atlas-database-user.ScopeDefinition.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodb-cdk/atlas-database-user.ScopeDefinition.property.type"></a>

```typescript
public readonly type: ScopeDefinitionType;
```

- *Type:* <a href="#@mongodb-cdk/atlas-database-user.ScopeDefinitionType">ScopeDefinitionType</a>

---



## Enums <a name="Enums" id="Enums"></a>

### CfnDatabaseUserPropsAwsiamType <a name="CfnDatabaseUserPropsAwsiamType" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType"></a>

Human-readable label that indicates whether the new database user authenticates with the Amazon Web Services (AWS) Identity and Access Management (IAM) credentials associated with the user or the user's role.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType.NONE">NONE</a></code> | NONE. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType.USER">USER</a></code> | USER. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType.ROLE">ROLE</a></code> | ROLE. |

---

##### `NONE` <a name="NONE" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType.NONE"></a>

NONE.

---


##### `USER` <a name="USER" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType.USER"></a>

USER.

---


##### `ROLE` <a name="ROLE" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsAwsiamType.ROLE"></a>

ROLE.

---


### CfnDatabaseUserPropsLdapAuthType <a name="CfnDatabaseUserPropsLdapAuthType" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType"></a>

Method by which the provided username is authenticated.

If no value is given, Atlas uses the default value of NONE.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType.NONE">NONE</a></code> | NONE. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType.USER">USER</a></code> | USER. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType.GROUP">GROUP</a></code> | GROUP. |

---

##### `NONE` <a name="NONE" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType.NONE"></a>

NONE.

---


##### `USER` <a name="USER" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType.USER"></a>

USER.

---


##### `GROUP` <a name="GROUP" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsLdapAuthType.GROUP"></a>

GROUP.

---


### CfnDatabaseUserPropsX509Type <a name="CfnDatabaseUserPropsX509Type" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type"></a>

Method that briefs who owns the certificate provided.

If no value is given while using X509Type, Atlas uses the default value of MANAGED.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type.NONE">NONE</a></code> | NONE. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type.MANAGED">MANAGED</a></code> | MANAGED. |
| <code><a href="#@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type.CUSTOMER">CUSTOMER</a></code> | CUSTOMER. |

---

##### `NONE` <a name="NONE" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type.NONE"></a>

NONE.

---


##### `MANAGED` <a name="MANAGED" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type.MANAGED"></a>

MANAGED.

---


##### `CUSTOMER` <a name="CUSTOMER" id="@mongodb-cdk/atlas-database-user.CfnDatabaseUserPropsX509Type.CUSTOMER"></a>

CUSTOMER.

---


### ScopeDefinitionType <a name="ScopeDefinitionType" id="@mongodb-cdk/atlas-database-user.ScopeDefinitionType"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-database-user.ScopeDefinitionType.CLUSTER">CLUSTER</a></code> | CLUSTER. |
| <code><a href="#@mongodb-cdk/atlas-database-user.ScopeDefinitionType.DATA_LAKE">DATA_LAKE</a></code> | DATA_LAKE. |

---

##### `CLUSTER` <a name="CLUSTER" id="@mongodb-cdk/atlas-database-user.ScopeDefinitionType.CLUSTER"></a>

CLUSTER.

---


##### `DATA_LAKE` <a name="DATA_LAKE" id="@mongodb-cdk/atlas-database-user.ScopeDefinitionType.DATA_LAKE"></a>

DATA_LAKE.

---

