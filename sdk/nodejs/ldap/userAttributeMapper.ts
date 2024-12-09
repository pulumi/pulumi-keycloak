// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing user attribute mappers for Keycloak users
 * federated via LDAP.
 *
 * The LDAP user attribute mapper can be used to map a single LDAP attribute
 * to an attribute on the Keycloak user model.
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
 * const ldapUserAttributeMapper = new keycloak.ldap.UserAttributeMapper("ldap_user_attribute_mapper", {
 *     realmId: realm.id,
 *     ldapUserFederationId: ldapUserFederation.id,
 *     name: "user-attribute-mapper",
 *     userModelAttribute: "foo",
 *     ldapAttribute: "bar",
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
 * $ pulumi import keycloak:ldap/userAttributeMapper:UserAttributeMapper ldap_user_attribute_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 */
export class UserAttributeMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserAttributeMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserAttributeMapperState, opts?: pulumi.CustomResourceOptions): UserAttributeMapper {
        return new UserAttributeMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:ldap/userAttributeMapper:UserAttributeMapper';

    /**
     * Returns true if the given object is an instance of UserAttributeMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserAttributeMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserAttributeMapper.__pulumiType;
    }

    /**
     * When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
     */
    public readonly alwaysReadValueFromLdap!: pulumi.Output<boolean | undefined>;
    /**
     * Default value to set in LDAP if `isMandatoryInLdap` is true and the value is empty.
     */
    public readonly attributeDefaultValue!: pulumi.Output<string | undefined>;
    /**
     * Should be true for binary LDAP attributes.
     */
    public readonly isBinaryAttribute!: pulumi.Output<boolean | undefined>;
    /**
     * When `true`, this attribute must exist in LDAP. Defaults to `false`.
     */
    public readonly isMandatoryInLdap!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the mapped attribute on the LDAP object.
     */
    public readonly ldapAttribute!: pulumi.Output<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    public readonly ldapUserFederationId!: pulumi.Output<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
     */
    public readonly readOnly!: pulumi.Output<boolean | undefined>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * Name of the user property or attribute you want to map the LDAP attribute into.
     */
    public readonly userModelAttribute!: pulumi.Output<string>;

    /**
     * Create a UserAttributeMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserAttributeMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserAttributeMapperArgs | UserAttributeMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserAttributeMapperState | undefined;
            resourceInputs["alwaysReadValueFromLdap"] = state ? state.alwaysReadValueFromLdap : undefined;
            resourceInputs["attributeDefaultValue"] = state ? state.attributeDefaultValue : undefined;
            resourceInputs["isBinaryAttribute"] = state ? state.isBinaryAttribute : undefined;
            resourceInputs["isMandatoryInLdap"] = state ? state.isMandatoryInLdap : undefined;
            resourceInputs["ldapAttribute"] = state ? state.ldapAttribute : undefined;
            resourceInputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["readOnly"] = state ? state.readOnly : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["userModelAttribute"] = state ? state.userModelAttribute : undefined;
        } else {
            const args = argsOrState as UserAttributeMapperArgs | undefined;
            if ((!args || args.ldapAttribute === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapAttribute'");
            }
            if ((!args || args.ldapUserFederationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.userModelAttribute === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userModelAttribute'");
            }
            resourceInputs["alwaysReadValueFromLdap"] = args ? args.alwaysReadValueFromLdap : undefined;
            resourceInputs["attributeDefaultValue"] = args ? args.attributeDefaultValue : undefined;
            resourceInputs["isBinaryAttribute"] = args ? args.isBinaryAttribute : undefined;
            resourceInputs["isMandatoryInLdap"] = args ? args.isMandatoryInLdap : undefined;
            resourceInputs["ldapAttribute"] = args ? args.ldapAttribute : undefined;
            resourceInputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["readOnly"] = args ? args.readOnly : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["userModelAttribute"] = args ? args.userModelAttribute : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserAttributeMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserAttributeMapper resources.
 */
export interface UserAttributeMapperState {
    /**
     * When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
     */
    alwaysReadValueFromLdap?: pulumi.Input<boolean>;
    /**
     * Default value to set in LDAP if `isMandatoryInLdap` is true and the value is empty.
     */
    attributeDefaultValue?: pulumi.Input<string>;
    /**
     * Should be true for binary LDAP attributes.
     */
    isBinaryAttribute?: pulumi.Input<boolean>;
    /**
     * When `true`, this attribute must exist in LDAP. Defaults to `false`.
     */
    isMandatoryInLdap?: pulumi.Input<boolean>;
    /**
     * Name of the mapped attribute on the LDAP object.
     */
    ldapAttribute?: pulumi.Input<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    ldapUserFederationId?: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    realmId?: pulumi.Input<string>;
    /**
     * Name of the user property or attribute you want to map the LDAP attribute into.
     */
    userModelAttribute?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserAttributeMapper resource.
 */
export interface UserAttributeMapperArgs {
    /**
     * When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
     */
    alwaysReadValueFromLdap?: pulumi.Input<boolean>;
    /**
     * Default value to set in LDAP if `isMandatoryInLdap` is true and the value is empty.
     */
    attributeDefaultValue?: pulumi.Input<string>;
    /**
     * Should be true for binary LDAP attributes.
     */
    isBinaryAttribute?: pulumi.Input<boolean>;
    /**
     * When `true`, this attribute must exist in LDAP. Defaults to `false`.
     */
    isMandatoryInLdap?: pulumi.Input<boolean>;
    /**
     * Name of the mapped attribute on the LDAP object.
     */
    ldapAttribute: pulumi.Input<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    ldapUserFederationId: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
     */
    readOnly?: pulumi.Input<boolean>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    realmId: pulumi.Input<string>;
    /**
     * Name of the user property or attribute you want to map the LDAP attribute into.
     */
    userModelAttribute: pulumi.Input<string>;
}
