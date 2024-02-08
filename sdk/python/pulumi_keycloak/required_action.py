# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RequiredActionArgs', 'RequiredAction']

@pulumi.input_type
class RequiredActionArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 default_action: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RequiredAction resource.
        :param pulumi.Input[str] alias: The alias of the action to attach as a required action.
        :param pulumi.Input[str] realm_id: The realm the required action exists in.
        :param pulumi.Input[bool] default_action: When `true`, the required action is set as the default action for new users. Defaults to `false`.
        :param pulumi.Input[bool] enabled: When `false`, the required action is not enabled for new users. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the required action.
        :param pulumi.Input[int] priority: The priority of the required action.
        """
        pulumi.set(__self__, "alias", alias)
        pulumi.set(__self__, "realm_id", realm_id)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        """
        The alias of the action to attach as a required action.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm the required action exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, the required action is set as the default action for new users. Defaults to `false`.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `false`, the required action is not enabled for new users. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the required action.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the required action.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)


@pulumi.input_type
class _RequiredActionState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RequiredAction resources.
        :param pulumi.Input[str] alias: The alias of the action to attach as a required action.
        :param pulumi.Input[bool] default_action: When `true`, the required action is set as the default action for new users. Defaults to `false`.
        :param pulumi.Input[bool] enabled: When `false`, the required action is not enabled for new users. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the required action.
        :param pulumi.Input[int] priority: The priority of the required action.
        :param pulumi.Input[str] realm_id: The realm the required action exists in.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        The alias of the action to attach as a required action.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, the required action is set as the default action for new users. Defaults to `false`.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `false`, the required action is not enabled for new users. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the required action.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of the required action.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm the required action exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class RequiredAction(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating and managing required actions within Keycloak.

        [Required actions](https://www.keycloak.org/docs/latest/server_admin/#con-required-actions_server_administration_guide) specify actions required before the first login of all new users.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        required_action = keycloak.RequiredAction("requiredAction",
            realm_id=realm.realm,
            alias="webauthn-register",
            enabled=True)
        ```

        ## Import

        Authentication executions can be imported using the formats: `{{realm}}/{{alias}}`.

         Example:

         bash

        ```sh
        $ pulumi import keycloak:index/requiredAction:RequiredAction required_action my-realm/my-default-action-alias
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The alias of the action to attach as a required action.
        :param pulumi.Input[bool] default_action: When `true`, the required action is set as the default action for new users. Defaults to `false`.
        :param pulumi.Input[bool] enabled: When `false`, the required action is not enabled for new users. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the required action.
        :param pulumi.Input[int] priority: The priority of the required action.
        :param pulumi.Input[str] realm_id: The realm the required action exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RequiredActionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing required actions within Keycloak.

        [Required actions](https://www.keycloak.org/docs/latest/server_admin/#con-required-actions_server_administration_guide) specify actions required before the first login of all new users.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        required_action = keycloak.RequiredAction("requiredAction",
            realm_id=realm.realm,
            alias="webauthn-register",
            enabled=True)
        ```

        ## Import

        Authentication executions can be imported using the formats: `{{realm}}/{{alias}}`.

         Example:

         bash

        ```sh
        $ pulumi import keycloak:index/requiredAction:RequiredAction required_action my-realm/my-default-action-alias
        ```

        :param str resource_name: The name of the resource.
        :param RequiredActionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RequiredActionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 default_action: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RequiredActionArgs.__new__(RequiredActionArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            __props__.__dict__["default_action"] = default_action
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["priority"] = priority
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(RequiredAction, __self__).__init__(
            'keycloak:index/requiredAction:RequiredAction',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            default_action: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'RequiredAction':
        """
        Get an existing RequiredAction resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The alias of the action to attach as a required action.
        :param pulumi.Input[bool] default_action: When `true`, the required action is set as the default action for new users. Defaults to `false`.
        :param pulumi.Input[bool] enabled: When `false`, the required action is not enabled for new users. Defaults to `false`.
        :param pulumi.Input[str] name: The name of the required action.
        :param pulumi.Input[int] priority: The priority of the required action.
        :param pulumi.Input[str] realm_id: The realm the required action exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RequiredActionState.__new__(_RequiredActionState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["default_action"] = default_action
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["priority"] = priority
        __props__.__dict__["realm_id"] = realm_id
        return RequiredAction(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        The alias of the action to attach as a required action.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, the required action is set as the default action for new users. Defaults to `false`.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When `false`, the required action is not enabled for new users. Defaults to `false`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the required action.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        The priority of the required action.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm the required action exists in.
        """
        return pulumi.get(self, "realm_id")

