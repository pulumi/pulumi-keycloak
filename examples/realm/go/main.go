package main

import (
	"github.com/pulumi/pulumi-keycloak/sdk/v5/go/keycloak"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		realm, err := keycloak.NewRealm(ctx, "new-go-realm", &keycloak.RealmArgs{
			Realm: pulumi.String("my-example-go-realm"),
		})
		if err != nil {
			return err
		}

		ctx.Export("realmId", realm.ID())

		return nil
	})
}
