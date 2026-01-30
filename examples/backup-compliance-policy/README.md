# How to create a MongoDB::Atlas::BackupCompliancePolicy 

## Step 1: Activate the backup compliance policy resource in cloudformation
   Step a: Create Role using [execution-role.yaml](https://github.com/mongodb/mongodbatlas-cloudformation-resources/blob/master/examples/execution-role.yaml) in CFN resources folder.

   Step b: Search for Mongodb::Atlas::BackupCompliancePolicy resource.

         (CloudFormation > Public extensions > choose 'Third party' > Search with " Execution name prefix = MongoDB " )
   Step c: Select and activate
         Enter the RoleArn that is created in step 1.

   Your BackupCompliancePolicy Resource is ready to use.

## Step 2: Create template using [backup-compliance-policy.json](backup-compliance-policy.json) or [backup-compliance-policy-simple.json](backup-compliance-policy-simple.json)
    Note: Make sure you are providing appropriate values for: 
    1. ProjectId
    2. AuthorizedEmail
    3. AuthorizedUserFirstName
    4. AuthorizedUserLastName
    5. Profile (optional)
    6. CopyProtectionEnabled (optional)
    7. EncryptionAtRestEnabled (optional)
    8. PitEnabled (optional)
    9. RestoreWindowDays (optional)
    10. Policy items (Hourly, Daily, Weekly, Monthly, Yearly) (optional)
