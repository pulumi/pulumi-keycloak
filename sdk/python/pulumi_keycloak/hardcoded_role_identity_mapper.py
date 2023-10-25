# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['HardcodedRoleIdentityMapperArgs', 'HardcodedRoleIdentityMapper']

@pulumi.input_type
class HardcodedRoleIdentityMapperArgs:
    def __init__(__self__, *,
                 identity_provider_alias: pulumi.Input[str],
                 realm: pulumi.Input[str],
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HardcodedRoleIdentityMapper resource.
        :param pulumi.Input[str] identity_provider_alias: IDP Alias
        :param pulumi.Input[str] realm: Realm Name
        :param pulumi.Input[str] name: IDP Mapper Name
        :param pulumi.Input[str] role: Role Name
        """
        HardcodedRoleIdentityMapperArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            identity_provider_alias=identity_provider_alias,
            realm=realm,
            extra_config=extra_config,
            name=name,
            role=role,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             identity_provider_alias: Optional[pulumi.Input[str]] = None,
             realm: Optional[pulumi.Input[str]] = None,
             extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             name: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if identity_provider_alias is None and 'identityProviderAlias' in kwargs:
            identity_provider_alias = kwargs['identityProviderAlias']
        if identity_provider_alias is None:
            raise TypeError("Missing 'identity_provider_alias' argument")
        if realm is None:
            raise TypeError("Missing 'realm' argument")
        if extra_config is None and 'extraConfig' in kwargs:
            extra_config = kwargs['extraConfig']

        _setter("identity_provider_alias", identity_provider_alias)
        _setter("realm", realm)
        if extra_config is not None:
            _setter("extra_config", extra_config)
        if name is not None:
            _setter("name", name)
        if role is not None:
            _setter("role", role)

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> pulumi.Input[str]:
        """
        IDP Alias
        """
        return pulumi.get(self, "identity_provider_alias")

    @identity_provider_alias.setter
    def identity_provider_alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_provider_alias", value)

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Input[str]:
        """
        Realm Name
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "extra_config")

    @extra_config.setter
    def extra_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IDP Mapper Name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Role Name
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class _HardcodedRoleIdentityMapperState:
    def __init__(__self__, *,
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HardcodedRoleIdentityMapper resources.
        :param pulumi.Input[str] identity_provider_alias: IDP Alias
        :param pulumi.Input[str] name: IDP Mapper Name
        :param pulumi.Input[str] realm: Realm Name
        :param pulumi.Input[str] role: Role Name
        """
        _HardcodedRoleIdentityMapperState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            extra_config=extra_config,
            identity_provider_alias=identity_provider_alias,
            name=name,
            realm=realm,
            role=role,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             identity_provider_alias: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             realm: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if extra_config is None and 'extraConfig' in kwargs:
            extra_config = kwargs['extraConfig']
        if identity_provider_alias is None and 'identityProviderAlias' in kwargs:
            identity_provider_alias = kwargs['identityProviderAlias']

        if extra_config is not None:
            _setter("extra_config", extra_config)
        if identity_provider_alias is not None:
            _setter("identity_provider_alias", identity_provider_alias)
        if name is not None:
            _setter("name", name)
        if realm is not None:
            _setter("realm", realm)
        if role is not None:
            _setter("role", role)

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "extra_config")

    @extra_config.setter
    def extra_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra_config", value)

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> Optional[pulumi.Input[str]]:
        """
        IDP Alias
        """
        return pulumi.get(self, "identity_provider_alias")

    @identity_provider_alias.setter
    def identity_provider_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_provider_alias", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        IDP Mapper Name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        """
        Realm Name
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Role Name
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class HardcodedRoleIdentityMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a HardcodedRoleIdentityMapper resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_provider_alias: IDP Alias
        :param pulumi.Input[str] name: IDP Mapper Name
        :param pulumi.Input[str] realm: Realm Name
        :param pulumi.Input[str] role: Role Name
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HardcodedRoleIdentityMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a HardcodedRoleIdentityMapper resource with the given unique name, props, and options.
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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            HardcodedRoleIdentityMapperArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
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
            extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            identity_provider_alias: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'HardcodedRoleIdentityMapper':
        """
        Get an existing HardcodedRoleIdentityMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] identity_provider_alias: IDP Alias
        :param pulumi.Input[str] name: IDP Mapper Name
        :param pulumi.Input[str] realm: Realm Name
        :param pulumi.Input[str] role: Role Name
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
    def extra_config(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "extra_config")

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> pulumi.Output[str]:
        """
        IDP Alias
        """
        return pulumi.get(self, "identity_provider_alias")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        IDP Mapper Name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[str]:
        """
        Realm Name
        """
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[Optional[str]]:
        """
        Role Name
        """
        return pulumi.get(self, "role")

