/*
 * Synthetics Monitoring API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 202101beta1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package syntheticsstub

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// A SyntheticsAdminServiceApiController binds http requests to an api service and writes the service results to the http response
type SyntheticsAdminServiceApiController struct {
	service SyntheticsAdminServiceApiServicer
}

// NewSyntheticsAdminServiceApiController creates a default api controller
func NewSyntheticsAdminServiceApiController(s SyntheticsAdminServiceApiServicer) Router {
	return &SyntheticsAdminServiceApiController{service: s}
}

// Routes returns all of the api route for the SyntheticsAdminServiceApiController
func (c *SyntheticsAdminServiceApiController) Routes() Routes {
	return Routes{
		{
			"AgentDelete",
			strings.ToUpper("Delete"),
			"/synthetics/v202101beta1/agents/{agent.id}",
			c.AgentDelete,
		},
		{
			"AgentGet",
			strings.ToUpper("Get"),
			"/synthetics/v202101beta1/agents/{agent.id}",
			c.AgentGet,
		},
		{
			"AgentPatch",
			strings.ToUpper("Patch"),
			"/synthetics/v202101beta1/agents/{agent.id}",
			c.AgentPatch,
		},
		{
			"AgentsList",
			strings.ToUpper("Get"),
			"/synthetics/v202101beta1/agents",
			c.AgentsList,
		},
		{
			"TestCreate",
			strings.ToUpper("Post"),
			"/synthetics/v202101beta1/tests",
			c.TestCreate,
		},
		{
			"TestDelete",
			strings.ToUpper("Delete"),
			"/synthetics/v202101beta1/tests/{id}",
			c.TestDelete,
		},
		{
			"TestGet",
			strings.ToUpper("Get"),
			"/synthetics/v202101beta1/tests/{id}",
			c.TestGet,
		},
		{
			"TestPatch",
			strings.ToUpper("Patch"),
			"/synthetics/v202101beta1/tests/{id}",
			c.TestPatch,
		},
		{
			"TestStatusUpdate",
			strings.ToUpper("Put"),
			"/synthetics/v202101beta1/tests/{id}/status",
			c.TestStatusUpdate,
		},
		{
			"TestsList",
			strings.ToUpper("Get"),
			"/synthetics/v202101beta1/tests",
			c.TestsList,
		},
	}
}

// AgentDelete - Delete an agent.
func (c *SyntheticsAdminServiceApiController) AgentDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agentId := params["agent.id"]
	result, err := c.service.AgentDelete(r.Context(), agentId)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AgentGet - Get information about an agent.
func (c *SyntheticsAdminServiceApiController) AgentGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agentId := params["agent.id"]
	result, err := c.service.AgentGet(r.Context(), agentId)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AgentPatch - Patch an agent.
func (c *SyntheticsAdminServiceApiController) AgentPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	agentId := params["agent.id"]
	v202101beta1PatchAgentRequest := &V202101beta1PatchAgentRequest{}
	if err := json.NewDecoder(r.Body).Decode(&v202101beta1PatchAgentRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.AgentPatch(r.Context(), agentId, *v202101beta1PatchAgentRequest)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// AgentsList - List Agents.
func (c *SyntheticsAdminServiceApiController) AgentsList(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.AgentsList(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// TestCreate - Create Synthetics Test.
func (c *SyntheticsAdminServiceApiController) TestCreate(w http.ResponseWriter, r *http.Request) {
	v202101beta1CreateTestRequest := &V202101beta1CreateTestRequest{}
	if err := json.NewDecoder(r.Body).Decode(&v202101beta1CreateTestRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.TestCreate(r.Context(), *v202101beta1CreateTestRequest)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// TestDelete - Delete an Synthetics Test.
func (c *SyntheticsAdminServiceApiController) TestDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := c.service.TestDelete(r.Context(), id)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// TestGet - Get information about Synthetics Test.
func (c *SyntheticsAdminServiceApiController) TestGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	result, err := c.service.TestGet(r.Context(), id)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// TestPatch - Patch a Synthetics Test.
func (c *SyntheticsAdminServiceApiController) TestPatch(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	v202101beta1PatchTestRequest := &V202101beta1PatchTestRequest{}
	if err := json.NewDecoder(r.Body).Decode(&v202101beta1PatchTestRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.TestPatch(r.Context(), id, *v202101beta1PatchTestRequest)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// TestStatusUpdate - Update a test status.
func (c *SyntheticsAdminServiceApiController) TestStatusUpdate(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	v202101beta1SetTestStatusRequest := &V202101beta1SetTestStatusRequest{}
	if err := json.NewDecoder(r.Body).Decode(&v202101beta1SetTestStatusRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := c.service.TestStatusUpdate(r.Context(), id, *v202101beta1SetTestStatusRequest)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// TestsList - List Synthetics Tests.
func (c *SyntheticsAdminServiceApiController) TestsList(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	preset, _ := strconv.ParseBool(query.Get("preset")) // if error -> preset = false
	result, err := c.service.TestsList(r.Context(), preset)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
