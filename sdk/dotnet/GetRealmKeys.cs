// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
        /// ## # keycloak.getRealmKeys data source
        /// 
        /// Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
        /// 
        /// Remarks:
        /// 
        /// - A key must meet all filter criteria
        /// - This datasource may return more than one value.
        /// - If no key matches the filter criteria, then an error is returned.
        /// 
        /// ### Argument Reference
        /// 
        /// The following arguments are supported:
        /// 
        /// - `realm_id` - (Required) The realm of which the keys are retrieved.
        /// - `algorithms` - (Optional) When specified, keys are filtered by algorithm (values for algorithm: `HS256`, `RS256`,`AES`, ...)
        /// - `status` - (Optional) When specified, keys are filtered by status (values for status: `ACTIVE`, `DISABLED` and `PASSIVE`)
        /// </summary>
        public static Task<GetRealmKeysResult> InvokeAsync(GetRealmKeysArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetRealmKeysResult>("keycloak:index/getRealmKeys:getRealmKeys", args ?? new GetRealmKeysArgs(), options.WithVersion());
    }


    public sealed class GetRealmKeysArgs : Pulumi.InvokeArgs
    {
        [Input("algorithms")]
        private List<string>? _algorithms;
        public List<string> Algorithms
        {
            get => _algorithms ?? (_algorithms = new List<string>());
            set => _algorithms = value;
        }

        [Input("realmId", required: true)]
        public string RealmId { get; set; } = null!;

        [Input("statuses")]
        private List<string>? _statuses;
        public List<string> Statuses
        {
            get => _statuses ?? (_statuses = new List<string>());
            set => _statuses = value;
        }

        public GetRealmKeysArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetRealmKeysResult
    {
        public readonly ImmutableArray<string> Algorithms;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetRealmKeysKeyResult> Keys;
        public readonly string RealmId;
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
