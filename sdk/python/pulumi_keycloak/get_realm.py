# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = [
    'GetRealmResult',
    'AwaitableGetRealmResult',
    'get_realm',
]

@pulumi.output_type
class GetRealmResult:
    """
    A collection of values returned by getRealm.
    """
    def __init__(__self__, access_code_lifespan=None, access_code_lifespan_login=None, access_code_lifespan_user_action=None, access_token_lifespan=None, access_token_lifespan_for_implicit_flow=None, account_theme=None, action_token_generated_by_admin_lifespan=None, action_token_generated_by_user_lifespan=None, admin_theme=None, attributes=None, browser_flow=None, client_authentication_flow=None, direct_grant_flow=None, display_name=None, display_name_html=None, docker_authentication_flow=None, duplicate_emails_allowed=None, edit_username_allowed=None, email_theme=None, enabled=None, id=None, internal_id=None, internationalizations=None, login_theme=None, login_with_email_allowed=None, offline_session_idle_timeout=None, offline_session_max_lifespan=None, password_policy=None, realm=None, refresh_token_max_reuse=None, registration_allowed=None, registration_email_as_username=None, registration_flow=None, remember_me=None, reset_credentials_flow=None, reset_password_allowed=None, security_defenses=None, smtp_servers=None, ssl_required=None, sso_session_idle_timeout=None, sso_session_max_lifespan=None, user_managed_access=None, verify_email=None):
        if access_code_lifespan and not isinstance(access_code_lifespan, str):
            raise TypeError("Expected argument 'access_code_lifespan' to be a str")
        pulumi.set(__self__, "access_code_lifespan", access_code_lifespan)
        if access_code_lifespan_login and not isinstance(access_code_lifespan_login, str):
            raise TypeError("Expected argument 'access_code_lifespan_login' to be a str")
        pulumi.set(__self__, "access_code_lifespan_login", access_code_lifespan_login)
        if access_code_lifespan_user_action and not isinstance(access_code_lifespan_user_action, str):
            raise TypeError("Expected argument 'access_code_lifespan_user_action' to be a str")
        pulumi.set(__self__, "access_code_lifespan_user_action", access_code_lifespan_user_action)
        if access_token_lifespan and not isinstance(access_token_lifespan, str):
            raise TypeError("Expected argument 'access_token_lifespan' to be a str")
        pulumi.set(__self__, "access_token_lifespan", access_token_lifespan)
        if access_token_lifespan_for_implicit_flow and not isinstance(access_token_lifespan_for_implicit_flow, str):
            raise TypeError("Expected argument 'access_token_lifespan_for_implicit_flow' to be a str")
        pulumi.set(__self__, "access_token_lifespan_for_implicit_flow", access_token_lifespan_for_implicit_flow)
        if account_theme and not isinstance(account_theme, str):
            raise TypeError("Expected argument 'account_theme' to be a str")
        pulumi.set(__self__, "account_theme", account_theme)
        if action_token_generated_by_admin_lifespan and not isinstance(action_token_generated_by_admin_lifespan, str):
            raise TypeError("Expected argument 'action_token_generated_by_admin_lifespan' to be a str")
        pulumi.set(__self__, "action_token_generated_by_admin_lifespan", action_token_generated_by_admin_lifespan)
        if action_token_generated_by_user_lifespan and not isinstance(action_token_generated_by_user_lifespan, str):
            raise TypeError("Expected argument 'action_token_generated_by_user_lifespan' to be a str")
        pulumi.set(__self__, "action_token_generated_by_user_lifespan", action_token_generated_by_user_lifespan)
        if admin_theme and not isinstance(admin_theme, str):
            raise TypeError("Expected argument 'admin_theme' to be a str")
        pulumi.set(__self__, "admin_theme", admin_theme)
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if browser_flow and not isinstance(browser_flow, str):
            raise TypeError("Expected argument 'browser_flow' to be a str")
        pulumi.set(__self__, "browser_flow", browser_flow)
        if client_authentication_flow and not isinstance(client_authentication_flow, str):
            raise TypeError("Expected argument 'client_authentication_flow' to be a str")
        pulumi.set(__self__, "client_authentication_flow", client_authentication_flow)
        if direct_grant_flow and not isinstance(direct_grant_flow, str):
            raise TypeError("Expected argument 'direct_grant_flow' to be a str")
        pulumi.set(__self__, "direct_grant_flow", direct_grant_flow)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if display_name_html and not isinstance(display_name_html, str):
            raise TypeError("Expected argument 'display_name_html' to be a str")
        pulumi.set(__self__, "display_name_html", display_name_html)
        if docker_authentication_flow and not isinstance(docker_authentication_flow, str):
            raise TypeError("Expected argument 'docker_authentication_flow' to be a str")
        pulumi.set(__self__, "docker_authentication_flow", docker_authentication_flow)
        if duplicate_emails_allowed and not isinstance(duplicate_emails_allowed, bool):
            raise TypeError("Expected argument 'duplicate_emails_allowed' to be a bool")
        pulumi.set(__self__, "duplicate_emails_allowed", duplicate_emails_allowed)
        if edit_username_allowed and not isinstance(edit_username_allowed, bool):
            raise TypeError("Expected argument 'edit_username_allowed' to be a bool")
        pulumi.set(__self__, "edit_username_allowed", edit_username_allowed)
        if email_theme and not isinstance(email_theme, str):
            raise TypeError("Expected argument 'email_theme' to be a str")
        pulumi.set(__self__, "email_theme", email_theme)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if internal_id and not isinstance(internal_id, str):
            raise TypeError("Expected argument 'internal_id' to be a str")
        pulumi.set(__self__, "internal_id", internal_id)
        if internationalizations and not isinstance(internationalizations, list):
            raise TypeError("Expected argument 'internationalizations' to be a list")
        pulumi.set(__self__, "internationalizations", internationalizations)
        if login_theme and not isinstance(login_theme, str):
            raise TypeError("Expected argument 'login_theme' to be a str")
        pulumi.set(__self__, "login_theme", login_theme)
        if login_with_email_allowed and not isinstance(login_with_email_allowed, bool):
            raise TypeError("Expected argument 'login_with_email_allowed' to be a bool")
        pulumi.set(__self__, "login_with_email_allowed", login_with_email_allowed)
        if offline_session_idle_timeout and not isinstance(offline_session_idle_timeout, str):
            raise TypeError("Expected argument 'offline_session_idle_timeout' to be a str")
        pulumi.set(__self__, "offline_session_idle_timeout", offline_session_idle_timeout)
        if offline_session_max_lifespan and not isinstance(offline_session_max_lifespan, str):
            raise TypeError("Expected argument 'offline_session_max_lifespan' to be a str")
        pulumi.set(__self__, "offline_session_max_lifespan", offline_session_max_lifespan)
        if password_policy and not isinstance(password_policy, str):
            raise TypeError("Expected argument 'password_policy' to be a str")
        pulumi.set(__self__, "password_policy", password_policy)
        if realm and not isinstance(realm, str):
            raise TypeError("Expected argument 'realm' to be a str")
        pulumi.set(__self__, "realm", realm)
        if refresh_token_max_reuse and not isinstance(refresh_token_max_reuse, float):
            raise TypeError("Expected argument 'refresh_token_max_reuse' to be a float")
        pulumi.set(__self__, "refresh_token_max_reuse", refresh_token_max_reuse)
        if registration_allowed and not isinstance(registration_allowed, bool):
            raise TypeError("Expected argument 'registration_allowed' to be a bool")
        pulumi.set(__self__, "registration_allowed", registration_allowed)
        if registration_email_as_username and not isinstance(registration_email_as_username, bool):
            raise TypeError("Expected argument 'registration_email_as_username' to be a bool")
        pulumi.set(__self__, "registration_email_as_username", registration_email_as_username)
        if registration_flow and not isinstance(registration_flow, str):
            raise TypeError("Expected argument 'registration_flow' to be a str")
        pulumi.set(__self__, "registration_flow", registration_flow)
        if remember_me and not isinstance(remember_me, bool):
            raise TypeError("Expected argument 'remember_me' to be a bool")
        pulumi.set(__self__, "remember_me", remember_me)
        if reset_credentials_flow and not isinstance(reset_credentials_flow, str):
            raise TypeError("Expected argument 'reset_credentials_flow' to be a str")
        pulumi.set(__self__, "reset_credentials_flow", reset_credentials_flow)
        if reset_password_allowed and not isinstance(reset_password_allowed, bool):
            raise TypeError("Expected argument 'reset_password_allowed' to be a bool")
        pulumi.set(__self__, "reset_password_allowed", reset_password_allowed)
        if security_defenses and not isinstance(security_defenses, list):
            raise TypeError("Expected argument 'security_defenses' to be a list")
        pulumi.set(__self__, "security_defenses", security_defenses)
        if smtp_servers and not isinstance(smtp_servers, list):
            raise TypeError("Expected argument 'smtp_servers' to be a list")
        pulumi.set(__self__, "smtp_servers", smtp_servers)
        if ssl_required and not isinstance(ssl_required, str):
            raise TypeError("Expected argument 'ssl_required' to be a str")
        pulumi.set(__self__, "ssl_required", ssl_required)
        if sso_session_idle_timeout and not isinstance(sso_session_idle_timeout, str):
            raise TypeError("Expected argument 'sso_session_idle_timeout' to be a str")
        pulumi.set(__self__, "sso_session_idle_timeout", sso_session_idle_timeout)
        if sso_session_max_lifespan and not isinstance(sso_session_max_lifespan, str):
            raise TypeError("Expected argument 'sso_session_max_lifespan' to be a str")
        pulumi.set(__self__, "sso_session_max_lifespan", sso_session_max_lifespan)
        if user_managed_access and not isinstance(user_managed_access, bool):
            raise TypeError("Expected argument 'user_managed_access' to be a bool")
        pulumi.set(__self__, "user_managed_access", user_managed_access)
        if verify_email and not isinstance(verify_email, bool):
            raise TypeError("Expected argument 'verify_email' to be a bool")
        pulumi.set(__self__, "verify_email", verify_email)

    @property
    @pulumi.getter(name="accessCodeLifespan")
    def access_code_lifespan(self) -> str:
        return pulumi.get(self, "access_code_lifespan")

    @property
    @pulumi.getter(name="accessCodeLifespanLogin")
    def access_code_lifespan_login(self) -> str:
        return pulumi.get(self, "access_code_lifespan_login")

    @property
    @pulumi.getter(name="accessCodeLifespanUserAction")
    def access_code_lifespan_user_action(self) -> str:
        return pulumi.get(self, "access_code_lifespan_user_action")

    @property
    @pulumi.getter(name="accessTokenLifespan")
    def access_token_lifespan(self) -> str:
        return pulumi.get(self, "access_token_lifespan")

    @property
    @pulumi.getter(name="accessTokenLifespanForImplicitFlow")
    def access_token_lifespan_for_implicit_flow(self) -> str:
        return pulumi.get(self, "access_token_lifespan_for_implicit_flow")

    @property
    @pulumi.getter(name="accountTheme")
    def account_theme(self) -> str:
        return pulumi.get(self, "account_theme")

    @property
    @pulumi.getter(name="actionTokenGeneratedByAdminLifespan")
    def action_token_generated_by_admin_lifespan(self) -> str:
        return pulumi.get(self, "action_token_generated_by_admin_lifespan")

    @property
    @pulumi.getter(name="actionTokenGeneratedByUserLifespan")
    def action_token_generated_by_user_lifespan(self) -> str:
        return pulumi.get(self, "action_token_generated_by_user_lifespan")

    @property
    @pulumi.getter(name="adminTheme")
    def admin_theme(self) -> str:
        return pulumi.get(self, "admin_theme")

    @property
    @pulumi.getter
    def attributes(self) -> Mapping[str, Any]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="browserFlow")
    def browser_flow(self) -> str:
        return pulumi.get(self, "browser_flow")

    @property
    @pulumi.getter(name="clientAuthenticationFlow")
    def client_authentication_flow(self) -> str:
        return pulumi.get(self, "client_authentication_flow")

    @property
    @pulumi.getter(name="directGrantFlow")
    def direct_grant_flow(self) -> str:
        return pulumi.get(self, "direct_grant_flow")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="displayNameHtml")
    def display_name_html(self) -> Optional[str]:
        return pulumi.get(self, "display_name_html")

    @property
    @pulumi.getter(name="dockerAuthenticationFlow")
    def docker_authentication_flow(self) -> str:
        return pulumi.get(self, "docker_authentication_flow")

    @property
    @pulumi.getter(name="duplicateEmailsAllowed")
    def duplicate_emails_allowed(self) -> bool:
        return pulumi.get(self, "duplicate_emails_allowed")

    @property
    @pulumi.getter(name="editUsernameAllowed")
    def edit_username_allowed(self) -> bool:
        return pulumi.get(self, "edit_username_allowed")

    @property
    @pulumi.getter(name="emailTheme")
    def email_theme(self) -> str:
        return pulumi.get(self, "email_theme")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="internalId")
    def internal_id(self) -> str:
        return pulumi.get(self, "internal_id")

    @property
    @pulumi.getter
    def internationalizations(self) -> List['outputs.GetRealmInternationalizationResult']:
        return pulumi.get(self, "internationalizations")

    @property
    @pulumi.getter(name="loginTheme")
    def login_theme(self) -> str:
        return pulumi.get(self, "login_theme")

    @property
    @pulumi.getter(name="loginWithEmailAllowed")
    def login_with_email_allowed(self) -> bool:
        return pulumi.get(self, "login_with_email_allowed")

    @property
    @pulumi.getter(name="offlineSessionIdleTimeout")
    def offline_session_idle_timeout(self) -> str:
        return pulumi.get(self, "offline_session_idle_timeout")

    @property
    @pulumi.getter(name="offlineSessionMaxLifespan")
    def offline_session_max_lifespan(self) -> str:
        return pulumi.get(self, "offline_session_max_lifespan")

    @property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> str:
        return pulumi.get(self, "password_policy")

    @property
    @pulumi.getter
    def realm(self) -> str:
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter(name="refreshTokenMaxReuse")
    def refresh_token_max_reuse(self) -> float:
        return pulumi.get(self, "refresh_token_max_reuse")

    @property
    @pulumi.getter(name="registrationAllowed")
    def registration_allowed(self) -> bool:
        return pulumi.get(self, "registration_allowed")

    @property
    @pulumi.getter(name="registrationEmailAsUsername")
    def registration_email_as_username(self) -> bool:
        return pulumi.get(self, "registration_email_as_username")

    @property
    @pulumi.getter(name="registrationFlow")
    def registration_flow(self) -> str:
        return pulumi.get(self, "registration_flow")

    @property
    @pulumi.getter(name="rememberMe")
    def remember_me(self) -> bool:
        return pulumi.get(self, "remember_me")

    @property
    @pulumi.getter(name="resetCredentialsFlow")
    def reset_credentials_flow(self) -> str:
        return pulumi.get(self, "reset_credentials_flow")

    @property
    @pulumi.getter(name="resetPasswordAllowed")
    def reset_password_allowed(self) -> bool:
        return pulumi.get(self, "reset_password_allowed")

    @property
    @pulumi.getter(name="securityDefenses")
    def security_defenses(self) -> List['outputs.GetRealmSecurityDefenseResult']:
        return pulumi.get(self, "security_defenses")

    @property
    @pulumi.getter(name="smtpServers")
    def smtp_servers(self) -> List['outputs.GetRealmSmtpServerResult']:
        return pulumi.get(self, "smtp_servers")

    @property
    @pulumi.getter(name="sslRequired")
    def ssl_required(self) -> str:
        return pulumi.get(self, "ssl_required")

    @property
    @pulumi.getter(name="ssoSessionIdleTimeout")
    def sso_session_idle_timeout(self) -> str:
        return pulumi.get(self, "sso_session_idle_timeout")

    @property
    @pulumi.getter(name="ssoSessionMaxLifespan")
    def sso_session_max_lifespan(self) -> str:
        return pulumi.get(self, "sso_session_max_lifespan")

    @property
    @pulumi.getter(name="userManagedAccess")
    def user_managed_access(self) -> bool:
        return pulumi.get(self, "user_managed_access")

    @property
    @pulumi.getter(name="verifyEmail")
    def verify_email(self) -> bool:
        return pulumi.get(self, "verify_email")


