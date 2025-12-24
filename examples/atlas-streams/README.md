# Using Atlas Streams with Cloudformation

Atlas Stream Processing is composed of multiple components, and users can leverage AWS CloudFormation to define a subset of these. To obtain more details on each of the components please refer to the [Atlas Stream Processing Documentation](https://www.mongodb.com/docs/atlas/atlas-sp/overview/#atlas-stream-processing-overview).

### Resources supported by AWS CloudFormation

- `MongoDB::Atlas::StreamInstance`: Enables creating, modifying, and deleting Stream Instances. As part of this resource, a computed `hostnames` attribute is available for connecting to the created instance.
- `MongoDB::Atlas::StreamConnection`: Enables creating, modifying, and deleting Stream Instance Connections, which serve as data sources and sinks for your instance.
- `MongoDB::Atlas::StreamProcessor`: Enables creating, modifying, and deleting Stream Processors, which define how data is processed in your stream instance using aggregation pipelines.

Connect to your stream instance defined in CloudFormation using the `hostnames` output attribute.
This value can then be used to connect to the stream instance using `mongosh`, as described in the [Get Started Tutorial](https://www.mongodb.com/docs/atlas/atlas-sp/tutorial/).
