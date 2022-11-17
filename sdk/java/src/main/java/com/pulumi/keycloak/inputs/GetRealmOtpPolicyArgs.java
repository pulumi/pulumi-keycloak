// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;


public final class GetRealmOtpPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetRealmOtpPolicyArgs Empty = new GetRealmOtpPolicyArgs();

    @Import(name="algorithm", required=true)
    private Output<String> algorithm;

    public Output<String> algorithm() {
        return this.algorithm;
    }

    @Import(name="digits", required=true)
    private Output<Integer> digits;

    public Output<Integer> digits() {
        return this.digits;
    }

    @Import(name="initialCounter", required=true)
    private Output<Integer> initialCounter;

    public Output<Integer> initialCounter() {
        return this.initialCounter;
    }

    @Import(name="lookAheadWindow", required=true)
    private Output<Integer> lookAheadWindow;

    public Output<Integer> lookAheadWindow() {
        return this.lookAheadWindow;
    }

    @Import(name="period", required=true)
    private Output<Integer> period;

    public Output<Integer> period() {
        return this.period;
    }

    @Import(name="type", required=true)
    private Output<String> type;

    public Output<String> type() {
        return this.type;
    }

    private GetRealmOtpPolicyArgs() {}

    private GetRealmOtpPolicyArgs(GetRealmOtpPolicyArgs $) {
        this.algorithm = $.algorithm;
        this.digits = $.digits;
        this.initialCounter = $.initialCounter;
        this.lookAheadWindow = $.lookAheadWindow;
        this.period = $.period;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRealmOtpPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRealmOtpPolicyArgs $;

        public Builder() {
            $ = new GetRealmOtpPolicyArgs();
        }

        public Builder(GetRealmOtpPolicyArgs defaults) {
            $ = new GetRealmOtpPolicyArgs(Objects.requireNonNull(defaults));
        }

        public Builder algorithm(Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        public Builder digits(Output<Integer> digits) {
            $.digits = digits;
            return this;
        }

        public Builder digits(Integer digits) {
            return digits(Output.of(digits));
        }

        public Builder initialCounter(Output<Integer> initialCounter) {
            $.initialCounter = initialCounter;
            return this;
        }

        public Builder initialCounter(Integer initialCounter) {
            return initialCounter(Output.of(initialCounter));
        }

        public Builder lookAheadWindow(Output<Integer> lookAheadWindow) {
            $.lookAheadWindow = lookAheadWindow;
            return this;
        }

        public Builder lookAheadWindow(Integer lookAheadWindow) {
            return lookAheadWindow(Output.of(lookAheadWindow));
        }

        public Builder period(Output<Integer> period) {
            $.period = period;
            return this;
        }

        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        public Builder type(String type) {
            return type(Output.of(type));
        }

        public GetRealmOtpPolicyArgs build() {
            $.algorithm = Objects.requireNonNull($.algorithm, "expected parameter 'algorithm' to be non-null");
            $.digits = Objects.requireNonNull($.digits, "expected parameter 'digits' to be non-null");
            $.initialCounter = Objects.requireNonNull($.initialCounter, "expected parameter 'initialCounter' to be non-null");
            $.lookAheadWindow = Objects.requireNonNull($.lookAheadWindow, "expected parameter 'lookAheadWindow' to be non-null");
            $.period = Objects.requireNonNull($.period, "expected parameter 'period' to be non-null");
            $.type = Objects.requireNonNull($.type, "expected parameter 'type' to be non-null");
            return $;
        }
    }

}