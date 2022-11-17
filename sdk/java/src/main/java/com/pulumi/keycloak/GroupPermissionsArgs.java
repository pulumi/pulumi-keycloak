// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.keycloak.inputs.GroupPermissionsManageMembersScopeArgs;
import com.pulumi.keycloak.inputs.GroupPermissionsManageMembershipScopeArgs;
import com.pulumi.keycloak.inputs.GroupPermissionsManageScopeArgs;
import com.pulumi.keycloak.inputs.GroupPermissionsViewMembersScopeArgs;
import com.pulumi.keycloak.inputs.GroupPermissionsViewScopeArgs;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupPermissionsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupPermissionsArgs Empty = new GroupPermissionsArgs();

    @Import(name="groupId", required=true)
    private Output<String> groupId;

    public Output<String> groupId() {
        return this.groupId;
    }

    @Import(name="manageMembersScope")
    private @Nullable Output<GroupPermissionsManageMembersScopeArgs> manageMembersScope;

    public Optional<Output<GroupPermissionsManageMembersScopeArgs>> manageMembersScope() {
        return Optional.ofNullable(this.manageMembersScope);
    }

    @Import(name="manageMembershipScope")
    private @Nullable Output<GroupPermissionsManageMembershipScopeArgs> manageMembershipScope;

    public Optional<Output<GroupPermissionsManageMembershipScopeArgs>> manageMembershipScope() {
        return Optional.ofNullable(this.manageMembershipScope);
    }

    @Import(name="manageScope")
    private @Nullable Output<GroupPermissionsManageScopeArgs> manageScope;

    public Optional<Output<GroupPermissionsManageScopeArgs>> manageScope() {
        return Optional.ofNullable(this.manageScope);
    }

    @Import(name="realmId", required=true)
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }

    @Import(name="viewMembersScope")
    private @Nullable Output<GroupPermissionsViewMembersScopeArgs> viewMembersScope;

    public Optional<Output<GroupPermissionsViewMembersScopeArgs>> viewMembersScope() {
        return Optional.ofNullable(this.viewMembersScope);
    }

    @Import(name="viewScope")
    private @Nullable Output<GroupPermissionsViewScopeArgs> viewScope;

    public Optional<Output<GroupPermissionsViewScopeArgs>> viewScope() {
        return Optional.ofNullable(this.viewScope);
    }

    private GroupPermissionsArgs() {}

    private GroupPermissionsArgs(GroupPermissionsArgs $) {
        this.groupId = $.groupId;
        this.manageMembersScope = $.manageMembersScope;
        this.manageMembershipScope = $.manageMembershipScope;
        this.manageScope = $.manageScope;
        this.realmId = $.realmId;
        this.viewMembersScope = $.viewMembersScope;
        this.viewScope = $.viewScope;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupPermissionsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupPermissionsArgs $;

        public Builder() {
            $ = new GroupPermissionsArgs();
        }

        public Builder(GroupPermissionsArgs defaults) {
            $ = new GroupPermissionsArgs(Objects.requireNonNull(defaults));
        }

        public Builder groupId(Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        public Builder manageMembersScope(@Nullable Output<GroupPermissionsManageMembersScopeArgs> manageMembersScope) {
            $.manageMembersScope = manageMembersScope;
            return this;
        }

        public Builder manageMembersScope(GroupPermissionsManageMembersScopeArgs manageMembersScope) {
            return manageMembersScope(Output.of(manageMembersScope));
        }

        public Builder manageMembershipScope(@Nullable Output<GroupPermissionsManageMembershipScopeArgs> manageMembershipScope) {
            $.manageMembershipScope = manageMembershipScope;
            return this;
        }

        public Builder manageMembershipScope(GroupPermissionsManageMembershipScopeArgs manageMembershipScope) {
            return manageMembershipScope(Output.of(manageMembershipScope));
        }

        public Builder manageScope(@Nullable Output<GroupPermissionsManageScopeArgs> manageScope) {
            $.manageScope = manageScope;
            return this;
        }

        public Builder manageScope(GroupPermissionsManageScopeArgs manageScope) {
            return manageScope(Output.of(manageScope));
        }

        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder viewMembersScope(@Nullable Output<GroupPermissionsViewMembersScopeArgs> viewMembersScope) {
            $.viewMembersScope = viewMembersScope;
            return this;
        }

        public Builder viewMembersScope(GroupPermissionsViewMembersScopeArgs viewMembersScope) {
            return viewMembersScope(Output.of(viewMembersScope));
        }

        public Builder viewScope(@Nullable Output<GroupPermissionsViewScopeArgs> viewScope) {
            $.viewScope = viewScope;
            return this;
        }

        public Builder viewScope(GroupPermissionsViewScopeArgs viewScope) {
            return viewScope(Output.of(viewScope));
        }

        public GroupPermissionsArgs build() {
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}