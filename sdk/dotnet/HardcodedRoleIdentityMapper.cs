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
    /// Allows for creating and managing hardcoded role mappers for Keycloak identity provider.
    /// 
    /// The identity provider hardcoded role mapper grants a specified Keycloak role to each Keycloak user from the LDAP provider.
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
    ///         Alias = "my-idp",
    ///         AuthorizationUrl = "https://authorizationurl.com",
    ///         ClientId = "clientID",
    ///         ClientSecret = "clientSecret",
    ///         TokenUrl = "https://tokenurl.com",
    ///     });
    /// 
    ///     var realmRole = new Keycloak.Role("realmRole", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Description = "My Realm Role",
    ///     });
    /// 
    ///     var oidcHardcodedRoleIdentityMapper = new Keycloak.HardcodedRoleIdentityMapper("oidcHardcodedRoleIdentityMapper", new()
    ///     {
    ///         Realm = realm.Id,
    ///         IdentityProviderAlias = oidcIdentityProvider.Alias,
    ///         Role = "my-realm-role",
    ///         ExtraConfig = 
    ///         {
    ///             { "syncMode", "INHERIT" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper")]
    public partial class HardcodedRoleIdentityMapper : global::Pulumi.CustomResource
    {
        [Output("extraConfig")]
        public Output<ImmutableDictionary<string, object>?> ExtraConfig { get; private set; } = null!;

        /// <summary>
        /// The IDP alias of the attribute to set.
        /// </summary>
        [Output("identityProviderAlias")]
        public Output<string> IdentityProviderAlias { get; private set; } = null!;

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm ID that this mapper will exist in.
        /// </summary>
        [Output("realm")]
        public Output<string> Realm { get; private set; } = null!;

        /// <summary>
        /// The name of the role which should be assigned to the users.
        /// </summary>
        [Output("role")]
        public Output<string?> Role { get; private set; } = null!;


        /// <summary>
        /// Create a HardcodedRoleIdentityMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HardcodedRoleIdentityMapper(string name, HardcodedRoleIdentityMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, args ?? new HardcodedRoleIdentityMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HardcodedRoleIdentityMapper(string name, Input<string> id, HardcodedRoleIdentityMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HardcodedRoleIdentityMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HardcodedRoleIdentityMapper Get(string name, Input<string> id, HardcodedRoleIdentityMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new HardcodedRoleIdentityMapper(name, id, state, options);
        }
    }

    public sealed class HardcodedRoleIdentityMapperArgs : global::Pulumi.ResourceArgs
    {
        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// The IDP alias of the attribute to set.
        /// </summary>
        [Input("identityProviderAlias", required: true)]
        public Input<string> IdentityProviderAlias { get; set; } = null!;

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm ID that this mapper will exist in.
        /// </summary>
        [Input("realm", required: true)]
        public Input<string> Realm { get; set; } = null!;

        /// <summary>
        /// The name of the role which should be assigned to the users.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public HardcodedRoleIdentityMapperArgs()
        {
        }
        public static new HardcodedRoleIdentityMapperArgs Empty => new HardcodedRoleIdentityMapperArgs();
    }

    public sealed class HardcodedRoleIdentityMapperState : global::Pulumi.ResourceArgs
    {
        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// The IDP alias of the attribute to set.
        /// </summary>
        [Input("identityProviderAlias")]
        public Input<string>? IdentityProviderAlias { get; set; }

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm ID that this mapper will exist in.
        /// </summary>
        [Input("realm")]
        public Input<string>? Realm { get; set; }

        /// <summary>
        /// The name of the role which should be assigned to the users.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public HardcodedRoleIdentityMapperState()
        {
        }
        public static new HardcodedRoleIdentityMapperState Empty => new HardcodedRoleIdentityMapperState();
    }
}
