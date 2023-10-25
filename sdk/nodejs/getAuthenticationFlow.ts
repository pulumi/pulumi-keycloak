// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * This data source can be used to fetch the ID of an authentication flow within Keycloak.
 */
export function getAuthenticationFlow(args: GetAuthenticationFlowArgs, opts?: pulumi.InvokeOptions): Promise<GetAuthenticationFlowResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("keycloak:index/getAuthenticationFlow:getAuthenticationFlow", {
        "alias": args.alias,
        "realmId": args.realmId,
    }, opts);
}

/**
 * A collection of arguments for invoking getAuthenticationFlow.
 */
export interface GetAuthenticationFlowArgs {
    /**
     * The alias of the flow.
     */
    alias: string;
    /**
     * The realm the authentication flow exists in.
     */
    realmId: string;
}

/**
 * A collection of values returned by getAuthenticationFlow.
 */
export interface GetAuthenticationFlowResult {
    readonly alias: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly realmId: string;
}
/**
 * This data source can be used to fetch the ID of an authentication flow within Keycloak.
 */
export function getAuthenticationFlowOutput(args: GetAuthenticationFlowOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAuthenticationFlowResult> {
    return pulumi.output(args).apply((a: any) => getAuthenticationFlow(a, opts))
}

/**
 * A collection of arguments for invoking getAuthenticationFlow.
 */
export interface GetAuthenticationFlowOutputArgs {
    /**
     * The alias of the flow.
     */
    alias: pulumi.Input<string>;
    /**
     * The realm the authentication flow exists in.
     */
    realmId: pulumi.Input<string>;
}
