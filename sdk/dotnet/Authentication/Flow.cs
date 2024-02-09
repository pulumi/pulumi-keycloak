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
    /// Allows for creating and managing an authentication flow within Keycloak.
    /// 
    /// [Authentication flows](https://www.keycloak.org/docs/11.0/server_admin/index.html#_authentication-flows) describe a sequence
    /// of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
    /// is a container for these actions, which are otherwise known as executions.
    /// 
    /// ## Example Usage
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
    ///         RealmName = "my-realm",
    ///         Enabled = true,
    ///     });
    /// 
    ///     var flow = new Keycloak.Authentication.Flow("flow", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Alias = "my-flow-alias",
    ///     });
    /// 
    ///     var execution = new Keycloak.Authentication.Execution("execution", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ParentFlowAlias = flow.Alias,
    ///         Authenticator = "identity-provider-redirector",
    ///         Requirement = "REQUIRED",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Authentication flows can be imported using the format `{{realmId}}/{{authenticationFlowId}}`. The authentication flow ID is
    /// 
    ///  typically a GUID which is autogenerated when the flow is created via Keycloak.
    /// 
    ///  Unfortunately, it is not trivial to retrieve the authentication flow ID from the UI. The best way to do this is to visit the
    /// 
    ///  "Authentication" page in Keycloak, and use the network tab of your browser to view the response of the API call to `/auth/admin/realms/${realm}/authentication/flows`,
    /// 
    ///  which will be a list of authentication flows.
    /// 
    ///  Example:
    /// 
    ///  bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:authentication/flow:Flow flow my-realm/e9a5641e-778c-4daf-89c0-f4ef617987d1
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:authentication/flow:Flow")]
    public partial class Flow : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The alias for this authentication flow.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// A description for the authentication flow.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        /// </summary>
        [Output("providerId")]
        public Output<string?> ProviderId { get; private set; } = null!;

        /// <summary>
        /// The realm that the authentication flow exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a Flow resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Flow(string name, FlowArgs args, CustomResourceOptions? options = null)
            : base("keycloak:authentication/flow:Flow", name, args ?? new FlowArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Flow(string name, Input<string> id, FlowState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:authentication/flow:Flow", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Flow resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Flow Get(string name, Input<string> id, FlowState? state = null, CustomResourceOptions? options = null)
        {
            return new Flow(name, id, state, options);
        }
    }

    public sealed class FlowArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias for this authentication flow.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// A description for the authentication flow.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The realm that the authentication flow exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public FlowArgs()
        {
        }
        public static new FlowArgs Empty => new FlowArgs();
    }

    public sealed class FlowState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias for this authentication flow.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// A description for the authentication flow.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The type of authentication flow to create. Valid choices include `basic-flow` and `client-flow`. Defaults to `basic-flow`.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The realm that the authentication flow exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public FlowState()
        {
        }
        public static new FlowState Empty => new FlowState();
    }
}
