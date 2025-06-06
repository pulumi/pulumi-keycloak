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
    public sealed class RealmSmtpServer
    {
        /// <summary>
        /// Enables authentication to the SMTP server.  This block supports the following arguments:
        /// </summary>
        public readonly Outputs.RealmSmtpServerAuth? Auth;
        /// <summary>
        /// The email address uses for bounces.
        /// </summary>
        public readonly string? EnvelopeFrom;
        /// <summary>
        /// The email address for the sender.
        /// </summary>
        public readonly string From;
        /// <summary>
        /// The display name of the sender email address.
        /// </summary>
        public readonly string? FromDisplayName;
        /// <summary>
        /// The host of the SMTP server.
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// The port of the SMTP server (defaults to 25).
        /// </summary>
        public readonly string? Port;
        /// <summary>
        /// The "reply to" email address.
        /// </summary>
        public readonly string? ReplyTo;
        /// <summary>
        /// The display name of the "reply to" email address.
        /// </summary>
        public readonly string? ReplyToDisplayName;
        /// <summary>
        /// When `true`, enables SSL. Defaults to `false`.
        /// </summary>
        public readonly bool? Ssl;
        /// <summary>
        /// When `true`, enables StartTLS. Defaults to `false`.
        /// </summary>
        public readonly bool? Starttls;

        [OutputConstructor]
        private RealmSmtpServer(
            Outputs.RealmSmtpServerAuth? auth,

            string? envelopeFrom,

            string from,

            string? fromDisplayName,

            string host,

            string? port,

            string? replyTo,

            string? replyToDisplayName,

            bool? ssl,

            bool? starttls)
        {
            Auth = auth;
            EnvelopeFrom = envelopeFrom;
            From = from;
            FromDisplayName = fromDisplayName;
            Host = host;
            Port = port;
            ReplyTo = replyTo;
            ReplyToDisplayName = replyToDisplayName;
            Ssl = ssl;
            Starttls = starttls;
        }
    }
}
