// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmUserProfileGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("annotations")]
        private InputMap<string>? _annotations;
        public InputMap<string> Annotations
        {
            get => _annotations ?? (_annotations = new InputMap<string>());
            set => _annotations = value;
        }

        [Input("displayDescription")]
        public Input<string>? DisplayDescription { get; set; }

        [Input("displayHeader")]
        public Input<string>? DisplayHeader { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public RealmUserProfileGroupArgs()
        {
        }
        public static new RealmUserProfileGroupArgs Empty => new RealmUserProfileGroupArgs();
    }
}
