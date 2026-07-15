# Contributing to the Pulumi ecosystem

Do you want to contribute to Pulumi? Awesome! We are so happy to have you.
We have a few tips and housekeeping items to help you get up and running.

## Code of Conduct

Please make sure to read and observe our [Code of Conduct](./CODE-OF-CONDUCT.md)

## Community Expectations

Please read about our [contribution guidelines here.](https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md#communications)

## Setting up your development environment

### Pulumi prerequisites

Please refer to the [main Pulumi repo](https://github.com/pulumi/pulumi/)'s [CONTRIBUTING.md file](
https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md#developing) for details on how to get set up with Pulumi.

## Committing Generated Code

You must generate and check in the SDKs on each pull request containing a code change, e.g. adding a new resource to `resources.go`.

1. Run `make build_sdks` from the root of this repository
1. Open a pull request containing all changes
1. *Note:* If a large number of seemingly-unrelated diffs are produced by `make build_sdks` (for example, lots of changes to comments unrelated to the change you are making), ensure that the latest dependencies for the provider are installed by running `go mod tidy` in the `provider/` directory of this repository.

## Running Integration Tests

Integration tests require a running Keycloak server. The repository includes a
Docker Compose environment for manual testing; it is not started by the regular
test suite or CI.

Start Keycloak from the repository root:

```sh
export KEYCLOAK_PORT=${KEYCLOAK_PORT:-8080}
docker compose -f docker-compose.test.yml up -d

until curl --fail --silent \
    http://localhost:${KEYCLOAK_PORT}/realms/master > /dev/null; do
    sleep 1
done
```

The environment uses Keycloak 26.6.3 by default. Export `KEYCLOAK_VERSION` to
test another version, or `KEYCLOAK_PORT` if port 8080 is unavailable, before
starting the environment:

```sh
export KEYCLOAK_VERSION=26.5.7
export KEYCLOAK_PORT=8081
docker compose -f docker-compose.test.yml up -d
```

Configure the provider to authenticate as the local bootstrap administrator:

```sh
export KEYCLOAK_URL=http://localhost:${KEYCLOAK_PORT:-8080}
export KEYCLOAK_CLIENT_ID=admin-cli
export KEYCLOAK_USER=keycloak
export KEYCLOAK_PASSWORD=password
unset KEYCLOAK_CLIENT_SECRET
```

You can now run a Pulumi program or focused reproduction against the local
server. These credentials and the development-mode server are for local testing
only.

Destroy resources created by the Pulumi program before stopping Keycloak. Then
remove the local environment and its data:

```sh
pulumi destroy
docker compose -f docker-compose.test.yml down --volumes
```

