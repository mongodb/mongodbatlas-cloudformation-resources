# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnFederatedSettingsOrgRoleMapping <a name="CfnFederatedSettingsOrgRoleMapping" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping"></a>

A CloudFormation `MongoDB::Atlas::FederatedSettingsOrgRoleMapping`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.Initializer"></a>

```typescript
import { CfnFederatedSettingsOrgRoleMapping } from '@mongodbatlas-awscdk/federated-settings-org-role-mapping'

new CfnFederatedSettingsOrgRoleMapping(scope: Construct, id: string, props: CfnFederatedSettingsOrgRoleMappingProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps">CfnFederatedSettingsOrgRoleMappingProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps">CfnFederatedSettingsOrgRoleMappingProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isConstruct"></a>

```typescript
import { CfnFederatedSettingsOrgRoleMapping } from '@mongodbatlas-awscdk/federated-settings-org-role-mapping'

CfnFederatedSettingsOrgRoleMapping.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isCfnElement"></a>

```typescript
import { CfnFederatedSettingsOrgRoleMapping } from '@mongodbatlas-awscdk/federated-settings-org-role-mapping'

CfnFederatedSettingsOrgRoleMapping.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isCfnResource"></a>

```typescript
import { CfnFederatedSettingsOrgRoleMapping } from '@mongodbatlas-awscdk/federated-settings-org-role-mapping'

CfnFederatedSettingsOrgRoleMapping.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::FederatedSettingsOrgRoleMapping.Id`. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps">CfnFederatedSettingsOrgRoleMappingProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::FederatedSettingsOrgRoleMapping.Id`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.props"></a>

```typescript
public readonly props: CfnFederatedSettingsOrgRoleMappingProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps">CfnFederatedSettingsOrgRoleMappingProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMapping.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/federated-settings-org-role-mapping'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnFederatedSettingsOrgRoleMappingProps <a name="CfnFederatedSettingsOrgRoleMappingProps" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps"></a>

Returns, adds, edits, and removes federation-related features such as role mappings and connected organization configurations.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.Initializer"></a>

```typescript
import { CfnFederatedSettingsOrgRoleMappingProps } from '@mongodbatlas-awscdk/federated-settings-org-role-mapping'

const cfnFederatedSettingsOrgRoleMappingProps: CfnFederatedSettingsOrgRoleMappingProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.federationSettingsId">federationSettingsId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your federation. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.orgId">orgId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.externalGroupName">externalGroupName</a></code> | <code>string</code> | Unique human-readable label that identifies the identity provider group to whichthis role mapping applies. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.roleAssignments">roleAssignments</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment">RoleAssignment</a>[]</code> | Atlas roles and the unique identifiers of the groups and organizations associated with each role. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `federationSettingsId`<sup>Required</sup> <a name="federationSettingsId" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.federationSettingsId"></a>

```typescript
public readonly federationSettingsId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your federation.

---

##### `orgId`<sup>Required</sup> <a name="orgId" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.orgId"></a>

```typescript
public readonly orgId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the organization that contains your projects.

---

##### `externalGroupName`<sup>Optional</sup> <a name="externalGroupName" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.externalGroupName"></a>

```typescript
public readonly externalGroupName: string;
```

- *Type:* string

Unique human-readable label that identifies the identity provider group to whichthis role mapping applies.

---

##### `roleAssignments`<sup>Optional</sup> <a name="roleAssignments" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.CfnFederatedSettingsOrgRoleMappingProps.property.roleAssignments"></a>

```typescript
public readonly roleAssignments: RoleAssignment[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment">RoleAssignment</a>[]

Atlas roles and the unique identifiers of the groups and organizations associated with each role.

---

### RoleAssignment <a name="RoleAssignment" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment.Initializer"></a>

```typescript
import { RoleAssignment } from '@mongodbatlas-awscdk/federated-settings-org-role-mapping'

const roleAssignment: RoleAssignment = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment.property.groupId">groupId</a></code> | <code>string</code> | List that contains comma-separated key value pairs to map zones to geographic regions. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment.property.orgId">orgId</a></code> | <code>string</code> | List that contains comma-separated key value pairs to map zones to geographic regions. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment.property.role">role</a></code> | <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole">RoleAssignmentRole</a></code> | *No description.* |

