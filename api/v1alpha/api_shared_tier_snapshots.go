/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas. To learn more, review the [Administration API overview](https://www.mongodb.com/docs/atlas/api/atlas-admin-api/). This OpenAPI specification covers all of the collections with the exception of Alerts, Alert Configurations, and Events. Refer to the [legacy documentation](https://www.mongodb.com/docs/atlas/reference/api-resources/) for the specifications of these resources.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1alpha

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


type SharedTierSnapshotsApi interface {

	/*
	DownloadSharedClusterBackup Download One M2 or M5 Cluster Snapshot

	Requests one snapshot for the specified shared cluster. This resource returns a `snapshotURL` that you can use to download the snapshot. This `snapshotURL` remains active for four hours after you make the request. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param clusterName Human-readable label that identifies the cluster.
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return SharedTierSnapshotsApiDownloadSharedClusterBackupRequest
	*/
	DownloadSharedClusterBackup(ctx context.Context, clusterName string, groupId string) SharedTierSnapshotsApiDownloadSharedClusterBackupRequest

	// DownloadSharedClusterBackupExecute executes the request
	//  @return TenantRestore
	DownloadSharedClusterBackupExecute(r SharedTierSnapshotsApiDownloadSharedClusterBackupRequest) (*TenantRestore, *http.Response, error)

	/*
	GetSharedClusterBackup Return One Snapshot for One M2 or M5 Cluster

	Returns details for one snapshot for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
	@return SharedTierSnapshotsApiGetSharedClusterBackupRequest
	*/
	GetSharedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) SharedTierSnapshotsApiGetSharedClusterBackupRequest

	// GetSharedClusterBackupExecute executes the request
	//  @return TenantSnapshot
	GetSharedClusterBackupExecute(r SharedTierSnapshotsApiGetSharedClusterBackupRequest) (*TenantSnapshot, *http.Response, error)

	/*
	ListSharedClusterBackups Return All Snapshots for One M2 or M5 Cluster

	Returns details for all snapshots for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return SharedTierSnapshotsApiListSharedClusterBackupsRequest
	*/
	ListSharedClusterBackups(ctx context.Context, groupId string, clusterName string) SharedTierSnapshotsApiListSharedClusterBackupsRequest

	// ListSharedClusterBackupsExecute executes the request
	//  @return PaginatedTenantSnapshotView
	ListSharedClusterBackupsExecute(r SharedTierSnapshotsApiListSharedClusterBackupsRequest) (*PaginatedTenantSnapshotView, *http.Response, error)
}

// SharedTierSnapshotsApiService SharedTierSnapshotsApi service
type SharedTierSnapshotsApiService service

type SharedTierSnapshotsApiDownloadSharedClusterBackupRequest struct {
	ctx context.Context
	ApiService SharedTierSnapshotsApi
	clusterName string
	groupId string
	tenantRestore *TenantRestore
	envelope *bool
	pretty *bool
}

// Snapshot to be downloaded.
func (r SharedTierSnapshotsApiDownloadSharedClusterBackupRequest) TenantRestore(tenantRestore TenantRestore) SharedTierSnapshotsApiDownloadSharedClusterBackupRequest {
	r.tenantRestore = &tenantRestore
	return r
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r SharedTierSnapshotsApiDownloadSharedClusterBackupRequest) Envelope(envelope bool) SharedTierSnapshotsApiDownloadSharedClusterBackupRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/Prettyprint\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;prettyprint&lt;/a&gt; format.
func (r SharedTierSnapshotsApiDownloadSharedClusterBackupRequest) Pretty(pretty bool) SharedTierSnapshotsApiDownloadSharedClusterBackupRequest {
	r.pretty = &pretty
	return r
}

func (r SharedTierSnapshotsApiDownloadSharedClusterBackupRequest) Execute() (*TenantRestore, *http.Response, error) {
	return r.ApiService.DownloadSharedClusterBackupExecute(r)
}

