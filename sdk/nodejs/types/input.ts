// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetRealmInternationalization {
    defaultLocale?: string;
    supportedLocales?: string[];
}

export interface GetRealmInternationalizationArgs {
    defaultLocale?: pulumi.Input<string>;
    supportedLocales?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GetRealmOtpPolicy {
    algorithm?: string;
    digits?: number;
    initialCounter?: number;
    lookAheadWindow?: number;
    period?: number;
    type?: string;
}

export interface GetRealmOtpPolicyArgs {
    algorithm?: pulumi.Input<string>;
    digits?: pulumi.Input<number>;
    initialCounter?: pulumi.Input<number>;
    lookAheadWindow?: pulumi.Input<number>;
    period?: pulumi.Input<number>;
    type?: pulumi.Input<string>;
}

export interface GetRealmSecurityDefense {
    bruteForceDetections?: inputs.GetRealmSecurityDefenseBruteForceDetection[];
    headers?: inputs.GetRealmSecurityDefenseHeader[];
}

export interface GetRealmSecurityDefenseArgs {
    bruteForceDetections?: pulumi.Input<pulumi.Input<inputs.GetRealmSecurityDefenseBruteForceDetectionArgs>[]>;
    headers?: pulumi.Input<pulumi.Input<inputs.GetRealmSecurityDefenseHeaderArgs>[]>;
}

export interface GetRealmSecurityDefenseBruteForceDetection {
    failureResetTimeSeconds?: number;
    maxFailureWaitSeconds?: number;
    maxLoginFailures?: number;
    minimumQuickLoginWaitSeconds?: number;
    permanentLockout?: boolean;
    quickLoginCheckMilliSeconds?: number;
    waitIncrementSeconds?: number;
}

export interface GetRealmSecurityDefenseBruteForceDetectionArgs {
    failureResetTimeSeconds?: pulumi.Input<number>;
    maxFailureWaitSeconds?: pulumi.Input<number>;
    maxLoginFailures?: pulumi.Input<number>;
    minimumQuickLoginWaitSeconds?: pulumi.Input<number>;
    permanentLockout?: pulumi.Input<boolean>;
    quickLoginCheckMilliSeconds?: pulumi.Input<number>;
    waitIncrementSeconds?: pulumi.Input<number>;
}

export interface GetRealmSecurityDefenseHeader {
    contentSecurityPolicy?: string;
    contentSecurityPolicyReportOnly?: string;
    referrerPolicy?: string;
    strictTransportSecurity?: string;
    xContentTypeOptions?: string;
    xFrameOptions?: string;
    xRobotsTag?: string;
    xXssProtection?: string;
}

export interface GetRealmSecurityDefenseHeaderArgs {
    contentSecurityPolicy?: pulumi.Input<string>;
    contentSecurityPolicyReportOnly?: pulumi.Input<string>;
    referrerPolicy?: pulumi.Input<string>;
    strictTransportSecurity?: pulumi.Input<string>;
    xContentTypeOptions?: pulumi.Input<string>;
    xFrameOptions?: pulumi.Input<string>;
    xRobotsTag?: pulumi.Input<string>;
    xXssProtection?: pulumi.Input<string>;
}

export interface GetRealmSmtpServer {
    auths?: inputs.GetRealmSmtpServerAuth[];
    envelopeFrom?: string;
    from?: string;
    fromDisplayName?: string;
    host?: string;
    port?: string;
    replyTo?: string;
    replyToDisplayName?: string;
    ssl?: boolean;
    starttls?: boolean;
}

export interface GetRealmSmtpServerArgs {
    auths?: pulumi.Input<pulumi.Input<inputs.GetRealmSmtpServerAuthArgs>[]>;
    envelopeFrom?: pulumi.Input<string>;
    from?: pulumi.Input<string>;
    fromDisplayName?: pulumi.Input<string>;
    host?: pulumi.Input<string>;
    port?: pulumi.Input<string>;
    replyTo?: pulumi.Input<string>;
    replyToDisplayName?: pulumi.Input<string>;
    ssl?: pulumi.Input<boolean>;
    starttls?: pulumi.Input<boolean>;
}

export interface GetRealmSmtpServerAuth {
    password?: string;
    username?: string;
}

export interface GetRealmSmtpServerAuthArgs {
    password?: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}

export interface GetRealmWebAuthnPasswordlessPolicy {
    acceptableAaguids?: string[];
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference?: string;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment?: string;
    avoidSameAuthenticatorRegister?: boolean;
    createTimeout?: number;
    relyingPartyEntityName?: string;
    relyingPartyId?: string;
    /**
     * Either Yes or No
     */
    requireResidentKey?: string;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, ES384, ES512 at the time of writing
     */
    signatureAlgorithms?: string[];
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: string;
}

export interface GetRealmWebAuthnPasswordlessPolicyArgs {
    acceptableAaguids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference?: pulumi.Input<string>;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment?: pulumi.Input<string>;
    avoidSameAuthenticatorRegister?: pulumi.Input<boolean>;
    createTimeout?: pulumi.Input<number>;
    relyingPartyEntityName?: pulumi.Input<string>;
    relyingPartyId?: pulumi.Input<string>;
    /**
     * Either Yes or No
     */
    requireResidentKey?: pulumi.Input<string>;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, ES384, ES512 at the time of writing
     */
    signatureAlgorithms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: pulumi.Input<string>;
}

export interface GetRealmWebAuthnPolicy {
    acceptableAaguids?: string[];
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference?: string;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment?: string;
    avoidSameAuthenticatorRegister?: boolean;
    createTimeout?: number;
    relyingPartyEntityName?: string;
    relyingPartyId?: string;
    /**
     * Either Yes or No
     */
    requireResidentKey?: string;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, ES384, ES512 at the time of writing
     */
    signatureAlgorithms?: string[];
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: string;
}

export interface GetRealmWebAuthnPolicyArgs {
    acceptableAaguids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference?: pulumi.Input<string>;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment?: pulumi.Input<string>;
    avoidSameAuthenticatorRegister?: pulumi.Input<boolean>;
    createTimeout?: pulumi.Input<number>;
    relyingPartyEntityName?: pulumi.Input<string>;
    relyingPartyId?: pulumi.Input<string>;
    /**
     * Either Yes or No
     */
    requireResidentKey?: pulumi.Input<string>;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, ES384, ES512 at the time of writing
     */
    signatureAlgorithms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: pulumi.Input<string>;
}

export interface GroupPermissionsManageMembersScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GroupPermissionsManageMembershipScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GroupPermissionsManageScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GroupPermissionsViewMembersScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface GroupPermissionsViewScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface RealmInternationalization {
    /**
     * The locale to use by default. This locale code must be present within the `supportedLocales` list.
     */
    defaultLocale: pulumi.Input<string>;
    /**
     * A list of [ISO 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes) locale codes that the realm should support.
     */
    supportedLocales: pulumi.Input<pulumi.Input<string>[]>;
}

export interface RealmOtpPolicy {
    /**
     * What hashing algorithm should be used to generate the OTP, Valid options are `HmacSHA1`,`HmacSHA256` and `HmacSHA512`. Defaults to `HmacSHA1`.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * How many digits the OTP have. Defaults to `6`.
     */
    digits?: pulumi.Input<number>;
    /**
     * What should the initial counter value be. Defaults to `2`.
     */
    initialCounter?: pulumi.Input<number>;
    /**
     * How far ahead should the server look just in case the token generator and server are out of time sync or counter sync. Defaults to `1`.
     */
    lookAheadWindow?: pulumi.Input<number>;
    /**
     * How many seconds should an OTP token be valid. Defaults to `30`.
     */
    period?: pulumi.Input<number>;
    /**
     * One Time Password Type, supported Values are `totp` for Time-Based One Time Password and `hotp` for Counter Based. Defaults to `totp`.
     */
    type?: pulumi.Input<string>;
}

export interface RealmSecurityDefenses {
    bruteForceDetection?: pulumi.Input<inputs.RealmSecurityDefensesBruteForceDetection>;
    headers?: pulumi.Input<inputs.RealmSecurityDefensesHeaders>;
}

export interface RealmSecurityDefensesBruteForceDetection {
    /**
     * When will failure count be reset?
     */
    failureResetTimeSeconds?: pulumi.Input<number>;
    maxFailureWaitSeconds?: pulumi.Input<number>;
    /**
     * How many failures before wait is triggered.
     */
    maxLoginFailures?: pulumi.Input<number>;
    /**
     * How long to wait after a quick login failure.
     * - `maxFailureWaitSeconds ` - (Optional) Max. time a user will be locked out.
     */
    minimumQuickLoginWaitSeconds?: pulumi.Input<number>;
    /**
     * When `true`, this will lock the user permanently when the user exceeds the maximum login failures.
     */
    permanentLockout?: pulumi.Input<boolean>;
    /**
     * Configures the amount of time, in milliseconds, for consecutive failures to lock a user out.
     */
    quickLoginCheckMilliSeconds?: pulumi.Input<number>;
    /**
     * This represents the amount of time a user should be locked out when the login failure threshold has been met.
     */
    waitIncrementSeconds?: pulumi.Input<number>;
}

export interface RealmSecurityDefensesHeaders {
    /**
     * Sets the Content Security Policy, which can be used for prevent pages from being included by non-origin iframes. More information can be found in the [W3C-CSP](https://www.w3.org/TR/CSP/) Abstract.
     */
    contentSecurityPolicy?: pulumi.Input<string>;
    /**
     * Used for testing Content Security Policies.
     */
    contentSecurityPolicyReportOnly?: pulumi.Input<string>;
    /**
     * The Referrer-Policy HTTP header controls how much referrer information (sent with the Referer header) should be included with requests.
     */
    referrerPolicy?: pulumi.Input<string>;
    /**
     * The Script-Transport-Security HTTP header tells browsers to always use HTTPS.
     */
    strictTransportSecurity?: pulumi.Input<string>;
    /**
     * Sets the X-Content-Type-Options, which can be used for prevent MIME-sniffing a response away from the declared content-type
     */
    xContentTypeOptions?: pulumi.Input<string>;
    /**
     * Sets the x-frame-option, which can be used to prevent pages from being included by non-origin iframes. More information can be found in the [RFC7034](https://tools.ietf.org/html/rfc7034)
     */
    xFrameOptions?: pulumi.Input<string>;
    /**
     * Prevent pages from appearing in search engines.
     */
    xRobotsTag?: pulumi.Input<string>;
    /**
     * This header configures the Cross-site scripting (XSS) filter in your browser.
     */
    xXssProtection?: pulumi.Input<string>;
}

export interface RealmSmtpServer {
    /**
     * Enables authentication to the SMTP server.  This block supports the following arguments:
     */
    auth?: pulumi.Input<inputs.RealmSmtpServerAuth>;
    /**
     * The email address uses for bounces.
     */
    envelopeFrom?: pulumi.Input<string>;
    /**
     * The email address for the sender.
     */
    from: pulumi.Input<string>;
    /**
     * The display name of the sender email address.
     */
    fromDisplayName?: pulumi.Input<string>;
    /**
     * The host of the SMTP server.
     */
    host: pulumi.Input<string>;
    /**
     * The port of the SMTP server (defaults to 25).
     */
    port?: pulumi.Input<string>;
    /**
     * The "reply to" email address.
     */
    replyTo?: pulumi.Input<string>;
    /**
     * The display name of the "reply to" email address.
     */
    replyToDisplayName?: pulumi.Input<string>;
    /**
     * When `true`, enables SSL. Defaults to `false`.
     */
    ssl?: pulumi.Input<boolean>;
    /**
     * When `true`, enables StartTLS. Defaults to `false`.
     */
    starttls?: pulumi.Input<boolean>;
}

export interface RealmSmtpServerAuth {
    /**
     * The SMTP server password.
     */
    password: pulumi.Input<string>;
    /**
     * The SMTP server username.
     */
    username: pulumi.Input<string>;
}

export interface RealmUserProfileAttribute {
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The display name of the attribute.
     */
    displayName?: pulumi.Input<string>;
    /**
     * A list of scopes. The attribute will only be enabled when these scopes are requested by clients.
     */
    enabledWhenScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of groups.
     */
    group?: pulumi.Input<string>;
    name: pulumi.Input<string>;
    /**
     * The permissions configuration information.
     */
    permissions?: pulumi.Input<inputs.RealmUserProfileAttributePermissions>;
    /**
     * A list of roles for which the attribute will be required.
     */
    requiredForRoles?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of scopes for which the attribute will be required.
     */
    requiredForScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of validators for the attribute.
     */
    validators?: pulumi.Input<pulumi.Input<inputs.RealmUserProfileAttributeValidator>[]>;
}

export interface RealmUserProfileAttributePermissions {
    /**
     * A list of profiles that will be able to edit the attribute. One of `admin`, `user`.
     */
    edits: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of profiles that will be able to view the attribute. One of `admin`, `user`.
     */
    views: pulumi.Input<pulumi.Input<string>[]>;
}

export interface RealmUserProfileAttributeValidator {
    /**
     * A map defining the configuration of the validator. Values can be a String or a json object.
     */
    config?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    name: pulumi.Input<string>;
}

export interface RealmUserProfileGroup {
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The display description of the group.
     */
    displayDescription?: pulumi.Input<string>;
    /**
     * The display header of the group.
     */
    displayHeader?: pulumi.Input<string>;
    name: pulumi.Input<string>;
}

export interface RealmWebAuthnPasswordlessPolicy {
    /**
     * A set of AAGUIDs for which an authenticator can be registered.
     */
    acceptableAaguids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference?: pulumi.Input<string>;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment?: pulumi.Input<string>;
    /**
     * When `true`, Keycloak will avoid registering the authenticator for WebAuthn if it has already been registered. Defaults to `false`.
     */
    avoidSameAuthenticatorRegister?: pulumi.Input<boolean>;
    /**
     * The timeout value for creating a user's public key credential in seconds. When set to `0`, this timeout option is not adapted. Defaults to `0`.
     */
    createTimeout?: pulumi.Input<number>;
    /**
     * A human-readable server name for the WebAuthn Relying Party. Defaults to `keycloak`.
     */
    relyingPartyEntityName?: pulumi.Input<string>;
    /**
     * The WebAuthn relying party ID.
     */
    relyingPartyId?: pulumi.Input<string>;
    /**
     * Either Yes or No
     */
    requireResidentKey?: pulumi.Input<string>;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
     */
    signatureAlgorithms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: pulumi.Input<string>;
}

export interface RealmWebAuthnPolicy {
    /**
     * A set of AAGUIDs for which an authenticator can be registered.
     */
    acceptableAaguids?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either none, indirect or direct
     */
    attestationConveyancePreference?: pulumi.Input<string>;
    /**
     * Either platform or cross-platform
     */
    authenticatorAttachment?: pulumi.Input<string>;
    /**
     * When `true`, Keycloak will avoid registering the authenticator for WebAuthn if it has already been registered. Defaults to `false`.
     */
    avoidSameAuthenticatorRegister?: pulumi.Input<boolean>;
    /**
     * The timeout value for creating a user's public key credential in seconds. When set to `0`, this timeout option is not adapted. Defaults to `0`.
     */
    createTimeout?: pulumi.Input<number>;
    /**
     * A human-readable server name for the WebAuthn Relying Party. Defaults to `keycloak`.
     */
    relyingPartyEntityName?: pulumi.Input<string>;
    /**
     * The WebAuthn relying party ID.
     */
    relyingPartyId?: pulumi.Input<string>;
    /**
     * Either Yes or No
     */
    requireResidentKey?: pulumi.Input<string>;
    /**
     * Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
     */
    signatureAlgorithms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: pulumi.Input<string>;
}

export interface UserFederatedIdentity {
    /**
     * The name of the identity provider
     */
    identityProvider: pulumi.Input<string>;
    /**
     * The ID of the user defined in the identity provider
     */
    userId: pulumi.Input<string>;
    /**
     * The username of the user defined in the identity provider
     */
    userName: pulumi.Input<string>;
}

export interface UserInitialPassword {
    /**
     * If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
     */
    temporary?: pulumi.Input<boolean>;
    /**
     * The initial password.
     */
    value: pulumi.Input<string>;
}

export interface UsersPermissionsImpersonateScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface UsersPermissionsManageGroupMembershipScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface UsersPermissionsManageScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface UsersPermissionsMapRolesScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface UsersPermissionsUserImpersonatedScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

export interface UsersPermissionsViewScope {
    decisionStrategy?: pulumi.Input<string>;
    description?: pulumi.Input<string>;
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}
export namespace ldap {
    export interface UserFederationCache {
        /**
         * Day of the week the entry will become invalid on
         */
        evictionDay?: pulumi.Input<number>;
        /**
         * Hour of day the entry will become invalid on.
         */
        evictionHour?: pulumi.Input<number>;
        /**
         * Minute of day the entry will become invalid on.
         */
        evictionMinute?: pulumi.Input<number>;
        /**
         * Max lifespan of cache entry (duration string).
         */
        maxLifespan?: pulumi.Input<string>;
        /**
         * Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
         */
        policy?: pulumi.Input<string>;
    }

