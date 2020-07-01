# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class IdentityProvider(pulumi.CustomResource):
    accepts_prompt_none_forward_from_client: pulumi.Output[bool]
    """
    This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In
    case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly
    returned to client, but the request with prompt=none will be forwarded to this identity provider.
    """
    add_read_token_role_on_create: pulumi.Output[bool]
    """
    Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
    """
    alias: pulumi.Output[str]
    """
    The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
    """
    authenticate_by_default: pulumi.Output[bool]
    """
    Enable/disable authenticate users by default.
    """
    authorization_url: pulumi.Output[str]
    """
    OIDC authorization URL.
    """
    backchannel_supported: pulumi.Output[bool]
    """
    Does the external IDP support backchannel logout?
    """
    client_id: pulumi.Output[str]
    """
    Client ID.
    """
    client_secret: pulumi.Output[str]
    """
    Client Secret.
    """
    default_scopes: pulumi.Output[str]
    """
    The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to 'openid'.
    """
    display_name: pulumi.Output[str]
    """
    Friendly name for Identity Providers.
    """
    enabled: pulumi.Output[bool]
    """
    Enable/disable this identity provider.
    """
    extra_config: pulumi.Output[dict]
    first_broker_login_flow_alias: pulumi.Output[str]
    """
    Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
    that there is not yet existing Keycloak account linked with the authenticated identity provider account.
    """
    hide_on_login_page: pulumi.Output[bool]
    """
    Hide On Login Page.
    """
    internal_id: pulumi.Output[str]
    """
    Internal Identity Provider Id
    """
    jwks_url: pulumi.Output[str]
    """
    JSON Web Key Set URL
    """
    link_only: pulumi.Output[bool]
    """
    If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
    want to allow login from the provider, but want to integrate with a provider
    """
    login_hint: pulumi.Output[str]
    """
    Login Hint.
    """
    logout_url: pulumi.Output[str]
    """
    Logout URL
    """
    post_broker_login_flow_alias: pulumi.Output[str]
    """
    Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
    additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
    you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
    authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
    """
    provider_id: pulumi.Output[str]
    """
    provider id, is always oidc, unless you have a custom implementation
    """
    realm: pulumi.Output[str]
    """
    Realm Name
    """
    store_token: pulumi.Output[bool]
    """
    Enable/disable if tokens must be stored after authenticating users.
    """
    token_url: pulumi.Output[str]
    """
    Token URL.
    """
    trust_email: pulumi.Output[bool]
    """
    If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
    """
    ui_locales: pulumi.Output[bool]
    """
    Pass current locale to identity provider
    """
    user_info_url: pulumi.Output[str]
    """
    User Info URL
    """
    validate_signature: pulumi.Output[bool]
    """
    Enable/disable signature validation of external IDP signatures.
    """
    def __init__(__self__, resource_name, opts=None, accepts_prompt_none_forward_from_client=None, add_read_token_role_on_create=None, alias=None, authenticate_by_default=None, authorization_url=None, backchannel_supported=None, client_id=None, client_secret=None, default_scopes=None, display_name=None, enabled=None, extra_config=None, first_broker_login_flow_alias=None, hide_on_login_page=None, jwks_url=None, link_only=None, login_hint=None, logout_url=None, post_broker_login_flow_alias=None, provider_id=None, realm=None, store_token=None, token_url=None, trust_email=None, ui_locales=None, user_info_url=None, validate_signature=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a IdentityProvider resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accepts_prompt_none_forward_from_client: This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In
               case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly
               returned to client, but the request with prompt=none will be forwarded to this identity provider.
        :param pulumi.Input[bool] add_read_token_role_on_create: Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
        :param pulumi.Input[str] alias: The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
        :param pulumi.Input[bool] authenticate_by_default: Enable/disable authenticate users by default.
        :param pulumi.Input[str] authorization_url: OIDC authorization URL.
        :param pulumi.Input[bool] backchannel_supported: Does the external IDP support backchannel logout?
        :param pulumi.Input[str] client_id: Client ID.
        :param pulumi.Input[str] client_secret: Client Secret.
        :param pulumi.Input[str] default_scopes: The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to 'openid'.
        :param pulumi.Input[str] display_name: Friendly name for Identity Providers.
        :param pulumi.Input[bool] enabled: Enable/disable this identity provider.
        :param pulumi.Input[str] first_broker_login_flow_alias: Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
               that there is not yet existing Keycloak account linked with the authenticated identity provider account.
        :param pulumi.Input[bool] hide_on_login_page: Hide On Login Page.
        :param pulumi.Input[str] jwks_url: JSON Web Key Set URL
        :param pulumi.Input[bool] link_only: If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
               want to allow login from the provider, but want to integrate with a provider
        :param pulumi.Input[str] login_hint: Login Hint.
        :param pulumi.Input[str] logout_url: Logout URL
        :param pulumi.Input[str] post_broker_login_flow_alias: Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
               additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
               you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
               authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
        :param pulumi.Input[str] provider_id: provider id, is always oidc, unless you have a custom implementation
        :param pulumi.Input[str] realm: Realm Name
        :param pulumi.Input[bool] store_token: Enable/disable if tokens must be stored after authenticating users.
        :param pulumi.Input[str] token_url: Token URL.
        :param pulumi.Input[bool] trust_email: If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
        :param pulumi.Input[bool] ui_locales: Pass current locale to identity provider
        :param pulumi.Input[str] user_info_url: User Info URL
        :param pulumi.Input[bool] validate_signature: Enable/disable signature validation of external IDP signatures.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['accepts_prompt_none_forward_from_client'] = accepts_prompt_none_forward_from_client
            __props__['add_read_token_role_on_create'] = add_read_token_role_on_create
            if alias is None:
                raise TypeError("Missing required property 'alias'")
            __props__['alias'] = alias
            __props__['authenticate_by_default'] = authenticate_by_default
            if authorization_url is None:
                raise TypeError("Missing required property 'authorization_url'")
            __props__['authorization_url'] = authorization_url
            __props__['backchannel_supported'] = backchannel_supported
            if client_id is None:
                raise TypeError("Missing required property 'client_id'")
            __props__['client_id'] = client_id
            if client_secret is None:
                raise TypeError("Missing required property 'client_secret'")
            __props__['client_secret'] = client_secret
            __props__['default_scopes'] = default_scopes
            __props__['display_name'] = display_name
            __props__['enabled'] = enabled
            __props__['extra_config'] = extra_config
            __props__['first_broker_login_flow_alias'] = first_broker_login_flow_alias
            __props__['hide_on_login_page'] = hide_on_login_page
            __props__['jwks_url'] = jwks_url
            __props__['link_only'] = link_only
            __props__['login_hint'] = login_hint
            __props__['logout_url'] = logout_url
            __props__['post_broker_login_flow_alias'] = post_broker_login_flow_alias
            __props__['provider_id'] = provider_id
            if realm is None:
                raise TypeError("Missing required property 'realm'")
            __props__['realm'] = realm
            __props__['store_token'] = store_token
            if token_url is None:
                raise TypeError("Missing required property 'token_url'")
            __props__['token_url'] = token_url
            __props__['trust_email'] = trust_email
            __props__['ui_locales'] = ui_locales
            __props__['user_info_url'] = user_info_url
            __props__['validate_signature'] = validate_signature
            __props__['internal_id'] = None
        super(IdentityProvider, __self__).__init__(
            'keycloak:oidc/identityProvider:IdentityProvider',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, accepts_prompt_none_forward_from_client=None, add_read_token_role_on_create=None, alias=None, authenticate_by_default=None, authorization_url=None, backchannel_supported=None, client_id=None, client_secret=None, default_scopes=None, display_name=None, enabled=None, extra_config=None, first_broker_login_flow_alias=None, hide_on_login_page=None, internal_id=None, jwks_url=None, link_only=None, login_hint=None, logout_url=None, post_broker_login_flow_alias=None, provider_id=None, realm=None, store_token=None, token_url=None, trust_email=None, ui_locales=None, user_info_url=None, validate_signature=None):
        """
        Get an existing IdentityProvider resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] accepts_prompt_none_forward_from_client: This is just used together with Identity Provider Authenticator or when kc_idp_hint points to this identity provider. In
               case that client sends a request with prompt=none and user is not yet authenticated, the error will not be directly
               returned to client, but the request with prompt=none will be forwarded to this identity provider.
        :param pulumi.Input[bool] add_read_token_role_on_create: Enable/disable if new users can read any stored tokens. This assigns the broker.read-token role.
        :param pulumi.Input[str] alias: The alias uniquely identifies an identity provider and it is also used to build the redirect uri.
        :param pulumi.Input[bool] authenticate_by_default: Enable/disable authenticate users by default.
        :param pulumi.Input[str] authorization_url: OIDC authorization URL.
        :param pulumi.Input[bool] backchannel_supported: Does the external IDP support backchannel logout?
        :param pulumi.Input[str] client_id: Client ID.
        :param pulumi.Input[str] client_secret: Client Secret.
        :param pulumi.Input[str] default_scopes: The scopes to be sent when asking for authorization. It can be a space-separated list of scopes. Defaults to 'openid'.
        :param pulumi.Input[str] display_name: Friendly name for Identity Providers.
        :param pulumi.Input[bool] enabled: Enable/disable this identity provider.
        :param pulumi.Input[str] first_broker_login_flow_alias: Alias of authentication flow, which is triggered after first login with this identity provider. Term 'First Login' means
               that there is not yet existing Keycloak account linked with the authenticated identity provider account.
        :param pulumi.Input[bool] hide_on_login_page: Hide On Login Page.
        :param pulumi.Input[str] internal_id: Internal Identity Provider Id
        :param pulumi.Input[str] jwks_url: JSON Web Key Set URL
        :param pulumi.Input[bool] link_only: If true, users cannot log in through this provider. They can only link to this provider. This is useful if you don't
               want to allow login from the provider, but want to integrate with a provider
        :param pulumi.Input[str] login_hint: Login Hint.
        :param pulumi.Input[str] logout_url: Logout URL
        :param pulumi.Input[str] post_broker_login_flow_alias: Alias of authentication flow, which is triggered after each login with this identity provider. Useful if you want
               additional verification of each user authenticated with this identity provider (for example OTP). Leave this empty if
               you don't want any additional authenticators to be triggered after login with this identity provider. Also note, that
               authenticator implementations must assume that user is already set in ClientSession as identity provider already set it.
        :param pulumi.Input[str] provider_id: provider id, is always oidc, unless you have a custom implementation
        :param pulumi.Input[str] realm: Realm Name
        :param pulumi.Input[bool] store_token: Enable/disable if tokens must be stored after authenticating users.
        :param pulumi.Input[str] token_url: Token URL.
        :param pulumi.Input[bool] trust_email: If enabled then email provided by this provider is not verified even if verification is enabled for the realm.
        :param pulumi.Input[bool] ui_locales: Pass current locale to identity provider
        :param pulumi.Input[str] user_info_url: User Info URL
        :param pulumi.Input[bool] validate_signature: Enable/disable signature validation of external IDP signatures.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accepts_prompt_none_forward_from_client"] = accepts_prompt_none_forward_from_client
        __props__["add_read_token_role_on_create"] = add_read_token_role_on_create
        __props__["alias"] = alias
        __props__["authenticate_by_default"] = authenticate_by_default
        __props__["authorization_url"] = authorization_url
        __props__["backchannel_supported"] = backchannel_supported
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["default_scopes"] = default_scopes
        __props__["display_name"] = display_name
        __props__["enabled"] = enabled
        __props__["extra_config"] = extra_config
        __props__["first_broker_login_flow_alias"] = first_broker_login_flow_alias
        __props__["hide_on_login_page"] = hide_on_login_page
        __props__["internal_id"] = internal_id
        __props__["jwks_url"] = jwks_url
        __props__["link_only"] = link_only
        __props__["login_hint"] = login_hint
        __props__["logout_url"] = logout_url
        __props__["post_broker_login_flow_alias"] = post_broker_login_flow_alias
        __props__["provider_id"] = provider_id
        __props__["realm"] = realm
        __props__["store_token"] = store_token
        __props__["token_url"] = token_url
        __props__["trust_email"] = trust_email
        __props__["ui_locales"] = ui_locales
        __props__["user_info_url"] = user_info_url
        __props__["validate_signature"] = validate_signature
        return IdentityProvider(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
