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
    /// Allows you to manage all group Scope Based Permissions https://www.keycloak.org/docs/latest/server_admin/#group.
    /// 
    /// This is part of a preview Keycloak feature: `admin_fine_grained_authz` (see https://www.keycloak.org/docs/latest/server_admin/#_fine_grain_permissions).
    /// This feature can be enabled with the Keycloak option `-Dkeycloak.profile.feature.admin_fine_grained_authz=enabled`. See the
    /// example `docker-compose.yml` file for an example.
    /// 
    /// When enabling Roles Permissions, Keycloak does several things automatically:
    /// 1. Enable Authorization on built-in `realm-management` client (if not already enabled).
    /// 2. Create a resource representing the role permissions.
    /// 3. Create scopes `view`, `manage`, `view-members`, `manage-members`, `manage-membership`.
    /// 4. Create all scope based permission for the scopes and role resource
    /// </summary>
    [KeycloakResourceType("keycloak:index/groupPermissions:GroupPermissions")]
    public partial class GroupPermissions : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Resource server id representing the realm management client on which this permission is managed
        /// </summary>
        [Output("authorizationResourceServerId")]
        public Output<string> AuthorizationResourceServerId { get; private set; } = null!;

        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        [Output("manageMembersScope")]
        public Output<Outputs.GroupPermissionsManageMembersScope?> ManageMembersScope { get; private set; } = null!;

        [Output("manageMembershipScope")]
        public Output<Outputs.GroupPermissionsManageMembershipScope?> ManageMembershipScope { get; private set; } = null!;

        [Output("manageScope")]
        public Output<Outputs.GroupPermissionsManageScope?> ManageScope { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("viewMembersScope")]
        public Output<Outputs.GroupPermissionsViewMembersScope?> ViewMembersScope { get; private set; } = null!;

        [Output("viewScope")]
        public Output<Outputs.GroupPermissionsViewScope?> ViewScope { get; private set; } = null!;


        /// <summary>
        /// Create a GroupPermissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupPermissions(string name, GroupPermissionsArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/groupPermissions:GroupPermissions", name, args ?? new GroupPermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupPermissions(string name, Input<string> id, GroupPermissionsState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/groupPermissions:GroupPermissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupPermissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupPermissions Get(string name, Input<string> id, GroupPermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupPermissions(name, id, state, options);
        }
    }

    public sealed class GroupPermissionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("manageMembersScope")]
        public Input<Inputs.GroupPermissionsManageMembersScopeArgs>? ManageMembersScope { get; set; }

        [Input("manageMembershipScope")]
        public Input<Inputs.GroupPermissionsManageMembershipScopeArgs>? ManageMembershipScope { get; set; }

        [Input("manageScope")]
        public Input<Inputs.GroupPermissionsManageScopeArgs>? ManageScope { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("viewMembersScope")]
        public Input<Inputs.GroupPermissionsViewMembersScopeArgs>? ViewMembersScope { get; set; }

        [Input("viewScope")]
        public Input<Inputs.GroupPermissionsViewScopeArgs>? ViewScope { get; set; }

        public GroupPermissionsArgs()
        {
        }
        public static new GroupPermissionsArgs Empty => new GroupPermissionsArgs();
    }

    public sealed class GroupPermissionsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Resource server id representing the realm management client on which this permission is managed
        /// </summary>
        [Input("authorizationResourceServerId")]
        public Input<string>? AuthorizationResourceServerId { get; set; }

        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("manageMembersScope")]
        public Input<Inputs.GroupPermissionsManageMembersScopeGetArgs>? ManageMembersScope { get; set; }

        [Input("manageMembershipScope")]
        public Input<Inputs.GroupPermissionsManageMembershipScopeGetArgs>? ManageMembershipScope { get; set; }

        [Input("manageScope")]
        public Input<Inputs.GroupPermissionsManageScopeGetArgs>? ManageScope { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("viewMembersScope")]
        public Input<Inputs.GroupPermissionsViewMembersScopeGetArgs>? ViewMembersScope { get; set; }

        [Input("viewScope")]
        public Input<Inputs.GroupPermissionsViewScopeGetArgs>? ViewScope { get; set; }

        public GroupPermissionsState()
        {
        }
        public static new GroupPermissionsState Empty => new GroupPermissionsState();
    }
}
