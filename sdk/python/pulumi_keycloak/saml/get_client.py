# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetClientResult',
    'AwaitableGetClientResult',
    'get_client',
    'get_client_output',
]

@pulumi.output_type
class GetClientResult:
    """
    A collection of values returned by getClient.
    """
    def __init__(__self__, assertion_consumer_post_url=None, assertion_consumer_redirect_url=None, authentication_flow_binding_overrides=None, base_url=None, canonicalization_method=None, client_id=None, client_signature_required=None, description=None, enabled=None, encrypt_assertions=None, encryption_certificate=None, encryption_certificate_sha1=None, extra_config=None, force_name_id_format=None, force_post_binding=None, front_channel_logout=None, full_scope_allowed=None, id=None, idp_initiated_sso_relay_state=None, idp_initiated_sso_url_name=None, include_authn_statement=None, login_theme=None, logout_service_post_binding_url=None, logout_service_redirect_binding_url=None, master_saml_processing_url=None, name=None, name_id_format=None, realm_id=None, root_url=None, saml_signature_key_name=None, sign_assertions=None, sign_documents=None, signature_algorithm=None, signature_key_name=None, signing_certificate=None, signing_certificate_sha1=None, signing_private_key=None, signing_private_key_sha1=None, valid_redirect_uris=None):
        if assertion_consumer_post_url and not isinstance(assertion_consumer_post_url, str):
            raise TypeError("Expected argument 'assertion_consumer_post_url' to be a str")
        pulumi.set(__self__, "assertion_consumer_post_url", assertion_consumer_post_url)
        if assertion_consumer_redirect_url and not isinstance(assertion_consumer_redirect_url, str):
            raise TypeError("Expected argument 'assertion_consumer_redirect_url' to be a str")
        pulumi.set(__self__, "assertion_consumer_redirect_url", assertion_consumer_redirect_url)
        if authentication_flow_binding_overrides and not isinstance(authentication_flow_binding_overrides, list):
            raise TypeError("Expected argument 'authentication_flow_binding_overrides' to be a list")
        pulumi.set(__self__, "authentication_flow_binding_overrides", authentication_flow_binding_overrides)
        if base_url and not isinstance(base_url, str):
            raise TypeError("Expected argument 'base_url' to be a str")
        pulumi.set(__self__, "base_url", base_url)
        if canonicalization_method and not isinstance(canonicalization_method, str):
            raise TypeError("Expected argument 'canonicalization_method' to be a str")
        pulumi.set(__self__, "canonicalization_method", canonicalization_method)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_signature_required and not isinstance(client_signature_required, bool):
            raise TypeError("Expected argument 'client_signature_required' to be a bool")
        pulumi.set(__self__, "client_signature_required", client_signature_required)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if encrypt_assertions and not isinstance(encrypt_assertions, bool):
            raise TypeError("Expected argument 'encrypt_assertions' to be a bool")
        pulumi.set(__self__, "encrypt_assertions", encrypt_assertions)
        if encryption_certificate and not isinstance(encryption_certificate, str):
            raise TypeError("Expected argument 'encryption_certificate' to be a str")
        pulumi.set(__self__, "encryption_certificate", encryption_certificate)
        if encryption_certificate_sha1 and not isinstance(encryption_certificate_sha1, str):
            raise TypeError("Expected argument 'encryption_certificate_sha1' to be a str")
        pulumi.set(__self__, "encryption_certificate_sha1", encryption_certificate_sha1)
        if extra_config and not isinstance(extra_config, dict):
            raise TypeError("Expected argument 'extra_config' to be a dict")
        pulumi.set(__self__, "extra_config", extra_config)
        if force_name_id_format and not isinstance(force_name_id_format, bool):
            raise TypeError("Expected argument 'force_name_id_format' to be a bool")
        pulumi.set(__self__, "force_name_id_format", force_name_id_format)
        if force_post_binding and not isinstance(force_post_binding, bool):
            raise TypeError("Expected argument 'force_post_binding' to be a bool")
        pulumi.set(__self__, "force_post_binding", force_post_binding)
        if front_channel_logout and not isinstance(front_channel_logout, bool):
            raise TypeError("Expected argument 'front_channel_logout' to be a bool")
        pulumi.set(__self__, "front_channel_logout", front_channel_logout)
        if full_scope_allowed and not isinstance(full_scope_allowed, bool):
            raise TypeError("Expected argument 'full_scope_allowed' to be a bool")
        pulumi.set(__self__, "full_scope_allowed", full_scope_allowed)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if idp_initiated_sso_relay_state and not isinstance(idp_initiated_sso_relay_state, str):
            raise TypeError("Expected argument 'idp_initiated_sso_relay_state' to be a str")
        pulumi.set(__self__, "idp_initiated_sso_relay_state", idp_initiated_sso_relay_state)
        if idp_initiated_sso_url_name and not isinstance(idp_initiated_sso_url_name, str):
            raise TypeError("Expected argument 'idp_initiated_sso_url_name' to be a str")
        pulumi.set(__self__, "idp_initiated_sso_url_name", idp_initiated_sso_url_name)
        if include_authn_statement and not isinstance(include_authn_statement, bool):
            raise TypeError("Expected argument 'include_authn_statement' to be a bool")
        pulumi.set(__self__, "include_authn_statement", include_authn_statement)
        if login_theme and not isinstance(login_theme, str):
            raise TypeError("Expected argument 'login_theme' to be a str")
        pulumi.set(__self__, "login_theme", login_theme)
        if logout_service_post_binding_url and not isinstance(logout_service_post_binding_url, str):
            raise TypeError("Expected argument 'logout_service_post_binding_url' to be a str")
        pulumi.set(__self__, "logout_service_post_binding_url", logout_service_post_binding_url)
        if logout_service_redirect_binding_url and not isinstance(logout_service_redirect_binding_url, str):
            raise TypeError("Expected argument 'logout_service_redirect_binding_url' to be a str")
        pulumi.set(__self__, "logout_service_redirect_binding_url", logout_service_redirect_binding_url)
        if master_saml_processing_url and not isinstance(master_saml_processing_url, str):
            raise TypeError("Expected argument 'master_saml_processing_url' to be a str")
        pulumi.set(__self__, "master_saml_processing_url", master_saml_processing_url)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if name_id_format and not isinstance(name_id_format, str):
            raise TypeError("Expected argument 'name_id_format' to be a str")
        pulumi.set(__self__, "name_id_format", name_id_format)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if root_url and not isinstance(root_url, str):
            raise TypeError("Expected argument 'root_url' to be a str")
        pulumi.set(__self__, "root_url", root_url)
        if saml_signature_key_name and not isinstance(saml_signature_key_name, str):
            raise TypeError("Expected argument 'saml_signature_key_name' to be a str")
        pulumi.set(__self__, "saml_signature_key_name", saml_signature_key_name)
        if sign_assertions and not isinstance(sign_assertions, bool):
            raise TypeError("Expected argument 'sign_assertions' to be a bool")
        pulumi.set(__self__, "sign_assertions", sign_assertions)
        if sign_documents and not isinstance(sign_documents, bool):
            raise TypeError("Expected argument 'sign_documents' to be a bool")
        pulumi.set(__self__, "sign_documents", sign_documents)
        if signature_algorithm and not isinstance(signature_algorithm, str):
            raise TypeError("Expected argument 'signature_algorithm' to be a str")
        pulumi.set(__self__, "signature_algorithm", signature_algorithm)
        if signature_key_name and not isinstance(signature_key_name, str):
            raise TypeError("Expected argument 'signature_key_name' to be a str")
        pulumi.set(__self__, "signature_key_name", signature_key_name)
        if signing_certificate and not isinstance(signing_certificate, str):
            raise TypeError("Expected argument 'signing_certificate' to be a str")
        pulumi.set(__self__, "signing_certificate", signing_certificate)
        if signing_certificate_sha1 and not isinstance(signing_certificate_sha1, str):
            raise TypeError("Expected argument 'signing_certificate_sha1' to be a str")
        pulumi.set(__self__, "signing_certificate_sha1", signing_certificate_sha1)
        if signing_private_key and not isinstance(signing_private_key, str):
            raise TypeError("Expected argument 'signing_private_key' to be a str")
        pulumi.set(__self__, "signing_private_key", signing_private_key)
        if signing_private_key_sha1 and not isinstance(signing_private_key_sha1, str):
            raise TypeError("Expected argument 'signing_private_key_sha1' to be a str")
        pulumi.set(__self__, "signing_private_key_sha1", signing_private_key_sha1)
        if valid_redirect_uris and not isinstance(valid_redirect_uris, list):
            raise TypeError("Expected argument 'valid_redirect_uris' to be a list")
        pulumi.set(__self__, "valid_redirect_uris", valid_redirect_uris)

    @property
    @pulumi.getter(name="assertionConsumerPostUrl")
    def assertion_consumer_post_url(self) -> str:
        return pulumi.get(self, "assertion_consumer_post_url")

    @property
    @pulumi.getter(name="assertionConsumerRedirectUrl")
    def assertion_consumer_redirect_url(self) -> str:
        return pulumi.get(self, "assertion_consumer_redirect_url")

    @property
    @pulumi.getter(name="authenticationFlowBindingOverrides")
    def authentication_flow_binding_overrides(self) -> Sequence['outputs.GetClientAuthenticationFlowBindingOverrideResult']:
        return pulumi.get(self, "authentication_flow_binding_overrides")

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> str:
        return pulumi.get(self, "base_url")

    @property
    @pulumi.getter(name="canonicalizationMethod")
    def canonicalization_method(self) -> str:
        return pulumi.get(self, "canonicalization_method")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSignatureRequired")
    def client_signature_required(self) -> bool:
        return pulumi.get(self, "client_signature_required")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="encryptAssertions")
    def encrypt_assertions(self) -> bool:
        return pulumi.get(self, "encrypt_assertions")

    @property
    @pulumi.getter(name="encryptionCertificate")
    def encryption_certificate(self) -> str:
        return pulumi.get(self, "encryption_certificate")

    @property
    @pulumi.getter(name="encryptionCertificateSha1")
    def encryption_certificate_sha1(self) -> str:
        return pulumi.get(self, "encryption_certificate_sha1")

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> Mapping[str, Any]:
        return pulumi.get(self, "extra_config")

    @property
    @pulumi.getter(name="forceNameIdFormat")
    def force_name_id_format(self) -> bool:
        return pulumi.get(self, "force_name_id_format")

    @property
    @pulumi.getter(name="forcePostBinding")
    def force_post_binding(self) -> bool:
        return pulumi.get(self, "force_post_binding")

    @property
    @pulumi.getter(name="frontChannelLogout")
    def front_channel_logout(self) -> bool:
        return pulumi.get(self, "front_channel_logout")

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
    @pulumi.getter(name="idpInitiatedSsoRelayState")
    def idp_initiated_sso_relay_state(self) -> str:
        return pulumi.get(self, "idp_initiated_sso_relay_state")

    @property
    @pulumi.getter(name="idpInitiatedSsoUrlName")
    def idp_initiated_sso_url_name(self) -> str:
        return pulumi.get(self, "idp_initiated_sso_url_name")

    @property
    @pulumi.getter(name="includeAuthnStatement")
    def include_authn_statement(self) -> bool:
        return pulumi.get(self, "include_authn_statement")

    @property
    @pulumi.getter(name="loginTheme")
    def login_theme(self) -> str:
        return pulumi.get(self, "login_theme")

    @property
    @pulumi.getter(name="logoutServicePostBindingUrl")
    def logout_service_post_binding_url(self) -> str:
        return pulumi.get(self, "logout_service_post_binding_url")

    @property
    @pulumi.getter(name="logoutServiceRedirectBindingUrl")
    def logout_service_redirect_binding_url(self) -> str:
        return pulumi.get(self, "logout_service_redirect_binding_url")

    @property
    @pulumi.getter(name="masterSamlProcessingUrl")
    def master_saml_processing_url(self) -> str:
        return pulumi.get(self, "master_saml_processing_url")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="nameIdFormat")
    def name_id_format(self) -> str:
        return pulumi.get(self, "name_id_format")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="rootUrl")
    def root_url(self) -> str:
        return pulumi.get(self, "root_url")

    @property
    @pulumi.getter(name="samlSignatureKeyName")
    def saml_signature_key_name(self) -> str:
        return pulumi.get(self, "saml_signature_key_name")

    @property
    @pulumi.getter(name="signAssertions")
    def sign_assertions(self) -> bool:
        return pulumi.get(self, "sign_assertions")

    @property
    @pulumi.getter(name="signDocuments")
    def sign_documents(self) -> bool:
        return pulumi.get(self, "sign_documents")

    @property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> str:
        return pulumi.get(self, "signature_algorithm")

    @property
    @pulumi.getter(name="signatureKeyName")
    def signature_key_name(self) -> str:
        return pulumi.get(self, "signature_key_name")

    @property
    @pulumi.getter(name="signingCertificate")
    def signing_certificate(self) -> str:
        return pulumi.get(self, "signing_certificate")

    @property
    @pulumi.getter(name="signingCertificateSha1")
    def signing_certificate_sha1(self) -> str:
        return pulumi.get(self, "signing_certificate_sha1")

    @property
    @pulumi.getter(name="signingPrivateKey")
    def signing_private_key(self) -> str:
        return pulumi.get(self, "signing_private_key")

    @property
    @pulumi.getter(name="signingPrivateKeySha1")
    def signing_private_key_sha1(self) -> str:
        return pulumi.get(self, "signing_private_key_sha1")

    @property
    @pulumi.getter(name="validRedirectUris")
    def valid_redirect_uris(self) -> Sequence[str]:
        return pulumi.get(self, "valid_redirect_uris")


