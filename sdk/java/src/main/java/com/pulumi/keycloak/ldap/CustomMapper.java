// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.ldap.CustomMapperArgs;
import com.pulumi.keycloak.ldap.inputs.CustomMapperState;
import java.lang.String;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing custom attribute mappers for Keycloak users federated via LDAP.
 * 
 * The LDAP custom mapper is implemented and deployed into Keycloak as a custom provider. This resource allows to
 * specify the custom id and custom implementation class of the self-implemented attribute mapper as well as additional
 * properties via config map.
 * 
 * The custom mapper should already be deployed into keycloak in order to be correctly configured.
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
 * import com.pulumi.keycloak.ldap.UserFederation;
 * import com.pulumi.keycloak.ldap.UserFederationArgs;
 * import com.pulumi.keycloak.ldap.CustomMapper;
 * import com.pulumi.keycloak.ldap.CustomMapperArgs;
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
 *         var ldapUserFederation = new UserFederation("ldapUserFederation", UserFederationArgs.builder()
 *             .name("openldap")
 *             .realmId(realm.id())
 *             .usernameLdapAttribute("cn")
 *             .rdnLdapAttribute("cn")
 *             .uuidLdapAttribute("entryDN")
 *             .userObjectClasses(            
 *                 "simpleSecurityObject",
 *                 "organizationalRole")
 *             .connectionUrl("ldap://openldap")
 *             .usersDn("dc=example,dc=org")
 *             .bindDn("cn=admin,dc=example,dc=org")
 *             .bindCredential("admin")
 *             .build());
 * 
 *         var customMapper = new CustomMapper("customMapper", CustomMapperArgs.builder()
 *             .name("custom-mapper")
 *             .realmId(openldap.realmId())
 *             .ldapUserFederationId(openldap.id())
 *             .providerId("custom-provider-registered-in-keycloak")
 *             .providerType("com.example.custom.ldap.mappers.CustomMapper")
 *             .config(Map.ofEntries(
 *                 Map.entry("attribute.name", "name"),
 *                 Map.entry("attribute.value", "value")
 *             ))
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
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
 * 
 * The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs.
 * 
 * Example:
 * 
 * bash
 * 
 * ```sh
 * $ pulumi import keycloak:ldap/customMapper:CustomMapper custom_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 * 
 */
@ResourceType(type="keycloak:ldap/customMapper:CustomMapper")
public class CustomMapper extends com.pulumi.resources.CustomResource {
    /**
     * A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
     * 
     */
    @Export(name="config", refs={Map.class,String.class}, tree="[0,1,1]")
    private Output</* @Nullable */ Map<String,String>> config;

    /**
     * @return A map with key / value pairs for configuring the LDAP mapper. The supported keys depend on the protocol mapper.
     * 
     */
    public Output<Optional<Map<String,String>>> config() {
        return Codegen.optional(this.config);
    }
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    @Export(name="ldapUserFederationId", refs={String.class}, tree="[0]")
    private Output<String> ldapUserFederationId;

    /**
     * @return The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    public Output<String> ldapUserFederationId() {
        return this.ldapUserFederationId;
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
     * The id of the LDAP mapper implemented in MapperFactory.
     * 
     */
    @Export(name="providerId", refs={String.class}, tree="[0]")
    private Output<String> providerId;

    /**
     * @return The id of the LDAP mapper implemented in MapperFactory.
     * 
     */
    public Output<String> providerId() {
        return this.providerId;
    }
    /**
     * The fully-qualified Java class name of the custom LDAP mapper.
     * 
     */
    @Export(name="providerType", refs={String.class}, tree="[0]")
    private Output<String> providerType;

    /**
     * @return The fully-qualified Java class name of the custom LDAP mapper.
     * 
     */
    public Output<String> providerType() {
        return this.providerType;
    }
    /**
     * The realm that this LDAP mapper will exist in.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm that this LDAP mapper will exist in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CustomMapper(java.lang.String name) {
        this(name, CustomMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CustomMapper(java.lang.String name, CustomMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CustomMapper(java.lang.String name, CustomMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/customMapper:CustomMapper", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private CustomMapper(java.lang.String name, Output<java.lang.String> id, @Nullable CustomMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/customMapper:CustomMapper", name, state, makeResourceOptions(options, id), false);
    }

    private static CustomMapperArgs makeArgs(CustomMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? CustomMapperArgs.Empty : args;
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
    public static CustomMapper get(java.lang.String name, Output<java.lang.String> id, @Nullable CustomMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CustomMapper(name, id, state, options);
    }
}
