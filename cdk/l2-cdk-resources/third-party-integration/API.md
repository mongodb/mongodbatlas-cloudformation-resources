# API Reference <a name="API Reference" id="api-reference"></a>

## Constructs <a name="Constructs" id="Constructs"></a>

### DatadogIntegration <a name="DatadogIntegration" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration"></a>

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.Initializer"></a>

```typescript
import { DatadogIntegration } from '@mongodbatlas-awscdk/atlas-integrations'

new DatadogIntegration(scope: Construct, id: string, props: DatadogIntegrationProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps">DatadogIntegrationProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps">DatadogIntegrationProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.isConstruct"></a>

```typescript
import { DatadogIntegration } from '@mongodbatlas-awscdk/atlas-integrations'

DatadogIntegration.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.property.cfnThirdPartyIntegration">cfnThirdPartyIntegration</a></code> | <code>@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `cfnThirdPartyIntegration`<sup>Required</sup> <a name="cfnThirdPartyIntegration" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegration.property.cfnThirdPartyIntegration"></a>

```typescript
public readonly cfnThirdPartyIntegration: CfnThirdPartyIntegration;
```

- *Type:* @mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration

---


### MicrosoftTeamsIntegration <a name="MicrosoftTeamsIntegration" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration"></a>

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.Initializer"></a>

```typescript
import { MicrosoftTeamsIntegration } from '@mongodbatlas-awscdk/atlas-integrations'

new MicrosoftTeamsIntegration(scope: Construct, id: string, props: MicrosoftTeamsIntegrationProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps">MicrosoftTeamsIntegrationProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps">MicrosoftTeamsIntegrationProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.isConstruct"></a>

```typescript
import { MicrosoftTeamsIntegration } from '@mongodbatlas-awscdk/atlas-integrations'

MicrosoftTeamsIntegration.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.property.cfnThirdPartyIntegration">cfnThirdPartyIntegration</a></code> | <code>@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `cfnThirdPartyIntegration`<sup>Required</sup> <a name="cfnThirdPartyIntegration" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegration.property.cfnThirdPartyIntegration"></a>

```typescript
public readonly cfnThirdPartyIntegration: CfnThirdPartyIntegration;
```

- *Type:* @mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration

---


### PagerDutyIntegration <a name="PagerDutyIntegration" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration"></a>

#### Initializers <a name="Initializers" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.Initializer"></a>

```typescript
import { PagerDutyIntegration } from '@mongodbatlas-awscdk/atlas-integrations'

new PagerDutyIntegration(scope: Construct, id: string, props: PagerDutyIntegrationProps)
```

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.Initializer.parameter.scope">scope</a></code> | <code>constructs.Construct</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.Initializer.parameter.id">id</a></code> | <code>string</code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.Initializer.parameter.props">props</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps">PagerDutyIntegrationProps</a></code> | *No description.* |

---

##### `scope`<sup>Required</sup> <a name="scope" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.Initializer.parameter.scope"></a>

- *Type:* constructs.Construct

---

##### `id`<sup>Required</sup> <a name="id" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.Initializer.parameter.id"></a>

- *Type:* string

---

##### `props`<sup>Required</sup> <a name="props" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.Initializer.parameter.props"></a>

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps">PagerDutyIntegrationProps</a>

---

#### Methods <a name="Methods" id="Methods"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.toString">toString</a></code> | Returns a string representation of this construct. |

---

##### `toString` <a name="toString" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.toString"></a>

```typescript
public toString(): string
```

Returns a string representation of this construct.

#### Static Functions <a name="Static Functions" id="Static Functions"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.isConstruct">isConstruct</a></code> | Checks if `x` is a construct. |

---

##### ~~`isConstruct`~~ <a name="isConstruct" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.isConstruct"></a>

```typescript
import { PagerDutyIntegration } from '@mongodbatlas-awscdk/atlas-integrations'

PagerDutyIntegration.isConstruct(x: any)
```

Checks if `x` is a construct.

###### `x`<sup>Required</sup> <a name="x" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.isConstruct.parameter.x"></a>

- *Type:* any

Any object.

---

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.property.node">node</a></code> | <code>constructs.Node</code> | The tree node. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.property.cfnThirdPartyIntegration">cfnThirdPartyIntegration</a></code> | <code>@mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration</code> | *No description.* |

---

##### `node`<sup>Required</sup> <a name="node" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.property.node"></a>

```typescript
public readonly node: Node;
```

- *Type:* constructs.Node

The tree node.

---

##### `cfnThirdPartyIntegration`<sup>Required</sup> <a name="cfnThirdPartyIntegration" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegration.property.cfnThirdPartyIntegration"></a>

```typescript
public readonly cfnThirdPartyIntegration: CfnThirdPartyIntegration;
```

- *Type:* @mongodbatlas-awscdk/third-party-integration.CfnThirdPartyIntegration

---


## Structs <a name="Structs" id="Structs"></a>

### DatadogIntegrationProps <a name="DatadogIntegrationProps" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.Initializer"></a>

```typescript
import { DatadogIntegrationProps } from '@mongodbatlas-awscdk/atlas-integrations'

const datadogIntegrationProps: DatadogIntegrationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.property.profile">profile</a></code> | <code>string</code> | Atlas API keys. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.property.apiKey">apiKey</a></code> | <code>string</code> | Key that allows MongoDB Cloud to access your Datadog account. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.property.region">region</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogRegion">DatadogRegion</a></code> | Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API. |

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

Atlas API keys.

---

##### `apiKey`<sup>Required</sup> <a name="apiKey" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.property.apiKey"></a>

```typescript
public readonly apiKey: string;
```

- *Type:* string

Key that allows MongoDB Cloud to access your Datadog account.

---

##### `region`<sup>Required</sup> <a name="region" id="@mongodbatlas-awscdk/atlas-integrations.DatadogIntegrationProps.property.region"></a>

```typescript
public readonly region: DatadogRegion;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogRegion">DatadogRegion</a>

