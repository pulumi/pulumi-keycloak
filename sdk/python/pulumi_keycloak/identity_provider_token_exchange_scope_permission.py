# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['IdentityProviderTokenExchangeScopePermissionArgs', 'IdentityProviderTokenExchangeScopePermission']

@pulumi.input_type
class IdentityProviderTokenExchangeScopePermissionArgs:
    def __init__(__self__, *,
                 clients: pulumi.Input[Sequence[pulumi.Input[str]]],
                 provider_alias: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 policy_type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IdentityProviderTokenExchangeScopePermission resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] clients: A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        :param pulumi.Input[str] provider_alias: Alias of the identity provider.
        :param pulumi.Input[str] realm_id: The realm that the identity provider exists in.
        :param pulumi.Input[str] policy_type: Defaults to "client" This is also the only value policy type supported by this provider.
        """
        pulumi.set(__self__, "clients", clients)
        pulumi.set(__self__, "provider_alias", provider_alias)
        pulumi.set(__self__, "realm_id", realm_id)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)

    @property
    @pulumi.getter
    def clients(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        """
        return pulumi.get(self, "clients")

    @clients.setter
    def clients(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "clients", value)

    @property
    @pulumi.getter(name="providerAlias")
    def provider_alias(self) -> pulumi.Input[str]:
        """
        Alias of the identity provider.
        """
        return pulumi.get(self, "provider_alias")

    @provider_alias.setter
    def provider_alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "provider_alias", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm that the identity provider exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to "client" This is also the only value policy type supported by this provider.
        """
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_type", value)


@pulumi.input_type
class _IdentityProviderTokenExchangeScopePermissionState:
    def __init__(__self__, *,
                 authorization_idp_resource_id: Optional[pulumi.Input[str]] = None,
                 authorization_resource_server_id: Optional[pulumi.Input[str]] = None,
                 authorization_token_exchange_scope_permission_id: Optional[pulumi.Input[str]] = None,
                 clients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy_id: Optional[pulumi.Input[str]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 provider_alias: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IdentityProviderTokenExchangeScopePermission resources.
        :param pulumi.Input[str] authorization_idp_resource_id: (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
        :param pulumi.Input[str] authorization_resource_server_id: (Computed) Resource server ID representing the realm management client on which this permission is managed.
        :param pulumi.Input[str] authorization_token_exchange_scope_permission_id: (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] clients: A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        :param pulumi.Input[str] policy_id: (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
        :param pulumi.Input[str] policy_type: Defaults to "client" This is also the only value policy type supported by this provider.
        :param pulumi.Input[str] provider_alias: Alias of the identity provider.
        :param pulumi.Input[str] realm_id: The realm that the identity provider exists in.
        """
        if authorization_idp_resource_id is not None:
            pulumi.set(__self__, "authorization_idp_resource_id", authorization_idp_resource_id)
        if authorization_resource_server_id is not None:
            pulumi.set(__self__, "authorization_resource_server_id", authorization_resource_server_id)
        if authorization_token_exchange_scope_permission_id is not None:
            pulumi.set(__self__, "authorization_token_exchange_scope_permission_id", authorization_token_exchange_scope_permission_id)
        if clients is not None:
            pulumi.set(__self__, "clients", clients)
        if policy_id is not None:
            pulumi.set(__self__, "policy_id", policy_id)
        if policy_type is not None:
            pulumi.set(__self__, "policy_type", policy_type)
        if provider_alias is not None:
            pulumi.set(__self__, "provider_alias", provider_alias)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="authorizationIdpResourceId")
    def authorization_idp_resource_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
        """
        return pulumi.get(self, "authorization_idp_resource_id")

    @authorization_idp_resource_id.setter
    def authorization_idp_resource_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_idp_resource_id", value)

    @property
    @pulumi.getter(name="authorizationResourceServerId")
    def authorization_resource_server_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) Resource server ID representing the realm management client on which this permission is managed.
        """
        return pulumi.get(self, "authorization_resource_server_id")

    @authorization_resource_server_id.setter
    def authorization_resource_server_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_resource_server_id", value)

    @property
    @pulumi.getter(name="authorizationTokenExchangeScopePermissionId")
    def authorization_token_exchange_scope_permission_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
        """
        return pulumi.get(self, "authorization_token_exchange_scope_permission_id")

    @authorization_token_exchange_scope_permission_id.setter
    def authorization_token_exchange_scope_permission_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authorization_token_exchange_scope_permission_id", value)

    @property
    @pulumi.getter
    def clients(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        """
        return pulumi.get(self, "clients")

    @clients.setter
    def clients(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "clients", value)

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> Optional[pulumi.Input[str]]:
        """
        (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
        """
        return pulumi.get(self, "policy_id")

    @policy_id.setter
    def policy_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_id", value)

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> Optional[pulumi.Input[str]]:
        """
        Defaults to "client" This is also the only value policy type supported by this provider.
        """
        return pulumi.get(self, "policy_type")

    @policy_type.setter
    def policy_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy_type", value)

    @property
    @pulumi.getter(name="providerAlias")
    def provider_alias(self) -> Optional[pulumi.Input[str]]:
        """
        Alias of the identity provider.
        """
        return pulumi.get(self, "provider_alias")

    @provider_alias.setter
    def provider_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_alias", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm that the identity provider exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class IdentityProviderTokenExchangeScopePermission(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 provider_alias: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        This resource can be imported using the format `{{realm_id}}/{{provider_alias}}`, where `provider_alias` is the alias that you assign to the identity provider upon creation. Examplebash

        ```sh
         $ pulumi import keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission oidc_idp_permission my-realm/myIdp
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] clients: A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        :param pulumi.Input[str] policy_type: Defaults to "client" This is also the only value policy type supported by this provider.
        :param pulumi.Input[str] provider_alias: Alias of the identity provider.
        :param pulumi.Input[str] realm_id: The realm that the identity provider exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IdentityProviderTokenExchangeScopePermissionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        This resource can be imported using the format `{{realm_id}}/{{provider_alias}}`, where `provider_alias` is the alias that you assign to the identity provider upon creation. Examplebash

        ```sh
         $ pulumi import keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission oidc_idp_permission my-realm/myIdp
        ```

        :param str resource_name: The name of the resource.
        :param IdentityProviderTokenExchangeScopePermissionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IdentityProviderTokenExchangeScopePermissionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 clients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy_type: Optional[pulumi.Input[str]] = None,
                 provider_alias: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IdentityProviderTokenExchangeScopePermissionArgs.__new__(IdentityProviderTokenExchangeScopePermissionArgs)

            if clients is None and not opts.urn:
                raise TypeError("Missing required property 'clients'")
            __props__.__dict__["clients"] = clients
            __props__.__dict__["policy_type"] = policy_type
            if provider_alias is None and not opts.urn:
                raise TypeError("Missing required property 'provider_alias'")
            __props__.__dict__["provider_alias"] = provider_alias
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            __props__.__dict__["authorization_idp_resource_id"] = None
            __props__.__dict__["authorization_resource_server_id"] = None
            __props__.__dict__["authorization_token_exchange_scope_permission_id"] = None
            __props__.__dict__["policy_id"] = None
        super(IdentityProviderTokenExchangeScopePermission, __self__).__init__(
            'keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authorization_idp_resource_id: Optional[pulumi.Input[str]] = None,
            authorization_resource_server_id: Optional[pulumi.Input[str]] = None,
            authorization_token_exchange_scope_permission_id: Optional[pulumi.Input[str]] = None,
            clients: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            policy_id: Optional[pulumi.Input[str]] = None,
            policy_type: Optional[pulumi.Input[str]] = None,
            provider_alias: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'IdentityProviderTokenExchangeScopePermission':
        """
        Get an existing IdentityProviderTokenExchangeScopePermission resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authorization_idp_resource_id: (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
        :param pulumi.Input[str] authorization_resource_server_id: (Computed) Resource server ID representing the realm management client on which this permission is managed.
        :param pulumi.Input[str] authorization_token_exchange_scope_permission_id: (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] clients: A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        :param pulumi.Input[str] policy_id: (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
        :param pulumi.Input[str] policy_type: Defaults to "client" This is also the only value policy type supported by this provider.
        :param pulumi.Input[str] provider_alias: Alias of the identity provider.
        :param pulumi.Input[str] realm_id: The realm that the identity provider exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IdentityProviderTokenExchangeScopePermissionState.__new__(_IdentityProviderTokenExchangeScopePermissionState)

        __props__.__dict__["authorization_idp_resource_id"] = authorization_idp_resource_id
        __props__.__dict__["authorization_resource_server_id"] = authorization_resource_server_id
        __props__.__dict__["authorization_token_exchange_scope_permission_id"] = authorization_token_exchange_scope_permission_id
        __props__.__dict__["clients"] = clients
        __props__.__dict__["policy_id"] = policy_id
        __props__.__dict__["policy_type"] = policy_type
        __props__.__dict__["provider_alias"] = provider_alias
        __props__.__dict__["realm_id"] = realm_id
        return IdentityProviderTokenExchangeScopePermission(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="authorizationIdpResourceId")
    def authorization_idp_resource_id(self) -> pulumi.Output[str]:
        """
        (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
        """
        return pulumi.get(self, "authorization_idp_resource_id")

    @property
    @pulumi.getter(name="authorizationResourceServerId")
    def authorization_resource_server_id(self) -> pulumi.Output[str]:
        """
        (Computed) Resource server ID representing the realm management client on which this permission is managed.
        """
        return pulumi.get(self, "authorization_resource_server_id")

    @property
    @pulumi.getter(name="authorizationTokenExchangeScopePermissionId")
    def authorization_token_exchange_scope_permission_id(self) -> pulumi.Output[str]:
        """
        (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
        """
        return pulumi.get(self, "authorization_token_exchange_scope_permission_id")

    @property
    @pulumi.getter
    def clients(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        """
        return pulumi.get(self, "clients")

    @property
    @pulumi.getter(name="policyId")
    def policy_id(self) -> pulumi.Output[str]:
        """
        (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
        """
        return pulumi.get(self, "policy_id")

    @property
    @pulumi.getter(name="policyType")
    def policy_type(self) -> pulumi.Output[Optional[str]]:
        """
        Defaults to "client" This is also the only value policy type supported by this provider.
        """
        return pulumi.get(self, "policy_type")

    @property
    @pulumi.getter(name="providerAlias")
    def provider_alias(self) -> pulumi.Output[str]:
        """
        Alias of the identity provider.
        """
        return pulumi.get(self, "provider_alias")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm that the identity provider exists in.
        """
        return pulumi.get(self, "realm_id")

