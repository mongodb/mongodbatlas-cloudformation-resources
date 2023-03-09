# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnX509AuthenticationDatabaseUser <a name="CfnX509AuthenticationDatabaseUser" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser"></a>

A CloudFormation `MongoDB::Atlas::X509AuthenticationDatabaseUser`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.Initializer"></a>

```typescript
import { CfnX509AuthenticationDatabaseUser } from '@mongodbatlas-awscdk/x509-authentication-database-user'

new CfnX509AuthenticationDatabaseUser(scope: Construct, id: string, props: CfnX509AuthenticationDatabaseUserProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps">CfnX509AuthenticationDatabaseUserProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps">CfnX509AuthenticationDatabaseUserProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isConstruct"></a>

```typescript
import { CfnX509AuthenticationDatabaseUser } from '@mongodbatlas-awscdk/x509-authentication-database-user'

CfnX509AuthenticationDatabaseUser.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isCfnElement"></a>

```typescript
import { CfnX509AuthenticationDatabaseUser } from '@mongodbatlas-awscdk/x509-authentication-database-user'

CfnX509AuthenticationDatabaseUser.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isCfnResource"></a>

```typescript
import { CfnX509AuthenticationDatabaseUser } from '@mongodbatlas-awscdk/x509-authentication-database-user'

CfnX509AuthenticationDatabaseUser.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.attrLinks">attrLinks</a></code> | <code>any[]</code> | Attribute `MongoDB::Atlas::X509AuthenticationDatabaseUser.Links`. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.attrResults">attrResults</a></code> | <code>any[]</code> | Attribute `MongoDB::Atlas::X509AuthenticationDatabaseUser.Results`. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps">CfnX509AuthenticationDatabaseUserProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrLinks`<sup>Required</sup> <a name="attrLinks" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.attrLinks"></a>

```typescript
public readonly attrLinks: any[];
```

- *Type:* any[]

Attribute `MongoDB::Atlas::X509AuthenticationDatabaseUser.Links`.

---

##### `attrResults`<sup>Required</sup> <a name="attrResults" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.attrResults"></a>

```typescript
public readonly attrResults: any[];
```

- *Type:* any[]

Attribute `MongoDB::Atlas::X509AuthenticationDatabaseUser.Results`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.props"></a>

```typescript
public readonly props: CfnX509AuthenticationDatabaseUserProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps">CfnX509AuthenticationDatabaseUserProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUser.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnX509AuthenticationDatabaseUserProps <a name="CfnX509AuthenticationDatabaseUserProps" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps"></a>

Generates one X.509 certificate for the specified MongoDB user. Atlas manages the certificate and MongoDB user that belong to one project. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

To get MongoDB Cloud to generate a managed certificate for a database user, set `"x509Type" : "MANAGED"` on the desired MongoDB Database User.

If you are managing your own Certificate Authority (CA) in Self-Managed X.509 mode, you must generate certificates for database users using your own CA.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.Initializer"></a>

```typescript
import { CfnX509AuthenticationDatabaseUserProps } from '@mongodbatlas-awscdk/x509-authentication-database-user'

const cfnX509AuthenticationDatabaseUserProps: CfnX509AuthenticationDatabaseUserProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.customerX509">customerX509</a></code> | <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CustomerX509">CustomerX509</a></code> | CustomerX509 represents Customer-managed X.509 configuration for an Atlas project. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.monthsUntilExpiration">monthsUntilExpiration</a></code> | <code>number</code> | A number of months that the created certificate is valid for before expiry, up to 24 months.default 3. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.profile">profile</a></code> | <code>string</code> | Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.projectId">projectId</a></code> | <code>string</code> | The unique identifier for the project . |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.totalCount">totalCount</a></code> | <code>number</code> | Total number of unexpired certificates returned in this response. |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.userName">userName</a></code> | <code>string</code> | Username of the database user to create a certificate for. |

---

##### `customerX509`<sup>Optional</sup> <a name="customerX509" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.customerX509"></a>

```typescript
public readonly customerX509: CustomerX509;
```

- *Type:* <a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CustomerX509">CustomerX509</a>

CustomerX509 represents Customer-managed X.509 configuration for an Atlas project.

---

##### `monthsUntilExpiration`<sup>Optional</sup> <a name="monthsUntilExpiration" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.monthsUntilExpiration"></a>

```typescript
public readonly monthsUntilExpiration: number;
```

- *Type:* number

A number of months that the created certificate is valid for before expiry, up to 24 months.default 3.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

The unique identifier for the project .

---

##### `totalCount`<sup>Optional</sup> <a name="totalCount" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.totalCount"></a>

```typescript
public readonly totalCount: number;
```

- *Type:* number

Total number of unexpired certificates returned in this response.

---

##### `userName`<sup>Optional</sup> <a name="userName" id="@mongodbatlas-awscdk/x509-authentication-database-user.CfnX509AuthenticationDatabaseUserProps.property.userName"></a>

```typescript
public readonly userName: string;
```

- *Type:* string

Username of the database user to create a certificate for.

---

### CustomerX509 <a name="CustomerX509" id="@mongodbatlas-awscdk/x509-authentication-database-user.CustomerX509"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/x509-authentication-database-user.CustomerX509.Initializer"></a>

```typescript
import { CustomerX509 } from '@mongodbatlas-awscdk/x509-authentication-database-user'

const customerX509: CustomerX509 = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/x509-authentication-database-user.CustomerX509.property.cas">cas</a></code> | <code>string</code> | PEM string containing one or more customer CAs for database user authentication. |

---

##### `cas`<sup>Optional</sup> <a name="cas" id="@mongodbatlas-awscdk/x509-authentication-database-user.CustomerX509.property.cas"></a>

```typescript
public readonly cas: string;
```

- *Type:* string

PEM string containing one or more customer CAs for database user authentication.

---



