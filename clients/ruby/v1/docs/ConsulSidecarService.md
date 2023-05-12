# NomadClient::ConsulSidecarService

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **disable_default_tcp_check** | **Boolean** |  | [optional] |
| **meta** | **Hash&lt;String, String&gt;** |  | [optional] |
| **port** | **String** |  | [optional] |
| **proxy** | [**ConsulProxy**](ConsulProxy.md) |  | [optional] |
| **tags** | **Array&lt;String&gt;** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::ConsulSidecarService.new(
  disable_default_tcp_check: null,
  meta: null,
  port: null,
  proxy: null,
  tags: null
)
```

