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
    public sealed class UserInitialPassword
    {
        public readonly bool? Temporary;
        public readonly string Value;

        [OutputConstructor]
        private UserInitialPassword(
            bool? temporary,

            string value)
        {
            Temporary = temporary;
            Value = value;
        }
    }
}
