// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package authentication

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
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
	case "keycloak:authentication/bindings:Bindings":
		r = &Bindings{}
	case "keycloak:authentication/execution:Execution":
		r = &Execution{}
	case "keycloak:authentication/executionConfig:ExecutionConfig":
		r = &ExecutionConfig{}
	case "keycloak:authentication/flow:Flow":
		r = &Flow{}
	case "keycloak:authentication/subflow:Subflow":
		r = &Subflow{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"keycloak",
		"authentication/bindings",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"authentication/execution",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"authentication/executionConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"authentication/flow",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"keycloak",
		"authentication/subflow",
		&module{version},
	)
}