---

##### `groupId`<sup>Optional</sup> <a name="groupId" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment.property.groupId"></a>

```typescript
public readonly groupId: string;
```

- *Type:* string

List that contains comma-separated key value pairs to map zones to geographic regions.

These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.

This parameter returns an empty object if no custom zones exist.

---

##### `orgId`<sup>Optional</sup> <a name="orgId" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment.property.orgId"></a>

```typescript
public readonly orgId: string;
```

- *Type:* string

List that contains comma-separated key value pairs to map zones to geographic regions.

These pairs map an ISO 3166-1a2 location code, with an ISO 3166-2 subdivision code when possible, to a unique 24-hexadecimal string that identifies the custom zone.

This parameter returns an empty object if no custom zones exist.

---

##### `role`<sup>Optional</sup> <a name="role" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignment.property.role"></a>

```typescript
public readonly role: RoleAssignmentRole;
```

- *Type:* <a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole">RoleAssignmentRole</a>

---



## Enums <a name="Enums" id="Enums"></a>

### RoleAssignmentRole <a name="RoleAssignmentRole" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_AUTOMATION_ADMIN">GLOBAL_AUTOMATION_ADMIN</a></code> | GLOBAL_AUTOMATION_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_BACKUP_ADMIN">GLOBAL_BACKUP_ADMIN</a></code> | GLOBAL_BACKUP_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_METERING_USER">GLOBAL_METERING_USER</a></code> | GLOBAL_METERING_USER. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_METRICS_INTERNAL_USER">GLOBAL_METRICS_INTERNAL_USER</a></code> | GLOBAL_METRICS_INTERNAL_USER. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_MONITORING_ADMIN">GLOBAL_MONITORING_ADMIN</a></code> | GLOBAL_MONITORING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_OWNER">GLOBAL_OWNER</a></code> | GLOBAL_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_READ_ONLY">GLOBAL_READ_ONLY</a></code> | GLOBAL_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_USER_ADMIN">GLOBAL_USER_ADMIN</a></code> | GLOBAL_USER_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_USER_READ_ONLY">GLOBAL_USER_READ_ONLY</a></code> | GLOBAL_USER_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_ACCOUNT_SUSPENSION_ADMIN">GLOBAL_ACCOUNT_SUSPENSION_ADMIN</a></code> | GLOBAL_ACCOUNT_SUSPENSION_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_BILLING_ADMIN">GLOBAL_BILLING_ADMIN</a></code> | GLOBAL_BILLING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_LEGAL_ADMIN">GLOBAL_LEGAL_ADMIN</a></code> | GLOBAL_LEGAL_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_FEATURE_FLAG_ADMIN">GLOBAL_FEATURE_FLAG_ADMIN</a></code> | GLOBAL_FEATURE_FLAG_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_ATLAS_TSE">GLOBAL_ATLAS_TSE</a></code> | GLOBAL_ATLAS_TSE. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_ATLAS_OPERATOR">GLOBAL_ATLAS_OPERATOR</a></code> | GLOBAL_ATLAS_OPERATOR. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_ATLAS_ADMIN">GLOBAL_ATLAS_ADMIN</a></code> | GLOBAL_ATLAS_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_STITCH_ADMIN">GLOBAL_STITCH_ADMIN</a></code> | GLOBAL_STITCH_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_CHARTS_ADMIN">GLOBAL_CHARTS_ADMIN</a></code> | GLOBAL_CHARTS_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_EXPERIMENT_ASSIGNMENT_USER">GLOBAL_EXPERIMENT_ASSIGNMENT_USER</a></code> | GLOBAL_EXPERIMENT_ASSIGNMENT_USER. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_STITCH_INTERNAL_ADMIN">GLOBAL_STITCH_INTERNAL_ADMIN</a></code> | GLOBAL_STITCH_INTERNAL_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_SECURITY_ADMIN">GLOBAL_SECURITY_ADMIN</a></code> | GLOBAL_SECURITY_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_QUERY_ENGINE_INTERNAL_ADMIN">GLOBAL_QUERY_ENGINE_INTERNAL_ADMIN</a></code> | GLOBAL_QUERY_ENGINE_INTERNAL_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_PROACTIVE_SUPPORT_ADMIN">GLOBAL_PROACTIVE_SUPPORT_ADMIN</a></code> | GLOBAL_PROACTIVE_SUPPORT_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_INFRASTRUCTURE_INTERNAL_ADMIN">GLOBAL_INFRASTRUCTURE_INTERNAL_ADMIN</a></code> | GLOBAL_INFRASTRUCTURE_INTERNAL_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_SALESFORCE_ADMIN">GLOBAL_SALESFORCE_ADMIN</a></code> | GLOBAL_SALESFORCE_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_SALESFORCE_READ_ONLY">GLOBAL_SALESFORCE_READ_ONLY</a></code> | GLOBAL_SALESFORCE_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_APP_SERVICES_CLUSTER_DEBUG_DATA_ACCESS">GLOBAL_APP_SERVICES_CLUSTER_DEBUG_DATA_ACCESS</a></code> | GLOBAL_APP_SERVICES_CLUSTER_DEBUG_DATA_ACCESS. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_MEMBER">ORG_MEMBER</a></code> | ORG_MEMBER. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_READ_ONLY">ORG_READ_ONLY</a></code> | ORG_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_BILLING_ADMIN">ORG_BILLING_ADMIN</a></code> | ORG_BILLING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_GROUP_CREATOR">ORG_GROUP_CREATOR</a></code> | ORG_GROUP_CREATOR. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_OWNER">ORG_OWNER</a></code> | ORG_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_AUTOMATION_ADMIN">GROUP_AUTOMATION_ADMIN</a></code> | GROUP_AUTOMATION_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_BACKUP_ADMIN">GROUP_BACKUP_ADMIN</a></code> | GROUP_BACKUP_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_MONITORING_ADMIN">GROUP_MONITORING_ADMIN</a></code> | GROUP_MONITORING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_OWNER">GROUP_OWNER</a></code> | GROUP_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_READ_ONLY">GROUP_READ_ONLY</a></code> | GROUP_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_USER_ADMIN">GROUP_USER_ADMIN</a></code> | GROUP_USER_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_BILLING_ADMIN">GROUP_BILLING_ADMIN</a></code> | GROUP_BILLING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_DATA_ACCESS_ADMIN">GROUP_DATA_ACCESS_ADMIN</a></code> | GROUP_DATA_ACCESS_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_DATA_ACCESS_READ_ONLY">GROUP_DATA_ACCESS_READ_ONLY</a></code> | GROUP_DATA_ACCESS_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_DATA_ACCESS_READ_WRITE">GROUP_DATA_ACCESS_READ_WRITE</a></code> | GROUP_DATA_ACCESS_READ_WRITE. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_CHARTS_ADMIN">GROUP_CHARTS_ADMIN</a></code> | GROUP_CHARTS_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_CLUSTER_MANAGER">GROUP_CLUSTER_MANAGER</a></code> | GROUP_CLUSTER_MANAGER. |
| <code><a href="#@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_SEARCH_INDEX_EDITOR">GROUP_SEARCH_INDEX_EDITOR</a></code> | GROUP_SEARCH_INDEX_EDITOR. |

