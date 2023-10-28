# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetUserResult',
    'AwaitableGetUserResult',
    'get_user',
    'get_user_output',
]

@pulumi.output_type
class GetUserResult:
    """
    A collection of values returned by getUser.
    """
    def __init__(__self__, attributes=None, email=None, email_verified=None, enabled=None, federated_identities=None, first_name=None, id=None, last_name=None, realm_id=None, username=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if email_verified and not isinstance(email_verified, bool):
            raise TypeError("Expected argument 'email_verified' to be a bool")
        pulumi.set(__self__, "email_verified", email_verified)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if federated_identities and not isinstance(federated_identities, list):
            raise TypeError("Expected argument 'federated_identities' to be a list")
        pulumi.set(__self__, "federated_identities", federated_identities)
        if first_name and not isinstance(first_name, str):
            raise TypeError("Expected argument 'first_name' to be a str")
        pulumi.set(__self__, "first_name", first_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_name and not isinstance(last_name, str):
            raise TypeError("Expected argument 'last_name' to be a str")
        pulumi.set(__self__, "last_name", last_name)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if username and not isinstance(username, str):
            raise TypeError("Expected argument 'username' to be a str")
        pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def attributes(self) -> Mapping[str, Any]:
        """
        (Computed) A map representing attributes for the user
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        (Computed) The user's email.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> bool:
        """
        (Computed) Whether the email address was validated or not. Default to `false`.
        """
        return pulumi.get(self, "email_verified")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        (Computed) When false, this user cannot log in. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="federatedIdentities")
    def federated_identities(self) -> Sequence[str]:
        """
        (Computed) The user's federated identities, if applicable. This block has the following schema:
        """
        return pulumi.get(self, "federated_identities")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> str:
        """
        (Computed) The user's first name.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> str:
        """
        (Computed) The user's last name.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def username(self) -> str:
        return pulumi.get(self, "username")


class AwaitableGetUserResult(GetUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserResult(
            attributes=self.attributes,
            email=self.email,
            email_verified=self.email_verified,
            enabled=self.enabled,
            federated_identities=self.federated_identities,
            first_name=self.first_name,
            id=self.id,
            last_name=self.last_name,
            realm_id=self.realm_id,
            username=self.username)


def get_user(realm_id: Optional[str] = None,
             username: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserResult:
    """
    This data source can be used to fetch properties of a user within Keycloak.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    master_realm = keycloak.get_realm(realm="master")
    default_admin_user = keycloak.get_user(realm_id=master_realm.id,
        username="keycloak")
    pulumi.export("keycloakUserId", default_admin_user.id)
    ```


    :param str realm_id: The realm this user belongs to.
    :param str username: The unique username of this user.
    """
    __args__ = dict()
    __args__['realmId'] = realm_id
    __args__['username'] = username
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:index/getUser:getUser', __args__, opts=opts, typ=GetUserResult).value

    return AwaitableGetUserResult(
        attributes=pulumi.get(__ret__, 'attributes'),
        email=pulumi.get(__ret__, 'email'),
        email_verified=pulumi.get(__ret__, 'email_verified'),
        enabled=pulumi.get(__ret__, 'enabled'),
        federated_identities=pulumi.get(__ret__, 'federated_identities'),
        first_name=pulumi.get(__ret__, 'first_name'),
        id=pulumi.get(__ret__, 'id'),
        last_name=pulumi.get(__ret__, 'last_name'),
        realm_id=pulumi.get(__ret__, 'realm_id'),
        username=pulumi.get(__ret__, 'username'))


@_utilities.lift_output_func(get_user)
def get_user_output(realm_id: Optional[pulumi.Input[str]] = None,
                    username: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserResult]:
    """
    This data source can be used to fetch properties of a user within Keycloak.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    master_realm = keycloak.get_realm(realm="master")
    default_admin_user = keycloak.get_user(realm_id=master_realm.id,
        username="keycloak")
    pulumi.export("keycloakUserId", default_admin_user.id)
    ```


    :param str realm_id: The realm this user belongs to.
    :param str username: The unique username of this user.
    """
    ...
