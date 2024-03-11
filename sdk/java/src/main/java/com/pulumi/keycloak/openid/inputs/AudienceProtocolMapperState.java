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


public final class AudienceProtocolMapperState extends com.pulumi.resources.ResourceArgs {

    public static final AudienceProtocolMapperState Empty = new AudienceProtocolMapperState();

    /**
     * Indicates if this claim should be added to the access token.
     * 
     */
    @Import(name="addToAccessToken")
    private @Nullable Output<Boolean> addToAccessToken;

    /**
     * @return Indicates if this claim should be added to the access token.
     * 
     */
    public Optional<Output<Boolean>> addToAccessToken() {
        return Optional.ofNullable(this.addToAccessToken);
    }

    /**
     * Indicates if this claim should be added to the id token.
     * 
     */
    @Import(name="addToIdToken")
    private @Nullable Output<Boolean> addToIdToken;

    /**
     * @return Indicates if this claim should be added to the id token.
     * 
     */
    public Optional<Output<Boolean>> addToIdToken() {
        return Optional.ofNullable(this.addToIdToken);
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

    /**
     * A client ID to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
     * 
     */
    @Import(name="includedClientAudience")
    private @Nullable Output<String> includedClientAudience;

    /**
     * @return A client ID to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
     * 
     */
    public Optional<Output<String>> includedClientAudience() {
        return Optional.ofNullable(this.includedClientAudience);
    }

    /**
     * A custom audience to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
     * 
     */
    @Import(name="includedCustomAudience")
    private @Nullable Output<String> includedCustomAudience;

    /**
     * @return A custom audience to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
     * 
     */
    public Optional<Output<String>> includedCustomAudience() {
        return Optional.ofNullable(this.includedCustomAudience);
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

    private AudienceProtocolMapperState() {}

    private AudienceProtocolMapperState(AudienceProtocolMapperState $) {
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
    public static Builder builder(AudienceProtocolMapperState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AudienceProtocolMapperState $;

        public Builder() {
            $ = new AudienceProtocolMapperState();
        }

        public Builder(AudienceProtocolMapperState defaults) {
            $ = new AudienceProtocolMapperState(Objects.requireNonNull(defaults));
        }

        /**
         * @param addToAccessToken Indicates if this claim should be added to the access token.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(@Nullable Output<Boolean> addToAccessToken) {
            $.addToAccessToken = addToAccessToken;
            return this;
        }

        /**
         * @param addToAccessToken Indicates if this claim should be added to the access token.
         * 
         * @return builder
         * 
         */
        public Builder addToAccessToken(Boolean addToAccessToken) {
            return addToAccessToken(Output.of(addToAccessToken));
        }

        /**
         * @param addToIdToken Indicates if this claim should be added to the id token.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(@Nullable Output<Boolean> addToIdToken) {
            $.addToIdToken = addToIdToken;
            return this;
        }

        /**
         * @param addToIdToken Indicates if this claim should be added to the id token.
         * 
         * @return builder
         * 
         */
        public Builder addToIdToken(Boolean addToIdToken) {
            return addToIdToken(Output.of(addToIdToken));
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

        /**
         * @param includedClientAudience A client ID to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
         * 
         * @return builder
         * 
         */
        public Builder includedClientAudience(@Nullable Output<String> includedClientAudience) {
            $.includedClientAudience = includedClientAudience;
            return this;
        }

        /**
         * @param includedClientAudience A client ID to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
         * 
         * @return builder
         * 
         */
        public Builder includedClientAudience(String includedClientAudience) {
            return includedClientAudience(Output.of(includedClientAudience));
        }

        /**
         * @param includedCustomAudience A custom audience to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
         * 
         * @return builder
         * 
         */
        public Builder includedCustomAudience(@Nullable Output<String> includedCustomAudience) {
            $.includedCustomAudience = includedCustomAudience;
            return this;
        }

        /**
         * @param includedCustomAudience A custom audience to include within the token&#39;s `aud` claim. Cannot be used with included_custom_audience
         * 
         * @return builder
         * 
         */
        public Builder includedCustomAudience(String includedCustomAudience) {
            return includedCustomAudience(Output.of(includedCustomAudience));
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

        public AudienceProtocolMapperState build() {
            return $;
        }
    }

}
