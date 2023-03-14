# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnProject <a name="CfnProject" id="@mongodbatlas-awscdk/project.CfnProject"></a>

A CloudFormation `MongoDB::Atlas::Project`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/project.CfnProject.Initializer"></a>

```typescript
import { CfnProject } from '@mongodbatlas-awscdk/project'

new CfnProject(scope: Construct, id: string, props: CfnProjectProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps">CfnProjectProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/project.CfnProject.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/project.CfnProject.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/project.CfnProject.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/project.CfnProjectProps">CfnProjectProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/project.CfnProject.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/project.CfnProject.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/project.CfnProject.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/project.CfnProject.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/project.CfnProject.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/project.CfnProject.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/project.CfnProject.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/project.CfnProject.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/project.CfnProject.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/project.CfnProject.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/project.CfnProject.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/project.CfnProject.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/project.CfnProject.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/project.CfnProject.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/project.CfnProject.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/project.CfnProject.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/project.CfnProject.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/project.CfnProject.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/project.CfnProject.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/project.CfnProject.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/project.CfnProject.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/project.CfnProject.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/project.CfnProject.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/project.CfnProject.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/project.CfnProject.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/project.CfnProject.isConstruct"></a>

```typescript
import { CfnProject } from '@mongodbatlas-awscdk/project'

CfnProject.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/project.CfnProject.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/project.CfnProject.isCfnElement"></a>

```typescript
import { CfnProject } from '@mongodbatlas-awscdk/project'

CfnProject.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/project.CfnProject.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/project.CfnProject.isCfnResource"></a>

```typescript
import { CfnProject } from '@mongodbatlas-awscdk/project'

CfnProject.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/project.CfnProject.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.attrClusterCount">attrClusterCount</a></code> | <code>number</code> | Attribute `MongoDB::Atlas::Project.ClusterCount`. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.attrCreated">attrCreated</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Project.Created`. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Project.Id`. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.attrProjectOwnerId">attrProjectOwnerId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Project.ProjectOwnerId`. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps">CfnProjectProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/project.CfnProject.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/project.CfnProject.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/project.CfnProject.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/project.CfnProject.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/project.CfnProject.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/project.CfnProject.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/project.CfnProject.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrClusterCount`<sup>Required</sup> <a name="attrClusterCount" id="@mongodbatlas-awscdk/project.CfnProject.property.attrClusterCount"></a>

```typescript
public readonly attrClusterCount: number;
```

- *Type:* number

Attribute `MongoDB::Atlas::Project.ClusterCount`.

---

##### `attrCreated`<sup>Required</sup> <a name="attrCreated" id="@mongodbatlas-awscdk/project.CfnProject.property.attrCreated"></a>

```typescript
public readonly attrCreated: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Project.Created`.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/project.CfnProject.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Project.Id`.

---

##### `attrProjectOwnerId`<sup>Required</sup> <a name="attrProjectOwnerId" id="@mongodbatlas-awscdk/project.CfnProject.property.attrProjectOwnerId"></a>

```typescript
public readonly attrProjectOwnerId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Project.ProjectOwnerId`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/project.CfnProject.property.props"></a>

```typescript
public readonly props: CfnProjectProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/project.CfnProjectProps">CfnProjectProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProject.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/project.CfnProject.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnProjectProps <a name="CfnProjectProps" id="@mongodbatlas-awscdk/project.CfnProjectProps"></a>

Retrieves or creates projects in any given Atlas organization.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/project.CfnProjectProps.Initializer"></a>

```typescript
import { CfnProjectProps } from '@mongodbatlas-awscdk/project'

const cfnProjectProps: CfnProjectProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps.property.name">name</a></code> | <code>string</code> | Name of the project to create. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps.property.orgId">orgId</a></code> | <code>string</code> | Unique identifier of the organization within which to create the project. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps.property.profile">profile</a></code> | <code>string</code> | Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used. |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps.property.projectApiKeys">projectApiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/project.ProjectApiKey">ProjectApiKey</a>[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps.property.projectSettings">projectSettings</a></code> | <code><a href="#@mongodbatlas-awscdk/project.ProjectSettings">ProjectSettings</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps.property.projectTeams">projectTeams</a></code> | <code><a href="#@mongodbatlas-awscdk/project.ProjectTeam">ProjectTeam</a>[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.CfnProjectProps.property.withDefaultAlertsSettings">withDefaultAlertsSettings</a></code> | <code>boolean</code> | Flag that indicates whether to create the project with default alert settings. |

---

##### `name`<sup>Required</sup> <a name="name" id="@mongodbatlas-awscdk/project.CfnProjectProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Name of the project to create.

---

##### `orgId`<sup>Required</sup> <a name="orgId" id="@mongodbatlas-awscdk/project.CfnProjectProps.property.orgId"></a>

```typescript
public readonly orgId: string;
```

- *Type:* string

Unique identifier of the organization within which to create the project.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/project.CfnProjectProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

Profile used to provide credentials information, (a secret with the cfn/atlas/profile/{Profile}, is required), if not provided default is used.

---

##### `projectApiKeys`<sup>Optional</sup> <a name="projectApiKeys" id="@mongodbatlas-awscdk/project.CfnProjectProps.property.projectApiKeys"></a>

