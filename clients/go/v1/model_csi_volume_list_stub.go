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
	"time"
)

// CSIVolumeListStub struct for CSIVolumeListStub
type CSIVolumeListStub struct {
	AccessMode *string `json:"AccessMode,omitempty"`
	AttachmentMode *string `json:"AttachmentMode,omitempty"`
	ControllerRequired *bool `json:"ControllerRequired,omitempty"`
	ControllersExpected *int32 `json:"ControllersExpected,omitempty"`
	ControllersHealthy *int32 `json:"ControllersHealthy,omitempty"`
	CreateIndex *int32 `json:"CreateIndex,omitempty"`
	CurrentReaders *int32 `json:"CurrentReaders,omitempty"`
	CurrentWriters *int32 `json:"CurrentWriters,omitempty"`
	ExternalID *string `json:"ExternalID,omitempty"`
	ID *string `json:"ID,omitempty"`
	ModifyIndex *int32 `json:"ModifyIndex,omitempty"`
	Name *string `json:"Name,omitempty"`
	Namespace *string `json:"Namespace,omitempty"`
	NodesExpected *int32 `json:"NodesExpected,omitempty"`
	NodesHealthy *int32 `json:"NodesHealthy,omitempty"`
	PluginID *string `json:"PluginID,omitempty"`
	Provider *string `json:"Provider,omitempty"`
	ResourceExhausted *time.Time `json:"ResourceExhausted,omitempty"`
	Schedulable *bool `json:"Schedulable,omitempty"`
	Topologies *[]CSITopology `json:"Topologies,omitempty"`
}

// NewCSIVolumeListStub instantiates a new CSIVolumeListStub object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCSIVolumeListStub() *CSIVolumeListStub {
	this := CSIVolumeListStub{}
	return &this
}

// NewCSIVolumeListStubWithDefaults instantiates a new CSIVolumeListStub object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCSIVolumeListStubWithDefaults() *CSIVolumeListStub {
	this := CSIVolumeListStub{}
	return &this
}

// GetAccessMode returns the AccessMode field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetAccessMode() string {
	if o == nil || o.AccessMode == nil {
		var ret string
		return ret
	}
	return *o.AccessMode
}

// GetAccessModeOk returns a tuple with the AccessMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetAccessModeOk() (*string, bool) {
	if o == nil || o.AccessMode == nil {
		return nil, false
	}
	return o.AccessMode, true
}

// HasAccessMode returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasAccessMode() bool {
	if o != nil && o.AccessMode != nil {
		return true
	}

	return false
}

// SetAccessMode gets a reference to the given string and assigns it to the AccessMode field.
func (o *CSIVolumeListStub) SetAccessMode(v string) {
	o.AccessMode = &v
}

// GetAttachmentMode returns the AttachmentMode field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetAttachmentMode() string {
	if o == nil || o.AttachmentMode == nil {
		var ret string
		return ret
	}
	return *o.AttachmentMode
}

// GetAttachmentModeOk returns a tuple with the AttachmentMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetAttachmentModeOk() (*string, bool) {
	if o == nil || o.AttachmentMode == nil {
		return nil, false
	}
	return o.AttachmentMode, true
}

// HasAttachmentMode returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasAttachmentMode() bool {
	if o != nil && o.AttachmentMode != nil {
		return true
	}

	return false
}

// SetAttachmentMode gets a reference to the given string and assigns it to the AttachmentMode field.
func (o *CSIVolumeListStub) SetAttachmentMode(v string) {
	o.AttachmentMode = &v
}

// GetControllerRequired returns the ControllerRequired field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetControllerRequired() bool {
	if o == nil || o.ControllerRequired == nil {
		var ret bool
		return ret
	}
	return *o.ControllerRequired
}

// GetControllerRequiredOk returns a tuple with the ControllerRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetControllerRequiredOk() (*bool, bool) {
	if o == nil || o.ControllerRequired == nil {
		return nil, false
	}
	return o.ControllerRequired, true
}

// HasControllerRequired returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasControllerRequired() bool {
	if o != nil && o.ControllerRequired != nil {
		return true
	}

	return false
}

// SetControllerRequired gets a reference to the given bool and assigns it to the ControllerRequired field.
func (o *CSIVolumeListStub) SetControllerRequired(v bool) {
	o.ControllerRequired = &v
}

// GetControllersExpected returns the ControllersExpected field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetControllersExpected() int32 {
	if o == nil || o.ControllersExpected == nil {
		var ret int32
		return ret
	}
	return *o.ControllersExpected
}

// GetControllersExpectedOk returns a tuple with the ControllersExpected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetControllersExpectedOk() (*int32, bool) {
	if o == nil || o.ControllersExpected == nil {
		return nil, false
	}
	return o.ControllersExpected, true
}

