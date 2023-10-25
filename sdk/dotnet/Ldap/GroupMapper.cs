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
    /// Allows for creating and managing group mappers for Keycloak users federated via LDAP.
    /// 
    /// The LDAP group mapper can be used to map an LDAP user's groups from some DN to Keycloak groups. This group mapper will also
    /// create the groups within Keycloak if they do not already exist.
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
    ///     var ldapGroupMapper = new Keycloak.Ldap.GroupMapper("ldapGroupMapper", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         LdapUserFederationId = ldapUserFederation.Id,
    ///         LdapGroupsDn = "dc=example,dc=org",
    ///         GroupNameLdapAttribute = "cn",
    ///         GroupObjectClasses = new[]
    ///         {
    ///             "groupOfNames",
    ///         },
    ///         MembershipAttributeType = "DN",
    ///         MembershipLdapAttribute = "member",
    ///         MembershipUserLdapAttribute = "cn",
    ///         MemberofLdapAttribute = "memberOf",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:ldap/groupMapper:GroupMapper ldap_group_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:ldap/groupMapper:GroupMapper")]
    public partial class GroupMapper : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
        /// </summary>
        [Output("dropNonExistingGroupsDuringSync")]
        public Output<bool?> DropNonExistingGroupsDuringSync { get; private set; } = null!;

        /// <summary>
        /// The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
        /// </summary>
        [Output("groupNameLdapAttribute")]
        public Output<string> GroupNameLdapAttribute { get; private set; } = null!;

        /// <summary>
        /// List of strings representing the object classes for the group. Must contain at least one.
        /// </summary>
        [Output("groupObjectClasses")]
        public Output<ImmutableArray<string>> GroupObjectClasses { get; private set; } = null!;

        /// <summary>
        /// When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
        /// </summary>
        [Output("groupsLdapFilter")]
        public Output<string?> GroupsLdapFilter { get; private set; } = null!;

        /// <summary>
        /// Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
        /// </summary>
        [Output("groupsPath")]
        public Output<string> GroupsPath { get; private set; } = null!;

        /// <summary>
        /// When `true`, missing groups in the hierarchy will be ignored.
        /// </summary>
        [Output("ignoreMissingGroups")]
        public Output<bool?> IgnoreMissingGroups { get; private set; } = null!;

        /// <summary>
        /// The LDAP DN where groups can be found.
        /// </summary>
        [Output("ldapGroupsDn")]
        public Output<string> LdapGroupsDn { get; private set; } = null!;

        /// <summary>
        /// The ID of the LDAP user federation provider to attach this mapper to.
        /// </summary>
        [Output("ldapUserFederationId")]
        public Output<string> LdapUserFederationId { get; private set; } = null!;

        /// <summary>
        /// Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
        /// </summary>
        [Output("mappedGroupAttributes")]
        public Output<ImmutableArray<string>> MappedGroupAttributes { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
        /// </summary>
        [Output("memberofLdapAttribute")]
        public Output<string?> MemberofLdapAttribute { get; private set; } = null!;

        /// <summary>
        /// Can be one of `DN` or `UID`. Defaults to `DN`.
        /// </summary>
        [Output("membershipAttributeType")]
        public Output<string?> MembershipAttributeType { get; private set; } = null!;

        /// <summary>
        /// The name of the LDAP attribute that is used for membership mappings.
        /// </summary>
        [Output("membershipLdapAttribute")]
        public Output<string> MembershipLdapAttribute { get; private set; } = null!;

        /// <summary>
        /// The name of the LDAP attribute on a user that is used for membership mappings.
        /// </summary>
        [Output("membershipUserLdapAttribute")]
        public Output<string> MembershipUserLdapAttribute { get; private set; } = null!;

        /// <summary>
        /// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
        /// </summary>
        [Output("preserveGroupInheritance")]
        public Output<bool?> PreserveGroupInheritance { get; private set; } = null!;

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
        /// </summary>
        [Output("userRolesRetrieveStrategy")]
        public Output<string?> UserRolesRetrieveStrategy { get; private set; } = null!;


        /// <summary>
        /// Create a GroupMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupMapper(string name, GroupMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:ldap/groupMapper:GroupMapper", name, args ?? new GroupMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupMapper(string name, Input<string> id, GroupMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:ldap/groupMapper:GroupMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupMapper Get(string name, Input<string> id, GroupMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupMapper(name, id, state, options);
        }
    }

    public sealed class GroupMapperArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
        /// </summary>
        [Input("dropNonExistingGroupsDuringSync")]
        public Input<bool>? DropNonExistingGroupsDuringSync { get; set; }

        /// <summary>
        /// The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
        /// </summary>
        [Input("groupNameLdapAttribute", required: true)]
        public Input<string> GroupNameLdapAttribute { get; set; } = null!;

        [Input("groupObjectClasses", required: true)]
        private InputList<string>? _groupObjectClasses;

        /// <summary>
        /// List of strings representing the object classes for the group. Must contain at least one.
        /// </summary>
        public InputList<string> GroupObjectClasses
        {
            get => _groupObjectClasses ?? (_groupObjectClasses = new InputList<string>());
            set => _groupObjectClasses = value;
        }

        /// <summary>
        /// When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
        /// </summary>
        [Input("groupsLdapFilter")]
        public Input<string>? GroupsLdapFilter { get; set; }

        /// <summary>
        /// Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
        /// </summary>
        [Input("groupsPath")]
        public Input<string>? GroupsPath { get; set; }

        /// <summary>
        /// When `true`, missing groups in the hierarchy will be ignored.
        /// </summary>
        [Input("ignoreMissingGroups")]
        public Input<bool>? IgnoreMissingGroups { get; set; }

        /// <summary>
        /// The LDAP DN where groups can be found.
        /// </summary>
        [Input("ldapGroupsDn", required: true)]
        public Input<string> LdapGroupsDn { get; set; } = null!;

        /// <summary>
        /// The ID of the LDAP user federation provider to attach this mapper to.
        /// </summary>
        [Input("ldapUserFederationId", required: true)]
        public Input<string> LdapUserFederationId { get; set; } = null!;

        [Input("mappedGroupAttributes")]
        private InputList<string>? _mappedGroupAttributes;

        /// <summary>
        /// Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
        /// </summary>
        public InputList<string> MappedGroupAttributes
        {
            get => _mappedGroupAttributes ?? (_mappedGroupAttributes = new InputList<string>());
            set => _mappedGroupAttributes = value;
        }

        /// <summary>
        /// Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
        /// </summary>
        [Input("memberofLdapAttribute")]
        public Input<string>? MemberofLdapAttribute { get; set; }

        /// <summary>
        /// Can be one of `DN` or `UID`. Defaults to `DN`.
        /// </summary>
        [Input("membershipAttributeType")]
        public Input<string>? MembershipAttributeType { get; set; }

        /// <summary>
        /// The name of the LDAP attribute that is used for membership mappings.
        /// </summary>
        [Input("membershipLdapAttribute", required: true)]
        public Input<string> MembershipLdapAttribute { get; set; } = null!;

        /// <summary>
        /// The name of the LDAP attribute on a user that is used for membership mappings.
        /// </summary>
        [Input("membershipUserLdapAttribute", required: true)]
        public Input<string> MembershipUserLdapAttribute { get; set; } = null!;

        /// <summary>
        /// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
        /// </summary>
        [Input("preserveGroupInheritance")]
        public Input<bool>? PreserveGroupInheritance { get; set; }

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
        /// </summary>
        [Input("userRolesRetrieveStrategy")]
        public Input<string>? UserRolesRetrieveStrategy { get; set; }

        public GroupMapperArgs()
        {
        }
        public static new GroupMapperArgs Empty => new GroupMapperArgs();
    }

    public sealed class GroupMapperState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `true`, groups that no longer exist within LDAP will be dropped in Keycloak during sync. Defaults to `false`.
        /// </summary>
        [Input("dropNonExistingGroupsDuringSync")]
        public Input<bool>? DropNonExistingGroupsDuringSync { get; set; }

        /// <summary>
        /// The name of the LDAP attribute that is used in group objects for the name and RDN of the group. Typically `cn`.
        /// </summary>
        [Input("groupNameLdapAttribute")]
        public Input<string>? GroupNameLdapAttribute { get; set; }

        [Input("groupObjectClasses")]
        private InputList<string>? _groupObjectClasses;

        /// <summary>
        /// List of strings representing the object classes for the group. Must contain at least one.
        /// </summary>
        public InputList<string> GroupObjectClasses
        {
            get => _groupObjectClasses ?? (_groupObjectClasses = new InputList<string>());
            set => _groupObjectClasses = value;
        }

        /// <summary>
        /// When specified, adds an additional custom filter to be used when querying for groups. Must start with `(` and end with `)`.
        /// </summary>
        [Input("groupsLdapFilter")]
        public Input<string>? GroupsLdapFilter { get; set; }

        /// <summary>
        /// Keycloak group path the LDAP groups are added to. For example if value `/Applications/App1` is used, then LDAP groups will be available in Keycloak under group `App1`, which is the child of top level group `Applications`. The configured group path must already exist in Keycloak when creating this mapper.
        /// </summary>
        [Input("groupsPath")]
        public Input<string>? GroupsPath { get; set; }

        /// <summary>
        /// When `true`, missing groups in the hierarchy will be ignored.
        /// </summary>
        [Input("ignoreMissingGroups")]
        public Input<bool>? IgnoreMissingGroups { get; set; }

        /// <summary>
        /// The LDAP DN where groups can be found.
        /// </summary>
        [Input("ldapGroupsDn")]
        public Input<string>? LdapGroupsDn { get; set; }

        /// <summary>
        /// The ID of the LDAP user federation provider to attach this mapper to.
        /// </summary>
        [Input("ldapUserFederationId")]
        public Input<string>? LdapUserFederationId { get; set; }

        [Input("mappedGroupAttributes")]
        private InputList<string>? _mappedGroupAttributes;

        /// <summary>
        /// Array of strings representing attributes on the LDAP group which will be mapped to attributes on the Keycloak group.
        /// </summary>
        public InputList<string> MappedGroupAttributes
        {
            get => _mappedGroupAttributes ?? (_mappedGroupAttributes = new InputList<string>());
            set => _mappedGroupAttributes = value;
        }

        /// <summary>
        /// Specifies the name of the LDAP attribute on the LDAP user that contains the groups the user is a member of. Defaults to `memberOf`.
        /// </summary>
        [Input("memberofLdapAttribute")]
        public Input<string>? MemberofLdapAttribute { get; set; }

        /// <summary>
        /// Can be one of `DN` or `UID`. Defaults to `DN`.
        /// </summary>
        [Input("membershipAttributeType")]
        public Input<string>? MembershipAttributeType { get; set; }

        /// <summary>
        /// The name of the LDAP attribute that is used for membership mappings.
        /// </summary>
        [Input("membershipLdapAttribute")]
        public Input<string>? MembershipLdapAttribute { get; set; }

        /// <summary>
        /// The name of the LDAP attribute on a user that is used for membership mappings.
        /// </summary>
        [Input("membershipUserLdapAttribute")]
        public Input<string>? MembershipUserLdapAttribute { get; set; }

        /// <summary>
        /// Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// Display name of this mapper when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// When `true`, group inheritance will be propagated from LDAP to Keycloak. When `false`, all LDAP groups will be propagated as top level groups within Keycloak.
        /// </summary>
        [Input("preserveGroupInheritance")]
        public Input<bool>? PreserveGroupInheritance { get; set; }

        /// <summary>
        /// The realm that this LDAP mapper will exist in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// Can be one of `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`, `GET_GROUPS_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_GROUPS_BY_MEMBER_ATTRIBUTE`.
        /// </summary>
        [Input("userRolesRetrieveStrategy")]
        public Input<string>? UserRolesRetrieveStrategy { get; set; }

        public GroupMapperState()
        {
        }
        public static new GroupMapperState Empty => new GroupMapperState();
    }
}
