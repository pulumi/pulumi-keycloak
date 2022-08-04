// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.saml.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientDefaultScopeState extends com.pulumi.resources.ResourceArgs {

    public static final ClientDefaultScopeState Empty = new ClientDefaultScopeState();

    /**
     * The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * An array of client scope names to attach to this client.
     * 
     */
    @Import(name="defaultScopes")
    private @Nullable Output<List<String>> defaultScopes;

    /**
     * @return An array of client scope names to attach to this client.
     * 
     */
    public Optional<Output<List<String>>> defaultScopes() {
        return Optional.ofNullable(this.defaultScopes);
    }

    /**
     * The realm this client and scopes exists in.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm this client and scopes exists in.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    private ClientDefaultScopeState() {}

    private ClientDefaultScopeState(ClientDefaultScopeState $) {
        this.clientId = $.clientId;
        this.defaultScopes = $.defaultScopes;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientDefaultScopeState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientDefaultScopeState $;

        public Builder() {
            $ = new ClientDefaultScopeState();
        }

        public Builder(ClientDefaultScopeState defaults) {
            $ = new ClientDefaultScopeState(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientId The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The ID of the client to attach default scopes to. Note that this is the unique ID of the client generated by Keycloak.
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param defaultScopes An array of client scope names to attach to this client.
         * 
         * @return builder
         * 
         */
        public Builder defaultScopes(@Nullable Output<List<String>> defaultScopes) {
            $.defaultScopes = defaultScopes;
            return this;
        }

        /**
         * @param defaultScopes An array of client scope names to attach to this client.
         * 
         * @return builder
         * 
         */
        public Builder defaultScopes(List<String> defaultScopes) {
            return defaultScopes(Output.of(defaultScopes));
        }

        /**
         * @param defaultScopes An array of client scope names to attach to this client.
         * 
         * @return builder
         * 
         */
        public Builder defaultScopes(String... defaultScopes) {
            return defaultScopes(List.of(defaultScopes));
        }

        /**
         * @param realmId The realm this client and scopes exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this client and scopes exists in.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public ClientDefaultScopeState build() {
            return $;
        }
    }

}
