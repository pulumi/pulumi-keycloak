# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = [
    'GetRealmResult',
    'AwaitableGetRealmResult',
    'get_realm',
    'get_realm_output',
]

@pulumi.output_type
class GetRealmResult:
    """
    A collection of values returned by getRealm.
    """
    def __init__(__self__, access_code_lifespan=None, access_code_lifespan_login=None, access_code_lifespan_user_action=None, access_token_lifespan=None, access_token_lifespan_for_implicit_flow=None, account_theme=None, action_token_generated_by_admin_lifespan=None, action_token_generated_by_user_lifespan=None, admin_theme=None, attributes=None, browser_flow=None, client_authentication_flow=None, client_session_idle_timeout=None, client_session_max_lifespan=None, default_default_client_scopes=None, default_optional_client_scopes=None, default_signature_algorithm=None, direct_grant_flow=None, display_name=None, display_name_html=None, docker_authentication_flow=None, duplicate_emails_allowed=None, edit_username_allowed=None, email_theme=None, enabled=None, id=None, internal_id=None, internationalizations=None, login_theme=None, login_with_email_allowed=None, oauth2_device_code_lifespan=None, oauth2_device_polling_interval=None, offline_session_idle_timeout=None, offline_session_max_lifespan=None, offline_session_max_lifespan_enabled=None, otp_policy=None, password_policy=None, realm=None, refresh_token_max_reuse=None, registration_allowed=None, registration_email_as_username=None, registration_flow=None, remember_me=None, reset_credentials_flow=None, reset_password_allowed=None, revoke_refresh_token=None, security_defenses=None, smtp_servers=None, ssl_required=None, sso_session_idle_timeout=None, sso_session_idle_timeout_remember_me=None, sso_session_max_lifespan=None, sso_session_max_lifespan_remember_me=None, user_managed_access=None, verify_email=None, web_authn_passwordless_policy=None, web_authn_policy=None):
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
        if client_session_idle_timeout and not isinstance(client_session_idle_timeout, str):
            raise TypeError("Expected argument 'client_session_idle_timeout' to be a str")
        pulumi.set(__self__, "client_session_idle_timeout", client_session_idle_timeout)
        if client_session_max_lifespan and not isinstance(client_session_max_lifespan, str):
            raise TypeError("Expected argument 'client_session_max_lifespan' to be a str")
        pulumi.set(__self__, "client_session_max_lifespan", client_session_max_lifespan)
        if default_default_client_scopes and not isinstance(default_default_client_scopes, list):
            raise TypeError("Expected argument 'default_default_client_scopes' to be a list")
        pulumi.set(__self__, "default_default_client_scopes", default_default_client_scopes)
        if default_optional_client_scopes and not isinstance(default_optional_client_scopes, list):
            raise TypeError("Expected argument 'default_optional_client_scopes' to be a list")
        pulumi.set(__self__, "default_optional_client_scopes", default_optional_client_scopes)
        if default_signature_algorithm and not isinstance(default_signature_algorithm, str):
            raise TypeError("Expected argument 'default_signature_algorithm' to be a str")
        pulumi.set(__self__, "default_signature_algorithm", default_signature_algorithm)
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
        if oauth2_device_code_lifespan and not isinstance(oauth2_device_code_lifespan, str):
            raise TypeError("Expected argument 'oauth2_device_code_lifespan' to be a str")
        pulumi.set(__self__, "oauth2_device_code_lifespan", oauth2_device_code_lifespan)
        if oauth2_device_polling_interval and not isinstance(oauth2_device_polling_interval, int):
            raise TypeError("Expected argument 'oauth2_device_polling_interval' to be a int")
        pulumi.set(__self__, "oauth2_device_polling_interval", oauth2_device_polling_interval)
        if offline_session_idle_timeout and not isinstance(offline_session_idle_timeout, str):
            raise TypeError("Expected argument 'offline_session_idle_timeout' to be a str")
        pulumi.set(__self__, "offline_session_idle_timeout", offline_session_idle_timeout)
        if offline_session_max_lifespan and not isinstance(offline_session_max_lifespan, str):
            raise TypeError("Expected argument 'offline_session_max_lifespan' to be a str")
        pulumi.set(__self__, "offline_session_max_lifespan", offline_session_max_lifespan)
        if offline_session_max_lifespan_enabled and not isinstance(offline_session_max_lifespan_enabled, bool):
            raise TypeError("Expected argument 'offline_session_max_lifespan_enabled' to be a bool")
        pulumi.set(__self__, "offline_session_max_lifespan_enabled", offline_session_max_lifespan_enabled)
        if otp_policy and not isinstance(otp_policy, dict):
            raise TypeError("Expected argument 'otp_policy' to be a dict")
        pulumi.set(__self__, "otp_policy", otp_policy)
        if password_policy and not isinstance(password_policy, str):
            raise TypeError("Expected argument 'password_policy' to be a str")
        pulumi.set(__self__, "password_policy", password_policy)
        if realm and not isinstance(realm, str):
            raise TypeError("Expected argument 'realm' to be a str")
        pulumi.set(__self__, "realm", realm)
        if refresh_token_max_reuse and not isinstance(refresh_token_max_reuse, int):
            raise TypeError("Expected argument 'refresh_token_max_reuse' to be a int")
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
        if revoke_refresh_token and not isinstance(revoke_refresh_token, bool):
            raise TypeError("Expected argument 'revoke_refresh_token' to be a bool")
        pulumi.set(__self__, "revoke_refresh_token", revoke_refresh_token)
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
        if sso_session_idle_timeout_remember_me and not isinstance(sso_session_idle_timeout_remember_me, str):
            raise TypeError("Expected argument 'sso_session_idle_timeout_remember_me' to be a str")
        pulumi.set(__self__, "sso_session_idle_timeout_remember_me", sso_session_idle_timeout_remember_me)
        if sso_session_max_lifespan and not isinstance(sso_session_max_lifespan, str):
            raise TypeError("Expected argument 'sso_session_max_lifespan' to be a str")
        pulumi.set(__self__, "sso_session_max_lifespan", sso_session_max_lifespan)
        if sso_session_max_lifespan_remember_me and not isinstance(sso_session_max_lifespan_remember_me, str):
            raise TypeError("Expected argument 'sso_session_max_lifespan_remember_me' to be a str")
        pulumi.set(__self__, "sso_session_max_lifespan_remember_me", sso_session_max_lifespan_remember_me)
        if user_managed_access and not isinstance(user_managed_access, bool):
            raise TypeError("Expected argument 'user_managed_access' to be a bool")
        pulumi.set(__self__, "user_managed_access", user_managed_access)
        if verify_email and not isinstance(verify_email, bool):
            raise TypeError("Expected argument 'verify_email' to be a bool")
        pulumi.set(__self__, "verify_email", verify_email)
        if web_authn_passwordless_policy and not isinstance(web_authn_passwordless_policy, dict):
            raise TypeError("Expected argument 'web_authn_passwordless_policy' to be a dict")
        pulumi.set(__self__, "web_authn_passwordless_policy", web_authn_passwordless_policy)
        if web_authn_policy and not isinstance(web_authn_policy, dict):
            raise TypeError("Expected argument 'web_authn_policy' to be a dict")
        pulumi.set(__self__, "web_authn_policy", web_authn_policy)

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
    @pulumi.getter(name="clientSessionIdleTimeout")
    def client_session_idle_timeout(self) -> str:
        return pulumi.get(self, "client_session_idle_timeout")

    @property
    @pulumi.getter(name="clientSessionMaxLifespan")
    def client_session_max_lifespan(self) -> str:
        return pulumi.get(self, "client_session_max_lifespan")

    @property
    @pulumi.getter(name="defaultDefaultClientScopes")
    def default_default_client_scopes(self) -> Sequence[str]:
        return pulumi.get(self, "default_default_client_scopes")

    @property
    @pulumi.getter(name="defaultOptionalClientScopes")
    def default_optional_client_scopes(self) -> Sequence[str]:
        return pulumi.get(self, "default_optional_client_scopes")

    @property
    @pulumi.getter(name="defaultSignatureAlgorithm")
    def default_signature_algorithm(self) -> str:
        return pulumi.get(self, "default_signature_algorithm")

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
    def internationalizations(self) -> Sequence['outputs.GetRealmInternationalizationResult']:
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
    @pulumi.getter(name="oauth2DeviceCodeLifespan")
    def oauth2_device_code_lifespan(self) -> str:
        return pulumi.get(self, "oauth2_device_code_lifespan")

    @property
    @pulumi.getter(name="oauth2DevicePollingInterval")
    def oauth2_device_polling_interval(self) -> int:
        return pulumi.get(self, "oauth2_device_polling_interval")

    @property
    @pulumi.getter(name="offlineSessionIdleTimeout")
    def offline_session_idle_timeout(self) -> str:
        return pulumi.get(self, "offline_session_idle_timeout")

    @property
    @pulumi.getter(name="offlineSessionMaxLifespan")
    def offline_session_max_lifespan(self) -> str:
        return pulumi.get(self, "offline_session_max_lifespan")

    @property
    @pulumi.getter(name="offlineSessionMaxLifespanEnabled")
    def offline_session_max_lifespan_enabled(self) -> bool:
        return pulumi.get(self, "offline_session_max_lifespan_enabled")

    @property
    @pulumi.getter(name="otpPolicy")
    def otp_policy(self) -> 'outputs.GetRealmOtpPolicyResult':
        return pulumi.get(self, "otp_policy")

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
    def refresh_token_max_reuse(self) -> int:
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
    @pulumi.getter(name="revokeRefreshToken")
    def revoke_refresh_token(self) -> bool:
        return pulumi.get(self, "revoke_refresh_token")

    @property
    @pulumi.getter(name="securityDefenses")
    def security_defenses(self) -> Sequence['outputs.GetRealmSecurityDefenseResult']:
        return pulumi.get(self, "security_defenses")

    @property
    @pulumi.getter(name="smtpServers")
    def smtp_servers(self) -> Sequence['outputs.GetRealmSmtpServerResult']:
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
    @pulumi.getter(name="ssoSessionIdleTimeoutRememberMe")
    def sso_session_idle_timeout_remember_me(self) -> str:
        return pulumi.get(self, "sso_session_idle_timeout_remember_me")

    @property
    @pulumi.getter(name="ssoSessionMaxLifespan")
    def sso_session_max_lifespan(self) -> str:
        return pulumi.get(self, "sso_session_max_lifespan")

    @property
    @pulumi.getter(name="ssoSessionMaxLifespanRememberMe")
    def sso_session_max_lifespan_remember_me(self) -> str:
        return pulumi.get(self, "sso_session_max_lifespan_remember_me")

    @property
    @pulumi.getter(name="userManagedAccess")
    def user_managed_access(self) -> bool:
        return pulumi.get(self, "user_managed_access")

    @property
    @pulumi.getter(name="verifyEmail")
    def verify_email(self) -> bool:
        return pulumi.get(self, "verify_email")

    @property
    @pulumi.getter(name="webAuthnPasswordlessPolicy")
    def web_authn_passwordless_policy(self) -> 'outputs.GetRealmWebAuthnPasswordlessPolicyResult':
        return pulumi.get(self, "web_authn_passwordless_policy")

    @property
    @pulumi.getter(name="webAuthnPolicy")
    def web_authn_policy(self) -> 'outputs.GetRealmWebAuthnPolicyResult':
        return pulumi.get(self, "web_authn_policy")


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
            client_session_idle_timeout=self.client_session_idle_timeout,
            client_session_max_lifespan=self.client_session_max_lifespan,
            default_default_client_scopes=self.default_default_client_scopes,
            default_optional_client_scopes=self.default_optional_client_scopes,
            default_signature_algorithm=self.default_signature_algorithm,
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
            oauth2_device_code_lifespan=self.oauth2_device_code_lifespan,
            oauth2_device_polling_interval=self.oauth2_device_polling_interval,
            offline_session_idle_timeout=self.offline_session_idle_timeout,
            offline_session_max_lifespan=self.offline_session_max_lifespan,
            offline_session_max_lifespan_enabled=self.offline_session_max_lifespan_enabled,
            otp_policy=self.otp_policy,
            password_policy=self.password_policy,
            realm=self.realm,
            refresh_token_max_reuse=self.refresh_token_max_reuse,
            registration_allowed=self.registration_allowed,
            registration_email_as_username=self.registration_email_as_username,
            registration_flow=self.registration_flow,
            remember_me=self.remember_me,
            reset_credentials_flow=self.reset_credentials_flow,
            reset_password_allowed=self.reset_password_allowed,
            revoke_refresh_token=self.revoke_refresh_token,
            security_defenses=self.security_defenses,
            smtp_servers=self.smtp_servers,
            ssl_required=self.ssl_required,
            sso_session_idle_timeout=self.sso_session_idle_timeout,
            sso_session_idle_timeout_remember_me=self.sso_session_idle_timeout_remember_me,
            sso_session_max_lifespan=self.sso_session_max_lifespan,
            sso_session_max_lifespan_remember_me=self.sso_session_max_lifespan_remember_me,
            user_managed_access=self.user_managed_access,
            verify_email=self.verify_email,
            web_authn_passwordless_policy=self.web_authn_passwordless_policy,
            web_authn_policy=self.web_authn_policy)


