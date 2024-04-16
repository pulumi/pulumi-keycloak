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
    /// ## # keycloak.ldap.UserAttributeMapper
    /// 
    /// Allows for creating and managing user attribute mappers for Keycloak users
    /// federated via LDAP.
    /// 
    /// The LDAP user attribute mapper can be used to map a single LDAP attribute
    /// to an attribute on the Keycloak user model.
    /// 
    /// ### Example Usage
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
    ///     var ldapUserAttributeMapper = new Keycloak.Ldap.UserAttributeMapper("ldap_user_attribute_mapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         LdapUserFederationId = ldapUserFederation.Id,
    ///         Name = "user-attribute-mapper",
    ///         UserModelAttribute = "foo",
    ///         LdapAttribute = "bar",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm that this LDAP mapper will exist in.
    /// - `ldap_user_federation_id` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
    /// - `name` - (Required) Display name of this mapper when displayed in the console.
    /// - `user_model_attribute` - (Required) Name of the user property or attribute you want to map the LDAP attribute into.
    /// - `ldap_attribute` - (Required) Name of the mapped attribute on the LDAP object.
    /// - `read_only` - (Optional) When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
    /// - `always_read_value_from_ldap` - (Optional) When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
    /// - `is_mandatory_in_ldap` - (Optional) When `true`, this attribute must exist in LDAP. Defaults to `false`.
    /// 
    /// ### Import
    /// 
    /// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
    /// The ID of the LDAP user federation provider and the mapper can be found within
    /// the Keycloak GUI, and they are typically GUIDs:
    /// </summary>
    [KeycloakResourceType("keycloak:ldap/userAttributeMapper:UserAttributeMapper")]
    public partial class UserAttributeMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When true, the value fetched from LDAP will override the value stored in Keycloak.
        /// </summary>
        [Output("alwaysReadValueFromLdap")]
        public Output<bool?> AlwaysReadValueFromLdap { get; private set; } = null!;

        /// <summary>
        /// Default value to set in LDAP if is_mandatory_in_ldap and the value is empty
        /// </summary>
        [Output("attributeDefaultValue")]
        public Output<string?> AttributeDefaultValue { get; private set; } = null!;

        /// <summary>
        /// Should be true for binary LDAP attributes
        /// </summary>
        [Output("isBinaryAttribute")]
        public Output<bool?> IsBinaryAttribute { get; private set; } = null!;

        /// <summary>
        /// When true, this attribute must exist in LDAP.
        /// </summary>
        [Output("isMandatoryInLdap")]
        public Output<bool?> IsMandatoryInLdap { get; private set; } = null!;

        /// <summary>
        /// Name of the mapped attribute on LDAP object.
        /// </summary>
        [Output("ldapAttribute")]
        public Output<string> LdapAttribute { get; private set; } = null!;

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

        /// <summary>
        /// When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
        /// </summary>
        [Output("readOnly")]
        public Output<bool?> ReadOnly { get; private set; } = null!;

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// Name of the UserModel property or attribute you want to map the LDAP attribute into.
        /// </summary>
        [Output("userModelAttribute")]
        public Output<string> UserModelAttribute { get; private set; } = null!;


        /// <summary>
        /// Create a UserAttributeMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserAttributeMapper(string name, UserAttributeMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:ldap/userAttributeMapper:UserAttributeMapper", name, args ?? new UserAttributeMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserAttributeMapper(string name, Input<string> id, UserAttributeMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:ldap/userAttributeMapper:UserAttributeMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserAttributeMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserAttributeMapper Get(string name, Input<string> id, UserAttributeMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new UserAttributeMapper(name, id, state, options);
        }
    }

    public sealed class UserAttributeMapperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, the value fetched from LDAP will override the value stored in Keycloak.
        /// </summary>
        [Input("alwaysReadValueFromLdap")]
        public Input<bool>? AlwaysReadValueFromLdap { get; set; }

        /// <summary>
        /// Default value to set in LDAP if is_mandatory_in_ldap and the value is empty
        /// </summary>
        [Input("attributeDefaultValue")]
        public Input<string>? AttributeDefaultValue { get; set; }

        /// <summary>
        /// Should be true for binary LDAP attributes
        /// </summary>
        [Input("isBinaryAttribute")]
        public Input<bool>? IsBinaryAttribute { get; set; }

        /// <summary>
        /// When true, this attribute must exist in LDAP.
        /// </summary>
        [Input("isMandatoryInLdap")]
        public Input<bool>? IsMandatoryInLdap { get; set; }

        /// <summary>
        /// Name of the mapped attribute on LDAP object.
        /// </summary>
        [Input("ldapAttribute", required: true)]
        public Input<string> LdapAttribute { get; set; } = null!;

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

        /// <summary>
        /// When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// Name of the UserModel property or attribute you want to map the LDAP attribute into.
        /// </summary>
        [Input("userModelAttribute", required: true)]
        public Input<string> UserModelAttribute { get; set; } = null!;

        public UserAttributeMapperArgs()
        {
        }
        public static new UserAttributeMapperArgs Empty => new UserAttributeMapperArgs();
    }

    public sealed class UserAttributeMapperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When true, the value fetched from LDAP will override the value stored in Keycloak.
        /// </summary>
        [Input("alwaysReadValueFromLdap")]
        public Input<bool>? AlwaysReadValueFromLdap { get; set; }

        /// <summary>
        /// Default value to set in LDAP if is_mandatory_in_ldap and the value is empty
        /// </summary>
        [Input("attributeDefaultValue")]
        public Input<string>? AttributeDefaultValue { get; set; }

        /// <summary>
        /// Should be true for binary LDAP attributes
        /// </summary>
        [Input("isBinaryAttribute")]
        public Input<bool>? IsBinaryAttribute { get; set; }

        /// <summary>
        /// When true, this attribute must exist in LDAP.
        /// </summary>
        [Input("isMandatoryInLdap")]
        public Input<bool>? IsMandatoryInLdap { get; set; }

        /// <summary>
        /// Name of the mapped attribute on LDAP object.
        /// </summary>
        [Input("ldapAttribute")]
        public Input<string>? LdapAttribute { get; set; }

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

        /// <summary>
        /// When true, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// Name of the UserModel property or attribute you want to map the LDAP attribute into.
        /// </summary>
        [Input("userModelAttribute")]
        public Input<string>? UserModelAttribute { get; set; }

        public UserAttributeMapperState()
        {
        }
        public static new UserAttributeMapperState Empty => new UserAttributeMapperState();
    }
}
