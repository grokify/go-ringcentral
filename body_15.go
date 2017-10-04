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

type Body15 struct {

	// Sender of an SMS message. The phoneNumber property must be filled to correspond to one of the account phone numbers which is allowed to send SMS
	From CallerInfo `json:"from,omitempty"`

	// Receiver of an SMS message. The phoneNumber property must be filled
	To []CallerInfo `json:"to,omitempty"`

	// Text of a message. Max length is 1000 symbols (2-byte UTF-16 encoded). If a character is encoded in 4 bytes in UTF-16 it is treated as 2 characters, thus restricting the maximum message length to 500 symbols
	Text string `json:"text,omitempty"`
}
