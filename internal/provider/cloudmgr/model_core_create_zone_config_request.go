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

// CoreCreateZoneConfigRequest struct for CoreCreateZoneConfigRequest
type CoreCreateZoneConfigRequest struct {
	Tenant *string `json:"tenant,omitempty"`
	Vxnet *string `json:"vxnet,omitempty"`
	Zone *string `json:"zone,omitempty"`
}

// NewCoreCreateZoneConfigRequest instantiates a new CoreCreateZoneConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreCreateZoneConfigRequest() *CoreCreateZoneConfigRequest {
	this := CoreCreateZoneConfigRequest{}
	return &this
}

// NewCoreCreateZoneConfigRequestWithDefaults instantiates a new CoreCreateZoneConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreCreateZoneConfigRequestWithDefaults() *CoreCreateZoneConfigRequest {
	this := CoreCreateZoneConfigRequest{}
	return &this
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *CoreCreateZoneConfigRequest) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneConfigRequest) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *CoreCreateZoneConfigRequest) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *CoreCreateZoneConfigRequest) SetTenant(v string) {
	o.Tenant = &v
}

// GetVxnet returns the Vxnet field value if set, zero value otherwise.
func (o *CoreCreateZoneConfigRequest) GetVxnet() string {
	if o == nil || o.Vxnet == nil {
		var ret string
		return ret
	}
	return *o.Vxnet
}

// GetVxnetOk returns a tuple with the Vxnet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneConfigRequest) GetVxnetOk() (*string, bool) {
	if o == nil || o.Vxnet == nil {
		return nil, false
	}
	return o.Vxnet, true
}

// HasVxnet returns a boolean if a field has been set.
func (o *CoreCreateZoneConfigRequest) HasVxnet() bool {
	if o != nil && o.Vxnet != nil {
		return true
	}

	return false
}

// SetVxnet gets a reference to the given string and assigns it to the Vxnet field.
func (o *CoreCreateZoneConfigRequest) SetVxnet(v string) {
	o.Vxnet = &v
}

// GetZone returns the Zone field value if set, zero value otherwise.
func (o *CoreCreateZoneConfigRequest) GetZone() string {
	if o == nil || o.Zone == nil {
		var ret string
		return ret
	}
	return *o.Zone
}

// GetZoneOk returns a tuple with the Zone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreCreateZoneConfigRequest) GetZoneOk() (*string, bool) {
	if o == nil || o.Zone == nil {
		return nil, false
	}
	return o.Zone, true
}

// HasZone returns a boolean if a field has been set.
func (o *CoreCreateZoneConfigRequest) HasZone() bool {
	if o != nil && o.Zone != nil {
		return true
	}

	return false
}

// SetZone gets a reference to the given string and assigns it to the Zone field.
func (o *CoreCreateZoneConfigRequest) SetZone(v string) {
	o.Zone = &v
}

func (o CoreCreateZoneConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	if o.Vxnet != nil {
		toSerialize["vxnet"] = o.Vxnet
	}
	if o.Zone != nil {
		toSerialize["zone"] = o.Zone
	}
	return json.Marshal(toSerialize)
}

type NullableCoreCreateZoneConfigRequest struct {
	value *CoreCreateZoneConfigRequest
	isSet bool
}

func (v NullableCoreCreateZoneConfigRequest) Get() *CoreCreateZoneConfigRequest {
	return v.value
}

func (v *NullableCoreCreateZoneConfigRequest) Set(val *CoreCreateZoneConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreCreateZoneConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreCreateZoneConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreCreateZoneConfigRequest(val *CoreCreateZoneConfigRequest) *NullableCoreCreateZoneConfigRequest {
	return &NullableCoreCreateZoneConfigRequest{value: val, isSet: true}
}

func (v NullableCoreCreateZoneConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreCreateZoneConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


