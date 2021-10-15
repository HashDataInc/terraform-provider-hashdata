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

// CoreDescribeWorkloadDetailResponse struct for CoreDescribeWorkloadDetailResponse
type CoreDescribeWorkloadDetailResponse struct {
	CommandId *int32 `json:"command_id,omitempty"`
	EpochId *int32 `json:"epoch_id,omitempty"`
	Id *string `json:"id,omitempty"`
	QueryState *string `json:"query_state,omitempty"`
	QueryText *string `json:"query_text,omitempty"`
	QueryType *string `json:"query_type,omitempty"`
	SegmentId *int32 `json:"segment_id,omitempty"`
	SenderId *int32 `json:"sender_id,omitempty"`
	Service *string `json:"service,omitempty"`
	SessionId *int32 `json:"session_id,omitempty"`
	SliceId *int32 `json:"slice_id,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewCoreDescribeWorkloadDetailResponse instantiates a new CoreDescribeWorkloadDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCoreDescribeWorkloadDetailResponse() *CoreDescribeWorkloadDetailResponse {
	this := CoreDescribeWorkloadDetailResponse{}
	return &this
}

// NewCoreDescribeWorkloadDetailResponseWithDefaults instantiates a new CoreDescribeWorkloadDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCoreDescribeWorkloadDetailResponseWithDefaults() *CoreDescribeWorkloadDetailResponse {
	this := CoreDescribeWorkloadDetailResponse{}
	return &this
}

// GetCommandId returns the CommandId field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetCommandId() int32 {
	if o == nil || o.CommandId == nil {
		var ret int32
		return ret
	}
	return *o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetCommandIdOk() (*int32, bool) {
	if o == nil || o.CommandId == nil {
		return nil, false
	}
	return o.CommandId, true
}

// HasCommandId returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasCommandId() bool {
	if o != nil && o.CommandId != nil {
		return true
	}

	return false
}

// SetCommandId gets a reference to the given int32 and assigns it to the CommandId field.
func (o *CoreDescribeWorkloadDetailResponse) SetCommandId(v int32) {
	o.CommandId = &v
}

// GetEpochId returns the EpochId field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetEpochId() int32 {
	if o == nil || o.EpochId == nil {
		var ret int32
		return ret
	}
	return *o.EpochId
}

// GetEpochIdOk returns a tuple with the EpochId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetEpochIdOk() (*int32, bool) {
	if o == nil || o.EpochId == nil {
		return nil, false
	}
	return o.EpochId, true
}

// HasEpochId returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasEpochId() bool {
	if o != nil && o.EpochId != nil {
		return true
	}

	return false
}

// SetEpochId gets a reference to the given int32 and assigns it to the EpochId field.
func (o *CoreDescribeWorkloadDetailResponse) SetEpochId(v int32) {
	o.EpochId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CoreDescribeWorkloadDetailResponse) SetId(v string) {
	o.Id = &v
}

// GetQueryState returns the QueryState field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetQueryState() string {
	if o == nil || o.QueryState == nil {
		var ret string
		return ret
	}
	return *o.QueryState
}

// GetQueryStateOk returns a tuple with the QueryState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetQueryStateOk() (*string, bool) {
	if o == nil || o.QueryState == nil {
		return nil, false
	}
	return o.QueryState, true
}

// HasQueryState returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasQueryState() bool {
	if o != nil && o.QueryState != nil {
		return true
	}

	return false
}

// SetQueryState gets a reference to the given string and assigns it to the QueryState field.
func (o *CoreDescribeWorkloadDetailResponse) SetQueryState(v string) {
	o.QueryState = &v
}

// GetQueryText returns the QueryText field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetQueryText() string {
	if o == nil || o.QueryText == nil {
		var ret string
		return ret
	}
	return *o.QueryText
}

// GetQueryTextOk returns a tuple with the QueryText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetQueryTextOk() (*string, bool) {
	if o == nil || o.QueryText == nil {
		return nil, false
	}
	return o.QueryText, true
}

// HasQueryText returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasQueryText() bool {
	if o != nil && o.QueryText != nil {
		return true
	}

	return false
}

// SetQueryText gets a reference to the given string and assigns it to the QueryText field.
func (o *CoreDescribeWorkloadDetailResponse) SetQueryText(v string) {
	o.QueryText = &v
}

// GetQueryType returns the QueryType field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetQueryType() string {
	if o == nil || o.QueryType == nil {
		var ret string
		return ret
	}
	return *o.QueryType
}

// GetQueryTypeOk returns a tuple with the QueryType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetQueryTypeOk() (*string, bool) {
	if o == nil || o.QueryType == nil {
		return nil, false
	}
	return o.QueryType, true
}

// HasQueryType returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasQueryType() bool {
	if o != nil && o.QueryType != nil {
		return true
	}

	return false
}

// SetQueryType gets a reference to the given string and assigns it to the QueryType field.
func (o *CoreDescribeWorkloadDetailResponse) SetQueryType(v string) {
	o.QueryType = &v
}

// GetSegmentId returns the SegmentId field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetSegmentId() int32 {
	if o == nil || o.SegmentId == nil {
		var ret int32
		return ret
	}
	return *o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetSegmentIdOk() (*int32, bool) {
	if o == nil || o.SegmentId == nil {
		return nil, false
	}
	return o.SegmentId, true
}

// HasSegmentId returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasSegmentId() bool {
	if o != nil && o.SegmentId != nil {
		return true
	}

	return false
}

// SetSegmentId gets a reference to the given int32 and assigns it to the SegmentId field.
func (o *CoreDescribeWorkloadDetailResponse) SetSegmentId(v int32) {
	o.SegmentId = &v
}

// GetSenderId returns the SenderId field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetSenderId() int32 {
	if o == nil || o.SenderId == nil {
		var ret int32
		return ret
	}
	return *o.SenderId
}

// GetSenderIdOk returns a tuple with the SenderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetSenderIdOk() (*int32, bool) {
	if o == nil || o.SenderId == nil {
		return nil, false
	}
	return o.SenderId, true
}

// HasSenderId returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasSenderId() bool {
	if o != nil && o.SenderId != nil {
		return true
	}

	return false
}

// SetSenderId gets a reference to the given int32 and assigns it to the SenderId field.
func (o *CoreDescribeWorkloadDetailResponse) SetSenderId(v int32) {
	o.SenderId = &v
}

// GetService returns the Service field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetService() string {
	if o == nil || o.Service == nil {
		var ret string
		return ret
	}
	return *o.Service
}

// GetServiceOk returns a tuple with the Service field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetServiceOk() (*string, bool) {
	if o == nil || o.Service == nil {
		return nil, false
	}
	return o.Service, true
}

// HasService returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasService() bool {
	if o != nil && o.Service != nil {
		return true
	}

	return false
}

// SetService gets a reference to the given string and assigns it to the Service field.
func (o *CoreDescribeWorkloadDetailResponse) SetService(v string) {
	o.Service = &v
}

// GetSessionId returns the SessionId field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetSessionId() int32 {
	if o == nil || o.SessionId == nil {
		var ret int32
		return ret
	}
	return *o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetSessionIdOk() (*int32, bool) {
	if o == nil || o.SessionId == nil {
		return nil, false
	}
	return o.SessionId, true
}

// HasSessionId returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasSessionId() bool {
	if o != nil && o.SessionId != nil {
		return true
	}

	return false
}

// SetSessionId gets a reference to the given int32 and assigns it to the SessionId field.
func (o *CoreDescribeWorkloadDetailResponse) SetSessionId(v int32) {
	o.SessionId = &v
}

// GetSliceId returns the SliceId field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetSliceId() int32 {
	if o == nil || o.SliceId == nil {
		var ret int32
		return ret
	}
	return *o.SliceId
}

// GetSliceIdOk returns a tuple with the SliceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetSliceIdOk() (*int32, bool) {
	if o == nil || o.SliceId == nil {
		return nil, false
	}
	return o.SliceId, true
}

// HasSliceId returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasSliceId() bool {
	if o != nil && o.SliceId != nil {
		return true
	}

	return false
}

// SetSliceId gets a reference to the given int32 and assigns it to the SliceId field.
func (o *CoreDescribeWorkloadDetailResponse) SetSliceId(v int32) {
	o.SliceId = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CoreDescribeWorkloadDetailResponse) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CoreDescribeWorkloadDetailResponse) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CoreDescribeWorkloadDetailResponse) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *CoreDescribeWorkloadDetailResponse) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o CoreDescribeWorkloadDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CommandId != nil {
		toSerialize["command_id"] = o.CommandId
	}
	if o.EpochId != nil {
		toSerialize["epoch_id"] = o.EpochId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.QueryState != nil {
		toSerialize["query_state"] = o.QueryState
	}
	if o.QueryText != nil {
		toSerialize["query_text"] = o.QueryText
	}
	if o.QueryType != nil {
		toSerialize["query_type"] = o.QueryType
	}
	if o.SegmentId != nil {
		toSerialize["segment_id"] = o.SegmentId
	}
	if o.SenderId != nil {
		toSerialize["sender_id"] = o.SenderId
	}
	if o.Service != nil {
		toSerialize["service"] = o.Service
	}
	if o.SessionId != nil {
		toSerialize["session_id"] = o.SessionId
	}
	if o.SliceId != nil {
		toSerialize["slice_id"] = o.SliceId
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableCoreDescribeWorkloadDetailResponse struct {
	value *CoreDescribeWorkloadDetailResponse
	isSet bool
}

func (v NullableCoreDescribeWorkloadDetailResponse) Get() *CoreDescribeWorkloadDetailResponse {
	return v.value
}

func (v *NullableCoreDescribeWorkloadDetailResponse) Set(val *CoreDescribeWorkloadDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreDescribeWorkloadDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreDescribeWorkloadDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreDescribeWorkloadDetailResponse(val *CoreDescribeWorkloadDetailResponse) *NullableCoreDescribeWorkloadDetailResponse {
	return &NullableCoreDescribeWorkloadDetailResponse{value: val, isSet: true}
}

func (v NullableCoreDescribeWorkloadDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreDescribeWorkloadDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

