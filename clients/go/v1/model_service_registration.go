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

// ServiceRegistration struct for ServiceRegistration
type ServiceRegistration struct {
	Address *string `json:"Address,omitempty"`
	AllocID *string `json:"AllocID,omitempty"`
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	Datacenter *string `json:"Datacenter,omitempty"`
	ID *string `json:"ID,omitempty"`
	JobID *string `json:"JobID,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	NodeID *string `json:"NodeID,omitempty"`
	Port *int32 `json:"Port,omitempty"`
	ServiceName *string `json:"ServiceName,omitempty"`
	Tags *[]string `json:"Tags,omitempty"`
}

// NewServiceRegistration instantiates a new ServiceRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceRegistration() *ServiceRegistration {
	this := ServiceRegistration{}
	return &this
}

// NewServiceRegistrationWithDefaults instantiates a new ServiceRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceRegistrationWithDefaults() *ServiceRegistration {
	this := ServiceRegistration{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *ServiceRegistration) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *ServiceRegistration) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *ServiceRegistration) SetAddress(v string) {
	o.Address = &v
}

// GetAllocID returns the AllocID field value if set, zero value otherwise.
func (o *ServiceRegistration) GetAllocID() string {
	if o == nil || o.AllocID == nil {
		var ret string
		return ret
	}
	return *o.AllocID
}

// GetAllocIDOk returns a tuple with the AllocID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetAllocIDOk() (*string, bool) {
	if o == nil || o.AllocID == nil {
		return nil, false
	}
	return o.AllocID, true
}

// HasAllocID returns a boolean if a field has been set.
func (o *ServiceRegistration) HasAllocID() bool {
	if o != nil && o.AllocID != nil {
		return true
	}

	return false
}

// SetAllocID gets a reference to the given string and assigns it to the AllocID field.
func (o *ServiceRegistration) SetAllocID(v string) {
	o.AllocID = &v
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *ServiceRegistration) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *ServiceRegistration) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *ServiceRegistration) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetDatacenter returns the Datacenter field value if set, zero value otherwise.
func (o *ServiceRegistration) GetDatacenter() string {
	if o == nil || o.Datacenter == nil {
		var ret string
		return ret
	}
	return *o.Datacenter
}

// GetDatacenterOk returns a tuple with the Datacenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetDatacenterOk() (*string, bool) {
	if o == nil || o.Datacenter == nil {
		return nil, false
	}
	return o.Datacenter, true
}

// HasDatacenter returns a boolean if a field has been set.
func (o *ServiceRegistration) HasDatacenter() bool {
	if o != nil && o.Datacenter != nil {
		return true
	}

	return false
}

// SetDatacenter gets a reference to the given string and assigns it to the Datacenter field.
func (o *ServiceRegistration) SetDatacenter(v string) {
	o.Datacenter = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *ServiceRegistration) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *ServiceRegistration) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *ServiceRegistration) SetID(v string) {
	o.ID = &v
}

// GetJobID returns the JobID field value if set, zero value otherwise.
func (o *ServiceRegistration) GetJobID() string {
	if o == nil || o.JobID == nil {
		var ret string
		return ret
	}
	return *o.JobID
}

// GetJobIDOk returns a tuple with the JobID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetJobIDOk() (*string, bool) {
	if o == nil || o.JobID == nil {
		return nil, false
	}
	return o.JobID, true
}

// HasJobID returns a boolean if a field has been set.
func (o *ServiceRegistration) HasJobID() bool {
	if o != nil && o.JobID != nil {
		return true
	}

	return false
}

// SetJobID gets a reference to the given string and assigns it to the JobID field.
func (o *ServiceRegistration) SetJobID(v string) {
	o.JobID = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *ServiceRegistration) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *ServiceRegistration) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *ServiceRegistration) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ServiceRegistration) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ServiceRegistration) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ServiceRegistration) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodeID returns the NodeID field value if set, zero value otherwise.
func (o *ServiceRegistration) GetNodeID() string {
	if o == nil || o.NodeID == nil {
		var ret string
		return ret
	}
	return *o.NodeID
}

// GetNodeIDOk returns a tuple with the NodeID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetNodeIDOk() (*string, bool) {
	if o == nil || o.NodeID == nil {
		return nil, false
	}
	return o.NodeID, true
}

// HasNodeID returns a boolean if a field has been set.
func (o *ServiceRegistration) HasNodeID() bool {
	if o != nil && o.NodeID != nil {
		return true
	}

	return false
}

// SetNodeID gets a reference to the given string and assigns it to the NodeID field.
func (o *ServiceRegistration) SetNodeID(v string) {
	o.NodeID = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *ServiceRegistration) GetPort() int32 {
	if o == nil || o.Port == nil {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetPortOk() (*int32, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *ServiceRegistration) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *ServiceRegistration) SetPort(v int32) {
	o.Port = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ServiceRegistration) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ServiceRegistration) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ServiceRegistration) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ServiceRegistration) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceRegistration) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ServiceRegistration) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ServiceRegistration) SetTags(v []string) {
	o.Tags = &v
}

func (o ServiceRegistration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["Address"] = o.Address
	}
	if o.AllocID != nil {
		toSerialize["AllocID"] = o.AllocID
	}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.Datacenter != nil {
		toSerialize["Datacenter"] = o.Datacenter
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.JobID != nil {
		toSerialize["JobID"] = o.JobID
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.NodeID != nil {
		toSerialize["NodeID"] = o.NodeID
	}
	if o.Port != nil {
		toSerialize["Port"] = o.Port
	}
	if o.ServiceName != nil {
		toSerialize["ServiceName"] = o.ServiceName
	}
	if o.Tags != nil {
		toSerialize["Tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableServiceRegistration struct {
	value *ServiceRegistration
	isSet bool
}

func (v NullableServiceRegistration) Get() *ServiceRegistration {
	return v.value
}

func (v *NullableServiceRegistration) Set(val *ServiceRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceRegistration(val *ServiceRegistration) *NullableServiceRegistration {
	return &NullableServiceRegistration{value: val, isSet: true}
}

func (v NullableServiceRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