class AwaitableGetClientResult(GetClientResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientResult(
            assertion_consumer_post_url=self.assertion_consumer_post_url,
            assertion_consumer_redirect_url=self.assertion_consumer_redirect_url,
            authentication_flow_binding_overrides=self.authentication_flow_binding_overrides,
            base_url=self.base_url,
            canonicalization_method=self.canonicalization_method,
            client_id=self.client_id,
            client_signature_required=self.client_signature_required,
            description=self.description,
            enabled=self.enabled,
            encrypt_assertions=self.encrypt_assertions,
            encryption_certificate=self.encryption_certificate,
            encryption_certificate_sha1=self.encryption_certificate_sha1,
            extra_config=self.extra_config,
            force_name_id_format=self.force_name_id_format,
            force_post_binding=self.force_post_binding,
            front_channel_logout=self.front_channel_logout,
            full_scope_allowed=self.full_scope_allowed,
            id=self.id,
            idp_initiated_sso_relay_state=self.idp_initiated_sso_relay_state,
            idp_initiated_sso_url_name=self.idp_initiated_sso_url_name,
            include_authn_statement=self.include_authn_statement,
            login_theme=self.login_theme,
            logout_service_post_binding_url=self.logout_service_post_binding_url,
            logout_service_redirect_binding_url=self.logout_service_redirect_binding_url,
            master_saml_processing_url=self.master_saml_processing_url,
            name=self.name,
            name_id_format=self.name_id_format,
            realm_id=self.realm_id,
            root_url=self.root_url,
            saml_signature_key_name=self.saml_signature_key_name,
            sign_assertions=self.sign_assertions,
            sign_documents=self.sign_documents,
            signature_algorithm=self.signature_algorithm,
            signature_key_name=self.signature_key_name,
            signing_certificate=self.signing_certificate,
            signing_certificate_sha1=self.signing_certificate_sha1,
            signing_private_key=self.signing_private_key,
            signing_private_key_sha1=self.signing_private_key_sha1,
            valid_redirect_uris=self.valid_redirect_uris)


def get_client(client_id: Optional[str] = None,
               realm_id: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientResult:
    """
    This data source can be used to fetch properties of a Keycloak client that uses the SAML protocol.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm_management = keycloak.saml.get_client(realm_id="my-realm",
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
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:saml/getClient:getClient', __args__, opts=opts, typ=GetClientResult).value

    return AwaitableGetClientResult(
        assertion_consumer_post_url=pulumi.get(__ret__, 'assertion_consumer_post_url'),
        assertion_consumer_redirect_url=pulumi.get(__ret__, 'assertion_consumer_redirect_url'),
        authentication_flow_binding_overrides=pulumi.get(__ret__, 'authentication_flow_binding_overrides'),
        base_url=pulumi.get(__ret__, 'base_url'),
        canonicalization_method=pulumi.get(__ret__, 'canonicalization_method'),
        client_id=pulumi.get(__ret__, 'client_id'),
        client_signature_required=pulumi.get(__ret__, 'client_signature_required'),
        description=pulumi.get(__ret__, 'description'),
        enabled=pulumi.get(__ret__, 'enabled'),
        encrypt_assertions=pulumi.get(__ret__, 'encrypt_assertions'),
        encryption_certificate=pulumi.get(__ret__, 'encryption_certificate'),
        encryption_certificate_sha1=pulumi.get(__ret__, 'encryption_certificate_sha1'),
        extra_config=pulumi.get(__ret__, 'extra_config'),
        force_name_id_format=pulumi.get(__ret__, 'force_name_id_format'),
        force_post_binding=pulumi.get(__ret__, 'force_post_binding'),
        front_channel_logout=pulumi.get(__ret__, 'front_channel_logout'),
        full_scope_allowed=pulumi.get(__ret__, 'full_scope_allowed'),
        id=pulumi.get(__ret__, 'id'),
        idp_initiated_sso_relay_state=pulumi.get(__ret__, 'idp_initiated_sso_relay_state'),
        idp_initiated_sso_url_name=pulumi.get(__ret__, 'idp_initiated_sso_url_name'),
        include_authn_statement=pulumi.get(__ret__, 'include_authn_statement'),
        login_theme=pulumi.get(__ret__, 'login_theme'),
        logout_service_post_binding_url=pulumi.get(__ret__, 'logout_service_post_binding_url'),
        logout_service_redirect_binding_url=pulumi.get(__ret__, 'logout_service_redirect_binding_url'),
        master_saml_processing_url=pulumi.get(__ret__, 'master_saml_processing_url'),
        name=pulumi.get(__ret__, 'name'),
        name_id_format=pulumi.get(__ret__, 'name_id_format'),
        realm_id=pulumi.get(__ret__, 'realm_id'),
        root_url=pulumi.get(__ret__, 'root_url'),
        saml_signature_key_name=pulumi.get(__ret__, 'saml_signature_key_name'),
        sign_assertions=pulumi.get(__ret__, 'sign_assertions'),
        sign_documents=pulumi.get(__ret__, 'sign_documents'),
        signature_algorithm=pulumi.get(__ret__, 'signature_algorithm'),
        signature_key_name=pulumi.get(__ret__, 'signature_key_name'),
        signing_certificate=pulumi.get(__ret__, 'signing_certificate'),
        signing_certificate_sha1=pulumi.get(__ret__, 'signing_certificate_sha1'),
        signing_private_key=pulumi.get(__ret__, 'signing_private_key'),
        signing_private_key_sha1=pulumi.get(__ret__, 'signing_private_key_sha1'),
        valid_redirect_uris=pulumi.get(__ret__, 'valid_redirect_uris'))


@_utilities.lift_output_func(get_client)
def get_client_output(client_id: Optional[pulumi.Input[str]] = None,
                      realm_id: Optional[pulumi.Input[str]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClientResult]:
    """
    This data source can be used to fetch properties of a Keycloak client that uses the SAML protocol.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm_management = keycloak.saml.get_client(realm_id="my-realm",
        client_id="realm-management")
    admin = keycloak.get_role(realm_id="my-realm",
        client_id=realm_management.id,
        name="realm-admin")
    ```


    :param str client_id: The client id (not its unique ID).
    :param str realm_id: The realm id.
    """
    ...
