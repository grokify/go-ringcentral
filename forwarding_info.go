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

type ForwardingInfo struct {

	// Specifies if the user's softphone(s) are notified before forwarding the incoming call to desk phones and forwarding numbers
	NotifyMySoftPhones bool `json:"notifyMySoftPhones,omitempty"`

	// Specifies if the administrator's softphone is notified before forwarding the incoming call to desk phones and forwarding numbers. The default value is 'False'
	NotifyAdminSoftPhones bool `json:"notifyAdminSoftPhones,omitempty"`

	// Number of rings before forwarding starts
	SoftPhonesRingCount int32 `json:"softPhonesRingCount,omitempty"`

	// Specifies the order in which forwarding numbers ring. 'Sequentially' means that forwarding numbers are ringing one at a time, in order of priority. 'Simultaneously' means that forwarding numbers are ring all at the same time
	RingingMode string `json:"ringingMode,omitempty"`

	// Information on a call forwarding rule
	Rules []RuleInfo `json:"rules,omitempty"`
}
