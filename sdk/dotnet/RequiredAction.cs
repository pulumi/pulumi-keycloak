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
    /// Allows for creating and managing required actions within Keycloak.
    /// 
    /// [Required actions](https://www.keycloak.org/docs/latest/server_admin/#con-required-actions_server_administration_guide) specify actions required before the first login of all new users.
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
    ///     var requiredAction = new Keycloak.RequiredAction("required_action", new()
    ///     {
    ///         RealmId = realm.RealmName,
    ///         Alias = "webauthn-register",
    ///         Enabled = true,
    ///         Name = "Webauthn Register",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Authentication executions can be imported using the formats: `{{realm}}/{{alias}}`.
    /// 
    /// Example:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:index/requiredAction:RequiredAction required_action my-realm/my-default-action-alias
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/requiredAction:RequiredAction")]
    public partial class RequiredAction : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The alias of the action to attach as a required action.
        /// </summary>
        [Output("alias")]
        public Output<string> Alias { get; private set; } = null!;

        /// <summary>
        /// When `true`, the required action is set as the default action for new users. Defaults to `false`.
        /// </summary>
        [Output("defaultAction")]
        public Output<bool?> DefaultAction { get; private set; } = null!;

        /// <summary>
        /// When `false`, the required action is not enabled for new users. Defaults to `false`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The name of the required action.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The priority of the required action.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// The realm the required action exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a RequiredAction resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RequiredAction(string name, RequiredActionArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/requiredAction:RequiredAction", name, args ?? new RequiredActionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RequiredAction(string name, Input<string> id, RequiredActionState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/requiredAction:RequiredAction", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RequiredAction resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RequiredAction Get(string name, Input<string> id, RequiredActionState? state = null, CustomResourceOptions? options = null)
        {
            return new RequiredAction(name, id, state, options);
        }
    }

    public sealed class RequiredActionArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias of the action to attach as a required action.
        /// </summary>
        [Input("alias", required: true)]
        public Input<string> Alias { get; set; } = null!;

        /// <summary>
        /// When `true`, the required action is set as the default action for new users. Defaults to `false`.
        /// </summary>
        [Input("defaultAction")]
        public Input<bool>? DefaultAction { get; set; }

        /// <summary>
        /// When `false`, the required action is not enabled for new users. Defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of the required action.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority of the required action.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The realm the required action exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public RequiredActionArgs()
        {
        }
        public static new RequiredActionArgs Empty => new RequiredActionArgs();
    }

    public sealed class RequiredActionState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The alias of the action to attach as a required action.
        /// </summary>
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        /// <summary>
        /// When `true`, the required action is set as the default action for new users. Defaults to `false`.
        /// </summary>
        [Input("defaultAction")]
        public Input<bool>? DefaultAction { get; set; }

        /// <summary>
        /// When `false`, the required action is not enabled for new users. Defaults to `false`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of the required action.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The priority of the required action.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The realm the required action exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public RequiredActionState()
        {
        }
        public static new RequiredActionState Empty => new RequiredActionState();
    }
}