// HasControllersExpected returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasControllersExpected() bool {
	if o != nil && o.ControllersExpected != nil {
		return true
	}

	return false
}

// SetControllersExpected gets a reference to the given int32 and assigns it to the ControllersExpected field.
func (o *CSIVolumeListStub) SetControllersExpected(v int32) {
	o.ControllersExpected = &v
}

// GetControllersHealthy returns the ControllersHealthy field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetControllersHealthy() int32 {
	if o == nil || o.ControllersHealthy == nil {
		var ret int32
		return ret
	}
	return *o.ControllersHealthy
}

// GetControllersHealthyOk returns a tuple with the ControllersHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetControllersHealthyOk() (*int32, bool) {
	if o == nil || o.ControllersHealthy == nil {
		return nil, false
	}
	return o.ControllersHealthy, true
}

// HasControllersHealthy returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasControllersHealthy() bool {
	if o != nil && o.ControllersHealthy != nil {
		return true
	}

	return false
}

// SetControllersHealthy gets a reference to the given int32 and assigns it to the ControllersHealthy field.
func (o *CSIVolumeListStub) SetControllersHealthy(v int32) {
	o.ControllersHealthy = &v
}

// GetCreateIndex returns the CreateIndex field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetCreateIndex() int32 {
	if o == nil || o.CreateIndex == nil {
		var ret int32
		return ret
	}
	return *o.CreateIndex
}

// GetCreateIndexOk returns a tuple with the CreateIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetCreateIndexOk() (*int32, bool) {
	if o == nil || o.CreateIndex == nil {
		return nil, false
	}
	return o.CreateIndex, true
}

// HasCreateIndex returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasCreateIndex() bool {
	if o != nil && o.CreateIndex != nil {
		return true
	}

	return false
}

// SetCreateIndex gets a reference to the given int32 and assigns it to the CreateIndex field.
func (o *CSIVolumeListStub) SetCreateIndex(v int32) {
	o.CreateIndex = &v
}

// GetCurrentReaders returns the CurrentReaders field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetCurrentReaders() int32 {
	if o == nil || o.CurrentReaders == nil {
		var ret int32
		return ret
	}
	return *o.CurrentReaders
}

// GetCurrentReadersOk returns a tuple with the CurrentReaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetCurrentReadersOk() (*int32, bool) {
	if o == nil || o.CurrentReaders == nil {
		return nil, false
	}
	return o.CurrentReaders, true
}

// HasCurrentReaders returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasCurrentReaders() bool {
	if o != nil && o.CurrentReaders != nil {
		return true
	}

	return false
}

// SetCurrentReaders gets a reference to the given int32 and assigns it to the CurrentReaders field.
func (o *CSIVolumeListStub) SetCurrentReaders(v int32) {
	o.CurrentReaders = &v
}

// GetCurrentWriters returns the CurrentWriters field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetCurrentWriters() int32 {
	if o == nil || o.CurrentWriters == nil {
		var ret int32
		return ret
	}
	return *o.CurrentWriters
}

// GetCurrentWritersOk returns a tuple with the CurrentWriters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetCurrentWritersOk() (*int32, bool) {
	if o == nil || o.CurrentWriters == nil {
		return nil, false
	}
	return o.CurrentWriters, true
}

// HasCurrentWriters returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasCurrentWriters() bool {
	if o != nil && o.CurrentWriters != nil {
		return true
	}

	return false
}

// SetCurrentWriters gets a reference to the given int32 and assigns it to the CurrentWriters field.
func (o *CSIVolumeListStub) SetCurrentWriters(v int32) {
	o.CurrentWriters = &v
}

// GetExternalID returns the ExternalID field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetExternalID() string {
	if o == nil || o.ExternalID == nil {
		var ret string
		return ret
	}
	return *o.ExternalID
}

// GetExternalIDOk returns a tuple with the ExternalID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetExternalIDOk() (*string, bool) {
	if o == nil || o.ExternalID == nil {
		return nil, false
	}
	return o.ExternalID, true
}

// HasExternalID returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasExternalID() bool {
	if o != nil && o.ExternalID != nil {
		return true
	}

	return false
}

// SetExternalID gets a reference to the given string and assigns it to the ExternalID field.
func (o *CSIVolumeListStub) SetExternalID(v string) {
	o.ExternalID = &v
}

// GetID returns the ID field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetID() string {
	if o == nil || o.ID == nil {
		var ret string
		return ret
	}
	return *o.ID
}

// GetIDOk returns a tuple with the ID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetIDOk() (*string, bool) {
	if o == nil || o.ID == nil {
		return nil, false
	}
	return o.ID, true
}

