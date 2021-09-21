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

// Template struct for Template
type Template struct {
	ChangeMode *string `json:"ChangeMode,omitempty"`
	ChangeSignal *string `json:"ChangeSignal,omitempty"`
	DestPath *string `json:"DestPath,omitempty"`
	EmbeddedTmpl *string `json:"EmbeddedTmpl,omitempty"`
	Envvars *bool `json:"Envvars,omitempty"`
	LeftDelim *string `json:"LeftDelim,omitempty"`
	Perms *string `json:"Perms,omitempty"`
	RightDelim *string `json:"RightDelim,omitempty"`
	SourcePath *string `json:"SourcePath,omitempty"`
	Splay *int64 `json:"Splay,omitempty"`
	VaultGrace *int64 `json:"VaultGrace,omitempty"`
}

// NewTemplate instantiates a new Template object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplate() *Template {
	this := Template{}
	return &this
}

// NewTemplateWithDefaults instantiates a new Template object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateWithDefaults() *Template {
	this := Template{}
	return &this
}

// GetChangeMode returns the ChangeMode field value if set, zero value otherwise.
func (o *Template) GetChangeMode() string {
	if o == nil || o.ChangeMode == nil {
		var ret string
		return ret
	}
	return *o.ChangeMode
}

// GetChangeModeOk returns a tuple with the ChangeMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetChangeModeOk() (*string, bool) {
	if o == nil || o.ChangeMode == nil {
		return nil, false
	}
	return o.ChangeMode, true
}

// HasChangeMode returns a boolean if a field has been set.
func (o *Template) HasChangeMode() bool {
	if o != nil && o.ChangeMode != nil {
		return true
	}

	return false
}

// SetChangeMode gets a reference to the given string and assigns it to the ChangeMode field.
func (o *Template) SetChangeMode(v string) {
	o.ChangeMode = &v
}

// GetChangeSignal returns the ChangeSignal field value if set, zero value otherwise.
func (o *Template) GetChangeSignal() string {
	if o == nil || o.ChangeSignal == nil {
		var ret string
		return ret
	}
	return *o.ChangeSignal
}

// GetChangeSignalOk returns a tuple with the ChangeSignal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetChangeSignalOk() (*string, bool) {
	if o == nil || o.ChangeSignal == nil {
		return nil, false
	}
	return o.ChangeSignal, true
}

// HasChangeSignal returns a boolean if a field has been set.
func (o *Template) HasChangeSignal() bool {
	if o != nil && o.ChangeSignal != nil {
		return true
	}

	return false
}

// SetChangeSignal gets a reference to the given string and assigns it to the ChangeSignal field.
func (o *Template) SetChangeSignal(v string) {
	o.ChangeSignal = &v
}

// GetDestPath returns the DestPath field value if set, zero value otherwise.
func (o *Template) GetDestPath() string {
	if o == nil || o.DestPath == nil {
		var ret string
		return ret
	}
	return *o.DestPath
}

// GetDestPathOk returns a tuple with the DestPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetDestPathOk() (*string, bool) {
	if o == nil || o.DestPath == nil {
		return nil, false
	}
	return o.DestPath, true
}

// HasDestPath returns a boolean if a field has been set.
func (o *Template) HasDestPath() bool {
	if o != nil && o.DestPath != nil {
		return true
	}

	return false
}

// SetDestPath gets a reference to the given string and assigns it to the DestPath field.
func (o *Template) SetDestPath(v string) {
	o.DestPath = &v
}

// GetEmbeddedTmpl returns the EmbeddedTmpl field value if set, zero value otherwise.
func (o *Template) GetEmbeddedTmpl() string {
	if o == nil || o.EmbeddedTmpl == nil {
		var ret string
		return ret
	}
	return *o.EmbeddedTmpl
}

// GetEmbeddedTmplOk returns a tuple with the EmbeddedTmpl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetEmbeddedTmplOk() (*string, bool) {
	if o == nil || o.EmbeddedTmpl == nil {
		return nil, false
	}
	return o.EmbeddedTmpl, true
}

// HasEmbeddedTmpl returns a boolean if a field has been set.
func (o *Template) HasEmbeddedTmpl() bool {
	if o != nil && o.EmbeddedTmpl != nil {
		return true
	}

	return false
}

// SetEmbeddedTmpl gets a reference to the given string and assigns it to the EmbeddedTmpl field.
func (o *Template) SetEmbeddedTmpl(v string) {
	o.EmbeddedTmpl = &v
}

// GetEnvvars returns the Envvars field value if set, zero value otherwise.
func (o *Template) GetEnvvars() bool {
	if o == nil || o.Envvars == nil {
		var ret bool
		return ret
	}
	return *o.Envvars
}

// GetEnvvarsOk returns a tuple with the Envvars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetEnvvarsOk() (*bool, bool) {
	if o == nil || o.Envvars == nil {
		return nil, false
	}
	return o.Envvars, true
}

// HasEnvvars returns a boolean if a field has been set.
func (o *Template) HasEnvvars() bool {
	if o != nil && o.Envvars != nil {
		return true
	}

	return false
}

// SetEnvvars gets a reference to the given bool and assigns it to the Envvars field.
func (o *Template) SetEnvvars(v bool) {
	o.Envvars = &v
}

// GetLeftDelim returns the LeftDelim field value if set, zero value otherwise.
func (o *Template) GetLeftDelim() string {
	if o == nil || o.LeftDelim == nil {
		var ret string
		return ret
	}
	return *o.LeftDelim
}

// GetLeftDelimOk returns a tuple with the LeftDelim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetLeftDelimOk() (*string, bool) {
	if o == nil || o.LeftDelim == nil {
		return nil, false
	}
	return o.LeftDelim, true
}

