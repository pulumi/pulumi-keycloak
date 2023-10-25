# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClientTimePolicyArgs', 'ClientTimePolicy']

@pulumi.input_type
class ClientTimePolicyArgs:
    def __init__(__self__, *,
                 decision_strategy: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 resource_server_id: pulumi.Input[str],
                 day_month: Optional[pulumi.Input[str]] = None,
                 day_month_end: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hour: Optional[pulumi.Input[str]] = None,
                 hour_end: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None,
                 minute: Optional[pulumi.Input[str]] = None,
                 minute_end: Optional[pulumi.Input[str]] = None,
                 month: Optional[pulumi.Input[str]] = None,
                 month_end: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before: Optional[pulumi.Input[str]] = None,
                 not_on_or_after: Optional[pulumi.Input[str]] = None,
                 year: Optional[pulumi.Input[str]] = None,
                 year_end: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ClientTimePolicy resource.
        """
        ClientTimePolicyArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            decision_strategy=decision_strategy,
            realm_id=realm_id,
            resource_server_id=resource_server_id,
            day_month=day_month,
            day_month_end=day_month_end,
            description=description,
            hour=hour,
            hour_end=hour_end,
            logic=logic,
            minute=minute,
            minute_end=minute_end,
            month=month,
            month_end=month_end,
            name=name,
            not_before=not_before,
            not_on_or_after=not_on_or_after,
            year=year,
            year_end=year_end,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             decision_strategy: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             resource_server_id: Optional[pulumi.Input[str]] = None,
             day_month: Optional[pulumi.Input[str]] = None,
             day_month_end: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             hour: Optional[pulumi.Input[str]] = None,
             hour_end: Optional[pulumi.Input[str]] = None,
             logic: Optional[pulumi.Input[str]] = None,
             minute: Optional[pulumi.Input[str]] = None,
             minute_end: Optional[pulumi.Input[str]] = None,
             month: Optional[pulumi.Input[str]] = None,
             month_end: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             not_before: Optional[pulumi.Input[str]] = None,
             not_on_or_after: Optional[pulumi.Input[str]] = None,
             year: Optional[pulumi.Input[str]] = None,
             year_end: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if decision_strategy is None and 'decisionStrategy' in kwargs:
            decision_strategy = kwargs['decisionStrategy']
        if decision_strategy is None:
            raise TypeError("Missing 'decision_strategy' argument")
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if realm_id is None:
            raise TypeError("Missing 'realm_id' argument")
        if resource_server_id is None and 'resourceServerId' in kwargs:
            resource_server_id = kwargs['resourceServerId']
        if resource_server_id is None:
            raise TypeError("Missing 'resource_server_id' argument")
        if day_month is None and 'dayMonth' in kwargs:
            day_month = kwargs['dayMonth']
        if day_month_end is None and 'dayMonthEnd' in kwargs:
            day_month_end = kwargs['dayMonthEnd']
        if hour_end is None and 'hourEnd' in kwargs:
            hour_end = kwargs['hourEnd']
        if minute_end is None and 'minuteEnd' in kwargs:
            minute_end = kwargs['minuteEnd']
        if month_end is None and 'monthEnd' in kwargs:
            month_end = kwargs['monthEnd']
        if not_before is None and 'notBefore' in kwargs:
            not_before = kwargs['notBefore']
        if not_on_or_after is None and 'notOnOrAfter' in kwargs:
            not_on_or_after = kwargs['notOnOrAfter']
        if year_end is None and 'yearEnd' in kwargs:
            year_end = kwargs['yearEnd']

        _setter("decision_strategy", decision_strategy)
        _setter("realm_id", realm_id)
        _setter("resource_server_id", resource_server_id)
        if day_month is not None:
            _setter("day_month", day_month)
        if day_month_end is not None:
            _setter("day_month_end", day_month_end)
        if description is not None:
            _setter("description", description)
        if hour is not None:
            _setter("hour", hour)
        if hour_end is not None:
            _setter("hour_end", hour_end)
        if logic is not None:
            _setter("logic", logic)
        if minute is not None:
            _setter("minute", minute)
        if minute_end is not None:
            _setter("minute_end", minute_end)
        if month is not None:
            _setter("month", month)
        if month_end is not None:
            _setter("month_end", month_end)
        if name is not None:
            _setter("name", name)
        if not_before is not None:
            _setter("not_before", not_before)
        if not_on_or_after is not None:
            _setter("not_on_or_after", not_on_or_after)
        if year is not None:
            _setter("year", year)
        if year_end is not None:
            _setter("year_end", year_end)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> pulumi.Input[str]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: pulumi.Input[str]):
        pulumi.set(self, "decision_strategy", value)

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
    @pulumi.getter(name="dayMonth")
    def day_month(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "day_month")

    @day_month.setter
    def day_month(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_month", value)

    @property
    @pulumi.getter(name="dayMonthEnd")
    def day_month_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "day_month_end")

    @day_month_end.setter
    def day_month_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_month_end", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def hour(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hour")

    @hour.setter
    def hour(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hour", value)

    @property
    @pulumi.getter(name="hourEnd")
    def hour_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hour_end")

    @hour_end.setter
    def hour_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hour_end", value)

    @property
    @pulumi.getter
    def logic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "logic")

    @logic.setter
    def logic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logic", value)

    @property
    @pulumi.getter
    def minute(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "minute")

    @minute.setter
    def minute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minute", value)

    @property
    @pulumi.getter(name="minuteEnd")
    def minute_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "minute_end")

    @minute_end.setter
    def minute_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minute_end", value)

    @property
    @pulumi.getter
    def month(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "month")

    @month.setter
    def month(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "month", value)

    @property
    @pulumi.getter(name="monthEnd")
    def month_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "month_end")

    @month_end.setter
    def month_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "month_end", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "not_before")

    @not_before.setter
    def not_before(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "not_before", value)

    @property
    @pulumi.getter(name="notOnOrAfter")
    def not_on_or_after(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "not_on_or_after")

    @not_on_or_after.setter
    def not_on_or_after(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "not_on_or_after", value)

    @property
    @pulumi.getter
    def year(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "year")

    @year.setter
    def year(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "year", value)

    @property
    @pulumi.getter(name="yearEnd")
    def year_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "year_end")

    @year_end.setter
    def year_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "year_end", value)


@pulumi.input_type
class _ClientTimePolicyState:
    def __init__(__self__, *,
                 day_month: Optional[pulumi.Input[str]] = None,
                 day_month_end: Optional[pulumi.Input[str]] = None,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hour: Optional[pulumi.Input[str]] = None,
                 hour_end: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None,
                 minute: Optional[pulumi.Input[str]] = None,
                 minute_end: Optional[pulumi.Input[str]] = None,
                 month: Optional[pulumi.Input[str]] = None,
                 month_end: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before: Optional[pulumi.Input[str]] = None,
                 not_on_or_after: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 resource_server_id: Optional[pulumi.Input[str]] = None,
                 year: Optional[pulumi.Input[str]] = None,
                 year_end: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClientTimePolicy resources.
        """
        _ClientTimePolicyState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            day_month=day_month,
            day_month_end=day_month_end,
            decision_strategy=decision_strategy,
            description=description,
            hour=hour,
            hour_end=hour_end,
            logic=logic,
            minute=minute,
            minute_end=minute_end,
            month=month,
            month_end=month_end,
            name=name,
            not_before=not_before,
            not_on_or_after=not_on_or_after,
            realm_id=realm_id,
            resource_server_id=resource_server_id,
            year=year,
            year_end=year_end,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             day_month: Optional[pulumi.Input[str]] = None,
             day_month_end: Optional[pulumi.Input[str]] = None,
             decision_strategy: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             hour: Optional[pulumi.Input[str]] = None,
             hour_end: Optional[pulumi.Input[str]] = None,
             logic: Optional[pulumi.Input[str]] = None,
             minute: Optional[pulumi.Input[str]] = None,
             minute_end: Optional[pulumi.Input[str]] = None,
             month: Optional[pulumi.Input[str]] = None,
             month_end: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             not_before: Optional[pulumi.Input[str]] = None,
             not_on_or_after: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             resource_server_id: Optional[pulumi.Input[str]] = None,
             year: Optional[pulumi.Input[str]] = None,
             year_end: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if day_month is None and 'dayMonth' in kwargs:
            day_month = kwargs['dayMonth']
        if day_month_end is None and 'dayMonthEnd' in kwargs:
            day_month_end = kwargs['dayMonthEnd']
        if decision_strategy is None and 'decisionStrategy' in kwargs:
            decision_strategy = kwargs['decisionStrategy']
        if hour_end is None and 'hourEnd' in kwargs:
            hour_end = kwargs['hourEnd']
        if minute_end is None and 'minuteEnd' in kwargs:
            minute_end = kwargs['minuteEnd']
        if month_end is None and 'monthEnd' in kwargs:
            month_end = kwargs['monthEnd']
        if not_before is None and 'notBefore' in kwargs:
            not_before = kwargs['notBefore']
        if not_on_or_after is None and 'notOnOrAfter' in kwargs:
            not_on_or_after = kwargs['notOnOrAfter']
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if resource_server_id is None and 'resourceServerId' in kwargs:
            resource_server_id = kwargs['resourceServerId']
        if year_end is None and 'yearEnd' in kwargs:
            year_end = kwargs['yearEnd']

        if day_month is not None:
            _setter("day_month", day_month)
        if day_month_end is not None:
            _setter("day_month_end", day_month_end)
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if description is not None:
            _setter("description", description)
        if hour is not None:
            _setter("hour", hour)
        if hour_end is not None:
            _setter("hour_end", hour_end)
        if logic is not None:
            _setter("logic", logic)
        if minute is not None:
            _setter("minute", minute)
        if minute_end is not None:
            _setter("minute_end", minute_end)
        if month is not None:
            _setter("month", month)
        if month_end is not None:
            _setter("month_end", month_end)
        if name is not None:
            _setter("name", name)
        if not_before is not None:
            _setter("not_before", not_before)
        if not_on_or_after is not None:
            _setter("not_on_or_after", not_on_or_after)
        if realm_id is not None:
            _setter("realm_id", realm_id)
        if resource_server_id is not None:
            _setter("resource_server_id", resource_server_id)
        if year is not None:
            _setter("year", year)
        if year_end is not None:
            _setter("year_end", year_end)

    @property
    @pulumi.getter(name="dayMonth")
    def day_month(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "day_month")

    @day_month.setter
    def day_month(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_month", value)

    @property
    @pulumi.getter(name="dayMonthEnd")
    def day_month_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "day_month_end")

    @day_month_end.setter
    def day_month_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "day_month_end", value)

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
    def hour(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hour")

    @hour.setter
    def hour(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hour", value)

    @property
    @pulumi.getter(name="hourEnd")
    def hour_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "hour_end")

    @hour_end.setter
    def hour_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "hour_end", value)

    @property
    @pulumi.getter
    def logic(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "logic")

    @logic.setter
    def logic(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "logic", value)

    @property
    @pulumi.getter
    def minute(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "minute")

    @minute.setter
    def minute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minute", value)

    @property
    @pulumi.getter(name="minuteEnd")
    def minute_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "minute_end")

    @minute_end.setter
    def minute_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "minute_end", value)

    @property
    @pulumi.getter
    def month(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "month")

    @month.setter
    def month(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "month", value)

    @property
    @pulumi.getter(name="monthEnd")
    def month_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "month_end")

    @month_end.setter
    def month_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "month_end", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "not_before")

    @not_before.setter
    def not_before(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "not_before", value)

    @property
    @pulumi.getter(name="notOnOrAfter")
    def not_on_or_after(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "not_on_or_after")

    @not_on_or_after.setter
    def not_on_or_after(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "not_on_or_after", value)

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

    @property
    @pulumi.getter
    def year(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "year")

    @year.setter
    def year(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "year", value)

    @property
    @pulumi.getter(name="yearEnd")
    def year_end(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "year_end")

    @year_end.setter
    def year_end(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "year_end", value)


class ClientTimePolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 day_month: Optional[pulumi.Input[str]] = None,
                 day_month_end: Optional[pulumi.Input[str]] = None,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hour: Optional[pulumi.Input[str]] = None,
                 hour_end: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None,
                 minute: Optional[pulumi.Input[str]] = None,
                 minute_end: Optional[pulumi.Input[str]] = None,
                 month: Optional[pulumi.Input[str]] = None,
                 month_end: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before: Optional[pulumi.Input[str]] = None,
                 not_on_or_after: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 resource_server_id: Optional[pulumi.Input[str]] = None,
                 year: Optional[pulumi.Input[str]] = None,
                 year_end: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ClientTimePolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClientTimePolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ClientTimePolicy resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ClientTimePolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientTimePolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ClientTimePolicyArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 day_month: Optional[pulumi.Input[str]] = None,
                 day_month_end: Optional[pulumi.Input[str]] = None,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 hour: Optional[pulumi.Input[str]] = None,
                 hour_end: Optional[pulumi.Input[str]] = None,
                 logic: Optional[pulumi.Input[str]] = None,
                 minute: Optional[pulumi.Input[str]] = None,
                 minute_end: Optional[pulumi.Input[str]] = None,
                 month: Optional[pulumi.Input[str]] = None,
                 month_end: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 not_before: Optional[pulumi.Input[str]] = None,
                 not_on_or_after: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 resource_server_id: Optional[pulumi.Input[str]] = None,
                 year: Optional[pulumi.Input[str]] = None,
                 year_end: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientTimePolicyArgs.__new__(ClientTimePolicyArgs)

            __props__.__dict__["day_month"] = day_month
            __props__.__dict__["day_month_end"] = day_month_end
            if decision_strategy is None and not opts.urn:
                raise TypeError("Missing required property 'decision_strategy'")
            __props__.__dict__["decision_strategy"] = decision_strategy
            __props__.__dict__["description"] = description
            __props__.__dict__["hour"] = hour
            __props__.__dict__["hour_end"] = hour_end
            __props__.__dict__["logic"] = logic
            __props__.__dict__["minute"] = minute
            __props__.__dict__["minute_end"] = minute_end
            __props__.__dict__["month"] = month
            __props__.__dict__["month_end"] = month_end
            __props__.__dict__["name"] = name
            __props__.__dict__["not_before"] = not_before
            __props__.__dict__["not_on_or_after"] = not_on_or_after
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            if resource_server_id is None and not opts.urn:
                raise TypeError("Missing required property 'resource_server_id'")
            __props__.__dict__["resource_server_id"] = resource_server_id
            __props__.__dict__["year"] = year
            __props__.__dict__["year_end"] = year_end
        super(ClientTimePolicy, __self__).__init__(
            'keycloak:openid/clientTimePolicy:ClientTimePolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            day_month: Optional[pulumi.Input[str]] = None,
            day_month_end: Optional[pulumi.Input[str]] = None,
            decision_strategy: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            hour: Optional[pulumi.Input[str]] = None,
            hour_end: Optional[pulumi.Input[str]] = None,
            logic: Optional[pulumi.Input[str]] = None,
            minute: Optional[pulumi.Input[str]] = None,
            minute_end: Optional[pulumi.Input[str]] = None,
            month: Optional[pulumi.Input[str]] = None,
            month_end: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            not_before: Optional[pulumi.Input[str]] = None,
            not_on_or_after: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            resource_server_id: Optional[pulumi.Input[str]] = None,
            year: Optional[pulumi.Input[str]] = None,
            year_end: Optional[pulumi.Input[str]] = None) -> 'ClientTimePolicy':
        """
        Get an existing ClientTimePolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientTimePolicyState.__new__(_ClientTimePolicyState)

        __props__.__dict__["day_month"] = day_month
        __props__.__dict__["day_month_end"] = day_month_end
        __props__.__dict__["decision_strategy"] = decision_strategy
        __props__.__dict__["description"] = description
        __props__.__dict__["hour"] = hour
        __props__.__dict__["hour_end"] = hour_end
        __props__.__dict__["logic"] = logic
        __props__.__dict__["minute"] = minute
        __props__.__dict__["minute_end"] = minute_end
        __props__.__dict__["month"] = month
        __props__.__dict__["month_end"] = month_end
        __props__.__dict__["name"] = name
        __props__.__dict__["not_before"] = not_before
        __props__.__dict__["not_on_or_after"] = not_on_or_after
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["resource_server_id"] = resource_server_id
        __props__.__dict__["year"] = year
        __props__.__dict__["year_end"] = year_end
        return ClientTimePolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dayMonth")
    def day_month(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "day_month")

    @property
    @pulumi.getter(name="dayMonthEnd")
    def day_month_end(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "day_month_end")

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
    def hour(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "hour")

    @property
    @pulumi.getter(name="hourEnd")
    def hour_end(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "hour_end")

    @property
    @pulumi.getter
    def logic(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "logic")

    @property
    @pulumi.getter
    def minute(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "minute")

    @property
    @pulumi.getter(name="minuteEnd")
    def minute_end(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "minute_end")

    @property
    @pulumi.getter
    def month(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "month")

    @property
    @pulumi.getter(name="monthEnd")
    def month_end(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "month_end")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "not_before")

    @property
    @pulumi.getter(name="notOnOrAfter")
    def not_on_or_after(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "not_on_or_after")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "resource_server_id")

    @property
    @pulumi.getter
    def year(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "year")

    @property
    @pulumi.getter(name="yearEnd")
    def year_end(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "year_end")

