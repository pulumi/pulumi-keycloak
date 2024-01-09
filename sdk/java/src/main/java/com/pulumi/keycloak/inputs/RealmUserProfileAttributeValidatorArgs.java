// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmUserProfileAttributeValidatorArgs extends com.pulumi.resources.ResourceArgs {

    public static final RealmUserProfileAttributeValidatorArgs Empty = new RealmUserProfileAttributeValidatorArgs();

    /**
     * A map defining the configuration of the validator. Values can be a String or a json object.
     * 
     */
    @Import(name="config")
    private @Nullable Output<Map<String,String>> config;

    /**
     * @return A map defining the configuration of the validator. Values can be a String or a json object.
     * 
     */
    public Optional<Output<Map<String,String>>> config() {
        return Optional.ofNullable(this.config);
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

    private RealmUserProfileAttributeValidatorArgs() {}

    private RealmUserProfileAttributeValidatorArgs(RealmUserProfileAttributeValidatorArgs $) {
        this.config = $.config;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmUserProfileAttributeValidatorArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmUserProfileAttributeValidatorArgs $;

        public Builder() {
            $ = new RealmUserProfileAttributeValidatorArgs();
        }

        public Builder(RealmUserProfileAttributeValidatorArgs defaults) {
            $ = new RealmUserProfileAttributeValidatorArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param config A map defining the configuration of the validator. Values can be a String or a json object.
         * 
         * @return builder
         * 
         */
        public Builder config(@Nullable Output<Map<String,String>> config) {
            $.config = config;
            return this;
        }

        /**
         * @param config A map defining the configuration of the validator. Values can be a String or a json object.
         * 
         * @return builder
         * 
         */
        public Builder config(Map<String,String> config) {
            return config(Output.of(config));
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

        public RealmUserProfileAttributeValidatorArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("RealmUserProfileAttributeValidatorArgs", "name");
            }
            return $;
        }
    }

}
