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
    'GetGroupResult',
    'AwaitableGetGroupResult',
    'get_group',
    'get_group_output',
]

@pulumi.output_type
class GetGroupResult:
    """
    A collection of values returned by getGroup.
    """
    def __init__(__self__, attributes=None, id=None, name=None, parent_id=None, path=None, realm_id=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def attributes(self) -> Mapping[str, Any]:
        return pulumi.get(self, "attributes")

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
    @pulumi.getter(name="parentId")
    def parent_id(self) -> str:
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")


class AwaitableGetGroupResult(GetGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetGroupResult(
            attributes=self.attributes,
            id=self.id,
            name=self.name,
            parent_id=self.parent_id,
            path=self.path,
            realm_id=self.realm_id)


def get_group(name: Optional[str] = None,
              realm_id: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetGroupResult:
    """
    This data source can be used to fetch properties of a Keycloak group for
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
    group = keycloak.get_group_output(realm_id=realm.id,
        name="group")
    group_roles = keycloak.GroupRoles("groupRoles",
        realm_id=realm.id,
        group_id=group.id,
        role_ids=[offline_access.id])
    ```


    :param str name: The name of the group. If there are multiple groups match `name`, the first result will be returned.
    :param str realm_id: The realm this group exists within.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['realmId'] = realm_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:index/getGroup:getGroup', __args__, opts=opts, typ=GetGroupResult).value

    return AwaitableGetGroupResult(
        attributes=pulumi.get(__ret__, 'attributes'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        parent_id=pulumi.get(__ret__, 'parent_id'),
        path=pulumi.get(__ret__, 'path'),
        realm_id=pulumi.get(__ret__, 'realm_id'))


@_utilities.lift_output_func(get_group)
def get_group_output(name: Optional[pulumi.Input[str]] = None,
                     realm_id: Optional[pulumi.Input[str]] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetGroupResult]:
    """
    This data source can be used to fetch properties of a Keycloak group for
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
    group = keycloak.get_group_output(realm_id=realm.id,
        name="group")
    group_roles = keycloak.GroupRoles("groupRoles",
        realm_id=realm.id,
        group_id=group.id,
        role_ids=[offline_access.id])
    ```


    :param str name: The name of the group. If there are multiple groups match `name`, the first result will be returned.
    :param str realm_id: The realm this group exists within.
    """
    ...
