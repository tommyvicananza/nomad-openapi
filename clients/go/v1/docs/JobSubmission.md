# JobSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Format** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**VariableFlags** | Pointer to **map[string]string** |  | [optional] 
**Variables** | Pointer to **string** |  | [optional] 

## Methods

### NewJobSubmission

`func NewJobSubmission() *JobSubmission`

NewJobSubmission instantiates a new JobSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobSubmissionWithDefaults

`func NewJobSubmissionWithDefaults() *JobSubmission`

NewJobSubmissionWithDefaults instantiates a new JobSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormat

`func (o *JobSubmission) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *JobSubmission) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *JobSubmission) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *JobSubmission) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetSource

`func (o *JobSubmission) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *JobSubmission) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *JobSubmission) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *JobSubmission) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetVariableFlags

`func (o *JobSubmission) GetVariableFlags() map[string]string`

GetVariableFlags returns the VariableFlags field if non-nil, zero value otherwise.

### GetVariableFlagsOk

`func (o *JobSubmission) GetVariableFlagsOk() (*map[string]string, bool)`

GetVariableFlagsOk returns a tuple with the VariableFlags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableFlags

`func (o *JobSubmission) SetVariableFlags(v map[string]string)`

SetVariableFlags sets VariableFlags field to given value.

### HasVariableFlags

`func (o *JobSubmission) HasVariableFlags() bool`

HasVariableFlags returns a boolean if a field has been set.

### GetVariables

`func (o *JobSubmission) GetVariables() string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *JobSubmission) GetVariablesOk() (*string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *JobSubmission) SetVariables(v string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *JobSubmission) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


