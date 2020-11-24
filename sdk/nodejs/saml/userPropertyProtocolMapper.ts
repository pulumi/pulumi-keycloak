// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing user property protocol mappers for SAML clients within Keycloak.
 *
 * SAML user property protocol mappers allow you to map properties of the Keycloak
 * user model to an attribute in a SAML assertion.
 *
 * Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
 * multiple different clients.
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
 * const samlClient = new keycloak.saml.Client("samlClient", {
 *     realmId: keycloak_realm.test.id,
 *     clientId: "saml-client",
 * });
 * const samlUserPropertyMapper = new keycloak.saml.UserPropertyProtocolMapper("samlUserPropertyMapper", {
 *     realmId: keycloak_realm.test.id,
 *     clientId: samlClient.id,
 *     userProperty: "email",
 *     samlAttributeName: "email",
 *     samlAttributeNameFormat: "Unspecified",
 * });
 * ```
 *
 * ## Import
 *
 * Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper saml_user_property_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 *
 * ```sh
 *  $ pulumi import keycloak:saml/userPropertyProtocolMapper:UserPropertyProtocolMapper saml_user_property_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
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

    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * An optional human-friendly name for this attribute.
     */
    public readonly friendlyName!: pulumi.Output<string | undefined>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * The name of the SAML attribute.
     */
    public readonly samlAttributeName!: pulumi.Output<string>;
    /**
     * The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
     */
    public readonly samlAttributeNameFormat!: pulumi.Output<string>;
    /**
     * The property of the Keycloak user model to map.
     */
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
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as UserPropertyProtocolMapperState | undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            inputs["friendlyName"] = state ? state.friendlyName : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["realmId"] = state ? state.realmId : undefined;
            inputs["samlAttributeName"] = state ? state.samlAttributeName : undefined;
            inputs["samlAttributeNameFormat"] = state ? state.samlAttributeNameFormat : undefined;
            inputs["userProperty"] = state ? state.userProperty : undefined;
        } else {
            const args = argsOrState as UserPropertyProtocolMapperArgs | undefined;
            if (!args || args.realmId === undefined) {
                throw new Error("Missing required property 'realmId'");
            }
            if (!args || args.samlAttributeName === undefined) {
                throw new Error("Missing required property 'samlAttributeName'");
            }
            if (!args || args.samlAttributeNameFormat === undefined) {
                throw new Error("Missing required property 'samlAttributeNameFormat'");
            }
            if (!args || args.userProperty === undefined) {
                throw new Error("Missing required property 'userProperty'");
            }
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            inputs["friendlyName"] = args ? args.friendlyName : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["realmId"] = args ? args.realmId : undefined;
            inputs["samlAttributeName"] = args ? args.samlAttributeName : undefined;
            inputs["samlAttributeNameFormat"] = args ? args.samlAttributeNameFormat : undefined;
            inputs["userProperty"] = args ? args.userProperty : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(UserPropertyProtocolMapper.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserPropertyProtocolMapper resources.
 */
export interface UserPropertyProtocolMapperState {
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * An optional human-friendly name for this attribute.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    readonly realmId?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute.
     */
    readonly samlAttributeName?: pulumi.Input<string>;
    /**
     * The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
     */
    readonly samlAttributeNameFormat?: pulumi.Input<string>;
    /**
     * The property of the Keycloak user model to map.
     */
    readonly userProperty?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserPropertyProtocolMapper resource.
 */
export interface UserPropertyProtocolMapperArgs {
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    readonly clientScopeId?: pulumi.Input<string>;
    /**
     * An optional human-friendly name for this attribute.
     */
    readonly friendlyName?: pulumi.Input<string>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    readonly realmId: pulumi.Input<string>;
    /**
     * The name of the SAML attribute.
     */
    readonly samlAttributeName: pulumi.Input<string>;
    /**
     * The SAML attribute Name Format. Can be one of `Unspecified`, `Basic`, or `URI Reference`.
     */
    readonly samlAttributeNameFormat: pulumi.Input<string>;
    /**
     * The property of the Keycloak user model to map.
     */
    readonly userProperty: pulumi.Input<string>;
}