def get_realm(attributes: Optional[Mapping[str, Any]] = None,
              default_default_client_scopes: Optional[Sequence[str]] = None,
              default_optional_client_scopes: Optional[Sequence[str]] = None,
              display_name_html: Optional[str] = None,
              internationalizations: Optional[Sequence[pulumi.InputType['GetRealmInternationalizationArgs']]] = None,
              otp_policy: Optional[pulumi.InputType['GetRealmOtpPolicyArgs']] = None,
              realm: Optional[str] = None,
              security_defenses: Optional[Sequence[pulumi.InputType['GetRealmSecurityDefenseArgs']]] = None,
              smtp_servers: Optional[Sequence[pulumi.InputType['GetRealmSmtpServerArgs']]] = None,
              web_authn_passwordless_policy: Optional[pulumi.InputType['GetRealmWebAuthnPasswordlessPolicyArgs']] = None,
              web_authn_policy: Optional[pulumi.InputType['GetRealmWebAuthnPolicyArgs']] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRealmResult:
    """
    This data source can be used to fetch properties of a Keycloak realm for
    usage with other resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.get_realm(realm="my-realm")
    group = keycloak.Role("group", realm_id=realm.id)
    ```


    :param str realm: The realm name.
    """
    __args__ = dict()
    __args__['attributes'] = attributes
    __args__['defaultDefaultClientScopes'] = default_default_client_scopes
    __args__['defaultOptionalClientScopes'] = default_optional_client_scopes
    __args__['displayNameHtml'] = display_name_html
    __args__['internationalizations'] = internationalizations
    __args__['otpPolicy'] = otp_policy
    __args__['realm'] = realm
    __args__['securityDefenses'] = security_defenses
    __args__['smtpServers'] = smtp_servers
    __args__['webAuthnPasswordlessPolicy'] = web_authn_passwordless_policy
    __args__['webAuthnPolicy'] = web_authn_policy
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
        client_session_idle_timeout=__ret__.client_session_idle_timeout,
        client_session_max_lifespan=__ret__.client_session_max_lifespan,
        default_default_client_scopes=__ret__.default_default_client_scopes,
        default_optional_client_scopes=__ret__.default_optional_client_scopes,
        default_signature_algorithm=__ret__.default_signature_algorithm,
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
        oauth2_device_code_lifespan=__ret__.oauth2_device_code_lifespan,
        oauth2_device_polling_interval=__ret__.oauth2_device_polling_interval,
        offline_session_idle_timeout=__ret__.offline_session_idle_timeout,
        offline_session_max_lifespan=__ret__.offline_session_max_lifespan,
        offline_session_max_lifespan_enabled=__ret__.offline_session_max_lifespan_enabled,
        otp_policy=__ret__.otp_policy,
        password_policy=__ret__.password_policy,
        realm=__ret__.realm,
        refresh_token_max_reuse=__ret__.refresh_token_max_reuse,
        registration_allowed=__ret__.registration_allowed,
        registration_email_as_username=__ret__.registration_email_as_username,
        registration_flow=__ret__.registration_flow,
        remember_me=__ret__.remember_me,
        reset_credentials_flow=__ret__.reset_credentials_flow,
        reset_password_allowed=__ret__.reset_password_allowed,
        revoke_refresh_token=__ret__.revoke_refresh_token,
        security_defenses=__ret__.security_defenses,
        smtp_servers=__ret__.smtp_servers,
        ssl_required=__ret__.ssl_required,
        sso_session_idle_timeout=__ret__.sso_session_idle_timeout,
        sso_session_idle_timeout_remember_me=__ret__.sso_session_idle_timeout_remember_me,
        sso_session_max_lifespan=__ret__.sso_session_max_lifespan,
        sso_session_max_lifespan_remember_me=__ret__.sso_session_max_lifespan_remember_me,
        user_managed_access=__ret__.user_managed_access,
        verify_email=__ret__.verify_email,
        web_authn_passwordless_policy=__ret__.web_authn_passwordless_policy,
        web_authn_policy=__ret__.web_authn_policy)


