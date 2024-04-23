# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 composite_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Role resource.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if composite_roles is not None:
            pulumi.set(__self__, "composite_roles", composite_roles)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="compositeRoles")
    def composite_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "composite_roles")

    @composite_roles.setter
    def composite_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "composite_roles", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 composite_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        """
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if composite_roles is not None:
            pulumi.set(__self__, "composite_roles", composite_roles)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="compositeRoles")
    def composite_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "composite_roles")

    @composite_roles.setter
    def composite_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "composite_roles", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 composite_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # Role

        Allows for creating and managing roles within Keycloak.

        Roles allow you define privileges within Keycloak and map them to users
        and groups.

        ### Example Usage (Realm role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        realm_role = keycloak.Role("realm_role",
            realm_id=realm.id,
            name="my-realm-role",
            description="My Realm Role")
        ```

        ### Example Usage (Client role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client = keycloak.openid.Client("client",
            realm_id=realm.id,
            client_id="client",
            name="client",
            enabled=True,
            access_type="BEARER-ONLY")
        client_role = keycloak.Role("client_role",
            realm_id=realm.id,
            client_id=client_keycloak_client["id"],
            name="my-client-role",
            description="My Client Role")
        ```

        ### Example Usage (Composite role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        # realm roles
        create_role = keycloak.Role("create_role",
            realm_id=realm.id,
            name="create")
        read_role = keycloak.Role("read_role",
            realm_id=realm.id,
            name="read")
        update_role = keycloak.Role("update_role",
            realm_id=realm.id,
            name="update")
        delete_role = keycloak.Role("delete_role",
            realm_id=realm.id,
            name="delete")
        # client role
        client = keycloak.openid.Client("client",
            realm_id=realm.id,
            client_id="client",
            name="client",
            enabled=True,
            access_type="BEARER-ONLY")
        client_role = keycloak.Role("client_role",
            realm_id=realm.id,
            client_id=client_keycloak_client["id"],
            name="my-client-role",
            description="My Client Role")
        admin_role = keycloak.Role("admin_role",
            realm_id=realm.id,
            name="admin",
            composite_roles=[
                "{keycloak_role.create_role.id}",
                "{keycloak_role.read_role.id}",
                "{keycloak_role.update_role.id}",
                "{keycloak_role.delete_role.id}",
                "{keycloak_role.client_role.id}",
            ])
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this role exists within.
        - `client_id` - (Optional) When specified, this role will be created as
          a client role attached to the client with the provided ID
        - `name` - (Required) The name of the role
        - `description` - (Optional) The description of the role
        - `composite_roles` - (Optional) When specified, this role will be a
          composite role, composed of all roles that have an ID present within
          this list.

        ### Import

        Roles can be imported using the format `{{realm_id}}/{{role_id}}`, where
        `role_id` is the unique ID that Keycloak assigns to the role. The ID is
        not easy to find in the GUI, but it appears in the URL when editing the
        role.

        Example:

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # Role

        Allows for creating and managing roles within Keycloak.

        Roles allow you define privileges within Keycloak and map them to users
        and groups.

        ### Example Usage (Realm role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        realm_role = keycloak.Role("realm_role",
            realm_id=realm.id,
            name="my-realm-role",
            description="My Realm Role")
        ```

        ### Example Usage (Client role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client = keycloak.openid.Client("client",
            realm_id=realm.id,
            client_id="client",
            name="client",
            enabled=True,
            access_type="BEARER-ONLY")
        client_role = keycloak.Role("client_role",
            realm_id=realm.id,
            client_id=client_keycloak_client["id"],
            name="my-client-role",
            description="My Client Role")
        ```

        ### Example Usage (Composite role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        # realm roles
        create_role = keycloak.Role("create_role",
            realm_id=realm.id,
            name="create")
        read_role = keycloak.Role("read_role",
            realm_id=realm.id,
            name="read")
        update_role = keycloak.Role("update_role",
            realm_id=realm.id,
            name="update")
        delete_role = keycloak.Role("delete_role",
            realm_id=realm.id,
            name="delete")
        # client role
        client = keycloak.openid.Client("client",
            realm_id=realm.id,
            client_id="client",
            name="client",
            enabled=True,
            access_type="BEARER-ONLY")
        client_role = keycloak.Role("client_role",
            realm_id=realm.id,
            client_id=client_keycloak_client["id"],
            name="my-client-role",
            description="My Client Role")
        admin_role = keycloak.Role("admin_role",
            realm_id=realm.id,
            name="admin",
            composite_roles=[
                "{keycloak_role.create_role.id}",
                "{keycloak_role.read_role.id}",
                "{keycloak_role.update_role.id}",
                "{keycloak_role.delete_role.id}",
                "{keycloak_role.client_role.id}",
            ])
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this role exists within.
        - `client_id` - (Optional) When specified, this role will be created as
          a client role attached to the client with the provided ID
        - `name` - (Required) The name of the role
        - `description` - (Optional) The description of the role
        - `composite_roles` - (Optional) When specified, this role will be a
          composite role, composed of all roles that have an ID present within
          this list.

        ### Import

        Roles can be imported using the format `{{realm_id}}/{{role_id}}`, where
        `role_id` is the unique ID that Keycloak assigns to the role. The ID is
        not easy to find in the GUI, but it appears in the URL when editing the
        role.

        Example:

        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 composite_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleArgs.__new__(RoleArgs)

            __props__.__dict__["attributes"] = attributes
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["composite_roles"] = composite_roles
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(Role, __self__).__init__(
            'keycloak:index/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            composite_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["attributes"] = attributes
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["composite_roles"] = composite_roles
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        return Role(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="compositeRoles")
    def composite_roles(self) -> pulumi.Output[Optional[Sequence[str]]]:
        return pulumi.get(self, "composite_roles")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm_id")

