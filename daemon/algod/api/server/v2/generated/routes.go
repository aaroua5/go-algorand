// Package generated provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/algorand/oapi-codegen DO NOT EDIT.
package generated

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"github.com/algorand/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get account information.
	// (GET /v2/accounts/{address})
	AccountInformation(ctx echo.Context, address string, params AccountInformationParams) error
	// Get a list of unconfirmed transactions currently in the transaction pool by address.
	// (GET /v2/accounts/{address}/transactions/pending)
	GetPendingTransactionsByAddress(ctx echo.Context, address string, params GetPendingTransactionsByAddressParams) error
	// Get application information.
	// (GET /v2/applications/{application-id})
	GetApplicationByID(ctx echo.Context, applicationId uint64) error
	// Get asset information.
	// (GET /v2/assets/{asset-id})
	GetAssetByID(ctx echo.Context, assetId uint64) error
	// Get the block for the given round.
	// (GET /v2/blocks/{round})
	GetBlock(ctx echo.Context, round uint64, params GetBlockParams) error
	// Get a Merkle proof for a transaction in a block.
	// (GET /v2/blocks/{round}/transactions/{txid}/proof)
	GetProof(ctx echo.Context, round uint64, txid string, params GetProofParams) error
	// Get the current supply reported by the ledger.
	// (GET /v2/ledger/supply)
	GetSupply(ctx echo.Context) error
	// Gets the current node status.
	// (GET /v2/status)
	GetStatus(ctx echo.Context) error
	// Gets the node status after waiting for the given round.
	// (GET /v2/status/wait-for-block-after/{round})
	WaitForBlock(ctx echo.Context, round uint64) error
	// Compile TEAL source code to binary, produce its hash
	// (POST /v2/teal/compile)
	TealCompile(ctx echo.Context) error
	// Provide debugging information for a transaction (or group).
	// (POST /v2/teal/dryrun)
	TealDryrun(ctx echo.Context) error
	// Broadcasts a raw transaction to the network.
	// (POST /v2/transactions)
	RawTransaction(ctx echo.Context) error
	// Get parameters for constructing a new transaction
	// (GET /v2/transactions/params)
	TransactionParams(ctx echo.Context) error
	// Get a list of unconfirmed transactions currently in the transaction pool.
	// (GET /v2/transactions/pending)
	GetPendingTransactions(ctx echo.Context, params GetPendingTransactionsParams) error
	// Get a specific pending transaction.
	// (GET /v2/transactions/pending/{txid})
	PendingTransactionInformation(ctx echo.Context, txid string, params PendingTransactionInformationParams) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// AccountInformation converts echo context to params.
func (w *ServerInterfaceWrapper) AccountInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params AccountInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.AccountInformation(ctx, address, params)
	return err
}

// GetPendingTransactionsByAddress converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactionsByAddress(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "address" -------------
	var address string

	err = runtime.BindStyledParameter("simple", false, "address", ctx.Param("address"), &address)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter address: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsByAddressParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactionsByAddress(ctx, address, params)
	return err
}

// GetApplicationByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetApplicationByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "application-id" -------------
	var applicationId uint64

	err = runtime.BindStyledParameter("simple", false, "application-id", ctx.Param("application-id"), &applicationId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter application-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetApplicationByID(ctx, applicationId)
	return err
}

// GetAssetByID converts echo context to params.
func (w *ServerInterfaceWrapper) GetAssetByID(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "asset-id" -------------
	var assetId uint64

	err = runtime.BindStyledParameter("simple", false, "asset-id", ctx.Param("asset-id"), &assetId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter asset-id: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetAssetByID(ctx, assetId)
	return err
}

// GetBlock converts echo context to params.
func (w *ServerInterfaceWrapper) GetBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetBlockParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetBlock(ctx, round, params)
	return err
}

// GetProof converts echo context to params.
func (w *ServerInterfaceWrapper) GetProof(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetProofParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetProof(ctx, round, txid, params)
	return err
}

