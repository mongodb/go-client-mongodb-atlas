/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type SharedTierRestoreJobsApi interface {

	/*
	CreateOneRestoreJobFromOneM2OrM5Cluster Create One Restore Job from One M2 or M5 Cluster

	Restores the specified cluster. MongoDB Cloud limits which clusters can be the target clusters of a restore. The target cluster can't use encryption at rest, run a major release MongoDB version different than the snapshot, or receive client requests during restores. MongoDB Cloud deletes all existing data on the target cluster prior to the restore operation. To use this resource, the requesting API Key must have the Atlas Project Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest
	*/
	CreateOneRestoreJobFromOneM2OrM5Cluster(ctx context.Context, clusterName string, groupId string) SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest

	// CreateOneRestoreJobFromOneM2OrM5ClusterExecute executes the request
	//  @return TenantRestore
	CreateOneRestoreJobFromOneM2OrM5ClusterExecute(r SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest) (*TenantRestore, *http.Response, error)

	/*
	ReturnAllRestoreJobsForOneM2OrM5Cluster Return All Restore Jobs for One M2 or M5 Cluster

	Returns all restore jobs for the specified M2 or M5 cluster. Restore jobs restore a cluster using a snapshot. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest
	*/
	ReturnAllRestoreJobsForOneM2OrM5Cluster(ctx context.Context, clusterName string, groupId string) SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest

	// ReturnAllRestoreJobsForOneM2OrM5ClusterExecute executes the request
	//  @return PaginatedTenantRestoreView
	ReturnAllRestoreJobsForOneM2OrM5ClusterExecute(r SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest) (*PaginatedTenantRestoreView, *http.Response, error)

	/*
	ReturnOneRestoreJobForOneM2OrM5Cluster Return One Restore Job for One M2 or M5 Cluster

	Returns the specified restore job. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param restoreId Unique 24-hexadecimal digit string that identifies the restore job to return.
	@return SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest
	*/
	ReturnOneRestoreJobForOneM2OrM5Cluster(ctx context.Context, clusterName string, groupId string, restoreId string) SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest

	// ReturnOneRestoreJobForOneM2OrM5ClusterExecute executes the request
	//  @return TenantRestore
	ReturnOneRestoreJobForOneM2OrM5ClusterExecute(r SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest) (*TenantRestore, *http.Response, error)
}

// SharedTierRestoreJobsApiService SharedTierRestoreJobsApi service
type SharedTierRestoreJobsApiService service

type SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest struct {
	ctx context.Context
	ApiService SharedTierRestoreJobsApi
	clusterName string
	groupId string
	tenantRestore *TenantRestore
	envelope *bool
	pretty *bool
}

// The restore job details.
func (r SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest) TenantRestore(tenantRestore TenantRestore) SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest {
	r.tenantRestore = &tenantRestore
	return r
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest) Envelope(envelope bool) SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/Prettyprint\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;prettyprint&lt;/a&gt; format.
func (r SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest) Pretty(pretty bool) SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest {
	r.pretty = &pretty
	return r
}

func (r SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest) Execute() (*TenantRestore, *http.Response, error) {
	return r.ApiService.CreateOneRestoreJobFromOneM2OrM5ClusterExecute(r)
}

