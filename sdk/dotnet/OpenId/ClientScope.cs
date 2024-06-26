// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId
{
    /// <summary>
    /// ## # keycloak.openid.ClientScope
    /// 
    /// Allows for creating and managing Keycloak client scopes that can be attached to
    /// clients that use the OpenID Connect protocol.
    /// 
    /// Client Scopes can be used to share common protocol and role mappings between multiple
    /// clients within a realm. They can also be used by clients to conditionally request
    /// claims or roles for a user based on the OAuth 2.0 `scope` parameter.
    /// 
    /// ### Example Usage
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
    ///     var openidClientScope = new Keycloak.OpenId.ClientScope("openid_client_scope", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "groups",
    ///         Description = "When requested, this scope will map a user's group memberships to a claim",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this client scope belongs to.
    /// - `name` - (Required) The display name of this client scope in the GUI.
    /// - `description` - (Optional) The description of this client scope in the GUI.
    /// - `consent_screen_text` - (Optional) When set, a consent screen will be displayed to users
    ///   authenticating to clients with this scope attached. The consent screen will display the string
    ///   value of this attribute.
    /// 
    /// ### Import
    /// 
    /// Client scopes can be imported using the format `{{realm_id}}/{{client_scope_id}}`, where `client_scope_id` is the unique ID that Keycloak
    /// assigns to the client scope upon creation. This value can be found in the URI when editing this client scope in the GUI, and is typically a GUID.
    /// 
    /// Example:
    /// </summary>
    [KeycloakResourceType("keycloak:openid/clientScope:ClientScope")]
    public partial class ClientScope : global::Pulumi.CustomResource
    {
        [Output("consentScreenText")]
        public Output<string?> ConsentScreenText { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("guiOrder")]
        public Output<int?> GuiOrder { get; private set; } = null!;

        [Output("includeInTokenScope")]
        public Output<bool?> IncludeInTokenScope { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a ClientScope resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClientScope(string name, ClientScopeArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientScope:ClientScope", name, args ?? new ClientScopeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClientScope(string name, Input<string> id, ClientScopeState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/clientScope:ClientScope", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClientScope resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClientScope Get(string name, Input<string> id, ClientScopeState? state = null, CustomResourceOptions? options = null)
        {
            return new ClientScope(name, id, state, options);
        }
    }

    public sealed class ClientScopeArgs : global::Pulumi.ResourceArgs
    {
        [Input("consentScreenText")]
        public Input<string>? ConsentScreenText { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("guiOrder")]
        public Input<int>? GuiOrder { get; set; }

        [Input("includeInTokenScope")]
        public Input<bool>? IncludeInTokenScope { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public ClientScopeArgs()
        {
        }
        public static new ClientScopeArgs Empty => new ClientScopeArgs();
    }

    public sealed class ClientScopeState : global::Pulumi.ResourceArgs
    {
        [Input("consentScreenText")]
        public Input<string>? ConsentScreenText { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("guiOrder")]
        public Input<int>? GuiOrder { get; set; }

        [Input("includeInTokenScope")]
        public Input<bool>? IncludeInTokenScope { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public ClientScopeState()
        {
        }
        public static new ClientScopeState Empty => new ClientScopeState();
    }
}