```typescript
public readonly projectApiKeys: ProjectApiKey[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/project.ProjectApiKey">ProjectApiKey</a>[]

---

##### `projectSettings`<sup>Optional</sup> <a name="projectSettings" id="@mongodbatlas-awscdk/project.CfnProjectProps.property.projectSettings"></a>

```typescript
public readonly projectSettings: ProjectSettings;
```

- *Type:* <a href="#@mongodbatlas-awscdk/project.ProjectSettings">ProjectSettings</a>

---

##### `projectTeams`<sup>Optional</sup> <a name="projectTeams" id="@mongodbatlas-awscdk/project.CfnProjectProps.property.projectTeams"></a>

```typescript
public readonly projectTeams: ProjectTeam[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/project.ProjectTeam">ProjectTeam</a>[]

---

##### `withDefaultAlertsSettings`<sup>Optional</sup> <a name="withDefaultAlertsSettings" id="@mongodbatlas-awscdk/project.CfnProjectProps.property.withDefaultAlertsSettings"></a>

```typescript
public readonly withDefaultAlertsSettings: boolean;
```

- *Type:* boolean

Flag that indicates whether to create the project with default alert settings.

---

### ProjectApiKey <a name="ProjectApiKey" id="@mongodbatlas-awscdk/project.ProjectApiKey"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/project.ProjectApiKey.Initializer"></a>

```typescript
import { ProjectApiKey } from '@mongodbatlas-awscdk/project'

const projectApiKey: ProjectApiKey = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectApiKey.property.key">key</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectApiKey.property.roleNames">roleNames</a></code> | <code>string[]</code> | *No description.* |

---

##### `key`<sup>Optional</sup> <a name="key" id="@mongodbatlas-awscdk/project.ProjectApiKey.property.key"></a>

```typescript
public readonly key: string;
```

- *Type:* string

---

##### `roleNames`<sup>Optional</sup> <a name="roleNames" id="@mongodbatlas-awscdk/project.ProjectApiKey.property.roleNames"></a>

```typescript
public readonly roleNames: string[];
```

- *Type:* string[]

---

### ProjectSettings <a name="ProjectSettings" id="@mongodbatlas-awscdk/project.ProjectSettings"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/project.ProjectSettings.Initializer"></a>

```typescript
import { ProjectSettings } from '@mongodbatlas-awscdk/project'

const projectSettings: ProjectSettings = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectSettings.property.isCollectDatabaseSpecificsStatisticsEnabled">isCollectDatabaseSpecificsStatisticsEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectSettings.property.isDataExplorerEnabled">isDataExplorerEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectSettings.property.isPerformanceAdvisorEnabled">isPerformanceAdvisorEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectSettings.property.isRealtimePerformancePanelEnabled">isRealtimePerformancePanelEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectSettings.property.isSchemaAdvisorEnabled">isSchemaAdvisorEnabled</a></code> | <code>boolean</code> | *No description.* |

---

##### `isCollectDatabaseSpecificsStatisticsEnabled`<sup>Optional</sup> <a name="isCollectDatabaseSpecificsStatisticsEnabled" id="@mongodbatlas-awscdk/project.ProjectSettings.property.isCollectDatabaseSpecificsStatisticsEnabled"></a>

```typescript
public readonly isCollectDatabaseSpecificsStatisticsEnabled: boolean;
```

- *Type:* boolean

---

##### `isDataExplorerEnabled`<sup>Optional</sup> <a name="isDataExplorerEnabled" id="@mongodbatlas-awscdk/project.ProjectSettings.property.isDataExplorerEnabled"></a>

```typescript
public readonly isDataExplorerEnabled: boolean;
```

- *Type:* boolean

---

##### `isPerformanceAdvisorEnabled`<sup>Optional</sup> <a name="isPerformanceAdvisorEnabled" id="@mongodbatlas-awscdk/project.ProjectSettings.property.isPerformanceAdvisorEnabled"></a>

```typescript
public readonly isPerformanceAdvisorEnabled: boolean;
```

- *Type:* boolean

---

##### `isRealtimePerformancePanelEnabled`<sup>Optional</sup> <a name="isRealtimePerformancePanelEnabled" id="@mongodbatlas-awscdk/project.ProjectSettings.property.isRealtimePerformancePanelEnabled"></a>

```typescript
public readonly isRealtimePerformancePanelEnabled: boolean;
```

- *Type:* boolean

---

##### `isSchemaAdvisorEnabled`<sup>Optional</sup> <a name="isSchemaAdvisorEnabled" id="@mongodbatlas-awscdk/project.ProjectSettings.property.isSchemaAdvisorEnabled"></a>

```typescript
public readonly isSchemaAdvisorEnabled: boolean;
```

- *Type:* boolean

---

### ProjectTeam <a name="ProjectTeam" id="@mongodbatlas-awscdk/project.ProjectTeam"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/project.ProjectTeam.Initializer"></a>

```typescript
import { ProjectTeam } from '@mongodbatlas-awscdk/project'

const projectTeam: ProjectTeam = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectTeam.property.roleNames">roleNames</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/project.ProjectTeam.property.teamId">teamId</a></code> | <code>string</code> | *No description.* |

---

##### `roleNames`<sup>Optional</sup> <a name="roleNames" id="@mongodbatlas-awscdk/project.ProjectTeam.property.roleNames"></a>

```typescript
public readonly roleNames: string[];
```

- *Type:* string[]

---

##### `teamId`<sup>Optional</sup> <a name="teamId" id="@mongodbatlas-awscdk/project.ProjectTeam.property.teamId"></a>

```typescript
public readonly teamId: string;
```

- *Type:* string

---



