from pulumi import export
import pulumi_keycloak as keycloak

realm = keycloak.Realm("new-python-realm",
  realm="my-example-python-realm"
)

export("realm_id", realm.id)
