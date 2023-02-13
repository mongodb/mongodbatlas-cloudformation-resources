# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnLdapVerify <a name="CfnLdapVerify" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify"></a>

A CloudFormation `MongoDB::Atlas::LDAPVerify`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.Initializer"></a>

```typescript
import { CfnLdapVerify } from '@mongodbatlas-awscdk/ldap-verify'

new CfnLdapVerify(scope: Construct, id: string, props: CfnLdapVerifyProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps">CfnLdapVerifyProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps">CfnLdapVerifyProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isConstruct"></a>

```typescript
import { CfnLdapVerify } from '@mongodbatlas-awscdk/ldap-verify'

CfnLdapVerify.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isCfnElement"></a>

```typescript
import { CfnLdapVerify } from '@mongodbatlas-awscdk/ldap-verify'

CfnLdapVerify.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isCfnResource"></a>

```typescript
import { CfnLdapVerify } from '@mongodbatlas-awscdk/ldap-verify'

CfnLdapVerify.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.attrRequestId">attrRequestId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::LDAPVerify.RequestId`. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps">CfnLdapVerifyProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrRequestId`<sup>Required</sup> <a name="attrRequestId" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.attrRequestId"></a>

```typescript
public readonly attrRequestId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::LDAPVerify.RequestId`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.props"></a>

```typescript
public readonly props: CfnLdapVerifyProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps">CfnLdapVerifyProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerify.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/ldap-verify.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/ldap-verify.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/ldap-verify'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/ldap-verify.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/ldap-verify.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnLdapVerifyProps <a name="CfnLdapVerifyProps" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps"></a>

Requests a verification of an LDAP configuration over TLS for an Atlas project.

Pass the requestId in the response object to the Verify |ldap| Configuration endpoint to get the status of a verification request. Atlas retains only the most recent request for each project.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.Initializer"></a>

```typescript
import { CfnLdapVerifyProps } from '@mongodbatlas-awscdk/ldap-verify'

const cfnLdapVerifyProps: CfnLdapVerifyProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/ldap-verify.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.bindPassword">bindPassword</a></code> | <code>string</code> | Password that MongoDB Cloud uses to authenticate the **bindUsername**. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.bindUsername">bindUsername</a></code> | <code>string</code> | Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.hostName">hostName</a></code> | <code>string</code> | Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.port">port</a></code> | <code>number</code> | Port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.authzQueryTemplate">authzQueryTemplate</a></code> | <code>string</code> | Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud runs to obtain the LDAP groups associated with the authenticated user. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.caCertificate">caCertificate</a></code> | <code>string</code> | Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.groupId">groupId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.status">status</a></code> | <code>string</code> | The current status of the LDAP over TLS/SSL configuration. |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.validations">validations</a></code> | <code><a href="#@mongodbatlas-awscdk/ldap-verify.Validation">Validation</a>[]</code> | List of validation messages related to the verification of the provided LDAP over TLS configuration details. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/ldap-verify.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `bindPassword`<sup>Required</sup> <a name="bindPassword" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.bindPassword"></a>

```typescript
public readonly bindPassword: string;
```

- *Type:* string

Password that MongoDB Cloud uses to authenticate the **bindUsername**.

---

##### `bindUsername`<sup>Required</sup> <a name="bindUsername" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.bindUsername"></a>

```typescript
public readonly bindUsername: string;
```

- *Type:* string

Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host.

LDAP distinguished names must be formatted according to RFC 2253.

---

##### `hostName`<sup>Required</sup> <a name="hostName" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.hostName"></a>

```typescript
public readonly hostName: string;
```

- *Type:* string

Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host.

This host must have access to the internet or have a Virtual Private Cloud (VPC) peering connection to your cluster.

---

##### `port`<sup>Required</sup> <a name="port" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.port"></a>

```typescript
public readonly port: number;
```

- *Type:* number

Port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections.

---

##### `authzQueryTemplate`<sup>Optional</sup> <a name="authzQueryTemplate" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.authzQueryTemplate"></a>

```typescript
public readonly authzQueryTemplate: string;
```

- *Type:* string

Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud runs to obtain the LDAP groups associated with the authenticated user.

MongoDB Cloud uses this parameter only for user authorization. Use the `{USER}` placeholder in the Uniform Resource Locator (URL) to substitute the authenticated username. The query relates to the host specified with the hostname. Format this query according to [RFC 4515](https://tools.ietf.org/search/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516).

---

##### `caCertificate`<sup>Optional</sup> <a name="caCertificate" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.caCertificate"></a>

```typescript
public readonly caCertificate: string;
```

- *Type:* string

Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host.

MongoDB Cloud allows self-signed certificates. To delete an assigned value, pass an empty string: `"caCertificate": ""`

---

##### `groupId`<sup>Optional</sup> <a name="groupId" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.groupId"></a>

```typescript
public readonly groupId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `status`<sup>Optional</sup> <a name="status" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.status"></a>

```typescript
public readonly status: string;
```

- *Type:* string

The current status of the LDAP over TLS/SSL configuration.

---

##### `validations`<sup>Optional</sup> <a name="validations" id="@mongodbatlas-awscdk/ldap-verify.CfnLdapVerifyProps.property.validations"></a>

```typescript
public readonly validations: Validation[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/ldap-verify.Validation">Validation</a>[]

List of validation messages related to the verification of the provided LDAP over TLS configuration details.

The array contains a document for each test that Atlas runs. Atlas stops running tests after the first failure.

---

### Validation <a name="Validation" id="@mongodbatlas-awscdk/ldap-verify.Validation"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/ldap-verify.Validation.Initializer"></a>

```typescript
import { Validation } from '@mongodbatlas-awscdk/ldap-verify'

const validation: Validation = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.Validation.property.status">status</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/ldap-verify.Validation.property.validationType">validationType</a></code> | <code>string</code> | *No description.* |

---

##### `status`<sup>Optional</sup> <a name="status" id="@mongodbatlas-awscdk/ldap-verify.Validation.property.status"></a>

```typescript
public readonly status: string;
```

- *Type:* string

---

##### `validationType`<sup>Optional</sup> <a name="validationType" id="@mongodbatlas-awscdk/ldap-verify.Validation.property.validationType"></a>

```typescript
public readonly validationType: string;
```

- *Type:* string

---