// HasID returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasID() bool {
	if o != nil && o.ID != nil {
		return true
	}

	return false
}

// SetID gets a reference to the given string and assigns it to the ID field.
func (o *CSIVolumeListStub) SetID(v string) {
	o.ID = &v
}

// GetModifyIndex returns the ModifyIndex field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetModifyIndex() int32 {
	if o == nil || o.ModifyIndex == nil {
		var ret int32
		return ret
	}
	return *o.ModifyIndex
}

// GetModifyIndexOk returns a tuple with the ModifyIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetModifyIndexOk() (*int32, bool) {
	if o == nil || o.ModifyIndex == nil {
		return nil, false
	}
	return o.ModifyIndex, true
}

// HasModifyIndex returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasModifyIndex() bool {
	if o != nil && o.ModifyIndex != nil {
		return true
	}

	return false
}

// SetModifyIndex gets a reference to the given int32 and assigns it to the ModifyIndex field.
func (o *CSIVolumeListStub) SetModifyIndex(v int32) {
	o.ModifyIndex = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CSIVolumeListStub) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *CSIVolumeListStub) SetNamespace(v string) {
	o.Namespace = &v
}

// GetNodesExpected returns the NodesExpected field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetNodesExpected() int32 {
	if o == nil || o.NodesExpected == nil {
		var ret int32
		return ret
	}
	return *o.NodesExpected
}

// GetNodesExpectedOk returns a tuple with the NodesExpected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetNodesExpectedOk() (*int32, bool) {
	if o == nil || o.NodesExpected == nil {
		return nil, false
	}
	return o.NodesExpected, true
}

// HasNodesExpected returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasNodesExpected() bool {
	if o != nil && o.NodesExpected != nil {
		return true
	}

	return false
}

// SetNodesExpected gets a reference to the given int32 and assigns it to the NodesExpected field.
func (o *CSIVolumeListStub) SetNodesExpected(v int32) {
	o.NodesExpected = &v
}

// GetNodesHealthy returns the NodesHealthy field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetNodesHealthy() int32 {
	if o == nil || o.NodesHealthy == nil {
		var ret int32
		return ret
	}
	return *o.NodesHealthy
}

// GetNodesHealthyOk returns a tuple with the NodesHealthy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetNodesHealthyOk() (*int32, bool) {
	if o == nil || o.NodesHealthy == nil {
		return nil, false
	}
	return o.NodesHealthy, true
}

// HasNodesHealthy returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasNodesHealthy() bool {
	if o != nil && o.NodesHealthy != nil {
		return true
	}

	return false
}

// SetNodesHealthy gets a reference to the given int32 and assigns it to the NodesHealthy field.
func (o *CSIVolumeListStub) SetNodesHealthy(v int32) {
	o.NodesHealthy = &v
}

// GetPluginID returns the PluginID field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetPluginID() string {
	if o == nil || o.PluginID == nil {
		var ret string
		return ret
	}
	return *o.PluginID
}

// GetPluginIDOk returns a tuple with the PluginID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetPluginIDOk() (*string, bool) {
	if o == nil || o.PluginID == nil {
		return nil, false
	}
	return o.PluginID, true
}

// HasPluginID returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasPluginID() bool {
	if o != nil && o.PluginID != nil {
		return true
	}

	return false
}

// SetPluginID gets a reference to the given string and assigns it to the PluginID field.
func (o *CSIVolumeListStub) SetPluginID(v string) {
	o.PluginID = &v
}

// GetProvider returns the Provider field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetProvider() string {
	if o == nil || o.Provider == nil {
		var ret string
		return ret
	}
	return *o.Provider
}

// GetProviderOk returns a tuple with the Provider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetProviderOk() (*string, bool) {
	if o == nil || o.Provider == nil {
		return nil, false
	}
	return o.Provider, true
}

// HasProvider returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasProvider() bool {
	if o != nil && o.Provider != nil {
		return true
	}

	return false
}

// SetProvider gets a reference to the given string and assigns it to the Provider field.
func (o *CSIVolumeListStub) SetProvider(v string) {
	o.Provider = &v
}

// GetResourceExhausted returns the ResourceExhausted field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetResourceExhausted() time.Time {
	if o == nil || o.ResourceExhausted == nil {
		var ret time.Time
		return ret
	}
	return *o.ResourceExhausted
}

// GetResourceExhaustedOk returns a tuple with the ResourceExhausted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetResourceExhaustedOk() (*time.Time, bool) {
	if o == nil || o.ResourceExhausted == nil {
		return nil, false
	}
	return o.ResourceExhausted, true
}

// HasResourceExhausted returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasResourceExhausted() bool {
	if o != nil && o.ResourceExhausted != nil {
		return true
	}

	return false
}

