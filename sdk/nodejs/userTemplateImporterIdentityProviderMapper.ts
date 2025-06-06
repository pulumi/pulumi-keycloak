// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing an username template importer identity provider mapper within Keycloak.
 *
 * The username template importer mapper can be used to map externally defined OIDC claims or SAML attributes with a template to the username of the imported Keycloak user:
 *
 * - Substitutions are enclosed in \${}. For example: '\${ALIAS}.\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
 *
 * > If you are using Keycloak 10 or higher, you will need to specify the `extraConfig` argument in order to define a `syncMode` for the mapper.
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
 *     alias: "oidc",
 *     authorizationUrl: "https://example.com/auth",
 *     tokenUrl: "https://example.com/token",
 *     clientId: "example_id",
 *     clientSecret: "example_token",
 *     defaultScopes: "openid random profile",
 * });
 * const usernameImporter = new keycloak.UserTemplateImporterIdentityProviderMapper("username_importer", {
 *     realm: realm.id,
 *     name: "username-template-importer",
 *     identityProviderAlias: oidc.alias,
 *     template: "${ALIAS}.${CLAIM.email}",
 *     extraConfig: {
 *         syncMode: "INHERIT",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak
 *
 * assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID.
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper username_importer my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
 * ```
 */
export class UserTemplateImporterIdentityProviderMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserTemplateImporterIdentityProviderMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserTemplateImporterIdentityProviderMapperState, opts?: pulumi.CustomResourceOptions): UserTemplateImporterIdentityProviderMapper {
        return new UserTemplateImporterIdentityProviderMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/userTemplateImporterIdentityProviderMapper:UserTemplateImporterIdentityProviderMapper';

    /**
     * Returns true if the given object is an instance of UserTemplateImporterIdentityProviderMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserTemplateImporterIdentityProviderMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserTemplateImporterIdentityProviderMapper.__pulumiType;
    }

    /**
     * Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     */
    public readonly extraConfig!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * The alias of the associated identity provider.
     */
    public readonly identityProviderAlias!: pulumi.Output<string>;
    /**
     * The name of the mapper.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of the realm.
     */
    public readonly realm!: pulumi.Output<string>;
    /**
     * Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
     */
    public readonly template!: pulumi.Output<string | undefined>;

    /**
     * Create a UserTemplateImporterIdentityProviderMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserTemplateImporterIdentityProviderMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserTemplateImporterIdentityProviderMapperArgs | UserTemplateImporterIdentityProviderMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserTemplateImporterIdentityProviderMapperState | undefined;
            resourceInputs["extraConfig"] = state ? state.extraConfig : undefined;
            resourceInputs["identityProviderAlias"] = state ? state.identityProviderAlias : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realm"] = state ? state.realm : undefined;
            resourceInputs["template"] = state ? state.template : undefined;
        } else {
            const args = argsOrState as UserTemplateImporterIdentityProviderMapperArgs | undefined;
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
            resourceInputs["template"] = args ? args.template : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserTemplateImporterIdentityProviderMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserTemplateImporterIdentityProviderMapper resources.
 */
export interface UserTemplateImporterIdentityProviderMapperState {
    /**
     * Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     */
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The alias of the associated identity provider.
     */
    identityProviderAlias?: pulumi.Input<string>;
    /**
     * The name of the mapper.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the realm.
     */
    realm?: pulumi.Input<string>;
    /**
     * Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
     */
    template?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserTemplateImporterIdentityProviderMapper resource.
 */
export interface UserTemplateImporterIdentityProviderMapperArgs {
    /**
     * Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     */
    extraConfig?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * The alias of the associated identity provider.
     */
    identityProviderAlias: pulumi.Input<string>;
    /**
     * The name of the mapper.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of the realm.
     */
    realm: pulumi.Input<string>;
    /**
     * Template to use to format the username to import. Substitutions are enclosed in \${}. For example: '\$\${ALIAS}.\$\${CLAIM.sub}'. ALIAS is the provider alias. CLAIM.\<NAME\> references an ID or Access token claim.
     */
    template?: pulumi.Input<string>;
}