@_utilities.lift_output_func(get_realm)
def get_realm_output(attributes: Optional[pulumi.Input[Optional[Mapping[str, Any]]]] = None,
                     default_default_client_scopes: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     default_optional_client_scopes: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                     display_name_html: Optional[pulumi.Input[Optional[str]]] = None,
                     internationalizations: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRealmInternationalizationArgs']]]]] = None,
                     otp_policy: Optional[pulumi.Input[Optional[pulumi.InputType['GetRealmOtpPolicyArgs']]]] = None,
                     realm: Optional[pulumi.Input[str]] = None,
                     security_defenses: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRealmSecurityDefenseArgs']]]]] = None,
                     smtp_servers: Optional[pulumi.Input[Optional[Sequence[pulumi.InputType['GetRealmSmtpServerArgs']]]]] = None,
                     web_authn_passwordless_policy: Optional[pulumi.Input[Optional[pulumi.InputType['GetRealmWebAuthnPasswordlessPolicyArgs']]]] = None,
                     web_authn_policy: Optional[pulumi.Input[Optional[pulumi.InputType['GetRealmWebAuthnPolicyArgs']]]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRealmResult]:
    """
    This data source can be used to fetch properties of a Keycloak realm for
    usage with other resources.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.get_realm(realm="my-realm")
    group = keycloak.Role("group", realm_id=realm.id)
    ```


    :param str realm: The realm name.
    """
    ...
