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

// ServiceCheck struct for ServiceCheck
type ServiceCheck struct {
	AddressMode *string `json:"AddressMode,omitempty"`
	Advertise *string `json:"Advertise,omitempty"`
	Args *[]string `json:"Args,omitempty"`
	Body *string `json:"Body,omitempty"`
	CheckRestart *CheckRestart `json:"CheckRestart,omitempty"`
	Command *string `json:"Command,omitempty"`
	Expose *bool `json:"Expose,omitempty"`
	FailuresBeforeCritical *int32 `json:"FailuresBeforeCritical,omitempty"`
	GRPCService *string `json:"GRPCService,omitempty"`
	GRPCUseTLS *bool `json:"GRPCUseTLS,omitempty"`
	Header *map[string][]string `json:"Header,omitempty"`
	InitialStatus *string `json:"InitialStatus,omitempty"`
	Interval *int64 `json:"Interval,omitempty"`
	Method *string `json:"Method,omitempty"`
	Name *string `json:"Name,omitempty"`
	OnUpdate *string `json:"OnUpdate,omitempty"`
	Path *string `json:"Path,omitempty"`
	PortLabel *string `json:"PortLabel,omitempty"`
	Protocol *string `json:"Protocol,omitempty"`
	SuccessBeforePassing *int32 `json:"SuccessBeforePassing,omitempty"`
	TLSSkipVerify *bool `json:"TLSSkipVerify,omitempty"`
	TaskName *string `json:"TaskName,omitempty"`
	Timeout *int64 `json:"Timeout,omitempty"`
	Type *string `json:"Type,omitempty"`
}

// NewServiceCheck instantiates a new ServiceCheck object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceCheck() *ServiceCheck {
	this := ServiceCheck{}
	return &this
}

// NewServiceCheckWithDefaults instantiates a new ServiceCheck object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceCheckWithDefaults() *ServiceCheck {
	this := ServiceCheck{}
	return &this
}

// GetAddressMode returns the AddressMode field value if set, zero value otherwise.
func (o *ServiceCheck) GetAddressMode() string {
	if o == nil || o.AddressMode == nil {
		var ret string
		return ret
	}
	return *o.AddressMode
}

// GetAddressModeOk returns a tuple with the AddressMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetAddressModeOk() (*string, bool) {
	if o == nil || o.AddressMode == nil {
		return nil, false
	}
	return o.AddressMode, true
}

// HasAddressMode returns a boolean if a field has been set.
func (o *ServiceCheck) HasAddressMode() bool {
	if o != nil && o.AddressMode != nil {
		return true
	}

	return false
}

// SetAddressMode gets a reference to the given string and assigns it to the AddressMode field.
func (o *ServiceCheck) SetAddressMode(v string) {
	o.AddressMode = &v
}

// GetAdvertise returns the Advertise field value if set, zero value otherwise.
func (o *ServiceCheck) GetAdvertise() string {
	if o == nil || o.Advertise == nil {
		var ret string
		return ret
	}
	return *o.Advertise
}

// GetAdvertiseOk returns a tuple with the Advertise field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetAdvertiseOk() (*string, bool) {
	if o == nil || o.Advertise == nil {
		return nil, false
	}
	return o.Advertise, true
}

// HasAdvertise returns a boolean if a field has been set.
func (o *ServiceCheck) HasAdvertise() bool {
	if o != nil && o.Advertise != nil {
		return true
	}

	return false
}

