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

// CoreListServiceTypeResponse struct for CoreListServiceTypeResponse
type CoreListServiceTypeResponse struct {
	Type *[]CoreDescribeServiceTypeResponse `json:"type,omitempty"`
}

// NewCoreListServiceTypeResponse instantiates a new CoreListServiceTypeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreListServiceTypeResponse() *CoreListServiceTypeResponse {
	this := CoreListServiceTypeResponse{}
	return &this
}

// NewCoreListServiceTypeResponseWithDefaults instantiates a new CoreListServiceTypeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreListServiceTypeResponseWithDefaults() *CoreListServiceTypeResponse {
	this := CoreListServiceTypeResponse{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CoreListServiceTypeResponse) GetType() []CoreDescribeServiceTypeResponse {
	if o == nil || o.Type == nil {
		var ret []CoreDescribeServiceTypeResponse
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListServiceTypeResponse) GetTypeOk() (*[]CoreDescribeServiceTypeResponse, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CoreListServiceTypeResponse) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given []CoreDescribeServiceTypeResponse and assigns it to the Type field.
func (o *CoreListServiceTypeResponse) SetType(v []CoreDescribeServiceTypeResponse) {
	o.Type = &v
}

func (o CoreListServiceTypeResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCoreListServiceTypeResponse struct {
	value *CoreListServiceTypeResponse
	isSet bool
}

func (v NullableCoreListServiceTypeResponse) Get() *CoreListServiceTypeResponse {
	return v.value
}

func (v *NullableCoreListServiceTypeResponse) Set(val *CoreListServiceTypeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreListServiceTypeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreListServiceTypeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreListServiceTypeResponse(val *CoreListServiceTypeResponse) *NullableCoreListServiceTypeResponse {
	return &NullableCoreListServiceTypeResponse{value: val, isSet: true}
}

func (v NullableCoreListServiceTypeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreListServiceTypeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


