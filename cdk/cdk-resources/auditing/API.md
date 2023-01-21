# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnAuditing <a name="CfnAuditing" id="@mongodb-cdk/atlas-auditing.CfnAuditing"></a>

A CloudFormation `MongoDB::Atlas::Auditing`.

#### Initializers <a name="Initializers" id="@mongodb-cdk/atlas-auditing.CfnAuditing.Initializer"></a>

```typescript
import { CfnAuditing } from '@mongodb-cdk/atlas-auditing'

new CfnAuditing(scope: Construct, id: string, props: CfnAuditingProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps">CfnAuditingProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodb-cdk/atlas-auditing.CfnAuditing.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodb-cdk/atlas-auditing.CfnAuditing.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-auditing.CfnAuditing.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps">CfnAuditingProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodb-cdk/atlas-auditing.CfnAuditing.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodb-cdk/atlas-auditing.CfnAuditing.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodb-cdk/atlas-auditing.CfnAuditing.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodb-cdk/atlas-auditing.CfnAuditing.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodb-cdk/atlas-auditing.CfnAuditing.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodb-cdk/atlas-auditing.CfnAuditing.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodb-cdk/atlas-auditing.CfnAuditing.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodb-cdk/atlas-auditing.CfnAuditing.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodb-cdk/atlas-auditing.CfnAuditing.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodb-cdk/atlas-auditing.CfnAuditing.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodb-cdk/atlas-auditing.CfnAuditing.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodb-cdk/atlas-auditing.CfnAuditing.isConstruct"></a>

```typescript
import { CfnAuditing } from '@mongodb-cdk/atlas-auditing'

CfnAuditing.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-auditing.CfnAuditing.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodb-cdk/atlas-auditing.CfnAuditing.isCfnElement"></a>

```typescript
import { CfnAuditing } from '@mongodb-cdk/atlas-auditing'

CfnAuditing.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodb-cdk/atlas-auditing.CfnAuditing.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodb-cdk/atlas-auditing.CfnAuditing.isCfnResource"></a>

```typescript
import { CfnAuditing } from '@mongodb-cdk/atlas-auditing'

CfnAuditing.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodb-cdk/atlas-auditing.CfnAuditing.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.props">props</a></code> | <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps">CfnAuditingProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.props"></a>

```typescript
public readonly props: CfnAuditingProps;
```

- *Type:* <a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps">CfnAuditingProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditing.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodb-cdk/atlas-auditing.CfnAuditing.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodb-cdk/atlas-auditing.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-auditing.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodb-cdk/atlas-auditing'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-auditing.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-auditing.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodb-cdk/atlas-auditing.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodb-cdk/atlas-auditing.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### CfnAuditingProps <a name="CfnAuditingProps" id="@mongodb-cdk/atlas-auditing.CfnAuditingProps"></a>

Returns and edits database auditing settings for MongoDB Cloud projects.

#### Initializer <a name="Initializer" id="@mongodb-cdk/atlas-auditing.CfnAuditingProps.Initializer"></a>

```typescript
import { CfnAuditingProps } from '@mongodb-cdk/atlas-auditing'

const cfnAuditingProps: CfnAuditingProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodb-cdk/atlas-auditing.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.groupId">groupId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.auditAuthorizationSuccess">auditAuthorizationSuccess</a></code> | <code>boolean</code> | Flag that indicates whether someone set auditing to track successful authentications. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.auditFilter">auditFilter</a></code> | <code>string</code> | JSON document that specifies which events to record. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.configurationType">configurationType</a></code> | <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType">CfnAuditingPropsConfigurationType</a></code> | Human-readable label that displays how to configure the audit filter. |

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodb-cdk/atlas-auditing.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `groupId`<sup>Required</sup> <a name="groupId" id="@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.groupId"></a>

```typescript
public readonly groupId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `auditAuthorizationSuccess`<sup>Optional</sup> <a name="auditAuthorizationSuccess" id="@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.auditAuthorizationSuccess"></a>

```typescript
public readonly auditAuthorizationSuccess: boolean;
```

- *Type:* boolean

Flag that indicates whether someone set auditing to track successful authentications.

This only applies to the `"atype" : "authCheck"` audit filter. Setting this parameter to `true` degrades cluster performance.

---

##### `auditFilter`<sup>Optional</sup> <a name="auditFilter" id="@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.auditFilter"></a>

```typescript
public readonly auditFilter: string;
```

- *Type:* string

JSON document that specifies which events to record.

Escape any characters that may prevent parsing, such as single or double quotes, using a backslash (`\`), for more information about audit filters refer to https://www.mongodb.com/docs/manual/tutorial/configure-audit-filters/.

---

##### `configurationType`<sup>Optional</sup> <a name="configurationType" id="@mongodb-cdk/atlas-auditing.CfnAuditingProps.property.configurationType"></a>

```typescript
public readonly configurationType: CfnAuditingPropsConfigurationType;
```

- *Type:* <a href="#@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType">CfnAuditingPropsConfigurationType</a>

Human-readable label that displays how to configure the audit filter.

---



## Enums <a name="Enums" id="Enums"></a>

### CfnAuditingPropsConfigurationType <a name="CfnAuditingPropsConfigurationType" id="@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType"></a>

Human-readable label that displays how to configure the audit filter.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType.FILTER_BUILDER">FILTER_BUILDER</a></code> | FILTER_BUILDER. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType.FILTER_JSON">FILTER_JSON</a></code> | FILTER_JSON. |
| <code><a href="#@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType.NONE">NONE</a></code> | NONE. |

---

##### `FILTER_BUILDER` <a name="FILTER_BUILDER" id="@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType.FILTER_BUILDER"></a>

FILTER_BUILDER.

---


##### `FILTER_JSON` <a name="FILTER_JSON" id="@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType.FILTER_JSON"></a>

FILTER_JSON.

---


##### `NONE` <a name="NONE" id="@mongodb-cdk/atlas-auditing.CfnAuditingPropsConfigurationType.NONE"></a>

NONE.

---

