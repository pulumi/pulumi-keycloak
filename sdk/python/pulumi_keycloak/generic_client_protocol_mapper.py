# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables


class GenericClientProtocolMapper(pulumi.CustomResource):
    client_id: pulumi.Output[str]
    """
    The mapper's associated client. Cannot be used at the same time as client_scope_id.
    """
    client_scope_id: pulumi.Output[str]
    """
    The mapper's associated client scope. Cannot be used at the same time as client_id.
    """
    config: pulumi.Output[dict]
    name: pulumi.Output[str]
    """
    A human-friendly name that will appear in the Keycloak console.
    """
    protocol: pulumi.Output[str]
    """
    The protocol of the client (openid-connect / saml).
    """
    protocol_mapper: pulumi.Output[str]
    """
    The type of the protocol mapper.
    """
    realm_id: pulumi.Output[str]
    """
    The realm id where the associated client or client scope exists.
    """
    def __init__(__self__, resource_name, opts=None, client_id=None, client_scope_id=None, config=None, name=None, protocol=None, protocol_mapper=None, realm_id=None, __props__=None, __name__=None, __opts__=None):
        """
        ## # GenericClientProtocolMapper

        Allows for creating and managing protocol mapper for both types of clients (openid-connect and saml) within Keycloak.

        There are two uses cases for using this resource:
        * If you implemented a custom protocol mapper, this resource can be used to configure it
        * If the provider doesn't support a particular protocol mapper, this resource can be used instead.

        Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
        Therefore, if possible, a specific mapper should be used.

        ### Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            enabled=True,
            realm="my-realm")
        saml_client = keycloak.saml.Client("samlClient",
            client_id="test-client",
            realm_id=realm.id)
        saml_hardcode_attribute_mapper = keycloak.GenericClientProtocolMapper("samlHardcodeAttributeMapper",
            client_id=saml_client.id,
            config={
                "attribute.name": "name",
                "attribute.nameformat": "Basic",
                "attribute.value": "value",
                "friendly.name": "display name",
            },
            protocol="saml",
            protocol_mapper="saml-hardcode-attribute-mapper",
            realm_id=realm.id)
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required) The client this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `protocol` - (Required) The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
        - `protocol_mapper` - (Required) The name of the protocol mapper. The protocol mapper must be
           compatible with the specified client.
        - `config` - (Required) A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] protocol: The protocol of the client (openid-connect / saml).
        :param pulumi.Input[str] protocol_mapper: The type of the protocol mapper.
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

            __props__['client_id'] = client_id
            __props__['client_scope_id'] = client_scope_id
            if config is None:
                raise TypeError("Missing required property 'config'")
            __props__['config'] = config
            __props__['name'] = name
            if protocol is None:
                raise TypeError("Missing required property 'protocol'")
            __props__['protocol'] = protocol
            if protocol_mapper is None:
                raise TypeError("Missing required property 'protocol_mapper'")
            __props__['protocol_mapper'] = protocol_mapper
            if realm_id is None:
                raise TypeError("Missing required property 'realm_id'")
            __props__['realm_id'] = realm_id
        super(GenericClientProtocolMapper, __self__).__init__(
            'keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, client_id=None, client_scope_id=None, config=None, name=None, protocol=None, protocol_mapper=None, realm_id=None):
        """
        Get an existing GenericClientProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] protocol: The protocol of the client (openid-connect / saml).
        :param pulumi.Input[str] protocol_mapper: The type of the protocol mapper.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["client_scope_id"] = client_scope_id
        __props__["config"] = config
        __props__["name"] = name
        __props__["protocol"] = protocol
        __props__["protocol_mapper"] = protocol_mapper
        __props__["realm_id"] = realm_id
        return GenericClientProtocolMapper(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
