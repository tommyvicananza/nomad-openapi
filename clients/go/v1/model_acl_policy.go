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

// ACLPolicy struct for ACLPolicy
type ACLPolicy struct {
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	Description *string `json:"Description,omitempty"`
	JobACL *JobACL `json:"JobACL,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	Name *string `json:"Name,omitempty"`
	Rules *string `json:"Rules,omitempty"`
}

// NewACLPolicy instantiates a new ACLPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACLPolicy() *ACLPolicy {
	this := ACLPolicy{}
	return &this
}

// NewACLPolicyWithDefaults instantiates a new ACLPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACLPolicyWithDefaults() *ACLPolicy {
	this := ACLPolicy{}
	return &this
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *ACLPolicy) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPolicy) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *ACLPolicy) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *ACLPolicy) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ACLPolicy) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ACLPolicy) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ACLPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetJobACL returns the JobACL field value if set, zero value otherwise.
func (o *ACLPolicy) GetJobACL() JobACL {
	if o == nil || o.JobACL == nil {
		var ret JobACL
		return ret
	}
	return *o.JobACL
}

// GetJobACLOk returns a tuple with the JobACL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPolicy) GetJobACLOk() (*JobACL, bool) {
	if o == nil || o.JobACL == nil {
		return nil, false
	}
	return o.JobACL, true
}

// HasJobACL returns a boolean if a field has been set.
func (o *ACLPolicy) HasJobACL() bool {
	if o != nil && o.JobACL != nil {
		return true
	}

	return false
}

// SetJobACL gets a reference to the given JobACL and assigns it to the JobACL field.
func (o *ACLPolicy) SetJobACL(v JobACL) {
	o.JobACL = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *ACLPolicy) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPolicy) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *ACLPolicy) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *ACLPolicy) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ACLPolicy) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPolicy) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ACLPolicy) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ACLPolicy) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ACLPolicy) GetRules() string {
	if o == nil || o.Rules == nil {
		var ret string
		return ret
	}
	return *o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACLPolicy) GetRulesOk() (*string, bool) {
	if o == nil || o.Rules == nil {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ACLPolicy) HasRules() bool {
	if o != nil && o.Rules != nil {
		return true
	}

	return false
}

// SetRules gets a reference to the given string and assigns it to the Rules field.
func (o *ACLPolicy) SetRules(v string) {
	o.Rules = &v
}

func (o ACLPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.JobACL != nil {
		toSerialize["JobACL"] = o.JobACL
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Rules != nil {
		toSerialize["Rules"] = o.Rules
	}
	return json.Marshal(toSerialize)
}

type NullableACLPolicy struct {
	value *ACLPolicy
	isSet bool
}

func (v NullableACLPolicy) Get() *ACLPolicy {
	return v.value
}

func (v *NullableACLPolicy) Set(val *ACLPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableACLPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableACLPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACLPolicy(val *ACLPolicy) *NullableACLPolicy {
	return &NullableACLPolicy{value: val, isSet: true}
}

func (v NullableACLPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACLPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


