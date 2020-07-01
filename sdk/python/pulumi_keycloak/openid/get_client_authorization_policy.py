# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetClientAuthorizationPolicyResult:
    """
    A collection of values returned by getClientAuthorizationPolicy.
    """
    def __init__(__self__, decision_strategy=None, id=None, logic=None, name=None, owner=None, policies=None, realm_id=None, resource_server_id=None, resources=None, scopes=None, type=None):
        if decision_strategy and not isinstance(decision_strategy, str):
            raise TypeError("Expected argument 'decision_strategy' to be a str")
        __self__.decision_strategy = decision_strategy
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if logic and not isinstance(logic, str):
            raise TypeError("Expected argument 'logic' to be a str")
        __self__.logic = logic
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        __self__.name = name
        if owner and not isinstance(owner, str):
            raise TypeError("Expected argument 'owner' to be a str")
        __self__.owner = owner
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        __self__.policies = policies
        if realm_id and not isinstance(realm_id, str):
            raise TypeError("Expected argument 'realm_id' to be a str")
        __self__.realm_id = realm_id
        if resource_server_id and not isinstance(resource_server_id, str):
            raise TypeError("Expected argument 'resource_server_id' to be a str")
        __self__.resource_server_id = resource_server_id
        if resources and not isinstance(resources, list):
            raise TypeError("Expected argument 'resources' to be a list")
        __self__.resources = resources
        if scopes and not isinstance(scopes, list):
            raise TypeError("Expected argument 'scopes' to be a list")
        __self__.scopes = scopes
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        __self__.type = type
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

def get_client_authorization_policy(logic=None,name=None,realm_id=None,resource_server_id=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['logic'] = logic
    __args__['name'] = name
    __args__['realmId'] = realm_id
    __args__['resourceServerId'] = resource_server_id
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('keycloak:openid/getClientAuthorizationPolicy:getClientAuthorizationPolicy', __args__, opts=opts).value

    return AwaitableGetClientAuthorizationPolicyResult(
        decision_strategy=__ret__.get('decisionStrategy'),
        id=__ret__.get('id'),
        logic=__ret__.get('logic'),
        name=__ret__.get('name'),
        owner=__ret__.get('owner'),
        policies=__ret__.get('policies'),
        realm_id=__ret__.get('realmId'),
        resource_server_id=__ret__.get('resourceServerId'),
        resources=__ret__.get('resources'),
        scopes=__ret__.get('scopes'),
        type=__ret__.get('type'))
