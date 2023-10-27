# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 client_id: pulumi.Input[str],
                 url: pulumi.Input[str],
                 additional_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 base_path: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 client_timeout: Optional[pulumi.Input[int]] = None,
                 initial_login: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 red_hat_sso: Optional[pulumi.Input[bool]] = None,
                 root_ca_certificate: Optional[pulumi.Input[str]] = None,
                 tls_insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] url: The base URL of the Keycloak instance, before `/auth`
        :param pulumi.Input[int] client_timeout: Timeout (in seconds) of the Keycloak client
        :param pulumi.Input[bool] initial_login: Whether or not to login to Keycloak instance on provider initialization
        :param pulumi.Input[bool] red_hat_sso: When true, the provider will treat the Keycloak instance as a Red Hat SSO server, specifically when parsing the version
               returned from the /serverinfo API endpoint.
        :param pulumi.Input[str] root_ca_certificate: Allows x509 calls using an unknown CA certificate (for development purposes)
        :param pulumi.Input[bool] tls_insecure_skip_verify: Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
               should be avoided.
        """
        pulumi.set(__self__, "client_id", client_id)
        pulumi.set(__self__, "url", url)
        if additional_headers is not None:
            pulumi.set(__self__, "additional_headers", additional_headers)
        if base_path is not None:
            pulumi.set(__self__, "base_path", base_path)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if client_timeout is None:
            client_timeout = (_utilities.get_env_int('KEYCLOAK_CLIENT_TIMEOUT') or 5)
        if client_timeout is not None:
            pulumi.set(__self__, "client_timeout", client_timeout)
        if initial_login is not None:
            pulumi.set(__self__, "initial_login", initial_login)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if realm is not None:
            pulumi.set(__self__, "realm", realm)
        if red_hat_sso is not None:
            pulumi.set(__self__, "red_hat_sso", red_hat_sso)
        if root_ca_certificate is not None:
            pulumi.set(__self__, "root_ca_certificate", root_ca_certificate)
        if tls_insecure_skip_verify is not None:
            pulumi.set(__self__, "tls_insecure_skip_verify", tls_insecure_skip_verify)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter
    def url(self) -> pulumi.Input[str]:
        """
        The base URL of the Keycloak instance, before `/auth`
        """
        return pulumi.get(self, "url")

    @url.setter
    def url(self, value: pulumi.Input[str]):
        pulumi.set(self, "url", value)

    @property
    @pulumi.getter(name="additionalHeaders")
    def additional_headers(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "additional_headers")

    @additional_headers.setter
    def additional_headers(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "additional_headers", value)

    @property
    @pulumi.getter(name="basePath")
    def base_path(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "base_path")

    @base_path.setter
    def base_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_path", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter(name="clientTimeout")
    def client_timeout(self) -> Optional[pulumi.Input[int]]:
        """
        Timeout (in seconds) of the Keycloak client
        """
        return pulumi.get(self, "client_timeout")

    @client_timeout.setter
    def client_timeout(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "client_timeout", value)

    @property
    @pulumi.getter(name="initialLogin")
    def initial_login(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether or not to login to Keycloak instance on provider initialization
        """
        return pulumi.get(self, "initial_login")

    @initial_login.setter
    def initial_login(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "initial_login", value)

    @property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "password", value)

    @property
    @pulumi.getter
    def realm(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "realm")

    @realm.setter
    def realm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "realm", value)

    @property
    @pulumi.getter(name="redHatSso")
    def red_hat_sso(self) -> Optional[pulumi.Input[bool]]:
        """
        When true, the provider will treat the Keycloak instance as a Red Hat SSO server, specifically when parsing the version
        returned from the /serverinfo API endpoint.
        """
        return pulumi.get(self, "red_hat_sso")

    @red_hat_sso.setter
    def red_hat_sso(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "red_hat_sso", value)

    @property
    @pulumi.getter(name="rootCaCertificate")
    def root_ca_certificate(self) -> Optional[pulumi.Input[str]]:
        """
        Allows x509 calls using an unknown CA certificate (for development purposes)
        """
        return pulumi.get(self, "root_ca_certificate")

    @root_ca_certificate.setter
    def root_ca_certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "root_ca_certificate", value)

    @property
    @pulumi.getter(name="tlsInsecureSkipVerify")
    def tls_insecure_skip_verify(self) -> Optional[pulumi.Input[bool]]:
        """
        Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
        should be avoided.
        """
        return pulumi.get(self, "tls_insecure_skip_verify")

    @tls_insecure_skip_verify.setter
    def tls_insecure_skip_verify(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "tls_insecure_skip_verify", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 base_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 client_timeout: Optional[pulumi.Input[int]] = None,
                 initial_login: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 red_hat_sso: Optional[pulumi.Input[bool]] = None,
                 root_ca_certificate: Optional[pulumi.Input[str]] = None,
                 tls_insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        The provider type for the keycloak package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] client_timeout: Timeout (in seconds) of the Keycloak client
        :param pulumi.Input[bool] initial_login: Whether or not to login to Keycloak instance on provider initialization
        :param pulumi.Input[bool] red_hat_sso: When true, the provider will treat the Keycloak instance as a Red Hat SSO server, specifically when parsing the version
               returned from the /serverinfo API endpoint.
        :param pulumi.Input[str] root_ca_certificate: Allows x509 calls using an unknown CA certificate (for development purposes)
        :param pulumi.Input[bool] tls_insecure_skip_verify: Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
               should be avoided.
        :param pulumi.Input[str] url: The base URL of the Keycloak instance, before `/auth`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the keycloak package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 additional_headers: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 base_path: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 client_timeout: Optional[pulumi.Input[int]] = None,
                 initial_login: Optional[pulumi.Input[bool]] = None,
                 password: Optional[pulumi.Input[str]] = None,
                 realm: Optional[pulumi.Input[str]] = None,
                 red_hat_sso: Optional[pulumi.Input[bool]] = None,
                 root_ca_certificate: Optional[pulumi.Input[str]] = None,
                 tls_insecure_skip_verify: Optional[pulumi.Input[bool]] = None,
                 url: Optional[pulumi.Input[str]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            __props__.__dict__["additional_headers"] = pulumi.Output.from_input(additional_headers).apply(pulumi.runtime.to_json) if additional_headers is not None else None
            __props__.__dict__["base_path"] = base_path
            if client_id is None and not opts.urn:
                raise TypeError("Missing required property 'client_id'")
            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_secret"] = client_secret
            if client_timeout is None:
                client_timeout = (_utilities.get_env_int('KEYCLOAK_CLIENT_TIMEOUT') or 5)
            __props__.__dict__["client_timeout"] = pulumi.Output.from_input(client_timeout).apply(pulumi.runtime.to_json) if client_timeout is not None else None
            __props__.__dict__["initial_login"] = pulumi.Output.from_input(initial_login).apply(pulumi.runtime.to_json) if initial_login is not None else None
            __props__.__dict__["password"] = password
            __props__.__dict__["realm"] = realm
            __props__.__dict__["red_hat_sso"] = pulumi.Output.from_input(red_hat_sso).apply(pulumi.runtime.to_json) if red_hat_sso is not None else None
            __props__.__dict__["root_ca_certificate"] = root_ca_certificate
            __props__.__dict__["tls_insecure_skip_verify"] = pulumi.Output.from_input(tls_insecure_skip_verify).apply(pulumi.runtime.to_json) if tls_insecure_skip_verify is not None else None
            if url is None and not opts.urn:
                raise TypeError("Missing required property 'url'")
            __props__.__dict__["url"] = url
            __props__.__dict__["username"] = username
        super(Provider, __self__).__init__(
            'keycloak',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="basePath")
    def base_path(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "base_path")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def password(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "password")

    @property
    @pulumi.getter
    def realm(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "realm")

    @property
    @pulumi.getter(name="rootCaCertificate")
    def root_ca_certificate(self) -> pulumi.Output[Optional[str]]:
        """
        Allows x509 calls using an unknown CA certificate (for development purposes)
        """
        return pulumi.get(self, "root_ca_certificate")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The base URL of the Keycloak instance, before `/auth`
        """
        return pulumi.get(self, "url")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "username")

