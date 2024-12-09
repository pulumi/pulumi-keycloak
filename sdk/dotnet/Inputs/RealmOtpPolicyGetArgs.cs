// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmOtpPolicyGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// What hashing algorithm should be used to generate the OTP, Valid options are `HmacSHA1`,`HmacSHA256` and `HmacSHA512`. Defaults to `HmacSHA1`.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// How many digits the OTP have. Defaults to `6`.
        /// </summary>
        [Input("digits")]
        public Input<int>? Digits { get; set; }

        /// <summary>
        /// What should the initial counter value be. Defaults to `2`.
        /// </summary>
        [Input("initialCounter")]
        public Input<int>? InitialCounter { get; set; }

        /// <summary>
        /// How far ahead should the server look just in case the token generator and server are out of time sync or counter sync. Defaults to `1`.
        /// </summary>
        [Input("lookAheadWindow")]
        public Input<int>? LookAheadWindow { get; set; }

        /// <summary>
        /// How many seconds should an OTP token be valid. Defaults to `30`.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// One Time Password Type, supported Values are `totp` for Time-Based One Time Password and `hotp` for Counter Based. Defaults to `totp`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public RealmOtpPolicyGetArgs()
        {
        }
        public static new RealmOtpPolicyGetArgs Empty => new RealmOtpPolicyGetArgs();
    }
}
