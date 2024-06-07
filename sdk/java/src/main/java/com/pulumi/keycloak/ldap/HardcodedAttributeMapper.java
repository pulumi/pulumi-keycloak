// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.ldap.HardcodedAttributeMapperArgs;
import com.pulumi.keycloak.ldap.inputs.HardcodedAttributeMapperState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing hardcoded attribute mappers for Keycloak users federated via LDAP.
 * 
 * The LDAP hardcoded attribute mapper will set the specified value to the LDAP attribute.
 * 
 * **NOTE**: This mapper only works when the `sync_registrations` attribute on the `keycloak.ldap.UserFederation` resource is set to `true`.
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
 * import com.pulumi.keycloak.ldap.HardcodedAttributeMapper;
 * import com.pulumi.keycloak.ldap.HardcodedAttributeMapperArgs;
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
 *             .syncRegistrations(true)
 *             .build());
 * 
 *         var assignBarToFoo = new HardcodedAttributeMapper("assignBarToFoo", HardcodedAttributeMapperArgs.builder()
 *             .realmId(realm.id())
 *             .ldapUserFederationId(ldapUserFederation.id())
 *             .name("assign-foo-to-bar")
 *             .attributeName("foo")
 *             .attributeValue("bar")
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
 * $ pulumi import keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper assign_bar_to_foo my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 * 
 */
@ResourceType(type="keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper")
public class HardcodedAttributeMapper extends com.pulumi.resources.CustomResource {
    /**
     * The name of the LDAP attribute to set.
     * 
     */
    @Export(name="attributeName", refs={String.class}, tree="[0]")
    private Output<String> attributeName;

    /**
     * @return The name of the LDAP attribute to set.
     * 
     */
    public Output<String> attributeName() {
        return this.attributeName;
    }
    /**
     * The value to set to the LDAP attribute. You can hardcode any value like &#39;foo&#39;.
     * 
     */
    @Export(name="attributeValue", refs={String.class}, tree="[0]")
    private Output<String> attributeValue;

    /**
     * @return The value to set to the LDAP attribute. You can hardcode any value like &#39;foo&#39;.
     * 
     */
    public Output<String> attributeValue() {
        return this.attributeValue;
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
    public HardcodedAttributeMapper(String name) {
        this(name, HardcodedAttributeMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HardcodedAttributeMapper(String name, HardcodedAttributeMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HardcodedAttributeMapper(String name, HardcodedAttributeMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper", name, args == null ? HardcodedAttributeMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private HardcodedAttributeMapper(String name, Output<String> id, @Nullable HardcodedAttributeMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/hardcodedAttributeMapper:HardcodedAttributeMapper", name, state, makeResourceOptions(options, id));
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
    public static HardcodedAttributeMapper get(String name, Output<String> id, @Nullable HardcodedAttributeMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HardcodedAttributeMapper(name, id, state, options);
    }
}
