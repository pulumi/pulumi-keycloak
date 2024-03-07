# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['FlowArgs', 'Flow']

@pulumi.input_type
class FlowArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Flow resource.
        :param pulumi.Input[str] alias: The alias for this authentication flow.
        :param pulumi.Input[str] realm_id: The realm that the authentication flow exists in.
        :param pulumi.Input[str] description: A description for the authentication flow.
        :param pulumi.Input[str] provider_id: The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        """
        pulumi.set(__self__, "alias", alias)
        pulumi.set(__self__, "realm_id", realm_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        """
        The alias for this authentication flow.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm that the authentication flow exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the authentication flow.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[str]]:
        """
        The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        """
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_id", value)


@pulumi.input_type
class _FlowState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Flow resources.
        :param pulumi.Input[str] alias: The alias for this authentication flow.
        :param pulumi.Input[str] description: A description for the authentication flow.
        :param pulumi.Input[str] provider_id: The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        :param pulumi.Input[str] realm_id: The realm that the authentication flow exists in.
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        The alias for this authentication flow.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the authentication flow.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[str]]:
        """
        The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        """
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm that the authentication flow exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class Flow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating and managing an authentication flow within Keycloak.

        [Authentication flows](https://www.keycloak.org/docs/11.0/server_admin/index.html#_authentication-flows) describe a sequence
        of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
        is a container for these actions, which are otherwise known as executions.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        flow = keycloak.authentication.Flow("flow",
            realm_id=realm.id,
            alias="my-flow-alias")
        execution = keycloak.authentication.Execution("execution",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="identity-provider-redirector",
            requirement="REQUIRED")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Authentication flows can be imported using the format `{{realmId}}/{{authenticationFlowId}}`. The authentication flow ID is

        typically a GUID which is autogenerated when the flow is created via Keycloak.

        Unfortunately, it is not trivial to retrieve the authentication flow ID from the UI. The best way to do this is to visit the

        "Authentication" page in Keycloak, and use the network tab of your browser to view the response of the API call to `/auth/admin/realms/${realm}/authentication/flows`,

        which will be a list of authentication flows.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:authentication/flow:Flow flow my-realm/e9a5641e-778c-4daf-89c0-f4ef617987d1
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The alias for this authentication flow.
        :param pulumi.Input[str] description: A description for the authentication flow.
        :param pulumi.Input[str] provider_id: The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        :param pulumi.Input[str] realm_id: The realm that the authentication flow exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FlowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing an authentication flow within Keycloak.

        [Authentication flows](https://www.keycloak.org/docs/11.0/server_admin/index.html#_authentication-flows) describe a sequence
        of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
        is a container for these actions, which are otherwise known as executions.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        flow = keycloak.authentication.Flow("flow",
            realm_id=realm.id,
            alias="my-flow-alias")
        execution = keycloak.authentication.Execution("execution",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="identity-provider-redirector",
            requirement="REQUIRED")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Authentication flows can be imported using the format `{{realmId}}/{{authenticationFlowId}}`. The authentication flow ID is

        typically a GUID which is autogenerated when the flow is created via Keycloak.

        Unfortunately, it is not trivial to retrieve the authentication flow ID from the UI. The best way to do this is to visit the

        "Authentication" page in Keycloak, and use the network tab of your browser to view the response of the API call to `/auth/admin/realms/${realm}/authentication/flows`,

        which will be a list of authentication flows.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:authentication/flow:Flow flow my-realm/e9a5641e-778c-4daf-89c0-f4ef617987d1
        ```

        :param str resource_name: The name of the resource.
        :param FlowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FlowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FlowArgs.__new__(FlowArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            __props__.__dict__["description"] = description
            __props__.__dict__["provider_id"] = provider_id
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(Flow, __self__).__init__(
            'keycloak:authentication/flow:Flow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            provider_id: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'Flow':
        """
        Get an existing Flow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The alias for this authentication flow.
        :param pulumi.Input[str] description: A description for the authentication flow.
        :param pulumi.Input[str] provider_id: The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        :param pulumi.Input[str] realm_id: The realm that the authentication flow exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FlowState.__new__(_FlowState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["description"] = description
        __props__.__dict__["provider_id"] = provider_id
        __props__.__dict__["realm_id"] = realm_id
        return Flow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        The alias for this authentication flow.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the authentication flow.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> pulumi.Output[Optional[str]]:
        """
        The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        """
        return pulumi.get(self, "provider_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm that the authentication flow exists in.
        """
        return pulumi.get(self, "realm_id")

