// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetClientDescriptionConverterProtocolMapper {
    private Map<String,Object> config;
    private String id;
    private String name;
    private String protocol;
    private String protocolMapper;

    private GetClientDescriptionConverterProtocolMapper() {}
    public Map<String,Object> config() {
        return this.config;
    }
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public String protocol() {
        return this.protocol;
    }
    public String protocolMapper() {
        return this.protocolMapper;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClientDescriptionConverterProtocolMapper defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,Object> config;
        private String id;
        private String name;
        private String protocol;
        private String protocolMapper;
        public Builder() {}
        public Builder(GetClientDescriptionConverterProtocolMapper defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.config = defaults.config;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.protocol = defaults.protocol;
    	      this.protocolMapper = defaults.protocolMapper;
        }

        @CustomType.Setter
        public Builder config(Map<String,Object> config) {
            this.config = Objects.requireNonNull(config);
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        @CustomType.Setter
        public Builder protocolMapper(String protocolMapper) {
            this.protocolMapper = Objects.requireNonNull(protocolMapper);
            return this;
        }
        public GetClientDescriptionConverterProtocolMapper build() {
            final var o = new GetClientDescriptionConverterProtocolMapper();
            o.config = config;
            o.id = id;
            o.name = name;
            o.protocol = protocol;
            o.protocolMapper = protocolMapper;
            return o;
        }
    }
}