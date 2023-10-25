# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetClientAuthorizationPolicyResult',
    'AwaitableGetClientAuthorizationPolicyResult',
    'get_client_authorization_policy',
    'get_client_authorization_policy_output',
]

@pulumi.output_type
class GetClientAuthorizationPolicyResult:
    """
    A collection of values returned by getClientAuthorizationPolicy.
    """
    def __init__(__self__, decision_strategy=None, id=None, logic=None, name=None, owner=None, policies=None, realm_id=None, resource_server_id=None, resources=None, scopes=None, type=None):
        if decision_strategy and not isinstance(decision_strategy, str):
            raise TypeError("Expected argument 'decision_strategy' to be a str")
        pulumi.set(__self__, "decision_strategy", decision_strategy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if logic and not isinstance(logic, str):
            raise TypeError("Expected argument 'logic' to be a str")
        pulumi.set(__self__, "logic", logic)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        pulumi.set(__self__, "owner", owner)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        pulumi.set(__self__, "realm_id", realm_id)
        if resource_server_id and not isinstance(resource_server_id, str):
            raise TypeError("Expected argument 'resource_server_id' to be a str")
        pulumi.set(__self__, "resource_server_id", resource_server_id)
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        pulumi.set(__self__, "resources", resources)
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        pulumi.set(__self__, "scopes", scopes)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> str:
        """
        (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
        """
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def logic(self) -> str:
        """
        (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
        """
        return pulumi.get(self, "logic")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> str:
        """
        (Computed) The ID of the owning resource. Applies to resources.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter
    def policies(self) -> Sequence[str]:
        """
        (Computed) The IDs of the policies that must be applied to scopes/resources for this policy/permission. Applies to policies and permissions.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> str:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="resourceServerId")
    def resource_server_id(self) -> str:
        return pulumi.get(self, "resource_server_id")

    @property
    @pulumi.getter
    def resources(self) -> Sequence[str]:
        """
        (Computed) The IDs of the resources that this permission applies to. Applies to resource-based permissions.
        """
        return pulumi.get(self, "resources")

    @property
    @pulumi.getter
    def scopes(self) -> Sequence[str]:
        """
        (Computed) The IDs of the scopes that this permission applies to. Applies to scope-based permissions.
        """
        return pulumi.get(self, "scopes")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        (Computed) The type of this policy / permission. For permissions, this could be `resource` or `scope`. For policies, this could be any type of authorization policy, such as `js`.
        """
        return pulumi.get(self, "type")


class AwaitableGetClientAuthorizationPolicyResult(GetClientAuthorizationPolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetClientAuthorizationPolicyResult(
            decision_strategy=self.decision_strategy,
            id=self.id,
            logic=self.logic,
            name=self.name,
            owner=self.owner,
            policies=self.policies,
            realm_id=self.realm_id,
            resource_server_id=self.resource_server_id,
            resources=self.resources,
            scopes=self.scopes,
            type=self.type)


def get_client_authorization_policy(name: Optional[str] = None,
                                    realm_id: Optional[str] = None,
                                    resource_server_id: Optional[str] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetClientAuthorizationPolicyResult:
    """
    This data source can be used to fetch policy and permission information for an OpenID client that has authorization enabled.


    :param str name: The name of the authorization policy.
    :param str realm_id: The realm this authorization policy exists within.
    :param str resource_server_id: The ID of the resource server this authorization policy is attached to.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['realmId'] = realm_id
    __args__['resourceServerId'] = resource_server_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('keycloak:openid/getClientAuthorizationPolicy:getClientAuthorizationPolicy', __args__, opts=opts, typ=GetClientAuthorizationPolicyResult).value

    return AwaitableGetClientAuthorizationPolicyResult(
        decision_strategy=pulumi.get(__ret__, 'decision_strategy'),
        id=pulumi.get(__ret__, 'id'),
        logic=pulumi.get(__ret__, 'logic'),
        name=pulumi.get(__ret__, 'name'),
        owner=pulumi.get(__ret__, 'owner'),
        policies=pulumi.get(__ret__, 'policies'),
        realm_id=pulumi.get(__ret__, 'realm_id'),
        resource_server_id=pulumi.get(__ret__, 'resource_server_id'),
        resources=pulumi.get(__ret__, 'resources'),
        scopes=pulumi.get(__ret__, 'scopes'),
        type=pulumi.get(__ret__, 'type'))


@_utilities.lift_output_func(get_client_authorization_policy)
def get_client_authorization_policy_output(name: Optional[pulumi.Input[str]] = None,
                                           realm_id: Optional[pulumi.Input[str]] = None,
                                           resource_server_id: Optional[pulumi.Input[str]] = None,
                                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetClientAuthorizationPolicyResult]:
    """
    This data source can be used to fetch policy and permission information for an OpenID client that has authorization enabled.


    :param str name: The name of the authorization policy.
    :param str realm_id: The realm this authorization policy exists within.
    :param str resource_server_id: The ID of the resource server this authorization policy is attached to.
    """
    ...
