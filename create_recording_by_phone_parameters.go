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

type CreateRecordingByPhoneParameters struct {

	// (Required)  The name of the recording.
	Name string `json:"name,omitempty"`

	// The caller id that the create recording message should be sent from.
	CalleridId Uuid `json:"callerid_id,omitempty"`

	// Add or remove the DialMyCalls intro message.
	Whitelabel bool `json:"whitelabel,omitempty"`

	// (Required)  The recipient's phone number who will record the message.
	Phone string `json:"phone,omitempty"`

	// The recipient's phone extension up to 10 numeric digits.
	Extension string `json:"extension,omitempty"`
}