// GetSupply converts echo context to params.
func (w *ServerInterfaceWrapper) GetSupply(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetSupply(ctx)
	return err
}

// GetStatus converts echo context to params.
func (w *ServerInterfaceWrapper) GetStatus(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetStatus(ctx)
	return err
}

// WaitForBlock converts echo context to params.
func (w *ServerInterfaceWrapper) WaitForBlock(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "round" -------------
	var round uint64

	err = runtime.BindStyledParameter("simple", false, "round", ctx.Param("round"), &round)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter round: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.WaitForBlock(ctx, round)
	return err
}

// TealCompile converts echo context to params.
func (w *ServerInterfaceWrapper) TealCompile(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealCompile(ctx)
	return err
}

// TealDryrun converts echo context to params.
func (w *ServerInterfaceWrapper) TealDryrun(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TealDryrun(ctx)
	return err
}

// RawTransaction converts echo context to params.
func (w *ServerInterfaceWrapper) RawTransaction(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.RawTransaction(ctx)
	return err
}

// TransactionParams converts echo context to params.
func (w *ServerInterfaceWrapper) TransactionParams(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.TransactionParams(ctx)
	return err
}

// GetPendingTransactions converts echo context to params.
func (w *ServerInterfaceWrapper) GetPendingTransactions(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"max":    true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetPendingTransactionsParams
	// ------------- Optional query parameter "max" -------------
	if paramValue := ctx.QueryParam("max"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "max", ctx.QueryParams(), &params.Max)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter max: %s", err))
	}

	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetPendingTransactions(ctx, params)
	return err
}

// PendingTransactionInformation converts echo context to params.
func (w *ServerInterfaceWrapper) PendingTransactionInformation(ctx echo.Context) error {

	validQueryParams := map[string]bool{
		"pretty": true,
		"format": true,
	}

	// Check for unknown query parameters.
	for name, _ := range ctx.QueryParams() {
		if _, ok := validQueryParams[name]; !ok {
			return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Unknown parameter detected: %s", name))
		}
	}

	var err error
	// ------------- Path parameter "txid" -------------
	var txid string

	err = runtime.BindStyledParameter("simple", false, "txid", ctx.Param("txid"), &txid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter txid: %s", err))
	}

	ctx.Set("api_key.Scopes", []string{""})

	// Parameter object where we will unmarshal all parameters from the context
	var params PendingTransactionInformationParams
	// ------------- Optional query parameter "format" -------------
	if paramValue := ctx.QueryParam("format"); paramValue != "" {

	}

	err = runtime.BindQueryParameter("form", true, false, "format", ctx.QueryParams(), &params.Format)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter format: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.PendingTransactionInformation(ctx, txid, params)
	return err
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}, si ServerInterface, m ...echo.MiddlewareFunc) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.GET("/v2/accounts/:address", wrapper.AccountInformation, m...)
	router.GET("/v2/accounts/:address/transactions/pending", wrapper.GetPendingTransactionsByAddress, m...)
	router.GET("/v2/applications/:application-id", wrapper.GetApplicationByID, m...)
	router.GET("/v2/assets/:asset-id", wrapper.GetAssetByID, m...)
	router.GET("/v2/blocks/:round", wrapper.GetBlock, m...)
	router.GET("/v2/blocks/:round/transactions/:txid/proof", wrapper.GetProof, m...)
	router.GET("/v2/ledger/supply", wrapper.GetSupply, m...)
	router.GET("/v2/status", wrapper.GetStatus, m...)
	router.GET("/v2/status/wait-for-block-after/:round", wrapper.WaitForBlock, m...)
	router.POST("/v2/teal/compile", wrapper.TealCompile, m...)
	router.POST("/v2/teal/dryrun", wrapper.TealDryrun, m...)
	router.POST("/v2/transactions", wrapper.RawTransaction, m...)
	router.GET("/v2/transactions/params", wrapper.TransactionParams, m...)
	router.GET("/v2/transactions/pending", wrapper.GetPendingTransactions, m...)
	router.GET("/v2/transactions/pending/:txid", wrapper.PendingTransactionInformation, m...)

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
