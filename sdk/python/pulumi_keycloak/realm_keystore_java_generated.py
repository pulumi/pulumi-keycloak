# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['RealmKeystoreJavaGeneratedArgs', 'RealmKeystoreJavaGenerated']

@pulumi.input_type
class RealmKeystoreJavaGeneratedArgs:
    def __init__(__self__, *,
                 key_alias: pulumi.Input[str],
                 key_password: pulumi.Input[str],
                 keystore: pulumi.Input[str],
                 keystore_password: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 active: Optional[pulumi.Input[bool]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a RealmKeystoreJavaGenerated resource.
        :param pulumi.Input[str] key_alias: Alias for the private key.
        :param pulumi.Input[str] key_password: Password for the private key.
        :param pulumi.Input[str] keystore: Path to keys file on keycloak instance.
        :param pulumi.Input[str] keystore_password: Password for the keys.
        :param pulumi.Input[str] realm_id: The realm this keystore exists in.
        :param pulumi.Input[bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[str] algorithm: Intended algorithm for the key. Defaults to `RS256`
        :param pulumi.Input[bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[int] priority: Priority for the provider. Defaults to `0`
        """
        RealmKeystoreJavaGeneratedArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            key_alias=key_alias,
            key_password=key_password,
            keystore=keystore,
            keystore_password=keystore_password,
            realm_id=realm_id,
            active=active,
            algorithm=algorithm,
            enabled=enabled,
            name=name,
            priority=priority,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             key_alias: Optional[pulumi.Input[str]] = None,
             key_password: Optional[pulumi.Input[str]] = None,
             keystore: Optional[pulumi.Input[str]] = None,
             keystore_password: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             active: Optional[pulumi.Input[bool]] = None,
             algorithm: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             name: Optional[pulumi.Input[str]] = None,
             priority: Optional[pulumi.Input[int]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if key_alias is None and 'keyAlias' in kwargs:
            key_alias = kwargs['keyAlias']
        if key_alias is None:
            raise TypeError("Missing 'key_alias' argument")
        if key_password is None and 'keyPassword' in kwargs:
            key_password = kwargs['keyPassword']
        if key_password is None:
            raise TypeError("Missing 'key_password' argument")
        if keystore is None:
            raise TypeError("Missing 'keystore' argument")
        if keystore_password is None and 'keystorePassword' in kwargs:
            keystore_password = kwargs['keystorePassword']
        if keystore_password is None:
            raise TypeError("Missing 'keystore_password' argument")
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if realm_id is None:
            raise TypeError("Missing 'realm_id' argument")

        _setter("key_alias", key_alias)
        _setter("key_password", key_password)
        _setter("keystore", keystore)
        _setter("keystore_password", keystore_password)
        _setter("realm_id", realm_id)
        if active is not None:
            _setter("active", active)
        if algorithm is not None:
            _setter("algorithm", algorithm)
        if enabled is not None:
            _setter("enabled", enabled)
        if name is not None:
            _setter("name", name)
        if priority is not None:
            _setter("priority", priority)

    @property
    @pulumi.getter(name="keyAlias")
    def key_alias(self) -> pulumi.Input[str]:
        """
        Alias for the private key.
        """
        return pulumi.get(self, "key_alias")

    @key_alias.setter
    def key_alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_alias", value)

    @property
    @pulumi.getter(name="keyPassword")
    def key_password(self) -> pulumi.Input[str]:
        """
        Password for the private key.
        """
        return pulumi.get(self, "key_password")

    @key_password.setter
    def key_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_password", value)

    @property
    @pulumi.getter
    def keystore(self) -> pulumi.Input[str]:
        """
        Path to keys file on keycloak instance.
        """
        return pulumi.get(self, "keystore")

    @keystore.setter
    def keystore(self, value: pulumi.Input[str]):
        pulumi.set(self, "keystore", value)

    @property
    @pulumi.getter(name="keystorePassword")
    def keystore_password(self) -> pulumi.Input[str]:
        """
        Password for the keys.
        """
        return pulumi.get(self, "keystore_password")

    @keystore_password.setter
    def keystore_password(self, value: pulumi.Input[str]):
        pulumi.set(self, "keystore_password", value)

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
        Intended algorithm for the key. Defaults to `RS256`
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


@pulumi.input_type
class _RealmKeystoreJavaGeneratedState:
    def __init__(__self__, *,
                 active: Optional[pulumi.Input[bool]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_alias: Optional[pulumi.Input[str]] = None,
                 key_password: Optional[pulumi.Input[str]] = None,
                 keystore: Optional[pulumi.Input[str]] = None,
                 keystore_password: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RealmKeystoreJavaGenerated resources.
        :param pulumi.Input[bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[str] algorithm: Intended algorithm for the key. Defaults to `RS256`
        :param pulumi.Input[bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[str] key_alias: Alias for the private key.
        :param pulumi.Input[str] key_password: Password for the private key.
        :param pulumi.Input[str] keystore: Path to keys file on keycloak instance.
        :param pulumi.Input[str] keystore_password: Password for the keys.
        :param pulumi.Input[str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[str] realm_id: The realm this keystore exists in.
        """
        _RealmKeystoreJavaGeneratedState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            active=active,
            algorithm=algorithm,
            enabled=enabled,
            key_alias=key_alias,
            key_password=key_password,
            keystore=keystore,
            keystore_password=keystore_password,
            name=name,
            priority=priority,
            realm_id=realm_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             active: Optional[pulumi.Input[bool]] = None,
             algorithm: Optional[pulumi.Input[str]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             key_alias: Optional[pulumi.Input[str]] = None,
             key_password: Optional[pulumi.Input[str]] = None,
             keystore: Optional[pulumi.Input[str]] = None,
             keystore_password: Optional[pulumi.Input[str]] = None,
             name: Optional[pulumi.Input[str]] = None,
             priority: Optional[pulumi.Input[int]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if key_alias is None and 'keyAlias' in kwargs:
            key_alias = kwargs['keyAlias']
        if key_password is None and 'keyPassword' in kwargs:
            key_password = kwargs['keyPassword']
        if keystore_password is None and 'keystorePassword' in kwargs:
            keystore_password = kwargs['keystorePassword']
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']

        if active is not None:
            _setter("active", active)
        if algorithm is not None:
            _setter("algorithm", algorithm)
        if enabled is not None:
            _setter("enabled", enabled)
        if key_alias is not None:
            _setter("key_alias", key_alias)
        if key_password is not None:
            _setter("key_password", key_password)
        if keystore is not None:
            _setter("keystore", keystore)
        if keystore_password is not None:
            _setter("keystore_password", keystore_password)
        if name is not None:
            _setter("name", name)
        if priority is not None:
            _setter("priority", priority)
        if realm_id is not None:
            _setter("realm_id", realm_id)

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
        Intended algorithm for the key. Defaults to `RS256`
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
    @pulumi.getter(name="keyAlias")
    def key_alias(self) -> Optional[pulumi.Input[str]]:
        """
        Alias for the private key.
        """
        return pulumi.get(self, "key_alias")

    @key_alias.setter
    def key_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_alias", value)

    @property
    @pulumi.getter(name="keyPassword")
    def key_password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for the private key.
        """
        return pulumi.get(self, "key_password")

    @key_password.setter
    def key_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_password", value)

    @property
    @pulumi.getter
    def keystore(self) -> Optional[pulumi.Input[str]]:
        """
        Path to keys file on keycloak instance.
        """
        return pulumi.get(self, "keystore")

    @keystore.setter
    def keystore(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "keystore", value)

    @property
    @pulumi.getter(name="keystorePassword")
    def keystore_password(self) -> Optional[pulumi.Input[str]]:
        """
        Password for the keys.
        """
        return pulumi.get(self, "keystore_password")

    @keystore_password.setter
    def keystore_password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "keystore_password", value)

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


class RealmKeystoreJavaGenerated(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_alias: Optional[pulumi.Input[str]] = None,
                 key_password: Optional[pulumi.Input[str]] = None,
                 keystore: Optional[pulumi.Input[str]] = None,
                 keystore_password: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating and managing `java-keystore` Realm keystores within Keycloak.

        A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm", realm="my-realm")
        java_keystore = keycloak.RealmKeystoreJavaGenerated("javaKeystore",
            realm_id=realm.id,
            enabled=True,
            active=True,
            keystore="<path to your keystore>",
            keystore_password="<password for keystore>",
            key_alias="<alias for the private key>",
            key_password="<password for the private key>",
            priority=100,
            algorithm="RS256")
        ```

        ## Import

        Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash

        ```sh
         $ pulumi import keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated java_keystore my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[str] algorithm: Intended algorithm for the key. Defaults to `RS256`
        :param pulumi.Input[bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[str] key_alias: Alias for the private key.
        :param pulumi.Input[str] key_password: Password for the private key.
        :param pulumi.Input[str] keystore: Path to keys file on keycloak instance.
        :param pulumi.Input[str] keystore_password: Password for the keys.
        :param pulumi.Input[str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[str] realm_id: The realm this keystore exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RealmKeystoreJavaGeneratedArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing `java-keystore` Realm keystores within Keycloak.

        A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm", realm="my-realm")
        java_keystore = keycloak.RealmKeystoreJavaGenerated("javaKeystore",
            realm_id=realm.id,
            enabled=True,
            active=True,
            keystore="<path to your keystore>",
            keystore_password="<password for keystore>",
            key_alias="<alias for the private key>",
            key_password="<password for the private key>",
            priority=100,
            algorithm="RS256")
        ```

        ## Import

        Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash

        ```sh
         $ pulumi import keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated java_keystore my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
        ```

        :param str resource_name: The name of the resource.
        :param RealmKeystoreJavaGeneratedArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RealmKeystoreJavaGeneratedArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            RealmKeystoreJavaGeneratedArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 active: Optional[pulumi.Input[bool]] = None,
                 algorithm: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 key_alias: Optional[pulumi.Input[str]] = None,
                 key_password: Optional[pulumi.Input[str]] = None,
                 keystore: Optional[pulumi.Input[str]] = None,
                 keystore_password: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RealmKeystoreJavaGeneratedArgs.__new__(RealmKeystoreJavaGeneratedArgs)

            __props__.__dict__["active"] = active
            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["enabled"] = enabled
            if key_alias is None and not opts.urn:
                raise TypeError("Missing required property 'key_alias'")
            __props__.__dict__["key_alias"] = key_alias
            if key_password is None and not opts.urn:
                raise TypeError("Missing required property 'key_password'")
            __props__.__dict__["key_password"] = key_password
            if keystore is None and not opts.urn:
                raise TypeError("Missing required property 'keystore'")
            __props__.__dict__["keystore"] = keystore
            if keystore_password is None and not opts.urn:
                raise TypeError("Missing required property 'keystore_password'")
            __props__.__dict__["keystore_password"] = keystore_password
            __props__.__dict__["name"] = name
            __props__.__dict__["priority"] = priority
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(RealmKeystoreJavaGenerated, __self__).__init__(
            'keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated',
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
            key_alias: Optional[pulumi.Input[str]] = None,
            key_password: Optional[pulumi.Input[str]] = None,
            keystore: Optional[pulumi.Input[str]] = None,
            keystore_password: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            priority: Optional[pulumi.Input[int]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'RealmKeystoreJavaGenerated':
        """
        Get an existing RealmKeystoreJavaGenerated resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] active: When `false`, key in not used for signing. Defaults to `true`.
        :param pulumi.Input[str] algorithm: Intended algorithm for the key. Defaults to `RS256`
        :param pulumi.Input[bool] enabled: When `false`, key is not accessible in this realm. Defaults to `true`.
        :param pulumi.Input[str] key_alias: Alias for the private key.
        :param pulumi.Input[str] key_password: Password for the private key.
        :param pulumi.Input[str] keystore: Path to keys file on keycloak instance.
        :param pulumi.Input[str] keystore_password: Password for the keys.
        :param pulumi.Input[str] name: Display name of provider when linked in admin console.
        :param pulumi.Input[int] priority: Priority for the provider. Defaults to `0`
        :param pulumi.Input[str] realm_id: The realm this keystore exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RealmKeystoreJavaGeneratedState.__new__(_RealmKeystoreJavaGeneratedState)

        __props__.__dict__["active"] = active
        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["key_alias"] = key_alias
        __props__.__dict__["key_password"] = key_password
        __props__.__dict__["keystore"] = keystore
        __props__.__dict__["keystore_password"] = keystore_password
        __props__.__dict__["name"] = name
        __props__.__dict__["priority"] = priority
        __props__.__dict__["realm_id"] = realm_id
        return RealmKeystoreJavaGenerated(resource_name, opts=opts, __props__=__props__)

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
        Intended algorithm for the key. Defaults to `RS256`
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
    @pulumi.getter(name="keyAlias")
    def key_alias(self) -> pulumi.Output[str]:
        """
        Alias for the private key.
        """
        return pulumi.get(self, "key_alias")

    @property
    @pulumi.getter(name="keyPassword")
    def key_password(self) -> pulumi.Output[str]:
        """
        Password for the private key.
        """
        return pulumi.get(self, "key_password")

    @property
    @pulumi.getter
    def keystore(self) -> pulumi.Output[str]:
        """
        Path to keys file on keycloak instance.
        """
        return pulumi.get(self, "keystore")

    @property
    @pulumi.getter(name="keystorePassword")
    def keystore_password(self) -> pulumi.Output[str]:
        """
        Password for the keys.
        """
        return pulumi.get(self, "keystore_password")

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

