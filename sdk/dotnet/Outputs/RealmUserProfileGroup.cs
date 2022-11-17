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
    public sealed class RealmUserProfileGroup
    {
        /// <summary>
        /// A map of annotations for the group.
        /// </summary>
        public readonly ImmutableDictionary<string, string>? Annotations;
        /// <summary>
        /// The display description of the group.
        /// </summary>
        public readonly string? DisplayDescription;
        /// <summary>
        /// The display header of the group.
        /// </summary>
        public readonly string? DisplayHeader;
        /// <summary>
        /// The name of the group.
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private RealmUserProfileGroup(
            ImmutableDictionary<string, string>? annotations,

            string? displayDescription,

            string? displayHeader,

            string name)
        {
            Annotations = annotations;
            DisplayDescription = displayDescription;
            DisplayHeader = displayHeader;
            Name = name;
        }
    }
}