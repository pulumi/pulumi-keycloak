// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.HardcodedGroupIdentityProviderMapperArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.HardcodedGroupIdentityProviderMapperState;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing hardcoded group mappers for Keycloak identity provider.
 * 
 * The identity provider hardcoded group mapper grants a specified Keycloak group to each Keycloak user from the identity provider.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="keycloak:index/hardcodedGroupIdentityProviderMapper:HardcodedGroupIdentityProviderMapper")
public class HardcodedGroupIdentityProviderMapper extends com.pulumi.resources.CustomResource {
    @Export(name="extraConfig", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> extraConfig;

    public Output<Optional<Map<String,String>>> extraConfig() {
        return Codegen.optional(this.extraConfig);
    }
    /**
     * The name of the group which should be assigned to the users.
     * 
     */
    @Export(name="group", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> group;

    /**
     * @return The name of the group which should be assigned to the users.
     * 
     */
    public Output<Optional<String>> group() {
        return Codegen.optional(this.group);
    }
    /**
     * The IDP alias of the attribute to set.
     * 
     */
    @Export(name="identityProviderAlias", refs={String.class}, tree="[0]")
    private Output<String> identityProviderAlias;

    /**
     * @return The IDP alias of the attribute to set.
     * 
     */
    public Output<String> identityProviderAlias() {
        return this.identityProviderAlias;
    }
    /**
     * Display name of this mapper when displayed in the console.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name of this mapper when displayed in the console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The realm ID that this mapper will exist in.
     * 
     */
    @Export(name="realm", refs={String.class}, tree="[0]")
    private Output<String> realm;

    /**
     * @return The realm ID that this mapper will exist in.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HardcodedGroupIdentityProviderMapper(java.lang.String name) {
        this(name, HardcodedGroupIdentityProviderMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HardcodedGroupIdentityProviderMapper(java.lang.String name, HardcodedGroupIdentityProviderMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HardcodedGroupIdentityProviderMapper(java.lang.String name, HardcodedGroupIdentityProviderMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/hardcodedGroupIdentityProviderMapper:HardcodedGroupIdentityProviderMapper", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HardcodedGroupIdentityProviderMapper(java.lang.String name, Output<java.lang.String> id, @Nullable HardcodedGroupIdentityProviderMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/hardcodedGroupIdentityProviderMapper:HardcodedGroupIdentityProviderMapper", name, state, makeResourceOptions(options, id), false);
    }

    private static HardcodedGroupIdentityProviderMapperArgs makeArgs(HardcodedGroupIdentityProviderMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HardcodedGroupIdentityProviderMapperArgs.Empty : args;
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<java.lang.String> id) {
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
    public static HardcodedGroupIdentityProviderMapper get(java.lang.String name, Output<java.lang.String> id, @Nullable HardcodedGroupIdentityProviderMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HardcodedGroupIdentityProviderMapper(name, id, state, options);
    }
}
