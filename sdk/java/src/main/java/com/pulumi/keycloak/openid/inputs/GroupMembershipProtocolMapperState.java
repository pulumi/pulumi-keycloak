// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMembershipProtocolMapperState extends com.pulumi.resources.ResourceArgs {

    public static final GroupMembershipProtocolMapperState Empty = new GroupMembershipProtocolMapperState();

    @Import(name="addToAccessToken")
    private @Nullable Output<Boolean> addToAccessToken;

    public Optional<Output<Boolean>> addToAccessToken() {
        return Optional.ofNullable(this.addToAccessToken);
    }

    @Import(name="addToIdToken")
    private @Nullable Output<Boolean> addToIdToken;

    public Optional<Output<Boolean>> addToIdToken() {
        return Optional.ofNullable(this.addToIdToken);
    }

    @Import(name="addToUserinfo")
    private @Nullable Output<Boolean> addToUserinfo;

    public Optional<Output<Boolean>> addToUserinfo() {
        return Optional.ofNullable(this.addToUserinfo);
    }

    @Import(name="claimName")
    private @Nullable Output<String> claimName;

    public Optional<Output<String>> claimName() {
        return Optional.ofNullable(this.claimName);
    }

    /**
     * The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
     * 
     */
    @Import(name="clientScopeId")
    private @Nullable Output<String> clientScopeId;

    /**
     * @return The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
     * 
     */
    public Optional<Output<String>> clientScopeId() {
        return Optional.ofNullable(this.clientScopeId);
    }

    @Import(name="fullPath")
    private @Nullable Output<Boolean> fullPath;

    public Optional<Output<Boolean>> fullPath() {
        return Optional.ofNullable(this.fullPath);
    }

    /**
     * A human-friendly name that will appear in the Keycloak console.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return A human-friendly name that will appear in the Keycloak console.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The realm id where the associated client or client scope exists.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm id where the associated client or client scope exists.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private GroupMembershipProtocolMapperState() {}

    private GroupMembershipProtocolMapperState(GroupMembershipProtocolMapperState $) {
        this.addToAccessToken = $.addToAccessToken;
        this.addToIdToken = $.addToIdToken;
        this.addToUserinfo = $.addToUserinfo;
        this.claimName = $.claimName;
        this.clientId = $.clientId;
        this.clientScopeId = $.clientScopeId;
        this.fullPath = $.fullPath;
        this.name = $.name;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMembershipProtocolMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMembershipProtocolMapperState $;

        public Builder() {
            $ = new GroupMembershipProtocolMapperState();
        }

        public Builder(GroupMembershipProtocolMapperState defaults) {
            $ = new GroupMembershipProtocolMapperState(Objects.requireNonNull(defaults));
        }

        public Builder addToAccessToken(@Nullable Output<Boolean> addToAccessToken) {
            $.addToAccessToken = addToAccessToken;
            return this;
        }

        public Builder addToAccessToken(Boolean addToAccessToken) {
            return addToAccessToken(Output.of(addToAccessToken));
        }

        public Builder addToIdToken(@Nullable Output<Boolean> addToIdToken) {
            $.addToIdToken = addToIdToken;
            return this;
        }

        public Builder addToIdToken(Boolean addToIdToken) {
            return addToIdToken(Output.of(addToIdToken));
        }

        public Builder addToUserinfo(@Nullable Output<Boolean> addToUserinfo) {
            $.addToUserinfo = addToUserinfo;
            return this;
        }

        public Builder addToUserinfo(Boolean addToUserinfo) {
            return addToUserinfo(Output.of(addToUserinfo));
        }

        public Builder claimName(@Nullable Output<String> claimName) {
            $.claimName = claimName;
            return this;
        }

        public Builder claimName(String claimName) {
            return claimName(Output.of(claimName));
        }

        /**
         * @param clientId The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The mapper&#39;s associated client. Cannot be used at the same time as client_scope_id.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param clientScopeId The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(@Nullable Output<String> clientScopeId) {
            $.clientScopeId = clientScopeId;
            return this;
        }

        /**
         * @param clientScopeId The mapper&#39;s associated client scope. Cannot be used at the same time as client_id.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(String clientScopeId) {
            return clientScopeId(Output.of(clientScopeId));
        }

        public Builder fullPath(@Nullable Output<Boolean> fullPath) {
            $.fullPath = fullPath;
            return this;
        }

        public Builder fullPath(Boolean fullPath) {
            return fullPath(Output.of(fullPath));
        }

        /**
         * @param name A human-friendly name that will appear in the Keycloak console.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name A human-friendly name that will appear in the Keycloak console.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm id where the associated client or client scope exists.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm id where the associated client or client scope exists.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public GroupMembershipProtocolMapperState build() {
            return $;
        }
    }

}
