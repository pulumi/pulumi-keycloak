# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['RealmKeystoreRsaArgs', 'RealmKeystoreRsa']

@pulumi.input_type
class RealmKeystoreRsaArgs:
    def __init__(__self__, *,
                 certificate: pulumi.Input[builtins.str],
                 private_key: pulumi.Input[builtins.str],
                 realm_id: pulumi.Input[builtins.str],
                 active: Optional[pulumi.Input[builtins.bool]] = None,
                 algorithm: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 provider_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a RealmKeystoreRsa resource.
        :param pulumi.Input[builtins.str] certificate: X509 Certificate encoded in PEM format.
        :param pulumi.Input[builtins.str] private_key: Private RSA Key encoded in PEM format.
        :param pulumi.Input[builtins.str] realm_id: The realm this keystore exists in.
        :param pulumi.Input[builtins.bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[builtins.str] algorithm: Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
        :param pulumi.Input[builtins.bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[builtins.str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[builtins.int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[builtins.str] provider_id: Use `rsa` for signing keys, `rsa-enc` for encryption keys
        """
        pulumi.set(__self__, "certificate", certificate)
        pulumi.set(__self__, "private_key", private_key)
        pulumi.set(__self__, "realm_id", realm_id)
        if active is not None:
            pulumi.set(__self__, "active", active)
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Input[builtins.str]:
        """
        X509 Certificate encoded in PEM format.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[builtins.str]:
        """
        Private RSA Key encoded in PEM format.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[builtins.str]:
        """
        The realm this keystore exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When `false`, key in not used for signing. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When `false`, key is not accessible in this realm. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Display name of provider when linked in admin console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Priority for the provider. Defaults to `0`
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Use `rsa` for signing keys, `rsa-enc` for encryption keys
        """
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "provider_id", value)


@pulumi.input_type
class _RealmKeystoreRsaState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[builtins.bool]] = None,
                 algorithm: Optional[pulumi.Input[builtins.str]] = None,
                 certificate: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 private_key: Optional[pulumi.Input[builtins.str]] = None,
                 provider_id: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering RealmKeystoreRsa resources.
        :param pulumi.Input[builtins.bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[builtins.str] algorithm: Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
        :param pulumi.Input[builtins.str] certificate: X509 Certificate encoded in PEM format.
        :param pulumi.Input[builtins.bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[builtins.str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[builtins.int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[builtins.str] private_key: Private RSA Key encoded in PEM format.
        :param pulumi.Input[builtins.str] provider_id: Use `rsa` for signing keys, `rsa-enc` for encryption keys
        :param pulumi.Input[builtins.str] realm_id: The realm this keystore exists in.
        """
        if active is not None:
            pulumi.set(__self__, "active", active)
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When `false`, key in not used for signing. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        X509 Certificate encoded in PEM format.
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When `false`, key is not accessible in this realm. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Display name of provider when linked in admin console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        Priority for the provider. Defaults to `0`
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Private RSA Key encoded in PEM format.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Use `rsa` for signing keys, `rsa-enc` for encryption keys
        """
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "provider_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The realm this keystore exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "realm_id", value)


@pulumi.type_token("keycloak:index/realmKeystoreRsa:RealmKeystoreRsa")
class RealmKeystoreRsa(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[builtins.bool]] = None,
                 algorithm: Optional[pulumi.Input[builtins.str]] = None,
                 certificate: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 private_key: Optional[pulumi.Input[builtins.str]] = None,
                 provider_id: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Allows for creating and managing `rsa` Realm keystores within Keycloak.

        A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.

        ## Import

        Realm keys can be imported using realm name and keystore id, you can find it in web UI.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:index/realmKeystoreRsa:RealmKeystoreRsa keystore_rsa my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[builtins.str] algorithm: Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
        :param pulumi.Input[builtins.str] certificate: X509 Certificate encoded in PEM format.
        :param pulumi.Input[builtins.bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[builtins.str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[builtins.int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[builtins.str] private_key: Private RSA Key encoded in PEM format.
        :param pulumi.Input[builtins.str] provider_id: Use `rsa` for signing keys, `rsa-enc` for encryption keys
        :param pulumi.Input[builtins.str] realm_id: The realm this keystore exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RealmKeystoreRsaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing `rsa` Realm keystores within Keycloak.

        A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.

        ## Import

        Realm keys can be imported using realm name and keystore id, you can find it in web UI.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:index/realmKeystoreRsa:RealmKeystoreRsa keystore_rsa my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
        ```

        :param str resource_name: The name of the resource.
        :param RealmKeystoreRsaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RealmKeystoreRsaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[builtins.bool]] = None,
                 algorithm: Optional[pulumi.Input[builtins.str]] = None,
                 certificate: Optional[pulumi.Input[builtins.str]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 name: Optional[pulumi.Input[builtins.str]] = None,
                 priority: Optional[pulumi.Input[builtins.int]] = None,
                 private_key: Optional[pulumi.Input[builtins.str]] = None,
                 provider_id: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RealmKeystoreRsaArgs.__new__(RealmKeystoreRsaArgs)

            __props__.__dict__["active"] = active
            __props__.__dict__["algorithm"] = algorithm
            if certificate is None and not opts.urn:
                raise TypeError("Missing required property 'certificate'")
            __props__.__dict__["certificate"] = certificate
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["priority"] = priority
            if private_key is None and not opts.urn:
                raise TypeError("Missing required property 'private_key'")
            __props__.__dict__["private_key"] = private_key
            __props__.__dict__["provider_id"] = provider_id
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(RealmKeystoreRsa, __self__).__init__(
            'keycloak:index/realmKeystoreRsa:RealmKeystoreRsa',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[builtins.bool]] = None,
            algorithm: Optional[pulumi.Input[builtins.str]] = None,
            certificate: Optional[pulumi.Input[builtins.str]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            name: Optional[pulumi.Input[builtins.str]] = None,
            priority: Optional[pulumi.Input[builtins.int]] = None,
            private_key: Optional[pulumi.Input[builtins.str]] = None,
            provider_id: Optional[pulumi.Input[builtins.str]] = None,
            realm_id: Optional[pulumi.Input[builtins.str]] = None) -> 'RealmKeystoreRsa':
        """
        Get an existing RealmKeystoreRsa resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[builtins.str] algorithm: Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
        :param pulumi.Input[builtins.str] certificate: X509 Certificate encoded in PEM format.
        :param pulumi.Input[builtins.bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[builtins.str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[builtins.int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[builtins.str] private_key: Private RSA Key encoded in PEM format.
        :param pulumi.Input[builtins.str] provider_id: Use `rsa` for signing keys, `rsa-enc` for encryption keys
        :param pulumi.Input[builtins.str] realm_id: The realm this keystore exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RealmKeystoreRsaState.__new__(_RealmKeystoreRsaState)

        __props__.__dict__["active"] = active
        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["priority"] = priority
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["provider_id"] = provider_id
        __props__.__dict__["realm_id"] = realm_id
        return RealmKeystoreRsa(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        When `false`, key in not used for signing. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Intended algorithm for the key. Defaults to `RS256`. Use `RSA-OAEP` for encryption keys
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[builtins.str]:
        """
        X509 Certificate encoded in PEM format.
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        When `false`, key is not accessible in this realm. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[builtins.str]:
        """
        Display name of provider when linked in admin console.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[builtins.int]]:
        """
        Priority for the provider. Defaults to `0`
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[builtins.str]:
        """
        Private RSA Key encoded in PEM format.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Use `rsa` for signing keys, `rsa-enc` for encryption keys
        """
        return pulumi.get(self, "provider_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[builtins.str]:
        """
        The realm this keystore exists in.
        """
        return pulumi.get(self, "realm_id")

