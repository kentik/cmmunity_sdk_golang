/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

type V202101beta1TraceHop struct {
	Agent bool `json:"agent,omitempty"`

	AgentId string `json:"agentId,omitempty"`

	Ttl int32 `json:"ttl,omitempty"`

	Ip string `json:"ip,omitempty"`

	Latencies []string `json:"latencies,omitempty"`

	Timeout bool `json:"timeout,omitempty"`
}
