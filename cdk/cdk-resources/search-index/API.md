# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnSearchIndex <a name="CfnSearchIndex" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex"></a>

A CloudFormation `MongoDB::Atlas::SearchIndex`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.Initializer"></a>

```typescript
import { CfnSearchIndex } from '@mongodbatlas-awscdk/search-index'

new CfnSearchIndex(scope: Construct, id: string, props: CfnSearchIndexProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps">CfnSearchIndexProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps">CfnSearchIndexProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.isConstruct"></a>

```typescript
import { CfnSearchIndex } from '@mongodbatlas-awscdk/search-index'

CfnSearchIndex.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.isCfnElement"></a>

```typescript
import { CfnSearchIndex } from '@mongodbatlas-awscdk/search-index'

CfnSearchIndex.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.isCfnResource"></a>

```typescript
import { CfnSearchIndex } from '@mongodbatlas-awscdk/search-index'

CfnSearchIndex.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.attrIndexId">attrIndexId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::SearchIndex.IndexId`. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.attrStatus">attrStatus</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::SearchIndex.Status`. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps">CfnSearchIndexProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrIndexId`<sup>Required</sup> <a name="attrIndexId" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.attrIndexId"></a>

```typescript
public readonly attrIndexId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::SearchIndex.IndexId`.

---

##### `attrStatus`<sup>Required</sup> <a name="attrStatus" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.attrStatus"></a>

```typescript
public readonly attrStatus: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::SearchIndex.Status`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.props"></a>

```typescript
public readonly props: CfnSearchIndexProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps">CfnSearchIndexProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/search-index.CfnSearchIndex.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiAtlasFtsAnalyzersViewManual <a name="ApiAtlasFtsAnalyzersViewManual" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.Initializer"></a>

```typescript
import { ApiAtlasFtsAnalyzersViewManual } from '@mongodbatlas-awscdk/search-index'

const apiAtlasFtsAnalyzersViewManual: ApiAtlasFtsAnalyzersViewManual = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.property.charFilters">charFilters</a></code> | <code>any[]</code> | Filters that examine text one character at a time and perform filtering operations. |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.property.name">name</a></code> | <code>string</code> | Human-readable name that identifies the custom analyzer. |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.property.tokenFilters">tokenFilters</a></code> | <code>any[]</code> | Filter that performs operations such as:. |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.property.tokenizer">tokenizer</a></code> | <code>any</code> | Tokenizer that you want to use to create tokens. |

---

##### `charFilters`<sup>Optional</sup> <a name="charFilters" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.property.charFilters"></a>

```typescript
public readonly charFilters: any[];
```

- *Type:* any[]

Filters that examine text one character at a time and perform filtering operations.

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable name that identifies the custom analyzer.

Names must be unique within an index, and must not start with any of the following strings:
- `lucene.`
- `builtin.`
- `mongodb.`

---

##### `tokenFilters`<sup>Optional</sup> <a name="tokenFilters" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.property.tokenFilters"></a>

```typescript
public readonly tokenFilters: any[];
```

- *Type:* any[]

Filter that performs operations such as:.

Stemming, which reduces related words, such as "talking", "talked", and "talks" to their root word "talk".

- Redaction, the removal of sensitive information from public documents.

---

##### `tokenizer`<sup>Optional</sup> <a name="tokenizer" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual.property.tokenizer"></a>

```typescript
public readonly tokenizer: any;
```

- *Type:* any

Tokenizer that you want to use to create tokens.

Tokens determine how Atlas Search splits up text into discrete chunks for indexing.

---

### ApiAtlasFtsMappingsViewManual <a name="ApiAtlasFtsMappingsViewManual" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsMappingsViewManual"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsMappingsViewManual.Initializer"></a>

