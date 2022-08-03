// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class GetRealmSecurityDefenseBruteForceDetection {
    private final Integer failureResetTimeSeconds;
    private final Integer maxFailureWaitSeconds;
    private final Integer maxLoginFailures;
    private final Integer minimumQuickLoginWaitSeconds;
    private final Boolean permanentLockout;
    private final Integer quickLoginCheckMilliSeconds;
    private final Integer waitIncrementSeconds;

    @CustomType.Constructor
    private GetRealmSecurityDefenseBruteForceDetection(
        @CustomType.Parameter("failureResetTimeSeconds") Integer failureResetTimeSeconds,
        @CustomType.Parameter("maxFailureWaitSeconds") Integer maxFailureWaitSeconds,
        @CustomType.Parameter("maxLoginFailures") Integer maxLoginFailures,
        @CustomType.Parameter("minimumQuickLoginWaitSeconds") Integer minimumQuickLoginWaitSeconds,
        @CustomType.Parameter("permanentLockout") Boolean permanentLockout,
        @CustomType.Parameter("quickLoginCheckMilliSeconds") Integer quickLoginCheckMilliSeconds,
        @CustomType.Parameter("waitIncrementSeconds") Integer waitIncrementSeconds) {
        this.failureResetTimeSeconds = failureResetTimeSeconds;
        this.maxFailureWaitSeconds = maxFailureWaitSeconds;
        this.maxLoginFailures = maxLoginFailures;
        this.minimumQuickLoginWaitSeconds = minimumQuickLoginWaitSeconds;
        this.permanentLockout = permanentLockout;
        this.quickLoginCheckMilliSeconds = quickLoginCheckMilliSeconds;
        this.waitIncrementSeconds = waitIncrementSeconds;
    }

    public Integer failureResetTimeSeconds() {
        return this.failureResetTimeSeconds;
    }
    public Integer maxFailureWaitSeconds() {
        return this.maxFailureWaitSeconds;
    }
    public Integer maxLoginFailures() {
        return this.maxLoginFailures;
    }
    public Integer minimumQuickLoginWaitSeconds() {
        return this.minimumQuickLoginWaitSeconds;
    }
    public Boolean permanentLockout() {
        return this.permanentLockout;
    }
    public Integer quickLoginCheckMilliSeconds() {
        return this.quickLoginCheckMilliSeconds;
    }
    public Integer waitIncrementSeconds() {
        return this.waitIncrementSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRealmSecurityDefenseBruteForceDetection defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private Integer failureResetTimeSeconds;
        private Integer maxFailureWaitSeconds;
        private Integer maxLoginFailures;
        private Integer minimumQuickLoginWaitSeconds;
        private Boolean permanentLockout;
        private Integer quickLoginCheckMilliSeconds;
        private Integer waitIncrementSeconds;

        public Builder() {
    	      // Empty
        }

        public Builder(GetRealmSecurityDefenseBruteForceDetection defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.failureResetTimeSeconds = defaults.failureResetTimeSeconds;
    	      this.maxFailureWaitSeconds = defaults.maxFailureWaitSeconds;
    	      this.maxLoginFailures = defaults.maxLoginFailures;
    	      this.minimumQuickLoginWaitSeconds = defaults.minimumQuickLoginWaitSeconds;
    	      this.permanentLockout = defaults.permanentLockout;
    	      this.quickLoginCheckMilliSeconds = defaults.quickLoginCheckMilliSeconds;
    	      this.waitIncrementSeconds = defaults.waitIncrementSeconds;
        }

        public Builder failureResetTimeSeconds(Integer failureResetTimeSeconds) {
            this.failureResetTimeSeconds = Objects.requireNonNull(failureResetTimeSeconds);
            return this;
        }
        public Builder maxFailureWaitSeconds(Integer maxFailureWaitSeconds) {
            this.maxFailureWaitSeconds = Objects.requireNonNull(maxFailureWaitSeconds);
            return this;
        }
        public Builder maxLoginFailures(Integer maxLoginFailures) {
            this.maxLoginFailures = Objects.requireNonNull(maxLoginFailures);
            return this;
        }
        public Builder minimumQuickLoginWaitSeconds(Integer minimumQuickLoginWaitSeconds) {
            this.minimumQuickLoginWaitSeconds = Objects.requireNonNull(minimumQuickLoginWaitSeconds);
            return this;
        }
        public Builder permanentLockout(Boolean permanentLockout) {
            this.permanentLockout = Objects.requireNonNull(permanentLockout);
            return this;
        }
        public Builder quickLoginCheckMilliSeconds(Integer quickLoginCheckMilliSeconds) {
            this.quickLoginCheckMilliSeconds = Objects.requireNonNull(quickLoginCheckMilliSeconds);
            return this;
        }
        public Builder waitIncrementSeconds(Integer waitIncrementSeconds) {
            this.waitIncrementSeconds = Objects.requireNonNull(waitIncrementSeconds);
            return this;
        }        public GetRealmSecurityDefenseBruteForceDetection build() {
            return new GetRealmSecurityDefenseBruteForceDetection(failureResetTimeSeconds, maxFailureWaitSeconds, maxLoginFailures, minimumQuickLoginWaitSeconds, permanentLockout, quickLoginCheckMilliSeconds, waitIncrementSeconds);
        }
    }
}
