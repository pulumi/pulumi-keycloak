// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetRealmSecurityDefenseHeader extends com.pulumi.resources.InvokeArgs {

    public static final GetRealmSecurityDefenseHeader Empty = new GetRealmSecurityDefenseHeader();

    @Import(name="contentSecurityPolicy", required=true)
    private String contentSecurityPolicy;

    public String contentSecurityPolicy() {
        return this.contentSecurityPolicy;
    }

    @Import(name="contentSecurityPolicyReportOnly", required=true)
    private String contentSecurityPolicyReportOnly;

    public String contentSecurityPolicyReportOnly() {
        return this.contentSecurityPolicyReportOnly;
    }

    @Import(name="strictTransportSecurity", required=true)
    private String strictTransportSecurity;

    public String strictTransportSecurity() {
        return this.strictTransportSecurity;
    }

    @Import(name="xContentTypeOptions", required=true)
    private String xContentTypeOptions;

    public String xContentTypeOptions() {
        return this.xContentTypeOptions;
    }

    @Import(name="xFrameOptions", required=true)
    private String xFrameOptions;

    public String xFrameOptions() {
        return this.xFrameOptions;
    }

    @Import(name="xRobotsTag", required=true)
    private String xRobotsTag;

    public String xRobotsTag() {
        return this.xRobotsTag;
    }

    @Import(name="xXssProtection", required=true)
    private String xXssProtection;

    public String xXssProtection() {
        return this.xXssProtection;
    }

    private GetRealmSecurityDefenseHeader() {}

    private GetRealmSecurityDefenseHeader(GetRealmSecurityDefenseHeader $) {
        this.contentSecurityPolicy = $.contentSecurityPolicy;
        this.contentSecurityPolicyReportOnly = $.contentSecurityPolicyReportOnly;
        this.strictTransportSecurity = $.strictTransportSecurity;
        this.xContentTypeOptions = $.xContentTypeOptions;
        this.xFrameOptions = $.xFrameOptions;
        this.xRobotsTag = $.xRobotsTag;
        this.xXssProtection = $.xXssProtection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRealmSecurityDefenseHeader defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRealmSecurityDefenseHeader $;

        public Builder() {
            $ = new GetRealmSecurityDefenseHeader();
        }

        public Builder(GetRealmSecurityDefenseHeader defaults) {
            $ = new GetRealmSecurityDefenseHeader(Objects.requireNonNull(defaults));
        }

        public Builder contentSecurityPolicy(String contentSecurityPolicy) {
            $.contentSecurityPolicy = contentSecurityPolicy;
            return this;
        }

        public Builder contentSecurityPolicyReportOnly(String contentSecurityPolicyReportOnly) {
            $.contentSecurityPolicyReportOnly = contentSecurityPolicyReportOnly;
            return this;
        }

        public Builder strictTransportSecurity(String strictTransportSecurity) {
            $.strictTransportSecurity = strictTransportSecurity;
            return this;
        }

        public Builder xContentTypeOptions(String xContentTypeOptions) {
            $.xContentTypeOptions = xContentTypeOptions;
            return this;
        }

        public Builder xFrameOptions(String xFrameOptions) {
            $.xFrameOptions = xFrameOptions;
            return this;
        }

        public Builder xRobotsTag(String xRobotsTag) {
            $.xRobotsTag = xRobotsTag;
            return this;
        }

        public Builder xXssProtection(String xXssProtection) {
            $.xXssProtection = xXssProtection;
            return this;
        }

        public GetRealmSecurityDefenseHeader build() {
            $.contentSecurityPolicy = Objects.requireNonNull($.contentSecurityPolicy, "expected parameter 'contentSecurityPolicy' to be non-null");
            $.contentSecurityPolicyReportOnly = Objects.requireNonNull($.contentSecurityPolicyReportOnly, "expected parameter 'contentSecurityPolicyReportOnly' to be non-null");
            $.strictTransportSecurity = Objects.requireNonNull($.strictTransportSecurity, "expected parameter 'strictTransportSecurity' to be non-null");
            $.xContentTypeOptions = Objects.requireNonNull($.xContentTypeOptions, "expected parameter 'xContentTypeOptions' to be non-null");
            $.xFrameOptions = Objects.requireNonNull($.xFrameOptions, "expected parameter 'xFrameOptions' to be non-null");
            $.xRobotsTag = Objects.requireNonNull($.xRobotsTag, "expected parameter 'xRobotsTag' to be non-null");
            $.xXssProtection = Objects.requireNonNull($.xXssProtection, "expected parameter 'xXssProtection' to be non-null");
            return $;
        }
    }

}