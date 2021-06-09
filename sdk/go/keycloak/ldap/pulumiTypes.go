// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ldap

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type UserFederationCache struct {
	// Day of the week the entry will become invalid on
	EvictionDay *int `pulumi:"evictionDay"`
	// Hour of day the entry will become invalid on.
	EvictionHour *int `pulumi:"evictionHour"`
	// Minute of day the entry will become invalid on.
	EvictionMinute *int `pulumi:"evictionMinute"`
	// Max lifespan of cache entry (duration string).
	MaxLifespan *string `pulumi:"maxLifespan"`
	// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	Policy *string `pulumi:"policy"`
}

// UserFederationCacheInput is an input type that accepts UserFederationCacheArgs and UserFederationCacheOutput values.
// You can construct a concrete instance of `UserFederationCacheInput` via:
//
//          UserFederationCacheArgs{...}
type UserFederationCacheInput interface {
	pulumi.Input

	ToUserFederationCacheOutput() UserFederationCacheOutput
	ToUserFederationCacheOutputWithContext(context.Context) UserFederationCacheOutput
}

type UserFederationCacheArgs struct {
	// Day of the week the entry will become invalid on
	EvictionDay pulumi.IntPtrInput `pulumi:"evictionDay"`
	// Hour of day the entry will become invalid on.
	EvictionHour pulumi.IntPtrInput `pulumi:"evictionHour"`
	// Minute of day the entry will become invalid on.
	EvictionMinute pulumi.IntPtrInput `pulumi:"evictionMinute"`
	// Max lifespan of cache entry (duration string).
	MaxLifespan pulumi.StringPtrInput `pulumi:"maxLifespan"`
	// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
	Policy pulumi.StringPtrInput `pulumi:"policy"`
}

func (UserFederationCacheArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserFederationCache)(nil)).Elem()
}

func (i UserFederationCacheArgs) ToUserFederationCacheOutput() UserFederationCacheOutput {
	return i.ToUserFederationCacheOutputWithContext(context.Background())
}

func (i UserFederationCacheArgs) ToUserFederationCacheOutputWithContext(ctx context.Context) UserFederationCacheOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFederationCacheOutput)
}

func (i UserFederationCacheArgs) ToUserFederationCachePtrOutput() UserFederationCachePtrOutput {
	return i.ToUserFederationCachePtrOutputWithContext(context.Background())
}

func (i UserFederationCacheArgs) ToUserFederationCachePtrOutputWithContext(ctx context.Context) UserFederationCachePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFederationCacheOutput).ToUserFederationCachePtrOutputWithContext(ctx)
}

// UserFederationCachePtrInput is an input type that accepts UserFederationCacheArgs, UserFederationCachePtr and UserFederationCachePtrOutput values.
// You can construct a concrete instance of `UserFederationCachePtrInput` via:
//
//          UserFederationCacheArgs{...}
//
//  or:
//
//          nil
type UserFederationCachePtrInput interface {
	pulumi.Input

	ToUserFederationCachePtrOutput() UserFederationCachePtrOutput
	ToUserFederationCachePtrOutputWithContext(context.Context) UserFederationCachePtrOutput
}

type userFederationCachePtrType UserFederationCacheArgs

func UserFederationCachePtr(v *UserFederationCacheArgs) UserFederationCachePtrInput {
	return (*userFederationCachePtrType)(v)
}

func (*userFederationCachePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFederationCache)(nil)).Elem()
}

func (i *userFederationCachePtrType) ToUserFederationCachePtrOutput() UserFederationCachePtrOutput {
	return i.ToUserFederationCachePtrOutputWithContext(context.Background())
}

func (i *userFederationCachePtrType) ToUserFederationCachePtrOutputWithContext(ctx context.Context) UserFederationCachePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFederationCachePtrOutput)
}

type UserFederationCacheOutput struct{ *pulumi.OutputState }

func (UserFederationCacheOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserFederationCache)(nil)).Elem()
}

