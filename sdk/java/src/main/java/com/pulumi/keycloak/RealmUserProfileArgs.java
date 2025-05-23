// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.keycloak.inputs.RealmUserProfileAttributeArgs;
import com.pulumi.keycloak.inputs.RealmUserProfileGroupArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RealmUserProfileArgs extends com.pulumi.resources.ResourceArgs {

    public static final RealmUserProfileArgs Empty = new RealmUserProfileArgs();

    /**
     * An ordered list of attributes.
     * 
     */
    @Import(name="attributes")
    private @Nullable Output<List<RealmUserProfileAttributeArgs>> attributes;

    /**
     * @return An ordered list of attributes.
     * 
     */
    public Optional<Output<List<RealmUserProfileAttributeArgs>>> attributes() {
        return Optional.ofNullable(this.attributes);
    }

    /**
     * A list of groups.
     * 
     */
    @Import(name="groups")
    private @Nullable Output<List<RealmUserProfileGroupArgs>> groups;

    /**
     * @return A list of groups.
     * 
     */
    public Optional<Output<List<RealmUserProfileGroupArgs>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * The ID of the realm the user profile applies to.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The ID of the realm the user profile applies to.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     * Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
     * 
     */
    @Import(name="unmanagedAttributePolicy")
    private @Nullable Output<String> unmanagedAttributePolicy;

    /**
     * @return Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
     * 
     */
    public Optional<Output<String>> unmanagedAttributePolicy() {
        return Optional.ofNullable(this.unmanagedAttributePolicy);
    }

    private RealmUserProfileArgs() {}

    private RealmUserProfileArgs(RealmUserProfileArgs $) {
        this.attributes = $.attributes;
        this.groups = $.groups;
        this.realmId = $.realmId;
        this.unmanagedAttributePolicy = $.unmanagedAttributePolicy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RealmUserProfileArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RealmUserProfileArgs $;

        public Builder() {
            $ = new RealmUserProfileArgs();
        }

        public Builder(RealmUserProfileArgs defaults) {
            $ = new RealmUserProfileArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attributes An ordered list of attributes.
         * 
         * @return builder
         * 
         */
        public Builder attributes(@Nullable Output<List<RealmUserProfileAttributeArgs>> attributes) {
            $.attributes = attributes;
            return this;
        }

        /**
         * @param attributes An ordered list of attributes.
         * 
         * @return builder
         * 
         */
        public Builder attributes(List<RealmUserProfileAttributeArgs> attributes) {
            return attributes(Output.of(attributes));
        }

        /**
         * @param attributes An ordered list of attributes.
         * 
         * @return builder
         * 
         */
        public Builder attributes(RealmUserProfileAttributeArgs... attributes) {
            return attributes(List.of(attributes));
        }

        /**
         * @param groups A list of groups.
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<List<RealmUserProfileGroupArgs>> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups A list of groups.
         * 
         * @return builder
         * 
         */
        public Builder groups(List<RealmUserProfileGroupArgs> groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param groups A list of groups.
         * 
         * @return builder
         * 
         */
        public Builder groups(RealmUserProfileGroupArgs... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param realmId The ID of the realm the user profile applies to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The ID of the realm the user profile applies to.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param unmanagedAttributePolicy Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
         * 
         * @return builder
         * 
         */
        public Builder unmanagedAttributePolicy(@Nullable Output<String> unmanagedAttributePolicy) {
            $.unmanagedAttributePolicy = unmanagedAttributePolicy;
            return this;
        }

        /**
         * @param unmanagedAttributePolicy Unmanaged attributes are user attributes not explicitly defined in the user profile configuration. By default, unmanaged attributes are not enabled. Value could be one of `DISABLED`, `ENABLED`, `ADMIN_EDIT` or `ADMIN_VIEW`. If value is not specified it means `DISABLED`
         * 
         * @return builder
         * 
         */
        public Builder unmanagedAttributePolicy(String unmanagedAttributePolicy) {
            return unmanagedAttributePolicy(Output.of(unmanagedAttributePolicy));
        }

        public RealmUserProfileArgs build() {
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("RealmUserProfileArgs", "realmId");
            }
            return $;
        }
    }

}
