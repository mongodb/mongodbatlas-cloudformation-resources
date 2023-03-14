# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### CfnThirdPartyIntegration <a name="CfnThirdPartyIntegration" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration"></a>

A CloudFormation `MongoDB::Atlas::ThirdPartyIntegration`.

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.Initializer"></a>

```typescript
import { CfnThirdPartyIntegration } from '@mongodbatlas-awscdk/third-party-integration'

new CfnThirdPartyIntegration(scope: Construct, id: string, props: CfnThirdPartyIntegrationProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | - scope in which this resource is defined. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.Initializer.parameter.id">id</a></code> | <code>string</code> | - scoped id of the resource. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps">CfnThirdPartyIntegrationProps</a></code> | - resource properties. |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

scope in which this resource is defined.

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.Initializer.parameter.id"></a>

- *Type:* string

scoped id of the resource.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps">CfnThirdPartyIntegrationProps</a>

resource properties.

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.overrideLogicalId">overrideLogicalId</a></code> | Overrides the auto-generated logical ID with a specific ID. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addDeletionOverride">addDeletionOverride</a></code> | Syntactic sugar for `addOverride(path, undefined)`. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addDependsOn">addDependsOn</a></code> | Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addMetadata">addMetadata</a></code> | Add a value to the CloudFormation Resource Metadata. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addOverride">addOverride</a></code> | Adds an override to the synthesized CloudFormation resource. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addPropertyDeletionOverride">addPropertyDeletionOverride</a></code> | Adds an override that deletes the value of a property from the resource definition. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addPropertyOverride">addPropertyOverride</a></code> | Adds an override to a resource property. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.applyRemovalPolicy">applyRemovalPolicy</a></code> | Sets the deletion policy of the resource based on the removal policy specified. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.getAtt">getAtt</a></code> | Returns a token for an runtime attribute of this resource. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.getMetadata">getMetadata</a></code> | Retrieve a value value from the CloudFormation Resource Metadata. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `overrideLogicalId` <a name="overrideLogicalId" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.overrideLogicalId"></a>

```typescript
public overrideLogicalId(newLogicalId: string): void
```

Overrides the auto-generated logical ID with a specific ID.

###### `newLogicalId`<sup>Required</sup> <a name="newLogicalId" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.overrideLogicalId.parameter.newLogicalId"></a>

- *Type:* string

The new logical ID to use for this stack element.

---

##### `addDeletionOverride` <a name="addDeletionOverride" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addDeletionOverride"></a>

```typescript
public addDeletionOverride(path: string): void
```

Syntactic sugar for `addOverride(path, undefined)`.

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addDeletionOverride.parameter.path"></a>

- *Type:* string

The path of the value to delete.

---

##### `addDependsOn` <a name="addDependsOn" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addDependsOn"></a>

```typescript
public addDependsOn(target: CfnResource): void
```

Indicates that this resource depends on another resource and cannot be provisioned unless the other resource has been successfully provisioned.

This can be used for resources across stacks (or nested stack) boundaries
and the dependency will automatically be transferred to the relevant scope.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addDependsOn.parameter.target"></a>

- *Type:* aws-cdk-lib.CfnResource

---

##### `addMetadata` <a name="addMetadata" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addMetadata.parameter.key"></a>

- *Type:* string

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addMetadata.parameter.value"></a>

- *Type:* any

---

##### `addOverride` <a name="addOverride" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addOverride"></a>

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

###### `path`<sup>Required</sup> <a name="path" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addOverride.parameter.path"></a>

- *Type:* string

The path of the property, you can use dot notation to override values in complex types.

Any intermdediate keys
will be created as needed.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addOverride.parameter.value"></a>

- *Type:* any

The value.

Could be primitive or complex.

---

##### `addPropertyDeletionOverride` <a name="addPropertyDeletionOverride" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addPropertyDeletionOverride"></a>

```typescript
public addPropertyDeletionOverride(propertyPath: string): void
```

Adds an override that deletes the value of a property from the resource definition.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addPropertyDeletionOverride.parameter.propertyPath"></a>

- *Type:* string

The path to the property.

---

##### `addPropertyOverride` <a name="addPropertyOverride" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addPropertyOverride"></a>

```typescript
public addPropertyOverride(propertyPath: string, value: any): void
```

Adds an override to a resource property.

Syntactic sugar for `addOverride("Properties.<...>", value)`.

###### `propertyPath`<sup>Required</sup> <a name="propertyPath" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addPropertyOverride.parameter.propertyPath"></a>

- *Type:* string

The path of the property.

---

###### `value`<sup>Required</sup> <a name="value" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.addPropertyOverride.parameter.value"></a>

- *Type:* any

The value.

---

##### `applyRemovalPolicy` <a name="applyRemovalPolicy" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.applyRemovalPolicy"></a>

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

###### `policy`<sup>Optional</sup> <a name="policy" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.applyRemovalPolicy.parameter.policy"></a>

- *Type:* aws-cdk-lib.RemovalPolicy

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.applyRemovalPolicy.parameter.options"></a>

- *Type:* aws-cdk-lib.RemovalPolicyOptions

---

##### `getAtt` <a name="getAtt" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.getAtt"></a>

```typescript
public getAtt(attributeName: string): Reference
```

Returns a token for an runtime attribute of this resource.

Ideally, use generated attribute accessors (e.g. `resource.arn`), but this can be used for future compatibility
in case there is no generated attribute.

###### `attributeName`<sup>Required</sup> <a name="attributeName" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.getAtt.parameter.attributeName"></a>

- *Type:* string

The name of the attribute.

---

##### `getMetadata` <a name="getMetadata" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.getMetadata"></a>

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

###### `key`<sup>Required</sup> <a name="key" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.getMetadata.parameter.key"></a>

- *Type:* string

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isCfnElement">isCfnElement</a></code> | Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template). |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isCfnResource">isCfnResource</a></code> | Check whether the given construct is a CfnResource. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isConstruct"></a>

```typescript
import { CfnThirdPartyIntegration } from '@mongodbatlas-awscdk/third-party-integration'

