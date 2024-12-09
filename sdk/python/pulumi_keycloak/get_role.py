# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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
from . import _utilities

__all__ = [
    'GetRoleResult',
    'AwaitableGetRoleResult',
    'get_role',
    'get_role_output',
]

@pulumi.output_type
class GetRoleResult:
    """
    A collection of values returned by getRole.
    """
    def __init__(__self__, attributes=None, client_id=None, composite_roles=None, description=None, id=None, name=None, realm_id=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if composite_roles and not isinstance(composite_roles, list):
            raise TypeError("Expected argument 'composite_roles' to be a list")
        pulumi.set(__self__, "composite_roles", composite_roles)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def attributes(self) -> Mapping[str, str]:
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[str]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="compositeRoles")
    def composite_roles(self) -> Sequence[str]:
        return pulumi.get(self, "composite_roles")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        (Computed) The description of the role.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")


class AwaitableGetRoleResult(GetRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleResult(
            attributes=self.attributes,
            client_id=self.client_id,
            composite_roles=self.composite_roles,
            description=self.description,
            id=self.id,
            name=self.name,
            realm_id=self.realm_id)


def get_role(client_id: Optional[str] = None,
             name: Optional[str] = None,
             realm_id: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleResult:
    """
    This data source can be used to fetch properties of a Keycloak role for
    usage with other resources, such as `GroupRoles`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.Realm("realm",
        realm="my-realm",
        enabled=True)
    offline_access = keycloak.get_role_output(realm_id=realm.id,
        name="offline_access")
    # use the data source
    group = keycloak.Group("group",
        realm_id=realm.id,
        name="group")
    group_roles = keycloak.GroupRoles("group_roles",
        realm_id=realm.id,
        group_id=group.id,
        role_ids=[offline_access.id])
    ```


    :param str client_id: When specified, this role is assumed to be a client role belonging to the client with the provided ID. The `id` attribute of a `keycloak_client` resource should be used here.
    :param str name: The name of the role.
    :param str realm_id: The realm this role exists within.
    """
    __args__ = dict()
    __args__['clientId'] = client_id
    __args__['name'] = name
    __args__['realmId'] = realm_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:index/getRole:getRole', __args__, opts=opts, typ=GetRoleResult).value

    return AwaitableGetRoleResult(
        attributes=pulumi.get(__ret__, 'attributes'),
        client_id=pulumi.get(__ret__, 'client_id'),
        composite_roles=pulumi.get(__ret__, 'composite_roles'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        realm_id=pulumi.get(__ret__, 'realm_id'))
def get_role_output(client_id: Optional[pulumi.Input[Optional[str]]] = None,
                    name: Optional[pulumi.Input[str]] = None,
                    realm_id: Optional[pulumi.Input[str]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRoleResult]:
    """
    This data source can be used to fetch properties of a Keycloak role for
    usage with other resources, such as `GroupRoles`.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_keycloak as keycloak

    realm = keycloak.Realm("realm",
        realm="my-realm",
        enabled=True)
    offline_access = keycloak.get_role_output(realm_id=realm.id,
        name="offline_access")
    # use the data source
    group = keycloak.Group("group",
        realm_id=realm.id,
        name="group")
    group_roles = keycloak.GroupRoles("group_roles",
        realm_id=realm.id,
        group_id=group.id,
        role_ids=[offline_access.id])
    ```


    :param str client_id: When specified, this role is assumed to be a client role belonging to the client with the provided ID. The `id` attribute of a `keycloak_client` resource should be used here.
    :param str name: The name of the role.
    :param str realm_id: The realm this role exists within.
    """
    __args__ = dict()
    __args__['clientId'] = client_id
    __args__['name'] = name
    __args__['realmId'] = realm_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('keycloak:index/getRole:getRole', __args__, opts=opts, typ=GetRoleResult)
    return __ret__.apply(lambda __response__: GetRoleResult(
        attributes=pulumi.get(__response__, 'attributes'),
        client_id=pulumi.get(__response__, 'client_id'),
        composite_roles=pulumi.get(__response__, 'composite_roles'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        realm_id=pulumi.get(__response__, 'realm_id')))
