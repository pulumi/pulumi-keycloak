// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Outputs
{

    [OutputType]
    public sealed class OrganizationDomain
    {
        /// <summary>
        /// The name of the organization.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Whether domain is verified or not. Default is false.
        /// </summary>
        public readonly bool? Verified;

        [OutputConstructor]
        private OrganizationDomain(
            string name,

            bool? verified)
        {
            Name = name;
            Verified = verified;
        }
    }
}
