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
from .. import _utilities

__all__ = ['ClientAuthorizationPermissionArgs', 'ClientAuthorizationPermission']

@pulumi.input_type
class ClientAuthorizationPermissionArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[builtins.str],
                 resource_server_id: pulumi.Input[builtins.str],
                 decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ClientAuthorizationPermission resource.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        pulumi.set(__self__, "resource_server_id", resource_server_id)
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "resource_server_id")

    @resource_server_id.setter
    def resource_server_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "resource_server_id", value)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _ClientAuthorizationPermissionState:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering ClientAuthorizationPermission resources.
        """
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if resource_server_id is not None:
            pulumi.set(__self__, "resource_server_id", resource_server_id)
        if resource_type is not None:
            pulumi.set(__self__, "resource_type", resource_type)
        if resources is not None:
            pulumi.set(__self__, "resources", resources)
        if scopes is not None:
            pulumi.set(__self__, "scopes", scopes)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "resource_server_id")

    @resource_server_id.setter
    def resource_server_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_server_id", value)

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "resource_type")

    @resource_type.setter
    def resource_type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "resource_type", value)

    @property
    @pulumi.getter
    def resources(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "resources")

    @resources.setter
    def resources(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "resources", value)

    @property
    @pulumi.getter
    def scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "scopes")

    @scopes.setter
    def scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "scopes", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "type", value)


@pulumi.type_token("keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission")
class ClientAuthorizationPermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Allows you to manage openid Client Authorization Permissions.

        ## Import

        Client authorization permissions can be imported using the format: `{{realmId}}/{{resourceServerId}}/{{permissionId}}`.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission test my-realm/3bd4a686-1062-4b59-97b8-e4e3f10b99da/63b3cde8-987d-4cd9-9306-1955579281d9
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClientAuthorizationPermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows you to manage openid Client Authorization Permissions.

        ## Import

        Client authorization permissions can be imported using the format: `{{realmId}}/{{resourceServerId}}/{{permissionId}}`.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission test my-realm/3bd4a686-1062-4b59-97b8-e4e3f10b99da/63b3cde8-987d-4cd9-9306-1955579281d9
        ```

        :param str resource_name: The name of the resource.
        :param ClientAuthorizationPermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientAuthorizationPermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_type: Optional[pulumi.Input[builtins.str]] = None,
                 resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 type: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientAuthorizationPermissionArgs.__new__(ClientAuthorizationPermissionArgs)

            __props__.__dict__["decision_strategy"] = decision_strategy
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["policies"] = policies
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            if resource_server_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_server_id'")
            __props__.__dict__["resource_server_id"] = resource_server_id
            __props__.__dict__["resource_type"] = resource_type
            __props__.__dict__["resources"] = resources
            __props__.__dict__["scopes"] = scopes
            __props__.__dict__["type"] = type
        super(ClientAuthorizationPermission, __self__).__init__(
            'keycloak:openid/clientAuthorizationPermission:ClientAuthorizationPermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            realm_id: Optional[pulumi.Input[builtins.str]] = None,
            resource_server_id: Optional[pulumi.Input[builtins.str]] = None,
            resource_type: Optional[pulumi.Input[builtins.str]] = None,
            resources: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            scopes: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            type: Optional[pulumi.Input[builtins.str]] = None) -> 'ClientAuthorizationPermission':
        """
        Get an existing ClientAuthorizationPermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientAuthorizationPermissionState.__new__(_ClientAuthorizationPermissionState)

        __props__.__dict__["decision_strategy"] = decision_strategy
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["policies"] = policies
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["resource_server_id"] = resource_server_id
        __props__.__dict__["resource_type"] = resource_type
        __props__.__dict__["resources"] = resources
        __props__.__dict__["scopes"] = scopes
        __props__.__dict__["type"] = type
        return ClientAuthorizationPermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "resource_server_id")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "resource_type")

    @property
    @pulumi.getter
    def resources(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def scopes(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "type")

