# V2 Private Endpoint Upgrade Guide

## Document purpose
With the Private Endpoint V2 we have introduced the next changes:

PrivateEndpoint resource has been decided into two resources:
- **MongoDB::Atlas::PrivateEndpointService**: new, resource responsible for creating a Private Endpoint Service, unconfigured, and then using it to configure a Private Endpoint with AWS or any other provider
- **MongoDB::Atlas::PrivateEndpoint**: the current resource is modified, it is responsible for adding a privateEndpoint to the Service

Users currently utilizing the previous version of **MongoDB::Atlas::PrivateEndpoint** may need to update their existing
stacks to accommodate the new split version. This document aims to provide guidance on upgrading your current stacks to
leverage the new split private endpoint version, all without the necessity of deleting and recreating any of your existing
MongoDB Atlas or AWS private endpoints

## Cloud Formation Limitations
before we start with the upgrade progress we need to understand the next limitations:

- For the migration we will be using the IMPORT feature provided by CloudFormation, the import process is not fully supported for third party resources, and it has the next limitations:
  
  - the import does not support UPDATE DELETE or CREATE resources, it can only READ existing resources
  - the import process does not support changes or additions on any Output, so all the outputs that we want to modify or any reference to the imported resources, must be removed or hardcoded before the import, and later updated 

## Update process:
In this example, we will walk through an recommended update procedure. We'll start with an existing stack that includes a project and a private endpoint (V1), and then proceed to upgrade it to utilize the new Splited private endpoint V2, all without needing to make any changes to your existing MongoDB Atlas resources

**Go from this:**
>MongoDB::Atlas::Project
	MongoDB::Atlas::PrivateEndpoint

**To this:**
>MongoDB::Atlas::Project
MongoDB::Atlas::PrivateEndpointService
AWS::EC2::VPCEndpoint
MongoDB::Atlas::PrivateEndpoint

### Steps:

- **Step-1 : Adding DeletionPolicy**:  Setting the property **"DeletionPolicy" : "Retain"** to all resources with the old MongoDB::Atlas::PrivateEndpoint resource
- **Step-2 : Remove References to original private endpoint**: Remove all MongoDB::Atlas::PrivateEndpoint resources, and removing or hard coding any reference or output of any old original private endoint
- **Step-3 : Import new resources**: Initiate an Import process for the new resources
- **Step-4 : Update the Outputs**: Update the current stack to reset any reference or output removed or hardcoded on step 2

We are going to start with the next example:

