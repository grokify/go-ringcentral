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
	"net/url"
	"strings"
)

type PhoneNumberPoolApi struct {
	Configuration *Configuration
}

func NewPhoneNumberPoolApi() *PhoneNumberPoolApi {
	configuration := NewConfiguration()
	return &PhoneNumberPoolApi{
		Configuration: configuration,
	}
}

func NewPhoneNumberPoolApiWithBasePath(basePath string) *PhoneNumberPoolApi {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &PhoneNumberPoolApi{
		Configuration: configuration,
	}
}

/**
 *
 * Look up Phone Number
 *
 * @param areaCode Area code of the location
 * @param countryCode Two-letter country code, complying with the ISO standard
 * @param countryId Internal identifier of a country; &#39;1&#39;- the US; &#39;39&#39; - Canada; &#39;224&#39; - the UK. The default value is &#39;1&#39;
 * @param exclude A string of digits (one and more) that should not appear among the last four digits (line part) of the phone numbers that will be returned. It is possible to specify several exclude parameters. If specified, it is taken into account in all returned phone numbers both in the phone numbers satisfying to parameters of lookup and in alternative phone numbers (in case when extendedSearch is specified)
 * @param extendedSearch If the value is &#39;False&#39;, then the returned numbers exactly correspond to the specified NXX, NPA and LINE or countryCode, areaCode and numberPattern parameters. If the value is &#39;True&#39;, then the resulting numbers are ranked and returned with the rank attribute values (1-10). The default value is &#39;False&#39;
 * @param line LINE pattern for vanity or wildcard search. Digits, Latin characters and asterisks are allowed (usually 4 characters)
 * @param numberPattern Phone number pattern (for wildcard or vanity search). For NANP countries (US, Canada) is concatenation of nxx (the first three digits) and line. If the first three characters are specified as not digits (e.g. 5** or CAT) then parameter extendedSearch will be ignored.
 * @param nxx NXX pattern for vanity or wildcard search. Digits, Latin characters and asterisks are allowed (usually 3 characters)
 * @param npa Area code (mandatory). For example, 800, 844, 855, 866, 877, 888 for North America; and 647 for Canada
 * @param paymentType Payment type. Default is &#39;Local&#39; (it should correlate with the npa provided)
 * @param perPage Indicates the page size (number of items). If not specified, the value is &#39;10&#39; by default
 * @param smsEnabled Specifies if SMS activation is available for the number. If specified, it is taken into account in all returned phone numbers both in the phone numbers satisfying to parameters of lookup and in alternative phone numbers (in case when extendedSearch is specified). If not specified, the value is null.
 * @return *InlineResponseDefault35
 */
func (a PhoneNumberPoolApi) RestapiV10NumberPoolLookupPost(areaCode int32, countryCode string, countryId string, exclude string, extendedSearch bool, line string, numberPattern string, nxx string, npa string, paymentType string, perPage int32, smsEnabled bool) (*InlineResponseDefault35, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/number-pool/lookup"

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
	localVarQueryParams.Add("areaCode", a.Configuration.APIClient.ParameterToString(areaCode, ""))
	localVarQueryParams.Add("countryCode", a.Configuration.APIClient.ParameterToString(countryCode, ""))
	localVarQueryParams.Add("countryId", a.Configuration.APIClient.ParameterToString(countryId, ""))
	localVarQueryParams.Add("exclude", a.Configuration.APIClient.ParameterToString(exclude, ""))
	localVarQueryParams.Add("extendedSearch", a.Configuration.APIClient.ParameterToString(extendedSearch, ""))
	localVarQueryParams.Add("line", a.Configuration.APIClient.ParameterToString(line, ""))
	localVarQueryParams.Add("numberPattern", a.Configuration.APIClient.ParameterToString(numberPattern, ""))
	localVarQueryParams.Add("nxx", a.Configuration.APIClient.ParameterToString(nxx, ""))
	localVarQueryParams.Add("npa", a.Configuration.APIClient.ParameterToString(npa, ""))
	localVarQueryParams.Add("paymentType", a.Configuration.APIClient.ParameterToString(paymentType, ""))
	localVarQueryParams.Add("perPage", a.Configuration.APIClient.ParameterToString(perPage, ""))
	localVarQueryParams.Add("smsEnabled", a.Configuration.APIClient.ParameterToString(smsEnabled, ""))

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
	var successPayload = new(InlineResponseDefault35)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10NumberPoolLookupPost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
 * Reserve Phone Number
 *
 * @param body
 * @return *InlineResponseDefault36
 */
func (a PhoneNumberPoolApi) RestapiV10NumberPoolReservePost(body Body19) (*InlineResponseDefault36, *APIResponse, error) {

	var localVarHttpMethod = strings.ToUpper("Post")
	// create path and map variables
	localVarPath := a.Configuration.BasePath + "/restapi/v1.0/number-pool/reserve"

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
	var successPayload = new(InlineResponseDefault36)
	localVarHttpResponse, err := a.Configuration.APIClient.CallAPI(localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)

	var localVarURL, _ = url.Parse(localVarPath)
	localVarURL.RawQuery = localVarQueryParams.Encode()
	var localVarAPIResponse = &APIResponse{Operation: "RestapiV10NumberPoolReservePost", Method: localVarHttpMethod, RequestURL: localVarURL.String()}
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
