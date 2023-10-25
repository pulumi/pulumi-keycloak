# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ClientScopeArgs', 'ClientScope']

@pulumi.input_type
class ClientScopeArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 consent_screen_text: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gui_order: Optional[pulumi.Input[int]] = None,
                 include_in_token_scope: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ClientScope resource.
        :param pulumi.Input[str] realm_id: The realm this client scope belongs to.
        :param pulumi.Input[str] consent_screen_text: When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        :param pulumi.Input[str] description: The description of this client scope in the GUI.
        :param pulumi.Input[int] gui_order: Specify order of the client scope in GUI (such as in Consent page) as integer.
        :param pulumi.Input[bool] include_in_token_scope: When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
        :param pulumi.Input[str] name: The display name of this client scope in the GUI.
        """
        ClientScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            realm_id=realm_id,
            consent_screen_text=consent_screen_text,
            description=description,
            gui_order=gui_order,
            include_in_token_scope=include_in_token_scope,
            name=name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             realm_id: Optional[pulumi.Input[str]] = None,
             consent_screen_text: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             gui_order: Optional[pulumi.Input[int]] = None,
             include_in_token_scope: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if realm_id is None:
            raise TypeError("Missing 'realm_id' argument")
        if consent_screen_text is None and 'consentScreenText' in kwargs:
            consent_screen_text = kwargs['consentScreenText']
        if gui_order is None and 'guiOrder' in kwargs:
            gui_order = kwargs['guiOrder']
        if include_in_token_scope is None and 'includeInTokenScope' in kwargs:
            include_in_token_scope = kwargs['includeInTokenScope']

        _setter("realm_id", realm_id)
        if consent_screen_text is not None:
            _setter("consent_screen_text", consent_screen_text)
        if description is not None:
            _setter("description", description)
        if gui_order is not None:
            _setter("gui_order", gui_order)
        if include_in_token_scope is not None:
            _setter("include_in_token_scope", include_in_token_scope)
        if name is not None:
            _setter("name", name)

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
    @pulumi.getter(name="includeInTokenScope")
    def include_in_token_scope(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
        """
        return pulumi.get(self, "include_in_token_scope")

    @include_in_token_scope.setter
    def include_in_token_scope(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_in_token_scope", value)

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
                 include_in_token_scope: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ClientScope resources.
        :param pulumi.Input[str] consent_screen_text: When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        :param pulumi.Input[str] description: The description of this client scope in the GUI.
        :param pulumi.Input[int] gui_order: Specify order of the client scope in GUI (such as in Consent page) as integer.
        :param pulumi.Input[bool] include_in_token_scope: When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
        :param pulumi.Input[str] name: The display name of this client scope in the GUI.
        :param pulumi.Input[str] realm_id: The realm this client scope belongs to.
        """
        _ClientScopeState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            consent_screen_text=consent_screen_text,
            description=description,
            gui_order=gui_order,
            include_in_token_scope=include_in_token_scope,
            name=name,
            realm_id=realm_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             consent_screen_text: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             gui_order: Optional[pulumi.Input[int]] = None,
             include_in_token_scope: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if consent_screen_text is None and 'consentScreenText' in kwargs:
            consent_screen_text = kwargs['consentScreenText']
        if gui_order is None and 'guiOrder' in kwargs:
            gui_order = kwargs['guiOrder']
        if include_in_token_scope is None and 'includeInTokenScope' in kwargs:
            include_in_token_scope = kwargs['includeInTokenScope']
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']

        if consent_screen_text is not None:
            _setter("consent_screen_text", consent_screen_text)
        if description is not None:
            _setter("description", description)
        if gui_order is not None:
            _setter("gui_order", gui_order)
        if include_in_token_scope is not None:
            _setter("include_in_token_scope", include_in_token_scope)
        if name is not None:
            _setter("name", name)
        if realm_id is not None:
            _setter("realm_id", realm_id)

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
    @pulumi.getter(name="includeInTokenScope")
    def include_in_token_scope(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
        """
        return pulumi.get(self, "include_in_token_scope")

    @include_in_token_scope.setter
    def include_in_token_scope(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "include_in_token_scope", value)

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
                 include_in_token_scope: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating and managing Keycloak client scopes that can be attached to clients that use the OpenID Connect protocol.

        Client Scopes can be used to share common protocol and role mappings between multiple clients within a realm. They can also
        be used by clients to conditionally request claims or roles for a user based on the OAuth 2.0 `scope` parameter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        openid_client_scope = keycloak.openid.ClientScope("openidClientScope",
            realm_id=realm.id,
            description="When requested, this scope will map a user's group memberships to a claim",
            include_in_token_scope=True,
            gui_order=1)
        ```

        ## Import

        Client scopes can be imported using the format `{{realm_id}}/{{client_scope_id}}`, where `client_scope_id` is the unique ID that Keycloak assigns to the client scope upon creation. This value can be found in the URI when editing this client scope in the GUI, and is typically a GUID. Examplebash

        ```sh
         $ pulumi import keycloak:openid/clientScope:ClientScope openid_client_scope my-realm/8e8f7fe1-df9b-40ed-bed3-4597aa0dac52
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] consent_screen_text: When set, a consent screen will be displayed to users authenticating to clients with this scope attached. The consent screen will display the string value of this attribute.
        :param pulumi.Input[str] description: The description of this client scope in the GUI.
        :param pulumi.Input[int] gui_order: Specify order of the client scope in GUI (such as in Consent page) as integer.
        :param pulumi.Input[bool] include_in_token_scope: When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
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
        Allows for creating and managing Keycloak client scopes that can be attached to clients that use the OpenID Connect protocol.

        Client Scopes can be used to share common protocol and role mappings between multiple clients within a realm. They can also
        be used by clients to conditionally request claims or roles for a user based on the OAuth 2.0 `scope` parameter.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        openid_client_scope = keycloak.openid.ClientScope("openidClientScope",
            realm_id=realm.id,
            description="When requested, this scope will map a user's group memberships to a claim",
            include_in_token_scope=True,
            gui_order=1)
        ```

        ## Import

        Client scopes can be imported using the format `{{realm_id}}/{{client_scope_id}}`, where `client_scope_id` is the unique ID that Keycloak assigns to the client scope upon creation. This value can be found in the URI when editing this client scope in the GUI, and is typically a GUID. Examplebash

        ```sh
         $ pulumi import keycloak:openid/clientScope:ClientScope openid_client_scope my-realm/8e8f7fe1-df9b-40ed-bed3-4597aa0dac52
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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ClientScopeArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 consent_screen_text: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 gui_order: Optional[pulumi.Input[int]] = None,
                 include_in_token_scope: Optional[pulumi.Input[bool]] = None,
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
            __props__.__dict__["include_in_token_scope"] = include_in_token_scope
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(ClientScope, __self__).__init__(
            'keycloak:openid/clientScope:ClientScope',
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
            include_in_token_scope: Optional[pulumi.Input[bool]] = None,
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
        :param pulumi.Input[bool] include_in_token_scope: When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
        :param pulumi.Input[str] name: The display name of this client scope in the GUI.
        :param pulumi.Input[str] realm_id: The realm this client scope belongs to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ClientScopeState.__new__(_ClientScopeState)

        __props__.__dict__["consent_screen_text"] = consent_screen_text
        __props__.__dict__["description"] = description
        __props__.__dict__["gui_order"] = gui_order
        __props__.__dict__["include_in_token_scope"] = include_in_token_scope
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
    @pulumi.getter(name="includeInTokenScope")
    def include_in_token_scope(self) -> pulumi.Output[Optional[bool]]:
        """
        When `true`, the name of this client scope will be added to the access token property 'scope' as well as to the Token Introspection Endpoint response.
        """
        return pulumi.get(self, "include_in_token_scope")

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

