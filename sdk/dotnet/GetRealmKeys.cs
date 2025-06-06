// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak
{
    public static class GetRealmKeys
    {
        /// <summary>
        /// Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
        /// 
        /// Remarks:
        /// 
        /// - A key must meet all filter criteria
        /// - This data source may return more than one value.
        /// - If no key matches the filter criteria, then an error will be returned.
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
        ///         Enabled = true,
        ///     });
        /// 
        ///     var realmKeys = Keycloak.GetRealmKeys.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         Algorithms = new[]
        ///         {
        ///             "AES",
        ///             "RS256",
        ///         },
        ///         Statuses = new[]
        ///         {
        ///             "ACTIVE",
        ///             "PASSIVE",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["certificate"] = realmKeys.Apply(getRealmKeysResult =&gt; getRealmKeysResult.Keys[0]?.Certificate),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetRealmKeysResult> InvokeAsync(GetRealmKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRealmKeysResult>("keycloak:index/getRealmKeys:getRealmKeys", args ?? new GetRealmKeysArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
        /// 
        /// Remarks:
        /// 
        /// - A key must meet all filter criteria
        /// - This data source may return more than one value.
        /// - If no key matches the filter criteria, then an error will be returned.
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
        ///         Enabled = true,
        ///     });
        /// 
        ///     var realmKeys = Keycloak.GetRealmKeys.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         Algorithms = new[]
        ///         {
        ///             "AES",
        ///             "RS256",
        ///         },
        ///         Statuses = new[]
        ///         {
        ///             "ACTIVE",
        ///             "PASSIVE",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["certificate"] = realmKeys.Apply(getRealmKeysResult =&gt; getRealmKeysResult.Keys[0]?.Certificate),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetRealmKeysResult> Invoke(GetRealmKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRealmKeysResult>("keycloak:index/getRealmKeys:getRealmKeys", args ?? new GetRealmKeysInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
        /// 
        /// Remarks:
        /// 
        /// - A key must meet all filter criteria
        /// - This data source may return more than one value.
        /// - If no key matches the filter criteria, then an error will be returned.
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
        ///         Enabled = true,
        ///     });
        /// 
        ///     var realmKeys = Keycloak.GetRealmKeys.Invoke(new()
        ///     {
        ///         RealmId = realm.Id,
        ///         Algorithms = new[]
        ///         {
        ///             "AES",
        ///             "RS256",
        ///         },
        ///         Statuses = new[]
        ///         {
        ///             "ACTIVE",
        ///             "PASSIVE",
        ///         },
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["certificate"] = realmKeys.Apply(getRealmKeysResult =&gt; getRealmKeysResult.Keys[0]?.Certificate),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetRealmKeysResult> Invoke(GetRealmKeysInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRealmKeysResult>("keycloak:index/getRealmKeys:getRealmKeys", args ?? new GetRealmKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRealmKeysArgs : global::Pulumi.InvokeArgs
    {
        [Input("algorithms")]
        private List<string>? _algorithms;

        /// <summary>
        /// When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
        /// </summary>
        public List<string> Algorithms
        {
            get => _algorithms ?? (_algorithms = new List<string>());
            set => _algorithms = value;
        }

        /// <summary>
        /// The realm from which the keys will be retrieved.
        /// </summary>
        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        [Input("statuses")]
        private List<string>? _statuses;

        /// <summary>
        /// When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
        /// </summary>
        public List<string> Statuses
        {
            get => _statuses ?? (_statuses = new List<string>());
            set => _statuses = value;
        }

        public GetRealmKeysArgs()
        {
        }
        public static new GetRealmKeysArgs Empty => new GetRealmKeysArgs();
    }

    public sealed class GetRealmKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("algorithms")]
        private InputList<string>? _algorithms;

        /// <summary>
        /// When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
        /// </summary>
        public InputList<string> Algorithms
        {
            get => _algorithms ?? (_algorithms = new InputList<string>());
            set => _algorithms = value;
        }

        /// <summary>
        /// The realm from which the keys will be retrieved.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        [Input("statuses")]
        private InputList<string>? _statuses;

        /// <summary>
        /// When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
        /// </summary>
        public InputList<string> Statuses
        {
            get => _statuses ?? (_statuses = new InputList<string>());
            set => _statuses = value;
        }

        public GetRealmKeysInvokeArgs()
        {
        }
        public static new GetRealmKeysInvokeArgs Empty => new GetRealmKeysInvokeArgs();
    }


    [OutputType]
    public sealed class GetRealmKeysResult
    {
        public readonly ImmutableArray<string> Algorithms;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// (Computed) A list of keys that match the filter criteria. Each key has the following attributes:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRealmKeysKeyResult> Keys;
        public readonly string RealmId;
        /// <summary>
        /// Key status (string)
        /// </summary>
        public readonly ImmutableArray<string> Statuses;

        [OutputConstructor]
        private GetRealmKeysResult(
            ImmutableArray<string> algorithms,

            string id,

            ImmutableArray<Outputs.GetRealmKeysKeyResult> keys,

            string realmId,

            ImmutableArray<string> statuses)
        {
            Algorithms = algorithms;
            Id = id;
            Keys = keys;
            RealmId = realmId;
            Statuses = statuses;
        }
    }
}
