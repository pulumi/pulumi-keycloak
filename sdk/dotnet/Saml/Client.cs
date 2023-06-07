// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Keycloak.Saml
{
    /// <summary>
    /// Allows for creating and managing Keycloak clients that use the SAML protocol.
    /// 
    /// Clients are entities that can use Keycloak for user authentication. Typically, clients are applications that redirect users
    /// to Keycloak for authentication in order to take advantage of Keycloak's user sessions for SSO.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Keycloak = Pulumi.Keycloak;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var realm = new Keycloak.Realm("realm", new()
    ///     {
    ///         RealmName = "my-realm",
    ///         Enabled = true,
    ///     });
    /// 
    ///     var samlClient = new Keycloak.Saml.Client("samlClient", new()
    ///     {
    ///         RealmId = realm.Id,
    ///         ClientId = "saml-client",
    ///         SignDocuments = false,
    ///         SignAssertions = true,
    ///         IncludeAuthnStatement = true,
    ///         SigningCertificate = File.ReadAllText("saml-cert.pem"),
    ///         SigningPrivateKey = File.ReadAllText("saml-key.pem"),
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `client_keycloak_id` is the unique ID that Keycloak assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID. Examplebash
    /// 
    /// ```sh
    ///  $ pulumi import keycloak:saml/client:Client saml_client my-realm/dcbc4c73-e478-4928-ae2e-d5e420223352
    /// ```
    /// </summary>
    [KeycloakResourceType("keycloak:saml/client:Client")]
    public partial class Client : global::Pulumi.CustomResource
    {
        /// <summary>
        /// SAML POST Binding URL for the client's assertion consumer service (login responses).
        /// </summary>
        [Output("assertionConsumerPostUrl")]
        public Output<string?> AssertionConsumerPostUrl { get; private set; } = null!;

        /// <summary>
        /// SAML Redirect Binding URL for the client's assertion consumer service (login responses).
        /// </summary>
        [Output("assertionConsumerRedirectUrl")]
        public Output<string?> AssertionConsumerRedirectUrl { get; private set; } = null!;

        /// <summary>
        /// Override realm authentication flow bindings
        /// </summary>
        [Output("authenticationFlowBindingOverrides")]
        public Output<Outputs.ClientAuthenticationFlowBindingOverrides?> AuthenticationFlowBindingOverrides { get; private set; } = null!;

        /// <summary>
        /// When specified, this URL will be used whenever Keycloak needs to link to this client.
        /// </summary>
        [Output("baseUrl")]
        public Output<string?> BaseUrl { get; private set; } = null!;

        /// <summary>
        /// The Canonicalization Method for XML signatures. Should be one of "EXCLUSIVE", "EXCLUSIVE_WITH_COMMENTS", "INCLUSIVE", or "INCLUSIVE_WITH_COMMENTS". Defaults to "EXCLUSIVE".
        /// </summary>
        [Output("canonicalizationMethod")]
        public Output<string?> CanonicalizationMethod { get; private set; } = null!;

        /// <summary>
        /// The unique ID of this client, referenced in the URI during authentication and in issued tokens.
        /// </summary>
        [Output("clientId")]
        public Output<string> ClientId { get; private set; } = null!;

        /// <summary>
        /// When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signing_certificate` and `signing_private_key`. Defaults to `true`.
        /// </summary>
        [Output("clientSignatureRequired")]
        public Output<bool?> ClientSignatureRequired { get; private set; } = null!;

        /// <summary>
        /// The description of this client in the GUI.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key. Defaults to `false`.
        /// </summary>
        [Output("encryptAssertions")]
        public Output<bool?> EncryptAssertions { get; private set; } = null!;

        /// <summary>
        /// If assertions for the client are encrypted, this certificate will be used for encryption.
        /// </summary>
        [Output("encryptionCertificate")]
        public Output<string> EncryptionCertificate { get; private set; } = null!;

        /// <summary>
        /// (Computed) The sha1sum fingerprint of the encryption certificate. If the encryption certificate is not in correct base64 format, this will be left empty.
        /// </summary>
        [Output("encryptionCertificateSha1")]
        public Output<string> EncryptionCertificateSha1 { get; private set; } = null!;

        [Output("extraConfig")]
        public Output<ImmutableDictionary<string, object>?> ExtraConfig { get; private set; } = null!;

        /// <summary>
        /// Ignore requested NameID subject format and use the one defined in `name_id_format` instead. Defaults to `false`.
        /// </summary>
        [Output("forceNameIdFormat")]
        public Output<bool?> ForceNameIdFormat { get; private set; } = null!;

        /// <summary>
        /// When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding. Defaults to `true`.
        /// </summary>
        [Output("forcePostBinding")]
        public Output<bool?> ForcePostBinding { get; private set; } = null!;

        /// <summary>
        /// When `true`, this client will require a browser redirect in order to perform a logout. Defaults to `true`.
        /// </summary>
        [Output("frontChannelLogout")]
        public Output<bool?> FrontChannelLogout { get; private set; } = null!;

        /// <summary>
        /// Allow to include all roles mappings in the access token
        /// </summary>
        [Output("fullScopeAllowed")]
        public Output<bool?> FullScopeAllowed { get; private set; } = null!;

        /// <summary>
        /// Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
        /// </summary>
        [Output("idpInitiatedSsoRelayState")]
        public Output<string?> IdpInitiatedSsoRelayState { get; private set; } = null!;

        /// <summary>
        /// URL fragment name to reference client when you want to do IDP Initiated SSO.
        /// </summary>
        [Output("idpInitiatedSsoUrlName")]
        public Output<string?> IdpInitiatedSsoUrlName { get; private set; } = null!;

        /// <summary>
        /// When `true`, an `AuthnStatement` will be included in the SAML response. Defaults to `true`.
        /// </summary>
        [Output("includeAuthnStatement")]
        public Output<bool?> IncludeAuthnStatement { get; private set; } = null!;

        /// <summary>
        /// The login theme of this client.
        /// </summary>
        [Output("loginTheme")]
        public Output<string?> LoginTheme { get; private set; } = null!;

        /// <summary>
        /// SAML POST Binding URL for the client's single logout service.
        /// </summary>
        [Output("logoutServicePostBindingUrl")]
        public Output<string?> LogoutServicePostBindingUrl { get; private set; } = null!;

        /// <summary>
        /// SAML Redirect Binding URL for the client's single logout service.
        /// </summary>
        [Output("logoutServiceRedirectBindingUrl")]
        public Output<string?> LogoutServiceRedirectBindingUrl { get; private set; } = null!;

        /// <summary>
        /// When specified, this URL will be used for all SAML requests.
        /// </summary>
        [Output("masterSamlProcessingUrl")]
        public Output<string?> MasterSamlProcessingUrl { get; private set; } = null!;

        /// <summary>
        /// The display name of this client in the GUI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Sets the Name ID format for the subject.
        /// </summary>
        [Output("nameIdFormat")]
        public Output<string> NameIdFormat { get; private set; } = null!;

        /// <summary>
        /// The realm this client is attached to.
        /// </summary>
        [Output("realmId")]
        public Output<string> RealmId { get; private set; } = null!;

        /// <summary>
        /// When specified, this value is prepended to all relative URLs.
        /// </summary>
        [Output("rootUrl")]
        public Output<string?> RootUrl { get; private set; } = null!;

        /// <summary>
        /// When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response. Defaults to `false`.
        /// </summary>
        [Output("signAssertions")]
        public Output<bool?> SignAssertions { get; private set; } = null!;

        /// <summary>
        /// When `true`, the SAML document will be signed by Keycloak using the realm's private key. Defaults to `true`.
        /// </summary>
        [Output("signDocuments")]
        public Output<bool?> SignDocuments { get; private set; } = null!;

        /// <summary>
        /// The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA256_MGF1, "RSA_SHA512", "RSA_SHA512_MGF1" or "DSA_SHA1".
        /// </summary>
        [Output("signatureAlgorithm")]
        public Output<string?> SignatureAlgorithm { get; private set; } = null!;

        /// <summary>
        /// The value of the `KeyName` element within the signed SAML document. Should be one of "NONE", "KEY_ID", or "CERT_SUBJECT". Defaults to "KEY_ID".
        /// </summary>
        [Output("signatureKeyName")]
        public Output<string?> SignatureKeyName { get; private set; } = null!;

        /// <summary>
        /// If documents or assertions from the client are signed, this certificate will be used to verify the signature.
        /// </summary>
        [Output("signingCertificate")]
        public Output<string> SigningCertificate { get; private set; } = null!;

        /// <summary>
        /// (Computed) The sha1sum fingerprint of the signing certificate. If the signing certificate is not in correct base64 format, this will be left empty.
        /// </summary>
        [Output("signingCertificateSha1")]
        public Output<string> SigningCertificateSha1 { get; private set; } = null!;

        /// <summary>
        /// If documents or assertions from the client are signed, this private key will be used to verify the signature.
        /// </summary>
        [Output("signingPrivateKey")]
        public Output<string> SigningPrivateKey { get; private set; } = null!;

        /// <summary>
        /// (Computed) The sha1sum fingerprint of the signing private key. If the signing private key is not in correct base64 format, this will be left empty.
        /// </summary>
        [Output("signingPrivateKeySha1")]
        public Output<string> SigningPrivateKeySha1 { get; private set; } = null!;

        /// <summary>
        /// When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
        /// </summary>
        [Output("validRedirectUris")]
        public Output<ImmutableArray<string>> ValidRedirectUris { get; private set; } = null!;


        /// <summary>
        /// Create a Client resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Client(string name, ClientArgs args, CustomResourceOptions? options = null)
            : base("keycloak:saml/client:Client", name, args ?? new ClientArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Client(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
            : base("keycloak:saml/client:Client", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Client resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Client Get(string name, Input<string> id, ClientState? state = null, CustomResourceOptions? options = null)
        {
            return new Client(name, id, state, options);
        }
    }

    public sealed class ClientArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SAML POST Binding URL for the client's assertion consumer service (login responses).
        /// </summary>
        [Input("assertionConsumerPostUrl")]
        public Input<string>? AssertionConsumerPostUrl { get; set; }

        /// <summary>
        /// SAML Redirect Binding URL for the client's assertion consumer service (login responses).
        /// </summary>
        [Input("assertionConsumerRedirectUrl")]
        public Input<string>? AssertionConsumerRedirectUrl { get; set; }

        /// <summary>
        /// Override realm authentication flow bindings
        /// </summary>
        [Input("authenticationFlowBindingOverrides")]
        public Input<Inputs.ClientAuthenticationFlowBindingOverridesArgs>? AuthenticationFlowBindingOverrides { get; set; }

        /// <summary>
        /// When specified, this URL will be used whenever Keycloak needs to link to this client.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// The Canonicalization Method for XML signatures. Should be one of "EXCLUSIVE", "EXCLUSIVE_WITH_COMMENTS", "INCLUSIVE", or "INCLUSIVE_WITH_COMMENTS". Defaults to "EXCLUSIVE".
        /// </summary>
        [Input("canonicalizationMethod")]
        public Input<string>? CanonicalizationMethod { get; set; }

        /// <summary>
        /// The unique ID of this client, referenced in the URI during authentication and in issued tokens.
        /// </summary>
        [Input("clientId", required: true)]
        public Input<string> ClientId { get; set; } = null!;

        /// <summary>
        /// When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signing_certificate` and `signing_private_key`. Defaults to `true`.
        /// </summary>
        [Input("clientSignatureRequired")]
        public Input<bool>? ClientSignatureRequired { get; set; }

        /// <summary>
        /// The description of this client in the GUI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key. Defaults to `false`.
        /// </summary>
        [Input("encryptAssertions")]
        public Input<bool>? EncryptAssertions { get; set; }

        /// <summary>
        /// If assertions for the client are encrypted, this certificate will be used for encryption.
        /// </summary>
        [Input("encryptionCertificate")]
        public Input<string>? EncryptionCertificate { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// Ignore requested NameID subject format and use the one defined in `name_id_format` instead. Defaults to `false`.
        /// </summary>
        [Input("forceNameIdFormat")]
        public Input<bool>? ForceNameIdFormat { get; set; }

        /// <summary>
        /// When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding. Defaults to `true`.
        /// </summary>
        [Input("forcePostBinding")]
        public Input<bool>? ForcePostBinding { get; set; }

        /// <summary>
        /// When `true`, this client will require a browser redirect in order to perform a logout. Defaults to `true`.
        /// </summary>
        [Input("frontChannelLogout")]
        public Input<bool>? FrontChannelLogout { get; set; }

        /// <summary>
        /// Allow to include all roles mappings in the access token
        /// </summary>
        [Input("fullScopeAllowed")]
        public Input<bool>? FullScopeAllowed { get; set; }

        /// <summary>
        /// Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
        /// </summary>
        [Input("idpInitiatedSsoRelayState")]
        public Input<string>? IdpInitiatedSsoRelayState { get; set; }

        /// <summary>
        /// URL fragment name to reference client when you want to do IDP Initiated SSO.
        /// </summary>
        [Input("idpInitiatedSsoUrlName")]
        public Input<string>? IdpInitiatedSsoUrlName { get; set; }

        /// <summary>
        /// When `true`, an `AuthnStatement` will be included in the SAML response. Defaults to `true`.
        /// </summary>
        [Input("includeAuthnStatement")]
        public Input<bool>? IncludeAuthnStatement { get; set; }

        /// <summary>
        /// The login theme of this client.
        /// </summary>
        [Input("loginTheme")]
        public Input<string>? LoginTheme { get; set; }

        /// <summary>
        /// SAML POST Binding URL for the client's single logout service.
        /// </summary>
        [Input("logoutServicePostBindingUrl")]
        public Input<string>? LogoutServicePostBindingUrl { get; set; }

        /// <summary>
        /// SAML Redirect Binding URL for the client's single logout service.
        /// </summary>
        [Input("logoutServiceRedirectBindingUrl")]
        public Input<string>? LogoutServiceRedirectBindingUrl { get; set; }

        /// <summary>
        /// When specified, this URL will be used for all SAML requests.
        /// </summary>
        [Input("masterSamlProcessingUrl")]
        public Input<string>? MasterSamlProcessingUrl { get; set; }

        /// <summary>
        /// The display name of this client in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets the Name ID format for the subject.
        /// </summary>
        [Input("nameIdFormat")]
        public Input<string>? NameIdFormat { get; set; }

        /// <summary>
        /// The realm this client is attached to.
        /// </summary>
        [Input("realmId", required: true)]
        public Input<string> RealmId { get; set; } = null!;

        /// <summary>
        /// When specified, this value is prepended to all relative URLs.
        /// </summary>
        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        /// <summary>
        /// When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response. Defaults to `false`.
        /// </summary>
        [Input("signAssertions")]
        public Input<bool>? SignAssertions { get; set; }

        /// <summary>
        /// When `true`, the SAML document will be signed by Keycloak using the realm's private key. Defaults to `true`.
        /// </summary>
        [Input("signDocuments")]
        public Input<bool>? SignDocuments { get; set; }

        /// <summary>
        /// The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA256_MGF1, "RSA_SHA512", "RSA_SHA512_MGF1" or "DSA_SHA1".
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        /// <summary>
        /// The value of the `KeyName` element within the signed SAML document. Should be one of "NONE", "KEY_ID", or "CERT_SUBJECT". Defaults to "KEY_ID".
        /// </summary>
        [Input("signatureKeyName")]
        public Input<string>? SignatureKeyName { get; set; }

        /// <summary>
        /// If documents or assertions from the client are signed, this certificate will be used to verify the signature.
        /// </summary>
        [Input("signingCertificate")]
        public Input<string>? SigningCertificate { get; set; }

        /// <summary>
        /// If documents or assertions from the client are signed, this private key will be used to verify the signature.
        /// </summary>
        [Input("signingPrivateKey")]
        public Input<string>? SigningPrivateKey { get; set; }

        [Input("validRedirectUris")]
        private InputList<string>? _validRedirectUris;

        /// <summary>
        /// When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
        /// </summary>
        public InputList<string> ValidRedirectUris
        {
            get => _validRedirectUris ?? (_validRedirectUris = new InputList<string>());
            set => _validRedirectUris = value;
        }

        public ClientArgs()
        {
        }
        public static new ClientArgs Empty => new ClientArgs();
    }

    public sealed class ClientState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// SAML POST Binding URL for the client's assertion consumer service (login responses).
        /// </summary>
        [Input("assertionConsumerPostUrl")]
        public Input<string>? AssertionConsumerPostUrl { get; set; }

        /// <summary>
        /// SAML Redirect Binding URL for the client's assertion consumer service (login responses).
        /// </summary>
        [Input("assertionConsumerRedirectUrl")]
        public Input<string>? AssertionConsumerRedirectUrl { get; set; }

        /// <summary>
        /// Override realm authentication flow bindings
        /// </summary>
        [Input("authenticationFlowBindingOverrides")]
        public Input<Inputs.ClientAuthenticationFlowBindingOverridesGetArgs>? AuthenticationFlowBindingOverrides { get; set; }

        /// <summary>
        /// When specified, this URL will be used whenever Keycloak needs to link to this client.
        /// </summary>
        [Input("baseUrl")]
        public Input<string>? BaseUrl { get; set; }

        /// <summary>
        /// The Canonicalization Method for XML signatures. Should be one of "EXCLUSIVE", "EXCLUSIVE_WITH_COMMENTS", "INCLUSIVE", or "INCLUSIVE_WITH_COMMENTS". Defaults to "EXCLUSIVE".
        /// </summary>
        [Input("canonicalizationMethod")]
        public Input<string>? CanonicalizationMethod { get; set; }

        /// <summary>
        /// The unique ID of this client, referenced in the URI during authentication and in issued tokens.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signing_certificate` and `signing_private_key`. Defaults to `true`.
        /// </summary>
        [Input("clientSignatureRequired")]
        public Input<bool>? ClientSignatureRequired { get; set; }

        /// <summary>
        /// The description of this client in the GUI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key. Defaults to `false`.
        /// </summary>
        [Input("encryptAssertions")]
        public Input<bool>? EncryptAssertions { get; set; }

        /// <summary>
        /// If assertions for the client are encrypted, this certificate will be used for encryption.
        /// </summary>
        [Input("encryptionCertificate")]
        public Input<string>? EncryptionCertificate { get; set; }

        /// <summary>
        /// (Computed) The sha1sum fingerprint of the encryption certificate. If the encryption certificate is not in correct base64 format, this will be left empty.
        /// </summary>
        [Input("encryptionCertificateSha1")]
        public Input<string>? EncryptionCertificateSha1 { get; set; }

        [Input("extraConfig")]
        private InputMap<object>? _extraConfig;
        public InputMap<object> ExtraConfig
        {
            get => _extraConfig ?? (_extraConfig = new InputMap<object>());
            set => _extraConfig = value;
        }

        /// <summary>
        /// Ignore requested NameID subject format and use the one defined in `name_id_format` instead. Defaults to `false`.
        /// </summary>
        [Input("forceNameIdFormat")]
        public Input<bool>? ForceNameIdFormat { get; set; }

        /// <summary>
        /// When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding. Defaults to `true`.
        /// </summary>
        [Input("forcePostBinding")]
        public Input<bool>? ForcePostBinding { get; set; }

        /// <summary>
        /// When `true`, this client will require a browser redirect in order to perform a logout. Defaults to `true`.
        /// </summary>
        [Input("frontChannelLogout")]
        public Input<bool>? FrontChannelLogout { get; set; }

        /// <summary>
        /// Allow to include all roles mappings in the access token
        /// </summary>
        [Input("fullScopeAllowed")]
        public Input<bool>? FullScopeAllowed { get; set; }

        /// <summary>
        /// Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
        /// </summary>
        [Input("idpInitiatedSsoRelayState")]
        public Input<string>? IdpInitiatedSsoRelayState { get; set; }

        /// <summary>
        /// URL fragment name to reference client when you want to do IDP Initiated SSO.
        /// </summary>
        [Input("idpInitiatedSsoUrlName")]
        public Input<string>? IdpInitiatedSsoUrlName { get; set; }

        /// <summary>
        /// When `true`, an `AuthnStatement` will be included in the SAML response. Defaults to `true`.
        /// </summary>
        [Input("includeAuthnStatement")]
        public Input<bool>? IncludeAuthnStatement { get; set; }

        /// <summary>
        /// The login theme of this client.
        /// </summary>
        [Input("loginTheme")]
        public Input<string>? LoginTheme { get; set; }

        /// <summary>
        /// SAML POST Binding URL for the client's single logout service.
        /// </summary>
        [Input("logoutServicePostBindingUrl")]
        public Input<string>? LogoutServicePostBindingUrl { get; set; }

        /// <summary>
        /// SAML Redirect Binding URL for the client's single logout service.
        /// </summary>
        [Input("logoutServiceRedirectBindingUrl")]
        public Input<string>? LogoutServiceRedirectBindingUrl { get; set; }

        /// <summary>
        /// When specified, this URL will be used for all SAML requests.
        /// </summary>
        [Input("masterSamlProcessingUrl")]
        public Input<string>? MasterSamlProcessingUrl { get; set; }

        /// <summary>
        /// The display name of this client in the GUI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Sets the Name ID format for the subject.
        /// </summary>
        [Input("nameIdFormat")]
        public Input<string>? NameIdFormat { get; set; }

        /// <summary>
        /// The realm this client is attached to.
        /// </summary>
        [Input("realmId")]
        public Input<string>? RealmId { get; set; }

        /// <summary>
        /// When specified, this value is prepended to all relative URLs.
        /// </summary>
        [Input("rootUrl")]
        public Input<string>? RootUrl { get; set; }

        /// <summary>
        /// When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response. Defaults to `false`.
        /// </summary>
        [Input("signAssertions")]
        public Input<bool>? SignAssertions { get; set; }

        /// <summary>
        /// When `true`, the SAML document will be signed by Keycloak using the realm's private key. Defaults to `true`.
        /// </summary>
        [Input("signDocuments")]
        public Input<bool>? SignDocuments { get; set; }

        /// <summary>
        /// The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA256_MGF1, "RSA_SHA512", "RSA_SHA512_MGF1" or "DSA_SHA1".
        /// </summary>
        [Input("signatureAlgorithm")]
        public Input<string>? SignatureAlgorithm { get; set; }

        /// <summary>
        /// The value of the `KeyName` element within the signed SAML document. Should be one of "NONE", "KEY_ID", or "CERT_SUBJECT". Defaults to "KEY_ID".
        /// </summary>
        [Input("signatureKeyName")]
        public Input<string>? SignatureKeyName { get; set; }

        /// <summary>
        /// If documents or assertions from the client are signed, this certificate will be used to verify the signature.
        /// </summary>
        [Input("signingCertificate")]
        public Input<string>? SigningCertificate { get; set; }

        /// <summary>
        /// (Computed) The sha1sum fingerprint of the signing certificate. If the signing certificate is not in correct base64 format, this will be left empty.
        /// </summary>
        [Input("signingCertificateSha1")]
        public Input<string>? SigningCertificateSha1 { get; set; }

        /// <summary>
        /// If documents or assertions from the client are signed, this private key will be used to verify the signature.
        /// </summary>
        [Input("signingPrivateKey")]
        public Input<string>? SigningPrivateKey { get; set; }

        /// <summary>
        /// (Computed) The sha1sum fingerprint of the signing private key. If the signing private key is not in correct base64 format, this will be left empty.
        /// </summary>
        [Input("signingPrivateKeySha1")]
        public Input<string>? SigningPrivateKeySha1 { get; set; }

        [Input("validRedirectUris")]
        private InputList<string>? _validRedirectUris;

        /// <summary>
        /// When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
        /// </summary>
        public InputList<string> ValidRedirectUris
        {
            get => _validRedirectUris ?? (_validRedirectUris = new InputList<string>());
            set => _validRedirectUris = value;
        }

        public ClientState()
        {
        }
        public static new ClientState Empty => new ClientState();
    }
}
