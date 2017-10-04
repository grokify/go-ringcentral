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

type ServerInfo struct {

	// Canonical URI of the API version
	Uri string `json:"uri,omitempty"`

	// Full API version information: uri, number, release date
	ApiVersions []VersionInfo `json:"apiVersions,omitempty"`

	// Server version
	ServerVersion string `json:"serverVersion,omitempty"`

	// Server revision
	ServerRevision string `json:"serverRevision,omitempty"`
}
