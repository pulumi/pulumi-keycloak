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
from .. import _utilities

__all__ = ['BindingsArgs', 'Bindings']

@pulumi.input_type
class BindingsArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[builtins.str],
                 browser_flow: Optional[pulumi.Input[builtins.str]] = None,
                 client_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
                 direct_grant_flow: Optional[pulumi.Input[builtins.str]] = None,
                 docker_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
                 first_broker_login_flow: Optional[pulumi.Input[builtins.str]] = None,
                 registration_flow: Optional[pulumi.Input[builtins.str]] = None,
                 reset_credentials_flow: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a Bindings resource.
        :param pulumi.Input[builtins.str] realm_id: The realm the authentication flow binding exists in.
        :param pulumi.Input[builtins.str] browser_flow: The alias of the flow to assign to the realm BrowserFlow.
        :param pulumi.Input[builtins.str] client_authentication_flow: The alias of the flow to assign to the realm ClientAuthenticationFlow.
        :param pulumi.Input[builtins.str] direct_grant_flow: The alias of the flow to assign to the realm DirectGrantFlow.
        :param pulumi.Input[builtins.str] docker_authentication_flow: The alias of the flow to assign to the realm DockerAuthenticationFlow.
        :param pulumi.Input[builtins.str] first_broker_login_flow: The alias of the flow to assign to the realm FirstBrokerLoginFlow (since Keycloak 24).
        :param pulumi.Input[builtins.str] registration_flow: The alias of the flow to assign to the realm RegistrationFlow.
        :param pulumi.Input[builtins.str] reset_credentials_flow: The alias of the flow to assign to the realm ResetCredentialsFlow.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        if browser_flow is not None:
            pulumi.set(__self__, "browser_flow", browser_flow)
        if client_authentication_flow is not None:
            pulumi.set(__self__, "client_authentication_flow", client_authentication_flow)
        if direct_grant_flow is not None:
            pulumi.set(__self__, "direct_grant_flow", direct_grant_flow)
        if docker_authentication_flow is not None:
            pulumi.set(__self__, "docker_authentication_flow", docker_authentication_flow)
        if first_broker_login_flow is not None:
            pulumi.set(__self__, "first_broker_login_flow", first_broker_login_flow)
        if registration_flow is not None:
            pulumi.set(__self__, "registration_flow", registration_flow)
        if reset_credentials_flow is not None:
            pulumi.set(__self__, "reset_credentials_flow", reset_credentials_flow)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[builtins.str]:
        """
        The realm the authentication flow binding exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="browserFlow")
    def browser_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm BrowserFlow.
        """
        return pulumi.get(self, "browser_flow")

    @browser_flow.setter
    def browser_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "browser_flow", value)

    @property
    @pulumi.getter(name="clientAuthenticationFlow")
    def client_authentication_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm ClientAuthenticationFlow.
        """
        return pulumi.get(self, "client_authentication_flow")

    @client_authentication_flow.setter
    def client_authentication_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "client_authentication_flow", value)

    @property
    @pulumi.getter(name="directGrantFlow")
    def direct_grant_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm DirectGrantFlow.
        """
        return pulumi.get(self, "direct_grant_flow")

    @direct_grant_flow.setter
    def direct_grant_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "direct_grant_flow", value)

    @property
    @pulumi.getter(name="dockerAuthenticationFlow")
    def docker_authentication_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm DockerAuthenticationFlow.
        """
        return pulumi.get(self, "docker_authentication_flow")

    @docker_authentication_flow.setter
    def docker_authentication_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "docker_authentication_flow", value)

    @property
    @pulumi.getter(name="firstBrokerLoginFlow")
    def first_broker_login_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm FirstBrokerLoginFlow (since Keycloak 24).
        """
        return pulumi.get(self, "first_broker_login_flow")

    @first_broker_login_flow.setter
    def first_broker_login_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "first_broker_login_flow", value)

    @property
    @pulumi.getter(name="registrationFlow")
    def registration_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm RegistrationFlow.
        """
        return pulumi.get(self, "registration_flow")

    @registration_flow.setter
    def registration_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "registration_flow", value)

    @property
    @pulumi.getter(name="resetCredentialsFlow")
    def reset_credentials_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm ResetCredentialsFlow.
        """
        return pulumi.get(self, "reset_credentials_flow")

    @reset_credentials_flow.setter
    def reset_credentials_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "reset_credentials_flow", value)


