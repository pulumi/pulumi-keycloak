// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;


public final class GetOrganizationArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOrganizationArgs Empty = new GetOrganizationArgs();

    /**
     * The organization name.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return The organization name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The name of the realm this organization exists within.
     * 
     */
    @Import(name="realm", required=true)
    private Output<String> realm;

    /**
     * @return The name of the realm this organization exists within.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }

    private GetOrganizationArgs() {}

    private GetOrganizationArgs(GetOrganizationArgs $) {
        this.name = $.name;
        this.realm = $.realm;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOrganizationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOrganizationArgs $;

        public Builder() {
            $ = new GetOrganizationArgs();
        }

        public Builder(GetOrganizationArgs defaults) {
            $ = new GetOrganizationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param name The organization name.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The organization name.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realm The name of the realm this organization exists within.
         * 
         * @return builder
         * 
         */
        public Builder realm(Output<String> realm) {
            $.realm = realm;
            return this;
        }

        /**
         * @param realm The name of the realm this organization exists within.
         * 
         * @return builder
         * 
         */
        public Builder realm(String realm) {
            return realm(Output.of(realm));
        }

        public GetOrganizationArgs build() {
            if ($.name == null) {
                throw new MissingRequiredPropertyException("GetOrganizationArgs", "name");
            }
            if ($.realm == null) {
                throw new MissingRequiredPropertyException("GetOrganizationArgs", "realm");
            }
            return $;
        }
    }

}
