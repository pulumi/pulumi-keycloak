// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetGroupResult {
    private Map<String,String> attributes;
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
    public Map<String,String> attributes() {
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
        private Map<String,String> attributes;
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
        public Builder attributes(Map<String,String> attributes) {
            if (attributes == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "attributes");
            }
            this.attributes = attributes;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder parentId(String parentId) {
            if (parentId == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "parentId");
            }
            this.parentId = parentId;
            return this;
        }
        @CustomType.Setter
        public Builder path(String path) {
            if (path == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "path");
            }
            this.path = path;
            return this;
        }
        @CustomType.Setter
        public Builder realmId(String realmId) {
            if (realmId == null) {
              throw new MissingRequiredPropertyException("GetGroupResult", "realmId");
            }
            this.realmId = realmId;
            return this;
        }
        public GetGroupResult build() {
            final var _resultValue = new GetGroupResult();
            _resultValue.attributes = attributes;
            _resultValue.id = id;
            _resultValue.name = name;
            _resultValue.parentId = parentId;
            _resultValue.path = path;
            _resultValue.realmId = realmId;
            return _resultValue;
        }
    }
}
