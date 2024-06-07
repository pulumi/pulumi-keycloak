# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserPropertyProtocolMapperArgs', 'UserPropertyProtocolMapper']

@pulumi.input_type
class UserPropertyProtocolMapperArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 saml_attribute_name: pulumi.Input[str],
                 saml_attribute_name_format: pulumi.Input[str],
                 user_property: pulumi.Input[str],
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UserPropertyProtocolMapper resource.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        pulumi.set(__self__, "saml_attribute_name", saml_attribute_name)
        pulumi.set(__self__, "saml_attribute_name_format", saml_attribute_name_format)
        pulumi.set(__self__, "user_property", user_property)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
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
    @pulumi.getter(name="samlAttributeName")
    def saml_attribute_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "saml_attribute_name")

    @saml_attribute_name.setter
    def saml_attribute_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "saml_attribute_name", value)

    @property
    @pulumi.getter(name="samlAttributeNameFormat")
    def saml_attribute_name_format(self) -> pulumi.Input[str]:
        return pulumi.get(self, "saml_attribute_name_format")

    @saml_attribute_name_format.setter
    def saml_attribute_name_format(self, value: pulumi.Input[str]):
        pulumi.set(self, "saml_attribute_name_format", value)

    @property
    @pulumi.getter(name="userProperty")
    def user_property(self) -> pulumi.Input[str]:
        return pulumi.get(self, "user_property")

    @user_property.setter
    def user_property(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_property", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_scope_id")

    @client_scope_id.setter
    def client_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_scope_id", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _UserPropertyProtocolMapperState:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 saml_attribute_name: Optional[pulumi.Input[str]] = None,
                 saml_attribute_name_format: Optional[pulumi.Input[str]] = None,
                 user_property: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserPropertyProtocolMapper resources.
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if friendly_name is not None:
            pulumi.set(__self__, "friendly_name", friendly_name)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if saml_attribute_name is not None:
            pulumi.set(__self__, "saml_attribute_name", saml_attribute_name)
        if saml_attribute_name_format is not None:
            pulumi.set(__self__, "saml_attribute_name_format", saml_attribute_name_format)
        if user_property is not None:
            pulumi.set(__self__, "user_property", user_property)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_scope_id")

    @client_scope_id.setter
    def client_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_scope_id", value)

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "friendly_name")

    @friendly_name.setter
    def friendly_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "friendly_name", value)

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

    @property
    @pulumi.getter(name="samlAttributeName")
    def saml_attribute_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "saml_attribute_name")

    @saml_attribute_name.setter
    def saml_attribute_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "saml_attribute_name", value)

    @property
    @pulumi.getter(name="samlAttributeNameFormat")
    def saml_attribute_name_format(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "saml_attribute_name_format")

    @saml_attribute_name_format.setter
    def saml_attribute_name_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "saml_attribute_name_format", value)

    @property
    @pulumi.getter(name="userProperty")
    def user_property(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "user_property")

    @user_property.setter
    def user_property(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_property", value)


class UserPropertyProtocolMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 saml_attribute_name: Optional[pulumi.Input[str]] = None,
                 saml_attribute_name_format: Optional[pulumi.Input[str]] = None,
                 user_property: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # saml.UserPropertyProtocolMapper

        Allows for creating and managing user property protocol mappers for
        SAML clients within Keycloak.

        SAML user property protocol mappers allow you to map properties of the Keycloak
        user model to an attribute in a SAML assertion. Protocol mappers
        can be defined for a single client, or they can be defined for a client scope which
        can be shared between multiple different clients.

        ### Example Usage (Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        saml_client = keycloak.saml.Client("saml_client",
            realm_id=test["id"],
            client_id="test-saml-client",
            name="test-saml-client")
        saml_user_property_mapper = keycloak.saml.UserPropertyProtocolMapper("saml_user_property_mapper",
            realm_id=test["id"],
            client_id=saml_client.id,
            name="email-user-property-mapper",
            user_property="email",
            saml_attribute_name="email",
            saml_attribute_name_format="Unspecified")
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required if `client_scope_id` is not specified) The SAML client this protocol mapper is attached to.
        - `client_scope_id` - (Required if `client_id` is not specified) The SAML client scope this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `user_property` - (Required) The property of the Keycloak user model to map.
        - `friendly_name` - (Optional) An optional human-friendly name for this attribute.
        - `saml_attribute_name` - (Required) The name of the SAML attribute.
        - `saml_attribute_name_format` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.

        ### Import

        Protocol mappers can be imported using one of the following formats:
        - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
        - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`

        Example:

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserPropertyProtocolMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # saml.UserPropertyProtocolMapper

        Allows for creating and managing user property protocol mappers for
        SAML clients within Keycloak.

        SAML user property protocol mappers allow you to map properties of the Keycloak
        user model to an attribute in a SAML assertion. Protocol mappers
        can be defined for a single client, or they can be defined for a client scope which
        can be shared between multiple different clients.

        ### Example Usage (Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        saml_client = keycloak.saml.Client("saml_client",
            realm_id=test["id"],
            client_id="test-saml-client",
            name="test-saml-client")
        saml_user_property_mapper = keycloak.saml.UserPropertyProtocolMapper("saml_user_property_mapper",
            realm_id=test["id"],
            client_id=saml_client.id,
            name="email-user-property-mapper",
            user_property="email",
            saml_attribute_name="email",
            saml_attribute_name_format="Unspecified")
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required if `client_scope_id` is not specified) The SAML client this protocol mapper is attached to.
        - `client_scope_id` - (Required if `client_id` is not specified) The SAML client scope this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `user_property` - (Required) The property of the Keycloak user model to map.
        - `friendly_name` - (Optional) An optional human-friendly name for this attribute.
        - `saml_attribute_name` - (Required) The name of the SAML attribute.
        - `saml_attribute_name_format` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.

        ### Import

        Protocol mappers can be imported using one of the following formats:
        - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
        - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`

        Example:

        :param str resource_name: The name of the resource.
        :param UserPropertyProtocolMapperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserPropertyProtocolMapperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 friendly_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 saml_attribute_name: Optional[pulumi.Input[str]] = None,
                 saml_attribute_name_format: Optional[pulumi.Input[str]] = None,
                 user_property: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserPropertyProtocolMapperArgs.__new__(UserPropertyProtocolMapperArgs)

            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_scope_id"] = client_scope_id
            __props__.__dict__["friendly_name"] = friendly_name
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            if saml_attribute_name is None and not opts.urn:
                raise TypeError("Missing required property 'saml_attribute_name'")
            __props__.__dict__["saml_attribute_name"] = saml_attribute_name
            if saml_attribute_name_format is None and not opts.urn:
                raise TypeError("Missing required property 'saml_attribute_name_format'")
            __props__.__dict__["saml_attribute_name_format"] = saml_attribute_name_format
            if user_property is None and not opts.urn:
                raise TypeError("Missing required property 'user_property'")
            __props__.__dict__["user_property"] = user_property
        super(UserPropertyProtocolMapper, __self__).__init__(
            'keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_scope_id: Optional[pulumi.Input[str]] = None,
            friendly_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            saml_attribute_name: Optional[pulumi.Input[str]] = None,
            saml_attribute_name_format: Optional[pulumi.Input[str]] = None,
            user_property: Optional[pulumi.Input[str]] = None) -> 'UserPropertyProtocolMapper':
        """
        Get an existing UserPropertyProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserPropertyProtocolMapperState.__new__(_UserPropertyProtocolMapperState)

        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_scope_id"] = client_scope_id
        __props__.__dict__["friendly_name"] = friendly_name
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["saml_attribute_name"] = saml_attribute_name
        __props__.__dict__["saml_attribute_name_format"] = saml_attribute_name_format
        __props__.__dict__["user_property"] = user_property
        return UserPropertyProtocolMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "client_scope_id")

    @property
    @pulumi.getter(name="friendlyName")
    def friendly_name(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "friendly_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="samlAttributeName")
    def saml_attribute_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "saml_attribute_name")

    @property
    @pulumi.getter(name="samlAttributeNameFormat")
    def saml_attribute_name_format(self) -> pulumi.Output[str]:
        return pulumi.get(self, "saml_attribute_name_format")

    @property
    @pulumi.getter(name="userProperty")
    def user_property(self) -> pulumi.Output[str]:
        return pulumi.get(self, "user_property")

