#!/bin/bash

# Wait for keycloak to come up
./testing/wait-for-local-keycloak.sh

# Run the terraform client
./testing/create-terraform-client.sh
