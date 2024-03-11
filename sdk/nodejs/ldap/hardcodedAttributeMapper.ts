// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing hardcoded attribute mappers for Keycloak users federated via LDAP.
 *
 * The LDAP hardcoded attribute mapper will set the specified value to the LDAP attribute.
 *
 * **NOTE**: This mapper only works when the `syncRegistrations` attribute on the `keycloak.ldap.UserFederation` resource is set to `true`.
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
 * const ldapUserFederation = new keycloak.ldap.UserFederation("ldapUserFederation", {
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
 *     syncRegistrations: true,
 * });
 * const assignBarToFoo = new keycloak.ldap.HardcodedAttributeMapper("assignBarToFoo", {
 *     realmId: realm.id,
 *     ldapUserFederationId: ldapUserFederation.id,
 *     attributeName: "foo",
 *     attributeValue: "bar",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
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
 * $ pulumi import keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper assign_bar_to_foo my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 */
export class HardcodedAttributeMapper extends pulumi.CustomResource {
    /**
     * Get an existing HardcodedAttributeMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: HardcodedAttributeMapperState, opts?: pulumi.CustomResourceOptions): HardcodedAttributeMapper {
        return new HardcodedAttributeMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper';

    /**
     * Returns true if the given object is an instance of HardcodedAttributeMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is HardcodedAttributeMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === HardcodedAttributeMapper.__pulumiType;
    }

    /**
     * The name of the LDAP attribute to set.
     */
    public readonly attributeName!: pulumi.Output<string>;
    /**
     * The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
     */
    public readonly attributeValue!: pulumi.Output<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    public readonly ldapUserFederationId!: pulumi.Output<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a HardcodedAttributeMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: HardcodedAttributeMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: HardcodedAttributeMapperArgs | HardcodedAttributeMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as HardcodedAttributeMapperState | undefined;
            resourceInputs["attributeName"] = state ? state.attributeName : undefined;
            resourceInputs["attributeValue"] = state ? state.attributeValue : undefined;
            resourceInputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as HardcodedAttributeMapperArgs | undefined;
            if ((!args || args.attributeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attributeName'");
            }
            if ((!args || args.attributeValue === undefined) && !opts.urn) {
                throw new Error("Missing required property 'attributeValue'");
            }
            if ((!args || args.ldapUserFederationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["attributeName"] = args ? args.attributeName : undefined;
            resourceInputs["attributeValue"] = args ? args.attributeValue : undefined;
            resourceInputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(HardcodedAttributeMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering HardcodedAttributeMapper resources.
 */
export interface HardcodedAttributeMapperState {
    /**
     * The name of the LDAP attribute to set.
     */
    attributeName?: pulumi.Input<string>;
    /**
     * The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
     */
    attributeValue?: pulumi.Input<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    ldapUserFederationId?: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a HardcodedAttributeMapper resource.
 */
export interface HardcodedAttributeMapperArgs {
    /**
     * The name of the LDAP attribute to set.
     */
    attributeName: pulumi.Input<string>;
    /**
     * The value to set to the LDAP attribute. You can hardcode any value like 'foo'.
     */
    attributeValue: pulumi.Input<string>;
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     */
    ldapUserFederationId: pulumi.Input<string>;
    /**
     * Display name of this mapper when displayed in the console.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm that this LDAP mapper will exist in.
     */
    realmId: pulumi.Input<string>;
}
