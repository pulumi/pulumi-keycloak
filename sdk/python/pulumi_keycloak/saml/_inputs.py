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


