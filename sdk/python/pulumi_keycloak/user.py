# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[builtins.str],
                 username: pulumi.Input[builtins.str],
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 email_verified: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]] = None,
                 first_name: Optional[pulumi.Input[builtins.str]] = None,
                 import_: Optional[pulumi.Input[builtins.bool]] = None,
                 initial_password: Optional[pulumi.Input['UserInitialPasswordArgs']] = None,
                 last_name: Optional[pulumi.Input[builtins.str]] = None,
                 required_actions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[builtins.str] realm_id: The realm this user belongs to.
        :param pulumi.Input[builtins.str] username: The unique username of this user.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] attributes: A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        :param pulumi.Input[builtins.str] email: The user's email.
        :param pulumi.Input[builtins.bool] email_verified: Whether the email address was validated or not. Default to `false`.
        :param pulumi.Input[builtins.bool] enabled: When false, this user cannot log in. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]] federated_identities: When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
        :param pulumi.Input[builtins.str] first_name: The user's first name.
        :param pulumi.Input[builtins.bool] import_: When `true`, the user with the specified `username` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with users that Keycloak creates automatically during realm creation, such as `admin`. Note, that the user will not be removed during destruction if `import` is `true`.
        :param pulumi.Input['UserInitialPasswordArgs'] initial_password: When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        :param pulumi.Input[builtins.str] last_name: The user's last name.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] required_actions: A list of required user actions.
        """
        pulumi.set(__self__, "realm_id", realm_id)
        pulumi.set(__self__, "username", username)
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if email_verified is not None:
            pulumi.set(__self__, "email_verified", email_verified)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if federated_identities is not None:
            pulumi.set(__self__, "federated_identities", federated_identities)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if import_ is not None:
            pulumi.set(__self__, "import_", import_)
        if initial_password is not None:
            pulumi.set(__self__, "initial_password", initial_password)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if required_actions is not None:
            pulumi.set(__self__, "required_actions", required_actions)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[builtins.str]:
        """
        The realm this user belongs to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[builtins.str]:
        """
        The unique username of this user.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The user's email.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the email address was validated or not. Default to `false`.
        """
        return pulumi.get(self, "email_verified")

    @email_verified.setter
    def email_verified(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "email_verified", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When false, this user cannot log in. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="federatedIdentities")
    def federated_identities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]]:
        """
        When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
        """
        return pulumi.get(self, "federated_identities")

    @federated_identities.setter
    def federated_identities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]]):
        pulumi.set(self, "federated_identities", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The user's first name.
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="import")
    def import_(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When `true`, the user with the specified `username` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with users that Keycloak creates automatically during realm creation, such as `admin`. Note, that the user will not be removed during destruction if `import` is `true`.
        """
        return pulumi.get(self, "import_")

    @import_.setter
    def import_(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "import_", value)

    @property
    @pulumi.getter(name="initialPassword")
    def initial_password(self) -> Optional[pulumi.Input['UserInitialPasswordArgs']]:
        """
        When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        """
        return pulumi.get(self, "initial_password")

    @initial_password.setter
    def initial_password(self, value: Optional[pulumi.Input['UserInitialPasswordArgs']]):
        pulumi.set(self, "initial_password", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The user's last name.
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter(name="requiredActions")
    def required_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of required user actions.
        """
        return pulumi.get(self, "required_actions")

    @required_actions.setter
    def required_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "required_actions", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 email_verified: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]] = None,
                 first_name: Optional[pulumi.Input[builtins.str]] = None,
                 import_: Optional[pulumi.Input[builtins.bool]] = None,
                 initial_password: Optional[pulumi.Input['UserInitialPasswordArgs']] = None,
                 last_name: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 required_actions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] attributes: A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        :param pulumi.Input[builtins.str] email: The user's email.
        :param pulumi.Input[builtins.bool] email_verified: Whether the email address was validated or not. Default to `false`.
        :param pulumi.Input[builtins.bool] enabled: When false, this user cannot log in. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]] federated_identities: When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
        :param pulumi.Input[builtins.str] first_name: The user's first name.
        :param pulumi.Input[builtins.bool] import_: When `true`, the user with the specified `username` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with users that Keycloak creates automatically during realm creation, such as `admin`. Note, that the user will not be removed during destruction if `import` is `true`.
        :param pulumi.Input['UserInitialPasswordArgs'] initial_password: When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        :param pulumi.Input[builtins.str] last_name: The user's last name.
        :param pulumi.Input[builtins.str] realm_id: The realm this user belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] required_actions: A list of required user actions.
        :param pulumi.Input[builtins.str] username: The unique username of this user.
        """
        if attributes is not None:
            pulumi.set(__self__, "attributes", attributes)
        if email is not None:
            pulumi.set(__self__, "email", email)
        if email_verified is not None:
            pulumi.set(__self__, "email_verified", email_verified)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if federated_identities is not None:
            pulumi.set(__self__, "federated_identities", federated_identities)
        if first_name is not None:
            pulumi.set(__self__, "first_name", first_name)
        if import_ is not None:
            pulumi.set(__self__, "import_", import_)
        if initial_password is not None:
            pulumi.set(__self__, "initial_password", initial_password)
        if last_name is not None:
            pulumi.set(__self__, "last_name", last_name)
        if realm_id is not None:
            pulumi.set(__self__, "realm_id", realm_id)
        if required_actions is not None:
            pulumi.set(__self__, "required_actions", required_actions)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The user's email.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the email address was validated or not. Default to `false`.
        """
        return pulumi.get(self, "email_verified")

    @email_verified.setter
    def email_verified(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "email_verified", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When false, this user cannot log in. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="federatedIdentities")
    def federated_identities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]]:
        """
        When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
        """
        return pulumi.get(self, "federated_identities")

    @federated_identities.setter
    def federated_identities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]]):
        pulumi.set(self, "federated_identities", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The user's first name.
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "first_name", value)

    @property
    @pulumi.getter(name="import")
    def import_(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        When `true`, the user with the specified `username` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with users that Keycloak creates automatically during realm creation, such as `admin`. Note, that the user will not be removed during destruction if `import` is `true`.
        """
        return pulumi.get(self, "import_")

    @import_.setter
    def import_(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "import_", value)

    @property
    @pulumi.getter(name="initialPassword")
    def initial_password(self) -> Optional[pulumi.Input['UserInitialPasswordArgs']]:
        """
        When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        """
        return pulumi.get(self, "initial_password")

    @initial_password.setter
    def initial_password(self, value: Optional[pulumi.Input['UserInitialPasswordArgs']]):
        pulumi.set(self, "initial_password", value)

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The user's last name.
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The realm this user belongs to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter(name="requiredActions")
    def required_actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of required user actions.
        """
        return pulumi.get(self, "required_actions")

    @required_actions.setter
    def required_actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "required_actions", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The unique username of this user.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "username", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 email_verified: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserFederatedIdentityArgs', 'UserFederatedIdentityArgsDict']]]]] = None,
                 first_name: Optional[pulumi.Input[builtins.str]] = None,
                 import_: Optional[pulumi.Input[builtins.bool]] = None,
                 initial_password: Optional[pulumi.Input[Union['UserInitialPasswordArgs', 'UserInitialPasswordArgsDict']]] = None,
                 last_name: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 required_actions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Allows for creating and managing Users within Keycloak.

        This resource was created primarily to enable the acceptance tests for the `Group` resource. Creating users within
        Keycloak is not recommended. Instead, users should be federated from external sources by configuring user federation providers
        or identity providers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        user = keycloak.User("user",
            realm_id=realm.id,
            username="bob",
            enabled=True,
            email="bob@domain.com",
            first_name="Bob",
            last_name="Bobson")
        user_with_initial_password = keycloak.User("user_with_initial_password",
            realm_id=realm.id,
            username="alice",
            enabled=True,
            email="alice@domain.com",
            first_name="Alice",
            last_name="Aliceberg",
            attributes={
                "foo": "bar",
                "multivalue": "value1##value2",
            },
            initial_password={
                "value": "some password",
                "temporary": True,
            })
        ```

        ## Import

        Users can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak

        assigns to the user upon creation. This value can be found in the GUI when editing the user.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:index/user:User user my-realm/60c3f971-b1d3-4b3a-9035-d16d7540a5e4
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] attributes: A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        :param pulumi.Input[builtins.str] email: The user's email.
        :param pulumi.Input[builtins.bool] email_verified: Whether the email address was validated or not. Default to `false`.
        :param pulumi.Input[builtins.bool] enabled: When false, this user cannot log in. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['UserFederatedIdentityArgs', 'UserFederatedIdentityArgsDict']]]] federated_identities: When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
        :param pulumi.Input[builtins.str] first_name: The user's first name.
        :param pulumi.Input[builtins.bool] import_: When `true`, the user with the specified `username` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with users that Keycloak creates automatically during realm creation, such as `admin`. Note, that the user will not be removed during destruction if `import` is `true`.
        :param pulumi.Input[Union['UserInitialPasswordArgs', 'UserInitialPasswordArgsDict']] initial_password: When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        :param pulumi.Input[builtins.str] last_name: The user's last name.
        :param pulumi.Input[builtins.str] realm_id: The realm this user belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] required_actions: A list of required user actions.
        :param pulumi.Input[builtins.str] username: The unique username of this user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows for creating and managing Users within Keycloak.

        This resource was created primarily to enable the acceptance tests for the `Group` resource. Creating users within
        Keycloak is not recommended. Instead, users should be federated from external sources by configuring user federation providers
        or identity providers.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_keycloak as keycloak

        realm = keycloak.Realm("realm",
            realm="my-realm",
            enabled=True)
        user = keycloak.User("user",
            realm_id=realm.id,
            username="bob",
            enabled=True,
            email="bob@domain.com",
            first_name="Bob",
            last_name="Bobson")
        user_with_initial_password = keycloak.User("user_with_initial_password",
            realm_id=realm.id,
            username="alice",
            enabled=True,
            email="alice@domain.com",
            first_name="Alice",
            last_name="Aliceberg",
            attributes={
                "foo": "bar",
                "multivalue": "value1##value2",
            },
            initial_password={
                "value": "some password",
                "temporary": True,
            })
        ```

        ## Import

        Users can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak

        assigns to the user upon creation. This value can be found in the GUI when editing the user.

        Example:

        bash

        ```sh
        $ pulumi import keycloak:index/user:User user my-realm/60c3f971-b1d3-4b3a-9035-d16d7540a5e4
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 email: Optional[pulumi.Input[builtins.str]] = None,
                 email_verified: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserFederatedIdentityArgs', 'UserFederatedIdentityArgsDict']]]]] = None,
                 first_name: Optional[pulumi.Input[builtins.str]] = None,
                 import_: Optional[pulumi.Input[builtins.bool]] = None,
                 initial_password: Optional[pulumi.Input[Union['UserInitialPasswordArgs', 'UserInitialPasswordArgsDict']]] = None,
                 last_name: Optional[pulumi.Input[builtins.str]] = None,
                 realm_id: Optional[pulumi.Input[builtins.str]] = None,
                 required_actions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 username: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["attributes"] = attributes
            __props__.__dict__["email"] = email
            __props__.__dict__["email_verified"] = email_verified
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["federated_identities"] = federated_identities
            __props__.__dict__["first_name"] = first_name
            __props__.__dict__["import_"] = import_
            __props__.__dict__["initial_password"] = initial_password
            __props__.__dict__["last_name"] = last_name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
            __props__.__dict__["required_actions"] = required_actions
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        super(User, __self__).__init__(
            'keycloak:index/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            attributes: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            email: Optional[pulumi.Input[builtins.str]] = None,
            email_verified: Optional[pulumi.Input[builtins.bool]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input[Union['UserFederatedIdentityArgs', 'UserFederatedIdentityArgsDict']]]]] = None,
            first_name: Optional[pulumi.Input[builtins.str]] = None,
            import_: Optional[pulumi.Input[builtins.bool]] = None,
            initial_password: Optional[pulumi.Input[Union['UserInitialPasswordArgs', 'UserInitialPasswordArgsDict']]] = None,
            last_name: Optional[pulumi.Input[builtins.str]] = None,
            realm_id: Optional[pulumi.Input[builtins.str]] = None,
            required_actions: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            username: Optional[pulumi.Input[builtins.str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] attributes: A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        :param pulumi.Input[builtins.str] email: The user's email.
        :param pulumi.Input[builtins.bool] email_verified: Whether the email address was validated or not. Default to `false`.
        :param pulumi.Input[builtins.bool] enabled: When false, this user cannot log in. Defaults to `true`.
        :param pulumi.Input[Sequence[pulumi.Input[Union['UserFederatedIdentityArgs', 'UserFederatedIdentityArgsDict']]]] federated_identities: When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
        :param pulumi.Input[builtins.str] first_name: The user's first name.
        :param pulumi.Input[builtins.bool] import_: When `true`, the user with the specified `username` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with users that Keycloak creates automatically during realm creation, such as `admin`. Note, that the user will not be removed during destruction if `import` is `true`.
        :param pulumi.Input[Union['UserInitialPasswordArgs', 'UserInitialPasswordArgsDict']] initial_password: When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        :param pulumi.Input[builtins.str] last_name: The user's last name.
        :param pulumi.Input[builtins.str] realm_id: The realm this user belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] required_actions: A list of required user actions.
        :param pulumi.Input[builtins.str] username: The unique username of this user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["attributes"] = attributes
        __props__.__dict__["email"] = email
        __props__.__dict__["email_verified"] = email_verified
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["federated_identities"] = federated_identities
        __props__.__dict__["first_name"] = first_name
        __props__.__dict__["import_"] = import_
        __props__.__dict__["initial_password"] = initial_password
        __props__.__dict__["last_name"] = last_name
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["required_actions"] = required_actions
        __props__.__dict__["username"] = username
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional[Mapping[str, builtins.str]]]:
        """
        A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The user's email.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Whether the email address was validated or not. Default to `false`.
        """
        return pulumi.get(self, "email_verified")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        When false, this user cannot log in. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="federatedIdentities")
    def federated_identities(self) -> pulumi.Output[Optional[Sequence['outputs.UserFederatedIdentity']]]:
        """
        When specified, the user will be linked to a federated identity provider. Refer to the federated user example for more details.
        """
        return pulumi.get(self, "federated_identities")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The user's first name.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter(name="import")
    def import_(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        When `true`, the user with the specified `username` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with users that Keycloak creates automatically during realm creation, such as `admin`. Note, that the user will not be removed during destruction if `import` is `true`.
        """
        return pulumi.get(self, "import_")

    @property
    @pulumi.getter(name="initialPassword")
    def initial_password(self) -> pulumi.Output[Optional['outputs.UserInitialPassword']]:
        """
        When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        """
        return pulumi.get(self, "initial_password")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The user's last name.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[builtins.str]:
        """
        The realm this user belongs to.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter(name="requiredActions")
    def required_actions(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        A list of required user actions.
        """
        return pulumi.get(self, "required_actions")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[builtins.str]:
        """
        The unique username of this user.
        """
        return pulumi.get(self, "username")

