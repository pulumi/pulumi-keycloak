// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// Allows for assigning realm roles to the service account of an openid client.
    /// You need to set `service_accounts_enabled` to `true` for the openid client that should be assigned the role.
    /// 
    /// If you'd like to attach client roles to a service account, please use the `keycloak.openid.ClientServiceAccountRole`
    /// resource.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var realm = new Keycloak.Realm("realm", new Keycloak.RealmArgs
    ///         {
    ///             RealmName = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var realmRole = new Keycloak.Role("realmRole", new Keycloak.RoleArgs
    ///         {
    ///             RealmId = realm.Id,
    ///         });
    ///         var client = new Keycloak.OpenId.Client("client", new Keycloak.OpenId.ClientArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ServiceAccountsEnabled = true,
    ///         });
    ///         var clientServiceAccountRole = new Keycloak.OpenId.ClientServiceAccountRealmRole("clientServiceAccountRole", new Keycloak.OpenId.ClientServiceAccountRealmRoleArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ServiceAccountUserId = client.ServiceAccountUserId,
    ///             Role = realmRole.Name,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the format `{{realmId}}/{{serviceAccountUserId}}/{{roleId}}`. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole client_service_account_role my-realm/489ba513-1ceb-49ba-ae0b-1ab1f5099ebf/c7230ab7-8e4e-4135-995d-e81b50696ad8
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole")]
    public partial class ClientServiceAccountRealmRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The realm that the client and role belong to.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// The name of the role that is assigned.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        /// </summary>
        [Output("serviceAccountUserId")]
        public Output<string> ServiceAccountUserId { get; private set; } = null!;


        /// <summary>
        /// Create a ClientServiceAccountRealmRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientServiceAccountRealmRole(string name, ClientServiceAccountRealmRoleArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole", name, args ?? new ClientServiceAccountRealmRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientServiceAccountRealmRole(string name, Input<string> id, ClientServiceAccountRealmRoleState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientServiceAccountRealmRole:ClientServiceAccountRealmRole", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ClientServiceAccountRealmRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientServiceAccountRealmRole Get(string name, Input<string> id, ClientServiceAccountRealmRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientServiceAccountRealmRole(name, id, state, options);
        }
    }

    public sealed class ClientServiceAccountRealmRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The realm that the client and role belong to.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// The name of the role that is assigned.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        /// </summary>
        [Input("serviceAccountUserId", required: true)]
        public Input<string> ServiceAccountUserId { get; set; } = null!;

        public ClientServiceAccountRealmRoleArgs()
        {
        }
    }

    public sealed class ClientServiceAccountRealmRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The realm that the client and role belong to.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// The name of the role that is assigned.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
        /// </summary>
        [Input("serviceAccountUserId")]
        public Input<string>? ServiceAccountUserId { get; set; }

        public ClientServiceAccountRealmRoleState()
        {
        }
    }
}
