# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClientInstallationProviderResult:
    """
    A collection of values returned by getClientInstallationProvider.
    """
    def __init__(__self__, client_id=None, id=None, provider_id=None, realm_id=None, value=None):
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        __self__.client_id = client_id
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if provider_id and not isinstance(provider_id, str):
            raise TypeError("Expected argument 'provider_id' to be a str")
        __self__.provider_id = provider_id
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        __self__.realm_id = realm_id
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        __self__.value = value
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

def get_client_installation_provider(client_id=None,provider_id=None,realm_id=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['clientId'] = client_id
    __args__['providerId'] = provider_id
    __args__['realmId'] = realm_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('keycloak:saml/getClientInstallationProvider:getClientInstallationProvider', __args__, opts=opts).value

    return AwaitableGetClientInstallationProviderResult(
        client_id=__ret__.get('clientId'),
        id=__ret__.get('id'),
        provider_id=__ret__.get('providerId'),
        realm_id=__ret__.get('realmId'),
        value=__ret__.get('value'))