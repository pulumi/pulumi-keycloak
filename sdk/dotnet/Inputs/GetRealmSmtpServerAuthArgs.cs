// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class GetRealmSmtpServerAuthArgs : Pulumi.InvokeArgs
    {
        [Input("password", required: true)]
        public string Password { get; set; } = null!;

        [Input("username", required: true)]
        public string Username { get; set; } = null!;

        public GetRealmSmtpServerAuthArgs()
        {
        }
    }
}
