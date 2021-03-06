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

// CloudmgrcommonIaasResource struct for CloudmgrcommonIaasResource
type CloudmgrcommonIaasResource struct {
	Az *string `json:"az,omitempty"`
	ClusterName *string `json:"clusterName,omitempty"`
	ClusterType *string `json:"clusterType,omitempty"`
	Ip *string `json:"ip,omitempty"`
	MachineTypeCode *string `json:"machineTypeCode,omitempty"`
	MachineTypeName *string `json:"machineTypeName,omitempty"`
	ProductTypeCode *string `json:"productTypeCode,omitempty"`
	ProjectList *[]CommonProjectInfo `json:"projectList,omitempty"`
	RegionCode *string `json:"regionCode,omitempty"`
	RegionName *string `json:"regionName,omitempty"`
	ResInfos *[]CommonResourceInfo `json:"resInfos,omitempty"`
	SubAccountId *string `json:"subAccountId,omitempty"`
	UserId *string `json:"userId,omitempty"`
}

// NewCloudmgrcommonIaasResource instantiates a new CloudmgrcommonIaasResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCloudmgrcommonIaasResource() *CloudmgrcommonIaasResource {
	this := CloudmgrcommonIaasResource{}
	return &this
}

// NewCloudmgrcommonIaasResourceWithDefaults instantiates a new CloudmgrcommonIaasResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCloudmgrcommonIaasResourceWithDefaults() *CloudmgrcommonIaasResource {
	this := CloudmgrcommonIaasResource{}
	return &this
}

// GetAz returns the Az field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetAz() string {
	if o == nil || o.Az == nil {
		var ret string
		return ret
	}
	return *o.Az
}

// GetAzOk returns a tuple with the Az field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetAzOk() (*string, bool) {
	if o == nil || o.Az == nil {
		return nil, false
	}
	return o.Az, true
}

// HasAz returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasAz() bool {
	if o != nil && o.Az != nil {
		return true
	}

	return false
}

// SetAz gets a reference to the given string and assigns it to the Az field.
func (o *CloudmgrcommonIaasResource) SetAz(v string) {
	o.Az = &v
}

// GetClusterName returns the ClusterName field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetClusterName() string {
	if o == nil || o.ClusterName == nil {
		var ret string
		return ret
	}
	return *o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetClusterNameOk() (*string, bool) {
	if o == nil || o.ClusterName == nil {
		return nil, false
	}
	return o.ClusterName, true
}

// HasClusterName returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasClusterName() bool {
	if o != nil && o.ClusterName != nil {
		return true
	}

	return false
}

// SetClusterName gets a reference to the given string and assigns it to the ClusterName field.
func (o *CloudmgrcommonIaasResource) SetClusterName(v string) {
	o.ClusterName = &v
}

// GetClusterType returns the ClusterType field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetClusterType() string {
	if o == nil || o.ClusterType == nil {
		var ret string
		return ret
	}
	return *o.ClusterType
}

// GetClusterTypeOk returns a tuple with the ClusterType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetClusterTypeOk() (*string, bool) {
	if o == nil || o.ClusterType == nil {
		return nil, false
	}
	return o.ClusterType, true
}

// HasClusterType returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasClusterType() bool {
	if o != nil && o.ClusterType != nil {
		return true
	}

	return false
}

// SetClusterType gets a reference to the given string and assigns it to the ClusterType field.
func (o *CloudmgrcommonIaasResource) SetClusterType(v string) {
	o.ClusterType = &v
}

// GetIp returns the Ip field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetIp() string {
	if o == nil || o.Ip == nil {
		var ret string
		return ret
	}
	return *o.Ip
}

// GetIpOk returns a tuple with the Ip field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetIpOk() (*string, bool) {
	if o == nil || o.Ip == nil {
		return nil, false
	}
	return o.Ip, true
}

// HasIp returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasIp() bool {
	if o != nil && o.Ip != nil {
		return true
	}

	return false
}

// SetIp gets a reference to the given string and assigns it to the Ip field.
func (o *CloudmgrcommonIaasResource) SetIp(v string) {
	o.Ip = &v
}

// GetMachineTypeCode returns the MachineTypeCode field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetMachineTypeCode() string {
	if o == nil || o.MachineTypeCode == nil {
		var ret string
		return ret
	}
	return *o.MachineTypeCode
}

