// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The provider type for the keycloak package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'keycloak';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }

    public readonly basePath!: pulumi.Output<string | undefined>;
    public readonly clientId!: pulumi.Output<string>;
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    public readonly password!: pulumi.Output<string | undefined>;
    public readonly realm!: pulumi.Output<string | undefined>;
    /**
     * Allows x509 calls using an unknown CA certificate (for development purposes)
     */
    public readonly rootCaCertificate!: pulumi.Output<string | undefined>;
    /**
     * The base URL of the Keycloak instance, before `/auth`
     */
    public readonly url!: pulumi.Output<string>;
    public readonly username!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["additionalHeaders"] = pulumi.output(args ? args.additionalHeaders : undefined).apply(JSON.stringify);
            resourceInputs["basePath"] = args ? args.basePath : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientSecret"] = args ? args.clientSecret : undefined;
            resourceInputs["clientTimeout"] = pulumi.output((args ? args.clientTimeout : undefined) ?? (utilities.getEnvNumber("KEYCLOAK_CLIENT_TIMEOUT") || 5)).apply(JSON.stringify);
            resourceInputs["initialLogin"] = pulumi.output(args ? args.initialLogin : undefined).apply(JSON.stringify);
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["rootCaCertificate"] = args ? args.rootCaCertificate : undefined;
            resourceInputs["tlsInsecureSkipVerify"] = pulumi.output(args ? args.tlsInsecureSkipVerify : undefined).apply(JSON.stringify);
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["username"] = args ? args.username : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    additionalHeaders?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    basePath?: pulumi.Input<string>;
    clientId: pulumi.Input<string>;
    clientSecret?: pulumi.Input<string>;
    /**
     * Timeout (in seconds) of the Keycloak client
     */
    clientTimeout?: pulumi.Input<number>;
    /**
     * Whether or not to login to Keycloak instance on provider initialization
     */
    initialLogin?: pulumi.Input<boolean>;
    password?: pulumi.Input<string>;
    realm?: pulumi.Input<string>;
    /**
     * Allows x509 calls using an unknown CA certificate (for development purposes)
     */
    rootCaCertificate?: pulumi.Input<string>;
    /**
     * Allows ignoring insecure certificates when set to true. Defaults to false. Disabling security check is dangerous and
     * should be avoided.
     */
    tlsInsecureSkipVerify?: pulumi.Input<boolean>;
    /**
     * The base URL of the Keycloak instance, before `/auth`
     */
    url: pulumi.Input<string>;
    username?: pulumi.Input<string>;
}