    export interface UserFederationKerberos {
        /**
         * The name of the kerberos realm, e.g. FOO.LOCAL.
         */
        kerberosRealm: pulumi.Input<string>;
        /**
         * Path to the kerberos keytab file on the server with credentials of the service principal.
         */
        keyTab: pulumi.Input<string>;
        /**
         * The kerberos server principal, e.g. 'HTTP/host.foo.com@FOO.LOCAL'.
         */
        serverPrincipal: pulumi.Input<string>;
        /**
         * Use kerberos login module instead of ldap service api. Defaults to `false`.
         */
        useKerberosForPasswordAuthentication?: pulumi.Input<boolean>;
    }
}

export namespace openid {
    export interface ClientAuthenticationFlowBindingOverrides {
        /**
         * Browser flow id, (flow needs to exist)
         */
        browserId?: pulumi.Input<string>;
        /**
         * Direct grant flow id (flow needs to exist)
         */
        directGrantId?: pulumi.Input<string>;
    }

    export interface ClientAuthorization {
        /**
         * When `true`, resources can be managed remotely by the resource server. Defaults to `false`.
         */
        allowRemoteResourceManagement?: pulumi.Input<boolean>;
        /**
         * Dictates how the policies associated with a given permission are evaluated and how a final decision is obtained. Could be one of `AFFIRMATIVE`, `CONSENSUS`, or `UNANIMOUS`. Applies to permissions.
         */
        decisionStrategy?: pulumi.Input<string>;
        /**
         * When `true`, defaults set by Keycloak will be respected. Defaults to `false`.
         */
        keepDefaults?: pulumi.Input<boolean>;
        /**
         * Dictates how policies are enforced when evaluating authorization requests. Can be one of `ENFORCING`, `PERMISSIVE`, or `DISABLED`.
         */
        policyEnforcementMode: pulumi.Input<string>;
    }

