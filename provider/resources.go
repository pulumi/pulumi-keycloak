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
	"fmt"
	"path"
	"strings"

	// embed is used to store bridge-metadata.json in the compiled binary
	_ "embed"

	"github.com/mrparkers/terraform-provider-keycloak/provider"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	tks "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"

	"github.com/pulumi/pulumi-keycloak/provider/v5/pkg/version"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "keycloak"
	// modules:
	mainMod   = "index"  // the main module
	openIDMod = "OpenId" // the openid module
	samlMod   = "Saml"   // the Saml module
)

func makeDataSource(mod string, res string) tokens.ModuleMember {
	mod = strings.ToLower(mod)
	return tfbridge.MakeDataSource(mainPkg, mod, res)
}

func makeResource(mod string, res string) tokens.Type {
	mod = strings.ToLower(mod)
	return tfbridge.MakeResource(mainPkg, mod, res)
}

func ref[T any](t T) *T { return &t }

//go:embed cmd/pulumi-resource-keycloak/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(provider.KeycloakProvider(nil))

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
		TFProviderLicense: ref(tfbridge.MITLicenseType),
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		UpstreamRepoPath:  "./upstream",
		Version:           version.Version,

		Config: map[string]*tfbridge.SchemaInfo{
			"client_timeout": {
				Default: &tfbridge.DefaultInfo{
					Value:   5,
					EnvVars: []string{"KEYCLOAK_CLIENT_TIMEOUT"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"keycloak_attribute_to_role_identity_provider_mapper": {
				Tok:  makeResource(mainMod, "AttributeToRoleIdentityMapper"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_custom_identity_provider_mapper": {
				Tok: makeResource(mainMod, "CustomIdentityProviderMapping"),
			},
			"keycloak_default_roles": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"default_roles": {
						CSharpName: "RoleNames",
					},
				},
			},
			"keycloak_hardcoded_attribute_identity_provider_mapper": {
				Tok:  makeResource(mainMod, "HardcodedAttributeIdentityProviderMapper"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_hardcoded_role_identity_provider_mapper": {
				Tok:  makeResource(mainMod, "HardcodedRoleIdentityMapper"),
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_realm": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"realm": {CSharpName: "RealmName"},
				},
			},
			"keycloak_users_permissions": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_realm_keystore_java_keystore": {
				Tok: makeResource(mainMod, "RealmKeystoreJavaGenerated"),
			},
			"keycloak_ldap_hardcoded_group_mapper": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_authorization_scope": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_authorization_permission": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_authorization_resource": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_aggregate_policy": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_client_policy": {Tok: makeResource(openIDMod, "ClientPolicy")},
			"keycloak_openid_client_group_policy": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_js_policy": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_role_policy": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_time_policy": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_user_policy": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_client_permissions": {
				Docs: &tfbridge.DocInfo{AllowMissing: true},
			},
			"keycloak_openid_audience_resolve_protocol_mapper": {
				Docs: &tfbridge.DocInfo{Source: "openid_audience_resolve_protocol_mapper.md"},
			},

			"keycloak_saml_client_default_scopes": {Tok: makeResource(samlMod, "ClientDefaultScope")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"keycloak_authentication_execution": {Tok: makeDataSource(mainMod, "getAuthenticationExecution")},
			"keycloak_authentication_flow":      {Tok: makeDataSource(mainMod, "getAuthenticationFlow")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject: struct{ Enabled bool }{true},
		},

		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				"authentication": "Authentication",
				"index":          "index",
				"keycloak":       "Keycloak",
				"ldap":           "Ldap",
				"oidc":           "Oidc",
				"openid":         "OpenId",
				"saml":           "Saml",
			},
		},
	}

	prov.MustComputeTokens(tks.KnownModules("keycloak_", mainMod, []string{
		"ldap_",
		"oidc_",
		"openid_",
		"saml_",
		"authentication_",
	}, tks.MakeStandard(mainPkg)))

	prov.MustApplyAutoAliases()

	prov.SetAutonaming(255, "-")

	return prov
}
