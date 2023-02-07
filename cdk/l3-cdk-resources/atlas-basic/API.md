# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### AtlasBasic <a name="AtlasBasic" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic"></a>

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer"></a>

```typescript
import { AtlasBasic } from '@mongodbatlas-awscdk/atlas-basic'

new AtlasBasic(scope: Construct, id: string, props: AtlasBasicProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps">AtlasBasicProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps">AtlasBasicProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.toString">toString</a></code> | Returns a string representation of this construct. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.addDependency">addDependency</a></code> | Add a dependency between this stack and another stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.addTransform">addTransform</a></code> | Add a Transform to this stack. A Transform is a macro that AWS CloudFormation uses to process your template. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.exportValue">exportValue</a></code> | Create a CloudFormation Export for a value. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.formatArn">formatArn</a></code> | Creates an ARN from components. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.getLogicalId">getLogicalId</a></code> | Allocates a stack-unique CloudFormation-compatible logical identity for a specific resource. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.renameLogicalId">renameLogicalId</a></code> | Rename a generated logical identities. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.reportMissingContextKey">reportMissingContextKey</a></code> | Indicate that a context key was expected. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.resolve">resolve</a></code> | Resolve a tokenized value in the context of the current stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.splitArn">splitArn</a></code> | Splits the provided ARN into its components. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.toJsonString">toJsonString</a></code> | Convert an object, potentially containing tokens, to a JSON string. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

##### `addDependency` <a name="addDependency" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.addDependency"></a>

```typescript
public addDependency(target: Stack, reason?: string): void
```

Add a dependency between this stack and another stack.

This can be used to define dependencies between any two stacks within an
app, and also supports nested stacks.

###### `target`<sup>Required</sup> <a name="target" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.addDependency.parameter.target"></a>

- *Type:* aws-cdk-lib.Stack

---

###### `reason`<sup>Optional</sup> <a name="reason" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.addDependency.parameter.reason"></a>

- *Type:* string

---

##### `addTransform` <a name="addTransform" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.addTransform"></a>

```typescript
public addTransform(transform: string): void
```

Add a Transform to this stack. A Transform is a macro that AWS CloudFormation uses to process your template.

Duplicate values are removed when stack is synthesized.

> [https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/transform-section-structure.html)

*Example*

```typescript
declare const stack: Stack;

