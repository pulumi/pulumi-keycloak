// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetGroupResult {
    private Map<String,Object> attributes;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    private String parentId;
    private String path;
    private String realmId;

    private GetGroupResult() {}
    public Map<String,Object> attributes() {
        return this.attributes;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String name() {
        return this.name;
    }
    public String parentId() {
        return this.parentId;
    }
    public String path() {
        return this.path;
    }
    public String realmId() {
        return this.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetGroupResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,Object> attributes;
        private String id;
        private String name;
        private String parentId;
        private String path;
        private String realmId;
        public Builder() {}
        public Builder(GetGroupResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attributes = defaults.attributes;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.parentId = defaults.parentId;
    	      this.path = defaults.path;
    	      this.realmId = defaults.realmId;
        }

        @CustomType.Setter
        public Builder attributes(Map<String,Object> attributes) {
            this.attributes = Objects.requireNonNull(attributes);
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
        public Builder parentId(String parentId) {
            this.parentId = Objects.requireNonNull(parentId);
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        @CustomType.Setter
        public Builder realmId(String realmId) {
            this.realmId = Objects.requireNonNull(realmId);
            return this;
        }
        public GetGroupResult build() {
            final var o = new GetGroupResult();
            o.attributes = attributes;
            o.id = id;
            o.name = name;
            o.parentId = parentId;
            o.path = path;
            o.realmId = realmId;
            return o;
        }
    }
}