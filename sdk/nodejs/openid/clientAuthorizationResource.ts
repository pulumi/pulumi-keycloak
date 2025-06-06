// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class ClientAuthorizationResource extends pulumi.CustomResource {
    /**
     * Get an existing ClientAuthorizationResource resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientAuthorizationResourceState, opts?: pulumi.CustomResourceOptions): ClientAuthorizationResource {
        return new ClientAuthorizationResource(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource';

    /**
     * Returns true if the given object is an instance of ClientAuthorizationResource.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientAuthorizationResource {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientAuthorizationResource.__pulumiType;
    }

    public readonly attributes!: pulumi.Output<{[key: string]: string} | undefined>;
    public readonly displayName!: pulumi.Output<string | undefined>;
    public readonly iconUri!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly ownerManagedAccess!: pulumi.Output<boolean | undefined>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly resourceServerId!: pulumi.Output<string>;
    public readonly scopes!: pulumi.Output<string[] | undefined>;
    public readonly type!: pulumi.Output<string | undefined>;
    public readonly uris!: pulumi.Output<string[] | undefined>;

    /**
     * Create a ClientAuthorizationResource resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientAuthorizationResourceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientAuthorizationResourceArgs | ClientAuthorizationResourceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientAuthorizationResourceState | undefined;
            resourceInputs["attributes"] = state ? state.attributes : undefined;
            resourceInputs["displayName"] = state ? state.displayName : undefined;
            resourceInputs["iconUri"] = state ? state.iconUri : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["ownerManagedAccess"] = state ? state.ownerManagedAccess : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["resourceServerId"] = state ? state.resourceServerId : undefined;
            resourceInputs["scopes"] = state ? state.scopes : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uris"] = state ? state.uris : undefined;
        } else {
            const args = argsOrState as ClientAuthorizationResourceArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.resourceServerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resourceServerId'");
            }
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["iconUri"] = args ? args.iconUri : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ownerManagedAccess"] = args ? args.ownerManagedAccess : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["resourceServerId"] = args ? args.resourceServerId : undefined;
            resourceInputs["scopes"] = args ? args.scopes : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["uris"] = args ? args.uris : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClientAuthorizationResource.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientAuthorizationResource resources.
 */
export interface ClientAuthorizationResourceState {
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    displayName?: pulumi.Input<string>;
    iconUri?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    ownerManagedAccess?: pulumi.Input<boolean>;
    realmId?: pulumi.Input<string>;
    resourceServerId?: pulumi.Input<string>;
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    type?: pulumi.Input<string>;
    uris?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a ClientAuthorizationResource resource.
 */
export interface ClientAuthorizationResourceArgs {
    attributes?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    displayName?: pulumi.Input<string>;
    iconUri?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    ownerManagedAccess?: pulumi.Input<boolean>;
    realmId: pulumi.Input<string>;
    resourceServerId: pulumi.Input<string>;
    scopes?: pulumi.Input<pulumi.Input<string>[]>;
    type?: pulumi.Input<string>;
    uris?: pulumi.Input<pulumi.Input<string>[]>;
}
