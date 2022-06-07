/*
Qovery API

- Qovery is the fastest way to deploy your full-stack apps on any Cloud provider. - ℹ️ The API is stable and still in development.

API version: 1.0.3
Contact: support+api+documentation@qovery.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package qovery

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// ContainerDeploymentRestrictionApiService ContainerDeploymentRestrictionApi service
type ContainerDeploymentRestrictionApiService service

type ApiCreateContainerDeploymentRestrictionRequest struct {
	ctx                                   context.Context
	ApiService                            *ContainerDeploymentRestrictionApiService
	containerId                           string
	containerDeploymentRestrictionRequest *ContainerDeploymentRestrictionRequest
}

func (r ApiCreateContainerDeploymentRestrictionRequest) ContainerDeploymentRestrictionRequest(containerDeploymentRestrictionRequest ContainerDeploymentRestrictionRequest) ApiCreateContainerDeploymentRestrictionRequest {
	r.containerDeploymentRestrictionRequest = &containerDeploymentRestrictionRequest
	return r
}

func (r ApiCreateContainerDeploymentRestrictionRequest) Execute() (*ContainerDeploymentRestriction, *http.Response, error) {
	return r.ApiService.CreateContainerDeploymentRestrictionExecute(r)
}

/*
CreateContainerDeploymentRestriction Create an container deployment restriction

Create an container deployment restriction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @return ApiCreateContainerDeploymentRestrictionRequest
*/
func (a *ContainerDeploymentRestrictionApiService) CreateContainerDeploymentRestriction(ctx context.Context, containerId string) ApiCreateContainerDeploymentRestrictionRequest {
	return ApiCreateContainerDeploymentRestrictionRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
	}
}

// Execute executes the request
//  @return ContainerDeploymentRestriction
func (a *ContainerDeploymentRestrictionApiService) CreateContainerDeploymentRestrictionExecute(r ApiCreateContainerDeploymentRestrictionRequest) (*ContainerDeploymentRestriction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ContainerDeploymentRestriction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerDeploymentRestrictionApiService.CreateContainerDeploymentRestriction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/deploymentRestriction"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.containerDeploymentRestrictionRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteContainerDeploymentRestrictionRequest struct {
	ctx         context.Context
	ApiService  *ContainerDeploymentRestrictionApiService
	containerId string
}

func (r ApiDeleteContainerDeploymentRestrictionRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteContainerDeploymentRestrictionExecute(r)
}

/*
DeleteContainerDeploymentRestriction Delete a container deployment restriction

Delete a container deployment restriction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @return ApiDeleteContainerDeploymentRestrictionRequest
*/
func (a *ContainerDeploymentRestrictionApiService) DeleteContainerDeploymentRestriction(ctx context.Context, containerId string) ApiDeleteContainerDeploymentRestrictionRequest {
	return ApiDeleteContainerDeploymentRestrictionRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
	}
}

// Execute executes the request
func (a *ContainerDeploymentRestrictionApiService) DeleteContainerDeploymentRestrictionExecute(r ApiDeleteContainerDeploymentRestrictionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerDeploymentRestrictionApiService.DeleteContainerDeploymentRestriction")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/deploymentRestriction/{deploymentRestrictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiEditContainerDeploymentRestrictionRequest struct {
	ctx                                   context.Context
	ApiService                            *ContainerDeploymentRestrictionApiService
	containerId                           string
	deploymentRestrictionId               string
	containerDeploymentRestrictionRequest *ContainerDeploymentRestrictionRequest
}

func (r ApiEditContainerDeploymentRestrictionRequest) ContainerDeploymentRestrictionRequest(containerDeploymentRestrictionRequest ContainerDeploymentRestrictionRequest) ApiEditContainerDeploymentRestrictionRequest {
	r.containerDeploymentRestrictionRequest = &containerDeploymentRestrictionRequest
	return r
}

func (r ApiEditContainerDeploymentRestrictionRequest) Execute() (*ContainerDeploymentRestriction, *http.Response, error) {
	return r.ApiService.EditContainerDeploymentRestrictionExecute(r)
}

/*
EditContainerDeploymentRestriction Edit a container deployment restriction

Edit a container deployment restriction

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @param deploymentRestrictionId Deployment Restriction ID
 @return ApiEditContainerDeploymentRestrictionRequest
*/
func (a *ContainerDeploymentRestrictionApiService) EditContainerDeploymentRestriction(ctx context.Context, containerId string, deploymentRestrictionId string) ApiEditContainerDeploymentRestrictionRequest {
	return ApiEditContainerDeploymentRestrictionRequest{
		ApiService:              a,
		ctx:                     ctx,
		containerId:             containerId,
		deploymentRestrictionId: deploymentRestrictionId,
	}
}

// Execute executes the request
//  @return ContainerDeploymentRestriction
func (a *ContainerDeploymentRestrictionApiService) EditContainerDeploymentRestrictionExecute(r ApiEditContainerDeploymentRestrictionRequest) (*ContainerDeploymentRestriction, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ContainerDeploymentRestriction
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerDeploymentRestrictionApiService.EditContainerDeploymentRestriction")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/deploymentRestriction/{deploymentRestrictionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"deploymentRestrictionId"+"}", url.PathEscape(parameterToString(r.deploymentRestrictionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.containerDeploymentRestrictionRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetContainerDeploymentRestrictionsRequest struct {
	ctx         context.Context
	ApiService  *ContainerDeploymentRestrictionApiService
	containerId string
}

func (r ApiGetContainerDeploymentRestrictionsRequest) Execute() (*ContainerDeploymentRestrictionResponseList, *http.Response, error) {
	return r.ApiService.GetContainerDeploymentRestrictionsExecute(r)
}

/*
GetContainerDeploymentRestrictions Get container deployment restrictions

Get container deployment restrictions

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param containerId Container ID
 @return ApiGetContainerDeploymentRestrictionsRequest
*/
func (a *ContainerDeploymentRestrictionApiService) GetContainerDeploymentRestrictions(ctx context.Context, containerId string) ApiGetContainerDeploymentRestrictionsRequest {
	return ApiGetContainerDeploymentRestrictionsRequest{
		ApiService:  a,
		ctx:         ctx,
		containerId: containerId,
	}
}

// Execute executes the request
//  @return ContainerDeploymentRestrictionResponseList
func (a *ContainerDeploymentRestrictionApiService) GetContainerDeploymentRestrictionsExecute(r ApiGetContainerDeploymentRestrictionsRequest) (*ContainerDeploymentRestrictionResponseList, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ContainerDeploymentRestrictionResponseList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ContainerDeploymentRestrictionApiService.GetContainerDeploymentRestrictions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/container/{containerId}/deploymentRestriction"
	localVarPath = strings.Replace(localVarPath, "{"+"containerId"+"}", url.PathEscape(parameterToString(r.containerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}