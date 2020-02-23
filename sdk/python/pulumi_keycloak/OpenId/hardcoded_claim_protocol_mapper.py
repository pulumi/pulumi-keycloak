# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class HardcodedClaimProtocolMapper(pulumi.CustomResource):
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
    claim_value: pulumi.Output[str]
    claim_value_type: pulumi.Output[str]
    """
    Claim type used when serializing tokens.
    """
    client_id: pulumi.Output[str]
    """
    The mapper's associated client. Cannot be used at the same time as client_scope_id.
    """
    client_scope_id: pulumi.Output[str]
    """
    The mapper's associated client scope. Cannot be used at the same time as client_id.
    """
    name: pulumi.Output[str]
    """
    A human-friendly name that will appear in the Keycloak console.
    """
    realm_id: pulumi.Output[str]
    """
    The realm id where the associated client or client scope exists.
    """
    def __init__(__self__, resource_name, opts=None, add_to_access_token=None, add_to_id_token=None, add_to_userinfo=None, claim_name=None, claim_value=None, claim_value_type=None, client_id=None, client_scope_id=None, name=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # OpenId.HardcodedClaimProtocolMapper

        Allows for creating and managing hardcoded claim protocol mappers within
        Keycloak.

        Hardcoded claim protocol mappers allow you to define a claim with a hardcoded
        value. Protocol mappers can be defined for a single client, or they can
        be defined for a client scope which can be shared between multiple different
        clients.

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
        - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `claim_name` - (Required) The name of the claim to insert into a token.
        - `claim_value` - (Required) The hardcoded value of the claim.
        - `claim_value_type` - (Optional) The claim type used when serializing JSON tokens. Can be one of `String`, `long`, `int`, or `boolean`. Defaults to `String`.
        - `add_to_id_token` - (Optional) Indicates if the property should be added as a claim to the id token. Defaults to `true`.
        - `add_to_access_token` - (Optional) Indicates if the property should be added as a claim to the access token. Defaults to `true`.
        - `add_to_userinfo` - (Optional) Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.

        > This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_hardcoded_claim_protocol_mapper.html.markdown.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the attribute should be a claim in the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the attribute should be a claim in the id token.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the attribute should appear in the userinfo response body.
        :param pulumi.Input[str] claim_value_type: Claim type used when serializing tokens.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
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
            if claim_value is None:
                raise TypeError("Missing required property 'claim_value'")
            __props__['claim_value'] = claim_value
            __props__['claim_value_type'] = claim_value_type
            __props__['client_id'] = client_id
            __props__['client_scope_id'] = client_scope_id
            __props__['name'] = name
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(HardcodedClaimProtocolMapper, __self__).__init__(
            'keycloak:OpenId/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, add_to_access_token=None, add_to_id_token=None, add_to_userinfo=None, claim_name=None, claim_value=None, claim_value_type=None, client_id=None, client_scope_id=None, name=None, realm_id=None):
        """
        Get an existing HardcodedClaimProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the attribute should be a claim in the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the attribute should be a claim in the id token.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the attribute should appear in the userinfo response body.
        :param pulumi.Input[str] claim_value_type: Claim type used when serializing tokens.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_to_access_token"] = add_to_access_token
        __props__["add_to_id_token"] = add_to_id_token
        __props__["add_to_userinfo"] = add_to_userinfo
        __props__["claim_name"] = claim_name
        __props__["claim_value"] = claim_value
        __props__["claim_value_type"] = claim_value_type
        __props__["client_id"] = client_id
        __props__["client_scope_id"] = client_scope_id
        __props__["name"] = name
        __props__["realm_id"] = realm_id
        return HardcodedClaimProtocolMapper(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

