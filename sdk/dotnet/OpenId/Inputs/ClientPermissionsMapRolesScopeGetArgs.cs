// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId.Inputs
{

    public sealed class ClientPermissionsMapRolesScopeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("decisionStrategy")]
        public Input<string>? DecisionStrategy { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public ClientPermissionsMapRolesScopeGetArgs()
        {
        }
        public static new ClientPermissionsMapRolesScopeGetArgs Empty => new ClientPermissionsMapRolesScopeGetArgs();
    }
}
