// Package v1 provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package v1

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// AWSS3UploadRequestOptions defines model for AWSS3UploadRequestOptions.
type AWSS3UploadRequestOptions map[string]interface{}

// AWSS3UploadStatus defines model for AWSS3UploadStatus.
type AWSS3UploadStatus struct {
	Url string `json:"url"`
}

// AWSUploadRequestOptions defines model for AWSUploadRequestOptions.
type AWSUploadRequestOptions struct {
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// AWSUploadStatus defines model for AWSUploadStatus.
type AWSUploadStatus struct {
	Ami    string `json:"ami"`
	Region string `json:"region"`
}

// ArchitectureItem defines model for ArchitectureItem.
type ArchitectureItem struct {
	Arch       string   `json:"arch"`
	ImageTypes []string `json:"image_types"`
}

// Architectures defines model for Architectures.
type Architectures []ArchitectureItem

// AzureUploadRequestOptions defines model for AzureUploadRequestOptions.
type AzureUploadRequestOptions struct {

	// Name of the resource group where the image should be uploaded.
	ResourceGroup string `json:"resource_group"`

	// ID of subscription where the image should be uploaded.
	SubscriptionId string `json:"subscription_id"`

	// ID of the tenant where the image should be uploaded. This link explains how
	// to find it in the Azure Portal:
	// https://docs.microsoft.com/en-us/azure/active-directory/fundamentals/active-directory-how-to-find-tenant
	TenantId string `json:"tenant_id"`
}

// AzureUploadStatus defines model for AzureUploadStatus.
type AzureUploadStatus struct {
	ImageName string `json:"image_name"`
}

// ComposeMetadata defines model for ComposeMetadata.
type ComposeMetadata struct {

	// ID (hash) of the built commit
	OstreeCommit *string `json:"ostree_commit,omitempty"`

	// Package list including NEVRA
	Packages *[]PackageMetadata `json:"packages,omitempty"`
}

// ComposeRequest defines model for ComposeRequest.
type ComposeRequest struct {
	Customizations *Customizations `json:"customizations,omitempty"`
	Distribution   string          `json:"distribution"`
	ImageRequests  []ImageRequest  `json:"image_requests"`
}

// ComposeResponse defines model for ComposeResponse.
type ComposeResponse struct {
	Id string `json:"id"`
}

// ComposeStatus defines model for ComposeStatus.
type ComposeStatus struct {
	ImageStatus ImageStatus `json:"image_status"`
}

// ComposesResponse defines model for ComposesResponse.
type ComposesResponse struct {
	Data  []ComposesResponseItem `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// ComposesResponseItem defines model for ComposesResponseItem.
type ComposesResponseItem struct {
	CreatedAt string      `json:"created_at"`
	Id        string      `json:"id"`
	Request   interface{} `json:"request"`
}

// Customizations defines model for Customizations.
type Customizations struct {
	Packages     *[]string     `json:"packages,omitempty"`
	Subscription *Subscription `json:"subscription,omitempty"`
}

// DistributionItem defines model for DistributionItem.
type DistributionItem struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

// Distributions defines model for Distributions.
type Distributions []DistributionItem

// GCPUploadRequestOptions defines model for GCPUploadRequestOptions.
type GCPUploadRequestOptions struct {

	// List of valid Google accounts to share the imported Compute Node image with.
	// Each string must contain a specifier of the account type. Valid formats are:
	//   - 'user:{emailid}': An email address that represents a specific
	//     Google account. For example, 'alice@example.com'.
	//   - 'serviceAccount:{emailid}': An email address that represents a
	//     service account. For example, 'my-other-app@appspot.gserviceaccount.com'.
	//   - 'group:{emailid}': An email address that represents a Google group.
	//     For example, 'admins@example.com'.
	//   - 'domain:{domain}': The G Suite domain (primary) that represents all
	//     the users of that domain. For example, 'google.com' or 'example.com'.
	//     If not specified, the imported Compute Node image is not shared with any
	//     account.
	ShareWithAccounts []string `json:"share_with_accounts"`
}

// GCPUploadStatus defines model for GCPUploadStatus.
type GCPUploadStatus struct {
	ImageName string `json:"image_name"`
	ProjectId string `json:"project_id"`
}

// HTTPError defines model for HTTPError.
type HTTPError struct {
	Detail string `json:"detail"`
	Title  string `json:"title"`
}

// HTTPErrorList defines model for HTTPErrorList.
type HTTPErrorList struct {
	Errors []HTTPError `json:"errors"`
}

// ImageRequest defines model for ImageRequest.
type ImageRequest struct {
	Architecture  string        `json:"architecture"`
	ImageType     ImageTypes    `json:"image_type"`
	Ostree        *OSTree       `json:"ostree,omitempty"`
	UploadRequest UploadRequest `json:"upload_request"`
}

// ImageStatus defines model for ImageStatus.
type ImageStatus struct {
	Status       string        `json:"status"`
	UploadStatus *UploadStatus `json:"upload_status,omitempty"`
}

// ImageTypes defines model for ImageTypes.
type ImageTypes string

// List of ImageTypes
const (
	ImageTypes_ami                 ImageTypes = "ami"
	ImageTypes_aws                 ImageTypes = "aws"
	ImageTypes_azure               ImageTypes = "azure"
	ImageTypes_edge_commit         ImageTypes = "edge-commit"
	ImageTypes_edge_installer      ImageTypes = "edge-installer"
	ImageTypes_gcp                 ImageTypes = "gcp"
	ImageTypes_rhel_edge_commit    ImageTypes = "rhel-edge-commit"
	ImageTypes_rhel_edge_installer ImageTypes = "rhel-edge-installer"
	ImageTypes_vhd                 ImageTypes = "vhd"
)

// OSTree defines model for OSTree.
type OSTree struct {
	Ref *string `json:"ref,omitempty"`
	Url *string `json:"url,omitempty"`
}

// Package defines model for Package.
type Package struct {
	Name    string `json:"name"`
	Summary string `json:"summary"`
}

// PackageMetadata defines model for PackageMetadata.
type PackageMetadata struct {
	Arch      string  `json:"arch"`
	Epoch     *string `json:"epoch,omitempty"`
	Name      string  `json:"name"`
	Release   string  `json:"release"`
	Sigmd5    string  `json:"sigmd5"`
	Signature *string `json:"signature,omitempty"`
	Type      string  `json:"type"`
	Version   string  `json:"version"`
}

// PackagesResponse defines model for PackagesResponse.
type PackagesResponse struct {
	Data  []Package `json:"data"`
	Links struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"links"`
	Meta struct {
		Count int `json:"count"`
	} `json:"meta"`
}

// Readiness defines model for Readiness.
type Readiness struct {
	Readiness string `json:"readiness"`
}

// Subscription defines model for Subscription.
type Subscription struct {
	ActivationKey string `json:"activation-key"`
	BaseUrl       string `json:"base-url"`
	Insights      bool   `json:"insights"`
	Organization  int    `json:"organization"`
	ServerUrl     string `json:"server-url"`
}

// UploadRequest defines model for UploadRequest.
type UploadRequest struct {
	Options interface{} `json:"options"`
	Type    UploadTypes `json:"type"`
}

// UploadStatus defines model for UploadStatus.
type UploadStatus struct {
	Options interface{} `json:"options"`
	Status  string      `json:"status"`
	Type    UploadTypes `json:"type"`
}

// UploadTypes defines model for UploadTypes.
type UploadTypes string

// List of UploadTypes
const (
	UploadTypes_aws    UploadTypes = "aws"
	UploadTypes_aws_s3 UploadTypes = "aws.s3"
	UploadTypes_azure  UploadTypes = "azure"
	UploadTypes_gcp    UploadTypes = "gcp"
)

// Version defines model for Version.
type Version struct {
	Version string `json:"version"`
}

// ComposeImageJSONBody defines parameters for ComposeImage.
type ComposeImageJSONBody ComposeRequest

// GetComposesParams defines parameters for GetComposes.
type GetComposesParams struct {

	// max amount of composes, default 100
	Limit *int `json:"limit,omitempty"`

	// composes page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// GetPackagesParams defines parameters for GetPackages.
type GetPackagesParams struct {

	// distribution to look up packages for
	Distribution string `json:"distribution"`

	// architecture to look up packages for
	Architecture string `json:"architecture"`

	// packages to look for
	Search string `json:"search"`

	// max amount of packages, default 100
	Limit *int `json:"limit,omitempty"`

	// packages page offset, default 0
	Offset *int `json:"offset,omitempty"`
}

// ComposeImageRequestBody defines body for ComposeImage for application/json ContentType.
type ComposeImageJSONRequestBody ComposeImageJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// get the architectures and their image types available for a given distribution
	// (GET /architectures/{distribution})
	GetArchitectures(ctx echo.Context, distribution string) error
	// compose image
	// (POST /compose)
	ComposeImage(ctx echo.Context) error
	// get a collection of previous compose requests for the logged in user
	// (GET /composes)
	GetComposes(ctx echo.Context, params GetComposesParams) error
	// get status of an image compose
	// (GET /composes/{composeId})
	GetComposeStatus(ctx echo.Context, composeId string) error
	// get metadata of an image compose
	// (GET /composes/{composeId}/metadata)
	GetComposeMetadata(ctx echo.Context, composeId string) error
	// get the available distributions
	// (GET /distributions)
	GetDistributions(ctx echo.Context) error
	// get the openapi json specification
	// (GET /openapi.json)
	GetOpenapiJson(ctx echo.Context) error

	// (GET /packages)
	GetPackages(ctx echo.Context, params GetPackagesParams) error
	// return the readiness
	// (GET /ready)
	GetReadiness(ctx echo.Context) error
	// get the service version
	// (GET /version)
	GetVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// GetArchitectures converts echo context to params.
func (w *ServerInterfaceWrapper) GetArchitectures(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "distribution" -------------
	var distribution string

	err = runtime.BindStyledParameter("simple", false, "distribution", ctx.Param("distribution"), &distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetArchitectures(ctx, distribution)
	return err
}

// ComposeImage converts echo context to params.
func (w *ServerInterfaceWrapper) ComposeImage(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ComposeImage(ctx)
	return err
}

// GetComposes converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposes(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetComposesParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposes(ctx, params)
	return err
}

// GetComposeStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeStatus(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeStatus(ctx, composeId)
	return err
}

// GetComposeMetadata converts echo context to params.
func (w *ServerInterfaceWrapper) GetComposeMetadata(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "composeId" -------------
	var composeId string

	err = runtime.BindStyledParameter("simple", false, "composeId", ctx.Param("composeId"), &composeId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter composeId: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetComposeMetadata(ctx, composeId)
	return err
}

// GetDistributions converts echo context to params.
func (w *ServerInterfaceWrapper) GetDistributions(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetDistributions(ctx)
	return err
}

// GetOpenapiJson converts echo context to params.
func (w *ServerInterfaceWrapper) GetOpenapiJson(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetOpenapiJson(ctx)
	return err
}

// GetPackages converts echo context to params.
func (w *ServerInterfaceWrapper) GetPackages(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPackagesParams
	// ------------- Required query parameter "distribution" -------------

	err = runtime.BindQueryParameter("form", true, true, "distribution", ctx.QueryParams(), &params.Distribution)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter distribution: %s", err))
	}

	// ------------- Required query parameter "architecture" -------------

	err = runtime.BindQueryParameter("form", true, true, "architecture", ctx.QueryParams(), &params.Architecture)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter architecture: %s", err))
	}

	// ------------- Required query parameter "search" -------------

	err = runtime.BindQueryParameter("form", true, true, "search", ctx.QueryParams(), &params.Search)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter search: %s", err))
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// ------------- Optional query parameter "offset" -------------

	err = runtime.BindQueryParameter("form", true, false, "offset", ctx.QueryParams(), &params.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter offset: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPackages(ctx, params)
	return err
}

// GetReadiness converts echo context to params.
func (w *ServerInterfaceWrapper) GetReadiness(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetReadiness(ctx)
	return err
}

// GetVersion converts echo context to params.
func (w *ServerInterfaceWrapper) GetVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetVersion(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/architectures/:distribution", wrapper.GetArchitectures)
	router.POST("/compose", wrapper.ComposeImage)
	router.GET("/composes", wrapper.GetComposes)
	router.GET("/composes/:composeId", wrapper.GetComposeStatus)
	router.GET("/composes/:composeId/metadata", wrapper.GetComposeMetadata)
	router.GET("/distributions", wrapper.GetDistributions)
	router.GET("/openapi.json", wrapper.GetOpenapiJson)
	router.GET("/packages", wrapper.GetPackages)
	router.GET("/ready", wrapper.GetReadiness)
	router.GET("/version", wrapper.GetVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file.
func GetSwagger() (*openapi3.Swagger, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	swagger, err := openapi3.NewSwaggerLoader().LoadSwaggerFromData(buf.Bytes())
	if err != nil {
		return nil, fmt.Errorf("error loading Swagger: %s", err)
	}
	return swagger, nil
}
