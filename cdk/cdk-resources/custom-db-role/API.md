# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnCustomDbRole <a name="CfnCustomDbRole" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole"></a>

A CloudFormation `MongoDB::Atlas::CustomDBRole`.

#### Initializers <a name="Initializers" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.Initializer"></a>

```typescript
import { CfnCustomDbRole } from '@mongodb-cdk/atlas-custom-db-role'

new CfnCustomDbRole(scope: Construct, id: string, props: CfnCustomDbRoleProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps">CfnCustomDbRoleProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps">CfnCustomDbRoleProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isConstruct"></a>

```typescript
import { CfnCustomDbRole } from '@mongodb-cdk/atlas-custom-db-role'

CfnCustomDbRole.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isCfnElement"></a>

```typescript
import { CfnCustomDbRole } from '@mongodb-cdk/atlas-custom-db-role'

CfnCustomDbRole.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isCfnResource"></a>

```typescript
import { CfnCustomDbRole } from '@mongodb-cdk/atlas-custom-db-role'

CfnCustomDbRole.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps">CfnCustomDbRoleProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.props"></a>

```typescript
public readonly props: CfnCustomDbRoleProps;
```

- *Type:* <a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps">CfnCustomDbRoleProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRole.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### Action <a name="Action" id="@mongodb-cdk/atlas-custom-db-role.Action"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-custom-db-role.Action.Initializer"></a>

