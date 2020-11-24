// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing an attribute importer identity provider mapper within Keycloak.
 *
 * The attribute importer mapper can be used to map attributes from externally defined users to attributes or properties of the imported Keycloak user:
 * - For the OIDC identity provider, this will map a claim on the ID or access token to an attribute for the imported Keycloak user.
 * - For the SAML identity provider, this will map a SAML attribute found within the assertion to an attribute for the imported Keycloak user.
 * - For social identity providers, this will map a JSON field from the user profile to an attribute for the imported Keycloak user.
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
 * const oidcIdentityProvider = new keycloak.oidc.IdentityProvider("oidcIdentityProvider", {
 *     realm: realm.id,
 *     alias: "oidc",
 *     authorizationUrl: "https://example.com/auth",
 *     tokenUrl: "https://example.com/token",
 *     clientId: "example_id",
 *     clientSecret: "example_token",
 *     defaultScopes: "openid random profile",
 * });
 * const oidcAttributeImporterIdentityProviderMapper = new keycloak.AttributeImporterIdentityProviderMapper("oidcAttributeImporterIdentityProviderMapper", {
 *     realm: realm.id,
 *     claimName: "my-email-claim",
 *     identityProviderAlias: oidcIdentityProvider.alias,
 *     userAttribute: "email",
 *     extraConfig: {
 *         syncMode: "INHERIT",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
 * ```
 */
export class AttributeImporterIdentityProviderMapper extends pulumi.CustomResource {
    /**
     * Get an existing AttributeImporterIdentityProviderMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AttributeImporterIdentityProviderMapperState, opts?: pulumi.CustomResourceOptions): AttributeImporterIdentityProviderMapper {
        return new AttributeImporterIdentityProviderMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/attributeImporterIdentityProviderMapper:AttributeImporterIdentityProviderMapper';

    /**
     * Returns true if the given object is an instance of AttributeImporterIdentityProviderMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AttributeImporterIdentityProviderMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AttributeImporterIdentityProviderMapper.__pulumiType;
    }

    /**
     * For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attributeName`.
     */
    public readonly attributeFriendlyName!: pulumi.Output<string | undefined>;
    /**
     * For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attributeFriendlyName`.
     */
    public readonly attributeName!: pulumi.Output<string | undefined>;
    /**
     * For OIDC based providers, this is the name of the claim to use.
     */
    public readonly claimName!: pulumi.Output<string | undefined>;
    /**
     * Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     */
    public readonly extraConfig!: pulumi.Output<{[key: string]: any} | undefined>;
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
     * The user attribute or property name to store the mapped result.
     */
    public readonly userAttribute!: pulumi.Output<string>;

    /**
     * Create a AttributeImporterIdentityProviderMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AttributeImporterIdentityProviderMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AttributeImporterIdentityProviderMapperArgs | AttributeImporterIdentityProviderMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AttributeImporterIdentityProviderMapperState | undefined;
            inputs["attributeFriendlyName"] = state ? state.attributeFriendlyName : undefined;
            inputs["attributeName"] = state ? state.attributeName : undefined;
            inputs["claimName"] = state ? state.claimName : undefined;
            inputs["extraConfig"] = state ? state.extraConfig : undefined;
            inputs["identityProviderAlias"] = state ? state.identityProviderAlias : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realm"] = state ? state.realm : undefined;
            inputs["userAttribute"] = state ? state.userAttribute : undefined;
        } else {
            const args = argsOrState as AttributeImporterIdentityProviderMapperArgs | undefined;
            if (!args || args.identityProviderAlias === undefined) {
                throw new Error("Missing required property 'identityProviderAlias'");
            }
            if (!args || args.realm === undefined) {
                throw new Error("Missing required property 'realm'");
            }
            if (!args || args.userAttribute === undefined) {
                throw new Error("Missing required property 'userAttribute'");
            }
            inputs["attributeFriendlyName"] = args ? args.attributeFriendlyName : undefined;
            inputs["attributeName"] = args ? args.attributeName : undefined;
            inputs["claimName"] = args ? args.claimName : undefined;
            inputs["extraConfig"] = args ? args.extraConfig : undefined;
            inputs["identityProviderAlias"] = args ? args.identityProviderAlias : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realm"] = args ? args.realm : undefined;
            inputs["userAttribute"] = args ? args.userAttribute : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AttributeImporterIdentityProviderMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AttributeImporterIdentityProviderMapper resources.
 */
export interface AttributeImporterIdentityProviderMapperState {
    /**
     * For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attributeName`.
     */
    readonly attributeFriendlyName?: pulumi.Input<string>;
    /**
     * For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attributeFriendlyName`.
     */
    readonly attributeName?: pulumi.Input<string>;
    /**
     * For OIDC based providers, this is the name of the claim to use.
     */
    readonly claimName?: pulumi.Input<string>;
    /**
     * Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     */
    readonly extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * The alias of the associated identity provider.
     */
    readonly identityProviderAlias?: pulumi.Input<string>;
    /**
     * The name of the mapper.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the realm.
     */
    readonly realm?: pulumi.Input<string>;
    /**
     * The user attribute or property name to store the mapped result.
     */
    readonly userAttribute?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AttributeImporterIdentityProviderMapper resource.
 */
export interface AttributeImporterIdentityProviderMapperArgs {
    /**
     * For SAML based providers, this is the friendly name of the attribute to search for in the assertion. Conflicts with `attributeName`.
     */
    readonly attributeFriendlyName?: pulumi.Input<string>;
    /**
     * For SAML based providers, this is the name of the attribute to search for in the assertion. Conflicts with `attributeFriendlyName`.
     */
    readonly attributeName?: pulumi.Input<string>;
    /**
     * For OIDC based providers, this is the name of the claim to use.
     */
    readonly claimName?: pulumi.Input<string>;
    /**
     * Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     */
    readonly extraConfig?: pulumi.Input<{[key: string]: any}>;
    /**
     * The alias of the associated identity provider.
     */
    readonly identityProviderAlias: pulumi.Input<string>;
    /**
     * The name of the mapper.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The name of the realm.
     */
    readonly realm: pulumi.Input<string>;
    /**
     * The user attribute or property name to store the mapped result.
     */
    readonly userAttribute: pulumi.Input<string>;
}
