// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class GetRealmSecurityDefenseBruteForceDetectionInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("failureResetTimeSeconds", required: true)]
        public Input<int> FailureResetTimeSeconds { get; set; } = null!;

        [Input("maxFailureWaitSeconds", required: true)]
        public Input<int> MaxFailureWaitSeconds { get; set; } = null!;

        [Input("maxLoginFailures", required: true)]
        public Input<int> MaxLoginFailures { get; set; } = null!;

        [Input("minimumQuickLoginWaitSeconds", required: true)]
        public Input<int> MinimumQuickLoginWaitSeconds { get; set; } = null!;

        [Input("permanentLockout", required: true)]
        public Input<bool> PermanentLockout { get; set; } = null!;

        [Input("quickLoginCheckMilliSeconds", required: true)]
        public Input<int> QuickLoginCheckMilliSeconds { get; set; } = null!;

        [Input("waitIncrementSeconds", required: true)]
        public Input<int> WaitIncrementSeconds { get; set; } = null!;

        public GetRealmSecurityDefenseBruteForceDetectionInputArgs()
        {
        }
        public static new GetRealmSecurityDefenseBruteForceDetectionInputArgs Empty => new GetRealmSecurityDefenseBruteForceDetectionInputArgs();
    }
}
