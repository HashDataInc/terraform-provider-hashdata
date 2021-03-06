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

// AccountCreateTeamRequest struct for AccountCreateTeamRequest
type AccountCreateTeamRequest struct {
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewAccountCreateTeamRequest instantiates a new AccountCreateTeamRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreateTeamRequest() *AccountCreateTeamRequest {
	this := AccountCreateTeamRequest{}
	return &this
}

// NewAccountCreateTeamRequestWithDefaults instantiates a new AccountCreateTeamRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreateTeamRequestWithDefaults() *AccountCreateTeamRequest {
	this := AccountCreateTeamRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AccountCreateTeamRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateTeamRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AccountCreateTeamRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AccountCreateTeamRequest) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountCreateTeamRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateTeamRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountCreateTeamRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountCreateTeamRequest) SetName(v string) {
	o.Name = &v
}

func (o AccountCreateTeamRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableAccountCreateTeamRequest struct {
	value *AccountCreateTeamRequest
	isSet bool
}

func (v NullableAccountCreateTeamRequest) Get() *AccountCreateTeamRequest {
	return v.value
}

func (v *NullableAccountCreateTeamRequest) Set(val *AccountCreateTeamRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreateTeamRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreateTeamRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreateTeamRequest(val *AccountCreateTeamRequest) *NullableAccountCreateTeamRequest {
	return &NullableAccountCreateTeamRequest{value: val, isSet: true}
}

func (v NullableAccountCreateTeamRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreateTeamRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