func (o UserFederationCacheOutput) ToUserFederationCacheOutput() UserFederationCacheOutput {
	return o
}

func (o UserFederationCacheOutput) ToUserFederationCacheOutputWithContext(ctx context.Context) UserFederationCacheOutput {
	return o
}

func (o UserFederationCacheOutput) ToUserFederationCachePtrOutput() UserFederationCachePtrOutput {
	return o.ToUserFederationCachePtrOutputWithContext(context.Background())
}

func (o UserFederationCacheOutput) ToUserFederationCachePtrOutputWithContext(ctx context.Context) UserFederationCachePtrOutput {
	return o.ApplyT(func(v UserFederationCache) *UserFederationCache {
		return &v
	}).(UserFederationCachePtrOutput)
}

// Day of the week the entry will become invalid on
func (o UserFederationCacheOutput) EvictionDay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v UserFederationCache) *int { return v.EvictionDay }).(pulumi.IntPtrOutput)
}

// Hour of day the entry will become invalid on.
func (o UserFederationCacheOutput) EvictionHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v UserFederationCache) *int { return v.EvictionHour }).(pulumi.IntPtrOutput)
}

// Minute of day the entry will become invalid on.
func (o UserFederationCacheOutput) EvictionMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v UserFederationCache) *int { return v.EvictionMinute }).(pulumi.IntPtrOutput)
}

// Max lifespan of cache entry (duration string).
func (o UserFederationCacheOutput) MaxLifespan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserFederationCache) *string { return v.MaxLifespan }).(pulumi.StringPtrOutput)
}

// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
func (o UserFederationCacheOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserFederationCache) *string { return v.Policy }).(pulumi.StringPtrOutput)
}

type UserFederationCachePtrOutput struct{ *pulumi.OutputState }

func (UserFederationCachePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFederationCache)(nil)).Elem()
}

func (o UserFederationCachePtrOutput) ToUserFederationCachePtrOutput() UserFederationCachePtrOutput {
	return o
}

func (o UserFederationCachePtrOutput) ToUserFederationCachePtrOutputWithContext(ctx context.Context) UserFederationCachePtrOutput {
	return o
}

func (o UserFederationCachePtrOutput) Elem() UserFederationCacheOutput {
	return o.ApplyT(func(v *UserFederationCache) UserFederationCache { return *v }).(UserFederationCacheOutput)
}

// Day of the week the entry will become invalid on
func (o UserFederationCachePtrOutput) EvictionDay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserFederationCache) *int {
		if v == nil {
			return nil
		}
		return v.EvictionDay
	}).(pulumi.IntPtrOutput)
}

// Hour of day the entry will become invalid on.
func (o UserFederationCachePtrOutput) EvictionHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserFederationCache) *int {
		if v == nil {
			return nil
		}
		return v.EvictionHour
	}).(pulumi.IntPtrOutput)
}

// Minute of day the entry will become invalid on.
func (o UserFederationCachePtrOutput) EvictionMinute() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *UserFederationCache) *int {
		if v == nil {
			return nil
		}
		return v.EvictionMinute
	}).(pulumi.IntPtrOutput)
}

// Max lifespan of cache entry (duration string).
func (o UserFederationCachePtrOutput) MaxLifespan() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFederationCache) *string {
		if v == nil {
			return nil
		}
		return v.MaxLifespan
	}).(pulumi.StringPtrOutput)
}

// Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
func (o UserFederationCachePtrOutput) Policy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFederationCache) *string {
		if v == nil {
			return nil
		}
		return v.Policy
	}).(pulumi.StringPtrOutput)
}

type UserFederationKerberos struct {
	// The name of the kerberos realm, e.g. FOO.LOCAL.
	KerberosRealm string `pulumi:"kerberosRealm"`
	// Path to the kerberos keytab file on the server with credentials of the service principal.
	KeyTab string `pulumi:"keyTab"`
	// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
	ServerPrincipal string `pulumi:"serverPrincipal"`
	// Use kerberos login module instead of ldap service api. Defaults to `false`.
	UseKerberosForPasswordAuthentication *bool `pulumi:"useKerberosForPasswordAuthentication"`
}

