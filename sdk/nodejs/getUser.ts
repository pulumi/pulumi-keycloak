// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can be used to fetch properties of a user within Keycloak.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const masterRealm = keycloak.getRealm({
 *     realm: "master",
 * });
 * const defaultAdminUser = masterRealm.then(masterRealm => keycloak.getUser({
 *     realmId: masterRealm.id,
 *     username: "keycloak",
 * }));
 * export const keycloakUserId = defaultAdminUser.then(defaultAdminUser => defaultAdminUser.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUser(args: GetUserArgs, opts?: pulumi.InvokeOptions): Promise<GetUserResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("keycloak:index/getUser:getUser", {
        "realmId": args.realmId,
        "username": args.username,
    }, opts);
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserArgs {
    /**
     * The realm this user belongs to.
     */
    realmId: string;
    /**
     * The unique username of this user.
     */
    username: string;
}

/**
 * A collection of values returned by getUser.
 */
export interface GetUserResult {
    /**
     * (Computed) A map representing attributes for the user
     */
    readonly attributes: {[key: string]: any};
    /**
     * (Computed) The user's email.
     */
    readonly email: string;
    /**
     * (Computed) Whether the email address was validated or not. Default to `false`.
     */
    readonly emailVerified: boolean;
    /**
     * (Computed) When false, this user cannot log in. Defaults to `true`.
     */
    readonly enabled: boolean;
    /**
     * (Computed) The user's federated identities, if applicable. This block has the following schema:
     */
    readonly federatedIdentities: string[];
    /**
     * (Computed) The user's first name.
     */
    readonly firstName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Computed) The user's last name.
     */
    readonly lastName: string;
    readonly realmId: string;
    readonly requiredActions: string[];
    readonly username: string;
}
/**
 * This data source can be used to fetch properties of a user within Keycloak.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const masterRealm = keycloak.getRealm({
 *     realm: "master",
 * });
 * const defaultAdminUser = masterRealm.then(masterRealm => keycloak.getUser({
 *     realmId: masterRealm.id,
 *     username: "keycloak",
 * }));
 * export const keycloakUserId = defaultAdminUser.then(defaultAdminUser => defaultAdminUser.id);
 * ```
 * <!--End PulumiCodeChooser -->
 */
export function getUserOutput(args: GetUserOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetUserResult> {
    return pulumi.output(args).apply((a: any) => getUser(a, opts))
}

/**
 * A collection of arguments for invoking getUser.
 */
export interface GetUserOutputArgs {
    /**
     * The realm this user belongs to.
     */
    realmId: pulumi.Input<string>;
    /**
     * The unique username of this user.
     */
    username: pulumi.Input<string>;
}
