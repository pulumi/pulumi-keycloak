[![Actions Status](https://github.com/pulumi/pulumi-keycloak/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-keycloak/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fkeycloak.svg)](https://www.npmjs.com/package/@pulumi/keycloak)
[![Python version](https://badge.fury.io/py/pulumi-keycloak.svg)](https://pypi.org/project/pulumi-keycloak)
[![NuGet version](https://badge.fury.io/nu/pulumi.keycloak.svg)](https://badge.fury.io/nu/pulumi.keycloak)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-keycloak/sdk/v5/go)](https://pkg.go.dev/github.com/pulumi/pulumi-keycloak/sdk/v5/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-keycloak/blob/master/LICENSE)

# Keycloak Resource Provider

The Keycloak resource provider for Pulumi lets you manage Keycloak resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/reference/cli/).

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

    $ go get github.com/pulumi/pulumi-keycloak/sdk/v5

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