// UserFederationKerberosInput is an input type that accepts UserFederationKerberosArgs and UserFederationKerberosOutput values.
// You can construct a concrete instance of `UserFederationKerberosInput` via:
//
//          UserFederationKerberosArgs{...}
type UserFederationKerberosInput interface {
	pulumi.Input

	ToUserFederationKerberosOutput() UserFederationKerberosOutput
	ToUserFederationKerberosOutputWithContext(context.Context) UserFederationKerberosOutput
}

type UserFederationKerberosArgs struct {
	// The name of the kerberos realm, e.g. FOO.LOCAL.
	KerberosRealm pulumi.StringInput `pulumi:"kerberosRealm"`
	// Path to the kerberos keytab file on the server with credentials of the service principal.
	KeyTab pulumi.StringInput `pulumi:"keyTab"`
	// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
	ServerPrincipal pulumi.StringInput `pulumi:"serverPrincipal"`
	// Use kerberos login module instead of ldap service api. Defaults to `false`.
	UseKerberosForPasswordAuthentication pulumi.BoolPtrInput `pulumi:"useKerberosForPasswordAuthentication"`
}

func (UserFederationKerberosArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserFederationKerberos)(nil)).Elem()
}

func (i UserFederationKerberosArgs) ToUserFederationKerberosOutput() UserFederationKerberosOutput {
	return i.ToUserFederationKerberosOutputWithContext(context.Background())
}

func (i UserFederationKerberosArgs) ToUserFederationKerberosOutputWithContext(ctx context.Context) UserFederationKerberosOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFederationKerberosOutput)
}

func (i UserFederationKerberosArgs) ToUserFederationKerberosPtrOutput() UserFederationKerberosPtrOutput {
	return i.ToUserFederationKerberosPtrOutputWithContext(context.Background())
}

func (i UserFederationKerberosArgs) ToUserFederationKerberosPtrOutputWithContext(ctx context.Context) UserFederationKerberosPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFederationKerberosOutput).ToUserFederationKerberosPtrOutputWithContext(ctx)
}

// UserFederationKerberosPtrInput is an input type that accepts UserFederationKerberosArgs, UserFederationKerberosPtr and UserFederationKerberosPtrOutput values.
// You can construct a concrete instance of `UserFederationKerberosPtrInput` via:
//
//          UserFederationKerberosArgs{...}
//
//  or:
//
//          nil
type UserFederationKerberosPtrInput interface {
	pulumi.Input

	ToUserFederationKerberosPtrOutput() UserFederationKerberosPtrOutput
	ToUserFederationKerberosPtrOutputWithContext(context.Context) UserFederationKerberosPtrOutput
}

type userFederationKerberosPtrType UserFederationKerberosArgs

func UserFederationKerberosPtr(v *UserFederationKerberosArgs) UserFederationKerberosPtrInput {
	return (*userFederationKerberosPtrType)(v)
}

func (*userFederationKerberosPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFederationKerberos)(nil)).Elem()
}

func (i *userFederationKerberosPtrType) ToUserFederationKerberosPtrOutput() UserFederationKerberosPtrOutput {
	return i.ToUserFederationKerberosPtrOutputWithContext(context.Background())
}

func (i *userFederationKerberosPtrType) ToUserFederationKerberosPtrOutputWithContext(ctx context.Context) UserFederationKerberosPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserFederationKerberosPtrOutput)
}

type UserFederationKerberosOutput struct{ *pulumi.OutputState }

func (UserFederationKerberosOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserFederationKerberos)(nil)).Elem()
}

func (o UserFederationKerberosOutput) ToUserFederationKerberosOutput() UserFederationKerberosOutput {
	return o
}

