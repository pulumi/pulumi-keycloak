// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating the "Audience Resolve" OIDC protocol mapper within Keycloak.
 *
 * This protocol mapper is useful to avoid manual management of audiences, instead relying on the presence of client roles
 * to imply which audiences are appropriate for the token. See the
 * [Keycloak docs](https://www.keycloak.org/docs/latest/server_admin/#_audience_resolve) for more details.
 *
 * ## Example Usage
 * ### Client)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const openidClient = new keycloak.openid.Client("openidClient", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     enabled: true,
 *     accessType: "CONFIDENTIAL",
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const audienceMapper = new keycloak.openid.AudienceResolveProtocolMappter("audienceMapper", {
 *     realmId: realm.id,
 *     clientId: openidClient.id,
 * });
 * ```
 * ### Client Scope)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as keycloak from "@pulumi/keycloak";
 *
 * const realm = new keycloak.Realm("realm", {
 *     realm: "my-realm",
 *     enabled: true,
 * });
 * const clientScope = new keycloak.openid.ClientScope("clientScope", {realmId: realm.id});
 * const audienceMapper = new keycloak.openid.AudienceProtocolMapper("audienceMapper", {
 *     realmId: realm.id,
 *     clientScopeId: clientScope.id,
 * });
 * ```
 *
 * ## Import
 *
 * Protocol mappers can be imported using one of the following formats- Client`{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}` - Client Scope`{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}` Examplebash
 *
 * ```sh
 *  $ pulumi import keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter audience_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 *
 * ```sh
 *  $ pulumi import keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter audience_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 */
export class AudienceResolveProtocolMappter extends pulumi.CustomResource {
    /**
     * Get an existing AudienceResolveProtocolMappter resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AudienceResolveProtocolMappterState, opts?: pulumi.CustomResourceOptions): AudienceResolveProtocolMappter {
        return new AudienceResolveProtocolMappter(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/audienceResolveProtocolMappter:AudienceResolveProtocolMappter';

    /**
     * Returns true if the given object is an instance of AudienceResolveProtocolMappter.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AudienceResolveProtocolMappter {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AudienceResolveProtocolMappter.__pulumiType;
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
     * The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a AudienceResolveProtocolMappter resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AudienceResolveProtocolMappterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AudienceResolveProtocolMappterArgs | AudienceResolveProtocolMappterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AudienceResolveProtocolMappterState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as AudienceResolveProtocolMappterArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AudienceResolveProtocolMappter.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AudienceResolveProtocolMappter resources.
 */
export interface AudienceResolveProtocolMappterState {
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
     */
    name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AudienceResolveProtocolMappter resource.
 */
export interface AudienceResolveProtocolMappterArgs {
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * The display name of this protocol mapper in the GUI. Defaults to "audience resolve".
     */
    name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    realmId: pulumi.Input<string>;
}