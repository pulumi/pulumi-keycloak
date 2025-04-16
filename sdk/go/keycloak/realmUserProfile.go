// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package keycloak

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Allows for managing Realm User Profiles within Keycloak.
//
// A user profile defines a schema for representing user attributes and how they are managed within a realm.
//
// Information for Keycloak versions < 24:
// The realm linked to the `RealmUserProfile` resource must have the user profile feature enabled.
// It can be done via the administration UI, or by setting the `userProfileEnabled` realm attribute to `true`.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-keycloak/sdk/v6/go/keycloak"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := keycloak.NewRealm(ctx, "realm", &keycloak.RealmArgs{
//				Realm: pulumi.String("my-realm"),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal([]string{
//				"opt1",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"key": "val",
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			tmpJSON2, err := json.Marshal(map[string]interface{}{
//				"key": "val",
//			})
//			if err != nil {
//				return err
//			}
//			json2 := string(tmpJSON2)
//			_, err = keycloak.NewRealmUserProfile(ctx, "userprofile", &keycloak.RealmUserProfileArgs{
//				RealmId:                  pulumi.Any(myRealm.Id),
//				UnmanagedAttributePolicy: pulumi.String("ENABLED"),
//				Attributes: keycloak.RealmUserProfileAttributeArray{
//					&keycloak.RealmUserProfileAttributeArgs{
//						Name:        pulumi.String("field1"),
//						DisplayName: pulumi.String("Field 1"),
//						Group:       pulumi.String("group1"),
//						MultiValued: pulumi.Bool(false),
//						EnabledWhenScopes: pulumi.StringArray{
//							pulumi.String("offline_access"),
//						},
//						RequiredForRoles: pulumi.StringArray{
//							pulumi.String("user"),
//						},
//						RequiredForScopes: pulumi.StringArray{
//							pulumi.String("offline_access"),
//						},
//						Permissions: &keycloak.RealmUserProfileAttributePermissionsArgs{
//							Views: pulumi.StringArray{
//								pulumi.String("admin"),
//								pulumi.String("user"),
//							},
//							Edits: pulumi.StringArray{
//								pulumi.String("admin"),
//								pulumi.String("user"),
//							},
//						},
//						Validators: keycloak.RealmUserProfileAttributeValidatorArray{
//							&keycloak.RealmUserProfileAttributeValidatorArgs{
//								Name: pulumi.String("person-name-prohibited-characters"),
//							},
//							&keycloak.RealmUserProfileAttributeValidatorArgs{
//								Name: pulumi.String("pattern"),
//								Config: pulumi.StringMap{
//									"pattern":       pulumi.String("^[a-z]+$"),
//									"error-message": pulumi.String("Nope"),
//								},
//							},
//						},
//						Annotations: pulumi.StringMap{
//							"foo": pulumi.String("bar"),
//						},
//					},
//					&keycloak.RealmUserProfileAttributeArgs{
//						Name: pulumi.String("field2"),
//						Validators: keycloak.RealmUserProfileAttributeValidatorArray{
//							&keycloak.RealmUserProfileAttributeValidatorArgs{
//								Name: pulumi.String("options"),
//								Config: pulumi.StringMap{
//									"options": pulumi.String(json0),
//								},
//							},
//						},
//						Annotations: pulumi.StringMap{
//							"foo": pulumi.String(json1),
//						},
//					},
//				},
//				Groups: keycloak.RealmUserProfileGroupArray{
//					&keycloak.RealmUserProfileGroupArgs{
//						Name:               pulumi.String("group1"),
//						DisplayHeader:      pulumi.String("Group 1"),
//						DisplayDescription: pulumi.String("A first group"),
//						Annotations: pulumi.StringMap{
//							"foo":  pulumi.String("bar"),
//							"foo2": pulumi.String(json2),
//						},
//					},
//					&keycloak.RealmUserProfileGroupArgs{
//						Name: pulumi.String("group2"),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// This resource currently does not support importing.
type RealmUserProfile struct {
	pulumi.CustomResourceState

	// An ordered list of attributes.
	Attributes RealmUserProfileAttributeArrayOutput `pulumi:"attributes"`
	// A list of groups.
	Groups RealmUserProfileGroupArrayOutput `pulumi:"groups"`
	// The ID of the realm the user profile applies to.
	RealmId pulumi.StringOutput `pulumi:"realmId"`
	// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
	UnmanagedAttributePolicy pulumi.StringPtrOutput `pulumi:"unmanagedAttributePolicy"`
}

// NewRealmUserProfile registers a new resource with the given unique name, arguments, and options.
func NewRealmUserProfile(ctx *pulumi.Context,
	name string, args *RealmUserProfileArgs, opts ...pulumi.ResourceOption) (*RealmUserProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RealmId == nil {
		return nil, errors.New("invalid value for required argument 'RealmId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RealmUserProfile
	err := ctx.RegisterResource("keycloak:index/realmUserProfile:RealmUserProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRealmUserProfile gets an existing RealmUserProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRealmUserProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RealmUserProfileState, opts ...pulumi.ResourceOption) (*RealmUserProfile, error) {
	var resource RealmUserProfile
	err := ctx.ReadResource("keycloak:index/realmUserProfile:RealmUserProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RealmUserProfile resources.
type realmUserProfileState struct {
	// An ordered list of attributes.
	Attributes []RealmUserProfileAttribute `pulumi:"attributes"`
	// A list of groups.
	Groups []RealmUserProfileGroup `pulumi:"groups"`
	// The ID of the realm the user profile applies to.
	RealmId *string `pulumi:"realmId"`
	// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
	UnmanagedAttributePolicy *string `pulumi:"unmanagedAttributePolicy"`
}

type RealmUserProfileState struct {
	// An ordered list of attributes.
	Attributes RealmUserProfileAttributeArrayInput
	// A list of groups.
	Groups RealmUserProfileGroupArrayInput
	// The ID of the realm the user profile applies to.
	RealmId pulumi.StringPtrInput
	// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
	UnmanagedAttributePolicy pulumi.StringPtrInput
}

func (RealmUserProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*realmUserProfileState)(nil)).Elem()
}

type realmUserProfileArgs struct {
	// An ordered list of attributes.
	Attributes []RealmUserProfileAttribute `pulumi:"attributes"`
	// A list of groups.
	Groups []RealmUserProfileGroup `pulumi:"groups"`
	// The ID of the realm the user profile applies to.
	RealmId string `pulumi:"realmId"`
	// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
	UnmanagedAttributePolicy *string `pulumi:"unmanagedAttributePolicy"`
}

// The set of arguments for constructing a RealmUserProfile resource.
type RealmUserProfileArgs struct {
	// An ordered list of attributes.
	Attributes RealmUserProfileAttributeArrayInput
	// A list of groups.
	Groups RealmUserProfileGroupArrayInput
	// The ID of the realm the user profile applies to.
	RealmId pulumi.StringInput
	// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
	UnmanagedAttributePolicy pulumi.StringPtrInput
}

func (RealmUserProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*realmUserProfileArgs)(nil)).Elem()
}

type RealmUserProfileInput interface {
	pulumi.Input

	ToRealmUserProfileOutput() RealmUserProfileOutput
	ToRealmUserProfileOutputWithContext(ctx context.Context) RealmUserProfileOutput
}

func (*RealmUserProfile) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmUserProfile)(nil)).Elem()
}

