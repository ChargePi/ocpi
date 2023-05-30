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

// checks if the Session type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Session{}

// Session struct for Session
type Session struct {
	CountryCode string `json:"country_code"`
	PartyId string `json:"party_id"`
	Id string `json:"id"`
	StartDateTime string `json:"start_date_time"`
	EndDateTime *string `json:"end_date_time,omitempty"`
	Kwh float32 `json:"kwh"`
	CdrToken CdrBodyCdrToken `json:"cdr_token"`
	AuthMethod string `json:"auth_method"`
	AuthorizationReference *string `json:"authorization_reference,omitempty"`
	LocationId string `json:"location_id"`
	EvseUid string `json:"evse_uid"`
	ConnectorId string `json:"connector_id"`
	MeterId *string `json:"meter_id,omitempty"`
	Currency string `json:"currency"`
	ChargingPeriods *SessionChargingPeriods `json:"charging_periods,omitempty"`
	TotalCosts *SessionTotalCosts `json:"total_costs,omitempty"`
	Status string `json:"status"`
	LastUpdated string `json:"last_updated"`
}

// NewSession instantiates a new Session object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSession(countryCode string, partyId string, id string, startDateTime string, kwh float32, cdrToken CdrBodyCdrToken, authMethod string, locationId string, evseUid string, connectorId string, currency string, status string, lastUpdated string) *Session {
	this := Session{}
	this.CountryCode = countryCode
	this.PartyId = partyId
	this.Id = id
	this.StartDateTime = startDateTime
	this.Kwh = kwh
	this.CdrToken = cdrToken
	this.AuthMethod = authMethod
	this.LocationId = locationId
	this.EvseUid = evseUid
	this.ConnectorId = connectorId
	this.Currency = currency
	this.Status = status
	this.LastUpdated = lastUpdated
	return &this
}

// NewSessionWithDefaults instantiates a new Session object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithDefaults() *Session {
	this := Session{}
	return &this
}

// GetCountryCode returns the CountryCode field value
func (o *Session) GetCountryCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value
// and a boolean to check if the value has been set.
func (o *Session) GetCountryCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CountryCode, true
}

// SetCountryCode sets field value
func (o *Session) SetCountryCode(v string) {
	o.CountryCode = v
}

// GetPartyId returns the PartyId field value
func (o *Session) GetPartyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PartyId
}

// GetPartyIdOk returns a tuple with the PartyId field value
// and a boolean to check if the value has been set.
func (o *Session) GetPartyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PartyId, true
}

// SetPartyId sets field value
func (o *Session) SetPartyId(v string) {
	o.PartyId = v
}

// GetId returns the Id field value
func (o *Session) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Session) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Session) SetId(v string) {
	o.Id = v
}

// GetStartDateTime returns the StartDateTime field value
func (o *Session) GetStartDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value
// and a boolean to check if the value has been set.
func (o *Session) GetStartDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartDateTime, true
}

// SetStartDateTime sets field value
func (o *Session) SetStartDateTime(v string) {
	o.StartDateTime = v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *Session) GetEndDateTime() string {
	if o == nil || IsNil(o.EndDateTime) {
		var ret string
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetEndDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.EndDateTime) {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *Session) HasEndDateTime() bool {
	if o != nil && !IsNil(o.EndDateTime) {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given string and assigns it to the EndDateTime field.
func (o *Session) SetEndDateTime(v string) {
	o.EndDateTime = &v
}

// GetKwh returns the Kwh field value
func (o *Session) GetKwh() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Kwh
}

// GetKwhOk returns a tuple with the Kwh field value
// and a boolean to check if the value has been set.
func (o *Session) GetKwhOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kwh, true
}

// SetKwh sets field value
func (o *Session) SetKwh(v float32) {
	o.Kwh = v
}

// GetCdrToken returns the CdrToken field value
func (o *Session) GetCdrToken() CdrBodyCdrToken {
	if o == nil {
		var ret CdrBodyCdrToken
		return ret
	}

	return o.CdrToken
}

// GetCdrTokenOk returns a tuple with the CdrToken field value
// and a boolean to check if the value has been set.
func (o *Session) GetCdrTokenOk() (*CdrBodyCdrToken, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CdrToken, true
}

// SetCdrToken sets field value
func (o *Session) SetCdrToken(v CdrBodyCdrToken) {
	o.CdrToken = v
}

// GetAuthMethod returns the AuthMethod field value
func (o *Session) GetAuthMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthMethod
}

