// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmWebAuthnPolicyGetArgs : Pulumi.ResourceArgs
    {
        [Input("acceptableAaguids")]
        private InputList<string>? _acceptableAaguids;

        /// <summary>
        /// A set of AAGUIDs for which an authenticator can be registered.
        /// </summary>
        public InputList<string> AcceptableAaguids
        {
            get => _acceptableAaguids ?? (_acceptableAaguids = new InputList<string>());
            set => _acceptableAaguids = value;
        }

        /// <summary>
        /// The preference of how to generate a WebAuthn attestation statement. Valid options are `not specified`, `none`, `indirect`, `direct`, or `enterprise`. Defaults to `not specified`.
        /// </summary>
        [Input("attestationConveyancePreference")]
        public Input<string>? AttestationConveyancePreference { get; set; }

        /// <summary>
        /// The acceptable attachment pattern for the WebAuthn authenticator. Valid options are `not specified`, `platform`, or `cross-platform`. Defaults to `not specified`.
        /// </summary>
        [Input("authenticatorAttachment")]
        public Input<string>? AuthenticatorAttachment { get; set; }

        /// <summary>
        /// When `true`, Keycloak will avoid registering the authenticator for WebAuthn if it has already been registered. Defaults to `false`.
        /// </summary>
        [Input("avoidSameAuthenticatorRegister")]
        public Input<bool>? AvoidSameAuthenticatorRegister { get; set; }

        /// <summary>
        /// The timeout value for creating a user's public key credential in seconds. When set to `0`, this timeout option is not adapted. Defaults to `0`.
        /// </summary>
        [Input("createTimeout")]
        public Input<int>? CreateTimeout { get; set; }

        /// <summary>
        /// A human readable server name for the WebAuthn Relying Party. Defaults to `keycloak`.
        /// </summary>
        [Input("relyingPartyEntityName")]
        public Input<string>? RelyingPartyEntityName { get; set; }

        /// <summary>
        /// The WebAuthn relying party ID.
        /// </summary>
        [Input("relyingPartyId")]
        public Input<string>? RelyingPartyId { get; set; }

        /// <summary>
        /// Specifies whether or not a public key should be created to represent the resident key. Valid options are `not specified`, `Yes`, or `No`. Defaults to `not specified`.
        /// </summary>
        [Input("requireResidentKey")]
        public Input<string>? RequireResidentKey { get; set; }

        [Input("signatureAlgorithms")]
        private InputList<string>? _signatureAlgorithms;

        /// <summary>
        /// A set of signature algorithms that should be used for the authentication assertion. Valid options at the time these docs were written are `ES256`, `ES384`, `ES512`, `RS256`, `RS384`, `RS512`, and `RS1`.
        /// </summary>
        public InputList<string> SignatureAlgorithms
        {
            get => _signatureAlgorithms ?? (_signatureAlgorithms = new InputList<string>());
            set => _signatureAlgorithms = value;
        }

        /// <summary>
        /// Specifies the policy for verifying a user logging in via WebAuthn. Valid options are `not specified`, `required`, `preferred`, or `discouraged`. Defaults to `not specified`.
        /// </summary>
        [Input("userVerificationRequirement")]
        public Input<string>? UserVerificationRequirement { get; set; }

        public RealmWebAuthnPolicyGetArgs()
        {
        }
    }
}
