// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AudienceProtocolMapperArgs extends com.pulumi.resources.ResourceArgs {

    public static final AudienceProtocolMapperArgs Empty = new AudienceProtocolMapperArgs();

    /**
     * Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     * 
     */
    @Import(name="addToAccessToken")
    private @Nullable Output<Boolean> addToAccessToken;

    /**
     * @return Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> addToAccessToken() {
        return Optional.ofNullable(this.addToAccessToken);
    }

    /**
     * Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     * 
     */
    @Import(name="addToIdToken")
    private @Nullable Output<Boolean> addToIdToken;

    /**
     * @return Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> addToIdToken() {
        return Optional.ofNullable(this.addToIdToken);
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
     * A client ID to include within the token&#39;s `aud` claim. Conflicts with `included_custom_audience`. One of `included_client_audience` or `included_custom_audience` must be specified.
     * 
     */
    @Import(name="includedClientAudience")
    private @Nullable Output<String> includedClientAudience;

    /**
     * @return A client ID to include within the token&#39;s `aud` claim. Conflicts with `included_custom_audience`. One of `included_client_audience` or `included_custom_audience` must be specified.
     * 
     */
    public Optional<Output<String>> includedClientAudience() {
        return Optional.ofNullable(this.includedClientAudience);
    }

    /**
     * A custom audience to include within the token&#39;s `aud` claim. Conflicts with `included_client_audience`. One of `included_client_audience` or `included_custom_audience` must be specified.
     * 
     */
    @Import(name="includedCustomAudience")
    private @Nullable Output<String> includedCustomAudience;

    /**
     * @return A custom audience to include within the token&#39;s `aud` claim. Conflicts with `included_client_audience`. One of `included_client_audience` or `included_custom_audience` must be specified.
     * 
     */
    public Optional<Output<String>> includedCustomAudience() {
        return Optional.ofNullable(this.includedCustomAudience);
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
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this protocol mapper exists within.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private AudienceProtocolMapperArgs() {}

    private AudienceProtocolMapperArgs(AudienceProtocolMapperArgs $) {
        this.addToAccessToken = $.addToAccessToken;
        this.addToIdToken = $.addToIdToken;
        this.clientId = $.clientId;
        this.clientScopeId = $.clientScopeId;
        this.includedClientAudience = $.includedClientAudience;
        this.includedCustomAudience = $.includedCustomAudience;
        this.name = $.name;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AudienceProtocolMapperArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AudienceProtocolMapperArgs $;

        public Builder() {
            $ = new AudienceProtocolMapperArgs();
        }

        public Builder(AudienceProtocolMapperArgs defaults) {
            $ = new AudienceProtocolMapperArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param addToAccessToken Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(@Nullable Output<Boolean> addToAccessToken) {
            $.addToAccessToken = addToAccessToken;
            return this;
        }

        /**
         * @param addToAccessToken Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(Boolean addToAccessToken) {
            return addToAccessToken(Output.of(addToAccessToken));
        }

        /**
         * @param addToIdToken Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(@Nullable Output<Boolean> addToIdToken) {
            $.addToIdToken = addToIdToken;
            return this;
        }

        /**
         * @param addToIdToken Indicates if the audience should be included in the `aud` claim for the id token. Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(Boolean addToIdToken) {
            return addToIdToken(Output.of(addToIdToken));
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
         * @param includedClientAudience A client ID to include within the token&#39;s `aud` claim. Conflicts with `included_custom_audience`. One of `included_client_audience` or `included_custom_audience` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder includedClientAudience(@Nullable Output<String> includedClientAudience) {
            $.includedClientAudience = includedClientAudience;
            return this;
        }

        /**
         * @param includedClientAudience A client ID to include within the token&#39;s `aud` claim. Conflicts with `included_custom_audience`. One of `included_client_audience` or `included_custom_audience` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder includedClientAudience(String includedClientAudience) {
            return includedClientAudience(Output.of(includedClientAudience));
        }

        /**
         * @param includedCustomAudience A custom audience to include within the token&#39;s `aud` claim. Conflicts with `included_client_audience`. One of `included_client_audience` or `included_custom_audience` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder includedCustomAudience(@Nullable Output<String> includedCustomAudience) {
            $.includedCustomAudience = includedCustomAudience;
            return this;
        }

        /**
         * @param includedCustomAudience A custom audience to include within the token&#39;s `aud` claim. Conflicts with `included_client_audience`. One of `included_client_audience` or `included_custom_audience` must be specified.
         * 
         * @return builder
         * 
         */
        public Builder includedCustomAudience(String includedCustomAudience) {
            return includedCustomAudience(Output.of(includedCustomAudience));
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
        public Builder realmId(Output<String> realmId) {
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

        public AudienceProtocolMapperArgs build() {
            $.realmId = Objects.requireNonNull($.realmId, "expected parameter 'realmId' to be non-null");
            return $;
        }
    }

}
