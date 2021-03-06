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

// AccountDescribeTeamOfUserResponse struct for AccountDescribeTeamOfUserResponse
type AccountDescribeTeamOfUserResponse struct {
	CreatedAt *string `json:"created_at,omitempty"`
	DestroyedAt *string `json:"destroyed_at,omitempty"`
	Id *string `json:"id,omitempty"`
	IsMaintainer *bool `json:"is_maintainer,omitempty"`
	Name *string `json:"name,omitempty"`
	Tenant *string `json:"tenant,omitempty"`
}

// NewAccountDescribeTeamOfUserResponse instantiates a new AccountDescribeTeamOfUserResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountDescribeTeamOfUserResponse() *AccountDescribeTeamOfUserResponse {
	this := AccountDescribeTeamOfUserResponse{}
	return &this
}

// NewAccountDescribeTeamOfUserResponseWithDefaults instantiates a new AccountDescribeTeamOfUserResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountDescribeTeamOfUserResponseWithDefaults() *AccountDescribeTeamOfUserResponse {
	this := AccountDescribeTeamOfUserResponse{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AccountDescribeTeamOfUserResponse) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDescribeTeamOfUserResponse) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AccountDescribeTeamOfUserResponse) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *AccountDescribeTeamOfUserResponse) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDestroyedAt returns the DestroyedAt field value if set, zero value otherwise.
func (o *AccountDescribeTeamOfUserResponse) GetDestroyedAt() string {
	if o == nil || o.DestroyedAt == nil {
		var ret string
		return ret
	}
	return *o.DestroyedAt
}

// GetDestroyedAtOk returns a tuple with the DestroyedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDescribeTeamOfUserResponse) GetDestroyedAtOk() (*string, bool) {
	if o == nil || o.DestroyedAt == nil {
		return nil, false
	}
	return o.DestroyedAt, true
}

// HasDestroyedAt returns a boolean if a field has been set.
func (o *AccountDescribeTeamOfUserResponse) HasDestroyedAt() bool {
	if o != nil && o.DestroyedAt != nil {
		return true
	}

	return false
}

// SetDestroyedAt gets a reference to the given string and assigns it to the DestroyedAt field.
func (o *AccountDescribeTeamOfUserResponse) SetDestroyedAt(v string) {
	o.DestroyedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccountDescribeTeamOfUserResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDescribeTeamOfUserResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccountDescribeTeamOfUserResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccountDescribeTeamOfUserResponse) SetId(v string) {
	o.Id = &v
}

// GetIsMaintainer returns the IsMaintainer field value if set, zero value otherwise.
func (o *AccountDescribeTeamOfUserResponse) GetIsMaintainer() bool {
	if o == nil || o.IsMaintainer == nil {
		var ret bool
		return ret
	}
	return *o.IsMaintainer
}

// GetIsMaintainerOk returns a tuple with the IsMaintainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDescribeTeamOfUserResponse) GetIsMaintainerOk() (*bool, bool) {
	if o == nil || o.IsMaintainer == nil {
		return nil, false
	}
	return o.IsMaintainer, true
}

// HasIsMaintainer returns a boolean if a field has been set.
func (o *AccountDescribeTeamOfUserResponse) HasIsMaintainer() bool {
	if o != nil && o.IsMaintainer != nil {
		return true
	}

	return false
}

// SetIsMaintainer gets a reference to the given bool and assigns it to the IsMaintainer field.
func (o *AccountDescribeTeamOfUserResponse) SetIsMaintainer(v bool) {
	o.IsMaintainer = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccountDescribeTeamOfUserResponse) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDescribeTeamOfUserResponse) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccountDescribeTeamOfUserResponse) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccountDescribeTeamOfUserResponse) SetName(v string) {
	o.Name = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise.
func (o *AccountDescribeTeamOfUserResponse) GetTenant() string {
	if o == nil || o.Tenant == nil {
		var ret string
		return ret
	}
	return *o.Tenant
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountDescribeTeamOfUserResponse) GetTenantOk() (*string, bool) {
	if o == nil || o.Tenant == nil {
		return nil, false
	}
	return o.Tenant, true
}

// HasTenant returns a boolean if a field has been set.
func (o *AccountDescribeTeamOfUserResponse) HasTenant() bool {
	if o != nil && o.Tenant != nil {
		return true
	}

	return false
}

// SetTenant gets a reference to the given string and assigns it to the Tenant field.
func (o *AccountDescribeTeamOfUserResponse) SetTenant(v string) {
	o.Tenant = &v
}

func (o AccountDescribeTeamOfUserResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DestroyedAt != nil {
		toSerialize["destroyed_at"] = o.DestroyedAt
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IsMaintainer != nil {
		toSerialize["is_maintainer"] = o.IsMaintainer
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Tenant != nil {
		toSerialize["tenant"] = o.Tenant
	}
	return json.Marshal(toSerialize)
}

type NullableAccountDescribeTeamOfUserResponse struct {
	value *AccountDescribeTeamOfUserResponse
	isSet bool
}

func (v NullableAccountDescribeTeamOfUserResponse) Get() *AccountDescribeTeamOfUserResponse {
	return v.value
}

func (v *NullableAccountDescribeTeamOfUserResponse) Set(val *AccountDescribeTeamOfUserResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountDescribeTeamOfUserResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountDescribeTeamOfUserResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountDescribeTeamOfUserResponse(val *AccountDescribeTeamOfUserResponse) *NullableAccountDescribeTeamOfUserResponse {
	return &NullableAccountDescribeTeamOfUserResponse{value: val, isSet: true}
}

func (v NullableAccountDescribeTeamOfUserResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountDescribeTeamOfUserResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