// GetMachineTypeCodeOk returns a tuple with the MachineTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetMachineTypeCodeOk() (*string, bool) {
	if o == nil || o.MachineTypeCode == nil {
		return nil, false
	}
	return o.MachineTypeCode, true
}

// HasMachineTypeCode returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasMachineTypeCode() bool {
	if o != nil && o.MachineTypeCode != nil {
		return true
	}

	return false
}

// SetMachineTypeCode gets a reference to the given string and assigns it to the MachineTypeCode field.
func (o *CloudmgrcommonIaasResource) SetMachineTypeCode(v string) {
	o.MachineTypeCode = &v
}

// GetMachineTypeName returns the MachineTypeName field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetMachineTypeName() string {
	if o == nil || o.MachineTypeName == nil {
		var ret string
		return ret
	}
	return *o.MachineTypeName
}

// GetMachineTypeNameOk returns a tuple with the MachineTypeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetMachineTypeNameOk() (*string, bool) {
	if o == nil || o.MachineTypeName == nil {
		return nil, false
	}
	return o.MachineTypeName, true
}

// HasMachineTypeName returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasMachineTypeName() bool {
	if o != nil && o.MachineTypeName != nil {
		return true
	}

	return false
}

// SetMachineTypeName gets a reference to the given string and assigns it to the MachineTypeName field.
func (o *CloudmgrcommonIaasResource) SetMachineTypeName(v string) {
	o.MachineTypeName = &v
}

// GetProductTypeCode returns the ProductTypeCode field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetProductTypeCode() string {
	if o == nil || o.ProductTypeCode == nil {
		var ret string
		return ret
	}
	return *o.ProductTypeCode
}

// GetProductTypeCodeOk returns a tuple with the ProductTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetProductTypeCodeOk() (*string, bool) {
	if o == nil || o.ProductTypeCode == nil {
		return nil, false
	}
	return o.ProductTypeCode, true
}

// HasProductTypeCode returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasProductTypeCode() bool {
	if o != nil && o.ProductTypeCode != nil {
		return true
	}

	return false
}

// SetProductTypeCode gets a reference to the given string and assigns it to the ProductTypeCode field.
func (o *CloudmgrcommonIaasResource) SetProductTypeCode(v string) {
	o.ProductTypeCode = &v
}

// GetProjectList returns the ProjectList field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetProjectList() []CommonProjectInfo {
	if o == nil || o.ProjectList == nil {
		var ret []CommonProjectInfo
		return ret
	}
	return *o.ProjectList
}

// GetProjectListOk returns a tuple with the ProjectList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetProjectListOk() (*[]CommonProjectInfo, bool) {
	if o == nil || o.ProjectList == nil {
		return nil, false
	}
	return o.ProjectList, true
}

// HasProjectList returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasProjectList() bool {
	if o != nil && o.ProjectList != nil {
		return true
	}

	return false
}

// SetProjectList gets a reference to the given []CommonProjectInfo and assigns it to the ProjectList field.
func (o *CloudmgrcommonIaasResource) SetProjectList(v []CommonProjectInfo) {
	o.ProjectList = &v
}

// GetRegionCode returns the RegionCode field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetRegionCode() string {
	if o == nil || o.RegionCode == nil {
		var ret string
		return ret
	}
	return *o.RegionCode
}

// GetRegionCodeOk returns a tuple with the RegionCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetRegionCodeOk() (*string, bool) {
	if o == nil || o.RegionCode == nil {
		return nil, false
	}
	return o.RegionCode, true
}

// HasRegionCode returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasRegionCode() bool {
	if o != nil && o.RegionCode != nil {
		return true
	}

	return false
}

// SetRegionCode gets a reference to the given string and assigns it to the RegionCode field.
func (o *CloudmgrcommonIaasResource) SetRegionCode(v string) {
	o.RegionCode = &v
}

// GetRegionName returns the RegionName field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetRegionName() string {
	if o == nil || o.RegionName == nil {
		var ret string
		return ret
	}
	return *o.RegionName
}

// GetRegionNameOk returns a tuple with the RegionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetRegionNameOk() (*string, bool) {
	if o == nil || o.RegionName == nil {
		return nil, false
	}
	return o.RegionName, true
}