stack.addTransform('AWS::Serverless-2016-10-31')
```


###### `transform`<sup>Required</sup> <a name="transform" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.addTransform.parameter.transform"></a>

- *Type:* string

The transform to add.

---

##### `exportValue` <a name="exportValue" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.exportValue"></a>

```typescript
public exportValue(exportedValue: any, options?: ExportValueOptions): string
```

Create a CloudFormation Export for a value.

Returns a string representing the corresponding `Fn.importValue()`
expression for this Export. You can control the name for the export by
passing the `name` option.

If you don't supply a value for `name`, the value you're exporting must be
a Resource attribute (for example: `bucket.bucketName`) and it will be
given the same name as the automatic cross-stack reference that would be created
if you used the attribute in another Stack.

One of the uses for this method is to *remove* the relationship between
two Stacks established by automatic cross-stack references. It will
temporarily ensure that the CloudFormation Export still exists while you
remove the reference from the consuming stack. After that, you can remove
the resource and the manual export.

## Example

Here is how the process works. Let's say there are two stacks,
`producerStack` and `consumerStack`, and `producerStack` has a bucket
called `bucket`, which is referenced by `consumerStack` (perhaps because
an AWS Lambda Function writes into it, or something like that).

It is not safe to remove `producerStack.bucket` because as the bucket is being
deleted, `consumerStack` might still be using it.

Instead, the process takes two deployments:

### Deployment 1: break the relationship

- Make sure `consumerStack` no longer references `bucket.bucketName` (maybe the consumer
   stack now uses its own bucket, or it writes to an AWS DynamoDB table, or maybe you just
   remove the Lambda Function altogether).
- In the `ProducerStack` class, call `this.exportValue(this.bucket.bucketName)`. This
   will make sure the CloudFormation Export continues to exist while the relationship
   between the two stacks is being broken.
- Deploy (this will effectively only change the `consumerStack`, but it's safe to deploy both).

### Deployment 2: remove the bucket resource

- You are now free to remove the `bucket` resource from `producerStack`.
- Don't forget to remove the `exportValue()` call as well.
- Deploy again (this time only the `producerStack` will be changed -- the bucket will be deleted).

###### `exportedValue`<sup>Required</sup> <a name="exportedValue" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.exportValue.parameter.exportedValue"></a>

- *Type:* any

---

###### `options`<sup>Optional</sup> <a name="options" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.exportValue.parameter.options"></a>

- *Type:* aws-cdk-lib.ExportValueOptions

---

##### `formatArn` <a name="formatArn" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.formatArn"></a>

```typescript
public formatArn(components: ArnComponents): string
```

Creates an ARN from components.

If `partition`, `region` or `account` are not specified, the stack's
partition, region and account will be used.

If any component is the empty string, an empty string will be inserted
into the generated ARN at the location that component corresponds to.

The ARN will be formatted as follows:

   arn:{partition}:{service}:{region}:{account}:{resource}{sep}}{resource-name}

The required ARN pieces that are omitted will be taken from the stack that
the 'scope' is attached to. If all ARN pieces are supplied, the supplied scope
can be 'undefined'.

###### `components`<sup>Required</sup> <a name="components" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.formatArn.parameter.components"></a>

- *Type:* aws-cdk-lib.ArnComponents

---

##### `getLogicalId` <a name="getLogicalId" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.getLogicalId"></a>

```typescript
public getLogicalId(element: CfnElement): string
```

Allocates a stack-unique CloudFormation-compatible logical identity for a specific resource.

This method is called when a `CfnElement` is created and used to render the
initial logical identity of resources. Logical ID renames are applied at
this stage.

This method uses the protected method `allocateLogicalId` to render the
logical ID for an element. To modify the naming scheme, extend the `Stack`
class and override this method.

###### `element`<sup>Required</sup> <a name="element" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.getLogicalId.parameter.element"></a>

- *Type:* aws-cdk-lib.CfnElement

The CloudFormation element for which a logical identity is needed.

---

##### `renameLogicalId` <a name="renameLogicalId" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.renameLogicalId"></a>

```typescript
public renameLogicalId(oldId: string, newId: string): void
```

Rename a generated logical identities.

To modify the naming scheme strategy, extend the `Stack` class and
override the `allocateLogicalId` method.

###### `oldId`<sup>Required</sup> <a name="oldId" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.renameLogicalId.parameter.oldId"></a>

- *Type:* string

---

###### `newId`<sup>Required</sup> <a name="newId" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.renameLogicalId.parameter.newId"></a>

- *Type:* string

---

##### `reportMissingContextKey` <a name="reportMissingContextKey" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.reportMissingContextKey"></a>

```typescript
public reportMissingContextKey(report: MissingContext): void
```

Indicate that a context key was expected.

Contains instructions which will be emitted into the cloud assembly on how
the key should be supplied.

###### `report`<sup>Required</sup> <a name="report" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.reportMissingContextKey.parameter.report"></a>

- *Type:* aws-cdk-lib.cloud_assembly_schema.MissingContext

The set of parameters needed to obtain the context.

---

##### `resolve` <a name="resolve" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.resolve"></a>

```typescript
public resolve(obj: any): any
```

Resolve a tokenized value in the context of the current stack.

###### `obj`<sup>Required</sup> <a name="obj" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.resolve.parameter.obj"></a>

- *Type:* any

---

##### `splitArn` <a name="splitArn" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.splitArn"></a>

```typescript
public splitArn(arn: string, arnFormat: ArnFormat): ArnComponents
```

Splits the provided ARN into its components.

Works both if 'arn' is a string like 'arn:aws:s3:::bucket',
and a Token representing a dynamic CloudFormation expression
(in which case the returned components will also be dynamic CloudFormation expressions,
encoded as Tokens).

###### `arn`<sup>Required</sup> <a name="arn" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.splitArn.parameter.arn"></a>

- *Type:* string

the ARN to split into its components.

---

###### `arnFormat`<sup>Required</sup> <a name="arnFormat" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.splitArn.parameter.arnFormat"></a>

- *Type:* aws-cdk-lib.ArnFormat

the expected format of 'arn' - depends on what format the service 'arn' represents uses.

---

##### `toJsonString` <a name="toJsonString" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.toJsonString"></a>

```typescript
public toJsonString(obj: any, space?: number): string
```

Convert an object, potentially containing tokens, to a JSON string.

###### `obj`<sup>Required</sup> <a name="obj" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.toJsonString.parameter.obj"></a>

- *Type:* any

---

###### `space`<sup>Optional</sup> <a name="space" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.toJsonString.parameter.space"></a>

- *Type:* number

---

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isStack">isStack</a></code> | Return whether the given object is a Stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.of">of</a></code> | Looks up the first stack scope in which `construct` is defined. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isConstruct"></a>

```typescript
import { AtlasBasic } from '@mongodbatlas-awscdk/atlas-basic'

AtlasBasic.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

##### `isStack` <a name="isStack" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isStack"></a>

```typescript
import { AtlasBasic } from '@mongodbatlas-awscdk/atlas-basic'

AtlasBasic.isStack(x: any)
```

Return whether the given object is a Stack.

We do attribute detection since we can't reliably use 'instanceof'.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.isStack.parameter.x"></a>

- *Type:* any

---

##### `of` <a name="of" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.of"></a>

```typescript
import { AtlasBasic } from '@mongodbatlas-awscdk/atlas-basic'

AtlasBasic.of(construct: IConstruct)
```

Looks up the first stack scope in which `construct` is defined.

Fails if there is no stack up the tree.

###### `construct`<sup>Required</sup> <a name="construct" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.of.parameter.construct"></a>

- *Type:* constructs.IConstruct

