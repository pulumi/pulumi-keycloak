// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml.Outputs
{

    [OutputType]
    public sealed class GetClientAuthenticationFlowBindingOverrideResult
    {
        public readonly string BrowserId;
        public readonly string DirectGrantId;

        [OutputConstructor]
        private GetClientAuthenticationFlowBindingOverrideResult(
            string browserId,

            string directGrantId)
        {
            BrowserId = browserId;
            DirectGrantId = directGrantId;
        }
    }
}
