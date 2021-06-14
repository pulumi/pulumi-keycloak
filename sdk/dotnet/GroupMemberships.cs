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
    /// ## Import
    /// 
    /// This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server. [1]providers/mrparkers/keycloak/latest/docs/resources/group_memberships
    /// </summary>
    [KeycloakResourceType("keycloak:index/groupMemberships:GroupMemberships")]
    public partial class GroupMemberships : Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the group this resource should manage memberships for.
        /// </summary>
        [Output("groupId")]
        public Output<string?> GroupId { get; private set; } = null!;

        /// <summary>
        /// A list of usernames that belong to this group.
        /// </summary>
        [Output("members")]
        public Output<ImmutableArray<string>> Members { get; private set; } = null!;

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMemberships resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMemberships(string name, GroupMembershipsArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/groupMemberships:GroupMemberships", name, args ?? new GroupMembershipsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMemberships(string name, Input<string> id, GroupMembershipsState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/groupMemberships:GroupMemberships", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMemberships resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMemberships Get(string name, Input<string> id, GroupMembershipsState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMemberships(name, id, state, options);
        }
    }

    public sealed class GroupMembershipsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the group this resource should manage memberships for.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("members", required: true)]
        private InputList<string>? _members;

        /// <summary>
        /// A list of usernames that belong to this group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GroupMembershipsArgs()
        {
        }
    }

    public sealed class GroupMembershipsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the group this resource should manage memberships for.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        [Input("members")]
        private InputList<string>? _members;

        /// <summary>
        /// A list of usernames that belong to this group.
        /// </summary>
        public InputList<string> Members
        {
            get => _members ?? (_members = new InputList<string>());
            set => _members = value;
        }

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public GroupMembershipsState()
        {
        }
    }
}
