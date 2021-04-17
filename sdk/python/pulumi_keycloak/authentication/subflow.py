# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SubflowArgs', 'Subflow']

@pulumi.input_type
class SubflowArgs:
    def __init__(__self__, *,
                 alias: pulumi.Input[str],
                 parent_flow_alias: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 authenticator: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 requirement: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Subflow resource.
        :param pulumi.Input[str] authenticator: Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
        """
        pulumi.set(__self__, "alias", alias)
        pulumi.set(__self__, "parent_flow_alias", parent_flow_alias)
        pulumi.set(__self__, "realm_id", realm_id)
        if authenticator is not None:
            pulumi.set(__self__, "authenticator", authenticator)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)
        if requirement is not None:
            pulumi.set(__self__, "requirement", requirement)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Input[str]:
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter(name="parentFlowAlias")
    def parent_flow_alias(self) -> pulumi.Input[str]:
        return pulumi.get(self, "parent_flow_alias")

    @parent_flow_alias.setter
    def parent_flow_alias(self, value: pulumi.Input[str]):
        pulumi.set(self, "parent_flow_alias", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def authenticator(self) -> Optional[pulumi.Input[str]]:
        """
        Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
        """
        return pulumi.get(self, "authenticator")

    @authenticator.setter
    def authenticator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authenticator", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_id", value)

    @property
    @pulumi.getter
    def requirement(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "requirement")

    @requirement.setter
    def requirement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "requirement", value)


@pulumi.input_type
class _SubflowState:
    def __init__(__self__, *,
                 alias: Optional[pulumi.Input[str]] = None,
                 authenticator: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 parent_flow_alias: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 requirement: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Subflow resources.
        :param pulumi.Input[str] authenticator: Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
        """
        if alias is not None:
            pulumi.set(__self__, "alias", alias)
        if authenticator is not None:
            pulumi.set(__self__, "authenticator", authenticator)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if parent_flow_alias is not None:
            pulumi.set(__self__, "parent_flow_alias", parent_flow_alias)
        if provider_id is not None:
            pulumi.set(__self__, "provider_id", provider_id)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if requirement is not None:
            pulumi.set(__self__, "requirement", requirement)

    @property
    @pulumi.getter
    def alias(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "alias")

    @alias.setter
    def alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alias", value)

    @property
    @pulumi.getter
    def authenticator(self) -> Optional[pulumi.Input[str]]:
        """
        Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
        """
        return pulumi.get(self, "authenticator")

    @authenticator.setter
    def authenticator(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "authenticator", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="parentFlowAlias")
    def parent_flow_alias(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "parent_flow_alias")

    @parent_flow_alias.setter
    def parent_flow_alias(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "parent_flow_alias", value)

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "provider_id")

    @provider_id.setter
    def provider_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "provider_id", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def requirement(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "requirement")

    @requirement.setter
    def requirement(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "requirement", value)


class Subflow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 authenticator: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 parent_flow_alias: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 requirement: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Subflow resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authenticator: Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubflowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Subflow resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SubflowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubflowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alias: Optional[pulumi.Input[str]] = None,
                 authenticator: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 parent_flow_alias: Optional[pulumi.Input[str]] = None,
                 provider_id: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 requirement: Optional[pulumi.Input[str]] = None,
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
            __props__ = SubflowArgs.__new__(SubflowArgs)

            if alias is None and not opts.urn:
                raise TypeError("Missing required property 'alias'")
            __props__.__dict__["alias"] = alias
            __props__.__dict__["authenticator"] = authenticator
            __props__.__dict__["description"] = description
            if parent_flow_alias is None and not opts.urn:
                raise TypeError("Missing required property 'parent_flow_alias'")
            __props__.__dict__["parent_flow_alias"] = parent_flow_alias
            __props__.__dict__["provider_id"] = provider_id
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            __props__.__dict__["requirement"] = requirement
        super(Subflow, __self__).__init__(
            'keycloak:authentication/subflow:Subflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alias: Optional[pulumi.Input[str]] = None,
            authenticator: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            parent_flow_alias: Optional[pulumi.Input[str]] = None,
            provider_id: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            requirement: Optional[pulumi.Input[str]] = None) -> 'Subflow':
        """
        Get an existing Subflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] authenticator: Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SubflowState.__new__(_SubflowState)

        __props__.__dict__["alias"] = alias
        __props__.__dict__["authenticator"] = authenticator
        __props__.__dict__["description"] = description
        __props__.__dict__["parent_flow_alias"] = parent_flow_alias
        __props__.__dict__["provider_id"] = provider_id
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["requirement"] = requirement
        return Subflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alias(self) -> pulumi.Output[str]:
        return pulumi.get(self, "alias")

    @property
    @pulumi.getter
    def authenticator(self) -> pulumi.Output[Optional[str]]:
        """
        Might be needed to be set with certain custom subflow with specific authenticator, in general this will remain empty
        """
        return pulumi.get(self, "authenticator")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="parentFlowAlias")
    def parent_flow_alias(self) -> pulumi.Output[str]:
        return pulumi.get(self, "parent_flow_alias")

    @property
    @pulumi.getter(name="providerId")
    def provider_id(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "provider_id")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def requirement(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "requirement")

