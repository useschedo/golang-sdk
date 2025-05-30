// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/schedosdk-go/internal/apijson"
	"github.com/stainless-sdks/schedosdk-go/internal/param"
	"github.com/stainless-sdks/schedosdk-go/internal/requestconfig"
	"github.com/stainless-sdks/schedosdk-go/option"
)

// EnvironmentService contains methods and other services that help with
// interacting with the Schedo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnvironmentService] method instead.
type EnvironmentService struct {
	Options []option.RequestOption
}

// NewEnvironmentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnvironmentService(opts ...option.RequestOption) (r *EnvironmentService) {
	r = &EnvironmentService{}
	r.Options = opts
	return
}

// Creates a new org environment
func (r *EnvironmentService) New(ctx context.Context, body EnvironmentNewParams, opts ...option.RequestOption) (res *Environment, err error) {
	opts = append(r.Options[:], opts...)
	path := "org/environments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves a list of environments for the current org
func (r *EnvironmentService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Environment, err error) {
	opts = append(r.Options[:], opts...)
	path := "org/environments"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an org environment
func (r *EnvironmentService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (res *Environment, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("org/environments/%v", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Environment struct {
	// ID of the ent.
	ID int64 `json:"id"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt string `json:"created_at"`
	// Edges holds the relations/edges for other nodes in the graph. The values are
	// being populated by the EnvironmentQuery when eager-loading is set.
	Edges EnvironmentEdges `json:"edges"`
	// Name holds the value of the "name" field.
	Name string `json:"name"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID int64 `json:"organization_id"`
	// A webhook signature key to validate incoming requests
	WebhookSignatureKey string          `json:"webhook_signature_key"`
	JSON                environmentJSON `json:"-"`
}

// environmentJSON contains the JSON metadata for the struct [Environment]
type environmentJSON struct {
	ID                  apijson.Field
	CreatedAt           apijson.Field
	Edges               apijson.Field
	Name                apijson.Field
	OrganizationID      apijson.Field
	WebhookSignatureKey apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *Environment) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentJSON) RawJSON() string {
	return r.raw
}

type EnvironmentEdges struct {
	// Organization that owns this environment
	Organization []Org                `json:"organization"`
	JSON         environmentEdgesJSON `json:"-"`
}

// environmentEdgesJSON contains the JSON metadata for the struct
// [EnvironmentEdges]
type environmentEdgesJSON struct {
	Organization apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *EnvironmentEdges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r environmentEdgesJSON) RawJSON() string {
	return r.raw
}

type EnvironmentNewParams struct {
	Name param.Field[string] `json:"name,required"`
}

func (r EnvironmentNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