/*
CreateOneRestoreJobFromOneM2OrM5Cluster Create One Restore Job from One M2 or M5 Cluster

Restores the specified cluster. MongoDB Cloud limits which clusters can be the target clusters of a restore. The target cluster can't use encryption at rest, run a major release MongoDB version different than the snapshot, or receive client requests during restores. MongoDB Cloud deletes all existing data on the target cluster prior to the restore operation. To use this resource, the requesting API Key must have the Atlas Project Admin role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterName Human-readable label that identifies the cluster.
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest
*/
func (a *SharedTierRestoreJobsApiService) CreateOneRestoreJobFromOneM2OrM5Cluster(ctx context.Context, clusterName string, groupId string) SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest {
	return SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest{
		ApiService: a,
		ctx: ctx,
		clusterName: clusterName,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return TenantRestore
func (a *SharedTierRestoreJobsApiService) CreateOneRestoreJobFromOneM2OrM5ClusterExecute(r SharedTierRestoreJobsApiCreateOneRestoreJobFromOneM2OrM5ClusterRequest) (*TenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.CreateOneRestoreJobFromOneM2OrM5Cluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restore"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if r.tenantRestore == nil {
		return localVarReturnValue, nil, reportError("tenantRestore is required and must be specified")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.tenantRestore
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest struct {
	ctx context.Context
	ApiService SharedTierRestoreJobsApi
	clusterName string
	groupId string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest) Envelope(envelope bool) SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/Prettyprint\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;prettyprint&lt;/a&gt; format.
func (r SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest) Pretty(pretty bool) SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest {
	r.pretty = &pretty
	return r
}

func (r SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest) Execute() (*PaginatedTenantRestoreView, *http.Response, error) {
	return r.ApiService.ReturnAllRestoreJobsForOneM2OrM5ClusterExecute(r)
}

/*
ReturnAllRestoreJobsForOneM2OrM5Cluster Return All Restore Jobs for One M2 or M5 Cluster

Returns all restore jobs for the specified M2 or M5 cluster. Restore jobs restore a cluster using a snapshot. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterName Human-readable label that identifies the cluster.
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest
*/
func (a *SharedTierRestoreJobsApiService) ReturnAllRestoreJobsForOneM2OrM5Cluster(ctx context.Context, clusterName string, groupId string) SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest {
	return SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest{
		ApiService: a,
		ctx: ctx,
		clusterName: clusterName,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return PaginatedTenantRestoreView
func (a *SharedTierRestoreJobsApiService) ReturnAllRestoreJobsForOneM2OrM5ClusterExecute(r SharedTierRestoreJobsApiReturnAllRestoreJobsForOneM2OrM5ClusterRequest) (*PaginatedTenantRestoreView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedTenantRestoreView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.ReturnAllRestoreJobsForOneM2OrM5Cluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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

type SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest struct {
	ctx context.Context
	ApiService SharedTierRestoreJobsApi
	clusterName string
	groupId string
	restoreId string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest) Envelope(envelope bool) SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/Prettyprint\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;prettyprint&lt;/a&gt; format.
func (r SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest) Pretty(pretty bool) SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest {
	r.pretty = &pretty
	return r
}

func (r SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest) Execute() (*TenantRestore, *http.Response, error) {
	return r.ApiService.ReturnOneRestoreJobForOneM2OrM5ClusterExecute(r)
}

/*
ReturnOneRestoreJobForOneM2OrM5Cluster Return One Restore Job for One M2 or M5 Cluster

Returns the specified restore job. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterName Human-readable label that identifies the cluster.
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @param restoreId Unique 24-hexadecimal digit string that identifies the restore job to return.
 @return SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest
*/
func (a *SharedTierRestoreJobsApiService) ReturnOneRestoreJobForOneM2OrM5Cluster(ctx context.Context, clusterName string, groupId string, restoreId string) SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest {
	return SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest{
		ApiService: a,
		ctx: ctx,
		clusterName: clusterName,
		groupId: groupId,
		restoreId: restoreId,
	}
}

// Execute executes the request
//  @return TenantRestore
func (a *SharedTierRestoreJobsApiService) ReturnOneRestoreJobForOneM2OrM5ClusterExecute(r SharedTierRestoreJobsApiReturnOneRestoreJobForOneM2OrM5ClusterRequest) (*TenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierRestoreJobsApiService.ReturnOneRestoreJobForOneM2OrM5Cluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/restores/{restoreId}"
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"restoreId"+"}", url.PathEscape(parameterToString(r.restoreId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.restoreId) < 24 {
		return localVarReturnValue, nil, reportError("restoreId must have at least 24 elements")
	}
	if strlen(r.restoreId) > 24 {
		return localVarReturnValue, nil, reportError("restoreId must have less than 24 elements")
	}

	if r.envelope != nil {
		localVarQueryParams.Add("envelope", parameterToString(*r.envelope, ""))
	}
	if r.pretty != nil {
		localVarQueryParams.Add("pretty", parameterToString(*r.pretty, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json", "application/json"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v ApiError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
            		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
            		newErr.model = v
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