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

// V202101beta1KnockTaskDefinition struct for V202101beta1KnockTaskDefinition
type V202101beta1KnockTaskDefinition struct {
	Target *string `json:"target,omitempty"`
	Period *int64  `json:"period,omitempty"`
	Expiry *int64  `json:"expiry,omitempty"`
	Count  *int64  `json:"count,omitempty"`
	Port   *int64  `json:"port,omitempty"`
}

// NewV202101beta1KnockTaskDefinition instantiates a new V202101beta1KnockTaskDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1KnockTaskDefinition() *V202101beta1KnockTaskDefinition {
	this := V202101beta1KnockTaskDefinition{}
	return &this
}

// NewV202101beta1KnockTaskDefinitionWithDefaults instantiates a new V202101beta1KnockTaskDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1KnockTaskDefinitionWithDefaults() *V202101beta1KnockTaskDefinition {
	this := V202101beta1KnockTaskDefinition{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *V202101beta1KnockTaskDefinition) GetTarget() string {
	if o == nil || o.Target == nil {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1KnockTaskDefinition) GetTargetOk() (*string, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *V202101beta1KnockTaskDefinition) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *V202101beta1KnockTaskDefinition) SetTarget(v string) {
	o.Target = &v
}

// GetPeriod returns the Period field value if set, zero value otherwise.
func (o *V202101beta1KnockTaskDefinition) GetPeriod() int64 {
	if o == nil || o.Period == nil {
		var ret int64
		return ret
	}
	return *o.Period
}

// GetPeriodOk returns a tuple with the Period field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1KnockTaskDefinition) GetPeriodOk() (*int64, bool) {
	if o == nil || o.Period == nil {
		return nil, false
	}
	return o.Period, true
}

// HasPeriod returns a boolean if a field has been set.
func (o *V202101beta1KnockTaskDefinition) HasPeriod() bool {
	if o != nil && o.Period != nil {
		return true
	}

	return false
}

// SetPeriod gets a reference to the given int64 and assigns it to the Period field.
func (o *V202101beta1KnockTaskDefinition) SetPeriod(v int64) {
	o.Period = &v
}

// GetExpiry returns the Expiry field value if set, zero value otherwise.
func (o *V202101beta1KnockTaskDefinition) GetExpiry() int64 {
	if o == nil || o.Expiry == nil {
		var ret int64
		return ret
	}
	return *o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1KnockTaskDefinition) GetExpiryOk() (*int64, bool) {
	if o == nil || o.Expiry == nil {
		return nil, false
	}
	return o.Expiry, true
}

// HasExpiry returns a boolean if a field has been set.
func (o *V202101beta1KnockTaskDefinition) HasExpiry() bool {
	if o != nil && o.Expiry != nil {
		return true
	}

	return false
}

// SetExpiry gets a reference to the given int64 and assigns it to the Expiry field.
func (o *V202101beta1KnockTaskDefinition) SetExpiry(v int64) {
	o.Expiry = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *V202101beta1KnockTaskDefinition) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1KnockTaskDefinition) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *V202101beta1KnockTaskDefinition) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *V202101beta1KnockTaskDefinition) SetCount(v int64) {
	o.Count = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *V202101beta1KnockTaskDefinition) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1KnockTaskDefinition) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *V202101beta1KnockTaskDefinition) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *V202101beta1KnockTaskDefinition) SetPort(v int64) {
	o.Port = &v
}

func (o V202101beta1KnockTaskDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	if o.Period != nil {
		toSerialize["period"] = o.Period
	}
	if o.Expiry != nil {
		toSerialize["expiry"] = o.Expiry
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1KnockTaskDefinition struct {
	value *V202101beta1KnockTaskDefinition
	isSet bool
}

func (v NullableV202101beta1KnockTaskDefinition) Get() *V202101beta1KnockTaskDefinition {
	return v.value
}

func (v *NullableV202101beta1KnockTaskDefinition) Set(val *V202101beta1KnockTaskDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1KnockTaskDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1KnockTaskDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1KnockTaskDefinition(val *V202101beta1KnockTaskDefinition) *NullableV202101beta1KnockTaskDefinition {
	return &NullableV202101beta1KnockTaskDefinition{value: val, isSet: true}
}

func (v NullableV202101beta1KnockTaskDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1KnockTaskDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
