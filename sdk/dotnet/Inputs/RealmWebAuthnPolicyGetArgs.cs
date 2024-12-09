// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmWebAuthnPolicyGetArgs : global::Pulumi.ResourceArgs
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
        /// Either none, indirect or direct
        /// </summary>
        [Input("attestationConveyancePreference")]
        public Input<string>? AttestationConveyancePreference { get; set; }

        /// <summary>
        /// Either platform or cross-platform
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
        /// Either Yes or No
        /// </summary>
        [Input("requireResidentKey")]
        public Input<string>? RequireResidentKey { get; set; }

        [Input("signatureAlgorithms")]
        private InputList<string>? _signatureAlgorithms;

        /// <summary>
        /// Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
        /// </summary>
        public InputList<string> SignatureAlgorithms
        {
            get => _signatureAlgorithms ?? (_signatureAlgorithms = new InputList<string>());
            set => _signatureAlgorithms = value;
        }

        /// <summary>
        /// Either required, preferred or discouraged
        /// </summary>
        [Input("userVerificationRequirement")]
        public Input<string>? UserVerificationRequirement { get; set; }

        public RealmWebAuthnPolicyGetArgs()
        {
        }
        public static new RealmWebAuthnPolicyGetArgs Empty => new RealmWebAuthnPolicyGetArgs();
    }
}
