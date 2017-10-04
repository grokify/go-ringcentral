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
	"time"
)

type Body8 struct {

	// Recipient information. Phone number property is mandatory. Optional for resend fax request
	To []CallerInfo `json:"to,omitempty"`

	// Fax resolution
	FaxResolution string `json:"faxResolution,omitempty"`

	// The datetime to send fax at, in ISO 8601 format including timezone, for example 2016-03-10T18:07:52.534Z. If time is not specified, the fax will be send immediately
	SendTime time.Time `json:"sendTime,omitempty"`

	// Optional. Cover page index. If not specified, the default cover page which is configured in \"Outbound Fax Settings\" is attached. See also 'Available Cover Pages' table below
	CoverIndex int32 `json:"coverIndex,omitempty"`

	// Optional. Cover page text, entered by the fax sender and printed on the cover page. Maximum length is limited to 1024 symbols
	CoverPageText string `json:"coverPageText,omitempty"`

	// Internal identifier of the original fax message which needs to be resent. Mandatory for resend fax request
	OriginalMessageId string `json:"originalMessageId,omitempty"`
}