CfnThirdPartyIntegration.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isCfnElement` <a name="isCfnElement" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isCfnElement"></a>

```typescript
import { CfnThirdPartyIntegration } from '@mongodbatlas-awscdk/third-party-integration'

CfnThirdPartyIntegration.isCfnElement(x: any)
```

Returns `true` if a construct is a stack element (i.e. part of the synthesized cloudformation template).

Uses duck-typing instead of `instanceof` to allow stack elements from different
versions of this library to be included in the same stack.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isCfnElement.parameter.x"></a>

- *Type:* any

---

##### `isCfnResource` <a name="isCfnResource" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isCfnResource"></a>

```typescript
import { CfnThirdPartyIntegration } from '@mongodbatlas-awscdk/third-party-integration'

CfnThirdPartyIntegration.isCfnResource(construct: IConstruct)
```

Check whether the given construct is a CfnResource.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.isCfnResource.parameter.construct"></a>

- *Type:* constructs.IConstruct

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.creationStack">creationStack</a></code> | <code>string[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.logicalId">logicalId</a></code> | <code>string</code> | The logical ID for this CloudFormation stack element. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.stack">stack</a></code> | <code>aws-cdk-lib.Stack</code> | The stack in which this element is defined. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.ref">ref</a></code> | <code>string</code> | Return a string that will be resolved to a CloudFormation `{ Ref }` for this element. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.cfnOptions">cfnOptions</a></code> | <code>aws-cdk-lib.ICfnResourceOptions</code> | Options for this resource, such as condition, update policy etc. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.cfnResourceType">cfnResourceType</a></code> | <code>string</code> | AWS resource type. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps">CfnThirdPartyIntegrationProps</a></code> | Resource props. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `creationStack`<sup>Required</sup> <a name="creationStack" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.creationStack"></a>

```typescript
public readonly creationStack: string[];
```

- *Type:* string[]

---

##### `logicalId`<sup>Required</sup> <a name="logicalId" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.logicalId"></a>

```typescript
public readonly logicalId: string;
```

- *Type:* string

The logical ID for this CloudFormation stack element.

The logical ID of the element
is calculated from the path of the resource node in the construct tree.

To override this value, use `overrideLogicalId(newLogicalId)`.

---

##### `stack`<sup>Required</sup> <a name="stack" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.stack"></a>

```typescript
public readonly stack: Stack;
```

- *Type:* aws-cdk-lib.Stack

The stack in which this element is defined.

CfnElements must be defined within a stack scope (directly or indirectly).

---

##### `ref`<sup>Required</sup> <a name="ref" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.ref"></a>

```typescript
public readonly ref: string;
```

- *Type:* string

Return a string that will be resolved to a CloudFormation `{ Ref }` for this element.

If, by any chance, the intrinsic reference of a resource is not a string, you could
coerce it to an IResolvable through `Lazy.any({ produce: resource.ref })`.

---

##### `cfnOptions`<sup>Required</sup> <a name="cfnOptions" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.cfnOptions"></a>

```typescript
public readonly cfnOptions: ICfnResourceOptions;
```

- *Type:* aws-cdk-lib.ICfnResourceOptions

Options for this resource, such as condition, update policy etc.

---

##### `cfnResourceType`<sup>Required</sup> <a name="cfnResourceType" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.cfnResourceType"></a>

```typescript
public readonly cfnResourceType: string;
```

- *Type:* string

AWS resource type.

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.props"></a>

```typescript
public readonly props: CfnThirdPartyIntegrationProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps">CfnThirdPartyIntegrationProps</a>

