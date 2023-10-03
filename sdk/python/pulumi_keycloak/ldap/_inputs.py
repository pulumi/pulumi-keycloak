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
    'UserFederationCacheArgs',
    'UserFederationKerberosArgs',
]

@pulumi.input_type
class UserFederationCacheArgs:
    def __init__(__self__, *,
                 eviction_day: Optional[pulumi.Input[int]] = None,
                 eviction_hour: Optional[pulumi.Input[int]] = None,
                 eviction_minute: Optional[pulumi.Input[int]] = None,
                 max_lifespan: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[int] eviction_day: Day of the week the entry will become invalid on
        :param pulumi.Input[int] eviction_hour: Hour of day the entry will become invalid on.
        :param pulumi.Input[int] eviction_minute: Minute of day the entry will become invalid on.
        :param pulumi.Input[str] max_lifespan: Max lifespan of cache entry (duration string).
        :param pulumi.Input[str] policy: Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
        """
        UserFederationCacheArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            eviction_day=eviction_day,
            eviction_hour=eviction_hour,
            eviction_minute=eviction_minute,
            max_lifespan=max_lifespan,
            policy=policy,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             eviction_day: Optional[pulumi.Input[int]] = None,
             eviction_hour: Optional[pulumi.Input[int]] = None,
             eviction_minute: Optional[pulumi.Input[int]] = None,
             max_lifespan: Optional[pulumi.Input[str]] = None,
             policy: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        if eviction_day is not None:
            _setter("eviction_day", eviction_day)
        if eviction_hour is not None:
            _setter("eviction_hour", eviction_hour)
        if eviction_minute is not None:
            _setter("eviction_minute", eviction_minute)
        if max_lifespan is not None:
            _setter("max_lifespan", max_lifespan)
        if policy is not None:
            _setter("policy", policy)

    @property
    @pulumi.getter(name="evictionDay")
    def eviction_day(self) -> Optional[pulumi.Input[int]]:
        """
        Day of the week the entry will become invalid on
        """
        return pulumi.get(self, "eviction_day")

    @eviction_day.setter
    def eviction_day(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "eviction_day", value)

    @property
    @pulumi.getter(name="evictionHour")
    def eviction_hour(self) -> Optional[pulumi.Input[int]]:
        """
        Hour of day the entry will become invalid on.
        """
        return pulumi.get(self, "eviction_hour")

    @eviction_hour.setter
    def eviction_hour(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "eviction_hour", value)

    @property
    @pulumi.getter(name="evictionMinute")
    def eviction_minute(self) -> Optional[pulumi.Input[int]]:
        """
        Minute of day the entry will become invalid on.
        """
        return pulumi.get(self, "eviction_minute")

    @eviction_minute.setter
    def eviction_minute(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "eviction_minute", value)

    @property
    @pulumi.getter(name="maxLifespan")
    def max_lifespan(self) -> Optional[pulumi.Input[str]]:
        """
        Max lifespan of cache entry (duration string).
        """
        return pulumi.get(self, "max_lifespan")

    @max_lifespan.setter
    def max_lifespan(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_lifespan", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


@pulumi.input_type
class UserFederationKerberosArgs:
    def __init__(__self__, *,
                 kerberos_realm: pulumi.Input[str],
                 key_tab: pulumi.Input[str],
                 server_principal: pulumi.Input[str],
                 use_kerberos_for_password_authentication: Optional[pulumi.Input[bool]] = None):
        """
        :param pulumi.Input[str] kerberos_realm: The name of the kerberos realm, e.g. FOO.LOCAL.
        :param pulumi.Input[str] key_tab: Path to the kerberos keytab file on the server with credentials of the service principal.
        :param pulumi.Input[str] server_principal: The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
        :param pulumi.Input[bool] use_kerberos_for_password_authentication: Use kerberos login module instead of ldap service api. Defaults to `false`.
        """
        UserFederationKerberosArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            kerberos_realm=kerberos_realm,
            key_tab=key_tab,
            server_principal=server_principal,
            use_kerberos_for_password_authentication=use_kerberos_for_password_authentication,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             kerberos_realm: pulumi.Input[str],
             key_tab: pulumi.Input[str],
             server_principal: pulumi.Input[str],
             use_kerberos_for_password_authentication: Optional[pulumi.Input[bool]] = None,
             opts: Optional[pulumi.ResourceOptions]=None):
        _setter("kerberos_realm", kerberos_realm)
        _setter("key_tab", key_tab)
        _setter("server_principal", server_principal)
        if use_kerberos_for_password_authentication is not None:
            _setter("use_kerberos_for_password_authentication", use_kerberos_for_password_authentication)

    @property
    @pulumi.getter(name="kerberosRealm")
    def kerberos_realm(self) -> pulumi.Input[str]:
        """
        The name of the kerberos realm, e.g. FOO.LOCAL.
        """
        return pulumi.get(self, "kerberos_realm")

    @kerberos_realm.setter
    def kerberos_realm(self, value: pulumi.Input[str]):
        pulumi.set(self, "kerberos_realm", value)

    @property
    @pulumi.getter(name="keyTab")
    def key_tab(self) -> pulumi.Input[str]:
        """
        Path to the kerberos keytab file on the server with credentials of the service principal.
        """
        return pulumi.get(self, "key_tab")

    @key_tab.setter
    def key_tab(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_tab", value)

    @property
    @pulumi.getter(name="serverPrincipal")
    def server_principal(self) -> pulumi.Input[str]:
        """
        The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
        """
        return pulumi.get(self, "server_principal")

    @server_principal.setter
    def server_principal(self, value: pulumi.Input[str]):
        pulumi.set(self, "server_principal", value)

    @property
    @pulumi.getter(name="useKerberosForPasswordAuthentication")
    def use_kerberos_for_password_authentication(self) -> Optional[pulumi.Input[bool]]:
        """
        Use kerberos login module instead of ldap service api. Defaults to `false`.
        """
        return pulumi.get(self, "use_kerberos_for_password_authentication")

    @use_kerberos_for_password_authentication.setter
    def use_kerberos_for_password_authentication(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_kerberos_for_password_authentication", value)


