// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class GetRealmWebAuthnPolicyArgs : global::Pulumi.InvokeArgs
    {
        [Input("acceptableAaguids", required: true)]
        private List<string>? _acceptableAaguids;
        public List<string> AcceptableAaguids
        {
            get => _acceptableAaguids ?? (_acceptableAaguids = new List<string>());
            set => _acceptableAaguids = value;
        }

        /// <summary>
        /// Either none, indirect or direct
        /// </summary>
        [Input("attestationConveyancePreference", required: true)]
        public string AttestationConveyancePreference { get; set; } = null!;

        /// <summary>
        /// Either platform or cross-platform
        /// </summary>
        [Input("authenticatorAttachment", required: true)]
        public string AuthenticatorAttachment { get; set; } = null!;

        [Input("avoidSameAuthenticatorRegister", required: true)]
        public bool AvoidSameAuthenticatorRegister { get; set; }

        [Input("createTimeout", required: true)]
        public int CreateTimeout { get; set; }

        [Input("extraOrigins", required: true)]
        private List<string>? _extraOrigins;
        public List<string> ExtraOrigins
        {
            get => _extraOrigins ?? (_extraOrigins = new List<string>());
            set => _extraOrigins = value;
        }

        [Input("relyingPartyEntityName", required: true)]
        public string RelyingPartyEntityName { get; set; } = null!;

        [Input("relyingPartyId", required: true)]
        public string RelyingPartyId { get; set; } = null!;

        /// <summary>
        /// Either Yes or No
        /// </summary>
        [Input("requireResidentKey", required: true)]
        public string RequireResidentKey { get; set; } = null!;

        [Input("signatureAlgorithms", required: true)]
        private List<string>? _signatureAlgorithms;

        /// <summary>
        /// Keycloak lists ES256, ES384, ES512, RS256, ES384, ES512 at the time of writing
        /// </summary>
        public List<string> SignatureAlgorithms
        {
            get => _signatureAlgorithms ?? (_signatureAlgorithms = new List<string>());
            set => _signatureAlgorithms = value;
        }

        /// <summary>
        /// Either required, preferred or discouraged
        /// </summary>
        [Input("userVerificationRequirement", required: true)]
        public string UserVerificationRequirement { get; set; } = null!;

        public GetRealmWebAuthnPolicyArgs()
        {
        }
        public static new GetRealmWebAuthnPolicyArgs Empty => new GetRealmWebAuthnPolicyArgs();
    }
}
