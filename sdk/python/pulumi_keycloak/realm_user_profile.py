# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RealmUserProfileArgs', 'RealmUserProfile']

@pulumi.input_type
class RealmUserProfileArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileAttributeArgs']]]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileGroupArgs']]]] = None):
        """
        The set of arguments for constructing a RealmUserProfile resource.
        :param pulumi.Input[str] realm_id: The ID of the realm the user profile applies to.
        :param pulumi.Input[Sequence[pulumi.Input['RealmUserProfileAttributeArgs']]] attributes: An ordered list of attributes.
        :param pulumi.Input[Sequence[pulumi.Input['RealmUserProfileGroupArgs']]] groups: A list of groups.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The ID of the realm the user profile applies to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileAttributeArgs']]]]:
        """
        An ordered list of attributes.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileAttributeArgs']]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileGroupArgs']]]]:
        """
        A list of groups.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileGroupArgs']]]]):
        pulumi.set(self, "groups", value)


@pulumi.input_type
class _RealmUserProfileState:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileAttributeArgs']]]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileGroupArgs']]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering RealmUserProfile resources.
        :param pulumi.Input[Sequence[pulumi.Input['RealmUserProfileAttributeArgs']]] attributes: An ordered list of attributes.
        :param pulumi.Input[Sequence[pulumi.Input['RealmUserProfileGroupArgs']]] groups: A list of groups.
        :param pulumi.Input[str] realm_id: The ID of the realm the user profile applies to.
        """
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileAttributeArgs']]]]:
        """
        An ordered list of attributes.
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileAttributeArgs']]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileGroupArgs']]]]:
        """
        A list of groups.
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['RealmUserProfileGroupArgs']]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the realm the user profile applies to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)


