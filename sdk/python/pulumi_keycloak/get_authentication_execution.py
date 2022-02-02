# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetAuthenticationExecutionResult',
    'AwaitableGetAuthenticationExecutionResult',
    'get_authentication_execution',
    'get_authentication_execution_output',
]

@pulumi.output_type
class GetAuthenticationExecutionResult:
    """
    A collection of values returned by getAuthenticationExecution.
    """
    def __init__(__self__, id=None, parent_flow_alias=None, provider_id=None, realm_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if parent_flow_alias and not isinstance(parent_flow_alias, str):
            raise TypeError("Expected argument 'parent_flow_alias' to be a str")
        pulumi.set(__self__, "parent_flow_alias", parent_flow_alias)
        if provider_id and not isinstance(provider_id, str):
            raise TypeError("Expected argument 'provider_id' to be a str")
        pulumi.set(__self__, "provider_id", provider_id)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="parentFlowAlias")
    def parent_flow_alias(self) -> str:
        return pulumi.get(self, "parent_flow_alias")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> str:
        return pulumi.get(self, "provider_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")


class AwaitableGetAuthenticationExecutionResult(GetAuthenticationExecutionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthenticationExecutionResult(
            id=self.id,
            parent_flow_alias=self.parent_flow_alias,
            provider_id=self.provider_id,
            realm_id=self.realm_id)


def get_authentication_execution(parent_flow_alias: Optional[str] = None,
                                 provider_id: Optional[str] = None,
                                 realm_id: Optional[str] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthenticationExecutionResult:
    """
    This data source can be used to fetch the ID of an authentication execution within Keycloak.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.Realm("realm",
        realm="my-realm",
        enabled=True)
    browser_auth_cookie = keycloak.get_authentication_execution_output(realm_id=realm.id,
        parent_flow_alias="browser",
        provider_id="auth-cookie")
    ```


    :param str parent_flow_alias: The alias of the flow this execution is attached to.
    :param str provider_id: The name of the provider. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools. This was previously known as the "authenticator".
    :param str realm_id: The realm the authentication execution exists in.
    """
    __args__ = dict()
    __args__['parentFlowAlias'] = parent_flow_alias
    __args__['providerId'] = provider_id
    __args__['realmId'] = realm_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('keycloak:index/getAuthenticationExecution:getAuthenticationExecution', __args__, opts=opts, typ=GetAuthenticationExecutionResult).value

    return AwaitableGetAuthenticationExecutionResult(
        id=__ret__.id,
        parent_flow_alias=__ret__.parent_flow_alias,
        provider_id=__ret__.provider_id,
        realm_id=__ret__.realm_id)


@_utilities.lift_output_func(get_authentication_execution)
def get_authentication_execution_output(parent_flow_alias: Optional[pulumi.Input[str]] = None,
                                        provider_id: Optional[pulumi.Input[str]] = None,
                                        realm_id: Optional[pulumi.Input[str]] = None,
                                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAuthenticationExecutionResult]:
    """
    This data source can be used to fetch the ID of an authentication execution within Keycloak.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.Realm("realm",
        realm="my-realm",
        enabled=True)
    browser_auth_cookie = keycloak.get_authentication_execution_output(realm_id=realm.id,
        parent_flow_alias="browser",
        provider_id="auth-cookie")
    ```


    :param str parent_flow_alias: The alias of the flow this execution is attached to.
    :param str provider_id: The name of the provider. This can be found by experimenting with the GUI and looking at HTTP requests within the network tab of your browser's development tools. This was previously known as the "authenticator".
    :param str realm_id: The realm the authentication execution exists in.
    """
    ...
