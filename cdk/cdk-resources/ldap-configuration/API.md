# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnLdapConfiguration <a name="CfnLdapConfiguration" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration"></a>

A CloudFormation `MongoDB::Atlas::LDAPConfiguration`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.Initializer"></a>

```typescript
import { CfnLdapConfiguration } from '@mongodbatlas-awscdk/ldap-configuration'

new CfnLdapConfiguration(scope: Construct, id: string, props: CfnLdapConfigurationProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps">CfnLdapConfigurationProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps">CfnLdapConfigurationProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isConstruct"></a>

```typescript
import { CfnLdapConfiguration } from '@mongodbatlas-awscdk/ldap-configuration'

CfnLdapConfiguration.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isCfnElement"></a>

```typescript
import { CfnLdapConfiguration } from '@mongodbatlas-awscdk/ldap-configuration'

CfnLdapConfiguration.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isCfnResource"></a>

```typescript
import { CfnLdapConfiguration } from '@mongodbatlas-awscdk/ldap-configuration'

CfnLdapConfiguration.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.attrGroupId">attrGroupId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::LDAPConfiguration.GroupId`. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps">CfnLdapConfigurationProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrGroupId`<sup>Required</sup> <a name="attrGroupId" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.attrGroupId"></a>

```typescript
public readonly attrGroupId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::LDAPConfiguration.GroupId`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.props"></a>

```typescript
public readonly props: CfnLdapConfigurationProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps">CfnLdapConfigurationProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfiguration.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiAtlasNdsUserToDnMappingView <a name="ApiAtlasNdsUserToDnMappingView" id="@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView.Initializer"></a>

```typescript
import { ApiAtlasNdsUserToDnMappingView } from '@mongodbatlas-awscdk/ldap-configuration'

const apiAtlasNdsUserToDnMappingView: ApiAtlasNdsUserToDnMappingView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView.property.ldapQuery">ldapQuery</a></code> | <code>string</code> | Lightweight Directory Access Protocol (LDAP) query template that inserts the LDAP name that the regular expression matches into an LDAP query Uniform Resource Identifier (URI). |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView.property.match">match</a></code> | <code>string</code> | Regular expression that MongoDB Cloud uses to match against the provided Lightweight Directory Access Protocol (LDAP) username. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView.property.substitution">substitution</a></code> | <code>string</code> | Lightweight Directory Access Protocol (LDAP) Distinguished Name (DN) template that converts the LDAP username that matches regular expression in the *match* parameter into an LDAP Distinguished Name (DN). |

---

##### `ldapQuery`<sup>Optional</sup> <a name="ldapQuery" id="@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView.property.ldapQuery"></a>

```typescript
public readonly ldapQuery: string;
```

- *Type:* string

Lightweight Directory Access Protocol (LDAP) query template that inserts the LDAP name that the regular expression matches into an LDAP query Uniform Resource Identifier (URI).

The formatting for the query must conform to [RFC 4515](https://datatracker.ietf.org/doc/html/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516).

---

##### `match`<sup>Optional</sup> <a name="match" id="@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView.property.match"></a>

```typescript
public readonly match: string;
```

- *Type:* string

Regular expression that MongoDB Cloud uses to match against the provided Lightweight Directory Access Protocol (LDAP) username.

Each parenthesis-enclosed section represents a regular expression capture group that the substitution or `ldapQuery` template uses.

---

##### `substitution`<sup>Optional</sup> <a name="substitution" id="@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView.property.substitution"></a>

```typescript
public readonly substitution: string;
```

- *Type:* string

Lightweight Directory Access Protocol (LDAP) Distinguished Name (DN) template that converts the LDAP username that matches regular expression in the *match* parameter into an LDAP Distinguished Name (DN).

---

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/ldap-configuration.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/ldap-configuration.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/ldap-configuration'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/ldap-configuration.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/ldap-configuration.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnLdapConfigurationProps <a name="CfnLdapConfigurationProps" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps"></a>

Returns, edits, verifies, and removes LDAP configurations.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.Initializer"></a>

```typescript
import { CfnLdapConfigurationProps } from '@mongodbatlas-awscdk/ldap-configuration'

