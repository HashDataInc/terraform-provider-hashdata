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

// OpsCountCloudmgrComponentResponse struct for OpsCountCloudmgrComponentResponse
type OpsCountCloudmgrComponentResponse struct {
	Health *int32 `json:"health,omitempty"`
	Sum *int32 `json:"sum,omitempty"`
	Unhealth *int32 `json:"unhealth,omitempty"`
}

// NewOpsCountCloudmgrComponentResponse instantiates a new OpsCountCloudmgrComponentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsCountCloudmgrComponentResponse() *OpsCountCloudmgrComponentResponse {
	this := OpsCountCloudmgrComponentResponse{}
	return &this
}

// NewOpsCountCloudmgrComponentResponseWithDefaults instantiates a new OpsCountCloudmgrComponentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsCountCloudmgrComponentResponseWithDefaults() *OpsCountCloudmgrComponentResponse {
	this := OpsCountCloudmgrComponentResponse{}
	return &this
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *OpsCountCloudmgrComponentResponse) GetHealth() int32 {
	if o == nil || o.Health == nil {
		var ret int32
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCountCloudmgrComponentResponse) GetHealthOk() (*int32, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *OpsCountCloudmgrComponentResponse) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given int32 and assigns it to the Health field.
func (o *OpsCountCloudmgrComponentResponse) SetHealth(v int32) {
	o.Health = &v
}

// GetSum returns the Sum field value if set, zero value otherwise.
func (o *OpsCountCloudmgrComponentResponse) GetSum() int32 {
	if o == nil || o.Sum == nil {
		var ret int32
		return ret
	}
	return *o.Sum
}

// GetSumOk returns a tuple with the Sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCountCloudmgrComponentResponse) GetSumOk() (*int32, bool) {
	if o == nil || o.Sum == nil {
		return nil, false
	}
	return o.Sum, true
}

// HasSum returns a boolean if a field has been set.
func (o *OpsCountCloudmgrComponentResponse) HasSum() bool {
	if o != nil && o.Sum != nil {
		return true
	}

	return false
}

// SetSum gets a reference to the given int32 and assigns it to the Sum field.
func (o *OpsCountCloudmgrComponentResponse) SetSum(v int32) {
	o.Sum = &v
}

// GetUnhealth returns the Unhealth field value if set, zero value otherwise.
func (o *OpsCountCloudmgrComponentResponse) GetUnhealth() int32 {
	if o == nil || o.Unhealth == nil {
		var ret int32
		return ret
	}
	return *o.Unhealth
}

// GetUnhealthOk returns a tuple with the Unhealth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCountCloudmgrComponentResponse) GetUnhealthOk() (*int32, bool) {
	if o == nil || o.Unhealth == nil {
		return nil, false
	}
	return o.Unhealth, true
}

// HasUnhealth returns a boolean if a field has been set.
func (o *OpsCountCloudmgrComponentResponse) HasUnhealth() bool {
	if o != nil && o.Unhealth != nil {
		return true
	}

	return false
}

// SetUnhealth gets a reference to the given int32 and assigns it to the Unhealth field.
func (o *OpsCountCloudmgrComponentResponse) SetUnhealth(v int32) {
	o.Unhealth = &v
}

func (o OpsCountCloudmgrComponentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Health != nil {
		toSerialize["health"] = o.Health
	}
	if o.Sum != nil {
		toSerialize["sum"] = o.Sum
	}
	if o.Unhealth != nil {
		toSerialize["unhealth"] = o.Unhealth
	}
	return json.Marshal(toSerialize)
}

type NullableOpsCountCloudmgrComponentResponse struct {
	value *OpsCountCloudmgrComponentResponse
	isSet bool
}

func (v NullableOpsCountCloudmgrComponentResponse) Get() *OpsCountCloudmgrComponentResponse {
	return v.value
}

func (v *NullableOpsCountCloudmgrComponentResponse) Set(val *OpsCountCloudmgrComponentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsCountCloudmgrComponentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsCountCloudmgrComponentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsCountCloudmgrComponentResponse(val *OpsCountCloudmgrComponentResponse) *NullableOpsCountCloudmgrComponentResponse {
	return &NullableOpsCountCloudmgrComponentResponse{value: val, isSet: true}
}

func (v NullableOpsCountCloudmgrComponentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsCountCloudmgrComponentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


