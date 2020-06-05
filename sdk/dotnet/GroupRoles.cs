// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public partial class GroupRoles : Pulumi.CustomResource
    {
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("roleIds")]
        public Output<ImmutableArray<string>> RoleIds { get; private set; } = null!;


        /// <summary>
        /// Create a GroupRoles resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupRoles(string name, GroupRolesArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/groupRoles:GroupRoles", name, args ?? new GroupRolesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupRoles(string name, Input<string> id, GroupRolesState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/groupRoles:GroupRoles", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupRoles resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupRoles Get(string name, Input<string> id, GroupRolesState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupRoles(name, id, state, options);
        }
    }

    public sealed class GroupRolesArgs : Pulumi.ResourceArgs
    {
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("roleIds", required: true)]
        private InputList<string>? _roleIds;
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        public GroupRolesArgs()
        {
        }
    }

    public sealed class GroupRolesState : Pulumi.ResourceArgs
    {
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("roleIds")]
        private InputList<string>? _roleIds;
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        public GroupRolesState()
        {
        }
    }
}
