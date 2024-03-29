// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetRealmOtpPolicy {
    private String algorithm;
    private Integer digits;
    private Integer initialCounter;
    private Integer lookAheadWindow;
    private Integer period;
    private String type;

    private GetRealmOtpPolicy() {}
    public String algorithm() {
        return this.algorithm;
    }
    public Integer digits() {
        return this.digits;
    }
    public Integer initialCounter() {
        return this.initialCounter;
    }
    public Integer lookAheadWindow() {
        return this.lookAheadWindow;
    }
    public Integer period() {
        return this.period;
    }
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRealmOtpPolicy defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String algorithm;
        private Integer digits;
        private Integer initialCounter;
        private Integer lookAheadWindow;
        private Integer period;
        private String type;
        public Builder() {}
        public Builder(GetRealmOtpPolicy defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithm = defaults.algorithm;
    	      this.digits = defaults.digits;
    	      this.initialCounter = defaults.initialCounter;
    	      this.lookAheadWindow = defaults.lookAheadWindow;
    	      this.period = defaults.period;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder algorithm(String algorithm) {
            if (algorithm == null) {
              throw new MissingRequiredPropertyException("GetRealmOtpPolicy", "algorithm");
            }
            this.algorithm = algorithm;
            return this;
        }
        @CustomType.Setter
        public Builder digits(Integer digits) {
            if (digits == null) {
              throw new MissingRequiredPropertyException("GetRealmOtpPolicy", "digits");
            }
            this.digits = digits;
            return this;
        }
        @CustomType.Setter
        public Builder initialCounter(Integer initialCounter) {
            if (initialCounter == null) {
              throw new MissingRequiredPropertyException("GetRealmOtpPolicy", "initialCounter");
            }
            this.initialCounter = initialCounter;
            return this;
        }
        @CustomType.Setter
        public Builder lookAheadWindow(Integer lookAheadWindow) {
            if (lookAheadWindow == null) {
              throw new MissingRequiredPropertyException("GetRealmOtpPolicy", "lookAheadWindow");
            }
            this.lookAheadWindow = lookAheadWindow;
            return this;
        }
        @CustomType.Setter
        public Builder period(Integer period) {
            if (period == null) {
              throw new MissingRequiredPropertyException("GetRealmOtpPolicy", "period");
            }
            this.period = period;
            return this;
        }
        @CustomType.Setter
        public Builder type(String type) {
            if (type == null) {
              throw new MissingRequiredPropertyException("GetRealmOtpPolicy", "type");
            }
            this.type = type;
            return this;
        }
        public GetRealmOtpPolicy build() {
            final var _resultValue = new GetRealmOtpPolicy();
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