---

##### `GLOBAL_AUTOMATION_ADMIN` <a name="GLOBAL_AUTOMATION_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_AUTOMATION_ADMIN"></a>

GLOBAL_AUTOMATION_ADMIN.

---


##### `GLOBAL_BACKUP_ADMIN` <a name="GLOBAL_BACKUP_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_BACKUP_ADMIN"></a>

GLOBAL_BACKUP_ADMIN.

---


##### `GLOBAL_METERING_USER` <a name="GLOBAL_METERING_USER" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_METERING_USER"></a>

GLOBAL_METERING_USER.

---


##### `GLOBAL_METRICS_INTERNAL_USER` <a name="GLOBAL_METRICS_INTERNAL_USER" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_METRICS_INTERNAL_USER"></a>

GLOBAL_METRICS_INTERNAL_USER.

---


##### `GLOBAL_MONITORING_ADMIN` <a name="GLOBAL_MONITORING_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_MONITORING_ADMIN"></a>

GLOBAL_MONITORING_ADMIN.

---


##### `GLOBAL_OWNER` <a name="GLOBAL_OWNER" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_OWNER"></a>

GLOBAL_OWNER.

---


##### `GLOBAL_READ_ONLY` <a name="GLOBAL_READ_ONLY" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_READ_ONLY"></a>

