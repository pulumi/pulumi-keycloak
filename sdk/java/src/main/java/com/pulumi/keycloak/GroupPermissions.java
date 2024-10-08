// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.GroupPermissionsArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.GroupPermissionsState;
import com.pulumi.keycloak.outputs.GroupPermissionsManageMembersScope;
import com.pulumi.keycloak.outputs.GroupPermissionsManageMembershipScope;
import com.pulumi.keycloak.outputs.GroupPermissionsManageScope;
import com.pulumi.keycloak.outputs.GroupPermissionsViewMembersScope;
import com.pulumi.keycloak.outputs.GroupPermissionsViewScope;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows you to manage all group Scope Based Permissions https://www.keycloak.org/docs/latest/server_admin/#group.
 * 
 * This is part of a preview Keycloak feature: `admin_fine_grained_authz` (see https://www.keycloak.org/docs/latest/server_admin/#_fine_grain_permissions).
 * This feature can be enabled with the Keycloak option `-Dkeycloak.profile.feature.admin_fine_grained_authz=enabled`. See the
 * example `docker-compose.yml` file for an example.
 * 
 * When enabling Roles Permissions, Keycloak does several things automatically:
 * 1. Enable Authorization on built-in `realm-management` client (if not already enabled).
 * 2. Create a resource representing the role permissions.
 * 3. Create scopes `view`, `manage`, `view-members`, `manage-members`, `manage-membership`.
 * 4. Create all scope based permission for the scopes and role resource
 * 
 */
@ResourceType(type="keycloak:index/groupPermissions:GroupPermissions")
public class GroupPermissions extends com.pulumi.resources.CustomResource {
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
    @Export(name="groupId", refs={String.class}, tree="[0]")
    private Output<String> groupId;

    public Output<String> groupId() {
        return this.groupId;
    }
    @Export(name="manageMembersScope", refs={GroupPermissionsManageMembersScope.class}, tree="[0]")
    private Output</* @Nullable */ GroupPermissionsManageMembersScope> manageMembersScope;

    public Output<Optional<GroupPermissionsManageMembersScope>> manageMembersScope() {
        return Codegen.optional(this.manageMembersScope);
    }
    @Export(name="manageMembershipScope", refs={GroupPermissionsManageMembershipScope.class}, tree="[0]")
    private Output</* @Nullable */ GroupPermissionsManageMembershipScope> manageMembershipScope;

    public Output<Optional<GroupPermissionsManageMembershipScope>> manageMembershipScope() {
        return Codegen.optional(this.manageMembershipScope);
    }
    @Export(name="manageScope", refs={GroupPermissionsManageScope.class}, tree="[0]")
    private Output</* @Nullable */ GroupPermissionsManageScope> manageScope;

    public Output<Optional<GroupPermissionsManageScope>> manageScope() {
        return Codegen.optional(this.manageScope);
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="viewMembersScope", refs={GroupPermissionsViewMembersScope.class}, tree="[0]")
    private Output</* @Nullable */ GroupPermissionsViewMembersScope> viewMembersScope;

    public Output<Optional<GroupPermissionsViewMembersScope>> viewMembersScope() {
        return Codegen.optional(this.viewMembersScope);
    }
    @Export(name="viewScope", refs={GroupPermissionsViewScope.class}, tree="[0]")
    private Output</* @Nullable */ GroupPermissionsViewScope> viewScope;

    public Output<Optional<GroupPermissionsViewScope>> viewScope() {
        return Codegen.optional(this.viewScope);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupPermissions(java.lang.String name) {
        this(name, GroupPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupPermissions(java.lang.String name, GroupPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupPermissions(java.lang.String name, GroupPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/groupPermissions:GroupPermissions", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private GroupPermissions(java.lang.String name, Output<java.lang.String> id, @Nullable GroupPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/groupPermissions:GroupPermissions", name, state, makeResourceOptions(options, id), false);
    }

    private static GroupPermissionsArgs makeArgs(GroupPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? GroupPermissionsArgs.Empty : args;
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
    public static GroupPermissions get(java.lang.String name, Output<java.lang.String> id, @Nullable GroupPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupPermissions(name, id, state, options);
    }
}
