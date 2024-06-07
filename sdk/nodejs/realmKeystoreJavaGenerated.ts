// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing `java-keystore` Realm keystores within Keycloak.
 *
 * A realm keystore manages generated key pairs that are used by Keycloak to perform cryptographic signatures and encryption.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {realm: "my-realm"});
 * const javaKeystore = new keycloak.RealmKeystoreJavaGenerated("java_keystore", {
 *     name: "my-java-keystore",
 *     realmId: realm.id,
 *     enabled: true,
 *     active: true,
 *     keystore: "<path to your keystore>",
 *     keystorePassword: "<password for keystore>",
 *     keyAlias: "<alias for the private key>",
 *     keyPassword: "<password for the private key>",
 *     priority: 100,
 *     algorithm: "RS256",
 * });
 * ```
 *
 * ## Import
 *
 * Realm keys can be imported using realm name and keystore id, you can find it in web UI.
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated java_keystore my-realm/618cfba7-49aa-4c09-9a19-2f699b576f0b
 * ```
 */
export class RealmKeystoreJavaGenerated extends pulumi.CustomResource {
    /**
     * Get an existing RealmKeystoreJavaGenerated resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RealmKeystoreJavaGeneratedState, opts?: pulumi.CustomResourceOptions): RealmKeystoreJavaGenerated {
        return new RealmKeystoreJavaGenerated(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/realmKeystoreJavaGenerated:RealmKeystoreJavaGenerated';

    /**
     * Returns true if the given object is an instance of RealmKeystoreJavaGenerated.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RealmKeystoreJavaGenerated {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RealmKeystoreJavaGenerated.__pulumiType;
    }

    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * Intended algorithm for the key. Defaults to `RS256`
     */
    public readonly algorithm!: pulumi.Output<string | undefined>;
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Alias for the private key.
     */
    public readonly keyAlias!: pulumi.Output<string>;
    /**
     * Password for the private key.
     */
    public readonly keyPassword!: pulumi.Output<string>;
    /**
     * Path to keys file on keycloak instance.
     */
    public readonly keystore!: pulumi.Output<string>;
    /**
     * Password for the keys.
     */
    public readonly keystorePassword!: pulumi.Output<string>;
    /**
     * Display name of provider when linked in admin console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Priority for the provider. Defaults to `0`
     */
    public readonly priority!: pulumi.Output<number | undefined>;
    /**
     * The realm this keystore exists in.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a RealmKeystoreJavaGenerated resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RealmKeystoreJavaGeneratedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RealmKeystoreJavaGeneratedArgs | RealmKeystoreJavaGeneratedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RealmKeystoreJavaGeneratedState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["keyAlias"] = state ? state.keyAlias : undefined;
            resourceInputs["keyPassword"] = state ? state.keyPassword : undefined;
            resourceInputs["keystore"] = state ? state.keystore : undefined;
            resourceInputs["keystorePassword"] = state ? state.keystorePassword : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["priority"] = state ? state.priority : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as RealmKeystoreJavaGeneratedArgs | undefined;
            if ((!args || args.keyAlias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyAlias'");
            }
            if ((!args || args.keyPassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keyPassword'");
            }
            if ((!args || args.keystore === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keystore'");
            }
            if ((!args || args.keystorePassword === undefined) && !opts.urn) {
                throw new Error("Missing required property 'keystorePassword'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["keyAlias"] = args ? args.keyAlias : undefined;
            resourceInputs["keyPassword"] = args ? args.keyPassword : undefined;
            resourceInputs["keystore"] = args ? args.keystore : undefined;
            resourceInputs["keystorePassword"] = args ? args.keystorePassword : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["priority"] = args ? args.priority : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RealmKeystoreJavaGenerated.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RealmKeystoreJavaGenerated resources.
 */
export interface RealmKeystoreJavaGeneratedState {
    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Intended algorithm for the key. Defaults to `RS256`
     */
    algorithm?: pulumi.Input<string>;
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Alias for the private key.
     */
    keyAlias?: pulumi.Input<string>;
    /**
     * Password for the private key.
     */
    keyPassword?: pulumi.Input<string>;
    /**
     * Path to keys file on keycloak instance.
     */
    keystore?: pulumi.Input<string>;
    /**
     * Password for the keys.
     */
    keystorePassword?: pulumi.Input<string>;
    /**
     * Display name of provider when linked in admin console.
     */
    name?: pulumi.Input<string>;
    /**
     * Priority for the provider. Defaults to `0`
     */
    priority?: pulumi.Input<number>;
    /**
     * The realm this keystore exists in.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RealmKeystoreJavaGenerated resource.
 */
export interface RealmKeystoreJavaGeneratedArgs {
    /**
     * When `false`, key in not used for signing. Defaults to `true`.
     */
    active?: pulumi.Input<boolean>;
    /**
     * Intended algorithm for the key. Defaults to `RS256`
     */
    algorithm?: pulumi.Input<string>;
    /**
     * When `false`, key is not accessible in this realm. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Alias for the private key.
     */
    keyAlias: pulumi.Input<string>;
    /**
     * Password for the private key.
     */
    keyPassword: pulumi.Input<string>;
    /**
     * Path to keys file on keycloak instance.
     */
    keystore: pulumi.Input<string>;
    /**
     * Password for the keys.
     */
    keystorePassword: pulumi.Input<string>;
    /**
     * Display name of provider when linked in admin console.
     */
    name?: pulumi.Input<string>;
    /**
     * Priority for the provider. Defaults to `0`
     */
    priority?: pulumi.Input<number>;
    /**
     * The realm this keystore exists in.
     */
    realmId: pulumi.Input<string>;
}