@pulumi.input_type
class _BindingsState:
    def __init__(__self__, *,
                 browser_flow: Optional[pulumi.Input[builtins.str]] = None,
                 client_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
                 direct_grant_flow: Optional[pulumi.Input[builtins.str]] = None,
                 docker_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
                 first_broker_login_flow: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 registration_flow: Optional[pulumi.Input[builtins.str]] = None,
                 reset_credentials_flow: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering Bindings resources.
        :param pulumi.Input[builtins.str] browser_flow: The alias of the flow to assign to the realm BrowserFlow.
        :param pulumi.Input[builtins.str] client_authentication_flow: The alias of the flow to assign to the realm ClientAuthenticationFlow.
        :param pulumi.Input[builtins.str] direct_grant_flow: The alias of the flow to assign to the realm DirectGrantFlow.
        :param pulumi.Input[builtins.str] docker_authentication_flow: The alias of the flow to assign to the realm DockerAuthenticationFlow.
        :param pulumi.Input[builtins.str] first_broker_login_flow: The alias of the flow to assign to the realm FirstBrokerLoginFlow (since Keycloak 24).
        :param pulumi.Input[builtins.str] realm_id: The realm the authentication flow binding exists in.
        :param pulumi.Input[builtins.str] registration_flow: The alias of the flow to assign to the realm RegistrationFlow.
        :param pulumi.Input[builtins.str] reset_credentials_flow: The alias of the flow to assign to the realm ResetCredentialsFlow.
        """
        if browser_flow is not None:
            pulumi.set(__self__, "browser_flow", browser_flow)
        if client_authentication_flow is not None:
            pulumi.set(__self__, "client_authentication_flow", client_authentication_flow)
        if direct_grant_flow is not None:
            pulumi.set(__self__, "direct_grant_flow", direct_grant_flow)
        if docker_authentication_flow is not None:
            pulumi.set(__self__, "docker_authentication_flow", docker_authentication_flow)
        if first_broker_login_flow is not None:
            pulumi.set(__self__, "first_broker_login_flow", first_broker_login_flow)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if registration_flow is not None:
            pulumi.set(__self__, "registration_flow", registration_flow)
        if reset_credentials_flow is not None:
            pulumi.set(__self__, "reset_credentials_flow", reset_credentials_flow)

    @property
    @pulumi.getter(name="browserFlow")
    def browser_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm BrowserFlow.
        """
        return pulumi.get(self, "browser_flow")

    @browser_flow.setter
    def browser_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "browser_flow", value)

    @property
    @pulumi.getter(name="clientAuthenticationFlow")
    def client_authentication_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm ClientAuthenticationFlow.
        """
        return pulumi.get(self, "client_authentication_flow")

    @client_authentication_flow.setter
    def client_authentication_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "client_authentication_flow", value)

    @property
    @pulumi.getter(name="directGrantFlow")
    def direct_grant_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm DirectGrantFlow.
        """
        return pulumi.get(self, "direct_grant_flow")

    @direct_grant_flow.setter
    def direct_grant_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "direct_grant_flow", value)

    @property
    @pulumi.getter(name="dockerAuthenticationFlow")
    def docker_authentication_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm DockerAuthenticationFlow.
        """
        return pulumi.get(self, "docker_authentication_flow")

    @docker_authentication_flow.setter
    def docker_authentication_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "docker_authentication_flow", value)

    @property
    @pulumi.getter(name="firstBrokerLoginFlow")
    def first_broker_login_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm FirstBrokerLoginFlow (since Keycloak 24).
        """
        return pulumi.get(self, "first_broker_login_flow")

    @first_broker_login_flow.setter
    def first_broker_login_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "first_broker_login_flow", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The realm the authentication flow binding exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="registrationFlow")
    def registration_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm RegistrationFlow.
        """
        return pulumi.get(self, "registration_flow")

    @registration_flow.setter
    def registration_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "registration_flow", value)

    @property
    @pulumi.getter(name="resetCredentialsFlow")
    def reset_credentials_flow(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The alias of the flow to assign to the realm ResetCredentialsFlow.
        """
        return pulumi.get(self, "reset_credentials_flow")

    @reset_credentials_flow.setter
    def reset_credentials_flow(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "reset_credentials_flow", value)


