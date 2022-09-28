// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class ClientRolePolicyRole {
    private String id;
    private Boolean required;

    private ClientRolePolicyRole() {}
    public String id() {
        return this.id;
    }
    public Boolean required() {
        return this.required;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(ClientRolePolicyRole defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private Boolean required;
        public Builder() {}
        public Builder(ClientRolePolicyRole defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.required = defaults.required;
        }

        @CustomType.Setter
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        @CustomType.Setter
        public Builder required(Boolean required) {
            this.required = Objects.requireNonNull(required);
            return this;
        }
        public ClientRolePolicyRole build() {
            final var o = new ClientRolePolicyRole();
            o.id = id;
            o.required = required;
            return o;
        }
    }
}
