// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmSecurityDefensesBruteForceDetectionGetArgs : Pulumi.ResourceArgs
    {
        [Input("failureResetTimeSeconds")]
        public Input<int>? FailureResetTimeSeconds { get; set; }

        [Input("maxFailureWaitSeconds")]
        public Input<int>? MaxFailureWaitSeconds { get; set; }

        [Input("maxLoginFailures")]
        public Input<int>? MaxLoginFailures { get; set; }

        [Input("minimumQuickLoginWaitSeconds")]
        public Input<int>? MinimumQuickLoginWaitSeconds { get; set; }

        [Input("permanentLockout")]
        public Input<bool>? PermanentLockout { get; set; }

        [Input("quickLoginCheckMilliSeconds")]
        public Input<int>? QuickLoginCheckMilliSeconds { get; set; }

        [Input("waitIncrementSeconds")]
        public Input<int>? WaitIncrementSeconds { get; set; }

        public RealmSecurityDefensesBruteForceDetectionGetArgs()
        {
        }
    }
}
