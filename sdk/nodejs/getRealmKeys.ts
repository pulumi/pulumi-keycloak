// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
 *
 * Remarks:
 *
 * - A key must meet all filter criteria
 * - This data source may return more than one value.
 * - If no key matches the filter criteria, then an error will be returned.
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
 * const realmKeys = keycloak.getRealmKeysOutput({
 *     realmId: realm.id,
 *     algorithms: [
 *         "AES",
 *         "RS256",
 *     ],
 *     statuses: [
 *         "ACTIVE",
 *         "PASSIVE",
 *     ],
 * });
 * export const certificate = realmKeys.apply(realmKeys => realmKeys.keys?.[0]?.certificate);
 * ```
 */
export function getRealmKeys(args: GetRealmKeysArgs, opts?: pulumi.InvokeOptions): Promise<GetRealmKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("keycloak:index/getRealmKeys:getRealmKeys", {
        "algorithms": args.algorithms,
        "realmId": args.realmId,
        "statuses": args.statuses,
    }, opts);
}

/**
 * A collection of arguments for invoking getRealmKeys.
 */
export interface GetRealmKeysArgs {
    /**
     * When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
     */
    algorithms?: string[];
    /**
     * The realm from which the keys will be retrieved.
     */
    realmId: string;
    /**
     * When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
     */
    statuses?: string[];
}

/**
 * A collection of values returned by getRealmKeys.
 */
export interface GetRealmKeysResult {
    readonly algorithms?: string[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * (Computed) A list of keys that match the filter criteria. Each key has the following attributes:
     */
    readonly keys: outputs.GetRealmKeysKey[];
    readonly realmId: string;
    /**
     * Key status (string)
     */
    readonly statuses?: string[];
}
/**
 * Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
 *
 * Remarks:
 *
 * - A key must meet all filter criteria
 * - This data source may return more than one value.
 * - If no key matches the filter criteria, then an error will be returned.
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
 * const realmKeys = keycloak.getRealmKeysOutput({
 *     realmId: realm.id,
 *     algorithms: [
 *         "AES",
 *         "RS256",
 *     ],
 *     statuses: [
 *         "ACTIVE",
 *         "PASSIVE",
 *     ],
 * });
 * export const certificate = realmKeys.apply(realmKeys => realmKeys.keys?.[0]?.certificate);
 * ```
 */
export function getRealmKeysOutput(args: GetRealmKeysOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRealmKeysResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("keycloak:index/getRealmKeys:getRealmKeys", {
        "algorithms": args.algorithms,
        "realmId": args.realmId,
        "statuses": args.statuses,
    }, opts);
}

/**
 * A collection of arguments for invoking getRealmKeys.
 */
export interface GetRealmKeysOutputArgs {
    /**
     * When specified, keys will be filtered by algorithm. The algorithms can be any of `HS256`, `RS256`,`AES`, etc.
     */
    algorithms?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The realm from which the keys will be retrieved.
     */
    realmId: pulumi.Input<string>;
    /**
     * When specified, keys will be filtered by status. The statuses can be any of `ACTIVE`, `DISABLED` and `PASSIVE`.
     */
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
}
