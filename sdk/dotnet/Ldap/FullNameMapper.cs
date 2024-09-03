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
    /// # keycloak.ldap.FullNameMapper
    /// 
    /// Allows for creating and managing full name mappers for Keycloak users federated
    /// via LDAP.
    /// 
    /// The LDAP full name mapper can map a user's full name from an LDAP attribute
    /// to the first and last name attributes of a Keycloak user.
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
    ///         RealmName = "test",
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
    ///     var ldapFullNameMapper = new Keycloak.Ldap.FullNameMapper("ldap_full_name_mapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         LdapUserFederationId = ldapUserFederation.Id,
    ///         Name = "full-name-mapper",
    ///         LdapFullNameAttribute = "cn",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm that this LDAP mapper will exist in.
    /// - `ldap_user_federation_id` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
    /// - `name` - (Required) Display name of this mapper when displayed in the console.
    /// - `ldap_full_name_attribute` - (Required) The name of the LDAP attribute containing the user's full name.
    /// - `read_only` - (Optional) When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
    /// - `write_only` - (Optional) When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
    /// 
    /// ### Import
    /// 
    /// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
    /// The ID of the LDAP user federation provider and the mapper can be found within
    /// the Keycloak GUI, and they are typically GUIDs:
    /// </summary>
    [KeycloakResourceType("keycloak:ldap/fullNameMapper:FullNameMapper")]
    public partial class FullNameMapper : global::Pulumi.CustomResource
    {
        [Output("ldapFullNameAttribute")]
        public Output<string> LdapFullNameAttribute { get; private set; } = null!;

        /// <summary>
        /// The ldap user federation provider to attach this mapper to.
        /// </summary>
        [Output("ldapUserFederationId")]
        public Output<string> LdapUserFederationId { get; private set; } = null!;

        /// <summary>
        /// Display name of the mapper when displayed in the console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("readOnly")]
        public Output<bool?> ReadOnly { get; private set; } = null!;

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        [Output("writeOnly")]
        public Output<bool?> WriteOnly { get; private set; } = null!;


        /// <summary>
        /// Create a FullNameMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FullNameMapper(string name, FullNameMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:ldap/fullNameMapper:FullNameMapper", name, args ?? new FullNameMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FullNameMapper(string name, Input<string> id, FullNameMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:ldap/fullNameMapper:FullNameMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing FullNameMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FullNameMapper Get(string name, Input<string> id, FullNameMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new FullNameMapper(name, id, state, options);
        }
    }

    public sealed class FullNameMapperArgs : global::Pulumi.ResourceArgs
    {
        [Input("ldapFullNameAttribute", required: true)]
        public Input<string> LdapFullNameAttribute { get; set; } = null!;

        /// <summary>
        /// The ldap user federation provider to attach this mapper to.
        /// </summary>
        [Input("ldapUserFederationId", required: true)]
        public Input<string> LdapUserFederationId { get; set; } = null!;

        /// <summary>
        /// Display name of the mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("writeOnly")]
        public Input<bool>? WriteOnly { get; set; }

        public FullNameMapperArgs()
        {
        }
        public static new FullNameMapperArgs Empty => new FullNameMapperArgs();
    }

    public sealed class FullNameMapperState : global::Pulumi.ResourceArgs
    {
        [Input("ldapFullNameAttribute")]
        public Input<string>? LdapFullNameAttribute { get; set; }

        /// <summary>
        /// The ldap user federation provider to attach this mapper to.
        /// </summary>
        [Input("ldapUserFederationId")]
        public Input<string>? LdapUserFederationId { get; set; }

        /// <summary>
        /// Display name of the mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("writeOnly")]
        public Input<bool>? WriteOnly { get; set; }

        public FullNameMapperState()
        {
        }
        public static new FullNameMapperState Empty => new FullNameMapperState();
    }
}
