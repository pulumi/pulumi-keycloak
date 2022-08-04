// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Authentication
{
    /// <summary>
    /// Allows for creating and managing an authentication subflow within Keycloak.
    /// 
    /// Like authentication flows, authentication subflows are containers for authentication executions.
    /// As its name implies, an authentication subflow is contained in an authentication flow.
    /// 
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
    ///             RealmName = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var flow = new Keycloak.Authentication.Flow("flow", new Keycloak.Authentication.FlowArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             Alias = "my-flow-alias",
    ///         });
    ///         var subflow = new Keycloak.Authentication.Subflow("subflow", new Keycloak.Authentication.SubflowArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             Alias = "my-subflow-alias",
    ///             ParentFlowAlias = flow.Alias,
    ///             ProviderId = "basic-flow",
    ///             Requirement = "ALTERNATIVE",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Authentication flows can be imported using the format `{{realmId}}/{{parentFlowAlias}}/{{authenticationSubflowId}}`. The authentication subflow ID is typically a GUID which is autogenerated when the subflow is created via Keycloak. Unfortunately, it is not trivial to retrieve the authentication subflow ID from the UI. The best way to do this is to visit the "Authentication" page in Keycloak, and use the network tab of your browser to view the response of the API call to `/auth/admin/realms/${realm}/authentication/flows/{flow}/executions`, which will be a list of executions, where the subflow will be. __The subflow ID is contained in the `flowID` field__ (not, as one could guess, the `id` field). Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:authentication/subflow:Subflow subflow my-realm/"Parent Flow"/3bad1172-bb5c-4a77-9615-c2606eb03081
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:authentication/subflow:Subflow")]
    public partial class Subflow : Pulumi.CustomResource
    {
        /// <summary>
        /// The alias for this authentication subflow.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// The name of the authenticator. Might be needed to be set with certain custom subflows with specific
        /// authenticators. In general this will remain empty.
        /// </summary>
        [Output("authenticator")]
        public Output<string?> Authenticator { get; private set; } = null!;

        /// <summary>
        /// A description for the authentication subflow.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The alias for the parent authentication flow.
        /// </summary>
        [Output("parentFlowAlias")]
        public Output<string> ParentFlowAlias { get; private set; } = null!;

        /// <summary>
        /// The type of authentication subflow to create. Valid choices include `basic-flow`, `form-flow`
        /// and `client-flow`. Defaults to `basic-flow`.
        /// </summary>
        [Output("providerId")]
        public Output<string?> ProviderId { get; private set; } = null!;

        /// <summary>
        /// The realm that the authentication subflow exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`,
        /// or `DISABLED`. Defaults to `DISABLED`.
        /// </summary>
        [Output("requirement")]
        public Output<string?> Requirement { get; private set; } = null!;


        /// <summary>
        /// Create a Subflow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Subflow(string name, SubflowArgs args, CustomResourceOptions? options = null)
            : base("keycloak:authentication/subflow:Subflow", name, args ?? new SubflowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Subflow(string name, Input<string> id, SubflowState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:authentication/subflow:Subflow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Subflow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Subflow Get(string name, Input<string> id, SubflowState? state = null, CustomResourceOptions? options = null)
        {
            return new Subflow(name, id, state, options);
        }
    }

    public sealed class SubflowArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias for this authentication subflow.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// The name of the authenticator. Might be needed to be set with certain custom subflows with specific
        /// authenticators. In general this will remain empty.
        /// </summary>
        [Input("authenticator")]
        public Input<string>? Authenticator { get; set; }

        /// <summary>
        /// A description for the authentication subflow.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The alias for the parent authentication flow.
        /// </summary>
        [Input("parentFlowAlias", required: true)]
        public Input<string> ParentFlowAlias { get; set; } = null!;

        /// <summary>
        /// The type of authentication subflow to create. Valid choices include `basic-flow`, `form-flow`
        /// and `client-flow`. Defaults to `basic-flow`.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The realm that the authentication subflow exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`,
        /// or `DISABLED`. Defaults to `DISABLED`.
        /// </summary>
        [Input("requirement")]
        public Input<string>? Requirement { get; set; }

        public SubflowArgs()
        {
        }
    }

    public sealed class SubflowState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias for this authentication subflow.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// The name of the authenticator. Might be needed to be set with certain custom subflows with specific
        /// authenticators. In general this will remain empty.
        /// </summary>
        [Input("authenticator")]
        public Input<string>? Authenticator { get; set; }

        /// <summary>
        /// A description for the authentication subflow.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The alias for the parent authentication flow.
        /// </summary>
        [Input("parentFlowAlias")]
        public Input<string>? ParentFlowAlias { get; set; }

        /// <summary>
        /// The type of authentication subflow to create. Valid choices include `basic-flow`, `form-flow`
        /// and `client-flow`. Defaults to `basic-flow`.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The realm that the authentication subflow exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`,
        /// or `DISABLED`. Defaults to `DISABLED`.
        /// </summary>
        [Input("requirement")]
        public Input<string>? Requirement { get; set; }

        public SubflowState()
        {
        }
    }
}
