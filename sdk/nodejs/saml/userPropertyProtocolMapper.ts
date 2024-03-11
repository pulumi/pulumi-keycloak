// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## # keycloak.saml.UserPropertyProtocolMapper
 *
 * Allows for creating and managing user property protocol mappers for
 * SAML clients within Keycloak.
 *
 * SAML user property protocol mappers allow you to map properties of the Keycloak
 * user model to an attribute in a SAML assertion. Protocol mappers
 * can be defined for a single client, or they can be defined for a client scope which
 * can be shared between multiple different clients.
 *
 * ### Example Usage (Client)
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     enabled: true,
 *     realm: "my-realm",
 * });
 * const samlClient = new keycloak.saml.Client("samlClient", {
 *     clientId: "test-saml-client",
 *     realmId: keycloak_realm.test.id,
 * });
 * const samlUserPropertyMapper = new keycloak.saml.UserPropertyProtocolMapper("samlUserPropertyMapper", {
 *     clientId: samlClient.id,
 *     realmId: keycloak_realm.test.id,
 *     samlAttributeName: "email",
 *     samlAttributeNameFormat: "Unspecified",
 *     userProperty: "email",
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ### Argument Reference
 *
 * The following arguments are supported:
 *
 * - `realmId` - (Required) The realm this protocol mapper exists within.
 * - `clientId` - (Required if `clientScopeId` is not specified) The SAML client this protocol mapper is attached to.
 * - `clientScopeId` - (Required if `clientId` is not specified) The SAML client scope this protocol mapper is attached to.
 * - `name` - (Required) The display name of this protocol mapper in the GUI.
 * - `userProperty` - (Required) The property of the Keycloak user model to map.
 * - `friendlyName` - (Optional) An optional human-friendly name for this attribute.
 * - `samlAttributeName` - (Required) The name of the SAML attribute.
 * - `samlAttributeNameFormat` - (Required) The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
 *
 * ### Import
 *
 * Protocol mappers can be imported using one of the following formats:
 * - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
 * - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
 *
 * Example:
 */
export class UserPropertyProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserPropertyProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserPropertyProtocolMapperState, opts?: pulumi.CustomResourceOptions): UserPropertyProtocolMapper {
        return new UserPropertyProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper';

    /**
     * Returns true if the given object is an instance of UserPropertyProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserPropertyProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserPropertyProtocolMapper.__pulumiType;
    }

    public readonly clientId!: pulumi.Output<string | undefined>;
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    public readonly name!: pulumi.Output<string>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly samlAttributeName!: pulumi.Output<string>;
    public readonly samlAttributeNameFormat!: pulumi.Output<string>;
    public readonly userProperty!: pulumi.Output<string>;

    /**
     * Create a UserPropertyProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserPropertyProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserPropertyProtocolMapperArgs | UserPropertyProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserPropertyProtocolMapperState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            resourceInputs["friendlyName"] = state ? state.friendlyName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["samlAttributeName"] = state ? state.samlAttributeName : undefined;
            resourceInputs["samlAttributeNameFormat"] = state ? state.samlAttributeNameFormat : undefined;
            resourceInputs["userProperty"] = state ? state.userProperty : undefined;
        } else {
            const args = argsOrState as UserPropertyProtocolMapperArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.samlAttributeName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'samlAttributeName'");
            }
            if ((!args || args.samlAttributeNameFormat === undefined) && !opts.urn) {
                throw new Error("Missing required property 'samlAttributeNameFormat'");
            }
            if ((!args || args.userProperty === undefined) && !opts.urn) {
                throw new Error("Missing required property 'userProperty'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            resourceInputs["friendlyName"] = args ? args.friendlyName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["samlAttributeName"] = args ? args.samlAttributeName : undefined;
            resourceInputs["samlAttributeNameFormat"] = args ? args.samlAttributeNameFormat : undefined;
            resourceInputs["userProperty"] = args ? args.userProperty : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserPropertyProtocolMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPropertyProtocolMapper resources.
 */
export interface UserPropertyProtocolMapperState {
    clientId?: pulumi.Input<string>;
    clientScopeId?: pulumi.Input<string>;
    friendlyName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    realmId?: pulumi.Input<string>;
    samlAttributeName?: pulumi.Input<string>;
    samlAttributeNameFormat?: pulumi.Input<string>;
    userProperty?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserPropertyProtocolMapper resource.
 */
export interface UserPropertyProtocolMapperArgs {
    clientId?: pulumi.Input<string>;
    clientScopeId?: pulumi.Input<string>;
    friendlyName?: pulumi.Input<string>;
    name?: pulumi.Input<string>;
    realmId: pulumi.Input<string>;
    samlAttributeName: pulumi.Input<string>;
    samlAttributeNameFormat: pulumi.Input<string>;
    userProperty: pulumi.Input<string>;
}
