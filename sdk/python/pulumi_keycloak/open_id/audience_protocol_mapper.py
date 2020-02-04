# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AudienceProtocolMapper(pulumi.CustomResource):
    add_to_access_token: pulumi.Output[bool]
    add_to_id_token: pulumi.Output[bool]
    client_id: pulumi.Output[str]
    client_scope_id: pulumi.Output[str]
    included_client_audience: pulumi.Output[str]
    included_custom_audience: pulumi.Output[str]
    name: pulumi.Output[str]
    realm_id: pulumi.Output[str]
    def __init__(__self__, resource_name, opts=None, add_to_access_token=None, add_to_id_token=None, client_id=None, client_scope_id=None, included_client_audience=None, included_custom_audience=None, name=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a AudienceProtocolMapper resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['add_to_access_token'] = add_to_access_token
            __props__['add_to_id_token'] = add_to_id_token
            __props__['client_id'] = client_id
            __props__['client_scope_id'] = client_scope_id
            __props__['included_client_audience'] = included_client_audience
            __props__['included_custom_audience'] = included_custom_audience
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(AudienceProtocolMapper, __self__).__init__(
            'keycloak:OpenId/audienceProtocolMapper:AudienceProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, add_to_access_token=None, add_to_id_token=None, client_id=None, client_scope_id=None, included_client_audience=None, included_custom_audience=None, name=None, realm_id=None):
        """
        Get an existing AudienceProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["add_to_access_token"] = add_to_access_token
        __props__["add_to_id_token"] = add_to_id_token
        __props__["client_id"] = client_id
        __props__["client_scope_id"] = client_scope_id
        __props__["included_client_audience"] = included_client_audience
        __props__["included_custom_audience"] = included_custom_audience
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        return AudienceProtocolMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

