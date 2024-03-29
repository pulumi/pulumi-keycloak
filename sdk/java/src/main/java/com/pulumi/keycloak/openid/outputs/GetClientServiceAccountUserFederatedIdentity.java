// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetClientServiceAccountUserFederatedIdentity {
    private String identityProvider;
    private String userId;
    private String userName;

    private GetClientServiceAccountUserFederatedIdentity() {}
    public String identityProvider() {
        return this.identityProvider;
    }
    public String userId() {
        return this.userId;
    }
    public String userName() {
        return this.userName;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClientServiceAccountUserFederatedIdentity defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String identityProvider;
        private String userId;
        private String userName;
        public Builder() {}
        public Builder(GetClientServiceAccountUserFederatedIdentity defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.identityProvider = defaults.identityProvider;
    	      this.userId = defaults.userId;
    	      this.userName = defaults.userName;
        }

        @CustomType.Setter
        public Builder identityProvider(String identityProvider) {
            if (identityProvider == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserFederatedIdentity", "identityProvider");
            }
            this.identityProvider = identityProvider;
            return this;
        }
        @CustomType.Setter
        public Builder userId(String userId) {
            if (userId == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserFederatedIdentity", "userId");
            }
            this.userId = userId;
            return this;
        }
        @CustomType.Setter
        public Builder userName(String userName) {
            if (userName == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserFederatedIdentity", "userName");
            }
            this.userName = userName;
            return this;
        }
        public GetClientServiceAccountUserFederatedIdentity build() {
            final var _resultValue = new GetClientServiceAccountUserFederatedIdentity();
            _resultValue.identityProvider = identityProvider;
            _resultValue.userId = userId;
            _resultValue.userName = userName;
            return _resultValue;
        }
    }
}
