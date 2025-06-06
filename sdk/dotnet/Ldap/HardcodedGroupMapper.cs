// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Ldap
{
    /// <summary>
    /// Allows for creating and managing hardcoded group mappers for Keycloak users federated via LDAP.
    /// 
    /// The LDAP hardcoded group mapper will grant a specified Keycloak group to each Keycloak user linked with LDAP.
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
    ///     var realmGroup = new Keycloak.Group("realm_group", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "my-group",
    ///     });
    /// 
    ///     var assignGroupToUsers = new Keycloak.Ldap.HardcodedGroupMapper("assign_group_to_users", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         LdapUserFederationId = ldapUserFederation.Id,
    ///         Name = "assign-group-to-users",
    ///         Group = realmGroup.Name,
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
    /// $ pulumi import keycloak:ldap/hardcodedGroupMapper:HardcodedGroupMapper assign_group_to_users my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:ldap/hardcodedGroupMapper:HardcodedGroupMapper")]
    public partial class HardcodedGroupMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the group which should be assigned to the users.
        /// </summary>
        [Output("group")]
        public Output<string> Group { get; private set; } = null!;

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
        /// Create a HardcodedGroupMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HardcodedGroupMapper(string name, HardcodedGroupMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:ldap/hardcodedGroupMapper:HardcodedGroupMapper", name, args ?? new HardcodedGroupMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HardcodedGroupMapper(string name, Input<string> id, HardcodedGroupMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:ldap/hardcodedGroupMapper:HardcodedGroupMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HardcodedGroupMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HardcodedGroupMapper Get(string name, Input<string> id, HardcodedGroupMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new HardcodedGroupMapper(name, id, state, options);
        }
    }

    public sealed class HardcodedGroupMapperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the group which should be assigned to the users.
        /// </summary>
        [Input("group", required: true)]
        public Input<string> Group { get; set; } = null!;

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

        public HardcodedGroupMapperArgs()
        {
        }
        public static new HardcodedGroupMapperArgs Empty => new HardcodedGroupMapperArgs();
    }

    public sealed class HardcodedGroupMapperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the group which should be assigned to the users.
        /// </summary>
        [Input("group")]
        public Input<string>? Group { get; set; }

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

        public HardcodedGroupMapperState()
        {
        }
        public static new HardcodedGroupMapperState Empty => new HardcodedGroupMapperState();
    }
}
