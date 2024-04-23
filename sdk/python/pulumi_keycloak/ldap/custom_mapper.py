# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['CustomMapperArgs', 'CustomMapper']

@pulumi.input_type
class CustomMapperArgs:
    def __init__(__self__, *,
                 ldap_user_federation_id: pulumi.Input[str],
                 provider_id: pulumi.Input[str],
                 provider_type: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomMapper resource.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] provider_id: The id of the LDAP mapper implemented in MapperFactory.
        :param pulumi.Input[str] provider_type: The fully-qualified Java class name of the custom LDAP mapper.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        :param pulumi.Input[Mapping[str, Any]] config: A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        """
        pulumi.set(__self__, "ldap_user_federation_id", ldap_user_federation_id)
        pulumi.set(__self__, "provider_id", provider_id)
        pulumi.set(__self__, "provider_type", provider_type)
        pulumi.set(__self__, "realm_id", realm_id)
        if config is not None:
            pulumi.set(__self__, "config", config)
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
    @pulumi.getter(name="providerId")
    def provider_id(self) -> pulumi.Input[str]:
        """
        The id of the LDAP mapper implemented in MapperFactory.
        """
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_id", value)

    @property
    @pulumi.getter(name="providerType")
    def provider_type(self) -> pulumi.Input[str]:
        """
        The fully-qualified Java class name of the custom LDAP mapper.
        """
        return pulumi.get(self, "provider_type")

    @provider_type.setter
    def provider_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_type", value)

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
    def config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "config", value)

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
class _CustomMapperState:
    def __init__(__self__, *,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 provider_type: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomMapper resources.
        :param pulumi.Input[Mapping[str, Any]] config: A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] provider_id: The id of the LDAP mapper implemented in MapperFactory.
        :param pulumi.Input[str] provider_type: The fully-qualified Java class name of the custom LDAP mapper.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        """
        if config is not None:
            pulumi.set(__self__, "config", config)
        if ldap_user_federation_id is not None:
            pulumi.set(__self__, "ldap_user_federation_id", ldap_user_federation_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)
        if provider_type is not None:
            pulumi.set(__self__, "provider_type", provider_type)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "config", value)

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
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[str]]:
        """
        The id of the LDAP mapper implemented in MapperFactory.
        """
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_id", value)

    @property
    @pulumi.getter(name="providerType")
    def provider_type(self) -> Optional[pulumi.Input[str]]:
        """
        The fully-qualified Java class name of the custom LDAP mapper.
        """
        return pulumi.get(self, "provider_type")

    @provider_type.setter
    def provider_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_type", value)

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


class CustomMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 provider_type: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating and managing custom attribute mappers for Keycloak users federated via LDAP.

        The LDAP custom mapper is implemented and deployed into Keycloak as a custom provider. This resource allows to
        specify the custom id and custom implementation class of the self-implemented attribute mapper as well as additional
        properties via config map.

        The custom mapper should already be deployed into keycloak in order to be correctly configured.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        ldap_user_federation = keycloak.ldap.UserFederation("ldap_user_federation",
            name="openldap",
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
        custom_mapper = keycloak.ldap.CustomMapper("custom_mapper",
            name="custom-mapper",
            realm_id=openldap["realmId"],
            ldap_user_federation_id=openldap["id"],
            provider_id="custom-provider-registered-in-keycloak",
            provider_type="com.example.custom.ldap.mappers.CustomMapper",
            config={
                "attribute.name": "name",
                "attribute.value": "value",
            })
        ```

        ## Import

        LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.

        The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:ldap/customMapper:CustomMapper custom_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] config: A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] provider_id: The id of the LDAP mapper implemented in MapperFactory.
        :param pulumi.Input[str] provider_type: The fully-qualified Java class name of the custom LDAP mapper.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing custom attribute mappers for Keycloak users federated via LDAP.

        The LDAP custom mapper is implemented and deployed into Keycloak as a custom provider. This resource allows to
        specify the custom id and custom implementation class of the self-implemented attribute mapper as well as additional
        properties via config map.

        The custom mapper should already be deployed into keycloak in order to be correctly configured.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        ldap_user_federation = keycloak.ldap.UserFederation("ldap_user_federation",
            name="openldap",
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
        custom_mapper = keycloak.ldap.CustomMapper("custom_mapper",
            name="custom-mapper",
            realm_id=openldap["realmId"],
            ldap_user_federation_id=openldap["id"],
            provider_id="custom-provider-registered-in-keycloak",
            provider_type="com.example.custom.ldap.mappers.CustomMapper",
            config={
                "attribute.name": "name",
                "attribute.value": "value",
            })
        ```

        ## Import

        LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.

        The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:ldap/customMapper:CustomMapper custom_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
        ```

        :param str resource_name: The name of the resource.
        :param CustomMapperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomMapperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 provider_type: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomMapperArgs.__new__(CustomMapperArgs)

            __props__.__dict__["config"] = config
            if ldap_user_federation_id is None and not opts.urn:
                raise TypeError("Missing required property 'ldap_user_federation_id'")
            __props__.__dict__["ldap_user_federation_id"] = ldap_user_federation_id
            __props__.__dict__["name"] = name
            if provider_id is None and not opts.urn:
                raise TypeError("Missing required property 'provider_id'")
            __props__.__dict__["provider_id"] = provider_id
            if provider_type is None and not opts.urn:
                raise TypeError("Missing required property 'provider_type'")
            __props__.__dict__["provider_type"] = provider_type
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(CustomMapper, __self__).__init__(
            'keycloak:ldap/customMapper:CustomMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            provider_id: Optional[pulumi.Input[str]] = None,
            provider_type: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'CustomMapper':
        """
        Get an existing CustomMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] config: A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] provider_id: The id of the LDAP mapper implemented in MapperFactory.
        :param pulumi.Input[str] provider_type: The fully-qualified Java class name of the custom LDAP mapper.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomMapperState.__new__(_CustomMapperState)

        __props__.__dict__["config"] = config
        __props__.__dict__["ldap_user_federation_id"] = ldap_user_federation_id
        __props__.__dict__["name"] = name
        __props__.__dict__["provider_id"] = provider_id
        __props__.__dict__["provider_type"] = provider_type
        __props__.__dict__["realm_id"] = realm_id
        return CustomMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        """
        return pulumi.get(self, "config")

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
    @pulumi.getter(name="providerId")
    def provider_id(self) -> pulumi.Output[str]:
        """
        The id of the LDAP mapper implemented in MapperFactory.
        """
        return pulumi.get(self, "provider_id")

    @property
    @pulumi.getter(name="providerType")
    def provider_type(self) -> pulumi.Output[str]:
        """
        The fully-qualified Java class name of the custom LDAP mapper.
        """
        return pulumi.get(self, "provider_type")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm that this LDAP mapper will exist in.
        """
        return pulumi.get(self, "realm_id")

