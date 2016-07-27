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

type ContactAttributes struct {

	// Contact's phone number.
	Phone string `json:"phone,omitempty"`

	// Contact's first name.
	Firstname string `json:"firstname,omitempty"`

	// Contact's last name.
	Lastname string `json:"lastname,omitempty"`

	// Contact's email address.
	Email string `json:"email,omitempty"`
}
