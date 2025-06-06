// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for creating and managing user session note protocol mappers within Keycloak.
 *
 * User session note protocol mappers map a custom user session note to a token claim.
 *
 * Protocol mappers can be defined for a single client, or they can be defined for a client scope which can be shared between
 * multiple different clients.
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
 * const userSessionNoteMapper = new keycloak.openid.UserSessionNoteProtocolMapper("user_session_note_mapper", {
 *     realmId: realm.id,
 *     clientId: openidClient.id,
 *     name: "user-session-note-mapper",
 *     claimName: "foo",
 *     claimValueType: "String",
 *     sessionNote: "bar",
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
 *     name: "client-scope",
 * });
 * const userSessionNoteMapper = new keycloak.openid.UserSessionNoteProtocolMapper("user_session_note_mapper", {
 *     realmId: realm.id,
 *     clientScopeId: clientScope.id,
 *     name: "user-session-note-mapper",
 *     claimName: "foo",
 *     claimValueType: "String",
 *     sessionNote: "bar",
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
 * $ pulumi import keycloak:openid/userSessionNoteProtocolMapper:UserSessionNoteProtocolMapper user_session_note_mapper my-realm/client/a7202154-8793-4656-b655-1dd18c181e14/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 *
 * ```sh
 * $ pulumi import keycloak:openid/userSessionNoteProtocolMapper:UserSessionNoteProtocolMapper user_session_note_mapper my-realm/client-scope/b799ea7e-73ee-4a73-990a-1eafebe8e20a/71602afa-f7d1-4788-8c49-ef8fd00af0f4
 * ```
 */
export class UserSessionNoteProtocolMapper extends pulumi.CustomResource {
    /**
     * Get an existing UserSessionNoteProtocolMapper resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserSessionNoteProtocolMapperState, opts?: pulumi.CustomResourceOptions): UserSessionNoteProtocolMapper {
        return new UserSessionNoteProtocolMapper(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/userSessionNoteProtocolMapper:UserSessionNoteProtocolMapper';

    /**
     * Returns true if the given object is an instance of UserSessionNoteProtocolMapper.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UserSessionNoteProtocolMapper {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UserSessionNoteProtocolMapper.__pulumiType;
    }

    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    public readonly addToAccessToken!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    public readonly addToIdToken!: pulumi.Output<boolean | undefined>;
    /**
     * The name of the claim to insert into a token.
     */
    public readonly claimName!: pulumi.Output<string>;
    /**
     * The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     */
    public readonly claimValueType!: pulumi.Output<string | undefined>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    public readonly clientScopeId!: pulumi.Output<string | undefined>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * String value being the name of stored user session note within the `UserSessionModel.note` map.
     */
    public readonly sessionNote!: pulumi.Output<string | undefined>;

    /**
     * Create a UserSessionNoteProtocolMapper resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserSessionNoteProtocolMapperArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserSessionNoteProtocolMapperArgs | UserSessionNoteProtocolMapperState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserSessionNoteProtocolMapperState | undefined;
            resourceInputs["addToAccessToken"] = state ? state.addToAccessToken : undefined;
            resourceInputs["addToIdToken"] = state ? state.addToIdToken : undefined;
            resourceInputs["claimName"] = state ? state.claimName : undefined;
            resourceInputs["claimValueType"] = state ? state.claimValueType : undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["clientScopeId"] = state ? state.clientScopeId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["sessionNote"] = state ? state.sessionNote : undefined;
        } else {
            const args = argsOrState as UserSessionNoteProtocolMapperArgs | undefined;
            if ((!args || args.claimName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'claimName'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["addToAccessToken"] = args ? args.addToAccessToken : undefined;
            resourceInputs["addToIdToken"] = args ? args.addToIdToken : undefined;
            resourceInputs["claimName"] = args ? args.claimName : undefined;
            resourceInputs["claimValueType"] = args ? args.claimValueType : undefined;
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["clientScopeId"] = args ? args.clientScopeId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["sessionNote"] = args ? args.sessionNote : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UserSessionNoteProtocolMapper.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UserSessionNoteProtocolMapper resources.
 */
export interface UserSessionNoteProtocolMapperState {
    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    addToIdToken?: pulumi.Input<boolean>;
    /**
     * The name of the claim to insert into a token.
     */
    claimName?: pulumi.Input<string>;
    /**
     * The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     */
    claimValueType?: pulumi.Input<string>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    realmId?: pulumi.Input<string>;
    /**
     * String value being the name of stored user session note within the `UserSessionModel.note` map.
     */
    sessionNote?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a UserSessionNoteProtocolMapper resource.
 */
export interface UserSessionNoteProtocolMapperArgs {
    /**
     * Indicates if the property should be added as a claim to the access token. Defaults to `true`.
     */
    addToAccessToken?: pulumi.Input<boolean>;
    /**
     * Indicates if the property should be added as a claim to the id token. Defaults to `true`.
     */
    addToIdToken?: pulumi.Input<boolean>;
    /**
     * The name of the claim to insert into a token.
     */
    claimName: pulumi.Input<string>;
    /**
     * The claim type used when serializing JSON tokens. Can be one of `String`, `JSON`, `long`, `int`, or `boolean`. Defaults to `String`.
     */
    claimValueType?: pulumi.Input<string>;
    /**
     * The client this protocol mapper should be attached to. Conflicts with `clientScopeId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `clientId`. One of `clientId` or `clientScopeId` must be specified.
     */
    clientScopeId?: pulumi.Input<string>;
    /**
     * The display name of this protocol mapper in the GUI.
     */
    name?: pulumi.Input<string>;
    /**
     * The realm this protocol mapper exists within.
     */
    realmId: pulumi.Input<string>;
    /**
     * String value being the name of stored user session note within the `UserSessionModel.note` map.
     */
    sessionNote?: pulumi.Input<string>;
}
