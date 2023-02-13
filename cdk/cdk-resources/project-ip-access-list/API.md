# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnProjectIpAccessList <a name="CfnProjectIpAccessList" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList"></a>

A CloudFormation `MongoDB::Atlas::ProjectIpAccessList`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.Initializer"></a>

```typescript
import { CfnProjectIpAccessList } from '@mongodbatlas-awscdk/project-ip-access-list'

new CfnProjectIpAccessList(scope: Construct, id: string, props: CfnProjectIpAccessListProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps">CfnProjectIpAccessListProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps">CfnProjectIpAccessListProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isConstruct"></a>

```typescript
import { CfnProjectIpAccessList } from '@mongodbatlas-awscdk/project-ip-access-list'

CfnProjectIpAccessList.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isCfnElement"></a>

```typescript
import { CfnProjectIpAccessList } from '@mongodbatlas-awscdk/project-ip-access-list'

CfnProjectIpAccessList.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isCfnResource"></a>

```typescript
import { CfnProjectIpAccessList } from '@mongodbatlas-awscdk/project-ip-access-list'

CfnProjectIpAccessList.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::ProjectIpAccessList.Id`. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps">CfnProjectIpAccessListProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::ProjectIpAccessList.Id`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.props"></a>

```typescript
public readonly props: CfnProjectIpAccessListProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps">CfnProjectIpAccessListProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessList.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### AccessListDefinition <a name="AccessListDefinition" id="@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.Initializer"></a>

```typescript
import { AccessListDefinition } from '@mongodbatlas-awscdk/project-ip-access-list'

const accessListDefinition: AccessListDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.awsSecurityGroup">awsSecurityGroup</a></code> | <code>string</code> | Unique string of the Amazon Web Services (AWS) security group that you want to add to the project's IP access list. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.cidrBlock">cidrBlock</a></code> | <code>string</code> | Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that you want to add to the project's IP access list. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.comment">comment</a></code> | <code>string</code> | Remark that explains the purpose or scope of this IP access list entry. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.deleteAfterDate">deleteAfterDate</a></code> | <code>string</code> | Date and time after which MongoDB Cloud deletes the temporary access list entry. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.ipAddress">ipAddress</a></code> | <code>string</code> | IP address that you want to add to the project's IP access list. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |

---

##### `awsSecurityGroup`<sup>Optional</sup> <a name="awsSecurityGroup" id="@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.awsSecurityGroup"></a>

```typescript
public readonly awsSecurityGroup: string;
```

- *Type:* string

Unique string of the Amazon Web Services (AWS) security group that you want to add to the project's IP access list.

Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. You must configure Virtual Private Connection (VPC) peering for your project before you can add an AWS security group to an IP access list. You cannot set AWS security groups as temporary access list entries. Don't set this parameter if you set cidrBlock or ipAddress.

---

##### `cidrBlock`<sup>Optional</sup> <a name="cidrBlock" id="@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.cidrBlock"></a>

```typescript
public readonly cidrBlock: string;
```

- *Type:* string

Range of IP addresses in Classless Inter-Domain Routing (CIDR) notation that you want to add to the project's IP access list.

Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. Don't set this parameter if you set awsSecurityGroup or ipAddress

---

##### `comment`<sup>Optional</sup> <a name="comment" id="@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.comment"></a>

```typescript
public readonly comment: string;
```

- *Type:* string

Remark that explains the purpose or scope of this IP access list entry.

---

##### `deleteAfterDate`<sup>Optional</sup> <a name="deleteAfterDate" id="@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.deleteAfterDate"></a>

```typescript
public readonly deleteAfterDate: string;
```

- *Type:* string

Date and time after which MongoDB Cloud deletes the temporary access list entry.

This parameter expresses its value in the ISO 8601 timestamp format in UTC and can include the time zone designation. The date must be later than the current date but no later than one week after you submit this request. The resource returns this parameter if you specified an expiration date when creating this IP access list entry.

---

##### `ipAddress`<sup>Optional</sup> <a name="ipAddress" id="@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.ipAddress"></a>

```typescript
public readonly ipAddress: string;
```

- *Type:* string

IP address that you want to add to the project's IP access list.

Your IP access list entry can be one awsSecurityGroup, one cidrBlock, or one ipAddress. Don't set this parameter if you set awsSecurityGroup or cidrBlock.

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/project-ip-access-list.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/project-ip-access-list.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/project-ip-access-list'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/project-ip-access-list.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/project-ip-access-list.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnProjectIpAccessListProps <a name="CfnProjectIpAccessListProps" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps"></a>

Returns, adds, edits, and removes network access limits to database deployments in MongoDB Cloud.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.Initializer"></a>

```typescript
import { CfnProjectIpAccessListProps } from '@mongodbatlas-awscdk/project-ip-access-list'

const cfnProjectIpAccessListProps: CfnProjectIpAccessListProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.accessList">accessList</a></code> | <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition">AccessListDefinition</a>[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.listOptions">listOptions</a></code> | <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.ListOptions">ListOptions</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.totalCount">totalCount</a></code> | <code>number</code> | Number of documents returned in this response. |

---

##### `accessList`<sup>Required</sup> <a name="accessList" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.accessList"></a>

```typescript
public readonly accessList: AccessListDefinition[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition">AccessListDefinition</a>[]

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/project-ip-access-list.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `listOptions`<sup>Optional</sup> <a name="listOptions" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.listOptions"></a>

```typescript
public readonly listOptions: ListOptions;
```

- *Type:* <a href="#@mongodbatlas-awscdk/project-ip-access-list.ListOptions">ListOptions</a>

---

##### `totalCount`<sup>Optional</sup> <a name="totalCount" id="@mongodbatlas-awscdk/project-ip-access-list.CfnProjectIpAccessListProps.property.totalCount"></a>

```typescript
public readonly totalCount: number;
```

- *Type:* number

Number of documents returned in this response.

---

### ListOptions <a name="ListOptions" id="@mongodbatlas-awscdk/project-ip-access-list.ListOptions"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/project-ip-access-list.ListOptions.Initializer"></a>

```typescript
import { ListOptions } from '@mongodbatlas-awscdk/project-ip-access-list'

const listOptions: ListOptions = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.ListOptions.property.includeCount">includeCount</a></code> | <code>boolean</code> | Flag that indicates whether the response returns the total number of items (totalCount) in the response. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.ListOptions.property.itemsPerPage">itemsPerPage</a></code> | <code>number</code> | Number of items that the response returns per page. |
| <code><a href="#@mongodbatlas-awscdk/project-ip-access-list.ListOptions.property.pageNum">pageNum</a></code> | <code>number</code> | Number of the page that displays the current set of the total objects that the response returns. |

---

##### `includeCount`<sup>Optional</sup> <a name="includeCount" id="@mongodbatlas-awscdk/project-ip-access-list.ListOptions.property.includeCount"></a>

```typescript
public readonly includeCount: boolean;
```

- *Type:* boolean

Flag that indicates whether the response returns the total number of items (totalCount) in the response.

---

##### `itemsPerPage`<sup>Optional</sup> <a name="itemsPerPage" id="@mongodbatlas-awscdk/project-ip-access-list.ListOptions.property.itemsPerPage"></a>

```typescript
public readonly itemsPerPage: number;
```

- *Type:* number

Number of items that the response returns per page.

---

##### `pageNum`<sup>Optional</sup> <a name="pageNum" id="@mongodbatlas-awscdk/project-ip-access-list.ListOptions.property.pageNum"></a>

```typescript
public readonly pageNum: number;
```

- *Type:* number

Number of the page that displays the current set of the total objects that the response returns.

---