func (i *RealmUserProfile) ToRealmUserProfileOutput() RealmUserProfileOutput {
	return i.ToRealmUserProfileOutputWithContext(context.Background())
}

func (i *RealmUserProfile) ToRealmUserProfileOutputWithContext(ctx context.Context) RealmUserProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmUserProfileOutput)
}

// RealmUserProfileArrayInput is an input type that accepts RealmUserProfileArray and RealmUserProfileArrayOutput values.
// You can construct a concrete instance of `RealmUserProfileArrayInput` via:
//
//	RealmUserProfileArray{ RealmUserProfileArgs{...} }
type RealmUserProfileArrayInput interface {
	pulumi.Input

	ToRealmUserProfileArrayOutput() RealmUserProfileArrayOutput
	ToRealmUserProfileArrayOutputWithContext(context.Context) RealmUserProfileArrayOutput
}

type RealmUserProfileArray []RealmUserProfileInput

func (RealmUserProfileArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmUserProfile)(nil)).Elem()
}

func (i RealmUserProfileArray) ToRealmUserProfileArrayOutput() RealmUserProfileArrayOutput {
	return i.ToRealmUserProfileArrayOutputWithContext(context.Background())
}

func (i RealmUserProfileArray) ToRealmUserProfileArrayOutputWithContext(ctx context.Context) RealmUserProfileArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmUserProfileArrayOutput)
}

