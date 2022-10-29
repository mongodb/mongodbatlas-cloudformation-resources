# MongoDB::Atlas::ProjectIpAccessList listOptions

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#pagenum" title="PageNum">PageNum</a>" : <i>Integer</i>,
    "<a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>" : <i>Integer</i>,
    "<a href="#includecount" title="IncludeCount">IncludeCount</a>" : <i>Boolean</i>
}
</pre>

### YAML

<pre>
<a href="#pagenum" title="PageNum">PageNum</a>: <i>Integer</i>
<a href="#itemsperpage" title="ItemsPerPage">ItemsPerPage</a>: <i>Integer</i>
<a href="#includecount" title="IncludeCount">IncludeCount</a>: <i>Boolean</i>
</pre>

## Properties

#### PageNum

For paginated result sets, page of results to retrieve.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ItemsPerPage

For paginated result sets, the number of results to include per page.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### IncludeCount

Flag that indicates whether Atlas returns the totalCount parameter in the response body.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

