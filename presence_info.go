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

type PresenceInfo struct {

	// Canonical URI of a presence info resource
	Uri string `json:"uri,omitempty"`

	// If 'True' enables other extensions to see the extension presence status
	AllowSeeMyPresence bool `json:"allowSeeMyPresence,omitempty"`

	// Extended DnD (Do not Disturb) status. Cannot be set for Department/Announcement/Voicemail (Take Messages Only)/Fax User/Shared Lines Group/Paging Only Group/IVR Menu/Application Extension/Park Location extensions. The 'DoNotAcceptDepartmentCalls' and 'TakeDepartmentCallsOnly' values are applicable only for extensions - members of a Department; if these values are set for department outsiders, the 400 Bad Request error code is returned. The 'TakeDepartmentCallsOnly' status can be set through the old RingCentral user interface and is available for some migrated accounts only.
	DndStatus string `json:"dndStatus,omitempty"`

	// Information on extension, for which this presence data is returned
	Extension PresenceInfoExtensionInfo `json:"extension,omitempty"`

	// Custom status message (as previously published by user)
	Message string `json:"message,omitempty"`

	// If 'True' enables the extension user to pick up a monitored line on hold
	PickUpCallsOnHold bool `json:"pickUpCallsOnHold,omitempty"`

	// Aggregated presence status, calculated from a number of sources
	PresenceStatus string `json:"presenceStatus,omitempty"`

	// If 'True' enables to ring extension phone, if any user monitored by this extension is ringing
	RingOnMonitoredCall bool `json:"ringOnMonitoredCall,omitempty"`

	// Telephony presence status
	TelephonyStatus string `json:"telephonyStatus,omitempty"`

	// User-defined presence status (as previously published by the user)
	UserStatus string `json:"userStatus,omitempty"`
}
