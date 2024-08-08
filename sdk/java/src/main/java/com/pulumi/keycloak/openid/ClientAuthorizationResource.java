// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.openid.ClientAuthorizationResourceArgs;
import com.pulumi.keycloak.openid.inputs.ClientAuthorizationResourceState;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

@ResourceType(type="keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource")
public class ClientAuthorizationResource extends com.pulumi.resources.CustomResource {
    @Export(name="attributes", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> attributes;

    public Output<Optional<Map<String,Object>>> attributes() {
        return Codegen.optional(this.attributes);
    }
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> displayName;

    public Output<Optional<String>> displayName() {
        return Codegen.optional(this.displayName);
    }
    @Export(name="iconUri", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> iconUri;

    public Output<Optional<String>> iconUri() {
        return Codegen.optional(this.iconUri);
    }
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }
    @Export(name="ownerManagedAccess", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ownerManagedAccess;

    public Output<Optional<Boolean>> ownerManagedAccess() {
        return Codegen.optional(this.ownerManagedAccess);
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
    @Export(name="scopes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> scopes;

    public Output<Optional<List<String>>> scopes() {
        return Codegen.optional(this.scopes);
    }
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> type;

    public Output<Optional<String>> type() {
        return Codegen.optional(this.type);
    }
    @Export(name="uris", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> uris;

    public Output<Optional<List<String>>> uris() {
        return Codegen.optional(this.uris);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public ClientAuthorizationResource(java.lang.String name) {
        this(name, ClientAuthorizationResourceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public ClientAuthorizationResource(java.lang.String name, ClientAuthorizationResourceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public ClientAuthorizationResource(java.lang.String name, ClientAuthorizationResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private ClientAuthorizationResource(java.lang.String name, Output<java.lang.String> id, @Nullable ClientAuthorizationResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:openid/clientAuthorizationResource:ClientAuthorizationResource", name, state, makeResourceOptions(options, id), false);
    }

    private static ClientAuthorizationResourceArgs makeArgs(ClientAuthorizationResourceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? ClientAuthorizationResourceArgs.Empty : args;
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
    public static ClientAuthorizationResource get(java.lang.String name, Output<java.lang.String> id, @Nullable ClientAuthorizationResourceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new ClientAuthorizationResource(name, id, state, options);
    }
}
