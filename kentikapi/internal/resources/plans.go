package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

type PlansAPI struct {
	BaseAPI
}

// NewPlansAPI is constructor.
func NewPlansAPI(transport api_connection.Transport, logPayloads bool) *PlansAPI {
	return &PlansAPI{
		BaseAPI{Transport: transport, LogPayloads: logPayloads},
	}
}

// GetAll plans.
func (a *PlansAPI) GetAll(ctx context.Context) ([]models.Plan, error) {
	utils.LogPayload(a.LogPayloads, "GetAll plans Kentik API request", "")
	var response api_payloads.GetAllPlansResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetAllPlans(), &response); err != nil {
		return []models.Plan{}, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll plans Kentik API response", response)

	return response.ToPlans()
}
