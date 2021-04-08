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

// V202101beta1TracerouteResult struct for V202101beta1TracerouteResult
type V202101beta1TracerouteResult struct {
	Time   *time.Time           `json:"time,omitempty"`
	Traces *[]V202101beta1Trace `json:"traces,omitempty"`
}

// NewV202101beta1TracerouteResult instantiates a new V202101beta1TracerouteResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1TracerouteResult() *V202101beta1TracerouteResult {
	this := V202101beta1TracerouteResult{}
	return &this
}

// NewV202101beta1TracerouteResultWithDefaults instantiates a new V202101beta1TracerouteResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1TracerouteResultWithDefaults() *V202101beta1TracerouteResult {
	this := V202101beta1TracerouteResult{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *V202101beta1TracerouteResult) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TracerouteResult) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *V202101beta1TracerouteResult) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *V202101beta1TracerouteResult) SetTime(v time.Time) {
	o.Time = &v
}

// GetTraces returns the Traces field value if set, zero value otherwise.
func (o *V202101beta1TracerouteResult) GetTraces() []V202101beta1Trace {
	if o == nil || o.Traces == nil {
		var ret []V202101beta1Trace
		return ret
	}
	return *o.Traces
}

// GetTracesOk returns a tuple with the Traces field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1TracerouteResult) GetTracesOk() (*[]V202101beta1Trace, bool) {
	if o == nil || o.Traces == nil {
		return nil, false
	}
	return o.Traces, true
}

// HasTraces returns a boolean if a field has been set.
func (o *V202101beta1TracerouteResult) HasTraces() bool {
	if o != nil && o.Traces != nil {
		return true
	}

	return false
}

// SetTraces gets a reference to the given []V202101beta1Trace and assigns it to the Traces field.
func (o *V202101beta1TracerouteResult) SetTraces(v []V202101beta1Trace) {
	o.Traces = &v
}

func (o V202101beta1TracerouteResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Traces != nil {
		toSerialize["traces"] = o.Traces
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1TracerouteResult struct {
	value *V202101beta1TracerouteResult
	isSet bool
}

func (v NullableV202101beta1TracerouteResult) Get() *V202101beta1TracerouteResult {
	return v.value
}

func (v *NullableV202101beta1TracerouteResult) Set(val *V202101beta1TracerouteResult) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1TracerouteResult) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1TracerouteResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1TracerouteResult(val *V202101beta1TracerouteResult) *NullableV202101beta1TracerouteResult {
	return &NullableV202101beta1TracerouteResult{value: val, isSet: true}
}

func (v NullableV202101beta1TracerouteResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1TracerouteResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