// RealmUserProfileMapInput is an input type that accepts RealmUserProfileMap and RealmUserProfileMapOutput values.
// You can construct a concrete instance of `RealmUserProfileMapInput` via:
//
//	RealmUserProfileMap{ "key": RealmUserProfileArgs{...} }
type RealmUserProfileMapInput interface {
	pulumi.Input

	ToRealmUserProfileMapOutput() RealmUserProfileMapOutput
	ToRealmUserProfileMapOutputWithContext(context.Context) RealmUserProfileMapOutput
}

type RealmUserProfileMap map[string]RealmUserProfileInput

func (RealmUserProfileMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmUserProfile)(nil)).Elem()
}

func (i RealmUserProfileMap) ToRealmUserProfileMapOutput() RealmUserProfileMapOutput {
	return i.ToRealmUserProfileMapOutputWithContext(context.Background())
}

func (i RealmUserProfileMap) ToRealmUserProfileMapOutputWithContext(ctx context.Context) RealmUserProfileMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RealmUserProfileMapOutput)
}

type RealmUserProfileOutput struct{ *pulumi.OutputState }

func (RealmUserProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RealmUserProfile)(nil)).Elem()
}

func (o RealmUserProfileOutput) ToRealmUserProfileOutput() RealmUserProfileOutput {
	return o
}

func (o RealmUserProfileOutput) ToRealmUserProfileOutputWithContext(ctx context.Context) RealmUserProfileOutput {
	return o
}

// An ordered list of attributes.
func (o RealmUserProfileOutput) Attributes() RealmUserProfileAttributeArrayOutput {
	return o.ApplyT(func(v *RealmUserProfile) RealmUserProfileAttributeArrayOutput { return v.Attributes }).(RealmUserProfileAttributeArrayOutput)
}

// A list of groups.
func (o RealmUserProfileOutput) Groups() RealmUserProfileGroupArrayOutput {
	return o.ApplyT(func(v *RealmUserProfile) RealmUserProfileGroupArrayOutput { return v.Groups }).(RealmUserProfileGroupArrayOutput)
}

// The ID of the realm the user profile applies to.
func (o RealmUserProfileOutput) RealmId() pulumi.StringOutput {
	return o.ApplyT(func(v *RealmUserProfile) pulumi.StringOutput { return v.RealmId }).(pulumi.StringOutput)
}

// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
func (o RealmUserProfileOutput) UnmanagedAttributePolicy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RealmUserProfile) pulumi.StringPtrOutput { return v.UnmanagedAttributePolicy }).(pulumi.StringPtrOutput)
}

type RealmUserProfileArrayOutput struct{ *pulumi.OutputState }

func (RealmUserProfileArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RealmUserProfile)(nil)).Elem()
}

func (o RealmUserProfileArrayOutput) ToRealmUserProfileArrayOutput() RealmUserProfileArrayOutput {
	return o
}

func (o RealmUserProfileArrayOutput) ToRealmUserProfileArrayOutputWithContext(ctx context.Context) RealmUserProfileArrayOutput {
	return o
}

func (o RealmUserProfileArrayOutput) Index(i pulumi.IntInput) RealmUserProfileOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RealmUserProfile {
		return vs[0].([]*RealmUserProfile)[vs[1].(int)]
	}).(RealmUserProfileOutput)
}

type RealmUserProfileMapOutput struct{ *pulumi.OutputState }

func (RealmUserProfileMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RealmUserProfile)(nil)).Elem()
}

func (o RealmUserProfileMapOutput) ToRealmUserProfileMapOutput() RealmUserProfileMapOutput {
	return o
}

func (o RealmUserProfileMapOutput) ToRealmUserProfileMapOutputWithContext(ctx context.Context) RealmUserProfileMapOutput {
	return o
}

func (o RealmUserProfileMapOutput) MapIndex(k pulumi.StringInput) RealmUserProfileOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RealmUserProfile {
		return vs[0].(map[string]*RealmUserProfile)[vs[1].(string)]
	}).(RealmUserProfileOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RealmUserProfileInput)(nil)).Elem(), &RealmUserProfile{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmUserProfileArrayInput)(nil)).Elem(), RealmUserProfileArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RealmUserProfileMapInput)(nil)).Elem(), RealmUserProfileMap{})
	pulumi.RegisterOutputType(RealmUserProfileOutput{})
	pulumi.RegisterOutputType(RealmUserProfileArrayOutput{})
	pulumi.RegisterOutputType(RealmUserProfileMapOutput{})
}
