/*
OCPI modules

Specification for OCPIs modules handlers

API version: 2.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package OCPI

import (
	"encoding/json"
)

// checks if the GenericResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericResponse{}

// GenericResponse struct for GenericResponse
type GenericResponse struct {
	StatusCode float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	TimeStamp string `json:"timeStamp"`
}

// NewGenericResponse instantiates a new GenericResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericResponse(statusCode float32, timeStamp string) *GenericResponse {
	this := GenericResponse{}
	this.StatusCode = statusCode
	this.TimeStamp = timeStamp
	return &this
}

// NewGenericResponseWithDefaults instantiates a new GenericResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericResponseWithDefaults() *GenericResponse {
	this := GenericResponse{}
	return &this
}

// GetStatusCode returns the StatusCode field value
func (o *GenericResponse) GetStatusCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *GenericResponse) GetStatusCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *GenericResponse) SetStatusCode(v float32) {
	o.StatusCode = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *GenericResponse) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericResponse) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *GenericResponse) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *GenericResponse) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetTimeStamp returns the TimeStamp field value
func (o *GenericResponse) GetTimeStamp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value
// and a boolean to check if the value has been set.
func (o *GenericResponse) GetTimeStampOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeStamp, true
}

// SetTimeStamp sets field value
func (o *GenericResponse) SetTimeStamp(v string) {
	o.TimeStamp = v
}

func (o GenericResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status_code"] = o.StatusCode
	if !IsNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	toSerialize["timeStamp"] = o.TimeStamp
	return toSerialize, nil
}

type NullableGenericResponse struct {
	value *GenericResponse
	isSet bool
}

func (v NullableGenericResponse) Get() *GenericResponse {
	return v.value
}

func (v *NullableGenericResponse) Set(val *GenericResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericResponse(val *GenericResponse) *NullableGenericResponse {
	return &NullableGenericResponse{value: val, isSet: true}
}

func (v NullableGenericResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


