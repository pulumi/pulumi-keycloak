# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetClientResult',
    'AwaitableGetClientResult',
    'get_client',
]

@pulumi.output_type
class GetClientResult:
    """
    A collection of values returned by getClient.
    """
    def __init__(__self__, access_token_lifespan=None, access_type=None, admin_url=None, authentication_flow_binding_overrides=None, authorizations=None, base_url=None, client_id=None, client_secret=None, consent_required=None, description=None, direct_access_grants_enabled=None, enabled=None, exclude_session_state_from_auth_response=None, full_scope_allowed=None, id=None, implicit_flow_enabled=None, login_theme=None, name=None, pkce_code_challenge_method=None, realm_id=None, resource_server_id=None, root_url=None, service_account_user_id=None, service_accounts_enabled=None, standard_flow_enabled=None, valid_redirect_uris=None, web_origins=None):
        if access_token_lifespan and not isinstance(access_token_lifespan, str):
            raise TypeError("Expected argument 'access_token_lifespan' to be a str")
        pulumi.set(__self__, "access_token_lifespan", access_token_lifespan)
        if access_type and not isinstance(access_type, str):
            raise TypeError("Expected argument 'access_type' to be a str")
        pulumi.set(__self__, "access_type", access_type)
        if admin_url and not isinstance(admin_url, str):
            raise TypeError("Expected argument 'admin_url' to be a str")
        pulumi.set(__self__, "admin_url", admin_url)
        if authentication_flow_binding_overrides and not isinstance(authentication_flow_binding_overrides, list):
            raise TypeError("Expected argument 'authentication_flow_binding_overrides' to be a list")
        pulumi.set(__self__, "authentication_flow_binding_overrides", authentication_flow_binding_overrides)
        if authorizations and not isinstance(authorizations, list):
            raise TypeError("Expected argument 'authorizations' to be a list")
        pulumi.set(__self__, "authorizations", authorizations)
        if base_url and not isinstance(base_url, str):
            raise TypeError("Expected argument 'base_url' to be a str")
        pulumi.set(__self__, "base_url", base_url)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if consent_required and not isinstance(consent_required, bool):
            raise TypeError("Expected argument 'consent_required' to be a bool")
        pulumi.set(__self__, "consent_required", consent_required)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if direct_access_grants_enabled and not isinstance(direct_access_grants_enabled, bool):
            raise TypeError("Expected argument 'direct_access_grants_enabled' to be a bool")
        pulumi.set(__self__, "direct_access_grants_enabled", direct_access_grants_enabled)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if exclude_session_state_from_auth_response and not isinstance(exclude_session_state_from_auth_response, bool):
            raise TypeError("Expected argument 'exclude_session_state_from_auth_response' to be a bool")
        pulumi.set(__self__, "exclude_session_state_from_auth_response", exclude_session_state_from_auth_response)
        if full_scope_allowed and not isinstance(full_scope_allowed, bool):
            raise TypeError("Expected argument 'full_scope_allowed' to be a bool")
        pulumi.set(__self__, "full_scope_allowed", full_scope_allowed)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if implicit_flow_enabled and not isinstance(implicit_flow_enabled, bool):
            raise TypeError("Expected argument 'implicit_flow_enabled' to be a bool")
        pulumi.set(__self__, "implicit_flow_enabled", implicit_flow_enabled)
        if login_theme and not isinstance(login_theme, str):
            raise TypeError("Expected argument 'login_theme' to be a str")
        pulumi.set(__self__, "login_theme", login_theme)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if pkce_code_challenge_method and not isinstance(pkce_code_challenge_method, str):
            raise TypeError("Expected argument 'pkce_code_challenge_method' to be a str")
        pulumi.set(__self__, "pkce_code_challenge_method", pkce_code_challenge_method)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if resource_server_id and not isinstance(resource_server_id, str):
            raise TypeError("Expected argument 'resource_server_id' to be a str")
        pulumi.set(__self__, "resource_server_id", resource_server_id)
        if root_url and not isinstance(root_url, str):
            raise TypeError("Expected argument 'root_url' to be a str")
        pulumi.set(__self__, "root_url", root_url)
        if service_account_user_id and not isinstance(service_account_user_id, str):
            raise TypeError("Expected argument 'service_account_user_id' to be a str")
        pulumi.set(__self__, "service_account_user_id", service_account_user_id)
        if service_accounts_enabled and not isinstance(service_accounts_enabled, bool):
            raise TypeError("Expected argument 'service_accounts_enabled' to be a bool")
        pulumi.set(__self__, "service_accounts_enabled", service_accounts_enabled)
        if standard_flow_enabled and not isinstance(standard_flow_enabled, bool):
            raise TypeError("Expected argument 'standard_flow_enabled' to be a bool")
        pulumi.set(__self__, "standard_flow_enabled", standard_flow_enabled)
        if valid_redirect_uris and not isinstance(valid_redirect_uris, list):
            raise TypeError("Expected argument 'valid_redirect_uris' to be a list")
        pulumi.set(__self__, "valid_redirect_uris", valid_redirect_uris)
        if web_origins and not isinstance(web_origins, list):
            raise TypeError("Expected argument 'web_origins' to be a list")
        pulumi.set(__self__, "web_origins", web_origins)

    @property
    @pulumi.getter(name="accessTokenLifespan")
    def access_token_lifespan(self) -> str:
        return pulumi.get(self, "access_token_lifespan")

    @property
    @pulumi.getter(name="accessType")
    def access_type(self) -> str:
        return pulumi.get(self, "access_type")

    @property
    @pulumi.getter(name="adminUrl")
    def admin_url(self) -> str:
        return pulumi.get(self, "admin_url")

    @property
    @pulumi.getter(name="authenticationFlowBindingOverrides")
    def authentication_flow_binding_overrides(self) -> Sequence['outputs.GetClientAuthenticationFlowBindingOverrideResult']:
        return pulumi.get(self, "authentication_flow_binding_overrides")

    @property
    @pulumi.getter
    def authorizations(self) -> Sequence['outputs.GetClientAuthorizationResult']:
        return pulumi.get(self, "authorizations")

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> str:
        return pulumi.get(self, "base_url")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> str:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter(name="consentRequired")
    def consent_required(self) -> bool:
        return pulumi.get(self, "consent_required")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="directAccessGrantsEnabled")
    def direct_access_grants_enabled(self) -> bool:
        return pulumi.get(self, "direct_access_grants_enabled")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="excludeSessionStateFromAuthResponse")
    def exclude_session_state_from_auth_response(self) -> bool:
        return pulumi.get(self, "exclude_session_state_from_auth_response")

    @property
    @pulumi.getter(name="fullScopeAllowed")
    def full_scope_allowed(self) -> bool:
        return pulumi.get(self, "full_scope_allowed")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="implicitFlowEnabled")
    def implicit_flow_enabled(self) -> bool:
        return pulumi.get(self, "implicit_flow_enabled")

    @property
    @pulumi.getter(name="loginTheme")
    def login_theme(self) -> str:
        return pulumi.get(self, "login_theme")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pkceCodeChallengeMethod")
    def pkce_code_challenge_method(self) -> str:
        return pulumi.get(self, "pkce_code_challenge_method")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> str:
        return pulumi.get(self, "resource_server_id")

    @property
    @pulumi.getter(name="rootUrl")
    def root_url(self) -> str:
        return pulumi.get(self, "root_url")

    @property
    @pulumi.getter(name="serviceAccountUserId")
    def service_account_user_id(self) -> str:
        return pulumi.get(self, "service_account_user_id")

    @property
    @pulumi.getter(name="serviceAccountsEnabled")
    def service_accounts_enabled(self) -> bool:
        return pulumi.get(self, "service_accounts_enabled")

    @property
    @pulumi.getter(name="standardFlowEnabled")
    def standard_flow_enabled(self) -> bool:
        return pulumi.get(self, "standard_flow_enabled")

    @property
    @pulumi.getter(name="validRedirectUris")
    def valid_redirect_uris(self) -> Sequence[str]:
        return pulumi.get(self, "valid_redirect_uris")

    @property
    @pulumi.getter(name="webOrigins")
    def web_origins(self) -> Sequence[str]:
        return pulumi.get(self, "web_origins")


