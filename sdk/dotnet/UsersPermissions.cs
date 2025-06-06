// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    /// <summary>
    /// Allows you to manage fine-grained permissions for all users in a realm: https://www.keycloak.org/docs/latest/server_admin/#_users-permissions
    /// 
    /// This is part of a preview Keycloak feature: `admin_fine_grained_authz` (see https://www.keycloak.org/docs/latest/server_admin/#_fine_grain_permissions).
    /// This feature can be enabled with the Keycloak option `-Dkeycloak.profile.feature.admin_fine_grained_authz=enabled`. See the
    /// example `docker-compose.yml` file for an example.
    /// 
    /// When enabling fine-grained permissions for users, Keycloak does several things automatically:
    /// 1. Enable Authorization on built-in `realm-management` client (if not already enabled).
    /// 2. Create a resource representing the users permissions.
    /// 3. Create scopes `view`, `manage`, `map-roles`, `manage-group-membership`, `impersonate`, and `user-impersonated`.
    /// 4. Create all scope based permission for the scopes and users resources.
    /// 
    /// &gt; This resource should only be created once per realm.
    /// </summary>
    [KeycloakResourceType("keycloak:index/usersPermissions:UsersPermissions")]
    public partial class UsersPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Resource server id representing the realm management client on which this permission is managed
        /// </summary>
        [Output("authorizationResourceServerId")]
        public Output<string> AuthorizationResourceServerId { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("impersonateScope")]
        public Output<Outputs.UsersPermissionsImpersonateScope?> ImpersonateScope { get; private set; } = null!;

        [Output("manageGroupMembershipScope")]
        public Output<Outputs.UsersPermissionsManageGroupMembershipScope?> ManageGroupMembershipScope { get; private set; } = null!;

        [Output("manageScope")]
        public Output<Outputs.UsersPermissionsManageScope?> ManageScope { get; private set; } = null!;

        [Output("mapRolesScope")]
        public Output<Outputs.UsersPermissionsMapRolesScope?> MapRolesScope { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("userImpersonatedScope")]
        public Output<Outputs.UsersPermissionsUserImpersonatedScope?> UserImpersonatedScope { get; private set; } = null!;

        [Output("viewScope")]
        public Output<Outputs.UsersPermissionsViewScope?> ViewScope { get; private set; } = null!;


        /// <summary>
        /// Create a UsersPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UsersPermissions(string name, UsersPermissionsArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/usersPermissions:UsersPermissions", name, args ?? new UsersPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UsersPermissions(string name, Input<string> id, UsersPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/usersPermissions:UsersPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UsersPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UsersPermissions Get(string name, Input<string> id, UsersPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new UsersPermissions(name, id, state, options);
        }
    }

    public sealed class UsersPermissionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("impersonateScope")]
        public Input<Inputs.UsersPermissionsImpersonateScopeArgs>? ImpersonateScope { get; set; }

        [Input("manageGroupMembershipScope")]
        public Input<Inputs.UsersPermissionsManageGroupMembershipScopeArgs>? ManageGroupMembershipScope { get; set; }

        [Input("manageScope")]
        public Input<Inputs.UsersPermissionsManageScopeArgs>? ManageScope { get; set; }

        [Input("mapRolesScope")]
        public Input<Inputs.UsersPermissionsMapRolesScopeArgs>? MapRolesScope { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("userImpersonatedScope")]
        public Input<Inputs.UsersPermissionsUserImpersonatedScopeArgs>? UserImpersonatedScope { get; set; }

        [Input("viewScope")]
        public Input<Inputs.UsersPermissionsViewScopeArgs>? ViewScope { get; set; }

        public UsersPermissionsArgs()
        {
        }
        public static new UsersPermissionsArgs Empty => new UsersPermissionsArgs();
    }

    public sealed class UsersPermissionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource server id representing the realm management client on which this permission is managed
        /// </summary>
        [Input("authorizationResourceServerId")]
        public Input<string>? AuthorizationResourceServerId { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("impersonateScope")]
        public Input<Inputs.UsersPermissionsImpersonateScopeGetArgs>? ImpersonateScope { get; set; }

        [Input("manageGroupMembershipScope")]
        public Input<Inputs.UsersPermissionsManageGroupMembershipScopeGetArgs>? ManageGroupMembershipScope { get; set; }

        [Input("manageScope")]
        public Input<Inputs.UsersPermissionsManageScopeGetArgs>? ManageScope { get; set; }

        [Input("mapRolesScope")]
        public Input<Inputs.UsersPermissionsMapRolesScopeGetArgs>? MapRolesScope { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("userImpersonatedScope")]
        public Input<Inputs.UsersPermissionsUserImpersonatedScopeGetArgs>? UserImpersonatedScope { get; set; }

        [Input("viewScope")]
        public Input<Inputs.UsersPermissionsViewScopeGetArgs>? ViewScope { get; set; }

        public UsersPermissionsState()
        {
        }
        public static new UsersPermissionsState Empty => new UsersPermissionsState();
    }
}
