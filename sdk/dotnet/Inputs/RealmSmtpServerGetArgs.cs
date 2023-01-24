// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmSmtpServerGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Enables authentication to the SMTP server.  This block supports the following arguments:
        /// </summary>
        [Input("auth")]
        public Input<Inputs.RealmSmtpServerAuthGetArgs>? Auth { get; set; }

        /// <summary>
        /// The email address uses for bounces.
        /// </summary>
        [Input("envelopeFrom")]
        public Input<string>? EnvelopeFrom { get; set; }

        /// <summary>
        /// The email address for the sender.
        /// </summary>
        [Input("from", required: true)]
        public Input<string> From { get; set; } = null!;

        /// <summary>
        /// The display name of the sender email address.
        /// </summary>
        [Input("fromDisplayName")]
        public Input<string>? FromDisplayName { get; set; }

        /// <summary>
        /// The host of the SMTP server.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// The port of the SMTP server (defaults to 25).
        /// </summary>
        [Input("port")]
        public Input<string>? Port { get; set; }

        /// <summary>
        /// The "reply to" email address.
        /// </summary>
        [Input("replyTo")]
        public Input<string>? ReplyTo { get; set; }

        /// <summary>
        /// The display name of the "reply to" email address.
        /// </summary>
        [Input("replyToDisplayName")]
        public Input<string>? ReplyToDisplayName { get; set; }

        /// <summary>
        /// When `true`, enables SSL. Defaults to `false`.
        /// </summary>
        [Input("ssl")]
        public Input<bool>? Ssl { get; set; }

        /// <summary>
        /// When `true`, enables StartTLS. Defaults to `false`.
        /// </summary>
        [Input("starttls")]
        public Input<bool>? Starttls { get; set; }

        public RealmSmtpServerGetArgs()
        {
        }
        public static new RealmSmtpServerGetArgs Empty => new RealmSmtpServerGetArgs();
    }
}