The construct to start the search from.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.account">account</a></code> | <code>string</code> | The AWS account into which this stack will be deployed. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.artifactId">artifactId</a></code> | <code>string</code> | The ID of the cloud assembly artifact for this stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.availabilityZones">availabilityZones</a></code> | <code>string[]</code> | Returns the list of AZs that are available in the AWS environment (account/region) associated with this stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.dependencies">dependencies</a></code> | <code>aws-cdk-lib.Stack[]</code> | Return the stacks this stack depends on. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.environment">environment</a></code> | <code>string</code> | The environment coordinates in which this stack is deployed. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.nested">nested</a></code> | <code>boolean</code> | Indicates if this is a nested stack, in which case `parentStack` will include a reference to it's parent. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.nestedStackParent">nestedStackParent</a></code> | <code>aws-cdk-lib.Stack</code> | If this is a nested stack, returns it's parent stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.nestedStackResource">nestedStackResource</a></code> | <code>aws-cdk-lib.CfnResource</code> | If this is a nested stack, this represents its `AWS::CloudFormation::Stack` resource. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.notificationArns">notificationArns</a></code> | <code>string[]</code> | Returns the list of notification Amazon Resource Names (ARNs) for the current stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.partition">partition</a></code> | <code>string</code> | The partition in which this stack is defined. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.region">region</a></code> | <code>string</code> | The AWS region into which this stack will be deployed (e.g. `us-west-2`). |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.stackId">stackId</a></code> | <code>string</code> | The ID of the stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.stackName">stackName</a></code> | <code>string</code> | The concrete CloudFormation physical stack name. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.synthesizer">synthesizer</a></code> | <code>aws-cdk-lib.IStackSynthesizer</code> | Synthesis method for this stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.tags">tags</a></code> | <code>aws-cdk-lib.TagManager</code> | Tags to be applied to the stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.templateFile">templateFile</a></code> | <code>string</code> | The name of the CloudFormation template file emitted to the output directory during synthesis. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.templateOptions">templateOptions</a></code> | <code>aws-cdk-lib.ITemplateOptions</code> | Options for CloudFormation template (like version, transform, description). |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.terminationProtection">terminationProtection</a></code> | <code>boolean</code> | Whether termination protection is enabled for this stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.urlSuffix">urlSuffix</a></code> | <code>string</code> | The Amazon domain suffix for the region in which this stack is defined. |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `account`<sup>Required</sup> <a name="account" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.account"></a>

```typescript
public readonly account: string;
```

- *Type:* string

The AWS account into which this stack will be deployed.

This value is resolved according to the following rules:

1. The value provided to `env.account` when the stack is defined. This can
    either be a concerete account (e.g. `585695031111`) or the
    `Aws.accountId` token.
3. `Aws.accountId`, which represents the CloudFormation intrinsic reference
    `{ "Ref": "AWS::AccountId" }` encoded as a string token.

Preferably, you should use the return value as an opaque string and not
attempt to parse it to implement your logic. If you do, you must first
check that it is a concerete value an not an unresolved token. If this
value is an unresolved token (`Token.isUnresolved(stack.account)` returns
`true`), this implies that the user wishes that this stack will synthesize
into a **account-agnostic template**. In this case, your code should either
fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
implement some other region-agnostic behavior.

---

##### `artifactId`<sup>Required</sup> <a name="artifactId" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.artifactId"></a>

```typescript
public readonly artifactId: string;
```

- *Type:* string

The ID of the cloud assembly artifact for this stack.

---

##### `availabilityZones`<sup>Required</sup> <a name="availabilityZones" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.availabilityZones"></a>

```typescript
public readonly availabilityZones: string[];
```

- *Type:* string[]

Returns the list of AZs that are available in the AWS environment (account/region) associated with this stack.

If the stack is environment-agnostic (either account and/or region are
tokens), this property will return an array with 2 tokens that will resolve
at deploy-time to the first two availability zones returned from CloudFormation's
`Fn::GetAZs` intrinsic function.

If they are not available in the context, returns a set of dummy values and
reports them as missing, and let the CLI resolve them by calling EC2
`DescribeAvailabilityZones` on the target environment.

To specify a different strategy for selecting availability zones override this method.

---

##### `dependencies`<sup>Required</sup> <a name="dependencies" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.dependencies"></a>

```typescript
public readonly dependencies: Stack[];
```

- *Type:* aws-cdk-lib.Stack[]

Return the stacks this stack depends on.

---

##### `environment`<sup>Required</sup> <a name="environment" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.environment"></a>

```typescript
public readonly environment: string;
```

- *Type:* string

The environment coordinates in which this stack is deployed.

In the form
`aws://account/region`. Use `stack.account` and `stack.region` to obtain
the specific values, no need to parse.

You can use this value to determine if two stacks are targeting the same
environment.

If either `stack.account` or `stack.region` are not concrete values (e.g.
`Aws.account` or `Aws.region`) the special strings `unknown-account` and/or
`unknown-region` will be used respectively to indicate this stack is
region/account-agnostic.

