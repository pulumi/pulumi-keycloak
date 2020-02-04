// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    public partial class GroupMembershipProtocolMapper : Pulumi.CustomResource
    {
        [Output("addToAccessToken")]
        public Output<bool?> AddToAccessToken { get; private set; } = null!;

        [Output("addToIdToken")]
        public Output<bool?> AddToIdToken { get; private set; } = null!;

        [Output("addToUserinfo")]
        public Output<bool?> AddToUserinfo { get; private set; } = null!;

        [Output("claimName")]
        public Output<string> ClaimName { get; private set; } = null!;

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Output("clientScopeId")]
        public Output<string?> ClientScopeId { get; private set; } = null!;

        [Output("fullPath")]
        public Output<bool?> FullPath { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMembershipProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMembershipProtocolMapper(string name, GroupMembershipProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:OpenId/groupMembershipProtocolMapper:GroupMembershipProtocolMapper", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private GroupMembershipProtocolMapper(string name, Input<string> id, GroupMembershipProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:OpenId/groupMembershipProtocolMapper:GroupMembershipProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMembershipProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMembershipProtocolMapper Get(string name, Input<string> id, GroupMembershipProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMembershipProtocolMapper(name, id, state, options);
        }
    }

    public sealed class GroupMembershipProtocolMapperArgs : Pulumi.ResourceArgs
    {
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

        [Input("addToUserinfo")]
        public Input<bool>? AddToUserinfo { get; set; }

        [Input("claimName", required: true)]
        public Input<string> ClaimName { get; set; } = null!;

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        [Input("fullPath")]
        public Input<bool>? FullPath { get; set; }

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GroupMembershipProtocolMapperArgs()
        {
        }
    }

    public sealed class GroupMembershipProtocolMapperState : Pulumi.ResourceArgs
    {
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

        [Input("addToUserinfo")]
        public Input<bool>? AddToUserinfo { get; set; }

        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        /// <summary>
        /// The mapper's associated client. Cannot be used at the same time as client_scope_id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The mapper's associated client scope. Cannot be used at the same time as client_id.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        [Input("fullPath")]
        public Input<bool>? FullPath { get; set; }

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public GroupMembershipProtocolMapperState()
        {
        }
    }
}
