// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// &gt; This content is derived from https://github.com/mrparkers/terraform-provider-keycloak/blob/master/website/docs/r/openid_client_optional_scopes.html.markdown.
    /// </summary>
    public partial class ClientOptionalScopes : Pulumi.CustomResource
    {
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        [Output("optionalScopes")]
        public Output<ImmutableArray<string>> OptionalScopes { get; private set; } = null!;

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
            : base("keycloak:OpenId/clientOptionalScopes:ClientOptionalScopes", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private ClientOptionalScopes(string name, Input<string> id, ClientOptionalScopesState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:OpenId/clientOptionalScopes:ClientOptionalScopes", name, state, MakeResourceOptions(options, id))
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
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        [Input("optionalScopes", required: true)]
        private InputList<string>? _optionalScopes;
        public InputList<string> OptionalScopes
        {
            get => _optionalScopes ?? (_optionalScopes = new InputList<string>());
            set => _optionalScopes = value;
        }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public ClientOptionalScopesArgs()
        {
        }
    }

    public sealed class ClientOptionalScopesState : Pulumi.ResourceArgs
    {
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("optionalScopes")]
        private InputList<string>? _optionalScopes;
        public InputList<string> OptionalScopes
        {
            get => _optionalScopes ?? (_optionalScopes = new InputList<string>());
            set => _optionalScopes = value;
        }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public ClientOptionalScopesState()
        {
        }
    }
}
