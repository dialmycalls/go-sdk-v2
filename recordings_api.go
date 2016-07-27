/* 
 * DialMyCalls API
 *
 * The DialMyCalls API
 *
 * OpenAPI spec version: 2.0.1
 * Contact: support@dialmycalls.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package dialmycalls

import (
	"strings"
	"fmt"
	"errors"
	"net/url"
	"encoding/json"
)

type Recordings struct {
	Configuration Configuration
}

func NewRecordings() *Recordings {
	configuration := NewConfiguration()
	return &Recordings{
		Configuration: *configuration,
	}
}

func NewRecordingsWithBasePath(basePath string) *Recordings {
	configuration := NewConfiguration()
	configuration.BasePath = basePath

	return &Recordings{
		Configuration: *configuration,
	}
}

/**
 * Create Recording (Text-to-Speech)
 * Create a new recording using text-to-speech. &lt;br&gt;&lt;br&gt; Returns a recording object on success, and returns an error otherwise. &lt;br&gt;&lt;br&gt; &#x60;&#x60;&#x60; curl -i -H \&quot;Content-Type: application/json\&quot; -X POST -d \&quot;{\\\&quot;name\\\&quot;: \\\&quot;Test Recording\\\&quot;, \\\&quot;gender\\\&quot;: \\\&quot;M\\\&quot;, \\\&quot;language\\\&quot;: \\\&quot;en\\\&quot;, \\\&quot;text\\\&quot;: \\\&quot;This is just a test.\\\&quot;}\&quot; https://$API_KEY@api.dialmycalls.com/2.0/recording/tts &#x60;&#x60;&#x60;
 *
 * @param createRecordingParameters Request body
 * @return *Object
 */
