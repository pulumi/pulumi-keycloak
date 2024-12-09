// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetClientDescriptionConverterProtocolMapper {
    config: {[key: string]: string};
    id: string;
    name: string;
    protocol: string;
    protocolMapper: string;
}

export interface GetRealmInternationalization {
    defaultLocale: string;
    supportedLocales: string[];
}

export interface GetRealmKeysKey {
    /**
     * Key algorithm (string)
     */
    algorithm: string;
    /**
     * Key certificate (string)
     */
    certificate: string;
    /**
     * Key ID (string)
     */
    kid: string;
    /**
     * Key provider ID (string)
     */
    providerId: string;
    /**
     * Key provider priority (int64)
     */
    providerPriority: number;
    /**
     * Key public key (string)
     */
    publicKey: string;
    /**
     * When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
     */
    status: string;
    /**
     * Key type (string)
     */
    type: string;
}

export interface GetRealmOtpPolicy {
    algorithm: string;
    digits: number;
    initialCounter: number;
    lookAheadWindow: number;
    period: number;
    type: string;
}

export interface GetRealmSecurityDefense {
    bruteForceDetections: outputs.GetRealmSecurityDefenseBruteForceDetection[];
    headers: outputs.GetRealmSecurityDefenseHeader[];
}

export interface GetRealmSecurityDefenseBruteForceDetection {
    failureResetTimeSeconds: number;
    maxFailureWaitSeconds: number;
    maxLoginFailures: number;
    minimumQuickLoginWaitSeconds: number;
    permanentLockout: boolean;
    quickLoginCheckMilliSeconds: number;
    waitIncrementSeconds: number;
}

export interface GetRealmSecurityDefenseHeader {
    contentSecurityPolicy: string;
    contentSecurityPolicyReportOnly: string;
    referrerPolicy: string;
    strictTransportSecurity: string;
    xContentTypeOptions: string;
    xFrameOptions: string;
    xRobotsTag: string;
    xXssProtection: string;
}

export interface GetRealmSmtpServer {
    auths: outputs.GetRealmSmtpServerAuth[];
    envelopeFrom: string;
    from: string;
    fromDisplayName: string;
    host: string;
    port: string;
    replyTo: string;
    replyToDisplayName: string;
    ssl: boolean;
    starttls: boolean;
}

export interface GetRealmSmtpServerAuth {
    password: string;
    username: string;
}

export interface GetRealmWebAuthnPasswordlessPolicy {
    acceptableAaguids: string[];
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference: string;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment: string;
    avoidSameAuthenticatorRegister: boolean;
    createTimeout: number;
    relyingPartyEntityName: string;
    relyingPartyId: string;
    /**
     * Either Yes or No
     */
    requireResidentKey: string;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, ES384, ES512 at the time of writing
     */
    signatureAlgorithms: string[];
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement: string;
}

export interface GetRealmWebAuthnPolicy {
    acceptableAaguids: string[];
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference: string;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment: string;
    avoidSameAuthenticatorRegister: boolean;
    createTimeout: number;
    relyingPartyEntityName: string;
    relyingPartyId: string;
    /**
     * Either Yes or No
     */
    requireResidentKey: string;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, ES384, ES512 at the time of writing
     */
    signatureAlgorithms: string[];
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement: string;
}

export interface GroupPermissionsManageMembersScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface GroupPermissionsManageMembershipScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface GroupPermissionsManageScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface GroupPermissionsViewMembersScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface GroupPermissionsViewScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface RealmInternationalization {
    /**
     * The locale to use by default. This locale code must be present within the `supportedLocales` list.
     */
    defaultLocale: string;
    /**
     * A list of [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) locale codes that the realm should support.
     */
    supportedLocales: string[];
}

