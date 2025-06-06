// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.outputs.RealmUserProfileAttributePermissions;
import com.pulumi.keycloak.outputs.RealmUserProfileAttributeValidator;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RealmUserProfileAttribute {
    private @Nullable Map<String,String> annotations;
    /**
     * @return The display name of the attribute.
     * 
     */
    private @Nullable String displayName;
    /**
     * @return A list of scopes. The attribute will only be enabled when these scopes are requested by clients.
     * 
     */
    private @Nullable List<String> enabledWhenScopes;
    /**
     * @return A list of groups.
     * 
     */
    private @Nullable String group;
    /**
     * @return If the attribute supports multiple values. Defaults to `false`.
     * 
     */
    private @Nullable Boolean multiValued;
    private String name;
    /**
     * @return The permissions configuration information.
     * 
     */
    private @Nullable RealmUserProfileAttributePermissions permissions;
    /**
     * @return A list of roles for which the attribute will be required.
     * 
     */
    private @Nullable List<String> requiredForRoles;
    /**
     * @return A list of scopes for which the attribute will be required.
     * 
     */
    private @Nullable List<String> requiredForScopes;
    /**
     * @return A list of validators for the attribute.
     * 
     */
    private @Nullable List<RealmUserProfileAttributeValidator> validators;

    private RealmUserProfileAttribute() {}
    public Map<String,String> annotations() {
        return this.annotations == null ? Map.of() : this.annotations;
    }
    /**
     * @return The display name of the attribute.
     * 
     */
    public Optional<String> displayName() {
        return Optional.ofNullable(this.displayName);
    }
    /**
     * @return A list of scopes. The attribute will only be enabled when these scopes are requested by clients.
     * 
     */
    public List<String> enabledWhenScopes() {
        return this.enabledWhenScopes == null ? List.of() : this.enabledWhenScopes;
    }
    /**
     * @return A list of groups.
     * 
     */
    public Optional<String> group() {
        return Optional.ofNullable(this.group);
    }
    /**
     * @return If the attribute supports multiple values. Defaults to `false`.
     * 
     */
    public Optional<Boolean> multiValued() {
        return Optional.ofNullable(this.multiValued);
    }
    public String name() {
        return this.name;
    }
    /**
     * @return The permissions configuration information.
     * 
     */
    public Optional<RealmUserProfileAttributePermissions> permissions() {
        return Optional.ofNullable(this.permissions);
    }
    /**
     * @return A list of roles for which the attribute will be required.
     * 
     */
    public List<String> requiredForRoles() {
        return this.requiredForRoles == null ? List.of() : this.requiredForRoles;
    }
    /**
     * @return A list of scopes for which the attribute will be required.
     * 
     */
    public List<String> requiredForScopes() {
        return this.requiredForScopes == null ? List.of() : this.requiredForScopes;
    }
    /**
     * @return A list of validators for the attribute.
     * 
     */
    public List<RealmUserProfileAttributeValidator> validators() {
        return this.validators == null ? List.of() : this.validators;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RealmUserProfileAttribute defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> annotations;
        private @Nullable String displayName;
        private @Nullable List<String> enabledWhenScopes;
        private @Nullable String group;
        private @Nullable Boolean multiValued;
        private String name;
        private @Nullable RealmUserProfileAttributePermissions permissions;
        private @Nullable List<String> requiredForRoles;
        private @Nullable List<String> requiredForScopes;
        private @Nullable List<RealmUserProfileAttributeValidator> validators;
        public Builder() {}
        public Builder(RealmUserProfileAttribute defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.displayName = defaults.displayName;
    	      this.enabledWhenScopes = defaults.enabledWhenScopes;
    	      this.group = defaults.group;
    	      this.multiValued = defaults.multiValued;
    	      this.name = defaults.name;
    	      this.permissions = defaults.permissions;
    	      this.requiredForRoles = defaults.requiredForRoles;
    	      this.requiredForScopes = defaults.requiredForScopes;
    	      this.validators = defaults.validators;
        }

        @CustomType.Setter
        public Builder annotations(@Nullable Map<String,String> annotations) {

            this.annotations = annotations;
            return this;
        }
        @CustomType.Setter
        public Builder displayName(@Nullable String displayName) {

            this.displayName = displayName;
            return this;
        }
        @CustomType.Setter
        public Builder enabledWhenScopes(@Nullable List<String> enabledWhenScopes) {

            this.enabledWhenScopes = enabledWhenScopes;
            return this;
        }
        public Builder enabledWhenScopes(String... enabledWhenScopes) {
            return enabledWhenScopes(List.of(enabledWhenScopes));
        }
        @CustomType.Setter
        public Builder group(@Nullable String group) {

            this.group = group;
            return this;
        }
        @CustomType.Setter
        public Builder multiValued(@Nullable Boolean multiValued) {

            this.multiValued = multiValued;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            if (name == null) {
              throw new MissingRequiredPropertyException("RealmUserProfileAttribute", "name");
            }
            this.name = name;
            return this;
        }
        @CustomType.Setter
        public Builder permissions(@Nullable RealmUserProfileAttributePermissions permissions) {

            this.permissions = permissions;
            return this;
        }
        @CustomType.Setter
        public Builder requiredForRoles(@Nullable List<String> requiredForRoles) {

            this.requiredForRoles = requiredForRoles;
            return this;
        }
        public Builder requiredForRoles(String... requiredForRoles) {
            return requiredForRoles(List.of(requiredForRoles));
        }
        @CustomType.Setter
        public Builder requiredForScopes(@Nullable List<String> requiredForScopes) {

            this.requiredForScopes = requiredForScopes;
            return this;
        }
        public Builder requiredForScopes(String... requiredForScopes) {
            return requiredForScopes(List.of(requiredForScopes));
        }
        @CustomType.Setter
        public Builder validators(@Nullable List<RealmUserProfileAttributeValidator> validators) {

            this.validators = validators;
            return this;
        }
        public Builder validators(RealmUserProfileAttributeValidator... validators) {
            return validators(List.of(validators));
        }
        public RealmUserProfileAttribute build() {
            final var _resultValue = new RealmUserProfileAttribute();
            _resultValue.annotations = annotations;
            _resultValue.displayName = displayName;
            _resultValue.enabledWhenScopes = enabledWhenScopes;
            _resultValue.group = group;
            _resultValue.multiValued = multiValued;
            _resultValue.name = name;
            _resultValue.permissions = permissions;
            _resultValue.requiredForRoles = requiredForRoles;
            _resultValue.requiredForScopes = requiredForScopes;
            _resultValue.validators = validators;
            return _resultValue;
        }
    }
}
