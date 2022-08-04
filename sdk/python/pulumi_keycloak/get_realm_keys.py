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
    'GetRealmKeysResult',
    'AwaitableGetRealmKeysResult',
    'get_realm_keys',
    'get_realm_keys_output',
]

@pulumi.output_type
class GetRealmKeysResult:
    """
    A collection of values returned by getRealmKeys.
    """
    def __init__(__self__, algorithms=None, id=None, keys=None, realm_id=None, statuses=None):
        if algorithms and not isinstance(algorithms, list):
            raise TypeError("Expected argument 'algorithms' to be a list")
        pulumi.set(__self__, "algorithms", algorithms)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if statuses and not isinstance(statuses, list):
            raise TypeError("Expected argument 'statuses' to be a list")
        pulumi.set(__self__, "statuses", statuses)

    @property
    @pulumi.getter
    def algorithms(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "algorithms")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def keys(self) -> Sequence['outputs.GetRealmKeysKeyResult']:
        """
        (Computed) A list of keys that match the filter criteria. Each key has the following attributes:
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def statuses(self) -> Optional[Sequence[str]]:
        """
        Key status (string)
        """
        return pulumi.get(self, "statuses")


class AwaitableGetRealmKeysResult(GetRealmKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRealmKeysResult(
            algorithms=self.algorithms,
            id=self.id,
            keys=self.keys,
            realm_id=self.realm_id,
            statuses=self.statuses)


def get_realm_keys(algorithms: Optional[Sequence[str]] = None,
                   realm_id: Optional[str] = None,
                   statuses: Optional[Sequence[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRealmKeysResult:
    """
    Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.

    Remarks:

    - A key must meet all filter criteria
    - This data source may return more than one value.
    - If no key matches the filter criteria, then an error will be returned.


    :param Sequence[str] algorithms: When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
    :param str realm_id: The realm from which the keys will be retrieved.
    :param Sequence[str] statuses: When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
    """
    __args__ = dict()
    __args__['algorithms'] = algorithms
    __args__['realmId'] = realm_id
    __args__['statuses'] = statuses
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:index/getRealmKeys:getRealmKeys', __args__, opts=opts, typ=GetRealmKeysResult).value

    return AwaitableGetRealmKeysResult(
        algorithms=__ret__.algorithms,
        id=__ret__.id,
        keys=__ret__.keys,
        realm_id=__ret__.realm_id,
        statuses=__ret__.statuses)


@_utilities.lift_output_func(get_realm_keys)
def get_realm_keys_output(algorithms: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          realm_id: Optional[pulumi.Input[str]] = None,
                          statuses: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRealmKeysResult]:
    """
    Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.

    Remarks:

    - A key must meet all filter criteria
    - This data source may return more than one value.
    - If no key matches the filter criteria, then an error will be returned.


    :param Sequence[str] algorithms: When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
    :param str realm_id: The realm from which the keys will be retrieved.
    :param Sequence[str] statuses: When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
    """
    ...
