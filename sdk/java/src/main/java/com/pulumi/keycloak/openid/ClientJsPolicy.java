// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.ClientJsPolicyArgs;
import com.pulumi.keycloak.openid.inputs.ClientJsPolicyState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="keycloak:openid/clientJsPolicy:ClientJsPolicy")
public class ClientJsPolicy extends com.pulumi.resources.CustomResource {
    @Export(name="code", refs={String.class}, tree="[0]")
    private Output<String> code;

    public Output<String> code() {
        return this.code;
    }
    @Export(name="decisionStrategy", refs={String.class}, tree="[0]")
    private Output<String> decisionStrategy;

    public Output<String> decisionStrategy() {
        return this.decisionStrategy;
    }
    @Export(name="description", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> description;

    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    @Export(name="logic", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> logic;

    public Output<Optional<String>> logic() {
        return Codegen.optional(this.logic);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }
    @Export(name="resourceServerId", refs={String.class}, tree="[0]")
    private Output<String> resourceServerId;

    public Output<String> resourceServerId() {
        return this.resourceServerId;
    }
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClientJsPolicy(String name) {
        this(name, ClientJsPolicyArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClientJsPolicy(String name, ClientJsPolicyArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClientJsPolicy(String name, ClientJsPolicyArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientJsPolicy:ClientJsPolicy", name, args == null ? ClientJsPolicyArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private ClientJsPolicy(String name, Output<String> id, @Nullable ClientJsPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientJsPolicy:ClientJsPolicy", name, state, makeResourceOptions(options, id));
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
    public static ClientJsPolicy get(String name, Output<String> id, @Nullable ClientJsPolicyState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClientJsPolicy(name, id, state, options);
    }
}
