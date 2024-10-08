// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.UsersPermissionsArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.UsersPermissionsState;
import com.pulumi.keycloak.outputs.UsersPermissionsImpersonateScope;
import com.pulumi.keycloak.outputs.UsersPermissionsManageGroupMembershipScope;
import com.pulumi.keycloak.outputs.UsersPermissionsManageScope;
import com.pulumi.keycloak.outputs.UsersPermissionsMapRolesScope;
import com.pulumi.keycloak.outputs.UsersPermissionsUserImpersonatedScope;
import com.pulumi.keycloak.outputs.UsersPermissionsViewScope;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows you to manage fine-grained permissions for all users in a realm: https://www.keycloak.org/docs/latest/server_admin/#_users-permissions
 * 
 * This is part of a preview Keycloak feature: `admin_fine_grained_authz` (see https://www.keycloak.org/docs/latest/server_admin/#_fine_grain_permissions).
 * This feature can be enabled with the Keycloak option `-Dkeycloak.profile.feature.admin_fine_grained_authz=enabled`. See the
 * example `docker-compose.yml` file for an example.
 * 
 * When enabling fine-grained permissions for users, Keycloak does several things automatically:
 * 1. Enable Authorization on built-in `realm-management` client (if not already enabled).
 * 2. Create a resource representing the users permissions.
 * 3. Create scopes `view`, `manage`, `map-roles`, `manage-group-membership`, `impersonate`, and `user-impersonated`.
 * 4. Create all scope based permission for the scopes and users resources.
 * 
 * &gt; This resource should only be created once per realm.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realm_id` - (Required) The realm in which to manage fine-grained user permissions.
 * 
 * Each of the scopes that can be managed are defined below:
 * 
 * - `view_scope` - (Optional) When specified, set the scope based view permission.
 * - `manage_scope` - (Optional) When specified, set the scope based manage permission.
 * - `map_roles_scope` - (Optional) When specified, set the scope based map_roles permission.
 * - `manage_group_membership_scope` - (Optional) When specified, set the scope based manage_group_membership permission.
 * - `impersonate_scope` - (Optional) When specified, set the scope based impersonate permission.
 * - `user_impersonated_scope` - (Optional) When specified, set the scope based user_impersonated permission.
 * 
 * The configuration block for each of these scopes supports the following arguments:
 * 
 * - `policies` - (Optional) Assigned policies to the permission. Each element within this list should be a policy ID.
 * - `description` - (Optional) Description of the permission.
 * - `decision_strategy` - (Optional) Decision strategy of the permission.
 * 
 * ### Attributes Reference
 * 
 * In addition to the arguments listed above, the following computed attributes are exported:
 * 
 * - `enabled` - When true, this indicates that fine-grained user permissions are enabled. This will always be `true`.
 * - `authorization_resource_server_id` - Resource server id representing the realm management client on which these permissions are managed.
 * 
 */
@ResourceType(type="keycloak:index/usersPermissions:UsersPermissions")
public class UsersPermissions extends com.pulumi.resources.CustomResource {
    /**
     * Resource server id representing the realm management client on which this permission is managed
     * 
     */
    @Export(name="authorizationResourceServerId", refs={String.class}, tree="[0]")
    private Output<String> authorizationResourceServerId;

    /**
     * @return Resource server id representing the realm management client on which this permission is managed
     * 
     */
    public Output<String> authorizationResourceServerId() {
        return this.authorizationResourceServerId;
    }
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> enabled;

    public Output<Boolean> enabled() {
        return this.enabled;
    }
    @Export(name="impersonateScope", refs={UsersPermissionsImpersonateScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsImpersonateScope> impersonateScope;

    public Output<Optional<UsersPermissionsImpersonateScope>> impersonateScope() {
        return Codegen.optional(this.impersonateScope);
    }
    @Export(name="manageGroupMembershipScope", refs={UsersPermissionsManageGroupMembershipScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsManageGroupMembershipScope> manageGroupMembershipScope;

    public Output<Optional<UsersPermissionsManageGroupMembershipScope>> manageGroupMembershipScope() {
        return Codegen.optional(this.manageGroupMembershipScope);
    }
    @Export(name="manageScope", refs={UsersPermissionsManageScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsManageScope> manageScope;

    public Output<Optional<UsersPermissionsManageScope>> manageScope() {
        return Codegen.optional(this.manageScope);
    }
    @Export(name="mapRolesScope", refs={UsersPermissionsMapRolesScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsMapRolesScope> mapRolesScope;

    public Output<Optional<UsersPermissionsMapRolesScope>> mapRolesScope() {
        return Codegen.optional(this.mapRolesScope);
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="userImpersonatedScope", refs={UsersPermissionsUserImpersonatedScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsUserImpersonatedScope> userImpersonatedScope;

    public Output<Optional<UsersPermissionsUserImpersonatedScope>> userImpersonatedScope() {
        return Codegen.optional(this.userImpersonatedScope);
    }
    @Export(name="viewScope", refs={UsersPermissionsViewScope.class}, tree="[0]")
    private Output</* @Nullable */ UsersPermissionsViewScope> viewScope;

    public Output<Optional<UsersPermissionsViewScope>> viewScope() {
        return Codegen.optional(this.viewScope);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UsersPermissions(java.lang.String name) {
        this(name, UsersPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UsersPermissions(java.lang.String name, UsersPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UsersPermissions(java.lang.String name, UsersPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/usersPermissions:UsersPermissions", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private UsersPermissions(java.lang.String name, Output<java.lang.String> id, @Nullable UsersPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/usersPermissions:UsersPermissions", name, state, makeResourceOptions(options, id), false);
    }

    private static UsersPermissionsArgs makeArgs(UsersPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? UsersPermissionsArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static UsersPermissions get(java.lang.String name, Output<java.lang.String> id, @Nullable UsersPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UsersPermissions(name, id, state, options);
    }
}
