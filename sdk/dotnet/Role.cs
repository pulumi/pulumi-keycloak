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
    /// Allows for creating and managing roles within Keycloak.
    /// 
    /// Roles allow you to define privileges within Keycloak and map them to users and groups.
    /// 
    /// ## Example Usage
    /// 
    /// ### Realm Role)
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
    ///         Attributes = 
    ///         {
    ///             { "key", "value" },
    ///             { "multivalue", "value1##value2" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Client Role)
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
    ///     var openidClient = new Keycloak.OpenId.Client("openid_client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "client",
    ///         Name = "client",
    ///         Enabled = true,
    ///         AccessType = "CONFIDENTIAL",
    ///         ValidRedirectUris = new[]
    ///         {
    ///             "http://localhost:8080/openid-callback",
    ///         },
    ///     });
    /// 
    ///     var clientRole = new Keycloak.Role("client_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = openidClientKeycloakClient.Id,
    ///         Name = "my-client-role",
    ///         Description = "My Client Role",
    ///         Attributes = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Composite Role)
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
    ///         Attributes = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     });
    /// 
    ///     var readRole = new Keycloak.Role("read_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "read",
    ///         Attributes = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     });
    /// 
    ///     var updateRole = new Keycloak.Role("update_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "update",
    ///         Attributes = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     });
    /// 
    ///     var deleteRole = new Keycloak.Role("delete_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "delete",
    ///         Attributes = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     });
    /// 
    ///     // client role
    ///     var openidClient = new Keycloak.OpenId.Client("openid_client", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "client",
    ///         Name = "client",
    ///         Enabled = true,
    ///         AccessType = "CONFIDENTIAL",
    ///         ValidRedirectUris = new[]
    ///         {
    ///             "http://localhost:8080/openid-callback",
    ///         },
    ///     });
    /// 
    ///     var clientRole = new Keycloak.Role("client_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = openidClientKeycloakClient.Id,
    ///         Name = "my-client-role",
    ///         Description = "My Client Role",
    ///         Attributes = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     });
    /// 
    ///     var adminRole = new Keycloak.Role("admin_role", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Name = "admin",
    ///         CompositeRoles = new[]
    ///         {
    ///             createRole.Id,
    ///             readRole.Id,
    ///             updateRole.Id,
    ///             deleteRole.Id,
    ///             clientRole.Id,
    ///         },
    ///         Attributes = 
    ///         {
    ///             { "key", "value" },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Roles can be imported using the format `{{realm_id}}/{{role_id}}`, where `role_id` is the unique ID that Keycloak assigns
    /// 
    /// to the role. The ID is not easy to find in the GUI, but it appears in the URL when editing the role.
    /// 
    /// Example:
    /// 
    /// bash
    /// 
    /// ```sh
    /// $ pulumi import keycloak:index/role:Role role my-realm/7e8cf32a-8acb-4d34-89c4-04fb1d10ccad
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/role:Role")]
    public partial class Role : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableDictionary<string, string>> Attributes { get; private set; } = null!;

        /// <summary>
        /// When specified, this role will be created as a client role attached to the client with the provided ID
        /// </summary>
        [Output("clientId")]
        public Output<string?> ClientId { get; private set; } = null!;

        /// <summary>
        /// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
        /// </summary>
        [Output("compositeRoles")]
        public Output<ImmutableArray<string>> CompositeRoles { get; private set; } = null!;

        /// <summary>
        /// The description of the role
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// When `true`, the role with the specified `name` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with roles that Keycloak creates automatically during realm creation, such as the client roles `create-client`, `view-realm`, ... for the client `realm-management` created per realm. Note, that the role will not be removed during destruction if `import` is `true`.
        /// </summary>
        [Output("import")]
        public Output<bool?> Import { get; private set; } = null!;

        /// <summary>
        /// The name of the role
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The realm this role exists within.
        /// </summary>
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

        /// <summary>
        /// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        /// </summary>
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        /// <summary>
        /// When specified, this role will be created as a client role attached to the client with the provided ID
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("compositeRoles")]
        private InputList<string>? _compositeRoles;

        /// <summary>
        /// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
        /// </summary>
        public InputList<string> CompositeRoles
        {
            get => _compositeRoles ?? (_compositeRoles = new InputList<string>());
            set => _compositeRoles = value;
        }

        /// <summary>
        /// The description of the role
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When `true`, the role with the specified `name` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with roles that Keycloak creates automatically during realm creation, such as the client roles `create-client`, `view-realm`, ... for the client `realm-management` created per realm. Note, that the role will not be removed during destruction if `import` is `true`.
        /// </summary>
        [Input("import")]
        public Input<bool>? Import { get; set; }

        /// <summary>
        /// The name of the role
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this role exists within.
        /// </summary>
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

        /// <summary>
        /// A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
        /// </summary>
        public InputMap<string> Attributes
        {
            get => _attributes ?? (_attributes = new InputMap<string>());
            set => _attributes = value;
        }

        /// <summary>
        /// When specified, this role will be created as a client role attached to the client with the provided ID
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        [Input("compositeRoles")]
        private InputList<string>? _compositeRoles;

        /// <summary>
        /// When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
        /// </summary>
        public InputList<string> CompositeRoles
        {
            get => _compositeRoles ?? (_compositeRoles = new InputList<string>());
            set => _compositeRoles = value;
        }

        /// <summary>
        /// The description of the role
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When `true`, the role with the specified `name` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with roles that Keycloak creates automatically during realm creation, such as the client roles `create-client`, `view-realm`, ... for the client `realm-management` created per realm. Note, that the role will not be removed during destruction if `import` is `true`.
        /// </summary>
        [Input("import")]
        public Input<bool>? Import { get; set; }

        /// <summary>
        /// The name of the role
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The realm this role exists within.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public RoleState()
        {
        }
        public static new RoleState Empty => new RoleState();
    }
}
