// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.inputs.GetRealmSmtpServerAuthArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class GetRealmSmtpServerArgs extends com.pulumi.resources.ResourceArgs {

    public static final GetRealmSmtpServerArgs Empty = new GetRealmSmtpServerArgs();

    @Import(name="auths", required=true)
    private Output<List<GetRealmSmtpServerAuthArgs>> auths;

    public Output<List<GetRealmSmtpServerAuthArgs>> auths() {
        return this.auths;
    }

    @Import(name="envelopeFrom", required=true)
    private Output<String> envelopeFrom;

    public Output<String> envelopeFrom() {
        return this.envelopeFrom;
    }

    @Import(name="from", required=true)
    private Output<String> from;

    public Output<String> from() {
        return this.from;
    }

    @Import(name="fromDisplayName", required=true)
    private Output<String> fromDisplayName;

    public Output<String> fromDisplayName() {
        return this.fromDisplayName;
    }

    @Import(name="host", required=true)
    private Output<String> host;

    public Output<String> host() {
        return this.host;
    }

    @Import(name="port", required=true)
    private Output<String> port;

    public Output<String> port() {
        return this.port;
    }

    @Import(name="replyTo", required=true)
    private Output<String> replyTo;

    public Output<String> replyTo() {
        return this.replyTo;
    }

    @Import(name="replyToDisplayName", required=true)
    private Output<String> replyToDisplayName;

    public Output<String> replyToDisplayName() {
        return this.replyToDisplayName;
    }

    @Import(name="ssl", required=true)
    private Output<Boolean> ssl;

    public Output<Boolean> ssl() {
        return this.ssl;
    }

    @Import(name="starttls", required=true)
    private Output<Boolean> starttls;

    public Output<Boolean> starttls() {
        return this.starttls;
    }

    private GetRealmSmtpServerArgs() {}

    private GetRealmSmtpServerArgs(GetRealmSmtpServerArgs $) {
        this.auths = $.auths;
        this.envelopeFrom = $.envelopeFrom;
        this.from = $.from;
        this.fromDisplayName = $.fromDisplayName;
        this.host = $.host;
        this.port = $.port;
        this.replyTo = $.replyTo;
        this.replyToDisplayName = $.replyToDisplayName;
        this.ssl = $.ssl;
        this.starttls = $.starttls;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetRealmSmtpServerArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetRealmSmtpServerArgs $;

        public Builder() {
            $ = new GetRealmSmtpServerArgs();
        }

        public Builder(GetRealmSmtpServerArgs defaults) {
            $ = new GetRealmSmtpServerArgs(Objects.requireNonNull(defaults));
        }

        public Builder auths(Output<List<GetRealmSmtpServerAuthArgs>> auths) {
            $.auths = auths;
            return this;
        }

        public Builder auths(List<GetRealmSmtpServerAuthArgs> auths) {
            return auths(Output.of(auths));
        }

        public Builder auths(GetRealmSmtpServerAuthArgs... auths) {
            return auths(List.of(auths));
        }

        public Builder envelopeFrom(Output<String> envelopeFrom) {
            $.envelopeFrom = envelopeFrom;
            return this;
        }

        public Builder envelopeFrom(String envelopeFrom) {
            return envelopeFrom(Output.of(envelopeFrom));
        }

        public Builder from(Output<String> from) {
            $.from = from;
            return this;
        }

        public Builder from(String from) {
            return from(Output.of(from));
        }

        public Builder fromDisplayName(Output<String> fromDisplayName) {
            $.fromDisplayName = fromDisplayName;
            return this;
        }

        public Builder fromDisplayName(String fromDisplayName) {
            return fromDisplayName(Output.of(fromDisplayName));
        }

        public Builder host(Output<String> host) {
            $.host = host;
            return this;
        }

        public Builder host(String host) {
            return host(Output.of(host));
        }

        public Builder port(Output<String> port) {
            $.port = port;
            return this;
        }

        public Builder port(String port) {
            return port(Output.of(port));
        }

        public Builder replyTo(Output<String> replyTo) {
            $.replyTo = replyTo;
            return this;
        }

        public Builder replyTo(String replyTo) {
            return replyTo(Output.of(replyTo));
        }

        public Builder replyToDisplayName(Output<String> replyToDisplayName) {
            $.replyToDisplayName = replyToDisplayName;
            return this;
        }

        public Builder replyToDisplayName(String replyToDisplayName) {
            return replyToDisplayName(Output.of(replyToDisplayName));
        }

        public Builder ssl(Output<Boolean> ssl) {
            $.ssl = ssl;
            return this;
        }

        public Builder ssl(Boolean ssl) {
            return ssl(Output.of(ssl));
        }

        public Builder starttls(Output<Boolean> starttls) {
            $.starttls = starttls;
            return this;
        }

        public Builder starttls(Boolean starttls) {
            return starttls(Output.of(starttls));
        }

        public GetRealmSmtpServerArgs build() {
            if ($.auths == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "auths");
            }
            if ($.envelopeFrom == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "envelopeFrom");
            }
            if ($.from == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "from");
            }
            if ($.fromDisplayName == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "fromDisplayName");
            }
            if ($.host == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "host");
            }
            if ($.port == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "port");
            }
            if ($.replyTo == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "replyTo");
            }
            if ($.replyToDisplayName == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "replyToDisplayName");
            }
            if ($.ssl == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "ssl");
            }
            if ($.starttls == null) {
                throw new MissingRequiredPropertyException("GetRealmSmtpServerArgs", "starttls");
            }
            return $;
        }
    }

}