export interface RealmOtpPolicy {
    /**
     * What hashing algorithm should be used to generate the OTP, Valid options are `HmacSHA1`,`HmacSHA256` and `HmacSHA512`. Defaults to `HmacSHA1`.
     */
    algorithm?: string;
    /**
     * How many digits the OTP have. Defaults to `6`.
     */
    digits?: number;
    /**
     * What should the initial counter value be. Defaults to `2`.
     */
    initialCounter?: number;
    /**
     * How far ahead should the server look just in case the token generator and server are out of time sync or counter sync. Defaults to `1`.
     */
    lookAheadWindow?: number;
    /**
     * How many seconds should an OTP token be valid. Defaults to `30`.
     */
    period?: number;
    /**
     * One Time Password Type, supported Values are `totp` for Time-Based One Time Password and `hotp` for Counter Based. Defaults to `totp`.
     */
    type?: string;
}

export interface RealmSecurityDefenses {
    bruteForceDetection?: outputs.RealmSecurityDefensesBruteForceDetection;
    headers?: outputs.RealmSecurityDefensesHeaders;
}

export interface RealmSecurityDefensesBruteForceDetection {
    /**
     * When will failure count be reset?
     */
    failureResetTimeSeconds?: number;
    maxFailureWaitSeconds?: number;
    /**
     * How many failures before wait is triggered.
     */
    maxLoginFailures?: number;
    /**
     * How long to wait after a quick login failure.
     * - `maxFailureWaitSeconds ` - (Optional) Max. time a user will be locked out.
     */
    minimumQuickLoginWaitSeconds?: number;
    /**
     * When `true`, this will lock the user permanently when the user exceeds the maximum login failures.
     */
    permanentLockout?: boolean;
    /**
     * Configures the amount of time, in milliseconds, for consecutive failures to lock a user out.
     */
    quickLoginCheckMilliSeconds?: number;
    /**
     * This represents the amount of time a user should be locked out when the login failure threshold has been met.
     */
    waitIncrementSeconds?: number;
}

export interface RealmSecurityDefensesHeaders {
    /**
     * Sets the Content Security Policy, which can be used for prevent pages from being included by non-origin iframes. More information can be found in the [W3C-CSP](https://www.w3.org/TR/CSP/) Abstract.
     */
    contentSecurityPolicy?: string;
    /**
     * Used for testing Content Security Policies.
     */
    contentSecurityPolicyReportOnly?: string;
    /**
     * The Referrer-Policy HTTP header controls how much referrer information (sent with the Referer header) should be included with requests.
     */
    referrerPolicy?: string;
    /**
     * The Script-Transport-Security HTTP header tells browsers to always use HTTPS.
     */
    strictTransportSecurity?: string;
    /**
     * Sets the X-Content-Type-Options, which can be used for prevent MIME-sniffing a response away from the declared content-type
     */
    xContentTypeOptions?: string;
    /**
     * Sets the x-frame-option, which can be used to prevent pages from being included by non-origin iframes. More information can be found in the [RFC7034](https://tools.ietf.org/html/rfc7034)
     */
    xFrameOptions?: string;
    /**
     * Prevent pages from appearing in search engines.
     */
    xRobotsTag?: string;
    /**
     * This header configures the Cross-site scripting (XSS) filter in your browser.
     */
    xXssProtection?: string;
}

export interface RealmSmtpServer {
    /**
     * Enables authentication to the SMTP server.  This block supports the following arguments:
     */
    auth?: outputs.RealmSmtpServerAuth;
    /**
     * The email address uses for bounces.
     */
    envelopeFrom?: string;
    /**
     * The email address for the sender.
     */
    from: string;
    /**
     * The display name of the sender email address.
     */
    fromDisplayName?: string;
    /**
     * The host of the SMTP server.
     */
    host: string;
    /**
     * The port of the SMTP server (defaults to 25).
     */
    port?: string;
    /**
     * The "reply to" email address.
     */
    replyTo?: string;
    /**
     * The display name of the "reply to" email address.
     */
    replyToDisplayName?: string;
    /**
     * When `true`, enables SSL. Defaults to `false`.
     */
    ssl?: boolean;
    /**
     * When `true`, enables StartTLS. Defaults to `false`.
     */
    starttls?: boolean;
}

export interface RealmSmtpServerAuth {
    /**
     * The SMTP server password.
     */
    password: string;
    /**
     * The SMTP server username.
     */
    username: string;
}

