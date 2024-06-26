// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmUserProfileAttributePermissionsArgs : global::Pulumi.ResourceArgs
    {
        [Input("edits", required: true)]
        private InputList<string>? _edits;
        public InputList<string> Edits
        {
            get => _edits ?? (_edits = new InputList<string>());
            set => _edits = value;
        }

        [Input("views", required: true)]
        private InputList<string>? _views;
        public InputList<string> Views
        {
            get => _views ?? (_views = new InputList<string>());
            set => _views = value;
        }

        public RealmUserProfileAttributePermissionsArgs()
        {
        }
        public static new RealmUserProfileAttributePermissionsArgs Empty => new RealmUserProfileAttributePermissionsArgs();
    }
}