// HasRegionName returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasRegionName() bool {
	if o != nil && o.RegionName != nil {
		return true
	}

	return false
}

// SetRegionName gets a reference to the given string and assigns it to the RegionName field.
func (o *CloudmgrcommonIaasResource) SetRegionName(v string) {
	o.RegionName = &v
}

// GetResInfos returns the ResInfos field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetResInfos() []CommonResourceInfo {
	if o == nil || o.ResInfos == nil {
		var ret []CommonResourceInfo
		return ret
	}
	return *o.ResInfos
}

// GetResInfosOk returns a tuple with the ResInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetResInfosOk() (*[]CommonResourceInfo, bool) {
	if o == nil || o.ResInfos == nil {
		return nil, false
	}
	return o.ResInfos, true
}

// HasResInfos returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasResInfos() bool {
	if o != nil && o.ResInfos != nil {
		return true
	}

	return false
}

// SetResInfos gets a reference to the given []CommonResourceInfo and assigns it to the ResInfos field.
func (o *CloudmgrcommonIaasResource) SetResInfos(v []CommonResourceInfo) {
	o.ResInfos = &v
}

// GetSubAccountId returns the SubAccountId field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetSubAccountId() string {
	if o == nil || o.SubAccountId == nil {
		var ret string
		return ret
	}
	return *o.SubAccountId
}

// GetSubAccountIdOk returns a tuple with the SubAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetSubAccountIdOk() (*string, bool) {
	if o == nil || o.SubAccountId == nil {
		return nil, false
	}
	return o.SubAccountId, true
}

// HasSubAccountId returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasSubAccountId() bool {
	if o != nil && o.SubAccountId != nil {
		return true
	}

	return false
}

// SetSubAccountId gets a reference to the given string and assigns it to the SubAccountId field.
func (o *CloudmgrcommonIaasResource) SetSubAccountId(v string) {
	o.SubAccountId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *CloudmgrcommonIaasResource) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CloudmgrcommonIaasResource) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *CloudmgrcommonIaasResource) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *CloudmgrcommonIaasResource) SetUserId(v string) {
	o.UserId = &v
}

func (o CloudmgrcommonIaasResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Az != nil {
		toSerialize["az"] = o.Az
	}
	if o.ClusterName != nil {
		toSerialize["clusterName"] = o.ClusterName
	}
	if o.ClusterType != nil {
		toSerialize["clusterType"] = o.ClusterType
	}
	if o.Ip != nil {
		toSerialize["ip"] = o.Ip
	}
	if o.MachineTypeCode != nil {
		toSerialize["machineTypeCode"] = o.MachineTypeCode
	}
	if o.MachineTypeName != nil {
		toSerialize["machineTypeName"] = o.MachineTypeName
	}
	if o.ProductTypeCode != nil {
		toSerialize["productTypeCode"] = o.ProductTypeCode
	}
	if o.ProjectList != nil {
		toSerialize["projectList"] = o.ProjectList
	}
	if o.RegionCode != nil {
		toSerialize["regionCode"] = o.RegionCode
	}
	if o.RegionName != nil {
		toSerialize["regionName"] = o.RegionName
	}
	if o.ResInfos != nil {
		toSerialize["resInfos"] = o.ResInfos
	}
	if o.SubAccountId != nil {
		toSerialize["subAccountId"] = o.SubAccountId
	}
	if o.UserId != nil {
		toSerialize["userId"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullableCloudmgrcommonIaasResource struct {
	value *CloudmgrcommonIaasResource
	isSet bool
}

func (v NullableCloudmgrcommonIaasResource) Get() *CloudmgrcommonIaasResource {
	return v.value
}

func (v *NullableCloudmgrcommonIaasResource) Set(val *CloudmgrcommonIaasResource) {
	v.value = val
	v.isSet = true
}

func (v NullableCloudmgrcommonIaasResource) IsSet() bool {
	return v.isSet
}

func (v *NullableCloudmgrcommonIaasResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCloudmgrcommonIaasResource(val *CloudmgrcommonIaasResource) *NullableCloudmgrcommonIaasResource {
	return &NullableCloudmgrcommonIaasResource{value: val, isSet: true}
}

func (v NullableCloudmgrcommonIaasResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCloudmgrcommonIaasResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


