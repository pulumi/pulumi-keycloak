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

__all__ = ['ClientUserPolicyArgs', 'ClientUserPolicy']

@pulumi.input_type
class ClientUserPolicyArgs:
    def __init__(__self__, *,
                 decision_strategy: pulumi.Input[builtins.str],
                 realm_id: pulumi.Input[builtins.str],
                 resource_server_id: pulumi.Input[builtins.str],
                 users: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 logic: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a ClientUserPolicy resource.
        """
        pulumi.set(__self__, "decision_strategy", decision_strategy)
        pulumi.set(__self__, "realm_id", realm_id)
        pulumi.set(__self__, "resource_server_id", resource_server_id)
        pulumi.set(__self__, "users", users)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if logic is not None:
            pulumi.set(__self__, "logic", logic)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> pulumi.Input[builtins.str]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "decision_strategy", value)

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
    @pulumi.getter
    def users(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "users", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def logic(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "logic")

    @logic.setter
    def logic(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "logic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ClientUserPolicyState:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 logic: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering ClientUserPolicy resources.
        """
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if logic is not None:
            pulumi.set(__self__, "logic", logic)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if resource_server_id is not None:
            pulumi.set(__self__, "resource_server_id", resource_server_id)
        if users is not None:
            pulumi.set(__self__, "users", users)

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
    def logic(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "logic")

    @logic.setter
    def logic(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "logic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

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
    @pulumi.getter
    def users(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        return pulumi.get(self, "users")

    @users.setter
    def users(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "users", value)


@pulumi.type_token("keycloak:openid/clientUserPolicy:ClientUserPolicy")
class ClientUserPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 logic: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Create a ClientUserPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClientUserPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ClientUserPolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ClientUserPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientUserPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
                 description: Optional[pulumi.Input[builtins.str]] = None,
                 logic: Optional[pulumi.Input[builtins.str]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 resource_server_id: Optional[pulumi.Input[builtins.str]] = None,
                 users: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientUserPolicyArgs.__new__(ClientUserPolicyArgs)

            if decision_strategy is None and not opts.urn:
                raise TypeError("Missing required property 'decision_strategy'")
            __props__.__dict__["decision_strategy"] = decision_strategy
            __props__.__dict__["description"] = description
            __props__.__dict__["logic"] = logic
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            if resource_server_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_server_id'")
            __props__.__dict__["resource_server_id"] = resource_server_id
            if users is None and not opts.urn:
                raise TypeError("Missing required property 'users'")
            __props__.__dict__["users"] = users
        super(ClientUserPolicy, __self__).__init__(
            'keycloak:openid/clientUserPolicy:ClientUserPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            decision_strategy: Optional[pulumi.Input[builtins.str]] = None,
            description: Optional[pulumi.Input[builtins.str]] = None,
            logic: Optional[pulumi.Input[builtins.str]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            realm_id: Optional[pulumi.Input[builtins.str]] = None,
            resource_server_id: Optional[pulumi.Input[builtins.str]] = None,
            users: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'ClientUserPolicy':
        """
        Get an existing ClientUserPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientUserPolicyState.__new__(_ClientUserPolicyState)

        __props__.__dict__["decision_strategy"] = decision_strategy
        __props__.__dict__["description"] = description
        __props__.__dict__["logic"] = logic
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["resource_server_id"] = resource_server_id
        __props__.__dict__["users"] = users
        return ClientUserPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def logic(self) -> pulumi.Output[Optional[builtins.str]]:
        return pulumi.get(self, "logic")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> pulumi.Output[builtins.str]:
        return pulumi.get(self, "resource_server_id")

    @property
    @pulumi.getter
    def users(self) -> pulumi.Output[Sequence[builtins.str]]:
        return pulumi.get(self, "users")

