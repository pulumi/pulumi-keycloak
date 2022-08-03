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
    private final Map<String,Object> config;
    private final String id;
    private final String name;
    private final String protocol;
    private final String protocolMapper;

    @CustomType.Constructor
    private GetClientDescriptionConverterProtocolMapper(
        @CustomType.Parameter("config") Map<String,Object> config,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("protocol") String protocol,
        @CustomType.Parameter("protocolMapper") String protocolMapper) {
        this.config = config;
        this.id = id;
        this.name = name;
        this.protocol = protocol;
        this.protocolMapper = protocolMapper;
    }

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

    public static final class Builder {
        private Map<String,Object> config;
        private String id;
        private String name;
        private String protocol;
        private String protocolMapper;

        public Builder() {
    	      // Empty
        }

        public Builder(GetClientDescriptionConverterProtocolMapper defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.config = defaults.config;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.protocol = defaults.protocol;
    	      this.protocolMapper = defaults.protocolMapper;
        }

        public Builder config(Map<String,Object> config) {
            this.config = Objects.requireNonNull(config);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder protocol(String protocol) {
            this.protocol = Objects.requireNonNull(protocol);
            return this;
        }
        public Builder protocolMapper(String protocolMapper) {
            this.protocolMapper = Objects.requireNonNull(protocolMapper);
            return this;
        }        public GetClientDescriptionConverterProtocolMapper build() {
            return new GetClientDescriptionConverterProtocolMapper(config, id, name, protocol, protocolMapper);
        }
    }
}