Two-letter code that indicates which regional URL MongoDB uses to access the Datadog API.

---

### MicrosoftTeamsIntegrationProps <a name="MicrosoftTeamsIntegrationProps" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps.Initializer"></a>

```typescript
import { MicrosoftTeamsIntegrationProps } from '@mongodbatlas-awscdk/atlas-integrations'

const microsoftTeamsIntegrationProps: MicrosoftTeamsIntegrationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps.property.profile">profile</a></code> | <code>string</code> | Atlas API keys. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps.property.microsoftTeamsWebhookUrl">microsoftTeamsWebhookUrl</a></code> | <code>string</code> | Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications. |

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

Atlas API keys.

---

##### `microsoftTeamsWebhookUrl`<sup>Required</sup> <a name="microsoftTeamsWebhookUrl" id="@mongodbatlas-awscdk/atlas-integrations.MicrosoftTeamsIntegrationProps.property.microsoftTeamsWebhookUrl"></a>

```typescript
public readonly microsoftTeamsWebhookUrl: string;
```

- *Type:* string

Endpoint web address of the Microsoft Teams webhook to which MongoDB Cloud sends notifications.

---

### PagerDutyIntegrationProps <a name="PagerDutyIntegrationProps" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.Initializer"></a>

```typescript
import { PagerDutyIntegrationProps } from '@mongodbatlas-awscdk/atlas-integrations'

const pagerDutyIntegrationProps: PagerDutyIntegrationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.property.profile">profile</a></code> | <code>string</code> | Atlas API keys. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.property.region">region</a></code> | <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyRegion">PagerDutyRegion</a></code> | PagerDuty region that indicates the API Uniform Resource Locator (URL) to use. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.property.serviceKey">serviceKey</a></code> | <code>string</code> | Service key associated with your PagerDuty account. |

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

Atlas API keys.

---

##### `region`<sup>Required</sup> <a name="region" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.property.region"></a>

```typescript
public readonly region: PagerDutyRegion;
```

- *Type:* <a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyRegion">PagerDutyRegion</a>

PagerDuty region that indicates the API Uniform Resource Locator (URL) to use.

---

##### `serviceKey`<sup>Required</sup> <a name="serviceKey" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyIntegrationProps.property.serviceKey"></a>

```typescript
public readonly serviceKey: string;
```

- *Type:* string

Service key associated with your PagerDuty account.

---

### ThirdPartyIntegrationProps <a name="ThirdPartyIntegrationProps" id="@mongodbatlas-awscdk/atlas-integrations.ThirdPartyIntegrationProps"></a>

#### Initializer <a name="Initializer" id="@mongodbatlas-awscdk/atlas-integrations.ThirdPartyIntegrationProps.Initializer"></a>

```typescript
import { ThirdPartyIntegrationProps } from '@mongodbatlas-awscdk/atlas-integrations'

const thirdPartyIntegrationProps: ThirdPartyIntegrationProps = { ... }
```

#### Properties <a name="Properties" id="Properties"></a>

| **Name** | **Type** | **Description** |
| --- | --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.ThirdPartyIntegrationProps.property.projectId">projectId</a></code> | <code>string</code> | Unique 24-hexadecimal digit string that identifies your project. |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.ThirdPartyIntegrationProps.property.profile">profile</a></code> | <code>string</code> | Atlas API keys. |

---

##### `projectId`<sup>Required</sup> <a name="projectId" id="@mongodbatlas-awscdk/atlas-integrations.ThirdPartyIntegrationProps.property.projectId"></a>

```typescript
public readonly projectId: string;
```

- *Type:* string

Unique 24-hexadecimal digit string that identifies your project.

---

##### `profile`<sup>Optional</sup> <a name="profile" id="@mongodbatlas-awscdk/atlas-integrations.ThirdPartyIntegrationProps.property.profile"></a>

```typescript
public readonly profile: string;
```

- *Type:* string

Atlas API keys.

---



## Enums <a name="Enums" id="Enums"></a>

### DatadogRegion <a name="DatadogRegion" id="@mongodbatlas-awscdk/atlas-integrations.DatadogRegion"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogRegion.US">US</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogRegion.EU">EU</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogRegion.US3">US3</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.DatadogRegion.US5">US5</a></code> | *No description.* |

---

##### `US` <a name="US" id="@mongodbatlas-awscdk/atlas-integrations.DatadogRegion.US"></a>

---


##### `EU` <a name="EU" id="@mongodbatlas-awscdk/atlas-integrations.DatadogRegion.EU"></a>

---


##### `US3` <a name="US3" id="@mongodbatlas-awscdk/atlas-integrations.DatadogRegion.US3"></a>

---


##### `US5` <a name="US5" id="@mongodbatlas-awscdk/atlas-integrations.DatadogRegion.US5"></a>

---


### PagerDutyRegion <a name="PagerDutyRegion" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyRegion"></a>

#### Members <a name="Members" id="Members"></a>

| **Name** | **Description** |
| --- | --- |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyRegion.US">US</a></code> | *No description.* |
| <code><a href="#@mongodbatlas-awscdk/atlas-integrations.PagerDutyRegion.EU">EU</a></code> | *No description.* |

---

##### `US` <a name="US" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyRegion.US"></a>

---


##### `EU` <a name="EU" id="@mongodbatlas-awscdk/atlas-integrations.PagerDutyRegion.EU"></a>

---

