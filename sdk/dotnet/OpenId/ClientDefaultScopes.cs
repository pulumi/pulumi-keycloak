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
    /// ## Example Usage
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
    ///     var client = new Keycloak.OpenId.Client("client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "test-client",
    ///         AccessType = "CONFIDENTIAL",
    ///     });
    /// 
    ///     var clientScope = new Keycloak.OpenId.ClientScope("client_scope", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "test-client-scope",
    ///     });
    /// 
    ///     var clientDefaultScopes = new Keycloak.OpenId.ClientDefaultScopes("client_default_scopes", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = client.Id,
    ///         DefaultScopes = new[]
    ///         {
    ///             "profile",
    ///             "email",
    ///             "roles",
    ///             "web-origins",
    ///             clientScope.Name,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import. Instead of importing, feel free to create this resource
    /// 
    /// as if it did not already exist on the server.
    /// </summary>
    [KeycloakResourceType("keycloak:openid/clientDefaultScopes:ClientDefaultScopes")]
    public partial class ClientDefaultScopes : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// An array of client scope names to attach to this client.
        /// </summary>
        [Output("defaultScopes")]
        public Output<ImmutableArray<string>> DefaultScopes { get; private set; } = null!;

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a ClientDefaultScopes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientDefaultScopes(string name, ClientDefaultScopesArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientDefaultScopes:ClientDefaultScopes", name, args ?? new ClientDefaultScopesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientDefaultScopes(string name, Input<string> id, ClientDefaultScopesState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientDefaultScopes:ClientDefaultScopes", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientDefaultScopes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientDefaultScopes Get(string name, Input<string> id, ClientDefaultScopesState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientDefaultScopes(name, id, state, options);
        }
    }

    public sealed class ClientDefaultScopesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("defaultScopes", required: true)]
        private InputList<string>? _defaultScopes;

        /// <summary>
        /// An array of client scope names to attach to this client.
        /// </summary>
        public InputList<string> DefaultScopes
        {
            get => _defaultScopes ?? (_defaultScopes = new InputList<string>());
            set => _defaultScopes = value;
        }

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public ClientDefaultScopesArgs()
        {
        }
        public static new ClientDefaultScopesArgs Empty => new ClientDefaultScopesArgs();
    }

    public sealed class ClientDefaultScopesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("defaultScopes")]
        private InputList<string>? _defaultScopes;

        /// <summary>
        /// An array of client scope names to attach to this client.
        /// </summary>
        public InputList<string> DefaultScopes
        {
            get => _defaultScopes ?? (_defaultScopes = new InputList<string>());
            set => _defaultScopes = value;
        }

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public ClientDefaultScopesState()
        {
        }
        public static new ClientDefaultScopesState Empty => new ClientDefaultScopesState();
    }
}
