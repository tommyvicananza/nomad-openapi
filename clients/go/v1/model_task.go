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

// Task struct for Task
type Task struct {
	Affinities *[]Affinity `json:"Affinities,omitempty"`
	Artifacts *[]TaskArtifact `json:"Artifacts,omitempty"`
	CSIPluginConfig *TaskCSIPluginConfig `json:"CSIPluginConfig,omitempty"`
	Config *map[string]interface{} `json:"Config,omitempty"`
	Constraints *[]Constraint `json:"Constraints,omitempty"`
	DispatchPayload *DispatchPayloadConfig `json:"DispatchPayload,omitempty"`
	Driver *string `json:"Driver,omitempty"`
	Env *map[string]string `json:"Env,omitempty"`
	Identity *WorkloadIdentity `json:"Identity,omitempty"`
	KillSignal *string `json:"KillSignal,omitempty"`
	KillTimeout *int64 `json:"KillTimeout,omitempty"`
	Kind *string `json:"Kind,omitempty"`
	Leader *bool `json:"Leader,omitempty"`
	Lifecycle *TaskLifecycle `json:"Lifecycle,omitempty"`
	LogConfig *LogConfig `json:"LogConfig,omitempty"`
	Meta *map[string]string `json:"Meta,omitempty"`
	Name *string `json:"Name,omitempty"`
	Resources *Resources `json:"Resources,omitempty"`
	RestartPolicy *RestartPolicy `json:"RestartPolicy,omitempty"`
	ScalingPolicies *[]ScalingPolicy `json:"ScalingPolicies,omitempty"`
	Services *[]Service `json:"Services,omitempty"`
	ShutdownDelay *int64 `json:"ShutdownDelay,omitempty"`
	Templates *[]Template `json:"Templates,omitempty"`
	User *string `json:"User,omitempty"`
	Vault *Vault `json:"Vault,omitempty"`
	VolumeMounts *[]VolumeMount `json:"VolumeMounts,omitempty"`
}

// NewTask instantiates a new Task object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTask() *Task {
	this := Task{}
	return &this
}

// NewTaskWithDefaults instantiates a new Task object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskWithDefaults() *Task {
	this := Task{}
	return &this
}

// GetAffinities returns the Affinities field value if set, zero value otherwise.
func (o *Task) GetAffinities() []Affinity {
	if o == nil || o.Affinities == nil {
		var ret []Affinity
		return ret
	}
	return *o.Affinities
}

// GetAffinitiesOk returns a tuple with the Affinities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetAffinitiesOk() (*[]Affinity, bool) {
	if o == nil || o.Affinities == nil {
		return nil, false
	}
	return o.Affinities, true
}

// HasAffinities returns a boolean if a field has been set.
func (o *Task) HasAffinities() bool {
	if o != nil && o.Affinities != nil {
		return true
	}

	return false
}

// SetAffinities gets a reference to the given []Affinity and assigns it to the Affinities field.
func (o *Task) SetAffinities(v []Affinity) {
	o.Affinities = &v
}

// GetArtifacts returns the Artifacts field value if set, zero value otherwise.
func (o *Task) GetArtifacts() []TaskArtifact {
	if o == nil || o.Artifacts == nil {
		var ret []TaskArtifact
		return ret
	}
	return *o.Artifacts
}

// GetArtifactsOk returns a tuple with the Artifacts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetArtifactsOk() (*[]TaskArtifact, bool) {
	if o == nil || o.Artifacts == nil {
		return nil, false
	}
	return o.Artifacts, true
}

// HasArtifacts returns a boolean if a field has been set.
func (o *Task) HasArtifacts() bool {
	if o != nil && o.Artifacts != nil {
		return true
	}

	return false
}

// SetArtifacts gets a reference to the given []TaskArtifact and assigns it to the Artifacts field.
func (o *Task) SetArtifacts(v []TaskArtifact) {
	o.Artifacts = &v
}

// GetCSIPluginConfig returns the CSIPluginConfig field value if set, zero value otherwise.
func (o *Task) GetCSIPluginConfig() TaskCSIPluginConfig {
	if o == nil || o.CSIPluginConfig == nil {
		var ret TaskCSIPluginConfig
		return ret
	}
	return *o.CSIPluginConfig
}

