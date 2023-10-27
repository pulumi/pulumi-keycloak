# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['GroupMembershipsArgs', 'GroupMemberships']

@pulumi.input_type
class GroupMembershipsArgs:
    def __init__(__self__, *,
                 members: pulumi.Input[Sequence[pulumi.Input[str]]],
                 realm_id: pulumi.Input[str],
                 group_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a GroupMemberships resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of usernames that belong to this group.
        :param pulumi.Input[str] realm_id: The realm this group exists in.
        :param pulumi.Input[str] group_id: The ID of the group this resource should manage memberships for.
        """
        pulumi.set(__self__, "members", members)
        pulumi.set(__self__, "realm_id", realm_id)
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)

    @property
    @pulumi.getter
    def members(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of usernames that belong to this group.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm this group exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the group this resource should manage memberships for.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)


@pulumi.input_type
class _GroupMembershipsState:
    def __init__(__self__, *,
                 group_id: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering GroupMemberships resources.
        :param pulumi.Input[str] group_id: The ID of the group this resource should manage memberships for.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of usernames that belong to this group.
        :param pulumi.Input[str] realm_id: The realm this group exists in.
        """
        if group_id is not None:
            pulumi.set(__self__, "group_id", group_id)
        if members is not None:
            pulumi.set(__self__, "members", members)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the group this resource should manage memberships for.
        """
        return pulumi.get(self, "group_id")

    @group_id.setter
    def group_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "group_id", value)

    @property
    @pulumi.getter
    def members(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of usernames that belong to this group.
        """
        return pulumi.get(self, "members")

    @members.setter
    def members(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "members", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm this group exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class GroupMemberships(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for managing a Keycloak group's members.

        Note that this resource attempts to be an **authoritative** source over group members. When this resource takes control
        over a group's members, users that are manually added to the group will be removed, and users that are manually removed
        from the group will be added upon the next run of `pulumi up`.

        Also note that you should not use `GroupMemberships` with a group has been assigned as a default group via
        `DefaultGroups`.

        This resource **should not** be used to control membership of a group that has its members federated from an external
        source via group mapping.

        To non-exclusively manage the group's of a user, see the [`UserGroups` resource][1]

        This resource paginates its data loading on refresh by 50 items.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        group = keycloak.Group("group", realm_id=realm.id)
        user = keycloak.User("user",
            realm_id=realm.id,
            username="my-user")
        group_members = keycloak.GroupMemberships("groupMembers",
            realm_id=realm.id,
            group_id=group.id,
            members=[user.username])
        ```

        ## Import

        This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server. [1]providers/mrparkers/keycloak/latest/docs/resources/group_memberships

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: The ID of the group this resource should manage memberships for.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of usernames that belong to this group.
        :param pulumi.Input[str] realm_id: The realm this group exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GroupMembershipsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for managing a Keycloak group's members.

        Note that this resource attempts to be an **authoritative** source over group members. When this resource takes control
        over a group's members, users that are manually added to the group will be removed, and users that are manually removed
        from the group will be added upon the next run of `pulumi up`.

        Also note that you should not use `GroupMemberships` with a group has been assigned as a default group via
        `DefaultGroups`.

        This resource **should not** be used to control membership of a group that has its members federated from an external
        source via group mapping.

        To non-exclusively manage the group's of a user, see the [`UserGroups` resource][1]

        This resource paginates its data loading on refresh by 50 items.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        group = keycloak.Group("group", realm_id=realm.id)
        user = keycloak.User("user",
            realm_id=realm.id,
            username="my-user")
        group_members = keycloak.GroupMemberships("groupMembers",
            realm_id=realm.id,
            group_id=group.id,
            members=[user.username])
        ```

        ## Import

        This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server. [1]providers/mrparkers/keycloak/latest/docs/resources/group_memberships

        :param str resource_name: The name of the resource.
        :param GroupMembershipsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GroupMembershipsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_id: Optional[pulumi.Input[str]] = None,
                 members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GroupMembershipsArgs.__new__(GroupMembershipsArgs)

            __props__.__dict__["group_id"] = group_id
            if members is None and not opts.urn:
                raise TypeError("Missing required property 'members'")
            __props__.__dict__["members"] = members
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(GroupMemberships, __self__).__init__(
            'keycloak:index/groupMemberships:GroupMemberships',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_id: Optional[pulumi.Input[str]] = None,
            members: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'GroupMemberships':
        """
        Get an existing GroupMemberships resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_id: The ID of the group this resource should manage memberships for.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] members: A list of usernames that belong to this group.
        :param pulumi.Input[str] realm_id: The realm this group exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GroupMembershipsState.__new__(_GroupMembershipsState)

        __props__.__dict__["group_id"] = group_id
        __props__.__dict__["members"] = members
        __props__.__dict__["realm_id"] = realm_id
        return GroupMemberships(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupId")
    def group_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the group this resource should manage memberships for.
        """
        return pulumi.get(self, "group_id")

    @property
    @pulumi.getter
    def members(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of usernames that belong to this group.
        """
        return pulumi.get(self, "members")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm this group exists in.
        """
        return pulumi.get(self, "realm_id")

