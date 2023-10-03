# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RealmEventsArgs', 'RealmEvents']

@pulumi.input_type
class RealmEventsArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 admin_events_details_enabled: Optional[pulumi.Input[bool]] = None,
                 admin_events_enabled: Optional[pulumi.Input[bool]] = None,
                 enabled_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 events_enabled: Optional[pulumi.Input[bool]] = None,
                 events_expiration: Optional[pulumi.Input[int]] = None,
                 events_listeners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a RealmEvents resource.
        :param pulumi.Input[str] realm_id: The name of the realm the event settings apply to.
        :param pulumi.Input[bool] admin_events_details_enabled: When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
        :param pulumi.Input[bool] admin_events_enabled: When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_event_types: The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
        :param pulumi.Input[bool] events_enabled: When `true`, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
        :param pulumi.Input[int] events_expiration: The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events_listeners: The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
        """
        RealmEventsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            realm_id=realm_id,
            admin_events_details_enabled=admin_events_details_enabled,
            admin_events_enabled=admin_events_enabled,
            enabled_event_types=enabled_event_types,
            events_enabled=events_enabled,
            events_expiration=events_expiration,
            events_listeners=events_listeners,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             realm_id: pulumi.Input[str],
             admin_events_details_enabled: Optional[pulumi.Input[bool]] = None,
             admin_events_enabled: Optional[pulumi.Input[bool]] = None,
             enabled_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             events_enabled: Optional[pulumi.Input[bool]] = None,
             events_expiration: Optional[pulumi.Input[int]] = None,
             events_listeners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("realm_id", realm_id)
        if admin_events_details_enabled is not None:
            _setter("admin_events_details_enabled", admin_events_details_enabled)
        if admin_events_enabled is not None:
            _setter("admin_events_enabled", admin_events_enabled)
        if enabled_event_types is not None:
            _setter("enabled_event_types", enabled_event_types)
        if events_enabled is not None:
            _setter("events_enabled", events_enabled)
        if events_expiration is not None:
            _setter("events_expiration", events_expiration)
        if events_listeners is not None:
            _setter("events_listeners", events_listeners)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The name of the realm the event settings apply to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="adminEventsDetailsEnabled")
    def admin_events_details_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
        """
        return pulumi.get(self, "admin_events_details_enabled")

    @admin_events_details_enabled.setter
    def admin_events_details_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_events_details_enabled", value)

    @property
    @pulumi.getter(name="adminEventsEnabled")
    def admin_events_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
        """
        return pulumi.get(self, "admin_events_enabled")

    @admin_events_enabled.setter
    def admin_events_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_events_enabled", value)

    @property
    @pulumi.getter(name="enabledEventTypes")
    def enabled_event_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
        """
        return pulumi.get(self, "enabled_event_types")

    @enabled_event_types.setter
    def enabled_event_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "enabled_event_types", value)

    @property
    @pulumi.getter(name="eventsEnabled")
    def events_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
        """
        return pulumi.get(self, "events_enabled")

    @events_enabled.setter
    def events_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "events_enabled", value)

    @property
    @pulumi.getter(name="eventsExpiration")
    def events_expiration(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
        """
        return pulumi.get(self, "events_expiration")

    @events_expiration.setter
    def events_expiration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "events_expiration", value)

    @property
    @pulumi.getter(name="eventsListeners")
    def events_listeners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
        """
        return pulumi.get(self, "events_listeners")

    @events_listeners.setter
    def events_listeners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "events_listeners", value)


@pulumi.input_type
class _RealmEventsState:
    def __init__(__self__, *,
                 admin_events_details_enabled: Optional[pulumi.Input[bool]] = None,
                 admin_events_enabled: Optional[pulumi.Input[bool]] = None,
                 enabled_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 events_enabled: Optional[pulumi.Input[bool]] = None,
                 events_expiration: Optional[pulumi.Input[int]] = None,
                 events_listeners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RealmEvents resources.
        :param pulumi.Input[bool] admin_events_details_enabled: When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
        :param pulumi.Input[bool] admin_events_enabled: When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_event_types: The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
        :param pulumi.Input[bool] events_enabled: When `true`, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
        :param pulumi.Input[int] events_expiration: The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events_listeners: The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
        :param pulumi.Input[str] realm_id: The name of the realm the event settings apply to.
        """
        _RealmEventsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            admin_events_details_enabled=admin_events_details_enabled,
            admin_events_enabled=admin_events_enabled,
            enabled_event_types=enabled_event_types,
            events_enabled=events_enabled,
            events_expiration=events_expiration,
            events_listeners=events_listeners,
            realm_id=realm_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             admin_events_details_enabled: Optional[pulumi.Input[bool]] = None,
             admin_events_enabled: Optional[pulumi.Input[bool]] = None,
             enabled_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             events_enabled: Optional[pulumi.Input[bool]] = None,
             events_expiration: Optional[pulumi.Input[int]] = None,
             events_listeners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if admin_events_details_enabled is not None:
            _setter("admin_events_details_enabled", admin_events_details_enabled)
        if admin_events_enabled is not None:
            _setter("admin_events_enabled", admin_events_enabled)
        if enabled_event_types is not None:
            _setter("enabled_event_types", enabled_event_types)
        if events_enabled is not None:
            _setter("events_enabled", events_enabled)
        if events_expiration is not None:
            _setter("events_expiration", events_expiration)
        if events_listeners is not None:
            _setter("events_listeners", events_listeners)
        if realm_id is not None:
            _setter("realm_id", realm_id)

    @property
    @pulumi.getter(name="adminEventsDetailsEnabled")
    def admin_events_details_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
        """
        return pulumi.get(self, "admin_events_details_enabled")

    @admin_events_details_enabled.setter
    def admin_events_details_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_events_details_enabled", value)

    @property
    @pulumi.getter(name="adminEventsEnabled")
    def admin_events_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
        """
        return pulumi.get(self, "admin_events_enabled")

    @admin_events_enabled.setter
    def admin_events_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "admin_events_enabled", value)

    @property
    @pulumi.getter(name="enabledEventTypes")
    def enabled_event_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
        """
        return pulumi.get(self, "enabled_event_types")

    @enabled_event_types.setter
    def enabled_event_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "enabled_event_types", value)

    @property
    @pulumi.getter(name="eventsEnabled")
    def events_enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
        """
        return pulumi.get(self, "events_enabled")

    @events_enabled.setter
    def events_enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "events_enabled", value)

    @property
    @pulumi.getter(name="eventsExpiration")
    def events_expiration(self) -> Optional[pulumi.Input[int]]:
        """
        The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
        """
        return pulumi.get(self, "events_expiration")

    @events_expiration.setter
    def events_expiration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "events_expiration", value)

    @property
    @pulumi.getter(name="eventsListeners")
    def events_listeners(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
        """
        return pulumi.get(self, "events_listeners")

    @events_listeners.setter
    def events_listeners(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "events_listeners", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the realm the event settings apply to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class RealmEvents(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_events_details_enabled: Optional[pulumi.Input[bool]] = None,
                 admin_events_enabled: Optional[pulumi.Input[bool]] = None,
                 enabled_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 events_enabled: Optional[pulumi.Input[bool]] = None,
                 events_expiration: Optional[pulumi.Input[int]] = None,
                 events_listeners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for managing Realm Events settings within Keycloak.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        realm_events = keycloak.RealmEvents("realmEvents",
            realm_id=realm.id,
            events_enabled=True,
            events_expiration=3600,
            admin_events_enabled=True,
            admin_events_details_enabled=True,
            enabled_event_types=[
                "LOGIN",
                "LOGOUT",
            ],
            events_listeners=["jboss-logging"])
        ```

        ## Import

        This resource currently does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_events_details_enabled: When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
        :param pulumi.Input[bool] admin_events_enabled: When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_event_types: The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
        :param pulumi.Input[bool] events_enabled: When `true`, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
        :param pulumi.Input[int] events_expiration: The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events_listeners: The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
        :param pulumi.Input[str] realm_id: The name of the realm the event settings apply to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RealmEventsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for managing Realm Events settings within Keycloak.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        realm_events = keycloak.RealmEvents("realmEvents",
            realm_id=realm.id,
            events_enabled=True,
            events_expiration=3600,
            admin_events_enabled=True,
            admin_events_details_enabled=True,
            enabled_event_types=[
                "LOGIN",
                "LOGOUT",
            ],
            events_listeners=["jboss-logging"])
        ```

        ## Import

        This resource currently does not support importing.

        :param str resource_name: The name of the resource.
        :param RealmEventsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RealmEventsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RealmEventsArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 admin_events_details_enabled: Optional[pulumi.Input[bool]] = None,
                 admin_events_enabled: Optional[pulumi.Input[bool]] = None,
                 enabled_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 events_enabled: Optional[pulumi.Input[bool]] = None,
                 events_expiration: Optional[pulumi.Input[int]] = None,
                 events_listeners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RealmEventsArgs.__new__(RealmEventsArgs)

            __props__.__dict__["admin_events_details_enabled"] = admin_events_details_enabled
            __props__.__dict__["admin_events_enabled"] = admin_events_enabled
            __props__.__dict__["enabled_event_types"] = enabled_event_types
            __props__.__dict__["events_enabled"] = events_enabled
            __props__.__dict__["events_expiration"] = events_expiration
            __props__.__dict__["events_listeners"] = events_listeners
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(RealmEvents, __self__).__init__(
            'keycloak:index/realmEvents:RealmEvents',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_events_details_enabled: Optional[pulumi.Input[bool]] = None,
            admin_events_enabled: Optional[pulumi.Input[bool]] = None,
            enabled_event_types: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            events_enabled: Optional[pulumi.Input[bool]] = None,
            events_expiration: Optional[pulumi.Input[int]] = None,
            events_listeners: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'RealmEvents':
        """
        Get an existing RealmEvents resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] admin_events_details_enabled: When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
        :param pulumi.Input[bool] admin_events_enabled: When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] enabled_event_types: The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
        :param pulumi.Input[bool] events_enabled: When `true`, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
        :param pulumi.Input[int] events_expiration: The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] events_listeners: The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
        :param pulumi.Input[str] realm_id: The name of the realm the event settings apply to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RealmEventsState.__new__(_RealmEventsState)

        __props__.__dict__["admin_events_details_enabled"] = admin_events_details_enabled
        __props__.__dict__["admin_events_enabled"] = admin_events_enabled
        __props__.__dict__["enabled_event_types"] = enabled_event_types
        __props__.__dict__["events_enabled"] = events_enabled
        __props__.__dict__["events_expiration"] = events_expiration
        __props__.__dict__["events_listeners"] = events_listeners
        __props__.__dict__["realm_id"] = realm_id
        return RealmEvents(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="adminEventsDetailsEnabled")
    def admin_events_details_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, saved admin events will included detailed information for create/update requests. Defaults to `false`.
        """
        return pulumi.get(self, "admin_events_details_enabled")

    @property
    @pulumi.getter(name="adminEventsEnabled")
    def admin_events_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
        """
        return pulumi.get(self, "admin_events_enabled")

    @property
    @pulumi.getter(name="enabledEventTypes")
    def enabled_event_types(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
        """
        return pulumi.get(self, "enabled_event_types")

    @property
    @pulumi.getter(name="eventsEnabled")
    def events_enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
        """
        return pulumi.get(self, "events_enabled")

    @property
    @pulumi.getter(name="eventsExpiration")
    def events_expiration(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
        """
        return pulumi.get(self, "events_expiration")

    @property
    @pulumi.getter(name="eventsListeners")
    def events_listeners(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
        """
        return pulumi.get(self, "events_listeners")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The name of the realm the event settings apply to.
        """
        return pulumi.get(self, "realm_id")

