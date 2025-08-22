package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ListSubjectsResponse represents the ListSubjectsResponse schema from the OpenAPI specification
type ListSubjectsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Subjects interface{} `json:"subjects,omitempty"`
}

// SubjectDetail represents the SubjectDetail schema from the OpenAPI specification
type SubjectDetail struct {
	Credentials interface{} `json:"credentials,omitempty"`
	Subjectid interface{} `json:"subjectId,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
	Instanceproperties interface{} `json:"instanceProperties,omitempty"`
	Subjectarn interface{} `json:"subjectArn,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Lastseenat interface{} `json:"lastSeenAt,omitempty"`
	X509subject interface{} `json:"x509Subject,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
}

// ResetNotificationSettingsRequest represents the ResetNotificationSettingsRequest schema from the OpenAPI specification
type ResetNotificationSettingsRequest struct {
	Notificationsettingkeys interface{} `json:"notificationSettingKeys"`
	Trustanchorid interface{} `json:"trustAnchorId"`
}

// InstanceProperty represents the InstanceProperty schema from the OpenAPI specification
type InstanceProperty struct {
	Failed interface{} `json:"failed,omitempty"`
	Properties interface{} `json:"properties,omitempty"` // A list of instanceProperty objects.
	Seenat interface{} `json:"seenAt,omitempty"`
}

// Source represents the Source schema from the OpenAPI specification
type Source struct {
	Sourcedata interface{} `json:"sourceData,omitempty"`
	Sourcetype interface{} `json:"sourceType,omitempty"`
}

// UpdateCrlRequest represents the UpdateCrlRequest schema from the OpenAPI specification
type UpdateCrlRequest struct {
	Crldata interface{} `json:"crlData,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// TrustAnchorDetail represents the TrustAnchorDetail schema from the OpenAPI specification
type TrustAnchorDetail struct {
	Trustanchorarn interface{} `json:"trustAnchorArn,omitempty"`
	Trustanchorid interface{} `json:"trustAnchorId,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Notificationsettings interface{} `json:"notificationSettings,omitempty"`
	Source interface{} `json:"source,omitempty"`
}

// ListProfilesResponse represents the ListProfilesResponse schema from the OpenAPI specification
type ListProfilesResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Profiles interface{} `json:"profiles,omitempty"`
}

// PutNotificationSettingsResponse represents the PutNotificationSettingsResponse schema from the OpenAPI specification
type PutNotificationSettingsResponse struct {
	Trustanchor TrustAnchorDetail `json:"trustAnchor"` // The state of the trust anchor after a read or write operation.
}

// NotificationSettingDetail represents the NotificationSettingDetail schema from the OpenAPI specification
type NotificationSettingDetail struct {
	Configuredby interface{} `json:"configuredBy,omitempty"`
	Enabled interface{} `json:"enabled"`
	Event interface{} `json:"event"`
	Threshold interface{} `json:"threshold,omitempty"`
	Channel interface{} `json:"channel,omitempty"`
}

// ProfileDetail represents the ProfileDetail schema from the OpenAPI specification
type ProfileDetail struct {
	Createdby interface{} `json:"createdBy,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Requireinstanceproperties interface{} `json:"requireInstanceProperties,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Managedpolicyarns interface{} `json:"managedPolicyArns,omitempty"`
	Profilearn interface{} `json:"profileArn,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
	Durationseconds interface{} `json:"durationSeconds,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Profileid interface{} `json:"profileId,omitempty"`
	Rolearns interface{} `json:"roleArns,omitempty"`
	Sessionpolicy interface{} `json:"sessionPolicy,omitempty"`
}

// CreateTrustAnchorRequest represents the CreateTrustAnchorRequest schema from the OpenAPI specification
type CreateTrustAnchorRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Name interface{} `json:"name"`
	Notificationsettings interface{} `json:"notificationSettings,omitempty"`
	Source interface{} `json:"source"`
}