func (a Recordings) CreateRecording(createRecordingParameters CreateRecordingParameters) (*Object, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/recording/tts"

	// verify the required parameter 'createRecordingParameters' is set
	if &createRecordingParameters == nil {
		return new(Object), nil, errors.New("Missing required parameter 'createRecordingParameters' when calling Recordings->CreateRecording")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (api_key) required

	// set key with prefix in header
	headerParams["X-Auth-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-ApiKey")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
"application/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &createRecordingParameters

	var successPayload = new(Object)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Create Recording (Phone)
 * Create a new recording by phone. &lt;br&gt;&lt;br&gt; Returns a recording object on success, and returns an error otherwise. &lt;br&gt;&lt;br&gt; &#x60;&#x60;&#x60; curl -i -H \&quot;Content-Type: application/json\&quot; -X POST -d \&quot;{\\\&quot;name\\\&quot;: \\\&quot;Test Recording\\\&quot;, \\\&quot;phone\\\&quot;: \\\&quot;5551234567\\\&quot;, \\\&quot;callerid_id\\\&quot;: \\\&quot;$CALLERID_ID\\\&quot;}\&quot; https://$API_KEY@api.dialmycalls.com/2.0/recording/phone &#x60;&#x60;&#x60;
 *
 * @param createRecordingByPhoneParameters Request body
 * @return *Object
 */
func (a Recordings) CreateRecordingByPhone(createRecordingByPhoneParameters CreateRecordingByPhoneParameters) (*Object, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/recording/phone"

	// verify the required parameter 'createRecordingByPhoneParameters' is set
	if &createRecordingByPhoneParameters == nil {
		return new(Object), nil, errors.New("Missing required parameter 'createRecordingByPhoneParameters' when calling Recordings->CreateRecordingByPhone")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (api_key) required

	// set key with prefix in header
	headerParams["X-Auth-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-ApiKey")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
"application/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &createRecordingByPhoneParameters

	var successPayload = new(Object)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Create Recording (URL)
 * Create a new recording from a URL. &lt;br&gt;&lt;br&gt; Returns a recording object on success, and returns an error otherwise. &lt;br&gt;&lt;br&gt; &#x60;&#x60;&#x60; curl -i -H \&quot;Content-Type: application/json\&quot; -X POST -d \&quot;{\\\&quot;name\\\&quot;: \\\&quot;Test Recording\\\&quot;, \\\&quot;url\\\&quot;: \\\&quot;https://ia700200.us.archive.org/1/items/testmp3testfile/mpthreetest.mp3\\\&quot;}\&quot; https://$API_KEY@api.dialmycalls.com/2.0/recording/url &#x60;&#x60;&#x60;
 *
 * @param createRecordingByUrlParameters Request body
 * @return *Object
 */
func (a Recordings) CreateRecordingByUrl(createRecordingByUrlParameters CreateRecordingByUrlParameters) (*Object, *APIResponse, error) {

	var httpMethod = "Post"
	// create path and map variables
	path := a.Configuration.BasePath + "/recording/url"

	// verify the required parameter 'createRecordingByUrlParameters' is set
	if &createRecordingByUrlParameters == nil {
		return new(Object), nil, errors.New("Missing required parameter 'createRecordingByUrlParameters' when calling Recordings->CreateRecordingByUrl")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (api_key) required

	// set key with prefix in header
	headerParams["X-Auth-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-ApiKey")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
"application/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &createRecordingByUrlParameters

	var successPayload = new(Object)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Delete Recording
 * Delete a recording. &lt;br&gt;&lt;br&gt; Returns the following if a valid identifier was provided, and returns an error otherwise. &lt;br&gt;&lt;br&gt; &#x60;&#x60;&#x60; curl -i -H \&quot;Content-Type: application/json\&quot; -X DELETE https://$API_KEY@api.dialmycalls.com/2.0/recording/$RECORDING_ID &#x60;&#x60;&#x60;
 *
 * @param recordingId RecordingId
 * @return *Object
 */
func (a Recordings) DeleteRecordingById(recordingId string) (*Object, *APIResponse, error) {

	var httpMethod = "Delete"
	// create path and map variables
	path := a.Configuration.BasePath + "/recording/{RecordingId}"
	path = strings.Replace(path, "{"+"RecordingId"+"}", fmt.Sprintf("%v", recordingId), -1)

	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		return new(Object), nil, errors.New("Missing required parameter 'recordingId' when calling Recordings->DeleteRecordingById")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (api_key) required

	// set key with prefix in header
	headerParams["X-Auth-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-ApiKey")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
"application/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Object)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Get Recording
 * Retrieve a recording. &lt;br&gt;&lt;br&gt; Returns a recording object if a valid identifier was provided, and returns an error otherwise. &lt;br&gt;&lt;br&gt; &#x60;&#x60;&#x60; curl -i -H \&quot;Content-Type: application/json\&quot; -X GET https://$API_KEY@api.dialmycalls.com/2.0/recording/$RECORDING_ID &#x60;&#x60;&#x60;
 *
 * @param recordingId RecordingId
 * @return *Object
 */
func (a Recordings) GetRecordingById(recordingId string) (*Object, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/recording/{RecordingId}"
	path = strings.Replace(path, "{"+"RecordingId"+"}", fmt.Sprintf("%v", recordingId), -1)

	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		return new(Object), nil, errors.New("Missing required parameter 'recordingId' when calling Recordings->GetRecordingById")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (api_key) required

	// set key with prefix in header
	headerParams["X-Auth-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-ApiKey")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
"application/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	var successPayload = new(Object)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * List Recordings
 * Retrieve a list of recordings. &lt;br&gt;&lt;br&gt; Returns a list of recording objects. &lt;br&gt;&lt;br&gt; &#x60;&#x60;&#x60; curl -i -H \&quot;Content-Type: application/json\&quot; -X GET https://$API_KEY@api.dialmycalls.com/2.0/recordings &#x60;&#x60;&#x60;
 *
 * @param range_ Range (ie \&quot;records&#x3D;201-300\&quot;) of recordings requested
 * @return *Object
 */
func (a Recordings) GetRecordings(range_ string) (*Object, *APIResponse, error) {

	var httpMethod = "Get"
	// create path and map variables
	path := a.Configuration.BasePath + "/recordings"


	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (api_key) required

	// set key with prefix in header
	headerParams["X-Auth-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-ApiKey")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
"application/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}

	// header params "Range"
	headerParams["Range"] = range_

	var successPayload = new(Object)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

/**
 * Update Recording
 * Update an existing recording. &lt;br&gt;&lt;br&gt; Returns a recording object if a valid identifier was provided and input validation passed, and returns an error otherwise. &lt;br&gt;&lt;br&gt; &#x60;&#x60;&#x60; curl -i -H \&quot;Content-Type: application/json\&quot; -X PUT -d \&quot;{\\\&quot;name\\\&quot;: \\\&quot;Test Recording Updated\\\&quot;}\&quot; https://$API_KEY@api.dialmycalls.com/2.0/recording/$RECORDING_ID &#x60;&#x60;&#x60;
 *
 * @param updateRecordingByIdParameters Request body
 * @param recordingId RecordingId
 * @return *Object
 */
func (a Recordings) UpdateRecordingById(updateRecordingByIdParameters UpdateRecordingByIdParameters, recordingId string) (*Object, *APIResponse, error) {

	var httpMethod = "Put"
	// create path and map variables
	path := a.Configuration.BasePath + "/recording/{RecordingId}"
	path = strings.Replace(path, "{"+"RecordingId"+"}", fmt.Sprintf("%v", recordingId), -1)

	// verify the required parameter 'updateRecordingByIdParameters' is set
	if &updateRecordingByIdParameters == nil {
		return new(Object), nil, errors.New("Missing required parameter 'updateRecordingByIdParameters' when calling Recordings->UpdateRecordingById")
	}
	// verify the required parameter 'recordingId' is set
	if &recordingId == nil {
		return new(Object), nil, errors.New("Missing required parameter 'recordingId' when calling Recordings->UpdateRecordingById")
	}

	headerParams := make(map[string]string)
	queryParams := url.Values{}
	formParams := make(map[string]string)
	var postBody interface{}
	var fileName string
	var fileBytes []byte
	// authentication (api_key) required

	// set key with prefix in header
	headerParams["X-Auth-ApiKey"] = a.Configuration.GetAPIKeyWithPrefix("X-Auth-ApiKey")

	// add default headers if any
	for key := range a.Configuration.DefaultHeader {
		headerParams[key] = a.Configuration.DefaultHeader[key]
	}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json", "application/xml",  }

	// set Content-Type header
	localVarHttpContentType := a.Configuration.APIClient.SelectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		headerParams["Content-Type"] = localVarHttpContentType
	}
	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
"application/xml",
	}

	// set Accept header
	localVarHttpHeaderAccept := a.Configuration.APIClient.SelectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		headerParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	postBody = &updateRecordingByIdParameters

	var successPayload = new(Object)
	httpResponse, err := a.Configuration.APIClient.CallAPI(path, httpMethod, postBody, headerParams, queryParams, formParams, fileName, fileBytes)
	if err != nil {
		return successPayload, NewAPIResponse(httpResponse.RawResponse), err
	}
	err = json.Unmarshal(httpResponse.Body(), &successPayload)
	return successPayload, NewAPIResponse(httpResponse.RawResponse), err
}

