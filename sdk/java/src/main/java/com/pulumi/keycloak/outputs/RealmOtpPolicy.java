// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RealmOtpPolicy {
    /**
     * @return What hashing algorithm should be used to generate the OTP.
     * 
     */
    private @Nullable String algorithm;
    private @Nullable Integer digits;
    private @Nullable Integer initialCounter;
    private @Nullable Integer lookAheadWindow;
    private @Nullable Integer period;
    /**
     * @return OTP Type, totp for Time-Based One Time Password or hotp for counter base one time password
     * 
     */
    private @Nullable String type;

    private RealmOtpPolicy() {}
    /**
     * @return What hashing algorithm should be used to generate the OTP.
     * 
     */
    public Optional<String> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }
    public Optional<Integer> digits() {
        return Optional.ofNullable(this.digits);
    }
    public Optional<Integer> initialCounter() {
        return Optional.ofNullable(this.initialCounter);
    }
    public Optional<Integer> lookAheadWindow() {
        return Optional.ofNullable(this.lookAheadWindow);
    }
    public Optional<Integer> period() {
        return Optional.ofNullable(this.period);
    }
    /**
     * @return OTP Type, totp for Time-Based One Time Password or hotp for counter base one time password
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RealmOtpPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String algorithm;
        private @Nullable Integer digits;
        private @Nullable Integer initialCounter;
        private @Nullable Integer lookAheadWindow;
        private @Nullable Integer period;
        private @Nullable String type;
        public Builder() {}
        public Builder(RealmOtpPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithm = defaults.algorithm;
    	      this.digits = defaults.digits;
    	      this.initialCounter = defaults.initialCounter;
    	      this.lookAheadWindow = defaults.lookAheadWindow;
    	      this.period = defaults.period;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder algorithm(@Nullable String algorithm) {

            this.algorithm = algorithm;
            return this;
        }
        @CustomType.Setter
        public Builder digits(@Nullable Integer digits) {

            this.digits = digits;
            return this;
        }
        @CustomType.Setter
        public Builder initialCounter(@Nullable Integer initialCounter) {

            this.initialCounter = initialCounter;
            return this;
        }
        @CustomType.Setter
        public Builder lookAheadWindow(@Nullable Integer lookAheadWindow) {

            this.lookAheadWindow = lookAheadWindow;
            return this;
        }
        @CustomType.Setter
        public Builder period(@Nullable Integer period) {

            this.period = period;
            return this;
        }
        @CustomType.Setter
        public Builder type(@Nullable String type) {

            this.type = type;
            return this;
        }
        public RealmOtpPolicy build() {
            final var _resultValue = new RealmOtpPolicy();
            _resultValue.algorithm = algorithm;
            _resultValue.digits = digits;
            _resultValue.initialCounter = initialCounter;
            _resultValue.lookAheadWindow = lookAheadWindow;
            _resultValue.period = period;
            _resultValue.type = type;
            return _resultValue;
        }
    }
}
