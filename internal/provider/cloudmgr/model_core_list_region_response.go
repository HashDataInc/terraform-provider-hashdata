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

// CoreListRegionResponse struct for CoreListRegionResponse
type CoreListRegionResponse struct {
	Content *[]CoreDescribeRegionResponse `json:"content,omitempty"`
	Count *int32 `json:"count,omitempty"`
	Total *int32 `json:"total,omitempty"`
}

// NewCoreListRegionResponse instantiates a new CoreListRegionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreListRegionResponse() *CoreListRegionResponse {
	this := CoreListRegionResponse{}
	return &this
}

// NewCoreListRegionResponseWithDefaults instantiates a new CoreListRegionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreListRegionResponseWithDefaults() *CoreListRegionResponse {
	this := CoreListRegionResponse{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *CoreListRegionResponse) GetContent() []CoreDescribeRegionResponse {
	if o == nil || o.Content == nil {
		var ret []CoreDescribeRegionResponse
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListRegionResponse) GetContentOk() (*[]CoreDescribeRegionResponse, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *CoreListRegionResponse) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given []CoreDescribeRegionResponse and assigns it to the Content field.
func (o *CoreListRegionResponse) SetContent(v []CoreDescribeRegionResponse) {
	o.Content = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *CoreListRegionResponse) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListRegionResponse) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *CoreListRegionResponse) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *CoreListRegionResponse) SetCount(v int32) {
	o.Count = &v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *CoreListRegionResponse) GetTotal() int32 {
	if o == nil || o.Total == nil {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreListRegionResponse) GetTotalOk() (*int32, bool) {
	if o == nil || o.Total == nil {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *CoreListRegionResponse) HasTotal() bool {
	if o != nil && o.Total != nil {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *CoreListRegionResponse) SetTotal(v int32) {
	o.Total = &v
}

func (o CoreListRegionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Total != nil {
		toSerialize["total"] = o.Total
	}
	return json.Marshal(toSerialize)
}

type NullableCoreListRegionResponse struct {
	value *CoreListRegionResponse
	isSet bool
}

func (v NullableCoreListRegionResponse) Get() *CoreListRegionResponse {
	return v.value
}

func (v *NullableCoreListRegionResponse) Set(val *CoreListRegionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreListRegionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreListRegionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreListRegionResponse(val *CoreListRegionResponse) *NullableCoreListRegionResponse {
	return &NullableCoreListRegionResponse{value: val, isSet: true}
}

func (v NullableCoreListRegionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreListRegionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

