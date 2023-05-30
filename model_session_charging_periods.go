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

// checks if the SessionChargingPeriods type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionChargingPeriods{}

// SessionChargingPeriods struct for SessionChargingPeriods
type SessionChargingPeriods struct {
	StartDateTime string `json:"start_date_time"`
	Dimensions *SessionChargingPeriodsDimensions `json:"dimensions,omitempty"`
	TariffId *string `json:"tariff_id,omitempty"`
}

// NewSessionChargingPeriods instantiates a new SessionChargingPeriods object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionChargingPeriods(startDateTime string) *SessionChargingPeriods {
	this := SessionChargingPeriods{}
	this.StartDateTime = startDateTime
	return &this
}

// NewSessionChargingPeriodsWithDefaults instantiates a new SessionChargingPeriods object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionChargingPeriodsWithDefaults() *SessionChargingPeriods {
	this := SessionChargingPeriods{}
	return &this
}

// GetStartDateTime returns the StartDateTime field value
func (o *SessionChargingPeriods) GetStartDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *SessionChargingPeriods) GetStartDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *SessionChargingPeriods) SetStartDateTime(v string) {
	o.StartDateTime = v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *SessionChargingPeriods) GetDimensions() SessionChargingPeriodsDimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret SessionChargingPeriodsDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionChargingPeriods) GetDimensionsOk() (*SessionChargingPeriodsDimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *SessionChargingPeriods) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given SessionChargingPeriodsDimensions and assigns it to the Dimensions field.
func (o *SessionChargingPeriods) SetDimensions(v SessionChargingPeriodsDimensions) {
	o.Dimensions = &v
}

// GetTariffId returns the TariffId field value if set, zero value otherwise.
func (o *SessionChargingPeriods) GetTariffId() string {
	if o == nil || IsNil(o.TariffId) {
		var ret string
		return ret
	}
	return *o.TariffId
}

// GetTariffIdOk returns a tuple with the TariffId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionChargingPeriods) GetTariffIdOk() (*string, bool) {
	if o == nil || IsNil(o.TariffId) {
		return nil, false
	}
	return o.TariffId, true
}

// HasTariffId returns a boolean if a field has been set.
func (o *SessionChargingPeriods) HasTariffId() bool {
	if o != nil && !IsNil(o.TariffId) {
		return true
	}

	return false
}

// SetTariffId gets a reference to the given string and assigns it to the TariffId field.
func (o *SessionChargingPeriods) SetTariffId(v string) {
	o.TariffId = &v
}

func (o SessionChargingPeriods) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionChargingPeriods) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["start_date_time"] = o.StartDateTime
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	if !IsNil(o.TariffId) {
		toSerialize["tariff_id"] = o.TariffId
	}
	return toSerialize, nil
}

type NullableSessionChargingPeriods struct {
	value *SessionChargingPeriods
	isSet bool
}

func (v NullableSessionChargingPeriods) Get() *SessionChargingPeriods {
	return v.value
}

func (v *NullableSessionChargingPeriods) Set(val *SessionChargingPeriods) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionChargingPeriods) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionChargingPeriods) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionChargingPeriods(val *SessionChargingPeriods) *NullableSessionChargingPeriods {
	return &NullableSessionChargingPeriods{value: val, isSet: true}
}

func (v NullableSessionChargingPeriods) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionChargingPeriods) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


