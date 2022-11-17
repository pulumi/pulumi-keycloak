// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.keycloak.outputs.GetRealmSecurityDefenseBruteForceDetection;
import com.pulumi.keycloak.outputs.GetRealmSecurityDefenseHeader;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRealmSecurityDefense {
    private List<GetRealmSecurityDefenseBruteForceDetection> bruteForceDetections;
    private List<GetRealmSecurityDefenseHeader> headers;

    private GetRealmSecurityDefense() {}
    public List<GetRealmSecurityDefenseBruteForceDetection> bruteForceDetections() {
        return this.bruteForceDetections;
    }
    public List<GetRealmSecurityDefenseHeader> headers() {
        return this.headers;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRealmSecurityDefense defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<GetRealmSecurityDefenseBruteForceDetection> bruteForceDetections;
        private List<GetRealmSecurityDefenseHeader> headers;
        public Builder() {}
        public Builder(GetRealmSecurityDefense defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bruteForceDetections = defaults.bruteForceDetections;
    	      this.headers = defaults.headers;
        }

        @CustomType.Setter
        public Builder bruteForceDetections(List<GetRealmSecurityDefenseBruteForceDetection> bruteForceDetections) {
            this.bruteForceDetections = Objects.requireNonNull(bruteForceDetections);
            return this;
        }
        public Builder bruteForceDetections(GetRealmSecurityDefenseBruteForceDetection... bruteForceDetections) {
            return bruteForceDetections(List.of(bruteForceDetections));
        }
        @CustomType.Setter
        public Builder headers(List<GetRealmSecurityDefenseHeader> headers) {
            this.headers = Objects.requireNonNull(headers);
            return this;
        }
        public Builder headers(GetRealmSecurityDefenseHeader... headers) {
            return headers(List.of(headers));
        }
        public GetRealmSecurityDefense build() {
            final var o = new GetRealmSecurityDefense();
            o.bruteForceDetections = bruteForceDetections;
            o.headers = headers;
            return o;
        }
    }
}