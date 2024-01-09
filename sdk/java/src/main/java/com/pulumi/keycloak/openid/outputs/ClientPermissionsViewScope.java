// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class ClientPermissionsViewScope {
    private @Nullable String decisionStrategy;
    private @Nullable String description;
    private @Nullable List<String> policies;

    private ClientPermissionsViewScope() {}
    public Optional<String> decisionStrategy() {
        return Optional.ofNullable(this.decisionStrategy);
    }
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    public List<String> policies() {
        return this.policies == null ? List.of() : this.policies;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClientPermissionsViewScope defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String decisionStrategy;
        private @Nullable String description;
        private @Nullable List<String> policies;
        public Builder() {}
        public Builder(ClientPermissionsViewScope defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.decisionStrategy = defaults.decisionStrategy;
    	      this.description = defaults.description;
    	      this.policies = defaults.policies;
        }

        @CustomType.Setter
        public Builder decisionStrategy(@Nullable String decisionStrategy) {

            this.decisionStrategy = decisionStrategy;
            return this;
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder policies(@Nullable List<String> policies) {

            this.policies = policies;
            return this;
        }
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }
        public ClientPermissionsViewScope build() {
            final var _resultValue = new ClientPermissionsViewScope();
            _resultValue.decisionStrategy = decisionStrategy;
            _resultValue.description = description;
            _resultValue.policies = policies;
            return _resultValue;
        }
    }
}
