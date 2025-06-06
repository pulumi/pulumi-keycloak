// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.OpenId.Outputs
{

    [OutputType]
    public sealed class ClientGroupPolicyGroup
    {
        public readonly bool ExtendChildren;
        public readonly string Id;
        public readonly string Path;

        [OutputConstructor]
        private ClientGroupPolicyGroup(
            bool extendChildren,

            string id,

            string path)
        {
            ExtendChildren = extendChildren;
            Id = id;
            Path = path;
        }
    }
}
