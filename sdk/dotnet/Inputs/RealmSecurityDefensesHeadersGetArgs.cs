// *** WARNING: this file was generated by pulumi-language-dotnet. ***
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
        /// <summary>
        /// Sets the Content Security Policy, which can be used for prevent pages from being included by non-origin iframes. More information can be found in the [W3C-CSP](https://www.w3.org/TR/CSP/) Abstract.
        /// </summary>
        [Input("contentSecurityPolicy")]
        public Input<string>? ContentSecurityPolicy { get; set; }

        /// <summary>
        /// Used for testing Content Security Policies.
        /// </summary>
        [Input("contentSecurityPolicyReportOnly")]
        public Input<string>? ContentSecurityPolicyReportOnly { get; set; }

        /// <summary>
        /// The Referrer-Policy HTTP header controls how much referrer information (sent with the Referer header) should be included with requests.
        /// </summary>
        [Input("referrerPolicy")]
        public Input<string>? ReferrerPolicy { get; set; }

        /// <summary>
        /// The Script-Transport-Security HTTP header tells browsers to always use HTTPS.
        /// </summary>
        [Input("strictTransportSecurity")]
        public Input<string>? StrictTransportSecurity { get; set; }

        /// <summary>
        /// Sets the X-Content-Type-Options, which can be used for prevent MIME-sniffing a response away from the declared content-type
        /// </summary>
        [Input("xContentTypeOptions")]
        public Input<string>? XContentTypeOptions { get; set; }

        /// <summary>
        /// Sets the x-frame-option, which can be used to prevent pages from being included by non-origin iframes. More information can be found in the [RFC7034](https://tools.ietf.org/html/rfc7034)
        /// </summary>
        [Input("xFrameOptions")]
        public Input<string>? XFrameOptions { get; set; }

        /// <summary>
        /// Prevent pages from appearing in search engines.
        /// </summary>
        [Input("xRobotsTag")]
        public Input<string>? XRobotsTag { get; set; }

        /// <summary>
        /// This header configures the Cross-site scripting (XSS) filter in your browser.
        /// </summary>
        [Input("xXssProtection")]
        public Input<string>? XXssProtection { get; set; }

        public RealmSecurityDefensesHeadersGetArgs()
        {
        }
        public static new RealmSecurityDefensesHeadersGetArgs Empty => new RealmSecurityDefensesHeadersGetArgs();
    }
}
