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
    /// Allows for creating and managing protocol mappers for both types of clients (openid-connect and saml) within Keycloak.
    /// 
    /// There are two uses cases for using this resource:
    /// * If you implemented a custom protocol mapper, this resource can be used to configure it
    /// * If the provider doesn't support a particular protocol mapper, this resource can be used instead.
    /// 
    /// Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
    /// Therefore, if possible, a specific mapper should be used instead.
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
    ///     var samlClient = new Keycloak.Saml.Client("samlClient", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "test-client",
    ///     });
    /// 
    ///     var samlHardcodeAttributeMapper = new Keycloak.GenericProtocolMapper("samlHardcodeAttributeMapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = samlClient.Id,
    ///         Protocol = "saml",
    ///         ProtocolMapper = "saml-hardcode-attribute-mapper",
    ///         Config = 
    ///         {
    ///             { "attribute.name", "name" },
    ///             { "attribute.nameformat", "Basic" },
    ///             { "attribute.value", "value" },
    ///             { "friendly.name", "display name" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Protocol mappers can be imported using the following format: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
    /// 
    ///  Example:
    /// 
    ///  bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:index/genericProtocolMapper:GenericProtocolMapper saml_hardcode_attribute_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/genericProtocolMapper:GenericProtocolMapper")]
    public partial class GenericProtocolMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The ID of the client this protocol mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// The ID of the client scope this protocol mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        /// </summary>
        [Output("clientScopeId")]
        public Output<string?> ClientScopeId { get; private set; } = null!;

        /// <summary>
        /// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
        /// </summary>
        [Output("config")]
        public Output<ImmutableDictionary<string, object>> Config { get; private set; } = null!;

        /// <summary>
        /// The display name of this protocol mapper in the GUI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
        /// </summary>
        [Output("protocolMapper")]
        public Output<string> ProtocolMapper { get; private set; } = null!;

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a GenericProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GenericProtocolMapper(string name, GenericProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/genericProtocolMapper:GenericProtocolMapper", name, args ?? new GenericProtocolMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GenericProtocolMapper(string name, Input<string> id, GenericProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/genericProtocolMapper:GenericProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GenericProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GenericProtocolMapper Get(string name, Input<string> id, GenericProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new GenericProtocolMapper(name, id, state, options);
        }
    }

    public sealed class GenericProtocolMapperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the client this protocol mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The ID of the client scope this protocol mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        [Input("config", required: true)]
        private InputMap<object>? _config;

        /// <summary>
        /// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
        /// </summary>
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// The display name of this protocol mapper in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
        /// </summary>
        [Input("protocolMapper", required: true)]
        public Input<string> ProtocolMapper { get; set; } = null!;

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GenericProtocolMapperArgs()
        {
        }
        public static new GenericProtocolMapperArgs Empty => new GenericProtocolMapperArgs();
    }

    public sealed class GenericProtocolMapperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ID of the client this protocol mapper should be added to. Conflicts with `client_scope_id`. This argument is required if `client_scope_id` is not set.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The ID of the client scope this protocol mapper should be added to. Conflicts with `client_id`. This argument is required if `client_id` is not set.
        /// </summary>
        [Input("clientScopeId")]
        public Input<string>? ClientScopeId { get; set; }

        [Input("config")]
        private InputMap<object>? _config;

        /// <summary>
        /// A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
        /// </summary>
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// The display name of this protocol mapper in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
        /// </summary>
        [Input("protocolMapper")]
        public Input<string>? ProtocolMapper { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public GenericProtocolMapperState()
        {
        }
        public static new GenericProtocolMapperState Empty => new GenericProtocolMapperState();
    }
}
