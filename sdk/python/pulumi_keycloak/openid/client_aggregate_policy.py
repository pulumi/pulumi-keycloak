# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClientAggregatePolicyArgs', 'ClientAggregatePolicy']

@pulumi.input_type
class ClientAggregatePolicyArgs:
    def __init__(__self__, *,
                 decision_strategy: pulumi.Input[str],
                 policies: pulumi.Input[Sequence[pulumi.Input[str]]],
                 realm_id: pulumi.Input[str],
                 resource_server_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ClientAggregatePolicy resource.
        """
        pulumi.set(__self__, "decision_strategy", decision_strategy)
        pulumi.set(__self__, "policies", policies)
        pulumi.set(__self__, "realm_id", realm_id)
        pulumi.set(__self__, "resource_server_id", resource_server_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if logic is not None:
            pulumi.set(__self__, "logic", logic)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> pulumi.Input[str]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: pulumi.Input[str]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "resource_server_id")

    @resource_server_id.setter
    def resource_server_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource_server_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def logic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "logic")

    @logic.setter
    def logic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ClientAggregatePolicyState:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 resource_server_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClientAggregatePolicy resources.
        """
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if logic is not None:
            pulumi.set(__self__, "logic", logic)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if resource_server_id is not None:
            pulumi.set(__self__, "resource_server_id", resource_server_id)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def logic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "logic")

    @logic.setter
    def logic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logic", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "resource_server_id")

    @resource_server_id.setter
    def resource_server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "resource_server_id", value)


class ClientAggregatePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 resource_server_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ClientAggregatePolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClientAggregatePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ClientAggregatePolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ClientAggregatePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientAggregatePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 resource_server_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientAggregatePolicyArgs.__new__(ClientAggregatePolicyArgs)

            if decision_strategy is None and not opts.urn:
                raise TypeError("Missing required property 'decision_strategy'")
            __props__.__dict__["decision_strategy"] = decision_strategy
            __props__.__dict__["description"] = description
            __props__.__dict__["logic"] = logic
            __props__.__dict__["name"] = name
            if policies is None and not opts.urn:
                raise TypeError("Missing required property 'policies'")
            __props__.__dict__["policies"] = policies
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            if resource_server_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_server_id'")
            __props__.__dict__["resource_server_id"] = resource_server_id
        super(ClientAggregatePolicy, __self__).__init__(
            'keycloak:openid/clientAggregatePolicy:ClientAggregatePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            decision_strategy: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            logic: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            resource_server_id: Optional[pulumi.Input[str]] = None) -> 'ClientAggregatePolicy':
        """
        Get an existing ClientAggregatePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientAggregatePolicyState.__new__(_ClientAggregatePolicyState)

        __props__.__dict__["decision_strategy"] = decision_strategy
        __props__.__dict__["description"] = description
        __props__.__dict__["logic"] = logic
        __props__.__dict__["name"] = name
        __props__.__dict__["policies"] = policies
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["resource_server_id"] = resource_server_id
        return ClientAggregatePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> pulumi.Output[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def logic(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "logic")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence[str]]:
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "resource_server_id")

