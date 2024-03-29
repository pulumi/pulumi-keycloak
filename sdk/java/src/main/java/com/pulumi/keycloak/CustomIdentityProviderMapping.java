// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.CustomIdentityProviderMappingArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.CustomIdentityProviderMappingState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.keycloak.Realm;
 * import com.pulumi.keycloak.RealmArgs;
 * import com.pulumi.keycloak.oidc.IdentityProvider;
 * import com.pulumi.keycloak.oidc.IdentityProviderArgs;
 * import com.pulumi.keycloak.CustomIdentityProviderMapping;
 * import com.pulumi.keycloak.CustomIdentityProviderMappingArgs;
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
 *         var realm = new Realm(&#34;realm&#34;, RealmArgs.builder()        
 *             .realm(&#34;my-realm&#34;)
 *             .enabled(true)
 *             .build());
 * 
 *         var oidcIdentityProvider = new IdentityProvider(&#34;oidcIdentityProvider&#34;, IdentityProviderArgs.builder()        
 *             .realm(realm.id())
 *             .alias(&#34;oidc&#34;)
 *             .authorizationUrl(&#34;https://example.com/auth&#34;)
 *             .tokenUrl(&#34;https://example.com/token&#34;)
 *             .clientId(&#34;example_id&#34;)
 *             .clientSecret(&#34;example_token&#34;)
 *             .defaultScopes(&#34;openid random profile&#34;)
 *             .build());
 * 
 *         var oidcCustomIdentityProviderMapping = new CustomIdentityProviderMapping(&#34;oidcCustomIdentityProviderMapping&#34;, CustomIdentityProviderMappingArgs.builder()        
 *             .realm(realm.id())
 *             .identityProviderAlias(oidcIdentityProvider.alias())
 *             .identityProviderMapper(&#34;%s-user-attribute-idp-mapper&#34;)
 *             .extraConfig(Map.ofEntries(
 *                 Map.entry(&#34;syncMode&#34;, &#34;INHERIT&#34;),
 *                 Map.entry(&#34;Claim&#34;, &#34;my-email-claim&#34;),
 *                 Map.entry(&#34;UserAttribute&#34;, &#34;email&#34;)
 *             ))
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Identity provider mappers can be imported using the format `{{realm_id}}/{{idp_alias}}/{{idp_mapper_id}}`, where `idp_alias` is the identity provider alias, and `idp_mapper_id` is the unique ID that Keycloak
 * 
 * assigns to the mapper upon creation. This value can be found in the URI when editing this mapper in the GUI, and is typically a GUID.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
 * ```
 * 
 */
@ResourceType(type="keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping")
public class CustomIdentityProviderMapping extends com.pulumi.resources.CustomResource {
    /**
     * Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     * 
     */
    @Export(name="extraConfig", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> extraConfig;

    /**
     * @return Key/value attributes to add to the identity provider mapper model that is persisted to Keycloak. This can be used to extend the base model with new Keycloak features.
     * 
     */
    public Output<Optional<Map<String,Object>>> extraConfig() {
        return Codegen.optional(this.extraConfig);
    }
    /**
     * The alias of the associated identity provider.
     * 
     */
    @Export(name="identityProviderAlias", refs={String.class}, tree="[0]")
    private Output<String> identityProviderAlias;

    /**
     * @return The alias of the associated identity provider.
     * 
     */
    public Output<String> identityProviderAlias() {
        return this.identityProviderAlias;
    }
    /**
     * The type of the identity provider mapper. This can be a format string that includes a `%!s(MISSING)` - this will be replaced by the provider id.
     * 
     */
    @Export(name="identityProviderMapper", refs={String.class}, tree="[0]")
    private Output<String> identityProviderMapper;

    /**
     * @return The type of the identity provider mapper. This can be a format string that includes a `%!s(MISSING)` - this will be replaced by the provider id.
     * 
     */
    public Output<String> identityProviderMapper() {
        return this.identityProviderMapper;
    }
    /**
     * The name of the mapper.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the mapper.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The name of the realm.
     * 
     */
    @Export(name="realm", refs={String.class}, tree="[0]")
    private Output<String> realm;

    /**
     * @return The name of the realm.
     * 
     */
    public Output<String> realm() {
        return this.realm;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomIdentityProviderMapping(String name) {
        this(name, CustomIdentityProviderMappingArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomIdentityProviderMapping(String name, CustomIdentityProviderMappingArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomIdentityProviderMapping(String name, CustomIdentityProviderMappingArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping", name, args == null ? CustomIdentityProviderMappingArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CustomIdentityProviderMapping(String name, Output<String> id, @Nullable CustomIdentityProviderMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/customIdentityProviderMapping:CustomIdentityProviderMapping", name, state, makeResourceOptions(options, id));
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
    public static CustomIdentityProviderMapping get(String name, Output<String> id, @Nullable CustomIdentityProviderMappingState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomIdentityProviderMapping(name, id, state, options);
    }
}