@pulumi.type_token("keycloak:authentication/bindings:Bindings")
class Bindings(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 browser_flow: Optional[pulumi.Input[builtins.str]] = None,
                 client_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
                 direct_grant_flow: Optional[pulumi.Input[builtins.str]] = None,
                 docker_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
                 first_broker_login_flow: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 registration_flow: Optional[pulumi.Input[builtins.str]] = None,
                 reset_credentials_flow: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Allows for creating and managing realm authentication flow bindings within Keycloak.

        [Authentication flows](https://www.keycloak.org/docs/latest/server_admin/index.html#_authentication-flows) describe a sequence
        of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
        is a container for these actions, which are otherwise known as executions.

        Realms assign authentication flows to supported user flows such as `registration` and `browser`. This resource allows the
        updating of realm authentication flow bindings to custom authentication flows created by `authentication.Flow`.

        Note that you can also use the `Realm` resource to assign authentication flow bindings at the realm level. This
        resource is useful if you would like to create a realm and an authentication flow, and assign this flow to the realm within
        a single run of `pulumi up`. In any case, do not attempt to use both the arguments within the `Realm` resource
        and this resource to manage authentication flow bindings, you should choose one or the other.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        flow = keycloak.authentication.Flow("flow",
            realm_id=realm.id,
            alias="my-flow-alias")
        # first execution
        execution_one = keycloak.authentication.Execution("execution_one",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="auth-cookie",
            requirement="ALTERNATIVE")
        # second execution
        execution_two = keycloak.authentication.Execution("execution_two",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="identity-provider-redirector",
            requirement="ALTERNATIVE",
            opts = pulumi.ResourceOptions(depends_on=[execution_one]))
        browser_authentication_binding = keycloak.authentication.Bindings("browser_authentication_binding",
            realm_id=realm.id,
            browser_flow=flow.alias)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] browser_flow: The alias of the flow to assign to the realm BrowserFlow.
        :param pulumi.Input[builtins.str] client_authentication_flow: The alias of the flow to assign to the realm ClientAuthenticationFlow.
        :param pulumi.Input[builtins.str] direct_grant_flow: The alias of the flow to assign to the realm DirectGrantFlow.
        :param pulumi.Input[builtins.str] docker_authentication_flow: The alias of the flow to assign to the realm DockerAuthenticationFlow.
        :param pulumi.Input[builtins.str] first_broker_login_flow: The alias of the flow to assign to the realm FirstBrokerLoginFlow (since Keycloak 24).
        :param pulumi.Input[builtins.str] realm_id: The realm the authentication flow binding exists in.
        :param pulumi.Input[builtins.str] registration_flow: The alias of the flow to assign to the realm RegistrationFlow.
        :param pulumi.Input[builtins.str] reset_credentials_flow: The alias of the flow to assign to the realm ResetCredentialsFlow.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BindingsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing realm authentication flow bindings within Keycloak.

        [Authentication flows](https://www.keycloak.org/docs/latest/server_admin/index.html#_authentication-flows) describe a sequence
        of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
        is a container for these actions, which are otherwise known as executions.

        Realms assign authentication flows to supported user flows such as `registration` and `browser`. This resource allows the
        updating of realm authentication flow bindings to custom authentication flows created by `authentication.Flow`.

        Note that you can also use the `Realm` resource to assign authentication flow bindings at the realm level. This
        resource is useful if you would like to create a realm and an authentication flow, and assign this flow to the realm within
        a single run of `pulumi up`. In any case, do not attempt to use both the arguments within the `Realm` resource
        and this resource to manage authentication flow bindings, you should choose one or the other.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        flow = keycloak.authentication.Flow("flow",
            realm_id=realm.id,
            alias="my-flow-alias")
        # first execution
        execution_one = keycloak.authentication.Execution("execution_one",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="auth-cookie",
            requirement="ALTERNATIVE")
        # second execution
        execution_two = keycloak.authentication.Execution("execution_two",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="identity-provider-redirector",
            requirement="ALTERNATIVE",
            opts = pulumi.ResourceOptions(depends_on=[execution_one]))
        browser_authentication_binding = keycloak.authentication.Bindings("browser_authentication_binding",
            realm_id=realm.id,
            browser_flow=flow.alias)
        ```

        :param str resource_name: The name of the resource.
        :param BindingsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BindingsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 browser_flow: Optional[pulumi.Input[builtins.str]] = None,
                 client_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
                 direct_grant_flow: Optional[pulumi.Input[builtins.str]] = None,
                 docker_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
                 first_broker_login_flow: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 registration_flow: Optional[pulumi.Input[builtins.str]] = None,
                 reset_credentials_flow: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BindingsArgs.__new__(BindingsArgs)

            __props__.__dict__["browser_flow"] = browser_flow
            __props__.__dict__["client_authentication_flow"] = client_authentication_flow
            __props__.__dict__["direct_grant_flow"] = direct_grant_flow
            __props__.__dict__["docker_authentication_flow"] = docker_authentication_flow
            __props__.__dict__["first_broker_login_flow"] = first_broker_login_flow
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            __props__.__dict__["registration_flow"] = registration_flow
            __props__.__dict__["reset_credentials_flow"] = reset_credentials_flow
        super(Bindings, __self__).__init__(
            'keycloak:authentication/bindings:Bindings',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            browser_flow: Optional[pulumi.Input[builtins.str]] = None,
            client_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
            direct_grant_flow: Optional[pulumi.Input[builtins.str]] = None,
            docker_authentication_flow: Optional[pulumi.Input[builtins.str]] = None,
            first_broker_login_flow: Optional[pulumi.Input[builtins.str]] = None,
            realm_id: Optional[pulumi.Input[builtins.str]] = None,
            registration_flow: Optional[pulumi.Input[builtins.str]] = None,
            reset_credentials_flow: Optional[pulumi.Input[builtins.str]] = None) -> 'Bindings':
        """
        Get an existing Bindings resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] browser_flow: The alias of the flow to assign to the realm BrowserFlow.
        :param pulumi.Input[builtins.str] client_authentication_flow: The alias of the flow to assign to the realm ClientAuthenticationFlow.
        :param pulumi.Input[builtins.str] direct_grant_flow: The alias of the flow to assign to the realm DirectGrantFlow.
        :param pulumi.Input[builtins.str] docker_authentication_flow: The alias of the flow to assign to the realm DockerAuthenticationFlow.
        :param pulumi.Input[builtins.str] first_broker_login_flow: The alias of the flow to assign to the realm FirstBrokerLoginFlow (since Keycloak 24).
        :param pulumi.Input[builtins.str] realm_id: The realm the authentication flow binding exists in.
        :param pulumi.Input[builtins.str] registration_flow: The alias of the flow to assign to the realm RegistrationFlow.
        :param pulumi.Input[builtins.str] reset_credentials_flow: The alias of the flow to assign to the realm ResetCredentialsFlow.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BindingsState.__new__(_BindingsState)

        __props__.__dict__["browser_flow"] = browser_flow
        __props__.__dict__["client_authentication_flow"] = client_authentication_flow
        __props__.__dict__["direct_grant_flow"] = direct_grant_flow
        __props__.__dict__["docker_authentication_flow"] = docker_authentication_flow
        __props__.__dict__["first_broker_login_flow"] = first_broker_login_flow
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["registration_flow"] = registration_flow
        __props__.__dict__["reset_credentials_flow"] = reset_credentials_flow
        return Bindings(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="browserFlow")
    def browser_flow(self) -> pulumi.Output[builtins.str]:
        """
        The alias of the flow to assign to the realm BrowserFlow.
        """
        return pulumi.get(self, "browser_flow")

    @property
    @pulumi.getter(name="clientAuthenticationFlow")
    def client_authentication_flow(self) -> pulumi.Output[builtins.str]:
        """
        The alias of the flow to assign to the realm ClientAuthenticationFlow.
        """
        return pulumi.get(self, "client_authentication_flow")

    @property
    @pulumi.getter(name="directGrantFlow")
    def direct_grant_flow(self) -> pulumi.Output[builtins.str]:
        """
        The alias of the flow to assign to the realm DirectGrantFlow.
        """
        return pulumi.get(self, "direct_grant_flow")

    @property
    @pulumi.getter(name="dockerAuthenticationFlow")
    def docker_authentication_flow(self) -> pulumi.Output[builtins.str]:
        """
        The alias of the flow to assign to the realm DockerAuthenticationFlow.
        """
        return pulumi.get(self, "docker_authentication_flow")

    @property
    @pulumi.getter(name="firstBrokerLoginFlow")
    def first_broker_login_flow(self) -> pulumi.Output[builtins.str]:
        """
        The alias of the flow to assign to the realm FirstBrokerLoginFlow (since Keycloak 24).
        """
        return pulumi.get(self, "first_broker_login_flow")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[builtins.str]:
        """
        The realm the authentication flow binding exists in.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="registrationFlow")
    def registration_flow(self) -> pulumi.Output[builtins.str]:
        """
        The alias of the flow to assign to the realm RegistrationFlow.
        """
        return pulumi.get(self, "registration_flow")

    @property
    @pulumi.getter(name="resetCredentialsFlow")
    def reset_credentials_flow(self) -> pulumi.Output[builtins.str]:
        """
        The alias of the flow to assign to the realm ResetCredentialsFlow.
        """
        return pulumi.get(self, "reset_credentials_flow")

