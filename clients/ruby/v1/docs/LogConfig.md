# NomadClient::LogConfig

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **disabled** | **Boolean** |  | [optional] |
| **enabled** | **Boolean** |  | [optional] |
| **max_file_size_mb** | **Integer** |  | [optional] |
| **max_files** | **Integer** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::LogConfig.new(
  disabled: null,
  enabled: null,
  max_file_size_mb: null,
  max_files: null
)
```

