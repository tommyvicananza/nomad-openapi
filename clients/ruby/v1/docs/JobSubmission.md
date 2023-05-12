# NomadClient::JobSubmission

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **format** | **String** |  | [optional] |
| **source** | **String** |  | [optional] |
| **variable_flags** | **Hash&lt;String, String&gt;** |  | [optional] |
| **variables** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobSubmission.new(
  format: null,
  source: null,
  variable_flags: null,
  variables: null
)
```

