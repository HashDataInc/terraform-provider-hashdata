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

// CommonProjectInfo struct for CommonProjectInfo
type CommonProjectInfo struct {
	ProjectId *string `json:"projectId,omitempty"`
	ProjectName *string `json:"projectName,omitempty"`
}

// NewCommonProjectInfo instantiates a new CommonProjectInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonProjectInfo() *CommonProjectInfo {
	this := CommonProjectInfo{}
	return &this
}

// NewCommonProjectInfoWithDefaults instantiates a new CommonProjectInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonProjectInfoWithDefaults() *CommonProjectInfo {
	this := CommonProjectInfo{}
	return &this
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *CommonProjectInfo) GetProjectId() string {
	if o == nil || o.ProjectId == nil {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonProjectInfo) GetProjectIdOk() (*string, bool) {
	if o == nil || o.ProjectId == nil {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *CommonProjectInfo) HasProjectId() bool {
	if o != nil && o.ProjectId != nil {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *CommonProjectInfo) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetProjectName returns the ProjectName field value if set, zero value otherwise.
func (o *CommonProjectInfo) GetProjectName() string {
	if o == nil || o.ProjectName == nil {
		var ret string
		return ret
	}
	return *o.ProjectName
}

// GetProjectNameOk returns a tuple with the ProjectName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonProjectInfo) GetProjectNameOk() (*string, bool) {
	if o == nil || o.ProjectName == nil {
		return nil, false
	}
	return o.ProjectName, true
}

// HasProjectName returns a boolean if a field has been set.
func (o *CommonProjectInfo) HasProjectName() bool {
	if o != nil && o.ProjectName != nil {
		return true
	}

	return false
}

// SetProjectName gets a reference to the given string and assigns it to the ProjectName field.
func (o *CommonProjectInfo) SetProjectName(v string) {
	o.ProjectName = &v
}

func (o CommonProjectInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ProjectId != nil {
		toSerialize["projectId"] = o.ProjectId
	}
	if o.ProjectName != nil {
		toSerialize["projectName"] = o.ProjectName
	}
	return json.Marshal(toSerialize)
}

type NullableCommonProjectInfo struct {
	value *CommonProjectInfo
	isSet bool
}

func (v NullableCommonProjectInfo) Get() *CommonProjectInfo {
	return v.value
}

func (v *NullableCommonProjectInfo) Set(val *CommonProjectInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonProjectInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonProjectInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonProjectInfo(val *CommonProjectInfo) *NullableCommonProjectInfo {
	return &NullableCommonProjectInfo{value: val, isSet: true}
}

func (v NullableCommonProjectInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonProjectInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


