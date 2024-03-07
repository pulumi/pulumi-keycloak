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
    /// Allows for creating and managing user attribute mappers for Keycloak users
    /// federated via LDAP.
    /// 
    /// The LDAP user attribute mapper can be used to map a single LDAP attribute
    /// to an attribute on the Keycloak user model.
    /// 
    /// ## Example Usage
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
    ///         RealmName = "my-realm",
    ///         Enabled = true,
    ///     });
    /// 
    ///     var ldapUserFederation = new Keycloak.Ldap.UserFederation("ldapUserFederation", new()
    ///     {
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
    ///     var ldapUserAttributeMapper = new Keycloak.Ldap.UserAttributeMapper("ldapUserAttributeMapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         LdapUserFederationId = ldapUserFederation.Id,
    ///         UserModelAttribute = "foo",
    ///         LdapAttribute = "bar",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
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
    /// $ pulumi import keycloak:ldap/userAttributeMapper:UserAttributeMapper ldap_user_attribute_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:ldap/userAttributeMapper:UserAttributeMapper")]
    public partial class UserAttributeMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
        /// </summary>
        [Output("alwaysReadValueFromLdap")]
        public Output<bool?> AlwaysReadValueFromLdap { get; private set; } = null!;

        /// <summary>
        /// Default value to set in LDAP if `is_mandatory_in_ldap` is true and the value is empty.
        /// </summary>
        [Output("attributeDefaultValue")]
        public Output<string?> AttributeDefaultValue { get; private set; } = null!;

        /// <summary>
        /// Should be true for binary LDAP attributes.
        /// </summary>
        [Output("isBinaryAttribute")]
        public Output<bool?> IsBinaryAttribute { get; private set; } = null!;

        /// <summary>
        /// When `true`, this attribute must exist in LDAP. Defaults to `false`.
        /// </summary>
        [Output("isMandatoryInLdap")]
        public Output<bool?> IsMandatoryInLdap { get; private set; } = null!;

        /// <summary>
        /// Name of the mapped attribute on the LDAP object.
        /// </summary>
        [Output("ldapAttribute")]
        public Output<string> LdapAttribute { get; private set; } = null!;

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
        /// When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
        /// </summary>
        [Output("readOnly")]
        public Output<bool?> ReadOnly { get; private set; } = null!;

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// Name of the user property or attribute you want to map the LDAP attribute into.
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
        /// When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
        /// </summary>
        [Input("alwaysReadValueFromLdap")]
        public Input<bool>? AlwaysReadValueFromLdap { get; set; }

        /// <summary>
        /// Default value to set in LDAP if `is_mandatory_in_ldap` is true and the value is empty.
        /// </summary>
        [Input("attributeDefaultValue")]
        public Input<string>? AttributeDefaultValue { get; set; }

        /// <summary>
        /// Should be true for binary LDAP attributes.
        /// </summary>
        [Input("isBinaryAttribute")]
        public Input<bool>? IsBinaryAttribute { get; set; }

        /// <summary>
        /// When `true`, this attribute must exist in LDAP. Defaults to `false`.
        /// </summary>
        [Input("isMandatoryInLdap")]
        public Input<bool>? IsMandatoryInLdap { get; set; }

        /// <summary>
        /// Name of the mapped attribute on the LDAP object.
        /// </summary>
        [Input("ldapAttribute", required: true)]
        public Input<string> LdapAttribute { get; set; } = null!;

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
        /// When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// Name of the user property or attribute you want to map the LDAP attribute into.
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
        /// When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
        /// </summary>
        [Input("alwaysReadValueFromLdap")]
        public Input<bool>? AlwaysReadValueFromLdap { get; set; }

        /// <summary>
        /// Default value to set in LDAP if `is_mandatory_in_ldap` is true and the value is empty.
        /// </summary>
        [Input("attributeDefaultValue")]
        public Input<string>? AttributeDefaultValue { get; set; }

        /// <summary>
        /// Should be true for binary LDAP attributes.
        /// </summary>
        [Input("isBinaryAttribute")]
        public Input<bool>? IsBinaryAttribute { get; set; }

        /// <summary>
        /// When `true`, this attribute must exist in LDAP. Defaults to `false`.
        /// </summary>
        [Input("isMandatoryInLdap")]
        public Input<bool>? IsMandatoryInLdap { get; set; }

        /// <summary>
        /// Name of the mapped attribute on the LDAP object.
        /// </summary>
        [Input("ldapAttribute")]
        public Input<string>? LdapAttribute { get; set; }

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
        /// When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
        /// </summary>
        [Input("readOnly")]
        public Input<bool>? ReadOnly { get; set; }

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// Name of the user property or attribute you want to map the LDAP attribute into.
        /// </summary>
        [Input("userModelAttribute")]
        public Input<string>? UserModelAttribute { get; set; }

        public UserAttributeMapperState()
        {
        }
        public static new UserAttributeMapperState Empty => new UserAttributeMapperState();
    }
}