GLOBAL_READ_ONLY.

---


##### `GLOBAL_USER_ADMIN` <a name="GLOBAL_USER_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_USER_ADMIN"></a>

GLOBAL_USER_ADMIN.

---


##### `GLOBAL_USER_READ_ONLY` <a name="GLOBAL_USER_READ_ONLY" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_USER_READ_ONLY"></a>

GLOBAL_USER_READ_ONLY.

---


##### `GLOBAL_ACCOUNT_SUSPENSION_ADMIN` <a name="GLOBAL_ACCOUNT_SUSPENSION_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_ACCOUNT_SUSPENSION_ADMIN"></a>

GLOBAL_ACCOUNT_SUSPENSION_ADMIN.

---


##### `GLOBAL_BILLING_ADMIN` <a name="GLOBAL_BILLING_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_BILLING_ADMIN"></a>

GLOBAL_BILLING_ADMIN.

---


##### `GLOBAL_LEGAL_ADMIN` <a name="GLOBAL_LEGAL_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_LEGAL_ADMIN"></a>

GLOBAL_LEGAL_ADMIN.

---


##### `GLOBAL_FEATURE_FLAG_ADMIN` <a name="GLOBAL_FEATURE_FLAG_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_FEATURE_FLAG_ADMIN"></a>

GLOBAL_FEATURE_FLAG_ADMIN.

---


##### `GLOBAL_ATLAS_TSE` <a name="GLOBAL_ATLAS_TSE" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_ATLAS_TSE"></a>

GLOBAL_ATLAS_TSE.

---


##### `GLOBAL_ATLAS_OPERATOR` <a name="GLOBAL_ATLAS_OPERATOR" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_ATLAS_OPERATOR"></a>

GLOBAL_ATLAS_OPERATOR.

---


##### `GLOBAL_ATLAS_ADMIN` <a name="GLOBAL_ATLAS_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_ATLAS_ADMIN"></a>

GLOBAL_ATLAS_ADMIN.

---


##### `GLOBAL_STITCH_ADMIN` <a name="GLOBAL_STITCH_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_STITCH_ADMIN"></a>

GLOBAL_STITCH_ADMIN.

---


##### `GLOBAL_CHARTS_ADMIN` <a name="GLOBAL_CHARTS_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_CHARTS_ADMIN"></a>

GLOBAL_CHARTS_ADMIN.

---


##### `GLOBAL_EXPERIMENT_ASSIGNMENT_USER` <a name="GLOBAL_EXPERIMENT_ASSIGNMENT_USER" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_EXPERIMENT_ASSIGNMENT_USER"></a>

GLOBAL_EXPERIMENT_ASSIGNMENT_USER.

---


##### `GLOBAL_STITCH_INTERNAL_ADMIN` <a name="GLOBAL_STITCH_INTERNAL_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_STITCH_INTERNAL_ADMIN"></a>

GLOBAL_STITCH_INTERNAL_ADMIN.

---


##### `GLOBAL_SECURITY_ADMIN` <a name="GLOBAL_SECURITY_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_SECURITY_ADMIN"></a>

GLOBAL_SECURITY_ADMIN.

---


##### `GLOBAL_QUERY_ENGINE_INTERNAL_ADMIN` <a name="GLOBAL_QUERY_ENGINE_INTERNAL_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_QUERY_ENGINE_INTERNAL_ADMIN"></a>

GLOBAL_QUERY_ENGINE_INTERNAL_ADMIN.

---


##### `GLOBAL_PROACTIVE_SUPPORT_ADMIN` <a name="GLOBAL_PROACTIVE_SUPPORT_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_PROACTIVE_SUPPORT_ADMIN"></a>

GLOBAL_PROACTIVE_SUPPORT_ADMIN.

---


##### `GLOBAL_INFRASTRUCTURE_INTERNAL_ADMIN` <a name="GLOBAL_INFRASTRUCTURE_INTERNAL_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_INFRASTRUCTURE_INTERNAL_ADMIN"></a>

GLOBAL_INFRASTRUCTURE_INTERNAL_ADMIN.

---


##### `GLOBAL_SALESFORCE_ADMIN` <a name="GLOBAL_SALESFORCE_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_SALESFORCE_ADMIN"></a>

