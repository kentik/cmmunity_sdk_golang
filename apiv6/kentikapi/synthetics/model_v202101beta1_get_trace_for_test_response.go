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
)

// V202101beta1GetTraceForTestResponse struct for V202101beta1GetTraceForTestResponse
type V202101beta1GetTraceForTestResponse struct {
	IpInfo  *[]V202101beta1IPInfo           `json:"ipInfo,omitempty"`
	Results *[]V202101beta1TracerouteResult `json:"results,omitempty"`
}

// NewV202101beta1GetTraceForTestResponse instantiates a new V202101beta1GetTraceForTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1GetTraceForTestResponse() *V202101beta1GetTraceForTestResponse {
	this := V202101beta1GetTraceForTestResponse{}
	return &this
}

// NewV202101beta1GetTraceForTestResponseWithDefaults instantiates a new V202101beta1GetTraceForTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1GetTraceForTestResponseWithDefaults() *V202101beta1GetTraceForTestResponse {
	this := V202101beta1GetTraceForTestResponse{}
	return &this
}

// GetIpInfo returns the IpInfo field value if set, zero value otherwise.
func (o *V202101beta1GetTraceForTestResponse) GetIpInfo() []V202101beta1IPInfo {
	if o == nil || o.IpInfo == nil {
		var ret []V202101beta1IPInfo
		return ret
	}
	return *o.IpInfo
}

// GetIpInfoOk returns a tuple with the IpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1GetTraceForTestResponse) GetIpInfoOk() (*[]V202101beta1IPInfo, bool) {
	if o == nil || o.IpInfo == nil {
		return nil, false
	}
	return o.IpInfo, true
}

// HasIpInfo returns a boolean if a field has been set.
func (o *V202101beta1GetTraceForTestResponse) HasIpInfo() bool {
	if o != nil && o.IpInfo != nil {
		return true
	}

	return false
}

// SetIpInfo gets a reference to the given []V202101beta1IPInfo and assigns it to the IpInfo field.
func (o *V202101beta1GetTraceForTestResponse) SetIpInfo(v []V202101beta1IPInfo) {
	o.IpInfo = &v
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *V202101beta1GetTraceForTestResponse) GetResults() []V202101beta1TracerouteResult {
	if o == nil || o.Results == nil {
		var ret []V202101beta1TracerouteResult
		return ret
	}
	return *o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1GetTraceForTestResponse) GetResultsOk() (*[]V202101beta1TracerouteResult, bool) {
	if o == nil || o.Results == nil {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *V202101beta1GetTraceForTestResponse) HasResults() bool {
	if o != nil && o.Results != nil {
		return true
	}

	return false
}

// SetResults gets a reference to the given []V202101beta1TracerouteResult and assigns it to the Results field.
func (o *V202101beta1GetTraceForTestResponse) SetResults(v []V202101beta1TracerouteResult) {
	o.Results = &v
}

func (o V202101beta1GetTraceForTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IpInfo != nil {
		toSerialize["ipInfo"] = o.IpInfo
	}
	if o.Results != nil {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1GetTraceForTestResponse struct {
	value *V202101beta1GetTraceForTestResponse
	isSet bool
}

func (v NullableV202101beta1GetTraceForTestResponse) Get() *V202101beta1GetTraceForTestResponse {
	return v.value
}

func (v *NullableV202101beta1GetTraceForTestResponse) Set(val *V202101beta1GetTraceForTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1GetTraceForTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1GetTraceForTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1GetTraceForTestResponse(val *V202101beta1GetTraceForTestResponse) *NullableV202101beta1GetTraceForTestResponse {
	return &NullableV202101beta1GetTraceForTestResponse{value: val, isSet: true}
}

func (v NullableV202101beta1GetTraceForTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1GetTraceForTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
