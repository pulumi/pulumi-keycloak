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
 *         var realm = new Realm("realm", RealmArgs.builder()
 *             .realm("my-realm")
 *             .enabled(true)
 *             .build());
 * 
 *         var oidc = new IdentityProvider("oidc", IdentityProviderArgs.builder()
 *             .realm(realm.id())
 *             .alias("my-idp")
 *             .authorizationUrl("https://authorizationurl.com")
 *             .clientId("clientID")
 *             .clientSecret("clientSecret")
 *             .tokenUrl("https://tokenurl.com")
 *             .build());
 * 
 *         var realmRole = new Role("realmRole", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .name("my-realm-role")
 *             .description("My Realm Role")
 *             .build());
 * 
 *         var oidcHardcodedRoleIdentityMapper = new HardcodedRoleIdentityMapper("oidcHardcodedRoleIdentityMapper", HardcodedRoleIdentityMapperArgs.builder()
 *             .realm(realm.id())
 *             .name("hardcodedRole")
 *             .identityProviderAlias(oidc.alias())
 *             .role("my-realm-role")
 *             .extraConfig(Map.of("syncMode", "INHERIT"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 */
@ResourceType(type="keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper")
public class HardcodedRoleIdentityMapper extends com.pulumi.resources.CustomResource {
    @Export(name="extraConfig", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> extraConfig;

    public Output<Optional<Map<String,String>>> extraConfig() {
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
    public HardcodedRoleIdentityMapper(java.lang.String name) {
        this(name, HardcodedRoleIdentityMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HardcodedRoleIdentityMapper(java.lang.String name, HardcodedRoleIdentityMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HardcodedRoleIdentityMapper(java.lang.String name, HardcodedRoleIdentityMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HardcodedRoleIdentityMapper(java.lang.String name, Output<java.lang.String> id, @Nullable HardcodedRoleIdentityMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/hardcodedRoleIdentityMapper:HardcodedRoleIdentityMapper", name, state, makeResourceOptions(options, id), false);
    }

    private static HardcodedRoleIdentityMapperArgs makeArgs(HardcodedRoleIdentityMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HardcodedRoleIdentityMapperArgs.Empty : args;
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
    public static HardcodedRoleIdentityMapper get(java.lang.String name, Output<java.lang.String> id, @Nullable HardcodedRoleIdentityMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HardcodedRoleIdentityMapper(name, id, state, options);
    }
}
