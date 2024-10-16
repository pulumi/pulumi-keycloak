# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['AudienceProtocolMapperArgs', 'AudienceProtocolMapper']

@pulumi.input_type
class AudienceProtocolMapperArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 included_client_audience: Optional[pulumi.Input[str]] = None,
                 included_custom_audience: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AudienceProtocolMapper resource.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        :param pulumi.Input[bool] add_to_access_token: Indicates if this claim should be added to the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if this claim should be added to the id token.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] included_client_audience: A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        :param pulumi.Input[str] included_custom_audience: A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        if add_to_access_token is not None:
            pulumi.set(__self__, "add_to_access_token", add_to_access_token)
        if add_to_id_token is not None:
            pulumi.set(__self__, "add_to_id_token", add_to_id_token)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if included_client_audience is not None:
            pulumi.set(__self__, "included_client_audience", included_client_audience)
        if included_custom_audience is not None:
            pulumi.set(__self__, "included_custom_audience", included_custom_audience)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm id where the associated client or client scope exists.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="addToAccessToken")
    def add_to_access_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if this claim should be added to the access token.
        """
        return pulumi.get(self, "add_to_access_token")

    @add_to_access_token.setter
    def add_to_access_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_access_token", value)

    @property
    @pulumi.getter(name="addToIdToken")
    def add_to_id_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if this claim should be added to the id token.
        """
        return pulumi.get(self, "add_to_id_token")

    @add_to_id_token.setter
    def add_to_id_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_id_token", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The mapper's associated client. Cannot be used at the same time as client_scope_id.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The mapper's associated client scope. Cannot be used at the same time as client_id.
        """
        return pulumi.get(self, "client_scope_id")

    @client_scope_id.setter
    def client_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_scope_id", value)

    @property
    @pulumi.getter(name="includedClientAudience")
    def included_client_audience(self) -> Optional[pulumi.Input[str]]:
        """
        A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        """
        return pulumi.get(self, "included_client_audience")

    @included_client_audience.setter
    def included_client_audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "included_client_audience", value)

    @property
    @pulumi.getter(name="includedCustomAudience")
    def included_custom_audience(self) -> Optional[pulumi.Input[str]]:
        """
        A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        """
        return pulumi.get(self, "included_custom_audience")

    @included_custom_audience.setter
    def included_custom_audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "included_custom_audience", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name that will appear in the Keycloak console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AudienceProtocolMapperState:
    def __init__(__self__, *,
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 included_client_audience: Optional[pulumi.Input[str]] = None,
                 included_custom_audience: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AudienceProtocolMapper resources.
        :param pulumi.Input[bool] add_to_access_token: Indicates if this claim should be added to the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if this claim should be added to the id token.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] included_client_audience: A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        :param pulumi.Input[str] included_custom_audience: A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        """
        if add_to_access_token is not None:
            pulumi.set(__self__, "add_to_access_token", add_to_access_token)
        if add_to_id_token is not None:
            pulumi.set(__self__, "add_to_id_token", add_to_id_token)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if included_client_audience is not None:
            pulumi.set(__self__, "included_client_audience", included_client_audience)
        if included_custom_audience is not None:
            pulumi.set(__self__, "included_custom_audience", included_custom_audience)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="addToAccessToken")
    def add_to_access_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if this claim should be added to the access token.
        """
        return pulumi.get(self, "add_to_access_token")

    @add_to_access_token.setter
    def add_to_access_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_access_token", value)

    @property
    @pulumi.getter(name="addToIdToken")
    def add_to_id_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if this claim should be added to the id token.
        """
        return pulumi.get(self, "add_to_id_token")

    @add_to_id_token.setter
    def add_to_id_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_id_token", value)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        The mapper's associated client. Cannot be used at the same time as client_scope_id.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The mapper's associated client scope. Cannot be used at the same time as client_id.
        """
        return pulumi.get(self, "client_scope_id")

    @client_scope_id.setter
    def client_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_scope_id", value)

    @property
    @pulumi.getter(name="includedClientAudience")
    def included_client_audience(self) -> Optional[pulumi.Input[str]]:
        """
        A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        """
        return pulumi.get(self, "included_client_audience")

    @included_client_audience.setter
    def included_client_audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "included_client_audience", value)

    @property
    @pulumi.getter(name="includedCustomAudience")
    def included_custom_audience(self) -> Optional[pulumi.Input[str]]:
        """
        A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        """
        return pulumi.get(self, "included_custom_audience")

    @included_custom_audience.setter
    def included_custom_audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "included_custom_audience", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly name that will appear in the Keycloak console.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm id where the associated client or client scope exists.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class AudienceProtocolMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 included_client_audience: Optional[pulumi.Input[str]] = None,
                 included_custom_audience: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # openid.AudienceProtocolMapper

        Allows for creating and managing audience protocol mappers within
        Keycloak. This mapper was added in Keycloak v4.6.0.Final.

        Audience protocol mappers allow you add audiences to the `aud` claim
        within issued tokens. The audience can be a custom string, or it can be
        mapped to the ID of a pre-existing client.

        ### Example Usage (Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        openid_client = keycloak.openid.Client("openid_client",
            realm_id=realm.id,
            client_id="test-client",
            name="test client",
            enabled=True,
            access_type="CONFIDENTIAL",
            valid_redirect_uris=["http://localhost:8080/openid-callback"])
        audience_mapper = keycloak.openid.AudienceProtocolMapper("audience_mapper",
            realm_id=realm.id,
            client_id=openid_client.id,
            name="audience-mapper",
            included_custom_audience="foo")
        ```

        ### Example Usage (Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("client_scope",
            realm_id=realm.id,
            name="test-client-scope")
        audience_mapper = keycloak.openid.AudienceProtocolMapper("audience_mapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id,
            name="audience-mapper",
            included_custom_audience="foo")
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
        - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `included_client_audience` - (Required if `included_custom_audience` is not specified) A client ID to include within the token's `aud` claim.
        - `included_custom_audience` - (Required if `included_client_audience` is not specified) A custom audience to include within the token's `aud` claim.
        - `add_to_id_token` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
        - `add_to_access_token` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.

        ### Import

        Protocol mappers can be imported using one of the following formats:
        - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
        - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`

        Example:

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if this claim should be added to the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if this claim should be added to the id token.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] included_client_audience: A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        :param pulumi.Input[str] included_custom_audience: A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AudienceProtocolMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # openid.AudienceProtocolMapper

        Allows for creating and managing audience protocol mappers within
        Keycloak. This mapper was added in Keycloak v4.6.0.Final.

        Audience protocol mappers allow you add audiences to the `aud` claim
        within issued tokens. The audience can be a custom string, or it can be
        mapped to the ID of a pre-existing client.

        ### Example Usage (Client)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        openid_client = keycloak.openid.Client("openid_client",
            realm_id=realm.id,
            client_id="test-client",
            name="test client",
            enabled=True,
            access_type="CONFIDENTIAL",
            valid_redirect_uris=["http://localhost:8080/openid-callback"])
        audience_mapper = keycloak.openid.AudienceProtocolMapper("audience_mapper",
            realm_id=realm.id,
            client_id=openid_client.id,
            name="audience-mapper",
            included_custom_audience="foo")
        ```

        ### Example Usage (Client Scope)

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("client_scope",
            realm_id=realm.id,
            name="test-client-scope")
        audience_mapper = keycloak.openid.AudienceProtocolMapper("audience_mapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id,
            name="audience-mapper",
            included_custom_audience="foo")
        ```

        ### Argument Reference

        The following arguments are supported:

        - `realm_id` - (Required) The realm this protocol mapper exists within.
        - `client_id` - (Required if `client_scope_id` is not specified) The client this protocol mapper is attached to.
        - `client_scope_id` - (Required if `client_id` is not specified) The client scope this protocol mapper is attached to.
        - `name` - (Required) The display name of this protocol mapper in the GUI.
        - `included_client_audience` - (Required if `included_custom_audience` is not specified) A client ID to include within the token's `aud` claim.
        - `included_custom_audience` - (Required if `included_client_audience` is not specified) A custom audience to include within the token's `aud` claim.
        - `add_to_id_token` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
        - `add_to_access_token` - (Optional) Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.

        ### Import

        Protocol mappers can be imported using one of the following formats:
        - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
        - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`

        Example:

        :param str resource_name: The name of the resource.
        :param AudienceProtocolMapperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AudienceProtocolMapperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 included_client_audience: Optional[pulumi.Input[str]] = None,
                 included_custom_audience: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AudienceProtocolMapperArgs.__new__(AudienceProtocolMapperArgs)

            __props__.__dict__["add_to_access_token"] = add_to_access_token
            __props__.__dict__["add_to_id_token"] = add_to_id_token
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_scope_id"] = client_scope_id
            __props__.__dict__["included_client_audience"] = included_client_audience
            __props__.__dict__["included_custom_audience"] = included_custom_audience
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(AudienceProtocolMapper, __self__).__init__(
            'keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_to_access_token: Optional[pulumi.Input[bool]] = None,
            add_to_id_token: Optional[pulumi.Input[bool]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_scope_id: Optional[pulumi.Input[str]] = None,
            included_client_audience: Optional[pulumi.Input[str]] = None,
            included_custom_audience: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'AudienceProtocolMapper':
        """
        Get an existing AudienceProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] add_to_access_token: Indicates if this claim should be added to the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if this claim should be added to the id token.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] included_client_audience: A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        :param pulumi.Input[str] included_custom_audience: A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AudienceProtocolMapperState.__new__(_AudienceProtocolMapperState)

        __props__.__dict__["add_to_access_token"] = add_to_access_token
        __props__.__dict__["add_to_id_token"] = add_to_id_token
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_scope_id"] = client_scope_id
        __props__.__dict__["included_client_audience"] = included_client_audience
        __props__.__dict__["included_custom_audience"] = included_custom_audience
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        return AudienceProtocolMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addToAccessToken")
    def add_to_access_token(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if this claim should be added to the access token.
        """
        return pulumi.get(self, "add_to_access_token")

    @property
    @pulumi.getter(name="addToIdToken")
    def add_to_id_token(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if this claim should be added to the id token.
        """
        return pulumi.get(self, "add_to_id_token")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The mapper's associated client. Cannot be used at the same time as client_scope_id.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientScopeId")
    def client_scope_id(self) -> pulumi.Output[Optional[str]]:
        """
        The mapper's associated client scope. Cannot be used at the same time as client_id.
        """
        return pulumi.get(self, "client_scope_id")

    @property
    @pulumi.getter(name="includedClientAudience")
    def included_client_audience(self) -> pulumi.Output[Optional[str]]:
        """
        A client ID to include within the token's `aud` claim. Cannot be used with included_custom_audience
        """
        return pulumi.get(self, "included_client_audience")

    @property
    @pulumi.getter(name="includedCustomAudience")
    def included_custom_audience(self) -> pulumi.Output[Optional[str]]:
        """
        A custom audience to include within the token's `aud` claim. Cannot be used with included_custom_audience
        """
        return pulumi.get(self, "included_custom_audience")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A human-friendly name that will appear in the Keycloak console.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm id where the associated client or client scope exists.
        """
        return pulumi.get(self, "realm_id")

