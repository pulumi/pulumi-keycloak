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
    /// ## # keycloak.ldap.HardcodedRoleMapper
    /// 
    /// This mapper will grant a specified Keycloak role to each Keycloak user linked with LDAP.
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm that this LDAP mapper will exist in.
    /// - `ldap_user_federation_id` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
    /// - `name` - (Required) Display name of this mapper when displayed in the console.
    /// - `role` - (Required) The role which should be assigned to the users.
    /// </summary>
    public partial class HardcodedRoleMapper : Pulumi.CustomResource
    {
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
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// Role to grant to user.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;


        /// <summary>
        /// Create a HardcodedRoleMapper resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public HardcodedRoleMapper(string name, HardcodedRoleMapperArgs args, CustomResourceOptions? options = null)
            : base("keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper", name, args ?? new HardcodedRoleMapperArgs(), MakeResourceOptions(options, ""))
        {
        }

        private HardcodedRoleMapper(string name, Input<string> id, HardcodedRoleMapperState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing HardcodedRoleMapper resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static HardcodedRoleMapper Get(string name, Input<string> id, HardcodedRoleMapperState? state = null, CustomResourceOptions? options = null)
        {
            return new HardcodedRoleMapper(name, id, state, options);
        }
    }

    public sealed class HardcodedRoleMapperArgs : Pulumi.ResourceArgs
    {
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
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// Role to grant to user.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        public HardcodedRoleMapperArgs()
        {
        }
    }

    public sealed class HardcodedRoleMapperState : Pulumi.ResourceArgs
    {
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
        /// The realm in which the ldap user federation provider exists.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// Role to grant to user.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        public HardcodedRoleMapperState()
        {
        }
    }
}
