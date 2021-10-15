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

// OpsCreateTransactionResponse struct for OpsCreateTransactionResponse
type OpsCreateTransactionResponse struct {
	Amount *string `json:"amount,omitempty"`
	Id *string `json:"id,omitempty"`
	OrderNum *string `json:"order_num,omitempty"`
	QrCode *string `json:"qr_code,omitempty"`
}

// NewOpsCreateTransactionResponse instantiates a new OpsCreateTransactionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpsCreateTransactionResponse() *OpsCreateTransactionResponse {
	this := OpsCreateTransactionResponse{}
	return &this
}

// NewOpsCreateTransactionResponseWithDefaults instantiates a new OpsCreateTransactionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpsCreateTransactionResponseWithDefaults() *OpsCreateTransactionResponse {
	this := OpsCreateTransactionResponse{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *OpsCreateTransactionResponse) GetAmount() string {
	if o == nil || o.Amount == nil {
		var ret string
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCreateTransactionResponse) GetAmountOk() (*string, bool) {
	if o == nil || o.Amount == nil {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *OpsCreateTransactionResponse) HasAmount() bool {
	if o != nil && o.Amount != nil {
		return true
	}

	return false
}

// SetAmount gets a reference to the given string and assigns it to the Amount field.
func (o *OpsCreateTransactionResponse) SetAmount(v string) {
	o.Amount = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OpsCreateTransactionResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCreateTransactionResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OpsCreateTransactionResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OpsCreateTransactionResponse) SetId(v string) {
	o.Id = &v
}

// GetOrderNum returns the OrderNum field value if set, zero value otherwise.
func (o *OpsCreateTransactionResponse) GetOrderNum() string {
	if o == nil || o.OrderNum == nil {
		var ret string
		return ret
	}
	return *o.OrderNum
}

// GetOrderNumOk returns a tuple with the OrderNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCreateTransactionResponse) GetOrderNumOk() (*string, bool) {
	if o == nil || o.OrderNum == nil {
		return nil, false
	}
	return o.OrderNum, true
}

// HasOrderNum returns a boolean if a field has been set.
func (o *OpsCreateTransactionResponse) HasOrderNum() bool {
	if o != nil && o.OrderNum != nil {
		return true
	}

	return false
}

// SetOrderNum gets a reference to the given string and assigns it to the OrderNum field.
func (o *OpsCreateTransactionResponse) SetOrderNum(v string) {
	o.OrderNum = &v
}

// GetQrCode returns the QrCode field value if set, zero value otherwise.
func (o *OpsCreateTransactionResponse) GetQrCode() string {
	if o == nil || o.QrCode == nil {
		var ret string
		return ret
	}
	return *o.QrCode
}

// GetQrCodeOk returns a tuple with the QrCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpsCreateTransactionResponse) GetQrCodeOk() (*string, bool) {
	if o == nil || o.QrCode == nil {
		return nil, false
	}
	return o.QrCode, true
}

// HasQrCode returns a boolean if a field has been set.
func (o *OpsCreateTransactionResponse) HasQrCode() bool {
	if o != nil && o.QrCode != nil {
		return true
	}

	return false
}

// SetQrCode gets a reference to the given string and assigns it to the QrCode field.
func (o *OpsCreateTransactionResponse) SetQrCode(v string) {
	o.QrCode = &v
}

func (o OpsCreateTransactionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Amount != nil {
		toSerialize["amount"] = o.Amount
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.OrderNum != nil {
		toSerialize["order_num"] = o.OrderNum
	}
	if o.QrCode != nil {
		toSerialize["qr_code"] = o.QrCode
	}
	return json.Marshal(toSerialize)
}

type NullableOpsCreateTransactionResponse struct {
	value *OpsCreateTransactionResponse
	isSet bool
}

func (v NullableOpsCreateTransactionResponse) Get() *OpsCreateTransactionResponse {
	return v.value
}

func (v *NullableOpsCreateTransactionResponse) Set(val *OpsCreateTransactionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableOpsCreateTransactionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableOpsCreateTransactionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpsCreateTransactionResponse(val *OpsCreateTransactionResponse) *NullableOpsCreateTransactionResponse {
	return &NullableOpsCreateTransactionResponse{value: val, isSet: true}
}

func (v NullableOpsCreateTransactionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpsCreateTransactionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

