## 1.0.9

### Enhancemetns:
* provider/destinations.go & provider/monitor.go: Add validation against API key auth

### Bug Fixes:
* provider/provider.go: Check for nil pointer on auth token when using API key auth

## 1.0.8 -> Skipped

## 1.0.7

### Enhancements:
* client/client.go: Added retry with exponential back off, accumulative 15 second backoff
* provider/provdier.go: API Keys are no longer required if using login cred auth
* provider/resources/view.go: `transforms` now supports json attributes for all types of transforms

### Bug Fixes
* provider/resources/view.go: Fix `transforms` json encoding string -> map[string]interface

## 1.0.6

### Features:
* provider/resources/objectgroup.go: `range` related filters now have support, although currently disabled
* provider/resources/objectgroup.go: `live_events` are now enabled
* provider/resources/objectgroup.go: `compression` type is now specifiable in `options`
* provider/resources/destinations.go: `destinations` are now a supported resource for kibana alerts
* provider/resources/monitors.go: `monitors` are now a supported resource for kibana alerts

### Enhancements:
* provider/resources/objectgroup.go: Pull apart `filter` to more reflect API. Extends supported `filter` types

### Bug Fixes
* client/bucket.go: Pass in bucket tagging header, enables more types on datasources/objectgroup & view
* provider/resources/objectgroup.go: Change `Compression` from computed to optional, moved into `options` block

## 1.0.5

### Enhancements:
* docs/resources/*: Documentation has been included for resources
* provider/resources/objectgroup.go: Removed `interval` and `index_retention.for_partition` for simplicity and redundancy

## 1.0.4

### Features:
* provider/resources/view.go: Enable users to provide multiple `preds` in views
* provider/provider.go: Enable the usage of API Key driven Auth

### Enhancements:
* provider/resources/view.go: Simplify `ViewRequestDTO`, rename to `ViewData`
* client/clientrequest.go: Add a note to JWT parsing error indicating failed authentication
* client/clientrequest.go: Break up request construction to allow for API Key driven auth 
  
## 1.0.3

### Enhancements:
* CHANGELOG.md: Adding
* provider/resources/usergroup.go: Convert `chaossearch_user_group` resource to take in json for `permissions` attribute

## 1.0.2

### Bug Fixes:
* provider/resources/index.go: Catch error thrown by `client.ReadIndexModel()`

## 1.0.1

### Enhancements:
* provider/resources/objectgroup.go: Add checking to ensure `index_retention.overall` and `target_active_index` have valid values
* provider/resources/index.go: Enable a `delete_timeout` to be configured for index deletion
* provider/resources/usergroup.gp: Flatten `user_group` schema for easier maintenance and readability
* provider/tests/provider_test.go: Add random UUID to to generated resources names

### Bug Fixes:
* provider/resources/view.go: Pass in `cacheable` into `CreateViewRequest`
* provider/resources/usergroup.go: Enable `user_group` ids to be stored in state, allowing for proper updates