---

##### `nested`<sup>Required</sup> <a name="nested" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.nested"></a>

```typescript
public readonly nested: boolean;
```

- *Type:* boolean

Indicates if this is a nested stack, in which case `parentStack` will include a reference to it's parent.

---

##### `nestedStackParent`<sup>Optional</sup> <a name="nestedStackParent" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.nestedStackParent"></a>

```typescript
public readonly nestedStackParent: Stack;
```

- *Type:* aws-cdk-lib.Stack

If this is a nested stack, returns it's parent stack.

---

##### `nestedStackResource`<sup>Optional</sup> <a name="nestedStackResource" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.nestedStackResource"></a>

```typescript
public readonly nestedStackResource: CfnResource;
```

- *Type:* aws-cdk-lib.CfnResource

If this is a nested stack, this represents its `AWS::CloudFormation::Stack` resource.

`undefined` for top-level (non-nested) stacks.

---

##### `notificationArns`<sup>Required</sup> <a name="notificationArns" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.notificationArns"></a>

```typescript
public readonly notificationArns: string[];
```

- *Type:* string[]

Returns the list of notification Amazon Resource Names (ARNs) for the current stack.

---

##### `partition`<sup>Required</sup> <a name="partition" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.partition"></a>

```typescript
public readonly partition: string;
```

- *Type:* string

The partition in which this stack is defined.

---

##### `region`<sup>Required</sup> <a name="region" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.region"></a>

```typescript
public readonly region: string;
```

- *Type:* string

The AWS region into which this stack will be deployed (e.g. `us-west-2`).

This value is resolved according to the following rules:

1. The value provided to `env.region` when the stack is defined. This can
    either be a concerete region (e.g. `us-west-2`) or the `Aws.region`
    token.
3. `Aws.region`, which is represents the CloudFormation intrinsic reference
    `{ "Ref": "AWS::Region" }` encoded as a string token.

Preferably, you should use the return value as an opaque string and not
attempt to parse it to implement your logic. If you do, you must first
check that it is a concerete value an not an unresolved token. If this
value is an unresolved token (`Token.isUnresolved(stack.region)` returns
`true`), this implies that the user wishes that this stack will synthesize
into a **region-agnostic template**. In this case, your code should either
fail (throw an error, emit a synth error using `Annotations.of(construct).addError()`) or
implement some other region-agnostic behavior.

---

##### `stackId`<sup>Required</sup> <a name="stackId" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.stackId"></a>

```typescript
public readonly stackId: string;
```

- *Type:* string

The ID of the stack.

---

*Example*

```typescript
// After resolving, looks like
'arn:aws:cloudformation:us-west-2:123456789012:stack/teststack/51af3dc0-da77-11e4-872e-1234567db123'
```


##### `stackName`<sup>Required</sup> <a name="stackName" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.stackName"></a>

```typescript
public readonly stackName: string;
```

- *Type:* string

The concrete CloudFormation physical stack name.

This is either the name defined explicitly in the `stackName` prop or
allocated based on the stack's location in the construct tree. Stacks that
are directly defined under the app use their construct `id` as their stack
name. Stacks that are defined deeper within the tree will use a hashed naming
scheme based on the construct path to ensure uniqueness.

If you wish to obtain the deploy-time AWS::StackName intrinsic,
you can use `Aws.stackName` directly.

---

##### `synthesizer`<sup>Required</sup> <a name="synthesizer" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.synthesizer"></a>

```typescript
public readonly synthesizer: IStackSynthesizer;
```

- *Type:* aws-cdk-lib.IStackSynthesizer

Synthesis method for this stack.

---

##### `tags`<sup>Required</sup> <a name="tags" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.tags"></a>

```typescript
public readonly tags: TagManager;
```

- *Type:* aws-cdk-lib.TagManager

Tags to be applied to the stack.

---

##### `templateFile`<sup>Required</sup> <a name="templateFile" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.templateFile"></a>

```typescript
public readonly templateFile: string;
```

- *Type:* string

The name of the CloudFormation template file emitted to the output directory during synthesis.

Example value: `MyStack.template.json`

---

##### `templateOptions`<sup>Required</sup> <a name="templateOptions" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.templateOptions"></a>

```typescript
public readonly templateOptions: ITemplateOptions;
```

- *Type:* aws-cdk-lib.ITemplateOptions

Options for CloudFormation template (like version, transform, description).

---

##### `terminationProtection`<sup>Optional</sup> <a name="terminationProtection" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.terminationProtection"></a>

```typescript
public readonly terminationProtection: boolean;
```

- *Type:* boolean

Whether termination protection is enabled for this stack.

---

##### `urlSuffix`<sup>Required</sup> <a name="urlSuffix" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasic.property.urlSuffix"></a>

```typescript
public readonly urlSuffix: string;
```

- *Type:* string

The Amazon domain suffix for the region in which this stack is defined.

---


## Structs <a name="Structs" id="Structs"></a>

### ApiKeyDefinition <a name="ApiKeyDefinition" id="@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.Initializer"></a>

