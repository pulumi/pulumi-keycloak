// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Allows for assigning client roles to the service account of an openid client.
 * You need to set `serviceAccountsEnabled` to `true` for the openid client that should be assigned the role.
 *
 * If you'd like to attach realm roles to a service account, please use the `keycloak.openid.ClientServiceAccountRealmRole`
 * resource.
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
 * // client1 provides a role to other clients
 * const client1 = new keycloak.openid.Client("client1", {
 *     realmId: realm.id,
 *     name: "client1",
 * });
 * const client1Role = new keycloak.Role("client1_role", {
 *     realmId: realm.id,
 *     clientId: client1.id,
 *     name: "my-client1-role",
 *     description: "A role that client1 provides",
 * });
 * // client2 is assigned the role of client1
 * const client2 = new keycloak.openid.Client("client2", {
 *     realmId: realm.id,
 *     name: "client2",
 *     serviceAccountsEnabled: true,
 * });
 * const client2ServiceAccountRole = new keycloak.openid.ClientServiceAccountRole("client2_service_account_role", {
 *     realmId: realm.id,
 *     serviceAccountUserId: client2.serviceAccountUserId,
 *     clientId: client1.id,
 *     role: client1Role.name,
 * });
 * ```
 *
 * ## Import
 *
 * This resource can be imported using the format `{{realmId}}/{{serviceAccountUserId}}/{{clientId}}/{{roleId}}`.
 *
 * Example:
 *
 * bash
 *
 * ```sh
 * $ pulumi import keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole client2_service_account_role my-realm/489ba513-1ceb-49ba-ae0b-1ab1f5099ebf/baf01820-0f8b-4494-9be2-fb3bc8a397a4/c7230ab7-8e4e-4135-995d-e81b50696ad8
 * ```
 */
export class ClientServiceAccountRole extends pulumi.CustomResource {
    /**
     * Get an existing ClientServiceAccountRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ClientServiceAccountRoleState, opts?: pulumi.CustomResourceOptions): ClientServiceAccountRole {
        return new ClientServiceAccountRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:openid/clientServiceAccountRole:ClientServiceAccountRole';

    /**
     * Returns true if the given object is an instance of ClientServiceAccountRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ClientServiceAccountRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ClientServiceAccountRole.__pulumiType;
    }

    /**
     * The id of the client that provides the role.
     */
    public readonly clientId!: pulumi.Output<string>;
    /**
     * The realm the clients and roles belong to.
     */
    public readonly realmId!: pulumi.Output<string>;
    /**
     * The name of the role that is assigned.
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
     */
    public readonly serviceAccountUserId!: pulumi.Output<string>;

    /**
     * Create a ClientServiceAccountRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ClientServiceAccountRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ClientServiceAccountRoleArgs | ClientServiceAccountRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ClientServiceAccountRoleState | undefined;
            resourceInputs["clientId"] = state ? state.clientId : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["serviceAccountUserId"] = state ? state.serviceAccountUserId : undefined;
        } else {
            const args = argsOrState as ClientServiceAccountRoleArgs | undefined;
            if ((!args || args.clientId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clientId'");
            }
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.serviceAccountUserId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountUserId'");
            }
            resourceInputs["clientId"] = args ? args.clientId : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["serviceAccountUserId"] = args ? args.serviceAccountUserId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ClientServiceAccountRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ClientServiceAccountRole resources.
 */
export interface ClientServiceAccountRoleState {
    /**
     * The id of the client that provides the role.
     */
    clientId?: pulumi.Input<string>;
    /**
     * The realm the clients and roles belong to.
     */
    realmId?: pulumi.Input<string>;
    /**
     * The name of the role that is assigned.
     */
    role?: pulumi.Input<string>;
    /**
     * The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
     */
    serviceAccountUserId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ClientServiceAccountRole resource.
 */
export interface ClientServiceAccountRoleArgs {
    /**
     * The id of the client that provides the role.
     */
    clientId: pulumi.Input<string>;
    /**
     * The realm the clients and roles belong to.
     */
    realmId: pulumi.Input<string>;
    /**
     * The name of the role that is assigned.
     */
    role: pulumi.Input<string>;
    /**
     * The id of the service account that is assigned the role (the service account of the client that "consumes" the role).
     */
    serviceAccountUserId: pulumi.Input<string>;
}
