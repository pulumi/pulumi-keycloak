// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # keycloak.saml.UserAttributeProtocolMapper
 *
 * Allows for creating and managing user attribute protocol mappers for
 * SAML clients within Keycloak.
 *
 * SAML user attribute protocol mappers allow you to map custom attributes defined
 * for a user within Keycloak to an attribute in a SAML assertion. Protocol mappers
 * can be defined for a single client, or they can be defined for a client scope which
 * can be shared between multiple different clients.
 *
 * ### Example Usage (Client)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const samlClient = new keycloak.saml.Client("saml_client", {
 *     clientId: "test-saml-client",
 *     realmId: keycloak_realm_test.id,
 * });
 * const samlUserAttributeMapper = new keycloak.saml.UserAttributeProtocolMapper("saml_user_attribute_mapper", {
 *     clientId: samlClient.id,
 *     realmId: keycloak_realm_test.id,
 *     samlAttributeName: "displayName",
 *     samlAttributeNameFormat: "Unspecified",
 *     userAttribute: "displayName",
 * });
 * ```
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm this protocol mapper exists within.
 * - `clientId` - (Required if `clientScopeId` is not specified) The SAML client this protocol mapper is attached to.
 * - `clientScopeId` - (Required if `clientId` is not specified) The SAML client scope this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the GUI.
 * - `userAttribute` - (Required) The custom user attribute to map.
 * - `friendlyName` - (Optional) An optional human-friendly name for this attribute.
 * - `samlAttributeName` - (Required) The name of the SAML attribute.
 * - `samlAttributeNameFormat` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
 */
export class UserAttributeProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserAttributeProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserAttributeProtocolMapperState, opts?: pulumi.CustomResourceOptions): UserAttributeProtocolMapper {
        return new UserAttributeProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:saml/userAttributeProtocolMapper:UserAttributeProtocolMapper';

    /**
     * Returns true if the given object is an instance of UserAttributeProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserAttributeProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserAttributeProtocolMapper.__pulumiType;
    }

    public readonly clientId!: pulumi.Output<string | undefined>;
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly samlAttributeName!: pulumi.Output<string>;
    public readonly samlAttributeNameFormat!: pulumi.Output<string>;
    public readonly userAttribute!: pulumi.Output<string>;

    /**
     * Create a UserAttributeProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserAttributeProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserAttributeProtocolMapperArgs | UserAttributeProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserAttributeProtocolMapperState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["friendlyName"] = state ? state.friendlyName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["samlAttributeName"] = state ? state.samlAttributeName : undefined;
            inputs["samlAttributeNameFormat"] = state ? state.samlAttributeNameFormat : undefined;
            inputs["userAttribute"] = state ? state.userAttribute : undefined;
        } else {
            const args = argsOrState as UserAttributeProtocolMapperArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.samlAttributeName === undefined) {
                throw new Error("Missing required property 'samlAttributeName'");
            }
            if (!args || args.samlAttributeNameFormat === undefined) {
                throw new Error("Missing required property 'samlAttributeNameFormat'");
            }
            if (!args || args.userAttribute === undefined) {
                throw new Error("Missing required property 'userAttribute'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["samlAttributeName"] = args ? args.samlAttributeName : undefined;
            inputs["samlAttributeNameFormat"] = args ? args.samlAttributeNameFormat : undefined;
            inputs["userAttribute"] = args ? args.userAttribute : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserAttributeProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserAttributeProtocolMapper resources.
 */
export interface UserAttributeProtocolMapperState {
    readonly clientId?: pulumi.Input<string>;
    readonly clientScopeId?: pulumi.Input<string>;
    readonly friendlyName?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly realmId?: pulumi.Input<string>;
    readonly samlAttributeName?: pulumi.Input<string>;
    readonly samlAttributeNameFormat?: pulumi.Input<string>;
    readonly userAttribute?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserAttributeProtocolMapper resource.
 */
export interface UserAttributeProtocolMapperArgs {
    readonly clientId?: pulumi.Input<string>;
    readonly clientScopeId?: pulumi.Input<string>;
    readonly friendlyName?: pulumi.Input<string>;
    readonly name?: pulumi.Input<string>;
    readonly realmId: pulumi.Input<string>;
    readonly samlAttributeName: pulumi.Input<string>;
    readonly samlAttributeNameFormat: pulumi.Input<string>;
    readonly userAttribute: pulumi.Input<string>;
}
