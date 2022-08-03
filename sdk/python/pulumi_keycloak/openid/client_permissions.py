# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['ClientPermissionsArgs', 'ClientPermissions']

@pulumi.input_type
class ClientPermissionsArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 configure_scope: Optional[pulumi.Input['ClientPermissionsConfigureScopeArgs']] = None,
                 manage_scope: Optional[pulumi.Input['ClientPermissionsManageScopeArgs']] = None,
                 map_roles_client_scope_scope: Optional[pulumi.Input['ClientPermissionsMapRolesClientScopeScopeArgs']] = None,
                 map_roles_composite_scope: Optional[pulumi.Input['ClientPermissionsMapRolesCompositeScopeArgs']] = None,
                 map_roles_scope: Optional[pulumi.Input['ClientPermissionsMapRolesScopeArgs']] = None,
                 token_exchange_scope: Optional[pulumi.Input['ClientPermissionsTokenExchangeScopeArgs']] = None,
                 view_scope: Optional[pulumi.Input['ClientPermissionsViewScopeArgs']] = None):
        """
        The set of arguments for constructing a ClientPermissions resource.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "realm_id", realm_id)
        if configure_scope is not None:
            pulumi.set(__self__, "configure_scope", configure_scope)
        if manage_scope is not None:
            pulumi.set(__self__, "manage_scope", manage_scope)
        if map_roles_client_scope_scope is not None:
            pulumi.set(__self__, "map_roles_client_scope_scope", map_roles_client_scope_scope)
        if map_roles_composite_scope is not None:
            pulumi.set(__self__, "map_roles_composite_scope", map_roles_composite_scope)
        if map_roles_scope is not None:
            pulumi.set(__self__, "map_roles_scope", map_roles_scope)
        if token_exchange_scope is not None:
            pulumi.set(__self__, "token_exchange_scope", token_exchange_scope)
        if view_scope is not None:
            pulumi.set(__self__, "view_scope", view_scope)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="configureScope")
    def configure_scope(self) -> Optional[pulumi.Input['ClientPermissionsConfigureScopeArgs']]:
        return pulumi.get(self, "configure_scope")

    @configure_scope.setter
    def configure_scope(self, value: Optional[pulumi.Input['ClientPermissionsConfigureScopeArgs']]):
        pulumi.set(self, "configure_scope", value)

    @property
    @pulumi.getter(name="manageScope")
    def manage_scope(self) -> Optional[pulumi.Input['ClientPermissionsManageScopeArgs']]:
        return pulumi.get(self, "manage_scope")

    @manage_scope.setter
    def manage_scope(self, value: Optional[pulumi.Input['ClientPermissionsManageScopeArgs']]):
        pulumi.set(self, "manage_scope", value)

    @property
    @pulumi.getter(name="mapRolesClientScopeScope")
    def map_roles_client_scope_scope(self) -> Optional[pulumi.Input['ClientPermissionsMapRolesClientScopeScopeArgs']]:
        return pulumi.get(self, "map_roles_client_scope_scope")

    @map_roles_client_scope_scope.setter
    def map_roles_client_scope_scope(self, value: Optional[pulumi.Input['ClientPermissionsMapRolesClientScopeScopeArgs']]):
        pulumi.set(self, "map_roles_client_scope_scope", value)

    @property
    @pulumi.getter(name="mapRolesCompositeScope")
    def map_roles_composite_scope(self) -> Optional[pulumi.Input['ClientPermissionsMapRolesCompositeScopeArgs']]:
        return pulumi.get(self, "map_roles_composite_scope")

    @map_roles_composite_scope.setter
    def map_roles_composite_scope(self, value: Optional[pulumi.Input['ClientPermissionsMapRolesCompositeScopeArgs']]):
        pulumi.set(self, "map_roles_composite_scope", value)

    @property
    @pulumi.getter(name="mapRolesScope")
    def map_roles_scope(self) -> Optional[pulumi.Input['ClientPermissionsMapRolesScopeArgs']]:
        return pulumi.get(self, "map_roles_scope")

    @map_roles_scope.setter
    def map_roles_scope(self, value: Optional[pulumi.Input['ClientPermissionsMapRolesScopeArgs']]):
        pulumi.set(self, "map_roles_scope", value)

    @property
    @pulumi.getter(name="tokenExchangeScope")
    def token_exchange_scope(self) -> Optional[pulumi.Input['ClientPermissionsTokenExchangeScopeArgs']]:
        return pulumi.get(self, "token_exchange_scope")

    @token_exchange_scope.setter
    def token_exchange_scope(self, value: Optional[pulumi.Input['ClientPermissionsTokenExchangeScopeArgs']]):
        pulumi.set(self, "token_exchange_scope", value)

    @property
    @pulumi.getter(name="viewScope")
    def view_scope(self) -> Optional[pulumi.Input['ClientPermissionsViewScopeArgs']]:
        return pulumi.get(self, "view_scope")

    @view_scope.setter
    def view_scope(self, value: Optional[pulumi.Input['ClientPermissionsViewScopeArgs']]):
        pulumi.set(self, "view_scope", value)


@pulumi.input_type
class _ClientPermissionsState:
    def __init__(__self__, *,
                 authorization_resource_server_id: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 configure_scope: Optional[pulumi.Input['ClientPermissionsConfigureScopeArgs']] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 manage_scope: Optional[pulumi.Input['ClientPermissionsManageScopeArgs']] = None,
                 map_roles_client_scope_scope: Optional[pulumi.Input['ClientPermissionsMapRolesClientScopeScopeArgs']] = None,
                 map_roles_composite_scope: Optional[pulumi.Input['ClientPermissionsMapRolesCompositeScopeArgs']] = None,
                 map_roles_scope: Optional[pulumi.Input['ClientPermissionsMapRolesScopeArgs']] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 token_exchange_scope: Optional[pulumi.Input['ClientPermissionsTokenExchangeScopeArgs']] = None,
                 view_scope: Optional[pulumi.Input['ClientPermissionsViewScopeArgs']] = None):
        """
        Input properties used for looking up and filtering ClientPermissions resources.
        :param pulumi.Input[str] authorization_resource_server_id: Resource server id representing the realm management client on which this permission is managed
        """
        if authorization_resource_server_id is not None:
            pulumi.set(__self__, "authorization_resource_server_id", authorization_resource_server_id)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if configure_scope is not None:
            pulumi.set(__self__, "configure_scope", configure_scope)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if manage_scope is not None:
            pulumi.set(__self__, "manage_scope", manage_scope)
        if map_roles_client_scope_scope is not None:
            pulumi.set(__self__, "map_roles_client_scope_scope", map_roles_client_scope_scope)
        if map_roles_composite_scope is not None:
            pulumi.set(__self__, "map_roles_composite_scope", map_roles_composite_scope)
        if map_roles_scope is not None:
            pulumi.set(__self__, "map_roles_scope", map_roles_scope)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if token_exchange_scope is not None:
            pulumi.set(__self__, "token_exchange_scope", token_exchange_scope)
        if view_scope is not None:
            pulumi.set(__self__, "view_scope", view_scope)

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
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="configureScope")
    def configure_scope(self) -> Optional[pulumi.Input['ClientPermissionsConfigureScopeArgs']]:
        return pulumi.get(self, "configure_scope")

    @configure_scope.setter
    def configure_scope(self, value: Optional[pulumi.Input['ClientPermissionsConfigureScopeArgs']]):
        pulumi.set(self, "configure_scope", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="manageScope")
    def manage_scope(self) -> Optional[pulumi.Input['ClientPermissionsManageScopeArgs']]:
        return pulumi.get(self, "manage_scope")

    @manage_scope.setter
    def manage_scope(self, value: Optional[pulumi.Input['ClientPermissionsManageScopeArgs']]):
        pulumi.set(self, "manage_scope", value)

    @property
    @pulumi.getter(name="mapRolesClientScopeScope")
    def map_roles_client_scope_scope(self) -> Optional[pulumi.Input['ClientPermissionsMapRolesClientScopeScopeArgs']]:
        return pulumi.get(self, "map_roles_client_scope_scope")

    @map_roles_client_scope_scope.setter
    def map_roles_client_scope_scope(self, value: Optional[pulumi.Input['ClientPermissionsMapRolesClientScopeScopeArgs']]):
        pulumi.set(self, "map_roles_client_scope_scope", value)

    @property
    @pulumi.getter(name="mapRolesCompositeScope")
    def map_roles_composite_scope(self) -> Optional[pulumi.Input['ClientPermissionsMapRolesCompositeScopeArgs']]:
        return pulumi.get(self, "map_roles_composite_scope")

    @map_roles_composite_scope.setter
    def map_roles_composite_scope(self, value: Optional[pulumi.Input['ClientPermissionsMapRolesCompositeScopeArgs']]):
        pulumi.set(self, "map_roles_composite_scope", value)

    @property
    @pulumi.getter(name="mapRolesScope")
    def map_roles_scope(self) -> Optional[pulumi.Input['ClientPermissionsMapRolesScopeArgs']]:
        return pulumi.get(self, "map_roles_scope")

    @map_roles_scope.setter
    def map_roles_scope(self, value: Optional[pulumi.Input['ClientPermissionsMapRolesScopeArgs']]):
        pulumi.set(self, "map_roles_scope", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="tokenExchangeScope")
    def token_exchange_scope(self) -> Optional[pulumi.Input['ClientPermissionsTokenExchangeScopeArgs']]:
        return pulumi.get(self, "token_exchange_scope")

    @token_exchange_scope.setter
    def token_exchange_scope(self, value: Optional[pulumi.Input['ClientPermissionsTokenExchangeScopeArgs']]):
        pulumi.set(self, "token_exchange_scope", value)

    @property
    @pulumi.getter(name="viewScope")
    def view_scope(self) -> Optional[pulumi.Input['ClientPermissionsViewScopeArgs']]:
        return pulumi.get(self, "view_scope")

    @view_scope.setter
    def view_scope(self, value: Optional[pulumi.Input['ClientPermissionsViewScopeArgs']]):
        pulumi.set(self, "view_scope", value)


class ClientPermissions(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 configure_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsConfigureScopeArgs']]] = None,
                 manage_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsManageScopeArgs']]] = None,
                 map_roles_client_scope_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesClientScopeScopeArgs']]] = None,
                 map_roles_composite_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesCompositeScopeArgs']]] = None,
                 map_roles_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesScopeArgs']]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 token_exchange_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsTokenExchangeScopeArgs']]] = None,
                 view_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsViewScopeArgs']]] = None,
                 __props__=None):
        """
        Create a ClientPermissions resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClientPermissionsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ClientPermissions resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ClientPermissionsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientPermissionsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 configure_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsConfigureScopeArgs']]] = None,
                 manage_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsManageScopeArgs']]] = None,
                 map_roles_client_scope_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesClientScopeScopeArgs']]] = None,
                 map_roles_composite_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesCompositeScopeArgs']]] = None,
                 map_roles_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesScopeArgs']]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 token_exchange_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsTokenExchangeScopeArgs']]] = None,
                 view_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsViewScopeArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientPermissionsArgs.__new__(ClientPermissionsArgs)

            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["configure_scope"] = configure_scope
            __props__.__dict__["manage_scope"] = manage_scope
            __props__.__dict__["map_roles_client_scope_scope"] = map_roles_client_scope_scope
            __props__.__dict__["map_roles_composite_scope"] = map_roles_composite_scope
            __props__.__dict__["map_roles_scope"] = map_roles_scope
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            __props__.__dict__["token_exchange_scope"] = token_exchange_scope
            __props__.__dict__["view_scope"] = view_scope
            __props__.__dict__["authorization_resource_server_id"] = None
            __props__.__dict__["enabled"] = None
        super(ClientPermissions, __self__).__init__(
            'keycloak:openid/clientPermissions:ClientPermissions',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization_resource_server_id: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            configure_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsConfigureScopeArgs']]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            manage_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsManageScopeArgs']]] = None,
            map_roles_client_scope_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesClientScopeScopeArgs']]] = None,
            map_roles_composite_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesCompositeScopeArgs']]] = None,
            map_roles_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsMapRolesScopeArgs']]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            token_exchange_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsTokenExchangeScopeArgs']]] = None,
            view_scope: Optional[pulumi.Input[pulumi.InputType['ClientPermissionsViewScopeArgs']]] = None) -> 'ClientPermissions':
        """
        Get an existing ClientPermissions resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_resource_server_id: Resource server id representing the realm management client on which this permission is managed
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientPermissionsState.__new__(_ClientPermissionsState)

        __props__.__dict__["authorization_resource_server_id"] = authorization_resource_server_id
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["configure_scope"] = configure_scope
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["manage_scope"] = manage_scope
        __props__.__dict__["map_roles_client_scope_scope"] = map_roles_client_scope_scope
        __props__.__dict__["map_roles_composite_scope"] = map_roles_composite_scope
        __props__.__dict__["map_roles_scope"] = map_roles_scope
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["token_exchange_scope"] = token_exchange_scope
        __props__.__dict__["view_scope"] = view_scope
        return ClientPermissions(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizationResourceServerId")
    def authorization_resource_server_id(self) -> pulumi.Output[str]:
        """
        Resource server id representing the realm management client on which this permission is managed
        """
        return pulumi.get(self, "authorization_resource_server_id")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="configureScope")
    def configure_scope(self) -> pulumi.Output[Optional['outputs.ClientPermissionsConfigureScope']]:
        return pulumi.get(self, "configure_scope")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="manageScope")
    def manage_scope(self) -> pulumi.Output[Optional['outputs.ClientPermissionsManageScope']]:
        return pulumi.get(self, "manage_scope")

    @property
    @pulumi.getter(name="mapRolesClientScopeScope")
    def map_roles_client_scope_scope(self) -> pulumi.Output[Optional['outputs.ClientPermissionsMapRolesClientScopeScope']]:
        return pulumi.get(self, "map_roles_client_scope_scope")

    @property
    @pulumi.getter(name="mapRolesCompositeScope")
    def map_roles_composite_scope(self) -> pulumi.Output[Optional['outputs.ClientPermissionsMapRolesCompositeScope']]:
        return pulumi.get(self, "map_roles_composite_scope")

    @property
    @pulumi.getter(name="mapRolesScope")
    def map_roles_scope(self) -> pulumi.Output[Optional['outputs.ClientPermissionsMapRolesScope']]:
        return pulumi.get(self, "map_roles_scope")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="tokenExchangeScope")
    def token_exchange_scope(self) -> pulumi.Output[Optional['outputs.ClientPermissionsTokenExchangeScope']]:
        return pulumi.get(self, "token_exchange_scope")

    @property
    @pulumi.getter(name="viewScope")
    def view_scope(self) -> pulumi.Output[Optional['outputs.ClientPermissionsViewScope']]:
        return pulumi.get(self, "view_scope")

