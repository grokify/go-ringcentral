/*
 * Ringcentral API
 *
 * RingCentral Connect Platform API
 *
 * OpenAPI spec version: v1.0
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package ringcentral

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

type AddressBookApi struct {
	Configuration *Configuration
}

func NewAddressBookApi() *AddressBookApi {
	configuration := NewConfiguration()
	return &AddressBookApi{
		Configuration: configuration,
	}
}

func NewAddressBookApiWithBasePath(basePath string) *AddressBookApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &AddressBookApi{
		Configuration: configuration,
	}
}

/**
 *
 * Delete Contact by ID
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param contactId Internal identifier of a contact record in the RingCentral database
 * @return void
 */
func (a AddressBookApi) RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdDelete(accountId string, extensionId string, contactId string) (*APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Delete")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", fmt.Sprintf("%v", contactId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != "" {
		localVarHeaderParams["Authorization"] = "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdDelete", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return localVarAPIResponse, err
	}
	return localVarAPIResponse, err
}

/**
 *
 * Get Contact by ID
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param contactId Internal identifier of a contact record in the RingCentral database
 * @return *PersonalContactInfo
 */
func (a AddressBookApi) RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdGet(accountId string, extensionId string, contactId string) (*PersonalContactInfo, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", fmt.Sprintf("%v", contactId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != "" {
		localVarHeaderParams["Authorization"] = "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(PersonalContactInfo)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 *
 * Update Contact by ID
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param contactId Internal identifier of a contact record in the RingCentral database
 * @param body
 * @return *PersonalContactInfo
 */
func (a AddressBookApi) RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdPut(accountId string, extensionId string, contactId string, body PersonalContactInfo) (*PersonalContactInfo, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Put")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact/{contactId}"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"contactId"+"}", fmt.Sprintf("%v", contactId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != "" {
		localVarHeaderParams["Authorization"] = "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(PersonalContactInfo)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactContactIdPut", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 *
 * Get Contact List
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param phoneNumber Phone number in E.164 (11-digits) format with or without plus &#39;+&#39;. Multiple values are supported
 * @param startsWith If specified, only contacts whose First name or Last name start with the mentioned substring are returned. Case-insensitive
 * @param sortBy Sorts results by the specified property. The default is &#39;First Name&#39;
 * @param page Indicates the page number to retrieve. Only positive number values are allowed. Default value is &#39;1&#39;
 * @param perPage Indicates the page size (number of items). If not specified, the value is &#39;100&#39; by default
 * @return *InlineResponseDefault9
 */
func (a AddressBookApi) RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactGet(accountId string, extensionId string, phoneNumber string, startsWith string, sortBy string, page int32, perPage int32) (*InlineResponseDefault9, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != "" {
		localVarHeaderParams["Authorization"] = "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("phoneNumber", a.Configuration.APIClient.ParameterToString(phoneNumber, ""))
	localVarQueryParams.Add("startsWith", a.Configuration.APIClient.ParameterToString(startsWith, ""))
	localVarQueryParams.Add("sortBy", a.Configuration.APIClient.ParameterToString(sortBy, ""))
	localVarQueryParams.Add("page", a.Configuration.APIClient.ParameterToString(page, ""))
	localVarQueryParams.Add("perPage", a.Configuration.APIClient.ParameterToString(perPage, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponseDefault9)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 *
 * Create New Contact
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param body
 * @return *PersonalContactInfo
 */
func (a AddressBookApi) RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactPost(accountId string, extensionId string, body PersonalContactInfo) (*PersonalContactInfo, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/contact"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != "" {
		localVarHeaderParams["Authorization"] = "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	var successPayload = new(PersonalContactInfo)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAddressBookContactPost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 *
 * Get Contact Group List
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @return *InlineResponseDefault10
 */
func (a AddressBookApi) RestapiV10AccountAccountIdExtensionExtensionIdAddressBookGroupGet(accountId string, extensionId string) (*InlineResponseDefault10, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book/group"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != "" {
		localVarHeaderParams["Authorization"] = "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponseDefault10)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAddressBookGroupGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}

/**
 *
 * Contacts Synchronization
 *
 * @param accountId Internal identifier of a RingCentral account or tilde (~) to indicate the account logged-in within the current session
 * @param extensionId Internal identifier of an extension or tilde (~) to indicate the extension assigned to the account logged-in within the current session
 * @param syncType Type of synchronization. The default value is &#39;FSync&#39;
 * @param syncToken Value of syncToken property of the last sync request response
 * @param perPage Number of records per page to be returned. The max number of records is 250, which is also the default. For FSync — if the number of records exceeds the parameter value (either specified or default), all of the pages can be retrieved in several requests. For ISync — if the number of records exceeds the page size, the number of incoming changes to this number is limited
 * @param pageId Internal identifier of a page. It can be obtained from the &#39;nextPageId&#39; parameter passed in response body
 * @return *InlineResponseDefault8
 */
func (a AddressBookApi) RestapiV10AccountAccountIdExtensionExtensionIdAddressBookSyncGet(accountId string, extensionId string, syncType string, syncToken string, perPage int32, pageId int32) (*InlineResponseDefault8, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Get")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/account/{accountId}/extension/{extensionId}/address-book-sync"
	localVarPath = strings.Replace(localVarPath, "{"+"accountId"+"}", fmt.Sprintf("%v", accountId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"extensionId"+"}", fmt.Sprintf("%v", extensionId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := make(map[string]string)
	var localVarPostBody interface{}
	var localVarFileName string
	var localVarFileBytes []byte
	// authentication '(oauth)' required
	// oauth required
	if a.Configuration.AccessToken != "" {
		localVarHeaderParams["Authorization"] = "Bearer " + a.Configuration.AccessToken
	}
	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		localVarHeaderParams[key] = a.Configuration.DefaultHeader[key]
	}
	localVarQueryParams.Add("syncType", a.Configuration.APIClient.ParameterToString(syncType, ""))
	localVarQueryParams.Add("syncToken", a.Configuration.APIClient.ParameterToString(syncToken, ""))
	localVarQueryParams.Add("perPage", a.Configuration.APIClient.ParameterToString(perPage, ""))
	localVarQueryParams.Add("pageId", a.Configuration.APIClient.ParameterToString(pageId, ""))

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(InlineResponseDefault8)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10AccountAccountIdExtensionExtensionIdAddressBookSyncGet", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
	if localVarHttpResponse != nil {
		localVarAPIResponse.Response = localVarHttpResponse.RawResponse
		localVarAPIResponse.Payload = localVarHttpResponse.Body()
	}

	if err != nil {
		return successPayload, localVarAPIResponse, err
	}
	err = json.Unmarshal(localVarHttpResponse.Body(), &successPayload)
	return successPayload, localVarAPIResponse, err
}
