// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Allows for creating and managing protocol mappers for both types of clients (openid-connect and saml) within Keycloak.
 *
 * There are two uses cases for using this resource:
 * * If you implemented a custom protocol mapper, this resource can be used to configure it
 * * If the provider doesn't support a particular protocol mapper, this resource can be used instead.
 *
 * Due to the generic nature of this mapper, it is less user-friendly and more prone to configuration errors.
 * Therefore, if possible, a specific mapper should be used instead.
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
 *     realmId: realm.id,
 *     clientId: "test-client",
 * });
 * const samlHardcodeAttributeMapper = new keycloak.GenericProtocolMapper("samlHardcodeAttributeMapper", {
 *     realmId: realm.id,
 *     clientId: samlClient.id,
 *     protocol: "saml",
 *     protocolMapper: "saml-hardcode-attribute-mapper",
 *     config: {
 *         "attribute.name": "name",
 *         "attribute.nameformat": "Basic",
 *         "attribute.value": "value",
 *         "friendly.name": "display name",
 *     },
 * });
 * ```
 *
 * ## Import
 *
 * Protocol mappers can be imported using the following format`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:index/genericProtocolMapper:GenericProtocolMapper saml_hardcode_attribute_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 */
export class GenericProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing GenericProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GenericProtocolMapperState, opts?: pulumi.CustomResourceOptions): GenericProtocolMapper {
        return new GenericProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/genericProtocolMapper:GenericProtocolMapper';

    /**
     * Returns true if the given object is an instance of GenericProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GenericProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GenericProtocolMapper.__pulumiType;
    }

    /**
     * The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
     */
    public readonly config!: pulumi.Output<{[key: string]: any}>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
     */
    public readonly protocolMapper!: pulumi.Output<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a GenericProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GenericProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GenericProtocolMapperArgs | GenericProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GenericProtocolMapperState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            resourceInputs["config"] = state ? state.config : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["protocolMapper"] = state ? state.protocolMapper : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as GenericProtocolMapperArgs | undefined;
            if ((!args || args.config === undefined) && !opts.urn) {
                throw new Error("Missing required property 'config'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            if ((!args || args.protocolMapper === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocolMapper'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            resourceInputs["config"] = args ? args.config : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["protocolMapper"] = args ? args.protocolMapper : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GenericProtocolMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GenericProtocolMapper resources.
 */
export interface GenericProtocolMapperState {
    /**
     * The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
     */
    config?: pulumi.Input<{[key: string]: any}>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
     */
    protocol?: pulumi.Input<string>;
    /**
     * The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
     */
    protocolMapper?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GenericProtocolMapper resource.
 */
export interface GenericProtocolMapperArgs {
    /**
     * The ID of the client this protocol mapper should be added to. Conflicts with `clientScopeId`. This argument is required if `clientScopeId` is not set.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The ID of the client scope this protocol mapper should be added to. Conflicts with `clientId`. This argument is required if `clientId` is not set.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * A map with key / value pairs for configuring the protocol mapper. The supported keys depends on the protocol mapper.
     */
    config: pulumi.Input<{[key: string]: any}>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    name?: pulumi.Input<string>;
    /**
     * The type of client (either `openid-connect` or `saml`). The type must match the type of the client.
     */
    protocol: pulumi.Input<string>;
    /**
     * The name of the protocol mapper. The protocol mapper must be compatible with the specified client.
     */
    protocolMapper: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    realmId: pulumi.Input<string>;
}