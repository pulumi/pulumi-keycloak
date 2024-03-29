// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.keycloak.inputs.RealmSecurityDefensesBruteForceDetectionArgs;
import com.pulumi.keycloak.inputs.RealmSecurityDefensesHeadersArgs;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmSecurityDefensesArgs extends com.pulumi.resources.ResourceArgs {

    public static final RealmSecurityDefensesArgs Empty = new RealmSecurityDefensesArgs();

    @Import(name="bruteForceDetection")
    private @Nullable Output<RealmSecurityDefensesBruteForceDetectionArgs> bruteForceDetection;

    public Optional<Output<RealmSecurityDefensesBruteForceDetectionArgs>> bruteForceDetection() {
        return Optional.ofNullable(this.bruteForceDetection);
    }

    @Import(name="headers")
    private @Nullable Output<RealmSecurityDefensesHeadersArgs> headers;

    public Optional<Output<RealmSecurityDefensesHeadersArgs>> headers() {
        return Optional.ofNullable(this.headers);
    }

    private RealmSecurityDefensesArgs() {}

    private RealmSecurityDefensesArgs(RealmSecurityDefensesArgs $) {
        this.bruteForceDetection = $.bruteForceDetection;
        this.headers = $.headers;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmSecurityDefensesArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmSecurityDefensesArgs $;

        public Builder() {
            $ = new RealmSecurityDefensesArgs();
        }

        public Builder(RealmSecurityDefensesArgs defaults) {
            $ = new RealmSecurityDefensesArgs(Objects.requireNonNull(defaults));
        }

        public Builder bruteForceDetection(@Nullable Output<RealmSecurityDefensesBruteForceDetectionArgs> bruteForceDetection) {
            $.bruteForceDetection = bruteForceDetection;
            return this;
        }

        public Builder bruteForceDetection(RealmSecurityDefensesBruteForceDetectionArgs bruteForceDetection) {
            return bruteForceDetection(Output.of(bruteForceDetection));
        }

        public Builder headers(@Nullable Output<RealmSecurityDefensesHeadersArgs> headers) {
            $.headers = headers;
            return this;
        }

        public Builder headers(RealmSecurityDefensesHeadersArgs headers) {
            return headers(Output.of(headers));
        }

        public RealmSecurityDefensesArgs build() {
            return $;
        }
    }

}
