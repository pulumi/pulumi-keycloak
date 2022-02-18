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
    /// Allows managing default realm roles within Keycloak.
    /// 
    /// Note: This feature was added in Keycloak v13, so this resource will not work on older versions of Keycloak.
    /// 
    /// ## Example Usage
    /// ### Realm Role)
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
    ///             Realm = "my-realm",
    ///             Enabled = true,
    ///         });
    ///         var defalutRoles = new Keycloak.DefaultRoles("defalutRoles", new Keycloak.DefaultRolesArgs
    ///         {
    ///             RealmId = realm.Id,
    ///             DefaultRoles = 
    ///             {
    ///                 "uma_authorization",
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Default roles can be imported using the format `{{realm_id}}/{{default_role_id}}`, where `default_role_id` is the unique ID of the composite role that Keycloak uses to control default realm level roles. The ID is not easy to find in the GUI, but it appears in the dev tools when editing the default roles. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:index/defaultRoles:DefaultRoles default_roles my-realm/a04c35c2-e95a-4dc5-bd32-e83a21be9e7d
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/defaultRoles:DefaultRoles")]
    public partial class DefaultRoles : Pulumi.CustomResource
    {
        /// <summary>
        /// Realm level roles assigned to new users by default.
        /// </summary>
        [Output("defaultRoles")]
        public Output<ImmutableArray<string>> RoleNames { get; private set; } = null!;

        /// <summary>
        /// The realm this role exists within.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a DefaultRoles resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DefaultRoles(string name, DefaultRolesArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/defaultRoles:DefaultRoles", name, args ?? new DefaultRolesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DefaultRoles(string name, Input<string> id, DefaultRolesState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/defaultRoles:DefaultRoles", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing DefaultRoles resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DefaultRoles Get(string name, Input<string> id, DefaultRolesState? state = null, CustomResourceOptions? options = null)
        {
            return new DefaultRoles(name, id, state, options);
        }
    }

    public sealed class DefaultRolesArgs : Pulumi.ResourceArgs
    {
        [Input("defaultRoles", required: true)]
        private InputList<string>? _defaultRoles;

        /// <summary>
        /// Realm level roles assigned to new users by default.
        /// </summary>
        public InputList<string> RoleNames
        {
            get => _defaultRoles ?? (_defaultRoles = new InputList<string>());
            set => _defaultRoles = value;
        }

        /// <summary>
        /// The realm this role exists within.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public DefaultRolesArgs()
        {
        }
    }

    public sealed class DefaultRolesState : Pulumi.ResourceArgs
    {
        [Input("defaultRoles")]
        private InputList<string>? _defaultRoles;

        /// <summary>
        /// Realm level roles assigned to new users by default.
        /// </summary>
        public InputList<string> RoleNames
        {
            get => _defaultRoles ?? (_defaultRoles = new InputList<string>());
            set => _defaultRoles = value;
        }

        /// <summary>
        /// The realm this role exists within.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public DefaultRolesState()
        {
        }
    }
}