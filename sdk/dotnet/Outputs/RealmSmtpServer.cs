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
    public sealed class RealmSmtpServer
    {
        public readonly Outputs.RealmSmtpServerAuth? Auth;
        public readonly string? EnvelopeFrom;
        public readonly string From;
        public readonly string? FromDisplayName;
        public readonly string Host;
        public readonly string? Port;
        public readonly string? ReplyTo;
        public readonly string? ReplyToDisplayName;
        public readonly bool? Ssl;
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
