// Package private provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package private

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Dispatch Playbooks
	// (POST /internal/dispatch)
	ApiInternalRunsCreate(ctx echo.Context) error
	// Cancel Playbook Runs
	// (POST /internal/v2/cancel)
	ApiInternalV2RunsCancel(ctx echo.Context) error
	// Obtain Connection Status of recipient(s) based on a list of host IDs
	// (POST /internal/v2/connection_status)
	ApiInternalHighlevelConnectionStatus(ctx echo.Context) error
	// Dispatch Playbooks
	// (POST /internal/v2/dispatch)
	ApiInternalV2RunsCreate(ctx echo.Context) error
	// Obtain connection status of recipient(s)
	// (POST /internal/v2/recipients/status)
	ApiInternalV2RecipientsStatus(ctx echo.Context) error
	// Get Version
	// (GET /internal/version)
	ApiInternalVersion(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// ApiInternalRunsCreate converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalRunsCreate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalRunsCreate(ctx)
	return err
}

// ApiInternalV2RunsCancel converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalV2RunsCancel(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalV2RunsCancel(ctx)
	return err
}

// ApiInternalHighlevelConnectionStatus converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalHighlevelConnectionStatus(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalHighlevelConnectionStatus(ctx)
	return err
}

// ApiInternalV2RunsCreate converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalV2RunsCreate(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalV2RunsCreate(ctx)
	return err
}

// ApiInternalV2RecipientsStatus converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalV2RecipientsStatus(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalV2RecipientsStatus(ctx)
	return err
}

// ApiInternalVersion converts echo context to params.
func (w *ServerInterfaceWrapper) ApiInternalVersion(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ApiInternalVersion(ctx)
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
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.POST(baseURL+"/internal/dispatch", wrapper.ApiInternalRunsCreate)
	router.POST(baseURL+"/internal/v2/cancel", wrapper.ApiInternalV2RunsCancel)
	router.POST(baseURL+"/internal/v2/connection_status", wrapper.ApiInternalHighlevelConnectionStatus)
	router.POST(baseURL+"/internal/v2/dispatch", wrapper.ApiInternalV2RunsCreate)
	router.POST(baseURL+"/internal/v2/recipients/status", wrapper.ApiInternalV2RecipientsStatus)
	router.GET(baseURL+"/internal/version", wrapper.ApiInternalVersion)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"AP//viQrb4okAAA=",
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