// ListCrlsResponse represents the ListCrlsResponse schema from the OpenAPI specification
type ListCrlsResponse struct {
	Crls interface{} `json:"crls,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// NotificationSettingKey represents the NotificationSettingKey schema from the OpenAPI specification
type NotificationSettingKey struct {
	Channel interface{} `json:"channel,omitempty"`
	Event interface{} `json:"event"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// UpdateTrustAnchorRequest represents the UpdateTrustAnchorRequest schema from the OpenAPI specification
type UpdateTrustAnchorRequest struct {
	Name interface{} `json:"name,omitempty"`
	Source interface{} `json:"source,omitempty"`
}

// ListTrustAnchorsResponse represents the ListTrustAnchorsResponse schema from the OpenAPI specification
type ListTrustAnchorsResponse struct {
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Trustanchors interface{} `json:"trustAnchors,omitempty"`
}

// ResetNotificationSettingsResponse represents the ResetNotificationSettingsResponse schema from the OpenAPI specification
type ResetNotificationSettingsResponse struct {
	Trustanchor TrustAnchorDetail `json:"trustAnchor"` // The state of the trust anchor after a read or write operation.
}

// CredentialSummary represents the CredentialSummary schema from the OpenAPI specification
type CredentialSummary struct {
	Failed interface{} `json:"failed,omitempty"`
	Issuer interface{} `json:"issuer,omitempty"`
	Seenat interface{} `json:"seenAt,omitempty"`
	Serialnumber interface{} `json:"serialNumber,omitempty"`
	X509certificatedata interface{} `json:"x509CertificateData,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
}

// InstancePropertyMap represents the InstancePropertyMap schema from the OpenAPI specification
type InstancePropertyMap struct {
}

// CrlDetail represents the CrlDetail schema from the OpenAPI specification
type CrlDetail struct {
	Crlid interface{} `json:"crlId,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Trustanchorarn interface{} `json:"trustAnchorArn,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
	Createdat interface{} `json:"createdAt,omitempty"`
	Crlarn interface{} `json:"crlArn,omitempty"`
	Crldata interface{} `json:"crlData,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Tags interface{} `json:"tags,omitempty"`
}

// SourceData represents the SourceData schema from the OpenAPI specification
type SourceData struct {
	Acmpcaarn interface{} `json:"acmPcaArn,omitempty"`
	X509certificatedata interface{} `json:"x509CertificateData,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"key"`
	Value interface{} `json:"value"`
}

// ScalarSubjectRequest represents the ScalarSubjectRequest schema from the OpenAPI specification
type ScalarSubjectRequest struct {
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"resourceArn"`
	Tags interface{} `json:"tags"`
}

// PutNotificationSettingsRequest represents the PutNotificationSettingsRequest schema from the OpenAPI specification
type PutNotificationSettingsRequest struct {
	Notificationsettings interface{} `json:"notificationSettings"`
	Trustanchorid interface{} `json:"trustAnchorId"`
}

// ImportCrlRequest represents the ImportCrlRequest schema from the OpenAPI specification
type ImportCrlRequest struct {
	Name interface{} `json:"name"`
	Tags interface{} `json:"tags,omitempty"`
	Trustanchorarn interface{} `json:"trustAnchorArn"`
	Crldata interface{} `json:"crlData"`
	Enabled interface{} `json:"enabled,omitempty"`
}

// TrustAnchorDetailResponse represents the TrustAnchorDetailResponse schema from the OpenAPI specification
type TrustAnchorDetailResponse struct {
	Trustanchor interface{} `json:"trustAnchor"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"resourceArn"`
	Tagkeys interface{} `json:"tagKeys"`
}

// ListRequest represents the ListRequest schema from the OpenAPI specification
type ListRequest struct {
}

// ScalarCrlRequest represents the ScalarCrlRequest schema from the OpenAPI specification
type ScalarCrlRequest struct {
}

// ScalarProfileRequest represents the ScalarProfileRequest schema from the OpenAPI specification
type ScalarProfileRequest struct {
}

// ProfileDetailResponse represents the ProfileDetailResponse schema from the OpenAPI specification
type ProfileDetailResponse struct {
	Profile interface{} `json:"profile,omitempty"`
}

// CrlDetailResponse represents the CrlDetailResponse schema from the OpenAPI specification
type CrlDetailResponse struct {
	Crl interface{} `json:"crl"`
}

// NotificationSetting represents the NotificationSetting schema from the OpenAPI specification
type NotificationSetting struct {
	Enabled interface{} `json:"enabled"`
	Event interface{} `json:"event"`
	Threshold interface{} `json:"threshold,omitempty"`
	Channel interface{} `json:"channel,omitempty"`
}

// SubjectDetailResponse represents the SubjectDetailResponse schema from the OpenAPI specification
type SubjectDetailResponse struct {
	Subject interface{} `json:"subject,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// SubjectSummary represents the SubjectSummary schema from the OpenAPI specification
type SubjectSummary struct {
	Createdat interface{} `json:"createdAt,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Lastseenat interface{} `json:"lastSeenAt,omitempty"`
	Subjectarn interface{} `json:"subjectArn,omitempty"`
	Subjectid interface{} `json:"subjectId,omitempty"`
	Updatedat interface{} `json:"updatedAt,omitempty"`
	X509subject interface{} `json:"x509Subject,omitempty"`
}

// CreateProfileRequest represents the CreateProfileRequest schema from the OpenAPI specification
type CreateProfileRequest struct {
	Tags interface{} `json:"tags,omitempty"`
	Durationseconds interface{} `json:"durationSeconds,omitempty"`
	Enabled interface{} `json:"enabled,omitempty"`
	Managedpolicyarns interface{} `json:"managedPolicyArns,omitempty"`
	Name interface{} `json:"name"`
	Requireinstanceproperties interface{} `json:"requireInstanceProperties,omitempty"`
	Rolearns interface{} `json:"roleArns"`
	Sessionpolicy interface{} `json:"sessionPolicy,omitempty"`
}

// UpdateProfileRequest represents the UpdateProfileRequest schema from the OpenAPI specification
type UpdateProfileRequest struct {
	Durationseconds interface{} `json:"durationSeconds,omitempty"`
	Managedpolicyarns interface{} `json:"managedPolicyArns,omitempty"`
	Name interface{} `json:"name,omitempty"`
	Rolearns interface{} `json:"roleArns,omitempty"`
	Sessionpolicy interface{} `json:"sessionPolicy,omitempty"`
}

// ScalarTrustAnchorRequest represents the ScalarTrustAnchorRequest schema from the OpenAPI specification
type ScalarTrustAnchorRequest struct {
}
