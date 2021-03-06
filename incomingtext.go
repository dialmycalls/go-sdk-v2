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

type Incomingtext struct {

	// Unique identifier for this text.
	Id Uuid `json:"id,omitempty"`

	// The phone number the text message was sent from.
	FromNumber string `json:"from_number,omitempty"`

	// The phone number the text message was sent to.
	ToNumber string `json:"to_number,omitempty"`

	// The text message that was sent.  This data is URL encoded.
	Message string `json:"message,omitempty"`

	// When the text message was created.
	CreatedAt string `json:"created_at,omitempty"`

	// When the text message was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}
