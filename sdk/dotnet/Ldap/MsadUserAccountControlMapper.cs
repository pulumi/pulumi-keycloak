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
    /// Allows for creating and managing MSAD user account control mappers for Keycloak
    /// users federated via LDAP.
    /// 
    /// The MSAD (Microsoft Active Directory) user account control mapper is specific
    /// to LDAP user federation providers that are pulling from AD, and it can propagate
    /// AD user state to Keycloak in order to enforce settings like expired passwords
    /// or disabled accounts.
    /// 
    /// ## Example Usage
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
    ///             RealmName = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var ldapUserFederation = new Keycloak.Ldap.UserFederation("ldapUserFederation", new Keycloak.Ldap.UserFederationArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             UsernameLdapAttribute = "cn",
    ///             RdnLdapAttribute = "cn",
    ///             UuidLdapAttribute = "objectGUID",
    ///             UserObjectClasses = 
    ///             {
    ///                 "person",
    ///                 "organizationalPerson",
    ///                 "user",
    ///             },
    ///             ConnectionUrl = "ldap://my-ad-server",
    ///             UsersDn = "dc=example,dc=org",
    ///             BindDn = "cn=admin,dc=example,dc=org",
    ///             BindCredential = "admin",
    ///         });
    ///         var msadUserAccountControlMapper = new Keycloak.Ldap.MsadUserAccountControlMapper("msadUserAccountControlMapper", new Keycloak.Ldap.MsadUserAccountControlMapperArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             LdapUserFederationId = ldapUserFederation.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper msad_user_account_control_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper")]
    public partial class MsadUserAccountControlMapper : Pulumi.CustomResource
    {
        /// <summary>
        /// When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
        /// </summary>
        [Output("ldapPasswordPolicyHintsEnabled")]
        public Output<bool?> LdapPasswordPolicyHintsEnabled { get; private set; } = null!;

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
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a MsadUserAccountControlMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MsadUserAccountControlMapper(string name, MsadUserAccountControlMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper", name, args ?? new MsadUserAccountControlMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MsadUserAccountControlMapper(string name, Input<string> id, MsadUserAccountControlMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MsadUserAccountControlMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MsadUserAccountControlMapper Get(string name, Input<string> id, MsadUserAccountControlMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new MsadUserAccountControlMapper(name, id, state, options);
        }
    }

    public sealed class MsadUserAccountControlMapperArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
        /// </summary>
        [Input("ldapPasswordPolicyHintsEnabled")]
        public Input<bool>? LdapPasswordPolicyHintsEnabled { get; set; }

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
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public MsadUserAccountControlMapperArgs()
        {
        }
    }

    public sealed class MsadUserAccountControlMapperState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
        /// </summary>
        [Input("ldapPasswordPolicyHintsEnabled")]
        public Input<bool>? LdapPasswordPolicyHintsEnabled { get; set; }

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
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public MsadUserAccountControlMapperState()
        {
        }
    }
}
