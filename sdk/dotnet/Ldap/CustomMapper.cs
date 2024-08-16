// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Ldap
{
    /// <summary>
    /// Allows for creating and managing custom attribute mappers for Keycloak users federated via LDAP.
    /// 
    /// The LDAP custom mapper is implemented and deployed into Keycloak as a custom provider. This resource allows to
    /// specify the custom id and custom implementation class of the self-implemented attribute mapper as well as additional
    /// properties via config map.
    /// 
    /// The custom mapper should already be deployed into keycloak in order to be correctly configured.
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
    ///     var ldapUserFederation = new Keycloak.Ldap.UserFederation("ldap_user_federation", new()
    ///     {
    ///         Name = "openldap",
    ///         RealmId = realm.Id,
    ///         UsernameLdapAttribute = "cn",
    ///         RdnLdapAttribute = "cn",
    ///         UuidLdapAttribute = "entryDN",
    ///         UserObjectClasses = new[]
    ///         {
    ///             "simpleSecurityObject",
    ///             "organizationalRole",
    ///         },
    ///         ConnectionUrl = "ldap://openldap",
    ///         UsersDn = "dc=example,dc=org",
    ///         BindDn = "cn=admin,dc=example,dc=org",
    ///         BindCredential = "admin",
    ///     });
    /// 
    ///     var customMapper = new Keycloak.Ldap.CustomMapper("custom_mapper", new()
    ///     {
    ///         Name = "custom-mapper",
    ///         RealmId = openldap.RealmId,
    ///         LdapUserFederationId = openldap.Id,
    ///         ProviderId = "custom-provider-registered-in-keycloak",
    ///         ProviderType = "com.example.custom.ldap.mappers.CustomMapper",
    ///         Config = 
    ///         {
    ///             { "attribute.name", "name" },
    ///             { "attribute.value", "value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
    /// 
    /// The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.
    /// 
    /// Example:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:ldap/customMapper:CustomMapper custom_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:ldap/customMapper:CustomMapper")]
    public partial class CustomMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        /// </summary>
        [Output("config")]
        public Output<ImmutableDictionary<string, string>?> Config { get; private set; } = null!;

        /// <summary>
        /// The ID of the LDAP user federation provider to attach this mapper to.
        /// </summary>
        [Output("ldapUserFederationId")]
        public Output<string> LdapUserFederationId { get; private set; } = null!;

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The id of the LDAP mapper implemented in MapperFactory.
        /// </summary>
        [Output("providerId")]
        public Output<string> ProviderId { get; private set; } = null!;

        /// <summary>
        /// The fully-qualified Java class name of the custom LDAP mapper.
        /// </summary>
        [Output("providerType")]
        public Output<string> ProviderType { get; private set; } = null!;

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a CustomMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomMapper(string name, CustomMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:ldap/customMapper:CustomMapper", name, args ?? new CustomMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomMapper(string name, Input<string> id, CustomMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:ldap/customMapper:CustomMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomMapper Get(string name, Input<string> id, CustomMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomMapper(name, id, state, options);
        }
    }

    public sealed class CustomMapperArgs : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<string>? _config;

        /// <summary>
        /// A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        /// </summary>
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        /// <summary>
        /// The ID of the LDAP user federation provider to attach this mapper to.
        /// </summary>
        [Input("ldapUserFederationId", required: true)]
        public Input<string> LdapUserFederationId { get; set; } = null!;

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the LDAP mapper implemented in MapperFactory.
        /// </summary>
        [Input("providerId", required: true)]
        public Input<string> ProviderId { get; set; } = null!;

        /// <summary>
        /// The fully-qualified Java class name of the custom LDAP mapper.
        /// </summary>
        [Input("providerType", required: true)]
        public Input<string> ProviderType { get; set; } = null!;

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public CustomMapperArgs()
        {
        }
        public static new CustomMapperArgs Empty => new CustomMapperArgs();
    }

    public sealed class CustomMapperState : global::Pulumi.ResourceArgs
    {
        [Input("config")]
        private InputMap<string>? _config;

        /// <summary>
        /// A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
        /// </summary>
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        /// <summary>
        /// The ID of the LDAP user federation provider to attach this mapper to.
        /// </summary>
        [Input("ldapUserFederationId")]
        public Input<string>? LdapUserFederationId { get; set; }

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the LDAP mapper implemented in MapperFactory.
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The fully-qualified Java class name of the custom LDAP mapper.
        /// </summary>
        [Input("providerType")]
        public Input<string>? ProviderType { get; set; }

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public CustomMapperState()
        {
        }
        public static new CustomMapperState Empty => new CustomMapperState();
    }
}