// HasLeftDelim returns a boolean if a field has been set.
func (o *Template) HasLeftDelim() bool {
	if o != nil && o.LeftDelim != nil {
		return true
	}

	return false
}

// SetLeftDelim gets a reference to the given string and assigns it to the LeftDelim field.
func (o *Template) SetLeftDelim(v string) {
	o.LeftDelim = &v
}

// GetPerms returns the Perms field value if set, zero value otherwise.
func (o *Template) GetPerms() string {
	if o == nil || o.Perms == nil {
		var ret string
		return ret
	}
	return *o.Perms
}

// GetPermsOk returns a tuple with the Perms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetPermsOk() (*string, bool) {
	if o == nil || o.Perms == nil {
		return nil, false
	}
	return o.Perms, true
}

// HasPerms returns a boolean if a field has been set.
func (o *Template) HasPerms() bool {
	if o != nil && o.Perms != nil {
		return true
	}

	return false
}

// SetPerms gets a reference to the given string and assigns it to the Perms field.
func (o *Template) SetPerms(v string) {
	o.Perms = &v
}

// GetRightDelim returns the RightDelim field value if set, zero value otherwise.
func (o *Template) GetRightDelim() string {
	if o == nil || o.RightDelim == nil {
		var ret string
		return ret
	}
	return *o.RightDelim
}

// GetRightDelimOk returns a tuple with the RightDelim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetRightDelimOk() (*string, bool) {
	if o == nil || o.RightDelim == nil {
		return nil, false
	}
	return o.RightDelim, true
}

// HasRightDelim returns a boolean if a field has been set.
func (o *Template) HasRightDelim() bool {
	if o != nil && o.RightDelim != nil {
		return true
	}

	return false
}

// SetRightDelim gets a reference to the given string and assigns it to the RightDelim field.
func (o *Template) SetRightDelim(v string) {
	o.RightDelim = &v
}

// GetSourcePath returns the SourcePath field value if set, zero value otherwise.
func (o *Template) GetSourcePath() string {
	if o == nil || o.SourcePath == nil {
		var ret string
		return ret
	}
	return *o.SourcePath
}

// GetSourcePathOk returns a tuple with the SourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetSourcePathOk() (*string, bool) {
	if o == nil || o.SourcePath == nil {
		return nil, false
	}
	return o.SourcePath, true
}

// HasSourcePath returns a boolean if a field has been set.
func (o *Template) HasSourcePath() bool {
	if o != nil && o.SourcePath != nil {
		return true
	}

	return false
}

// SetSourcePath gets a reference to the given string and assigns it to the SourcePath field.
func (o *Template) SetSourcePath(v string) {
	o.SourcePath = &v
}

// GetSplay returns the Splay field value if set, zero value otherwise.
func (o *Template) GetSplay() int64 {
	if o == nil || o.Splay == nil {
		var ret int64
		return ret
	}
	return *o.Splay
}

// GetSplayOk returns a tuple with the Splay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetSplayOk() (*int64, bool) {
	if o == nil || o.Splay == nil {
		return nil, false
	}
	return o.Splay, true
}

// HasSplay returns a boolean if a field has been set.
func (o *Template) HasSplay() bool {
	if o != nil && o.Splay != nil {
		return true
	}

	return false
}

// SetSplay gets a reference to the given int64 and assigns it to the Splay field.
func (o *Template) SetSplay(v int64) {
	o.Splay = &v
}

// GetVaultGrace returns the VaultGrace field value if set, zero value otherwise.
func (o *Template) GetVaultGrace() int64 {
	if o == nil || o.VaultGrace == nil {
		var ret int64
		return ret
	}
	return *o.VaultGrace
}

// GetVaultGraceOk returns a tuple with the VaultGrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Template) GetVaultGraceOk() (*int64, bool) {
	if o == nil || o.VaultGrace == nil {
		return nil, false
	}
	return o.VaultGrace, true
}

// HasVaultGrace returns a boolean if a field has been set.
func (o *Template) HasVaultGrace() bool {
	if o != nil && o.VaultGrace != nil {
		return true
	}

	return false
}

// SetVaultGrace gets a reference to the given int64 and assigns it to the VaultGrace field.
func (o *Template) SetVaultGrace(v int64) {
	o.VaultGrace = &v
}

func (o Template) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeMode != nil {
		toSerialize["ChangeMode"] = o.ChangeMode
	}
	if o.ChangeSignal != nil {
		toSerialize["ChangeSignal"] = o.ChangeSignal
	}
	if o.DestPath != nil {
		toSerialize["DestPath"] = o.DestPath
	}
	if o.EmbeddedTmpl != nil {
		toSerialize["EmbeddedTmpl"] = o.EmbeddedTmpl
	}
	if o.Envvars != nil {
		toSerialize["Envvars"] = o.Envvars
	}
	if o.LeftDelim != nil {
		toSerialize["LeftDelim"] = o.LeftDelim
	}
	if o.Perms != nil {
		toSerialize["Perms"] = o.Perms
	}
	if o.RightDelim != nil {
		toSerialize["RightDelim"] = o.RightDelim
	}
	if o.SourcePath != nil {
		toSerialize["SourcePath"] = o.SourcePath
	}
	if o.Splay != nil {
		toSerialize["Splay"] = o.Splay
	}
	if o.VaultGrace != nil {
		toSerialize["VaultGrace"] = o.VaultGrace
	}
	return json.Marshal(toSerialize)
}

type NullableTemplate struct {
	value *Template
	isSet bool
}

func (v NullableTemplate) Get() *Template {
	return v.value
}

func (v *NullableTemplate) Set(val *Template) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplate(val *Template) *NullableTemplate {
	return &NullableTemplate{value: val, isSet: true}
}

func (v NullableTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

