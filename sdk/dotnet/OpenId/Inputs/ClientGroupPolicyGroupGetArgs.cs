// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId.Inputs
{

    public sealed class ClientGroupPolicyGroupGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("extendChildren", required: true)]
        public Input<bool> ExtendChildren { get; set; } = null!;

        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public ClientGroupPolicyGroupGetArgs()
        {
        }
        public static new ClientGroupPolicyGroupGetArgs Empty => new ClientGroupPolicyGroupGetArgs();
    }
}
