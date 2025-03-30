// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/useschedo/golang-sdk/internal/apijson"
	"github.com/useschedo/golang-sdk/internal/param"
	"github.com/useschedo/golang-sdk/internal/requestconfig"
	"github.com/useschedo/golang-sdk/option"
)

// ApikeyService contains methods and other services that help with interacting
// with the Schedo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApikeyService] method instead.
type ApikeyService struct {
	Options []option.RequestOption
}

// NewApikeyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewApikeyService(opts ...option.RequestOption) (r *ApikeyService) {
	r = &ApikeyService{}
	r.Options = opts
	return
}

// Generates a new API Key to access Schedo.dev API
func (r *ApikeyService) New(ctx context.Context, body ApikeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = append(r.Options[:], opts...)
	path := "apikeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a list of API Keys for the organization
func (r *ApikeyService) List(ctx context.Context, opts ...option.RequestOption) (res *[][]APIKey, err error) {
	opts = append(r.Options[:], opts...)
	path := "apikeys"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revokes an API Key making it inactive
func (r *ApikeyService) Revoke(ctx context.Context, id int64, opts ...option.RequestOption) (res *[][]APIKey, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("apikeys/revoke/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type APIKey struct {
	// ID of the ent.
	ID int64 `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at"`
	// Edges holds the relations/edges for other nodes in the graph. The values are
	// being populated by the ApiKeyQuery when eager-loading is set.
	Edges APIKeyEdges `json:"edges"`
	// EnvironmentID holds the value of the "environment_id" field.
	EnvironmentID int64 `json:"environment_id"`
	// Internal holds the value of the "internal" field.
	Internal bool `json:"internal"`
	// Key holds the value of the "key" field.
	Key string `json:"key"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID int64 `json:"organization_id"`
	// Revoked holds the value of the "revoked" field.
	Revoked bool       `json:"revoked"`
	JSON    apiKeyJSON `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Edges          apijson.Field
	EnvironmentID  apijson.Field
	Internal       apijson.Field
	Key            apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	Revoked        apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

type APIKeyEdges = interface{}

type ApikeyNewParams struct {
	EnvironmentID param.Field[int64]  `json:"environment_id,required"`
	Name          param.Field[string] `json:"name,required"`
}

func (r ApikeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