/*
DownloadSharedClusterBackup Download One M2 or M5 Cluster Snapshot

Requests one snapshot for the specified shared cluster. This resource returns a `snapshotURL` that you can use to download the snapshot. This `snapshotURL` remains active for four hours after you make the request. To use this resource, the requesting API Key must have the Project Atlas Admin role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param clusterName Human-readable label that identifies the cluster.
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return SharedTierSnapshotsApiDownloadSharedClusterBackupRequest
*/
func (a *SharedTierSnapshotsApiService) DownloadSharedClusterBackup(ctx context.Context, clusterName string, groupId string) SharedTierSnapshotsApiDownloadSharedClusterBackupRequest {
	return SharedTierSnapshotsApiDownloadSharedClusterBackupRequest{
		ApiService: a,
		ctx: ctx,
		clusterName: clusterName,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return TenantRestore
func (a *SharedTierSnapshotsApiService) DownloadSharedClusterBackupExecute(r SharedTierSnapshotsApiDownloadSharedClusterBackupRequest) (*TenantRestore, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantRestore
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.DownloadSharedClusterBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/download"
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
		if localVarHTTPResponse.StatusCode == 401 {
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

type SharedTierSnapshotsApiGetSharedClusterBackupRequest struct {
	ctx context.Context
	ApiService SharedTierSnapshotsApi
	groupId string
	clusterName string
	snapshotId string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r SharedTierSnapshotsApiGetSharedClusterBackupRequest) Envelope(envelope bool) SharedTierSnapshotsApiGetSharedClusterBackupRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/Prettyprint\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;prettyprint&lt;/a&gt; format.
func (r SharedTierSnapshotsApiGetSharedClusterBackupRequest) Pretty(pretty bool) SharedTierSnapshotsApiGetSharedClusterBackupRequest {
	r.pretty = &pretty
	return r
}

func (r SharedTierSnapshotsApiGetSharedClusterBackupRequest) Execute() (*TenantSnapshot, *http.Response, error) {
	return r.ApiService.GetSharedClusterBackupExecute(r)
}

/*
GetSharedClusterBackup Return One Snapshot for One M2 or M5 Cluster

Returns details for one snapshot for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @param clusterName Human-readable label that identifies the cluster.
 @param snapshotId Unique 24-hexadecimal digit string that identifies the desired snapshot.
 @return SharedTierSnapshotsApiGetSharedClusterBackupRequest
*/
func (a *SharedTierSnapshotsApiService) GetSharedClusterBackup(ctx context.Context, groupId string, clusterName string, snapshotId string) SharedTierSnapshotsApiGetSharedClusterBackupRequest {
	return SharedTierSnapshotsApiGetSharedClusterBackupRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		clusterName: clusterName,
		snapshotId: snapshotId,
	}
}

// Execute executes the request
//  @return TenantSnapshot
func (a *SharedTierSnapshotsApiService) GetSharedClusterBackupExecute(r SharedTierSnapshotsApiGetSharedClusterBackupRequest) (*TenantSnapshot, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TenantSnapshot
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.GetSharedClusterBackup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots/{snapshotId}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snapshotId"+"}", url.PathEscape(parameterToString(r.snapshotId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
	}
	if strlen(r.snapshotId) < 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have at least 24 elements")
	}
	if strlen(r.snapshotId) > 24 {
		return localVarReturnValue, nil, reportError("snapshotId must have less than 24 elements")
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

type SharedTierSnapshotsApiListSharedClusterBackupsRequest struct {
	ctx context.Context
	ApiService SharedTierSnapshotsApi
	groupId string
	clusterName string
	envelope *bool
	pretty *bool
}

// Flag that indicates whether Application wraps the response in an &#x60;envelope&#x60; JSON object. Some API clients cannot access the HTTP response headers or status code. To remediate this, set envelope&#x3D;true in the query. Endpoints that return a list of results use the results object as an envelope. Application adds the status parameter to the response body.
func (r SharedTierSnapshotsApiListSharedClusterBackupsRequest) Envelope(envelope bool) SharedTierSnapshotsApiListSharedClusterBackupsRequest {
	r.envelope = &envelope
	return r
}

// Flag that indicates whether the response body should be in the &lt;a href&#x3D;\&quot;https://en.wikipedia.org/wiki/Prettyprint\&quot; target&#x3D;\&quot;_blank\&quot; rel&#x3D;\&quot;noopener noreferrer\&quot;&gt;prettyprint&lt;/a&gt; format.
func (r SharedTierSnapshotsApiListSharedClusterBackupsRequest) Pretty(pretty bool) SharedTierSnapshotsApiListSharedClusterBackupsRequest {
	r.pretty = &pretty
	return r
}

func (r SharedTierSnapshotsApiListSharedClusterBackupsRequest) Execute() (*PaginatedTenantSnapshotView, *http.Response, error) {
	return r.ApiService.ListSharedClusterBackupsExecute(r)
}

/*
ListSharedClusterBackups Return All Snapshots for One M2 or M5 Cluster

Returns details for all snapshots for the specified shared cluster. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @param clusterName Human-readable label that identifies the cluster.
 @return SharedTierSnapshotsApiListSharedClusterBackupsRequest
*/
func (a *SharedTierSnapshotsApiService) ListSharedClusterBackups(ctx context.Context, groupId string, clusterName string) SharedTierSnapshotsApiListSharedClusterBackupsRequest {
	return SharedTierSnapshotsApiListSharedClusterBackupsRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		clusterName: clusterName,
	}
}

// Execute executes the request
//  @return PaginatedTenantSnapshotView
func (a *SharedTierSnapshotsApiService) ListSharedClusterBackupsExecute(r SharedTierSnapshotsApiListSharedClusterBackupsRequest) (*PaginatedTenantSnapshotView, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedTenantSnapshotView
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SharedTierSnapshotsApiService.ListSharedClusterBackups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/backup/tenant/snapshots"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterToString(r.groupId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(parameterToString(r.clusterName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}
	if strlen(r.clusterName) < 1 {
		return localVarReturnValue, nil, reportError("clusterName must have at least 1 elements")
	}
	if strlen(r.clusterName) > 64 {
		return localVarReturnValue, nil, reportError("clusterName must have less than 64 elements")
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