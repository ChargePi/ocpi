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

// checks if the ChargingProfileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargingProfileResponse{}

// ChargingProfileResponse struct for ChargingProfileResponse
type ChargingProfileResponse struct {
	Result ChargingProfileResponseType `json:"result"`
	Timeout int32 `json:"timeout"`
}

// NewChargingProfileResponse instantiates a new ChargingProfileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargingProfileResponse(result ChargingProfileResponseType, timeout int32) *ChargingProfileResponse {
	this := ChargingProfileResponse{}
	this.Result = result
	this.Timeout = timeout
	return &this
}

// NewChargingProfileResponseWithDefaults instantiates a new ChargingProfileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargingProfileResponseWithDefaults() *ChargingProfileResponse {
	this := ChargingProfileResponse{}
	return &this
}

// GetResult returns the Result field value
func (o *ChargingProfileResponse) GetResult() ChargingProfileResponseType {
	if o == nil {
		var ret ChargingProfileResponseType
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *ChargingProfileResponse) GetResultOk() (*ChargingProfileResponseType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *ChargingProfileResponse) SetResult(v ChargingProfileResponseType) {
	o.Result = v
}

// GetTimeout returns the Timeout field value
func (o *ChargingProfileResponse) GetTimeout() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value
// and a boolean to check if the value has been set.
func (o *ChargingProfileResponse) GetTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timeout, true
}

// SetTimeout sets field value
func (o *ChargingProfileResponse) SetTimeout(v int32) {
	o.Timeout = v
}

func (o ChargingProfileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargingProfileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["result"] = o.Result
	toSerialize["timeout"] = o.Timeout
	return toSerialize, nil
}

type NullableChargingProfileResponse struct {
	value *ChargingProfileResponse
	isSet bool
}

func (v NullableChargingProfileResponse) Get() *ChargingProfileResponse {
	return v.value
}

func (v *NullableChargingProfileResponse) Set(val *ChargingProfileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingProfileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingProfileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingProfileResponse(val *ChargingProfileResponse) *NullableChargingProfileResponse {
	return &NullableChargingProfileResponse{value: val, isSet: true}
}

func (v NullableChargingProfileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingProfileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


