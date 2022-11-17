// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetRoleResult {
    private Map<String,Object> attributes;
    private @Nullable String clientId;
    private List<String> compositeRoles;
    /**
     * @return (Computed) The description of the role.
     * 
     */
    private String description;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String name;
    private String realmId;

    private GetRoleResult() {}
    public Map<String,Object> attributes() {
        return this.attributes;
    }
    public Optional<String> clientId() {
        return Optional.ofNullable(this.clientId);
    }
    public List<String> compositeRoles() {
        return this.compositeRoles;
    }
    /**
     * @return (Computed) The description of the role.
     * 
     */
    public String description() {
        return this.description;
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
    public String realmId() {
        return this.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRoleResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,Object> attributes;
        private @Nullable String clientId;
        private List<String> compositeRoles;
        private String description;
        private String id;
        private String name;
        private String realmId;
        public Builder() {}
        public Builder(GetRoleResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attributes = defaults.attributes;
    	      this.clientId = defaults.clientId;
    	      this.compositeRoles = defaults.compositeRoles;
    	      this.description = defaults.description;
    	      this.id = defaults.id;
    	      this.name = defaults.name;
    	      this.realmId = defaults.realmId;
        }

        @CustomType.Setter
        public Builder attributes(Map<String,Object> attributes) {
            this.attributes = Objects.requireNonNull(attributes);
            return this;
        }
        @CustomType.Setter
        public Builder clientId(@Nullable String clientId) {
            this.clientId = clientId;
            return this;
        }
        @CustomType.Setter
        public Builder compositeRoles(List<String> compositeRoles) {
            this.compositeRoles = Objects.requireNonNull(compositeRoles);
            return this;
        }
        public Builder compositeRoles(String... compositeRoles) {
            return compositeRoles(List.of(compositeRoles));
        }
        @CustomType.Setter
        public Builder description(String description) {
            this.description = Objects.requireNonNull(description);
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
        public Builder realmId(String realmId) {
            this.realmId = Objects.requireNonNull(realmId);
            return this;
        }
        public GetRoleResult build() {
            final var o = new GetRoleResult();
            o.attributes = attributes;
            o.clientId = clientId;
            o.compositeRoles = compositeRoles;
            o.description = description;
            o.id = id;
            o.name = name;
            o.realmId = realmId;
            return o;
        }
    }
}