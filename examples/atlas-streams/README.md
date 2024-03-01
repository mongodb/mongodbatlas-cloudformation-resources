# Using Atlas Streams with Cloudformation

Atlas Stream Processing is composed of multiple components, and users can leverage AWS CloudFormation to define a subset of these. To obtain more details on each of the components please refer to the [Atlas Stream Processing Documentation](https://www.mongodb.com/docs/atlas/atlas-sp/overview/#atlas-stream-processing-overview).



### Resources supported by AWS CloudFormation

- `MongoDB::Atlas::StreamInstance`: Enables creating, modifying, and deleting Stream Instances. as part of this resource, a computed `hostnames` attribute is available for connecting to the created instance.
- `MongoDB::Atlas::StreamConnection`: Enables creating, modifying, and deleting Stream Instance Connections, which serve as data sources and sinks for your instance.


> **NOTE:**  
> - Atlas Streams functionality is currently in **[Public Preview](https://www.mongodb.com/blog/post/atlas-stream-processing-now-in-public-preview)**
> - Please review [Limitations](https://www.mongodb.com/docs/atlas/atlas-sp/limitations/#std-label-atlas-sp-limitations) of Atlas Streams Processing during this preview period.


### Managing Stream Processors

Once a stream instance and its connections have been defined, `Stream Processors` can be created to define how your data will be processed in your instance. There are currently no resources defined in CloudFormation to provide this configuration. To obtain information on how this can be configured refer to [Manage Stream Processors](https://www.mongodb.com/docs/atlas/atlas-sp/manage-stream-processor/#manage-stream-processors).

Connect to your stream instance defined in CloudFormation using the `hostnames` output attribute.
This value can then be used to connect to the stream instance using `mongosh`, as described in the [Get Started Tutorial](https://www.mongodb.com/docs/atlas/atlas-sp/tutorial/). 
