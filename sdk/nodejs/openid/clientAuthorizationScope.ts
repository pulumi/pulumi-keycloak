// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ClientAuthorizationScope extends pulumi.CustomResource {
    /**
     * Get an existing ClientAuthorizationScope resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientAuthorizationScopeState, opts?: pulumi.CustomResourceOptions): ClientAuthorizationScope {
        return new ClientAuthorizationScope(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientAuthorizationScope:ClientAuthorizationScope';

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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientAuthorizationScopeState | undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["iconUri"] = state ? state.iconUri : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["resourceServerId"] = state ? state.resourceServerId : undefined;
        } else {
            const args = argsOrState as ClientAuthorizationScopeArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.resourceServerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceServerId'");
            }
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["iconUri"] = args ? args.iconUri : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["resourceServerId"] = args ? args.resourceServerId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClientAuthorizationScope.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientAuthorizationScope resources.
 */
export interface ClientAuthorizationScopeState {
    displayName?: pulumi.Input<string>;
    iconUri?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    realmId?: pulumi.Input<string>;
    resourceServerId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientAuthorizationScope resource.
 */
export interface ClientAuthorizationScopeArgs {
    displayName?: pulumi.Input<string>;
    iconUri?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    realmId: pulumi.Input<string>;
    resourceServerId: pulumi.Input<string>;
}
