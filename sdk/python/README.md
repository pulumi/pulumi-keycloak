[![Build Status](https://travis-ci.com/pulumi/pulumi-keycloak.svg?token=eHg7Zp5zdDDJfTjY8ejq&branch=master)](https://travis-ci.com/pulumi/pulumi-keycloak)

# Keycloak Resource Provider

The Keycloak resource provider for Pulumi lets you manage Keycloak resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://www.mailgun.com//).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/keycloak

or `yarn`:

    $ yarn add @pulumi/keycloak

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_keycloak

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-keycloak/sdk/go/...

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Keycloak

## Configuration

The following configuration points are available:

- `keycloak:clientId` - (Required) The client_id for the client that was created in the "Keycloak Setup" section. 
  Use the admin-cli client if you are using the password grant. Defaults to the environment variable `KEYCLOAK_CLIENT_ID`.
- `keycloak:url` - (Required) - The URL of the Keycloak instance, before /auth/admin. Defaults to the environment 
  variable `KEYCLOAK_URL`.
- `keycloak:clientSecret` - (Optional) The secret for the client used by the provider for authentication via the client
  credentials grant. This can be found or changed using the "Credentials" tab in the client settings. Defaults to the 
  environment variable `KEYCLOAK_CLIENT_SECRET`. This attribute is required when using the client credentials grant,
  and cannot be set when using the password grant.
- `keycloak:username`- (Optional) The username of the user used by the provider for authentication via the password grant.
  Defaults to environment variable `KEYCLOAK_USER`. This attribute is required when using the password grant, and cannot 
  be set when using the client credentials grant.
- `keycloak:password`- (Optional) The password of the user used by the provider for authentication via the password grant. Defaults to 
  environment variable `KEYCLOAK_PASSWORD`. This attribute is required when using the password grant, and cannot be set when
  using the client credentials grant.
- `keycloak:realm` - (Optional) The realm used by the provider for authentication. Defaults to environment variable 
  `KEYCLOAK_REALM`, or `master` if the environment variable is not specified.
- `keycloak:initialLogin` - (Optional) Optionally avoid Keycloak login during provider setup, for when Keycloak itself 
  is being provisioned by terraform. Defaults to `true`, which is the original method.
- `keycloak:clientTimeout` - (Optional) Sets the timeout of the client when addressing Keycloak, in seconds. Defaults to `5`.

## Reference

For further information, please visit [the Keycloak provider docs](https://www.pulumi.com/docs/intro/cloud-providers/keycloak) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/keycloak).
