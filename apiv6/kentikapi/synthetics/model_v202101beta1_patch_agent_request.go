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

// V202101beta1PatchAgentRequest struct for V202101beta1PatchAgentRequest
type V202101beta1PatchAgentRequest struct {
	Agent      *V202101beta1Agent `json:"agent,omitempty"`
	UpdateMask *string            `json:"updateMask,omitempty"`
}

// NewV202101beta1PatchAgentRequest instantiates a new V202101beta1PatchAgentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV202101beta1PatchAgentRequest() *V202101beta1PatchAgentRequest {
	this := V202101beta1PatchAgentRequest{}
	return &this
}

// NewV202101beta1PatchAgentRequestWithDefaults instantiates a new V202101beta1PatchAgentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV202101beta1PatchAgentRequestWithDefaults() *V202101beta1PatchAgentRequest {
	this := V202101beta1PatchAgentRequest{}
	return &this
}

// GetAgent returns the Agent field value if set, zero value otherwise.
func (o *V202101beta1PatchAgentRequest) GetAgent() V202101beta1Agent {
	if o == nil || o.Agent == nil {
		var ret V202101beta1Agent
		return ret
	}
	return *o.Agent
}

// GetAgentOk returns a tuple with the Agent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1PatchAgentRequest) GetAgentOk() (*V202101beta1Agent, bool) {
	if o == nil || o.Agent == nil {
		return nil, false
	}
	return o.Agent, true
}

// HasAgent returns a boolean if a field has been set.
func (o *V202101beta1PatchAgentRequest) HasAgent() bool {
	if o != nil && o.Agent != nil {
		return true
	}

	return false
}

// SetAgent gets a reference to the given V202101beta1Agent and assigns it to the Agent field.
func (o *V202101beta1PatchAgentRequest) SetAgent(v V202101beta1Agent) {
	o.Agent = &v
}

// GetUpdateMask returns the UpdateMask field value if set, zero value otherwise.
func (o *V202101beta1PatchAgentRequest) GetUpdateMask() string {
	if o == nil || o.UpdateMask == nil {
		var ret string
		return ret
	}
	return *o.UpdateMask
}

// GetUpdateMaskOk returns a tuple with the UpdateMask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V202101beta1PatchAgentRequest) GetUpdateMaskOk() (*string, bool) {
	if o == nil || o.UpdateMask == nil {
		return nil, false
	}
	return o.UpdateMask, true
}

// HasUpdateMask returns a boolean if a field has been set.
func (o *V202101beta1PatchAgentRequest) HasUpdateMask() bool {
	if o != nil && o.UpdateMask != nil {
		return true
	}

	return false
}

// SetUpdateMask gets a reference to the given string and assigns it to the UpdateMask field.
func (o *V202101beta1PatchAgentRequest) SetUpdateMask(v string) {
	o.UpdateMask = &v
}

func (o V202101beta1PatchAgentRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Agent != nil {
		toSerialize["agent"] = o.Agent
	}
	if o.UpdateMask != nil {
		toSerialize["updateMask"] = o.UpdateMask
	}
	return json.Marshal(toSerialize)
}

type NullableV202101beta1PatchAgentRequest struct {
	value *V202101beta1PatchAgentRequest
	isSet bool
}

func (v NullableV202101beta1PatchAgentRequest) Get() *V202101beta1PatchAgentRequest {
	return v.value
}

func (v *NullableV202101beta1PatchAgentRequest) Set(val *V202101beta1PatchAgentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableV202101beta1PatchAgentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableV202101beta1PatchAgentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV202101beta1PatchAgentRequest(val *V202101beta1PatchAgentRequest) *NullableV202101beta1PatchAgentRequest {
	return &NullableV202101beta1PatchAgentRequest{value: val, isSet: true}
}

func (v NullableV202101beta1PatchAgentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV202101beta1PatchAgentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
