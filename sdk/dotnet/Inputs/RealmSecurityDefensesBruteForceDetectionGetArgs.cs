// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmSecurityDefensesBruteForceDetectionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// When will failure count be reset?
        /// </summary>
        [Input("failureResetTimeSeconds")]
        public Input<int>? FailureResetTimeSeconds { get; set; }

        [Input("maxFailureWaitSeconds")]
        public Input<int>? MaxFailureWaitSeconds { get; set; }

        /// <summary>
        /// How many failures before wait is triggered.
        /// </summary>
        [Input("maxLoginFailures")]
        public Input<int>? MaxLoginFailures { get; set; }

        /// <summary>
        /// How long to wait after a quick login failure.
        /// - `max_failure_wait_seconds ` - (Optional) Max. time a user will be locked out.
        /// </summary>
        [Input("minimumQuickLoginWaitSeconds")]
        public Input<int>? MinimumQuickLoginWaitSeconds { get; set; }

        /// <summary>
        /// When `true`, this will lock the user permanently when the user exceeds the maximum login failures.
        /// </summary>
        [Input("permanentLockout")]
        public Input<bool>? PermanentLockout { get; set; }

        /// <summary>
        /// Configures the amount of time, in milliseconds, for consecutive failures to lock a user out.
        /// </summary>
        [Input("quickLoginCheckMilliSeconds")]
        public Input<int>? QuickLoginCheckMilliSeconds { get; set; }

        /// <summary>
        /// This represents the amount of time a user should be locked out when the login failure threshold has been met.
        /// </summary>
        [Input("waitIncrementSeconds")]
        public Input<int>? WaitIncrementSeconds { get; set; }

        public RealmSecurityDefensesBruteForceDetectionGetArgs()
        {
        }
        public static new RealmSecurityDefensesBruteForceDetectionGetArgs Empty => new RealmSecurityDefensesBruteForceDetectionGetArgs();
    }
}
