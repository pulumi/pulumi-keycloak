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
    public sealed class RealmUserProfileAttributePermissions
    {
        public readonly ImmutableArray<string> Edits;
        public readonly ImmutableArray<string> Views;

        [OutputConstructor]
        private RealmUserProfileAttributePermissions(
            ImmutableArray<string> edits,

            ImmutableArray<string> views)
        {
            Edits = edits;
            Views = views;
        }
    }
}
