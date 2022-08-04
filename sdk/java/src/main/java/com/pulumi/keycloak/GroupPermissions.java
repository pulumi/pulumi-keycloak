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

@ResourceType(type="keycloak:index/groupPermissions:GroupPermissions")
public class GroupPermissions extends com.pulumi.resources.CustomResource {
    /**
     * Resource server id representing the realm management client on which this permission is managed
     * 
     */
    @Export(name="authorizationResourceServerId", type=String.class, parameters={})
    private Output<String> authorizationResourceServerId;

    /**
     * @return Resource server id representing the realm management client on which this permission is managed
     * 
     */
    public Output<String> authorizationResourceServerId() {
        return this.authorizationResourceServerId;
    }
    @Export(name="enabled", type=Boolean.class, parameters={})
    private Output<Boolean> enabled;

    public Output<Boolean> enabled() {
        return this.enabled;
    }
    @Export(name="groupId", type=String.class, parameters={})
    private Output<String> groupId;

    public Output<String> groupId() {
        return this.groupId;
    }
    @Export(name="manageMembersScope", type=GroupPermissionsManageMembersScope.class, parameters={})
    private Output</* @Nullable */ GroupPermissionsManageMembersScope> manageMembersScope;

    public Output<Optional<GroupPermissionsManageMembersScope>> manageMembersScope() {
        return Codegen.optional(this.manageMembersScope);
    }
    @Export(name="manageMembershipScope", type=GroupPermissionsManageMembershipScope.class, parameters={})
    private Output</* @Nullable */ GroupPermissionsManageMembershipScope> manageMembershipScope;

    public Output<Optional<GroupPermissionsManageMembershipScope>> manageMembershipScope() {
        return Codegen.optional(this.manageMembershipScope);
    }
    @Export(name="manageScope", type=GroupPermissionsManageScope.class, parameters={})
    private Output</* @Nullable */ GroupPermissionsManageScope> manageScope;

    public Output<Optional<GroupPermissionsManageScope>> manageScope() {
        return Codegen.optional(this.manageScope);
    }
    @Export(name="realmId", type=String.class, parameters={})
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="viewMembersScope", type=GroupPermissionsViewMembersScope.class, parameters={})
    private Output</* @Nullable */ GroupPermissionsViewMembersScope> viewMembersScope;

    public Output<Optional<GroupPermissionsViewMembersScope>> viewMembersScope() {
        return Codegen.optional(this.viewMembersScope);
    }
    @Export(name="viewScope", type=GroupPermissionsViewScope.class, parameters={})
    private Output</* @Nullable */ GroupPermissionsViewScope> viewScope;

    public Output<Optional<GroupPermissionsViewScope>> viewScope() {
        return Codegen.optional(this.viewScope);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public GroupPermissions(String name) {
        this(name, GroupPermissionsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public GroupPermissions(String name, GroupPermissionsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public GroupPermissions(String name, GroupPermissionsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/groupPermissions:GroupPermissions", name, args == null ? GroupPermissionsArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private GroupPermissions(String name, Output<String> id, @Nullable GroupPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/groupPermissions:GroupPermissions", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
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
    public static GroupPermissions get(String name, Output<String> id, @Nullable GroupPermissionsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new GroupPermissions(name, id, state, options);
    }
}
