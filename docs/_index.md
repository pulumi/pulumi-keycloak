---
title: Keycloak Provider
meta_desc: Provides an overview on how to configure the Pulumi Keycloak provider.
layout: package
---
## Installation

The keycloak provider is available as a package in all Pulumi languages:

* JavaScript/TypeScript: [`@pulumi/keycloak`](https://www.npmjs.com/package/@pulumi/keycloak)
* Python: [`pulumi-keycloak`](https://pypi.org/project/pulumi-keycloak/)
* Go: [`github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak`](https://github.com/pulumi/pulumi-keycloak)
* .NET: [`Pulumi.Keycloak`](https://www.nuget.org/packages/Pulumi.Keycloak)
* Java: [`com.pulumi/keycloak`](https://central.sonatype.com/artifact/com.pulumi/keycloak)
## Overview
The Keycloak provider can be used to interact with [Keycloak](https://www.keycloak.org/).
## A note for users of the legacy Wildfly distribution

Recently, Keycloak has been updated to use Quarkus over the legacy Wildfly distribution. The only significant change here
that affects this Pulumi provider is the removal of `/auth` from the default context path for the Keycloak API.

If you are using the legacy Wildfly distribution of Keycloak, you will need to set the `basePath` provider argument to
`/auth`. This can also be done by using the `KEYCLOAK_BASE_PATH` environment variable.
## Keycloak Setup

This Pulumi provider can be configured to use the [client credentials](https://www.oauth.com/oauth2-servers/access-tokens/client-credentials/)
or [password](https://www.oauth.com/oauth2-servers/access-tokens/password-grant/) grant types. If you aren't
sure which to use, the client credentials grant is recommended, as it was designed for machine to machine authentication.
### Client Credentials Grant Setup (recommended)

1. Create a new client using the `openid-connect` protocol. This client can be created in the `master` realm if you would
   like to manage your entire Keycloak instance, or in any other realm if you only want to manage that realm.
2. Update the client you just created:
   1. Set `Access Type` to `confidential`.
   2. Set `Standard Flow Enabled` to `OFF`.
   3. Set `Direct Access Grants Enabled` to `OFF`
   4. Set `Service Accounts Enabled` to `ON`.
3. Grant required roles for managing Keycloak via the `Service Account Roles` tab in the client you created in step 1, see Assigning Roles section below.
### Password Grant Setup

These steps will assume that you are using the `admin-cli` client, which is already correctly configured for this type
of authentication. Do not follow these steps if you have already followed the steps for the client credentials grant.

1. Create or identify the user whose credentials will be used for authentication.
2. Edit this user in the "Users" section of the management console and assign roles using the "Role Mappings" tab.
### Assigning Roles

There are many ways that roles can be assigned to manage Keycloak. Here are a couple of common scenarios accompanied
by suggested roles to assign. This is not an exhaustive list, and there is often more than one way to assign a particular set
of permissions.

- Managing the entire Keycloak instance: Assign the `admin` role to a user or service account within the `master` realm.
- Managing the entire `foo` realm: Assign the `realm-admin` client role from the `realm-management` client to a user or service
  account within the `foo` realm.
- Managing clients for all realms within the entire Keycloak instance: Assign the `create-client` client role from each of
  the realm clients to a user or service account within the `master` realm. For example, given a Keycloak instance with realms
  `master`, `foo`, and `bar`, assign the `create-client` client role from the clients `master-realm`, `foo-realm`, and `bar-realm`.
## Example Usage (client credentials grant)

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    keycloak:clientId:
        value: pulumi
    keycloak:clientSecret:
        value: 884e0f95-0f42-4a63-9b1f-94274655669e
    keycloak:url:
        value: http://localhost:8080

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    keycloak:clientId:
        value: pulumi
    keycloak:clientSecret:
        value: 884e0f95-0f42-4a63-9b1f-94274655669e
    keycloak:url:
        value: http://localhost:8080

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    keycloak:clientId:
        value: pulumi
    keycloak:clientSecret:
        value: 884e0f95-0f42-4a63-9b1f-94274655669e
    keycloak:url:
        value: http://localhost:8080

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    keycloak:clientId:
        value: pulumi
    keycloak:clientSecret:
        value: 884e0f95-0f42-4a63-9b1f-94274655669e
    keycloak:url:
        value: http://localhost:8080

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    keycloak:clientId:
        value: pulumi
    keycloak:clientSecret:
        value: 884e0f95-0f42-4a63-9b1f-94274655669e
    keycloak:url:
        value: http://localhost:8080

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    keycloak:clientId:
        value: pulumi
    keycloak:clientSecret:
        value: 884e0f95-0f42-4a63-9b1f-94274655669e
    keycloak:url:
        value: http://localhost:8080

```

{{% /choosable %}}
{{< /chooser >}}
## Example Usage (password grant)

{{< chooser language "typescript,python,go,csharp,java,yaml" >}}
{{% choosable language typescript %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: nodejs
config:
    keycloak:clientId:
        value: admin-cli
    keycloak:password:
        value: password
    keycloak:url:
        value: http://localhost:8080
    keycloak:username:
        value: keycloak

```

{{% /choosable %}}
{{% choosable language python %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: python
config:
    keycloak:clientId:
        value: admin-cli
    keycloak:password:
        value: password
    keycloak:url:
        value: http://localhost:8080
    keycloak:username:
        value: keycloak

```

{{% /choosable %}}
{{% choosable language csharp %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: dotnet
config:
    keycloak:clientId:
        value: admin-cli
    keycloak:password:
        value: password
    keycloak:url:
        value: http://localhost:8080
    keycloak:username:
        value: keycloak

```

{{% /choosable %}}
{{% choosable language go %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: go
config:
    keycloak:clientId:
        value: admin-cli
    keycloak:password:
        value: password
    keycloak:url:
        value: http://localhost:8080
    keycloak:username:
        value: keycloak

```

{{% /choosable %}}
{{% choosable language yaml %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: yaml
config:
    keycloak:clientId:
        value: admin-cli
    keycloak:password:
        value: password
    keycloak:url:
        value: http://localhost:8080
    keycloak:username:
        value: keycloak

```

{{% /choosable %}}
{{% choosable language java %}}
```yaml
# Pulumi.yaml provider configuration file
name: configuration-example
runtime: java
config:
    keycloak:clientId:
        value: admin-cli
    keycloak:password:
        value: password
    keycloak:url:
        value: http://localhost:8080
    keycloak:username:
        value: keycloak

```

{{% /choosable %}}
{{< /chooser >}}
## Configuration Reference

The following configuration inputs are supported:

- `clientId` - (Required) The `clientId` for the client that was created in the "Keycloak Setup" section. Use the `admin-cli` client if you are using the password grant. Defaults to the environment variable `KEYCLOAK_CLIENT_ID`.
- `url` - (Required) The URL of the Keycloak instance, before `/auth/admin`. Defaults to the environment variable `KEYCLOAK_URL`.
- `clientSecret` - (Optional) The secret for the client used by the provider for authentication via the client credentials grant. This can be found or changed using the "Credentials" tab in the client settings. Defaults to the environment variable `KEYCLOAK_CLIENT_SECRET`. This attribute is required when using the client credentials grant, and cannot be set when using the password grant.
- `username` - (Optional) The username of the user used by the provider for authentication via the password grant. Defaults to the environment variable `KEYCLOAK_USER`. This attribute is required when using the password grant, and cannot be set when using the client credentials grant.
- `password` - (Optional) The password of the user used by the provider for authentication via the password grant. Defaults to the environment variable `KEYCLOAK_PASSWORD`. This attribute is required when using the password grant, and cannot be set when using the client credentials grant.
- `realm` - (Optional) The realm used by the provider for authentication. Defaults to the environment variable `KEYCLOAK_REALM`, or `master` if the environment variable is not specified.
- `initialLogin` - (Optional) Optionally avoid Keycloak login during provider setup, for when Keycloak itself is being provisioned by pulumi. Defaults to true, which is the original method.
- `clientTimeout` - (Optional) Sets the timeout of the client when addressing Keycloak, in seconds. Defaults to the environment variable `KEYCLOAK_CLIENT_TIMEOUT`, or `5` if the environment variable is not specified.
- `tlsInsecureSkipVerify` - (Optional) Allows ignoring insecure certificates when set to `true`. Defaults to `false`. Disabling this security check is dangerous and should only be done in local or test environments.
- `rootCaCertificate` - (Optional) Allows x509 calls using an unknown CA certificate (for development purposes)
- `basePath` - (Optional) The base path used for accessing the Keycloak REST API.  Defaults to the environment variable `KEYCLOAK_BASE_PATH`, or an empty string if the environment variable is not specified. Note that users of the legacy distribution of Keycloak will need to set this attribute to `/auth`.
- `additionalHeaders` - (Optional) A map of custom HTTP headers to add to each request to the Keycloak API.