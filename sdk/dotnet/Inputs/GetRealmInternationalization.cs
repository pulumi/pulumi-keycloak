// *** WARNING: this file was generated by pulumi-language-dotnet. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Inputs
{

    public sealed class GetRealmInternationalizationArgs : global::Pulumi.InvokeArgs
    {
        [Input("defaultLocale", required: true)]
        public string DefaultLocale { get; set; } = null!;

        [Input("supportedLocales", required: true)]
        private List<string>? _supportedLocales;
        public List<string> SupportedLocales
        {
            get => _supportedLocales ?? (_supportedLocales = new List<string>());
            set => _supportedLocales = value;
        }

        public GetRealmInternationalizationArgs()
        {
        }
        public static new GetRealmInternationalizationArgs Empty => new GetRealmInternationalizationArgs();
    }
}