const cfnLdapConfigurationProps: CfnLdapConfigurationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/ldap-configuration.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.bindPassword">bindPassword</a></code> | <code>string</code> | Password that MongoDB Cloud uses to authenticate the **bindUsername**. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.bindUsername">bindUsername</a></code> | <code>string</code> | Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.hostname">hostname</a></code> | <code>string</code> | Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.port">port</a></code> | <code>number</code> | Port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.authenticationEnabled">authenticationEnabled</a></code> | <code>boolean</code> | Flag that indicates whether users can authenticate using an Lightweight Directory Access Protocol (LDAP) host. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.authorizationEnabled">authorizationEnabled</a></code> | <code>boolean</code> | Flag that indicates whether users can authorize access to MongoDB Cloud resources using an Lightweight Directory Access Protocol (LDAP) host. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.authzQueryTemplate">authzQueryTemplate</a></code> | <code>string</code> | Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud runs to obtain the LDAP groups associated with the authenticated user. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.caCertificate">caCertificate</a></code> | <code>string</code> | Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.status">status</a></code> | <code>string</code> | The current status of the LDAP over TLS/SSL configuration. |
| <code><a href="#@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.userToDnMapping">userToDnMapping</a></code> | <code><a href="#@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView">ApiAtlasNdsUserToDnMappingView</a>[]</code> | User-to-Distinguished Name (DN) map that MongoDB Cloud uses to transform a Lightweight Directory Access Protocol (LDAP) username into an LDAP DN. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/ldap-configuration.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `bindPassword`<sup>Required</sup> <a name="bindPassword" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.bindPassword"></a>

```typescript
public readonly bindPassword: string;
```

- *Type:* string

Password that MongoDB Cloud uses to authenticate the **bindUsername**.

---

##### `bindUsername`<sup>Required</sup> <a name="bindUsername" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.bindUsername"></a>

```typescript
public readonly bindUsername: string;
```

- *Type:* string

Full Distinguished Name (DN) of the Lightweight Directory Access Protocol (LDAP) user that MongoDB Cloud uses to connect to the LDAP host.

LDAP distinguished names must be formatted according to RFC 2253.

---

##### `hostname`<sup>Required</sup> <a name="hostname" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.hostname"></a>

```typescript
public readonly hostname: string;
```

- *Type:* string

Human-readable label that identifies the hostname or Internet Protocol (IP) address of the Lightweight Directory Access Protocol (LDAP) host.

This host must have access to the internet or have a Virtual Private Cloud (VPC) peering connection to your cluster.

---

##### `port`<sup>Required</sup> <a name="port" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.port"></a>

```typescript
public readonly port: number;
```

- *Type:* number

Port to which the Lightweight Directory Access Protocol (LDAP) host listens for client connections.

---

##### `authenticationEnabled`<sup>Optional</sup> <a name="authenticationEnabled" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.authenticationEnabled"></a>

```typescript
public readonly authenticationEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether users can authenticate using an Lightweight Directory Access Protocol (LDAP) host.

---

##### `authorizationEnabled`<sup>Optional</sup> <a name="authorizationEnabled" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.authorizationEnabled"></a>

```typescript
public readonly authorizationEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether users can authorize access to MongoDB Cloud resources using an Lightweight Directory Access Protocol (LDAP) host.

---

##### `authzQueryTemplate`<sup>Optional</sup> <a name="authzQueryTemplate" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.authzQueryTemplate"></a>

```typescript
public readonly authzQueryTemplate: string;
```

- *Type:* string

Lightweight Directory Access Protocol (LDAP) query template that MongoDB Cloud runs to obtain the LDAP groups associated with the authenticated user.

MongoDB Cloud uses this parameter only for user authorization. Use the `{USER}` placeholder in the Uniform Resource Locator (URL) to substitute the authenticated username. The query relates to the host specified with the hostname. Format this query according to [RFC 4515](https://tools.ietf.org/search/rfc4515) and [RFC 4516](https://datatracker.ietf.org/doc/html/rfc4516).

---

##### `caCertificate`<sup>Optional</sup> <a name="caCertificate" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.caCertificate"></a>

```typescript
public readonly caCertificate: string;
```

- *Type:* string

Certificate Authority (CA) certificate that MongoDB Cloud uses to verify the identity of the Lightweight Directory Access Protocol (LDAP) host.

MongoDB Cloud allows self-signed certificates. To delete an assigned value, pass an empty string: `"caCertificate": ""`

---

##### `status`<sup>Optional</sup> <a name="status" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.status"></a>

```typescript
public readonly status: string;
```

- *Type:* string

The current status of the LDAP over TLS/SSL configuration.

---

##### `userToDnMapping`<sup>Optional</sup> <a name="userToDnMapping" id="@mongodbatlas-awscdk/ldap-configuration.CfnLdapConfigurationProps.property.userToDnMapping"></a>

```typescript
public readonly userToDnMapping: ApiAtlasNdsUserToDnMappingView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/ldap-configuration.ApiAtlasNdsUserToDnMappingView">ApiAtlasNdsUserToDnMappingView</a>[]

User-to-Distinguished Name (DN) map that MongoDB Cloud uses to transform a Lightweight Directory Access Protocol (LDAP) username into an LDAP DN.

---



