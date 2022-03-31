// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
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
    ///         var tokenExchangeRealm = new Keycloak.Realm("tokenExchangeRealm", new Keycloak.RealmArgs
    ///         {
    ///             Realm = "token-exchange_destination_realm",
    ///             Enabled = true,
    ///         });
    ///         var tokenExchangeMyOidcIdp = new Keycloak.Oidc.IdentityProvider("tokenExchangeMyOidcIdp", new Keycloak.Oidc.IdentityProviderArgs
    ///         {
    ///             Realm = tokenExchangeRealm.Id,
    ///             Alias = "myIdp",
    ///             AuthorizationUrl = "http://localhost:8080/auth/realms/someRealm/protocol/openid-connect/auth",
    ///             TokenUrl = "http://localhost:8080/auth/realms/someRealm/protocol/openid-connect/token",
    ///             ClientId = "clientId",
    ///             ClientSecret = "secret",
    ///             DefaultScopes = "openid",
    ///         });
    ///         var token_exchangeWebappClient = new Keycloak.OpenId.Client("token-exchangeWebappClient", new Keycloak.OpenId.ClientArgs
    ///         {
    ///             RealmId = tokenExchangeRealm.Id,
    ///             ClientId = "webapp_client",
    ///             ClientSecret = "secret",
    ///             Description = "a webapp client on the destination realm",
    ///             AccessType = "CONFIDENTIAL",
    ///             StandardFlowEnabled = true,
    ///             ValidRedirectUris = 
    ///             {
    ///                 "http://localhost:8080/*",
    ///             },
    ///         });
    ///         //relevant part
    ///         var oidcIdpPermission = new Keycloak.IdentityProviderTokenExchangeScopePermission("oidcIdpPermission", new Keycloak.IdentityProviderTokenExchangeScopePermissionArgs
    ///         {
    ///             RealmId = tokenExchangeRealm.Id,
    ///             ProviderAlias = tokenExchangeMyOidcIdp.Alias,
    ///             PolicyType = "client",
    ///             Clients = 
    ///             {
    ///                 token_exchangeWebappClient.Id,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the format `{{realm_id}}/{{provider_alias}}`, where `provider_alias` is the alias that you assign to the identity provider upon creation. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission oidc_idp_permission my-realm/myIdp
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission")]
    public partial class IdentityProviderTokenExchangeScopePermission : Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
        /// </summary>
        [Output("authorizationIdpResourceId")]
        public Output<string> AuthorizationIdpResourceId { get; private set; } = null!;

        /// <summary>
        /// (Computed) Resource server ID representing the realm management client on which this permission is managed.
        /// </summary>
        [Output("authorizationResourceServerId")]
        public Output<string> AuthorizationResourceServerId { get; private set; } = null!;

        /// <summary>
        /// (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
        /// </summary>
        [Output("authorizationTokenExchangeScopePermissionId")]
        public Output<string> AuthorizationTokenExchangeScopePermissionId { get; private set; } = null!;

        /// <summary>
        /// A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        /// </summary>
        [Output("clients")]
        public Output<ImmutableArray<string>> Clients { get; private set; } = null!;

        /// <summary>
        /// (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
        /// </summary>
        [Output("policyId")]
        public Output<string> PolicyId { get; private set; } = null!;

        /// <summary>
        /// Defaults to "client" This is also the only value policy type supported by this provider.
        /// </summary>
        [Output("policyType")]
        public Output<string?> PolicyType { get; private set; } = null!;

        /// <summary>
        /// Alias of the identity provider.
        /// </summary>
        [Output("providerAlias")]
        public Output<string> ProviderAlias { get; private set; } = null!;

        /// <summary>
        /// The realm that the identity provider exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a IdentityProviderTokenExchangeScopePermission resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IdentityProviderTokenExchangeScopePermission(string name, IdentityProviderTokenExchangeScopePermissionArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission", name, args ?? new IdentityProviderTokenExchangeScopePermissionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IdentityProviderTokenExchangeScopePermission(string name, Input<string> id, IdentityProviderTokenExchangeScopePermissionState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/identityProviderTokenExchangeScopePermission:IdentityProviderTokenExchangeScopePermission", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IdentityProviderTokenExchangeScopePermission resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IdentityProviderTokenExchangeScopePermission Get(string name, Input<string> id, IdentityProviderTokenExchangeScopePermissionState? state = null, CustomResourceOptions? options = null)
        {
            return new IdentityProviderTokenExchangeScopePermission(name, id, state, options);
        }
    }

    public sealed class IdentityProviderTokenExchangeScopePermissionArgs : Pulumi.ResourceArgs
    {
        [Input("clients", required: true)]
        private InputList<string>? _clients;

        /// <summary>
        /// A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        /// </summary>
        public InputList<string> Clients
        {
            get => _clients ?? (_clients = new InputList<string>());
            set => _clients = value;
        }

        /// <summary>
        /// Defaults to "client" This is also the only value policy type supported by this provider.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Alias of the identity provider.
        /// </summary>
        [Input("providerAlias", required: true)]
        public Input<string> ProviderAlias { get; set; } = null!;

        /// <summary>
        /// The realm that the identity provider exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public IdentityProviderTokenExchangeScopePermissionArgs()
        {
        }
    }

    public sealed class IdentityProviderTokenExchangeScopePermissionState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) Resource ID representing the identity provider, this automatically created by keycloak.
        /// </summary>
        [Input("authorizationIdpResourceId")]
        public Input<string>? AuthorizationIdpResourceId { get; set; }

        /// <summary>
        /// (Computed) Resource server ID representing the realm management client on which this permission is managed.
        /// </summary>
        [Input("authorizationResourceServerId")]
        public Input<string>? AuthorizationResourceServerId { get; set; }

        /// <summary>
        /// (Computed) Permission ID representing the Permission with scope 'Token Exchange' and the resource 'authorization_idp_resource_id', this automatically created by keycloak, the policy ID will be set on this permission.
        /// </summary>
        [Input("authorizationTokenExchangeScopePermissionId")]
        public Input<string>? AuthorizationTokenExchangeScopePermissionId { get; set; }

        [Input("clients")]
        private InputList<string>? _clients;

        /// <summary>
        /// A list of IDs of the clients for which a policy will be created and set on scope based token exchange permission.
        /// </summary>
        public InputList<string> Clients
        {
            get => _clients ?? (_clients = new InputList<string>());
            set => _clients = value;
        }

        /// <summary>
        /// (Computed) Policy ID that will be set on the scope based token exchange permission automatically created by enabling permissions on the reference identity provider.
        /// </summary>
        [Input("policyId")]
        public Input<string>? PolicyId { get; set; }

        /// <summary>
        /// Defaults to "client" This is also the only value policy type supported by this provider.
        /// </summary>
        [Input("policyType")]
        public Input<string>? PolicyType { get; set; }

        /// <summary>
        /// Alias of the identity provider.
        /// </summary>
        [Input("providerAlias")]
        public Input<string>? ProviderAlias { get; set; }

        /// <summary>
        /// The realm that the identity provider exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public IdentityProviderTokenExchangeScopePermissionState()
        {
        }
    }
}
