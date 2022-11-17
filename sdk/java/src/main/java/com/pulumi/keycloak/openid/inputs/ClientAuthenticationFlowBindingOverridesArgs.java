// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientAuthenticationFlowBindingOverridesArgs extends com.pulumi.resources.ResourceArgs {

    public static final ClientAuthenticationFlowBindingOverridesArgs Empty = new ClientAuthenticationFlowBindingOverridesArgs();

    /**
     * Browser flow id, (flow needs to exist)
     * 
     */
    @Import(name="browserId")
    private @Nullable Output<String> browserId;

    /**
     * @return Browser flow id, (flow needs to exist)
     * 
     */
    public Optional<Output<String>> browserId() {
        return Optional.ofNullable(this.browserId);
    }

    /**
     * Direct grant flow id (flow needs to exist)
     * 
     */
    @Import(name="directGrantId")
    private @Nullable Output<String> directGrantId;

    /**
     * @return Direct grant flow id (flow needs to exist)
     * 
     */
    public Optional<Output<String>> directGrantId() {
        return Optional.ofNullable(this.directGrantId);
    }

    private ClientAuthenticationFlowBindingOverridesArgs() {}

    private ClientAuthenticationFlowBindingOverridesArgs(ClientAuthenticationFlowBindingOverridesArgs $) {
        this.browserId = $.browserId;
        this.directGrantId = $.directGrantId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientAuthenticationFlowBindingOverridesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientAuthenticationFlowBindingOverridesArgs $;

        public Builder() {
            $ = new ClientAuthenticationFlowBindingOverridesArgs();
        }

        public Builder(ClientAuthenticationFlowBindingOverridesArgs defaults) {
            $ = new ClientAuthenticationFlowBindingOverridesArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param browserId Browser flow id, (flow needs to exist)
         * 
         * @return builder
         * 
         */
        public Builder browserId(@Nullable Output<String> browserId) {
            $.browserId = browserId;
            return this;
        }

        /**
         * @param browserId Browser flow id, (flow needs to exist)
         * 
         * @return builder
         * 
         */
        public Builder browserId(String browserId) {
            return browserId(Output.of(browserId));
        }

        /**
         * @param directGrantId Direct grant flow id (flow needs to exist)
         * 
         * @return builder
         * 
         */
        public Builder directGrantId(@Nullable Output<String> directGrantId) {
            $.directGrantId = directGrantId;
            return this;
        }

        /**
         * @param directGrantId Direct grant flow id (flow needs to exist)
         * 
         * @return builder
         * 
         */
        public Builder directGrantId(String directGrantId) {
            return directGrantId(Output.of(directGrantId));
        }

        public ClientAuthenticationFlowBindingOverridesArgs build() {
            return $;
        }
    }

}