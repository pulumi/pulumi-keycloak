# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'GetClientDescriptionConverterResult',
    'AwaitableGetClientDescriptionConverterResult',
    'get_client_description_converter',
    'get_client_description_converter_output',
]

@pulumi.output_type
class GetClientDescriptionConverterResult:
    """
    A collection of values returned by getClientDescriptionConverter.
    """
    def __init__(__self__, access=None, admin_url=None, attributes=None, authentication_flow_binding_overrides=None, authorization_services_enabled=None, authorization_settings=None, base_url=None, bearer_only=None, body=None, client_authenticator_type=None, client_id=None, consent_required=None, default_client_scopes=None, default_roles=None, description=None, direct_access_grants_enabled=None, enabled=None, frontchannel_logout=None, full_scope_allowed=None, id=None, implicit_flow_enabled=None, name=None, not_before=None, optional_client_scopes=None, origin=None, protocol=None, protocol_mappers=None, public_client=None, realm_id=None, redirect_uris=None, registered_nodes=None, registration_access_token=None, root_url=None, secret=None, service_accounts_enabled=None, standard_flow_enabled=None, surrogate_auth_required=None, web_origins=None):
        if access and not isinstance(access, dict):
            raise TypeError("Expected argument 'access' to be a dict")
        pulumi.set(__self__, "access", access)
        if admin_url and not isinstance(admin_url, str):
            raise TypeError("Expected argument 'admin_url' to be a str")
        pulumi.set(__self__, "admin_url", admin_url)
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if authentication_flow_binding_overrides and not isinstance(authentication_flow_binding_overrides, dict):
            raise TypeError("Expected argument 'authentication_flow_binding_overrides' to be a dict")
        pulumi.set(__self__, "authentication_flow_binding_overrides", authentication_flow_binding_overrides)
        if authorization_services_enabled and not isinstance(authorization_services_enabled, bool):
            raise TypeError("Expected argument 'authorization_services_enabled' to be a bool")
        pulumi.set(__self__, "authorization_services_enabled", authorization_services_enabled)
        if authorization_settings and not isinstance(authorization_settings, dict):
            raise TypeError("Expected argument 'authorization_settings' to be a dict")
        pulumi.set(__self__, "authorization_settings", authorization_settings)
        if base_url and not isinstance(base_url, str):
            raise TypeError("Expected argument 'base_url' to be a str")
        pulumi.set(__self__, "base_url", base_url)
        if bearer_only and not isinstance(bearer_only, bool):
            raise TypeError("Expected argument 'bearer_only' to be a bool")
        pulumi.set(__self__, "bearer_only", bearer_only)
        if body and not isinstance(body, str):
            raise TypeError("Expected argument 'body' to be a str")
        pulumi.set(__self__, "body", body)
        if client_authenticator_type and not isinstance(client_authenticator_type, str):
            raise TypeError("Expected argument 'client_authenticator_type' to be a str")
        pulumi.set(__self__, "client_authenticator_type", client_authenticator_type)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if consent_required and not isinstance(consent_required, str):
            raise TypeError("Expected argument 'consent_required' to be a str")
        pulumi.set(__self__, "consent_required", consent_required)
        if default_client_scopes and not isinstance(default_client_scopes, list):
            raise TypeError("Expected argument 'default_client_scopes' to be a list")
        pulumi.set(__self__, "default_client_scopes", default_client_scopes)
        if default_roles and not isinstance(default_roles, list):
            raise TypeError("Expected argument 'default_roles' to be a list")
        pulumi.set(__self__, "default_roles", default_roles)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if direct_access_grants_enabled and not isinstance(direct_access_grants_enabled, bool):
            raise TypeError("Expected argument 'direct_access_grants_enabled' to be a bool")
        pulumi.set(__self__, "direct_access_grants_enabled", direct_access_grants_enabled)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if frontchannel_logout and not isinstance(frontchannel_logout, bool):
            raise TypeError("Expected argument 'frontchannel_logout' to be a bool")
        pulumi.set(__self__, "frontchannel_logout", frontchannel_logout)
        if full_scope_allowed and not isinstance(full_scope_allowed, bool):
            raise TypeError("Expected argument 'full_scope_allowed' to be a bool")
        pulumi.set(__self__, "full_scope_allowed", full_scope_allowed)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if implicit_flow_enabled and not isinstance(implicit_flow_enabled, bool):
            raise TypeError("Expected argument 'implicit_flow_enabled' to be a bool")
        pulumi.set(__self__, "implicit_flow_enabled", implicit_flow_enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if not_before and not isinstance(not_before, int):
            raise TypeError("Expected argument 'not_before' to be a int")
        pulumi.set(__self__, "not_before", not_before)
        if optional_client_scopes and not isinstance(optional_client_scopes, list):
            raise TypeError("Expected argument 'optional_client_scopes' to be a list")
        pulumi.set(__self__, "optional_client_scopes", optional_client_scopes)
        if origin and not isinstance(origin, str):
            raise TypeError("Expected argument 'origin' to be a str")
        pulumi.set(__self__, "origin", origin)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if protocol_mappers and not isinstance(protocol_mappers, list):
            raise TypeError("Expected argument 'protocol_mappers' to be a list")
        pulumi.set(__self__, "protocol_mappers", protocol_mappers)
        if public_client and not isinstance(public_client, bool):
            raise TypeError("Expected argument 'public_client' to be a bool")
        pulumi.set(__self__, "public_client", public_client)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if redirect_uris and not isinstance(redirect_uris, list):
            raise TypeError("Expected argument 'redirect_uris' to be a list")
        pulumi.set(__self__, "redirect_uris", redirect_uris)
        if registered_nodes and not isinstance(registered_nodes, dict):
            raise TypeError("Expected argument 'registered_nodes' to be a dict")
        pulumi.set(__self__, "registered_nodes", registered_nodes)
        if registration_access_token and not isinstance(registration_access_token, str):
            raise TypeError("Expected argument 'registration_access_token' to be a str")
        pulumi.set(__self__, "registration_access_token", registration_access_token)
        if root_url and not isinstance(root_url, str):
            raise TypeError("Expected argument 'root_url' to be a str")
        pulumi.set(__self__, "root_url", root_url)
        if secret and not isinstance(secret, str):
            raise TypeError("Expected argument 'secret' to be a str")
        pulumi.set(__self__, "secret", secret)
        if service_accounts_enabled and not isinstance(service_accounts_enabled, bool):
            raise TypeError("Expected argument 'service_accounts_enabled' to be a bool")
        pulumi.set(__self__, "service_accounts_enabled", service_accounts_enabled)
        if standard_flow_enabled and not isinstance(standard_flow_enabled, bool):
            raise TypeError("Expected argument 'standard_flow_enabled' to be a bool")
        pulumi.set(__self__, "standard_flow_enabled", standard_flow_enabled)
        if surrogate_auth_required and not isinstance(surrogate_auth_required, bool):
            raise TypeError("Expected argument 'surrogate_auth_required' to be a bool")
        pulumi.set(__self__, "surrogate_auth_required", surrogate_auth_required)
        if web_origins and not isinstance(web_origins, list):
            raise TypeError("Expected argument 'web_origins' to be a list")
        pulumi.set(__self__, "web_origins", web_origins)

    @property
    @pulumi.getter
    def access(self) -> Mapping[str, Any]:
        return pulumi.get(self, "access")

    @property
    @pulumi.getter(name="adminUrl")
    def admin_url(self) -> str:
        return pulumi.get(self, "admin_url")

    @property
    @pulumi.getter
    def attributes(self) -> Mapping[str, Any]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="authenticationFlowBindingOverrides")
    def authentication_flow_binding_overrides(self) -> Mapping[str, Any]:
        return pulumi.get(self, "authentication_flow_binding_overrides")

    @property
    @pulumi.getter(name="authorizationServicesEnabled")
    def authorization_services_enabled(self) -> bool:
        return pulumi.get(self, "authorization_services_enabled")

    @property
    @pulumi.getter(name="authorizationSettings")
    def authorization_settings(self) -> Mapping[str, Any]:
        return pulumi.get(self, "authorization_settings")

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> str:
        return pulumi.get(self, "base_url")

    @property
    @pulumi.getter(name="bearerOnly")
    def bearer_only(self) -> bool:
        return pulumi.get(self, "bearer_only")

    @property
    @pulumi.getter
    def body(self) -> str:
        return pulumi.get(self, "body")

    @property
    @pulumi.getter(name="clientAuthenticatorType")
    def client_authenticator_type(self) -> str:
        return pulumi.get(self, "client_authenticator_type")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="consentRequired")
    def consent_required(self) -> str:
        return pulumi.get(self, "consent_required")

    @property
    @pulumi.getter(name="defaultClientScopes")
    def default_client_scopes(self) -> Sequence[str]:
        return pulumi.get(self, "default_client_scopes")

    @property
    @pulumi.getter(name="defaultRoles")
    def default_roles(self) -> Sequence[str]:
        return pulumi.get(self, "default_roles")

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
    @pulumi.getter(name="frontchannelLogout")
    def frontchannel_logout(self) -> bool:
        return pulumi.get(self, "frontchannel_logout")

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
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notBefore")
    def not_before(self) -> int:
        return pulumi.get(self, "not_before")

    @property
    @pulumi.getter(name="optionalClientScopes")
    def optional_client_scopes(self) -> Sequence[str]:
        return pulumi.get(self, "optional_client_scopes")

    @property
    @pulumi.getter
    def origin(self) -> str:
        return pulumi.get(self, "origin")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="protocolMappers")
    def protocol_mappers(self) -> Sequence['outputs.GetClientDescriptionConverterProtocolMapperResult']:
        return pulumi.get(self, "protocol_mappers")

    @property
    @pulumi.getter(name="publicClient")
    def public_client(self) -> bool:
        return pulumi.get(self, "public_client")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="redirectUris")
    def redirect_uris(self) -> Sequence[str]:
        return pulumi.get(self, "redirect_uris")

    @property
    @pulumi.getter(name="registeredNodes")
    def registered_nodes(self) -> Mapping[str, Any]:
        return pulumi.get(self, "registered_nodes")

    @property
    @pulumi.getter(name="registrationAccessToken")
    def registration_access_token(self) -> str:
        return pulumi.get(self, "registration_access_token")

    @property
    @pulumi.getter(name="rootUrl")
    def root_url(self) -> str:
        return pulumi.get(self, "root_url")

    @property
    @pulumi.getter
    def secret(self) -> str:
        return pulumi.get(self, "secret")

    @property
    @pulumi.getter(name="serviceAccountsEnabled")
    def service_accounts_enabled(self) -> bool:
        return pulumi.get(self, "service_accounts_enabled")

    @property
    @pulumi.getter(name="standardFlowEnabled")
    def standard_flow_enabled(self) -> bool:
        return pulumi.get(self, "standard_flow_enabled")

    @property
    @pulumi.getter(name="surrogateAuthRequired")
    def surrogate_auth_required(self) -> bool:
        return pulumi.get(self, "surrogate_auth_required")

    @property
    @pulumi.getter(name="webOrigins")
    def web_origins(self) -> Sequence[str]:
        return pulumi.get(self, "web_origins")


