// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Keycloak
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("keycloak");
        public static ImmutableDictionary<string, string>? AdditionalHeaders { get; set; } = __config.GetObject<ImmutableDictionary<string, string>>("additionalHeaders");

        public static string? BasePath { get; set; } = __config.Get("basePath");

        public static string? ClientId { get; set; } = __config.Get("clientId");

        public static string? ClientSecret { get; set; } = __config.Get("clientSecret");

        /// <summary>
        /// Timeout (in seconds) of the Keycloak client
        /// </summary>
        public static int? ClientTimeout { get; set; } = __config.GetInt32("clientTimeout") ?? Utilities.GetEnvInt32("KEYCLOAK_CLIENT_TIMEOUT") ?? 5;

        /// <summary>
        /// Whether or not to login to Keycloak instance on provider initialization
        /// </summary>
        public static bool? InitialLogin { get; set; } = __config.GetBoolean("initialLogin");

        public static string? Password { get; set; } = __config.Get("password");

        public static string? Realm { get; set; } = __config.Get("realm");

        /// <summary>
        /// Allows x509 calls using an unknown CA certificate (for development purposes)
        /// </summary>
        public static string? RootCaCertificate { get; set; } = __config.Get("rootCaCertificate");

        /// <summary>
        /// Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
        /// should be avoided.
        /// </summary>
        public static bool? TlsInsecureSkipVerify { get; set; } = __config.GetBoolean("tlsInsecureSkipVerify");

        /// <summary>
        /// The base URL of the Keycloak instance, before `/auth`
        /// </summary>
        public static string? Url { get; set; } = __config.Get("url");

        public static string? Username { get; set; } = __config.Get("username");

    }
}
