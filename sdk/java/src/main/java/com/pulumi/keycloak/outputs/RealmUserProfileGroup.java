// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class RealmUserProfileGroup {
    /**
     * @return A map of annotations for the group.
     * 
     */
    private final @Nullable Map<String,String> annotations;
    /**
     * @return The display description of the group.
     * 
     */
    private final @Nullable String displayDescription;
    /**
     * @return The display header of the group.
     * 
     */
    private final @Nullable String displayHeader;
    /**
     * @return The name of the group.
     * 
     */
    private final String name;

    @CustomType.Constructor
    private RealmUserProfileGroup(
        @CustomType.Parameter("annotations") @Nullable Map<String,String> annotations,
        @CustomType.Parameter("displayDescription") @Nullable String displayDescription,
        @CustomType.Parameter("displayHeader") @Nullable String displayHeader,
        @CustomType.Parameter("name") String name) {
        this.annotations = annotations;
        this.displayDescription = displayDescription;
        this.displayHeader = displayHeader;
        this.name = name;
    }

    /**
     * @return A map of annotations for the group.
     * 
     */
    public Map<String,String> annotations() {
        return this.annotations == null ? Map.of() : this.annotations;
    }
    /**
     * @return The display description of the group.
     * 
     */
    public Optional<String> displayDescription() {
        return Optional.ofNullable(this.displayDescription);
    }
    /**
     * @return The display header of the group.
     * 
     */
    public Optional<String> displayHeader() {
        return Optional.ofNullable(this.displayHeader);
    }
    /**
     * @return The name of the group.
     * 
     */
    public String name() {
        return this.name;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(RealmUserProfileGroup defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable Map<String,String> annotations;
        private @Nullable String displayDescription;
        private @Nullable String displayHeader;
        private String name;

        public Builder() {
    	      // Empty
        }

        public Builder(RealmUserProfileGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.displayDescription = defaults.displayDescription;
    	      this.displayHeader = defaults.displayHeader;
    	      this.name = defaults.name;
        }

        public Builder annotations(@Nullable Map<String,String> annotations) {
            this.annotations = annotations;
            return this;
        }
        public Builder displayDescription(@Nullable String displayDescription) {
            this.displayDescription = displayDescription;
            return this;
        }
        public Builder displayHeader(@Nullable String displayHeader) {
            this.displayHeader = displayHeader;
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }        public RealmUserProfileGroup build() {
            return new RealmUserProfileGroup(annotations, displayDescription, displayHeader, name);
        }
    }
}