// GetAuthMethodOk returns a tuple with the AuthMethod field value
// and a boolean to check if the value has been set.
func (o *Session) GetAuthMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthMethod, true
}

// SetAuthMethod sets field value
func (o *Session) SetAuthMethod(v string) {
	o.AuthMethod = v
}

// GetAuthorizationReference returns the AuthorizationReference field value if set, zero value otherwise.
func (o *Session) GetAuthorizationReference() string {
	if o == nil || IsNil(o.AuthorizationReference) {
		var ret string
		return ret
	}
	return *o.AuthorizationReference
}

// GetAuthorizationReferenceOk returns a tuple with the AuthorizationReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetAuthorizationReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorizationReference) {
		return nil, false
	}
	return o.AuthorizationReference, true
}

// HasAuthorizationReference returns a boolean if a field has been set.
func (o *Session) HasAuthorizationReference() bool {
	if o != nil && !IsNil(o.AuthorizationReference) {
		return true
	}

	return false
}

// SetAuthorizationReference gets a reference to the given string and assigns it to the AuthorizationReference field.
func (o *Session) SetAuthorizationReference(v string) {
	o.AuthorizationReference = &v
}

// GetLocationId returns the LocationId field value
func (o *Session) GetLocationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value
// and a boolean to check if the value has been set.
func (o *Session) GetLocationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocationId, true
}

// SetLocationId sets field value
func (o *Session) SetLocationId(v string) {
	o.LocationId = v
}

// GetEvseUid returns the EvseUid field value
func (o *Session) GetEvseUid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EvseUid
}

// GetEvseUidOk returns a tuple with the EvseUid field value
// and a boolean to check if the value has been set.
func (o *Session) GetEvseUidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EvseUid, true
}

// SetEvseUid sets field value
func (o *Session) SetEvseUid(v string) {
	o.EvseUid = v
}

// GetConnectorId returns the ConnectorId field value
func (o *Session) GetConnectorId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorId
}

// GetConnectorIdOk returns a tuple with the ConnectorId field value
// and a boolean to check if the value has been set.
func (o *Session) GetConnectorIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorId, true
}

// SetConnectorId sets field value
func (o *Session) SetConnectorId(v string) {
	o.ConnectorId = v
}

// GetMeterId returns the MeterId field value if set, zero value otherwise.
func (o *Session) GetMeterId() string {
	if o == nil || IsNil(o.MeterId) {
		var ret string
		return ret
	}
	return *o.MeterId
}

// GetMeterIdOk returns a tuple with the MeterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetMeterIdOk() (*string, bool) {
	if o == nil || IsNil(o.MeterId) {
		return nil, false
	}
	return o.MeterId, true
}

// HasMeterId returns a boolean if a field has been set.
func (o *Session) HasMeterId() bool {
	if o != nil && !IsNil(o.MeterId) {
		return true
	}

	return false
}

// SetMeterId gets a reference to the given string and assigns it to the MeterId field.
func (o *Session) SetMeterId(v string) {
	o.MeterId = &v
}

// GetCurrency returns the Currency field value
func (o *Session) GetCurrency() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value
// and a boolean to check if the value has been set.
func (o *Session) GetCurrencyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Currency, true
}

// SetCurrency sets field value
func (o *Session) SetCurrency(v string) {
	o.Currency = v
}

