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
from . import _utilities

__all__ = ['DefaultGroupsArgs', 'DefaultGroups']

@pulumi.input_type
class DefaultGroupsArgs:
    def __init__(__self__, *,
                 group_ids: pulumi.Input[Sequence[pulumi.Input[builtins.str]]],
                 realm_id: pulumi.Input[builtins.str]):
        """
        The set of arguments for constructing a DefaultGroups resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] group_ids: A set of group ids that should be default groups on the realm referenced by `realm_id`.
        :param pulumi.Input[builtins.str] realm_id: The realm this group exists in.
        """
        pulumi.set(__self__, "group_ids", group_ids)
        pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> pulumi.Input[Sequence[pulumi.Input[builtins.str]]]:
        """
        A set of group ids that should be default groups on the realm referenced by `realm_id`.
        """
        return pulumi.get(self, "group_ids")

    @group_ids.setter
    def group_ids(self, value: pulumi.Input[Sequence[pulumi.Input[builtins.str]]]):
        pulumi.set(self, "group_ids", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[builtins.str]:
        """
        The realm this group exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "realm_id", value)


@pulumi.input_type
class _DefaultGroupsState:
    def __init__(__self__, *,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering DefaultGroups resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] group_ids: A set of group ids that should be default groups on the realm referenced by `realm_id`.
        :param pulumi.Input[builtins.str] realm_id: The realm this group exists in.
        """
        if group_ids is not None:
            pulumi.set(__self__, "group_ids", group_ids)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A set of group ids that should be default groups on the realm referenced by `realm_id`.
        """
        return pulumi.get(self, "group_ids")

    @group_ids.setter
    def group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "group_ids", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The realm this group exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "realm_id", value)


@pulumi.type_token("keycloak:index/defaultGroups:DefaultGroups")
class DefaultGroups(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Allows for managing a realm's default groups.

        > You should not use `DefaultGroups` with a group whose members are managed by `GroupMemberships`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        group = keycloak.Group("group",
            realm_id=realm.id,
            name="my-group")
        default = keycloak.DefaultGroups("default",
            realm_id=realm.id,
            group_ids=[group.id])
        ```

        ## Import

        Default groups can be imported using the format `{{realm_id}}` where `realm_id` is the realm the group exists in.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:index/defaultGroups:DefaultGroups default my-realm
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] group_ids: A set of group ids that should be default groups on the realm referenced by `realm_id`.
        :param pulumi.Input[builtins.str] realm_id: The realm this group exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DefaultGroupsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for managing a realm's default groups.

        > You should not use `DefaultGroups` with a group whose members are managed by `GroupMemberships`.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        group = keycloak.Group("group",
            realm_id=realm.id,
            name="my-group")
        default = keycloak.DefaultGroups("default",
            realm_id=realm.id,
            group_ids=[group.id])
        ```

        ## Import

        Default groups can be imported using the format `{{realm_id}}` where `realm_id` is the realm the group exists in.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:index/defaultGroups:DefaultGroups default my-realm
        ```

        :param str resource_name: The name of the resource.
        :param DefaultGroupsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DefaultGroupsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DefaultGroupsArgs.__new__(DefaultGroupsArgs)

            if group_ids is None and not opts.urn:
                raise TypeError("Missing required property 'group_ids'")
            __props__.__dict__["group_ids"] = group_ids
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(DefaultGroups, __self__).__init__(
            'keycloak:index/defaultGroups:DefaultGroups',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            realm_id: Optional[pulumi.Input[builtins.str]] = None) -> 'DefaultGroups':
        """
        Get an existing DefaultGroups resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] group_ids: A set of group ids that should be default groups on the realm referenced by `realm_id`.
        :param pulumi.Input[builtins.str] realm_id: The realm this group exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _DefaultGroupsState.__new__(_DefaultGroupsState)

        __props__.__dict__["group_ids"] = group_ids
        __props__.__dict__["realm_id"] = realm_id
        return DefaultGroups(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        A set of group ids that should be default groups on the realm referenced by `realm_id`.
        """
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[builtins.str]:
        """
        The realm this group exists in.
        """
        return pulumi.get(self, "realm_id")

