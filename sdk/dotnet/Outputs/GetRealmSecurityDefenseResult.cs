// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Outputs
{

    [OutputType]
    public sealed class GetRealmSecurityDefenseResult
    {
        public readonly ImmutableArray<Outputs.GetRealmSecurityDefenseBruteForceDetectionResult> BruteForceDetections;
        public readonly ImmutableArray<Outputs.GetRealmSecurityDefenseHeaderResult> Headers;

        [OutputConstructor]
        private GetRealmSecurityDefenseResult(
            ImmutableArray<Outputs.GetRealmSecurityDefenseBruteForceDetectionResult> bruteForceDetections,

            ImmutableArray<Outputs.GetRealmSecurityDefenseHeaderResult> headers)
        {
            BruteForceDetections = bruteForceDetections;
            Headers = headers;
        }
    }
}