// GetChargingPeriods returns the ChargingPeriods field value if set, zero value otherwise.
func (o *Session) GetChargingPeriods() SessionChargingPeriods {
	if o == nil || IsNil(o.ChargingPeriods) {
		var ret SessionChargingPeriods
		return ret
	}
	return *o.ChargingPeriods
}

// GetChargingPeriodsOk returns a tuple with the ChargingPeriods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetChargingPeriodsOk() (*SessionChargingPeriods, bool) {
	if o == nil || IsNil(o.ChargingPeriods) {
		return nil, false
	}
	return o.ChargingPeriods, true
}

// HasChargingPeriods returns a boolean if a field has been set.
func (o *Session) HasChargingPeriods() bool {
	if o != nil && !IsNil(o.ChargingPeriods) {
		return true
	}

	return false
}

// SetChargingPeriods gets a reference to the given SessionChargingPeriods and assigns it to the ChargingPeriods field.
func (o *Session) SetChargingPeriods(v SessionChargingPeriods) {
	o.ChargingPeriods = &v
}

// GetTotalCosts returns the TotalCosts field value if set, zero value otherwise.
func (o *Session) GetTotalCosts() SessionTotalCosts {
	if o == nil || IsNil(o.TotalCosts) {
		var ret SessionTotalCosts
		return ret
	}
	return *o.TotalCosts
}

// GetTotalCostsOk returns a tuple with the TotalCosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Session) GetTotalCostsOk() (*SessionTotalCosts, bool) {
	if o == nil || IsNil(o.TotalCosts) {
		return nil, false
	}
	return o.TotalCosts, true
}

// HasTotalCosts returns a boolean if a field has been set.
func (o *Session) HasTotalCosts() bool {
	if o != nil && !IsNil(o.TotalCosts) {
		return true
	}

	return false
}

// SetTotalCosts gets a reference to the given SessionTotalCosts and assigns it to the TotalCosts field.
func (o *Session) SetTotalCosts(v SessionTotalCosts) {
	o.TotalCosts = &v
}

// GetStatus returns the Status field value
func (o *Session) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Session) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Session) SetStatus(v string) {
	o.Status = v
}

// GetLastUpdated returns the LastUpdated field value
func (o *Session) GetLastUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *Session) GetLastUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *Session) SetLastUpdated(v string) {
	o.LastUpdated = v
}

func (o Session) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Session) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["country_code"] = o.CountryCode
	toSerialize["party_id"] = o.PartyId
	toSerialize["id"] = o.Id
	toSerialize["start_date_time"] = o.StartDateTime
	if !IsNil(o.EndDateTime) {
		toSerialize["end_date_time"] = o.EndDateTime
	}
	toSerialize["kwh"] = o.Kwh
	toSerialize["cdr_token"] = o.CdrToken
	toSerialize["auth_method"] = o.AuthMethod
	if !IsNil(o.AuthorizationReference) {
		toSerialize["authorization_reference"] = o.AuthorizationReference
	}
	toSerialize["location_id"] = o.LocationId
	toSerialize["evse_uid"] = o.EvseUid
	toSerialize["connector_id"] = o.ConnectorId
	if !IsNil(o.MeterId) {
		toSerialize["meter_id"] = o.MeterId
	}
	toSerialize["currency"] = o.Currency
	if !IsNil(o.ChargingPeriods) {
		toSerialize["charging_periods"] = o.ChargingPeriods
	}
	if !IsNil(o.TotalCosts) {
		toSerialize["total_costs"] = o.TotalCosts
	}
	toSerialize["status"] = o.Status
	toSerialize["last_updated"] = o.LastUpdated
	return toSerialize, nil
}

type NullableSession struct {
	value *Session
	isSet bool
}

func (v NullableSession) Get() *Session {
	return v.value
}

func (v *NullableSession) Set(val *Session) {
	v.value = val
	v.isSet = true
}

func (v NullableSession) IsSet() bool {
	return v.isSet
}

func (v *NullableSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSession(val *Session) *NullableSession {
	return &NullableSession{value: val, isSet: true}
}

func (v NullableSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


