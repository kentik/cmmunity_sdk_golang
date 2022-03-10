//nolint:dupl
package resources

import (
	"context"

	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_connection"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_endpoints"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/api_payloads"
	"github.com/kentik/community_sdk_golang/kentikapi/internal/utils"
	"github.com/kentik/community_sdk_golang/kentikapi/models"
)

// UsersAPI aggregates Users API methods.
type UsersAPI struct {
	BaseAPI
}

// NewUsersAPI creates new UsersAPI.
func NewUsersAPI(transport api_connection.Transport, logPayloads bool) *UsersAPI {
	return &UsersAPI{BaseAPI{Transport: transport, LogPayloads: logPayloads}}
}

// GetAll lists users.
func (a *UsersAPI) GetAll(ctx context.Context) ([]models.User, error) {
	utils.LogPayload(a.LogPayloads, "GetAll users Kentik API request", "")
	var response api_payloads.GetAllUsersResponse
	if err := a.GetAndValidate(ctx, api_endpoints.UsersPath, &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "GetAll users Kentik API response", response)

	return response.ToUsers(), nil
}

// Get retrieves user with given ID.
func (a *UsersAPI) Get(ctx context.Context, id models.ID) (*models.User, error) {
	utils.LogPayload(a.LogPayloads, "Get user Kentik API request ID", id)
	var response api_payloads.GetUserResponse
	if err := a.GetAndValidate(ctx, api_endpoints.GetUserPath(id), &response); err != nil {
		return nil, err
	}
	utils.LogPayload(a.LogPayloads, "Get user Kentik API response", response)

	return response.ToUser(), nil
}

// Create creates new user.
func (a *UsersAPI) Create(ctx context.Context, user models.User) (*models.User, error) {
	request := api_payloads.CreateUserRequest{User: api_payloads.UserToPayload(user)}
	utils.LogPayload(a.LogPayloads, "Create user Kentik API request", request)
	var response api_payloads.CreateUserResponse
	err := a.PostAndValidate(
		ctx,
		api_endpoints.UserPath,
		request,
		&response,
	)
	utils.LogPayload(a.LogPayloads, "Create user Kentik API response", response)
	if err != nil {
		return nil, err
	}

	return response.ToUser(), nil
}

// Update updates the user.
func (a *UsersAPI) Update(ctx context.Context, user models.User) (*models.User, error) {
	request := api_payloads.UpdateUserRequest{User: api_payloads.UserToPayload(user)}
	utils.LogPayload(a.LogPayloads, "Update user Kentik API request", request)
	var response api_payloads.UpdateUserResponse
	err := a.UpdateAndValidate(
		ctx,
		api_endpoints.GetUserPath(user.ID),
		request,
		&response,
	)
	utils.LogPayload(a.LogPayloads, "Update user Kentik API response", response)
	if err != nil {
		return nil, err
	}

	return response.ToUser(), err
}

// Delete removes user with given ID.
func (a *UsersAPI) Delete(ctx context.Context, id models.ID) error {
	utils.LogPayload(a.LogPayloads, "Delete user Kentik API request ID", id)
	if err := a.DeleteAndValidate(ctx, api_endpoints.GetUserPath(id), nil); err != nil {
		return err
	}

	return nil
}
