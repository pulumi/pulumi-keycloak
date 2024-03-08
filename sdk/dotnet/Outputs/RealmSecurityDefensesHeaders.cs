// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Outputs
{

    [OutputType]
    public sealed class RealmSecurityDefensesHeaders
    {
        public readonly string? ContentSecurityPolicy;
        public readonly string? ContentSecurityPolicyReportOnly;
        public readonly string? ReferrerPolicy;
        public readonly string? StrictTransportSecurity;
        public readonly string? XContentTypeOptions;
        public readonly string? XFrameOptions;
        public readonly string? XRobotsTag;
        public readonly string? XXssProtection;

        [OutputConstructor]
        private RealmSecurityDefensesHeaders(
            string? contentSecurityPolicy,

            string? contentSecurityPolicyReportOnly,

            string? referrerPolicy,

            string? strictTransportSecurity,

            string? xContentTypeOptions,

            string? xFrameOptions,

            string? xRobotsTag,

            string? xXssProtection)
        {
            ContentSecurityPolicy = contentSecurityPolicy;
            ContentSecurityPolicyReportOnly = contentSecurityPolicyReportOnly;
            ReferrerPolicy = referrerPolicy;
            StrictTransportSecurity = strictTransportSecurity;
            XContentTypeOptions = xContentTypeOptions;
            XFrameOptions = xFrameOptions;
            XRobotsTag = xRobotsTag;
            XXssProtection = xXssProtection;
        }
    }
}
