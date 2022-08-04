# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClientScopeArgs', 'ClientScope']

@pulumi.input_type
class ClientScopeArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 consent_screen_text: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gui_order: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ClientScope resource.
        :param pulumi.Input[str] realm_id: The realm this client scope belongs to.
        :param pulumi.Input[str] consent_screen_text: When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        :param pulumi.Input[str] description: The description of this client scope in the GUI.
        :param pulumi.Input[int] gui_order: Specify order of the client scope in GUI (such as in Consent page) as integer.
        :param pulumi.Input[str] name: The display name of this client scope in the GUI.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        if consent_screen_text is not None:
            pulumi.set(__self__, "consent_screen_text", consent_screen_text)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if gui_order is not None:
            pulumi.set(__self__, "gui_order", gui_order)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm this client scope belongs to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="consentScreenText")
    def consent_screen_text(self) -> Optional[pulumi.Input[str]]:
        """
        When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        """
        return pulumi.get(self, "consent_screen_text")

    @consent_screen_text.setter
    def consent_screen_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consent_screen_text", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this client scope in the GUI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="guiOrder")
    def gui_order(self) -> Optional[pulumi.Input[int]]:
        """
        Specify order of the client scope in GUI (such as in Consent page) as integer.
        """
        return pulumi.get(self, "gui_order")

    @gui_order.setter
    def gui_order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gui_order", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of this client scope in the GUI.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ClientScopeState:
    def __init__(__self__, *,
                 consent_screen_text: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gui_order: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClientScope resources.
        :param pulumi.Input[str] consent_screen_text: When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        :param pulumi.Input[str] description: The description of this client scope in the GUI.
        :param pulumi.Input[int] gui_order: Specify order of the client scope in GUI (such as in Consent page) as integer.
        :param pulumi.Input[str] name: The display name of this client scope in the GUI.
        :param pulumi.Input[str] realm_id: The realm this client scope belongs to.
        """
        if consent_screen_text is not None:
            pulumi.set(__self__, "consent_screen_text", consent_screen_text)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if gui_order is not None:
            pulumi.set(__self__, "gui_order", gui_order)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="consentScreenText")
    def consent_screen_text(self) -> Optional[pulumi.Input[str]]:
        """
        When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        """
        return pulumi.get(self, "consent_screen_text")

    @consent_screen_text.setter
    def consent_screen_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "consent_screen_text", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of this client scope in the GUI.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="guiOrder")
    def gui_order(self) -> Optional[pulumi.Input[int]]:
        """
        Specify order of the client scope in GUI (such as in Consent page) as integer.
        """
        return pulumi.get(self, "gui_order")

    @gui_order.setter
    def gui_order(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "gui_order", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of this client scope in the GUI.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm this client scope belongs to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class ClientScope(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_screen_text: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gui_order: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating and managing Keycloak client scopes that can be attached to clients that use the SAML protocol.

        Client Scopes can be used to share common protocol and role mappings between multiple clients within a realm.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        saml_client_scope = keycloak.saml.ClientScope("samlClientScope",
            realm_id=realm.id,
            description="This scope will map a user's group memberships to SAML assertion",
            gui_order=1)
        ```

        ## Import

        Client scopes can be imported using the format `{{realm_id}}/{{client_scope_id}}`, where `client_scope_id` is the unique ID that Keycloak assigns to the client scope upon creation. This value can be found in the URI when editing this client scope in the GUI, and is typically a GUID. Examplebash

        ```sh
         $ pulumi import keycloak:saml/clientScope:ClientScope saml_client_scope my-realm/e8a5d115-6985-4de3-a0f5-732e1be4525e
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consent_screen_text: When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        :param pulumi.Input[str] description: The description of this client scope in the GUI.
        :param pulumi.Input[int] gui_order: Specify order of the client scope in GUI (such as in Consent page) as integer.
        :param pulumi.Input[str] name: The display name of this client scope in the GUI.
        :param pulumi.Input[str] realm_id: The realm this client scope belongs to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ClientScopeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing Keycloak client scopes that can be attached to clients that use the SAML protocol.

        Client Scopes can be used to share common protocol and role mappings between multiple clients within a realm.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        saml_client_scope = keycloak.saml.ClientScope("samlClientScope",
            realm_id=realm.id,
            description="This scope will map a user's group memberships to SAML assertion",
            gui_order=1)
        ```

        ## Import

        Client scopes can be imported using the format `{{realm_id}}/{{client_scope_id}}`, where `client_scope_id` is the unique ID that Keycloak assigns to the client scope upon creation. This value can be found in the URI when editing this client scope in the GUI, and is typically a GUID. Examplebash

        ```sh
         $ pulumi import keycloak:saml/clientScope:ClientScope saml_client_scope my-realm/e8a5d115-6985-4de3-a0f5-732e1be4525e
        ```

        :param str resource_name: The name of the resource.
        :param ClientScopeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ClientScopeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_screen_text: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gui_order: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ClientScopeArgs.__new__(ClientScopeArgs)

            __props__.__dict__["consent_screen_text"] = consent_screen_text
            __props__.__dict__["description"] = description
            __props__.__dict__["gui_order"] = gui_order
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(ClientScope, __self__).__init__(
            'keycloak:saml/clientScope:ClientScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            consent_screen_text: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            gui_order: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'ClientScope':
        """
        Get an existing ClientScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consent_screen_text: When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        :param pulumi.Input[str] description: The description of this client scope in the GUI.
        :param pulumi.Input[int] gui_order: Specify order of the client scope in GUI (such as in Consent page) as integer.
        :param pulumi.Input[str] name: The display name of this client scope in the GUI.
        :param pulumi.Input[str] realm_id: The realm this client scope belongs to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientScopeState.__new__(_ClientScopeState)

        __props__.__dict__["consent_screen_text"] = consent_screen_text
        __props__.__dict__["description"] = description
        __props__.__dict__["gui_order"] = gui_order
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        return ClientScope(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="consentScreenText")
    def consent_screen_text(self) -> pulumi.Output[Optional[str]]:
        """
        When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        """
        return pulumi.get(self, "consent_screen_text")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        The description of this client scope in the GUI.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="guiOrder")
    def gui_order(self) -> pulumi.Output[Optional[int]]:
        """
        Specify order of the client scope in GUI (such as in Consent page) as integer.
        """
        return pulumi.get(self, "gui_order")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The display name of this client scope in the GUI.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm this client scope belongs to.
        """
        return pulumi.get(self, "realm_id")

