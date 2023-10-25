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
    /// Allows for creating and managing hardcoded claim protocol mappers within Keycloak.
    /// 
    /// Hardcoded claim protocol mappers allow you to define a claim with a hardcoded value.
    /// 
    /// Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
    /// multiple different clients.
    /// 
    /// ## Example Usage
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
    ///     var openidClient = new Keycloak.OpenId.Client("openidClient", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "client",
    ///         Enabled = true,
    ///         AccessType = "CONFIDENTIAL",
    ///         ValidRedirectUris = new[]
    ///         {
    ///             "http://localhost:8080/openid-callback",
    ///         },
    ///     });
    /// 
    ///     var hardcodedClaimMapper = new Keycloak.OpenId.HardcodedClaimProtocolMapper("hardcodedClaimMapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = openidClient.Id,
    ///         ClaimName = "foo",
    ///         ClaimValue = "bar",
    ///     });
    /// 
    /// });
    /// ```
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
    ///     var clientScope = new Keycloak.OpenId.ClientScope("clientScope", new()
    ///     {
    ///         RealmId = realm.Id,
    ///     });
    /// 
    ///     var hardcodedClaimMapper = new Keycloak.OpenId.HardcodedClaimProtocolMapper("hardcodedClaimMapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientScopeId = clientScope.Id,
    ///         ClaimName = "foo",
    ///         ClaimValue = "bar",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper hardcoded_claim_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper hardcoded_claim_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper")]
    public partial class HardcodedClaimProtocolMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
        /// </summary>
        [Output("addToAccessToken")]
        public Output<bool?> AddToAccessToken { get; private set; } = null!;

        /// <summary>
        /// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
        /// </summary>
        [Output("addToIdToken")]
        public Output<bool?> AddToIdToken { get; private set; } = null!;

        /// <summary>
        /// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
        /// </summary>
        [Output("addToUserinfo")]
        public Output<bool?> AddToUserinfo { get; private set; } = null!;

        /// <summary>
        /// The name of the claim to insert into a token.
        /// </summary>
        [Output("claimName")]
        public Output<string> ClaimName { get; private set; } = null!;

        /// <summary>
        /// The hardcoded value of the claim.
        /// </summary>
        [Output("claimValue")]
        public Output<string> ClaimValue { get; private set; } = null!;

        /// <summary>
        /// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
        /// </summary>
        [Output("claimValueType")]
        public Output<string?> ClaimValueType { get; private set; } = null!;

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
        /// Create a HardcodedClaimProtocolMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HardcodedClaimProtocolMapper(string name, HardcodedClaimProtocolMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper", name, args ?? new HardcodedClaimProtocolMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HardcodedClaimProtocolMapper(string name, Input<string> id, HardcodedClaimProtocolMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:openid/hardcodedClaimProtocolMapper:HardcodedClaimProtocolMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HardcodedClaimProtocolMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HardcodedClaimProtocolMapper Get(string name, Input<string> id, HardcodedClaimProtocolMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new HardcodedClaimProtocolMapper(name, id, state, options);
        }
    }

    public sealed class HardcodedClaimProtocolMapperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
        /// </summary>
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        /// <summary>
        /// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
        /// </summary>
        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

        /// <summary>
        /// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
        /// </summary>
        [Input("addToUserinfo")]
        public Input<bool>? AddToUserinfo { get; set; }

        /// <summary>
        /// The name of the claim to insert into a token.
        /// </summary>
        [Input("claimName", required: true)]
        public Input<string> ClaimName { get; set; } = null!;

        /// <summary>
        /// The hardcoded value of the claim.
        /// </summary>
        [Input("claimValue", required: true)]
        public Input<string> ClaimValue { get; set; } = null!;

        /// <summary>
        /// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
        /// </summary>
        [Input("claimValueType")]
        public Input<string>? ClaimValueType { get; set; }

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
        /// The display name of this protocol mapper in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public HardcodedClaimProtocolMapperArgs()
        {
        }
        public static new HardcodedClaimProtocolMapperArgs Empty => new HardcodedClaimProtocolMapperArgs();
    }

    public sealed class HardcodedClaimProtocolMapperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the property should be added as a claim to the access token. Defaults to `true`.
        /// </summary>
        [Input("addToAccessToken")]
        public Input<bool>? AddToAccessToken { get; set; }

        /// <summary>
        /// Indicates if the property should be added as a claim to the id token. Defaults to `true`.
        /// </summary>
        [Input("addToIdToken")]
        public Input<bool>? AddToIdToken { get; set; }

        /// <summary>
        /// Indicates if the property should be added as a claim to the UserInfo response body. Defaults to `true`.
        /// </summary>
        [Input("addToUserinfo")]
        public Input<bool>? AddToUserinfo { get; set; }

        /// <summary>
        /// The name of the claim to insert into a token.
        /// </summary>
        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        /// <summary>
        /// The hardcoded value of the claim.
        /// </summary>
        [Input("claimValue")]
        public Input<string>? ClaimValue { get; set; }

        /// <summary>
        /// The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
        /// </summary>
        [Input("claimValueType")]
        public Input<string>? ClaimValueType { get; set; }

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
        /// The display name of this protocol mapper in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this protocol mapper exists within.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public HardcodedClaimProtocolMapperState()
        {
        }
        public static new HardcodedClaimProtocolMapperState Empty => new HardcodedClaimProtocolMapperState();
    }
}