```typescript
import { ApiAtlasFtsMappingsViewManual } from '@mongodbatlas-awscdk/search-index'

const apiAtlasFtsMappingsViewManual: ApiAtlasFtsMappingsViewManual = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsMappingsViewManual.property.dynamic">dynamic</a></code> | <code>boolean</code> | Flag that indicates whether the index uses dynamic or static mappings. |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsMappingsViewManual.property.fields">fields</a></code> | <code>string</code> | One or more field specifications for the Atlas Search index. |

---

##### `dynamic`<sup>Optional</sup> <a name="dynamic" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsMappingsViewManual.property.dynamic"></a>

```typescript
public readonly dynamic: boolean;
```

- *Type:* boolean

Flag that indicates whether the index uses dynamic or static mappings.

Required if **mappings.fields** is omitted.

---

##### `fields`<sup>Optional</sup> <a name="fields" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsMappingsViewManual.property.fields"></a>

```typescript
public readonly fields: string;
```

- *Type:* string

One or more field specifications for the Atlas Search index.

Required if **mappings.dynamic** is omitted or set to **false**.

---

### ApiAtlasFtsSynonymMappingDefinitionView <a name="ApiAtlasFtsSynonymMappingDefinitionView" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView.Initializer"></a>

```typescript
import { ApiAtlasFtsSynonymMappingDefinitionView } from '@mongodbatlas-awscdk/search-index'

const apiAtlasFtsSynonymMappingDefinitionView: ApiAtlasFtsSynonymMappingDefinitionView = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView.property.analyzer">analyzer</a></code> | <code>string</code> | Specific pre-defined method chosen to apply to the synonyms to be searched. |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView.property.name">name</a></code> | <code>string</code> | Human-readable label that identifies the synonym definition. |
| <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView.property.source">source</a></code> | <code><a href="#@mongodbatlas-awscdk/search-index.SynonymSource">SynonymSource</a></code> | Data set that stores the mapping one or more words map to one or more synonyms of those words. |

---

##### `analyzer`<sup>Optional</sup> <a name="analyzer" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView.property.analyzer"></a>

```typescript
public readonly analyzer: string;
```

- *Type:* string

Specific pre-defined method chosen to apply to the synonyms to be searched.

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable label that identifies the synonym definition.

Each **synonym.name** must be unique within the same index definition.

---

##### `source`<sup>Optional</sup> <a name="source" id="@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView.property.source"></a>

```typescript
public readonly source: SynonymSource;
```

- *Type:* <a href="#@mongodbatlas-awscdk/search-index.SynonymSource">SynonymSource</a>

Data set that stores the mapping one or more words map to one or more synonyms of those words.

---

### CfnSearchIndexProps <a name="CfnSearchIndexProps" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps"></a>

Returns, adds, edits, and removes Atlas Search indexes.

Also returns and updates user-defined analyzers.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.Initializer"></a>

