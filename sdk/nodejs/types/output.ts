// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";

export interface GetClientDescriptionConverterProtocolMapper {
    config: {[key: string]: any};
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
    algorithm: string;
    certificate: string;
    kid: string;
    providerId: string;
    providerPriority: number;
    publicKey: string;
    status: string;
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
    defaultLocale: string;
    supportedLocales: string[];
}

export interface RealmOtpPolicy {
    /**
     * What hashing algorithm should be used to generate the OTP.
     */
    algorithm?: string;
    digits?: number;
    initialCounter?: number;
    lookAheadWindow?: number;
    period?: number;
    /**
     * OTP Type, totp for Time-Based One Time Password or hotp for counter base one time password
     */
    type?: string;
}

export interface RealmSecurityDefenses {
    bruteForceDetection?: outputs.RealmSecurityDefensesBruteForceDetection;
    headers?: outputs.RealmSecurityDefensesHeaders;
}

export interface RealmSecurityDefensesBruteForceDetection {
    failureResetTimeSeconds?: number;
    maxFailureWaitSeconds?: number;
    maxLoginFailures?: number;
    minimumQuickLoginWaitSeconds?: number;
    permanentLockout?: boolean;
    quickLoginCheckMilliSeconds?: number;
    waitIncrementSeconds?: number;
}

export interface RealmSecurityDefensesHeaders {
    contentSecurityPolicy?: string;
    contentSecurityPolicyReportOnly?: string;
    referrerPolicy?: string;
    strictTransportSecurity?: string;
    xContentTypeOptions?: string;
    xFrameOptions?: string;
    xRobotsTag?: string;
    xXssProtection?: string;
}

export interface RealmSmtpServer {
    auth?: outputs.RealmSmtpServerAuth;
    envelopeFrom?: string;
    from: string;
    fromDisplayName?: string;
    host: string;
    port?: string;
    replyTo?: string;
    replyToDisplayName?: string;
    ssl?: boolean;
    starttls?: boolean;
}

export interface RealmSmtpServerAuth {
    password: string;
    username: string;
}

export interface RealmUserProfileAttribute {
    annotations?: {[key: string]: string};
    displayName?: string;
    enabledWhenScopes?: string[];
    /**
     * A list of groups.
     */
    group?: string;
    name: string;
    permissions?: outputs.RealmUserProfileAttributePermissions;
    requiredForRoles?: string[];
    requiredForScopes?: string[];
    validators?: outputs.RealmUserProfileAttributeValidator[];
}

export interface RealmUserProfileAttributePermissions {
    edits: string[];
    views: string[];
}

export interface RealmUserProfileAttributeValidator {
    config?: {[key: string]: string};
    name: string;
}

export interface RealmUserProfileGroup {
    annotations?: {[key: string]: string};
    displayDescription?: string;
    displayHeader?: string;
    name: string;
}

export interface RealmWebAuthnPasswordlessPolicy {
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
     * Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
     */
    signatureAlgorithms: string[];
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: string;
}

export interface RealmWebAuthnPolicy {
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
     * Keycloak lists ES256, ES384, ES512, RS256, RS384, RS512, RS1 at the time of writing
     */
    signatureAlgorithms: string[];
    /**
     * Either required, preferred or discouraged
     */
    userVerificationRequirement?: string;
}

export interface UserFederatedIdentity {
    identityProvider: string;
    userId: string;
    userName: string;
}

export interface UserInitialPassword {
    temporary?: boolean;
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
         * Day of the week the entry will become invalid on.
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
        policy?: string;
    }

    export interface UserFederationKerberos {
        /**
         * The name of the kerberos realm, e.g. FOO.LOCAL
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
        browserId?: string;
        directGrantId?: string;
    }

    export interface ClientAuthorization {
        allowRemoteResourceManagement?: boolean;
        decisionStrategy?: string;
        keepDefaults?: boolean;
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
        browserId?: string;
        directGrantId?: string;
    }

    export interface GetClientAuthenticationFlowBindingOverride {
        browserId: string;
        directGrantId: string;
    }

}
