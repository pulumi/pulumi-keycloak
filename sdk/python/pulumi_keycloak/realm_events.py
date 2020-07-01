# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class RealmEvents(pulumi.CustomResource):
    admin_events_details_enabled: pulumi.Output[bool]
    admin_events_enabled: pulumi.Output[bool]
    enabled_event_types: pulumi.Output[list]
    events_enabled: pulumi.Output[bool]
    events_expiration: pulumi.Output[float]
    events_listeners: pulumi.Output[list]
    realm_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, admin_events_details_enabled=None, admin_events_enabled=None, enabled_event_types=None, events_enabled=None, events_expiration=None, events_listeners=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # RealmEvents

        Allows for managing Realm Events settings within Keycloak.

        ### Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm", realm="test")
        realm_events = keycloak.RealmEvents("realmEvents",
            admin_events_details_enabled=True,
            admin_events_enabled=True,
            enabled_event_types=[
                "LOGIN",
                "LOGOUT",
            ],
            events_enabled=True,
            events_expiration=3600,
            events_listeners=["jboss-logging"],
            realm_id=realm.id)
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The name of the realm the event settings apply to.
        - `admin_events_enabled` - (Optional) When true, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
        - `admin_events_details_enabled` - (Optional) When true, saved admin events will included detailed information for create/update requests. Defaults to `false`.
        - `events_enabled` - (Optional) When true, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
        - `events_expiration` - (Optional) The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
        - `enabled_event_types` - (Optional) The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
        - `events_listeners` - (Optional) The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.

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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['admin_events_details_enabled'] = admin_events_details_enabled
            __props__['admin_events_enabled'] = admin_events_enabled
            __props__['enabled_event_types'] = enabled_event_types
            __props__['events_enabled'] = events_enabled
            __props__['events_expiration'] = events_expiration
            __props__['events_listeners'] = events_listeners
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(RealmEvents, __self__).__init__(
            'keycloak:index/realmEvents:RealmEvents',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, admin_events_details_enabled=None, admin_events_enabled=None, enabled_event_types=None, events_enabled=None, events_expiration=None, events_listeners=None, realm_id=None):
        """
        Get an existing RealmEvents resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["admin_events_details_enabled"] = admin_events_details_enabled
        __props__["admin_events_enabled"] = admin_events_enabled
        __props__["enabled_event_types"] = enabled_event_types
        __props__["events_enabled"] = events_enabled
        __props__["events_expiration"] = events_expiration
        __props__["events_listeners"] = events_listeners
        __props__["realm_id"] = realm_id
        return RealmEvents(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
