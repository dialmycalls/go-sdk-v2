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

type PushToTalk struct {

	// The add-on type. Option: ptt
	Type_ string `json:"type,omitempty"`

	// Phone number to connect the caller to.
	CalleridId Uuid `json:"callerid_id,omitempty"`

	// The ringtone to play to the user when connecting. Options: elevator_music, caribbean_music, classical_music, digital_ringing, old_phone_ringing, goofy_music
	Ringtone string `json:"ringtone,omitempty"`

	// Calls per minute throttling.
	Cpm int32 `json:"cpm,omitempty"`

	// Add a generic add-on message.
	AddMessage Object `json:"add_message,omitempty"`
}
