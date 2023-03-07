# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnOnlineArchive <a name="CfnOnlineArchive" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive"></a>

A CloudFormation `MongoDB::Atlas::OnlineArchive`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.Initializer"></a>

```typescript
import { CfnOnlineArchive } from '@mongodbatlas-awscdk/online-archive'

new CfnOnlineArchive(scope: Construct, id: string, props: CfnOnlineArchiveProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps">CfnOnlineArchiveProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps">CfnOnlineArchiveProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isConstruct"></a>

```typescript
import { CfnOnlineArchive } from '@mongodbatlas-awscdk/online-archive'

CfnOnlineArchive.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isCfnElement"></a>

```typescript
import { CfnOnlineArchive } from '@mongodbatlas-awscdk/online-archive'

CfnOnlineArchive.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isCfnResource"></a>

```typescript
import { CfnOnlineArchive } from '@mongodbatlas-awscdk/online-archive'

CfnOnlineArchive.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.attrArchiveId">attrArchiveId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::OnlineArchive.ArchiveId`. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.attrProjectId">attrProjectId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::OnlineArchive.ProjectId`. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.attrState">attrState</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::OnlineArchive.State`. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.attrTotalCount">attrTotalCount</a></code> | <code>number</code> | Attribute `MongoDB::Atlas::OnlineArchive.TotalCount`. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps">CfnOnlineArchiveProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrArchiveId`<sup>Required</sup> <a name="attrArchiveId" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.attrArchiveId"></a>

```typescript
public readonly attrArchiveId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::OnlineArchive.ArchiveId`.

---

##### `attrProjectId`<sup>Required</sup> <a name="attrProjectId" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.attrProjectId"></a>

```typescript
public readonly attrProjectId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::OnlineArchive.ProjectId`.

---

##### `attrState`<sup>Required</sup> <a name="attrState" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.attrState"></a>

```typescript
public readonly attrState: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::OnlineArchive.State`.

---

##### `attrTotalCount`<sup>Required</sup> <a name="attrTotalCount" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.attrTotalCount"></a>

```typescript
public readonly attrTotalCount: number;
```

- *Type:* number

Attribute `MongoDB::Atlas::OnlineArchive.TotalCount`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.props"></a>

```typescript
public readonly props: CfnOnlineArchiveProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps">CfnOnlineArchiveProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchive.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnOnlineArchiveProps <a name="CfnOnlineArchiveProps" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps"></a>

Returns, adds, edits, or removes an online archive.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.Initializer"></a>

```typescript
import { CfnOnlineArchiveProps } from '@mongodbatlas-awscdk/online-archive'

const cfnOnlineArchiveProps: CfnOnlineArchiveProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.clusterName">clusterName</a></code> | <code>string</code> | Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.criteria">criteria</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaView">CriteriaView</a></code> | Rules by which MongoDB MongoDB Cloud archives data. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.collectionType">collectionType</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchivePropsCollectionType">CfnOnlineArchivePropsCollectionType</a></code> | Classification of MongoDB database collection that you want to return. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.collName">collName</a></code> | <code>string</code> | Human-readable label that identifies the collection for which you created the online archive. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.dbName">dbName</a></code> | <code>string</code> | Human-readable label of the database that contains the collection that contains the online archive. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.includeCount">includeCount</a></code> | <code>boolean</code> | Flag that indicates whether the response returns the total number of items (**totalCount**) in the response. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.itemsPerPage">itemsPerPage</a></code> | <code>number</code> | Number of items that the response returns per page. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.pageNum">pageNum</a></code> | <code>number</code> | Number of the page that displays the current set of the total objects that the response returns. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.partitionFields">partitionFields</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldView">PartitionFieldView</a>[]</code> | List that contains document parameters to use to logically divide data within a collection. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.profile">profile</a></code> | <code>string</code> | The profile is defined in AWS Secret manager. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.schedule">schedule</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleView">ScheduleView</a></code> | Regular frequency and duration when archiving process occurs. |

---

##### `clusterName`<sup>Required</sup> <a name="clusterName" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

Human-readable label that identifies the cluster that contains the collection from which you want to remove an online archive.

---

##### `criteria`<sup>Required</sup> <a name="criteria" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.criteria"></a>

```typescript
public readonly criteria: CriteriaView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.CriteriaView">CriteriaView</a>