class RealmUserProfile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileAttributeArgs']]]]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileGroupArgs']]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Allows for managing Realm User Profiles within Keycloak.

        A user profile defines a schema for representing user attributes and how they are managed within a realm.
        This is a preview feature, hence not fully supported and disabled by default.
        To enable it, start the server with one of the following flags:
        - WildFly distribution: `-Dkeycloak.profile.feature.declarative_user_profile=enabled`
        - Quarkus distribution: `--features=preview` or `--features=declarative-user-profile`

        The realm linked to the `RealmUserProfile` resource must have the user profile feature enabled.
        It can be done via the administration UI, or by setting the `userProfileEnabled` realm attribute to `true`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import json
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            attributes={
                "userProfileEnabled": True,
            })
        userprofile = keycloak.RealmUserProfile("userprofile",
            realm_id=keycloak_realm["my_realm"]["id"],
            attributes=[
                keycloak.RealmUserProfileAttributeArgs(
                    name="field1",
                    display_name="Field 1",
                    group="group1",
                    enabled_when_scopes=["offline_access"],
                    required_for_roles=["user"],
                    required_for_scopes=["offline_access"],
                    permissions=keycloak.RealmUserProfileAttributePermissionsArgs(
                        views=[
                            "admin",
                            "user",
                        ],
                        edits=[
                            "admin",
                            "user",
                        ],
                    ),
                    validators=[
                        keycloak.RealmUserProfileAttributeValidatorArgs(
                            name="person-name-prohibited-characters",
                        ),
                        keycloak.RealmUserProfileAttributeValidatorArgs(
                            name="pattern",
                            config={
                                "pattern": "^[a-z]+$",
                                "error-message": "Nope",
                            },
                        ),
                    ],
                    annotations={
                        "foo": "bar",
                    },
                ),
                keycloak.RealmUserProfileAttributeArgs(
                    name="field2",
                    validators=[keycloak.RealmUserProfileAttributeValidatorArgs(
                        name="options",
                        config={
                            "options": json.dumps(["opt1"]),
                        },
                    )],
                    annotations={
                        "foo": json.dumps({
                            "key": "val",
                        }),
                    },
                ),
            ],
            groups=[
                keycloak.RealmUserProfileGroupArgs(
                    name="group1",
                    display_header="Group 1",
                    display_description="A first group",
                    annotations={
                        "foo": "bar",
                        "foo2": json.dumps({
                            "key": "val",
                        }),
                    },
                ),
                keycloak.RealmUserProfileGroupArgs(
                    name="group2",
                ),
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        This resource currently does not support importing.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileAttributeArgs']]]] attributes: An ordered list of attributes.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileGroupArgs']]]] groups: A list of groups.
        :param pulumi.Input[str] realm_id: The ID of the realm the user profile applies to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RealmUserProfileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for managing Realm User Profiles within Keycloak.

        A user profile defines a schema for representing user attributes and how they are managed within a realm.
        This is a preview feature, hence not fully supported and disabled by default.
        To enable it, start the server with one of the following flags:
        - WildFly distribution: `-Dkeycloak.profile.feature.declarative_user_profile=enabled`
        - Quarkus distribution: `--features=preview` or `--features=declarative-user-profile`

        The realm linked to the `RealmUserProfile` resource must have the user profile feature enabled.
        It can be done via the administration UI, or by setting the `userProfileEnabled` realm attribute to `true`.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import json
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            attributes={
                "userProfileEnabled": True,
            })
        userprofile = keycloak.RealmUserProfile("userprofile",
            realm_id=keycloak_realm["my_realm"]["id"],
            attributes=[
                keycloak.RealmUserProfileAttributeArgs(
                    name="field1",
                    display_name="Field 1",
                    group="group1",
                    enabled_when_scopes=["offline_access"],
                    required_for_roles=["user"],
                    required_for_scopes=["offline_access"],
                    permissions=keycloak.RealmUserProfileAttributePermissionsArgs(
                        views=[
                            "admin",
                            "user",
                        ],
                        edits=[
                            "admin",
                            "user",
                        ],
                    ),
                    validators=[
                        keycloak.RealmUserProfileAttributeValidatorArgs(
                            name="person-name-prohibited-characters",
                        ),
                        keycloak.RealmUserProfileAttributeValidatorArgs(
                            name="pattern",
                            config={
                                "pattern": "^[a-z]+$",
                                "error-message": "Nope",
                            },
                        ),
                    ],
                    annotations={
                        "foo": "bar",
                    },
                ),
                keycloak.RealmUserProfileAttributeArgs(
                    name="field2",
                    validators=[keycloak.RealmUserProfileAttributeValidatorArgs(
                        name="options",
                        config={
                            "options": json.dumps(["opt1"]),
                        },
                    )],
                    annotations={
                        "foo": json.dumps({
                            "key": "val",
                        }),
                    },
                ),
            ],
            groups=[
                keycloak.RealmUserProfileGroupArgs(
                    name="group1",
                    display_header="Group 1",
                    display_description="A first group",
                    annotations={
                        "foo": "bar",
                        "foo2": json.dumps({
                            "key": "val",
                        }),
                    },
                ),
                keycloak.RealmUserProfileGroupArgs(
                    name="group2",
                ),
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        This resource currently does not support importing.

        :param str resource_name: The name of the resource.
        :param RealmUserProfileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RealmUserProfileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileAttributeArgs']]]]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileGroupArgs']]]]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RealmUserProfileArgs.__new__(RealmUserProfileArgs)

            __props__.__dict__["attributes"] = attributes
            __props__.__dict__["groups"] = groups
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
        super(RealmUserProfile, __self__).__init__(
            'keycloak:index/realmUserProfile:RealmUserProfile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attributes: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileAttributeArgs']]]]] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileGroupArgs']]]]] = None,
            realm_id: Optional[pulumi.Input[str]] = None) -> 'RealmUserProfile':
        """
        Get an existing RealmUserProfile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileAttributeArgs']]]] attributes: An ordered list of attributes.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['RealmUserProfileGroupArgs']]]] groups: A list of groups.
        :param pulumi.Input[str] realm_id: The ID of the realm the user profile applies to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RealmUserProfileState.__new__(_RealmUserProfileState)

        __props__.__dict__["attributes"] = attributes
        __props__.__dict__["groups"] = groups
        __props__.__dict__["realm_id"] = realm_id
        return RealmUserProfile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional[Sequence['outputs.RealmUserProfileAttribute']]]:
        """
        An ordered list of attributes.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Optional[Sequence['outputs.RealmUserProfileGroup']]]:
        """
        A list of groups.
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The ID of the realm the user profile applies to.
        """
        return pulumi.get(self, "realm_id")

