// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.ldap.UserAttributeMapperArgs;
import com.pulumi.keycloak.ldap.inputs.UserAttributeMapperState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing user attribute mappers for Keycloak users
 * federated via LDAP.
 * 
 * The LDAP user attribute mapper can be used to map a single LDAP attribute
 * to an attribute on the Keycloak user model.
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
 * import com.pulumi.keycloak.ldap.UserFederation;
 * import com.pulumi.keycloak.ldap.UserFederationArgs;
 * import com.pulumi.keycloak.ldap.UserAttributeMapper;
 * import com.pulumi.keycloak.ldap.UserAttributeMapperArgs;
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
 *         var ldapUserFederation = new UserFederation(&#34;ldapUserFederation&#34;, UserFederationArgs.builder()        
 *             .realmId(realm.id())
 *             .usernameLdapAttribute(&#34;cn&#34;)
 *             .rdnLdapAttribute(&#34;cn&#34;)
 *             .uuidLdapAttribute(&#34;entryDN&#34;)
 *             .userObjectClasses(            
 *                 &#34;simpleSecurityObject&#34;,
 *                 &#34;organizationalRole&#34;)
 *             .connectionUrl(&#34;ldap://openldap&#34;)
 *             .usersDn(&#34;dc=example,dc=org&#34;)
 *             .bindDn(&#34;cn=admin,dc=example,dc=org&#34;)
 *             .bindCredential(&#34;admin&#34;)
 *             .build());
 * 
 *         var ldapUserAttributeMapper = new UserAttributeMapper(&#34;ldapUserAttributeMapper&#34;, UserAttributeMapperArgs.builder()        
 *             .realmId(realm.id())
 *             .ldapUserFederationId(ldapUserFederation.id())
 *             .userModelAttribute(&#34;foo&#34;)
 *             .ldapAttribute(&#34;bar&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`. The ID of the LDAP user federation provider and the mapper can be found within the Keycloak GUI, and they are typically GUIDs. Examplebash
 * 
 * ```sh
 *  $ pulumi import keycloak:ldap/userAttributeMapper:UserAttributeMapper ldap_user_attribute_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 * 
 */
@ResourceType(type="keycloak:ldap/userAttributeMapper:UserAttributeMapper")
public class UserAttributeMapper extends com.pulumi.resources.CustomResource {
    /**
     * When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
     * 
     */
    @Export(name="alwaysReadValueFromLdap", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> alwaysReadValueFromLdap;

    /**
     * @return When `true`, the value fetched from LDAP will override the value stored in Keycloak. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> alwaysReadValueFromLdap() {
        return Codegen.optional(this.alwaysReadValueFromLdap);
    }
    /**
     * When `true`, this attribute must exist in LDAP. Defaults to `false`.
     * 
     */
    @Export(name="isMandatoryInLdap", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> isMandatoryInLdap;

    /**
     * @return When `true`, this attribute must exist in LDAP. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> isMandatoryInLdap() {
        return Codegen.optional(this.isMandatoryInLdap);
    }
    /**
     * Name of the mapped attribute on the LDAP object.
     * 
     */
    @Export(name="ldapAttribute", type=String.class, parameters={})
    private Output<String> ldapAttribute;

    /**
     * @return Name of the mapped attribute on the LDAP object.
     * 
     */
    public Output<String> ldapAttribute() {
        return this.ldapAttribute;
    }
    /**
     * The ID of the LDAP user federation provider to attach this mapper to.
     * 
     */
    @Export(name="ldapUserFederationId", type=String.class, parameters={})
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
    @Export(name="name", type=String.class, parameters={})
    private Output<String> name;

    /**
     * @return Display name of this mapper when displayed in the console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
     * 
     */
    @Export(name="readOnly", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> readOnly;

    /**
     * @return When `true`, this attribute is not saved back to LDAP when the user attribute is updated in Keycloak. Defaults to `false`.
     * 
     */
    public Output<Optional<Boolean>> readOnly() {
        return Codegen.optional(this.readOnly);
    }
    /**
     * The realm that this LDAP mapper will exist in.
     * 
     */
    @Export(name="realmId", type=String.class, parameters={})
    private Output<String> realmId;

    /**
     * @return The realm that this LDAP mapper will exist in.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }
    /**
     * Name of the user property or attribute you want to map the LDAP attribute into.
     * 
     */
    @Export(name="userModelAttribute", type=String.class, parameters={})
    private Output<String> userModelAttribute;

    /**
     * @return Name of the user property or attribute you want to map the LDAP attribute into.
     * 
     */
    public Output<String> userModelAttribute() {
        return this.userModelAttribute;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public UserAttributeMapper(String name) {
        this(name, UserAttributeMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public UserAttributeMapper(String name, UserAttributeMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public UserAttributeMapper(String name, UserAttributeMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/userAttributeMapper:UserAttributeMapper", name, args == null ? UserAttributeMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private UserAttributeMapper(String name, Output<String> id, @Nullable UserAttributeMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/userAttributeMapper:UserAttributeMapper", name, state, makeResourceOptions(options, id));
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
    public static UserAttributeMapper get(String name, Output<String> id, @Nullable UserAttributeMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new UserAttributeMapper(name, id, state, options);
    }
}