// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class GetUserRealmRolesResult {
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    private String realmId;
    /**
     * @return (Computed) A list of realm roles that belong to this user.
     * 
     */
    private List<String> roleNames;
    private String userId;

    private GetUserRealmRolesResult() {}
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String realmId() {
        return this.realmId;
    }
    /**
     * @return (Computed) A list of realm roles that belong to this user.
     * 
     */
    public List<String> roleNames() {
        return this.roleNames;
    }
    public String userId() {
        return this.userId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetUserRealmRolesResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String id;
        private String realmId;
        private List<String> roleNames;
        private String userId;
        public Builder() {}
        public Builder(GetUserRealmRolesResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.id = defaults.id;
    	      this.realmId = defaults.realmId;
    	      this.roleNames = defaults.roleNames;
    	      this.userId = defaults.userId;
        }

        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetUserRealmRolesResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder realmId(String realmId) {
            if (realmId == null) {
              throw new MissingRequiredPropertyException("GetUserRealmRolesResult", "realmId");
            }
            this.realmId = realmId;
            return this;
        }
        @CustomType.Setter
        public Builder roleNames(List<String> roleNames) {
            if (roleNames == null) {
              throw new MissingRequiredPropertyException("GetUserRealmRolesResult", "roleNames");
            }
            this.roleNames = roleNames;
            return this;
        }
        public Builder roleNames(String... roleNames) {
            return roleNames(List.of(roleNames));
        }
        @CustomType.Setter
        public Builder userId(String userId) {
            if (userId == null) {
              throw new MissingRequiredPropertyException("GetUserRealmRolesResult", "userId");
            }
            this.userId = userId;
            return this;
        }
        public GetUserRealmRolesResult build() {
            final var _resultValue = new GetUserRealmRolesResult();
            _resultValue.id = id;
            _resultValue.realmId = realmId;
            _resultValue.roleNames = roleNames;
            _resultValue.userId = userId;
            return _resultValue;
        }
    }
}