```typescript
import { ApiKeyDefinition } from '@mongodbatlas-awscdk/atlas-basic'

const apiKeyDefinition: ApiKeyDefinition = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.property.privateKey">privateKey</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.property.publicKey">publicKey</a></code> | <code>string</code> | *No description.* |

---

##### `privateKey`<sup>Optional</sup> <a name="privateKey" id="@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.property.privateKey"></a>

```typescript
public readonly privateKey: string;
```

- *Type:* string

---

##### `publicKey`<sup>Optional</sup> <a name="publicKey" id="@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition.property.publicKey"></a>

```typescript
public readonly publicKey: string;
```

- *Type:* string

---

### AtlasBasicProps <a name="AtlasBasicProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.Initializer"></a>

```typescript
import { AtlasBasicProps } from '@mongodbatlas-awscdk/atlas-basic'

const atlasBasicProps: AtlasBasicProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.analyticsReporting">analyticsReporting</a></code> | <code>boolean</code> | Include runtime versioning information in this Stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.description">description</a></code> | <code>string</code> | A description of the stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.env">env</a></code> | <code>aws-cdk-lib.Environment</code> | The AWS environment (account/region) where this stack will be deployed. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.stackName">stackName</a></code> | <code>string</code> | Name to deploy the stack with. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.synthesizer">synthesizer</a></code> | <code>aws-cdk-lib.IStackSynthesizer</code> | Synthesis method to use while deploying this stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.tags">tags</a></code> | <code>{[ key: string ]: string}</code> | Stack tags that will be applied to all the taggable resources and the stack itself. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.terminationProtection">terminationProtection</a></code> | <code>boolean</code> | Whether to enable termination protection for this stack. |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.apiKeys">apiKeys</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition">ApiKeyDefinition</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.clusterProps">clusterProps</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps">ClusterProps</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.projectProps">projectProps</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps">ProjectProps</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.dbUserProps">dbUserProps</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps">DatabaseUserProps</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.ipAccessListProps">ipAccessListProps</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps">IpAccessListProps</a></code> | *No description.* |

---

##### `analyticsReporting`<sup>Optional</sup> <a name="analyticsReporting" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.analyticsReporting"></a>

```typescript
public readonly analyticsReporting: boolean;
```

- *Type:* boolean
- *Default:* `analyticsReporting` setting of containing `App`, or value of 'aws:cdk:version-reporting' context key

Include runtime versioning information in this Stack.

---

##### `description`<sup>Optional</sup> <a name="description" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.description"></a>

```typescript
public readonly description: string;
```

- *Type:* string
- *Default:* No description.

A description of the stack.

---

##### `env`<sup>Optional</sup> <a name="env" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.env"></a>

```typescript
public readonly env: Environment;
```

- *Type:* aws-cdk-lib.Environment
- *Default:* The environment of the containing `Stage` if available, otherwise create the stack will be environment-agnostic.

The AWS environment (account/region) where this stack will be deployed.

Set the `region`/`account` fields of `env` to either a concrete value to
select the indicated environment (recommended for production stacks), or to
the values of environment variables
`CDK_DEFAULT_REGION`/`CDK_DEFAULT_ACCOUNT` to let the target environment
depend on the AWS credentials/configuration that the CDK CLI is executed
under (recommended for development stacks).

If the `Stack` is instantiated inside a `Stage`, any undefined
`region`/`account` fields from `env` will default to the same field on the
encompassing `Stage`, if configured there.

If either `region` or `account` are not set nor inherited from `Stage`, the
Stack will be considered "*environment-agnostic*"". Environment-agnostic
stacks can be deployed to any environment but may not be able to take
advantage of all features of the CDK. For example, they will not be able to
use environmental context lookups such as `ec2.Vpc.fromLookup` and will not
automatically translate Service Principals to the right format based on the
environment's AWS partition, and other such enhancements.

---

*Example*

```typescript
// Use a concrete account and region to deploy this stack to:
// `.account` and `.region` will simply return these values.
new Stack(app, 'Stack1', {
  env: {
    account: '123456789012',
    region: 'us-east-1'
  },
});

// Use the CLI's current credentials to determine the target environment:
// `.account` and `.region` will reflect the account+region the CLI
// is configured to use (based on the user CLI credentials)
new Stack(app, 'Stack2', {
  env: {
    account: process.env.CDK_DEFAULT_ACCOUNT,
    region: process.env.CDK_DEFAULT_REGION
  },
});

// Define multiple stacks stage associated with an environment
const myStage = new Stage(app, 'MyStage', {
  env: {
    account: '123456789012',
    region: 'us-east-1'
  }
});

// both of these stacks will use the stage's account/region:
// `.account` and `.region` will resolve to the concrete values as above
new MyStack(myStage, 'Stack1');
new YourStack(myStage, 'Stack2');

// Define an environment-agnostic stack:
// `.account` and `.region` will resolve to `{ "Ref": "AWS::AccountId" }` and `{ "Ref": "AWS::Region" }` respectively.
// which will only resolve to actual values by CloudFormation during deployment.
new MyStack(app, 'Stack1');
```


