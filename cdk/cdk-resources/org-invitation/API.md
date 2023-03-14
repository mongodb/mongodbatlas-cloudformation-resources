# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnOrgInvitation <a name="CfnOrgInvitation" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation"></a>

A CloudFormation `MongoDB::Atlas::OrgInvitation`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.Initializer"></a>

```typescript
import { CfnOrgInvitation } from '@mongodbatlas-awscdk/org-invitation'

new CfnOrgInvitation(scope: Construct, id: string, props: CfnOrgInvitationProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps">CfnOrgInvitationProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps">CfnOrgInvitationProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isConstruct"></a>

```typescript
import { CfnOrgInvitation } from '@mongodbatlas-awscdk/org-invitation'

CfnOrgInvitation.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isCfnElement"></a>

```typescript
import { CfnOrgInvitation } from '@mongodbatlas-awscdk/org-invitation'

CfnOrgInvitation.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isCfnResource"></a>

```typescript
import { CfnOrgInvitation } from '@mongodbatlas-awscdk/org-invitation'

CfnOrgInvitation.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.attrCreatedAt">attrCreatedAt</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::OrgInvitation.CreatedAt`. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.attrExpiresAt">attrExpiresAt</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::OrgInvitation.ExpiresAt`. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::OrgInvitation.Id`. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.attrInviterUsername">attrInviterUsername</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::OrgInvitation.InviterUsername`. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps">CfnOrgInvitationProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrCreatedAt`<sup>Required</sup> <a name="attrCreatedAt" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.attrCreatedAt"></a>

```typescript
public readonly attrCreatedAt: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::OrgInvitation.CreatedAt`.

---

##### `attrExpiresAt`<sup>Required</sup> <a name="attrExpiresAt" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.attrExpiresAt"></a>

```typescript
public readonly attrExpiresAt: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::OrgInvitation.ExpiresAt`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::OrgInvitation.Id`.

---

##### `attrInviterUsername`<sup>Required</sup> <a name="attrInviterUsername" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.attrInviterUsername"></a>

```typescript
public readonly attrInviterUsername: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::OrgInvitation.InviterUsername`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.props"></a>

```typescript
public readonly props: CfnOrgInvitationProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps">CfnOrgInvitationProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitation.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnOrgInvitationProps <a name="CfnOrgInvitationProps" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps"></a>

Returns, adds, and edits organizational units in MongoDB Cloud.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.Initializer"></a>

```typescript
import { CfnOrgInvitationProps } from '@mongodbatlas-awscdk/org-invitation'

const cfnOrgInvitationProps: CfnOrgInvitationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.profile">profile</a></code> | <code>string</code> | Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.includeCount">includeCount</a></code> | <code>boolean</code> | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.invitationId">invitationId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the invitation. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.itemsPerPage">itemsPerPage</a></code> | <code>number</code> | Number of items that the response returns per page. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.orgId">orgId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies the organization that contains your projects. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.orgName">orgName</a></code> | <code>string</code> | Human-readable label that identifies this organization. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.pageNum">pageNum</a></code> | <code>number</code> | Number of the page that displays the current set of the total objects that the response returns. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.roles">roles</a></code> | <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles">CfnOrgInvitationPropsRoles</a>[]</code> | One or more organization or project level roles to assign to the MongoDB Cloud user. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.teamIds">teamIds</a></code> | <code>string[]</code> | List of unique 24-hexadecimal digit strings that identifies each team. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.totalCount">totalCount</a></code> | <code>number</code> | Number of documents returned in this response. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.username">username</a></code> | <code>string</code> | Email address of the MongoDB Cloud user invited to join the organization. |

---

##### `profile`<sup>Required</sup> <a name="profile" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.

---

##### `includeCount`<sup>Optional</sup> <a name="includeCount" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.includeCount"></a>

```typescript
public readonly includeCount: boolean;
```

- *Type:* boolean

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

---

##### `invitationId`<sup>Optional</sup> <a name="invitationId" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.invitationId"></a>

```typescript
public readonly invitationId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the invitation.

---

##### `itemsPerPage`<sup>Optional</sup> <a name="itemsPerPage" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.itemsPerPage"></a>

```typescript
public readonly itemsPerPage: number;
```

- *Type:* number

Number of items that the response returns per page.

---

##### `orgId`<sup>Optional</sup> <a name="orgId" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.orgId"></a>

```typescript
public readonly orgId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies the organization that contains your projects.

---

##### `orgName`<sup>Optional</sup> <a name="orgName" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.orgName"></a>

```typescript
public readonly orgName: string;
```

- *Type:* string

Human-readable label that identifies this organization.

---

##### `pageNum`<sup>Optional</sup> <a name="pageNum" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.pageNum"></a>

