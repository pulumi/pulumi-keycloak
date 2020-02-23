// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package keycloak

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// ## # .AttributeImporterIdentityProviderMapper
// 
// Allows to create and manage identity provider mappers within Keycloak.
// 
// ### Argument Reference
// 
// The following arguments are supported:
// 
// - `realm` - (Required) The name of the realm.
// - `name` - (Required) The name of the mapper.
// - `identityProviderAlias` - (Required) The alias of the associated identity provider.
// - `userAttribute` - (Required) The user attribute name to store SAML attribute.
// - `attributeName` - (Optional) The Name of attribute to search for in assertion. You can leave this blank and specify a friendly name instead.
// - `attributeFriendlyName` - (Optional) The friendly name of attribute to search for in assertion.  You can leave this blank and specify an attribute name instead.
// - `claimName` - (Optional) The claim name.
// 
// > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/attribute_importer_identity_provider_mapper.html.markdown.
type AttributeImporterIdentityProviderMapper struct {
	pulumi.CustomResourceState

	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrOutput `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName pulumi.StringPtrOutput `pulumi:"attributeName"`
	// Claim Name
	ClaimName pulumi.StringPtrOutput `pulumi:"claimName"`
	// IDP Alias
	IdentityProviderAlias pulumi.StringOutput `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name pulumi.StringOutput `pulumi:"name"`
	// Realm Name
	Realm pulumi.StringOutput `pulumi:"realm"`
	// User Attribute
	UserAttribute pulumi.StringOutput `pulumi:"userAttribute"`
}

// NewAttributeImporterIdentityProviderMapper registers a new resource with the given unique name, arguments, and options.
func NewAttributeImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, args *AttributeImporterIdentityProviderMapperArgs, opts ...pulumi.ResourceOption) (*AttributeImporterIdentityProviderMapper, error) {
	if args == nil || args.IdentityProviderAlias == nil {
		return nil, errors.New("missing required argument 'IdentityProviderAlias'")
	}
	if args == nil || args.Realm == nil {
		return nil, errors.New("missing required argument 'Realm'")
	}
	if args == nil || args.UserAttribute == nil {
		return nil, errors.New("missing required argument 'UserAttribute'")
	}
	if args == nil {
		args = &AttributeImporterIdentityProviderMapperArgs{}
	}
	var resource AttributeImporterIdentityProviderMapper
	err := ctx.RegisterResource("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAttributeImporterIdentityProviderMapper gets an existing AttributeImporterIdentityProviderMapper resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAttributeImporterIdentityProviderMapper(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AttributeImporterIdentityProviderMapperState, opts ...pulumi.ResourceOption) (*AttributeImporterIdentityProviderMapper, error) {
	var resource AttributeImporterIdentityProviderMapper
	err := ctx.ReadResource("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AttributeImporterIdentityProviderMapper resources.
type attributeImporterIdentityProviderMapperState struct {
	// Attribute Friendly Name
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName *string `pulumi:"attributeName"`
	// Claim Name
	ClaimName *string `pulumi:"claimName"`
	// IDP Alias
	IdentityProviderAlias *string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm *string `pulumi:"realm"`
	// User Attribute
	UserAttribute *string `pulumi:"userAttribute"`
}

type AttributeImporterIdentityProviderMapperState struct {
	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrInput
	// Attribute Name
	AttributeName pulumi.StringPtrInput
	// Claim Name
	ClaimName pulumi.StringPtrInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringPtrInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringPtrInput
	// User Attribute
	UserAttribute pulumi.StringPtrInput
}

func (AttributeImporterIdentityProviderMapperState) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeImporterIdentityProviderMapperState)(nil)).Elem()
}

type attributeImporterIdentityProviderMapperArgs struct {
	// Attribute Friendly Name
	AttributeFriendlyName *string `pulumi:"attributeFriendlyName"`
	// Attribute Name
	AttributeName *string `pulumi:"attributeName"`
	// Claim Name
	ClaimName *string `pulumi:"claimName"`
	// IDP Alias
	IdentityProviderAlias string `pulumi:"identityProviderAlias"`
	// IDP Mapper Name
	Name *string `pulumi:"name"`
	// Realm Name
	Realm string `pulumi:"realm"`
	// User Attribute
	UserAttribute string `pulumi:"userAttribute"`
}

// The set of arguments for constructing a AttributeImporterIdentityProviderMapper resource.
type AttributeImporterIdentityProviderMapperArgs struct {
	// Attribute Friendly Name
	AttributeFriendlyName pulumi.StringPtrInput
	// Attribute Name
	AttributeName pulumi.StringPtrInput
	// Claim Name
	ClaimName pulumi.StringPtrInput
	// IDP Alias
	IdentityProviderAlias pulumi.StringInput
	// IDP Mapper Name
	Name pulumi.StringPtrInput
	// Realm Name
	Realm pulumi.StringInput
	// User Attribute
	UserAttribute pulumi.StringInput
}

func (AttributeImporterIdentityProviderMapperArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*attributeImporterIdentityProviderMapperArgs)(nil)).Elem()
}

