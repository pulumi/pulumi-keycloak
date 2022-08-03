// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupRolesArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupRolesArgs Empty = new GroupRolesArgs();

    /**
     * Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the group will be removed. Defaults to `true`.
     * 
     */
    @Import(name="exhaustive")
    private @Nullable Output<Boolean> exhaustive;

    /**
     * @return Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the group will be removed. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> exhaustive() {
        return Optional.ofNullable(this.exhaustive);
    }

    /**
     * The ID of the group this resource should manage roles for.
     * 
     */
    @Import(name="groupId", required=true)
    private Output<String> groupId;

    /**
     * @return The ID of the group this resource should manage roles for.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
    }

    /**
     * The realm this group exists in.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this group exists in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * A list of role IDs to map to the group.
     * 
     */
    @Import(name="roleIds", required=true)
    private Output<List<String>> roleIds;

    /**
     * @return A list of role IDs to map to the group.
     * 
     */
    public Output<List<String>> roleIds() {
        return this.roleIds;
    }

    private GroupRolesArgs() {}

    private GroupRolesArgs(GroupRolesArgs $) {
        this.exhaustive = $.exhaustive;
        this.groupId = $.groupId;
        this.realmId = $.realmId;
        this.roleIds = $.roleIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupRolesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupRolesArgs $;

        public Builder() {
            $ = new GroupRolesArgs();
        }

        public Builder(GroupRolesArgs defaults) {
            $ = new GroupRolesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param exhaustive Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the group will be removed. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder exhaustive(@Nullable Output<Boolean> exhaustive) {
            $.exhaustive = exhaustive;
            return this;
        }

        /**
         * @param exhaustive Indicates if the list of roles is exhaustive. In this case, roles that are manually added to the group will be removed. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder exhaustive(Boolean exhaustive) {
            return exhaustive(Output.of(exhaustive));
        }

        /**
         * @param groupId The ID of the group this resource should manage roles for.
         * 
         * @return builder
         * 
         */
        public Builder groupId(Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the group this resource should manage roles for.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param realmId The realm this group exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this group exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param roleIds A list of role IDs to map to the group.
         * 
         * @return builder
         * 
         */
        public Builder roleIds(Output<List<String>> roleIds) {
            $.roleIds = roleIds;
            return this;
        }

        /**
         * @param roleIds A list of role IDs to map to the group.
         * 
         * @return builder
         * 
         */
        public Builder roleIds(List<String> roleIds) {
            return roleIds(Output.of(roleIds));
        }

        /**
         * @param roleIds A list of role IDs to map to the group.
         * 
         * @return builder
         * 
         */
        public Builder roleIds(String... roleIds) {
            return roleIds(List.of(roleIds));
        }

        public GroupRolesArgs build() {
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            $.roleIds = Objects.requireNonNull($.roleIds, "expected parameter 'roleIds' to be non-null");
            return $;
        }
    }

}
