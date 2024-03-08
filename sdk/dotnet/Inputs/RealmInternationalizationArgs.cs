// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class RealmInternationalizationArgs : global::Pulumi.ResourceArgs
    {
        [Input("defaultLocale", required: true)]
        public Input<string> DefaultLocale { get; set; } = null!;

        [Input("supportedLocales", required: true)]
        private InputList<string>? _supportedLocales;
        public InputList<string> SupportedLocales
        {
            get => _supportedLocales ?? (_supportedLocales = new InputList<string>());
            set => _supportedLocales = value;
        }

        public RealmInternationalizationArgs()
        {
        }
        public static new RealmInternationalizationArgs Empty => new RealmInternationalizationArgs();
    }
}