// SetAdvertise gets a reference to the given string and assigns it to the Advertise field.
func (o *ServiceCheck) SetAdvertise(v string) {
	o.Advertise = &v
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *ServiceCheck) GetArgs() []string {
	if o == nil || o.Args == nil {
		var ret []string
		return ret
	}
	return *o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetArgsOk() (*[]string, bool) {
	if o == nil || o.Args == nil {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *ServiceCheck) HasArgs() bool {
	if o != nil && o.Args != nil {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *ServiceCheck) SetArgs(v []string) {
	o.Args = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *ServiceCheck) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *ServiceCheck) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *ServiceCheck) SetBody(v string) {
	o.Body = &v
}

// GetCheckRestart returns the CheckRestart field value if set, zero value otherwise.
func (o *ServiceCheck) GetCheckRestart() CheckRestart {
	if o == nil || o.CheckRestart == nil {
		var ret CheckRestart
		return ret
	}
	return *o.CheckRestart
}

// GetCheckRestartOk returns a tuple with the CheckRestart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetCheckRestartOk() (*CheckRestart, bool) {
	if o == nil || o.CheckRestart == nil {
		return nil, false
	}
	return o.CheckRestart, true
}

// HasCheckRestart returns a boolean if a field has been set.
func (o *ServiceCheck) HasCheckRestart() bool {
	if o != nil && o.CheckRestart != nil {
		return true
	}

	return false
}

// SetCheckRestart gets a reference to the given CheckRestart and assigns it to the CheckRestart field.
func (o *ServiceCheck) SetCheckRestart(v CheckRestart) {
	o.CheckRestart = &v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *ServiceCheck) GetCommand() string {
	if o == nil || o.Command == nil {
		var ret string
		return ret
	}
	return *o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetCommandOk() (*string, bool) {
	if o == nil || o.Command == nil {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *ServiceCheck) HasCommand() bool {
	if o != nil && o.Command != nil {
		return true
	}

	return false
}

// SetCommand gets a reference to the given string and assigns it to the Command field.
func (o *ServiceCheck) SetCommand(v string) {
	o.Command = &v
}

// GetExpose returns the Expose field value if set, zero value otherwise.
func (o *ServiceCheck) GetExpose() bool {
	if o == nil || o.Expose == nil {
		var ret bool
		return ret
	}
	return *o.Expose
}

// GetExposeOk returns a tuple with the Expose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetExposeOk() (*bool, bool) {
	if o == nil || o.Expose == nil {
		return nil, false
	}
	return o.Expose, true
}

// HasExpose returns a boolean if a field has been set.
func (o *ServiceCheck) HasExpose() bool {
	if o != nil && o.Expose != nil {
		return true
	}

	return false
}

// SetExpose gets a reference to the given bool and assigns it to the Expose field.
func (o *ServiceCheck) SetExpose(v bool) {
	o.Expose = &v
}

// GetFailuresBeforeCritical returns the FailuresBeforeCritical field value if set, zero value otherwise.
func (o *ServiceCheck) GetFailuresBeforeCritical() int32 {
	if o == nil || o.FailuresBeforeCritical == nil {
		var ret int32
		return ret
	}
	return *o.FailuresBeforeCritical
}

// GetFailuresBeforeCriticalOk returns a tuple with the FailuresBeforeCritical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetFailuresBeforeCriticalOk() (*int32, bool) {
	if o == nil || o.FailuresBeforeCritical == nil {
		return nil, false
	}
	return o.FailuresBeforeCritical, true
}

// HasFailuresBeforeCritical returns a boolean if a field has been set.
func (o *ServiceCheck) HasFailuresBeforeCritical() bool {
	if o != nil && o.FailuresBeforeCritical != nil {
		return true
	}

	return false
}

// SetFailuresBeforeCritical gets a reference to the given int32 and assigns it to the FailuresBeforeCritical field.
func (o *ServiceCheck) SetFailuresBeforeCritical(v int32) {
	o.FailuresBeforeCritical = &v
}

// GetGRPCService returns the GRPCService field value if set, zero value otherwise.
func (o *ServiceCheck) GetGRPCService() string {
	if o == nil || o.GRPCService == nil {
		var ret string
		return ret
	}
	return *o.GRPCService
}

// GetGRPCServiceOk returns a tuple with the GRPCService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetGRPCServiceOk() (*string, bool) {
	if o == nil || o.GRPCService == nil {
		return nil, false
	}
	return o.GRPCService, true
}

