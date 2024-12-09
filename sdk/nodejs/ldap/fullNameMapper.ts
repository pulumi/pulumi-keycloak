// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing full name mappers for Keycloak users federated via LDAP.
 *
 * The LDAP full name mapper can map a user's full name from an LDAP attribute to the first and last name attributes of a
 * Keycloak user.
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
 * const ldapUserFederation = new keycloak.ldap.UserFederation("ldap_user_federation", {
 *     name: "openldap",
 *     realmId: realm.id,
 *     usernameLdapAttribute: "cn",
 *     rdnLdapAttribute: "cn",
 *     uuidLdapAttribute: "entryDN",
 *     userObjectClasses: [
 *         "simpleSecurityObject",
 *         "organizationalRole",
 *     ],
 *     connectionUrl: "ldap://openldap",
 *     usersDn: "dc=example,dc=org",
 *     bindDn: "cn=admin,dc=example,dc=org",
 *     bindCredential: "admin",
 * });
 * const ldapFullNameMapper = new keycloak.ldap.FullNameMapper("ldap_full_name_mapper", {
 *     realmId: realm.id,
 *     ldapUserFederationId: ldapUserFederation.id,
 *     name: "full-name-mapper",
 *     ldapFullNameAttribute: "cn",
 * });
 * ```
 *
 * ## Import
 *
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
 *
 * The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:ldap/fullNameMapper:FullNameMapper ldap_full_name_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 */
export class FullNameMapper extends pulumi.CustomResource {
    /**
     * Get an existing FullNameMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: FullNameMapperState, opts?: pulumi.CustomResourceOptions): FullNameMapper {
        return new FullNameMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:ldap/fullNameMapper:FullNameMapper';

    /**
     * Returns true if the given object is an instance of FullNameMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is FullNameMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === FullNameMapper.__pulumiType;
    }

    /**
     * The name of the LDAP attribute containing the user's full name.
     */
    public readonly ldapFullNameAttribute!: pulumi.Output<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    public readonly ldapUserFederationId!: pulumi.Output<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
     */
    public readonly readOnly!: pulumi.Output<boolean | undefined>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
     */
    public readonly writeOnly!: pulumi.Output<boolean | undefined>;

    /**
     * Create a FullNameMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FullNameMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: FullNameMapperArgs | FullNameMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as FullNameMapperState | undefined;
            resourceInputs["ldapFullNameAttribute"] = state ? state.ldapFullNameAttribute : undefined;
            resourceInputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["writeOnly"] = state ? state.writeOnly : undefined;
        } else {
            const args = argsOrState as FullNameMapperArgs | undefined;
            if ((!args || args.ldapFullNameAttribute === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapFullNameAttribute'");
            }
            if ((!args || args.ldapUserFederationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["ldapFullNameAttribute"] = args ? args.ldapFullNameAttribute : undefined;
            resourceInputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["readOnly"] = args ? args.readOnly : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["writeOnly"] = args ? args.writeOnly : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(FullNameMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering FullNameMapper resources.
 */
export interface FullNameMapperState {
    /**
     * The name of the LDAP attribute containing the user's full name.
     */
    ldapFullNameAttribute?: pulumi.Input<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    ldapUserFederationId?: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    realmId?: pulumi.Input<string>;
    /**
     * When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
     */
    writeOnly?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a FullNameMapper resource.
 */
export interface FullNameMapperArgs {
    /**
     * The name of the LDAP attribute containing the user's full name.
     */
    ldapFullNameAttribute: pulumi.Input<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    ldapUserFederationId: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * When `true`, updates to a user within Keycloak will not be written back to LDAP. Defaults to `false`.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    realmId: pulumi.Input<string>;
    /**
     * When `true`, this mapper will only be used to write updates to LDAP. Defaults to `false`.
     */
    writeOnly?: pulumi.Input<boolean>;
}
