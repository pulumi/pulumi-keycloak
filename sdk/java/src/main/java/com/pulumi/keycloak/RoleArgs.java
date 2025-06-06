// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final RoleArgs Empty = new RoleArgs();

    /**
     * A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     * 
     */
    @Import(name="attributes")
    private @Nullable Output<Map<String,String>> attributes;

    /**
     * @return A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
     * 
     */
    public Optional<Output<Map<String,String>>> attributes() {
        return Optional.ofNullable(this.attributes);
    }

    /**
     * When specified, this role will be created as a client role attached to the client with the provided ID
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return When specified, this role will be created as a client role attached to the client with the provided ID
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
     * 
     */
    @Import(name="compositeRoles")
    private @Nullable Output<List<String>> compositeRoles;

    /**
     * @return When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
     * 
     */
    public Optional<Output<List<String>>> compositeRoles() {
        return Optional.ofNullable(this.compositeRoles);
    }

    /**
     * The description of the role
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the role
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * When `true`, the role with the specified `name` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with roles that Keycloak creates automatically during realm creation, such as the client roles `create-client`, `view-realm`, ... for the client `realm-management` created per realm. Note, that the role will not be removed during destruction if `import` is `true`.
     * 
     */
    @Import(name="import")
    private @Nullable Output<Boolean> import_;

    /**
     * @return When `true`, the role with the specified `name` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with roles that Keycloak creates automatically during realm creation, such as the client roles `create-client`, `view-realm`, ... for the client `realm-management` created per realm. Note, that the role will not be removed during destruction if `import` is `true`.
     * 
     */
    public Optional<Output<Boolean>> import_() {
        return Optional.ofNullable(this.import_);
    }

    /**
     * The name of the role
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the role
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The realm this role exists within.
     * 
     */
    @Import(name="realmId", required=true)
    private Output<String> realmId;

    /**
     * @return The realm this role exists within.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    private RoleArgs() {}

    private RoleArgs(RoleArgs $) {
        this.attributes = $.attributes;
        this.clientId = $.clientId;
        this.compositeRoles = $.compositeRoles;
        this.description = $.description;
        this.import_ = $.import_;
        this.name = $.name;
        this.realmId = $.realmId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RoleArgs $;

        public Builder() {
            $ = new RoleArgs();
        }

        public Builder(RoleArgs defaults) {
            $ = new RoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param attributes A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
         * 
         * @return builder
         * 
         */
        public Builder attributes(@Nullable Output<Map<String,String>> attributes) {
            $.attributes = attributes;
            return this;
        }

        /**
         * @param attributes A map representing attributes for the role. In order to add multivalue attributes, use `##` to seperate the values. Max length for each value is 255 chars
         * 
         * @return builder
         * 
         */
        public Builder attributes(Map<String,String> attributes) {
            return attributes(Output.of(attributes));
        }

        /**
         * @param clientId When specified, this role will be created as a client role attached to the client with the provided ID
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId When specified, this role will be created as a client role attached to the client with the provided ID
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param compositeRoles When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
         * 
         * @return builder
         * 
         */
        public Builder compositeRoles(@Nullable Output<List<String>> compositeRoles) {
            $.compositeRoles = compositeRoles;
            return this;
        }

        /**
         * @param compositeRoles When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
         * 
         * @return builder
         * 
         */
        public Builder compositeRoles(List<String> compositeRoles) {
            return compositeRoles(Output.of(compositeRoles));
        }

        /**
         * @param compositeRoles When specified, this role will be a composite role, composed of all roles that have an ID present within this list.
         * 
         * @return builder
         * 
         */
        public Builder compositeRoles(String... compositeRoles) {
            return compositeRoles(List.of(compositeRoles));
        }

        /**
         * @param description The description of the role
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the role
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param import_ When `true`, the role with the specified `name` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with roles that Keycloak creates automatically during realm creation, such as the client roles `create-client`, `view-realm`, ... for the client `realm-management` created per realm. Note, that the role will not be removed during destruction if `import` is `true`.
         * 
         * @return builder
         * 
         */
        public Builder import_(@Nullable Output<Boolean> import_) {
            $.import_ = import_;
            return this;
        }

        /**
         * @param import_ When `true`, the role with the specified `name` is assumed to already exist, and it will be imported into state instead of being created. This attribute is useful when dealing with roles that Keycloak creates automatically during realm creation, such as the client roles `create-client`, `view-realm`, ... for the client `realm-management` created per realm. Note, that the role will not be removed during destruction if `import` is `true`.
         * 
         * @return builder
         * 
         */
        public Builder import_(Boolean import_) {
            return import_(Output.of(import_));
        }

        /**
         * @param name The name of the role
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the role
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param realmId The realm this role exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this role exists within.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public RoleArgs build() {
            if ($.realmId == null) {
                throw new MissingRequiredPropertyException("RoleArgs", "realmId");
            }
            return $;
        }
    }

}