class AwaitableGetRealmResult(GetRealmResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRealmResult(
            access_code_lifespan=self.access_code_lifespan,
            access_code_lifespan_login=self.access_code_lifespan_login,
            access_code_lifespan_user_action=self.access_code_lifespan_user_action,
            access_token_lifespan=self.access_token_lifespan,
            access_token_lifespan_for_implicit_flow=self.access_token_lifespan_for_implicit_flow,
            account_theme=self.account_theme,
            action_token_generated_by_admin_lifespan=self.action_token_generated_by_admin_lifespan,
            action_token_generated_by_user_lifespan=self.action_token_generated_by_user_lifespan,
            admin_theme=self.admin_theme,
            attributes=self.attributes,
            browser_flow=self.browser_flow,
            client_authentication_flow=self.client_authentication_flow,
            direct_grant_flow=self.direct_grant_flow,
            display_name=self.display_name,
            display_name_html=self.display_name_html,
            docker_authentication_flow=self.docker_authentication_flow,
            duplicate_emails_allowed=self.duplicate_emails_allowed,
            edit_username_allowed=self.edit_username_allowed,
            email_theme=self.email_theme,
            enabled=self.enabled,
            id=self.id,
            internal_id=self.internal_id,
            internationalizations=self.internationalizations,
            login_theme=self.login_theme,
            login_with_email_allowed=self.login_with_email_allowed,
            offline_session_idle_timeout=self.offline_session_idle_timeout,
            offline_session_max_lifespan=self.offline_session_max_lifespan,
            password_policy=self.password_policy,
            realm=self.realm,
            refresh_token_max_reuse=self.refresh_token_max_reuse,
            registration_allowed=self.registration_allowed,
            registration_email_as_username=self.registration_email_as_username,
            registration_flow=self.registration_flow,
            remember_me=self.remember_me,
            reset_credentials_flow=self.reset_credentials_flow,
            reset_password_allowed=self.reset_password_allowed,
            security_defenses=self.security_defenses,
            smtp_servers=self.smtp_servers,
            ssl_required=self.ssl_required,
            sso_session_idle_timeout=self.sso_session_idle_timeout,
            sso_session_max_lifespan=self.sso_session_max_lifespan,
            user_managed_access=self.user_managed_access,
            verify_email=self.verify_email)


def get_realm(attributes: Optional[Mapping[str, Any]] = None,
              display_name_html: Optional[str] = None,
              internationalizations: Optional[List[pulumi.InputType['GetRealmInternationalizationArgs']]] = None,
              realm: Optional[str] = None,
              security_defenses: Optional[List[pulumi.InputType['GetRealmSecurityDefenseArgs']]] = None,
              smtp_servers: Optional[List[pulumi.InputType['GetRealmSmtpServerArgs']]] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRealmResult:
    """
    ## # Realm data source

    This data source can be used to fetch properties of a Keycloak realm for
    usage with other resources.

    ### Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.get_realm(realm="my-realm")
    group = keycloak.Role("group", realm_id=data["keycloak_realm"]["id"])
    ```

    ### Argument Reference

    The following arguments are supported:

    - `realm` - (Required) The realm name.

    ### Attributes Reference

    See the docs for the `Realm` resource for details on the exported attributes.
    """
    __args__ = dict()
    __args__['attributes'] = attributes
    __args__['displayNameHtml'] = display_name_html
    __args__['internationalizations'] = internationalizations
    __args__['realm'] = realm
    __args__['securityDefenses'] = security_defenses
    __args__['smtpServers'] = smtp_servers
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('keycloak:index/getRealm:getRealm', __args__, opts=opts, typ=GetRealmResult).value

    return AwaitableGetRealmResult(
        access_code_lifespan=__ret__.access_code_lifespan,
        access_code_lifespan_login=__ret__.access_code_lifespan_login,
        access_code_lifespan_user_action=__ret__.access_code_lifespan_user_action,
        access_token_lifespan=__ret__.access_token_lifespan,
        access_token_lifespan_for_implicit_flow=__ret__.access_token_lifespan_for_implicit_flow,
        account_theme=__ret__.account_theme,
        action_token_generated_by_admin_lifespan=__ret__.action_token_generated_by_admin_lifespan,
        action_token_generated_by_user_lifespan=__ret__.action_token_generated_by_user_lifespan,
        admin_theme=__ret__.admin_theme,
        attributes=__ret__.attributes,
        browser_flow=__ret__.browser_flow,
        client_authentication_flow=__ret__.client_authentication_flow,
        direct_grant_flow=__ret__.direct_grant_flow,
        display_name=__ret__.display_name,
        display_name_html=__ret__.display_name_html,
        docker_authentication_flow=__ret__.docker_authentication_flow,
        duplicate_emails_allowed=__ret__.duplicate_emails_allowed,
        edit_username_allowed=__ret__.edit_username_allowed,
        email_theme=__ret__.email_theme,
        enabled=__ret__.enabled,
        id=__ret__.id,
        internal_id=__ret__.internal_id,
        internationalizations=__ret__.internationalizations,
        login_theme=__ret__.login_theme,
        login_with_email_allowed=__ret__.login_with_email_allowed,
        offline_session_idle_timeout=__ret__.offline_session_idle_timeout,
        offline_session_max_lifespan=__ret__.offline_session_max_lifespan,
        password_policy=__ret__.password_policy,
        realm=__ret__.realm,
        refresh_token_max_reuse=__ret__.refresh_token_max_reuse,
        registration_allowed=__ret__.registration_allowed,
        registration_email_as_username=__ret__.registration_email_as_username,
        registration_flow=__ret__.registration_flow,
        remember_me=__ret__.remember_me,
        reset_credentials_flow=__ret__.reset_credentials_flow,
        reset_password_allowed=__ret__.reset_password_allowed,
        security_defenses=__ret__.security_defenses,
        smtp_servers=__ret__.smtp_servers,
        ssl_required=__ret__.ssl_required,
        sso_session_idle_timeout=__ret__.sso_session_idle_timeout,
        sso_session_max_lifespan=__ret__.sso_session_max_lifespan,
        user_managed_access=__ret__.user_managed_access,
        verify_email=__ret__.verify_email)
