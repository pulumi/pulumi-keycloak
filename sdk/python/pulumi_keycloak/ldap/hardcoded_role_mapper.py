# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HardcodedRoleMapperArgs', 'HardcodedRoleMapper']

@pulumi.input_type
class HardcodedRoleMapperArgs:
    def __init__(__self__, *,
                 ldap_user_federation_id: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 role: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HardcodedRoleMapper resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        :param pulumi.Input[str] role: The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        """
        pulumi.set(__self__, "ldap_user_federation_id", ldap_user_federation_id)
        pulumi.set(__self__, "realm_id", realm_id)
        pulumi.set(__self__, "role", role)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="ldapUserFederationId")
    def ldap_user_federation_id(self) -> pulumi.Input[str]:
        """
        The ID of the LDAP user federation provider to attach this mapper to.
        """
        return pulumi.get(self, "ldap_user_federation_id")

    @ldap_user_federation_id.setter
    def ldap_user_federation_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ldap_user_federation_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm that this LDAP mapper will exist in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of this mapper when displayed in the console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _HardcodedRoleMapperState:
    def __init__(__self__, *,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HardcodedRoleMapper resources.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        :param pulumi.Input[str] role: The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
        """
        if ldap_user_federation_id is not None:
            pulumi.set(__self__, "ldap_user_federation_id", ldap_user_federation_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="ldapUserFederationId")
    def ldap_user_federation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the LDAP user federation provider to attach this mapper to.
        """
        return pulumi.get(self, "ldap_user_federation_id")

    @ldap_user_federation_id.setter
    def ldap_user_federation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ldap_user_federation_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of this mapper when displayed in the console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm that this LDAP mapper will exist in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class HardcodedRoleMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating and managing hardcoded role mappers for Keycloak users federated via LDAP.

        The LDAP hardcoded role mapper will grant a specified Keycloak role to each Keycloak user linked with LDAP.

        ## Example Usage
        ### Realm Role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        ldap_user_federation = keycloak.ldap.UserFederation("ldapUserFederation",
            realm_id=realm.id,
            username_ldap_attribute="cn",
            rdn_ldap_attribute="cn",
            uuid_ldap_attribute="entryDN",
            user_object_classes=[
                "simpleSecurityObject",
                "organizationalRole",
            ],
            connection_url="ldap://openldap",
            users_dn="dc=example,dc=org",
            bind_dn="cn=admin,dc=example,dc=org",
            bind_credential="admin")
        realm_admin_role = keycloak.Role("realmAdminRole",
            realm_id=realm.id,
            description="My Realm Role")
        assign_admin_role_to_all_users = keycloak.ldap.HardcodedRoleMapper("assignAdminRoleToAllUsers",
            realm_id=realm.id,
            ldap_user_federation_id=ldap_user_federation.id,
            role=realm_admin_role.name)
        ```
        ### Client Role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        ldap_user_federation = keycloak.ldap.UserFederation("ldapUserFederation",
            realm_id=realm.id,
            username_ldap_attribute="cn",
            rdn_ldap_attribute="cn",
            uuid_ldap_attribute="entryDN",
            user_object_classes=[
                "simpleSecurityObject",
                "organizationalRole",
            ],
            connection_url="ldap://openldap",
            users_dn="dc=example,dc=org",
            bind_dn="cn=admin,dc=example,dc=org",
            bind_credential="admin")
        realm_management = keycloak.openid.get_client_output(realm_id=realm.id,
            client_id="realm-management")
        create_client = pulumi.Output.all(realm.id, realm_management).apply(lambda id, realm_management: keycloak.get_role_output(realm_id=id,
            client_id=realm_management.id,
            name="create-client"))
        assign_admin_role_to_all_users = keycloak.ldap.HardcodedRoleMapper("assignAdminRoleToAllUsers",
            realm_id=realm.id,
            ldap_user_federation_id=ldap_user_federation.id,
            role=pulumi.Output.all(realm_management, create_client).apply(lambda realm_management, create_client: f"{realm_management.client_id}.{create_client.name}"))
        ```

        ## Import

        LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.

         The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.

         Example:

         bash

        ```sh
        $ pulumi import keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper assign_admin_role_to_all_users my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        :param pulumi.Input[str] role: The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HardcodedRoleMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing hardcoded role mappers for Keycloak users federated via LDAP.

        The LDAP hardcoded role mapper will grant a specified Keycloak role to each Keycloak user linked with LDAP.

        ## Example Usage
        ### Realm Role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        ldap_user_federation = keycloak.ldap.UserFederation("ldapUserFederation",
            realm_id=realm.id,
            username_ldap_attribute="cn",
            rdn_ldap_attribute="cn",
            uuid_ldap_attribute="entryDN",
            user_object_classes=[
                "simpleSecurityObject",
                "organizationalRole",
            ],
            connection_url="ldap://openldap",
            users_dn="dc=example,dc=org",
            bind_dn="cn=admin,dc=example,dc=org",
            bind_credential="admin")
        realm_admin_role = keycloak.Role("realmAdminRole",
            realm_id=realm.id,
            description="My Realm Role")
        assign_admin_role_to_all_users = keycloak.ldap.HardcodedRoleMapper("assignAdminRoleToAllUsers",
            realm_id=realm.id,
            ldap_user_federation_id=ldap_user_federation.id,
            role=realm_admin_role.name)
        ```
        ### Client Role)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        ldap_user_federation = keycloak.ldap.UserFederation("ldapUserFederation",
            realm_id=realm.id,
            username_ldap_attribute="cn",
            rdn_ldap_attribute="cn",
            uuid_ldap_attribute="entryDN",
            user_object_classes=[
                "simpleSecurityObject",
                "organizationalRole",
            ],
            connection_url="ldap://openldap",
            users_dn="dc=example,dc=org",
            bind_dn="cn=admin,dc=example,dc=org",
            bind_credential="admin")
        realm_management = keycloak.openid.get_client_output(realm_id=realm.id,
            client_id="realm-management")
        create_client = pulumi.Output.all(realm.id, realm_management).apply(lambda id, realm_management: keycloak.get_role_output(realm_id=id,
            client_id=realm_management.id,
            name="create-client"))
        assign_admin_role_to_all_users = keycloak.ldap.HardcodedRoleMapper("assignAdminRoleToAllUsers",
            realm_id=realm.id,
            ldap_user_federation_id=ldap_user_federation.id,
            role=pulumi.Output.all(realm_management, create_client).apply(lambda realm_management, create_client: f"{realm_management.client_id}.{create_client.name}"))
        ```

        ## Import

        LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.

         The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.

         Example:

         bash

        ```sh
        $ pulumi import keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper assign_admin_role_to_all_users my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
        ```

        :param str resource_name: The name of the resource.
        :param HardcodedRoleMapperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HardcodedRoleMapperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HardcodedRoleMapperArgs.__new__(HardcodedRoleMapperArgs)

            if ldap_user_federation_id is None and not opts.urn:
                raise TypeError("Missing required property 'ldap_user_federation_id'")
            __props__.__dict__["ldap_user_federation_id"] = ldap_user_federation_id
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
        super(HardcodedRoleMapper, __self__).__init__(
            'keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'HardcodedRoleMapper':
        """
        Get an existing HardcodedRoleMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        :param pulumi.Input[str] role: The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HardcodedRoleMapperState.__new__(_HardcodedRoleMapperState)

        __props__.__dict__["ldap_user_federation_id"] = ldap_user_federation_id
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["role"] = role
        return HardcodedRoleMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="ldapUserFederationId")
    def ldap_user_federation_id(self) -> pulumi.Output[str]:
        """
        The ID of the LDAP user federation provider to attach this mapper to.
        """
        return pulumi.get(self, "ldap_user_federation_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Display name of this mapper when displayed in the console.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm that this LDAP mapper will exist in.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
        """
        return pulumi.get(self, "role")

