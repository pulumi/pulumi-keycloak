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
    /// Allows for creating and managing `aes-generated` Realm keystores within Keycloak.
    /// 
    /// A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
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
    ///     });
    /// 
    ///     var keystoreAesGenerated = new Keycloak.RealmKeystoreAesGenerated("keystoreAesGenerated", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         Enabled = true,
    ///         Active = true,
    ///         Priority = 100,
    ///         SecretSize = 16,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Realm keys can be imported using realm name and keystore id, you can find it in web UI. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:index/realmKeystoreAesGenerated:RealmKeystoreAesGenerated keystore_aes_generated my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:index/realmKeystoreAesGenerated:RealmKeystoreAesGenerated")]
    public partial class RealmKeystoreAesGenerated : global::Pulumi.CustomResource
    {
        /// <summary>
        /// When `false`, key in not used for signing. Defaults to `true`.
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// When `false`, key is not accessible in this realm. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Display name of provider when linked in admin console.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Priority for the provider. Defaults to `0`
        /// </summary>
        [Output("priority")]
        public Output<int?> Priority { get; private set; } = null!;

        /// <summary>
        /// The realm this keystore exists in.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
        /// </summary>
        [Output("secretSize")]
        public Output<int?> SecretSize { get; private set; } = null!;


        /// <summary>
        /// Create a RealmKeystoreAesGenerated resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RealmKeystoreAesGenerated(string name, RealmKeystoreAesGeneratedArgs args, CustomResourceOptions? options = null)
            : base("keycloak:index/realmKeystoreAesGenerated:RealmKeystoreAesGenerated", name, args ?? new RealmKeystoreAesGeneratedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RealmKeystoreAesGenerated(string name, Input<string> id, RealmKeystoreAesGeneratedState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:index/realmKeystoreAesGenerated:RealmKeystoreAesGenerated", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RealmKeystoreAesGenerated resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RealmKeystoreAesGenerated Get(string name, Input<string> id, RealmKeystoreAesGeneratedState? state = null, CustomResourceOptions? options = null)
        {
            return new RealmKeystoreAesGenerated(name, id, state, options);
        }
    }

    public sealed class RealmKeystoreAesGeneratedArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `false`, key in not used for signing. Defaults to `true`.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// When `false`, key is not accessible in this realm. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Display name of provider when linked in admin console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority for the provider. Defaults to `0`
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The realm this keystore exists in.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
        /// </summary>
        [Input("secretSize")]
        public Input<int>? SecretSize { get; set; }

        public RealmKeystoreAesGeneratedArgs()
        {
        }
        public static new RealmKeystoreAesGeneratedArgs Empty => new RealmKeystoreAesGeneratedArgs();
    }

    public sealed class RealmKeystoreAesGeneratedState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When `false`, key in not used for signing. Defaults to `true`.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// When `false`, key is not accessible in this realm. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Display name of provider when linked in admin console.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Priority for the provider. Defaults to `0`
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// The realm this keystore exists in.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// Size in bytes for the generated AES Key. Size 16 is for AES-128, Size 24 for AES-192 and Size 32 for AES-256. WARN: Bigger keys then 128 bits are not allowed on some JDK implementations. Defaults to `16`.
        /// </summary>
        [Input("secretSize")]
        public Input<int>? SecretSize { get; set; }

        public RealmKeystoreAesGeneratedState()
        {
        }
        public static new RealmKeystoreAesGeneratedState Empty => new RealmKeystoreAesGeneratedState();
    }
}
