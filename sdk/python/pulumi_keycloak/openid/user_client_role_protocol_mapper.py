# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class UserClientRoleProtocolMapper(pulumi.CustomResource):
    add_to_access_token: pulumi.Output[bool]
    """
    Indicates if the attribute should be a claim in the access token.
    """
    add_to_id_token: pulumi.Output[bool]
    """
    Indicates if the attribute should be a claim in the id token.
    """
    add_to_userinfo: pulumi.Output[bool]
    """
    Indicates if the attribute should appear in the userinfo response body.
    """
    claim_name: pulumi.Output[str]
    claim_value_type: pulumi.Output[str]
    """
    Claim type used when serializing tokens.
    """
    client_id: pulumi.Output[str]
    """
    The mapper's associated client. Cannot be used at the same time as client_scope_id.
    """
    client_id_for_role_mappings: pulumi.Output[str]
    """
    Client ID for role mappings.
    """
    client_role_prefix: pulumi.Output[str]
    """
    Prefix that will be added to each client role.
    """
    client_scope_id: pulumi.Output[str]
    """
    The mapper's associated client scope. Cannot be used at the same time as client_id.
    """
    multivalued: pulumi.Output[bool]
    """
    Indicates whether this attribute is a single value or an array of values.
    """
    name: pulumi.Output[str]
    """
    A human-friendly name that will appear in the Keycloak console.
    """
    realm_id: pulumi.Output[str]
    """
    The realm id where the associated client or client scope exists.
    """
    def __init__(__self__, resource_name, opts=None, add_to_access_token=None, add_to_id_token=None, add_to_userinfo=None, claim_name=None, claim_value_type=None, client_id=None, client_id_for_role_mappings=None, client_role_prefix=None, client_scope_id=None, multivalued=None, name=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a UserClientRoleProtocolMapper resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the attribute should be a claim in the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the attribute should be a claim in the id token.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the attribute should appear in the userinfo response body.
        :param pulumi.Input[str] claim_value_type: Claim type used when serializing tokens.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_id_for_role_mappings: Client ID for role mappings.
        :param pulumi.Input[str] client_role_prefix: Prefix that will be added to each client role.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[bool] multivalued: Indicates whether this attribute is a single value or an array of values.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
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
            __props__['add_to_userinfo'] = add_to_userinfo
            if claim_name is None:
                raise TypeError("Missing required property 'claim_name'")
            __props__['claim_name'] = claim_name
            __props__['claim_value_type'] = claim_value_type
            __props__['client_id'] = client_id
            __props__['client_id_for_role_mappings'] = client_id_for_role_mappings
            __props__['client_role_prefix'] = client_role_prefix
            __props__['client_scope_id'] = client_scope_id
            __props__['multivalued'] = multivalued
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(UserClientRoleProtocolMapper, __self__).__init__(
            'keycloak:openid/userClientRoleProtocolMapper:UserClientRoleProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, add_to_access_token=None, add_to_id_token=None, add_to_userinfo=None, claim_name=None, claim_value_type=None, client_id=None, client_id_for_role_mappings=None, client_role_prefix=None, client_scope_id=None, multivalued=None, name=None, realm_id=None):
        """
        Get an existing UserClientRoleProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the attribute should be a claim in the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the attribute should be a claim in the id token.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the attribute should appear in the userinfo response body.
        :param pulumi.Input[str] claim_value_type: Claim type used when serializing tokens.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_id_for_role_mappings: Client ID for role mappings.
        :param pulumi.Input[str] client_role_prefix: Prefix that will be added to each client role.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[bool] multivalued: Indicates whether this attribute is a single value or an array of values.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_to_access_token"] = add_to_access_token
        __props__["add_to_id_token"] = add_to_id_token
        __props__["add_to_userinfo"] = add_to_userinfo
        __props__["claim_name"] = claim_name
        __props__["claim_value_type"] = claim_value_type
        __props__["client_id"] = client_id
        __props__["client_id_for_role_mappings"] = client_id_for_role_mappings
        __props__["client_role_prefix"] = client_role_prefix
        __props__["client_scope_id"] = client_scope_id
        __props__["multivalued"] = multivalued
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        return UserClientRoleProtocolMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

