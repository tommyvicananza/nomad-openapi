/*
 * Nomad
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.1.4
 * Contact: support@hashicorp.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// ConsulSidecarService struct for ConsulSidecarService
type ConsulSidecarService struct {
	DisableDefaultTCPCheck *bool `json:"DisableDefaultTCPCheck,omitempty"`
	Meta *map[string]string `json:"Meta,omitempty"`
	Port *string `json:"Port,omitempty"`
	Proxy *ConsulProxy `json:"Proxy,omitempty"`
	Tags *[]string `json:"Tags,omitempty"`
}

// NewConsulSidecarService instantiates a new ConsulSidecarService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConsulSidecarService() *ConsulSidecarService {
	this := ConsulSidecarService{}
	return &this
}

// NewConsulSidecarServiceWithDefaults instantiates a new ConsulSidecarService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConsulSidecarServiceWithDefaults() *ConsulSidecarService {
	this := ConsulSidecarService{}
	return &this
}

// GetDisableDefaultTCPCheck returns the DisableDefaultTCPCheck field value if set, zero value otherwise.
func (o *ConsulSidecarService) GetDisableDefaultTCPCheck() bool {
	if o == nil || o.DisableDefaultTCPCheck == nil {
		var ret bool
		return ret
	}
	return *o.DisableDefaultTCPCheck
}

// GetDisableDefaultTCPCheckOk returns a tuple with the DisableDefaultTCPCheck field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulSidecarService) GetDisableDefaultTCPCheckOk() (*bool, bool) {
	if o == nil || o.DisableDefaultTCPCheck == nil {
		return nil, false
	}
	return o.DisableDefaultTCPCheck, true
}

// HasDisableDefaultTCPCheck returns a boolean if a field has been set.
func (o *ConsulSidecarService) HasDisableDefaultTCPCheck() bool {
	if o != nil && o.DisableDefaultTCPCheck != nil {
		return true
	}

	return false
}

// SetDisableDefaultTCPCheck gets a reference to the given bool and assigns it to the DisableDefaultTCPCheck field.
func (o *ConsulSidecarService) SetDisableDefaultTCPCheck(v bool) {
	o.DisableDefaultTCPCheck = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ConsulSidecarService) GetMeta() map[string]string {
	if o == nil || o.Meta == nil {
		var ret map[string]string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulSidecarService) GetMetaOk() (*map[string]string, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ConsulSidecarService) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]string and assigns it to the Meta field.
func (o *ConsulSidecarService) SetMeta(v map[string]string) {
	o.Meta = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ConsulSidecarService) GetPort() string {
	if o == nil || o.Port == nil {
		var ret string
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulSidecarService) GetPortOk() (*string, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ConsulSidecarService) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given string and assigns it to the Port field.
func (o *ConsulSidecarService) SetPort(v string) {
	o.Port = &v
}

// GetProxy returns the Proxy field value if set, zero value otherwise.
func (o *ConsulSidecarService) GetProxy() ConsulProxy {
	if o == nil || o.Proxy == nil {
		var ret ConsulProxy
		return ret
	}
	return *o.Proxy
}

// GetProxyOk returns a tuple with the Proxy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulSidecarService) GetProxyOk() (*ConsulProxy, bool) {
	if o == nil || o.Proxy == nil {
		return nil, false
	}
	return o.Proxy, true
}

// HasProxy returns a boolean if a field has been set.
func (o *ConsulSidecarService) HasProxy() bool {
	if o != nil && o.Proxy != nil {
		return true
	}

	return false
}

// SetProxy gets a reference to the given ConsulProxy and assigns it to the Proxy field.
func (o *ConsulSidecarService) SetProxy(v ConsulProxy) {
	o.Proxy = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ConsulSidecarService) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConsulSidecarService) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ConsulSidecarService) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ConsulSidecarService) SetTags(v []string) {
	o.Tags = &v
}

func (o ConsulSidecarService) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisableDefaultTCPCheck != nil {
		toSerialize["DisableDefaultTCPCheck"] = o.DisableDefaultTCPCheck
	}
	if o.Meta != nil {
		toSerialize["Meta"] = o.Meta
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.Proxy != nil {
		toSerialize["Proxy"] = o.Proxy
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableConsulSidecarService struct {
	value *ConsulSidecarService
	isSet bool
}

func (v NullableConsulSidecarService) Get() *ConsulSidecarService {
	return v.value
}

func (v *NullableConsulSidecarService) Set(val *ConsulSidecarService) {
	v.value = val
	v.isSet = true
}

func (v NullableConsulSidecarService) IsSet() bool {
	return v.isSet
}

func (v *NullableConsulSidecarService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConsulSidecarService(val *ConsulSidecarService) *NullableConsulSidecarService {
	return &NullableConsulSidecarService{value: val, isSet: true}
}

func (v NullableConsulSidecarService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConsulSidecarService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


