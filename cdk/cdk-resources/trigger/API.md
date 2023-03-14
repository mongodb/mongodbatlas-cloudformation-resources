# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnTrigger <a name="CfnTrigger" id="@mongodbatlas-awscdk/trigger.CfnTrigger"></a>

A CloudFormation `MongoDB::Atlas::Trigger`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/trigger.CfnTrigger.Initializer"></a>

```typescript
import { CfnTrigger } from '@mongodbatlas-awscdk/trigger'

new CfnTrigger(scope: Construct, id: string, props: CfnTriggerProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps">CfnTriggerProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/trigger.CfnTrigger.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/trigger.CfnTrigger.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/trigger.CfnTrigger.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps">CfnTriggerProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/trigger.CfnTrigger.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/trigger.CfnTrigger.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/trigger.CfnTrigger.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/trigger.CfnTrigger.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/trigger.CfnTrigger.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/trigger.CfnTrigger.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/trigger.CfnTrigger.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/trigger.CfnTrigger.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/trigger.CfnTrigger.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/trigger.CfnTrigger.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/trigger.CfnTrigger.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/trigger.CfnTrigger.isConstruct"></a>

```typescript
import { CfnTrigger } from '@mongodbatlas-awscdk/trigger'

CfnTrigger.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/trigger.CfnTrigger.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/trigger.CfnTrigger.isCfnElement"></a>

```typescript
import { CfnTrigger } from '@mongodbatlas-awscdk/trigger'

CfnTrigger.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/trigger.CfnTrigger.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/trigger.CfnTrigger.isCfnResource"></a>

```typescript
import { CfnTrigger } from '@mongodbatlas-awscdk/trigger'

CfnTrigger.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/trigger.CfnTrigger.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.attrId">attrId</a></code> | <code>string</code> | Attribute `MongoDB::Atlas::Trigger.Id`. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps">CfnTriggerProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `attrId`<sup>Required</sup> <a name="attrId" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.attrId"></a>

```typescript
public readonly attrId: string;
```

- *Type:* string

Attribute `MongoDB::Atlas::Trigger.Id`.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.props"></a>

```typescript
public readonly props: CfnTriggerProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps">CfnTriggerProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTrigger.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/trigger.CfnTrigger.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### AuthConfig <a name="AuthConfig" id="@mongodbatlas-awscdk/trigger.AuthConfig"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.AuthConfig.Initializer"></a>

```typescript
import { AuthConfig } from '@mongodbatlas-awscdk/trigger'

const authConfig: AuthConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfig.property.operationType">operationType</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigOperationType">AuthConfigOperationType</a></code> | The type of authentication event that the trigger listens for. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfig.property.providers">providers</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders">AuthConfigProviders</a>[]</code> | The type(s) of authentication provider that the trigger listens to. |

---

##### `operationType`<sup>Required</sup> <a name="operationType" id="@mongodbatlas-awscdk/trigger.AuthConfig.property.operationType"></a>