##### `stackName`<sup>Optional</sup> <a name="stackName" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.stackName"></a>

```typescript
public readonly stackName: string;
```

- *Type:* string
- *Default:* Derived from construct path.

Name to deploy the stack with.

---

##### `synthesizer`<sup>Optional</sup> <a name="synthesizer" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.synthesizer"></a>

```typescript
public readonly synthesizer: IStackSynthesizer;
```

- *Type:* aws-cdk-lib.IStackSynthesizer
- *Default:* `DefaultStackSynthesizer` if the `@aws-cdk/core:newStyleStackSynthesis` feature flag is set, `LegacyStackSynthesizer` otherwise.

Synthesis method to use while deploying this stack.

---

##### `tags`<sup>Optional</sup> <a name="tags" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.tags"></a>

```typescript
public readonly tags: {[ key: string ]: string};
```

- *Type:* {[ key: string ]: string}
- *Default:* {}

Stack tags that will be applied to all the taggable resources and the stack itself.

---

##### `terminationProtection`<sup>Optional</sup> <a name="terminationProtection" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.terminationProtection"></a>

```typescript
public readonly terminationProtection: boolean;
```

- *Type:* boolean
- *Default:* false

Whether to enable termination protection for this stack.

---

##### `apiKeys`<sup>Required</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.ApiKeyDefinition">ApiKeyDefinition</a>

---

##### `clusterProps`<sup>Required</sup> <a name="clusterProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.clusterProps"></a>

```typescript
public readonly clusterProps: ClusterProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps">ClusterProps</a>

---

##### `projectProps`<sup>Required</sup> <a name="projectProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.projectProps"></a>

```typescript
public readonly projectProps: ProjectProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps">ProjectProps</a>

---

##### `dbUserProps`<sup>Optional</sup> <a name="dbUserProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.dbUserProps"></a>

```typescript
public readonly dbUserProps: DatabaseUserProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps">DatabaseUserProps</a>

---

##### `ipAccessListProps`<sup>Optional</sup> <a name="ipAccessListProps" id="@mongodbatlas-awscdk/atlas-basic.AtlasBasicProps.property.ipAccessListProps"></a>

```typescript
public readonly ipAccessListProps: IpAccessListProps;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps">IpAccessListProps</a>

---

### ClusterProps <a name="ClusterProps" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.Initializer"></a>

