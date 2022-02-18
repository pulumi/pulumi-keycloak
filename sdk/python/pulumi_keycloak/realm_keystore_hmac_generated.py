# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RealmKeystoreHmacGeneratedArgs', 'RealmKeystoreHmacGenerated']

@pulumi.input_type
class RealmKeystoreHmacGeneratedArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 active: Optional[pulumi.Input[bool]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 secret_size: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RealmKeystoreHmacGenerated resource.
        :param pulumi.Input[str] realm_id: The realm this keystore exists in.
        :param pulumi.Input[bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[str] algorithm: Intended algorithm for the key. Defaults to `HS256`
        :param pulumi.Input[bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[int] secret_size: Size in bytes for the generated secret. Defaults to `64`.
        """
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
        if secret_size is not None:
            pulumi.set(__self__, "secret_size", secret_size)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm this keystore exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        When `false`, key in not used for signing. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Intended algorithm for the key. Defaults to `HS256`
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `false`, key is not accessible in this realm. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of provider when linked in admin console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        Priority for the provider. Defaults to `0`
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="secretSize")
    def secret_size(self) -> Optional[pulumi.Input[int]]:
        """
        Size in bytes for the generated secret. Defaults to `64`.
        """
        return pulumi.get(self, "secret_size")

    @secret_size.setter
    def secret_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "secret_size", value)


@pulumi.input_type
class _RealmKeystoreHmacGeneratedState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 secret_size: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering RealmKeystoreHmacGenerated resources.
        :param pulumi.Input[bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[str] algorithm: Intended algorithm for the key. Defaults to `HS256`
        :param pulumi.Input[bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[str] realm_id: The realm this keystore exists in.
        :param pulumi.Input[int] secret_size: Size in bytes for the generated secret. Defaults to `64`.
        """
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
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if secret_size is not None:
            pulumi.set(__self__, "secret_size", secret_size)

    @property
    @pulumi.getter
    def active(self) -> Optional[pulumi.Input[bool]]:
        """
        When `false`, key in not used for signing. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @active.setter
    def active(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "active", value)

    @property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[str]]:
        """
        Intended algorithm for the key. Defaults to `HS256`
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "algorithm", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When `false`, key is not accessible in this realm. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Display name of provider when linked in admin console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        Priority for the provider. Defaults to `0`
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm this keystore exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="secretSize")
    def secret_size(self) -> Optional[pulumi.Input[int]]:
        """
        Size in bytes for the generated secret. Defaults to `64`.
        """
        return pulumi.get(self, "secret_size")

    @secret_size.setter
    def secret_size(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "secret_size", value)


class RealmKeystoreHmacGenerated(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 secret_size: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Allows for creating and managing `hmac-generated` Realm keystores within Keycloak.

        A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm", realm="my-realm")
        keystore_hmac_generated = keycloak.RealmKeystoreHmacGenerated("keystoreHmacGenerated",
            realm_id=realm.realm,
            enabled=True,
            active=True,
            priority=100,
            algorithm="HS256",
            secret_size=64)
        ```

        ## Import

        Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash

        ```sh
         $ pulumi import keycloak:index/realmKeystoreHmacGenerated:RealmKeystoreHmacGenerated keystore_hmac_generated my-realm/my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[str] algorithm: Intended algorithm for the key. Defaults to `HS256`
        :param pulumi.Input[bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[str] realm_id: The realm this keystore exists in.
        :param pulumi.Input[int] secret_size: Size in bytes for the generated secret. Defaults to `64`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RealmKeystoreHmacGeneratedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing `hmac-generated` Realm keystores within Keycloak.

        A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm", realm="my-realm")
        keystore_hmac_generated = keycloak.RealmKeystoreHmacGenerated("keystoreHmacGenerated",
            realm_id=realm.realm,
            enabled=True,
            active=True,
            priority=100,
            algorithm="HS256",
            secret_size=64)
        ```

        ## Import

        Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash

        ```sh
         $ pulumi import keycloak:index/realmKeystoreHmacGenerated:RealmKeystoreHmacGenerated keystore_hmac_generated my-realm/my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
        ```

        :param str resource_name: The name of the resource.
        :param RealmKeystoreHmacGeneratedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RealmKeystoreHmacGeneratedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 secret_size: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RealmKeystoreHmacGeneratedArgs.__new__(RealmKeystoreHmacGeneratedArgs)

            __props__.__dict__["active"] = active
            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["name"] = name
            __props__.__dict__["priority"] = priority
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            __props__.__dict__["secret_size"] = secret_size
        super(RealmKeystoreHmacGenerated, __self__).__init__(
            'keycloak:index/realmKeystoreHmacGenerated:RealmKeystoreHmacGenerated',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active: Optional[pulumi.Input[bool]] = None,
            algorithm: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            secret_size: Optional[pulumi.Input[int]] = None) -> 'RealmKeystoreHmacGenerated':
        """
        Get an existing RealmKeystoreHmacGenerated resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[str] algorithm: Intended algorithm for the key. Defaults to `HS256`
        :param pulumi.Input[bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[str] realm_id: The realm this keystore exists in.
        :param pulumi.Input[int] secret_size: Size in bytes for the generated secret. Defaults to `64`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RealmKeystoreHmacGeneratedState.__new__(_RealmKeystoreHmacGeneratedState)

        __props__.__dict__["active"] = active
        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["name"] = name
        __props__.__dict__["priority"] = priority
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["secret_size"] = secret_size
        return RealmKeystoreHmacGenerated(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def active(self) -> pulumi.Output[Optional[bool]]:
        """
        When `false`, key in not used for signing. Defaults to `true`.
        """
        return pulumi.get(self, "active")

    @property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[Optional[str]]:
        """
        Intended algorithm for the key. Defaults to `HS256`
        """
        return pulumi.get(self, "algorithm")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When `false`, key is not accessible in this realm. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Display name of provider when linked in admin console.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[Optional[int]]:
        """
        Priority for the provider. Defaults to `0`
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm this keystore exists in.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="secretSize")
    def secret_size(self) -> pulumi.Output[Optional[int]]:
        """
        Size in bytes for the generated secret. Defaults to `64`.
        """
        return pulumi.get(self, "secret_size")