```typescript
import { CfnSearchIndexProps } from '@mongodbatlas-awscdk/search-index'

const cfnSearchIndexProps: CfnSearchIndexProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.clusterName">clusterName</a></code> | <code>string</code> | Name of the cluster that contains the database and collection with one or more Application Search indexes. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.collectionName">collectionName</a></code> | <code>string</code> | Human-readable label that identifies the collection that contains one or more Atlas Search indexes. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.database">database</a></code> | <code>string</code> | Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.mappings">mappings</a></code> | <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsMappingsViewManual">ApiAtlasFtsMappingsViewManual</a></code> | Index specifications for the collection's fields. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.analyzer">analyzer</a></code> | <code>string</code> | Specific pre-defined method chosen to convert database field text into searchable words. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.analyzers">analyzers</a></code> | <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual">ApiAtlasFtsAnalyzersViewManual</a>[]</code> | List of user-defined methods to convert database field text into searchable words. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.name">name</a></code> | <code>string</code> | Human-readable label that identifies this index. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.profile">profile</a></code> | <code>string</code> | The profile is defined in AWS Secret manager. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.searchAnalyzer">searchAnalyzer</a></code> | <code>string</code> | Method applied to identify words when searching this index. |
| <code><a href="#@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.synonyms">synonyms</a></code> | <code><a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView">ApiAtlasFtsSynonymMappingDefinitionView</a>[]</code> | Rule sets that map words to their synonyms in this index. |

---

##### `clusterName`<sup>Required</sup> <a name="clusterName" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.clusterName"></a>

```typescript
public readonly clusterName: string;
```

- *Type:* string

Name of the cluster that contains the database and collection with one or more Application Search indexes.

---

##### `collectionName`<sup>Required</sup> <a name="collectionName" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.collectionName"></a>

```typescript
public readonly collectionName: string;
```

- *Type:* string

Human-readable label that identifies the collection that contains one or more Atlas Search indexes.

---

##### `database`<sup>Required</sup> <a name="database" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.database"></a>

```typescript
public readonly database: string;
```

- *Type:* string

Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.

---

##### `mappings`<sup>Required</sup> <a name="mappings" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.mappings"></a>

```typescript
public readonly mappings: ApiAtlasFtsMappingsViewManual;
```

- *Type:* <a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsMappingsViewManual">ApiAtlasFtsMappingsViewManual</a>

Index specifications for the collection's fields.

---

##### `analyzer`<sup>Optional</sup> <a name="analyzer" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.analyzer"></a>

```typescript
public readonly analyzer: string;
```

- *Type:* string

Specific pre-defined method chosen to convert database field text into searchable words.

This conversion reduces the text of fields into the smallest units of text. These units are called a **term** or **token**. This process, known as tokenization, involves a variety of changes made to the text in fields:

- extracting words
- removing punctuation
- removing accents
- changing to lowercase
- removing common words
- reducing words to their root form (stemming)
- changing words to their base form (lemmatization)
MongoDB Cloud uses the selected process to build the Atlas Search index.

---

##### `analyzers`<sup>Optional</sup> <a name="analyzers" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.analyzers"></a>

```typescript
public readonly analyzers: ApiAtlasFtsAnalyzersViewManual[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsAnalyzersViewManual">ApiAtlasFtsAnalyzersViewManual</a>[]

List of user-defined methods to convert database field text into searchable words.

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

Human-readable label that identifies this index.

Within each namespace, names of all indexes in the namespace must be unique.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

The profile is defined in AWS Secret manager.

See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `searchAnalyzer`<sup>Optional</sup> <a name="searchAnalyzer" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.searchAnalyzer"></a>

```typescript
public readonly searchAnalyzer: string;
```

- *Type:* string

Method applied to identify words when searching this index.

---

##### `synonyms`<sup>Optional</sup> <a name="synonyms" id="@mongodbatlas-awscdk/search-index.CfnSearchIndexProps.property.synonyms"></a>

```typescript
public readonly synonyms: ApiAtlasFtsSynonymMappingDefinitionView[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/search-index.ApiAtlasFtsSynonymMappingDefinitionView">ApiAtlasFtsSynonymMappingDefinitionView</a>[]

Rule sets that map words to their synonyms in this index.

---

### SynonymSource <a name="SynonymSource" id="@mongodbatlas-awscdk/search-index.SynonymSource"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/search-index.SynonymSource.Initializer"></a>

```typescript
import { SynonymSource } from '@mongodbatlas-awscdk/search-index'

const synonymSource: SynonymSource = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/search-index.SynonymSource.property.collection">collection</a></code> | <code>string</code> | Human-readable label that identifies the MongoDB collection that stores words and their applicable synonyms. |

---

##### `collection`<sup>Optional</sup> <a name="collection" id="@mongodbatlas-awscdk/search-index.SynonymSource.property.collection"></a>

```typescript
public readonly collection: string;
```

- *Type:* string

Human-readable label that identifies the MongoDB collection that stores words and their applicable synonyms.

---



