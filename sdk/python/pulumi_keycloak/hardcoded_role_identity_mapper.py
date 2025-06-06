# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
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

__all__ = ['HardcodedRoleIdentityMapperArgs', 'HardcodedRoleIdentityMapper']

@pulumi.input_type
class HardcodedRoleIdentityMapperArgs:
    def __init__(__self__, *,
                 identity_provider_alias: pulumi.Input[builtins.str],
                 realm: pulumi.Input[builtins.str],
                 extra_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a HardcodedRoleIdentityMapper resource.
        :param pulumi.Input[builtins.str] identity_provider_alias: The IDP alias of the attribute to set.
        :param pulumi.Input[builtins.str] realm: The realm ID that this mapper will exist in.
        :param pulumi.Input[builtins.str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[builtins.str] role: The name of the role which should be assigned to the users.
        """
        pulumi.set(__self__, "identity_provider_alias", identity_provider_alias)
        pulumi.set(__self__, "realm", realm)
        if extra_config is not None:
            pulumi.set(__self__, "extra_config", extra_config)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> pulumi.Input[builtins.str]:
        """
        The IDP alias of the attribute to set.
        """
        return pulumi.get(self, "identity_provider_alias")

    @identity_provider_alias.setter
    def identity_provider_alias(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "identity_provider_alias", value)

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Input[builtins.str]:
        """
        The realm ID that this mapper will exist in.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "extra_config")

    @extra_config.setter
    def extra_config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "extra_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Display name of this mapper when displayed in the console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the role which should be assigned to the users.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class _HardcodedRoleIdentityMapperState:
    def __init__(__self__, *,
                 extra_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 realm: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering HardcodedRoleIdentityMapper resources.
        :param pulumi.Input[builtins.str] identity_provider_alias: The IDP alias of the attribute to set.
        :param pulumi.Input[builtins.str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[builtins.str] realm: The realm ID that this mapper will exist in.
        :param pulumi.Input[builtins.str] role: The name of the role which should be assigned to the users.
        """
        if extra_config is not None:
            pulumi.set(__self__, "extra_config", extra_config)
        if identity_provider_alias is not None:
            pulumi.set(__self__, "identity_provider_alias", identity_provider_alias)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "extra_config")

    @extra_config.setter
    def extra_config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "extra_config", value)

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The IDP alias of the attribute to set.
        """
        return pulumi.get(self, "identity_provider_alias")

    @identity_provider_alias.setter
    def identity_provider_alias(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "identity_provider_alias", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Display name of this mapper when displayed in the console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The realm ID that this mapper will exist in.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The name of the role which should be assigned to the users.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role", value)


@pulumi.type_token("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper")
class HardcodedRoleIdentityMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extra_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 realm: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Allows for creating and managing hardcoded role mappers for Keycloak identity provider.

        The identity provider hardcoded role mapper grants a specified Keycloak role to each Keycloak user from the LDAP provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        oidc = keycloak.oidc.IdentityProvider("oidc",
            realm=realm.id,
            alias="my-idp",
            authorization_url="https://authorizationurl.com",
            client_id="clientID",
            client_secret="clientSecret",
            token_url="https://tokenurl.com")
        realm_role = keycloak.Role("realm_role",
            realm_id=realm.id,
            name="my-realm-role",
            description="My Realm Role")
        oidc_hardcoded_role_identity_mapper = keycloak.HardcodedRoleIdentityMapper("oidc",
            realm=realm.id,
            name="hardcodedRole",
            identity_provider_alias=oidc.alias,
            role="my-realm-role",
            extra_config={
                "syncMode": "INHERIT",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] identity_provider_alias: The IDP alias of the attribute to set.
        :param pulumi.Input[builtins.str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[builtins.str] realm: The realm ID that this mapper will exist in.
        :param pulumi.Input[builtins.str] role: The name of the role which should be assigned to the users.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HardcodedRoleIdentityMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing hardcoded role mappers for Keycloak identity provider.

        The identity provider hardcoded role mapper grants a specified Keycloak role to each Keycloak user from the LDAP provider.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        oidc = keycloak.oidc.IdentityProvider("oidc",
            realm=realm.id,
            alias="my-idp",
            authorization_url="https://authorizationurl.com",
            client_id="clientID",
            client_secret="clientSecret",
            token_url="https://tokenurl.com")
        realm_role = keycloak.Role("realm_role",
            realm_id=realm.id,
            name="my-realm-role",
            description="My Realm Role")
        oidc_hardcoded_role_identity_mapper = keycloak.HardcodedRoleIdentityMapper("oidc",
            realm=realm.id,
            name="hardcodedRole",
            identity_provider_alias=oidc.alias,
            role="my-realm-role",
            extra_config={
                "syncMode": "INHERIT",
            })
        ```

        :param str resource_name: The name of the resource.
        :param HardcodedRoleIdentityMapperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HardcodedRoleIdentityMapperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extra_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 realm: Optional[pulumi.Input[builtins.str]] = None,
                 role: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HardcodedRoleIdentityMapperArgs.__new__(HardcodedRoleIdentityMapperArgs)

            __props__.__dict__["extra_config"] = extra_config
            if identity_provider_alias is None and not opts.urn:
                raise TypeError("Missing required property 'identity_provider_alias'")
            __props__.__dict__["identity_provider_alias"] = identity_provider_alias
            __props__.__dict__["name"] = name
            if realm is None and not opts.urn:
                raise TypeError("Missing required property 'realm'")
            __props__.__dict__["realm"] = realm
            __props__.__dict__["role"] = role
        super(HardcodedRoleIdentityMapper, __self__).__init__(
            'keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            extra_config: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            identity_provider_alias: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            realm: Optional[pulumi.Input[builtins.str]] = None,
            role: Optional[pulumi.Input[builtins.str]] = None) -> 'HardcodedRoleIdentityMapper':
        """
        Get an existing HardcodedRoleIdentityMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] identity_provider_alias: The IDP alias of the attribute to set.
        :param pulumi.Input[builtins.str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[builtins.str] realm: The realm ID that this mapper will exist in.
        :param pulumi.Input[builtins.str] role: The name of the role which should be assigned to the users.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HardcodedRoleIdentityMapperState.__new__(_HardcodedRoleIdentityMapperState)

        __props__.__dict__["extra_config"] = extra_config
        __props__.__dict__["identity_provider_alias"] = identity_provider_alias
        __props__.__dict__["name"] = name
        __props__.__dict__["realm"] = realm
        __props__.__dict__["role"] = role
        return HardcodedRoleIdentityMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        return pulumi.get(self, "extra_config")

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> pulumi.Output[builtins.str]:
        """
        The IDP alias of the attribute to set.
        """
        return pulumi.get(self, "identity_provider_alias")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Display name of this mapper when displayed in the console.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[builtins.str]:
        """
        The realm ID that this mapper will exist in.
        """
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The name of the role which should be assigned to the users.
        """
        return pulumi.get(self, "role")

