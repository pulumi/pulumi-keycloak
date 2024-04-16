# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['HardcodedClaimProtocolMapperArgs', 'HardcodedClaimProtocolMapper']

@pulumi.input_type
class HardcodedClaimProtocolMapperArgs:
    def __init__(__self__, *,
                 claim_name: pulumi.Input[str],
                 claim_value: pulumi.Input[str],
                 realm_id: pulumi.Input[str],
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 add_to_userinfo: Optional[pulumi.Input[bool]] = None,
                 claim_value_type: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a HardcodedClaimProtocolMapper resource.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the attribute should be a claim in the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the attribute should be a claim in the id token.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the attribute should appear in the userinfo response body.
        :param pulumi.Input[str] claim_value_type: Claim type used when serializing tokens.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        """
        pulumi.set(__self__, "claim_name", claim_name)
        pulumi.set(__self__, "claim_value", claim_value)
        pulumi.set(__self__, "realm_id", realm_id)
        if add_to_access_token is not None:
            pulumi.set(__self__, "add_to_access_token", add_to_access_token)
        if add_to_id_token is not None:
            pulumi.set(__self__, "add_to_id_token", add_to_id_token)
        if add_to_userinfo is not None:
            pulumi.set(__self__, "add_to_userinfo", add_to_userinfo)
        if claim_value_type is not None:
            pulumi.set(__self__, "claim_value_type", claim_value_type)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="claimName")
    def claim_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "claim_name")

    @claim_name.setter
    def claim_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "claim_name", value)

    @property
    @pulumi.getter(name="claimValue")
    def claim_value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "claim_value")

    @claim_value.setter
    def claim_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "claim_value", value)

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
        Indicates if the attribute should be a claim in the access token.
        """
        return pulumi.get(self, "add_to_access_token")

    @add_to_access_token.setter
    def add_to_access_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_access_token", value)

    @property
    @pulumi.getter(name="addToIdToken")
    def add_to_id_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the attribute should be a claim in the id token.
        """
        return pulumi.get(self, "add_to_id_token")

    @add_to_id_token.setter
    def add_to_id_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_id_token", value)

    @property
    @pulumi.getter(name="addToUserinfo")
    def add_to_userinfo(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the attribute should appear in the userinfo response body.
        """
        return pulumi.get(self, "add_to_userinfo")

    @add_to_userinfo.setter
    def add_to_userinfo(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_userinfo", value)

    @property
    @pulumi.getter(name="claimValueType")
    def claim_value_type(self) -> Optional[pulumi.Input[str]]:
        """
        Claim type used when serializing tokens.
        """
        return pulumi.get(self, "claim_value_type")

    @claim_value_type.setter
    def claim_value_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "claim_value_type", value)

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
class _HardcodedClaimProtocolMapperState:
    def __init__(__self__, *,
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 add_to_userinfo: Optional[pulumi.Input[bool]] = None,
                 claim_name: Optional[pulumi.Input[str]] = None,
                 claim_value: Optional[pulumi.Input[str]] = None,
                 claim_value_type: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering HardcodedClaimProtocolMapper resources.
        :param pulumi.Input[bool] add_to_access_token: Indicates if the attribute should be a claim in the access token.
        :param pulumi.Input[bool] add_to_id_token: Indicates if the attribute should be a claim in the id token.
        :param pulumi.Input[bool] add_to_userinfo: Indicates if the attribute should appear in the userinfo response body.
        :param pulumi.Input[str] claim_value_type: Claim type used when serializing tokens.
        :param pulumi.Input[str] client_id: The mapper's associated client. Cannot be used at the same time as client_scope_id.
        :param pulumi.Input[str] client_scope_id: The mapper's associated client scope. Cannot be used at the same time as client_id.
        :param pulumi.Input[str] name: A human-friendly name that will appear in the Keycloak console.
        :param pulumi.Input[str] realm_id: The realm id where the associated client or client scope exists.
        """
        if add_to_access_token is not None:
            pulumi.set(__self__, "add_to_access_token", add_to_access_token)
        if add_to_id_token is not None:
            pulumi.set(__self__, "add_to_id_token", add_to_id_token)
        if add_to_userinfo is not None:
            pulumi.set(__self__, "add_to_userinfo", add_to_userinfo)
        if claim_name is not None:
            pulumi.set(__self__, "claim_name", claim_name)
        if claim_value is not None:
            pulumi.set(__self__, "claim_value", claim_value)
        if claim_value_type is not None:
            pulumi.set(__self__, "claim_value_type", claim_value_type)
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_scope_id is not None:
            pulumi.set(__self__, "client_scope_id", client_scope_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter(name="addToAccessToken")
    def add_to_access_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the attribute should be a claim in the access token.
        """
        return pulumi.get(self, "add_to_access_token")

    @add_to_access_token.setter
    def add_to_access_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_access_token", value)

    @property
    @pulumi.getter(name="addToIdToken")
    def add_to_id_token(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the attribute should be a claim in the id token.
        """
        return pulumi.get(self, "add_to_id_token")

    @add_to_id_token.setter
    def add_to_id_token(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_id_token", value)

    @property
    @pulumi.getter(name="addToUserinfo")
    def add_to_userinfo(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates if the attribute should appear in the userinfo response body.
        """
        return pulumi.get(self, "add_to_userinfo")

    @add_to_userinfo.setter
    def add_to_userinfo(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "add_to_userinfo", value)

    @property
    @pulumi.getter(name="claimName")
    def claim_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "claim_name")

    @claim_name.setter
    def claim_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "claim_name", value)

    @property
    @pulumi.getter(name="claimValue")
    def claim_value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "claim_value")

    @claim_value.setter
    def claim_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "claim_value", value)

    @property
    @pulumi.getter(name="claimValueType")
    def claim_value_type(self) -> Optional[pulumi.Input[str]]:
        """
        Claim type used when serializing tokens.
        """
        return pulumi.get(self, "claim_value_type")

    @claim_value_type.setter
    def claim_value_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "claim_value_type", value)

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


