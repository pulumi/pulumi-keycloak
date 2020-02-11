import * as keycloak from "@pulumi/keycloak";

const realm = new keycloak.Realm("new-realm", {
    realm: "my-example-realm"
});

export const realmId = realm.id;
