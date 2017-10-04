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

type ExtensionListEvent struct {

	// Internal identifier of an extension
	ExtensionId string `json:"extensionId,omitempty"`

	// Type of extension info change
	EventType string `json:"eventType,omitempty"`
}
