// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UsersPermissionsImpersonateScope {
    private final @Nullable String decisionStrategy;
    private final @Nullable String description;
    private final @Nullable List<String> policies;

    @CustomType.Constructor
    private UsersPermissionsImpersonateScope(
        @CustomType.Parameter("decisionStrategy") @Nullable String decisionStrategy,
        @CustomType.Parameter("description") @Nullable String description,
        @CustomType.Parameter("policies") @Nullable List<String> policies) {
        this.decisionStrategy = decisionStrategy;
        this.description = description;
        this.policies = policies;
    }

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

    public static Builder builder(UsersPermissionsImpersonateScope defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String decisionStrategy;
        private @Nullable String description;
        private @Nullable List<String> policies;

        public Builder() {
    	      // Empty
        }

        public Builder(UsersPermissionsImpersonateScope defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.decisionStrategy = defaults.decisionStrategy;
    	      this.description = defaults.description;
    	      this.policies = defaults.policies;
        }

        public Builder decisionStrategy(@Nullable String decisionStrategy) {
            this.decisionStrategy = decisionStrategy;
            return this;
        }
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        public Builder policies(@Nullable List<String> policies) {
            this.policies = policies;
            return this;
        }
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }        public UsersPermissionsImpersonateScope build() {
            return new UsersPermissionsImpersonateScope(decisionStrategy, description, policies);
        }
    }
}
