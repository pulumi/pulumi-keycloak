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
    /// Allows for creating and managing Groups within Keycloak.
    /// 
    /// Groups provide a logical wrapping for users within Keycloak. Users within a group can share attributes and roles, and
    /// group membership can be mapped to a claim.
    /// 
    /// Attributes can also be defined on Groups.
    /// 
    /// Groups can also be federated from external data sources, such as LDAP or Active Directory. This resource **should not**
    /// be used to manage groups that were created this way.
    /// 
    /// ## Example Usage
    /// 
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
    ///     var parentGroup = new Keycloak.Group("parentGroup", new()
    ///     {
    ///         RealmId = realm.Id,
    ///     });
    /// 
    ///     var childGroup = new Keycloak.Group("childGroup", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ParentId = parentGroup.Id,
    ///     });
    /// 
    ///     var childGroupWithOptionalAttributes = new Keycloak.Group("childGroupWithOptionalAttributes", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ParentId = parentGroup.Id,
    ///         Attributes = 
    ///         {
    ///             { "foo", "bar" },
    ///             { "multivalue", "value1##value2" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Groups can be imported using the format `{{realm_id}}/{{group_id}}`, where `group_id` is the unique ID that Keycloak assigns to the group upon creation. This value can be found in the URI when editing this group in the GUI, and is typically a GUID. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:index/group:Group child_group my-realm/934a4a4e-28bd-4703-a0fa-332df153aabd
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/group:Group")]
    public partial class Group : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableDictionary<string, object>?> Attributes { get; private set; } = null!;

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of this group's parent. If omitted, this group will be defined at the root level.
        /// </summary>
        [Output("parentId")]
        public Output<string?> ParentId { get; private set; } = null!;

        /// <summary>
        /// (Computed) The complete path of the group. For example, the child group's path in the example configuration would be `/parent-group/child-group`.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/group:Group", name, args ?? new GroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;

        /// <summary>
        /// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        /// </summary>
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of this group's parent. If omitted, this group will be defined at the root level.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public GroupArgs()
        {
        }
        public static new GroupArgs Empty => new GroupArgs();
    }

    public sealed class GroupState : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<object>? _attributes;

        /// <summary>
        /// A map representing attributes for the group. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        /// </summary>
        public InputMap<object> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<object>());
            set => _attributes = value;
        }

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of this group's parent. If omitted, this group will be defined at the root level.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// (Computed) The complete path of the group. For example, the child group's path in the example configuration would be `/parent-group/child-group`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The realm this group exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public GroupState()
        {
        }
        public static new GroupState Empty => new GroupState();
    }
}
