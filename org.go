// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"context"
	"net/http"

	"github.com/useschedo/golang-sdk/internal/apijson"
	"github.com/useschedo/golang-sdk/internal/requestconfig"
	"github.com/useschedo/golang-sdk/option"
)

// OrgService contains methods and other services that help with interacting with
// the Schedo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgService] method instead.
type OrgService struct {
	Options []option.RequestOption
}

// NewOrgService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgService(opts ...option.RequestOption) (r *OrgService) {
	r = &OrgService{}
	r.Options = opts
	return
}

// Retrieves information about current org
func (r *OrgService) Get(ctx context.Context, opts ...option.RequestOption) (res *Org, err error) {
	opts = append(r.Options[:], opts...)
	path := "org"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type Org struct {
	// ID of the ent.
	ID int64 `json:"id"`
	// External Clerk organization ID
	ClerkID string `json:"clerk_id"`
	// Edges holds the relations/edges for other nodes in the graph. The values are
	// being populated by the OrgQuery when eager-loading is set.
	Edges OrgEdges `json:"edges"`
	// Organization name
	Name string `json:"name"`
	// SingleUserOrg holds the value of the "single_user_org" field.
	SingleUserOrg bool    `json:"single_user_org"`
	JSON          orgJSON `json:"-"`
}

// orgJSON contains the JSON metadata for the struct [Org]
type orgJSON struct {
	ID            apijson.Field
	ClerkID       apijson.Field
	Edges         apijson.Field
	Name          apijson.Field
	SingleUserOrg apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Org) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgJSON) RawJSON() string {
	return r.raw
}

type OrgEdges struct {
	// APIKeys holds the value of the api_keys edge.
	APIKeys []APIKey     `json:"api_keys"`
	JSON    orgEdgesJSON `json:"-"`
}

// orgEdgesJSON contains the JSON metadata for the struct [OrgEdges]
type orgEdgesJSON struct {
	APIKeys     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgEdges) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgEdgesJSON) RawJSON() string {
	return r.raw
}
