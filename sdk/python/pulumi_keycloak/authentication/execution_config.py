# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['ExecutionConfigArgs', 'ExecutionConfig']

@pulumi.input_type
class ExecutionConfigArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 config: pulumi.Input[Mapping[str, pulumi.Input[str]]],
                 execution_id: pulumi.Input[str],
                 realm_id: pulumi.Input[str]):
        """
        The set of arguments for constructing a ExecutionConfig resource.
        :param pulumi.Input[str] alias: The name of the configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] config: The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
        :param pulumi.Input[str] execution_id: The authentication execution this configuration is attached to.
        :param pulumi.Input[str] realm_id: The realm the authentication execution exists in.
        """
        ExecutionConfigArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            alias=alias,
            config=config,
            execution_id=execution_id,
            realm_id=realm_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             alias: Optional[pulumi.Input[str]] = None,
             config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             execution_id: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if alias is None:
            raise TypeError("Missing 'alias' argument")
        if config is None:
            raise TypeError("Missing 'config' argument")
        if execution_id is None and 'executionId' in kwargs:
            execution_id = kwargs['executionId']
        if execution_id is None:
            raise TypeError("Missing 'execution_id' argument")
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if realm_id is None:
            raise TypeError("Missing 'realm_id' argument")

        _setter("alias", alias)
        _setter("config", config)
        _setter("execution_id", execution_id)
        _setter("realm_id", realm_id)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        """
        The name of the configuration.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input[Mapping[str, pulumi.Input[str]]]:
        """
        The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input[Mapping[str, pulumi.Input[str]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> pulumi.Input[str]:
        """
        The authentication execution this configuration is attached to.
        """
        return pulumi.get(self, "execution_id")

    @execution_id.setter
    def execution_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "execution_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm the authentication execution exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)


@pulumi.input_type
class _ExecutionConfigState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 execution_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ExecutionConfig resources.
        :param pulumi.Input[str] alias: The name of the configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] config: The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
        :param pulumi.Input[str] execution_id: The authentication execution this configuration is attached to.
        :param pulumi.Input[str] realm_id: The realm the authentication execution exists in.
        """
        _ExecutionConfigState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            alias=alias,
            config=config,
            execution_id=execution_id,
            realm_id=realm_id,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             alias: Optional[pulumi.Input[str]] = None,
             config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
             execution_id: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions]=None,
             **kwargs):
        if execution_id is None and 'executionId' in kwargs:
            execution_id = kwargs['executionId']
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']

        if alias is not None:
            _setter("alias", alias)
        if config is not None:
            _setter("config", config)
        if execution_id is not None:
            _setter("execution_id", execution_id)
        if realm_id is not None:
            _setter("realm_id", realm_id)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the configuration.
        """
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> Optional[pulumi.Input[str]]:
        """
        The authentication execution this configuration is attached to.
        """
        return pulumi.get(self, "execution_id")

    @execution_id.setter
    def execution_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "execution_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm the authentication execution exists in.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class ExecutionConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 execution_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for managing an authentication execution's configuration. If a particular authentication execution supports additional
        configuration (such as with the `identity-provider-redirector` execution), this can be managed with this resource.

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
        execution = keycloak.authentication.Execution("execution",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="identity-provider-redirector")
        config = keycloak.authentication.ExecutionConfig("config",
            realm_id=realm.id,
            execution_id=execution.id,
            alias="my-config-alias",
            config={
                "defaultProvider": "my-config-default-idp",
            })
        ```

        ## Import

        Configurations can be imported using the format `{{realm}}/{{authenticationExecutionId}}/{{authenticationExecutionConfigId}}`. If the `authenticationExecutionId` is incorrect, the import will still be successful. A subsequent apply will change the `authenticationExecutionId` to the correct one, which causes the configuration to be replaced. Examplebash

        ```sh
         $ pulumi import keycloak:authentication/executionConfig:ExecutionConfig config my-realm/be081463-ddbf-4b42-9eff-9c97886f24ff/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The name of the configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] config: The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
        :param pulumi.Input[str] execution_id: The authentication execution this configuration is attached to.
        :param pulumi.Input[str] realm_id: The realm the authentication execution exists in.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExecutionConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for managing an authentication execution's configuration. If a particular authentication execution supports additional
        configuration (such as with the `identity-provider-redirector` execution), this can be managed with this resource.

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
        execution = keycloak.authentication.Execution("execution",
            realm_id=realm.id,
            parent_flow_alias=flow.alias,
            authenticator="identity-provider-redirector")
        config = keycloak.authentication.ExecutionConfig("config",
            realm_id=realm.id,
            execution_id=execution.id,
            alias="my-config-alias",
            config={
                "defaultProvider": "my-config-default-idp",
            })
        ```

        ## Import

        Configurations can be imported using the format `{{realm}}/{{authenticationExecutionId}}/{{authenticationExecutionConfigId}}`. If the `authenticationExecutionId` is incorrect, the import will still be successful. A subsequent apply will change the `authenticationExecutionId` to the correct one, which causes the configuration to be replaced. Examplebash

        ```sh
         $ pulumi import keycloak:authentication/executionConfig:ExecutionConfig config my-realm/be081463-ddbf-4b42-9eff-9c97886f24ff/30559fcf-6fb8-45ea-8c46-2b86f46ebc17
        ```

        :param str resource_name: The name of the resource.
        :param ExecutionConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExecutionConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            ExecutionConfigArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 execution_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExecutionConfigArgs.__new__(ExecutionConfigArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            if execution_id is None and not opts.urn:
                raise TypeError("Missing required property 'execution_id'")
            __props__.__dict__["execution_id"] = execution_id
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(ExecutionConfig, __self__).__init__(
            'keycloak:authentication/executionConfig:ExecutionConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            config: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            execution_id: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'ExecutionConfig':
        """
        Get an existing ExecutionConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alias: The name of the configuration.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] config: The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
        :param pulumi.Input[str] execution_id: The authentication execution this configuration is attached to.
        :param pulumi.Input[str] realm_id: The realm the authentication execution exists in.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ExecutionConfigState.__new__(_ExecutionConfigState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["config"] = config
        __props__.__dict__["execution_id"] = execution_id
        __props__.__dict__["realm_id"] = realm_id
        return ExecutionConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        """
        The name of the configuration.
        """
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The configuration. Keys are specific to each configurable authentication execution and not checked when applying.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="executionId")
    def execution_id(self) -> pulumi.Output[str]:
        """
        The authentication execution this configuration is attached to.
        """
        return pulumi.get(self, "execution_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm the authentication execution exists in.
        """
        return pulumi.get(self, "realm_id")

