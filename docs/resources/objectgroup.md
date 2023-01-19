# Object Group Resource

> Object groups organize and associate similar data files in your cloud storage for ChaosSearch indexing.

Creates an _Object Group_ to structure your data for Indexing

Check out the _Object Group_ documentation here: [Object Group Docs](https://docs.chaossearch.io/docs/creating-object-groups)

## Example Usage

```hcl
resource "chaossearch_object_group" "create-object-group" {
  bucket = "tf-provider"
  source = "chaossearch-tf-provider-test"
  live_events = "arn:partition:service:region:account-id:resource-id"
  format {
    type             = "CSV"
    column_delimiter = ","
    row_delimiter    = "\n"
    header_row       = true
    field_selection  = jsonencode([
        {
            "excludes": [
                "data",
                "bigobject"
            ],
            "type": "blacklist"
        }
    ])
    array_selection  = jsonencode([
        {
            "excludes": [
                "object.ids",
            ],
            "type": "blacklist"
        }
    ])
  }
  index_retention {
    overall       = -1
  }
  filter {
    field = "key"
    prefix = "ec"
  }
  filter {
    field = "key"
    regex = ".*"
  }
  filter {
    field = "storageClass"
    equals = "STANDARD"
  }
  options {
    compression = "GZIP"
  }
}
```

## Argument Reference

-   `bucket` - **(Required)** Name of the object group
-   `source` - **(Required)** Name of the bucket where your data is stored
-   `live_events` - **(Optional)** The SQS Arn for live event streaming
-   `format` - **(Optional)** A config block used for file formatting specifics

    -   `type` - **(Optional)** Specifies the type of file
    -   `column_delimiter` - **(Optional)** Specifies the character for separating columns
    -   `row_delimiter` - **(Optional)** Specifies the character for separating rows
    -   `header_row` - **(Optional)** Specifies if the file includes a header row
    -   `array_flatten_depth` - **(Optional)** How deeply nested should arrays be allowed to get before parsing stops. Defaults to unlimited.
    -   `strip_prefix` - **(Optional)** By default, all fields will be prefixed with 'root.'. If this is set to true, that prefix will be disabled.
    -   `horizontal` - **(Optional)** If true, array fields will be turned into new columns on each flattened message. Ifx false, array fields will be broadcast ito multiple flattened rows for each array item.
    -   `array_selection` - **(Optional)** Specify array fields to leave as flat, unparsed strings when indexing
    -   `field_selection` - **(Optional)** Specify object fields to leave as flat, unparsed strings when indexing

-   `index_retention` - **(Optional)** Config block for specifying how long an index is retained
    -   Only applies on update
    -   `overall` - **(Optional)** Takes the amount of days an index is retained
        -   use `-1` for an indefinite amount of time
-   `filter` - **(Optional)** Config block for housing filtering
    -   Note: Make sure that `prefix`, `regex` and `equals` are all broken into their own filter block
    -   `field` - **(Required)** What field the filter applies to
        -   Can be `key` and `storageClass`
    -   `prefix` - **(Optional)** Used with `key` field. The prefix the field must match for the file
    -   `regex` - **(Optional)** Used with `key` field. The regex for filtering files
    -   `equals` - **(Optional)** Used with `storageClass` field. Supplies the `storageClass` type of the S3 bucket
        -   Can be `STANDARD`, `STANDARD_IA`, `INTELLIGENT_TIERING`, `ONEZONE_IA`, `GLACIER`, `DEEP_ARCHIVE`, `REDUCED_REDUNDANCY`
-   `options` - **(Optional)** A config block for housing advanced settings
    -   `compression` - **(Optional)** Form of file compression being used
        -   Can either be `GZIP` or `SNAPPY`
