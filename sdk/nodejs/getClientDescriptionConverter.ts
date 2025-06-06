// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * This data source uses the [ClientDescriptionConverter](https://www.keycloak.org/docs-api/latest/javadocs/org/keycloak/exportimport/ClientDescriptionConverter.html) API to convert a generic client description into a Keycloak
 * client. This data can then be used to manage the client within Keycloak.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const samlClient = keycloak.getClientDescriptionConverterOutput({
 *     realmId: realm.id,
 *     body: `\x09<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" validUntil="2021-04-17T12:41:46Z" cacheDuration="PT604800S" entityID="FakeEntityId">
 *     <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
 *         <md:KeyDescriptor use="signing">
 * \x09\x09\x09<ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
 * \x09\x09\x09\x09<ds:X509Data>
 * \x09\x09\x09\x09\x09<ds:X509Certificate>MIICyDCCAjGgAwIBAgIBADANBgkqhkiG9w0BAQ0FADCBgDELMAkGA1UEBhMCdXMx
 * \x09\x09\x09\x09\x09CzAJBgNVBAgMAklBMSQwIgYDVQQKDBt0ZXJyYWZvcm0tcHJvdmlkZXIta2V5Y2xv
 * \x09\x09\x09\x09\x09YWsxHDAaBgNVBAMME21ycGFya2Vycy5naXRodWIuaW8xIDAeBgkqhkiG9w0BCQEW
 * \x09\x09\x09\x09\x09EW1pY2hhZWxAcGFya2VyLmdnMB4XDTE5MDEwODE0NDYzNloXDTI5MDEwNTE0NDYz
 * \x09\x09\x09\x09\x09NlowgYAxCzAJBgNVBAYTAnVzMQswCQYDVQQIDAJJQTEkMCIGA1UECgwbdGVycmFm
 * \x09\x09\x09\x09\x09b3JtLXByb3ZpZGVyLWtleWNsb2FrMRwwGgYDVQQDDBNtcnBhcmtlcnMuZ2l0aHVi
 * \x09\x09\x09\x09\x09LmlvMSAwHgYJKoZIhvcNAQkBFhFtaWNoYWVsQHBhcmtlci5nZzCBnzANBgkqhkiG
 * \x09\x09\x09\x09\x099w0BAQEFAAOBjQAwgYkCgYEAxuZny7uyYxGVPtpie14gNQC4tT9sAvO2sVNDhuoe
 * \x09\x09\x09\x09\x09qIKLRpNwkHnwQmwe5OxSh9K0BPHp/DNuuVWUqvo4tniEYn3jBr7FwLYLTKojQIxj
 * \x09\x09\x09\x09\x0953S1UTT9EXq3eP5HsHMD0QnTuca2nlNYUDBm6ud2fQj0Jt5qLx86EbEC28N56IRv
 * \x09\x09\x09\x09\x09GX8CAwEAAaNQME4wHQYDVR0OBBYEFMLnbQh77j7vhGTpAhKpDhCrBsPZMB8GA1Ud
 * \x09\x09\x09\x09\x09IwQYMBaAFMLnbQh77j7vhGTpAhKpDhCrBsPZMAwGA1UdEwQFMAMBAf8wDQYJKoZI
 * \x09\x09\x09\x09\x09hvcNAQENBQADgYEAB8wGrAQY0pAfwbnYSyBt4STbebeRTu1/q1ucfrtc3qsegcd5
 * \x09\x09\x09\x09\x09n01xTR+T2uZJwqHFPpFjr4IPORiHx3+4BWCweslPD53qBjKUPXcbMO1Revjef6Tj
 * \x09\x09\x09\x09\x09K3K0AuJ94fxgXVoT61Nzu/a6Lj6RhzU/Dao9mlSbJY+YSbm+ZBpsuRUQ84s=</ds:X509Certificate>
 * \x09\x09\x09\x09</ds:X509Data>
 * \x09\x09\x09</ds:KeyInfo>
 * \x09\x09</md:KeyDescriptor>
 * \x09\x09<md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
 *         <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://localhost/acs/saml/" index="1"/>
 *     </md:SPSSODescriptor>
 * </md:EntityDescriptor>
 * `,
 * });
 * const samlClientClient = new keycloak.saml.Client("saml_client", {
 *     realmId: realm.id,
 *     clientId: samlClient.apply(samlClient => samlClient.clientId),
 * });
 * ```
 */
export function getClientDescriptionConverter(args: GetClientDescriptionConverterArgs, opts?: pulumi.InvokeOptions): Promise<GetClientDescriptionConverterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("keycloak:index/getClientDescriptionConverter:getClientDescriptionConverter", {
        "body": args.body,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getClientDescriptionConverter.
 */
export interface GetClientDescriptionConverterArgs {
    /**
     * The body of the request to convert.
     */
    body: string;
    /**
     * The realm to use for the client description converter API call.
     */
    realmId: string;
}

/**
 * A collection of values returned by getClientDescriptionConverter.
 */
export interface GetClientDescriptionConverterResult {
    readonly access: {[key: string]: string};
    readonly adminUrl: string;
    readonly alwaysDisplayInConsole: boolean;
    readonly attributes: {[key: string]: string};
    readonly authenticationFlowBindingOverrides: {[key: string]: string};
    readonly authorizationServicesEnabled: boolean;
    readonly authorizationSettings: {[key: string]: string};
    readonly baseUrl: string;
    readonly bearerOnly: boolean;
    readonly body: string;
    readonly clientAuthenticatorType: string;
    readonly clientId: string;
    readonly consentRequired: string;
    readonly defaultClientScopes: string[];
    readonly defaultRoles: string[];
    readonly description: string;
    readonly directAccessGrantsEnabled: boolean;
    readonly enabled: boolean;
    readonly frontchannelLogout: boolean;
    readonly fullScopeAllowed: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly implicitFlowEnabled: boolean;
    readonly name: string;
    readonly notBefore: number;
    readonly optionalClientScopes: string[];
    readonly origin: string;
    readonly protocol: string;
    readonly protocolMappers: outputs.GetClientDescriptionConverterProtocolMapper[];
    readonly publicClient: boolean;
    readonly realmId: string;
    readonly redirectUris: string[];
    readonly registeredNodes: {[key: string]: string};
    readonly registrationAccessToken: string;
    readonly rootUrl: string;
    readonly secret: string;
    readonly serviceAccountsEnabled: boolean;
    readonly standardFlowEnabled: boolean;
    readonly surrogateAuthRequired: boolean;
    readonly webOrigins: string[];
}
/**
 * This data source uses the [ClientDescriptionConverter](https://www.keycloak.org/docs-api/latest/javadocs/org/keycloak/exportimport/ClientDescriptionConverter.html) API to convert a generic client description into a Keycloak
 * client. This data can then be used to manage the client within Keycloak.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const samlClient = keycloak.getClientDescriptionConverterOutput({
 *     realmId: realm.id,
 *     body: `\x09<md:EntityDescriptor xmlns:md="urn:oasis:names:tc:SAML:2.0:metadata" validUntil="2021-04-17T12:41:46Z" cacheDuration="PT604800S" entityID="FakeEntityId">
 *     <md:SPSSODescriptor AuthnRequestsSigned="false" WantAssertionsSigned="false" protocolSupportEnumeration="urn:oasis:names:tc:SAML:2.0:protocol">
 *         <md:KeyDescriptor use="signing">
 * \x09\x09\x09<ds:KeyInfo xmlns:ds="http://www.w3.org/2000/09/xmldsig#">
 * \x09\x09\x09\x09<ds:X509Data>
 * \x09\x09\x09\x09\x09<ds:X509Certificate>MIICyDCCAjGgAwIBAgIBADANBgkqhkiG9w0BAQ0FADCBgDELMAkGA1UEBhMCdXMx
 * \x09\x09\x09\x09\x09CzAJBgNVBAgMAklBMSQwIgYDVQQKDBt0ZXJyYWZvcm0tcHJvdmlkZXIta2V5Y2xv
 * \x09\x09\x09\x09\x09YWsxHDAaBgNVBAMME21ycGFya2Vycy5naXRodWIuaW8xIDAeBgkqhkiG9w0BCQEW
 * \x09\x09\x09\x09\x09EW1pY2hhZWxAcGFya2VyLmdnMB4XDTE5MDEwODE0NDYzNloXDTI5MDEwNTE0NDYz
 * \x09\x09\x09\x09\x09NlowgYAxCzAJBgNVBAYTAnVzMQswCQYDVQQIDAJJQTEkMCIGA1UECgwbdGVycmFm
 * \x09\x09\x09\x09\x09b3JtLXByb3ZpZGVyLWtleWNsb2FrMRwwGgYDVQQDDBNtcnBhcmtlcnMuZ2l0aHVi
 * \x09\x09\x09\x09\x09LmlvMSAwHgYJKoZIhvcNAQkBFhFtaWNoYWVsQHBhcmtlci5nZzCBnzANBgkqhkiG
 * \x09\x09\x09\x09\x099w0BAQEFAAOBjQAwgYkCgYEAxuZny7uyYxGVPtpie14gNQC4tT9sAvO2sVNDhuoe
 * \x09\x09\x09\x09\x09qIKLRpNwkHnwQmwe5OxSh9K0BPHp/DNuuVWUqvo4tniEYn3jBr7FwLYLTKojQIxj
 * \x09\x09\x09\x09\x0953S1UTT9EXq3eP5HsHMD0QnTuca2nlNYUDBm6ud2fQj0Jt5qLx86EbEC28N56IRv
 * \x09\x09\x09\x09\x09GX8CAwEAAaNQME4wHQYDVR0OBBYEFMLnbQh77j7vhGTpAhKpDhCrBsPZMB8GA1Ud
 * \x09\x09\x09\x09\x09IwQYMBaAFMLnbQh77j7vhGTpAhKpDhCrBsPZMAwGA1UdEwQFMAMBAf8wDQYJKoZI
 * \x09\x09\x09\x09\x09hvcNAQENBQADgYEAB8wGrAQY0pAfwbnYSyBt4STbebeRTu1/q1ucfrtc3qsegcd5
 * \x09\x09\x09\x09\x09n01xTR+T2uZJwqHFPpFjr4IPORiHx3+4BWCweslPD53qBjKUPXcbMO1Revjef6Tj
 * \x09\x09\x09\x09\x09K3K0AuJ94fxgXVoT61Nzu/a6Lj6RhzU/Dao9mlSbJY+YSbm+ZBpsuRUQ84s=</ds:X509Certificate>
 * \x09\x09\x09\x09</ds:X509Data>
 * \x09\x09\x09</ds:KeyInfo>
 * \x09\x09</md:KeyDescriptor>
 * \x09\x09<md:NameIDFormat>urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified</md:NameIDFormat>
 *         <md:AssertionConsumerService Binding="urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST" Location="https://localhost/acs/saml/" index="1"/>
 *     </md:SPSSODescriptor>
 * </md:EntityDescriptor>
 * `,
 * });
 * const samlClientClient = new keycloak.saml.Client("saml_client", {
 *     realmId: realm.id,
 *     clientId: samlClient.apply(samlClient => samlClient.clientId),
 * });
 * ```
 */
export function getClientDescriptionConverterOutput(args: GetClientDescriptionConverterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetClientDescriptionConverterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("keycloak:index/getClientDescriptionConverter:getClientDescriptionConverter", {
        "body": args.body,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getClientDescriptionConverter.
 */
export interface GetClientDescriptionConverterOutputArgs {
    /**
     * The body of the request to convert.
     */
    body: pulumi.Input<string>;
    /**
     * The realm to use for the client description converter API call.
     */
    realmId: pulumi.Input<string>;
}
