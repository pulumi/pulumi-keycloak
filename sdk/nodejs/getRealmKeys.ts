// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## # keycloak.getRealmKeys data source
 *
 * Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
 *
 * Remarks:
 *
 * - A key must meet all filter criteria
 * - This datasource may return more than one value.
 * - If no key matches the filter criteria, then an error is returned.
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
    algorithms?: string[];
    realmId: string;
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
    readonly keys: outputs.GetRealmKeysKey[];
    readonly realmId: string;
    readonly statuses?: string[];
}
/**
 * ## # keycloak.getRealmKeys data source
 *
 * Use this data source to get the keys of a realm. Keys can be filtered by algorithm and status.
 *
 * Remarks:
 *
 * - A key must meet all filter criteria
 * - This datasource may return more than one value.
 * - If no key matches the filter criteria, then an error is returned.
 */
export function getRealmKeysOutput(args: GetRealmKeysOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetRealmKeysResult> {
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
    algorithms?: pulumi.Input<pulumi.Input<string>[]>;
    realmId: pulumi.Input<string>;
    statuses?: pulumi.Input<pulumi.Input<string>[]>;
}