Rules by which MongoDB MongoDB Cloud archives data.

Use the **criteria.type** field to choose how MongoDB Cloud selects data to archive. Choose data using the age of the data or a MongoDB query.
**"criteria.type": "DATE"** selects documents to archive based on a date.
**"criteria.type": "CUSTOM"** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **"criteria.type": "CUSTOM"** when **"collectionType": "TIMESERIES"**.

---

##### `collectionType`<sup>Optional</sup> <a name="collectionType" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.collectionType"></a>

```typescript
public readonly collectionType: CfnOnlineArchivePropsCollectionType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchivePropsCollectionType">CfnOnlineArchivePropsCollectionType</a>

Classification of MongoDB database collection that you want to return.

If you set this parameter to `TIMESERIES`, set `"criteria.type" : "date"` and `"criteria.dateFormat" : "ISODATE"`.

---

##### `collName`<sup>Optional</sup> <a name="collName" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.collName"></a>

```typescript
public readonly collName: string;
```

- *Type:* string

Human-readable label that identifies the collection for which you created the online archive.

---

##### `dbName`<sup>Optional</sup> <a name="dbName" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.dbName"></a>

```typescript
public readonly dbName: string;
```

- *Type:* string

Human-readable label of the database that contains the collection that contains the online archive.

---

##### `includeCount`<sup>Optional</sup> <a name="includeCount" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.includeCount"></a>

```typescript
public readonly includeCount: boolean;
```

- *Type:* boolean

Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.

---

##### `itemsPerPage`<sup>Optional</sup> <a name="itemsPerPage" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.itemsPerPage"></a>

```typescript
public readonly itemsPerPage: number;
```

- *Type:* number

Number of items that the response returns per page.

---

##### `pageNum`<sup>Optional</sup> <a name="pageNum" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.pageNum"></a>

```typescript
public readonly pageNum: number;
```

- *Type:* number

Number of the page that displays the current set of the total objects that the response returns.

---

##### `partitionFields`<sup>Optional</sup> <a name="partitionFields" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.partitionFields"></a>

```typescript
public readonly partitionFields: PartitionFieldView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldView">PartitionFieldView</a>[]

List that contains document parameters to use to logically divide data within a collection.

Partitions provide a coarse level of filtering of the underlying collection data. To divide your data, specify up to two parameters that you frequently query. Any queries that don't use these parameters result in a full collection scan of all archived documents. This takes more time and increase your costs.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

The profile is defined in AWS Secret manager.

See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

---

##### `schedule`<sup>Optional</sup> <a name="schedule" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchiveProps.property.schedule"></a>

```typescript
public readonly schedule: ScheduleView;
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.ScheduleView">ScheduleView</a>

Regular frequency and duration when archiving process occurs.

---

### CriteriaView <a name="CriteriaView" id="@mongodbatlas-awscdk/online-archive.CriteriaView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/online-archive.CriteriaView.Initializer"></a>

```typescript
import { CriteriaView } from '@mongodbatlas-awscdk/online-archive'

const criteriaView: CriteriaView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaView.property.dateField">dateField</a></code> | <code>string</code> | Indexed database parameter that stores the date that determines when data moves to the online archive. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaView.property.dateFormat">dateFormat</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat">CriteriaViewDateFormat</a></code> | Syntax used to write the date after which data moves to the online archive. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaView.property.expireAfterDays">expireAfterDays</a></code> | <code>number</code> | Number of days after the value in the criteria.dateField when MongoDB Cloud archives data in the specified cluster. Set this parameter when you set "criteria.type" : "DATE". |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaView.property.query">query</a></code> | <code>string</code> | MongoDB find query that selects documents to archive. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaView.property.type">type</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewType">CriteriaViewType</a></code> | Means by which MongoDB Cloud selects data to archive. |

---

##### `dateField`<sup>Optional</sup> <a name="dateField" id="@mongodbatlas-awscdk/online-archive.CriteriaView.property.dateField"></a>

```typescript
public readonly dateField: string;
```

- *Type:* string

Indexed database parameter that stores the date that determines when data moves to the online archive.

MongoDB Cloud archives the data when the current date exceeds the date in this database parameter plus the number of days specified through the expireAfterDays parameter. Set this parameter when you set "criteria.type" : "DATE".

---

