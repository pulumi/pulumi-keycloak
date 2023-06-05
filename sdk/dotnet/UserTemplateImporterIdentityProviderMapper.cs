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
    /// Allows for creating and managing an username template importer identity provider mapper within Keycloak.
    /// 
    /// The username template importer mapper can be used to map externally defined OIDC claims or SAML attributes with a template to the username of the imported Keycloak user:
    /// 
    /// - Substitutions are enclosed in \${}. For example: '\${ALIAS}.\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\&lt;NAME\&gt; references an ID or Access token claim.
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
    ///     var oidc = new Keycloak.Oidc.IdentityProvider("oidc", new()
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
    ///     var usernameImporter = new Keycloak.UserTemplateImporterIdentityProviderMapper("usernameImporter", new()
    ///     {
    ///         Realm = realm.Id,
    ///         IdentityProviderAlias = oidc.Alias,
    ///         Template = "${ALIAS}.${CLAIM.email}",
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
    /// Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper username_importer my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper")]
    public partial class UserTemplateImporterIdentityProviderMapper : global::Pulumi.CustomResource
    {
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
        /// Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\&lt;NAME\&gt; references an ID or Access token claim.
        /// </summary>
        [Output("template")]
        public Output<string?> Template { get; private set; } = null!;


        /// <summary>
        /// Create a UserTemplateImporterIdentityProviderMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserTemplateImporterIdentityProviderMapper(string name, UserTemplateImporterIdentityProviderMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper", name, args ?? new UserTemplateImporterIdentityProviderMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserTemplateImporterIdentityProviderMapper(string name, Input<string> id, UserTemplateImporterIdentityProviderMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserTemplateImporterIdentityProviderMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserTemplateImporterIdentityProviderMapper Get(string name, Input<string> id, UserTemplateImporterIdentityProviderMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new UserTemplateImporterIdentityProviderMapper(name, id, state, options);
        }
    }

    public sealed class UserTemplateImporterIdentityProviderMapperArgs : global::Pulumi.ResourceArgs
    {
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
        /// Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\&lt;NAME\&gt; references an ID or Access token claim.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public UserTemplateImporterIdentityProviderMapperArgs()
        {
        }
        public static new UserTemplateImporterIdentityProviderMapperArgs Empty => new UserTemplateImporterIdentityProviderMapperArgs();
    }

    public sealed class UserTemplateImporterIdentityProviderMapperState : global::Pulumi.ResourceArgs
    {
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
        /// Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\&lt;NAME\&gt; references an ID or Access token claim.
        /// </summary>
        [Input("template")]
        public Input<string>? Template { get; set; }

        public UserTemplateImporterIdentityProviderMapperState()
        {
        }
        public static new UserTemplateImporterIdentityProviderMapperState Empty => new UserTemplateImporterIdentityProviderMapperState();
    }
}
