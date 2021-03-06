// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper":
		r = &AttributeImporterIdentityProviderMapper{}
	case "keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper":
		r = &AttributeToRoleIdentityMapper{}
	case "keycloak:index/customUserFederation:CustomUserFederation":
		r = &CustomUserFederation{}
	case "keycloak:index/defaultGroups:DefaultGroups":
		r = &DefaultGroups{}
	case "keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper":
		r = &GenericClientProtocolMapper{}
	case "keycloak:index/genericClientRoleMapper:GenericClientRoleMapper":
		r = &GenericClientRoleMapper{}
	case "keycloak:index/group:Group":
		r = &Group{}
	case "keycloak:index/groupMemberships:GroupMemberships":
		r = &GroupMemberships{}
	case "keycloak:index/groupRoles:GroupRoles":
		r = &GroupRoles{}
	case "keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper":
		r = &HardcodedAttributeIdentityProviderMapper{}
	case "keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper":
		r = &HardcodedRoleIdentityMapper{}
	case "keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission":
		r = &IdentityProviderTokenExchangeScopePermission{}
	case "keycloak:index/realm:Realm":
		r = &Realm{}
	case "keycloak:index/realmEvents:RealmEvents":
		r = &RealmEvents{}
	case "keycloak:index/requiredAction:RequiredAction":
		r = &RequiredAction{}
	case "keycloak:index/role:Role":
		r = &Role{}
	case "keycloak:index/user:User":
		r = &User{}
	case "keycloak:index/userGroups:UserGroups":
		r = &UserGroups{}
	case "keycloak:index/userRoles:UserRoles":
		r = &UserRoles{}
	case "keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper":
		r = &UserTemplateImporterIdentityProviderMapper{}
	case "keycloak:index/usersPermissions:UsersPermissions":
		r = &UsersPermissions{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:keycloak" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/attributeImporterIdentityProviderMapper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/attributeToRoleIdentityMapper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/customUserFederation",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/defaultGroups",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/genericClientProtocolMapper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/genericClientRoleMapper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/group",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/groupMemberships",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/groupRoles",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/hardcodedAttributeIdentityProviderMapper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/hardcodedRoleIdentityMapper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/identityProviderTokenExchangeScopePermission",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/realm",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/realmEvents",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/requiredAction",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/role",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/user",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/userGroups",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/userRoles",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/userTemplateImporterIdentityProviderMapper",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"index/usersPermissions",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"keycloak",
		&pkg{version},
	)
}