// GetCSIPluginConfigOk returns a tuple with the CSIPluginConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetCSIPluginConfigOk() (*TaskCSIPluginConfig, bool) {
	if o == nil || o.CSIPluginConfig == nil {
		return nil, false
	}
	return o.CSIPluginConfig, true
}

// HasCSIPluginConfig returns a boolean if a field has been set.
func (o *Task) HasCSIPluginConfig() bool {
	if o != nil && o.CSIPluginConfig != nil {
		return true
	}

	return false
}

// SetCSIPluginConfig gets a reference to the given TaskCSIPluginConfig and assigns it to the CSIPluginConfig field.
func (o *Task) SetCSIPluginConfig(v TaskCSIPluginConfig) {
	o.CSIPluginConfig = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *Task) GetConfig() map[string]interface{} {
	if o == nil || o.Config == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetConfigOk() (*map[string]interface{}, bool) {
	if o == nil || o.Config == nil {
		return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *Task) HasConfig() bool {
	if o != nil && o.Config != nil {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]interface{} and assigns it to the Config field.
func (o *Task) SetConfig(v map[string]interface{}) {
	o.Config = &v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *Task) GetConstraints() []Constraint {
	if o == nil || o.Constraints == nil {
		var ret []Constraint
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetConstraintsOk() (*[]Constraint, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *Task) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given []Constraint and assigns it to the Constraints field.
func (o *Task) SetConstraints(v []Constraint) {
	o.Constraints = &v
}

// GetDispatchPayload returns the DispatchPayload field value if set, zero value otherwise.
func (o *Task) GetDispatchPayload() DispatchPayloadConfig {
	if o == nil || o.DispatchPayload == nil {
		var ret DispatchPayloadConfig
		return ret
	}
	return *o.DispatchPayload
}

// GetDispatchPayloadOk returns a tuple with the DispatchPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDispatchPayloadOk() (*DispatchPayloadConfig, bool) {
	if o == nil || o.DispatchPayload == nil {
		return nil, false
	}
	return o.DispatchPayload, true
}

// HasDispatchPayload returns a boolean if a field has been set.
func (o *Task) HasDispatchPayload() bool {
	if o != nil && o.DispatchPayload != nil {
		return true
	}

	return false
}

// SetDispatchPayload gets a reference to the given DispatchPayloadConfig and assigns it to the DispatchPayload field.
func (o *Task) SetDispatchPayload(v DispatchPayloadConfig) {
	o.DispatchPayload = &v
}

// GetDriver returns the Driver field value if set, zero value otherwise.
func (o *Task) GetDriver() string {
	if o == nil || o.Driver == nil {
		var ret string
		return ret
	}
	return *o.Driver
}

// GetDriverOk returns a tuple with the Driver field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetDriverOk() (*string, bool) {
	if o == nil || o.Driver == nil {
		return nil, false
	}
	return o.Driver, true
}

// HasDriver returns a boolean if a field has been set.
func (o *Task) HasDriver() bool {
	if o != nil && o.Driver != nil {
		return true
	}

	return false
}

// SetDriver gets a reference to the given string and assigns it to the Driver field.
func (o *Task) SetDriver(v string) {
	o.Driver = &v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *Task) GetEnv() map[string]string {
	if o == nil || o.Env == nil {
		var ret map[string]string
		return ret
	}
	return *o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetEnvOk() (*map[string]string, bool) {
	if o == nil || o.Env == nil {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *Task) HasEnv() bool {
	if o != nil && o.Env != nil {
		return true
	}

	return false
}

// SetEnv gets a reference to the given map[string]string and assigns it to the Env field.
func (o *Task) SetEnv(v map[string]string) {
	o.Env = &v
}

// GetIdentity returns the Identity field value if set, zero value otherwise.
func (o *Task) GetIdentity() WorkloadIdentity {
	if o == nil || o.Identity == nil {
		var ret WorkloadIdentity
		return ret
	}
	return *o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetIdentityOk() (*WorkloadIdentity, bool) {
	if o == nil || o.Identity == nil {
		return nil, false
	}
	return o.Identity, true
}

// HasIdentity returns a boolean if a field has been set.
func (o *Task) HasIdentity() bool {
	if o != nil && o.Identity != nil {
		return true
	}

	return false
}

// SetIdentity gets a reference to the given WorkloadIdentity and assigns it to the Identity field.
func (o *Task) SetIdentity(v WorkloadIdentity) {
	o.Identity = &v
}

// GetKillSignal returns the KillSignal field value if set, zero value otherwise.
func (o *Task) GetKillSignal() string {
	if o == nil || o.KillSignal == nil {
		var ret string
		return ret
	}
	return *o.KillSignal
}

// GetKillSignalOk returns a tuple with the KillSignal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetKillSignalOk() (*string, bool) {
	if o == nil || o.KillSignal == nil {
		return nil, false
	}
	return o.KillSignal, true
}

// HasKillSignal returns a boolean if a field has been set.
func (o *Task) HasKillSignal() bool {
	if o != nil && o.KillSignal != nil {
		return true
	}

	return false
}

// SetKillSignal gets a reference to the given string and assigns it to the KillSignal field.
func (o *Task) SetKillSignal(v string) {
	o.KillSignal = &v
}

// GetKillTimeout returns the KillTimeout field value if set, zero value otherwise.
func (o *Task) GetKillTimeout() int64 {
	if o == nil || o.KillTimeout == nil {
		var ret int64
		return ret
	}
	return *o.KillTimeout
}

// GetKillTimeoutOk returns a tuple with the KillTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetKillTimeoutOk() (*int64, bool) {
	if o == nil || o.KillTimeout == nil {
		return nil, false
	}
	return o.KillTimeout, true
}

// HasKillTimeout returns a boolean if a field has been set.
func (o *Task) HasKillTimeout() bool {
	if o != nil && o.KillTimeout != nil {
		return true
	}

	return false
}

// SetKillTimeout gets a reference to the given int64 and assigns it to the KillTimeout field.
func (o *Task) SetKillTimeout(v int64) {
	o.KillTimeout = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Task) GetKind() string {
	if o == nil || o.Kind == nil {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetKindOk() (*string, bool) {
	if o == nil || o.Kind == nil {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Task) HasKind() bool {
	if o != nil && o.Kind != nil {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Task) SetKind(v string) {
	o.Kind = &v
}

// GetLeader returns the Leader field value if set, zero value otherwise.
func (o *Task) GetLeader() bool {
	if o == nil || o.Leader == nil {
		var ret bool
		return ret
	}
	return *o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetLeaderOk() (*bool, bool) {
	if o == nil || o.Leader == nil {
		return nil, false
	}
	return o.Leader, true
}

// HasLeader returns a boolean if a field has been set.
func (o *Task) HasLeader() bool {
	if o != nil && o.Leader != nil {
		return true
	}

	return false
}

// SetLeader gets a reference to the given bool and assigns it to the Leader field.
func (o *Task) SetLeader(v bool) {
	o.Leader = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *Task) GetLifecycle() TaskLifecycle {
	if o == nil || o.Lifecycle == nil {
		var ret TaskLifecycle
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetLifecycleOk() (*TaskLifecycle, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *Task) HasLifecycle() bool {
	if o != nil && o.Lifecycle != nil {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given TaskLifecycle and assigns it to the Lifecycle field.
func (o *Task) SetLifecycle(v TaskLifecycle) {
	o.Lifecycle = &v
}

// GetLogConfig returns the LogConfig field value if set, zero value otherwise.
func (o *Task) GetLogConfig() LogConfig {
	if o == nil || o.LogConfig == nil {
		var ret LogConfig
		return ret
	}
	return *o.LogConfig
}

// GetLogConfigOk returns a tuple with the LogConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetLogConfigOk() (*LogConfig, bool) {
	if o == nil || o.LogConfig == nil {
		return nil, false
	}
	return o.LogConfig, true
}

// HasLogConfig returns a boolean if a field has been set.
func (o *Task) HasLogConfig() bool {
	if o != nil && o.LogConfig != nil {
		return true
	}

	return false
}

// SetLogConfig gets a reference to the given LogConfig and assigns it to the LogConfig field.
func (o *Task) SetLogConfig(v LogConfig) {
	o.LogConfig = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *Task) GetMeta() map[string]string {
	if o == nil || o.Meta == nil {
		var ret map[string]string
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetMetaOk() (*map[string]string, bool) {
	if o == nil || o.Meta == nil {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *Task) HasMeta() bool {
	if o != nil && o.Meta != nil {
		return true
	}

	return false
}

// SetMeta gets a reference to the given map[string]string and assigns it to the Meta field.
func (o *Task) SetMeta(v map[string]string) {
	o.Meta = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Task) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Task) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Task) SetName(v string) {
	o.Name = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *Task) GetResources() Resources {
	if o == nil || o.Resources == nil {
		var ret Resources
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetResourcesOk() (*Resources, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *Task) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given Resources and assigns it to the Resources field.
func (o *Task) SetResources(v Resources) {
	o.Resources = &v
}

// GetRestartPolicy returns the RestartPolicy field value if set, zero value otherwise.
func (o *Task) GetRestartPolicy() RestartPolicy {
	if o == nil || o.RestartPolicy == nil {
		var ret RestartPolicy
		return ret
	}
	return *o.RestartPolicy
}

// GetRestartPolicyOk returns a tuple with the RestartPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetRestartPolicyOk() (*RestartPolicy, bool) {
	if o == nil || o.RestartPolicy == nil {
		return nil, false
	}
	return o.RestartPolicy, true
}

// HasRestartPolicy returns a boolean if a field has been set.
func (o *Task) HasRestartPolicy() bool {
	if o != nil && o.RestartPolicy != nil {
		return true
	}

	return false
}

// SetRestartPolicy gets a reference to the given RestartPolicy and assigns it to the RestartPolicy field.
func (o *Task) SetRestartPolicy(v RestartPolicy) {
	o.RestartPolicy = &v
}

// GetScalingPolicies returns the ScalingPolicies field value if set, zero value otherwise.
func (o *Task) GetScalingPolicies() []ScalingPolicy {
	if o == nil || o.ScalingPolicies == nil {
		var ret []ScalingPolicy
		return ret
	}
	return *o.ScalingPolicies
}

// GetScalingPoliciesOk returns a tuple with the ScalingPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetScalingPoliciesOk() (*[]ScalingPolicy, bool) {
	if o == nil || o.ScalingPolicies == nil {
		return nil, false
	}
	return o.ScalingPolicies, true
}

// HasScalingPolicies returns a boolean if a field has been set.
func (o *Task) HasScalingPolicies() bool {
	if o != nil && o.ScalingPolicies != nil {
		return true
	}

	return false
}

// SetScalingPolicies gets a reference to the given []ScalingPolicy and assigns it to the ScalingPolicies field.
func (o *Task) SetScalingPolicies(v []ScalingPolicy) {
	o.ScalingPolicies = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *Task) GetServices() []Service {
	if o == nil || o.Services == nil {
		var ret []Service
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetServicesOk() (*[]Service, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *Task) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []Service and assigns it to the Services field.
func (o *Task) SetServices(v []Service) {
	o.Services = &v
}

// GetShutdownDelay returns the ShutdownDelay field value if set, zero value otherwise.
func (o *Task) GetShutdownDelay() int64 {
	if o == nil || o.ShutdownDelay == nil {
		var ret int64
		return ret
	}
	return *o.ShutdownDelay
}

// GetShutdownDelayOk returns a tuple with the ShutdownDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetShutdownDelayOk() (*int64, bool) {
	if o == nil || o.ShutdownDelay == nil {
		return nil, false
	}
	return o.ShutdownDelay, true
}

// HasShutdownDelay returns a boolean if a field has been set.
func (o *Task) HasShutdownDelay() bool {
	if o != nil && o.ShutdownDelay != nil {
		return true
	}

	return false
}

// SetShutdownDelay gets a reference to the given int64 and assigns it to the ShutdownDelay field.
func (o *Task) SetShutdownDelay(v int64) {
	o.ShutdownDelay = &v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *Task) GetTemplates() []Template {
	if o == nil || o.Templates == nil {
		var ret []Template
		return ret
	}
	return *o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetTemplatesOk() (*[]Template, bool) {
	if o == nil || o.Templates == nil {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *Task) HasTemplates() bool {
	if o != nil && o.Templates != nil {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []Template and assigns it to the Templates field.
func (o *Task) SetTemplates(v []Template) {
	o.Templates = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *Task) GetUser() string {
	if o == nil || o.User == nil {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetUserOk() (*string, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *Task) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *Task) SetUser(v string) {
	o.User = &v
}

// GetVault returns the Vault field value if set, zero value otherwise.
func (o *Task) GetVault() Vault {
	if o == nil || o.Vault == nil {
		var ret Vault
		return ret
	}
	return *o.Vault
}

// GetVaultOk returns a tuple with the Vault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetVaultOk() (*Vault, bool) {
	if o == nil || o.Vault == nil {
		return nil, false
	}
	return o.Vault, true
}

// HasVault returns a boolean if a field has been set.
func (o *Task) HasVault() bool {
	if o != nil && o.Vault != nil {
		return true
	}

	return false
}

// SetVault gets a reference to the given Vault and assigns it to the Vault field.
func (o *Task) SetVault(v Vault) {
	o.Vault = &v
}

// GetVolumeMounts returns the VolumeMounts field value if set, zero value otherwise.
func (o *Task) GetVolumeMounts() []VolumeMount {
	if o == nil || o.VolumeMounts == nil {
		var ret []VolumeMount
		return ret
	}
	return *o.VolumeMounts
}

// GetVolumeMountsOk returns a tuple with the VolumeMounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Task) GetVolumeMountsOk() (*[]VolumeMount, bool) {
	if o == nil || o.VolumeMounts == nil {
		return nil, false
	}
	return o.VolumeMounts, true
}

// HasVolumeMounts returns a boolean if a field has been set.
func (o *Task) HasVolumeMounts() bool {
	if o != nil && o.VolumeMounts != nil {
		return true
	}

	return false
}

// SetVolumeMounts gets a reference to the given []VolumeMount and assigns it to the VolumeMounts field.
func (o *Task) SetVolumeMounts(v []VolumeMount) {
	o.VolumeMounts = &v
}

func (o Task) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Affinities != nil {
		toSerialize["Affinities"] = o.Affinities
	}
	if o.Artifacts != nil {
		toSerialize["Artifacts"] = o.Artifacts
	}
	if o.CSIPluginConfig != nil {
		toSerialize["CSIPluginConfig"] = o.CSIPluginConfig
	}
	if o.Config != nil {
		toSerialize["Config"] = o.Config
	}
	if o.Constraints != nil {
		toSerialize["Constraints"] = o.Constraints
	}
	if o.DispatchPayload != nil {
		toSerialize["DispatchPayload"] = o.DispatchPayload
	}
	if o.Driver != nil {
		toSerialize["Driver"] = o.Driver
	}
	if o.Env != nil {
		toSerialize["Env"] = o.Env
	}
	if o.Identity != nil {
		toSerialize["Identity"] = o.Identity
	}
	if o.KillSignal != nil {
		toSerialize["KillSignal"] = o.KillSignal
	}
	if o.KillTimeout != nil {
		toSerialize["KillTimeout"] = o.KillTimeout
	}
	if o.Kind != nil {
		toSerialize["Kind"] = o.Kind
	}
	if o.Leader != nil {
		toSerialize["Leader"] = o.Leader
	}
	if o.Lifecycle != nil {
		toSerialize["Lifecycle"] = o.Lifecycle
	}
	if o.LogConfig != nil {
		toSerialize["LogConfig"] = o.LogConfig
	}
	if o.Meta != nil {
		toSerialize["Meta"] = o.Meta
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Resources != nil {
		toSerialize["Resources"] = o.Resources
	}
	if o.RestartPolicy != nil {
		toSerialize["RestartPolicy"] = o.RestartPolicy
	}
	if o.ScalingPolicies != nil {
		toSerialize["ScalingPolicies"] = o.ScalingPolicies
	}
	if o.Services != nil {
		toSerialize["Services"] = o.Services
	}
	if o.ShutdownDelay != nil {
		toSerialize["ShutdownDelay"] = o.ShutdownDelay
	}
	if o.Templates != nil {
		toSerialize["Templates"] = o.Templates
	}
	if o.User != nil {
		toSerialize["User"] = o.User
	}
	if o.Vault != nil {
		toSerialize["Vault"] = o.Vault
	}
	if o.VolumeMounts != nil {
		toSerialize["VolumeMounts"] = o.VolumeMounts
	}
	return json.Marshal(toSerialize)
}

type NullableTask struct {
	value *Task
	isSet bool
}

func (v NullableTask) Get() *Task {
	return v.value
}

func (v *NullableTask) Set(val *Task) {
	v.value = val
	v.isSet = true
}

func (v NullableTask) IsSet() bool {
	return v.isSet
}

func (v *NullableTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTask(val *Task) *NullableTask {
	return &NullableTask{value: val, isSet: true}
}

func (v NullableTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


