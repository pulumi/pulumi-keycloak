# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['RealmDefaultClientScopesArgs', 'RealmDefaultClientScopes']

@pulumi.input_type
class RealmDefaultClientScopesArgs:
    def __init__(__self__, *,
                 default_scopes: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 realm_id: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a RealmDefaultClientScopes resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] default_scopes: An array of default client scope names that should be used when creating new Keycloak clients.
        :param pulumi.Input[builtins.str] realm_id: The realm this client and scopes exists in.
        """
        pulumi.set(__self__, "default_scopes", default_scopes)
        pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="defaultScopes")
    def default_scopes(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        An array of default client scope names that should be used when creating new Keycloak clients.
        """
        return pulumi.get(self, "default_scopes")

    @default_scopes.setter
    def default_scopes(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "default_scopes", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[builtins.str]:
        """
        The realm this client and scopes exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "realm_id", value)


@pulumi.input_type
class _RealmDefaultClientScopesState:
    def __init__(__self__, *,
                 default_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering RealmDefaultClientScopes resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] default_scopes: An array of default client scope names that should be used when creating new Keycloak clients.
        :param pulumi.Input[builtins.str] realm_id: The realm this client and scopes exists in.
        """
        if default_scopes is not None:
            pulumi.set(__self__, "default_scopes", default_scopes)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="defaultScopes")
    def default_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        An array of default client scope names that should be used when creating new Keycloak clients.
        """
        return pulumi.get(self, "default_scopes")

    @default_scopes.setter
    def default_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "default_scopes", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The realm this client and scopes exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "realm_id", value)


class RealmDefaultClientScopes(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Allows you to manage the set of default client scopes for a Keycloak realm, which are used when new clients are created.

        Note that this resource attempts to be an **authoritative** source over the default client scopes for a Keycloak realm,
        so any Keycloak defaults and manual adjustments will be overwritten.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("client_scope",
            realm_id=realm.id,
            name="test-client-scope")
        default_scopes = keycloak.RealmDefaultClientScopes("default_scopes",
            realm_id=realm.id,
            default_scopes=[
                "profile",
                "email",
                "roles",
                "web-origins",
                client_scope.name,
            ])
        ```

        ## Import

        This resource does not support import. Instead of importing, feel free to create this resource

        as if it did not already exist on the server.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] default_scopes: An array of default client scope names that should be used when creating new Keycloak clients.
        :param pulumi.Input[builtins.str] realm_id: The realm this client and scopes exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RealmDefaultClientScopesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows you to manage the set of default client scopes for a Keycloak realm, which are used when new clients are created.

        Note that this resource attempts to be an **authoritative** source over the default client scopes for a Keycloak realm,
        so any Keycloak defaults and manual adjustments will be overwritten.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("client_scope",
            realm_id=realm.id,
            name="test-client-scope")
        default_scopes = keycloak.RealmDefaultClientScopes("default_scopes",
            realm_id=realm.id,
            default_scopes=[
                "profile",
                "email",
                "roles",
                "web-origins",
                client_scope.name,
            ])
        ```

        ## Import

        This resource does not support import. Instead of importing, feel free to create this resource

        as if it did not already exist on the server.

        :param str resource_name: The name of the resource.
        :param RealmDefaultClientScopesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RealmDefaultClientScopesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RealmDefaultClientScopesArgs.__new__(RealmDefaultClientScopesArgs)

            if default_scopes is None and not opts.urn:
                raise TypeError("Missing required property 'default_scopes'")
            __props__.__dict__["default_scopes"] = default_scopes
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(RealmDefaultClientScopes, __self__).__init__(
            'keycloak:index/realmDefaultClientScopes:RealmDefaultClientScopes',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            realm_id: Optional[pulumi.Input[builtins.str]] = None) -> 'RealmDefaultClientScopes':
        """
        Get an existing RealmDefaultClientScopes resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] default_scopes: An array of default client scope names that should be used when creating new Keycloak clients.
        :param pulumi.Input[builtins.str] realm_id: The realm this client and scopes exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RealmDefaultClientScopesState.__new__(_RealmDefaultClientScopesState)

        __props__.__dict__["default_scopes"] = default_scopes
        __props__.__dict__["realm_id"] = realm_id
        return RealmDefaultClientScopes(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultScopes")
    def default_scopes(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        An array of default client scope names that should be used when creating new Keycloak clients.
        """
        return pulumi.get(self, "default_scopes")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[builtins.str]:
        """
        The realm this client and scopes exists in.
        """
        return pulumi.get(self, "realm_id")

