# Model Code Generator for go-hubspot

This tool automatically creates Go struct and internal names slice from CSV files output by HubSpot's property export function.
Please check Usage and Notes before use.

## Usage

1. Prepare HubSpot properties CSV file.
  (Export from [Settings] -> [Properties] -> [Export all properties])
2. Run following command.
3. `{lowercase_ObjectName}_model.go` will be generated.

```shell
$ make generate OBJECT=<ObjectName> FILEPATH=<PropertiesFilePath>
```

### Sample

```shell
$ make generate OBJECT=Contact FILEPATH=contact.csv
```

## Notes

- HubSpot CSV columns must be in the following order.
  This is the format output from HubSpot.

| Column Name          | Use |
|----------------------|-----|
| Name                 | ◯   |
| Internal name        | ◯   |
| Type                 | ◯   |
| Description          |     |
| Group name           |     |
| Form field           |     |
| Options              |     |
| Read only value      |     |
| Read only definition |     |
| Calculated           |     |
| External options     |     |
| Deleted              |     |
| Hubspot defined      |     |
| Created user         |     |
| Usages               |     |

- If a file with the same name already exists, it will be overwritten.
- HubSpot may or may not have clear delimiters, such as `hs_createdate` and `lifecyclestage`, and the name generated from the latter will not match Go's naming conventions.
  (e.g. `lifecyclestage` -> `Lifecyclestage`, Go's naming conventions `LifeCycleStage`)
- `$ go generate` does not allow dynamic parameters to be passed, so `$ go run` is used to run it.