// HasGRPCService returns a boolean if a field has been set.
func (o *ServiceCheck) HasGRPCService() bool {
	if o != nil && o.GRPCService != nil {
		return true
	}

	return false
}

// SetGRPCService gets a reference to the given string and assigns it to the GRPCService field.
func (o *ServiceCheck) SetGRPCService(v string) {
	o.GRPCService = &v
}

// GetGRPCUseTLS returns the GRPCUseTLS field value if set, zero value otherwise.
func (o *ServiceCheck) GetGRPCUseTLS() bool {
	if o == nil || o.GRPCUseTLS == nil {
		var ret bool
		return ret
	}
	return *o.GRPCUseTLS
}

// GetGRPCUseTLSOk returns a tuple with the GRPCUseTLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetGRPCUseTLSOk() (*bool, bool) {
	if o == nil || o.GRPCUseTLS == nil {
		return nil, false
	}
	return o.GRPCUseTLS, true
}

// HasGRPCUseTLS returns a boolean if a field has been set.
func (o *ServiceCheck) HasGRPCUseTLS() bool {
	if o != nil && o.GRPCUseTLS != nil {
		return true
	}

	return false
}

// SetGRPCUseTLS gets a reference to the given bool and assigns it to the GRPCUseTLS field.
func (o *ServiceCheck) SetGRPCUseTLS(v bool) {
	o.GRPCUseTLS = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *ServiceCheck) GetHeader() map[string][]string {
	if o == nil || o.Header == nil {
		var ret map[string][]string
		return ret
	}
	return *o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetHeaderOk() (*map[string][]string, bool) {
	if o == nil || o.Header == nil {
		return nil, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *ServiceCheck) HasHeader() bool {
	if o != nil && o.Header != nil {
		return true
	}

	return false
}

// SetHeader gets a reference to the given map[string][]string and assigns it to the Header field.
func (o *ServiceCheck) SetHeader(v map[string][]string) {
	o.Header = &v
}

// GetInitialStatus returns the InitialStatus field value if set, zero value otherwise.
func (o *ServiceCheck) GetInitialStatus() string {
	if o == nil || o.InitialStatus == nil {
		var ret string
		return ret
	}
	return *o.InitialStatus
}

// GetInitialStatusOk returns a tuple with the InitialStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetInitialStatusOk() (*string, bool) {
	if o == nil || o.InitialStatus == nil {
		return nil, false
	}
	return o.InitialStatus, true
}

// HasInitialStatus returns a boolean if a field has been set.
func (o *ServiceCheck) HasInitialStatus() bool {
	if o != nil && o.InitialStatus != nil {
		return true
	}

	return false
}

// SetInitialStatus gets a reference to the given string and assigns it to the InitialStatus field.
func (o *ServiceCheck) SetInitialStatus(v string) {
	o.InitialStatus = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *ServiceCheck) GetInterval() int64 {
	if o == nil || o.Interval == nil {
		var ret int64
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetIntervalOk() (*int64, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *ServiceCheck) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int64 and assigns it to the Interval field.
func (o *ServiceCheck) SetInterval(v int64) {
	o.Interval = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ServiceCheck) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ServiceCheck) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ServiceCheck) SetMethod(v string) {
	o.Method = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceCheck) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceCheck) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceCheck) SetName(v string) {
	o.Name = &v
}

// GetOnUpdate returns the OnUpdate field value if set, zero value otherwise.
func (o *ServiceCheck) GetOnUpdate() string {
	if o == nil || o.OnUpdate == nil {
		var ret string
		return ret
	}
	return *o.OnUpdate
}

// GetOnUpdateOk returns a tuple with the OnUpdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetOnUpdateOk() (*string, bool) {
	if o == nil || o.OnUpdate == nil {
		return nil, false
	}
	return o.OnUpdate, true
}

