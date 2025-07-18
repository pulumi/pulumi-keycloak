// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.openid.outputs.GetClientServiceAccountUserFederatedIdentity;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;

@CustomType
public final class GetClientServiceAccountUserResult {
    /**
     * @return (Computed) The service account user&#39;s attributes.
     * 
     */
    private Map<String,String> attributes;
    private String clientId;
    /**
     * @return (Computed) The service account user&#39;s email.
     * 
     */
    private String email;
    private Boolean emailVerified;
    /**
     * @return (Computed) Whether the service account user is enabled.
     * 
     */
    private Boolean enabled;
    /**
     * @return (Computed) This attribute exists in order to adhere to the spec of a Keycloak user, but a service account user will never have a federated identity, so this will always be `null`.
     * 
     */
    private List<GetClientServiceAccountUserFederatedIdentity> federatedIdentities;
    /**
     * @return (Computed) The service account user&#39;s first name.
     * 
     */
    private String firstName;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return (Computed) The service account user&#39;s last name.
     * 
     */
    private String lastName;
    private String realmId;
    private List<String> requiredActions;
    /**
     * @return (Computed) The service account user&#39;s username.
     * 
     */
    private String username;

    private GetClientServiceAccountUserResult() {}
    /**
     * @return (Computed) The service account user&#39;s attributes.
     * 
     */
    public Map<String,String> attributes() {
        return this.attributes;
    }
    public String clientId() {
        return this.clientId;
    }
    /**
     * @return (Computed) The service account user&#39;s email.
     * 
     */
    public String email() {
        return this.email;
    }
    public Boolean emailVerified() {
        return this.emailVerified;
    }
    /**
     * @return (Computed) Whether the service account user is enabled.
     * 
     */
    public Boolean enabled() {
        return this.enabled;
    }
    /**
     * @return (Computed) This attribute exists in order to adhere to the spec of a Keycloak user, but a service account user will never have a federated identity, so this will always be `null`.
     * 
     */
    public List<GetClientServiceAccountUserFederatedIdentity> federatedIdentities() {
        return this.federatedIdentities;
    }
    /**
     * @return (Computed) The service account user&#39;s first name.
     * 
     */
    public String firstName() {
        return this.firstName;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (Computed) The service account user&#39;s last name.
     * 
     */
    public String lastName() {
        return this.lastName;
    }
    public String realmId() {
        return this.realmId;
    }
    public List<String> requiredActions() {
        return this.requiredActions;
    }
    /**
     * @return (Computed) The service account user&#39;s username.
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetClientServiceAccountUserResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private Map<String,String> attributes;
        private String clientId;
        private String email;
        private Boolean emailVerified;
        private Boolean enabled;
        private List<GetClientServiceAccountUserFederatedIdentity> federatedIdentities;
        private String firstName;
        private String id;
        private String lastName;
        private String realmId;
        private List<String> requiredActions;
        private String username;
        public Builder() {}
        public Builder(GetClientServiceAccountUserResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.attributes = defaults.attributes;
    	      this.clientId = defaults.clientId;
    	      this.email = defaults.email;
    	      this.emailVerified = defaults.emailVerified;
    	      this.enabled = defaults.enabled;
    	      this.federatedIdentities = defaults.federatedIdentities;
    	      this.firstName = defaults.firstName;
    	      this.id = defaults.id;
    	      this.lastName = defaults.lastName;
    	      this.realmId = defaults.realmId;
    	      this.requiredActions = defaults.requiredActions;
    	      this.username = defaults.username;
        }

        @CustomType.Setter
        public Builder attributes(Map<String,String> attributes) {
            if (attributes == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "attributes");
            }
            this.attributes = attributes;
            return this;
        }
        @CustomType.Setter
        public Builder clientId(String clientId) {
            if (clientId == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "clientId");
            }
            this.clientId = clientId;
            return this;
        }
        @CustomType.Setter
        public Builder email(String email) {
            if (email == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "email");
            }
            this.email = email;
            return this;
        }
        @CustomType.Setter
        public Builder emailVerified(Boolean emailVerified) {
            if (emailVerified == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "emailVerified");
            }
            this.emailVerified = emailVerified;
            return this;
        }
        @CustomType.Setter
        public Builder enabled(Boolean enabled) {
            if (enabled == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "enabled");
            }
            this.enabled = enabled;
            return this;
        }
        @CustomType.Setter
        public Builder federatedIdentities(List<GetClientServiceAccountUserFederatedIdentity> federatedIdentities) {
            if (federatedIdentities == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "federatedIdentities");
            }
            this.federatedIdentities = federatedIdentities;
            return this;
        }
        public Builder federatedIdentities(GetClientServiceAccountUserFederatedIdentity... federatedIdentities) {
            return federatedIdentities(List.of(federatedIdentities));
        }
        @CustomType.Setter
        public Builder firstName(String firstName) {
            if (firstName == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "firstName");
            }
            this.firstName = firstName;
            return this;
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder lastName(String lastName) {
            if (lastName == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "lastName");
            }
            this.lastName = lastName;
            return this;
        }
        @CustomType.Setter
        public Builder realmId(String realmId) {
            if (realmId == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "realmId");
            }
            this.realmId = realmId;
            return this;
        }
        @CustomType.Setter
        public Builder requiredActions(List<String> requiredActions) {
            if (requiredActions == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "requiredActions");
            }
            this.requiredActions = requiredActions;
            return this;
        }
        public Builder requiredActions(String... requiredActions) {
            return requiredActions(List.of(requiredActions));
        }
        @CustomType.Setter
        public Builder username(String username) {
            if (username == null) {
              throw new MissingRequiredPropertyException("GetClientServiceAccountUserResult", "username");
            }
            this.username = username;
            return this;
        }
        public GetClientServiceAccountUserResult build() {
            final var _resultValue = new GetClientServiceAccountUserResult();
            _resultValue.attributes = attributes;
            _resultValue.clientId = clientId;
            _resultValue.email = email;
            _resultValue.emailVerified = emailVerified;
            _resultValue.enabled = enabled;
            _resultValue.federatedIdentities = federatedIdentities;
            _resultValue.firstName = firstName;
            _resultValue.id = id;
            _resultValue.lastName = lastName;
            _resultValue.realmId = realmId;
            _resultValue.requiredActions = requiredActions;
            _resultValue.username = username;
            return _resultValue;
        }
    }
}
