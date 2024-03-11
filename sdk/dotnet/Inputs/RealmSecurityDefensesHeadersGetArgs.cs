// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmSecurityDefensesHeadersGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("contentSecurityPolicy")]
        public Input<string>? ContentSecurityPolicy { get; set; }

        [Input("contentSecurityPolicyReportOnly")]
        public Input<string>? ContentSecurityPolicyReportOnly { get; set; }

        [Input("referrerPolicy")]
        public Input<string>? ReferrerPolicy { get; set; }

        [Input("strictTransportSecurity")]
        public Input<string>? StrictTransportSecurity { get; set; }

        [Input("xContentTypeOptions")]
        public Input<string>? XContentTypeOptions { get; set; }

        [Input("xFrameOptions")]
        public Input<string>? XFrameOptions { get; set; }

        [Input("xRobotsTag")]
        public Input<string>? XRobotsTag { get; set; }

        [Input("xXssProtection")]
        public Input<string>? XXssProtection { get; set; }

        public RealmSecurityDefensesHeadersGetArgs()
        {
        }
        public static new RealmSecurityDefensesHeadersGetArgs Empty => new RealmSecurityDefensesHeadersGetArgs();
    }
}
