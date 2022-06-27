// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmUserProfileGroupArgs : Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;

        /// <summary>
        /// A map of annotations for the group.
        /// </summary>
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        /// <summary>
        /// The display description of the group.
        /// </summary>
        [Input("displayDescription")]
        public Input<string>? DisplayDescription { get; set; }

        /// <summary>
        /// The display header of the group.
        /// </summary>
        [Input("displayHeader")]
        public Input<string>? DisplayHeader { get; set; }

        /// <summary>
        /// The name of the group.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public RealmUserProfileGroupArgs()
        {
        }
    }
}
