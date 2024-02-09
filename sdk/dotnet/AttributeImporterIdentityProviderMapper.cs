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
    /// Allows for creating and managing an attribute importer identity provider mapper within Keycloak.
    /// 
    /// The attribute importer mapper can be used to map attributes from externally defined users to attributes or properties of the imported Keycloak user:
    /// - For the OIDC identity provider, this will map a claim on the ID or access token to an attribute for the imported Keycloak user.
    /// - For the SAML identity provider, this will map a SAML attribute found within the assertion to an attribute for the imported Keycloak user.
    /// - For social identity providers, this will map a JSON field from the user profile to an attribute for the imported Keycloak user.
    /// 
    /// &gt; If you are using Keycloak 10 or higher, you will need to specify the `extra_config` argument in order to define a `syncMode` for the mapper.
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
    ///     var oidcIdentityProvider = new Keycloak.Oidc.IdentityProvider("oidcIdentityProvider", new()
    ///     {
    ///         Realm = realm.Id,
    ///         Alias = "oidc",
    ///         AuthorizationUrl = "https://example.com/auth",
    ///         TokenUrl = "https://example.com/token",
    ///         ClientId = "example_id",
    ///         ClientSecret = "example_token",
    ///         DefaultScopes = "openid random profile",
    ///     });
    /// 
    ///     var oidcAttributeImporterIdentityProviderMapper = new Keycloak.AttributeImporterIdentityProviderMapper("oidcAttributeImporterIdentityProviderMapper", new()
    ///     {
    ///         Realm = realm.Id,
    ///         ClaimName = "my-email-claim",
    ///         IdentityProviderAlias = oidcIdentityProvider.Alias,
    ///         UserAttribute = "email",
    ///         ExtraConfig = 
    ///         {
    ///             { "syncMode", "INHERIT" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak
    /// 
    ///  assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID.
    /// 
    ///  Example:
    /// 
    ///  bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper")]
    public partial class AttributeImporterIdentityProviderMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attribute_name`.
        /// </summary>
        [Output("attributeFriendlyName")]
        public Output<string?> AttributeFriendlyName { get; private set; } = null!;

        /// <summary>
        /// For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attribute_friendly_name`.
        /// </summary>
        [Output("attributeName")]
        public Output<string?> AttributeName { get; private set; } = null!;

        /// <summary>
        /// For OIDC based providers, this is the name of the claim to use.
        /// </summary>
        [Output("claimName")]
        public Output<string?> ClaimName { get; private set; } = null!;

        /// <summary>
        /// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        /// </summary>
        [Output("extraConfig")]
        public Output<ImmutableDictionary<string, object>?> ExtraConfig { get; private set; } = null!;

        /// <summary>
        /// The alias of the associated identity provider.
        /// </summary>
        [Output("identityProviderAlias")]
        public Output<string> IdentityProviderAlias { get; private set; } = null!;

        /// <summary>
        /// The name of the mapper.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of the realm.
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// The user attribute or property name to store the mapped result.
        /// </summary>
        [Output("userAttribute")]
        public Output<string> UserAttribute { get; private set; } = null!;


        /// <summary>
        /// Create a AttributeImporterIdentityProviderMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AttributeImporterIdentityProviderMapper(string name, AttributeImporterIdentityProviderMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, args ?? new AttributeImporterIdentityProviderMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AttributeImporterIdentityProviderMapper(string name, Input<string> id, AttributeImporterIdentityProviderMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AttributeImporterIdentityProviderMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AttributeImporterIdentityProviderMapper Get(string name, Input<string> id, AttributeImporterIdentityProviderMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new AttributeImporterIdentityProviderMapper(name, id, state, options);
        }
    }

    public sealed class AttributeImporterIdentityProviderMapperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attribute_name`.
        /// </summary>
        [Input("attributeFriendlyName")]
        public Input<string>? AttributeFriendlyName { get; set; }

        /// <summary>
        /// For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attribute_friendly_name`.
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// For OIDC based providers, this is the name of the claim to use.
        /// </summary>
        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;

        /// <summary>
        /// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        /// </summary>
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// The alias of the associated identity provider.
        /// </summary>
        [Input("identityProviderAlias", required: true)]
        public Input<string> IdentityProviderAlias { get; set; } = null!;

        /// <summary>
        /// The name of the mapper.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the realm.
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// The user attribute or property name to store the mapped result.
        /// </summary>
        [Input("userAttribute", required: true)]
        public Input<string> UserAttribute { get; set; } = null!;

        public AttributeImporterIdentityProviderMapperArgs()
        {
        }
        public static new AttributeImporterIdentityProviderMapperArgs Empty => new AttributeImporterIdentityProviderMapperArgs();
    }

    public sealed class AttributeImporterIdentityProviderMapperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attribute_name`.
        /// </summary>
        [Input("attributeFriendlyName")]
        public Input<string>? AttributeFriendlyName { get; set; }

        /// <summary>
        /// For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attribute_friendly_name`.
        /// </summary>
        [Input("attributeName")]
        public Input<string>? AttributeName { get; set; }

        /// <summary>
        /// For OIDC based providers, this is the name of the claim to use.
        /// </summary>
        [Input("claimName")]
        public Input<string>? ClaimName { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;

        /// <summary>
        /// Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
        /// </summary>
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// The alias of the associated identity provider.
        /// </summary>
        [Input("identityProviderAlias")]
        public Input<string>? IdentityProviderAlias { get; set; }

        /// <summary>
        /// The name of the mapper.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of the realm.
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// The user attribute or property name to store the mapped result.
        /// </summary>
        [Input("userAttribute")]
        public Input<string>? UserAttribute { get; set; }

        public AttributeImporterIdentityProviderMapperState()
        {
        }
        public static new AttributeImporterIdentityProviderMapperState Empty => new AttributeImporterIdentityProviderMapperState();
    }
}