GLOBAL_SALESFORCE_ADMIN.

---


##### `GLOBAL_SALESFORCE_READ_ONLY` <a name="GLOBAL_SALESFORCE_READ_ONLY" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_SALESFORCE_READ_ONLY"></a>

GLOBAL_SALESFORCE_READ_ONLY.

---


##### `GLOBAL_APP_SERVICES_CLUSTER_DEBUG_DATA_ACCESS` <a name="GLOBAL_APP_SERVICES_CLUSTER_DEBUG_DATA_ACCESS" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GLOBAL_APP_SERVICES_CLUSTER_DEBUG_DATA_ACCESS"></a>

GLOBAL_APP_SERVICES_CLUSTER_DEBUG_DATA_ACCESS.

---


##### `ORG_MEMBER` <a name="ORG_MEMBER" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_MEMBER"></a>

ORG_MEMBER.

---


##### `ORG_READ_ONLY` <a name="ORG_READ_ONLY" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_READ_ONLY"></a>

ORG_READ_ONLY.

---


##### `ORG_BILLING_ADMIN` <a name="ORG_BILLING_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_BILLING_ADMIN"></a>

ORG_BILLING_ADMIN.

---


##### `ORG_GROUP_CREATOR` <a name="ORG_GROUP_CREATOR" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_GROUP_CREATOR"></a>

ORG_GROUP_CREATOR.

---


##### `ORG_OWNER` <a name="ORG_OWNER" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.ORG_OWNER"></a>

ORG_OWNER.

---


##### `GROUP_AUTOMATION_ADMIN` <a name="GROUP_AUTOMATION_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_AUTOMATION_ADMIN"></a>

GROUP_AUTOMATION_ADMIN.

---


##### `GROUP_BACKUP_ADMIN` <a name="GROUP_BACKUP_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_BACKUP_ADMIN"></a>

GROUP_BACKUP_ADMIN.

---


##### `GROUP_MONITORING_ADMIN` <a name="GROUP_MONITORING_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_MONITORING_ADMIN"></a>

GROUP_MONITORING_ADMIN.

---


##### `GROUP_OWNER` <a name="GROUP_OWNER" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_OWNER"></a>

GROUP_OWNER.

---


##### `GROUP_READ_ONLY` <a name="GROUP_READ_ONLY" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_READ_ONLY"></a>

GROUP_READ_ONLY.

---


##### `GROUP_USER_ADMIN` <a name="GROUP_USER_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_USER_ADMIN"></a>

GROUP_USER_ADMIN.

---


##### `GROUP_BILLING_ADMIN` <a name="GROUP_BILLING_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_BILLING_ADMIN"></a>

GROUP_BILLING_ADMIN.

---


##### `GROUP_DATA_ACCESS_ADMIN` <a name="GROUP_DATA_ACCESS_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_DATA_ACCESS_ADMIN"></a>

GROUP_DATA_ACCESS_ADMIN.

---


##### `GROUP_DATA_ACCESS_READ_ONLY` <a name="GROUP_DATA_ACCESS_READ_ONLY" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_DATA_ACCESS_READ_ONLY"></a>

GROUP_DATA_ACCESS_READ_ONLY.

---


##### `GROUP_DATA_ACCESS_READ_WRITE` <a name="GROUP_DATA_ACCESS_READ_WRITE" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_DATA_ACCESS_READ_WRITE"></a>

GROUP_DATA_ACCESS_READ_WRITE.

---


##### `GROUP_CHARTS_ADMIN` <a name="GROUP_CHARTS_ADMIN" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_CHARTS_ADMIN"></a>

GROUP_CHARTS_ADMIN.

---


##### `GROUP_CLUSTER_MANAGER` <a name="GROUP_CLUSTER_MANAGER" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_CLUSTER_MANAGER"></a>

GROUP_CLUSTER_MANAGER.

---


##### `GROUP_SEARCH_INDEX_EDITOR` <a name="GROUP_SEARCH_INDEX_EDITOR" id="@mongodbatlas-awscdk/federated-settings-org-role-mapping.RoleAssignmentRole.GROUP_SEARCH_INDEX_EDITOR"></a>

GROUP_SEARCH_INDEX_EDITOR.

---

