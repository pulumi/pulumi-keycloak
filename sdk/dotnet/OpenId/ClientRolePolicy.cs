// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    [KeycloakResourceType("keycloak:openid/clientRolePolicy:ClientRolePolicy")]
    public partial class ClientRolePolicy : global::Pulumi.CustomResource
    {
        [Output("decisionStrategy")]
        public Output<string?> DecisionStrategy { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("logic")]
        public Output<string?> Logic { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("resourceServerId")]
        public Output<string> ResourceServerId { get; private set; } = null!;

        [Output("roles")]
        public Output<ImmutableArray<Outputs.ClientRolePolicyRole>> Roles { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ClientRolePolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientRolePolicy(string name, ClientRolePolicyArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientRolePolicy:ClientRolePolicy", name, args ?? new ClientRolePolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientRolePolicy(string name, Input<string> id, ClientRolePolicyState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientRolePolicy:ClientRolePolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientRolePolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientRolePolicy Get(string name, Input<string> id, ClientRolePolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientRolePolicy(name, id, state, options);
        }
    }

    public sealed class ClientRolePolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("decisionStrategy")]
        public Input<string>? DecisionStrategy { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("logic")]
        public Input<string>? Logic { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("resourceServerId", required: true)]
        public Input<string> ResourceServerId { get; set; } = null!;

        [Input("roles", required: true)]
        private InputList<Inputs.ClientRolePolicyRoleArgs>? _roles;
        public InputList<Inputs.ClientRolePolicyRoleArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.ClientRolePolicyRoleArgs>());
            set => _roles = value;
        }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ClientRolePolicyArgs()
        {
        }
        public static new ClientRolePolicyArgs Empty => new ClientRolePolicyArgs();
    }

    public sealed class ClientRolePolicyState : global::Pulumi.ResourceArgs
    {
        [Input("decisionStrategy")]
        public Input<string>? DecisionStrategy { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("logic")]
        public Input<string>? Logic { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("resourceServerId")]
        public Input<string>? ResourceServerId { get; set; }

        [Input("roles")]
        private InputList<Inputs.ClientRolePolicyRoleGetArgs>? _roles;
        public InputList<Inputs.ClientRolePolicyRoleGetArgs> Roles
        {
            get => _roles ?? (_roles = new InputList<Inputs.ClientRolePolicyRoleGetArgs>());
            set => _roles = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        public ClientRolePolicyState()
        {
        }
        public static new ClientRolePolicyState Empty => new ClientRolePolicyState();
    }
}
