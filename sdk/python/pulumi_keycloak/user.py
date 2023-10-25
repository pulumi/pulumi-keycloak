# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Callable, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 realm_id: pulumi.Input[str],
                 username: pulumi.Input[str],
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 email_verified: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 initial_password: Optional[pulumi.Input['UserInitialPasswordArgs']] = None,
                 last_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] realm_id: The realm this user belongs to.
        :param pulumi.Input[str] username: The unique username of this user.
        :param pulumi.Input[Mapping[str, Any]] attributes: A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        :param pulumi.Input[str] email: The user's email.
        :param pulumi.Input[bool] email_verified: Whether the email address was validated or not. Default to `false`.
        :param pulumi.Input[bool] enabled: When false, this user cannot log in. Defaults to `true`.
        :param pulumi.Input[str] first_name: The user's first name.
        :param pulumi.Input['UserInitialPasswordArgs'] initial_password: When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        :param pulumi.Input[str] last_name: The user's last name.
        """
        UserArgs._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            realm_id=realm_id,
            username=username,
            attributes=attributes,
            email=email,
            email_verified=email_verified,
            enabled=enabled,
            federated_identities=federated_identities,
            first_name=first_name,
            initial_password=initial_password,
            last_name=last_name,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             realm_id: Optional[pulumi.Input[str]] = None,
             username: Optional[pulumi.Input[str]] = None,
             attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             email: Optional[pulumi.Input[str]] = None,
             email_verified: Optional[pulumi.Input[bool]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]] = None,
             first_name: Optional[pulumi.Input[str]] = None,
             initial_password: Optional[pulumi.Input['UserInitialPasswordArgs']] = None,
             last_name: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']
        if realm_id is None:
            raise TypeError("Missing 'realm_id' argument")
        if username is None:
            raise TypeError("Missing 'username' argument")
        if email_verified is None and 'emailVerified' in kwargs:
            email_verified = kwargs['emailVerified']
        if federated_identities is None and 'federatedIdentities' in kwargs:
            federated_identities = kwargs['federatedIdentities']
        if first_name is None and 'firstName' in kwargs:
            first_name = kwargs['firstName']
        if initial_password is None and 'initialPassword' in kwargs:
            initial_password = kwargs['initialPassword']
        if last_name is None and 'lastName' in kwargs:
            last_name = kwargs['lastName']

        _setter("realm_id", realm_id)
        _setter("username", username)
        if attributes is not None:
            _setter("attributes", attributes)
        if email is not None:
            _setter("email", email)
        if email_verified is not None:
            _setter("email_verified", email_verified)
        if enabled is not None:
            _setter("enabled", enabled)
        if federated_identities is not None:
            _setter("federated_identities", federated_identities)
        if first_name is not None:
            _setter("first_name", first_name)
        if initial_password is not None:
            _setter("initial_password", initial_password)
        if last_name is not None:
            _setter("last_name", last_name)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Input[str]:
        """
        The realm this user belongs to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The unique username of this user.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The user's email.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the email address was validated or not. Default to `false`.
        """
        return pulumi.get(self, "email_verified")

    @email_verified.setter
    def email_verified(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "email_verified", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When false, this user cannot log in. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="federatedIdentities")
    def federated_identities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]]:
        return pulumi.get(self, "federated_identities")

    @federated_identities.setter
    def federated_identities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]]):
        pulumi.set(self, "federated_identities", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        The user's first name.
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

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
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        The user's last name.
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 email_verified: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 initial_password: Optional[pulumi.Input['UserInitialPasswordArgs']] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[Mapping[str, Any]] attributes: A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        :param pulumi.Input[str] email: The user's email.
        :param pulumi.Input[bool] email_verified: Whether the email address was validated or not. Default to `false`.
        :param pulumi.Input[bool] enabled: When false, this user cannot log in. Defaults to `true`.
        :param pulumi.Input[str] first_name: The user's first name.
        :param pulumi.Input['UserInitialPasswordArgs'] initial_password: When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        :param pulumi.Input[str] last_name: The user's last name.
        :param pulumi.Input[str] realm_id: The realm this user belongs to.
        :param pulumi.Input[str] username: The unique username of this user.
        """
        _UserState._configure(
            lambda key, value: pulumi.set(__self__, key, value),
            attributes=attributes,
            email=email,
            email_verified=email_verified,
            enabled=enabled,
            federated_identities=federated_identities,
            first_name=first_name,
            initial_password=initial_password,
            last_name=last_name,
            realm_id=realm_id,
            username=username,
        )
    @staticmethod
    def _configure(
             _setter: Callable[[Any, Any], None],
             attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
             email: Optional[pulumi.Input[str]] = None,
             email_verified: Optional[pulumi.Input[bool]] = None,
             enabled: Optional[pulumi.Input[bool]] = None,
             federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]] = None,
             first_name: Optional[pulumi.Input[str]] = None,
             initial_password: Optional[pulumi.Input['UserInitialPasswordArgs']] = None,
             last_name: Optional[pulumi.Input[str]] = None,
             realm_id: Optional[pulumi.Input[str]] = None,
             username: Optional[pulumi.Input[str]] = None,
             opts: Optional[pulumi.ResourceOptions] = None,
             **kwargs):
        if email_verified is None and 'emailVerified' in kwargs:
            email_verified = kwargs['emailVerified']
        if federated_identities is None and 'federatedIdentities' in kwargs:
            federated_identities = kwargs['federatedIdentities']
        if first_name is None and 'firstName' in kwargs:
            first_name = kwargs['firstName']
        if initial_password is None and 'initialPassword' in kwargs:
            initial_password = kwargs['initialPassword']
        if last_name is None and 'lastName' in kwargs:
            last_name = kwargs['lastName']
        if realm_id is None and 'realmId' in kwargs:
            realm_id = kwargs['realmId']

        if attributes is not None:
            _setter("attributes", attributes)
        if email is not None:
            _setter("email", email)
        if email_verified is not None:
            _setter("email_verified", email_verified)
        if enabled is not None:
            _setter("enabled", enabled)
        if federated_identities is not None:
            _setter("federated_identities", federated_identities)
        if first_name is not None:
            _setter("first_name", first_name)
        if initial_password is not None:
            _setter("initial_password", initial_password)
        if last_name is not None:
            _setter("last_name", last_name)
        if realm_id is not None:
            _setter("realm_id", realm_id)
        if username is not None:
            _setter("username", username)

    @property
    @pulumi.getter
    def attributes(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        """
        return pulumi.get(self, "attributes")

    @attributes.setter
    def attributes(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "attributes", value)

    @property
    @pulumi.getter
    def email(self) -> Optional[pulumi.Input[str]]:
        """
        The user's email.
        """
        return pulumi.get(self, "email")

    @email.setter
    def email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "email", value)

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether the email address was validated or not. Default to `false`.
        """
        return pulumi.get(self, "email_verified")

    @email_verified.setter
    def email_verified(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "email_verified", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        When false, this user cannot log in. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="federatedIdentities")
    def federated_identities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]]:
        return pulumi.get(self, "federated_identities")

    @federated_identities.setter
    def federated_identities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['UserFederatedIdentityArgs']]]]):
        pulumi.set(self, "federated_identities", value)

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> Optional[pulumi.Input[str]]:
        """
        The user's first name.
        """
        return pulumi.get(self, "first_name")

    @first_name.setter
    def first_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "first_name", value)

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
    def last_name(self) -> Optional[pulumi.Input[str]]:
        """
        The user's last name.
        """
        return pulumi.get(self, "last_name")

    @last_name.setter
    def last_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "last_name", value)

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> Optional[pulumi.Input[str]]:
        """
        The realm this user belongs to.
        """
        return pulumi.get(self, "realm_id")

    @realm_id.setter
    def realm_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm_id", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The unique username of this user.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 email_verified: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserFederatedIdentityArgs']]]]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 initial_password: Optional[pulumi.Input[pulumi.InputType['UserInitialPasswordArgs']]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
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
        user_with_initial_password = keycloak.User("userWithInitialPassword",
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
            initial_password=keycloak.UserInitialPasswordArgs(
                value="some password",
                temporary=True,
            ))
        ```

        ## Import

        Users can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak assigns to the user upon creation. This value can be found in the GUI when editing the user. Examplebash

        ```sh
         $ pulumi import keycloak:index/user:User user my-realm/60c3f971-b1d3-4b3a-9035-d16d7540a5e4
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] attributes: A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        :param pulumi.Input[str] email: The user's email.
        :param pulumi.Input[bool] email_verified: Whether the email address was validated or not. Default to `false`.
        :param pulumi.Input[bool] enabled: When false, this user cannot log in. Defaults to `true`.
        :param pulumi.Input[str] first_name: The user's first name.
        :param pulumi.Input[pulumi.InputType['UserInitialPasswordArgs']] initial_password: When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        :param pulumi.Input[str] last_name: The user's last name.
        :param pulumi.Input[str] realm_id: The realm this user belongs to.
        :param pulumi.Input[str] username: The unique username of this user.
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
        user_with_initial_password = keycloak.User("userWithInitialPassword",
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
            initial_password=keycloak.UserInitialPasswordArgs(
                value="some password",
                temporary=True,
            ))
        ```

        ## Import

        Users can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak assigns to the user upon creation. This value can be found in the GUI when editing the user. Examplebash

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
            kwargs = kwargs or {}
            def _setter(key, value):
                kwargs[key] = value
            UserArgs._configure(_setter, **kwargs)
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 email: Optional[pulumi.Input[str]] = None,
                 email_verified: Optional[pulumi.Input[bool]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserFederatedIdentityArgs']]]]] = None,
                 first_name: Optional[pulumi.Input[str]] = None,
                 initial_password: Optional[pulumi.Input[pulumi.InputType['UserInitialPasswordArgs']]] = None,
                 last_name: Optional[pulumi.Input[str]] = None,
                 realm_id: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
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
            initial_password = _utilities.configure(initial_password, UserInitialPasswordArgs, True)
            __props__.__dict__["initial_password"] = initial_password
            __props__.__dict__["last_name"] = last_name
            if realm_id is None and not opts.urn:
                raise TypeError("Missing required property 'realm_id'")
            __props__.__dict__["realm_id"] = realm_id
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
            attributes: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            email: Optional[pulumi.Input[str]] = None,
            email_verified: Optional[pulumi.Input[bool]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            federated_identities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['UserFederatedIdentityArgs']]]]] = None,
            first_name: Optional[pulumi.Input[str]] = None,
            initial_password: Optional[pulumi.Input[pulumi.InputType['UserInitialPasswordArgs']]] = None,
            last_name: Optional[pulumi.Input[str]] = None,
            realm_id: Optional[pulumi.Input[str]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] attributes: A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        :param pulumi.Input[str] email: The user's email.
        :param pulumi.Input[bool] email_verified: Whether the email address was validated or not. Default to `false`.
        :param pulumi.Input[bool] enabled: When false, this user cannot log in. Defaults to `true`.
        :param pulumi.Input[str] first_name: The user's first name.
        :param pulumi.Input[pulumi.InputType['UserInitialPasswordArgs']] initial_password: When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        :param pulumi.Input[str] last_name: The user's last name.
        :param pulumi.Input[str] realm_id: The realm this user belongs to.
        :param pulumi.Input[str] username: The unique username of this user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["attributes"] = attributes
        __props__.__dict__["email"] = email
        __props__.__dict__["email_verified"] = email_verified
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["federated_identities"] = federated_identities
        __props__.__dict__["first_name"] = first_name
        __props__.__dict__["initial_password"] = initial_password
        __props__.__dict__["last_name"] = last_name
        __props__.__dict__["realm_id"] = realm_id
        __props__.__dict__["username"] = username
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def attributes(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        A map representing attributes for the user. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter
    def email(self) -> pulumi.Output[Optional[str]]:
        """
        The user's email.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="emailVerified")
    def email_verified(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether the email address was validated or not. Default to `false`.
        """
        return pulumi.get(self, "email_verified")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        When false, this user cannot log in. Defaults to `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="federatedIdentities")
    def federated_identities(self) -> pulumi.Output[Optional[Sequence['outputs.UserFederatedIdentity']]]:
        return pulumi.get(self, "federated_identities")

    @property
    @pulumi.getter(name="firstName")
    def first_name(self) -> pulumi.Output[Optional[str]]:
        """
        The user's first name.
        """
        return pulumi.get(self, "first_name")

    @property
    @pulumi.getter(name="initialPassword")
    def initial_password(self) -> pulumi.Output[Optional['outputs.UserInitialPassword']]:
        """
        When given, the user's initial password will be set. This attribute is only respected during initial user creation.
        """
        return pulumi.get(self, "initial_password")

    @property
    @pulumi.getter(name="lastName")
    def last_name(self) -> pulumi.Output[Optional[str]]:
        """
        The user's last name.
        """
        return pulumi.get(self, "last_name")

    @property
    @pulumi.getter(name="realmId")
    def realm_id(self) -> pulumi.Output[str]:
        """
        The realm this user belongs to.
        """
        return pulumi.get(self, "realm_id")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The unique username of this user.
        """
        return pulumi.get(self, "username")

