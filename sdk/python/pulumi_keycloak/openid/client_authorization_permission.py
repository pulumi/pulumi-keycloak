# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['ClientAuthorizationPermission']


class ClientAuthorizationPermission(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 resource_server_id: Optional[pulumi.Input[str]] = None,
                 resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a ClientAuthorizationPermission resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
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

            __props__['decision_strategy'] = decision_strategy
            __props__['description'] = description
            __props__['name'] = name
            __props__['policies'] = policies
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if resource_server_id is None:
                raise TypeError("Missing required property 'resource_server_id'")
            __props__['resource_server_id'] = resource_server_id
            __props__['resources'] = resources
            __props__['scopes'] = scopes
            __props__['type'] = type
        super(ClientAuthorizationPermission, __self__).__init__(
            'keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            decision_strategy: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            resource_server_id: Optional[pulumi.Input[str]] = None,
            resources: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            scopes: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'ClientAuthorizationPermission':
        """
        Get an existing ClientAuthorizationPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["decision_strategy"] = decision_strategy
        __props__["description"] = description
        __props__["name"] = name
        __props__["policies"] = policies
        __props__["realm_id"] = realm_id
        __props__["resource_server_id"] = resource_server_id
        __props__["resources"] = resources
        __props__["scopes"] = scopes
        __props__["type"] = type
        return ClientAuthorizationPermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> Optional[List[str]]:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> str:
        return pulumi.get(self, "resource_server_id")

    @property
    @pulumi.getter
    def resources(self) -> Optional[List[str]]:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def scopes(self) -> Optional[List[str]]:
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def type(self) -> Optional[str]:
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

