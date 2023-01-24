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
    /// ## Example Usage
    /// ### Exhaustive Groups)
    /// ```csharp
    /// using System.Collections.Generic;
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
    ///     var @group = new Keycloak.Group("group", new()
    ///     {
    ///         RealmId = realm.Id,
    ///     });
    /// 
    ///     var user = new Keycloak.User("user", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Username = "my-user",
    ///     });
    /// 
    ///     var userGroups = new Keycloak.UserGroups("userGroups", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         UserId = user.Id,
    ///         GroupIds = new[]
    ///         {
    ///             @group.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ### Non Exhaustive Groups)
    /// ```csharp
    /// using System.Collections.Generic;
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
    ///     var groupFoo = new Keycloak.Group("groupFoo", new()
    ///     {
    ///         RealmId = realm.Id,
    ///     });
    /// 
    ///     var groupBar = new Keycloak.Group("groupBar", new()
    ///     {
    ///         RealmId = realm.Id,
    ///     });
    /// 
    ///     var user = new Keycloak.User("user", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Username = "my-user",
    ///     });
    /// 
    ///     var userGroupsAssociation1UserGroups = new Keycloak.UserGroups("userGroupsAssociation1UserGroups", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         UserId = user.Id,
    ///         Exhaustive = false,
    ///         GroupIds = new[]
    ///         {
    ///             groupFoo.Id,
    ///         },
    ///     });
    /// 
    ///     var userGroupsAssociation1Index_userGroupsUserGroups = new Keycloak.UserGroups("userGroupsAssociation1Index/userGroupsUserGroups", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         UserId = user.Id,
    ///         Exhaustive = false,
    ///         GroupIds = new[]
    ///         {
    ///             groupBar.Id,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource does not support import. Instead of importing, feel free to create this resource as if it did not already exist on the server.
    /// </summary>
    [KeycloakResourceType("keycloak:index/userGroups:UserGroups")]
    public partial class UserGroups : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
        /// </summary>
        [Output("exhaustive")]
        public Output<bool?> Exhaustive { get; private set; } = null!;

        /// <summary>
        /// A list of group IDs that the user is member of.
        /// </summary>
        [Output("groupIds")]
        public Output<ImmutableArray<string>> GroupIds { get; private set; } = null!;

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// The ID of the user this resource should manage groups for.
        /// </summary>
        [Output("userId")]
        public Output<string> UserId { get; private set; } = null!;


        /// <summary>
        /// Create a UserGroups resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UserGroups(string name, UserGroupsArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/userGroups:UserGroups", name, args ?? new UserGroupsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UserGroups(string name, Input<string> id, UserGroupsState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/userGroups:UserGroups", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UserGroups resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UserGroups Get(string name, Input<string> id, UserGroupsState? state = null, CustomResourceOptions? options = null)
        {
            return new UserGroups(name, id, state, options);
        }
    }

    public sealed class UserGroupsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
        /// </summary>
        [Input("exhaustive")]
        public Input<bool>? Exhaustive { get; set; }

        [Input("groupIds", required: true)]
        private InputList<string>? _groupIds;

        /// <summary>
        /// A list of group IDs that the user is member of.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// The ID of the user this resource should manage groups for.
        /// </summary>
        [Input("userId", required: true)]
        public Input<string> UserId { get; set; } = null!;

        public UserGroupsArgs()
        {
        }
        public static new UserGroupsArgs Empty => new UserGroupsArgs();
    }

    public sealed class UserGroupsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Indicates if the list of the user's groups is exhaustive. In this case, groups that are manually added to the user will be removed. Defaults to `true`.
        /// </summary>
        [Input("exhaustive")]
        public Input<bool>? Exhaustive { get; set; }

        [Input("groupIds")]
        private InputList<string>? _groupIds;

        /// <summary>
        /// A list of group IDs that the user is member of.
        /// </summary>
        public InputList<string> GroupIds
        {
            get => _groupIds ?? (_groupIds = new InputList<string>());
            set => _groupIds = value;
        }

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// The ID of the user this resource should manage groups for.
        /// </summary>
        [Input("userId")]
        public Input<string>? UserId { get; set; }

        public UserGroupsState()
        {
        }
        public static new UserGroupsState Empty => new UserGroupsState();
    }
}
