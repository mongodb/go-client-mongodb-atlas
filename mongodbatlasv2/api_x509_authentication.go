/*
MongoDB Atlas Administration API

The MongoDB Atlas Administration API allows developers to manage all components in MongoDB Atlas.   The Atlas Administration API authenticates using HTTP Digest Authentication. Provide a programmatic API public key and corresponding private key as the username and password when constructing the HTTP request. For example, with [curl](https://en.wikipedia.org/wiki/CURL): `curl --user \"{PUBLIC-KEY}:{PRIVATE-KEY}\" --digest`   To learn more, see [Get Started with the Atlas Administration API](https://www.mongodb.com/docs/atlas/configure-api-access/). For support, see [MongoDB Support](https://www.mongodb.com/support/get-started)

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mongodbatlasv2

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type X509AuthenticationApi interface {

	/*
	CreateDatabaseUserCertificate Create One X.509 Certificate for One MongoDB User

	Generates one X.509 certificate for the specified MongoDB user. Atlas manages the certificate and MongoDB user that belong to one project. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

To get MongoDB Cloud to generate a managed certificate for a database user, set `"x509Type" : "MANAGED"` on the desired MongoDB Database User.

If you are managing your own Certificate Authority (CA) in Self-Managed X.509 mode, you must generate certificates for database users using your own CA.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param username Human-readable label that represents the MongoDB database user account for whom to create a certificate.
	@return X509AuthenticationApiCreateDatabaseUserCertificateRequest
	*/
	CreateDatabaseUserCertificate(ctx context.Context, groupId string, username string) X509AuthenticationApiCreateDatabaseUserCertificateRequest

	// CreateDatabaseUserCertificateExecute executes the request
	CreateDatabaseUserCertificateExecute(r X509AuthenticationApiCreateDatabaseUserCertificateRequest) (*http.Response, error)

	/*
	DisableCustomerManagedX509 Disable Customer-Managed X.509

	Clears the customer-managed X.509 settings on a project, including the uploaded Certificate Authority, which disables self-managed X.509.

 Updating this configuration triggers a rolling restart of the database. You must have the Project Owner role to use this endpoint. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return X509AuthenticationApiDisableCustomerManagedX509Request
	*/
	DisableCustomerManagedX509(ctx context.Context, groupId string) X509AuthenticationApiDisableCustomerManagedX509Request

	// DisableCustomerManagedX509Execute executes the request
	//  @return UserSecurity
	DisableCustomerManagedX509Execute(r X509AuthenticationApiDisableCustomerManagedX509Request) (*UserSecurity, *http.Response, error)

	/*
	ListDatabaseUserCertificates Return All X.509 Certificates Assigned to One MongoDB User

	Returns all unexpired X.509 certificates for the specified MongoDB user. This MongoDB user belongs to one project. Atlas manages these certificates and the MongoDB user. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param username Human-readable label that represents the MongoDB database user account whose certificates you want to return.
	@return X509AuthenticationApiListDatabaseUserCertificatesRequest
	*/
	ListDatabaseUserCertificates(ctx context.Context, groupId string, username string) X509AuthenticationApiListDatabaseUserCertificatesRequest

	// ListDatabaseUserCertificatesExecute executes the request
	//  @return PaginatedUserCert
	ListDatabaseUserCertificatesExecute(r X509AuthenticationApiListDatabaseUserCertificatesRequest) (*PaginatedUserCert, *http.Response, error)
}

// X509AuthenticationApiService X509AuthenticationApi service
type X509AuthenticationApiService service

type X509AuthenticationApiCreateDatabaseUserCertificateRequest struct {
	ctx context.Context
	ApiService X509AuthenticationApi
	groupId string
	username string
	userCert *UserCert
}
type X509AuthenticationApiCreateDatabaseUserCertificateQueryParams struct {
		groupId string
		username string
		userCert *UserCert
}

// Generates one X.509 certificate for the specified MongoDB user.
func (r X509AuthenticationApiCreateDatabaseUserCertificateRequest) UserCert(userCert UserCert) X509AuthenticationApiCreateDatabaseUserCertificateRequest {
	r.userCert = &userCert
	return r
}

func (r X509AuthenticationApiCreateDatabaseUserCertificateRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateDatabaseUserCertificateExecute(r)
}

/*
CreateDatabaseUserCertificate Create One X.509 Certificate for One MongoDB User

Generates one X.509 certificate for the specified MongoDB user. Atlas manages the certificate and MongoDB user that belong to one project. To use this resource, the requesting API Key must have the Project Owner role. This resource doesn't require the API Key to have an Access List.

To get MongoDB Cloud to generate a managed certificate for a database user, set `"x509Type" : "MANAGED"` on the desired MongoDB Database User.

If you are managing your own Certificate Authority (CA) in Self-Managed X.509 mode, you must generate certificates for database users using your own CA.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @param username Human-readable label that represents the MongoDB database user account for whom to create a certificate.
 @return X509AuthenticationApiCreateDatabaseUserCertificateRequest
*/
func (a *X509AuthenticationApiService) CreateDatabaseUserCertificate(ctx context.Context, groupId string, username string) X509AuthenticationApiCreateDatabaseUserCertificateRequest {
	return X509AuthenticationApiCreateDatabaseUserCertificateRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		username: username,
	}
}

