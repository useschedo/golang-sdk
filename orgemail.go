// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package schedo

import (
	"context"
	"net/http"

	"github.com/stainless-sdks/schedosdk-go/internal/apijson"
	"github.com/stainless-sdks/schedosdk-go/internal/requestconfig"
	"github.com/stainless-sdks/schedosdk-go/option"
)

// OrgEmailService contains methods and other services that help with interacting
// with the Schedo API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgEmailService] method instead.
type OrgEmailService struct {
	Options []option.RequestOption
}

// NewOrgEmailService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgEmailService(opts ...option.RequestOption) (r *OrgEmailService) {
	r = &OrgEmailService{}
	r.Options = opts
	return
}

// Adds an email to the organization
func (r *OrgEmailService) New(ctx context.Context, body OrgEmailNewParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "org/emails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Gets all emails for the organization
func (r *OrgEmailService) List(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "org/emails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Removes an email from the organization
func (r *OrgEmailService) Delete(ctx context.Context, body OrgEmailDeleteParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	path := "org/emails"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type OrgEmailNewParams struct {
	Body string `json:"body,required"`
}

func (r OrgEmailNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type OrgEmailDeleteParams struct {
	Body string `json:"body,required"`
}

func (r OrgEmailDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}