Resource props.

---

#### Constants <a name="Constants" id="Constants"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.CFN_RESOURCE_TYPE_NAME">CFN_RESOURCE_TYPE_NAME</a></code> | <code>string</code> | The CloudFormation resource type name for this resource class. |

---

##### `CFN_RESOURCE_TYPE_NAME`<sup>Required</sup> <a name="CFN_RESOURCE_TYPE_NAME" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration.property.CFN_RESOURCE_TYPE_NAME"></a>

```typescript
public readonly CFN_RESOURCE_TYPE_NAME: string;
```

- *Type:* string

The CloudFormation resource type name for this resource class.

---

## Structs <a name="Structs" id="Structs"></a>

### CfnThirdPartyIntegrationProps <a name="CfnThirdPartyIntegrationProps" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps"></a>

Returns, adds, edits, and removes third-party service integration configurations.

MongoDB Cloud sends alerts to each third-party service that you configure.

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.Initializer"></a>

```typescript
import { CfnThirdPartyIntegrationProps } from '@mongodbatlas-awscdk/third-party-integration'

const cfnThirdPartyIntegrationProps: CfnThirdPartyIntegrationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.apiKey">apiKey</a></code> | <code>string</code> | Key that allows MongoDB Cloud to access your Opsgenie/Datadog account. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.apiToken">apiToken</a></code> | <code>string</code> | Key that allows MongoDB Cloud to access your Slack account. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.channelName">channelName</a></code> | <code>string</code> | Name of the Slack channel to which MongoDB Cloud sends alert notifications. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.enabled">enabled</a></code> | <code>boolean</code> | Flag that indicates whether someone has activated the Prometheus integration. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.listenAddress">listenAddress</a></code> | <code>string</code> | Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.microsoftTeamsWebhookUrl">microsoftTeamsWebhookUrl</a></code> | <code>string</code> | Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.password">password</a></code> | <code>string</code> | Password required for your integration with Prometheus. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.profile">profile</a></code> | <code>string</code> | The profile is defined in AWS Secret manager. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.region">region</a></code> | <code>string</code> | Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie/Datadog API. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.routingKey">routingKey</a></code> | <code>string</code> | Routing key associated with your Splunk On-Call account. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.scheme">scheme</a></code> | <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsScheme">CfnThirdPartyIntegrationPropsScheme</a></code> | Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.secret">secret</a></code> | <code>string</code> | Parameter returned if someone configure this webhook with a secret. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.serviceDiscovery">serviceDiscovery</a></code> | <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsServiceDiscovery">CfnThirdPartyIntegrationPropsServiceDiscovery</a></code> | Desired method to discover the Prometheus service. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.serviceKey">serviceKey</a></code> | <code>string</code> | Service key associated with your PagerDuty account. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.teamName">teamName</a></code> | <code>string</code> | Human-readable label that identifies your Slack team. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.tlsPemPath">tlsPemPath</a></code> | <code>string</code> | Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.type">type</a></code> | <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType">CfnThirdPartyIntegrationPropsType</a></code> | Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.url">url</a></code> | <code>string</code> | Endpoint web address to which MongoDB Cloud sends notifications. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.userName">userName</a></code> | <code>string</code> | Human-readable label that identifies your Prometheus incoming webhook. |

---

##### `apiKey`<sup>Optional</sup> <a name="apiKey" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.apiKey"></a>

```typescript
public readonly apiKey: string;
```

- *Type:* string

Key that allows MongoDB Cloud to access your Opsgenie/Datadog account.

---

##### `apiToken`<sup>Optional</sup> <a name="apiToken" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.apiToken"></a>

```typescript
public readonly apiToken: string;
```

- *Type:* string

Key that allows MongoDB Cloud to access your Slack account.

---

##### `channelName`<sup>Optional</sup> <a name="channelName" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.channelName"></a>

```typescript
public readonly channelName: string;
```

- *Type:* string

Name of the Slack channel to which MongoDB Cloud sends alert notifications.

---

##### `enabled`<sup>Optional</sup> <a name="enabled" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.enabled"></a>

```typescript
public readonly enabled: boolean;
```

- *Type:* boolean

Flag that indicates whether someone has activated the Prometheus integration.

---

##### `listenAddress`<sup>Optional</sup> <a name="listenAddress" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.listenAddress"></a>

```typescript
public readonly listenAddress: string;
```

- *Type:* string

Combination of IPv4 address and Internet Assigned Numbers Authority (IANA) port or the IANA port alone to which Prometheus binds to ingest MongoDB metrics.

---

##### `microsoftTeamsWebhookUrl`<sup>Optional</sup> <a name="microsoftTeamsWebhookUrl" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.microsoftTeamsWebhookUrl"></a>

```typescript
public readonly microsoftTeamsWebhookUrl: string;
```

- *Type:* string

Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.

---

##### `password`<sup>Optional</sup> <a name="password" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.password"></a>

```typescript
public readonly password: string;
```

- *Type:* string

Password required for your integration with Prometheus.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

The profile is defined in AWS Secret manager.

See [Secret Manager Profile setup](../../../examples/profile-secret.yaml).

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `region`<sup>Optional</sup> <a name="region" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

Two-letter code that indicates which regional URL MongoDB uses to access the Opsgenie/Datadog API.

---

##### `routingKey`<sup>Optional</sup> <a name="routingKey" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.routingKey"></a>

```typescript
public readonly routingKey: string;
```

- *Type:* string

Routing key associated with your Splunk On-Call account.

---

##### `scheme`<sup>Optional</sup> <a name="scheme" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.scheme"></a>

```typescript
public readonly scheme: CfnThirdPartyIntegrationPropsScheme;
```

- *Type:* <a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsScheme">CfnThirdPartyIntegrationPropsScheme</a>

Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud.

---

##### `secret`<sup>Optional</sup> <a name="secret" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.secret"></a>

```typescript
public readonly secret: string;
```

- *Type:* string

Parameter returned if someone configure this webhook with a secret.

---

##### `serviceDiscovery`<sup>Optional</sup> <a name="serviceDiscovery" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.serviceDiscovery"></a>

```typescript
public readonly serviceDiscovery: CfnThirdPartyIntegrationPropsServiceDiscovery;
```

- *Type:* <a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsServiceDiscovery">CfnThirdPartyIntegrationPropsServiceDiscovery</a>

Desired method to discover the Prometheus service.

---

##### `serviceKey`<sup>Optional</sup> <a name="serviceKey" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.serviceKey"></a>

```typescript
public readonly serviceKey: string;
```

- *Type:* string

Service key associated with your PagerDuty account.

---

##### `teamName`<sup>Optional</sup> <a name="teamName" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.teamName"></a>

```typescript
public readonly teamName: string;
```

- *Type:* string

Human-readable label that identifies your Slack team.

Set this parameter when you configure a legacy Slack integration.

---

##### `tlsPemPath`<sup>Optional</sup> <a name="tlsPemPath" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.tlsPemPath"></a>

```typescript
public readonly tlsPemPath: string;
```

- *Type:* string

Root-relative path to the Transport Layer Security (TLS) Privacy Enhanced Mail (PEM) key and certificate file on the host.

---

##### `type`<sup>Optional</sup> <a name="type" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.type"></a>

```typescript
public readonly type: CfnThirdPartyIntegrationPropsType;
```

- *Type:* <a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType">CfnThirdPartyIntegrationPropsType</a>

Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud.

The value must match the third-party service integration type.

---

##### `url`<sup>Optional</sup> <a name="url" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.url"></a>

```typescript
public readonly url: string;
```

- *Type:* string

Endpoint web address to which MongoDB Cloud sends notifications.

---

##### `userName`<sup>Optional</sup> <a name="userName" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationProps.property.userName"></a>

```typescript
public readonly userName: string;
```

- *Type:* string

Human-readable label that identifies your Prometheus incoming webhook.

---



## Enums <a name="Enums" id="Enums"></a>

### CfnThirdPartyIntegrationPropsScheme <a name="CfnThirdPartyIntegrationPropsScheme" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsScheme"></a>

Security Scheme to apply to HyperText Transfer Protocol (HTTP) traffic between Prometheus and MongoDB Cloud.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsScheme.HTTP">HTTP</a></code> | http. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsScheme.HTTPS">HTTPS</a></code> | https. |

---

##### `HTTP` <a name="HTTP" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsScheme.HTTP"></a>

http.

---


##### `HTTPS` <a name="HTTPS" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsScheme.HTTPS"></a>

https.

---


### CfnThirdPartyIntegrationPropsServiceDiscovery <a name="CfnThirdPartyIntegrationPropsServiceDiscovery" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsServiceDiscovery"></a>

Desired method to discover the Prometheus service.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsServiceDiscovery.HTTP">HTTP</a></code> | http. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsServiceDiscovery.FILE">FILE</a></code> | file. |

---

##### `HTTP` <a name="HTTP" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsServiceDiscovery.HTTP"></a>

http.

---


##### `FILE` <a name="FILE" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsServiceDiscovery.FILE"></a>

file.

---


### CfnThirdPartyIntegrationPropsType <a name="CfnThirdPartyIntegrationPropsType" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType"></a>

Human-readable label that identifies the service to which you want to integrate with MongoDB Cloud.

The value must match the third-party service integration type.

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.PAGER_DUTY">PAGER_DUTY</a></code> | PAGER_DUTY. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.MICROSOFT_TEAMS">MICROSOFT_TEAMS</a></code> | MICROSOFT_TEAMS. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.SLACK">SLACK</a></code> | SLACK. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.DATADOG">DATADOG</a></code> | DATADOG. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.OPS_GENIE">OPS_GENIE</a></code> | OPS_GENIE. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.VICTOR_OPS">VICTOR_OPS</a></code> | VICTOR_OPS. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.WEBHOOK">WEBHOOK</a></code> | WEBHOOK. |
| <code><a href="#@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.PROMETHEUS">PROMETHEUS</a></code> | PROMETHEUS. |

---

##### `PAGER_DUTY` <a name="PAGER_DUTY" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.PAGER_DUTY"></a>

PAGER_DUTY.

---


##### `MICROSOFT_TEAMS` <a name="MICROSOFT_TEAMS" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.MICROSOFT_TEAMS"></a>

MICROSOFT_TEAMS.

---


##### `SLACK` <a name="SLACK" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.SLACK"></a>

SLACK.

---


##### `DATADOG` <a name="DATADOG" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.DATADOG"></a>

DATADOG.

---


##### `OPS_GENIE` <a name="OPS_GENIE" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.OPS_GENIE"></a>

OPS_GENIE.

---


##### `VICTOR_OPS` <a name="VICTOR_OPS" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.VICTOR_OPS"></a>

VICTOR_OPS.

---


##### `WEBHOOK` <a name="WEBHOOK" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.WEBHOOK"></a>

WEBHOOK.

---


##### `PROMETHEUS` <a name="PROMETHEUS" id="@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegrationPropsType.PROMETHEUS"></a>

PROMETHEUS.

---