export interface RealmUserProfileAttribute {
    annotations?: {[key: string]: string};
    /**
     * The display name of the attribute.
     */
    displayName?: string;
    /**
     * A list of scopes. The attribute will only be enabled when these scopes are requested by clients.
     */
    enabledWhenScopes?: string[];
    /**
     * A list of groups.
     */
    group?: string;
    name: string;
    /**
     * The permissions configuration information.
     */
    permissions?: outputs.RealmUserProfileAttributePermissions;
    /**
     * A list of roles for which the attribute will be required.
     */
    requiredForRoles?: string[];
    /**
     * A list of scopes for which the attribute will be required.
     */
    requiredForScopes?: string[];
    /**
     * A list of validators for the attribute.
     */
    validators?: outputs.RealmUserProfileAttributeValidator[];
}

export interface RealmUserProfileAttributePermissions {
    /**
     * A list of profiles that will be able to edit the attribute. One of `admin`, `user`.
     */
    edits: string[];
    /**
     * A list of profiles that will be able to view the attribute. One of `admin`, `user`.
     */
    views: string[];
}

export interface RealmUserProfileAttributeValidator {
    /**
     * A map defining the configuration of the validator. Values can be a String or a json object.
     */
    config?: {[key: string]: string};
    name: string;
}

export interface RealmUserProfileGroup {
    annotations?: {[key: string]: string};
    /**
     * The display description of the group.
     */
    displayDescription?: string;
    /**
     * The display header of the group.
     */
    displayHeader?: string;
    name: string;
}

export interface RealmWebAuthnPasswordlessPolicy {
    /**
     * A set of AAGUIDs for which an authenticator can be registered.
     */
    acceptableAaguids?: string[];
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference?: string;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment?: string;
    /**
     * When `true`, Keycloak will avoid registering the authenticator for WebAuthn if it has already been registered. Defaults to `false`.
     */
    avoidSameAuthenticatorRegister?: boolean;
    /**
     * The timeout value for creating a user's public key credential in seconds. When set to `0`, this timeout option is not adapted. Defaults to `0`.
     */
    createTimeout?: number;
    /**
     * A human readable server name for the WebAuthn Relying Party. Defaults to `keycloak`.
     */
    relyingPartyEntityName?: string;
    /**
     * The WebAuthn relying party ID.
     */
    relyingPartyId?: string;
    /**
     * Either Yes or No
     */
    requireResidentKey?: string;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
     */
    signatureAlgorithms: string[];
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: string;
}

export interface RealmWebAuthnPolicy {
    /**
     * A set of AAGUIDs for which an authenticator can be registered.
     */
    acceptableAaguids?: string[];
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference?: string;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment?: string;
    /**
     * When `true`, Keycloak will avoid registering the authenticator for WebAuthn if it has already been registered. Defaults to `false`.
     */
    avoidSameAuthenticatorRegister?: boolean;
    /**
     * The timeout value for creating a user's public key credential in seconds. When set to `0`, this timeout option is not adapted. Defaults to `0`.
     */
    createTimeout?: number;
    /**
     * A human readable server name for the WebAuthn Relying Party. Defaults to `keycloak`.
     */
    relyingPartyEntityName?: string;
    /**
     * The WebAuthn relying party ID.
     */
    relyingPartyId?: string;
    /**
     * Either Yes or No
     */
    requireResidentKey?: string;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
     */
    signatureAlgorithms: string[];
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: string;
}

export interface UserFederatedIdentity {
    /**
     * The name of the identity provider
     */
    identityProvider: string;
    /**
     * The ID of the user defined in the identity provider
     */
    userId: string;
    /**
     * The user name of the user defined in the identity provider
     */
    userName: string;
}

export interface UserInitialPassword {
    /**
     * If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
     */
    temporary?: boolean;
    /**
     * The initial password.
     */
    value: string;
}

export interface UsersPermissionsImpersonateScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface UsersPermissionsManageGroupMembershipScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface UsersPermissionsManageScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface UsersPermissionsMapRolesScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface UsersPermissionsUserImpersonatedScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export interface UsersPermissionsViewScope {
    decisionStrategy?: string;
    description?: string;
    policies?: string[];
}

export namespace ldap {
    export interface UserFederationCache {
        /**
         * Day of the week the entry will become invalid on
         */
        evictionDay?: number;
        /**
         * Hour of day the entry will become invalid on.
         */
        evictionHour?: number;
        /**
         * Minute of day the entry will become invalid on.
         */
        evictionMinute?: number;
        /**
         * Max lifespan of cache entry (duration string).
         */
        maxLifespan?: string;
        /**
         * Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
         */
        policy?: string;
    }

    export interface UserFederationKerberos {
        /**
         * The name of the kerberos realm, e.g. FOO.LOCAL.
         */
        kerberosRealm: string;
        /**
         * Path to the kerberos keytab file on the server with credentials of the service principal.
         */
        keyTab: string;
        /**
         * The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
         */
        serverPrincipal: string;
        /**
         * Use kerberos login module instead of ldap service api. Defaults to `false`.
         */
        useKerberosForPasswordAuthentication?: boolean;
    }

}

export namespace openid {
    export interface ClientAuthenticationFlowBindingOverrides {
        /**
         * Browser flow id, (flow needs to exist)
         */
        browserId?: string;
        /**
         * Direct grant flow id (flow needs to exist)
         */
        directGrantId?: string;
    }

    export interface ClientAuthorization {
        /**
         * When `true`, resources can be managed remotely by the resource server. Defaults to `false`.
         */
        allowRemoteResourceManagement?: boolean;
        /**
         * Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
         */
        decisionStrategy?: string;
        /**
         * When `true`, defaults set by Keycloak will be respected. Defaults to `false`.
         */
        keepDefaults?: boolean;
        /**
         * Dictates how policies are enforced when evaluating authorization requests. Can be one of `ENFORCING`, `PERMISSIVE`, or `DISABLED`.
         */
        policyEnforcementMode: string;
    }

    export interface ClientGroupPolicyGroup {
        extendChildren: boolean;
        id: string;
        path: string;
    }

    export interface ClientPermissionsConfigureScope {
        decisionStrategy?: string;
        description?: string;
        policies?: string[];
    }

    export interface ClientPermissionsManageScope {
        decisionStrategy?: string;
        description?: string;
        policies?: string[];
    }

    export interface ClientPermissionsMapRolesClientScopeScope {
        decisionStrategy?: string;
        description?: string;
        policies?: string[];
    }

    export interface ClientPermissionsMapRolesCompositeScope {
        decisionStrategy?: string;
        description?: string;
        policies?: string[];
    }

    export interface ClientPermissionsMapRolesScope {
        decisionStrategy?: string;
        description?: string;
        policies?: string[];
    }

    export interface ClientPermissionsTokenExchangeScope {
        decisionStrategy?: string;
        description?: string;
        policies?: string[];
    }

    export interface ClientPermissionsViewScope {
        decisionStrategy?: string;
        description?: string;
        policies?: string[];
    }

    export interface ClientRolePolicyRole {
        id: string;
        required: boolean;
    }

    export interface GetClientAuthenticationFlowBindingOverride {
        browserId: string;
        directGrantId: string;
    }

    export interface GetClientAuthorization {
        allowRemoteResourceManagement: boolean;
        decisionStrategy: string;
        keepDefaults: boolean;
        policyEnforcementMode: string;
    }

    export interface GetClientServiceAccountUserFederatedIdentity {
        identityProvider: string;
        userId: string;
        userName: string;
    }

}

export namespace saml {
    export interface ClientAuthenticationFlowBindingOverrides {
        /**
         * Browser flow id, (flow needs to exist)
         */
        browserId?: string;
        /**
         * Direct grant flow id (flow needs to exist)
         */
        directGrantId?: string;
    }

    export interface GetClientAuthenticationFlowBindingOverride {
        browserId: string;
        directGrantId: string;
    }

}
