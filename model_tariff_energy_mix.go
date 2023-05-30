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

// checks if the TariffEnergyMix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TariffEnergyMix{}

// TariffEnergyMix struct for TariffEnergyMix
type TariffEnergyMix struct {
	IsGreenEnergy bool `json:"is_green_energy"`
	EnergySources *CdrTariffsEnergyMixEnergySources `json:"energy_sources,omitempty"`
	EnvironImpact *TariffEnergyMixEnvironImpact `json:"environ_impact,omitempty"`
	SupplierName *string `json:"supplier_name,omitempty"`
	EnergyProductName *string `json:"energy_product_name,omitempty"`
}

// NewTariffEnergyMix instantiates a new TariffEnergyMix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTariffEnergyMix(isGreenEnergy bool) *TariffEnergyMix {
	this := TariffEnergyMix{}
	this.IsGreenEnergy = isGreenEnergy
	return &this
}

// NewTariffEnergyMixWithDefaults instantiates a new TariffEnergyMix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTariffEnergyMixWithDefaults() *TariffEnergyMix {
	this := TariffEnergyMix{}
	return &this
}

// GetIsGreenEnergy returns the IsGreenEnergy field value
func (o *TariffEnergyMix) GetIsGreenEnergy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsGreenEnergy
}

// GetIsGreenEnergyOk returns a tuple with the IsGreenEnergy field value
// and a boolean to check if the value has been set.
func (o *TariffEnergyMix) GetIsGreenEnergyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsGreenEnergy, true
}

// SetIsGreenEnergy sets field value
func (o *TariffEnergyMix) SetIsGreenEnergy(v bool) {
	o.IsGreenEnergy = v
}

// GetEnergySources returns the EnergySources field value if set, zero value otherwise.
func (o *TariffEnergyMix) GetEnergySources() CdrTariffsEnergyMixEnergySources {
	if o == nil || IsNil(o.EnergySources) {
		var ret CdrTariffsEnergyMixEnergySources
		return ret
	}
	return *o.EnergySources
}

// GetEnergySourcesOk returns a tuple with the EnergySources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffEnergyMix) GetEnergySourcesOk() (*CdrTariffsEnergyMixEnergySources, bool) {
	if o == nil || IsNil(o.EnergySources) {
		return nil, false
	}
	return o.EnergySources, true
}

// HasEnergySources returns a boolean if a field has been set.
func (o *TariffEnergyMix) HasEnergySources() bool {
	if o != nil && !IsNil(o.EnergySources) {
		return true
	}

	return false
}

// SetEnergySources gets a reference to the given CdrTariffsEnergyMixEnergySources and assigns it to the EnergySources field.
func (o *TariffEnergyMix) SetEnergySources(v CdrTariffsEnergyMixEnergySources) {
	o.EnergySources = &v
}

// GetEnvironImpact returns the EnvironImpact field value if set, zero value otherwise.
func (o *TariffEnergyMix) GetEnvironImpact() TariffEnergyMixEnvironImpact {
	if o == nil || IsNil(o.EnvironImpact) {
		var ret TariffEnergyMixEnvironImpact
		return ret
	}
	return *o.EnvironImpact
}

// GetEnvironImpactOk returns a tuple with the EnvironImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffEnergyMix) GetEnvironImpactOk() (*TariffEnergyMixEnvironImpact, bool) {
	if o == nil || IsNil(o.EnvironImpact) {
		return nil, false
	}
	return o.EnvironImpact, true
}

// HasEnvironImpact returns a boolean if a field has been set.
func (o *TariffEnergyMix) HasEnvironImpact() bool {
	if o != nil && !IsNil(o.EnvironImpact) {
		return true
	}

	return false
}

// SetEnvironImpact gets a reference to the given TariffEnergyMixEnvironImpact and assigns it to the EnvironImpact field.
func (o *TariffEnergyMix) SetEnvironImpact(v TariffEnergyMixEnvironImpact) {
	o.EnvironImpact = &v
}

// GetSupplierName returns the SupplierName field value if set, zero value otherwise.
func (o *TariffEnergyMix) GetSupplierName() string {
	if o == nil || IsNil(o.SupplierName) {
		var ret string
		return ret
	}
	return *o.SupplierName
}

// GetSupplierNameOk returns a tuple with the SupplierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffEnergyMix) GetSupplierNameOk() (*string, bool) {
	if o == nil || IsNil(o.SupplierName) {
		return nil, false
	}
	return o.SupplierName, true
}

// HasSupplierName returns a boolean if a field has been set.
func (o *TariffEnergyMix) HasSupplierName() bool {
	if o != nil && !IsNil(o.SupplierName) {
		return true
	}

	return false
}

// SetSupplierName gets a reference to the given string and assigns it to the SupplierName field.
func (o *TariffEnergyMix) SetSupplierName(v string) {
	o.SupplierName = &v
}

// GetEnergyProductName returns the EnergyProductName field value if set, zero value otherwise.
func (o *TariffEnergyMix) GetEnergyProductName() string {
	if o == nil || IsNil(o.EnergyProductName) {
		var ret string
		return ret
	}
	return *o.EnergyProductName
}

// GetEnergyProductNameOk returns a tuple with the EnergyProductName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TariffEnergyMix) GetEnergyProductNameOk() (*string, bool) {
	if o == nil || IsNil(o.EnergyProductName) {
		return nil, false
	}
	return o.EnergyProductName, true
}

// HasEnergyProductName returns a boolean if a field has been set.
func (o *TariffEnergyMix) HasEnergyProductName() bool {
	if o != nil && !IsNil(o.EnergyProductName) {
		return true
	}

	return false
}

// SetEnergyProductName gets a reference to the given string and assigns it to the EnergyProductName field.
func (o *TariffEnergyMix) SetEnergyProductName(v string) {
	o.EnergyProductName = &v
}

func (o TariffEnergyMix) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TariffEnergyMix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["is_green_energy"] = o.IsGreenEnergy
	if !IsNil(o.EnergySources) {
		toSerialize["energy_sources"] = o.EnergySources
	}
	if !IsNil(o.EnvironImpact) {
		toSerialize["environ_impact"] = o.EnvironImpact
	}
	if !IsNil(o.SupplierName) {
		toSerialize["supplier_name"] = o.SupplierName
	}
	if !IsNil(o.EnergyProductName) {
		toSerialize["energy_product_name"] = o.EnergyProductName
	}
	return toSerialize, nil
}

type NullableTariffEnergyMix struct {
	value *TariffEnergyMix
	isSet bool
}

func (v NullableTariffEnergyMix) Get() *TariffEnergyMix {
	return v.value
}

func (v *NullableTariffEnergyMix) Set(val *TariffEnergyMix) {
	v.value = val
	v.isSet = true
}

func (v NullableTariffEnergyMix) IsSet() bool {
	return v.isSet
}

func (v *NullableTariffEnergyMix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTariffEnergyMix(val *TariffEnergyMix) *NullableTariffEnergyMix {
	return &NullableTariffEnergyMix{value: val, isSet: true}
}

func (v NullableTariffEnergyMix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTariffEnergyMix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


