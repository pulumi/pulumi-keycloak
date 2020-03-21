using Pulumi;
using Keycloak = Pulumi.Keycloak;

class KeyCloakStack : Stack
{
    public KeyCloakStack()
    {
        var realm = new Keycloak.Realm("new-csharp-realm", new Keycloak.RealmArgs
        {
            RealmName = "my-example-csharp-realm",
        });
    }
}