```typescript
import { ClusterProps } from '@mongodbatlas-awscdk/atlas-basic'

const clusterProps: ClusterProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.advancedSettings">advancedSettings</a></code> | <code>@mongodbatlas-awscdk/cluster.ProcessArgs</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.apiKeys">apiKeys</a></code> | <code>@mongodbatlas-awscdk/cluster.ApiKeyDefinition</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.backupEnabled">backupEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.biConnector">biConnector</a></code> | <code>@mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.clusterType">clusterType</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.connectionStrings">connectionStrings</a></code> | <code>@mongodbatlas-awscdk/cluster.ConnectionStrings</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.diskSizeGb">diskSizeGb</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.encryptionAtRestProvider">encryptionAtRestProvider</a></code> | <code>@mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.labels">labels</a></code> | <code>@mongodbatlas-awscdk/cluster.CfnClusterPropsLabels[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.mongoDbMajorVersion">mongoDbMajorVersion</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.paused">paused</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.pitEnabled">pitEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.projectId">projectId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.replicationSpecs">replicationSpecs</a></code> | <code>@mongodbatlas-awscdk/cluster.AdvancedReplicationSpec[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.rootCertType">rootCertType</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.terminationProtectionEnabled">terminationProtectionEnabled</a></code> | <code>boolean</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.versionReleaseSystem">versionReleaseSystem</a></code> | <code>string</code> | *No description.* |

---

##### `advancedSettings`<sup>Optional</sup> <a name="advancedSettings" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.advancedSettings"></a>

```typescript
public readonly advancedSettings: ProcessArgs;
```

- *Type:* @mongodbatlas-awscdk/cluster.ProcessArgs

---

##### `apiKeys`<sup>Optional</sup> <a name="apiKeys" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.apiKeys"></a>

```typescript
public readonly apiKeys: ApiKeyDefinition;
```

- *Type:* @mongodbatlas-awscdk/cluster.ApiKeyDefinition

---

##### `backupEnabled`<sup>Optional</sup> <a name="backupEnabled" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.backupEnabled"></a>

```typescript
public readonly backupEnabled: boolean;
```

- *Type:* boolean

---

##### `biConnector`<sup>Optional</sup> <a name="biConnector" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.biConnector"></a>

```typescript
public readonly biConnector: CfnClusterPropsBiConnector;
```

- *Type:* @mongodbatlas-awscdk/cluster.CfnClusterPropsBiConnector

---

##### `clusterType`<sup>Optional</sup> <a name="clusterType" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.clusterType"></a>

```typescript
public readonly clusterType: string;
```

- *Type:* string

---

##### `connectionStrings`<sup>Optional</sup> <a name="connectionStrings" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.connectionStrings"></a>

```typescript
public readonly connectionStrings: ConnectionStrings;
```

- *Type:* @mongodbatlas-awscdk/cluster.ConnectionStrings

---

##### `diskSizeGb`<sup>Optional</sup> <a name="diskSizeGb" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.diskSizeGb"></a>

```typescript
public readonly diskSizeGb: number;
```

- *Type:* number

---

##### `encryptionAtRestProvider`<sup>Optional</sup> <a name="encryptionAtRestProvider" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.encryptionAtRestProvider"></a>

```typescript
public readonly encryptionAtRestProvider: CfnClusterPropsEncryptionAtRestProvider;
```

- *Type:* @mongodbatlas-awscdk/cluster.CfnClusterPropsEncryptionAtRestProvider

---

##### `labels`<sup>Optional</sup> <a name="labels" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.labels"></a>

```typescript
public readonly labels: CfnClusterPropsLabels[];
```

- *Type:* @mongodbatlas-awscdk/cluster.CfnClusterPropsLabels[]

---

##### `mongoDbMajorVersion`<sup>Optional</sup> <a name="mongoDbMajorVersion" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.mongoDbMajorVersion"></a>

```typescript
public readonly mongoDbMajorVersion: string;
```

- *Type:* string

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

---

##### `paused`<sup>Optional</sup> <a name="paused" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.paused"></a>

```typescript
public readonly paused: boolean;
```

- *Type:* boolean

---

##### `pitEnabled`<sup>Optional</sup> <a name="pitEnabled" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.pitEnabled"></a>

```typescript
public readonly pitEnabled: boolean;
```

- *Type:* boolean

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

---

##### `replicationSpecs`<sup>Optional</sup> <a name="replicationSpecs" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.replicationSpecs"></a>

```typescript
public readonly replicationSpecs: AdvancedReplicationSpec[];
```

- *Type:* @mongodbatlas-awscdk/cluster.AdvancedReplicationSpec[]

---

##### `rootCertType`<sup>Optional</sup> <a name="rootCertType" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.rootCertType"></a>

```typescript
public readonly rootCertType: string;
```

- *Type:* string

---

##### `terminationProtectionEnabled`<sup>Optional</sup> <a name="terminationProtectionEnabled" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.terminationProtectionEnabled"></a>

```typescript
public readonly terminationProtectionEnabled: boolean;
```

- *Type:* boolean

---

##### `versionReleaseSystem`<sup>Optional</sup> <a name="versionReleaseSystem" id="@mongodbatlas-awscdk/atlas-basic.ClusterProps.property.versionReleaseSystem"></a>

```typescript
public readonly versionReleaseSystem: string;
```

- *Type:* string

---

### DatabaseUserProps <a name="DatabaseUserProps" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.Initializer"></a>

```typescript
import { DatabaseUserProps } from '@mongodbatlas-awscdk/atlas-basic'

