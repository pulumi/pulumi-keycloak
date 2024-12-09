// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMembershipsState extends com.pulumi.resources.ResourceArgs {

    public static final GroupMembershipsState Empty = new GroupMembershipsState();

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
    @Import(name="members")
    private @Nullable Output<List<String>> members;

    /**
     * @return A list of usernames that belong to this group.
     * 
     */
    public Optional<Output<List<String>>> members() {
        return Optional.ofNullable(this.members);
    }

    /**
     * The realm this group exists in.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm this group exists in.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private GroupMembershipsState() {}

    private GroupMembershipsState(GroupMembershipsState $) {
        this.groupId = $.groupId;
        this.members = $.members;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMembershipsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMembershipsState $;

        public Builder() {
            $ = new GroupMembershipsState();
        }

        public Builder(GroupMembershipsState defaults) {
            $ = new GroupMembershipsState(Objects.requireNonNull(defaults));
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
        public Builder members(@Nullable Output<List<String>> members) {
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
        public Builder realmId(@Nullable Output<String> realmId) {
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

        public GroupMembershipsState build() {
            return $;
        }
    }

}
