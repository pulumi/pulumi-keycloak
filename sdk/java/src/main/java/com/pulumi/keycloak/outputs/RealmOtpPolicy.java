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
     * @return What hashing algorithm should be used to generate the OTP, Valid options are `HmacSHA1`,`HmacSHA256` and `HmacSHA512`. Defaults to `HmacSHA1`.
     * 
     */
    private @Nullable String algorithm;
    /**
     * @return How many digits the OTP have. Defaults to `6`.
     * 
     */
    private @Nullable Integer digits;
    /**
     * @return What should the initial counter value be. Defaults to `2`.
     * 
     */
    private @Nullable Integer initialCounter;
    /**
     * @return How far ahead should the server look just in case the token generator and server are out of time sync or counter sync. Defaults to `1`.
     * 
     */
    private @Nullable Integer lookAheadWindow;
    /**
     * @return How many seconds should an OTP token be valid. Defaults to `30`.
     * 
     */
    private @Nullable Integer period;
    /**
     * @return One Time Password Type, supported Values are `totp` for Time-Based One Time Password and `hotp` for Counter Based. Defaults to `totp`.
     * 
     */
    private @Nullable String type;

    private RealmOtpPolicy() {}
    /**
     * @return What hashing algorithm should be used to generate the OTP, Valid options are `HmacSHA1`,`HmacSHA256` and `HmacSHA512`. Defaults to `HmacSHA1`.
     * 
     */
    public Optional<String> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }
    /**
     * @return How many digits the OTP have. Defaults to `6`.
     * 
     */
    public Optional<Integer> digits() {
        return Optional.ofNullable(this.digits);
    }
    /**
     * @return What should the initial counter value be. Defaults to `2`.
     * 
     */
    public Optional<Integer> initialCounter() {
        return Optional.ofNullable(this.initialCounter);
    }
    /**
     * @return How far ahead should the server look just in case the token generator and server are out of time sync or counter sync. Defaults to `1`.
     * 
     */
    public Optional<Integer> lookAheadWindow() {
        return Optional.ofNullable(this.lookAheadWindow);
    }
    /**
     * @return How many seconds should an OTP token be valid. Defaults to `30`.
     * 
     */
    public Optional<Integer> period() {
        return Optional.ofNullable(this.period);
    }
    /**
     * @return One Time Password Type, supported Values are `totp` for Time-Based One Time Password and `hotp` for Counter Based. Defaults to `totp`.
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
            final var o = new RealmOtpPolicy();
            o.algorithm = algorithm;
            o.digits = digits;
            o.initialCounter = initialCounter;
            o.lookAheadWindow = lookAheadWindow;
            o.period = period;
            o.type = type;
            return o;
        }
    }
}