// SetResourceExhausted gets a reference to the given time.Time and assigns it to the ResourceExhausted field.
func (o *CSIVolumeListStub) SetResourceExhausted(v time.Time) {
	o.ResourceExhausted = &v
}

// GetSchedulable returns the Schedulable field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetSchedulable() bool {
	if o == nil || o.Schedulable == nil {
		var ret bool
		return ret
	}
	return *o.Schedulable
}

// GetSchedulableOk returns a tuple with the Schedulable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetSchedulableOk() (*bool, bool) {
	if o == nil || o.Schedulable == nil {
		return nil, false
	}
	return o.Schedulable, true
}

// HasSchedulable returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasSchedulable() bool {
	if o != nil && o.Schedulable != nil {
		return true
	}

	return false
}

// SetSchedulable gets a reference to the given bool and assigns it to the Schedulable field.
func (o *CSIVolumeListStub) SetSchedulable(v bool) {
	o.Schedulable = &v
}

// GetTopologies returns the Topologies field value if set, zero value otherwise.
func (o *CSIVolumeListStub) GetTopologies() []CSITopology {
	if o == nil || o.Topologies == nil {
		var ret []CSITopology
		return ret
	}
	return *o.Topologies
}

// GetTopologiesOk returns a tuple with the Topologies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CSIVolumeListStub) GetTopologiesOk() (*[]CSITopology, bool) {
	if o == nil || o.Topologies == nil {
		return nil, false
	}
	return o.Topologies, true
}

// HasTopologies returns a boolean if a field has been set.
func (o *CSIVolumeListStub) HasTopologies() bool {
	if o != nil && o.Topologies != nil {
		return true
	}

	return false
}

// SetTopologies gets a reference to the given []CSITopology and assigns it to the Topologies field.
func (o *CSIVolumeListStub) SetTopologies(v []CSITopology) {
	o.Topologies = &v
}

func (o CSIVolumeListStub) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessMode != nil {
		toSerialize["AccessMode"] = o.AccessMode
	}
	if o.AttachmentMode != nil {
		toSerialize["AttachmentMode"] = o.AttachmentMode
	}
	if o.ControllerRequired != nil {
		toSerialize["ControllerRequired"] = o.ControllerRequired
	}
	if o.ControllersExpected != nil {
		toSerialize["ControllersExpected"] = o.ControllersExpected
	}
	if o.ControllersHealthy != nil {
		toSerialize["ControllersHealthy"] = o.ControllersHealthy
	}
	if o.CreateIndex != nil {
		toSerialize["CreateIndex"] = o.CreateIndex
	}
	if o.CurrentReaders != nil {
		toSerialize["CurrentReaders"] = o.CurrentReaders
	}
	if o.CurrentWriters != nil {
		toSerialize["CurrentWriters"] = o.CurrentWriters
	}
	if o.ExternalID != nil {
		toSerialize["ExternalID"] = o.ExternalID
	}
	if o.ID != nil {
		toSerialize["ID"] = o.ID
	}
	if o.ModifyIndex != nil {
		toSerialize["ModifyIndex"] = o.ModifyIndex
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Namespace != nil {
		toSerialize["Namespace"] = o.Namespace
	}
	if o.NodesExpected != nil {
		toSerialize["NodesExpected"] = o.NodesExpected
	}
	if o.NodesHealthy != nil {
		toSerialize["NodesHealthy"] = o.NodesHealthy
	}
	if o.PluginID != nil {
		toSerialize["PluginID"] = o.PluginID
	}
	if o.Provider != nil {
		toSerialize["Provider"] = o.Provider
	}
	if o.ResourceExhausted != nil {
		toSerialize["ResourceExhausted"] = o.ResourceExhausted
	}
	if o.Schedulable != nil {
		toSerialize["Schedulable"] = o.Schedulable
	}
	if o.Topologies != nil {
		toSerialize["Topologies"] = o.Topologies
	}
	return json.Marshal(toSerialize)
}

type NullableCSIVolumeListStub struct {
	value *CSIVolumeListStub
	isSet bool
}

func (v NullableCSIVolumeListStub) Get() *CSIVolumeListStub {
	return v.value
}

func (v *NullableCSIVolumeListStub) Set(val *CSIVolumeListStub) {
	v.value = val
	v.isSet = true
}

func (v NullableCSIVolumeListStub) IsSet() bool {
	return v.isSet
}

func (v *NullableCSIVolumeListStub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCSIVolumeListStub(val *CSIVolumeListStub) *NullableCSIVolumeListStub {
	return &NullableCSIVolumeListStub{value: val, isSet: true}
}

func (v NullableCSIVolumeListStub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCSIVolumeListStub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


