# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'ClientAuthenticationFlowBindingOverrides',
    'ClientAuthorization',
    'ClientGroupPolicyGroup',
    'ClientRolePolicyRole',
    'GetClientAuthenticationFlowBindingOverrideResult',
    'GetClientAuthorizationResult',
    'GetClientServiceAccountUserFederatedIdentityResult',
]

@pulumi.output_type
class ClientAuthenticationFlowBindingOverrides(dict):
    def __init__(__self__, *,
                 browser_id: Optional[str] = None,
                 direct_grant_id: Optional[str] = None):
        """
        :param str browser_id: Browser flow id, (flow needs to exist)
        :param str direct_grant_id: Direct grant flow id (flow needs to exist)
        """
        if browser_id is not None:
            pulumi.set(__self__, "browser_id", browser_id)
        if direct_grant_id is not None:
            pulumi.set(__self__, "direct_grant_id", direct_grant_id)

    @property
    @pulumi.getter(name="browserId")
    def browser_id(self) -> Optional[str]:
        """
        Browser flow id, (flow needs to exist)
        """
        return pulumi.get(self, "browser_id")

    @property
    @pulumi.getter(name="directGrantId")
    def direct_grant_id(self) -> Optional[str]:
        """
        Direct grant flow id (flow needs to exist)
        """
        return pulumi.get(self, "direct_grant_id")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClientAuthorization(dict):
    def __init__(__self__, *,
                 policy_enforcement_mode: str,
                 allow_remote_resource_management: Optional[bool] = None,
                 keep_defaults: Optional[bool] = None):
        """
        :param str policy_enforcement_mode: Dictates how policies are enforced when evaluating authorization requests. Can be one of `ENFORCING`, `PERMISSIVE`, or `DISABLED`.
        :param bool allow_remote_resource_management: When `true`, resources can be managed remotely by the resource server. Defaults to `false`.
        :param bool keep_defaults: When `true`, defaults set by Keycloak will be respected. Defaults to `false`.
        """
        pulumi.set(__self__, "policy_enforcement_mode", policy_enforcement_mode)
        if allow_remote_resource_management is not None:
            pulumi.set(__self__, "allow_remote_resource_management", allow_remote_resource_management)
        if keep_defaults is not None:
            pulumi.set(__self__, "keep_defaults", keep_defaults)

    @property
    @pulumi.getter(name="policyEnforcementMode")
    def policy_enforcement_mode(self) -> str:
        """
        Dictates how policies are enforced when evaluating authorization requests. Can be one of `ENFORCING`, `PERMISSIVE`, or `DISABLED`.
        """
        return pulumi.get(self, "policy_enforcement_mode")

    @property
    @pulumi.getter(name="allowRemoteResourceManagement")
    def allow_remote_resource_management(self) -> Optional[bool]:
        """
        When `true`, resources can be managed remotely by the resource server. Defaults to `false`.
        """
        return pulumi.get(self, "allow_remote_resource_management")

    @property
    @pulumi.getter(name="keepDefaults")
    def keep_defaults(self) -> Optional[bool]:
        """
        When `true`, defaults set by Keycloak will be respected. Defaults to `false`.
        """
        return pulumi.get(self, "keep_defaults")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClientGroupPolicyGroup(dict):
    def __init__(__self__, *,
                 extend_children: bool,
                 id: str,
                 path: str):
        pulumi.set(__self__, "extend_children", extend_children)
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter(name="extendChildren")
    def extend_children(self) -> bool:
        return pulumi.get(self, "extend_children")

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ClientRolePolicyRole(dict):
    def __init__(__self__, *,
                 id: str,
                 required: bool):
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "required", required)

    @property
    @pulumi.getter
    def id(self) -> str:
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def required(self) -> bool:
        return pulumi.get(self, "required")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetClientAuthenticationFlowBindingOverrideResult(dict):
    def __init__(__self__, *,
                 browser_id: str,
                 direct_grant_id: str):
        pulumi.set(__self__, "browser_id", browser_id)
        pulumi.set(__self__, "direct_grant_id", direct_grant_id)

    @property
    @pulumi.getter(name="browserId")
    def browser_id(self) -> str:
        return pulumi.get(self, "browser_id")

    @property
    @pulumi.getter(name="directGrantId")
    def direct_grant_id(self) -> str:
        return pulumi.get(self, "direct_grant_id")


@pulumi.output_type
class GetClientAuthorizationResult(dict):
    def __init__(__self__, *,
                 allow_remote_resource_management: bool,
                 keep_defaults: bool,
                 policy_enforcement_mode: str):
        pulumi.set(__self__, "allow_remote_resource_management", allow_remote_resource_management)
        pulumi.set(__self__, "keep_defaults", keep_defaults)
        pulumi.set(__self__, "policy_enforcement_mode", policy_enforcement_mode)

    @property
    @pulumi.getter(name="allowRemoteResourceManagement")
    def allow_remote_resource_management(self) -> bool:
        return pulumi.get(self, "allow_remote_resource_management")

    @property
    @pulumi.getter(name="keepDefaults")
    def keep_defaults(self) -> bool:
        return pulumi.get(self, "keep_defaults")

    @property
    @pulumi.getter(name="policyEnforcementMode")
    def policy_enforcement_mode(self) -> str:
        return pulumi.get(self, "policy_enforcement_mode")


@pulumi.output_type
class GetClientServiceAccountUserFederatedIdentityResult(dict):
    def __init__(__self__, *,
                 identity_provider: str,
                 user_id: str,
                 user_name: str):
        pulumi.set(__self__, "identity_provider", identity_provider)
        pulumi.set(__self__, "user_id", user_id)
        pulumi.set(__self__, "user_name", user_name)

    @property
    @pulumi.getter(name="identityProvider")
    def identity_provider(self) -> str:
        return pulumi.get(self, "identity_provider")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> str:
        return pulumi.get(self, "user_id")

    @property
    @pulumi.getter(name="userName")
    def user_name(self) -> str:
        return pulumi.get(self, "user_name")

