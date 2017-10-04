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

type AttachmentInfo struct {

	// Link to custom data attachment
	Uri string `json:"uri,omitempty"`

	// Type of custom data attachment, see also MIME Types
	ContentType string `json:"contentType,omitempty"`
}
