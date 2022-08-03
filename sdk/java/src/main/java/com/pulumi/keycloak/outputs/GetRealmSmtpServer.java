// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.keycloak.outputs.GetRealmSmtpServerAuth;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetRealmSmtpServer {
    private final List<GetRealmSmtpServerAuth> auths;
    private final String envelopeFrom;
    private final String from;
    private final String fromDisplayName;
    private final String host;
    private final String port;
    private final String replyTo;
    private final String replyToDisplayName;
    private final Boolean ssl;
    private final Boolean starttls;

    @CustomType.Constructor
    private GetRealmSmtpServer(
        @CustomType.Parameter("auths") List<GetRealmSmtpServerAuth> auths,
        @CustomType.Parameter("envelopeFrom") String envelopeFrom,
        @CustomType.Parameter("from") String from,
        @CustomType.Parameter("fromDisplayName") String fromDisplayName,
        @CustomType.Parameter("host") String host,
        @CustomType.Parameter("port") String port,
        @CustomType.Parameter("replyTo") String replyTo,
        @CustomType.Parameter("replyToDisplayName") String replyToDisplayName,
        @CustomType.Parameter("ssl") Boolean ssl,
        @CustomType.Parameter("starttls") Boolean starttls) {
        this.auths = auths;
        this.envelopeFrom = envelopeFrom;
        this.from = from;
        this.fromDisplayName = fromDisplayName;
        this.host = host;
        this.port = port;
        this.replyTo = replyTo;
        this.replyToDisplayName = replyToDisplayName;
        this.ssl = ssl;
        this.starttls = starttls;
    }

    public List<GetRealmSmtpServerAuth> auths() {
        return this.auths;
    }
    public String envelopeFrom() {
        return this.envelopeFrom;
    }
    public String from() {
        return this.from;
    }
    public String fromDisplayName() {
        return this.fromDisplayName;
    }
    public String host() {
        return this.host;
    }
    public String port() {
        return this.port;
    }
    public String replyTo() {
        return this.replyTo;
    }
    public String replyToDisplayName() {
        return this.replyToDisplayName;
    }
    public Boolean ssl() {
        return this.ssl;
    }
    public Boolean starttls() {
        return this.starttls;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRealmSmtpServer defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private List<GetRealmSmtpServerAuth> auths;
        private String envelopeFrom;
        private String from;
        private String fromDisplayName;
        private String host;
        private String port;
        private String replyTo;
        private String replyToDisplayName;
        private Boolean ssl;
        private Boolean starttls;

        public Builder() {
    	      // Empty
        }

        public Builder(GetRealmSmtpServer defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.auths = defaults.auths;
    	      this.envelopeFrom = defaults.envelopeFrom;
    	      this.from = defaults.from;
    	      this.fromDisplayName = defaults.fromDisplayName;
    	      this.host = defaults.host;
    	      this.port = defaults.port;
    	      this.replyTo = defaults.replyTo;
    	      this.replyToDisplayName = defaults.replyToDisplayName;
    	      this.ssl = defaults.ssl;
    	      this.starttls = defaults.starttls;
        }

        public Builder auths(List<GetRealmSmtpServerAuth> auths) {
            this.auths = Objects.requireNonNull(auths);
            return this;
        }
        public Builder auths(GetRealmSmtpServerAuth... auths) {
            return auths(List.of(auths));
        }
        public Builder envelopeFrom(String envelopeFrom) {
            this.envelopeFrom = Objects.requireNonNull(envelopeFrom);
            return this;
        }
        public Builder from(String from) {
            this.from = Objects.requireNonNull(from);
            return this;
        }
        public Builder fromDisplayName(String fromDisplayName) {
            this.fromDisplayName = Objects.requireNonNull(fromDisplayName);
            return this;
        }
        public Builder host(String host) {
            this.host = Objects.requireNonNull(host);
            return this;
        }
        public Builder port(String port) {
            this.port = Objects.requireNonNull(port);
            return this;
        }
        public Builder replyTo(String replyTo) {
            this.replyTo = Objects.requireNonNull(replyTo);
            return this;
        }
        public Builder replyToDisplayName(String replyToDisplayName) {
            this.replyToDisplayName = Objects.requireNonNull(replyToDisplayName);
            return this;
        }
        public Builder ssl(Boolean ssl) {
            this.ssl = Objects.requireNonNull(ssl);
            return this;
        }
        public Builder starttls(Boolean starttls) {
            this.starttls = Objects.requireNonNull(starttls);
            return this;
        }        public GetRealmSmtpServer build() {
            return new GetRealmSmtpServer(auths, envelopeFrom, from, fromDisplayName, host, port, replyTo, replyToDisplayName, ssl, starttls);
        }
    }
}