// HasOnUpdate returns a boolean if a field has been set.
func (o *ServiceCheck) HasOnUpdate() bool {
	if o != nil && o.OnUpdate != nil {
		return true
	}

	return false
}

// SetOnUpdate gets a reference to the given string and assigns it to the OnUpdate field.
func (o *ServiceCheck) SetOnUpdate(v string) {
	o.OnUpdate = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ServiceCheck) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ServiceCheck) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ServiceCheck) SetPath(v string) {
	o.Path = &v
}

// GetPortLabel returns the PortLabel field value if set, zero value otherwise.
func (o *ServiceCheck) GetPortLabel() string {
	if o == nil || o.PortLabel == nil {
		var ret string
		return ret
	}
	return *o.PortLabel
}

// GetPortLabelOk returns a tuple with the PortLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetPortLabelOk() (*string, bool) {
	if o == nil || o.PortLabel == nil {
		return nil, false
	}
	return o.PortLabel, true
}

// HasPortLabel returns a boolean if a field has been set.
func (o *ServiceCheck) HasPortLabel() bool {
	if o != nil && o.PortLabel != nil {
		return true
	}

	return false
}

// SetPortLabel gets a reference to the given string and assigns it to the PortLabel field.
func (o *ServiceCheck) SetPortLabel(v string) {
	o.PortLabel = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ServiceCheck) GetProtocol() string {
	if o == nil || o.Protocol == nil {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetProtocolOk() (*string, bool) {
	if o == nil || o.Protocol == nil {
		return nil, false
	}
	return o.Protocol, true
}

// HasProtocol returns a boolean if a field has been set.
func (o *ServiceCheck) HasProtocol() bool {
	if o != nil && o.Protocol != nil {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ServiceCheck) SetProtocol(v string) {
	o.Protocol = &v
}

// GetSuccessBeforePassing returns the SuccessBeforePassing field value if set, zero value otherwise.
func (o *ServiceCheck) GetSuccessBeforePassing() int32 {
	if o == nil || o.SuccessBeforePassing == nil {
		var ret int32
		return ret
	}
	return *o.SuccessBeforePassing
}

// GetSuccessBeforePassingOk returns a tuple with the SuccessBeforePassing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetSuccessBeforePassingOk() (*int32, bool) {
	if o == nil || o.SuccessBeforePassing == nil {
		return nil, false
	}
	return o.SuccessBeforePassing, true
}

// HasSuccessBeforePassing returns a boolean if a field has been set.
func (o *ServiceCheck) HasSuccessBeforePassing() bool {
	if o != nil && o.SuccessBeforePassing != nil {
		return true
	}

	return false
}

// SetSuccessBeforePassing gets a reference to the given int32 and assigns it to the SuccessBeforePassing field.
func (o *ServiceCheck) SetSuccessBeforePassing(v int32) {
	o.SuccessBeforePassing = &v
}

// GetTLSSkipVerify returns the TLSSkipVerify field value if set, zero value otherwise.
func (o *ServiceCheck) GetTLSSkipVerify() bool {
	if o == nil || o.TLSSkipVerify == nil {
		var ret bool
		return ret
	}
	return *o.TLSSkipVerify
}

// GetTLSSkipVerifyOk returns a tuple with the TLSSkipVerify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetTLSSkipVerifyOk() (*bool, bool) {
	if o == nil || o.TLSSkipVerify == nil {
		return nil, false
	}
	return o.TLSSkipVerify, true
}

// HasTLSSkipVerify returns a boolean if a field has been set.
func (o *ServiceCheck) HasTLSSkipVerify() bool {
	if o != nil && o.TLSSkipVerify != nil {
		return true
	}

	return false
}

// SetTLSSkipVerify gets a reference to the given bool and assigns it to the TLSSkipVerify field.
func (o *ServiceCheck) SetTLSSkipVerify(v bool) {
	o.TLSSkipVerify = &v
}

// GetTaskName returns the TaskName field value if set, zero value otherwise.
func (o *ServiceCheck) GetTaskName() string {
	if o == nil || o.TaskName == nil {
		var ret string
		return ret
	}
	return *o.TaskName
}

// GetTaskNameOk returns a tuple with the TaskName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetTaskNameOk() (*string, bool) {
	if o == nil || o.TaskName == nil {
		return nil, false
	}
	return o.TaskName, true
}