class AwaitableGetClientDescriptionConverterResult(GetClientDescriptionConverterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientDescriptionConverterResult(
            access=self.access,
            admin_url=self.admin_url,
            attributes=self.attributes,
            authentication_flow_binding_overrides=self.authentication_flow_binding_overrides,
            authorization_services_enabled=self.authorization_services_enabled,
            authorization_settings=self.authorization_settings,
            base_url=self.base_url,
            bearer_only=self.bearer_only,
            body=self.body,
            client_authenticator_type=self.client_authenticator_type,
            client_id=self.client_id,
            consent_required=self.consent_required,
            default_client_scopes=self.default_client_scopes,
            default_roles=self.default_roles,
            description=self.description,
            direct_access_grants_enabled=self.direct_access_grants_enabled,
            enabled=self.enabled,
            frontchannel_logout=self.frontchannel_logout,
            full_scope_allowed=self.full_scope_allowed,
            id=self.id,
            implicit_flow_enabled=self.implicit_flow_enabled,
            name=self.name,
            not_before=self.not_before,
            optional_client_scopes=self.optional_client_scopes,
            origin=self.origin,
            protocol=self.protocol,
            protocol_mappers=self.protocol_mappers,
            public_client=self.public_client,
            realm_id=self.realm_id,
            redirect_uris=self.redirect_uris,
            registered_nodes=self.registered_nodes,
            registration_access_token=self.registration_access_token,
            root_url=self.root_url,
            secret=self.secret,
            service_accounts_enabled=self.service_accounts_enabled,
            standard_flow_enabled=self.standard_flow_enabled,
            surrogate_auth_required=self.surrogate_auth_required,
            web_origins=self.web_origins)


def get_client_description_converter(body: Optional[str] = None,
                                     realm_id: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientDescriptionConverterResult:
    """
    This data source uses the [ClientDescriptionConverter](https://www.keycloak.org/docs-api/6.0/javadocs/org/keycloak/exportimport/ClientDescriptionConverter.html) API to convert a generic client description into a Keycloak
    client. This data can then be used to manage the client within Keycloak.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.Realm("realm",
        realm="my-realm",
        enabled=True)
    saml_client = keycloak.get_client_description_converter_output(realm_id=realm.id,
        body=\"\"\"\\x09<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" validUntil="2021-04-17T12:41:46Z" cacheDuration="PT604800S" entityID="FakeEntityId">
        <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
            <md:KeyDescriptor use="signing">
    \\x09\\x09\\x09<ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
    \\x09\\x09\\x09\\x09<ds:X509Data>
    \\x09\\x09\\x09\\x09\\x09<ds:X509Certificate>MIICyDCCAjGgAwIBAgIBADANBgkqhkiG9w0BAQ0FADCBgDELMAkGA1UEBhMCdXMx
    \\x09\\x09\\x09\\x09\\x09CzAJBgNVBAgMAklBMSQwIgYDVQQKDBt0ZXJyYWZvcm0tcHJvdmlkZXIta2V5Y2xv
    \\x09\\x09\\x09\\x09\\x09YWsxHDAaBgNVBAMME21ycGFya2Vycy5naXRodWIuaW8xIDAeBgkqhkiG9w0BCQEW
    \\x09\\x09\\x09\\x09\\x09EW1pY2hhZWxAcGFya2VyLmdnMB4XDTE5MDEwODE0NDYzNloXDTI5MDEwNTE0NDYz
    \\x09\\x09\\x09\\x09\\x09NlowgYAxCzAJBgNVBAYTAnVzMQswCQYDVQQIDAJJQTEkMCIGA1UECgwbdGVycmFm
    \\x09\\x09\\x09\\x09\\x09b3JtLXByb3ZpZGVyLWtleWNsb2FrMRwwGgYDVQQDDBNtcnBhcmtlcnMuZ2l0aHVi
    \\x09\\x09\\x09\\x09\\x09LmlvMSAwHgYJKoZIhvcNAQkBFhFtaWNoYWVsQHBhcmtlci5nZzCBnzANBgkqhkiG
    \\x09\\x09\\x09\\x09\\x099w0BAQEFAAOBjQAwgYkCgYEAxuZny7uyYxGVPtpie14gNQC4tT9sAvO2sVNDhuoe
    \\x09\\x09\\x09\\x09\\x09qIKLRpNwkHnwQmwe5OxSh9K0BPHp/DNuuVWUqvo4tniEYn3jBr7FwLYLTKojQIxj
    \\x09\\x09\\x09\\x09\\x0953S1UTT9EXq3eP5HsHMD0QnTuca2nlNYUDBm6ud2fQj0Jt5qLx86EbEC28N56IRv
    \\x09\\x09\\x09\\x09\\x09GX8CAwEAAaNQME4wHQYDVR0OBBYEFMLnbQh77j7vhGTpAhKpDhCrBsPZMB8GA1Ud
    \\x09\\x09\\x09\\x09\\x09IwQYMBaAFMLnbQh77j7vhGTpAhKpDhCrBsPZMAwGA1UdEwQFMAMBAf8wDQYJKoZI
    \\x09\\x09\\x09\\x09\\x09hvcNAQENBQADgYEAB8wGrAQY0pAfwbnYSyBt4STbebeRTu1/q1ucfrtc3qsegcd5
    \\x09\\x09\\x09\\x09\\x09n01xTR+T2uZJwqHFPpFjr4IPORiHx3+4BWCweslPD53qBjKUPXcbMO1Revjef6Tj
    \\x09\\x09\\x09\\x09\\x09K3K0AuJ94fxgXVoT61Nzu/a6Lj6RhzU/Dao9mlSbJY+YSbm+ZBpsuRUQ84s=</ds:X509Certificate>
    \\x09\\x09\\x09\\x09</ds:X509Data>
    \\x09\\x09\\x09</ds:KeyInfo>
    \\x09\\x09</md:KeyDescriptor>
    \\x09\\x09<md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
            <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://localhost/acs/saml/" index="1"/>
        </md:SPSSODescriptor>
    </md:EntityDescriptor>
    \"\"\")
    saml_client_client = keycloak.saml.Client("saml_client",
        realm_id=realm.id,
        client_id=saml_client.client_id)
    ```
    <!--End PulumiCodeChooser -->


    :param str body: The body of the request to convert.
    :param str realm_id: The realm to use for the client description converter API call.
    """
    __args__ = dict()
    __args__['body'] = body
    __args__['realmId'] = realm_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:index/getClientDescriptionConverter:getClientDescriptionConverter', __args__, opts=opts, typ=GetClientDescriptionConverterResult).value

    return AwaitableGetClientDescriptionConverterResult(
        access=pulumi.get(__ret__, 'access'),
        admin_url=pulumi.get(__ret__, 'admin_url'),
        attributes=pulumi.get(__ret__, 'attributes'),
        authentication_flow_binding_overrides=pulumi.get(__ret__, 'authentication_flow_binding_overrides'),
        authorization_services_enabled=pulumi.get(__ret__, 'authorization_services_enabled'),
        authorization_settings=pulumi.get(__ret__, 'authorization_settings'),
        base_url=pulumi.get(__ret__, 'base_url'),
        bearer_only=pulumi.get(__ret__, 'bearer_only'),
        body=pulumi.get(__ret__, 'body'),
        client_authenticator_type=pulumi.get(__ret__, 'client_authenticator_type'),
        client_id=pulumi.get(__ret__, 'client_id'),
        consent_required=pulumi.get(__ret__, 'consent_required'),
        default_client_scopes=pulumi.get(__ret__, 'default_client_scopes'),
        default_roles=pulumi.get(__ret__, 'default_roles'),
        description=pulumi.get(__ret__, 'description'),
        direct_access_grants_enabled=pulumi.get(__ret__, 'direct_access_grants_enabled'),
        enabled=pulumi.get(__ret__, 'enabled'),
        frontchannel_logout=pulumi.get(__ret__, 'frontchannel_logout'),
        full_scope_allowed=pulumi.get(__ret__, 'full_scope_allowed'),
        id=pulumi.get(__ret__, 'id'),
        implicit_flow_enabled=pulumi.get(__ret__, 'implicit_flow_enabled'),
        name=pulumi.get(__ret__, 'name'),
        not_before=pulumi.get(__ret__, 'not_before'),
        optional_client_scopes=pulumi.get(__ret__, 'optional_client_scopes'),
        origin=pulumi.get(__ret__, 'origin'),
        protocol=pulumi.get(__ret__, 'protocol'),
        protocol_mappers=pulumi.get(__ret__, 'protocol_mappers'),
        public_client=pulumi.get(__ret__, 'public_client'),
        realm_id=pulumi.get(__ret__, 'realm_id'),
        redirect_uris=pulumi.get(__ret__, 'redirect_uris'),
        registered_nodes=pulumi.get(__ret__, 'registered_nodes'),
        registration_access_token=pulumi.get(__ret__, 'registration_access_token'),
        root_url=pulumi.get(__ret__, 'root_url'),
        secret=pulumi.get(__ret__, 'secret'),
        service_accounts_enabled=pulumi.get(__ret__, 'service_accounts_enabled'),
        standard_flow_enabled=pulumi.get(__ret__, 'standard_flow_enabled'),
        surrogate_auth_required=pulumi.get(__ret__, 'surrogate_auth_required'),
        web_origins=pulumi.get(__ret__, 'web_origins'))


@_utilities.lift_output_func(get_client_description_converter)
def get_client_description_converter_output(body: Optional[pulumi.Input[str]] = None,
                                            realm_id: Optional[pulumi.Input[str]] = None,
                                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClientDescriptionConverterResult]:
    """
    This data source uses the [ClientDescriptionConverter](https://www.keycloak.org/docs-api/6.0/javadocs/org/keycloak/exportimport/ClientDescriptionConverter.html) API to convert a generic client description into a Keycloak
    client. This data can then be used to manage the client within Keycloak.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.Realm("realm",
        realm="my-realm",
        enabled=True)
    saml_client = keycloak.get_client_description_converter_output(realm_id=realm.id,
        body=\"\"\"\\x09<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" validUntil="2021-04-17T12:41:46Z" cacheDuration="PT604800S" entityID="FakeEntityId">
        <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
            <md:KeyDescriptor use="signing">
    \\x09\\x09\\x09<ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
    \\x09\\x09\\x09\\x09<ds:X509Data>
    \\x09\\x09\\x09\\x09\\x09<ds:X509Certificate>MIICyDCCAjGgAwIBAgIBADANBgkqhkiG9w0BAQ0FADCBgDELMAkGA1UEBhMCdXMx
    \\x09\\x09\\x09\\x09\\x09CzAJBgNVBAgMAklBMSQwIgYDVQQKDBt0ZXJyYWZvcm0tcHJvdmlkZXIta2V5Y2xv
    \\x09\\x09\\x09\\x09\\x09YWsxHDAaBgNVBAMME21ycGFya2Vycy5naXRodWIuaW8xIDAeBgkqhkiG9w0BCQEW
    \\x09\\x09\\x09\\x09\\x09EW1pY2hhZWxAcGFya2VyLmdnMB4XDTE5MDEwODE0NDYzNloXDTI5MDEwNTE0NDYz
    \\x09\\x09\\x09\\x09\\x09NlowgYAxCzAJBgNVBAYTAnVzMQswCQYDVQQIDAJJQTEkMCIGA1UECgwbdGVycmFm
    \\x09\\x09\\x09\\x09\\x09b3JtLXByb3ZpZGVyLWtleWNsb2FrMRwwGgYDVQQDDBNtcnBhcmtlcnMuZ2l0aHVi
    \\x09\\x09\\x09\\x09\\x09LmlvMSAwHgYJKoZIhvcNAQkBFhFtaWNoYWVsQHBhcmtlci5nZzCBnzANBgkqhkiG
    \\x09\\x09\\x09\\x09\\x099w0BAQEFAAOBjQAwgYkCgYEAxuZny7uyYxGVPtpie14gNQC4tT9sAvO2sVNDhuoe
    \\x09\\x09\\x09\\x09\\x09qIKLRpNwkHnwQmwe5OxSh9K0BPHp/DNuuVWUqvo4tniEYn3jBr7FwLYLTKojQIxj
    \\x09\\x09\\x09\\x09\\x0953S1UTT9EXq3eP5HsHMD0QnTuca2nlNYUDBm6ud2fQj0Jt5qLx86EbEC28N56IRv
    \\x09\\x09\\x09\\x09\\x09GX8CAwEAAaNQME4wHQYDVR0OBBYEFMLnbQh77j7vhGTpAhKpDhCrBsPZMB8GA1Ud
    \\x09\\x09\\x09\\x09\\x09IwQYMBaAFMLnbQh77j7vhGTpAhKpDhCrBsPZMAwGA1UdEwQFMAMBAf8wDQYJKoZI
    \\x09\\x09\\x09\\x09\\x09hvcNAQENBQADgYEAB8wGrAQY0pAfwbnYSyBt4STbebeRTu1/q1ucfrtc3qsegcd5
    \\x09\\x09\\x09\\x09\\x09n01xTR+T2uZJwqHFPpFjr4IPORiHx3+4BWCweslPD53qBjKUPXcbMO1Revjef6Tj
    \\x09\\x09\\x09\\x09\\x09K3K0AuJ94fxgXVoT61Nzu/a6Lj6RhzU/Dao9mlSbJY+YSbm+ZBpsuRUQ84s=</ds:X509Certificate>
    \\x09\\x09\\x09\\x09</ds:X509Data>
    \\x09\\x09\\x09</ds:KeyInfo>
    \\x09\\x09</md:KeyDescriptor>
    \\x09\\x09<md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
            <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://localhost/acs/saml/" index="1"/>
        </md:SPSSODescriptor>
    </md:EntityDescriptor>
    \"\"\")
    saml_client_client = keycloak.saml.Client("saml_client",
        realm_id=realm.id,
        client_id=saml_client.client_id)
    ```
    <!--End PulumiCodeChooser -->


    :param str body: The body of the request to convert.
    :param str realm_id: The realm to use for the client description converter API call.
    """
    ...
