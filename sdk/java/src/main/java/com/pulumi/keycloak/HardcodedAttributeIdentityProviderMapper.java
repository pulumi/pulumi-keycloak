// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.HardcodedAttributeIdentityProviderMapperArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.HardcodedAttributeIdentityProviderMapperState;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper")
public class HardcodedAttributeIdentityProviderMapper extends com.pulumi.resources.CustomResource {
    /**
     * OIDC Claim
     * 
     */
    @Export(name="attributeName", type=String.class, parameters={})
    private Output</* @Nullable */ String> attributeName;

    /**
     * @return OIDC Claim
     * 
     */
    public Output<Optional<String>> attributeName() {
        return Codegen.optional(this.attributeName);
    }
    /**
     * User Attribute
     * 
     */
    @Export(name="attributeValue", type=String.class, parameters={})
    private Output</* @Nullable */ String> attributeValue;

    /**
     * @return User Attribute
     * 
     */
    public Output<Optional<String>> attributeValue() {
        return Codegen.optional(this.attributeValue);
    }
    @Export(name="extraConfig", type=Map.class, parameters={String.class, Object.class})
    private Output</* @Nullable */ Map<String,Object>> extraConfig;

    public Output<Optional<Map<String,Object>>> extraConfig() {
        return Codegen.optional(this.extraConfig);
    }
    /**
     * IDP Alias
     * 
     */
    @Export(name="identityProviderAlias", type=String.class, parameters={})
    private Output<String> identityProviderAlias;

    /**
     * @return IDP Alias
     * 
     */
    public Output<String> identityProviderAlias() {
        return this.identityProviderAlias;
    }
    /**
     * IDP Mapper Name
     * 
     */
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return IDP Mapper Name
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Realm Name
     * 
     */
    @Export(name="realm", type=String.class, parameters={})
    private Output<String> realm;

    /**
     * @return Realm Name
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }
    /**
     * Is Attribute Related To a User Session
     * 
     */
    @Export(name="userSession", type=Boolean.class, parameters={})
    private Output<Boolean> userSession;

    /**
     * @return Is Attribute Related To a User Session
     * 
     */
    public Output<Boolean> userSession() {
        return this.userSession;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HardcodedAttributeIdentityProviderMapper(String name) {
        this(name, HardcodedAttributeIdentityProviderMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HardcodedAttributeIdentityProviderMapper(String name, HardcodedAttributeIdentityProviderMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HardcodedAttributeIdentityProviderMapper(String name, HardcodedAttributeIdentityProviderMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper", name, args == null ? HardcodedAttributeIdentityProviderMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HardcodedAttributeIdentityProviderMapper(String name, Output<String> id, @Nullable HardcodedAttributeIdentityProviderMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/hardcodedAttributeIdentityProviderMapper:HardcodedAttributeIdentityProviderMapper", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static HardcodedAttributeIdentityProviderMapper get(String name, Output<String> id, @Nullable HardcodedAttributeIdentityProviderMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HardcodedAttributeIdentityProviderMapper(name, id, state, options);
    }
}