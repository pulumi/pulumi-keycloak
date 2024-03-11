// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Allows you to manage fine-grained permissions for all users in a realm: https://www.keycloak.org/docs/latest/server_admin/#_users-permissions
 *
 * This is part of a preview Keycloak feature: `adminFineGrainedAuthz` (see https://www.keycloak.org/docs/latest/server_admin/#_fine_grain_permissions).
 * This feature can be enabled with the Keycloak option `-Dkeycloak.profile.feature.admin_fine_grained_authz=enabled`. See the
 * example `docker-compose.yml` file for an example.
 *
 * When enabling fine-grained permissions for users, Keycloak does several things automatically:
 * 1. Enable Authorization on built-in `realm-management` client (if not already enabled).
 * 2. Create a resource representing the users permissions.
 * 3. Create scopes `view`, `manage`, `map-roles`, `manage-group-membership`, `impersonate`, and `user-impersonated`.
 * 4. Create all scope based permission for the scopes and users resources.
 *
 * > This resource should only be created once per realm.
 */
export class UsersPermissions extends pulumi.CustomResource {
    /**
     * Get an existing UsersPermissions resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UsersPermissionsState, opts?: pulumi.CustomResourceOptions): UsersPermissions {
        return new UsersPermissions(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'keycloak:index/usersPermissions:UsersPermissions';

    /**
     * Returns true if the given object is an instance of UsersPermissions.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is UsersPermissions {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === UsersPermissions.__pulumiType;
    }

    /**
     * Resource server id representing the realm management client on which this permission is managed
     */
    public /*out*/ readonly authorizationResourceServerId!: pulumi.Output<string>;
    public /*out*/ readonly enabled!: pulumi.Output<boolean>;
    public readonly impersonateScope!: pulumi.Output<outputs.UsersPermissionsImpersonateScope | undefined>;
    public readonly manageGroupMembershipScope!: pulumi.Output<outputs.UsersPermissionsManageGroupMembershipScope | undefined>;
    public readonly manageScope!: pulumi.Output<outputs.UsersPermissionsManageScope | undefined>;
    public readonly mapRolesScope!: pulumi.Output<outputs.UsersPermissionsMapRolesScope | undefined>;
    public readonly realmId!: pulumi.Output<string>;
    public readonly userImpersonatedScope!: pulumi.Output<outputs.UsersPermissionsUserImpersonatedScope | undefined>;
    public readonly viewScope!: pulumi.Output<outputs.UsersPermissionsViewScope | undefined>;

    /**
     * Create a UsersPermissions resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UsersPermissionsArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UsersPermissionsArgs | UsersPermissionsState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UsersPermissionsState | undefined;
            resourceInputs["authorizationResourceServerId"] = state ? state.authorizationResourceServerId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["impersonateScope"] = state ? state.impersonateScope : undefined;
            resourceInputs["manageGroupMembershipScope"] = state ? state.manageGroupMembershipScope : undefined;
            resourceInputs["manageScope"] = state ? state.manageScope : undefined;
            resourceInputs["mapRolesScope"] = state ? state.mapRolesScope : undefined;
            resourceInputs["realmId"] = state ? state.realmId : undefined;
            resourceInputs["userImpersonatedScope"] = state ? state.userImpersonatedScope : undefined;
            resourceInputs["viewScope"] = state ? state.viewScope : undefined;
        } else {
            const args = argsOrState as UsersPermissionsArgs | undefined;
            if ((!args || args.realmId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'realmId'");
            }
            resourceInputs["impersonateScope"] = args ? args.impersonateScope : undefined;
            resourceInputs["manageGroupMembershipScope"] = args ? args.manageGroupMembershipScope : undefined;
            resourceInputs["manageScope"] = args ? args.manageScope : undefined;
            resourceInputs["mapRolesScope"] = args ? args.mapRolesScope : undefined;
            resourceInputs["realmId"] = args ? args.realmId : undefined;
            resourceInputs["userImpersonatedScope"] = args ? args.userImpersonatedScope : undefined;
            resourceInputs["viewScope"] = args ? args.viewScope : undefined;
            resourceInputs["authorizationResourceServerId"] = undefined /*out*/;
            resourceInputs["enabled"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(UsersPermissions.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering UsersPermissions resources.
 */
export interface UsersPermissionsState {
    /**
     * Resource server id representing the realm management client on which this permission is managed
     */
    authorizationResourceServerId?: pulumi.Input<string>;
    enabled?: pulumi.Input<boolean>;
    impersonateScope?: pulumi.Input<inputs.UsersPermissionsImpersonateScope>;
    manageGroupMembershipScope?: pulumi.Input<inputs.UsersPermissionsManageGroupMembershipScope>;
    manageScope?: pulumi.Input<inputs.UsersPermissionsManageScope>;
    mapRolesScope?: pulumi.Input<inputs.UsersPermissionsMapRolesScope>;
    realmId?: pulumi.Input<string>;
    userImpersonatedScope?: pulumi.Input<inputs.UsersPermissionsUserImpersonatedScope>;
    viewScope?: pulumi.Input<inputs.UsersPermissionsViewScope>;
}

/**
 * The set of arguments for constructing a UsersPermissions resource.
 */
export interface UsersPermissionsArgs {
    impersonateScope?: pulumi.Input<inputs.UsersPermissionsImpersonateScope>;
    manageGroupMembershipScope?: pulumi.Input<inputs.UsersPermissionsManageGroupMembershipScope>;
    manageScope?: pulumi.Input<inputs.UsersPermissionsManageScope>;
    mapRolesScope?: pulumi.Input<inputs.UsersPermissionsMapRolesScope>;
    realmId: pulumi.Input<string>;
    userImpersonatedScope?: pulumi.Input<inputs.UsersPermissionsUserImpersonatedScope>;
    viewScope?: pulumi.Input<inputs.UsersPermissionsViewScope>;
}
