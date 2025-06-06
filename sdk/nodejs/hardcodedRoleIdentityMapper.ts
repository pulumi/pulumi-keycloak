// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing hardcoded role mappers for Keycloak identity provider.
 *
 * The identity provider hardcoded role mapper grants a specified Keycloak role to each Keycloak user from the LDAP provider.
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
 * const oidc = new keycloak.oidc.IdentityProvider("oidc", {
 *     realm: realm.id,
 *     alias: "my-idp",
 *     authorizationUrl: "https://authorizationurl.com",
 *     clientId: "clientID",
 *     clientSecret: "clientSecret",
 *     tokenUrl: "https://tokenurl.com",
 * });
 * const realmRole = new keycloak.Role("realm_role", {
 *     realmId: realm.id,
 *     name: "my-realm-role",
 *     description: "My Realm Role",
 * });
 * const oidcHardcodedRoleIdentityMapper = new keycloak.HardcodedRoleIdentityMapper("oidc", {
 *     realm: realm.id,
 *     name: "hardcodedRole",
 *     identityProviderAlias: oidc.alias,
 *     role: "my-realm-role",
 *     extraConfig: {
 *         syncMode: "INHERIT",
 *     },
 * });
 * ```
 */
export class HardcodedRoleIdentityMapper extends pulumi.CustomResource {
    /**
     * Get an existing HardcodedRoleIdentityMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HardcodedRoleIdentityMapperState, opts?: pulumi.CustomResourceOptions): HardcodedRoleIdentityMapper {
        return new HardcodedRoleIdentityMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper';

    /**
     * Returns true if the given object is an instance of HardcodedRoleIdentityMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HardcodedRoleIdentityMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HardcodedRoleIdentityMapper.__pulumiType;
    }

    public readonly extraConfig!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The IDP alias of the attribute to set.
     */
    public readonly identityProviderAlias!: pulumi.Output<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm ID that this mapper will exist in.
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * The name of the role which should be assigned to the users.
     */
    public readonly role!: pulumi.Output<string | undefined>;

    /**
     * Create a HardcodedRoleIdentityMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HardcodedRoleIdentityMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HardcodedRoleIdentityMapperArgs | HardcodedRoleIdentityMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HardcodedRoleIdentityMapperState | undefined;
            resourceInputs["extraConfig"] = state ? state.extraConfig : undefined;
            resourceInputs["identityProviderAlias"] = state ? state.identityProviderAlias : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
        } else {
            const args = argsOrState as HardcodedRoleIdentityMapperArgs | undefined;
            if ((!args || args.identityProviderAlias === undefined) && !opts.urn) {
                throw new Error("Missing required property 'identityProviderAlias'");
            }
            if ((!args || args.realm === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realm'");
            }
            resourceInputs["extraConfig"] = args ? args.extraConfig : undefined;
            resourceInputs["identityProviderAlias"] = args ? args.identityProviderAlias : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realm"] = args ? args.realm : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HardcodedRoleIdentityMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HardcodedRoleIdentityMapper resources.
 */
export interface HardcodedRoleIdentityMapperState {
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The IDP alias of the attribute to set.
     */
    identityProviderAlias?: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm ID that this mapper will exist in.
     */
    realm?: pulumi.Input<string>;
    /**
     * The name of the role which should be assigned to the users.
     */
    role?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HardcodedRoleIdentityMapper resource.
 */
export interface HardcodedRoleIdentityMapperArgs {
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The IDP alias of the attribute to set.
     */
    identityProviderAlias: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm ID that this mapper will exist in.
     */
    realm: pulumi.Input<string>;
    /**
     * The name of the role which should be assigned to the users.
     */
    role?: pulumi.Input<string>;
}
