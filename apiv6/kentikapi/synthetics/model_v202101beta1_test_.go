/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package synthetics

import (
	"encoding/json"
	"time"
)

// V202101beta1Test struct for V202101beta1Test
type V202101beta1Test struct {
	Id            *string                   `json:"id,omitempty"`
	Name          *string                   `json:"name,omitempty"`
	Type          *string                   `json:"type,omitempty"`
	DeviceId      *string                   `json:"deviceId,omitempty"`
	Status        *V202101beta1TestStatus   `json:"status,omitempty"`
	Settings      *V202101beta1TestSettings `json:"settings,omitempty"`
	ExpiresOn     *time.Time                `json:"expiresOn,omitempty"`
	Cdate         *time.Time                `json:"cdate,omitempty"`
	Edate         *time.Time                `json:"edate,omitempty"`
	CreatedBy     *V202101beta1UserInfo     `json:"createdBy,omitempty"`
	LastUpdatedBy *V202101beta1UserInfo     `json:"lastUpdatedBy,omitempty"`
}

// NewV202101beta1Test instantiates a new V202101beta1Test object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1Test() *V202101beta1Test {
	this := V202101beta1Test{}
	var status V202101beta1TestStatus = V202101BETA1TESTSTATUS_UNSPECIFIED
	this.Status = &status
	return &this
}

// NewV202101beta1TestWithDefaults instantiates a new V202101beta1Test object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1TestWithDefaults() *V202101beta1Test {
	this := V202101beta1Test{}
	var status V202101beta1TestStatus = V202101BETA1TESTSTATUS_UNSPECIFIED
	this.Status = &status
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V202101beta1Test) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V202101beta1Test) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V202101beta1Test) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V202101beta1Test) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V202101beta1Test) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V202101beta1Test) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V202101beta1Test) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V202101beta1Test) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *V202101beta1Test) SetType(v string) {
	o.Type = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *V202101beta1Test) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *V202101beta1Test) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *V202101beta1Test) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *V202101beta1Test) GetStatus() V202101beta1TestStatus {
	if o == nil || o.Status == nil {
		var ret V202101beta1TestStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetStatusOk() (*V202101beta1TestStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *V202101beta1Test) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given V202101beta1TestStatus and assigns it to the Status field.
func (o *V202101beta1Test) SetStatus(v V202101beta1TestStatus) {
	o.Status = &v
}

// GetSettings returns the Settings field value if set, zero value otherwise.
func (o *V202101beta1Test) GetSettings() V202101beta1TestSettings {
	if o == nil || o.Settings == nil {
		var ret V202101beta1TestSettings
		return ret
	}
	return *o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetSettingsOk() (*V202101beta1TestSettings, bool) {
	if o == nil || o.Settings == nil {
		return nil, false
	}
	return o.Settings, true
}

// HasSettings returns a boolean if a field has been set.
func (o *V202101beta1Test) HasSettings() bool {
	if o != nil && o.Settings != nil {
		return true
	}

	return false
}

// SetSettings gets a reference to the given V202101beta1TestSettings and assigns it to the Settings field.
func (o *V202101beta1Test) SetSettings(v V202101beta1TestSettings) {
	o.Settings = &v
}

// GetExpiresOn returns the ExpiresOn field value if set, zero value otherwise.
func (o *V202101beta1Test) GetExpiresOn() time.Time {
	if o == nil || o.ExpiresOn == nil {
		var ret time.Time
		return ret
	}
	return *o.ExpiresOn
}

// GetExpiresOnOk returns a tuple with the ExpiresOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetExpiresOnOk() (*time.Time, bool) {
	if o == nil || o.ExpiresOn == nil {
		return nil, false
	}
	return o.ExpiresOn, true
}

// HasExpiresOn returns a boolean if a field has been set.
func (o *V202101beta1Test) HasExpiresOn() bool {
	if o != nil && o.ExpiresOn != nil {
		return true
	}

	return false
}

// SetExpiresOn gets a reference to the given time.Time and assigns it to the ExpiresOn field.
func (o *V202101beta1Test) SetExpiresOn(v time.Time) {
	o.ExpiresOn = &v
}

// GetCdate returns the Cdate field value if set, zero value otherwise.
func (o *V202101beta1Test) GetCdate() time.Time {
	if o == nil || o.Cdate == nil {
		var ret time.Time
		return ret
	}
	return *o.Cdate
}

// GetCdateOk returns a tuple with the Cdate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetCdateOk() (*time.Time, bool) {
	if o == nil || o.Cdate == nil {
		return nil, false
	}
	return o.Cdate, true
}

// HasCdate returns a boolean if a field has been set.
func (o *V202101beta1Test) HasCdate() bool {
	if o != nil && o.Cdate != nil {
		return true
	}

	return false
}

// SetCdate gets a reference to the given time.Time and assigns it to the Cdate field.
func (o *V202101beta1Test) SetCdate(v time.Time) {
	o.Cdate = &v
}

// GetEdate returns the Edate field value if set, zero value otherwise.
func (o *V202101beta1Test) GetEdate() time.Time {
	if o == nil || o.Edate == nil {
		var ret time.Time
		return ret
	}
	return *o.Edate
}

// GetEdateOk returns a tuple with the Edate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetEdateOk() (*time.Time, bool) {
	if o == nil || o.Edate == nil {
		return nil, false
	}
	return o.Edate, true
}

