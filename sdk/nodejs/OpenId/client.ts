// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export class Client extends pulumi.CustomResource {
    /**
     * Get an existing Client resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientState, opts?: pulumi.CustomResourceOptions): Client {
        return new Client(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:OpenId/client:Client';

    /**
     * Returns true if the given object is an instance of Client.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Client {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Client.__pulumiType;
    }

    public readonly accessType!: pulumi.Output<string>;
    public readonly authorization!: pulumi.Output<outputs.OpenId.ClientAuthorization | undefined>;
    public readonly clientId!: pulumi.Output<string>;
    public readonly clientSecret!: pulumi.Output<string>;
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly directAccessGrantsEnabled!: pulumi.Output<boolean | undefined>;
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    public readonly implicitFlowEnabled!: pulumi.Output<boolean | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly pkceCodeChallengeMethod!: pulumi.Output<string | undefined>;
    public readonly realmId!: pulumi.Output<string>;
    public /*out*/ readonly resourceServerId!: pulumi.Output<string>;
    public /*out*/ readonly serviceAccountUserId!: pulumi.Output<string>;
    public readonly serviceAccountsEnabled!: pulumi.Output<boolean | undefined>;
    public readonly standardFlowEnabled!: pulumi.Output<boolean | undefined>;
    public readonly validRedirectUris!: pulumi.Output<string[] | undefined>;
    public readonly webOrigins!: pulumi.Output<string[] | undefined>;

    /**
     * Create a Client resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientArgs | ClientState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as ClientState | undefined;
            inputs["accessType"] = state ? state.accessType : undefined;
            inputs["authorization"] = state ? state.authorization : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["directAccessGrantsEnabled"] = state ? state.directAccessGrantsEnabled : undefined;
            inputs["enabled"] = state ? state.enabled : undefined;
            inputs["implicitFlowEnabled"] = state ? state.implicitFlowEnabled : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pkceCodeChallengeMethod"] = state ? state.pkceCodeChallengeMethod : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["resourceServerId"] = state ? state.resourceServerId : undefined;
            inputs["serviceAccountUserId"] = state ? state.serviceAccountUserId : undefined;
            inputs["serviceAccountsEnabled"] = state ? state.serviceAccountsEnabled : undefined;
            inputs["standardFlowEnabled"] = state ? state.standardFlowEnabled : undefined;
            inputs["validRedirectUris"] = state ? state.validRedirectUris : undefined;
            inputs["webOrigins"] = state ? state.webOrigins : undefined;
        } else {
            const args = argsOrState as ClientArgs | undefined;
            if (!args || args.accessType === undefined) {
                throw new Error("Missing required property 'accessType'");
            }
            if (!args || args.clientId === undefined) {
                throw new Error("Missing required property 'clientId'");
            }
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            inputs["accessType"] = args ? args.accessType : undefined;
            inputs["authorization"] = args ? args.authorization : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["directAccessGrantsEnabled"] = args ? args.directAccessGrantsEnabled : undefined;
            inputs["enabled"] = args ? args.enabled : undefined;
            inputs["implicitFlowEnabled"] = args ? args.implicitFlowEnabled : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pkceCodeChallengeMethod"] = args ? args.pkceCodeChallengeMethod : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["serviceAccountsEnabled"] = args ? args.serviceAccountsEnabled : undefined;
            inputs["standardFlowEnabled"] = args ? args.standardFlowEnabled : undefined;
            inputs["validRedirectUris"] = args ? args.validRedirectUris : undefined;
            inputs["webOrigins"] = args ? args.webOrigins : undefined;
            inputs["resourceServerId"] = undefined /*out*/;
            inputs["serviceAccountUserId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Client.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Client resources.
 */
export interface ClientState {
    readonly accessType?: pulumi.Input<string>;
    readonly authorization?: pulumi.Input<inputs.OpenId.ClientAuthorization>;
    readonly clientId?: pulumi.Input<string>;
    readonly clientSecret?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly directAccessGrantsEnabled?: pulumi.Input<boolean>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly implicitFlowEnabled?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly pkceCodeChallengeMethod?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
    readonly resourceServerId?: pulumi.Input<string>;
    readonly serviceAccountUserId?: pulumi.Input<string>;
    readonly serviceAccountsEnabled?: pulumi.Input<boolean>;
    readonly standardFlowEnabled?: pulumi.Input<boolean>;
    readonly validRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    readonly webOrigins?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Client resource.
 */
export interface ClientArgs {
    readonly accessType: pulumi.Input<string>;
    readonly authorization?: pulumi.Input<inputs.OpenId.ClientAuthorization>;
    readonly clientId: pulumi.Input<string>;
    readonly clientSecret?: pulumi.Input<string>;
    readonly description?: pulumi.Input<string>;
    readonly directAccessGrantsEnabled?: pulumi.Input<boolean>;
    readonly enabled?: pulumi.Input<boolean>;
    readonly implicitFlowEnabled?: pulumi.Input<boolean>;
    readonly name?: pulumi.Input<string>;
    readonly pkceCodeChallengeMethod?: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
    readonly serviceAccountsEnabled?: pulumi.Input<boolean>;
    readonly standardFlowEnabled?: pulumi.Input<boolean>;
    readonly validRedirectUris?: pulumi.Input<pulumi.Input<string>[]>;
    readonly webOrigins?: pulumi.Input<pulumi.Input<string>[]>;
}
