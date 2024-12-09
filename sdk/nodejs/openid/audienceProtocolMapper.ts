// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing audience protocol mappers within Keycloak.
 *
 * Audience protocol mappers allow you add audiences to the `aud` claim within issued tokens. The audience can be a custom
 * string, or it can be mapped to the ID of a pre-existing client.
 *
 * ## Example Usage
 *
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
 * const openidClient = new keycloak.openid.Client("openid_client", {
 *     realmId: realm.id,
 *     clientId: "client",
 *     name: "client",
 *     enabled: true,
 *     accessType: "CONFIDENTIAL",
 *     validRedirectUris: ["http://localhost:8080/openid-callback"],
 * });
 * const audienceMapper = new keycloak.openid.AudienceProtocolMapper("audience_mapper", {
 *     realmId: realm.id,
 *     clientId: openidClient.id,
 *     name: "audience-mapper",
 *     includedCustomAudience: "foo",
 * });
 * ```
 *
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
 * const clientScope = new keycloak.openid.ClientScope("client_scope", {
 *     realmId: realm.id,
 *     name: "test-client-scope",
 * });
 * const audienceMapper = new keycloak.openid.AudienceProtocolMapper("audience_mapper", {
 *     realmId: realm.id,
 *     clientScopeId: clientScope.id,
 *     name: "audience-mapper",
 *     includedCustomAudience: "foo",
 * });
 * ```
 *
 * ## Import
 *
 * Protocol mappers can be imported using one of the following formats:
 *
 * - Client: `{{realm_id}}/client/{{client_keycloak_id}}/{{protocol_mapper_id}}`
 *
 * - Client Scope: `{{realm_id}}/client-scope/{{client_scope_keycloak_id}}/{{protocol_mapper_id}}`
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper audience_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 *
 * ```sh
 * $ pulumi import keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper audience_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 */
export class AudienceProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing AudienceProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AudienceProtocolMapperState, opts?: pulumi.CustomResourceOptions): AudienceProtocolMapper {
        return new AudienceProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/audienceProtocolMapper:AudienceProtocolMapper';

    /**
     * Returns true if the given object is an instance of AudienceProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AudienceProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AudienceProtocolMapper.__pulumiType;
    }

    /**
     * Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     */
    public readonly addToAccessToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     */
    public readonly addToIdToken!: pulumi.Output<boolean | undefined>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * A client ID to include within the token's `aud` claim. Conflicts with `includedCustomAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
     */
    public readonly includedClientAudience!: pulumi.Output<string | undefined>;
    /**
     * A custom audience to include within the token's `aud` claim. Conflicts with `includedClientAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
     */
    public readonly includedCustomAudience!: pulumi.Output<string | undefined>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    public readonly realmId!: pulumi.Output<string>;

    /**
     * Create a AudienceProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AudienceProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AudienceProtocolMapperArgs | AudienceProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AudienceProtocolMapperState | undefined;
            resourceInputs["addToAccessToken"] = state ? state.addToAccessToken : undefined;
            resourceInputs["addToIdToken"] = state ? state.addToIdToken : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            resourceInputs["includedClientAudience"] = state ? state.includedClientAudience : undefined;
            resourceInputs["includedCustomAudience"] = state ? state.includedCustomAudience : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
        } else {
            const args = argsOrState as AudienceProtocolMapperArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["addToAccessToken"] = args ? args.addToAccessToken : undefined;
            resourceInputs["addToIdToken"] = args ? args.addToIdToken : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            resourceInputs["includedClientAudience"] = args ? args.includedClientAudience : undefined;
            resourceInputs["includedCustomAudience"] = args ? args.includedCustomAudience : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AudienceProtocolMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AudienceProtocolMapper resources.
 */
export interface AudienceProtocolMapperState {
    /**
     * Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     */
    addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     */
    addToIdToken?: pulumi.Input<boolean>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * A client ID to include within the token's `aud` claim. Conflicts with `includedCustomAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
     */
    includedClientAudience?: pulumi.Input<string>;
    /**
     * A custom audience to include within the token's `aud` claim. Conflicts with `includedClientAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
     */
    includedCustomAudience?: pulumi.Input<string>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    realmId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AudienceProtocolMapper resource.
 */
export interface AudienceProtocolMapperArgs {
    /**
     * Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     */
    addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     */
    addToIdToken?: pulumi.Input<boolean>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * A client ID to include within the token's `aud` claim. Conflicts with `includedCustomAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
     */
    includedClientAudience?: pulumi.Input<string>;
    /**
     * A custom audience to include within the token's `aud` claim. Conflicts with `includedClientAudience`. One of `includedClientAudience` or `includedCustomAudience` must be specified.
     */
    includedCustomAudience?: pulumi.Input<string>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    realmId: pulumi.Input<string>;
}
