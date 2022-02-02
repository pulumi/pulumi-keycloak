// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing MSAD-LDS user account control mappers for Keycloak
 * users federated via LDAP.
 *
 * The MSAD-LDS (Microsoft Active Directory Lightweight Directory Service) user account control mapper is specific
 * to LDAP user federation providers that are pulling from AD-LDS, and it can propagate
 * AD-LDS user state to Keycloak in order to enforce settings like expired passwords
 * or disabled accounts.
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
 * const ldapUserFederation = new keycloak.ldap.UserFederation("ldapUserFederation", {
 *     realmId: realm.id,
 *     usernameLdapAttribute: "cn",
 *     rdnLdapAttribute: "cn",
 *     uuidLdapAttribute: "objectGUID",
 *     userObjectClasses: [
 *         "person",
 *         "organizationalPerson",
 *         "user",
 *     ],
 *     connectionUrl: "ldap://my-ad-server",
 *     usersDn: "dc=example,dc=org",
 *     bindDn: "cn=admin,dc=example,dc=org",
 *     bindCredential: "admin",
 * });
 * const msadLdsUserAccountControlMapper = new keycloak.ldap.MsadLdsUserAccountControlMapper("msadLdsUserAccountControlMapper", {
 *     realmId: realm.id,
 *     ldapUserFederationId: ldapUserFederation.id,
 * });
 * ```
 *
 * ## Import
 *
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper msad_lds_user_account_control_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 */
export class MsadLdsUserAccountControlMapper extends pulumi.CustomResource {
    /**
     * Get an existing MsadLdsUserAccountControlMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MsadLdsUserAccountControlMapperState, opts?: pulumi.CustomResourceOptions): MsadLdsUserAccountControlMapper {
        return new MsadLdsUserAccountControlMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:ldap/msadLdsUserAccountControlMapper:MsadLdsUserAccountControlMapper';

    /**
     * Returns true if the given object is an instance of MsadLdsUserAccountControlMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MsadLdsUserAccountControlMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MsadLdsUserAccountControlMapper.__pulumiType;
    }

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
     * Create a MsadLdsUserAccountControlMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MsadLdsUserAccountControlMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MsadLdsUserAccountControlMapperArgs | MsadLdsUserAccountControlMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MsadLdsUserAccountControlMapperState | undefined;
            resourceInputs["ldapUserFederationId"] = state ? state.ldapUserFederationId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as MsadLdsUserAccountControlMapperArgs | undefined;
            if ((!args || args.ldapUserFederationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ldapUserFederationId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["ldapUserFederationId"] = args ? args.ldapUserFederationId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MsadLdsUserAccountControlMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MsadLdsUserAccountControlMapper resources.
 */
export interface MsadLdsUserAccountControlMapperState {
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
 * The set of arguments for constructing a MsadLdsUserAccountControlMapper resource.
 */
export interface MsadLdsUserAccountControlMapperArgs {
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