```typescript
import { Action } from '@mongodb-cdk/atlas-custom-db-role'

const action: Action = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.Action.property.action">action</a></code> | <code>string</code> | Human-readable label that identifies the privilege action. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.Action.property.resources">resources</a></code> | <code><a href="#@mongodb-cdk/atlas-custom-db-role.Resource">Resource</a>[]</code> | List of resources on which you grant the action. |

---

##### `action`<sup>Optional</sup> <a name="action" id="@mongodb-cdk/atlas-custom-db-role.Action.property.action"></a>

```typescript
public readonly action: string;
```

- *Type:* string

Human-readable label that identifies the privilege action.

---

##### `resources`<sup>Optional</sup> <a name="resources" id="@mongodb-cdk/atlas-custom-db-role.Action.property.resources"></a>

```typescript
public readonly resources: Resource[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-custom-db-role.Resource">Resource</a>[]

List of resources on which you grant the action.

---

### ApiKey <a name="ApiKey" id="@mongodb-cdk/atlas-custom-db-role.ApiKey"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-custom-db-role.ApiKey.Initializer"></a>

```typescript
import { ApiKey } from '@mongodb-cdk/atlas-custom-db-role'

const apiKey: ApiKey = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.ApiKey.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.ApiKey.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodb-cdk/atlas-custom-db-role.ApiKey.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodb-cdk/atlas-custom-db-role.ApiKey.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnCustomDbRoleProps <a name="CfnCustomDbRoleProps" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps"></a>

Returns, adds, edits, and removes custom database user privilege roles.

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.Initializer"></a>

```typescript
import { CfnCustomDbRoleProps } from '@mongodb-cdk/atlas-custom-db-role'

const cfnCustomDbRoleProps: CfnCustomDbRoleProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.roleName">roleName</a></code> | <code>string</code> | Human-readable label that identifies the role for the request. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.actions">actions</a></code> | <code><a href="#@mongodb-cdk/atlas-custom-db-role.Action">Action</a>[]</code> | List of the individual privilege actions that the role grants. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodb-cdk/atlas-custom-db-role.ApiKey">ApiKey</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.groupId">groupId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.inheritedRoles">inheritedRoles</a></code> | <code><a href="#@mongodb-cdk/atlas-custom-db-role.InheritedRole">InheritedRole</a>[]</code> | List of the built-in roles that this custom role inherits. |

---

##### `roleName`<sup>Required</sup> <a name="roleName" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.roleName"></a>

```typescript
public readonly roleName: string;
```

- *Type:* string

Human-readable label that identifies the role for the request.

This name must be unique for this custom role in this project.

---

##### `actions`<sup>Optional</sup> <a name="actions" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.actions"></a>

```typescript
public readonly actions: Action[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-custom-db-role.Action">Action</a>[]

List of the individual privilege actions that the role grants.

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKey;
```

- *Type:* <a href="#@mongodb-cdk/atlas-custom-db-role.ApiKey">ApiKey</a>

---

##### `groupId`<sup>Optional</sup> <a name="groupId" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.groupId"></a>

```typescript
public readonly groupId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `inheritedRoles`<sup>Optional</sup> <a name="inheritedRoles" id="@mongodb-cdk/atlas-custom-db-role.CfnCustomDbRoleProps.property.inheritedRoles"></a>

```typescript
public readonly inheritedRoles: InheritedRole[];
```

- *Type:* <a href="#@mongodb-cdk/atlas-custom-db-role.InheritedRole">InheritedRole</a>[]

List of the built-in roles that this custom role inherits.

---

### InheritedRole <a name="InheritedRole" id="@mongodb-cdk/atlas-custom-db-role.InheritedRole"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-custom-db-role.InheritedRole.Initializer"></a>

```typescript
import { InheritedRole } from '@mongodb-cdk/atlas-custom-db-role'

const inheritedRole: InheritedRole = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.InheritedRole.property.db">db</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.InheritedRole.property.role">role</a></code> | <code>string</code> | *No description.* |

---

##### `db`<sup>Optional</sup> <a name="db" id="@mongodb-cdk/atlas-custom-db-role.InheritedRole.property.db"></a>

```typescript
public readonly db: string;
```

- *Type:* string

---

##### `role`<sup>Optional</sup> <a name="role" id="@mongodb-cdk/atlas-custom-db-role.InheritedRole.property.role"></a>

```typescript
public readonly role: string;
```

- *Type:* string

---

### Resource <a name="Resource" id="@mongodb-cdk/atlas-custom-db-role.Resource"></a>

List of resources on which you grant the action.

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-custom-db-role.Resource.Initializer"></a>

```typescript
import { Resource } from '@mongodb-cdk/atlas-custom-db-role'

const resource: Resource = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.Resource.property.cluster">cluster</a></code> | <code>boolean</code> | Flag that indicates whether to grant the action on the cluster resource. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.Resource.property.collection">collection</a></code> | <code>string</code> | Human-readable label that identifies the collection on which you grant the action to one MongoDB user. |
| <code><a href="#@mongodb-cdk/atlas-custom-db-role.Resource.property.db">db</a></code> | <code>string</code> | Human-readable label that identifies the database on which you grant the action to one MongoDB user. |

---

##### `cluster`<sup>Optional</sup> <a name="cluster" id="@mongodb-cdk/atlas-custom-db-role.Resource.property.cluster"></a>

```typescript
public readonly cluster: boolean;
```

- *Type:* boolean

Flag that indicates whether to grant the action on the cluster resource.

If true, MongoDB Cloud ignores the actions.resources.collection and actions.resources.db parameters.

---

##### `collection`<sup>Optional</sup> <a name="collection" id="@mongodb-cdk/atlas-custom-db-role.Resource.property.collection"></a>

```typescript
public readonly collection: string;
```

- *Type:* string

Human-readable label that identifies the collection on which you grant the action to one MongoDB user.

If you don't set this parameter, you grant the action to all collections in the database specified in the actions.resources.db parameter. If you set "actions.resources.cluster" : true, MongoDB Cloud ignores this parameter.

---

##### `db`<sup>Optional</sup> <a name="db" id="@mongodb-cdk/atlas-custom-db-role.Resource.property.db"></a>

```typescript
public readonly db: string;
```

- *Type:* string

Human-readable label that identifies the database on which you grant the action to one MongoDB user.

If you set "actions.resources.cluster" : true, MongoDB Cloud ignores this parameter.

---



