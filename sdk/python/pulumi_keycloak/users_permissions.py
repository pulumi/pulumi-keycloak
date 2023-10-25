# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['UsersPermissionsArgs', 'UsersPermissions']

@pulumi.input_type
class UsersPermissionsArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 impersonate_scope: Optional[pulumi.Input['UsersPermissionsImpersonateScopeArgs']] = None,
                 manage_group_membership_scope: Optional[pulumi.Input['UsersPermissionsManageGroupMembershipScopeArgs']] = None,
                 manage_scope: Optional[pulumi.Input['UsersPermissionsManageScopeArgs']] = None,
                 map_roles_scope: Optional[pulumi.Input['UsersPermissionsMapRolesScopeArgs']] = None,
                 user_impersonated_scope: Optional[pulumi.Input['UsersPermissionsUserImpersonatedScopeArgs']] = None,
                 view_scope: Optional[pulumi.Input['UsersPermissionsViewScopeArgs']] = None):
        """
        The set of arguments for constructing a UsersPermissions resource.
        """
        UsersPermissionsArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            realm_id=realm_id,
            impersonate_scope=impersonate_scope,
            manage_group_membership_scope=manage_group_membership_scope,
            manage_scope=manage_scope,
            map_roles_scope=map_roles_scope,
            user_impersonated_scope=user_impersonated_scope,
            view_scope=view_scope,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             realm_id: Optional[pulumi.Input[str]] = None,
             impersonate_scope: Optional[pulumi.Input['UsersPermissionsImpersonateScopeArgs']] = None,
             manage_group_membership_scope: Optional[pulumi.Input['UsersPermissionsManageGroupMembershipScopeArgs']] = None,
             manage_scope: Optional[pulumi.Input['UsersPermissionsManageScopeArgs']] = None,
             map_roles_scope: Optional[pulumi.Input['UsersPermissionsMapRolesScopeArgs']] = None,
             user_impersonated_scope: Optional[pulumi.Input['UsersPermissionsUserImpersonatedScopeArgs']] = None,
             view_scope: Optional[pulumi.Input['UsersPermissionsViewScopeArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if realm_id is None:
            raise TypeError("Missing 'realm_id' argument")
        if impersonate_scope is None and 'impersonateScope' in kwargs:
            impersonate_scope = kwargs['impersonateScope']
        if manage_group_membership_scope is None and 'manageGroupMembershipScope' in kwargs:
            manage_group_membership_scope = kwargs['manageGroupMembershipScope']
        if manage_scope is None and 'manageScope' in kwargs:
            manage_scope = kwargs['manageScope']
        if map_roles_scope is None and 'mapRolesScope' in kwargs:
            map_roles_scope = kwargs['mapRolesScope']
        if user_impersonated_scope is None and 'userImpersonatedScope' in kwargs:
            user_impersonated_scope = kwargs['userImpersonatedScope']
        if view_scope is None and 'viewScope' in kwargs:
            view_scope = kwargs['viewScope']

        _setter("realm_id", realm_id)
        if impersonate_scope is not None:
            _setter("impersonate_scope", impersonate_scope)
        if manage_group_membership_scope is not None:
            _setter("manage_group_membership_scope", manage_group_membership_scope)
        if manage_scope is not None:
            _setter("manage_scope", manage_scope)
        if map_roles_scope is not None:
            _setter("map_roles_scope", map_roles_scope)
        if user_impersonated_scope is not None:
            _setter("user_impersonated_scope", user_impersonated_scope)
        if view_scope is not None:
            _setter("view_scope", view_scope)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="impersonateScope")
    def impersonate_scope(self) -> Optional[pulumi.Input['UsersPermissionsImpersonateScopeArgs']]:
        return pulumi.get(self, "impersonate_scope")

    @impersonate_scope.setter
    def impersonate_scope(self, value: Optional[pulumi.Input['UsersPermissionsImpersonateScopeArgs']]):
        pulumi.set(self, "impersonate_scope", value)

    @property
    @pulumi.getter(name="manageGroupMembershipScope")
    def manage_group_membership_scope(self) -> Optional[pulumi.Input['UsersPermissionsManageGroupMembershipScopeArgs']]:
        return pulumi.get(self, "manage_group_membership_scope")

    @manage_group_membership_scope.setter
    def manage_group_membership_scope(self, value: Optional[pulumi.Input['UsersPermissionsManageGroupMembershipScopeArgs']]):
        pulumi.set(self, "manage_group_membership_scope", value)

    @property
    @pulumi.getter(name="manageScope")
    def manage_scope(self) -> Optional[pulumi.Input['UsersPermissionsManageScopeArgs']]:
        return pulumi.get(self, "manage_scope")

    @manage_scope.setter
    def manage_scope(self, value: Optional[pulumi.Input['UsersPermissionsManageScopeArgs']]):
        pulumi.set(self, "manage_scope", value)

    @property
    @pulumi.getter(name="mapRolesScope")
    def map_roles_scope(self) -> Optional[pulumi.Input['UsersPermissionsMapRolesScopeArgs']]:
        return pulumi.get(self, "map_roles_scope")

    @map_roles_scope.setter
    def map_roles_scope(self, value: Optional[pulumi.Input['UsersPermissionsMapRolesScopeArgs']]):
        pulumi.set(self, "map_roles_scope", value)

    @property
    @pulumi.getter(name="userImpersonatedScope")
    def user_impersonated_scope(self) -> Optional[pulumi.Input['UsersPermissionsUserImpersonatedScopeArgs']]:
        return pulumi.get(self, "user_impersonated_scope")

    @user_impersonated_scope.setter
    def user_impersonated_scope(self, value: Optional[pulumi.Input['UsersPermissionsUserImpersonatedScopeArgs']]):
        pulumi.set(self, "user_impersonated_scope", value)

    @property
    @pulumi.getter(name="viewScope")
    def view_scope(self) -> Optional[pulumi.Input['UsersPermissionsViewScopeArgs']]:
        return pulumi.get(self, "view_scope")

    @view_scope.setter
    def view_scope(self, value: Optional[pulumi.Input['UsersPermissionsViewScopeArgs']]):
        pulumi.set(self, "view_scope", value)


@pulumi.input_type
class _UsersPermissionsState:
    def __init__(__self__, *,
                 authorization_resource_server_id: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 impersonate_scope: Optional[pulumi.Input['UsersPermissionsImpersonateScopeArgs']] = None,
                 manage_group_membership_scope: Optional[pulumi.Input['UsersPermissionsManageGroupMembershipScopeArgs']] = None,
                 manage_scope: Optional[pulumi.Input['UsersPermissionsManageScopeArgs']] = None,
                 map_roles_scope: Optional[pulumi.Input['UsersPermissionsMapRolesScopeArgs']] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 user_impersonated_scope: Optional[pulumi.Input['UsersPermissionsUserImpersonatedScopeArgs']] = None,
                 view_scope: Optional[pulumi.Input['UsersPermissionsViewScopeArgs']] = None):
        """
        Input properties used for looking up and filtering UsersPermissions resources.
        :param pulumi.Input[str] authorization_resource_server_id: Resource server id representing the realm management client on which this permission is managed
        """
        _UsersPermissionsState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            authorization_resource_server_id=authorization_resource_server_id,
            enabled=enabled,
            impersonate_scope=impersonate_scope,
            manage_group_membership_scope=manage_group_membership_scope,
            manage_scope=manage_scope,
            map_roles_scope=map_roles_scope,
            realm_id=realm_id,
            user_impersonated_scope=user_impersonated_scope,
            view_scope=view_scope,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             authorization_resource_server_id: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             impersonate_scope: Optional[pulumi.Input['UsersPermissionsImpersonateScopeArgs']] = None,
             manage_group_membership_scope: Optional[pulumi.Input['UsersPermissionsManageGroupMembershipScopeArgs']] = None,
             manage_scope: Optional[pulumi.Input['UsersPermissionsManageScopeArgs']] = None,
             map_roles_scope: Optional[pulumi.Input['UsersPermissionsMapRolesScopeArgs']] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             user_impersonated_scope: Optional[pulumi.Input['UsersPermissionsUserImpersonatedScopeArgs']] = None,
             view_scope: Optional[pulumi.Input['UsersPermissionsViewScopeArgs']] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if authorization_resource_server_id is None and 'authorizationResourceServerId' in kwargs:
            authorization_resource_server_id = kwargs['authorizationResourceServerId']
        if impersonate_scope is None and 'impersonateScope' in kwargs:
            impersonate_scope = kwargs['impersonateScope']
        if manage_group_membership_scope is None and 'manageGroupMembershipScope' in kwargs:
            manage_group_membership_scope = kwargs['manageGroupMembershipScope']
        if manage_scope is None and 'manageScope' in kwargs:
            manage_scope = kwargs['manageScope']
        if map_roles_scope is None and 'mapRolesScope' in kwargs:
            map_roles_scope = kwargs['mapRolesScope']
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if user_impersonated_scope is None and 'userImpersonatedScope' in kwargs:
            user_impersonated_scope = kwargs['userImpersonatedScope']
        if view_scope is None and 'viewScope' in kwargs:
            view_scope = kwargs['viewScope']

        if authorization_resource_server_id is not None:
            _setter("authorization_resource_server_id", authorization_resource_server_id)
        if enabled is not None:
            _setter("enabled", enabled)
        if impersonate_scope is not None:
            _setter("impersonate_scope", impersonate_scope)
        if manage_group_membership_scope is not None:
            _setter("manage_group_membership_scope", manage_group_membership_scope)
        if manage_scope is not None:
            _setter("manage_scope", manage_scope)
        if map_roles_scope is not None:
            _setter("map_roles_scope", map_roles_scope)
        if realm_id is not None:
            _setter("realm_id", realm_id)
        if user_impersonated_scope is not None:
            _setter("user_impersonated_scope", user_impersonated_scope)
        if view_scope is not None:
            _setter("view_scope", view_scope)

    @property
    @pulumi.getter(name="authorizationResourceServerId")
    def authorization_resource_server_id(self) -> Optional[pulumi.Input[str]]:
        """
        Resource server id representing the realm management client on which this permission is managed
        """
        return pulumi.get(self, "authorization_resource_server_id")

    @authorization_resource_server_id.setter
    def authorization_resource_server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_resource_server_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="impersonateScope")
    def impersonate_scope(self) -> Optional[pulumi.Input['UsersPermissionsImpersonateScopeArgs']]:
        return pulumi.get(self, "impersonate_scope")

    @impersonate_scope.setter
    def impersonate_scope(self, value: Optional[pulumi.Input['UsersPermissionsImpersonateScopeArgs']]):
        pulumi.set(self, "impersonate_scope", value)

    @property
    @pulumi.getter(name="manageGroupMembershipScope")
    def manage_group_membership_scope(self) -> Optional[pulumi.Input['UsersPermissionsManageGroupMembershipScopeArgs']]:
        return pulumi.get(self, "manage_group_membership_scope")

    @manage_group_membership_scope.setter
    def manage_group_membership_scope(self, value: Optional[pulumi.Input['UsersPermissionsManageGroupMembershipScopeArgs']]):
        pulumi.set(self, "manage_group_membership_scope", value)

    @property
    @pulumi.getter(name="manageScope")
    def manage_scope(self) -> Optional[pulumi.Input['UsersPermissionsManageScopeArgs']]:
        return pulumi.get(self, "manage_scope")

    @manage_scope.setter
    def manage_scope(self, value: Optional[pulumi.Input['UsersPermissionsManageScopeArgs']]):
        pulumi.set(self, "manage_scope", value)

    @property
    @pulumi.getter(name="mapRolesScope")
    def map_roles_scope(self) -> Optional[pulumi.Input['UsersPermissionsMapRolesScopeArgs']]:
        return pulumi.get(self, "map_roles_scope")

    @map_roles_scope.setter
    def map_roles_scope(self, value: Optional[pulumi.Input['UsersPermissionsMapRolesScopeArgs']]):
        pulumi.set(self, "map_roles_scope", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="userImpersonatedScope")
    def user_impersonated_scope(self) -> Optional[pulumi.Input['UsersPermissionsUserImpersonatedScopeArgs']]:
        return pulumi.get(self, "user_impersonated_scope")

    @user_impersonated_scope.setter
    def user_impersonated_scope(self, value: Optional[pulumi.Input['UsersPermissionsUserImpersonatedScopeArgs']]):
        pulumi.set(self, "user_impersonated_scope", value)

    @property
    @pulumi.getter(name="viewScope")
    def view_scope(self) -> Optional[pulumi.Input['UsersPermissionsViewScopeArgs']]:
        return pulumi.get(self, "view_scope")

    @view_scope.setter
    def view_scope(self, value: Optional[pulumi.Input['UsersPermissionsViewScopeArgs']]):
        pulumi.set(self, "view_scope", value)


class UsersPermissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 impersonate_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsImpersonateScopeArgs']]] = None,
                 manage_group_membership_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsManageGroupMembershipScopeArgs']]] = None,
                 manage_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsManageScopeArgs']]] = None,
                 map_roles_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsMapRolesScopeArgs']]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 user_impersonated_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsUserImpersonatedScopeArgs']]] = None,
                 view_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsViewScopeArgs']]] = None,
                 __props__=None):
        """
        Create a UsersPermissions resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UsersPermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a UsersPermissions resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param UsersPermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UsersPermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            UsersPermissionsArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 impersonate_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsImpersonateScopeArgs']]] = None,
                 manage_group_membership_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsManageGroupMembershipScopeArgs']]] = None,
                 manage_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsManageScopeArgs']]] = None,
                 map_roles_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsMapRolesScopeArgs']]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 user_impersonated_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsUserImpersonatedScopeArgs']]] = None,
                 view_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsViewScopeArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UsersPermissionsArgs.__new__(UsersPermissionsArgs)

            impersonate_scope = _utilities.configure(impersonate_scope, UsersPermissionsImpersonateScopeArgs, True)
            __props__.__dict__["impersonate_scope"] = impersonate_scope
            manage_group_membership_scope = _utilities.configure(manage_group_membership_scope, UsersPermissionsManageGroupMembershipScopeArgs, True)
            __props__.__dict__["manage_group_membership_scope"] = manage_group_membership_scope
            manage_scope = _utilities.configure(manage_scope, UsersPermissionsManageScopeArgs, True)
            __props__.__dict__["manage_scope"] = manage_scope
            map_roles_scope = _utilities.configure(map_roles_scope, UsersPermissionsMapRolesScopeArgs, True)
            __props__.__dict__["map_roles_scope"] = map_roles_scope
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            user_impersonated_scope = _utilities.configure(user_impersonated_scope, UsersPermissionsUserImpersonatedScopeArgs, True)
            __props__.__dict__["user_impersonated_scope"] = user_impersonated_scope
            view_scope = _utilities.configure(view_scope, UsersPermissionsViewScopeArgs, True)
            __props__.__dict__["view_scope"] = view_scope
            __props__.__dict__["authorization_resource_server_id"] = None
            __props__.__dict__["enabled"] = None
        super(UsersPermissions, __self__).__init__(
            'keycloak:index/usersPermissions:UsersPermissions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization_resource_server_id: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            impersonate_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsImpersonateScopeArgs']]] = None,
            manage_group_membership_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsManageGroupMembershipScopeArgs']]] = None,
            manage_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsManageScopeArgs']]] = None,
            map_roles_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsMapRolesScopeArgs']]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            user_impersonated_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsUserImpersonatedScopeArgs']]] = None,
            view_scope: Optional[pulumi.Input[pulumi.InputType['UsersPermissionsViewScopeArgs']]] = None) -> 'UsersPermissions':
        """
        Get an existing UsersPermissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_resource_server_id: Resource server id representing the realm management client on which this permission is managed
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UsersPermissionsState.__new__(_UsersPermissionsState)

        __props__.__dict__["authorization_resource_server_id"] = authorization_resource_server_id
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["impersonate_scope"] = impersonate_scope
        __props__.__dict__["manage_group_membership_scope"] = manage_group_membership_scope
        __props__.__dict__["manage_scope"] = manage_scope
        __props__.__dict__["map_roles_scope"] = map_roles_scope
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["user_impersonated_scope"] = user_impersonated_scope
        __props__.__dict__["view_scope"] = view_scope
        return UsersPermissions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizationResourceServerId")
    def authorization_resource_server_id(self) -> pulumi.Output[str]:
        """
        Resource server id representing the realm management client on which this permission is managed
        """
        return pulumi.get(self, "authorization_resource_server_id")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="impersonateScope")
    def impersonate_scope(self) -> pulumi.Output[Optional['outputs.UsersPermissionsImpersonateScope']]:
        return pulumi.get(self, "impersonate_scope")

    @property
    @pulumi.getter(name="manageGroupMembershipScope")
    def manage_group_membership_scope(self) -> pulumi.Output[Optional['outputs.UsersPermissionsManageGroupMembershipScope']]:
        return pulumi.get(self, "manage_group_membership_scope")

    @property
    @pulumi.getter(name="manageScope")
    def manage_scope(self) -> pulumi.Output[Optional['outputs.UsersPermissionsManageScope']]:
        return pulumi.get(self, "manage_scope")

    @property
    @pulumi.getter(name="mapRolesScope")
    def map_roles_scope(self) -> pulumi.Output[Optional['outputs.UsersPermissionsMapRolesScope']]:
        return pulumi.get(self, "map_roles_scope")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="userImpersonatedScope")
    def user_impersonated_scope(self) -> pulumi.Output[Optional['outputs.UsersPermissionsUserImpersonatedScope']]:
        return pulumi.get(self, "user_impersonated_scope")

    @property
    @pulumi.getter(name="viewScope")
    def view_scope(self) -> pulumi.Output[Optional['outputs.UsersPermissionsViewScope']]:
        return pulumi.get(self, "view_scope")

