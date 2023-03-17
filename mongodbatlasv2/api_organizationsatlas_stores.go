 package latest

import (
	"context"
	openapiclient "github.com/mongodb/go-client-mongodb-atlas"
)

type OrganizationsApiStore struct {
}

func OrganizationsApiServiceCreateOrganization(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.CreateOrganization(context.Background()).Execute()
}
func OrganizationsApiServiceCreateOrganizationInvitation(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.CreateOrganizationInvitation(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceDeleteOrganization(apiClient *openapiclient.Client) {
	httpRes, err := apiClient.OrganizationsApi.DeleteOrganization(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceDeleteOrganizationInvitation(apiClient *openapiclient.Client) {
	httpRes, err := apiClient.OrganizationsApi.DeleteOrganizationInvitation(context.Background(), orgId, invitationId).Execute()
}
func OrganizationsApiServiceGetOrganization(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.GetOrganization(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceGetOrganizationInvitation(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.GetOrganizationInvitation(context.Background(), orgId, invitationId).Execute()
}
func OrganizationsApiServiceGetOrganizationSettings(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.GetOrganizationSettings(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceListOrganizationInvitations(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.ListOrganizationInvitations(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceListOrganizationProjects(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.ListOrganizationProjects(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceListOrganizationUsers(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.ListOrganizationUsers(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceListOrganizations(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.ListOrganizations(context.Background()).Execute()
}
func OrganizationsApiServiceRenameOrganization(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.RenameOrganization(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceUpdateOrganizationInvitation(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.UpdateOrganizationInvitation(context.Background(), orgId).Execute()
}
func OrganizationsApiServiceUpdateOrganizationInvitationById(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.UpdateOrganizationInvitationById(context.Background(), orgId, invitationId).Execute()
}
func OrganizationsApiServiceUpdateOrganizationSettings(apiClient *openapiclient.Client) {
	resp, httpRes, err := apiClient.OrganizationsApi.UpdateOrganizationSettings(context.Background(), orgId).Execute()
}
}
