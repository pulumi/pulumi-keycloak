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
    /// # keycloak.Role
    /// 
    /// Allows for creating and managing roles within Keycloak.
    /// 
    /// Roles allow you define privileges within Keycloak and map them to users
    /// and groups.
    /// 
    /// ### Example Usage (Realm role)
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
    ///     var realmRole = new Keycloak.Role("realm_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "my-realm-role",
    ///         Description = "My Realm Role",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example Usage (Client role)
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
    ///     var client = new Keycloak.OpenId.Client("client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "client",
    ///         Name = "client",
    ///         Enabled = true,
    ///         AccessType = "BEARER-ONLY",
    ///     });
    /// 
    ///     var clientRole = new Keycloak.Role("client_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = clientKeycloakClient.Id,
    ///         Name = "my-client-role",
    ///         Description = "My Client Role",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Example Usage (Composite role)
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
    ///     // realm roles
    ///     var createRole = new Keycloak.Role("create_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "create",
    ///     });
    /// 
    ///     var readRole = new Keycloak.Role("read_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "read",
    ///     });
    /// 
    ///     var updateRole = new Keycloak.Role("update_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "update",
    ///     });
    /// 
    ///     var deleteRole = new Keycloak.Role("delete_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "delete",
    ///     });
    /// 
    ///     // client role
    ///     var client = new Keycloak.OpenId.Client("client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "client",
    ///         Name = "client",
    ///         Enabled = true,
    ///         AccessType = "BEARER-ONLY",
    ///     });
    /// 
    ///     var clientRole = new Keycloak.Role("client_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = clientKeycloakClient.Id,
    ///         Name = "my-client-role",
    ///         Description = "My Client Role",
    ///     });
    /// 
    ///     var adminRole = new Keycloak.Role("admin_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "admin",
    ///         CompositeRoles = new[]
    ///         {
    ///             "{keycloak_role.create_role.id}",
    ///             "{keycloak_role.read_role.id}",
    ///             "{keycloak_role.update_role.id}",
    ///             "{keycloak_role.delete_role.id}",
    ///             "{keycloak_role.client_role.id}",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm this role exists within.
    /// - `client_id` - (Optional) When specified, this role will be created as
    ///   a client role attached to the client with the provided ID
    /// - `name` - (Required) The name of the role
    /// - `description` - (Optional) The description of the role
    /// - `composite_roles` - (Optional) When specified, this role will be a
    ///   composite role, composed of all roles that have an ID present within
    ///   this list.
    /// 
    /// ### Import
    /// 
    /// Roles can be imported using the format `{{realm_id}}/{{role_id}}`, where
    /// `role_id` is the unique ID that Keycloak assigns to the role. The ID is
    /// not easy to find in the GUI, but it appears in the URL when editing the
    /// role.
    /// 
    /// Example:
    /// </summary>
    [KeycloakResourceType("keycloak:index/role:Role")]
    public partial class Role : global::Pulumi.CustomResource
    {
        [Output("attributes")]
        public Output<ImmutableDictionary<string, string>?> Attributes { get; private set; } = null!;

        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        [Output("compositeRoles")]
        public Output<ImmutableArray<string>> CompositeRoles { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a Role resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Role(string name, RoleArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/role:Role", name, args ?? new RoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Role(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/role:Role", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Role resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Role Get(string name, Input<string> id, RoleState? state = null, CustomResourceOptions? options = null)
        {
            return new Role(name, id, state, options);
        }
    }

    public sealed class RoleArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("compositeRoles")]
        private InputList<string>? _compositeRoles;
        public InputList<string> CompositeRoles
        {
            get => _compositeRoles ?? (_compositeRoles = new InputList<string>());
            set => _compositeRoles = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public RoleArgs()
        {
        }
        public static new RoleArgs Empty => new RoleArgs();
    }

    public sealed class RoleState : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputMap<string>? _attributes;
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("compositeRoles")]
        private InputList<string>? _compositeRoles;
        public InputList<string> CompositeRoles
        {
            get => _compositeRoles ?? (_compositeRoles = new InputList<string>());
            set => _compositeRoles = value;
        }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public RoleState()
        {
        }
        public static new RoleState Empty => new RoleState();
    }
}
