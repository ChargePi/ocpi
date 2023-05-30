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

// checks if the ChargingPreferencesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChargingPreferencesResponse{}

// ChargingPreferencesResponse struct for ChargingPreferencesResponse
type ChargingPreferencesResponse struct {
	ChargingPreferences string `json:"charging_preferences"`
	StatusCode float32 `json:"status_code"`
	StatusMessage *string `json:"status_message,omitempty"`
	TimeStamp *string `json:"timeStamp,omitempty"`
}

// NewChargingPreferencesResponse instantiates a new ChargingPreferencesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChargingPreferencesResponse(chargingPreferences string, statusCode float32) *ChargingPreferencesResponse {
	this := ChargingPreferencesResponse{}
	this.ChargingPreferences = chargingPreferences
	this.StatusCode = statusCode
	return &this
}

// NewChargingPreferencesResponseWithDefaults instantiates a new ChargingPreferencesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChargingPreferencesResponseWithDefaults() *ChargingPreferencesResponse {
	this := ChargingPreferencesResponse{}
	return &this
}

// GetChargingPreferences returns the ChargingPreferences field value
func (o *ChargingPreferencesResponse) GetChargingPreferences() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ChargingPreferences
}

// GetChargingPreferencesOk returns a tuple with the ChargingPreferences field value
// and a boolean to check if the value has been set.
func (o *ChargingPreferencesResponse) GetChargingPreferencesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChargingPreferences, true
}

// SetChargingPreferences sets field value
func (o *ChargingPreferencesResponse) SetChargingPreferences(v string) {
	o.ChargingPreferences = v
}

// GetStatusCode returns the StatusCode field value
func (o *ChargingPreferencesResponse) GetStatusCode() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *ChargingPreferencesResponse) GetStatusCodeOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *ChargingPreferencesResponse) SetStatusCode(v float32) {
	o.StatusCode = v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *ChargingPreferencesResponse) GetStatusMessage() string {
	if o == nil || IsNil(o.StatusMessage) {
		var ret string
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingPreferencesResponse) GetStatusMessageOk() (*string, bool) {
	if o == nil || IsNil(o.StatusMessage) {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *ChargingPreferencesResponse) HasStatusMessage() bool {
	if o != nil && !IsNil(o.StatusMessage) {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given string and assigns it to the StatusMessage field.
func (o *ChargingPreferencesResponse) SetStatusMessage(v string) {
	o.StatusMessage = &v
}

// GetTimeStamp returns the TimeStamp field value if set, zero value otherwise.
func (o *ChargingPreferencesResponse) GetTimeStamp() string {
	if o == nil || IsNil(o.TimeStamp) {
		var ret string
		return ret
	}
	return *o.TimeStamp
}

// GetTimeStampOk returns a tuple with the TimeStamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChargingPreferencesResponse) GetTimeStampOk() (*string, bool) {
	if o == nil || IsNil(o.TimeStamp) {
		return nil, false
	}
	return o.TimeStamp, true
}

// HasTimeStamp returns a boolean if a field has been set.
func (o *ChargingPreferencesResponse) HasTimeStamp() bool {
	if o != nil && !IsNil(o.TimeStamp) {
		return true
	}

	return false
}

// SetTimeStamp gets a reference to the given string and assigns it to the TimeStamp field.
func (o *ChargingPreferencesResponse) SetTimeStamp(v string) {
	o.TimeStamp = &v
}

func (o ChargingPreferencesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChargingPreferencesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["charging_preferences"] = o.ChargingPreferences
	toSerialize["status_code"] = o.StatusCode
	if !IsNil(o.StatusMessage) {
		toSerialize["status_message"] = o.StatusMessage
	}
	if !IsNil(o.TimeStamp) {
		toSerialize["timeStamp"] = o.TimeStamp
	}
	return toSerialize, nil
}

type NullableChargingPreferencesResponse struct {
	value *ChargingPreferencesResponse
	isSet bool
}

func (v NullableChargingPreferencesResponse) Get() *ChargingPreferencesResponse {
	return v.value
}

func (v *NullableChargingPreferencesResponse) Set(val *ChargingPreferencesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingPreferencesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingPreferencesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingPreferencesResponse(val *ChargingPreferencesResponse) *NullableChargingPreferencesResponse {
	return &NullableChargingPreferencesResponse{value: val, isSet: true}
}

func (v NullableChargingPreferencesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingPreferencesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


