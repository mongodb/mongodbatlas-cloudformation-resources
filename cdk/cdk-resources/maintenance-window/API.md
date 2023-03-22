# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnMaintenanceWindow <a name="CfnMaintenanceWindow" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow"></a>

A CloudFormation `MongoDB::Atlas::MaintenanceWindow`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.Initializer"></a>

```typescript
import { CfnMaintenanceWindow } from '@mongodbatlas-awscdk/maintenance-window'

new CfnMaintenanceWindow(scope: Construct, id: string, props: CfnMaintenanceWindowProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps">CfnMaintenanceWindowProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps">CfnMaintenanceWindowProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isConstruct"></a>

```typescript
import { CfnMaintenanceWindow } from '@mongodbatlas-awscdk/maintenance-window'

CfnMaintenanceWindow.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isCfnElement"></a>

```typescript
import { CfnMaintenanceWindow } from '@mongodbatlas-awscdk/maintenance-window'

CfnMaintenanceWindow.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isCfnResource"></a>

```typescript
import { CfnMaintenanceWindow } from '@mongodbatlas-awscdk/maintenance-window'

CfnMaintenanceWindow.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps">CfnMaintenanceWindowProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.props"></a>

```typescript
public readonly props: CfnMaintenanceWindowProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps">CfnMaintenanceWindowProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindow.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnMaintenanceWindowProps <a name="CfnMaintenanceWindowProps" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps"></a>

The maintenanceWindow resource provides access to retrieve or update the current Atlas project maintenance window.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.Initializer"></a>

```typescript
import { CfnMaintenanceWindowProps } from '@mongodbatlas-awscdk/maintenance-window'

const cfnMaintenanceWindowProps: CfnMaintenanceWindowProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.hourOfDay">hourOfDay</a></code> | <code>number</code> | Zero-based integer that represents the hour of the of the day that the maintenance window starts according to a 24-hour clock. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.autoDeferOnceEnabled">autoDeferOnceEnabled</a></code> | <code>boolean</code> | Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.dayOfWeek">dayOfWeek</a></code> | <code>number</code> | One-based integer that represents the day of the week that the maintenance window starts. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.profile">profile</a></code> | <code>string</code> | The profile is defined in AWS Secret manager. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.startAsap">startAsap</a></code> | <code>boolean</code> | Flag that indicates whether MongoDB Cloud starts the maintenance window immediately upon receiving this request. |

---

##### `hourOfDay`<sup>Required</sup> <a name="hourOfDay" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.hourOfDay"></a>

```typescript
public readonly hourOfDay: number;
```

- *Type:* number

Zero-based integer that represents the hour of the of the day that the maintenance window starts according to a 24-hour clock.

Use `0` for midnight and `12` for noon.

---

##### `autoDeferOnceEnabled`<sup>Optional</sup> <a name="autoDeferOnceEnabled" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.autoDeferOnceEnabled"></a>

```typescript
public readonly autoDeferOnceEnabled: boolean;
```

- *Type:* boolean

Flag that indicates whether MongoDB Cloud should defer all maintenance windows for one week after you enable them.

---

##### `dayOfWeek`<sup>Optional</sup> <a name="dayOfWeek" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.dayOfWeek"></a>

```typescript
public readonly dayOfWeek: number;
```

- *Type:* number

One-based integer that represents the day of the week that the maintenance window starts.

| Value | Day of Week |
|---|---|
| `1` | Sunday |
| `2` | Monday |
| `3` | Tuesday |
| `4` | Wednesday |
| `5` | Thursday |
| `6` | Friday |
| `7` | Saturday |

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

The profile is defined in AWS Secret manager.

See [Secret Manager Profile setup](../../../examples/profile-secret.yaml)

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `startAsap`<sup>Optional</sup> <a name="startAsap" id="@mongodbatlas-awscdk/maintenance-window.CfnMaintenanceWindowProps.property.startAsap"></a>

```typescript
public readonly startAsap: boolean;
```

- *Type:* boolean

Flag that indicates whether MongoDB Cloud starts the maintenance window immediately upon receiving this request.

To start the maintenance window immediately for your project, MongoDB Cloud must have maintenance scheduled and you must set a maintenance window. This flag resets to `false` after MongoDB Cloud completes maintenance.

---



