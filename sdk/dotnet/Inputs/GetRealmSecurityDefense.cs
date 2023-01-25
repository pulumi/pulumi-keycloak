// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class GetRealmSecurityDefenseArgs : global::Pulumi.InvokeArgs
    {
        [Input("bruteForceDetections", required: true)]
        private List<Inputs.GetRealmSecurityDefenseBruteForceDetectionArgs>? _bruteForceDetections;
        public List<Inputs.GetRealmSecurityDefenseBruteForceDetectionArgs> BruteForceDetections
        {
            get => _bruteForceDetections ?? (_bruteForceDetections = new List<Inputs.GetRealmSecurityDefenseBruteForceDetectionArgs>());
            set => _bruteForceDetections = value;
        }

        [Input("headers", required: true)]
        private List<Inputs.GetRealmSecurityDefenseHeaderArgs>? _headers;
        public List<Inputs.GetRealmSecurityDefenseHeaderArgs> Headers
        {
            get => _headers ?? (_headers = new List<Inputs.GetRealmSecurityDefenseHeaderArgs>());
            set => _headers = value;
        }

        public GetRealmSecurityDefenseArgs()
        {
        }
        public static new GetRealmSecurityDefenseArgs Empty => new GetRealmSecurityDefenseArgs();
    }
}
