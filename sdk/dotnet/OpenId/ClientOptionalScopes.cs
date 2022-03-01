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
    ///             Realm = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var client = new Keycloak.OpenId.Client("client", new Keycloak.OpenId.ClientArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = "test-client",
    ///             AccessType = "CONFIDENTIAL",
    ///         });
    ///         var clientScope = new Keycloak.OpenId.ClientScope("clientScope", new Keycloak.OpenId.ClientScopeArgs
    ///         {
    ///             RealmId = realm.Id,
    ///         });
    ///         var clientOptionalScopes = new Keycloak.OpenId.ClientOptionalScopes("clientOptionalScopes", new Keycloak.OpenId.ClientOptionalScopesArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             ClientId = client.Id,
    ///             OptionalScopes = 
    ///             {
    ///                 "address",
    ///                 "phone",
    ///                 "offline_access",
    ///                 "microprofile-jwt",
    ///                 clientScope.Name,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server.
    /// </summary>
    [KeycloakResourceType("keycloak:openid/clientOptionalScopes:ClientOptionalScopes")]
    public partial class ClientOptionalScopes : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the client to attach optional scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// An array of client scope names to attach to this client as optional scopes.
        /// </summary>
        [Output("optionalScopes")]
        public Output<ImmutableArray<string>> OptionalScopes { get; private set; } = null!;

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a ClientOptionalScopes resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientOptionalScopes(string name, ClientOptionalScopesArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientOptionalScopes:ClientOptionalScopes", name, args ?? new ClientOptionalScopesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientOptionalScopes(string name, Input<string> id, ClientOptionalScopesState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientOptionalScopes:ClientOptionalScopes", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientOptionalScopes resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientOptionalScopes Get(string name, Input<string> id, ClientOptionalScopesState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientOptionalScopes(name, id, state, options);
        }
    }

    public sealed class ClientOptionalScopesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the client to attach optional scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("optionalScopes", required: true)]
        private InputList<string>? _optionalScopes;

        /// <summary>
        /// An array of client scope names to attach to this client as optional scopes.
        /// </summary>
        public InputList<string> OptionalScopes
        {
            get => _optionalScopes ?? (_optionalScopes = new InputList<string>());
            set => _optionalScopes = value;
        }

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public ClientOptionalScopesArgs()
        {
        }
    }

    public sealed class ClientOptionalScopesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the client to attach optional scopes to. Note that this is the unique ID of the client generated by Keycloak.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("optionalScopes")]
        private InputList<string>? _optionalScopes;

        /// <summary>
        /// An array of client scope names to attach to this client as optional scopes.
        /// </summary>
        public InputList<string> OptionalScopes
        {
            get => _optionalScopes ?? (_optionalScopes = new InputList<string>());
            set => _optionalScopes = value;
        }

        /// <summary>
        /// The realm this client and scopes exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public ClientOptionalScopesState()
        {
        }
    }
}
