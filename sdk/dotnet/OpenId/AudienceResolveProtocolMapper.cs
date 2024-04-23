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
    /// Allows for creating the "Audience Resolve" OIDC protocol mapper within Keycloak.
    /// 
    /// This protocol mapper is useful to avoid manual management of audiences, instead relying on the presence of client roles
    /// to imply which audiences are appropriate for the token. See the
    /// [Keycloak docs](https://www.keycloak.org/docs/latest/server_admin/#_audience_resolve) for more details.
    /// 
    /// ## Example Usage
    /// 
    /// ### Client)
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
    ///     var openidClient = new Keycloak.OpenId.Client("openid_client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "client",
    ///         Name = "client",
    ///         Enabled = true,
    ///         AccessType = "CONFIDENTIAL",
    ///         ValidRedirectUris = new[]
    ///         {
    ///             "http://localhost:8080/openid-callback",
    ///         },
    ///     });
    /// 
    ///     var audienceMapper = new Keycloak.OpenId.AudienceResolveProtocolMapper("audience_mapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = openidClient.Id,
    ///         Name = "my-audience-resolve-mapper",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Client Scope)
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
    ///     var clientScope = new Keycloak.OpenId.ClientScope("client_scope", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "test-client-scope",
    ///     });
    /// 
    ///     var audienceMapper = new Keycloak.OpenId.AudienceProtocolMapper("audience_mapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientScopeId = clientScope.Id,
    ///     });
    /// 
    /// });
    /// ```
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
    /// $ pulumi import keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper audience_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// 
    /// ```sh
    /// $ pulumi import keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper audience_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper")]
    public partial class AudienceResolveProtocolMapper : global::Pulumi.CustomResource
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
        /// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a AudienceResolveProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AudienceResolveProtocolMapper(string name, AudienceResolveProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper", name, args ?? new AudienceResolveProtocolMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AudienceResolveProtocolMapper(string name, Input<string> id, AudienceResolveProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/audienceResolveProtocolMapper:AudienceResolveProtocolMapper", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                Aliases =
                {
                    new global::Pulumi.Alias { Type = "keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter" },
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AudienceResolveProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AudienceResolveProtocolMapper Get(string name, Input<string> id, AudienceResolveProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new AudienceResolveProtocolMapper(name, id, state, options);
        }
    }

    public sealed class AudienceResolveProtocolMapperArgs : global::Pulumi.ResourceArgs
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
        /// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public AudienceResolveProtocolMapperArgs()
        {
        }
        public static new AudienceResolveProtocolMapperArgs Empty => new AudienceResolveProtocolMapperArgs();
    }

    public sealed class AudienceResolveProtocolMapperState : global::Pulumi.ResourceArgs
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
        /// The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public AudienceResolveProtocolMapperState()
        {
        }
        public static new AudienceResolveProtocolMapperState Empty => new AudienceResolveProtocolMapperState();
    }
}
