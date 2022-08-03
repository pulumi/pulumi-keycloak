// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RealmSecurityDefensesBruteForceDetection {
    /**
     * @return When will failure count be reset?
     * 
     */
    private final @Nullable Integer failureResetTimeSeconds;
    private final @Nullable Integer maxFailureWaitSeconds;
    /**
     * @return How many failures before wait is triggered.
     * 
     */
    private final @Nullable Integer maxLoginFailures;
    /**
     * @return How long to wait after a quick login failure.
     * - ` max_failure_wait_seconds  ` - (Optional) Max. time a user will be locked out.
     * 
     */
    private final @Nullable Integer minimumQuickLoginWaitSeconds;
    /**
     * @return When `true`, this will lock the user permanently when the user exceeds the maximum login failures.
     * 
     */
    private final @Nullable Boolean permanentLockout;
    /**
     * @return Configures the amount of time, in milliseconds, for consecutive failures to lock a user out.
     * 
     */
    private final @Nullable Integer quickLoginCheckMilliSeconds;
    /**
     * @return This represents the amount of time a user should be locked out when the login failure threshold has been met.
     * 
     */
    private final @Nullable Integer waitIncrementSeconds;

    @CustomType.Constructor
    private RealmSecurityDefensesBruteForceDetection(
        @CustomType.Parameter("failureResetTimeSeconds") @Nullable Integer failureResetTimeSeconds,
        @CustomType.Parameter("maxFailureWaitSeconds") @Nullable Integer maxFailureWaitSeconds,
        @CustomType.Parameter("maxLoginFailures") @Nullable Integer maxLoginFailures,
        @CustomType.Parameter("minimumQuickLoginWaitSeconds") @Nullable Integer minimumQuickLoginWaitSeconds,
        @CustomType.Parameter("permanentLockout") @Nullable Boolean permanentLockout,
        @CustomType.Parameter("quickLoginCheckMilliSeconds") @Nullable Integer quickLoginCheckMilliSeconds,
        @CustomType.Parameter("waitIncrementSeconds") @Nullable Integer waitIncrementSeconds) {
        this.failureResetTimeSeconds = failureResetTimeSeconds;
        this.maxFailureWaitSeconds = maxFailureWaitSeconds;
        this.maxLoginFailures = maxLoginFailures;
        this.minimumQuickLoginWaitSeconds = minimumQuickLoginWaitSeconds;
        this.permanentLockout = permanentLockout;
        this.quickLoginCheckMilliSeconds = quickLoginCheckMilliSeconds;
        this.waitIncrementSeconds = waitIncrementSeconds;
    }

    /**
     * @return When will failure count be reset?
     * 
     */
    public Optional<Integer> failureResetTimeSeconds() {
        return Optional.ofNullable(this.failureResetTimeSeconds);
    }
    public Optional<Integer> maxFailureWaitSeconds() {
        return Optional.ofNullable(this.maxFailureWaitSeconds);
    }
    /**
     * @return How many failures before wait is triggered.
     * 
     */
    public Optional<Integer> maxLoginFailures() {
        return Optional.ofNullable(this.maxLoginFailures);
    }
    /**
     * @return How long to wait after a quick login failure.
     * - ` max_failure_wait_seconds  ` - (Optional) Max. time a user will be locked out.
     * 
     */
    public Optional<Integer> minimumQuickLoginWaitSeconds() {
        return Optional.ofNullable(this.minimumQuickLoginWaitSeconds);
    }
    /**
     * @return When `true`, this will lock the user permanently when the user exceeds the maximum login failures.
     * 
     */
    public Optional<Boolean> permanentLockout() {
        return Optional.ofNullable(this.permanentLockout);
    }
    /**
     * @return Configures the amount of time, in milliseconds, for consecutive failures to lock a user out.
     * 
     */
    public Optional<Integer> quickLoginCheckMilliSeconds() {
        return Optional.ofNullable(this.quickLoginCheckMilliSeconds);
    }
    /**
     * @return This represents the amount of time a user should be locked out when the login failure threshold has been met.
     * 
     */
    public Optional<Integer> waitIncrementSeconds() {
        return Optional.ofNullable(this.waitIncrementSeconds);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RealmSecurityDefensesBruteForceDetection defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Integer failureResetTimeSeconds;
        private @Nullable Integer maxFailureWaitSeconds;
        private @Nullable Integer maxLoginFailures;
        private @Nullable Integer minimumQuickLoginWaitSeconds;
        private @Nullable Boolean permanentLockout;
        private @Nullable Integer quickLoginCheckMilliSeconds;
        private @Nullable Integer waitIncrementSeconds;

        public Builder() {
    	      // Empty
        }

        public Builder(RealmSecurityDefensesBruteForceDetection defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.failureResetTimeSeconds = defaults.failureResetTimeSeconds;
    	      this.maxFailureWaitSeconds = defaults.maxFailureWaitSeconds;
    	      this.maxLoginFailures = defaults.maxLoginFailures;
    	      this.minimumQuickLoginWaitSeconds = defaults.minimumQuickLoginWaitSeconds;
    	      this.permanentLockout = defaults.permanentLockout;
    	      this.quickLoginCheckMilliSeconds = defaults.quickLoginCheckMilliSeconds;
    	      this.waitIncrementSeconds = defaults.waitIncrementSeconds;
        }

        public Builder failureResetTimeSeconds(@Nullable Integer failureResetTimeSeconds) {
            this.failureResetTimeSeconds = failureResetTimeSeconds;
            return this;
        }
        public Builder maxFailureWaitSeconds(@Nullable Integer maxFailureWaitSeconds) {
            this.maxFailureWaitSeconds = maxFailureWaitSeconds;
            return this;
        }
        public Builder maxLoginFailures(@Nullable Integer maxLoginFailures) {
            this.maxLoginFailures = maxLoginFailures;
            return this;
        }
        public Builder minimumQuickLoginWaitSeconds(@Nullable Integer minimumQuickLoginWaitSeconds) {
            this.minimumQuickLoginWaitSeconds = minimumQuickLoginWaitSeconds;
            return this;
        }
        public Builder permanentLockout(@Nullable Boolean permanentLockout) {
            this.permanentLockout = permanentLockout;
            return this;
        }
        public Builder quickLoginCheckMilliSeconds(@Nullable Integer quickLoginCheckMilliSeconds) {
            this.quickLoginCheckMilliSeconds = quickLoginCheckMilliSeconds;
            return this;
        }
        public Builder waitIncrementSeconds(@Nullable Integer waitIncrementSeconds) {
            this.waitIncrementSeconds = waitIncrementSeconds;
            return this;
        }        public RealmSecurityDefensesBruteForceDetection build() {
            return new RealmSecurityDefensesBruteForceDetection(failureResetTimeSeconds, maxFailureWaitSeconds, maxLoginFailures, minimumQuickLoginWaitSeconds, permanentLockout, quickLoginCheckMilliSeconds, waitIncrementSeconds);
        }
    }
}
