// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing an authentication subflow within Keycloak.
 *
 * Like authentication flows, authentication subflows are containers for authentication executions.
 * As its name implies, an authentication subflow is contained in an authentication flow.
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
 * const flow = new keycloak.authentication.Flow("flow", {
 *     realmId: realm.id,
 *     alias: "my-flow-alias",
 * });
 * const subflow = new keycloak.authentication.Subflow("subflow", {
 *     realmId: realm.id,
 *     alias: "my-subflow-alias",
 *     parentFlowAlias: flow.alias,
 *     providerId: "basic-flow",
 *     requirement: "ALTERNATIVE",
 * });
 * ```
 *
 * ## Import
 *
 * Authentication flows can be imported using the format `{{realmId}}/{{parentFlowAlias}}/{{authenticationSubflowId}}`. The authentication subflow ID is typically a GUID which is autogenerated when the subflow is created via Keycloak. Unfortunately, it is not trivial to retrieve the authentication subflow ID from the UI. The best way to do this is to visit the "Authentication" page in Keycloak, and use the network tab of your browser to view the response of the API call to `/auth/admin/realms/${realm}/authentication/flows/{flow}/executions`, which will be a list of executions, where the subflow will be. __The subflow ID is contained in the `flowID` field__ (not, as one could guess, the `id` field). Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:authentication/subflow:Subflow subflow my-realm/"Parent Flow"/3bad1172-bb5c-4a77-9615-c2606eb03081
 * ```
 */
export class Subflow extends pulumi.CustomResource {
    /**
     * Get an existing Subflow resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubflowState, opts?: pulumi.CustomResourceOptions): Subflow {
        return new Subflow(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:authentication/subflow:Subflow';

    /**
     * Returns true if the given object is an instance of Subflow.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Subflow {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Subflow.__pulumiType;
    }

    /**
     * The alias for this authentication subflow.
     */
    public readonly alias!: pulumi.Output<string>;
    /**
     * The name of the authenticator. Might be needed to be set with certain custom subflows with specific
     * authenticators. In general this will remain empty.
     */
    public readonly authenticator!: pulumi.Output<string | undefined>;
    /**
     * A description for the authentication subflow.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The alias for the parent authentication flow.
     */
    public readonly parentFlowAlias!: pulumi.Output<string>;
    /**
     * The type of authentication subflow to create. Valid choices include `basic-flow`, `form-flow`
     * and `client-flow`. Defaults to `basic-flow`.
     */
    public readonly providerId!: pulumi.Output<string | undefined>;
    /**
     * The realm that the authentication subflow exists in.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`,
     * or `DISABLED`. Defaults to `DISABLED`.
     */
    public readonly requirement!: pulumi.Output<string | undefined>;

    /**
     * Create a Subflow resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubflowArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubflowArgs | SubflowState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubflowState | undefined;
            resourceInputs["alias"] = state ? state.alias : undefined;
            resourceInputs["authenticator"] = state ? state.authenticator : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["parentFlowAlias"] = state ? state.parentFlowAlias : undefined;
            resourceInputs["providerId"] = state ? state.providerId : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["requirement"] = state ? state.requirement : undefined;
        } else {
            const args = argsOrState as SubflowArgs | undefined;
            if ((!args || args.alias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'alias'");
            }
            if ((!args || args.parentFlowAlias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parentFlowAlias'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["alias"] = args ? args.alias : undefined;
            resourceInputs["authenticator"] = args ? args.authenticator : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["parentFlowAlias"] = args ? args.parentFlowAlias : undefined;
            resourceInputs["providerId"] = args ? args.providerId : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["requirement"] = args ? args.requirement : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Subflow.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Subflow resources.
 */
export interface SubflowState {
    /**
     * The alias for this authentication subflow.
     */
    alias?: pulumi.Input<string>;
    /**
     * The name of the authenticator. Might be needed to be set with certain custom subflows with specific
     * authenticators. In general this will remain empty.
     */
    authenticator?: pulumi.Input<string>;
    /**
     * A description for the authentication subflow.
     */
    description?: pulumi.Input<string>;
    /**
     * The alias for the parent authentication flow.
     */
    parentFlowAlias?: pulumi.Input<string>;
    /**
     * The type of authentication subflow to create. Valid choices include `basic-flow`, `form-flow`
     * and `client-flow`. Defaults to `basic-flow`.
     */
    providerId?: pulumi.Input<string>;
    /**
     * The realm that the authentication subflow exists in.
     */
    realmId?: pulumi.Input<string>;
    /**
     * The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`,
     * or `DISABLED`. Defaults to `DISABLED`.
     */
    requirement?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Subflow resource.
 */
export interface SubflowArgs {
    /**
     * The alias for this authentication subflow.
     */
    alias: pulumi.Input<string>;
    /**
     * The name of the authenticator. Might be needed to be set with certain custom subflows with specific
     * authenticators. In general this will remain empty.
     */
    authenticator?: pulumi.Input<string>;
    /**
     * A description for the authentication subflow.
     */
    description?: pulumi.Input<string>;
    /**
     * The alias for the parent authentication flow.
     */
    parentFlowAlias: pulumi.Input<string>;
    /**
     * The type of authentication subflow to create. Valid choices include `basic-flow`, `form-flow`
     * and `client-flow`. Defaults to `basic-flow`.
     */
    providerId?: pulumi.Input<string>;
    /**
     * The realm that the authentication subflow exists in.
     */
    realmId: pulumi.Input<string>;
    /**
     * The requirement setting, which can be one of `REQUIRED`, `ALTERNATIVE`, `OPTIONAL`, `CONDITIONAL`,
     * or `DISABLED`. Defaults to `DISABLED`.
     */
    requirement?: pulumi.Input<string>;
}
