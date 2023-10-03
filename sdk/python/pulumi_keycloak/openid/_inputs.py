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
    'ClientAuthenticationFlowBindingOverridesArgs',
    'ClientAuthorizationArgs',
    'ClientGroupPolicyGroupArgs',
    'ClientPermissionsConfigureScopeArgs',
    'ClientPermissionsManageScopeArgs',
    'ClientPermissionsMapRolesClientScopeScopeArgs',
    'ClientPermissionsMapRolesCompositeScopeArgs',
    'ClientPermissionsMapRolesScopeArgs',
    'ClientPermissionsTokenExchangeScopeArgs',
    'ClientPermissionsViewScopeArgs',
    'ClientRolePolicyRoleArgs',
]

@pulumi.input_type
class ClientAuthenticationFlowBindingOverridesArgs:
    def __init__(__self__, *,
                 browser_id: Optional[pulumi.Input[str]] = None,
                 direct_grant_id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] browser_id: Browser flow id, (flow needs to exist)
        :param pulumi.Input[str] direct_grant_id: Direct grant flow id (flow needs to exist)
        """
        ClientAuthenticationFlowBindingOverridesArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            browser_id=browser_id,
            direct_grant_id=direct_grant_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             browser_id: Optional[pulumi.Input[str]] = None,
             direct_grant_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if browser_id is not None:
            _setter("browser_id", browser_id)
        if direct_grant_id is not None:
            _setter("direct_grant_id", direct_grant_id)

    @property
    @pulumi.getter(name="browserId")
    def browser_id(self) -> Optional[pulumi.Input[str]]:
        """
        Browser flow id, (flow needs to exist)
        """
        return pulumi.get(self, "browser_id")

    @browser_id.setter
    def browser_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "browser_id", value)

    @property
    @pulumi.getter(name="directGrantId")
    def direct_grant_id(self) -> Optional[pulumi.Input[str]]:
        """
        Direct grant flow id (flow needs to exist)
        """
        return pulumi.get(self, "direct_grant_id")

    @direct_grant_id.setter
    def direct_grant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "direct_grant_id", value)


@pulumi.input_type
class ClientAuthorizationArgs:
    def __init__(__self__, *,
                 policy_enforcement_mode: pulumi.Input[str],
                 allow_remote_resource_management: Optional[pulumi.Input[bool]] = None,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 keep_defaults: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] policy_enforcement_mode: Dictates how policies are enforced when evaluating authorization requests. Can be one of `ENFORCING`, `PERMISSIVE`, or `DISABLED`.
        :param pulumi.Input[bool] allow_remote_resource_management: When `true`, resources can be managed remotely by the resource server. Defaults to `false`.
        :param pulumi.Input[str] decision_strategy: Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
        :param pulumi.Input[bool] keep_defaults: When `true`, defaults set by Keycloak will be respected. Defaults to `false`.
        """
        ClientAuthorizationArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            policy_enforcement_mode=policy_enforcement_mode,
            allow_remote_resource_management=allow_remote_resource_management,
            decision_strategy=decision_strategy,
            keep_defaults=keep_defaults,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             policy_enforcement_mode: pulumi.Input[str],
             allow_remote_resource_management: Optional[pulumi.Input[bool]] = None,
             decision_strategy: Optional[pulumi.Input[str]] = None,
             keep_defaults: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("policy_enforcement_mode", policy_enforcement_mode)
        if allow_remote_resource_management is not None:
            _setter("allow_remote_resource_management", allow_remote_resource_management)
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if keep_defaults is not None:
            _setter("keep_defaults", keep_defaults)

    @property
    @pulumi.getter(name="policyEnforcementMode")
    def policy_enforcement_mode(self) -> pulumi.Input[str]:
        """
        Dictates how policies are enforced when evaluating authorization requests. Can be one of `ENFORCING`, `PERMISSIVE`, or `DISABLED`.
        """
        return pulumi.get(self, "policy_enforcement_mode")

    @policy_enforcement_mode.setter
    def policy_enforcement_mode(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy_enforcement_mode", value)

    @property
    @pulumi.getter(name="allowRemoteResourceManagement")
    def allow_remote_resource_management(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, resources can be managed remotely by the resource server. Defaults to `false`.
        """
        return pulumi.get(self, "allow_remote_resource_management")

    @allow_remote_resource_management.setter
    def allow_remote_resource_management(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_remote_resource_management", value)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        """
        Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
        """
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter(name="keepDefaults")
    def keep_defaults(self) -> Optional[pulumi.Input[bool]]:
        """
        When `true`, defaults set by Keycloak will be respected. Defaults to `false`.
        """
        return pulumi.get(self, "keep_defaults")

    @keep_defaults.setter
    def keep_defaults(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "keep_defaults", value)


@pulumi.input_type
class ClientGroupPolicyGroupArgs:
    def __init__(__self__, *,
                 extend_children: pulumi.Input[bool],
                 id: pulumi.Input[str],
                 path: pulumi.Input[str]):
        ClientGroupPolicyGroupArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            extend_children=extend_children,
            id=id,
            path=path,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             extend_children: pulumi.Input[bool],
             id: pulumi.Input[str],
             path: pulumi.Input[str],
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("extend_children", extend_children)
        _setter("id", id)
        _setter("path", path)

    @property
    @pulumi.getter(name="extendChildren")
    def extend_children(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "extend_children")

    @extend_children.setter
    def extend_children(self, value: pulumi.Input[bool]):
        pulumi.set(self, "extend_children", value)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)


@pulumi.input_type
class ClientPermissionsConfigureScopeArgs:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        ClientPermissionsConfigureScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            decision_strategy=decision_strategy,
            description=description,
            policies=policies,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             decision_strategy: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if description is not None:
            _setter("description", description)
        if policies is not None:
            _setter("policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class ClientPermissionsManageScopeArgs:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        ClientPermissionsManageScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            decision_strategy=decision_strategy,
            description=description,
            policies=policies,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             decision_strategy: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if description is not None:
            _setter("description", description)
        if policies is not None:
            _setter("policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class ClientPermissionsMapRolesClientScopeScopeArgs:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        ClientPermissionsMapRolesClientScopeScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            decision_strategy=decision_strategy,
            description=description,
            policies=policies,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             decision_strategy: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if description is not None:
            _setter("description", description)
        if policies is not None:
            _setter("policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class ClientPermissionsMapRolesCompositeScopeArgs:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        ClientPermissionsMapRolesCompositeScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            decision_strategy=decision_strategy,
            description=description,
            policies=policies,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             decision_strategy: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if description is not None:
            _setter("description", description)
        if policies is not None:
            _setter("policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class ClientPermissionsMapRolesScopeArgs:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        ClientPermissionsMapRolesScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            decision_strategy=decision_strategy,
            description=description,
            policies=policies,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             decision_strategy: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if description is not None:
            _setter("description", description)
        if policies is not None:
            _setter("policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class ClientPermissionsTokenExchangeScopeArgs:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        ClientPermissionsTokenExchangeScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            decision_strategy=decision_strategy,
            description=description,
            policies=policies,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             decision_strategy: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if description is not None:
            _setter("description", description)
        if policies is not None:
            _setter("policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class ClientPermissionsViewScopeArgs:
    def __init__(__self__, *,
                 decision_strategy: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        ClientPermissionsViewScopeArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            decision_strategy=decision_strategy,
            description=description,
            policies=policies,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             decision_strategy: Optional[pulumi.Input[str]] = None,
             description: Optional[pulumi.Input[str]] = None,
             policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if decision_strategy is not None:
            _setter("decision_strategy", decision_strategy)
        if description is not None:
            _setter("description", description)
        if policies is not None:
            _setter("policies", policies)

    @property
    @pulumi.getter(name="decisionStrategy")
    def decision_strategy(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "decision_strategy")

    @decision_strategy.setter
    def decision_strategy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "decision_strategy", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class ClientRolePolicyRoleArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 required: pulumi.Input[bool]):
        ClientRolePolicyRoleArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            id=id,
            required=required,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             id: pulumi.Input[str],
             required: pulumi.Input[bool],
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("id", id)
        _setter("required", required)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def required(self) -> pulumi.Input[bool]:
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: pulumi.Input[bool]):
        pulumi.set(self, "required", value)


