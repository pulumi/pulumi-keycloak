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
    'GetUserRealmRolesResult',
    'AwaitableGetUserRealmRolesResult',
    'get_user_realm_roles',
    'get_user_realm_roles_output',
]

@pulumi.output_type
class GetUserRealmRolesResult:
    """
    A collection of values returned by getUserRealmRoles.
    """
    def __init__(__self__, id=None, realm_id=None, role_names=None, user_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if role_names and not isinstance(role_names, list):
            raise TypeError("Expected argument 'role_names' to be a list")
        pulumi.set(__self__, "role_names", role_names)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="roleNames")
    def role_names(self) -> Sequence[str]:
        """
        (Computed) A list of realm roles that belong to this user.
        """
        return pulumi.get(self, "role_names")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        return pulumi.get(self, "user_id")


class AwaitableGetUserRealmRolesResult(GetUserRealmRolesResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetUserRealmRolesResult(
            id=self.id,
            realm_id=self.realm_id,
            role_names=self.role_names,
            user_id=self.user_id)


def get_user_realm_roles(realm_id: Optional[str] = None,
                         user_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetUserRealmRolesResult:
    """
    This data source can be used to fetch the realm roles of a user within Keycloak.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    master_realm = keycloak.get_realm(realm="master")
    default_admin_user = keycloak.get_user(realm_id=master_realm.id,
        username="keycloak")
    user_realm_roles = keycloak.get_user_realm_roles(realm_id=master_realm.id,
        user_id=default_admin_user.id)
    pulumi.export("keycloakUserRoleNames", user_realm_roles.role_names)
    ```
    <!--End PulumiCodeChooser -->


    :param str realm_id: The realm this user belongs to.
    :param str user_id: The ID of the user to query realm roles for.
    """
    __args__ = dict()
    __args__['realmId'] = realm_id
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:index/getUserRealmRoles:getUserRealmRoles', __args__, opts=opts, typ=GetUserRealmRolesResult).value

    return AwaitableGetUserRealmRolesResult(
        id=pulumi.get(__ret__, 'id'),
        realm_id=pulumi.get(__ret__, 'realm_id'),
        role_names=pulumi.get(__ret__, 'role_names'),
        user_id=pulumi.get(__ret__, 'user_id'))


@_utilities.lift_output_func(get_user_realm_roles)
def get_user_realm_roles_output(realm_id: Optional[pulumi.Input[str]] = None,
                                user_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetUserRealmRolesResult]:
    """
    This data source can be used to fetch the realm roles of a user within Keycloak.

    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    master_realm = keycloak.get_realm(realm="master")
    default_admin_user = keycloak.get_user(realm_id=master_realm.id,
        username="keycloak")
    user_realm_roles = keycloak.get_user_realm_roles(realm_id=master_realm.id,
        user_id=default_admin_user.id)
    pulumi.export("keycloakUserRoleNames", user_realm_roles.role_names)
    ```
    <!--End PulumiCodeChooser -->


    :param str realm_id: The realm this user belongs to.
    :param str user_id: The ID of the user to query realm roles for.
    """
    ...