class HardcodedClaimProtocolMapper(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 add_to_userinfo: Optional[pulumi.Input[bool]] = None,
                 claim_name: Optional[pulumi.Input[str]] = None,
                 claim_value: Optional[pulumi.Input[str]] = None,
                 claim_value_type: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## # openid.HardcodedClaimProtocolMapper

        Allows for creating and managing hardcoded claim protocol mappers within
        Keycloak.

        Hardcoded claim protocol mappers allow you to define a claim with a hardcoded
        value. Protocol mappers can be defined for a single client, or they can
        be defined for a client scope which can be shared between multiple different
        clients.

        ### Example Usage (Client)

        <!--Start PulumiCodeChooser -->
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
        hardcoded_claim_mapper = keycloak.openid.HardcodedClaimProtocolMapper("hardcoded_claim_mapper",
            realm_id=realm.id,
            client_id=openid_client.id,
            name="hardcoded-claim-mapper",
            claim_name="foo",
            claim_value="bar")
        ```
        <!--End PulumiCodeChooser -->

        ### Example Usage (Client Scope)

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("client_scope",
            realm_id=realm.id,
            name="test-client-scope")
        hardcoded_claim_mapper = keycloak.openid.HardcodedClaimProtocolMapper("hardcoded_claim_mapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id,
            name="hardcoded-claim-mapper",
            claim_name="foo",
            claim_value="bar")
        ```
        <!--End PulumiCodeChooser -->

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

        ### Import

        Protocol mappers can be imported using one of the following formats:
        - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
        - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`

        Example:

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
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: HardcodedClaimProtocolMapperArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## # openid.HardcodedClaimProtocolMapper

        Allows for creating and managing hardcoded claim protocol mappers within
        Keycloak.

        Hardcoded claim protocol mappers allow you to define a claim with a hardcoded
        value. Protocol mappers can be defined for a single client, or they can
        be defined for a client scope which can be shared between multiple different
        clients.

        ### Example Usage (Client)

        <!--Start PulumiCodeChooser -->
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
        hardcoded_claim_mapper = keycloak.openid.HardcodedClaimProtocolMapper("hardcoded_claim_mapper",
            realm_id=realm.id,
            client_id=openid_client.id,
            name="hardcoded-claim-mapper",
            claim_name="foo",
            claim_value="bar")
        ```
        <!--End PulumiCodeChooser -->

        ### Example Usage (Client Scope)

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        client_scope = keycloak.openid.ClientScope("client_scope",
            realm_id=realm.id,
            name="test-client-scope")
        hardcoded_claim_mapper = keycloak.openid.HardcodedClaimProtocolMapper("hardcoded_claim_mapper",
            realm_id=realm.id,
            client_scope_id=client_scope.id,
            name="hardcoded-claim-mapper",
            claim_name="foo",
            claim_value="bar")
        ```
        <!--End PulumiCodeChooser -->

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

        ### Import

        Protocol mappers can be imported using one of the following formats:
        - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
        - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`

        Example:

        :param str resource_name: The name of the resource.
        :param HardcodedClaimProtocolMapperArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(HardcodedClaimProtocolMapperArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_to_access_token: Optional[pulumi.Input[bool]] = None,
                 add_to_id_token: Optional[pulumi.Input[bool]] = None,
                 add_to_userinfo: Optional[pulumi.Input[bool]] = None,
                 claim_name: Optional[pulumi.Input[str]] = None,
                 claim_value: Optional[pulumi.Input[str]] = None,
                 claim_value_type: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_scope_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = HardcodedClaimProtocolMapperArgs.__new__(HardcodedClaimProtocolMapperArgs)

            __props__.__dict__["add_to_access_token"] = add_to_access_token
            __props__.__dict__["add_to_id_token"] = add_to_id_token
            __props__.__dict__["add_to_userinfo"] = add_to_userinfo
            if claim_name is None and not opts.urn:
                raise TypeError("Missing required property 'claim_name'")
            __props__.__dict__["claim_name"] = claim_name
            if claim_value is None and not opts.urn:
                raise TypeError("Missing required property 'claim_value'")
            __props__.__dict__["claim_value"] = claim_value
            __props__.__dict__["claim_value_type"] = claim_value_type
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_scope_id"] = client_scope_id
            __props__.__dict__["name"] = name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(HardcodedClaimProtocolMapper, __self__).__init__(
            'keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_to_access_token: Optional[pulumi.Input[bool]] = None,
            add_to_id_token: Optional[pulumi.Input[bool]] = None,
            add_to_userinfo: Optional[pulumi.Input[bool]] = None,
            claim_name: Optional[pulumi.Input[str]] = None,
            claim_value: Optional[pulumi.Input[str]] = None,
            claim_value_type: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_scope_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'HardcodedClaimProtocolMapper':
        """
        Get an existing HardcodedClaimProtocolMapper resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
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

        __props__ = _HardcodedClaimProtocolMapperState.__new__(_HardcodedClaimProtocolMapperState)

        __props__.__dict__["add_to_access_token"] = add_to_access_token
        __props__.__dict__["add_to_id_token"] = add_to_id_token
        __props__.__dict__["add_to_userinfo"] = add_to_userinfo
        __props__.__dict__["claim_name"] = claim_name
        __props__.__dict__["claim_value"] = claim_value
        __props__.__dict__["claim_value_type"] = claim_value_type
        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_scope_id"] = client_scope_id
        __props__.__dict__["name"] = name
        __props__.__dict__["realm_id"] = realm_id
        return HardcodedClaimProtocolMapper(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addToAccessToken")
    def add_to_access_token(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the attribute should be a claim in the access token.
        """
        return pulumi.get(self, "add_to_access_token")

    @property
    @pulumi.getter(name="addToIdToken")
    def add_to_id_token(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the attribute should be a claim in the id token.
        """
        return pulumi.get(self, "add_to_id_token")

    @property
    @pulumi.getter(name="addToUserinfo")
    def add_to_userinfo(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates if the attribute should appear in the userinfo response body.
        """
        return pulumi.get(self, "add_to_userinfo")

    @property
    @pulumi.getter(name="claimName")
    def claim_name(self) -> pulumi.Output[str]:
        return pulumi.get(self, "claim_name")

    @property
    @pulumi.getter(name="claimValue")
    def claim_value(self) -> pulumi.Output[str]:
        return pulumi.get(self, "claim_value")

    @property
    @pulumi.getter(name="claimValueType")
    def claim_value_type(self) -> pulumi.Output[Optional[str]]:
        """
        Claim type used when serializing tokens.
        """
        return pulumi.get(self, "claim_value_type")

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

