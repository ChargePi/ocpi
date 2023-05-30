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

// checks if the LocationsData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationsData{}

// LocationsData struct for LocationsData
type LocationsData struct {
	CountryCode *string `json:"country_code,omitempty"`
	PartyId *string `json:"party_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Publish *bool `json:"publish,omitempty"`
	PublishAllowedTo *LocationsDataPublishAllowedTo `json:"publish_allowed_to,omitempty"`
	Name *string `json:"name,omitempty"`
	Address string `json:"address"`
	City string `json:"city"`
	PostalCode *string `json:"postal_code,omitempty"`
	State *string `json:"state,omitempty"`
	Country string `json:"country"`
	Coordinates CdrCdrLocationCoordinates `json:"coordinates"`
	RelatedLocations *LocationsDataRelatedLocations `json:"related_locations,omitempty"`
	ParkingType *string `json:"parking_type,omitempty"`
	Evses *Evse `json:"evses,omitempty"`
	Directions *CdrTariffsTariffAltText `json:"directions,omitempty"`
	Operator *BusinessDetails `json:"operator,omitempty"`
	Suboperator *BusinessDetails `json:"suboperator,omitempty"`
	Owner *BusinessDetails `json:"owner,omitempty"`
	Facilities *string `json:"facilities,omitempty"`
	TimeZone string `json:"time_zone"`
	OpeningTimes *LocationsDataOpeningTimes `json:"opening_times,omitempty"`
	ChargingWhenClosed *bool `json:"charging_when_closed,omitempty"`
	Images *Image `json:"images,omitempty"`
	EnergyMix *CdrTariffsEnergyMix `json:"energy_mix,omitempty"`
	LastUpdated string `json:"last_updated"`
}

// NewLocationsData instantiates a new LocationsData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationsData(address string, city string, country string, coordinates CdrCdrLocationCoordinates, timeZone string, lastUpdated string) *LocationsData {
	this := LocationsData{}
	this.Address = address
	this.City = city
	this.Country = country
	this.Coordinates = coordinates
	this.TimeZone = timeZone
	this.LastUpdated = lastUpdated
	return &this
}

// NewLocationsDataWithDefaults instantiates a new LocationsData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationsDataWithDefaults() *LocationsData {
	this := LocationsData{}
	return &this
}

// GetCountryCode returns the CountryCode field value if set, zero value otherwise.
func (o *LocationsData) GetCountryCode() string {
	if o == nil || IsNil(o.CountryCode) {
		var ret string
		return ret
	}
	return *o.CountryCode
}

// GetCountryCodeOk returns a tuple with the CountryCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetCountryCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CountryCode) {
		return nil, false
	}
	return o.CountryCode, true
}

// HasCountryCode returns a boolean if a field has been set.
func (o *LocationsData) HasCountryCode() bool {
	if o != nil && !IsNil(o.CountryCode) {
		return true
	}

	return false
}

// SetCountryCode gets a reference to the given string and assigns it to the CountryCode field.
func (o *LocationsData) SetCountryCode(v string) {
	o.CountryCode = &v
}

// GetPartyId returns the PartyId field value if set, zero value otherwise.
func (o *LocationsData) GetPartyId() string {
	if o == nil || IsNil(o.PartyId) {
		var ret string
		return ret
	}
	return *o.PartyId
}

// GetPartyIdOk returns a tuple with the PartyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetPartyIdOk() (*string, bool) {
	if o == nil || IsNil(o.PartyId) {
		return nil, false
	}
	return o.PartyId, true
}

// HasPartyId returns a boolean if a field has been set.
func (o *LocationsData) HasPartyId() bool {
	if o != nil && !IsNil(o.PartyId) {
		return true
	}

	return false
}

// SetPartyId gets a reference to the given string and assigns it to the PartyId field.
func (o *LocationsData) SetPartyId(v string) {
	o.PartyId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *LocationsData) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *LocationsData) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *LocationsData) SetId(v string) {
	o.Id = &v
}

// GetPublish returns the Publish field value if set, zero value otherwise.
func (o *LocationsData) GetPublish() bool {
	if o == nil || IsNil(o.Publish) {
		var ret bool
		return ret
	}
	return *o.Publish
}

// GetPublishOk returns a tuple with the Publish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetPublishOk() (*bool, bool) {
	if o == nil || IsNil(o.Publish) {
		return nil, false
	}
	return o.Publish, true
}

// HasPublish returns a boolean if a field has been set.
func (o *LocationsData) HasPublish() bool {
	if o != nil && !IsNil(o.Publish) {
		return true
	}

	return false
}

// SetPublish gets a reference to the given bool and assigns it to the Publish field.
func (o *LocationsData) SetPublish(v bool) {
	o.Publish = &v
}

// GetPublishAllowedTo returns the PublishAllowedTo field value if set, zero value otherwise.
func (o *LocationsData) GetPublishAllowedTo() LocationsDataPublishAllowedTo {
	if o == nil || IsNil(o.PublishAllowedTo) {
		var ret LocationsDataPublishAllowedTo
		return ret
	}
	return *o.PublishAllowedTo
}

// GetPublishAllowedToOk returns a tuple with the PublishAllowedTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetPublishAllowedToOk() (*LocationsDataPublishAllowedTo, bool) {
	if o == nil || IsNil(o.PublishAllowedTo) {
		return nil, false
	}
	return o.PublishAllowedTo, true
}

// HasPublishAllowedTo returns a boolean if a field has been set.
func (o *LocationsData) HasPublishAllowedTo() bool {
	if o != nil && !IsNil(o.PublishAllowedTo) {
		return true
	}

	return false
}

// SetPublishAllowedTo gets a reference to the given LocationsDataPublishAllowedTo and assigns it to the PublishAllowedTo field.
func (o *LocationsData) SetPublishAllowedTo(v LocationsDataPublishAllowedTo) {
	o.PublishAllowedTo = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LocationsData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LocationsData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LocationsData) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value
func (o *LocationsData) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *LocationsData) GetAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *LocationsData) SetAddress(v string) {
	o.Address = v
}

// GetCity returns the City field value
func (o *LocationsData) GetCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.City
}

// GetCityOk returns a tuple with the City field value
// and a boolean to check if the value has been set.
func (o *LocationsData) GetCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.City, true
}

// SetCity sets field value
func (o *LocationsData) SetCity(v string) {
	o.City = v
}

// GetPostalCode returns the PostalCode field value if set, zero value otherwise.
func (o *LocationsData) GetPostalCode() string {
	if o == nil || IsNil(o.PostalCode) {
		var ret string
		return ret
	}
	return *o.PostalCode
}

// GetPostalCodeOk returns a tuple with the PostalCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetPostalCodeOk() (*string, bool) {
	if o == nil || IsNil(o.PostalCode) {
		return nil, false
	}
	return o.PostalCode, true
}

// HasPostalCode returns a boolean if a field has been set.
func (o *LocationsData) HasPostalCode() bool {
	if o != nil && !IsNil(o.PostalCode) {
		return true
	}

	return false
}

// SetPostalCode gets a reference to the given string and assigns it to the PostalCode field.
func (o *LocationsData) SetPostalCode(v string) {
	o.PostalCode = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *LocationsData) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *LocationsData) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *LocationsData) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value
func (o *LocationsData) GetCountry() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Country
}

// GetCountryOk returns a tuple with the Country field value
// and a boolean to check if the value has been set.
func (o *LocationsData) GetCountryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Country, true
}

// SetCountry sets field value
func (o *LocationsData) SetCountry(v string) {
	o.Country = v
}

// GetCoordinates returns the Coordinates field value
func (o *LocationsData) GetCoordinates() CdrCdrLocationCoordinates {
	if o == nil {
		var ret CdrCdrLocationCoordinates
		return ret
	}

	return o.Coordinates
}

// GetCoordinatesOk returns a tuple with the Coordinates field value
// and a boolean to check if the value has been set.
func (o *LocationsData) GetCoordinatesOk() (*CdrCdrLocationCoordinates, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Coordinates, true
}

// SetCoordinates sets field value
func (o *LocationsData) SetCoordinates(v CdrCdrLocationCoordinates) {
	o.Coordinates = v
}

// GetRelatedLocations returns the RelatedLocations field value if set, zero value otherwise.
func (o *LocationsData) GetRelatedLocations() LocationsDataRelatedLocations {
	if o == nil || IsNil(o.RelatedLocations) {
		var ret LocationsDataRelatedLocations
		return ret
	}
	return *o.RelatedLocations
}

// GetRelatedLocationsOk returns a tuple with the RelatedLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetRelatedLocationsOk() (*LocationsDataRelatedLocations, bool) {
	if o == nil || IsNil(o.RelatedLocations) {
		return nil, false
	}
	return o.RelatedLocations, true
}

// HasRelatedLocations returns a boolean if a field has been set.
func (o *LocationsData) HasRelatedLocations() bool {
	if o != nil && !IsNil(o.RelatedLocations) {
		return true
	}

	return false
}

// SetRelatedLocations gets a reference to the given LocationsDataRelatedLocations and assigns it to the RelatedLocations field.
func (o *LocationsData) SetRelatedLocations(v LocationsDataRelatedLocations) {
	o.RelatedLocations = &v
}

// GetParkingType returns the ParkingType field value if set, zero value otherwise.
func (o *LocationsData) GetParkingType() string {
	if o == nil || IsNil(o.ParkingType) {
		var ret string
		return ret
	}
	return *o.ParkingType
}

// GetParkingTypeOk returns a tuple with the ParkingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetParkingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParkingType) {
		return nil, false
	}
	return o.ParkingType, true
}

// HasParkingType returns a boolean if a field has been set.
func (o *LocationsData) HasParkingType() bool {
	if o != nil && !IsNil(o.ParkingType) {
		return true
	}

	return false
}

// SetParkingType gets a reference to the given string and assigns it to the ParkingType field.
func (o *LocationsData) SetParkingType(v string) {
	o.ParkingType = &v
}

// GetEvses returns the Evses field value if set, zero value otherwise.
func (o *LocationsData) GetEvses() Evse {
	if o == nil || IsNil(o.Evses) {
		var ret Evse
		return ret
	}
	return *o.Evses
}

// GetEvsesOk returns a tuple with the Evses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetEvsesOk() (*Evse, bool) {
	if o == nil || IsNil(o.Evses) {
		return nil, false
	}
	return o.Evses, true
}

// HasEvses returns a boolean if a field has been set.
func (o *LocationsData) HasEvses() bool {
	if o != nil && !IsNil(o.Evses) {
		return true
	}

	return false
}

// SetEvses gets a reference to the given Evse and assigns it to the Evses field.
func (o *LocationsData) SetEvses(v Evse) {
	o.Evses = &v
}

// GetDirections returns the Directions field value if set, zero value otherwise.
func (o *LocationsData) GetDirections() CdrTariffsTariffAltText {
	if o == nil || IsNil(o.Directions) {
		var ret CdrTariffsTariffAltText
		return ret
	}
	return *o.Directions
}

// GetDirectionsOk returns a tuple with the Directions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetDirectionsOk() (*CdrTariffsTariffAltText, bool) {
	if o == nil || IsNil(o.Directions) {
		return nil, false
	}
	return o.Directions, true
}

// HasDirections returns a boolean if a field has been set.
func (o *LocationsData) HasDirections() bool {
	if o != nil && !IsNil(o.Directions) {
		return true
	}

	return false
}

// SetDirections gets a reference to the given CdrTariffsTariffAltText and assigns it to the Directions field.
func (o *LocationsData) SetDirections(v CdrTariffsTariffAltText) {
	o.Directions = &v
}

// GetOperator returns the Operator field value if set, zero value otherwise.
func (o *LocationsData) GetOperator() BusinessDetails {
	if o == nil || IsNil(o.Operator) {
		var ret BusinessDetails
		return ret
	}
	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetOperatorOk() (*BusinessDetails, bool) {
	if o == nil || IsNil(o.Operator) {
		return nil, false
	}
	return o.Operator, true
}

// HasOperator returns a boolean if a field has been set.
func (o *LocationsData) HasOperator() bool {
	if o != nil && !IsNil(o.Operator) {
		return true
	}

	return false
}

// SetOperator gets a reference to the given BusinessDetails and assigns it to the Operator field.
func (o *LocationsData) SetOperator(v BusinessDetails) {
	o.Operator = &v
}

// GetSuboperator returns the Suboperator field value if set, zero value otherwise.
func (o *LocationsData) GetSuboperator() BusinessDetails {
	if o == nil || IsNil(o.Suboperator) {
		var ret BusinessDetails
		return ret
	}
	return *o.Suboperator
}

// GetSuboperatorOk returns a tuple with the Suboperator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetSuboperatorOk() (*BusinessDetails, bool) {
	if o == nil || IsNil(o.Suboperator) {
		return nil, false
	}
	return o.Suboperator, true
}

// HasSuboperator returns a boolean if a field has been set.
func (o *LocationsData) HasSuboperator() bool {
	if o != nil && !IsNil(o.Suboperator) {
		return true
	}

	return false
}

// SetSuboperator gets a reference to the given BusinessDetails and assigns it to the Suboperator field.
func (o *LocationsData) SetSuboperator(v BusinessDetails) {
	o.Suboperator = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *LocationsData) GetOwner() BusinessDetails {
	if o == nil || IsNil(o.Owner) {
		var ret BusinessDetails
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetOwnerOk() (*BusinessDetails, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *LocationsData) HasOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given BusinessDetails and assigns it to the Owner field.
func (o *LocationsData) SetOwner(v BusinessDetails) {
	o.Owner = &v
}

// GetFacilities returns the Facilities field value if set, zero value otherwise.
func (o *LocationsData) GetFacilities() string {
	if o == nil || IsNil(o.Facilities) {
		var ret string
		return ret
	}
	return *o.Facilities
}

// GetFacilitiesOk returns a tuple with the Facilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetFacilitiesOk() (*string, bool) {
	if o == nil || IsNil(o.Facilities) {
		return nil, false
	}
	return o.Facilities, true
}

// HasFacilities returns a boolean if a field has been set.
func (o *LocationsData) HasFacilities() bool {
	if o != nil && !IsNil(o.Facilities) {
		return true
	}

	return false
}

// SetFacilities gets a reference to the given string and assigns it to the Facilities field.
func (o *LocationsData) SetFacilities(v string) {
	o.Facilities = &v
}

// GetTimeZone returns the TimeZone field value
func (o *LocationsData) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *LocationsData) GetTimeZoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *LocationsData) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetOpeningTimes returns the OpeningTimes field value if set, zero value otherwise.
func (o *LocationsData) GetOpeningTimes() LocationsDataOpeningTimes {
	if o == nil || IsNil(o.OpeningTimes) {
		var ret LocationsDataOpeningTimes
		return ret
	}
	return *o.OpeningTimes
}

// GetOpeningTimesOk returns a tuple with the OpeningTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetOpeningTimesOk() (*LocationsDataOpeningTimes, bool) {
	if o == nil || IsNil(o.OpeningTimes) {
		return nil, false
	}
	return o.OpeningTimes, true
}

// HasOpeningTimes returns a boolean if a field has been set.
func (o *LocationsData) HasOpeningTimes() bool {
	if o != nil && !IsNil(o.OpeningTimes) {
		return true
	}

	return false
}

// SetOpeningTimes gets a reference to the given LocationsDataOpeningTimes and assigns it to the OpeningTimes field.
func (o *LocationsData) SetOpeningTimes(v LocationsDataOpeningTimes) {
	o.OpeningTimes = &v
}

// GetChargingWhenClosed returns the ChargingWhenClosed field value if set, zero value otherwise.
func (o *LocationsData) GetChargingWhenClosed() bool {
	if o == nil || IsNil(o.ChargingWhenClosed) {
		var ret bool
		return ret
	}
	return *o.ChargingWhenClosed
}

// GetChargingWhenClosedOk returns a tuple with the ChargingWhenClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetChargingWhenClosedOk() (*bool, bool) {
	if o == nil || IsNil(o.ChargingWhenClosed) {
		return nil, false
	}
	return o.ChargingWhenClosed, true
}

// HasChargingWhenClosed returns a boolean if a field has been set.
func (o *LocationsData) HasChargingWhenClosed() bool {
	if o != nil && !IsNil(o.ChargingWhenClosed) {
		return true
	}

	return false
}

// SetChargingWhenClosed gets a reference to the given bool and assigns it to the ChargingWhenClosed field.
func (o *LocationsData) SetChargingWhenClosed(v bool) {
	o.ChargingWhenClosed = &v
}

// GetImages returns the Images field value if set, zero value otherwise.
func (o *LocationsData) GetImages() Image {
	if o == nil || IsNil(o.Images) {
		var ret Image
		return ret
	}
	return *o.Images
}

// GetImagesOk returns a tuple with the Images field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetImagesOk() (*Image, bool) {
	if o == nil || IsNil(o.Images) {
		return nil, false
	}
	return o.Images, true
}

// HasImages returns a boolean if a field has been set.
func (o *LocationsData) HasImages() bool {
	if o != nil && !IsNil(o.Images) {
		return true
	}

	return false
}

// SetImages gets a reference to the given Image and assigns it to the Images field.
func (o *LocationsData) SetImages(v Image) {
	o.Images = &v
}

// GetEnergyMix returns the EnergyMix field value if set, zero value otherwise.
func (o *LocationsData) GetEnergyMix() CdrTariffsEnergyMix {
	if o == nil || IsNil(o.EnergyMix) {
		var ret CdrTariffsEnergyMix
		return ret
	}
	return *o.EnergyMix
}

// GetEnergyMixOk returns a tuple with the EnergyMix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationsData) GetEnergyMixOk() (*CdrTariffsEnergyMix, bool) {
	if o == nil || IsNil(o.EnergyMix) {
		return nil, false
	}
	return o.EnergyMix, true
}

// HasEnergyMix returns a boolean if a field has been set.
func (o *LocationsData) HasEnergyMix() bool {
	if o != nil && !IsNil(o.EnergyMix) {
		return true
	}

	return false
}

// SetEnergyMix gets a reference to the given CdrTariffsEnergyMix and assigns it to the EnergyMix field.
func (o *LocationsData) SetEnergyMix(v CdrTariffsEnergyMix) {
	o.EnergyMix = &v
}

// GetLastUpdated returns the LastUpdated field value
func (o *LocationsData) GetLastUpdated() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
func (o *LocationsData) GetLastUpdatedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastUpdated, true
}

// SetLastUpdated sets field value
func (o *LocationsData) SetLastUpdated(v string) {
	o.LastUpdated = v
}

func (o LocationsData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationsData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CountryCode) {
		toSerialize["country_code"] = o.CountryCode
	}
	if !IsNil(o.PartyId) {
		toSerialize["party_id"] = o.PartyId
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Publish) {
		toSerialize["publish"] = o.Publish
	}
	if !IsNil(o.PublishAllowedTo) {
		toSerialize["publish_allowed_to"] = o.PublishAllowedTo
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["address"] = o.Address
	toSerialize["city"] = o.City
	if !IsNil(o.PostalCode) {
		toSerialize["postal_code"] = o.PostalCode
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	toSerialize["country"] = o.Country
	toSerialize["coordinates"] = o.Coordinates
	if !IsNil(o.RelatedLocations) {
		toSerialize["related_locations"] = o.RelatedLocations
	}
	if !IsNil(o.ParkingType) {
		toSerialize["parking_type"] = o.ParkingType
	}
	if !IsNil(o.Evses) {
		toSerialize["evses"] = o.Evses
	}
	if !IsNil(o.Directions) {
		toSerialize["directions"] = o.Directions
	}
	if !IsNil(o.Operator) {
		toSerialize["operator"] = o.Operator
	}
	if !IsNil(o.Suboperator) {
		toSerialize["suboperator"] = o.Suboperator
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.Facilities) {
		toSerialize["facilities"] = o.Facilities
	}
	toSerialize["time_zone"] = o.TimeZone
	if !IsNil(o.OpeningTimes) {
		toSerialize["opening_times"] = o.OpeningTimes
	}
	if !IsNil(o.ChargingWhenClosed) {
		toSerialize["charging_when_closed"] = o.ChargingWhenClosed
	}
	if !IsNil(o.Images) {
		toSerialize["images"] = o.Images
	}
	if !IsNil(o.EnergyMix) {
		toSerialize["energy_mix"] = o.EnergyMix
	}
	toSerialize["last_updated"] = o.LastUpdated
	return toSerialize, nil
}

type NullableLocationsData struct {
	value *LocationsData
	isSet bool
}

func (v NullableLocationsData) Get() *LocationsData {
	return v.value
}

func (v *NullableLocationsData) Set(val *LocationsData) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationsData) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationsData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationsData(val *LocationsData) *NullableLocationsData {
	return &NullableLocationsData{value: val, isSet: true}
}

func (v NullableLocationsData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationsData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


