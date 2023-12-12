// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class UserFederationCache {
    /**
     * @return Day of the week the entry will become invalid on
     * 
     */
    private @Nullable Integer evictionDay;
    /**
     * @return Hour of day the entry will become invalid on.
     * 
     */
    private @Nullable Integer evictionHour;
    /**
     * @return Minute of day the entry will become invalid on.
     * 
     */
    private @Nullable Integer evictionMinute;
    /**
     * @return Max lifespan of cache entry (duration string).
     * 
     */
    private @Nullable String maxLifespan;
    /**
     * @return Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     * 
     */
    private @Nullable String policy;

    private UserFederationCache() {}
    /**
     * @return Day of the week the entry will become invalid on
     * 
     */
    public Optional<Integer> evictionDay() {
        return Optional.ofNullable(this.evictionDay);
    }
    /**
     * @return Hour of day the entry will become invalid on.
     * 
     */
    public Optional<Integer> evictionHour() {
        return Optional.ofNullable(this.evictionHour);
    }
    /**
     * @return Minute of day the entry will become invalid on.
     * 
     */
    public Optional<Integer> evictionMinute() {
        return Optional.ofNullable(this.evictionMinute);
    }
    /**
     * @return Max lifespan of cache entry (duration string).
     * 
     */
    public Optional<String> maxLifespan() {
        return Optional.ofNullable(this.maxLifespan);
    }
    /**
     * @return Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     * 
     */
    public Optional<String> policy() {
        return Optional.ofNullable(this.policy);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(UserFederationCache defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Integer evictionDay;
        private @Nullable Integer evictionHour;
        private @Nullable Integer evictionMinute;
        private @Nullable String maxLifespan;
        private @Nullable String policy;
        public Builder() {}
        public Builder(UserFederationCache defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.evictionDay = defaults.evictionDay;
    	      this.evictionHour = defaults.evictionHour;
    	      this.evictionMinute = defaults.evictionMinute;
    	      this.maxLifespan = defaults.maxLifespan;
    	      this.policy = defaults.policy;
        }

        @CustomType.Setter
        public Builder evictionDay(@Nullable Integer evictionDay) {
            this.evictionDay = evictionDay;
            return this;
        }
        @CustomType.Setter
        public Builder evictionHour(@Nullable Integer evictionHour) {
            this.evictionHour = evictionHour;
            return this;
        }
        @CustomType.Setter
        public Builder evictionMinute(@Nullable Integer evictionMinute) {
            this.evictionMinute = evictionMinute;
            return this;
        }
        @CustomType.Setter
        public Builder maxLifespan(@Nullable String maxLifespan) {
            this.maxLifespan = maxLifespan;
            return this;
        }
        @CustomType.Setter
        public Builder policy(@Nullable String policy) {
            this.policy = policy;
            return this;
        }
        public UserFederationCache build() {
            final var _resultValue = new UserFederationCache();
            _resultValue.evictionDay = evictionDay;
            _resultValue.evictionHour = evictionHour;
            _resultValue.evictionMinute = evictionMinute;
            _resultValue.maxLifespan = maxLifespan;
            _resultValue.policy = policy;
            return _resultValue;
        }
    }
}
