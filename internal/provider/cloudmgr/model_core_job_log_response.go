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

// CoreJobLogResponse struct for CoreJobLogResponse
type CoreJobLogResponse struct {
	Content *string `json:"content,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewCoreJobLogResponse instantiates a new CoreJobLogResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreJobLogResponse() *CoreJobLogResponse {
	this := CoreJobLogResponse{}
	return &this
}

// NewCoreJobLogResponseWithDefaults instantiates a new CoreJobLogResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreJobLogResponseWithDefaults() *CoreJobLogResponse {
	this := CoreJobLogResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CoreJobLogResponse) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreJobLogResponse) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CoreJobLogResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *CoreJobLogResponse) SetContent(v string) {
	o.Content = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreJobLogResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreJobLogResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreJobLogResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoreJobLogResponse) SetId(v string) {
	o.Id = &v
}

func (o CoreJobLogResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCoreJobLogResponse struct {
	value *CoreJobLogResponse
	isSet bool
}

func (v NullableCoreJobLogResponse) Get() *CoreJobLogResponse {
	return v.value
}

func (v *NullableCoreJobLogResponse) Set(val *CoreJobLogResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreJobLogResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreJobLogResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreJobLogResponse(val *CoreJobLogResponse) *NullableCoreJobLogResponse {
	return &NullableCoreJobLogResponse{value: val, isSet: true}
}

func (v NullableCoreJobLogResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreJobLogResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


