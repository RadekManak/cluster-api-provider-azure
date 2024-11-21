//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservicefleet

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FleetUpdateStrategiesClient contains the methods for the FleetUpdateStrategies group.
// Don't use this type directly, use NewFleetUpdateStrategiesClient() instead.
type FleetUpdateStrategiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFleetUpdateStrategiesClient creates a new instance of FleetUpdateStrategiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFleetUpdateStrategiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FleetUpdateStrategiesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FleetUpdateStrategiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a FleetUpdateStrategy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - updateStrategyName - The name of the UpdateStrategy resource.
//   - resource - Resource create parameters.
//   - options - FleetUpdateStrategiesClientBeginCreateOrUpdateOptions contains the optional parameters for the FleetUpdateStrategiesClient.BeginCreateOrUpdate
//     method.
func (client *FleetUpdateStrategiesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, resource FleetUpdateStrategy, options *FleetUpdateStrategiesClientBeginCreateOrUpdateOptions) (*runtime.Poller[FleetUpdateStrategiesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, fleetName, updateStrategyName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FleetUpdateStrategiesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FleetUpdateStrategiesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a FleetUpdateStrategy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *FleetUpdateStrategiesClient) createOrUpdate(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, resource FleetUpdateStrategy, options *FleetUpdateStrategiesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FleetUpdateStrategiesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, fleetName, updateStrategyName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FleetUpdateStrategiesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, resource FleetUpdateStrategy, options *FleetUpdateStrategiesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateStrategies/{updateStrategyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if updateStrategyName == "" {
		return nil, errors.New("parameter updateStrategyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateStrategyName}", url.PathEscape(updateStrategyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header["If-None-Match"] = []string{*options.IfNoneMatch}
	}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a FleetUpdateStrategy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - updateStrategyName - The name of the UpdateStrategy resource.
//   - options - FleetUpdateStrategiesClientBeginDeleteOptions contains the optional parameters for the FleetUpdateStrategiesClient.BeginDelete
//     method.
func (client *FleetUpdateStrategiesClient) BeginDelete(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, options *FleetUpdateStrategiesClientBeginDeleteOptions) (*runtime.Poller[FleetUpdateStrategiesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, fleetName, updateStrategyName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FleetUpdateStrategiesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FleetUpdateStrategiesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a FleetUpdateStrategy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *FleetUpdateStrategiesClient) deleteOperation(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, options *FleetUpdateStrategiesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FleetUpdateStrategiesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, fleetName, updateStrategyName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FleetUpdateStrategiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, options *FleetUpdateStrategiesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateStrategies/{updateStrategyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if updateStrategyName == "" {
		return nil, errors.New("parameter updateStrategyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateStrategyName}", url.PathEscape(updateStrategyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	return req, nil
}

// Get - Get a FleetUpdateStrategy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - updateStrategyName - The name of the UpdateStrategy resource.
//   - options - FleetUpdateStrategiesClientGetOptions contains the optional parameters for the FleetUpdateStrategiesClient.Get
//     method.
func (client *FleetUpdateStrategiesClient) Get(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, options *FleetUpdateStrategiesClientGetOptions) (FleetUpdateStrategiesClientGetResponse, error) {
	var err error
	const operationName = "FleetUpdateStrategiesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, fleetName, updateStrategyName, options)
	if err != nil {
		return FleetUpdateStrategiesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FleetUpdateStrategiesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FleetUpdateStrategiesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FleetUpdateStrategiesClient) getCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, updateStrategyName string, options *FleetUpdateStrategiesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateStrategies/{updateStrategyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	if updateStrategyName == "" {
		return nil, errors.New("parameter updateStrategyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{updateStrategyName}", url.PathEscape(updateStrategyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FleetUpdateStrategiesClient) getHandleResponse(resp *http.Response) (FleetUpdateStrategiesClientGetResponse, error) {
	result := FleetUpdateStrategiesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetUpdateStrategy); err != nil {
		return FleetUpdateStrategiesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByFleetPager - List FleetUpdateStrategy resources by Fleet
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - fleetName - The name of the Fleet resource.
//   - options - FleetUpdateStrategiesClientListByFleetOptions contains the optional parameters for the FleetUpdateStrategiesClient.NewListByFleetPager
//     method.
func (client *FleetUpdateStrategiesClient) NewListByFleetPager(resourceGroupName string, fleetName string, options *FleetUpdateStrategiesClientListByFleetOptions) *runtime.Pager[FleetUpdateStrategiesClientListByFleetResponse] {
	return runtime.NewPager(runtime.PagingHandler[FleetUpdateStrategiesClientListByFleetResponse]{
		More: func(page FleetUpdateStrategiesClientListByFleetResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FleetUpdateStrategiesClientListByFleetResponse) (FleetUpdateStrategiesClientListByFleetResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FleetUpdateStrategiesClient.NewListByFleetPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByFleetCreateRequest(ctx, resourceGroupName, fleetName, options)
			}, nil)
			if err != nil {
				return FleetUpdateStrategiesClientListByFleetResponse{}, err
			}
			return client.listByFleetHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByFleetCreateRequest creates the ListByFleet request.
func (client *FleetUpdateStrategiesClient) listByFleetCreateRequest(ctx context.Context, resourceGroupName string, fleetName string, options *FleetUpdateStrategiesClientListByFleetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/fleets/{fleetName}/updateStrategies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if fleetName == "" {
		return nil, errors.New("parameter fleetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fleetName}", url.PathEscape(fleetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByFleetHandleResponse handles the ListByFleet response.
func (client *FleetUpdateStrategiesClient) listByFleetHandleResponse(resp *http.Response) (FleetUpdateStrategiesClientListByFleetResponse, error) {
	result := FleetUpdateStrategiesClientListByFleetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FleetUpdateStrategyListResult); err != nil {
		return FleetUpdateStrategiesClientListByFleetResponse{}, err
	}
	return result, nil
}