// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.CustomUserFederationArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.CustomUserFederationState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.CustomUserFederation
 * 
 * Allows for creating and managing custom user federation providers within Keycloak.
 * 
 * A custom user federation provider is an implementation of Keycloak&#39;s
 * [User Storage SPI](https://www.keycloak.org/docs/4.2/server_development/index.html#_user-storage-spi).
 * An example of this implementation can be found here.
 * 
 * ### Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.CustomUserFederation;
 * import com.pulumi.keycloak.CustomUserFederationArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var realm = new Realm("realm", RealmArgs.builder()        
 *             .realm("test")
 *             .enabled(true)
 *             .build());
 * 
 *         var customUserFederation = new CustomUserFederation("customUserFederation", CustomUserFederationArgs.builder()        
 *             .name("custom")
 *             .realmId(realm.id())
 *             .providerId("custom")
 *             .enabled(true)
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Argument Reference
 * 
 * The following arguments are supported:
 * 
 * - `realm_id` - (Required) The realm that this provider will provide user federation for.
 * - `name` - (Required) Display name of the provider when displayed in the console.
 * - `provider_id` - (Required) The unique ID of the custom provider, specified in the `getId` implementation for the `UserStorageProviderFactory` interface.
 * - `enabled` - (Optional) When `false`, this provider will not be used when performing queries for users. Defaults to `true`.
 * - `priority` - (Optional) Priority of this provider when looking up users. Lower values are first. Defaults to `0`.
 * - `cache_policy` - (Optional) Can be one of `DEFAULT`, `EVICT_DAILY`, `EVICT_WEEKLY`, `MAX_LIFESPAN`, or `NO_CACHE`. Defaults to `DEFAULT`.
 * 
 * ### Import
 * 
 * Custom user federation providers can be imported using the format `{{realm_id}}/{{custom_user_federation_id}}`.
 * The ID of the custom user federation provider can be found within the Keycloak GUI and is typically a GUID:
 * 
 */
@ResourceType(type="keycloak:index/customUserFederation:CustomUserFederation")
public class CustomUserFederation extends com.pulumi.resources.CustomResource {
    @Export(name="cachePolicy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> cachePolicy;

    public Output<Optional<String>> cachePolicy() {
        return Codegen.optional(this.cachePolicy);
    }
    /**
     * How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users
     * sync.
     * 
     */
    @Export(name="changedSyncPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> changedSyncPeriod;

    /**
     * @return How frequently Keycloak should sync changed users, in seconds. Omit this property to disable periodic changed users
     * sync.
     * 
     */
    public Output<Optional<Integer>> changedSyncPeriod() {
        return Codegen.optional(this.changedSyncPeriod);
    }
    @Export(name="config", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> config;

    public Output<Optional<Map<String,Object>>> config() {
        return Codegen.optional(this.config);
    }
    /**
     * When false, this provider will not be used when performing queries for users.
     * 
     */
    @Export(name="enabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> enabled;

    /**
     * @return When false, this provider will not be used when performing queries for users.
     * 
     */
    public Output<Optional<Boolean>> enabled() {
        return Codegen.optional(this.enabled);
    }
    /**
     * How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
     * 
     */
    @Export(name="fullSyncPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> fullSyncPeriod;

    /**
     * @return How frequently Keycloak should sync all users, in seconds. Omit this property to disable periodic full sync.
     * 
     */
    public Output<Optional<Integer>> fullSyncPeriod() {
        return Codegen.optional(this.fullSyncPeriod);
    }
    /**
     * Display name of the provider when displayed in the console.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name of the provider when displayed in the console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The parent_id of the generated component. will use realm_id if not specified.
     * 
     */
    @Export(name="parentId", refs={String.class}, tree="[0]")
    private Output<String> parentId;

    /**
     * @return The parent_id of the generated component. will use realm_id if not specified.
     * 
     */
    public Output<String> parentId() {
        return this.parentId;
    }
    /**
     * Priority of this provider when looking up users. Lower values are first.
     * 
     */
    @Export(name="priority", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> priority;

    /**
     * @return Priority of this provider when looking up users. Lower values are first.
     * 
     */
    public Output<Optional<Integer>> priority() {
        return Codegen.optional(this.priority);
    }
    /**
     * The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
     * interface
     * 
     */
    @Export(name="providerId", refs={String.class}, tree="[0]")
    private Output<String> providerId;

    /**
     * @return The unique ID of the custom provider, specified in the `getId` implementation for the UserStorageProviderFactory
     * interface
     * 
     */
    public Output<String> providerId() {
        return this.providerId;
    }
    /**
     * The realm (name) this provider will provide user federation for.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm (name) this provider will provide user federation for.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomUserFederation(String name) {
        this(name, CustomUserFederationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomUserFederation(String name, CustomUserFederationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomUserFederation(String name, CustomUserFederationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/customUserFederation:CustomUserFederation", name, args == null ? CustomUserFederationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CustomUserFederation(String name, Output<String> id, @Nullable CustomUserFederationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/customUserFederation:CustomUserFederation", name, state, makeResourceOptions(options, id));
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
    public static CustomUserFederation get(String name, Output<String> id, @Nullable CustomUserFederationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomUserFederation(name, id, state, options);
    }
}