// Execute executes the request
func (a *X509AuthenticationApiService) CreateDatabaseUserCertificateExecute(r X509AuthenticationApiCreateDatabaseUserCertificateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "X509AuthenticationApiService.CreateDatabaseUserCertificate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return nil, reportError("groupId must have less than 24 elements")
	}
	if r.userCert == nil {
		return nil, reportError("userCert is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

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
	localVarPostBody = r.userCert
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type X509AuthenticationApiDisableCustomerManagedX509Request struct {
	ctx context.Context
	ApiService X509AuthenticationApi
	groupId string
}
type X509AuthenticationApiDisableCustomerManagedX509QueryParams struct {
		groupId string
}

func (r X509AuthenticationApiDisableCustomerManagedX509Request) Execute() (*UserSecurity, *http.Response, error) {
	return r.ApiService.DisableCustomerManagedX509Execute(r)
}

/*
DisableCustomerManagedX509 Disable Customer-Managed X.509

Clears the customer-managed X.509 settings on a project, including the uploaded Certificate Authority, which disables self-managed X.509.

 Updating this configuration triggers a rolling restart of the database. You must have the Project Owner role to use this endpoint. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @return X509AuthenticationApiDisableCustomerManagedX509Request
*/
func (a *X509AuthenticationApiService) DisableCustomerManagedX509(ctx context.Context, groupId string) X509AuthenticationApiDisableCustomerManagedX509Request {
	return X509AuthenticationApiDisableCustomerManagedX509Request{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
	}
}

// Execute executes the request
//  @return UserSecurity
func (a *X509AuthenticationApiService) DisableCustomerManagedX509Execute(r X509AuthenticationApiDisableCustomerManagedX509Request) (*UserSecurity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserSecurity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "X509AuthenticationApiService.DisableCustomerManagedX509")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/userSecurity/customerX509"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
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

type X509AuthenticationApiListDatabaseUserCertificatesRequest struct {
	ctx context.Context
	ApiService X509AuthenticationApi
	groupId string
	username string
	includeCount *bool
	itemsPerPage *int32
	pageNum *int32
}
type X509AuthenticationApiListDatabaseUserCertificatesQueryParams struct {
		groupId string
		username string
		includeCount *bool
		itemsPerPage *int32
		pageNum *int32
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r X509AuthenticationApiListDatabaseUserCertificatesRequest) IncludeCount(includeCount bool) X509AuthenticationApiListDatabaseUserCertificatesRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r X509AuthenticationApiListDatabaseUserCertificatesRequest) ItemsPerPage(itemsPerPage int32) X509AuthenticationApiListDatabaseUserCertificatesRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r X509AuthenticationApiListDatabaseUserCertificatesRequest) PageNum(pageNum int32) X509AuthenticationApiListDatabaseUserCertificatesRequest {
	r.pageNum = &pageNum
	return r
}

func (r X509AuthenticationApiListDatabaseUserCertificatesRequest) Execute() (*PaginatedUserCert, *http.Response, error) {
	return r.ApiService.ListDatabaseUserCertificatesExecute(r)
}

/*
ListDatabaseUserCertificates Return All X.509 Certificates Assigned to One MongoDB User

Returns all unexpired X.509 certificates for the specified MongoDB user. This MongoDB user belongs to one project. Atlas manages these certificates and the MongoDB user. To use this resource, the requesting API Key must have the Project Read Only role. This resource doesn't require the API Key to have an Access List.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
 @param username Human-readable label that represents the MongoDB database user account whose certificates you want to return.
 @return X509AuthenticationApiListDatabaseUserCertificatesRequest
*/
func (a *X509AuthenticationApiService) ListDatabaseUserCertificates(ctx context.Context, groupId string, username string) X509AuthenticationApiListDatabaseUserCertificatesRequest {
	return X509AuthenticationApiListDatabaseUserCertificatesRequest{
		ApiService: a,
		ctx: ctx,
		groupId: groupId,
		username: username,
	}
}

// Execute executes the request
//  @return PaginatedUserCert
func (a *X509AuthenticationApiService) ListDatabaseUserCertificatesExecute(r X509AuthenticationApiListDatabaseUserCertificatesRequest) (*PaginatedUserCert, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PaginatedUserCert
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "X509AuthenticationApiService.ListDatabaseUserCertificates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/databaseUsers/{username}/certs"
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(parameterValueToString(r.groupId, "groupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.groupId) < 24 {
		return localVarReturnValue, nil, reportError("groupId must have at least 24 elements")
	}
	if strlen(r.groupId) > 24 {
		return localVarReturnValue, nil, reportError("groupId must have less than 24 elements")
	}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int32 = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int32 = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
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

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		var v Error
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, localVarHTTPMethod, localVarPath, v)
		newErr.model = v
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