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
    /// ## # keycloak.openid.ClientPermissions
    /// 
    /// Allows you to manage all openid client Scope Based Permissions.
    /// 
    /// This is part of a preview keycloak feature. You need to enable this feature to be able to use this resource. More
    /// information about enabling the preview feature can be found
    /// here: https://www.keycloak.org/docs/latest/securing_apps/index.html#_token-exchange
    /// 
    /// When enabling Openid Client Permissions, Keycloak does several things automatically:
    /// 
    /// 1. Enable Authorization on build-in realm-management client
    /// 2. Create scopes "view", "manage", "configure", "map-roles", "map-roles-client-scope", "map-roles-composite", "
    ///    token-exchange"
    /// 3. Create a resource representing the openid client
    /// 4. Create all scope based permission for the scopes and openid client resource
    /// 
    /// If the realm-management Authorization is not enable, you have to ceate a dependency (`depends_on`) with the policy and
    /// the openid client.
    /// 
    /// ### Example Usage
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
    ///         RealmName = "realm",
    ///     });
    /// 
    ///     var myOpenidClient = new Keycloak.OpenId.Client("myOpenidClient", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "my_openid_client",
    ///         ClientSecret = "secret",
    ///         AccessType = "CONFIDENTIAL",
    ///         StandardFlowEnabled = true,
    ///         ValidRedirectUris = new[]
    ///         {
    ///             "http://localhost:8080/*",
    ///         },
    ///     });
    /// 
    ///     var realmManagement = Keycloak.OpenId.GetClient.Invoke(new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "realm-management",
    ///     });
    /// 
    ///     var testUser = new Keycloak.User("testUser", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Username = "test-user",
    ///         Email = "test-user@fakedomain.com",
    ///         FirstName = "Testy",
    ///         LastName = "Tester",
    ///     });
    /// 
    ///     var testClientUserPolicy = new Keycloak.OpenId.ClientUserPolicy("testClientUserPolicy", new()
    ///     {
    ///         ResourceServerId = realmManagement.Apply(getClientResult =&gt; getClientResult.Id),
    ///         RealmId = realm.Id,
    ///         Users = new[]
    ///         {
    ///             testUser.Id,
    ///         },
    ///         Logic = "POSITIVE",
    ///         DecisionStrategy = "UNANIMOUS",
    ///     }, new CustomResourceOptions
    ///     {
    ///         DependsOn = new[]
    ///         {
    ///             myOpenidClient,
    ///         },
    ///     });
    /// 
    ///     var myPermission = new Keycloak.OpenId.ClientPermissions("myPermission", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = myOpenidClient.Id,
    ///         ViewScope = new Keycloak.OpenId.Inputs.ClientPermissionsViewScopeArgs
    ///         {
    ///             Policies = new[]
    ///             {
    ///                 testClientUserPolicy.Id,
    ///             },
    ///             Description = "my description",
    ///             DecisionStrategy = "UNANIMOUS",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this group exists in.
    /// - `client_id` - (Required) The id of the client that provides the role.
    /// 
    /// #### Permission Scopes
    /// 
    /// Permission scopes can be defined using the following attributes:
    /// 
    /// - `view_scope`
    /// - `manage_scope`
    /// - `configure_scope`
    /// - `map_roles_scope`
    /// - `map_roles_client_scope_scope`
    /// - `map_roles_composite_scope`
    /// - `token_exchange_scope`
    /// 
    /// Each of these attributes have the following schema:
    /// 
    /// - `policies` - (Optional) A list of policy IDs
    /// - `description` - (Optional) A description for the permission scope
    /// - `decision_strategy` - (Optional) The decision strategy, can be one of `UNANIMOUS`, `AFFIRMATIVE`, or `CONSENSUS`.
    /// 
    /// ### Attributes Reference
    /// 
    /// In addition to the arguments listed above, the following computed attributes are exported:
    /// 
    /// - `authorization_resource_server_id` - Resource server id representing the realm management client on which this
    ///   permission is managed.
    /// </summary>
    [KeycloakResourceType("keycloak:openid/clientPermissions:ClientPermissions")]
    public partial class ClientPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Resource server id representing the realm management client on which this permission is managed
        /// </summary>
        [Output("authorizationResourceServerId")]
        public Output<string> AuthorizationResourceServerId { get; private set; } = null!;

        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        [Output("configureScope")]
        public Output<Outputs.ClientPermissionsConfigureScope?> ConfigureScope { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("manageScope")]
        public Output<Outputs.ClientPermissionsManageScope?> ManageScope { get; private set; } = null!;

        [Output("mapRolesClientScopeScope")]
        public Output<Outputs.ClientPermissionsMapRolesClientScopeScope?> MapRolesClientScopeScope { get; private set; } = null!;

        [Output("mapRolesCompositeScope")]
        public Output<Outputs.ClientPermissionsMapRolesCompositeScope?> MapRolesCompositeScope { get; private set; } = null!;

        [Output("mapRolesScope")]
        public Output<Outputs.ClientPermissionsMapRolesScope?> MapRolesScope { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("tokenExchangeScope")]
        public Output<Outputs.ClientPermissionsTokenExchangeScope?> TokenExchangeScope { get; private set; } = null!;

        [Output("viewScope")]
        public Output<Outputs.ClientPermissionsViewScope?> ViewScope { get; private set; } = null!;


        /// <summary>
        /// Create a ClientPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientPermissions(string name, ClientPermissionsArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientPermissions:ClientPermissions", name, args ?? new ClientPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientPermissions(string name, Input<string> id, ClientPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientPermissions:ClientPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientPermissions Get(string name, Input<string> id, ClientPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientPermissions(name, id, state, options);
        }
    }

    public sealed class ClientPermissionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("configureScope")]
        public Input<Inputs.ClientPermissionsConfigureScopeArgs>? ConfigureScope { get; set; }

        [Input("manageScope")]
        public Input<Inputs.ClientPermissionsManageScopeArgs>? ManageScope { get; set; }

        [Input("mapRolesClientScopeScope")]
        public Input<Inputs.ClientPermissionsMapRolesClientScopeScopeArgs>? MapRolesClientScopeScope { get; set; }

        [Input("mapRolesCompositeScope")]
        public Input<Inputs.ClientPermissionsMapRolesCompositeScopeArgs>? MapRolesCompositeScope { get; set; }

        [Input("mapRolesScope")]
        public Input<Inputs.ClientPermissionsMapRolesScopeArgs>? MapRolesScope { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("tokenExchangeScope")]
        public Input<Inputs.ClientPermissionsTokenExchangeScopeArgs>? TokenExchangeScope { get; set; }

        [Input("viewScope")]
        public Input<Inputs.ClientPermissionsViewScopeArgs>? ViewScope { get; set; }

        public ClientPermissionsArgs()
        {
        }
        public static new ClientPermissionsArgs Empty => new ClientPermissionsArgs();
    }

    public sealed class ClientPermissionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource server id representing the realm management client on which this permission is managed
        /// </summary>
        [Input("authorizationResourceServerId")]
        public Input<string>? AuthorizationResourceServerId { get; set; }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("configureScope")]
        public Input<Inputs.ClientPermissionsConfigureScopeGetArgs>? ConfigureScope { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("manageScope")]
        public Input<Inputs.ClientPermissionsManageScopeGetArgs>? ManageScope { get; set; }

        [Input("mapRolesClientScopeScope")]
        public Input<Inputs.ClientPermissionsMapRolesClientScopeScopeGetArgs>? MapRolesClientScopeScope { get; set; }

        [Input("mapRolesCompositeScope")]
        public Input<Inputs.ClientPermissionsMapRolesCompositeScopeGetArgs>? MapRolesCompositeScope { get; set; }

        [Input("mapRolesScope")]
        public Input<Inputs.ClientPermissionsMapRolesScopeGetArgs>? MapRolesScope { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("tokenExchangeScope")]
        public Input<Inputs.ClientPermissionsTokenExchangeScopeGetArgs>? TokenExchangeScope { get; set; }

        [Input("viewScope")]
        public Input<Inputs.ClientPermissionsViewScopeGetArgs>? ViewScope { get; set; }

        public ClientPermissionsState()
        {
        }
        public static new ClientPermissionsState Empty => new ClientPermissionsState();
    }
}
