// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package keycloak

import (
	"strings"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/mrparkers/terraform-provider-keycloak/provider"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "keycloak"
	// modules:
	mainMod           = "index"          // the y module
	ldapMod           = "Ldap"           // the ldap module
	oidcMod           = "Oidc"           // the oidc module
	openIDMod         = "OpenId"         // the openid module
	samlMod           = "Saml"           // the Saml module
	authenticationMod = "Authentication" // the Authentication module
)

var namespaceMap = map[string]string{
	"keycloak": "Keycloak",
}

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	moduleName := strings.ToLower(mod)
	namespaceMap[moduleName] = mod
	fn := string(unicode.ToLower(rune(mem[0]))) + mem[1:]
	token := moduleName + "/" + fn
	return tokens.ModuleMember(mainPkg + ":" + token + ":" + mem)
}

func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

func makeDataSource(mod string, res string) tokens.ModuleMember {
	return makeMember(mod, res)
}

func makeResource(mod string, res string) tokens.Type {
	return makeType(mod, res)
}

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := provider.KeycloakProvider()

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "keycloak",
		GitHubOrg:         "mrparkers",
		Description:       "A Pulumi package for creating and managing keycloak cloud resources.",
		Keywords:          []string{"pulumi", "keycloak"},
		License:           "Apache-2.0",
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/pulumi/pulumi-keycloak",
		TFProviderLicense: refProviderLicense(tfbridge.MITLicenseType),
		Config: map[string]*tfbridge.SchemaInfo{
			"client_id": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KEYCLOAK_CLIENT_ID"},
				},
			},
			"client_secret": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KEYCLOAK_CLIENT_SECRET"},
				},
			},
			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KEYCLOAK_USER"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KEYCLOAK_PASSWORD"},
				},
			},
			"realm": {
				Default: &tfbridge.DefaultInfo{
					Value:   "master",
					EnvVars: []string{"KEYCLOAK_REALM"},
				},
			},
			"url": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KEYCLOAK_URL"},
				},
			},
			"client_timeout": {
				Default: &tfbridge.DefaultInfo{
					Value:   5,
					EnvVars: []string{"KEYCLOAK_CLIENT_TIMEOUT"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			"keycloak_attribute_importer_identity_provider_mapper": {
				Tok: makeResource(mainMod, "AttributeImporterIdentityProviderMapper"),
			},
			"keycloak_attribute_to_role_identity_provider_mapper": {
				Tok: makeResource(mainMod, "AttributeToRoleIdentityMapper"),
			},
			"keycloak_custom_user_federation":         {Tok: makeResource(mainMod, "CustomUserFederation")},
			"keycloak_default_groups":                 {Tok: makeResource(mainMod, "DefaultGroups")},
			"keycloak_generic_client_protocol_mapper": {Tok: makeResource(mainMod, "GenericClientProtocolMapper")},
			"keycloak_group":                          {Tok: makeResource(mainMod, "Group")},
			"keycloak_group_memberships":              {Tok: makeResource(mainMod, "GroupMemberships")},
			"keycloak_group_roles":                    {Tok: makeResource(mainMod, "GroupRoles")},
			"keycloak_hardcoded_attribute_identity_provider_mapper": {
				Tok: makeResource(mainMod, "HardcodedAttributeIdentityProviderMapper"),
			},
			"keycloak_hardcoded_role_identity_provider_mapper": {
				Tok: makeResource(mainMod, "HardcodedRoleIdentityMapper"),
			},
			"keycloak_realm": {
				Tok: makeResource(mainMod, "Realm"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"realm": {
						CSharpName: "RealmName",
					},
				},
			},
			"keycloak_required_action": {Tok: makeResource(mainMod, "RequiredAction")},
			"keycloak_role":            {Tok: makeResource(mainMod, "Role")},
			"keycloak_user":            {Tok: makeResource(mainMod, "User")},
			"keycloak_user_template_importer_identity_provider_mapper": {
				Tok: makeResource(mainMod, "UserTemplateImporterIdentityProviderMapper"),
			},
			"keycloak_realm_events":               {Tok: makeResource(mainMod, "RealmEvents")},
			"keycloak_generic_client_role_mapper": {Tok: makeResource(mainMod, "GenericClientRoleMapper")},

			"keycloak_ldap_full_name_mapper": {Tok: makeResource(ldapMod, "FullNameMapper")},
			"keycloak_ldap_group_mapper":     {Tok: makeResource(ldapMod, "GroupMapper")},
			"keycloak_ldap_msad_user_account_control_mapper": {
				Tok: makeResource(ldapMod, "MsadUserAccountControlMapper"),
			},
			"keycloak_ldap_user_attribute_mapper": {Tok: makeResource(ldapMod, "UserAttributeMapper")},
			"keycloak_ldap_user_federation":       {Tok: makeResource(ldapMod, "UserFederation")},
			"keycloak_ldap_hardcoded_role_mapper": {Tok: makeResource(ldapMod, "HardcodedRoleMapper")},
			"keycloak_ldap_msad_lds_user_account_control_mapper": {
				Tok: makeResource(ldapMod, "MsadLdsUserAccountControlMapper"),
			},
			"keycloak_ldap_hardcoded_group_mapper": {Tok: makeResource(ldapMod, "HardcodedGroupMapper")},
			"keycloak_ldap_role_mapper":            {Tok: makeResource(ldapMod, "RoleMapper")},

			"keycloak_oidc_identity_provider":        {Tok: makeResource(oidcMod, "IdentityProvider")},
			"keycloak_oidc_google_identity_provider": {Tok: makeResource(oidcMod, "GoogleIdentityProvider")},

			"keycloak_openid_audience_protocol_mapper": {Tok: makeResource(openIDMod, "AudienceProtocolMapper")},
			"keycloak_openid_client":                   {Tok: makeResource(openIDMod, "Client")},
			"keycloak_openid_client_authorization_permission": {
				Tok: makeResource(openIDMod, "ClientAuthorizationPermission"),
			},
			"keycloak_openid_client_authorization_resource": {Tok: makeResource(openIDMod, "ClientAuthorizationResource")},
			"keycloak_openid_client_authorization_scope":    {Tok: makeResource(openIDMod, "ClientAuthorizationScope")},
			"keycloak_openid_client_default_scopes":         {Tok: makeResource(openIDMod, "ClientDefaultScopes")},
			"keycloak_openid_client_optional_scopes":        {Tok: makeResource(openIDMod, "ClientOptionalScopes")},
			"keycloak_openid_client_scope":                  {Tok: makeResource(openIDMod, "ClientScope")},
			"keycloak_openid_client_service_account_role":   {Tok: makeResource(openIDMod, "ClientServiceAccountRole")},
			"keycloak_openid_full_name_protocol_mapper":     {Tok: makeResource(openIDMod, "FullNameProtocolMapper")},
			"keycloak_openid_group_membership_protocol_mapper": {
				Tok: makeResource(openIDMod, "GroupMembershipProtocolMapper"),
			},
			"keycloak_openid_hardcoded_claim_protocol_mapper": {
				Tok: makeResource(openIDMod, "HardcodedClaimProtocolMapper"),
			},
			"keycloak_openid_hardcoded_role_protocol_mapper": {
				Tok: makeResource(openIDMod, "HardcodedRoleProtocolMapper"),
			},
			"keycloak_openid_user_attribute_protocol_mapper": {
				Tok: makeResource(openIDMod, "UserAttributeProtocolMapper"),
			},
			"keycloak_openid_user_property_protocol_mapper": {
				Tok: makeResource(openIDMod, "UserPropertyProtocolMapper"),
			},
			"keycloak_openid_user_realm_role_protocol_mapper": {
				Tok: makeResource(openIDMod, "UserRealmRoleProtocolMapper"),
			},
			"keycloak_openid_client_service_account_realm_role": {
				Tok: makeResource(openIDMod, "ClientServiceAccountRealmRole"),
			},
			"keycloak_openid_client_aggregate_policy": {Tok: makeResource(openIDMod, "ClientAggregatePolicy")},
			"keycloak_openid_client_client_policy":    {Tok: makeResource(openIDMod, "ClientPolicy")},
			"keycloak_openid_client_group_policy":     {Tok: makeResource(openIDMod, "ClientGroupPolicy")},
			"keycloak_openid_client_js_policy":        {Tok: makeResource(openIDMod, "ClientJsPolicy")},
			"keycloak_openid_client_role_policy":      {Tok: makeResource(openIDMod, "ClientRolePolicy")},
			"keycloak_openid_client_time_policy":      {Tok: makeResource(openIDMod, "ClientTimePolicy")},
			"keycloak_openid_client_user_policy":      {Tok: makeResource(openIDMod, "ClientUserPolicy")},
			"keycloak_openid_user_client_role_protocol_mapper": {
				Tok: makeResource(openIDMod, "UserClientRoleProtocolMapper"),
			},
			"keycloak_openid_user_session_note_protocol_mapper": {
				Tok: makeResource(openIDMod, "UserSessionNoteProtocolMapper"),
			},

			"keycloak_saml_client":                         {Tok: makeResource(samlMod, "Client")},
			"keycloak_saml_identity_provider":              {Tok: makeResource(samlMod, "IdentityProvider")},
			"keycloak_saml_user_attribute_protocol_mapper": {Tok: makeResource(samlMod, "UserAttributeProtocolMapper")},
			"keycloak_saml_user_property_protocol_mapper":  {Tok: makeResource(samlMod, "UserPropertyProtocolMapper")},

			"keycloak_authentication_execution":        {Tok: makeResource(authenticationMod, "Execution")},
			"keycloak_authentication_execution_config": {Tok: makeResource(authenticationMod, "ExecutionConfig")},
			"keycloak_authentication_flow":             {Tok: makeResource(authenticationMod, "Flow")},
			"keycloak_authentication_subflow":          {Tok: makeResource(authenticationMod, "Subflow")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"keycloak_group":      {Tok: makeDataSource(mainMod, "getGroup")},
			"keycloak_realm":      {Tok: makeDataSource(mainMod, "getRealm")},
			"keycloak_realm_keys": {Tok: makeDataSource(mainMod, "getRealmKeys")},
			"keycloak_role":       {Tok: makeDataSource(mainMod, "getRole")},

			"keycloak_openid_client":                      {Tok: makeDataSource(openIDMod, "getClient")},
			"keycloak_openid_client_authorization_policy": {Tok: makeDataSource(openIDMod, "getClientAuthorizationPolicy")},
			"keycloak_openid_client_service_account_user": {Tok: makeDataSource(openIDMod, "getClientServiceAccountUser")},

			"keycloak_saml_client_installation_provider": {Tok: makeDataSource(samlMod, "getClientInstallationProvider")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=2.0.0,<3.0.0",
			},
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: namespaceMap,
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
