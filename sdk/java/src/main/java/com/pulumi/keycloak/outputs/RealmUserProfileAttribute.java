// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.keycloak.outputs.RealmUserProfileAttributePermissions;
import com.pulumi.keycloak.outputs.RealmUserProfileAttributeValidator;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RealmUserProfileAttribute {
    /**
     * @return A map of annotations for the group.
     * 
     */
    private final @Nullable Map<String,String> annotations;
    /**
     * @return The display name of the attribute.
     * 
     */
    private final @Nullable String displayName;
    /**
     * @return A list of scopes. The attribute will only be enabled when these scopes are requested by clients.
     * 
     */
    private final @Nullable List<String> enabledWhenScopes;
    /**
     * @return The group that the attribute belong to.
     * 
     */
    private final @Nullable String group;
    /**
     * @return The name of the group.
     * 
     */
    private final String name;
    /**
     * @return The permissions configuration information.
     * 
     */
    private final @Nullable RealmUserProfileAttributePermissions permissions;
    /**
     * @return A list of roles for which the attribute will be required.
     * 
     */
    private final @Nullable List<String> requiredForRoles;
    /**
     * @return A list of scopes for which the attribute will be required.
     * 
     */
    private final @Nullable List<String> requiredForScopes;
    /**
     * @return A list of validators for the attribute.
     * 
     */
    private final @Nullable List<RealmUserProfileAttributeValidator> validators;

    @CustomType.Constructor
    private RealmUserProfileAttribute(
        @CustomType.Parameter("annotations") @Nullable Map<String,String> annotations,
        @CustomType.Parameter("displayName") @Nullable String displayName,
        @CustomType.Parameter("enabledWhenScopes") @Nullable List<String> enabledWhenScopes,
        @CustomType.Parameter("group") @Nullable String group,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("permissions") @Nullable RealmUserProfileAttributePermissions permissions,
        @CustomType.Parameter("requiredForRoles") @Nullable List<String> requiredForRoles,
        @CustomType.Parameter("requiredForScopes") @Nullable List<String> requiredForScopes,
        @CustomType.Parameter("validators") @Nullable List<RealmUserProfileAttributeValidator> validators) {
        this.annotations = annotations;
        this.displayName = displayName;
        this.enabledWhenScopes = enabledWhenScopes;
        this.group = group;
        this.name = name;
        this.permissions = permissions;
        this.requiredForRoles = requiredForRoles;
        this.requiredForScopes = requiredForScopes;
        this.validators = validators;
    }

    /**
     * @return A map of annotations for the group.
     * 
     */
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
     * @return The group that the attribute belong to.
     * 
     */
    public Optional<String> group() {
        return Optional.ofNullable(this.group);
    }
    /**
     * @return The name of the group.
     * 
     */
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

    public static final class Builder {
        private @Nullable Map<String,String> annotations;
        private @Nullable String displayName;
        private @Nullable List<String> enabledWhenScopes;
        private @Nullable String group;
        private String name;
        private @Nullable RealmUserProfileAttributePermissions permissions;
        private @Nullable List<String> requiredForRoles;
        private @Nullable List<String> requiredForScopes;
        private @Nullable List<RealmUserProfileAttributeValidator> validators;

        public Builder() {
    	      // Empty
        }

        public Builder(RealmUserProfileAttribute defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.displayName = defaults.displayName;
    	      this.enabledWhenScopes = defaults.enabledWhenScopes;
    	      this.group = defaults.group;
    	      this.name = defaults.name;
    	      this.permissions = defaults.permissions;
    	      this.requiredForRoles = defaults.requiredForRoles;
    	      this.requiredForScopes = defaults.requiredForScopes;
    	      this.validators = defaults.validators;
        }

        public Builder annotations(@Nullable Map<String,String> annotations) {
            this.annotations = annotations;
            return this;
        }
        public Builder displayName(@Nullable String displayName) {
            this.displayName = displayName;
            return this;
        }
        public Builder enabledWhenScopes(@Nullable List<String> enabledWhenScopes) {
            this.enabledWhenScopes = enabledWhenScopes;
            return this;
        }
        public Builder enabledWhenScopes(String... enabledWhenScopes) {
            return enabledWhenScopes(List.of(enabledWhenScopes));
        }
        public Builder group(@Nullable String group) {
            this.group = group;
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder permissions(@Nullable RealmUserProfileAttributePermissions permissions) {
            this.permissions = permissions;
            return this;
        }
        public Builder requiredForRoles(@Nullable List<String> requiredForRoles) {
            this.requiredForRoles = requiredForRoles;
            return this;
        }
        public Builder requiredForRoles(String... requiredForRoles) {
            return requiredForRoles(List.of(requiredForRoles));
        }
        public Builder requiredForScopes(@Nullable List<String> requiredForScopes) {
            this.requiredForScopes = requiredForScopes;
            return this;
        }
        public Builder requiredForScopes(String... requiredForScopes) {
            return requiredForScopes(List.of(requiredForScopes));
        }
        public Builder validators(@Nullable List<RealmUserProfileAttributeValidator> validators) {
            this.validators = validators;
            return this;
        }
        public Builder validators(RealmUserProfileAttributeValidator... validators) {
            return validators(List.of(validators));
        }        public RealmUserProfileAttribute build() {
            return new RealmUserProfileAttribute(annotations, displayName, enabledWhenScopes, group, name, permissions, requiredForRoles, requiredForScopes, validators);
        }
    }
}
