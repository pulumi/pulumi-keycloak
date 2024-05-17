// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.AttributeToRoleIdentityMapperArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.AttributeToRoleIdentityMapperState;
import java.lang.Object;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing an attribute to role identity provider mapper within Keycloak.
 * 
 * &gt; If you are using Keycloak 10 or higher, you will need to specify the `extra_config` argument in order to define a `syncMode` for the mapper.
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
 * import com.pulumi.keycloak.AttributeToRoleIdentityMapper;
 * import com.pulumi.keycloak.AttributeToRoleIdentityMapperArgs;
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
 *             .alias("oidc")
 *             .authorizationUrl("https://example.com/auth")
 *             .tokenUrl("https://example.com/token")
 *             .clientId("example_id")
 *             .clientSecret("example_token")
 *             .defaultScopes("openid random profile")
 *             .build());
 * 
 *         var realmRole = new Role("realmRole", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .name("my-realm-role")
 *             .description("My Realm Role")
 *             .build());
 * 
 *         var oidcAttributeToRoleIdentityMapper = new AttributeToRoleIdentityMapper("oidcAttributeToRoleIdentityMapper", AttributeToRoleIdentityMapperArgs.builder()
 *             .realm(realm.id())
 *             .name("role-attribute")
 *             .identityProviderAlias(oidc.alias())
 *             .role("my-realm-role")
 *             .claimName("my-claim")
 *             .claimValue("my-value")
 *             .extraConfig(Map.of("syncMode", "INHERIT"))
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
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
 * $ pulumi import keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper test_mapper my-realm/my-mapper/f446db98-7133-4e30-b18a-3d28fde7ca1b
 * ```
 * 
 */
@ResourceType(type="keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper")
public class AttributeToRoleIdentityMapper extends com.pulumi.resources.CustomResource {
    /**
     * Attribute Friendly Name. Conflicts with `attribute_name`.
     * 
     */
    @Export(name="attributeFriendlyName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> attributeFriendlyName;

    /**
     * @return Attribute Friendly Name. Conflicts with `attribute_name`.
     * 
     */
    public Output<Optional<String>> attributeFriendlyName() {
        return Codegen.optional(this.attributeFriendlyName);
    }
    /**
     * Attribute Name.
     * 
     */
    @Export(name="attributeName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> attributeName;

    /**
     * @return Attribute Name.
     * 
     */
    public Output<Optional<String>> attributeName() {
        return Codegen.optional(this.attributeName);
    }
    /**
     * Attribute Value.
     * 
     */
    @Export(name="attributeValue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> attributeValue;

    /**
     * @return Attribute Value.
     * 
     */
    public Output<Optional<String>> attributeValue() {
        return Codegen.optional(this.attributeValue);
    }
    /**
     * OIDC Claim Name
     * 
     */
    @Export(name="claimName", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> claimName;

    /**
     * @return OIDC Claim Name
     * 
     */
    public Output<Optional<String>> claimName() {
        return Codegen.optional(this.claimName);
    }
    /**
     * OIDC Claim Value
     * 
     */
    @Export(name="claimValue", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> claimValue;

    /**
     * @return OIDC Claim Value
     * 
     */
    public Output<Optional<String>> claimValue() {
        return Codegen.optional(this.claimValue);
    }
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
     * Role Name.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return Role Name.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AttributeToRoleIdentityMapper(String name) {
        this(name, AttributeToRoleIdentityMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AttributeToRoleIdentityMapper(String name, AttributeToRoleIdentityMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AttributeToRoleIdentityMapper(String name, AttributeToRoleIdentityMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper", name, args == null ? AttributeToRoleIdentityMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AttributeToRoleIdentityMapper(String name, Output<String> id, @Nullable AttributeToRoleIdentityMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/attributeToRoleIdentityMapper:AttributeToRoleIdentityMapper", name, state, makeResourceOptions(options, id));
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
    public static AttributeToRoleIdentityMapper get(String name, Output<String> id, @Nullable AttributeToRoleIdentityMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AttributeToRoleIdentityMapper(name, id, state, options);
    }
}
