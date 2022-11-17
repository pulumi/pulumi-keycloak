// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMembershipsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupMembershipsArgs Empty = new GroupMembershipsArgs();

    /**
     * The ID of the group this resource should manage memberships for.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return The ID of the group this resource should manage memberships for.
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * A list of usernames that belong to this group.
     * 
     */
    @Import(name="members", required=true)
    private Output<List<String>> members;

    /**
     * @return A list of usernames that belong to this group.
     * 
     */
    public Output<List<String>> members() {
        return this.members;
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

    private GroupMembershipsArgs() {}

    private GroupMembershipsArgs(GroupMembershipsArgs $) {
        this.groupId = $.groupId;
        this.members = $.members;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMembershipsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMembershipsArgs $;

        public Builder() {
            $ = new GroupMembershipsArgs();
        }

        public Builder(GroupMembershipsArgs defaults) {
            $ = new GroupMembershipsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groupId The ID of the group this resource should manage memberships for.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId The ID of the group this resource should manage memberships for.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param members A list of usernames that belong to this group.
         * 
         * @return builder
         * 
         */
        public Builder members(Output<List<String>> members) {
            $.members = members;
            return this;
        }

        /**
         * @param members A list of usernames that belong to this group.
         * 
         * @return builder
         * 
         */
        public Builder members(List<String> members) {
            return members(Output.of(members));
        }

        /**
         * @param members A list of usernames that belong to this group.
         * 
         * @return builder
         * 
         */
        public Builder members(String... members) {
            return members(List.of(members));
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

        public GroupMembershipsArgs build() {
            $.members = Objects.requireNonNull($.members, "expected parameter 'members' to be non-null");
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}