```typescript
public readonly operationType: AuthConfigOperationType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.AuthConfigOperationType">AuthConfigOperationType</a>

The type of authentication event that the trigger listens for.

---

##### `providers`<sup>Required</sup> <a name="providers" id="@mongodbatlas-awscdk/trigger.AuthConfig.property.providers"></a>

```typescript
public readonly providers: AuthConfigProviders[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders">AuthConfigProviders</a>[]

The type(s) of authentication provider that the trigger listens to.

---

### CfnTriggerProps <a name="CfnTriggerProps" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps"></a>

View and manage your application's [triggers](https://www.mongodb.com/docs/atlas/app-services/triggers/overview/).

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.Initializer"></a>

```typescript
import { CfnTriggerProps } from '@mongodbatlas-awscdk/trigger'

const cfnTriggerProps: CfnTriggerProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.appId">appId</a></code> | <code>string</code> | App Services Application ID. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.projectId">projectId</a></code> | <code>string</code> | Project Id for application services. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.authTrigger">authTrigger</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfig">AuthConfig</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.databaseTrigger">databaseTrigger</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig">DatabaseConfig</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.disabled">disabled</a></code> | <code>boolean</code> | If `true`, the trigger is disabled and does not listen for events or execute. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.eventProcessors">eventProcessors</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.Event">Event</a></code> | An object where each field name is an event processor ID and each value is an object that configures its corresponding event processor. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.functionId">functionId</a></code> | <code>string</code> | The ID of the function that the trigger calls when it fires. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.functionName">functionName</a></code> | <code>string</code> | The name of the function that the trigger calls when it fires, i.e. the function described by `function_id`. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.name">name</a></code> | <code>string</code> | The trigger's name. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.profile">profile</a></code> | <code>string</code> | The profile is defined in AWS Secret manager. |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.scheduleTrigger">scheduleTrigger</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.ScheduleConfig">ScheduleConfig</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.type">type</a></code> | <code>string</code> | The trigger's type. |

---

##### `appId`<sup>Required</sup> <a name="appId" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.appId"></a>

```typescript
public readonly appId: string;
```

- *Type:* string

App Services Application ID.

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Project Id for application services.

---

##### `authTrigger`<sup>Optional</sup> <a name="authTrigger" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.authTrigger"></a>

```typescript
public readonly authTrigger: AuthConfig;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.AuthConfig">AuthConfig</a>

---

##### `databaseTrigger`<sup>Optional</sup> <a name="databaseTrigger" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.databaseTrigger"></a>

```typescript
public readonly databaseTrigger: DatabaseConfig;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig">DatabaseConfig</a>

---

##### `disabled`<sup>Optional</sup> <a name="disabled" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.disabled"></a>

```typescript
public readonly disabled: boolean;
```

- *Type:* boolean

If `true`, the trigger is disabled and does not listen for events or execute.

---

##### `eventProcessors`<sup>Optional</sup> <a name="eventProcessors" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.eventProcessors"></a>

```typescript
public readonly eventProcessors: Event;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.Event">Event</a>

An object where each field name is an event processor ID and each value is an object that configures its corresponding event processor.

For an example configuration object, see
[Send Trigger Events to AWS
EventBridge](https://www.mongodb.com/docs/realm/triggers/examples/send-events-aws-eventbridge#std-label-event_processor_example).

---

##### `functionId`<sup>Optional</sup> <a name="functionId" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.functionId"></a>

```typescript
public readonly functionId: string;
```

- *Type:* string

The ID of the function that the trigger calls when it fires.

This value is the same as `event_processors.FUNCTION.function_id`.
You can either define the value here or in `event_processors.FUNCTION.function_id`.
The App Services backend duplicates the value to the configuration location where you did not define it.

For example, if you define `function_id`, the backend duplicates it to `event_processors.FUNCTION.function_id`.

---

##### `functionName`<sup>Optional</sup> <a name="functionName" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.functionName"></a>

```typescript
public readonly functionName: string;
```

- *Type:* string

The name of the function that the trigger calls when it fires, i.e. the function described by `function_id`.

This value is the same as `event_processors.FUNCTION.function_name`.
You can either define the value here or in `event_processors.FUNCTION.function_name`.
The App Services backend duplicates the value to the configuration location where you did not define it.

For example, if you define `function_name`, the backend duplicates it to `event_processors.FUNCTION.function_name`.

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

The trigger's name.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

The profile is defined in AWS Secret manager.

See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

---

##### `scheduleTrigger`<sup>Optional</sup> <a name="scheduleTrigger" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.scheduleTrigger"></a>

```typescript
public readonly scheduleTrigger: ScheduleConfig;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.ScheduleConfig">ScheduleConfig</a>

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodbatlas-awscdk/trigger.CfnTriggerProps.property.type"></a>

```typescript
public readonly type: string;
```

- *Type:* string

The trigger's type.

---

### DatabaseConfig <a name="DatabaseConfig" id="@mongodbatlas-awscdk/trigger.DatabaseConfig"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.Initializer"></a>

```typescript
import { DatabaseConfig } from '@mongodbatlas-awscdk/trigger'

const databaseConfig: DatabaseConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.collection">collection</a></code> | <code>string</code> | The name of a collection in the specified database. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.database">database</a></code> | <code>string</code> | The name of a database in the linked data source. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.fullDocument">fullDocument</a></code> | <code>boolean</code> | If `true`, indicates that `UPDATE` change events should include the most current [majority-committed](https://www.mongodb.com/docs/manual/reference/read-concern-majority/) version of the modified document in the `fullDocument` field. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.fullDocumentBeforeChange">fullDocumentBeforeChange</a></code> | <code>boolean</code> | If true, indicates that `UPDATE` change events should include a snapshot of the modified document from immediately before the update was applied. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.match">match</a></code> | <code>any</code> | A [$match](https://www.mongodb.com/docs/manual/reference/operator/aggregation/match) expression filters change events. The trigger will only fire if the expression evaluates to true for a given change event. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.operationTypes">operationTypes</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes">DatabaseConfigOperationTypes</a>[]</code> | The type(s) of MongoDB change event that the trigger listens for. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.serviceId">serviceId</a></code> | <code>string</code> | The _id value of a linked MongoDB data source. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.skipCatchupEvents">skipCatchupEvents</a></code> | <code>boolean</code> | If `true`, enabling the Trigger after it was disabled will not invoke events that occurred while the Trigger was disabled. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.tolerateResumeErrors">tolerateResumeErrors</a></code> | <code>boolean</code> | If `true`, when this Trigger's resume token cannot be found in the cluster's oplog, the Trigger automatically resumes processing events at the next relevant change stream event. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfig.property.unordered">unordered</a></code> | <code>boolean</code> | If `true`, event ordering is disabled and this Trigger can process events in parallel. |

---

##### `collection`<sup>Optional</sup> <a name="collection" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.collection"></a>

```typescript
public readonly collection: string;
```

- *Type:* string

The name of a collection in the specified database.

The
trigger listens to events from this collection.

---

##### `database`<sup>Optional</sup> <a name="database" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.database"></a>

```typescript
public readonly database: string;
```

- *Type:* string

The name of a database in the linked data source.

---

##### `fullDocument`<sup>Optional</sup> <a name="fullDocument" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.fullDocument"></a>

```typescript
public readonly fullDocument: boolean;
```

- *Type:* boolean

If `true`, indicates that `UPDATE` change events should include the most current [majority-committed](https://www.mongodb.com/docs/manual/reference/read-concern-majority/) version of the modified document in the `fullDocument` field.

---

##### `fullDocumentBeforeChange`<sup>Optional</sup> <a name="fullDocumentBeforeChange" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.fullDocumentBeforeChange"></a>

```typescript
public readonly fullDocumentBeforeChange: boolean;
```

- *Type:* boolean

If true, indicates that `UPDATE` change events should include a snapshot of the modified document from immediately before the update was applied.

You must enable [document
preimages](https://www.mongodb.com/docs/atlas/app-services/mongodb/preimages/)
for your cluster to include these snapshots.

---

##### `match`<sup>Optional</sup> <a name="match" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.match"></a>

```typescript
public readonly match: any;
```

- *Type:* any

A [$match](https://www.mongodb.com/docs/manual/reference/operator/aggregation/match) expression filters change events. The trigger will only fire if the expression evaluates to true for a given change event.

---

##### `operationTypes`<sup>Optional</sup> <a name="operationTypes" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.operationTypes"></a>

```typescript
public readonly operationTypes: DatabaseConfigOperationTypes[];
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes">DatabaseConfigOperationTypes</a>[]

The type(s) of MongoDB change event that the trigger listens for.

---

##### `serviceId`<sup>Optional</sup> <a name="serviceId" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.serviceId"></a>

```typescript
public readonly serviceId: string;
```

- *Type:* string

The _id value of a linked MongoDB data source.

See [Get a Data Source](#operation/adminGetService).

---

##### `skipCatchupEvents`<sup>Optional</sup> <a name="skipCatchupEvents" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.skipCatchupEvents"></a>

```typescript
public readonly skipCatchupEvents: boolean;
```

- *Type:* boolean

If `true`, enabling the Trigger after it was disabled will not invoke events that occurred while the Trigger was disabled.

---

##### `tolerateResumeErrors`<sup>Optional</sup> <a name="tolerateResumeErrors" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.tolerateResumeErrors"></a>

```typescript
public readonly tolerateResumeErrors: boolean;
```

- *Type:* boolean

If `true`, when this Trigger's resume token cannot be found in the cluster's oplog, the Trigger automatically resumes processing events at the next relevant change stream event.

All change stream events from when the Trigger was suspended until the Trigger
resumes execution do not have the Trigger fire for them.

---

##### `unordered`<sup>Optional</sup> <a name="unordered" id="@mongodbatlas-awscdk/trigger.DatabaseConfig.property.unordered"></a>

```typescript
public readonly unordered: boolean;
```

- *Type:* boolean

If `true`, event ordering is disabled and this Trigger can process events in parallel.

If `false`, event
ordering is enabled and the Trigger executes events
serially.

---

### Event <a name="Event" id="@mongodbatlas-awscdk/trigger.Event"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.Event.Initializer"></a>

```typescript
import { Event } from '@mongodbatlas-awscdk/trigger'

const event: Event = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.Event.property.awseventbridge">awseventbridge</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.EventAwseventbridge">EventAwseventbridge</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/trigger.Event.property.function">function</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.EventFunction">EventFunction</a></code> | *No description.* |

---

##### `awseventbridge`<sup>Optional</sup> <a name="awseventbridge" id="@mongodbatlas-awscdk/trigger.Event.property.awseventbridge"></a>

```typescript
public readonly awseventbridge: EventAwseventbridge;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.EventAwseventbridge">EventAwseventbridge</a>

---

##### `function`<sup>Optional</sup> <a name="function" id="@mongodbatlas-awscdk/trigger.Event.property.function"></a>

```typescript
public readonly function: EventFunction;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.EventFunction">EventFunction</a>

---

### EventAwseventbridge <a name="EventAwseventbridge" id="@mongodbatlas-awscdk/trigger.EventAwseventbridge"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.EventAwseventbridge.Initializer"></a>

```typescript
import { EventAwseventbridge } from '@mongodbatlas-awscdk/trigger'

const eventAwseventbridge: EventAwseventbridge = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.EventAwseventbridge.property.awsConfig">awsConfig</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig">EventAwseventbridgeAwsConfig</a></code> | *No description.* |

---

##### `awsConfig`<sup>Optional</sup> <a name="awsConfig" id="@mongodbatlas-awscdk/trigger.EventAwseventbridge.property.awsConfig"></a>

```typescript
public readonly awsConfig: EventAwseventbridgeAwsConfig;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig">EventAwseventbridgeAwsConfig</a>

---

### EventAwseventbridgeAwsConfig <a name="EventAwseventbridgeAwsConfig" id="@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig.Initializer"></a>

```typescript
import { EventAwseventbridgeAwsConfig } from '@mongodbatlas-awscdk/trigger'

const eventAwseventbridgeAwsConfig: EventAwseventbridgeAwsConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig.property.accountId">accountId</a></code> | <code>string</code> | An AWS Account ID. |
| <code><a href="#@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig.property.extendedJsonEnabled">extendedJsonEnabled</a></code> | <code>boolean</code> | If `true`, event objects are serialized using EJSON. |
| <code><a href="#@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig.property.region">region</a></code> | <code>string</code> | An AWS region. |

---

##### `accountId`<sup>Optional</sup> <a name="accountId" id="@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig.property.accountId"></a>

```typescript
public readonly accountId: string;
```

- *Type:* string

An AWS Account ID.

---

##### `extendedJsonEnabled`<sup>Optional</sup> <a name="extendedJsonEnabled" id="@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig.property.extendedJsonEnabled"></a>

```typescript
public readonly extendedJsonEnabled: boolean;
```

- *Type:* boolean

If `true`, event objects are serialized using EJSON.

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/trigger.EventAwseventbridgeAwsConfig.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

An AWS region.

---

### EventFunction <a name="EventFunction" id="@mongodbatlas-awscdk/trigger.EventFunction"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.EventFunction.Initializer"></a>

```typescript
import { EventFunction } from '@mongodbatlas-awscdk/trigger'

const eventFunction: EventFunction = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.EventFunction.property.funcConfig">funcConfig</a></code> | <code><a href="#@mongodbatlas-awscdk/trigger.EventFunctionFuncConfig">EventFunctionFuncConfig</a></code> | *No description.* |

---

##### `funcConfig`<sup>Optional</sup> <a name="funcConfig" id="@mongodbatlas-awscdk/trigger.EventFunction.property.funcConfig"></a>

```typescript
public readonly funcConfig: EventFunctionFuncConfig;
```

- *Type:* <a href="#@mongodbatlas-awscdk/trigger.EventFunctionFuncConfig">EventFunctionFuncConfig</a>

---

### EventFunctionFuncConfig <a name="EventFunctionFuncConfig" id="@mongodbatlas-awscdk/trigger.EventFunctionFuncConfig"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.EventFunctionFuncConfig.Initializer"></a>

```typescript
import { EventFunctionFuncConfig } from '@mongodbatlas-awscdk/trigger'

const eventFunctionFuncConfig: EventFunctionFuncConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.EventFunctionFuncConfig.property.functionId">functionId</a></code> | <code>string</code> | The ID of the function that the trigger calls when it fires. |
| <code><a href="#@mongodbatlas-awscdk/trigger.EventFunctionFuncConfig.property.functionName">functionName</a></code> | <code>string</code> | The name of the function that the trigger calls when it fires, i.e. the function described by `function_id`. |

---

##### `functionId`<sup>Optional</sup> <a name="functionId" id="@mongodbatlas-awscdk/trigger.EventFunctionFuncConfig.property.functionId"></a>

```typescript
public readonly functionId: string;
```

- *Type:* string

The ID of the function that the trigger calls when it fires.

This value is the same as the root-level `function_id`.
You can either define the value here or in `function_id`.
The App Services backend duplicates the value to the configuration location where you did not define it.

For example, if you define `event_processors.FUNCTION.function_id`, the backend duplicates it to `function_id`.

---

##### `functionName`<sup>Optional</sup> <a name="functionName" id="@mongodbatlas-awscdk/trigger.EventFunctionFuncConfig.property.functionName"></a>

```typescript
public readonly functionName: string;
```

- *Type:* string

The name of the function that the trigger calls when it fires, i.e. the function described by `function_id`.

This value is the same as the root-level `function_name`.
You can either define the value here or in `function_name`.
The App Services backend duplicates the value to the configuration location where you did not define it.

For example, if you define `event_processors.FUNCTION.function_name`, the backend duplicates it to `function_name`.

---

### ScheduleConfig <a name="ScheduleConfig" id="@mongodbatlas-awscdk/trigger.ScheduleConfig"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/trigger.ScheduleConfig.Initializer"></a>

```typescript
import { ScheduleConfig } from '@mongodbatlas-awscdk/trigger'

const scheduleConfig: ScheduleConfig = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.ScheduleConfig.property.schedule">schedule</a></code> | <code>string</code> | A [cron expression](https://www.mongodb.com/docs/atlas/app-services/triggers/scheduled-triggers/#cron-expressions) that specifies when the trigger executes. |
| <code><a href="#@mongodbatlas-awscdk/trigger.ScheduleConfig.property.skipcatchupEvents">skipcatchupEvents</a></code> | <code>boolean</code> | If `true`, enabling the trigger after it was disabled will not invoke events that occurred while the trigger was disabled. |

---

##### `schedule`<sup>Optional</sup> <a name="schedule" id="@mongodbatlas-awscdk/trigger.ScheduleConfig.property.schedule"></a>

```typescript
public readonly schedule: string;
```

- *Type:* string

A [cron expression](https://www.mongodb.com/docs/atlas/app-services/triggers/scheduled-triggers/#cron-expressions) that specifies when the trigger executes.

---

##### `skipcatchupEvents`<sup>Optional</sup> <a name="skipcatchupEvents" id="@mongodbatlas-awscdk/trigger.ScheduleConfig.property.skipcatchupEvents"></a>

```typescript
public readonly skipcatchupEvents: boolean;
```

- *Type:* boolean

If `true`, enabling the trigger after it was disabled will not invoke events that occurred while the trigger was disabled.

---



## Enums <a name="Enums" id="Enums"></a>

### AuthConfigOperationType <a name="AuthConfigOperationType" id="@mongodbatlas-awscdk/trigger.AuthConfigOperationType"></a>

The type of authentication event that the trigger listens for.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigOperationType.LOGIN">LOGIN</a></code> | LOGIN. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigOperationType.CREATE">CREATE</a></code> | CREATE. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigOperationType.DELETE">DELETE</a></code> | DELETE. |

---

##### `LOGIN` <a name="LOGIN" id="@mongodbatlas-awscdk/trigger.AuthConfigOperationType.LOGIN"></a>

LOGIN.

---


##### `CREATE` <a name="CREATE" id="@mongodbatlas-awscdk/trigger.AuthConfigOperationType.CREATE"></a>

CREATE.

---


##### `DELETE` <a name="DELETE" id="@mongodbatlas-awscdk/trigger.AuthConfigOperationType.DELETE"></a>

DELETE.

---


### AuthConfigProviders <a name="AuthConfigProviders" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders.ANON_USER">ANON_USER</a></code> | anon-user. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders.API_KEY">API_KEY</a></code> | api-key. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders.CUSTOM_TOKEN">CUSTOM_TOKEN</a></code> | custom-token. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders.CUSTOM_FUNCTION">CUSTOM_FUNCTION</a></code> | custom-function. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders.LOCAL_USERPASS">LOCAL_USERPASS</a></code> | local-userpass. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders.OAUTH2_APPLE">OAUTH2_APPLE</a></code> | oauth2-apple. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders.OAUTH2_FACEBOOK">OAUTH2_FACEBOOK</a></code> | oauth2-facebook. |
| <code><a href="#@mongodbatlas-awscdk/trigger.AuthConfigProviders.OAUTH2_GOOGLE">OAUTH2_GOOGLE</a></code> | oauth2-google. |

---

##### `ANON_USER` <a name="ANON_USER" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders.ANON_USER"></a>

anon-user.

---


##### `API_KEY` <a name="API_KEY" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders.API_KEY"></a>

api-key.

---


##### `CUSTOM_TOKEN` <a name="CUSTOM_TOKEN" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders.CUSTOM_TOKEN"></a>

custom-token.

---


##### `CUSTOM_FUNCTION` <a name="CUSTOM_FUNCTION" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders.CUSTOM_FUNCTION"></a>

custom-function.

---


##### `LOCAL_USERPASS` <a name="LOCAL_USERPASS" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders.LOCAL_USERPASS"></a>

local-userpass.

---


##### `OAUTH2_APPLE` <a name="OAUTH2_APPLE" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders.OAUTH2_APPLE"></a>

oauth2-apple.

---


##### `OAUTH2_FACEBOOK` <a name="OAUTH2_FACEBOOK" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders.OAUTH2_FACEBOOK"></a>

oauth2-facebook.

---


##### `OAUTH2_GOOGLE` <a name="OAUTH2_GOOGLE" id="@mongodbatlas-awscdk/trigger.AuthConfigProviders.OAUTH2_GOOGLE"></a>

oauth2-google.

---


### DatabaseConfigOperationTypes <a name="DatabaseConfigOperationTypes" id="@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes.INSERT">INSERT</a></code> | INSERT. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes.UPDATE">UPDATE</a></code> | UPDATE. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes.REPLACE">REPLACE</a></code> | REPLACE. |
| <code><a href="#@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes.DELETE">DELETE</a></code> | DELETE. |

---

##### `INSERT` <a name="INSERT" id="@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes.INSERT"></a>

INSERT.

---


##### `UPDATE` <a name="UPDATE" id="@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes.UPDATE"></a>

UPDATE.

---


##### `REPLACE` <a name="REPLACE" id="@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes.REPLACE"></a>

REPLACE.

---


##### `DELETE` <a name="DELETE" id="@mongodbatlas-awscdk/trigger.DatabaseConfigOperationTypes.DELETE"></a>

DELETE.

---

