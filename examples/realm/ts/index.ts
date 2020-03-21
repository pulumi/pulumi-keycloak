import * as keycloak from "@pulumi/keycloak";

const realm = new keycloak.Realm("new-typescript-realm", {
    realm: "my-example-typescript-realm"
});

export const realmId = realm.id;
