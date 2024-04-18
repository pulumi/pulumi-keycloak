// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing realm authentication flow bindings within Keycloak.
 *
 * [Authentication flows](https://www.keycloak.org/docs/latest/server_admin/index.html#_authentication-flows) describe a sequence
 * of actions that a user or service must perform in order to be authenticated to Keycloak. The authentication flow itself
 * is a container for these actions, which are otherwise known as executions.
 *
 * Realms assign authentication flows to supported user flows such as `registration` and `browser`. This resource allows the
 * updating of realm authentication flow bindings to custom authentication flows created by `keycloak.authentication.Flow`.
 *
 * Note that you can also use the `keycloak.Realm` resource to assign authentication flow bindings at the realm level. This
 * resource is useful if you would like to create a realm and an authentication flow, and assign this flow to the realm within
 * a single run of `pulumi up`. In any case, do not attempt to use both the arguments within the `keycloak.Realm` resource
 * and this resource to manage authentication flow bindings, you should choose one or the other.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
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
 * // first execution
 * const executionOne = new keycloak.authentication.Execution("execution_one", {
 *     realmId: realm.id,
 *     parentFlowAlias: flow.alias,
 *     authenticator: "auth-cookie",
 *     requirement: "ALTERNATIVE",
 * });
 * // second execution
 * const executionTwo = new keycloak.authentication.Execution("execution_two", {
 *     realmId: realm.id,
 *     parentFlowAlias: flow.alias,
 *     authenticator: "identity-provider-redirector",
 *     requirement: "ALTERNATIVE",
 * }, {
 *     dependsOn: [executionOne],
 * });
 * const browserAuthenticationBinding = new keycloak.authentication.Bindings("browser_authentication_binding", {
 *     realmId: realm.id,
 *     browserFlow: flow.alias,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class Bindings extends pulumi.CustomResource {
    /**
     * Get an existing Bindings resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BindingsState, opts?: pulumi.CustomResourceOptions): Bindings {
        return new Bindings(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:authentication/bindings:Bindings';

    /**
     * Returns true if the given object is an instance of Bindings.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Bindings {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Bindings.__pulumiType;
    }

    /**
     * The alias of the flow to assign to the realm BrowserFlow.
     */
    public readonly browserFlow!: pulumi.Output<string>;
    /**
     * The alias of the flow to assign to the realm ClientAuthenticationFlow.
     */
    public readonly clientAuthenticationFlow!: pulumi.Output<string>;
    /**
     * The alias of the flow to assign to the realm DirectGrantFlow.
     */
    public readonly directGrantFlow!: pulumi.Output<string>;
    /**
     * The alias of the flow to assign to the realm DockerAuthenticationFlow.
     */
    public readonly dockerAuthenticationFlow!: pulumi.Output<string>;
    /**
     * The realm the authentication flow binding exists in.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * The alias of the flow to assign to the realm RegistrationFlow.
     */
    public readonly registrationFlow!: pulumi.Output<string>;
    /**
     * The alias of the flow to assign to the realm ResetCredentialsFlow.
     */
    public readonly resetCredentialsFlow!: pulumi.Output<string>;

    /**
     * Create a Bindings resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BindingsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BindingsArgs | BindingsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BindingsState | undefined;
            resourceInputs["browserFlow"] = state ? state.browserFlow : undefined;
            resourceInputs["clientAuthenticationFlow"] = state ? state.clientAuthenticationFlow : undefined;
            resourceInputs["directGrantFlow"] = state ? state.directGrantFlow : undefined;
            resourceInputs["dockerAuthenticationFlow"] = state ? state.dockerAuthenticationFlow : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["registrationFlow"] = state ? state.registrationFlow : undefined;
            resourceInputs["resetCredentialsFlow"] = state ? state.resetCredentialsFlow : undefined;
        } else {
            const args = argsOrState as BindingsArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["browserFlow"] = args ? args.browserFlow : undefined;
            resourceInputs["clientAuthenticationFlow"] = args ? args.clientAuthenticationFlow : undefined;
            resourceInputs["directGrantFlow"] = args ? args.directGrantFlow : undefined;
            resourceInputs["dockerAuthenticationFlow"] = args ? args.dockerAuthenticationFlow : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["registrationFlow"] = args ? args.registrationFlow : undefined;
            resourceInputs["resetCredentialsFlow"] = args ? args.resetCredentialsFlow : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Bindings.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Bindings resources.
 */
export interface BindingsState {
    /**
     * The alias of the flow to assign to the realm BrowserFlow.
     */
    browserFlow?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm ClientAuthenticationFlow.
     */
    clientAuthenticationFlow?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm DirectGrantFlow.
     */
    directGrantFlow?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm DockerAuthenticationFlow.
     */
    dockerAuthenticationFlow?: pulumi.Input<string>;
    /**
     * The realm the authentication flow binding exists in.
     */
    realmId?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm RegistrationFlow.
     */
    registrationFlow?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm ResetCredentialsFlow.
     */
    resetCredentialsFlow?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Bindings resource.
 */
export interface BindingsArgs {
    /**
     * The alias of the flow to assign to the realm BrowserFlow.
     */
    browserFlow?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm ClientAuthenticationFlow.
     */
    clientAuthenticationFlow?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm DirectGrantFlow.
     */
    directGrantFlow?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm DockerAuthenticationFlow.
     */
    dockerAuthenticationFlow?: pulumi.Input<string>;
    /**
     * The realm the authentication flow binding exists in.
     */
    realmId: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm RegistrationFlow.
     */
    registrationFlow?: pulumi.Input<string>;
    /**
     * The alias of the flow to assign to the realm ResetCredentialsFlow.
     */
    resetCredentialsFlow?: pulumi.Input<string>;
}
