// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UserInitialPassword {
    /**
     * @return If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
     * 
     */
    private final @Nullable Boolean temporary;
    /**
     * @return The initial password.
     * 
     */
    private final String value;

    @CustomType.Constructor
    private UserInitialPassword(
        @CustomType.Parameter("temporary") @Nullable Boolean temporary,
        @CustomType.Parameter("value") String value) {
        this.temporary = temporary;
        this.value = value;
    }

    /**
     * @return If set to `true`, the initial password is set up for renewal on first use. Default to `false`.
     * 
     */
    public Optional<Boolean> temporary() {
        return Optional.ofNullable(this.temporary);
    }
    /**
     * @return The initial password.
     * 
     */
    public String value() {
        return this.value;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UserInitialPassword defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Boolean temporary;
        private String value;

        public Builder() {
    	      // Empty
        }

        public Builder(UserInitialPassword defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.temporary = defaults.temporary;
    	      this.value = defaults.value;
        }

        public Builder temporary(@Nullable Boolean temporary) {
            this.temporary = temporary;
            return this;
        }
        public Builder value(String value) {
            this.value = Objects.requireNonNull(value);
            return this;
        }        public UserInitialPassword build() {
            return new UserInitialPassword(temporary, value);
        }
    }
}
