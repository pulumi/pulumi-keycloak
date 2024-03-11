# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'ClientAuthenticationFlowBindingOverrides',
    'ClientAuthorization',
    'ClientGroupPolicyGroup',
    'ClientPermissionsConfigureScope',
    'ClientPermissionsManageScope',
    'ClientPermissionsMapRolesClientScopeScope',
    'ClientPermissionsMapRolesCompositeScope',
    'ClientPermissionsMapRolesScope',
    'ClientPermissionsTokenExchangeScope',
    'ClientPermissionsViewScope',
    'ClientRolePolicyRole',
    'GetClientAuthenticationFlowBindingOverrideResult',
    'GetClientAuthorizationResult',
    'GetClientServiceAccountUserFederatedIdentityResult',
]

@pulumi.output_type
class ClientAuthenticationFlowBindingOverrides(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "browserId":
            suggest = "browser_id"
        elif key == "directGrantId":
            suggest = "direct_grant_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientAuthenticationFlowBindingOverrides. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientAuthenticationFlowBindingOverrides.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientAuthenticationFlowBindingOverrides.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 browser_id: Optional[str] = None,
                 direct_grant_id: Optional[str] = None):
        if browser_id is not None:
            pulumi.set(__self__, "browser_id", browser_id)
        if direct_grant_id is not None:
            pulumi.set(__self__, "direct_grant_id", direct_grant_id)

    @property
    @pulumi.getter(name="browserId")
    def browser_id(self) -> Optional[str]:
        return pulumi.get(self, "browser_id")

    @property
    @pulumi.getter(name="directGrantId")
    def direct_grant_id(self) -> Optional[str]:
        return pulumi.get(self, "direct_grant_id")


@pulumi.output_type
class ClientAuthorization(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "policyEnforcementMode":
            suggest = "policy_enforcement_mode"
        elif key == "allowRemoteResourceManagement":
            suggest = "allow_remote_resource_management"
        elif key == "decisionStrategy":
            suggest = "decision_strategy"
        elif key == "keepDefaults":
            suggest = "keep_defaults"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientAuthorization. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientAuthorization.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientAuthorization.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 policy_enforcement_mode: str,
                 allow_remote_resource_management: Optional[bool] = None,
                 decision_strategy: Optional[str] = None,
                 keep_defaults: Optional[bool] = None):
        pulumi.set(__self__, "policy_enforcement_mode", policy_enforcement_mode)
        if allow_remote_resource_management is not None:
            pulumi.set(__self__, "allow_remote_resource_management", allow_remote_resource_management)
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if keep_defaults is not None:
            pulumi.set(__self__, "keep_defaults", keep_defaults)

    @property
    @pulumi.getter(name="policyEnforcementMode")
    def policy_enforcement_mode(self) -> str:
        return pulumi.get(self, "policy_enforcement_mode")

    @property
    @pulumi.getter(name="allowRemoteResourceManagement")
    def allow_remote_resource_management(self) -> Optional[bool]:
        return pulumi.get(self, "allow_remote_resource_management")

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter(name="keepDefaults")
    def keep_defaults(self) -> Optional[bool]:
        return pulumi.get(self, "keep_defaults")


@pulumi.output_type
class ClientGroupPolicyGroup(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "extendChildren":
            suggest = "extend_children"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientGroupPolicyGroup. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientGroupPolicyGroup.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientGroupPolicyGroup.__key_warning(key)
        return super().get(key, default)

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


@pulumi.output_type
class ClientPermissionsConfigureScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "decisionStrategy":
            suggest = "decision_strategy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientPermissionsConfigureScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientPermissionsConfigureScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientPermissionsConfigureScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 decision_strategy: Optional[str] = None,
                 description: Optional[str] = None,
                 policies: Optional[Sequence[str]] = None):
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "policies")


@pulumi.output_type
class ClientPermissionsManageScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "decisionStrategy":
            suggest = "decision_strategy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientPermissionsManageScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientPermissionsManageScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientPermissionsManageScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 decision_strategy: Optional[str] = None,
                 description: Optional[str] = None,
                 policies: Optional[Sequence[str]] = None):
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "policies")


@pulumi.output_type
class ClientPermissionsMapRolesClientScopeScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "decisionStrategy":
            suggest = "decision_strategy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientPermissionsMapRolesClientScopeScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientPermissionsMapRolesClientScopeScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientPermissionsMapRolesClientScopeScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 decision_strategy: Optional[str] = None,
                 description: Optional[str] = None,
                 policies: Optional[Sequence[str]] = None):
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "policies")


@pulumi.output_type
class ClientPermissionsMapRolesCompositeScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "decisionStrategy":
            suggest = "decision_strategy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientPermissionsMapRolesCompositeScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientPermissionsMapRolesCompositeScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientPermissionsMapRolesCompositeScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 decision_strategy: Optional[str] = None,
                 description: Optional[str] = None,
                 policies: Optional[Sequence[str]] = None):
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "policies")


@pulumi.output_type
class ClientPermissionsMapRolesScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "decisionStrategy":
            suggest = "decision_strategy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientPermissionsMapRolesScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientPermissionsMapRolesScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientPermissionsMapRolesScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 decision_strategy: Optional[str] = None,
                 description: Optional[str] = None,
                 policies: Optional[Sequence[str]] = None):
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "policies")


@pulumi.output_type
class ClientPermissionsTokenExchangeScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "decisionStrategy":
            suggest = "decision_strategy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientPermissionsTokenExchangeScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientPermissionsTokenExchangeScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientPermissionsTokenExchangeScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 decision_strategy: Optional[str] = None,
                 description: Optional[str] = None,
                 policies: Optional[Sequence[str]] = None):
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "policies")


@pulumi.output_type
class ClientPermissionsViewScope(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "decisionStrategy":
            suggest = "decision_strategy"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ClientPermissionsViewScope. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ClientPermissionsViewScope.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ClientPermissionsViewScope.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 decision_strategy: Optional[str] = None,
                 description: Optional[str] = None,
                 policies: Optional[Sequence[str]] = None):
        if decision_strategy is not None:
            pulumi.set(__self__, "decision_strategy", decision_strategy)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[str]:
        return pulumi.get(self, "decision_strategy")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def policies(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "policies")


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
                 decision_strategy: str,
                 keep_defaults: bool,
                 policy_enforcement_mode: str):
        pulumi.set(__self__, "allow_remote_resource_management", allow_remote_resource_management)
        pulumi.set(__self__, "decision_strategy", decision_strategy)
        pulumi.set(__self__, "keep_defaults", keep_defaults)
        pulumi.set(__self__, "policy_enforcement_mode", policy_enforcement_mode)

    @property
    @pulumi.getter(name="allowRemoteResourceManagement")
    def allow_remote_resource_management(self) -> bool:
        return pulumi.get(self, "allow_remote_resource_management")

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> str:
        return pulumi.get(self, "decision_strategy")

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


