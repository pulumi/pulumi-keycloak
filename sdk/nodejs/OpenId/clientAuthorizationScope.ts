// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class ClientAuthorizationScope extends pulumi.CustomResource {
    /**
     * Get an existing ClientAuthorizationScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientAuthorizationScopeState, opts?: pulumi.CustomResourceOptions): ClientAuthorizationScope {
        return new ClientAuthorizationScope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:OpenId/clientAuthorizationScope:ClientAuthorizationScope';

    /**
     * Returns true if the given object is an instance of ClientAuthorizationScope.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientAuthorizationScope {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientAuthorizationScope.__pulumiType;
    }

    public readonly displayName!: pulumi.Output<string | undefined>;
    public readonly iconUri!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly resourceServerId!: pulumi.Output<string>;

    /**
     * Create a ClientAuthorizationScope resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientAuthorizationScopeArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientAuthorizationScopeArgs | ClientAuthorizationScopeState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClientAuthorizationScopeState | undefined;
            inputs["displayName"] = state ? state.displayName : undefined;
            inputs["iconUri"] = state ? state.iconUri : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["resourceServerId"] = state ? state.resourceServerId : undefined;
        } else {
            const args = argsOrState as ClientAuthorizationScopeArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.resourceServerId === undefined) {
                throw new Error("Missing required property 'resourceServerId'");
            }
            inputs["displayName"] = args ? args.displayName : undefined;
            inputs["iconUri"] = args ? args.iconUri : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["resourceServerId"] = args ? args.resourceServerId : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(ClientAuthorizationScope.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientAuthorizationScope resources.
 */
export interface ClientAuthorizationScopeState {
    readonly displayName?: pulumi.Input<string>;
    readonly iconUri?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
    readonly resourceServerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientAuthorizationScope resource.
 */
export interface ClientAuthorizationScopeArgs {
    readonly displayName?: pulumi.Input<string>;
    readonly iconUri?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
    readonly resourceServerId: pulumi.Input<string>;
}
