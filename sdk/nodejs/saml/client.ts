// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing Keycloak clients that use the SAML protocol.
 *
 * Clients are entities that can use Keycloak for user authentication. Typically, clients are applications that redirect users
 * to Keycloak for authentication in order to take advantage of Keycloak's user sessions for SSO.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as fs from "fs";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const samlClient = new keycloak.saml.Client("samlClient", {
 *     realmId: realm.id,
 *     clientId: "saml-client",
 *     signDocuments: false,
 *     signAssertions: true,
 *     includeAuthnStatement: true,
 *     signingCertificate: fs.readFileSync("saml-cert.pem"),
 *     signingPrivateKey: fs.readFileSync("saml-key.pem"),
 * });
 * ```
 *
 * ## Import
 *
 * Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `client_keycloak_id` is the unique ID that Keycloak assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:saml/client:Client saml_client my-realm/dcbc4c73-e478-4928-ae2e-d5e420223352
 * ```
 */
export class Client extends pulumi.CustomResource {
    /**
     * Get an existing Client resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientState, opts?: pulumi.CustomResourceOptions): Client {
        return new Client(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:saml/client:Client';

    /**
     * Returns true if the given object is an instance of Client.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Client {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Client.__pulumiType;
    }

    /**
     * SAML POST Binding URL for the client's assertion consumer service (login responses).
     */
    public readonly assertionConsumerPostUrl!: pulumi.Output<string | undefined>;
    /**
     * SAML Redirect Binding URL for the client's assertion consumer service (login responses).
     */
    public readonly assertionConsumerRedirectUrl!: pulumi.Output<string | undefined>;
    /**
     * Override realm authentication flow bindings
     */
    public readonly authenticationFlowBindingOverrides!: pulumi.Output<outputs.saml.ClientAuthenticationFlowBindingOverrides | undefined>;
    /**
     * When specified, this URL will be used whenever Keycloak needs to link to this client.
     */
    public readonly baseUrl!: pulumi.Output<string | undefined>;
    /**
     * The Canonicalization Method for XML signatures. Should be one of "EXCLUSIVE", "EXCLUSIVE_WITH_COMMENTS", "INCLUSIVE", or "INCLUSIVE_WITH_COMMENTS". Defaults to "EXCLUSIVE".
     */
    public readonly canonicalizationMethod!: pulumi.Output<string | undefined>;
    /**
     * The unique ID of this client, referenced in the URI during authentication and in issued tokens.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signingCertificate` and `signingPrivateKey`. Defaults to `true`.
     */
    public readonly clientSignatureRequired!: pulumi.Output<boolean | undefined>;
    /**
     * The description of this client in the GUI.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key. Defaults to `false`.
     */
    public readonly encryptAssertions!: pulumi.Output<boolean | undefined>;
    /**
     * If assertions for the client are encrypted, this certificate will be used for encryption.
     */
    public readonly encryptionCertificate!: pulumi.Output<string>;
    /**
     * (Computed) The sha1sum fingerprint of the encryption certificate. If the encryption certificate is not in correct base64 format, this will be left empty.
     */
    public /*out*/ readonly encryptionCertificateSha1!: pulumi.Output<string>;
    public readonly extraConfig!: pulumi.Output<{[key: string]: any} | undefined>;
    /**
     * Ignore requested NameID subject format and use the one defined in `nameIdFormat` instead. Defaults to `false`.
     */
    public readonly forceNameIdFormat!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding. Defaults to `true`.
     */
    public readonly forcePostBinding!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, this client will require a browser redirect in order to perform a logout. Defaults to `true`.
     */
    public readonly frontChannelLogout!: pulumi.Output<boolean | undefined>;
    /**
     * - Allow to include all roles mappings in the access token
     */
    public readonly fullScopeAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
     */
    public readonly idpInitiatedSsoRelayState!: pulumi.Output<string | undefined>;
    /**
     * URL fragment name to reference client when you want to do IDP Initiated SSO.
     */
    public readonly idpInitiatedSsoUrlName!: pulumi.Output<string | undefined>;
    /**
     * When `true`, an `AuthnStatement` will be included in the SAML response. Defaults to `true`.
     */
    public readonly includeAuthnStatement!: pulumi.Output<boolean | undefined>;
    /**
     * The login theme of this client.
     */
    public readonly loginTheme!: pulumi.Output<string | undefined>;
    /**
     * SAML POST Binding URL for the client's single logout service.
     */
    public readonly logoutServicePostBindingUrl!: pulumi.Output<string | undefined>;
    /**
     * SAML Redirect Binding URL for the client's single logout service.
     */
    public readonly logoutServiceRedirectBindingUrl!: pulumi.Output<string | undefined>;
    /**
     * When specified, this URL will be used for all SAML requests.
     */
    public readonly masterSamlProcessingUrl!: pulumi.Output<string | undefined>;
    /**
     * The display name of this client in the GUI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Sets the Name ID format for the subject.
     */
    public readonly nameIdFormat!: pulumi.Output<string>;
    /**
     * The realm this client is attached to.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * When specified, this value is prepended to all relative URLs.
     */
    public readonly rootUrl!: pulumi.Output<string | undefined>;
    /**
     * When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response. Defaults to `false`.
     */
    public readonly signAssertions!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, the SAML document will be signed by Keycloak using the realm's private key. Defaults to `true`.
     */
    public readonly signDocuments!: pulumi.Output<boolean | undefined>;
    /**
     * The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA512", or "DSA_SHA1".
     */
    public readonly signatureAlgorithm!: pulumi.Output<string | undefined>;
    /**
     * The value of the `KeyName` element within the signed SAML document. Should be one of "NONE", "KEY_ID", or "CERT_SUBJECT". Defaults to "KEY_ID".
     */
    public readonly signatureKeyName!: pulumi.Output<string | undefined>;
    /**
     * If documents or assertions from the client are signed, this certificate will be used to verify the signature.
     */
    public readonly signingCertificate!: pulumi.Output<string>;
    /**
     * (Computed) The sha1sum fingerprint of the signing certificate. If the signing certificate is not in correct base64 format, this will be left empty.
     */
    public /*out*/ readonly signingCertificateSha1!: pulumi.Output<string>;
    /**
     * If documents or assertions from the client are signed, this private key will be used to verify the signature.
     */
    public readonly signingPrivateKey!: pulumi.Output<string>;
    /**
     * (Computed) The sha1sum fingerprint of the signing private key. If the signing private key is not in correct base64 format, this will be left empty.
     */
    public /*out*/ readonly signingPrivateKeySha1!: pulumi.Output<string>;
    /**
     * When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
     */
    public readonly validRedirectUris!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Client resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientArgs | ClientState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientState | undefined;
            resourceInputs["assertionConsumerPostUrl"] = state ? state.assertionConsumerPostUrl : undefined;
            resourceInputs["assertionConsumerRedirectUrl"] = state ? state.assertionConsumerRedirectUrl : undefined;
            resourceInputs["authenticationFlowBindingOverrides"] = state ? state.authenticationFlowBindingOverrides : undefined;
            resourceInputs["baseUrl"] = state ? state.baseUrl : undefined;
            resourceInputs["canonicalizationMethod"] = state ? state.canonicalizationMethod : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientSignatureRequired"] = state ? state.clientSignatureRequired : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["encryptAssertions"] = state ? state.encryptAssertions : undefined;
            resourceInputs["encryptionCertificate"] = state ? state.encryptionCertificate : undefined;
            resourceInputs["encryptionCertificateSha1"] = state ? state.encryptionCertificateSha1 : undefined;
            resourceInputs["extraConfig"] = state ? state.extraConfig : undefined;
            resourceInputs["forceNameIdFormat"] = state ? state.forceNameIdFormat : undefined;
            resourceInputs["forcePostBinding"] = state ? state.forcePostBinding : undefined;
            resourceInputs["frontChannelLogout"] = state ? state.frontChannelLogout : undefined;
            resourceInputs["fullScopeAllowed"] = state ? state.fullScopeAllowed : undefined;
            resourceInputs["idpInitiatedSsoRelayState"] = state ? state.idpInitiatedSsoRelayState : undefined;
            resourceInputs["idpInitiatedSsoUrlName"] = state ? state.idpInitiatedSsoUrlName : undefined;
            resourceInputs["includeAuthnStatement"] = state ? state.includeAuthnStatement : undefined;
            resourceInputs["loginTheme"] = state ? state.loginTheme : undefined;
            resourceInputs["logoutServicePostBindingUrl"] = state ? state.logoutServicePostBindingUrl : undefined;
            resourceInputs["logoutServiceRedirectBindingUrl"] = state ? state.logoutServiceRedirectBindingUrl : undefined;
            resourceInputs["masterSamlProcessingUrl"] = state ? state.masterSamlProcessingUrl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["nameIdFormat"] = state ? state.nameIdFormat : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["rootUrl"] = state ? state.rootUrl : undefined;
            resourceInputs["signAssertions"] = state ? state.signAssertions : undefined;
            resourceInputs["signDocuments"] = state ? state.signDocuments : undefined;
            resourceInputs["signatureAlgorithm"] = state ? state.signatureAlgorithm : undefined;
            resourceInputs["signatureKeyName"] = state ? state.signatureKeyName : undefined;
            resourceInputs["signingCertificate"] = state ? state.signingCertificate : undefined;
            resourceInputs["signingCertificateSha1"] = state ? state.signingCertificateSha1 : undefined;
            resourceInputs["signingPrivateKey"] = state ? state.signingPrivateKey : undefined;
            resourceInputs["signingPrivateKeySha1"] = state ? state.signingPrivateKeySha1 : undefined;
            resourceInputs["validRedirectUris"] = state ? state.validRedirectUris : undefined;
        } else {
            const args = argsOrState as ClientArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["assertionConsumerPostUrl"] = args ? args.assertionConsumerPostUrl : undefined;
            resourceInputs["assertionConsumerRedirectUrl"] = args ? args.assertionConsumerRedirectUrl : undefined;
            resourceInputs["authenticationFlowBindingOverrides"] = args ? args.authenticationFlowBindingOverrides : undefined;
            resourceInputs["baseUrl"] = args ? args.baseUrl : undefined;
            resourceInputs["canonicalizationMethod"] = args ? args.canonicalizationMethod : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSignatureRequired"] = args ? args.clientSignatureRequired : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["encryptAssertions"] = args ? args.encryptAssertions : undefined;
            resourceInputs["encryptionCertificate"] = args ? args.encryptionCertificate : undefined;
            resourceInputs["extraConfig"] = args ? args.extraConfig : undefined;
            resourceInputs["forceNameIdFormat"] = args ? args.forceNameIdFormat : undefined;
            resourceInputs["forcePostBinding"] = args ? args.forcePostBinding : undefined;
            resourceInputs["frontChannelLogout"] = args ? args.frontChannelLogout : undefined;
            resourceInputs["fullScopeAllowed"] = args ? args.fullScopeAllowed : undefined;
            resourceInputs["idpInitiatedSsoRelayState"] = args ? args.idpInitiatedSsoRelayState : undefined;
            resourceInputs["idpInitiatedSsoUrlName"] = args ? args.idpInitiatedSsoUrlName : undefined;
            resourceInputs["includeAuthnStatement"] = args ? args.includeAuthnStatement : undefined;
            resourceInputs["loginTheme"] = args ? args.loginTheme : undefined;
            resourceInputs["logoutServicePostBindingUrl"] = args ? args.logoutServicePostBindingUrl : undefined;
            resourceInputs["logoutServiceRedirectBindingUrl"] = args ? args.logoutServiceRedirectBindingUrl : undefined;
            resourceInputs["masterSamlProcessingUrl"] = args ? args.masterSamlProcessingUrl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["nameIdFormat"] = args ? args.nameIdFormat : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["rootUrl"] = args ? args.rootUrl : undefined;
            resourceInputs["signAssertions"] = args ? args.signAssertions : undefined;
            resourceInputs["signDocuments"] = args ? args.signDocuments : undefined;
            resourceInputs["signatureAlgorithm"] = args ? args.signatureAlgorithm : undefined;
            resourceInputs["signatureKeyName"] = args ? args.signatureKeyName : undefined;
            resourceInputs["signingCertificate"] = args ? args.signingCertificate : undefined;
            resourceInputs["signingPrivateKey"] = args ? args.signingPrivateKey : undefined;
            resourceInputs["validRedirectUris"] = args ? args.validRedirectUris : undefined;
            resourceInputs["encryptionCertificateSha1"] = undefined /*out*/;
            resourceInputs["signingCertificateSha1"] = undefined /*out*/;
            resourceInputs["signingPrivateKeySha1"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Client.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Client resources.
 */
export interface ClientState {
    /**
     * SAML POST Binding URL for the client's assertion consumer service (login responses).
     */
    assertionConsumerPostUrl?: pulumi.Input<string>;
    /**
     * SAML Redirect Binding URL for the client's assertion consumer service (login responses).
     */
    assertionConsumerRedirectUrl?: pulumi.Input<string>;
    /**
     * Override realm authentication flow bindings
     */
    authenticationFlowBindingOverrides?: pulumi.Input<inputs.saml.ClientAuthenticationFlowBindingOverrides>;
    /**
     * When specified, this URL will be used whenever Keycloak needs to link to this client.
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * The Canonicalization Method for XML signatures. Should be one of "EXCLUSIVE", "EXCLUSIVE_WITH_COMMENTS", "INCLUSIVE", or "INCLUSIVE_WITH_COMMENTS". Defaults to "EXCLUSIVE".
     */
    canonicalizationMethod?: pulumi.Input<string>;
    /**
     * The unique ID of this client, referenced in the URI during authentication and in issued tokens.
     */
    clientId?: pulumi.Input<string>;
    /**
     * When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signingCertificate` and `signingPrivateKey`. Defaults to `true`.
     */
    clientSignatureRequired?: pulumi.Input<boolean>;
    /**
     * The description of this client in the GUI.
     */
    description?: pulumi.Input<string>;
    /**
     * When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key. Defaults to `false`.
     */
    encryptAssertions?: pulumi.Input<boolean>;
    /**
     * If assertions for the client are encrypted, this certificate will be used for encryption.
     */
    encryptionCertificate?: pulumi.Input<string>;
    /**
     * (Computed) The sha1sum fingerprint of the encryption certificate. If the encryption certificate is not in correct base64 format, this will be left empty.
     */
    encryptionCertificateSha1?: pulumi.Input<string>;
    extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * Ignore requested NameID subject format and use the one defined in `nameIdFormat` instead. Defaults to `false`.
     */
    forceNameIdFormat?: pulumi.Input<boolean>;
    /**
     * When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding. Defaults to `true`.
     */
    forcePostBinding?: pulumi.Input<boolean>;
    /**
     * When `true`, this client will require a browser redirect in order to perform a logout. Defaults to `true`.
     */
    frontChannelLogout?: pulumi.Input<boolean>;
    /**
     * - Allow to include all roles mappings in the access token
     */
    fullScopeAllowed?: pulumi.Input<boolean>;
    /**
     * Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
     */
    idpInitiatedSsoRelayState?: pulumi.Input<string>;
    /**
     * URL fragment name to reference client when you want to do IDP Initiated SSO.
     */
    idpInitiatedSsoUrlName?: pulumi.Input<string>;
    /**
     * When `true`, an `AuthnStatement` will be included in the SAML response. Defaults to `true`.
     */
    includeAuthnStatement?: pulumi.Input<boolean>;
    /**
     * The login theme of this client.
     */
    loginTheme?: pulumi.Input<string>;
    /**
     * SAML POST Binding URL for the client's single logout service.
     */
    logoutServicePostBindingUrl?: pulumi.Input<string>;
    /**
     * SAML Redirect Binding URL for the client's single logout service.
     */
    logoutServiceRedirectBindingUrl?: pulumi.Input<string>;
    /**
     * When specified, this URL will be used for all SAML requests.
     */
    masterSamlProcessingUrl?: pulumi.Input<string>;
    /**
     * The display name of this client in the GUI.
     */
    name?: pulumi.Input<string>;
    /**
     * Sets the Name ID format for the subject.
     */
    nameIdFormat?: pulumi.Input<string>;
    /**
     * The realm this client is attached to.
     */
    realmId?: pulumi.Input<string>;
    /**
     * When specified, this value is prepended to all relative URLs.
     */
    rootUrl?: pulumi.Input<string>;
    /**
     * When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response. Defaults to `false`.
     */
    signAssertions?: pulumi.Input<boolean>;
    /**
     * When `true`, the SAML document will be signed by Keycloak using the realm's private key. Defaults to `true`.
     */
    signDocuments?: pulumi.Input<boolean>;
    /**
     * The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA512", or "DSA_SHA1".
     */
    signatureAlgorithm?: pulumi.Input<string>;
    /**
     * The value of the `KeyName` element within the signed SAML document. Should be one of "NONE", "KEY_ID", or "CERT_SUBJECT". Defaults to "KEY_ID".
     */
    signatureKeyName?: pulumi.Input<string>;
    /**
     * If documents or assertions from the client are signed, this certificate will be used to verify the signature.
     */
    signingCertificate?: pulumi.Input<string>;
    /**
     * (Computed) The sha1sum fingerprint of the signing certificate. If the signing certificate is not in correct base64 format, this will be left empty.
     */
    signingCertificateSha1?: pulumi.Input<string>;
    /**
     * If documents or assertions from the client are signed, this private key will be used to verify the signature.
     */
    signingPrivateKey?: pulumi.Input<string>;
    /**
     * (Computed) The sha1sum fingerprint of the signing private key. If the signing private key is not in correct base64 format, this will be left empty.
     */
    signingPrivateKeySha1?: pulumi.Input<string>;
    /**
     * When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
     */
    validRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Client resource.
 */
export interface ClientArgs {
    /**
     * SAML POST Binding URL for the client's assertion consumer service (login responses).
     */
    assertionConsumerPostUrl?: pulumi.Input<string>;
    /**
     * SAML Redirect Binding URL for the client's assertion consumer service (login responses).
     */
    assertionConsumerRedirectUrl?: pulumi.Input<string>;
    /**
     * Override realm authentication flow bindings
     */
    authenticationFlowBindingOverrides?: pulumi.Input<inputs.saml.ClientAuthenticationFlowBindingOverrides>;
    /**
     * When specified, this URL will be used whenever Keycloak needs to link to this client.
     */
    baseUrl?: pulumi.Input<string>;
    /**
     * The Canonicalization Method for XML signatures. Should be one of "EXCLUSIVE", "EXCLUSIVE_WITH_COMMENTS", "INCLUSIVE", or "INCLUSIVE_WITH_COMMENTS". Defaults to "EXCLUSIVE".
     */
    canonicalizationMethod?: pulumi.Input<string>;
    /**
     * The unique ID of this client, referenced in the URI during authentication and in issued tokens.
     */
    clientId: pulumi.Input<string>;
    /**
     * When `true`, Keycloak will expect that documents originating from a client will be signed using the certificate and/or key configured via `signingCertificate` and `signingPrivateKey`. Defaults to `true`.
     */
    clientSignatureRequired?: pulumi.Input<boolean>;
    /**
     * The description of this client in the GUI.
     */
    description?: pulumi.Input<string>;
    /**
     * When false, this client will not be able to initiate a login or obtain access tokens. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * When `true`, the SAML assertions will be encrypted by Keycloak using the client's public key. Defaults to `false`.
     */
    encryptAssertions?: pulumi.Input<boolean>;
    /**
     * If assertions for the client are encrypted, this certificate will be used for encryption.
     */
    encryptionCertificate?: pulumi.Input<string>;
    extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * Ignore requested NameID subject format and use the one defined in `nameIdFormat` instead. Defaults to `false`.
     */
    forceNameIdFormat?: pulumi.Input<boolean>;
    /**
     * When `true`, Keycloak will always respond to an authentication request via the SAML POST Binding. Defaults to `true`.
     */
    forcePostBinding?: pulumi.Input<boolean>;
    /**
     * When `true`, this client will require a browser redirect in order to perform a logout. Defaults to `true`.
     */
    frontChannelLogout?: pulumi.Input<boolean>;
    /**
     * - Allow to include all roles mappings in the access token
     */
    fullScopeAllowed?: pulumi.Input<boolean>;
    /**
     * Relay state you want to send with SAML request when you want to do IDP Initiated SSO.
     */
    idpInitiatedSsoRelayState?: pulumi.Input<string>;
    /**
     * URL fragment name to reference client when you want to do IDP Initiated SSO.
     */
    idpInitiatedSsoUrlName?: pulumi.Input<string>;
    /**
     * When `true`, an `AuthnStatement` will be included in the SAML response. Defaults to `true`.
     */
    includeAuthnStatement?: pulumi.Input<boolean>;
    /**
     * The login theme of this client.
     */
    loginTheme?: pulumi.Input<string>;
    /**
     * SAML POST Binding URL for the client's single logout service.
     */
    logoutServicePostBindingUrl?: pulumi.Input<string>;
    /**
     * SAML Redirect Binding URL for the client's single logout service.
     */
    logoutServiceRedirectBindingUrl?: pulumi.Input<string>;
    /**
     * When specified, this URL will be used for all SAML requests.
     */
    masterSamlProcessingUrl?: pulumi.Input<string>;
    /**
     * The display name of this client in the GUI.
     */
    name?: pulumi.Input<string>;
    /**
     * Sets the Name ID format for the subject.
     */
    nameIdFormat?: pulumi.Input<string>;
    /**
     * The realm this client is attached to.
     */
    realmId: pulumi.Input<string>;
    /**
     * When specified, this value is prepended to all relative URLs.
     */
    rootUrl?: pulumi.Input<string>;
    /**
     * When `true`, the SAML assertions will be signed by Keycloak using the realm's private key, and embedded within the SAML XML Auth response. Defaults to `false`.
     */
    signAssertions?: pulumi.Input<boolean>;
    /**
     * When `true`, the SAML document will be signed by Keycloak using the realm's private key. Defaults to `true`.
     */
    signDocuments?: pulumi.Input<boolean>;
    /**
     * The signature algorithm used to sign documents. Should be one of "RSA_SHA1", "RSA_SHA256", "RSA_SHA512", or "DSA_SHA1".
     */
    signatureAlgorithm?: pulumi.Input<string>;
    /**
     * The value of the `KeyName` element within the signed SAML document. Should be one of "NONE", "KEY_ID", or "CERT_SUBJECT". Defaults to "KEY_ID".
     */
    signatureKeyName?: pulumi.Input<string>;
    /**
     * If documents or assertions from the client are signed, this certificate will be used to verify the signature.
     */
    signingCertificate?: pulumi.Input<string>;
    /**
     * If documents or assertions from the client are signed, this private key will be used to verify the signature.
     */
    signingPrivateKey?: pulumi.Input<string>;
    /**
     * When specified, Keycloak will use this list to validate given Assertion Consumer URLs specified in the authentication request.
     */
    validRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}
