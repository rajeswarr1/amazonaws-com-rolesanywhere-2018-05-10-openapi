package main

import (
	"github.com/iam-roles-anywhere/mcp-server/config"
	"github.com/iam-roles-anywhere/mcp-server/models"
	tools_profiles "github.com/iam-roles-anywhere/mcp-server/tools/profiles"
	tools_trustanchor "github.com/iam-roles-anywhere/mcp-server/tools/trustanchor"
	tools_crls "github.com/iam-roles-anywhere/mcp-server/tools/crls"
	tools_untagresource "github.com/iam-roles-anywhere/mcp-server/tools/untagresource"
	tools_subject "github.com/iam-roles-anywhere/mcp-server/tools/subject"
	tools_profile "github.com/iam-roles-anywhere/mcp-server/tools/profile"
	tools_put_notifications_settings "github.com/iam-roles-anywhere/mcp-server/tools/put_notifications_settings"
	tools_crl "github.com/iam-roles-anywhere/mcp-server/tools/crl"
	tools_reset_notifications_settings "github.com/iam-roles-anywhere/mcp-server/tools/reset_notifications_settings"
	tools_trustanchors "github.com/iam-roles-anywhere/mcp-server/tools/trustanchors"
	tools_subjects "github.com/iam-roles-anywhere/mcp-server/tools/subjects"
	tools_listtagsforresource_resourcearn "github.com/iam-roles-anywhere/mcp-server/tools/listtagsforresource_resourcearn"
	tools_tagresource "github.com/iam-roles-anywhere/mcp-server/tools/tagresource"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_profiles.CreateCreateprofileTool(cfg),
		tools_profiles.CreateListprofilesTool(cfg),
		tools_trustanchor.CreateGettrustanchorTool(cfg),
		tools_trustanchor.CreateUpdatetrustanchorTool(cfg),
		tools_trustanchor.CreateDeletetrustanchorTool(cfg),
		tools_crls.CreateListcrlsTool(cfg),
		tools_crls.CreateImportcrlTool(cfg),
		tools_untagresource.CreateUntagresourceTool(cfg),
		tools_subject.CreateGetsubjectTool(cfg),
		tools_profile.CreateDeleteprofileTool(cfg),
		tools_profile.CreateGetprofileTool(cfg),
		tools_profile.CreateUpdateprofileTool(cfg),
		tools_profile.CreateDisableprofileTool(cfg),
		tools_put_notifications_settings.CreatePutnotificationsettingsTool(cfg),
		tools_crl.CreateUpdatecrlTool(cfg),
		tools_crl.CreateDeletecrlTool(cfg),
		tools_crl.CreateGetcrlTool(cfg),
		tools_profile.CreateEnableprofileTool(cfg),
		tools_crl.CreateEnablecrlTool(cfg),
		tools_reset_notifications_settings.CreateResetnotificationsettingsTool(cfg),
		tools_trustanchors.CreateCreatetrustanchorTool(cfg),
		tools_trustanchors.CreateListtrustanchorsTool(cfg),
		tools_subjects.CreateListsubjectsTool(cfg),
		tools_trustanchor.CreateDisabletrustanchorTool(cfg),
		tools_trustanchor.CreateEnabletrustanchorTool(cfg),
		tools_crl.CreateDisablecrlTool(cfg),
		tools_listtagsforresource_resourcearn.CreateListtagsforresourceTool(cfg),
		tools_tagresource.CreateTagresourceTool(cfg),
	}
}
