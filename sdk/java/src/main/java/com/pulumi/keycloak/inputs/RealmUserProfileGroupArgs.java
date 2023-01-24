// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmUserProfileGroupArgs extends com.pulumi.resources.ResourceArgs {

    public static final RealmUserProfileGroupArgs Empty = new RealmUserProfileGroupArgs();

    /**
     * A map of annotations for the attribute. Values can be a String or a json object.
     * 
     */
    @Import(name="annotations")
    private @Nullable Output<Map<String,String>> annotations;

    /**
     * @return A map of annotations for the attribute. Values can be a String or a json object.
     * 
     */
    public Optional<Output<Map<String,String>>> annotations() {
        return Optional.ofNullable(this.annotations);
    }

    /**
     * The display description of the group.
     * 
     */
    @Import(name="displayDescription")
    private @Nullable Output<String> displayDescription;

    /**
     * @return The display description of the group.
     * 
     */
    public Optional<Output<String>> displayDescription() {
        return Optional.ofNullable(this.displayDescription);
    }

    /**
     * The display header of the group.
     * 
     */
    @Import(name="displayHeader")
    private @Nullable Output<String> displayHeader;

    /**
     * @return The display header of the group.
     * 
     */
    public Optional<Output<String>> displayHeader() {
        return Optional.ofNullable(this.displayHeader);
    }

    /**
     * The name of the attribute.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The name of the attribute.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    private RealmUserProfileGroupArgs() {}

    private RealmUserProfileGroupArgs(RealmUserProfileGroupArgs $) {
        this.annotations = $.annotations;
        this.displayDescription = $.displayDescription;
        this.displayHeader = $.displayHeader;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmUserProfileGroupArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmUserProfileGroupArgs $;

        public Builder() {
            $ = new RealmUserProfileGroupArgs();
        }

        public Builder(RealmUserProfileGroupArgs defaults) {
            $ = new RealmUserProfileGroupArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param annotations A map of annotations for the attribute. Values can be a String or a json object.
         * 
         * @return builder
         * 
         */
        public Builder annotations(@Nullable Output<Map<String,String>> annotations) {
            $.annotations = annotations;
            return this;
        }

        /**
         * @param annotations A map of annotations for the attribute. Values can be a String or a json object.
         * 
         * @return builder
         * 
         */
        public Builder annotations(Map<String,String> annotations) {
            return annotations(Output.of(annotations));
        }

        /**
         * @param displayDescription The display description of the group.
         * 
         * @return builder
         * 
         */
        public Builder displayDescription(@Nullable Output<String> displayDescription) {
            $.displayDescription = displayDescription;
            return this;
        }

        /**
         * @param displayDescription The display description of the group.
         * 
         * @return builder
         * 
         */
        public Builder displayDescription(String displayDescription) {
            return displayDescription(Output.of(displayDescription));
        }

        /**
         * @param displayHeader The display header of the group.
         * 
         * @return builder
         * 
         */
        public Builder displayHeader(@Nullable Output<String> displayHeader) {
            $.displayHeader = displayHeader;
            return this;
        }

        /**
         * @param displayHeader The display header of the group.
         * 
         * @return builder
         * 
         */
        public Builder displayHeader(String displayHeader) {
            return displayHeader(Output.of(displayHeader));
        }

        /**
         * @param name The name of the attribute.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the attribute.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public RealmUserProfileGroupArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            return $;
        }
    }

}
