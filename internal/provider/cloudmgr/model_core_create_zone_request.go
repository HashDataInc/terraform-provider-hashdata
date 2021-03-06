/*
 * Cloud Manager API
 *
 * Cloud Manager Restful API Documentation.
 *
 * API version: v2.0
 * Contact: wang@hashdata.cn
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudmgr

import (
	"encoding/json"
)

// CoreCreateZoneRequest struct for CoreCreateZoneRequest
type CoreCreateZoneRequest struct {
	Description *string `json:"description,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	Name *string `json:"name,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
	Region *string `json:"region,omitempty"`
	Tenant *string `json:"tenant,omitempty"`
	Vendor *string `json:"vendor,omitempty"`
	Zone *string `json:"zone,omitempty"`
}

// NewCoreCreateZoneRequest instantiates a new CoreCreateZoneRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCreateZoneRequest() *CoreCreateZoneRequest {
	this := CoreCreateZoneRequest{}
	return &this
}

// NewCoreCreateZoneRequestWithDefaults instantiates a new CoreCreateZoneRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCreateZoneRequestWithDefaults() *CoreCreateZoneRequest {
	this := CoreCreateZoneRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CoreCreateZoneRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CoreCreateZoneRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CoreCreateZoneRequest) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *CoreCreateZoneRequest) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneRequest) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *CoreCreateZoneRequest) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *CoreCreateZoneRequest) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CoreCreateZoneRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CoreCreateZoneRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CoreCreateZoneRequest) SetName(v string) {
	o.Name = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *CoreCreateZoneRequest) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneRequest) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *CoreCreateZoneRequest) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *CoreCreateZoneRequest) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *CoreCreateZoneRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *CoreCreateZoneRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *CoreCreateZoneRequest) SetRegion(v string) {
	o.Region = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *CoreCreateZoneRequest) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneRequest) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *CoreCreateZoneRequest) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *CoreCreateZoneRequest) SetTenant(v string) {
	o.Tenant = &v
}

// GetVendor returns the Vendor field value if set, zero value otherwise.
func (o *CoreCreateZoneRequest) GetVendor() string {
	if o == nil || o.Vendor == nil {
		var ret string
		return ret
	}
	return *o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneRequest) GetVendorOk() (*string, bool) {
	if o == nil || o.Vendor == nil {
		return nil, false
	}
	return o.Vendor, true
}

// HasVendor returns a boolean if a field has been set.
func (o *CoreCreateZoneRequest) HasVendor() bool {
	if o != nil && o.Vendor != nil {
		return true
	}

	return false
}

// SetVendor gets a reference to the given string and assigns it to the Vendor field.
func (o *CoreCreateZoneRequest) SetVendor(v string) {
	o.Vendor = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *CoreCreateZoneRequest) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneRequest) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *CoreCreateZoneRequest) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *CoreCreateZoneRequest) SetZone(v string) {
	o.Zone = &v
}

func (o CoreCreateZoneRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.Vendor != nil {
		toSerialize["vendor"] = o.Vendor
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	return json.Marshal(toSerialize)
}

type NullableCoreCreateZoneRequest struct {
	value *CoreCreateZoneRequest
	isSet bool
}

func (v NullableCoreCreateZoneRequest) Get() *CoreCreateZoneRequest {
	return v.value
}

func (v *NullableCoreCreateZoneRequest) Set(val *CoreCreateZoneRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCreateZoneRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCreateZoneRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCreateZoneRequest(val *CoreCreateZoneRequest) *NullableCoreCreateZoneRequest {
	return &NullableCoreCreateZoneRequest{value: val, isSet: true}
}

func (v NullableCoreCreateZoneRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCreateZoneRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


