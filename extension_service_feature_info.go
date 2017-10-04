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

type ExtensionServiceFeatureInfo struct {

	// Feature status; shows feature availability for an extension
	Enabled bool `json:"enabled,omitempty"`

	// Feature name, see all available values in Service Feature List
	FeatureName string `json:"featureName,omitempty"`

	// Reason of limitation for a particular service feature. Returned only if the enabled parameter value is 'False', see Service Feature Limitations and Reasons. When retrieving service features for an extension, the reasons for the limitations, if any, are returned in response
	Reason string `json:"reason,omitempty"`
}
