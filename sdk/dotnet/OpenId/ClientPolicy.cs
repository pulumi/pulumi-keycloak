// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// This resource can be used to create client policy.
    /// 
    /// ## Example Usage
    /// 
    /// In this example, we'll create a new OpenID client, then enabled permissions for the client. A client without permissions disabled cannot be assigned by a client policy. We'll use the `keycloak.openid.ClientPolicy` resource to create a new client policy, which could be applied to many clients, for a realm and a resource_server_id.
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var realm = new Keycloak.Realm("realm", new()
    ///     {
    ///         RealmName = "my-realm",
    ///         Enabled = true,
    ///     });
    /// 
    ///     var openidClient = new Keycloak.OpenId.Client("openid_client", new()
    ///     {
    ///         ClientId = "openid_client",
    ///         Name = "openid_client",
    ///         RealmId = realm.Id,
    ///         AccessType = "CONFIDENTIAL",
    ///         ServiceAccountsEnabled = true,
    ///     });
    /// 
    ///     var myPermission = new Keycloak.OpenId.ClientPermissions("my_permission", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = openidClient.Id,
    ///     });
    /// 
    ///     var realmManagement = Keycloak.OpenId.GetClient.Invoke(new()
    ///     {
    ///         RealmId = "my-realm",
    ///         ClientId = "realm-management",
    ///     });
    /// 
    ///     var tokenExchange = new Keycloak.OpenId.ClientPolicy("token_exchange", new()
    ///     {
    ///         ResourceServerId = realmManagement.Apply(getClientResult =&gt; getClientResult.Id),
    ///         RealmId = realm.Id,
    ///         Name = "my-policy",
    ///         Logic = "POSITIVE",
    ///         DecisionStrategy = "UNANIMOUS",
    ///         Clients = new[]
    ///         {
    ///             openidClient.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:openid/clientPolicy:ClientPolicy")]
    public partial class ClientPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The clients allowed by this client policy.
        /// </summary>
        [Output("clients")]
        public Output<ImmutableArray<string>> Clients { get; private set; } = null!;

        /// <summary>
        /// (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
        /// </summary>
        [Output("decisionStrategy")]
        public Output<string?> DecisionStrategy { get; private set; } = null!;

        /// <summary>
        /// The description of this client policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
        /// </summary>
        [Output("logic")]
        public Output<string?> Logic { get; private set; } = null!;

        /// <summary>
        /// The name of this client policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm this client policy exists within.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// The ID of the resource server this client policy is attached to.
        /// </summary>
        [Output("resourceServerId")]
        public Output<string> ResourceServerId { get; private set; } = null!;


        /// <summary>
        /// Create a ClientPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientPolicy(string name, ClientPolicyArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientPolicy:ClientPolicy", name, args ?? new ClientPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientPolicy(string name, Input<string> id, ClientPolicyState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientPolicy:ClientPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientPolicy Get(string name, Input<string> id, ClientPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientPolicy(name, id, state, options);
        }
    }

    public sealed class ClientPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("clients", required: true)]
        private InputList<string>? _clients;

        /// <summary>
        /// The clients allowed by this client policy.
        /// </summary>
        public InputList<string> Clients
        {
            get => _clients ?? (_clients = new InputList<string>());
            set => _clients = value;
        }

        /// <summary>
        /// (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
        /// </summary>
        [Input("decisionStrategy")]
        public Input<string>? DecisionStrategy { get; set; }

        /// <summary>
        /// The description of this client policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
        /// </summary>
        [Input("logic")]
        public Input<string>? Logic { get; set; }

        /// <summary>
        /// The name of this client policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this client policy exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// The ID of the resource server this client policy is attached to.
        /// </summary>
        [Input("resourceServerId", required: true)]
        public Input<string> ResourceServerId { get; set; } = null!;

        public ClientPolicyArgs()
        {
        }
        public static new ClientPolicyArgs Empty => new ClientPolicyArgs();
    }

    public sealed class ClientPolicyState : global::Pulumi.ResourceArgs
    {
        [Input("clients")]
        private InputList<string>? _clients;

        /// <summary>
        /// The clients allowed by this client policy.
        /// </summary>
        public InputList<string> Clients
        {
            get => _clients ?? (_clients = new InputList<string>());
            set => _clients = value;
        }

        /// <summary>
        /// (Computed) Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
        /// </summary>
        [Input("decisionStrategy")]
        public Input<string>? DecisionStrategy { get; set; }

        /// <summary>
        /// The description of this client policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Computed) Dictates how the policy decision should be made. Can be either `POSITIVE` or `NEGATIVE`. Applies to policies.
        /// </summary>
        [Input("logic")]
        public Input<string>? Logic { get; set; }

        /// <summary>
        /// The name of this client policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this client policy exists within.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// The ID of the resource server this client policy is attached to.
        /// </summary>
        [Input("resourceServerId")]
        public Input<string>? ResourceServerId { get; set; }

        public ClientPolicyState()
        {
        }
        public static new ClientPolicyState Empty => new ClientPolicyState();
    }
}
