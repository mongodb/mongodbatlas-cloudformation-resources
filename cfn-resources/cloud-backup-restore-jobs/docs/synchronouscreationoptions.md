# MongoDB::Atlas::CloudBackUpRestoreJobs SynchronousCreationOptions

Options that needs to be set to control the synchronous creation flow, this options need to be set if EnableSynchronousCreation is se to TRUE

## Syntax

To declare this entity in your AWS CloudFormation template, use the following syntax:

### JSON

<pre>
{
    "<a href="#timeoutinseconds" title="TimeOutInSeconds">TimeOutInSeconds</a>" : <i>Integer</i>,
    "<a href="#callbackdelayseconds" title="CallbackDelaySeconds">CallbackDelaySeconds</a>" : <i>Integer</i>,
    "<a href="#returnsuccessiftimeout" title="ReturnSuccessIfTimeOut">ReturnSuccessIfTimeOut</a>" : <i>Boolean</i>
}
</pre>

### YAML

<pre>
<a href="#timeoutinseconds" title="TimeOutInSeconds">TimeOutInSeconds</a>: <i>Integer</i>
<a href="#callbackdelayseconds" title="CallbackDelaySeconds">CallbackDelaySeconds</a>: <i>Integer</i>
<a href="#returnsuccessiftimeout" title="ReturnSuccessIfTimeOut">ReturnSuccessIfTimeOut</a>: <i>Boolean</i>
</pre>

## Properties

#### TimeOutInSeconds

The amount of time the process will wait until exiting with a success, default (1200 seconds)

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### CallbackDelaySeconds

Represents the time interval, measured in seconds, for the synchronous process to wait before checking again to verify if the job has been completed. example: if set to 20, it will chek every 20 seconds if the resource is completed, default (30 seconds)

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

#### ReturnSuccessIfTimeOut

if set to true, the process will return success, in the event of a timeOut default false

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/using-cfn-updating-stacks-update-behaviors.html#update-no-interrupt)