// HasEdate returns a boolean if a field has been set.
func (o *V202101beta1Test) HasEdate() bool {
	if o != nil && o.Edate != nil {
		return true
	}

	return false
}

// SetEdate gets a reference to the given time.Time and assigns it to the Edate field.
func (o *V202101beta1Test) SetEdate(v time.Time) {
	o.Edate = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *V202101beta1Test) GetCreatedBy() V202101beta1UserInfo {
	if o == nil || o.CreatedBy == nil {
		var ret V202101beta1UserInfo
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetCreatedByOk() (*V202101beta1UserInfo, bool) {
	if o == nil || o.CreatedBy == nil {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *V202101beta1Test) HasCreatedBy() bool {
	if o != nil && o.CreatedBy != nil {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given V202101beta1UserInfo and assigns it to the CreatedBy field.
func (o *V202101beta1Test) SetCreatedBy(v V202101beta1UserInfo) {
	o.CreatedBy = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *V202101beta1Test) GetLastUpdatedBy() V202101beta1UserInfo {
	if o == nil || o.LastUpdatedBy == nil {
		var ret V202101beta1UserInfo
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1Test) GetLastUpdatedByOk() (*V202101beta1UserInfo, bool) {
	if o == nil || o.LastUpdatedBy == nil {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *V202101beta1Test) HasLastUpdatedBy() bool {
	if o != nil && o.LastUpdatedBy != nil {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given V202101beta1UserInfo and assigns it to the LastUpdatedBy field.
func (o *V202101beta1Test) SetLastUpdatedBy(v V202101beta1UserInfo) {
	o.LastUpdatedBy = &v
}

func (o V202101beta1Test) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Settings != nil {
		toSerialize["settings"] = o.Settings
	}
	if o.ExpiresOn != nil {
		toSerialize["expiresOn"] = o.ExpiresOn
	}
	if o.Cdate != nil {
		toSerialize["cdate"] = o.Cdate
	}
	if o.Edate != nil {
		toSerialize["edate"] = o.Edate
	}
	if o.CreatedBy != nil {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if o.LastUpdatedBy != nil {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1Test struct {
	value *V202101beta1Test
	isSet bool
}

func (v NullableV202101beta1Test) Get() *V202101beta1Test {
	return v.value
}

func (v *NullableV202101beta1Test) Set(val *V202101beta1Test) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1Test) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1Test) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1Test(val *V202101beta1Test) *NullableV202101beta1Test {
	return &NullableV202101beta1Test{value: val, isSet: true}
}

func (v NullableV202101beta1Test) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1Test) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
