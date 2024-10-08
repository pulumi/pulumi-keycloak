// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * ## # keycloak.saml.Client
 *
 * Allows for creating and managing Keycloak clients that use the SAML protocol.
 *
 * Clients are entities that can use Keycloak for user authentication. Typically,
 * clients are applications that redirect users to Keycloak for authentication
 * in order to take advantage of Keycloak's user sessions for SSO.
 *
 * ### Import
 *
 * Clients can be imported using the format `{{realm_id}}/{{client_keycloak_id}}`, where `clientKeycloakId` is the unique ID that Keycloak
 * assigns to the client upon creation. This value can be found in the URI when editing this client in the GUI, and is typically a GUID.
 *
 * Example:
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

    public readonly assertionConsumerPostUrl!: pulumi.Output<string | undefined>;
    public readonly assertionConsumerRedirectUrl!: pulumi.Output<string | undefined>;
    public readonly authenticationFlowBindingOverrides!: pulumi.Output<outputs.saml.ClientAuthenticationFlowBindingOverrides | undefined>;
    public readonly baseUrl!: pulumi.Output<string | undefined>;
    public readonly canonicalizationMethod!: pulumi.Output<string | undefined>;
    public readonly clientId!: pulumi.Output<string>;
    public readonly clientSignatureRequired!: pulumi.Output<boolean | undefined>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly encryptAssertions!: pulumi.Output<boolean | undefined>;
    public readonly encryptionCertificate!: pulumi.Output<string>;
    public /*out*/ readonly encryptionCertificateSha1!: pulumi.Output<string>;
    public readonly extraConfig!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly forceNameIdFormat!: pulumi.Output<boolean | undefined>;
    public readonly forcePostBinding!: pulumi.Output<boolean | undefined>;
    public readonly frontChannelLogout!: pulumi.Output<boolean | undefined>;
    public readonly fullScopeAllowed!: pulumi.Output<boolean | undefined>;
    public readonly idpInitiatedSsoRelayState!: pulumi.Output<string | undefined>;
    public readonly idpInitiatedSsoUrlName!: pulumi.Output<string | undefined>;
    public readonly includeAuthnStatement!: pulumi.Output<boolean | undefined>;
    public readonly loginTheme!: pulumi.Output<string | undefined>;
    public readonly logoutServicePostBindingUrl!: pulumi.Output<string | undefined>;
    public readonly logoutServiceRedirectBindingUrl!: pulumi.Output<string | undefined>;
    public readonly masterSamlProcessingUrl!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly nameIdFormat!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly rootUrl!: pulumi.Output<string | undefined>;
    public readonly signAssertions!: pulumi.Output<boolean | undefined>;
    public readonly signDocuments!: pulumi.Output<boolean | undefined>;
    public readonly signatureAlgorithm!: pulumi.Output<string | undefined>;
    public readonly signatureKeyName!: pulumi.Output<string | undefined>;
    public readonly signingCertificate!: pulumi.Output<string>;
    public /*out*/ readonly signingCertificateSha1!: pulumi.Output<string>;
    public readonly signingPrivateKey!: pulumi.Output<string>;
    public /*out*/ readonly signingPrivateKeySha1!: pulumi.Output<string>;
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
    assertionConsumerPostUrl?: pulumi.Input<string>;
    assertionConsumerRedirectUrl?: pulumi.Input<string>;
    authenticationFlowBindingOverrides?: pulumi.Input<inputs.saml.ClientAuthenticationFlowBindingOverrides>;
    baseUrl?: pulumi.Input<string>;
    canonicalizationMethod?: pulumi.Input<string>;
    clientId?: pulumi.Input<string>;
    clientSignatureRequired?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    encryptAssertions?: pulumi.Input<boolean>;
    encryptionCertificate?: pulumi.Input<string>;
    encryptionCertificateSha1?: pulumi.Input<string>;
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    forceNameIdFormat?: pulumi.Input<boolean>;
    forcePostBinding?: pulumi.Input<boolean>;
    frontChannelLogout?: pulumi.Input<boolean>;
    fullScopeAllowed?: pulumi.Input<boolean>;
    idpInitiatedSsoRelayState?: pulumi.Input<string>;
    idpInitiatedSsoUrlName?: pulumi.Input<string>;
    includeAuthnStatement?: pulumi.Input<boolean>;
    loginTheme?: pulumi.Input<string>;
    logoutServicePostBindingUrl?: pulumi.Input<string>;
    logoutServiceRedirectBindingUrl?: pulumi.Input<string>;
    masterSamlProcessingUrl?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    nameIdFormat?: pulumi.Input<string>;
    realmId?: pulumi.Input<string>;
    rootUrl?: pulumi.Input<string>;
    signAssertions?: pulumi.Input<boolean>;
    signDocuments?: pulumi.Input<boolean>;
    signatureAlgorithm?: pulumi.Input<string>;
    signatureKeyName?: pulumi.Input<string>;
    signingCertificate?: pulumi.Input<string>;
    signingCertificateSha1?: pulumi.Input<string>;
    signingPrivateKey?: pulumi.Input<string>;
    signingPrivateKeySha1?: pulumi.Input<string>;
    validRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Client resource.
 */
export interface ClientArgs {
    assertionConsumerPostUrl?: pulumi.Input<string>;
    assertionConsumerRedirectUrl?: pulumi.Input<string>;
    authenticationFlowBindingOverrides?: pulumi.Input<inputs.saml.ClientAuthenticationFlowBindingOverrides>;
    baseUrl?: pulumi.Input<string>;
    canonicalizationMethod?: pulumi.Input<string>;
    clientId: pulumi.Input<string>;
    clientSignatureRequired?: pulumi.Input<boolean>;
    description?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    encryptAssertions?: pulumi.Input<boolean>;
    encryptionCertificate?: pulumi.Input<string>;
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    forceNameIdFormat?: pulumi.Input<boolean>;
    forcePostBinding?: pulumi.Input<boolean>;
    frontChannelLogout?: pulumi.Input<boolean>;
    fullScopeAllowed?: pulumi.Input<boolean>;
    idpInitiatedSsoRelayState?: pulumi.Input<string>;
    idpInitiatedSsoUrlName?: pulumi.Input<string>;
    includeAuthnStatement?: pulumi.Input<boolean>;
    loginTheme?: pulumi.Input<string>;
    logoutServicePostBindingUrl?: pulumi.Input<string>;
    logoutServiceRedirectBindingUrl?: pulumi.Input<string>;
    masterSamlProcessingUrl?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    nameIdFormat?: pulumi.Input<string>;
    realmId: pulumi.Input<string>;
    rootUrl?: pulumi.Input<string>;
    signAssertions?: pulumi.Input<boolean>;
    signDocuments?: pulumi.Input<boolean>;
    signatureAlgorithm?: pulumi.Input<string>;
    signatureKeyName?: pulumi.Input<string>;
    signingCertificate?: pulumi.Input<string>;
    signingPrivateKey?: pulumi.Input<string>;
    validRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
}
