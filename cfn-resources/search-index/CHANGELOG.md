# Changelog

## (2023-10-12)

**(POSSIBLE CHANGE) FIXED MAPPING INPUTS**

Originally, the mapping fields were inputted like this:
``` json
mappings: {
fields: [
"ownershipDetails.assetName:string"
],
},
``` 

the problem with this approach  is that the user will only be able to input one field and the time, and also only the type
was specified, some data types like "autocomplete" may require extra parameters. (like analyzer, maxGrams, minGrams, etc),
also a field would have accepted only one data type, preventing the user to setting multiple types to a single field. 

we required a way to input dynamic structures to each property, but since cloud formation does not provide a way to
input dynamic objects as parameters, we decided to input the data type list as a stringify json,
representing a list of types like is shown in the MongoDB atlas UI

Original:
``` json
mappings: {
fields: [
"property_type:string"
],
},
``` 

New:
``` json
mappings: {
fields: [
"property_type:[{\"type\":\"autocomplete\"},{\"type\":\"string\"}}]"
],
},
```

Update process: previous stacks could be updated by just providing a new stack, with the new field structure, json structure
for current mappings can be reviewed by editing the current indexes with the JSON editor 
(select your cluster -> Search tab -> select the index from the list -> select Edit index Definition dropdown , and select With Json Editor)

**REMOVED MAPPINGS FROM CREATE ONLY PROPERTIES**

Mappings property is no longer CREATE ONLY, this change will allow the user to initiate a stack update, modifying the 
mappings without replacing the resource.