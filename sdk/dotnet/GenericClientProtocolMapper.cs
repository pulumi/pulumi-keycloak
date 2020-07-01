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
    /// ## # keycloak.GenericClientProtocolMapper
    /// 
    /// Allows for creating and managing protocol mapper for both types of clients (openid-connect and saml) within Keycloak.
    /// 
    /// There are two uses cases for using this resource:
    /// * If you implemented a custom protocol mapper, this resource can be used to configure it
    /// * If the provider doesn't support a particular protocol mapper, this resource can be used instead.
    /// 
    /// Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
    /// Therefore, if possible, a specific mapper should be used.
    /// 
    /// ### Example Usage
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
    ///             Enabled = true,
    ///             Realm = "my-realm",
    ///         });
    ///         var samlClient = new Keycloak.Saml.Client("samlClient", new Keycloak.Saml.ClientArgs
    ///         {
    ///             ClientId = "test-client",
    ///             RealmId = realm.Id,
    ///         });
    ///         var samlHardcodeAttributeMapper = new Keycloak.GenericClientProtocolMapper("samlHardcodeAttributeMapper", new Keycloak.GenericClientProtocolMapperArgs
    ///         {
    ///             ClientId = samlClient.Id,
    ///             Config = 
    ///             {
    ///                 { "attribute.name", "name" },
    ///                 { "attribute.nameformat", "Basic" },
    ///                 { "attribute.value", "value" },
    ///                 { "friendly.name", "display name" },
    ///             },
    ///             Protocol = "saml",
    ///             ProtocolMapper = "saml-hardcode-attribute-mapper",
    ///             RealmId = realm.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this protocol mapper exists within.
    /// - `client_id` - (Required) The client this protocol mapper is attached to.
    /// - `name` - (Required) The display name of this protocol mapper in the GUI.
    /// - `protocol` - (Required) The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
    /// - `protocol_mapper` - (Required) The name of the protocol mapper. The protocol mapper must be
    ///    compatible with the specified client.
    /// - `config` - (Required) A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
    /// </summary>
    public partial class GenericClientProtocolMapper : Pulumi.CustomResource
    {
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

        [Output("config")]
        public Output<ImmutableDictionary<string, object>> Config { get; private set; } = null!;

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The protocol of the client (openid-connect / saml).
        /// </summary>
        [Output("protocol")]
        public Output<string> Protocol { get; private set; } = null!;

        /// <summary>
        /// The type of the protocol mapper.
        /// </summary>
        [Output("protocolMapper")]
        public Output<string> ProtocolMapper { get; private set; } = null!;

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a GenericClientProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GenericClientProtocolMapper(string name, GenericClientProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper", name, args ?? new GenericClientProtocolMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GenericClientProtocolMapper(string name, Input<string> id, GenericClientProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/genericClientProtocolMapper:GenericClientProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GenericClientProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GenericClientProtocolMapper Get(string name, Input<string> id, GenericClientProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new GenericClientProtocolMapper(name, id, state, options);
        }
    }

    public sealed class GenericClientProtocolMapperArgs : Pulumi.ResourceArgs
    {
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

        [Input("config", required: true)]
        private InputMap<object>? _config;
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol of the client (openid-connect / saml).
        /// </summary>
        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        /// <summary>
        /// The type of the protocol mapper.
        /// </summary>
        [Input("protocolMapper", required: true)]
        public Input<string> ProtocolMapper { get; set; } = null!;

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GenericClientProtocolMapperArgs()
        {
        }
    }

    public sealed class GenericClientProtocolMapperState : Pulumi.ResourceArgs
    {
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

        [Input("config")]
        private InputMap<object>? _config;
        public InputMap<object> Config
        {
            get => _config ?? (_config = new InputMap<object>());
            set => _config = value;
        }

        /// <summary>
        /// A human-friendly name that will appear in the Keycloak console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The protocol of the client (openid-connect / saml).
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// The type of the protocol mapper.
        /// </summary>
        [Input("protocolMapper")]
        public Input<string>? ProtocolMapper { get; set; }

        /// <summary>
        /// The realm id where the associated client or client scope exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public GenericClientProtocolMapperState()
        {
        }
    }
}
