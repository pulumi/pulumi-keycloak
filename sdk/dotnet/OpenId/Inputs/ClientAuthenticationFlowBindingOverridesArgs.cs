// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId.Inputs
{

    public sealed class ClientAuthenticationFlowBindingOverridesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Browser flow id, (flow needs to exist)
        /// </summary>
        [Input("browserId")]
        public Input<string>? BrowserId { get; set; }

        /// <summary>
        /// Direct grant flow id (flow needs to exist)
        /// </summary>
        [Input("directGrantId")]
        public Input<string>? DirectGrantId { get; set; }

        public ClientAuthenticationFlowBindingOverridesArgs()
        {
        }
        public static new ClientAuthenticationFlowBindingOverridesArgs Empty => new ClientAuthenticationFlowBindingOverridesArgs();
    }
}
