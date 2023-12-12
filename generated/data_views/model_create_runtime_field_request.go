/*
Data views

OpenAPI schema for data view endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package data_views

import (
	"encoding/json"
)

// checks if the CreateRuntimeFieldRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateRuntimeFieldRequest{}

// CreateRuntimeFieldRequest struct for CreateRuntimeFieldRequest
type CreateRuntimeFieldRequest struct {
	// The name for a runtime field.
	Name interface{} `json:"name"`
	// The runtime field definition object.
	RuntimeField interface{} `json:"runtimeField"`
}

// NewCreateRuntimeFieldRequest instantiates a new CreateRuntimeFieldRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateRuntimeFieldRequest(name interface{}, runtimeField interface{}) *CreateRuntimeFieldRequest {
	this := CreateRuntimeFieldRequest{}
	this.Name = name
	this.RuntimeField = runtimeField
	return &this
}

// NewCreateRuntimeFieldRequestWithDefaults instantiates a new CreateRuntimeFieldRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateRuntimeFieldRequestWithDefaults() *CreateRuntimeFieldRequest {
	this := CreateRuntimeFieldRequest{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CreateRuntimeFieldRequest) GetName() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRuntimeFieldRequest) GetNameOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateRuntimeFieldRequest) SetName(v interface{}) {
	o.Name = v
}

// GetRuntimeField returns the RuntimeField field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *CreateRuntimeFieldRequest) GetRuntimeField() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.RuntimeField
}

// GetRuntimeFieldOk returns a tuple with the RuntimeField field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateRuntimeFieldRequest) GetRuntimeFieldOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RuntimeField) {
		return nil, false
	}
	return &o.RuntimeField, true
}

// SetRuntimeField sets field value
func (o *CreateRuntimeFieldRequest) SetRuntimeField(v interface{}) {
	o.RuntimeField = v
}

func (o CreateRuntimeFieldRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateRuntimeFieldRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RuntimeField != nil {
		toSerialize["runtimeField"] = o.RuntimeField
	}
	return toSerialize, nil
}

type NullableCreateRuntimeFieldRequest struct {
	value *CreateRuntimeFieldRequest
	isSet bool
}

func (v NullableCreateRuntimeFieldRequest) Get() *CreateRuntimeFieldRequest {
	return v.value
}

func (v *NullableCreateRuntimeFieldRequest) Set(val *CreateRuntimeFieldRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateRuntimeFieldRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateRuntimeFieldRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateRuntimeFieldRequest(val *CreateRuntimeFieldRequest) *NullableCreateRuntimeFieldRequest {
	return &NullableCreateRuntimeFieldRequest{value: val, isSet: true}
}

func (v NullableCreateRuntimeFieldRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateRuntimeFieldRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}