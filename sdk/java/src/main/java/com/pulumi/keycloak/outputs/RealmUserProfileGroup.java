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
    private @Nullable Map<String,String> annotations;
    /**
     * @return The display description of the group.
     * 
     */
    private @Nullable String displayDescription;
    /**
     * @return The display header of the group.
     * 
     */
    private @Nullable String displayHeader;
    /**
     * @return The name of the group.
     * 
     */
    private String name;

    private RealmUserProfileGroup() {}
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
    @CustomType.Builder
    public static final class Builder {
        private @Nullable Map<String,String> annotations;
        private @Nullable String displayDescription;
        private @Nullable String displayHeader;
        private String name;
        public Builder() {}
        public Builder(RealmUserProfileGroup defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.annotations = defaults.annotations;
    	      this.displayDescription = defaults.displayDescription;
    	      this.displayHeader = defaults.displayHeader;
    	      this.name = defaults.name;
        }

        @CustomType.Setter
        public Builder annotations(@Nullable Map<String,String> annotations) {
            this.annotations = annotations;
            return this;
        }
        @CustomType.Setter
        public Builder displayDescription(@Nullable String displayDescription) {
            this.displayDescription = displayDescription;
            return this;
        }
        @CustomType.Setter
        public Builder displayHeader(@Nullable String displayHeader) {
            this.displayHeader = displayHeader;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public RealmUserProfileGroup build() {
            final var o = new RealmUserProfileGroup();
            o.annotations = annotations;
            o.displayDescription = displayDescription;
            o.displayHeader = displayHeader;
            o.name = name;
            return o;
        }
    }
}