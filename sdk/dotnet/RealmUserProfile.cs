// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    /// <summary>
    /// Allows for managing Realm User Profiles within Keycloak.
    /// 
    /// A user profile defines a schema for representing user attributes and how they are managed within a realm.
    /// 
    /// Information for Keycloak versions &lt; 24:
    /// The realm linked to the `keycloak.RealmUserProfile` resource must have the user profile feature enabled.
    /// It can be done via the administration UI, or by setting the `userProfileEnabled` realm attribute to `true`.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var realm = new Keycloak.Realm("realm", new()
    ///     {
    ///         RealmName = "my-realm",
    ///     });
    /// 
    ///     var userprofile = new Keycloak.RealmUserProfile("userprofile", new()
    ///     {
    ///         RealmId = myRealm.Id,
    ///         UnmanagedAttributePolicy = "ENABLED",
    ///         Attributes = new[]
    ///         {
    ///             new Keycloak.Inputs.RealmUserProfileAttributeArgs
    ///             {
    ///                 Name = "field1",
    ///                 DisplayName = "Field 1",
    ///                 Group = "group1",
    ///                 MultiValued = false,
    ///                 EnabledWhenScopes = new[]
    ///                 {
    ///                     "offline_access",
    ///                 },
    ///                 RequiredForRoles = new[]
    ///                 {
    ///                     "user",
    ///                 },
    ///                 RequiredForScopes = new[]
    ///                 {
    ///                     "offline_access",
    ///                 },
    ///                 Permissions = new Keycloak.Inputs.RealmUserProfileAttributePermissionsArgs
    ///                 {
    ///                     Views = new[]
    ///                     {
    ///                         "admin",
    ///                         "user",
    ///                     },
    ///                     Edits = new[]
    ///                     {
    ///                         "admin",
    ///                         "user",
    ///                     },
    ///                 },
    ///                 Validators = new[]
    ///                 {
    ///                     new Keycloak.Inputs.RealmUserProfileAttributeValidatorArgs
    ///                     {
    ///                         Name = "person-name-prohibited-characters",
    ///                     },
    ///                     new Keycloak.Inputs.RealmUserProfileAttributeValidatorArgs
    ///                     {
    ///                         Name = "pattern",
    ///                         Config = 
    ///                         {
    ///                             { "pattern", "^[a-z]+$" },
    ///                             { "error-message", "Nope" },
    ///                         },
    ///                     },
    ///                 },
    ///                 Annotations = 
    ///                 {
    ///                     { "foo", "bar" },
    ///                 },
    ///             },
    ///             new Keycloak.Inputs.RealmUserProfileAttributeArgs
    ///             {
    ///                 Name = "field2",
    ///                 Validators = new[]
    ///                 {
    ///                     new Keycloak.Inputs.RealmUserProfileAttributeValidatorArgs
    ///                     {
    ///                         Name = "options",
    ///                         Config = 
    ///                         {
    ///                             { "options", JsonSerializer.Serialize(new[]
    ///                             {
    ///                                 "opt1",
    ///                             }) },
    ///                         },
    ///                     },
    ///                 },
    ///                 Annotations = 
    ///                 {
    ///                     { "foo", JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["key"] = "val",
    ///                     }) },
    ///                 },
    ///             },
    ///         },
    ///         Groups = new[]
    ///         {
    ///             new Keycloak.Inputs.RealmUserProfileGroupArgs
    ///             {
    ///                 Name = "group1",
    ///                 DisplayHeader = "Group 1",
    ///                 DisplayDescription = "A first group",
    ///                 Annotations = 
    ///                 {
    ///                     { "foo", "bar" },
    ///                     { "foo2", JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["key"] = "val",
    ///                     }) },
    ///                 },
    ///             },
    ///             new Keycloak.Inputs.RealmUserProfileGroupArgs
    ///             {
    ///                 Name = "group2",
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// This resource currently does not support importing.
    /// </summary>
    [KeycloakResourceType("keycloak:index/realmUserProfile:RealmUserProfile")]
    public partial class RealmUserProfile : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An ordered list of attributes.
        /// </summary>
        [Output("attributes")]
        public Output<ImmutableArray<Outputs.RealmUserProfileAttribute>> Attributes { get; private set; } = null!;

        /// <summary>
        /// A list of groups.
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<Outputs.RealmUserProfileGroup>> Groups { get; private set; } = null!;

        /// <summary>
        /// The ID of the realm the user profile applies to.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
        /// </summary>
        [Output("unmanagedAttributePolicy")]
        public Output<string?> UnmanagedAttributePolicy { get; private set; } = null!;


        /// <summary>
        /// Create a RealmUserProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RealmUserProfile(string name, RealmUserProfileArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/realmUserProfile:RealmUserProfile", name, args ?? new RealmUserProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RealmUserProfile(string name, Input<string> id, RealmUserProfileState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/realmUserProfile:RealmUserProfile", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RealmUserProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RealmUserProfile Get(string name, Input<string> id, RealmUserProfileState? state = null, CustomResourceOptions? options = null)
        {
            return new RealmUserProfile(name, id, state, options);
        }
    }

    public sealed class RealmUserProfileArgs : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<Inputs.RealmUserProfileAttributeArgs>? _attributes;

        /// <summary>
        /// An ordered list of attributes.
        /// </summary>
        public InputList<Inputs.RealmUserProfileAttributeArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.RealmUserProfileAttributeArgs>());
            set => _attributes = value;
        }

        [Input("groups")]
        private InputList<Inputs.RealmUserProfileGroupArgs>? _groups;

        /// <summary>
        /// A list of groups.
        /// </summary>
        public InputList<Inputs.RealmUserProfileGroupArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.RealmUserProfileGroupArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// The ID of the realm the user profile applies to.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
        /// </summary>
        [Input("unmanagedAttributePolicy")]
        public Input<string>? UnmanagedAttributePolicy { get; set; }

        public RealmUserProfileArgs()
        {
        }
        public static new RealmUserProfileArgs Empty => new RealmUserProfileArgs();
    }

    public sealed class RealmUserProfileState : global::Pulumi.ResourceArgs
    {
        [Input("attributes")]
        private InputList<Inputs.RealmUserProfileAttributeGetArgs>? _attributes;

        /// <summary>
        /// An ordered list of attributes.
        /// </summary>
        public InputList<Inputs.RealmUserProfileAttributeGetArgs> Attributes
        {
            get => _attributes ?? (_attributes = new InputList<Inputs.RealmUserProfileAttributeGetArgs>());
            set => _attributes = value;
        }

        [Input("groups")]
        private InputList<Inputs.RealmUserProfileGroupGetArgs>? _groups;

        /// <summary>
        /// A list of groups.
        /// </summary>
        public InputList<Inputs.RealmUserProfileGroupGetArgs> Groups
        {
            get => _groups ?? (_groups = new InputList<Inputs.RealmUserProfileGroupGetArgs>());
            set => _groups = value;
        }

        /// <summary>
        /// The ID of the realm the user profile applies to.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
        /// </summary>
        [Input("unmanagedAttributePolicy")]
        public Input<string>? UnmanagedAttributePolicy { get; set; }

        public RealmUserProfileState()
        {
        }
        public static new RealmUserProfileState Empty => new RealmUserProfileState();
    }
}