// HasTaskName returns a boolean if a field has been set.
func (o *ServiceCheck) HasTaskName() bool {
	if o != nil && o.TaskName != nil {
		return true
	}

	return false
}

// SetTaskName gets a reference to the given string and assigns it to the TaskName field.
func (o *ServiceCheck) SetTaskName(v string) {
	o.TaskName = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ServiceCheck) GetTimeout() int64 {
	if o == nil || o.Timeout == nil {
		var ret int64
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetTimeoutOk() (*int64, bool) {
	if o == nil || o.Timeout == nil {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ServiceCheck) HasTimeout() bool {
	if o != nil && o.Timeout != nil {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int64 and assigns it to the Timeout field.
func (o *ServiceCheck) SetTimeout(v int64) {
	o.Timeout = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceCheck) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCheck) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceCheck) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceCheck) SetType(v string) {
	o.Type = &v
}

func (o ServiceCheck) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressMode != nil {
		toSerialize["AddressMode"] = o.AddressMode
	}
	if o.Advertise != nil {
		toSerialize["Advertise"] = o.Advertise
	}
	if o.Args != nil {
		toSerialize["Args"] = o.Args
	}
	if o.Body != nil {
		toSerialize["Body"] = o.Body
	}
	if o.CheckRestart != nil {
		toSerialize["CheckRestart"] = o.CheckRestart
	}
	if o.Command != nil {
		toSerialize["Command"] = o.Command
	}
	if o.Expose != nil {
		toSerialize["Expose"] = o.Expose
	}
	if o.FailuresBeforeCritical != nil {
		toSerialize["FailuresBeforeCritical"] = o.FailuresBeforeCritical
	}
	if o.GRPCService != nil {
		toSerialize["GRPCService"] = o.GRPCService
	}
	if o.GRPCUseTLS != nil {
		toSerialize["GRPCUseTLS"] = o.GRPCUseTLS
	}
	if o.Header != nil {
		toSerialize["Header"] = o.Header
	}
	if o.InitialStatus != nil {
		toSerialize["InitialStatus"] = o.InitialStatus
	}
	if o.Interval != nil {
		toSerialize["Interval"] = o.Interval
	}
	if o.Method != nil {
		toSerialize["Method"] = o.Method
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OnUpdate != nil {
		toSerialize["OnUpdate"] = o.OnUpdate
	}
	if o.Path != nil {
		toSerialize["Path"] = o.Path
	}
	if o.PortLabel != nil {
		toSerialize["PortLabel"] = o.PortLabel
	}
	if o.Protocol != nil {
		toSerialize["Protocol"] = o.Protocol
	}
	if o.SuccessBeforePassing != nil {
		toSerialize["SuccessBeforePassing"] = o.SuccessBeforePassing
	}
	if o.TLSSkipVerify != nil {
		toSerialize["TLSSkipVerify"] = o.TLSSkipVerify
	}
	if o.TaskName != nil {
		toSerialize["TaskName"] = o.TaskName
	}
	if o.Timeout != nil {
		toSerialize["Timeout"] = o.Timeout
	}
	if o.Type != nil {
		toSerialize["Type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableServiceCheck struct {
	value *ServiceCheck
	isSet bool
}

func (v NullableServiceCheck) Get() *ServiceCheck {
	return v.value
}

func (v *NullableServiceCheck) Set(val *ServiceCheck) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceCheck) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceCheck) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceCheck(val *ServiceCheck) *NullableServiceCheck {
	return &NullableServiceCheck{value: val, isSet: true}
}

func (v NullableServiceCheck) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceCheck) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


