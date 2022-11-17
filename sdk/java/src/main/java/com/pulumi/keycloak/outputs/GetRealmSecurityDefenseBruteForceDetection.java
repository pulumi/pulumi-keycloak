// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.util.Objects;

@CustomType
public final class GetRealmSecurityDefenseBruteForceDetection {
    private Integer failureResetTimeSeconds;
    private Integer maxFailureWaitSeconds;
    private Integer maxLoginFailures;
    private Integer minimumQuickLoginWaitSeconds;
    private Boolean permanentLockout;
    private Integer quickLoginCheckMilliSeconds;
    private Integer waitIncrementSeconds;

    private GetRealmSecurityDefenseBruteForceDetection() {}
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
    @CustomType.Builder
    public static final class Builder {
        private Integer failureResetTimeSeconds;
        private Integer maxFailureWaitSeconds;
        private Integer maxLoginFailures;
        private Integer minimumQuickLoginWaitSeconds;
        private Boolean permanentLockout;
        private Integer quickLoginCheckMilliSeconds;
        private Integer waitIncrementSeconds;
        public Builder() {}
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

        @CustomType.Setter
        public Builder failureResetTimeSeconds(Integer failureResetTimeSeconds) {
            this.failureResetTimeSeconds = Objects.requireNonNull(failureResetTimeSeconds);
            return this;
        }
        @CustomType.Setter
        public Builder maxFailureWaitSeconds(Integer maxFailureWaitSeconds) {
            this.maxFailureWaitSeconds = Objects.requireNonNull(maxFailureWaitSeconds);
            return this;
        }
        @CustomType.Setter
        public Builder maxLoginFailures(Integer maxLoginFailures) {
            this.maxLoginFailures = Objects.requireNonNull(maxLoginFailures);
            return this;
        }
        @CustomType.Setter
        public Builder minimumQuickLoginWaitSeconds(Integer minimumQuickLoginWaitSeconds) {
            this.minimumQuickLoginWaitSeconds = Objects.requireNonNull(minimumQuickLoginWaitSeconds);
            return this;
        }
        @CustomType.Setter
        public Builder permanentLockout(Boolean permanentLockout) {
            this.permanentLockout = Objects.requireNonNull(permanentLockout);
            return this;
        }
        @CustomType.Setter
        public Builder quickLoginCheckMilliSeconds(Integer quickLoginCheckMilliSeconds) {
            this.quickLoginCheckMilliSeconds = Objects.requireNonNull(quickLoginCheckMilliSeconds);
            return this;
        }
        @CustomType.Setter
        public Builder waitIncrementSeconds(Integer waitIncrementSeconds) {
            this.waitIncrementSeconds = Objects.requireNonNull(waitIncrementSeconds);
            return this;
        }
        public GetRealmSecurityDefenseBruteForceDetection build() {
            final var o = new GetRealmSecurityDefenseBruteForceDetection();
            o.failureResetTimeSeconds = failureResetTimeSeconds;
            o.maxFailureWaitSeconds = maxFailureWaitSeconds;
            o.maxLoginFailures = maxLoginFailures;
            o.minimumQuickLoginWaitSeconds = minimumQuickLoginWaitSeconds;
            o.permanentLockout = permanentLockout;
            o.quickLoginCheckMilliSeconds = quickLoginCheckMilliSeconds;
            o.waitIncrementSeconds = waitIncrementSeconds;
            return o;
        }
    }
}