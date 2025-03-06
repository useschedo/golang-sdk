// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/Schedo-go/internal/apijson"
	"github.com/stainless-sdks/Schedo-go/internal/param"
	"github.com/stainless-sdks/Schedo-go/internal/requestconfig"
	"github.com/stainless-sdks/Schedo-go/option"
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

type APIKey struct {
	// Key holds the value of the "key" field.
	ID string `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at"`
	// Edges holds the relations/edges for other nodes in the graph. The values are
	// being populated by the ApiKeyQuery when eager-loading is set.
	Edges *APIKeyEdges `json:"edges"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID int64 `json:"organization_id"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt string     `json:"updated_at"`
	JSON      apiKeyJSON `json:"-"`
}

// apiKeyJSON contains the JSON metadata for the struct [APIKey]
type apiKeyJSON struct {
	ID             apijson.Field
	CreatedAt      apijson.Field
	Edges          apijson.Field
	Name           apijson.Field
	OrganizationID apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *APIKey) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyJSON) RawJSON() string {
	return r.raw
}

type APIKeyEdges struct {
	// Organization holds the value of the organization edge.
	Organization *Org            `json:"organization"`
	JSON         apiKeyEdgesJSON `json:"-"`
}

// apiKeyEdgesJSON contains the JSON metadata for the struct [APIKeyEdges]
type apiKeyEdgesJSON struct {
	Organization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *APIKeyEdges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r apiKeyEdgesJSON) RawJSON() string {
	return r.raw
}

type ApikeyNewParams struct {
	Name param.Field[string] `json:"name,required"`
}

func (r ApikeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
