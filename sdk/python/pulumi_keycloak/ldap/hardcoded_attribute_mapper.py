# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HardcodedAttributeMapperArgs', 'HardcodedAttributeMapper']

@pulumi.input_type
class HardcodedAttributeMapperArgs:
    def __init__(__self__, *,
                 attribute_name: pulumi.Input[str],
                 attribute_value: pulumi.Input[str],
                 ldap_user_federation_id: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HardcodedAttributeMapper resource.
        :param pulumi.Input[str] attribute_name: The name of the LDAP attribute to set.
        :param pulumi.Input[str] attribute_value: The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        """
        HardcodedAttributeMapperArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            attribute_name=attribute_name,
            attribute_value=attribute_value,
            ldap_user_federation_id=ldap_user_federation_id,
            realm_id=realm_id,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             attribute_name: pulumi.Input[str],
             attribute_value: pulumi.Input[str],
             ldap_user_federation_id: pulumi.Input[str],
             realm_id: pulumi.Input[str],
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("attribute_name", attribute_name)
        _setter("attribute_value", attribute_value)
        _setter("ldap_user_federation_id", ldap_user_federation_id)
        _setter("realm_id", realm_id)
        if name is not None:
            _setter("name", name)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> pulumi.Input[str]:
        """
        The name of the LDAP attribute to set.
        """
        return pulumi.get(self, "attribute_name")

    @attribute_name.setter
    def attribute_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute_name", value)

    @property
    @pulumi.getter(name="attributeValue")
    def attribute_value(self) -> pulumi.Input[str]:
        """
        The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
        """
        return pulumi.get(self, "attribute_value")

    @attribute_value.setter
    def attribute_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute_value", value)

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
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of this mapper when displayed in the console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _HardcodedAttributeMapperState:
    def __init__(__self__, *,
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HardcodedAttributeMapper resources.
        :param pulumi.Input[str] attribute_name: The name of the LDAP attribute to set.
        :param pulumi.Input[str] attribute_value: The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        """
        _HardcodedAttributeMapperState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            attribute_name=attribute_name,
            attribute_value=attribute_value,
            ldap_user_federation_id=ldap_user_federation_id,
            name=name,
            realm_id=realm_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             attribute_name: Optional[pulumi.Input[str]] = None,
             attribute_value: Optional[pulumi.Input[str]] = None,
             ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if attribute_name is not None:
            _setter("attribute_name", attribute_name)
        if attribute_value is not None:
            _setter("attribute_value", attribute_value)
        if ldap_user_federation_id is not None:
            _setter("ldap_user_federation_id", ldap_user_federation_id)
        if name is not None:
            _setter("name", name)
        if realm_id is not None:
            _setter("realm_id", realm_id)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the LDAP attribute to set.
        """
        return pulumi.get(self, "attribute_name")

    @attribute_name.setter
    def attribute_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_name", value)

    @property
    @pulumi.getter(name="attributeValue")
    def attribute_value(self) -> Optional[pulumi.Input[str]]:
        """
        The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
        """
        return pulumi.get(self, "attribute_value")

    @attribute_value.setter
    def attribute_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "attribute_value", value)

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


class HardcodedAttributeMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating and managing hardcoded attribute mappers for Keycloak users federated via LDAP.

        The LDAP hardcoded attribute mapper will set the specified value to the LDAP attribute.

        **NOTE**: This mapper only works when the `sync_registrations` attribute on the `ldap.UserFederation` resource is set to `true`.

        ## Example Usage

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
            bind_credential="admin",
            sync_registrations=True)
        assign_bar_to_foo = keycloak.ldap.HardcodedAttributeMapper("assignBarToFoo",
            realm_id=realm.id,
            ldap_user_federation_id=ldap_user_federation.id,
            attribute_name="foo",
            attribute_value="bar")
        ```

        ## Import

        LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash

        ```sh
         $ pulumi import keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper assign_bar_to_foo my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attribute_name: The name of the LDAP attribute to set.
        :param pulumi.Input[str] attribute_value: The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HardcodedAttributeMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing hardcoded attribute mappers for Keycloak users federated via LDAP.

        The LDAP hardcoded attribute mapper will set the specified value to the LDAP attribute.

        **NOTE**: This mapper only works when the `sync_registrations` attribute on the `ldap.UserFederation` resource is set to `true`.

        ## Example Usage

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
            bind_credential="admin",
            sync_registrations=True)
        assign_bar_to_foo = keycloak.ldap.HardcodedAttributeMapper("assignBarToFoo",
            realm_id=realm.id,
            ldap_user_federation_id=ldap_user_federation.id,
            attribute_name="foo",
            attribute_value="bar")
        ```

        ## Import

        LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash

        ```sh
         $ pulumi import keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper assign_bar_to_foo my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
        ```

        :param str resource_name: The name of the resource.
        :param HardcodedAttributeMapperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HardcodedAttributeMapperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            HardcodedAttributeMapperArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attribute_name: Optional[pulumi.Input[str]] = None,
                 attribute_value: Optional[pulumi.Input[str]] = None,
                 ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HardcodedAttributeMapperArgs.__new__(HardcodedAttributeMapperArgs)

            if attribute_name is None and not opts.urn:
                raise TypeError("Missing required property 'attribute_name'")
            __props__.__dict__["attribute_name"] = attribute_name
            if attribute_value is None and not opts.urn:
                raise TypeError("Missing required property 'attribute_value'")
            __props__.__dict__["attribute_value"] = attribute_value
            if ldap_user_federation_id is None and not opts.urn:
                raise TypeError("Missing required property 'ldap_user_federation_id'")
            __props__.__dict__["ldap_user_federation_id"] = ldap_user_federation_id
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(HardcodedAttributeMapper, __self__).__init__(
            'keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attribute_name: Optional[pulumi.Input[str]] = None,
            attribute_value: Optional[pulumi.Input[str]] = None,
            ldap_user_federation_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'HardcodedAttributeMapper':
        """
        Get an existing HardcodedAttributeMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attribute_name: The name of the LDAP attribute to set.
        :param pulumi.Input[str] attribute_value: The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
        :param pulumi.Input[str] ldap_user_federation_id: The ID of the LDAP user federation provider to attach this mapper to.
        :param pulumi.Input[str] name: Display name of this mapper when displayed in the console.
        :param pulumi.Input[str] realm_id: The realm that this LDAP mapper will exist in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _HardcodedAttributeMapperState.__new__(_HardcodedAttributeMapperState)

        __props__.__dict__["attribute_name"] = attribute_name
        __props__.__dict__["attribute_value"] = attribute_value
        __props__.__dict__["ldap_user_federation_id"] = ldap_user_federation_id
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        return HardcodedAttributeMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="attributeName")
    def attribute_name(self) -> pulumi.Output[str]:
        """
        The name of the LDAP attribute to set.
        """
        return pulumi.get(self, "attribute_name")

    @property
    @pulumi.getter(name="attributeValue")
    def attribute_value(self) -> pulumi.Output[str]:
        """
        The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
        """
        return pulumi.get(self, "attribute_value")

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

