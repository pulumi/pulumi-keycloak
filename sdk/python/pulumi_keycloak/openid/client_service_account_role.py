# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClientServiceAccountRoleArgs', 'ClientServiceAccountRole']

@pulumi.input_type
class ClientServiceAccountRoleArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 role: pulumi.Input[str],
                 service_account_user_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ClientServiceAccountRole resource.
        :param pulumi.Input[str] client_id: The id of the client that provides the role.
        :param pulumi.Input[str] realm_id: The realm the clients and roles belong to.
        :param pulumi.Input[str] role: The name of the role that is assigned.
        :param pulumi.Input[str] service_account_user_id: The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        """
        ClientServiceAccountRoleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            client_id=client_id,
            realm_id=realm_id,
            role=role,
            service_account_user_id=service_account_user_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             client_id: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             service_account_user_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if client_id is None and 'clientId' in kwargs:
            client_id = kwargs['clientId']
        if client_id is None:
            raise TypeError("Missing 'client_id' argument")
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if realm_id is None:
            raise TypeError("Missing 'realm_id' argument")
        if role is None:
            raise TypeError("Missing 'role' argument")
        if service_account_user_id is None and 'serviceAccountUserId' in kwargs:
            service_account_user_id = kwargs['serviceAccountUserId']
        if service_account_user_id is None:
            raise TypeError("Missing 'service_account_user_id' argument")

        _setter("client_id", client_id)
        _setter("realm_id", realm_id)
        _setter("role", role)
        _setter("service_account_user_id", service_account_user_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        """
        The id of the client that provides the role.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm the clients and roles belong to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The name of the role that is assigned.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="serviceAccountUserId")
    def service_account_user_id(self) -> pulumi.Input[str]:
        """
        The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        """
        return pulumi.get(self, "service_account_user_id")

    @service_account_user_id.setter
    def service_account_user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_user_id", value)


@pulumi.input_type
class _ClientServiceAccountRoleState:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 service_account_user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClientServiceAccountRole resources.
        :param pulumi.Input[str] client_id: The id of the client that provides the role.
        :param pulumi.Input[str] realm_id: The realm the clients and roles belong to.
        :param pulumi.Input[str] role: The name of the role that is assigned.
        :param pulumi.Input[str] service_account_user_id: The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        """
        _ClientServiceAccountRoleState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            client_id=client_id,
            realm_id=realm_id,
            role=role,
            service_account_user_id=service_account_user_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             client_id: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             role: Optional[pulumi.Input[str]] = None,
             service_account_user_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if client_id is None and 'clientId' in kwargs:
            client_id = kwargs['clientId']
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if service_account_user_id is None and 'serviceAccountUserId' in kwargs:
            service_account_user_id = kwargs['serviceAccountUserId']

        if client_id is not None:
            _setter("client_id", client_id)
        if realm_id is not None:
            _setter("realm_id", realm_id)
        if role is not None:
            _setter("role", role)
        if service_account_user_id is not None:
            _setter("service_account_user_id", service_account_user_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the client that provides the role.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm the clients and roles belong to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the role that is assigned.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="serviceAccountUserId")
    def service_account_user_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        """
        return pulumi.get(self, "service_account_user_id")

    @service_account_user_id.setter
    def service_account_user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_user_id", value)


class ClientServiceAccountRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 service_account_user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for assigning client roles to the service account of an openid client.
        You need to set `service_accounts_enabled` to `true` for the openid client that should be assigned the role.

        If you'd like to attach realm roles to a service account, please use the `openid.ClientServiceAccountRealmRole`
        resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        # client1 provides a role to other clients
        client1 = keycloak.openid.Client("client1", realm_id=realm.id)
        client1_role = keycloak.Role("client1Role",
            realm_id=realm.id,
            client_id=client1.id,
            description="A role that client1 provides")
        # client2 is assigned the role of client1
        client2 = keycloak.openid.Client("client2",
            realm_id=realm.id,
            service_accounts_enabled=True)
        client2_service_account_role = keycloak.openid.ClientServiceAccountRole("client2ServiceAccountRole",
            realm_id=realm.id,
            service_account_user_id=client2.service_account_user_id,
            client_id=client1.id,
            role=client1_role.name)
        ```

        ## Import

        This resource can be imported using the format `{{realmId}}/{{serviceAccountUserId}}/{{clientId}}/{{roleId}}`. Examplebash

        ```sh
         $ pulumi import keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole client2_service_account_role my-realm/489ba513-1ceb-49ba-ae0b-1ab1f5099ebf/baf01820-0f8b-4494-9be2-fb3bc8a397a4/c7230ab7-8e4e-4135-995d-e81b50696ad8
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The id of the client that provides the role.
        :param pulumi.Input[str] realm_id: The realm the clients and roles belong to.
        :param pulumi.Input[str] role: The name of the role that is assigned.
        :param pulumi.Input[str] service_account_user_id: The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClientServiceAccountRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for assigning client roles to the service account of an openid client.
        You need to set `service_accounts_enabled` to `true` for the openid client that should be assigned the role.

        If you'd like to attach realm roles to a service account, please use the `openid.ClientServiceAccountRealmRole`
        resource.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        # client1 provides a role to other clients
        client1 = keycloak.openid.Client("client1", realm_id=realm.id)
        client1_role = keycloak.Role("client1Role",
            realm_id=realm.id,
            client_id=client1.id,
            description="A role that client1 provides")
        # client2 is assigned the role of client1
        client2 = keycloak.openid.Client("client2",
            realm_id=realm.id,
            service_accounts_enabled=True)
        client2_service_account_role = keycloak.openid.ClientServiceAccountRole("client2ServiceAccountRole",
            realm_id=realm.id,
            service_account_user_id=client2.service_account_user_id,
            client_id=client1.id,
            role=client1_role.name)
        ```

        ## Import

        This resource can be imported using the format `{{realmId}}/{{serviceAccountUserId}}/{{clientId}}/{{roleId}}`. Examplebash

        ```sh
         $ pulumi import keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole client2_service_account_role my-realm/489ba513-1ceb-49ba-ae0b-1ab1f5099ebf/baf01820-0f8b-4494-9be2-fb3bc8a397a4/c7230ab7-8e4e-4135-995d-e81b50696ad8
        ```

        :param str resource_name: The name of the resource.
        :param ClientServiceAccountRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientServiceAccountRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ClientServiceAccountRoleArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 service_account_user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientServiceAccountRoleArgs.__new__(ClientServiceAccountRoleArgs)

            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            if service_account_user_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_user_id'")
            __props__.__dict__["service_account_user_id"] = service_account_user_id
        super(ClientServiceAccountRole, __self__).__init__(
            'keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            service_account_user_id: Optional[pulumi.Input[str]] = None) -> 'ClientServiceAccountRole':
        """
        Get an existing ClientServiceAccountRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The id of the client that provides the role.
        :param pulumi.Input[str] realm_id: The realm the clients and roles belong to.
        :param pulumi.Input[str] role: The name of the role that is assigned.
        :param pulumi.Input[str] service_account_user_id: The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientServiceAccountRoleState.__new__(_ClientServiceAccountRoleState)

        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["role"] = role
        __props__.__dict__["service_account_user_id"] = service_account_user_id
        return ClientServiceAccountRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        The id of the client that provides the role.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm the clients and roles belong to.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The name of the role that is assigned.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="serviceAccountUserId")
    def service_account_user_id(self) -> pulumi.Output[str]:
        """
        The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        """
        return pulumi.get(self, "service_account_user_id")