```typescript
public readonly pageNum: number;
```

- *Type:* number

Number of the page that displays the current set of the total objects that the response returns.

---

##### `roles`<sup>Optional</sup> <a name="roles" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.roles"></a>

```typescript
public readonly roles: CfnOrgInvitationPropsRoles[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles">CfnOrgInvitationPropsRoles</a>[]

One or more organization or project level roles to assign to the MongoDB Cloud user.

---

##### `teamIds`<sup>Optional</sup> <a name="teamIds" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.teamIds"></a>

```typescript
public readonly teamIds: string[];
```

- *Type:* string[]

List of unique 24-hexadecimal digit strings that identifies each team.

---

##### `totalCount`<sup>Optional</sup> <a name="totalCount" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.totalCount"></a>

```typescript
public readonly totalCount: number;
```

- *Type:* number

Number of documents returned in this response.

---

##### `username`<sup>Optional</sup> <a name="username" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationProps.property.username"></a>

```typescript
public readonly username: string;
```

- *Type:* string

Email address of the MongoDB Cloud user invited to join the organization.

---



## Enums <a name="Enums" id="Enums"></a>

### CfnOrgInvitationPropsRoles <a name="CfnOrgInvitationPropsRoles" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_OWNER">ORG_OWNER</a></code> | ORG_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_MEMBER">ORG_MEMBER</a></code> | ORG_MEMBER. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_GROUP_CREATOR">ORG_GROUP_CREATOR</a></code> | ORG_GROUP_CREATOR. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_BILLING_ADMIN">ORG_BILLING_ADMIN</a></code> | ORG_BILLING_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_READ_ONLY">ORG_READ_ONLY</a></code> | ORG_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_CLUSTER_MANAGER">GROUP_CLUSTER_MANAGER</a></code> | GROUP_CLUSTER_MANAGER. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_DATA_ACCESS_ADMIN">GROUP_DATA_ACCESS_ADMIN</a></code> | GROUP_DATA_ACCESS_ADMIN. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_DATA_ACCESS_READ_ONLY">GROUP_DATA_ACCESS_READ_ONLY</a></code> | GROUP_DATA_ACCESS_READ_ONLY. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_DATA_ACCESS_READ_WRITE">GROUP_DATA_ACCESS_READ_WRITE</a></code> | GROUP_DATA_ACCESS_READ_WRITE. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_OWNER">GROUP_OWNER</a></code> | GROUP_OWNER. |
| <code><a href="#@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_READ_ONLY">GROUP_READ_ONLY</a></code> | GROUP_READ_ONLY. |

---

##### `ORG_OWNER` <a name="ORG_OWNER" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_OWNER"></a>

ORG_OWNER.

---


##### `ORG_MEMBER` <a name="ORG_MEMBER" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_MEMBER"></a>

ORG_MEMBER.

---


##### `ORG_GROUP_CREATOR` <a name="ORG_GROUP_CREATOR" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_GROUP_CREATOR"></a>

ORG_GROUP_CREATOR.

---


##### `ORG_BILLING_ADMIN` <a name="ORG_BILLING_ADMIN" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_BILLING_ADMIN"></a>

ORG_BILLING_ADMIN.

---


##### `ORG_READ_ONLY` <a name="ORG_READ_ONLY" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.ORG_READ_ONLY"></a>

ORG_READ_ONLY.

---


##### `GROUP_CLUSTER_MANAGER` <a name="GROUP_CLUSTER_MANAGER" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_CLUSTER_MANAGER"></a>

GROUP_CLUSTER_MANAGER.

---


##### `GROUP_DATA_ACCESS_ADMIN` <a name="GROUP_DATA_ACCESS_ADMIN" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_DATA_ACCESS_ADMIN"></a>

GROUP_DATA_ACCESS_ADMIN.

---


##### `GROUP_DATA_ACCESS_READ_ONLY` <a name="GROUP_DATA_ACCESS_READ_ONLY" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_DATA_ACCESS_READ_ONLY"></a>

GROUP_DATA_ACCESS_READ_ONLY.

---


##### `GROUP_DATA_ACCESS_READ_WRITE` <a name="GROUP_DATA_ACCESS_READ_WRITE" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_DATA_ACCESS_READ_WRITE"></a>

GROUP_DATA_ACCESS_READ_WRITE.

---


##### `GROUP_OWNER` <a name="GROUP_OWNER" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_OWNER"></a>

GROUP_OWNER.

---


##### `GROUP_READ_ONLY` <a name="GROUP_READ_ONLY" id="@mongodbatlas-awscdk/org-invitation.CfnOrgInvitationPropsRoles.GROUP_READ_ONLY"></a>

GROUP_READ_ONLY.

---

