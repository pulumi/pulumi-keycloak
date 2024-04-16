// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml
{
    /// <summary>
    /// Allows for creating and managing script protocol mappers for SAML clients within Keycloak.
    /// 
    /// Script protocol mappers evaluate a JavaScript function to produce an attribute value based on context information.
    /// 
    /// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
    /// multiple different clients.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
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
    ///     var samlClient = new Keycloak.Saml.Client("saml_client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "saml-client",
    ///         Name = "saml-client",
    ///     });
    /// 
    ///     var samlScriptMapper = new Keycloak.Saml.ScriptProtocolMapper("saml_script_mapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = samlClient.Id,
    ///         Name = "script-mapper",
    ///         Script = "exports = 'foo';",
    ///         SamlAttributeName = "displayName",
    ///         SamlAttributeNameFormat = "Unspecified",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Protocol mappers can be imported using one of the following formats:
    /// 
    /// - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
    /// 
    /// - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
    /// 
    /// Example:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper saml_script_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper saml_script_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper")]
    public partial class ScriptProtocolMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Output("clientScopeId")]
        public Output<string?> ClientScopeId { get; private set; } = null!;

        /// <summary>
        /// An optional human-friendly name for this attribute.
        /// </summary>
        [Output("friendlyName")]
        public Output<string?> FriendlyName { get; private set; } = null!;

        /// <summary>
        /// The display name of this protocol mapper in the GUI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// The name of the SAML attribute.
        /// </summary>
        [Output("samlAttributeName")]
        public Output<string> SamlAttributeName { get; private set; } = null!;

        /// <summary>
        /// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
        /// </summary>
        [Output("samlAttributeNameFormat")]
        public Output<string> SamlAttributeNameFormat { get; private set; } = null!;

        /// <summary>
        /// JavaScript code to compute the attribute value.
        /// </summary>
        [Output("script")]
        public Output<string> Script { get; private set; } = null!;

        /// <summary>
        /// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
        /// </summary>
        [Output("singleValueAttribute")]
        public Output<bool?> SingleValueAttribute { get; private set; } = null!;


        /// <summary>
        /// Create a ScriptProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ScriptProtocolMapper(string name, ScriptProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper", name, args ?? new ScriptProtocolMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ScriptProtocolMapper(string name, Input<string> id, ScriptProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:saml/scriptProtocolMapper:ScriptProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ScriptProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ScriptProtocolMapper Get(string name, Input<string> id, ScriptProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new ScriptProtocolMapper(name, id, state, options);
        }
    }

    public sealed class ScriptProtocolMapperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// An optional human-friendly name for this attribute.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// The display name of this protocol mapper in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// The name of the SAML attribute.
        /// </summary>
        [Input("samlAttributeName", required: true)]
        public Input<string> SamlAttributeName { get; set; } = null!;

        /// <summary>
        /// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
        /// </summary>
        [Input("samlAttributeNameFormat", required: true)]
        public Input<string> SamlAttributeNameFormat { get; set; } = null!;

        /// <summary>
        /// JavaScript code to compute the attribute value.
        /// </summary>
        [Input("script", required: true)]
        public Input<string> Script { get; set; } = null!;

        /// <summary>
        /// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
        /// </summary>
        [Input("singleValueAttribute")]
        public Input<bool>? SingleValueAttribute { get; set; }

        public ScriptProtocolMapperArgs()
        {
        }
        public static new ScriptProtocolMapperArgs Empty => new ScriptProtocolMapperArgs();
    }

    public sealed class ScriptProtocolMapperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        /// <summary>
        /// An optional human-friendly name for this attribute.
        /// </summary>
        [Input("friendlyName")]
        public Input<string>? FriendlyName { get; set; }

        /// <summary>
        /// The display name of this protocol mapper in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// The name of the SAML attribute.
        /// </summary>
        [Input("samlAttributeName")]
        public Input<string>? SamlAttributeName { get; set; }

        /// <summary>
        /// The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
        /// </summary>
        [Input("samlAttributeNameFormat")]
        public Input<string>? SamlAttributeNameFormat { get; set; }

        /// <summary>
        /// JavaScript code to compute the attribute value.
        /// </summary>
        [Input("script")]
        public Input<string>? Script { get; set; }

        /// <summary>
        /// When `true`, all values will be stored under one attribute with multiple attribute values. Defaults to `true`.
        /// </summary>
        [Input("singleValueAttribute")]
        public Input<bool>? SingleValueAttribute { get; set; }

        public ScriptProtocolMapperState()
        {
        }
        public static new ScriptProtocolMapperState Empty => new ScriptProtocolMapperState();
    }
}
