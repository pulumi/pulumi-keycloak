// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClientScopeResult {
    private String consentScreenText;
    private String description;
    private Integer guiOrder;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private Boolean includeInTokenScope;
    private String name;
    private String realmId;

    private GetClientScopeResult() {}
    public String consentScreenText() {
        return this.consentScreenText;
    }
    public String description() {
        return this.description;
    }
    public Integer guiOrder() {
        return this.guiOrder;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public Boolean includeInTokenScope() {
        return this.includeInTokenScope;
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

    public static Builder builder(GetClientScopeResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String consentScreenText;
        private String description;
        private Integer guiOrder;
        private String id;
        private Boolean includeInTokenScope;
        private String name;
        private String realmId;
        public Builder() {}
        public Builder(GetClientScopeResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.consentScreenText = defaults.consentScreenText;
    	      this.description = defaults.description;
    	      this.guiOrder = defaults.guiOrder;
    	      this.id = defaults.id;
    	      this.includeInTokenScope = defaults.includeInTokenScope;
    	      this.name = defaults.name;
    	      this.realmId = defaults.realmId;
        }

        @CustomType.Setter
        public Builder consentScreenText(String consentScreenText) {
            if (consentScreenText == null) {
              throw new MissingRequiredPropertyException("GetClientScopeResult", "consentScreenText");
            }
            this.consentScreenText = consentScreenText;
            return this;
        }
        @CustomType.Setter
        public Builder description(String description) {
            if (description == null) {
              throw new MissingRequiredPropertyException("GetClientScopeResult", "description");
            }
            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder guiOrder(Integer guiOrder) {
            if (guiOrder == null) {
              throw new MissingRequiredPropertyException("GetClientScopeResult", "guiOrder");
            }
            this.guiOrder = guiOrder;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetClientScopeResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder includeInTokenScope(Boolean includeInTokenScope) {
            if (includeInTokenScope == null) {
              throw new MissingRequiredPropertyException("GetClientScopeResult", "includeInTokenScope");
            }
            this.includeInTokenScope = includeInTokenScope;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("GetClientScopeResult", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder realmId(String realmId) {
            if (realmId == null) {
              throw new MissingRequiredPropertyException("GetClientScopeResult", "realmId");
            }
            this.realmId = realmId;
            return this;
        }
        public GetClientScopeResult build() {
            final var _resultValue = new GetClientScopeResult();
            _resultValue.consentScreenText = consentScreenText;
            _resultValue.description = description;
            _resultValue.guiOrder = guiOrder;
            _resultValue.id = id;
            _resultValue.includeInTokenScope = includeInTokenScope;
            _resultValue.name = name;
            _resultValue.realmId = realmId;
            return _resultValue;
        }
    }
}
