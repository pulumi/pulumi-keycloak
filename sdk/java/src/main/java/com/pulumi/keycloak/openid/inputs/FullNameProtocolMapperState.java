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


public final class FullNameProtocolMapperState extends com.pulumi.resources.ResourceArgs {

    public static final FullNameProtocolMapperState Empty = new FullNameProtocolMapperState();

    /**
     * Indicates if the user&#39;s full name should be added as a claim to the access token. Defaults to `true`.
     * 
     */
    @Import(name="addToAccessToken")
    private @Nullable Output<Boolean> addToAccessToken;

    /**
     * @return Indicates if the user&#39;s full name should be added as a claim to the access token. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> addToAccessToken() {
        return Optional.ofNullable(this.addToAccessToken);
    }

    /**
     * Indicates if the user&#39;s full name should be added as a claim to the id token. Defaults to `true`.
     * 
     */
    @Import(name="addToIdToken")
    private @Nullable Output<Boolean> addToIdToken;

    /**
     * @return Indicates if the user&#39;s full name should be added as a claim to the id token. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> addToIdToken() {
        return Optional.ofNullable(this.addToIdToken);
    }

    /**
     * Indicates if the user&#39;s full name should be added as a claim to the UserInfo response body. Defaults to `true`.
     * 
     */
    @Import(name="addToUserinfo")
    private @Nullable Output<Boolean> addToUserinfo;

    /**
     * @return Indicates if the user&#39;s full name should be added as a claim to the UserInfo response body. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> addToUserinfo() {
        return Optional.ofNullable(this.addToUserinfo);
    }

    /**
     * The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    @Import(name="clientScopeId")
    private @Nullable Output<String> clientScopeId;

    /**
     * @return The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
     * 
     */
    public Optional<Output<String>> clientScopeId() {
        return Optional.ofNullable(this.clientScopeId);
    }

    /**
     * The display name of this protocol mapper in the GUI.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The display name of this protocol mapper in the GUI.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The realm this protocol mapper exists within.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm this protocol mapper exists within.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private FullNameProtocolMapperState() {}

    private FullNameProtocolMapperState(FullNameProtocolMapperState $) {
        this.addToAccessToken = $.addToAccessToken;
        this.addToIdToken = $.addToIdToken;
        this.addToUserinfo = $.addToUserinfo;
        this.clientId = $.clientId;
        this.clientScopeId = $.clientScopeId;
        this.name = $.name;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(FullNameProtocolMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private FullNameProtocolMapperState $;

        public Builder() {
            $ = new FullNameProtocolMapperState();
        }

        public Builder(FullNameProtocolMapperState defaults) {
            $ = new FullNameProtocolMapperState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addToAccessToken Indicates if the user&#39;s full name should be added as a claim to the access token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(@Nullable Output<Boolean> addToAccessToken) {
            $.addToAccessToken = addToAccessToken;
            return this;
        }

        /**
         * @param addToAccessToken Indicates if the user&#39;s full name should be added as a claim to the access token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(Boolean addToAccessToken) {
            return addToAccessToken(Output.of(addToAccessToken));
        }

        /**
         * @param addToIdToken Indicates if the user&#39;s full name should be added as a claim to the id token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(@Nullable Output<Boolean> addToIdToken) {
            $.addToIdToken = addToIdToken;
            return this;
        }

        /**
         * @param addToIdToken Indicates if the user&#39;s full name should be added as a claim to the id token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(Boolean addToIdToken) {
            return addToIdToken(Output.of(addToIdToken));
        }

        /**
         * @param addToUserinfo Indicates if the user&#39;s full name should be added as a claim to the UserInfo response body. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToUserinfo(@Nullable Output<Boolean> addToUserinfo) {
            $.addToUserinfo = addToUserinfo;
            return this;
        }

        /**
         * @param addToUserinfo Indicates if the user&#39;s full name should be added as a claim to the UserInfo response body. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToUserinfo(Boolean addToUserinfo) {
            return addToUserinfo(Output.of(addToUserinfo));
        }

        /**
         * @param clientId The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The client this protocol mapper should be attached to. Conflicts with `client_scope_id`. One of `client_id` or `client_scope_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param clientScopeId The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(@Nullable Output<String> clientScopeId) {
            $.clientScopeId = clientScopeId;
            return this;
        }

        /**
         * @param clientScopeId The client scope this protocol mapper should be attached to. Conflicts with `client_id`. One of `client_id` or `client_scope_id` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder clientScopeId(String clientScopeId) {
            return clientScopeId(Output.of(clientScopeId));
        }

        /**
         * @param name The display name of this protocol mapper in the GUI.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The display name of this protocol mapper in the GUI.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm this protocol mapper exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this protocol mapper exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public FullNameProtocolMapperState build() {
            return $;
        }
    }

}