const databaseUserProps: DatabaseUserProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.awsiamType">awsiamType</a></code> | <code>@mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsAwsiamType</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.databaseName">databaseName</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.deleteAfterDate">deleteAfterDate</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.labels">labels</a></code> | <code>@mongodbatlas-awscdk/database-user.LabelDefinition[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.ldapAuthType">ldapAuthType</a></code> | <code>@mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsLdapAuthType</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.password">password</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.projectId">projectId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.roles">roles</a></code> | <code>@mongodbatlas-awscdk/database-user.RoleDefinition[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.scopes">scopes</a></code> | <code>@mongodbatlas-awscdk/database-user.ScopeDefinition[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.username">username</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.x509Type">x509Type</a></code> | <code>@mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsX509Type</code> | *No description.* |

---

##### `awsiamType`<sup>Optional</sup> <a name="awsiamType" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.awsiamType"></a>

```typescript
public readonly awsiamType: CfnDatabaseUserPropsAwsiamType;
```

- *Type:* @mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsAwsiamType

---

##### `databaseName`<sup>Optional</sup> <a name="databaseName" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.databaseName"></a>

```typescript
public readonly databaseName: string;
```

- *Type:* string

---

##### `deleteAfterDate`<sup>Optional</sup> <a name="deleteAfterDate" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.deleteAfterDate"></a>

```typescript
public readonly deleteAfterDate: string;
```

- *Type:* string

---

##### `labels`<sup>Optional</sup> <a name="labels" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.labels"></a>

```typescript
public readonly labels: LabelDefinition[];
```

- *Type:* @mongodbatlas-awscdk/database-user.LabelDefinition[]

---

##### `ldapAuthType`<sup>Optional</sup> <a name="ldapAuthType" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.ldapAuthType"></a>

```typescript
public readonly ldapAuthType: CfnDatabaseUserPropsLdapAuthType;
```

- *Type:* @mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsLdapAuthType

---

##### `password`<sup>Optional</sup> <a name="password" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.password"></a>

```typescript
public readonly password: string;
```

- *Type:* string

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

---

##### `roles`<sup>Optional</sup> <a name="roles" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.roles"></a>

```typescript
public readonly roles: RoleDefinition[];
```

- *Type:* @mongodbatlas-awscdk/database-user.RoleDefinition[]

---

##### `scopes`<sup>Optional</sup> <a name="scopes" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.scopes"></a>

```typescript
public readonly scopes: ScopeDefinition[];
```

- *Type:* @mongodbatlas-awscdk/database-user.ScopeDefinition[]

---

##### `username`<sup>Optional</sup> <a name="username" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.username"></a>

```typescript
public readonly username: string;
```

- *Type:* string

---

##### `x509Type`<sup>Optional</sup> <a name="x509Type" id="@mongodbatlas-awscdk/atlas-basic.DatabaseUserProps.property.x509Type"></a>

```typescript
public readonly x509Type: CfnDatabaseUserPropsX509Type;
```

- *Type:* @mongodbatlas-awscdk/database-user.CfnDatabaseUserPropsX509Type

---

### IpAccessListProps <a name="IpAccessListProps" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.Initializer"></a>

```typescript
import { IpAccessListProps } from '@mongodbatlas-awscdk/atlas-basic'

const ipAccessListProps: IpAccessListProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.accessList">accessList</a></code> | <code>@mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.listOptions">listOptions</a></code> | <code>@mongodbatlas-awscdk/project-ip-access-list.ListOptions</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.projectId">projectId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.totalCount">totalCount</a></code> | <code>number</code> | *No description.* |

---

##### `accessList`<sup>Required</sup> <a name="accessList" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.accessList"></a>

```typescript
public readonly accessList: AccessListDefinition[];
```

- *Type:* @mongodbatlas-awscdk/project-ip-access-list.AccessListDefinition[]

---

##### `listOptions`<sup>Optional</sup> <a name="listOptions" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.listOptions"></a>

```typescript
public readonly listOptions: ListOptions;
```

- *Type:* @mongodbatlas-awscdk/project-ip-access-list.ListOptions

---

##### `projectId`<sup>Optional</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

---

##### `totalCount`<sup>Optional</sup> <a name="totalCount" id="@mongodbatlas-awscdk/atlas-basic.IpAccessListProps.property.totalCount"></a>

```typescript
public readonly totalCount: number;
```

- *Type:* number

---

### ProjectProps <a name="ProjectProps" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.Initializer"></a>

```typescript
import { ProjectProps } from '@mongodbatlas-awscdk/atlas-basic'

const projectProps: ProjectProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.orgId">orgId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.clusterCount">clusterCount</a></code> | <code>number</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.name">name</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectApiKeys">projectApiKeys</a></code> | <code>@mongodbatlas-awscdk/project.ProjectApiKey[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectOwnerId">projectOwnerId</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectSettings">projectSettings</a></code> | <code>@mongodbatlas-awscdk/project.ProjectSettings</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectTeams">projectTeams</a></code> | <code>@mongodbatlas-awscdk/project.ProjectTeam[]</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.withDefaultAlertsSettings">withDefaultAlertsSettings</a></code> | <code>boolean</code> | *No description.* |

---

##### `orgId`<sup>Required</sup> <a name="orgId" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.orgId"></a>

```typescript
public readonly orgId: string;
```

- *Type:* string

---

##### `clusterCount`<sup>Optional</sup> <a name="clusterCount" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.clusterCount"></a>

```typescript
public readonly clusterCount: number;
```

- *Type:* number

---

##### `name`<sup>Optional</sup> <a name="name" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.name"></a>

```typescript
public readonly name: string;
```

- *Type:* string

---

##### `projectApiKeys`<sup>Optional</sup> <a name="projectApiKeys" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectApiKeys"></a>

```typescript
public readonly projectApiKeys: ProjectApiKey[];
```

- *Type:* @mongodbatlas-awscdk/project.ProjectApiKey[]

---

##### `projectOwnerId`<sup>Optional</sup> <a name="projectOwnerId" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectOwnerId"></a>

```typescript
public readonly projectOwnerId: string;
```

- *Type:* string

---

##### `projectSettings`<sup>Optional</sup> <a name="projectSettings" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectSettings"></a>

```typescript
public readonly projectSettings: ProjectSettings;
```

- *Type:* @mongodbatlas-awscdk/project.ProjectSettings

---

##### `projectTeams`<sup>Optional</sup> <a name="projectTeams" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.projectTeams"></a>

```typescript
public readonly projectTeams: ProjectTeam[];
```

- *Type:* @mongodbatlas-awscdk/project.ProjectTeam[]

---

##### `withDefaultAlertsSettings`<sup>Optional</sup> <a name="withDefaultAlertsSettings" id="@mongodbatlas-awscdk/atlas-basic.ProjectProps.property.withDefaultAlertsSettings"></a>

```typescript
public readonly withDefaultAlertsSettings: boolean;
```

- *Type:* boolean

---



