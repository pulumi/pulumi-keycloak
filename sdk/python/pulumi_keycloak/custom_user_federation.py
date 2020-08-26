# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['CustomUserFederation']


class CustomUserFederation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cache_policy: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parent_id: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[float]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a CustomUserFederation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: When false, this provider will not be used when performing queries for users.
        :param pulumi.Input[str] name: Display name of the provider when displayed in the console.
        :param pulumi.Input[str] parent_id: The parent_id of the generated component. will use realm_id if not specified.
        :param pulumi.Input[float] priority: Priority of this provider when looking up users. Lower values are first.
        :param pulumi.Input[str] provider_id: The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
               interface
        :param pulumi.Input[str] realm_id: The realm (name) this provider will provide user federation for.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['cache_policy'] = cache_policy
            __props__['config'] = config
            __props__['enabled'] = enabled
            __props__['name'] = name
            __props__['parent_id'] = parent_id
            __props__['priority'] = priority
            if provider_id is None:
                raise TypeError("Missing required property 'provider_id'")
            __props__['provider_id'] = provider_id
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(CustomUserFederation, __self__).__init__(
            'keycloak:index/customUserFederation:CustomUserFederation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cache_policy: Optional[pulumi.Input[str]] = None,
            config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            parent_id: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[float]] = None,
            provider_id: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'CustomUserFederation':
        """
        Get an existing CustomUserFederation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] enabled: When false, this provider will not be used when performing queries for users.
        :param pulumi.Input[str] name: Display name of the provider when displayed in the console.
        :param pulumi.Input[str] parent_id: The parent_id of the generated component. will use realm_id if not specified.
        :param pulumi.Input[float] priority: Priority of this provider when looking up users. Lower values are first.
        :param pulumi.Input[str] provider_id: The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
               interface
        :param pulumi.Input[str] realm_id: The realm (name) this provider will provide user federation for.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["cache_policy"] = cache_policy
        __props__["config"] = config
        __props__["enabled"] = enabled
        __props__["name"] = name
        __props__["parent_id"] = parent_id
        __props__["priority"] = priority
        __props__["provider_id"] = provider_id
        __props__["realm_id"] = realm_id
        return CustomUserFederation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cachePolicy")
    def cache_policy(self) -> Optional[str]:
        return pulumi.get(self, "cache_policy")

    @property
    @pulumi.getter
    def config(self) -> Optional[Mapping[str, Any]]:
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def enabled(self) -> Optional[bool]:
        """
        When false, this provider will not be used when performing queries for users.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Display name of the provider when displayed in the console.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> Optional[str]:
        """
        The parent_id of the generated component. will use realm_id if not specified.
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def priority(self) -> Optional[float]:
        """
        Priority of this provider when looking up users. Lower values are first.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> str:
        """
        The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
        interface
        """
        return pulumi.get(self, "provider_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        """
        The realm (name) this provider will provide user federation for.
        """
        return pulumi.get(self, "realm_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