##### `dateFormat`<sup>Optional</sup> <a name="dateFormat" id="@mongodbatlas-awscdk/online-archive.CriteriaView.property.dateFormat"></a>

```typescript
public readonly dateFormat: CriteriaViewDateFormat;
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat">CriteriaViewDateFormat</a>

Syntax used to write the date after which data moves to the online archive.

Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when "criteria.type" : "DATE". You must set "criteria.type" : "DATE" if "collectionType": "TIMESERIES".

---

##### `expireAfterDays`<sup>Optional</sup> <a name="expireAfterDays" id="@mongodbatlas-awscdk/online-archive.CriteriaView.property.expireAfterDays"></a>

```typescript
public readonly expireAfterDays: number;
```

- *Type:* number

Number of days after the value in the criteria.dateField when MongoDB Cloud archives data in the specified cluster. Set this parameter when you set "criteria.type" : "DATE".

---

##### `query`<sup>Optional</sup> <a name="query" id="@mongodbatlas-awscdk/online-archive.CriteriaView.property.query"></a>

```typescript
public readonly query: string;
```

- *Type:* string

MongoDB find query that selects documents to archive.

The specified query follows the syntax of the db.collection.find(query) command. This query can't use the empty document ({}) to return all documents. Set this parameter when "criteria.type" : "CUSTOM".

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodbatlas-awscdk/online-archive.CriteriaView.property.type"></a>

```typescript
public readonly type: CriteriaViewType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewType">CriteriaViewType</a>

Means by which MongoDB Cloud selects data to archive.

Data can be chosen using the age of the data or a MongoDB query.
**DATE** selects documents to archive based on a date.
**CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **CUSTOM** when `"collectionType": "TIMESERIES"`.

---

### PartitionFieldView <a name="PartitionFieldView" id="@mongodbatlas-awscdk/online-archive.PartitionFieldView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/online-archive.PartitionFieldView.Initializer"></a>

```typescript
import { PartitionFieldView } from '@mongodbatlas-awscdk/online-archive'

const partitionFieldView: PartitionFieldView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldView.property.fieldName">fieldName</a></code> | <code>string</code> | Human-readable label that identifies the parameter that MongoDB Cloud uses to partition data. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldView.property.fieldType">fieldType</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType">PartitionFieldViewFieldType</a></code> | Data type of the parameter that that MongoDB Cloud uses to partition data. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldView.property.order">order</a></code> | <code>number</code> | Sequence in which MongoDB Cloud slices the collection data to create partitions. |

---

##### `fieldName`<sup>Optional</sup> <a name="fieldName" id="@mongodbatlas-awscdk/online-archive.PartitionFieldView.property.fieldName"></a>

```typescript
public readonly fieldName: string;
```

- *Type:* string

Human-readable label that identifies the parameter that MongoDB Cloud uses to partition data.

To specify a nested parameter, use the dot notation.

---

##### `fieldType`<sup>Optional</sup> <a name="fieldType" id="@mongodbatlas-awscdk/online-archive.PartitionFieldView.property.fieldType"></a>

```typescript
public readonly fieldType: PartitionFieldViewFieldType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType">PartitionFieldViewFieldType</a>

Data type of the parameter that that MongoDB Cloud uses to partition data.

