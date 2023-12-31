// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.HardcodedRoleIdentityMapperArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.HardcodedRoleIdentityMapperState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing hardcoded role mappers for Keycloak identity provider.
 * 
 * The identity provider hardcoded role mapper grants a specified Keycloak role to each Keycloak user from the LDAP provider.
 * 
 * ## Example Usage
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
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
 * import com.pulumi.keycloak.HardcodedRoleIdentityMapper;
 * import com.pulumi.keycloak.HardcodedRoleIdentityMapperArgs;
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
 *             .alias(&#34;my-idp&#34;)
 *             .authorizationUrl(&#34;https://authorizationurl.com&#34;)
 *             .clientId(&#34;clientID&#34;)
 *             .clientSecret(&#34;clientSecret&#34;)
 *             .tokenUrl(&#34;https://tokenurl.com&#34;)
 *             .build());
 * 
 *         var realmRole = new Role(&#34;realmRole&#34;, RoleArgs.builder()        
 *             .realmId(realm.id())
 *             .description(&#34;My Realm Role&#34;)
 *             .build());
 * 
 *         var oidcHardcodedRoleIdentityMapper = new HardcodedRoleIdentityMapper(&#34;oidcHardcodedRoleIdentityMapper&#34;, HardcodedRoleIdentityMapperArgs.builder()        
 *             .realm(realm.id())
 *             .identityProviderAlias(oidcIdentityProvider.alias())
 *             .role(&#34;my-realm-role&#34;)
 *             .extraConfig(Map.of(&#34;syncMode&#34;, &#34;INHERIT&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper")
public class HardcodedRoleIdentityMapper extends com.pulumi.resources.CustomResource {
    @Export(name="extraConfig", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> extraConfig;

    public Output<Optional<Map<String,Object>>> extraConfig() {
        return Codegen.optional(this.extraConfig);
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
     * The name of the role which should be assigned to the users.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> role;

    /**
     * @return The name of the role which should be assigned to the users.
     * 
     */
    public Output<Optional<String>> role() {
        return Codegen.optional(this.role);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HardcodedRoleIdentityMapper(String name) {
        this(name, HardcodedRoleIdentityMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HardcodedRoleIdentityMapper(String name, HardcodedRoleIdentityMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HardcodedRoleIdentityMapper(String name, HardcodedRoleIdentityMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, args == null ? HardcodedRoleIdentityMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HardcodedRoleIdentityMapper(String name, Output<String> id, @Nullable HardcodedRoleIdentityMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, state, makeResourceOptions(options, id));
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
    public static HardcodedRoleIdentityMapper get(String name, Output<String> id, @Nullable HardcodedRoleIdentityMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HardcodedRoleIdentityMapper(name, id, state, options);
    }
}