    export interface ClientGroupPolicyGroup {
        extendChildren: pulumi.Input<boolean>;
        id: pulumi.Input<string>;
        path: pulumi.Input<string>;
    }

    export interface ClientPermissionsConfigureScope {
        decisionStrategy?: pulumi.Input<string>;
        description?: pulumi.Input<string>;
        policies?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface ClientPermissionsManageScope {
        decisionStrategy?: pulumi.Input<string>;
        description?: pulumi.Input<string>;
        policies?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface ClientPermissionsMapRolesClientScopeScope {
        decisionStrategy?: pulumi.Input<string>;
        description?: pulumi.Input<string>;
        policies?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface ClientPermissionsMapRolesCompositeScope {
        decisionStrategy?: pulumi.Input<string>;
        description?: pulumi.Input<string>;
        policies?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface ClientPermissionsMapRolesScope {
        decisionStrategy?: pulumi.Input<string>;
        description?: pulumi.Input<string>;
        policies?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface ClientPermissionsTokenExchangeScope {
        decisionStrategy?: pulumi.Input<string>;
        description?: pulumi.Input<string>;
        policies?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface ClientPermissionsViewScope {
        decisionStrategy?: pulumi.Input<string>;
        description?: pulumi.Input<string>;
        policies?: pulumi.Input<pulumi.Input<string>[]>;
    }

    export interface ClientRolePolicyRole {
        id: pulumi.Input<string>;
        required: pulumi.Input<boolean>;
    }

}

export namespace saml {
    export interface ClientAuthenticationFlowBindingOverrides {
        /**
         * Browser flow id, (flow needs to exist)
         */
        browserId?: pulumi.Input<string>;
        /**
         * Direct grant flow id (flow needs to exist)
         */
        directGrantId?: pulumi.Input<string>;
    }

}
