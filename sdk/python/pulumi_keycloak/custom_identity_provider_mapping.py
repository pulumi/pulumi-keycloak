# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['CustomIdentityProviderMappingArgs', 'CustomIdentityProviderMapping']

@pulumi.input_type
class CustomIdentityProviderMappingArgs:
    def __init__(__self__, *,
                 identity_provider_alias: pulumi.Input[str],
                 identity_provider_mapper: pulumi.Input[str],
                 realm: pulumi.Input[str],
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a CustomIdentityProviderMapping resource.
        :param pulumi.Input[str] identity_provider_alias: The alias of the associated identity provider.
        :param pulumi.Input[str] identity_provider_mapper: The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
        :param pulumi.Input[str] realm: The name of the realm.
        :param pulumi.Input[Mapping[str, Any]] extra_config: Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        :param pulumi.Input[str] name: The name of the mapper.
        """
        pulumi.set(__self__, "identity_provider_alias", identity_provider_alias)
        pulumi.set(__self__, "identity_provider_mapper", identity_provider_mapper)
        pulumi.set(__self__, "realm", realm)
        if extra_config is not None:
            pulumi.set(__self__, "extra_config", extra_config)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> pulumi.Input[str]:
        """
        The alias of the associated identity provider.
        """
        return pulumi.get(self, "identity_provider_alias")

    @identity_provider_alias.setter
    def identity_provider_alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_provider_alias", value)

    @property
    @pulumi.getter(name="identityProviderMapper")
    def identity_provider_mapper(self) -> pulumi.Input[str]:
        """
        The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
        """
        return pulumi.get(self, "identity_provider_mapper")

    @identity_provider_mapper.setter
    def identity_provider_mapper(self, value: pulumi.Input[str]):
        pulumi.set(self, "identity_provider_mapper", value)

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Input[str]:
        """
        The name of the realm.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        """
        return pulumi.get(self, "extra_config")

    @extra_config.setter
    def extra_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra_config", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the mapper.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _CustomIdentityProviderMappingState:
    def __init__(__self__, *,
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[str]] = None,
                 identity_provider_mapper: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering CustomIdentityProviderMapping resources.
        :param pulumi.Input[Mapping[str, Any]] extra_config: Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        :param pulumi.Input[str] identity_provider_alias: The alias of the associated identity provider.
        :param pulumi.Input[str] identity_provider_mapper: The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
        :param pulumi.Input[str] name: The name of the mapper.
        :param pulumi.Input[str] realm: The name of the realm.
        """
        if extra_config is not None:
            pulumi.set(__self__, "extra_config", extra_config)
        if identity_provider_alias is not None:
            pulumi.set(__self__, "identity_provider_alias", identity_provider_alias)
        if identity_provider_mapper is not None:
            pulumi.set(__self__, "identity_provider_mapper", identity_provider_mapper)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        """
        return pulumi.get(self, "extra_config")

    @extra_config.setter
    def extra_config(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "extra_config", value)

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> Optional[pulumi.Input[str]]:
        """
        The alias of the associated identity provider.
        """
        return pulumi.get(self, "identity_provider_alias")

    @identity_provider_alias.setter
    def identity_provider_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_provider_alias", value)

    @property
    @pulumi.getter(name="identityProviderMapper")
    def identity_provider_mapper(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
        """
        return pulumi.get(self, "identity_provider_mapper")

    @identity_provider_mapper.setter
    def identity_provider_mapper(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "identity_provider_mapper", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the mapper.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the realm.
        """
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)


class CustomIdentityProviderMapping(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[str]] = None,
                 identity_provider_mapper: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        oidc = keycloak.oidc.IdentityProvider("oidc",
            realm=realm.id,
            alias="oidc",
            authorization_url="https://example.com/auth",
            token_url="https://example.com/token",
            client_id="example_id",
            client_secret="example_token",
            default_scopes="openid random profile")
        oidc_custom_identity_provider_mapping = keycloak.CustomIdentityProviderMapping("oidc",
            realm=realm.id,
            name="email-attribute-importer",
            identity_provider_alias=oidc.alias,
            identity_provider_mapper="%s-user-attribute-idp-mapper",
            extra_config={
                "syncMode": "INHERIT",
                "Claim": "my-email-claim",
                "UserAttribute": "email",
            })
        ```

        ## Import

        Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak

        assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] extra_config: Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        :param pulumi.Input[str] identity_provider_alias: The alias of the associated identity provider.
        :param pulumi.Input[str] identity_provider_mapper: The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
        :param pulumi.Input[str] name: The name of the mapper.
        :param pulumi.Input[str] realm: The name of the realm.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomIdentityProviderMappingArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        oidc = keycloak.oidc.IdentityProvider("oidc",
            realm=realm.id,
            alias="oidc",
            authorization_url="https://example.com/auth",
            token_url="https://example.com/token",
            client_id="example_id",
            client_secret="example_token",
            default_scopes="openid random profile")
        oidc_custom_identity_provider_mapping = keycloak.CustomIdentityProviderMapping("oidc",
            realm=realm.id,
            name="email-attribute-importer",
            identity_provider_alias=oidc.alias,
            identity_provider_mapper="%s-user-attribute-idp-mapper",
            extra_config={
                "syncMode": "INHERIT",
                "Claim": "my-email-claim",
                "UserAttribute": "email",
            })
        ```

        ## Import

        Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak

        assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
        ```

        :param str resource_name: The name of the resource.
        :param CustomIdentityProviderMappingArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomIdentityProviderMappingArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 identity_provider_alias: Optional[pulumi.Input[str]] = None,
                 identity_provider_mapper: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomIdentityProviderMappingArgs.__new__(CustomIdentityProviderMappingArgs)

            __props__.__dict__["extra_config"] = extra_config
            if identity_provider_alias is None and not opts.urn:
                raise TypeError("Missing required property 'identity_provider_alias'")
            __props__.__dict__["identity_provider_alias"] = identity_provider_alias
            if identity_provider_mapper is None and not opts.urn:
                raise TypeError("Missing required property 'identity_provider_mapper'")
            __props__.__dict__["identity_provider_mapper"] = identity_provider_mapper
            __props__.__dict__["name"] = name
            if realm is None and not opts.urn:
                raise TypeError("Missing required property 'realm'")
            __props__.__dict__["realm"] = realm
        super(CustomIdentityProviderMapping, __self__).__init__(
            'keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            extra_config: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            identity_provider_alias: Optional[pulumi.Input[str]] = None,
            identity_provider_mapper: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm: Optional[pulumi.Input[str]] = None) -> 'CustomIdentityProviderMapping':
        """
        Get an existing CustomIdentityProviderMapping resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] extra_config: Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        :param pulumi.Input[str] identity_provider_alias: The alias of the associated identity provider.
        :param pulumi.Input[str] identity_provider_mapper: The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
        :param pulumi.Input[str] name: The name of the mapper.
        :param pulumi.Input[str] realm: The name of the realm.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CustomIdentityProviderMappingState.__new__(_CustomIdentityProviderMappingState)

        __props__.__dict__["extra_config"] = extra_config
        __props__.__dict__["identity_provider_alias"] = identity_provider_alias
        __props__.__dict__["identity_provider_mapper"] = identity_provider_mapper
        __props__.__dict__["name"] = name
        __props__.__dict__["realm"] = realm
        return CustomIdentityProviderMapping(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="extraConfig")
    def extra_config(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        """
        return pulumi.get(self, "extra_config")

    @property
    @pulumi.getter(name="identityProviderAlias")
    def identity_provider_alias(self) -> pulumi.Output[str]:
        """
        The alias of the associated identity provider.
        """
        return pulumi.get(self, "identity_provider_alias")

    @property
    @pulumi.getter(name="identityProviderMapper")
    def identity_provider_mapper(self) -> pulumi.Output[str]:
        """
        The type of the identity provider mapper. This can be a format string that includes a `%s` - this will be replaced by the provider id.
        """
        return pulumi.get(self, "identity_provider_mapper")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the mapper.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[str]:
        """
        The name of the realm.
        """
        return pulumi.get(self, "realm")