func (o UserFederationKerberosOutput) ToUserFederationKerberosOutputWithContext(ctx context.Context) UserFederationKerberosOutput {
	return o
}

func (o UserFederationKerberosOutput) ToUserFederationKerberosPtrOutput() UserFederationKerberosPtrOutput {
	return o.ToUserFederationKerberosPtrOutputWithContext(context.Background())
}

func (o UserFederationKerberosOutput) ToUserFederationKerberosPtrOutputWithContext(ctx context.Context) UserFederationKerberosPtrOutput {
	return o.ApplyT(func(v UserFederationKerberos) *UserFederationKerberos {
		return &v
	}).(UserFederationKerberosPtrOutput)
}

// The name of the kerberos realm, e.g. FOO.LOCAL.
func (o UserFederationKerberosOutput) KerberosRealm() pulumi.StringOutput {
	return o.ApplyT(func(v UserFederationKerberos) string { return v.KerberosRealm }).(pulumi.StringOutput)
}

// Path to the kerberos keytab file on the server with credentials of the service principal.
func (o UserFederationKerberosOutput) KeyTab() pulumi.StringOutput {
	return o.ApplyT(func(v UserFederationKerberos) string { return v.KeyTab }).(pulumi.StringOutput)
}

// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
func (o UserFederationKerberosOutput) ServerPrincipal() pulumi.StringOutput {
	return o.ApplyT(func(v UserFederationKerberos) string { return v.ServerPrincipal }).(pulumi.StringOutput)
}

// Use kerberos login module instead of ldap service api. Defaults to `false`.
func (o UserFederationKerberosOutput) UseKerberosForPasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v UserFederationKerberos) *bool { return v.UseKerberosForPasswordAuthentication }).(pulumi.BoolPtrOutput)
}

type UserFederationKerberosPtrOutput struct{ *pulumi.OutputState }

func (UserFederationKerberosPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserFederationKerberos)(nil)).Elem()
}

func (o UserFederationKerberosPtrOutput) ToUserFederationKerberosPtrOutput() UserFederationKerberosPtrOutput {
	return o
}

func (o UserFederationKerberosPtrOutput) ToUserFederationKerberosPtrOutputWithContext(ctx context.Context) UserFederationKerberosPtrOutput {
	return o
}

func (o UserFederationKerberosPtrOutput) Elem() UserFederationKerberosOutput {
	return o.ApplyT(func(v *UserFederationKerberos) UserFederationKerberos { return *v }).(UserFederationKerberosOutput)
}

// The name of the kerberos realm, e.g. FOO.LOCAL.
func (o UserFederationKerberosPtrOutput) KerberosRealm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFederationKerberos) *string {
		if v == nil {
			return nil
		}
		return &v.KerberosRealm
	}).(pulumi.StringPtrOutput)
}

// Path to the kerberos keytab file on the server with credentials of the service principal.
func (o UserFederationKerberosPtrOutput) KeyTab() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFederationKerberos) *string {
		if v == nil {
			return nil
		}
		return &v.KeyTab
	}).(pulumi.StringPtrOutput)
}

// The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
func (o UserFederationKerberosPtrOutput) ServerPrincipal() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserFederationKerberos) *string {
		if v == nil {
			return nil
		}
		return &v.ServerPrincipal
	}).(pulumi.StringPtrOutput)
}

// Use kerberos login module instead of ldap service api. Defaults to `false`.
func (o UserFederationKerberosPtrOutput) UseKerberosForPasswordAuthentication() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *UserFederationKerberos) *bool {
		if v == nil {
			return nil
		}
		return v.UseKerberosForPasswordAuthentication
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(UserFederationCacheOutput{})
	pulumi.RegisterOutputType(UserFederationCachePtrOutput{})
	pulumi.RegisterOutputType(UserFederationKerberosOutput{})
	pulumi.RegisterOutputType(UserFederationKerberosPtrOutput{})
}