Partition parameters of type [UUID](http://bsonspec.org/spec.html) must be of binary subtype 4. MongoDB Cloud skips partition parameters of type UUID with subtype 3.

---

##### `order`<sup>Optional</sup> <a name="order" id="@mongodbatlas-awscdk/online-archive.PartitionFieldView.property.order"></a>

```typescript
public readonly order: number;
```

- *Type:* number

Sequence in which MongoDB Cloud slices the collection data to create partitions.

The resource expresses this sequence starting with zero. The value of the **criteria.dateField** parameter defaults as the first item in the partition sequence.

---

### ScheduleView <a name="ScheduleView" id="@mongodbatlas-awscdk/online-archive.ScheduleView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/online-archive.ScheduleView.Initializer"></a>

```typescript
import { ScheduleView } from '@mongodbatlas-awscdk/online-archive'

const scheduleView: ScheduleView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleView.property.dayOfMonth">dayOfMonth</a></code> | <code>number</code> | Day of the month when the scheduled archive starts. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleView.property.dayOfWeek">dayOfWeek</a></code> | <code>number</code> | Day of the month when the scheduled archive starts. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleView.property.endHour">endHour</a></code> | <code>number</code> | Hour of the day when the scheduled window to run one online archive ends. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleView.property.endMinute">endMinute</a></code> | <code>number</code> | Minute of the hour when the scheduled window to run one online archive ends. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleView.property.startHour">startHour</a></code> | <code>number</code> | Hour of the day when the when the scheduled window to run one online archive starts. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleView.property.startMinute">startMinute</a></code> | <code>number</code> | Minute of the hour when the scheduled window to run one online archive starts. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleView.property.type">type</a></code> | <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleViewType">ScheduleViewType</a></code> | *No description.* |

---

##### `dayOfMonth`<sup>Optional</sup> <a name="dayOfMonth" id="@mongodbatlas-awscdk/online-archive.ScheduleView.property.dayOfMonth"></a>

```typescript
public readonly dayOfMonth: number;
```

- *Type:* number

Day of the month when the scheduled archive starts.

---

##### `dayOfWeek`<sup>Optional</sup> <a name="dayOfWeek" id="@mongodbatlas-awscdk/online-archive.ScheduleView.property.dayOfWeek"></a>

```typescript
public readonly dayOfWeek: number;
```

- *Type:* number

Day of the month when the scheduled archive starts.

---

##### `endHour`<sup>Optional</sup> <a name="endHour" id="@mongodbatlas-awscdk/online-archive.ScheduleView.property.endHour"></a>

```typescript
public readonly endHour: number;
```

- *Type:* number

Hour of the day when the scheduled window to run one online archive ends.

---

##### `endMinute`<sup>Optional</sup> <a name="endMinute" id="@mongodbatlas-awscdk/online-archive.ScheduleView.property.endMinute"></a>

```typescript
public readonly endMinute: number;
```

- *Type:* number

Minute of the hour when the scheduled window to run one online archive ends.

---

##### `startHour`<sup>Optional</sup> <a name="startHour" id="@mongodbatlas-awscdk/online-archive.ScheduleView.property.startHour"></a>

```typescript
public readonly startHour: number;
```

- *Type:* number

Hour of the day when the when the scheduled window to run one online archive starts.

---

##### `startMinute`<sup>Optional</sup> <a name="startMinute" id="@mongodbatlas-awscdk/online-archive.ScheduleView.property.startMinute"></a>

```typescript
public readonly startMinute: number;
```

- *Type:* number

Minute of the hour when the scheduled window to run one online archive starts.

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodbatlas-awscdk/online-archive.ScheduleView.property.type"></a>

```typescript
public readonly type: ScheduleViewType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/online-archive.ScheduleViewType">ScheduleViewType</a>

---



## Enums <a name="Enums" id="Enums"></a>

### CfnOnlineArchivePropsCollectionType <a name="CfnOnlineArchivePropsCollectionType" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchivePropsCollectionType"></a>

Classification of MongoDB database collection that you want to return.

If you set this parameter to `TIMESERIES`, set `"criteria.type" : "date"` and `"criteria.dateFormat" : "ISODATE"`.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchivePropsCollectionType.STANDARD">STANDARD</a></code> | STANDARD. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CfnOnlineArchivePropsCollectionType.TIMESERIES">TIMESERIES</a></code> | TIMESERIES. |

---

##### `STANDARD` <a name="STANDARD" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchivePropsCollectionType.STANDARD"></a>

STANDARD.

---


##### `TIMESERIES` <a name="TIMESERIES" id="@mongodbatlas-awscdk/online-archive.CfnOnlineArchivePropsCollectionType.TIMESERIES"></a>

TIMESERIES.

---


### CriteriaViewDateFormat <a name="CriteriaViewDateFormat" id="@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat"></a>

Syntax used to write the date after which data moves to the online archive.

Date can be expressed as ISO 8601 or Epoch timestamps. The Epoch timestamp can be expressed as nanoseconds, milliseconds, or seconds. Set this parameter when "criteria.type" : "DATE". You must set "criteria.type" : "DATE" if "collectionType": "TIMESERIES".

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat.ISODATE">ISODATE</a></code> | ISODATE. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat.EPOCH_SECONDS">EPOCH_SECONDS</a></code> | EPOCH_SECONDS. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat.EPOCH_MILLIS">EPOCH_MILLIS</a></code> | EPOCH_MILLIS. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat.EPOCH_NANOSECONDS">EPOCH_NANOSECONDS</a></code> | EPOCH_NANOSECONDS. |

---

##### `ISODATE` <a name="ISODATE" id="@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat.ISODATE"></a>

ISODATE.

---


##### `EPOCH_SECONDS` <a name="EPOCH_SECONDS" id="@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat.EPOCH_SECONDS"></a>

EPOCH_SECONDS.

---


##### `EPOCH_MILLIS` <a name="EPOCH_MILLIS" id="@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat.EPOCH_MILLIS"></a>

EPOCH_MILLIS.

---


##### `EPOCH_NANOSECONDS` <a name="EPOCH_NANOSECONDS" id="@mongodbatlas-awscdk/online-archive.CriteriaViewDateFormat.EPOCH_NANOSECONDS"></a>

EPOCH_NANOSECONDS.

---


### CriteriaViewType <a name="CriteriaViewType" id="@mongodbatlas-awscdk/online-archive.CriteriaViewType"></a>

Means by which MongoDB Cloud selects data to archive.

Data can be chosen using the age of the data or a MongoDB query.
**DATE** selects documents to archive based on a date.
**CUSTOM** selects documents to archive based on a custom JSON query. MongoDB Cloud doesn't support **CUSTOM** when `"collectionType": "TIMESERIES"`.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewType.DATE">DATE</a></code> | DATE. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.CriteriaViewType.CUSTOM">CUSTOM</a></code> | CUSTOM. |

---

##### `DATE` <a name="DATE" id="@mongodbatlas-awscdk/online-archive.CriteriaViewType.DATE"></a>

DATE.

---


##### `CUSTOM` <a name="CUSTOM" id="@mongodbatlas-awscdk/online-archive.CriteriaViewType.CUSTOM"></a>

CUSTOM.

---


### PartitionFieldViewFieldType <a name="PartitionFieldViewFieldType" id="@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType"></a>

Data type of the parameter that that MongoDB Cloud uses to partition data.

Partition parameters of type [UUID](http://bsonspec.org/spec.html) must be of binary subtype 4. MongoDB Cloud skips partition parameters of type UUID with subtype 3.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.DATE">DATE</a></code> | date. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.INT">INT</a></code> | int. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.LONG">LONG</a></code> | long. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.OBJECT_ID">OBJECT_ID</a></code> | objectId. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.STRING">STRING</a></code> | string. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.UUID">UUID</a></code> | uuid. |

---

##### `DATE` <a name="DATE" id="@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.DATE"></a>

date.

---


##### `INT` <a name="INT" id="@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.INT"></a>

int.

---


##### `LONG` <a name="LONG" id="@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.LONG"></a>

long.

---


##### `OBJECT_ID` <a name="OBJECT_ID" id="@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.OBJECT_ID"></a>

objectId.

---


##### `STRING` <a name="STRING" id="@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.STRING"></a>

string.

---


##### `UUID` <a name="UUID" id="@mongodbatlas-awscdk/online-archive.PartitionFieldViewFieldType.UUID"></a>

uuid.

---


### ScheduleViewType <a name="ScheduleViewType" id="@mongodbatlas-awscdk/online-archive.ScheduleViewType"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleViewType.DAILY">DAILY</a></code> | DAILY. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleViewType.MONTHLY">MONTHLY</a></code> | MONTHLY. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleViewType.DEFAULT">DEFAULT</a></code> | DEFAULT. |
| <code><a href="#@mongodbatlas-awscdk/online-archive.ScheduleViewType.WEEKLY">WEEKLY</a></code> | WEEKLY. |

---

##### `DAILY` <a name="DAILY" id="@mongodbatlas-awscdk/online-archive.ScheduleViewType.DAILY"></a>

DAILY.

---


##### `MONTHLY` <a name="MONTHLY" id="@mongodbatlas-awscdk/online-archive.ScheduleViewType.MONTHLY"></a>

MONTHLY.

---


##### `DEFAULT` <a name="DEFAULT" id="@mongodbatlas-awscdk/online-archive.ScheduleViewType.DEFAULT"></a>

DEFAULT.

---


##### `WEEKLY` <a name="WEEKLY" id="@mongodbatlas-awscdk/online-archive.ScheduleViewType.WEEKLY"></a>

WEEKLY.

---