class AwaitableGetClientResult(GetClientResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientResult(
            access_token_lifespan=self.access_token_lifespan,
            access_type=self.access_type,
            admin_url=self.admin_url,
            authentication_flow_binding_overrides=self.authentication_flow_binding_overrides,
            authorizations=self.authorizations,
            base_url=self.base_url,
            client_id=self.client_id,
            client_secret=self.client_secret,
            consent_required=self.consent_required,
            description=self.description,
            direct_access_grants_enabled=self.direct_access_grants_enabled,
            enabled=self.enabled,
            exclude_session_state_from_auth_response=self.exclude_session_state_from_auth_response,
            full_scope_allowed=self.full_scope_allowed,
            id=self.id,
            implicit_flow_enabled=self.implicit_flow_enabled,
            login_theme=self.login_theme,
            name=self.name,
            pkce_code_challenge_method=self.pkce_code_challenge_method,
            realm_id=self.realm_id,
            resource_server_id=self.resource_server_id,
            root_url=self.root_url,
            service_account_user_id=self.service_account_user_id,
            service_accounts_enabled=self.service_accounts_enabled,
            standard_flow_enabled=self.standard_flow_enabled,
            valid_redirect_uris=self.valid_redirect_uris,
            web_origins=self.web_origins)


def get_client(client_id: Optional[str] = None,
               realm_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientResult:
    """
    This data source can be used to fetch properties of a Keycloak OpenID client for usage with other resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm_management = keycloak.openid.get_client(realm_id="my-realm",
        client_id="realm-management")
    admin = keycloak.get_role(realm_id="my-realm",
        client_id=realm_management.id,
        name="realm-admin")
    ```


    :param str client_id: The client id (not its unique ID).
    :param str realm_id: The realm id.
    """
    __args__ = dict()
    __args__['clientId'] = client_id
    __args__['realmId'] = realm_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('keycloak:openid/getClient:getClient', __args__, opts=opts, typ=GetClientResult).value

    return AwaitableGetClientResult(
        access_token_lifespan=__ret__.access_token_lifespan,
        access_type=__ret__.access_type,
        admin_url=__ret__.admin_url,
        authentication_flow_binding_overrides=__ret__.authentication_flow_binding_overrides,
        authorizations=__ret__.authorizations,
        base_url=__ret__.base_url,
        client_id=__ret__.client_id,
        client_secret=__ret__.client_secret,
        consent_required=__ret__.consent_required,
        description=__ret__.description,
        direct_access_grants_enabled=__ret__.direct_access_grants_enabled,
        enabled=__ret__.enabled,
        exclude_session_state_from_auth_response=__ret__.exclude_session_state_from_auth_response,
        full_scope_allowed=__ret__.full_scope_allowed,
        id=__ret__.id,
        implicit_flow_enabled=__ret__.implicit_flow_enabled,
        login_theme=__ret__.login_theme,
        name=__ret__.name,
        pkce_code_challenge_method=__ret__.pkce_code_challenge_method,
        realm_id=__ret__.realm_id,
        resource_server_id=__ret__.resource_server_id,
        root_url=__ret__.root_url,
        service_account_user_id=__ret__.service_account_user_id,
        service_accounts_enabled=__ret__.service_accounts_enabled,
        standard_flow_enabled=__ret__.standard_flow_enabled,
        valid_redirect_uris=__ret__.valid_redirect_uris,
        web_origins=__ret__.web_origins)
