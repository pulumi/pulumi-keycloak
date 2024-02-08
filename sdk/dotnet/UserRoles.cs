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
    /// Allows you to manage roles assigned to a Keycloak user.
    /// 
    /// If `exhaustive` is true, this resource attempts to be an **authoritative** source over user roles: roles that are manually added to the user will be removed, and roles that are manually removed from the
    /// user will be added upon the next run of `pulumi up`.
    /// If `exhaustive` is false, this resource is a partial assignation of roles to a user. As a result, you can use multiple `keycloak.UserRoles` for the same `user_id`.
    /// 
    /// Note that when assigning composite roles to a user, you may see a non-empty plan following a `pulumi up` if you assign
    /// a role and a composite that includes that role to the same user.
    /// 
    /// ## Example Usage
    /// ### Exhaustive Roles)
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
    ///     var realmRole = new Keycloak.Role("realmRole", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Description = "My Realm Role",
    ///     });
    /// 
    ///     var client = new Keycloak.OpenId.Client("client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "client",
    ///         Enabled = true,
    ///         AccessType = "BEARER-ONLY",
    ///     });
    /// 
    ///     var clientRole = new Keycloak.Role("clientRole", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = keycloak_client.Client.Id,
    ///         Description = "My Client Role",
    ///     });
    /// 
    ///     var user = new Keycloak.User("user", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Username = "bob",
    ///         Enabled = true,
    ///         Email = "bob@domain.com",
    ///         FirstName = "Bob",
    ///         LastName = "Bobson",
    ///     });
    /// 
    ///     var userRoles = new Keycloak.UserRoles("userRoles", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         UserId = user.Id,
    ///         RoleIds = new[]
    ///         {
    ///             realmRole.Id,
    ///             clientRole.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource can be imported using the format `{{realm_id}}/{{user_id}}`, where `user_id` is the unique ID that Keycloak
    /// 
    ///  assigns to the user upon creation. This value can be found in the GUI when editing the user, and is typically a GUID.
    /// 
    ///  Example:
    /// 
    ///  bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:index/userRoles:UserRoles user_roles my-realm/b0ae6924-1bd5-4655-9e38-dae7c5e42924
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/userRoles:UserRoles")]
    public partial class UserRoles : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
        /// </summary>
        [Output("exhaustive")]
        public Output<bool?> Exhaustive { get; private set; } = null!;

        /// <summary>
        /// The realm this user exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// A list of role IDs to map to the user
        /// </summary>
        [Output("roleIds")]
        public Output<ImmutableArray<string>> RoleIds { get; private set; } = null!;

        /// <summary>
        /// The ID of the user this resource should manage roles for.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserRoles resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserRoles(string name, UserRolesArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/userRoles:UserRoles", name, args ?? new UserRolesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserRoles(string name, Input<string> id, UserRolesState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/userRoles:UserRoles", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserRoles resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserRoles Get(string name, Input<string> id, UserRolesState? state = null, CustomResourceOptions? options = null)
        {
            return new UserRoles(name, id, state, options);
        }
    }

    public sealed class UserRolesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
        /// </summary>
        [Input("exhaustive")]
        public Input<bool>? Exhaustive { get; set; }

        /// <summary>
        /// The realm this user exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("roleIds", required: true)]
        private InputList<string>? _roleIds;

        /// <summary>
        /// A list of role IDs to map to the user
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        /// <summary>
        /// The ID of the user this resource should manage roles for.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserRolesArgs()
        {
        }
        public static new UserRolesArgs Empty => new UserRolesArgs();
    }

    public sealed class UserRolesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the user will be removed. Defaults to `true`.
        /// </summary>
        [Input("exhaustive")]
        public Input<bool>? Exhaustive { get; set; }

        /// <summary>
        /// The realm this user exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        [Input("roleIds")]
        private InputList<string>? _roleIds;

        /// <summary>
        /// A list of role IDs to map to the user
        /// </summary>
        public InputList<string> RoleIds
        {
            get => _roleIds ?? (_roleIds = new InputList<string>());
            set => _roleIds = value;
        }

        /// <summary>
        /// The ID of the user this resource should manage roles for.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public UserRolesState()
        {
        }
        public static new UserRolesState Empty => new UserRolesState();
    }
}
