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
    /// ## # keycloak.CustomUserFederation
    /// 
    /// Allows for creating and managing custom user federation providers within Keycloak.
    /// 
    /// A custom user federation provider is an implementation of Keycloak's
    /// [User Storage SPI](https://www.keycloak.org/docs/4.2/server_development/index.html#_user-storage-spi).
    /// An example of this implementation can be found here.
    /// 
    /// ### Example Usage
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
    ///         RealmName = "test",
    ///         Enabled = true,
    ///     });
    /// 
    ///     var customUserFederation = new Keycloak.CustomUserFederation("custom_user_federation", new()
    ///     {
    ///         Name = "custom",
    ///         RealmId = realm.Id,
    ///         ProviderId = "custom",
    ///         Enabled = true,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Argument Reference
    /// 
    /// The following arguments are supported:
    /// 
    /// - `realm_id` - (Required) The realm that this provider will provide user federation for.
    /// - `name` - (Required) Display name of the provider when displayed in the console.
    /// - `provider_id` - (Required) The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
    /// - `enabled` - (Optional) When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
    /// - `priority` - (Optional) Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
    /// - `cache_policy` - (Optional) Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
    /// 
    /// ### Import
    /// 
    /// Custom user federation providers can be imported using the format `{{realm_id}}/{{custom_user_federation_id}}`.
    /// The ID of the custom user federation provider can be found within the Keycloak GUI and is typically a GUID:
    /// </summary>
    [KeycloakResourceType("keycloak:index/customUserFederation:CustomUserFederation")]
    public partial class CustomUserFederation : global::Pulumi.CustomResource
    {
        [Output("cachePolicy")]
        public Output<string?> CachePolicy { get; private set; } = null!;

        /// <summary>
        /// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users
        /// sync.
        /// </summary>
        [Output("changedSyncPeriod")]
        public Output<int?> ChangedSyncPeriod { get; private set; } = null!;

        [Output("config")]
        public Output<ImmutableDictionary<string, string>?> Config { get; private set; } = null!;

        /// <summary>
        /// When false, this provider will not be used when performing queries for users.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
        /// </summary>
        [Output("fullSyncPeriod")]
        public Output<int?> FullSyncPeriod { get; private set; } = null!;

        /// <summary>
        /// Display name of the provider when displayed in the console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The parent_id of the generated component. will use realm_id if not specified.
        /// </summary>
        [Output("parentId")]
        public Output<string> ParentId { get; private set; } = null!;

        /// <summary>
        /// Priority of this provider when looking up users. Lower values are first.
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
        /// interface
        /// </summary>
        [Output("providerId")]
        public Output<string> ProviderId { get; private set; } = null!;

        /// <summary>
        /// The realm (name) this provider will provide user federation for.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;


        /// <summary>
        /// Create a CustomUserFederation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomUserFederation(string name, CustomUserFederationArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/customUserFederation:CustomUserFederation", name, args ?? new CustomUserFederationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomUserFederation(string name, Input<string> id, CustomUserFederationState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/customUserFederation:CustomUserFederation", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CustomUserFederation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomUserFederation Get(string name, Input<string> id, CustomUserFederationState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomUserFederation(name, id, state, options);
        }
    }

    public sealed class CustomUserFederationArgs : global::Pulumi.ResourceArgs
    {
        [Input("cachePolicy")]
        public Input<string>? CachePolicy { get; set; }

        /// <summary>
        /// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users
        /// sync.
        /// </summary>
        [Input("changedSyncPeriod")]
        public Input<int>? ChangedSyncPeriod { get; set; }

        [Input("config")]
        private InputMap<string>? _config;
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        /// <summary>
        /// When false, this provider will not be used when performing queries for users.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
        /// </summary>
        [Input("fullSyncPeriod")]
        public Input<int>? FullSyncPeriod { get; set; }

        /// <summary>
        /// Display name of the provider when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent_id of the generated component. will use realm_id if not specified.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// Priority of this provider when looking up users. Lower values are first.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
        /// interface
        /// </summary>
        [Input("providerId", required: true)]
        public Input<string> ProviderId { get; set; } = null!;

        /// <summary>
        /// The realm (name) this provider will provide user federation for.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        public CustomUserFederationArgs()
        {
        }
        public static new CustomUserFederationArgs Empty => new CustomUserFederationArgs();
    }

    public sealed class CustomUserFederationState : global::Pulumi.ResourceArgs
    {
        [Input("cachePolicy")]
        public Input<string>? CachePolicy { get; set; }

        /// <summary>
        /// How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users
        /// sync.
        /// </summary>
        [Input("changedSyncPeriod")]
        public Input<int>? ChangedSyncPeriod { get; set; }

        [Input("config")]
        private InputMap<string>? _config;
        public InputMap<string> Config
        {
            get => _config ?? (_config = new InputMap<string>());
            set => _config = value;
        }

        /// <summary>
        /// When false, this provider will not be used when performing queries for users.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
        /// </summary>
        [Input("fullSyncPeriod")]
        public Input<int>? FullSyncPeriod { get; set; }

        /// <summary>
        /// Display name of the provider when displayed in the console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The parent_id of the generated component. will use realm_id if not specified.
        /// </summary>
        [Input("parentId")]
        public Input<string>? ParentId { get; set; }

        /// <summary>
        /// Priority of this provider when looking up users. Lower values are first.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
        /// interface
        /// </summary>
        [Input("providerId")]
        public Input<string>? ProviderId { get; set; }

        /// <summary>
        /// The realm (name) this provider will provide user federation for.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        public CustomUserFederationState()
        {
        }
        public static new CustomUserFederationState Empty => new CustomUserFederationState();
    }
}
