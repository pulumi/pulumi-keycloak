# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class ClientAuthorizationResource(pulumi.CustomResource):
    attributes: pulumi.Output[dict]
    display_name: pulumi.Output[str]
    icon_uri: pulumi.Output[str]
    name: pulumi.Output[str]
    owner_managed_access: pulumi.Output[bool]
    realm_id: pulumi.Output[str]
    resource_server_id: pulumi.Output[str]
    scopes: pulumi.Output[list]
    type: pulumi.Output[str]
    uris: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, attributes=None, display_name=None, icon_uri=None, name=None, owner_managed_access=None, realm_id=None, resource_server_id=None, scopes=None, type=None, uris=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a ClientAuthorizationResource resource with the given unique name, props, and options.
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

            __props__['attributes'] = attributes
            __props__['display_name'] = display_name
            __props__['icon_uri'] = icon_uri
            __props__['name'] = name
            __props__['owner_managed_access'] = owner_managed_access
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
            if resource_server_id is None:
                raise TypeError("Missing required property 'resource_server_id'")
            __props__['resource_server_id'] = resource_server_id
            __props__['scopes'] = scopes
            __props__['type'] = type
            __props__['uris'] = uris
        super(ClientAuthorizationResource, __self__).__init__(
            'keycloak:OpenId/clientAuthorizationResource:ClientAuthorizationResource',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, attributes=None, display_name=None, icon_uri=None, name=None, owner_managed_access=None, realm_id=None, resource_server_id=None, scopes=None, type=None, uris=None):
        """
        Get an existing ClientAuthorizationResource resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["attributes"] = attributes
        __props__["display_name"] = display_name
        __props__["icon_uri"] = icon_uri
        __props__["name"] = name
        __props__["owner_managed_access"] = owner_managed_access
        __props__["realm_id"] = realm_id
        __props__["resource_server_id"] = resource_server_id
        __props__["scopes"] = scopes
        __props__["type"] = type
        __props__["uris"] = uris
        return ClientAuthorizationResource(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

