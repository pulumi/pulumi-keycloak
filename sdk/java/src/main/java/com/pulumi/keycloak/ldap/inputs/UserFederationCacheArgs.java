// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserFederationCacheArgs extends com.pulumi.resources.ResourceArgs {

    public static final UserFederationCacheArgs Empty = new UserFederationCacheArgs();

    /**
     * Day of the week the entry will become invalid on
     * 
     */
    @Import(name="evictionDay")
    private @Nullable Output<Integer> evictionDay;

    /**
     * @return Day of the week the entry will become invalid on
     * 
     */
    public Optional<Output<Integer>> evictionDay() {
        return Optional.ofNullable(this.evictionDay);
    }

    /**
     * Hour of day the entry will become invalid on.
     * 
     */
    @Import(name="evictionHour")
    private @Nullable Output<Integer> evictionHour;

    /**
     * @return Hour of day the entry will become invalid on.
     * 
     */
    public Optional<Output<Integer>> evictionHour() {
        return Optional.ofNullable(this.evictionHour);
    }

    /**
     * Minute of day the entry will become invalid on.
     * 
     */
    @Import(name="evictionMinute")
    private @Nullable Output<Integer> evictionMinute;

    /**
     * @return Minute of day the entry will become invalid on.
     * 
     */
    public Optional<Output<Integer>> evictionMinute() {
        return Optional.ofNullable(this.evictionMinute);
    }

    /**
     * Max lifespan of cache entry (duration string).
     * 
     */
    @Import(name="maxLifespan")
    private @Nullable Output<String> maxLifespan;

    /**
     * @return Max lifespan of cache entry (duration string).
     * 
     */
    public Optional<Output<String>> maxLifespan() {
        return Optional.ofNullable(this.maxLifespan);
    }

    /**
     * Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     * 
     */
    @Import(name="policy")
    private @Nullable Output<String> policy;

    /**
     * @return Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
     * 
     */
    public Optional<Output<String>> policy() {
        return Optional.ofNullable(this.policy);
    }

    private UserFederationCacheArgs() {}

    private UserFederationCacheArgs(UserFederationCacheArgs $) {
        this.evictionDay = $.evictionDay;
        this.evictionHour = $.evictionHour;
        this.evictionMinute = $.evictionMinute;
        this.maxLifespan = $.maxLifespan;
        this.policy = $.policy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserFederationCacheArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserFederationCacheArgs $;

        public Builder() {
            $ = new UserFederationCacheArgs();
        }

        public Builder(UserFederationCacheArgs defaults) {
            $ = new UserFederationCacheArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param evictionDay Day of the week the entry will become invalid on
         * 
         * @return builder
         * 
         */
        public Builder evictionDay(@Nullable Output<Integer> evictionDay) {
            $.evictionDay = evictionDay;
            return this;
        }

        /**
         * @param evictionDay Day of the week the entry will become invalid on
         * 
         * @return builder
         * 
         */
        public Builder evictionDay(Integer evictionDay) {
            return evictionDay(Output.of(evictionDay));
        }

        /**
         * @param evictionHour Hour of day the entry will become invalid on.
         * 
         * @return builder
         * 
         */
        public Builder evictionHour(@Nullable Output<Integer> evictionHour) {
            $.evictionHour = evictionHour;
            return this;
        }

        /**
         * @param evictionHour Hour of day the entry will become invalid on.
         * 
         * @return builder
         * 
         */
        public Builder evictionHour(Integer evictionHour) {
            return evictionHour(Output.of(evictionHour));
        }

        /**
         * @param evictionMinute Minute of day the entry will become invalid on.
         * 
         * @return builder
         * 
         */
        public Builder evictionMinute(@Nullable Output<Integer> evictionMinute) {
            $.evictionMinute = evictionMinute;
            return this;
        }

        /**
         * @param evictionMinute Minute of day the entry will become invalid on.
         * 
         * @return builder
         * 
         */
        public Builder evictionMinute(Integer evictionMinute) {
            return evictionMinute(Output.of(evictionMinute));
        }

        /**
         * @param maxLifespan Max lifespan of cache entry (duration string).
         * 
         * @return builder
         * 
         */
        public Builder maxLifespan(@Nullable Output<String> maxLifespan) {
            $.maxLifespan = maxLifespan;
            return this;
        }

        /**
         * @param maxLifespan Max lifespan of cache entry (duration string).
         * 
         * @return builder
         * 
         */
        public Builder maxLifespan(String maxLifespan) {
            return maxLifespan(Output.of(maxLifespan));
        }

        /**
         * @param policy Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
         * 
         * @return builder
         * 
         */
        public Builder policy(@Nullable Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        public UserFederationCacheArgs build() {
            return $;
        }
    }

}
