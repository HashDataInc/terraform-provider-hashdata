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

// CoreDescribeInternetResponse struct for CoreDescribeInternetResponse
type CoreDescribeInternetResponse struct {
	Bandwidth *int32 `json:"bandwidth,omitempty"`
	ElasticIp *bool `json:"elastic_ip,omitempty"`
	ElasticIpId *string `json:"elastic_ip_id,omitempty"`
	PublicIp *string `json:"public_ip,omitempty"`
}

// NewCoreDescribeInternetResponse instantiates a new CoreDescribeInternetResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDescribeInternetResponse() *CoreDescribeInternetResponse {
	this := CoreDescribeInternetResponse{}
	return &this
}

// NewCoreDescribeInternetResponseWithDefaults instantiates a new CoreDescribeInternetResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDescribeInternetResponseWithDefaults() *CoreDescribeInternetResponse {
	this := CoreDescribeInternetResponse{}
	return &this
}

// GetBandwidth returns the Bandwidth field value if set, zero value otherwise.
func (o *CoreDescribeInternetResponse) GetBandwidth() int32 {
	if o == nil || o.Bandwidth == nil {
		var ret int32
		return ret
	}
	return *o.Bandwidth
}

// GetBandwidthOk returns a tuple with the Bandwidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeInternetResponse) GetBandwidthOk() (*int32, bool) {
	if o == nil || o.Bandwidth == nil {
		return nil, false
	}
	return o.Bandwidth, true
}

// HasBandwidth returns a boolean if a field has been set.
func (o *CoreDescribeInternetResponse) HasBandwidth() bool {
	if o != nil && o.Bandwidth != nil {
		return true
	}

	return false
}

// SetBandwidth gets a reference to the given int32 and assigns it to the Bandwidth field.
func (o *CoreDescribeInternetResponse) SetBandwidth(v int32) {
	o.Bandwidth = &v
}

// GetElasticIp returns the ElasticIp field value if set, zero value otherwise.
func (o *CoreDescribeInternetResponse) GetElasticIp() bool {
	if o == nil || o.ElasticIp == nil {
		var ret bool
		return ret
	}
	return *o.ElasticIp
}

// GetElasticIpOk returns a tuple with the ElasticIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeInternetResponse) GetElasticIpOk() (*bool, bool) {
	if o == nil || o.ElasticIp == nil {
		return nil, false
	}
	return o.ElasticIp, true
}

// HasElasticIp returns a boolean if a field has been set.
func (o *CoreDescribeInternetResponse) HasElasticIp() bool {
	if o != nil && o.ElasticIp != nil {
		return true
	}

	return false
}

// SetElasticIp gets a reference to the given bool and assigns it to the ElasticIp field.
func (o *CoreDescribeInternetResponse) SetElasticIp(v bool) {
	o.ElasticIp = &v
}

// GetElasticIpId returns the ElasticIpId field value if set, zero value otherwise.
func (o *CoreDescribeInternetResponse) GetElasticIpId() string {
	if o == nil || o.ElasticIpId == nil {
		var ret string
		return ret
	}
	return *o.ElasticIpId
}

// GetElasticIpIdOk returns a tuple with the ElasticIpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeInternetResponse) GetElasticIpIdOk() (*string, bool) {
	if o == nil || o.ElasticIpId == nil {
		return nil, false
	}
	return o.ElasticIpId, true
}

// HasElasticIpId returns a boolean if a field has been set.
func (o *CoreDescribeInternetResponse) HasElasticIpId() bool {
	if o != nil && o.ElasticIpId != nil {
		return true
	}

	return false
}

// SetElasticIpId gets a reference to the given string and assigns it to the ElasticIpId field.
func (o *CoreDescribeInternetResponse) SetElasticIpId(v string) {
	o.ElasticIpId = &v
}

// GetPublicIp returns the PublicIp field value if set, zero value otherwise.
func (o *CoreDescribeInternetResponse) GetPublicIp() string {
	if o == nil || o.PublicIp == nil {
		var ret string
		return ret
	}
	return *o.PublicIp
}

// GetPublicIpOk returns a tuple with the PublicIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeInternetResponse) GetPublicIpOk() (*string, bool) {
	if o == nil || o.PublicIp == nil {
		return nil, false
	}
	return o.PublicIp, true
}

// HasPublicIp returns a boolean if a field has been set.
func (o *CoreDescribeInternetResponse) HasPublicIp() bool {
	if o != nil && o.PublicIp != nil {
		return true
	}

	return false
}

// SetPublicIp gets a reference to the given string and assigns it to the PublicIp field.
func (o *CoreDescribeInternetResponse) SetPublicIp(v string) {
	o.PublicIp = &v
}

func (o CoreDescribeInternetResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Bandwidth != nil {
		toSerialize["bandwidth"] = o.Bandwidth
	}
	if o.ElasticIp != nil {
		toSerialize["elastic_ip"] = o.ElasticIp
	}
	if o.ElasticIpId != nil {
		toSerialize["elastic_ip_id"] = o.ElasticIpId
	}
	if o.PublicIp != nil {
		toSerialize["public_ip"] = o.PublicIp
	}
	return json.Marshal(toSerialize)
}

type NullableCoreDescribeInternetResponse struct {
	value *CoreDescribeInternetResponse
	isSet bool
}

func (v NullableCoreDescribeInternetResponse) Get() *CoreDescribeInternetResponse {
	return v.value
}

func (v *NullableCoreDescribeInternetResponse) Set(val *CoreDescribeInternetResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDescribeInternetResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDescribeInternetResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDescribeInternetResponse(val *CoreDescribeInternetResponse) *NullableCoreDescribeInternetResponse {
	return &NullableCoreDescribeInternetResponse{value: val, isSet: true}
}

func (v NullableCoreDescribeInternetResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDescribeInternetResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


