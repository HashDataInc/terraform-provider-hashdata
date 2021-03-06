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

// CoreWorkloadTagResponse struct for CoreWorkloadTagResponse
type CoreWorkloadTagResponse struct {
	Id *string `json:"id,omitempty"`
	Tags *map[string]interface{} `json:"tags,omitempty"`
}

// NewCoreWorkloadTagResponse instantiates a new CoreWorkloadTagResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreWorkloadTagResponse() *CoreWorkloadTagResponse {
	this := CoreWorkloadTagResponse{}
	return &this
}

// NewCoreWorkloadTagResponseWithDefaults instantiates a new CoreWorkloadTagResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreWorkloadTagResponseWithDefaults() *CoreWorkloadTagResponse {
	this := CoreWorkloadTagResponse{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreWorkloadTagResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWorkloadTagResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreWorkloadTagResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoreWorkloadTagResponse) SetId(v string) {
	o.Id = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *CoreWorkloadTagResponse) GetTags() map[string]interface{} {
	if o == nil || o.Tags == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreWorkloadTagResponse) GetTagsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *CoreWorkloadTagResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given map[string]interface{} and assigns it to the Tags field.
func (o *CoreWorkloadTagResponse) SetTags(v map[string]interface{}) {
	o.Tags = &v
}

func (o CoreWorkloadTagResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableCoreWorkloadTagResponse struct {
	value *CoreWorkloadTagResponse
	isSet bool
}

func (v NullableCoreWorkloadTagResponse) Get() *CoreWorkloadTagResponse {
	return v.value
}

func (v *NullableCoreWorkloadTagResponse) Set(val *CoreWorkloadTagResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreWorkloadTagResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreWorkloadTagResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreWorkloadTagResponse(val *CoreWorkloadTagResponse) *NullableCoreWorkloadTagResponse {
	return &NullableCoreWorkloadTagResponse{value: val, isSet: true}
}

func (v NullableCoreWorkloadTagResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreWorkloadTagResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


