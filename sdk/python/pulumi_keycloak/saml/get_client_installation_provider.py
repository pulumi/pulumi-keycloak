# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'GetClientInstallationProviderResult',
    'AwaitableGetClientInstallationProviderResult',
    'get_client_installation_provider',
    'get_client_installation_provider_output',
]

@pulumi.output_type
class GetClientInstallationProviderResult:
    """
    A collection of values returned by getClientInstallationProvider.
    """
    def __init__(__self__, client_id=None, id=None, provider_id=None, realm_id=None, value=None):
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if provider_id and not isinstance(provider_id, str):
            raise TypeError("Expected argument 'provider_id' to be a str")
        pulumi.set(__self__, "provider_id", provider_id)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> builtins.str:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> builtins.str:
        return pulumi.get(self, "provider_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> builtins.str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def value(self) -> builtins.str:
        """
        (Computed) The returned document needed for SAML installation.
        """
        return pulumi.get(self, "value")


class AwaitableGetClientInstallationProviderResult(GetClientInstallationProviderResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientInstallationProviderResult(
            client_id=self.client_id,
            id=self.id,
            provider_id=self.provider_id,
            realm_id=self.realm_id,
            value=self.value)


def get_client_installation_provider(client_id: Optional[builtins.str] = None,
                                     provider_id: Optional[builtins.str] = None,
                                     realm_id: Optional[builtins.str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientInstallationProviderResult:
    """
    This data source can be used to retrieve Installation Provider of a SAML Client.


    :param builtins.str client_id: The ID of the SAML client. The `id` attribute of a `keycloak_client` resource should be used here.
    :param builtins.str provider_id: The ID of the SAML installation provider. Could be one of `saml-idp-descriptor`, `keycloak-saml`, `saml-sp-descriptor`, `keycloak-saml-subsystem`, `mod-auth-mellon`, etc.
    :param builtins.str realm_id: The realm that the SAML client exists within.
    """
    __args__ = dict()
    __args__['clientId'] = client_id
    __args__['providerId'] = provider_id
    __args__['realmId'] = realm_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:saml/getClientInstallationProvider:getClientInstallationProvider', __args__, opts=opts, typ=GetClientInstallationProviderResult).value

    return AwaitableGetClientInstallationProviderResult(
        client_id=pulumi.get(__ret__, 'client_id'),
        id=pulumi.get(__ret__, 'id'),
        provider_id=pulumi.get(__ret__, 'provider_id'),
        realm_id=pulumi.get(__ret__, 'realm_id'),
        value=pulumi.get(__ret__, 'value'))
def get_client_installation_provider_output(client_id: Optional[pulumi.Input[builtins.str]] = None,
                                            provider_id: Optional[pulumi.Input[builtins.str]] = None,
                                            realm_id: Optional[pulumi.Input[builtins.str]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetClientInstallationProviderResult]:
    """
    This data source can be used to retrieve Installation Provider of a SAML Client.


    :param builtins.str client_id: The ID of the SAML client. The `id` attribute of a `keycloak_client` resource should be used here.
    :param builtins.str provider_id: The ID of the SAML installation provider. Could be one of `saml-idp-descriptor`, `keycloak-saml`, `saml-sp-descriptor`, `keycloak-saml-subsystem`, `mod-auth-mellon`, etc.
    :param builtins.str realm_id: The realm that the SAML client exists within.
    """
    __args__ = dict()
    __args__['clientId'] = client_id
    __args__['providerId'] = provider_id
    __args__['realmId'] = realm_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('keycloak:saml/getClientInstallationProvider:getClientInstallationProvider', __args__, opts=opts, typ=GetClientInstallationProviderResult)
    return __ret__.apply(lambda __response__: GetClientInstallationProviderResult(
        client_id=pulumi.get(__response__, 'client_id'),
        id=pulumi.get(__response__, 'id'),
        provider_id=pulumi.get(__response__, 'provider_id'),
        realm_id=pulumi.get(__response__, 'realm_id'),
        value=pulumi.get(__response__, 'value')))
