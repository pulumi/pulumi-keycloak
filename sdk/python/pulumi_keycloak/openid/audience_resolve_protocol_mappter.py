# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AudienceResolveProtocolMappterArgs', 'AudienceResolveProtocolMappter']

@pulumi.input_type
class AudienceResolveProtocolMappterArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AudienceResolveProtocolMappter resource.
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
        :param pulumi.Input[str] client_id: The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] client_scope_id: The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        """
        pulumi.set(__self__, "realm_id", realm_id)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm this protocol mapper exists within.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_scope_id")

    @client_scope_id.setter
    def client_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_scope_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AudienceResolveProtocolMappterState:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AudienceResolveProtocolMappter resources.
        :param pulumi.Input[str] client_id: The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] client_scope_id: The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_scope_id")

    @client_scope_id.setter
    def client_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_scope_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm this protocol mapper exists within.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


warnings.warn("""keycloak.openid/audienceresolveprotocolmappter.AudienceResolveProtocolMappter has been deprecated in favor of keycloak.openid/audienceresolveprotocolmapper.AudienceResolveProtocolMapper""", DeprecationWarning)


class AudienceResolveProtocolMappter(pulumi.CustomResource):
    warnings.warn("""keycloak.openid/audienceresolveprotocolmappter.AudienceResolveProtocolMappter has been deprecated in favor of keycloak.openid/audienceresolveprotocolmapper.AudienceResolveProtocolMapper""", DeprecationWarning)

    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for creating the "Audience Resolve" OIDC protocol mapper within Keycloak.

        This protocol mapper is useful to avoid manual management of audiences, instead relying on the presence of client roles
        to imply which audiences are appropriate for the token. See the
        [Keycloak docs](https://www.keycloak.org/docs/latest/server_admin/#_audience_resolve) for more details.

        ## Example Usage
        ### Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        openid_client = keycloak.openid.Client("openidClient",
            realm_id=realm.id,
            client_id="client",
            enabled=True,
            access_type="CONFIDENTIAL",
            valid_redirect_uris=["http://localhost:8080/openid-callback"])
        audience_mapper = keycloak.openid.AudienceResolveProtocolMapper("audienceMapper",
            realm_id=realm.id,
            client_id=openid_client.id)
        ```
        ### Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("clientScope", realm_id=realm.id)
        audience_mapper = keycloak.openid.AudienceProtocolMapper("audienceMapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id)
        ```

        ## Import

        Protocol mappers can be imported using one of the following formats:

         - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`

         - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`

         Example:

         bash

        ```sh
        $ pulumi import keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter audience_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
        ```

        ```sh
        $ pulumi import keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter audience_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] client_scope_id: The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AudienceResolveProtocolMappterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating the "Audience Resolve" OIDC protocol mapper within Keycloak.

        This protocol mapper is useful to avoid manual management of audiences, instead relying on the presence of client roles
        to imply which audiences are appropriate for the token. See the
        [Keycloak docs](https://www.keycloak.org/docs/latest/server_admin/#_audience_resolve) for more details.

        ## Example Usage
        ### Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        openid_client = keycloak.openid.Client("openidClient",
            realm_id=realm.id,
            client_id="client",
            enabled=True,
            access_type="CONFIDENTIAL",
            valid_redirect_uris=["http://localhost:8080/openid-callback"])
        audience_mapper = keycloak.openid.AudienceResolveProtocolMapper("audienceMapper",
            realm_id=realm.id,
            client_id=openid_client.id)
        ```
        ### Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("clientScope", realm_id=realm.id)
        audience_mapper = keycloak.openid.AudienceProtocolMapper("audienceMapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id)
        ```

        ## Import

        Protocol mappers can be imported using one of the following formats:

         - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`

         - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`

         Example:

         bash

        ```sh
        $ pulumi import keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter audience_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
        ```

        ```sh
        $ pulumi import keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter audience_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
        ```

        :param str resource_name: The name of the resource.
        :param AudienceResolveProtocolMappterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AudienceResolveProtocolMappterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        pulumi.log.warn("""AudienceResolveProtocolMappter is deprecated: keycloak.openid/audienceresolveprotocolmappter.AudienceResolveProtocolMappter has been deprecated in favor of keycloak.openid/audienceresolveprotocolmapper.AudienceResolveProtocolMapper""")
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AudienceResolveProtocolMappterArgs.__new__(AudienceResolveProtocolMappterArgs)

            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_scope_id"] = client_scope_id
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(AudienceResolveProtocolMappter, __self__).__init__(
            'keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_scope_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'AudienceResolveProtocolMappter':
        """
        Get an existing AudienceResolveProtocolMappter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] client_scope_id: The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        :param pulumi.Input[str] name: The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        :param pulumi.Input[str] realm_id: The realm this protocol mapper exists within.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AudienceResolveProtocolMappterState.__new__(_AudienceResolveProtocolMappterState)

        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_scope_id"] = client_scope_id
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        return AudienceResolveProtocolMappter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        """
        return pulumi.get(self, "client_scope_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm this protocol mapper exists within.
        """
        return pulumi.get(self, "realm_id")

