// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.ldap.HardcodedRoleMapperArgs;
import com.pulumi.keycloak.ldap.inputs.HardcodedRoleMapperState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing hardcoded role mappers for Keycloak users federated via LDAP.
 * 
 * The LDAP hardcoded role mapper will grant a specified Keycloak role to each Keycloak user linked with LDAP.
 * 
 * ## Example Usage
 * 
 * ### Realm Role)
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
 * import com.pulumi.keycloak.Role;
 * import com.pulumi.keycloak.RoleArgs;
 * import com.pulumi.keycloak.ldap.HardcodedRoleMapper;
 * import com.pulumi.keycloak.ldap.HardcodedRoleMapperArgs;
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
 *         var realmAdminRole = new Role("realmAdminRole", RoleArgs.builder()
 *             .realmId(realm.id())
 *             .name("my-admin-role")
 *             .description("My Realm Role")
 *             .build());
 * 
 *         var assignAdminRoleToAllUsers = new HardcodedRoleMapper("assignAdminRoleToAllUsers", HardcodedRoleMapperArgs.builder()
 *             .realmId(realm.id())
 *             .ldapUserFederationId(ldapUserFederation.id())
 *             .name("assign-admin-role-to-all-users")
 *             .role(realmAdminRole.name())
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ### Client Role)
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
 * import com.pulumi.keycloak.openid.OpenidFunctions;
 * import com.pulumi.keycloak.openid.inputs.GetClientArgs;
 * import com.pulumi.keycloak.KeycloakFunctions;
 * import com.pulumi.keycloak.inputs.GetRoleArgs;
 * import com.pulumi.keycloak.ldap.HardcodedRoleMapper;
 * import com.pulumi.keycloak.ldap.HardcodedRoleMapperArgs;
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
 *         // data sources aren't technically necessary here, but they are helpful for demonstration purposes
 *         final var realmManagement = OpenidFunctions.getClient(GetClientArgs.builder()
 *             .realmId(realm.id())
 *             .clientId("realm-management")
 *             .build());
 * 
 *         final var createClient = Output.tuple(realm.id(), realmManagement).applyValue(values -> {
 *             var id = values.t1;
 *             var realmManagement = values.t2;
 *             return KeycloakFunctions.getRole(GetRoleArgs.builder()
 *                 .realmId(id)
 *                 .clientId(realmManagement.id())
 *                 .name("create-client")
 *                 .build());
 *         });
 * 
 *         var assignAdminRoleToAllUsers = new HardcodedRoleMapper("assignAdminRoleToAllUsers", HardcodedRoleMapperArgs.builder()
 *             .realmId(realm.id())
 *             .ldapUserFederationId(ldapUserFederation.id())
 *             .name("assign-admin-role-to-all-users")
 *             .role(Output.tuple(realmManagement, createClient).applyValue(values -> {
 *                 var realmManagement = values.t1;
 *                 var createClient = values.t2;
 *                 return String.format("%s.%s", realmManagement.clientId(),createClient.name());
 *             }))
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
 * $ pulumi import keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper assign_admin_role_to_all_users my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 * 
 */
@ResourceType(type="keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper")
public class HardcodedRoleMapper extends com.pulumi.resources.CustomResource {
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
     * The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The name of the role which should be assigned to the users. Client roles should use the format `{{client_id}}.{{client_role_name}}`.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public HardcodedRoleMapper(java.lang.String name) {
        this(name, HardcodedRoleMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public HardcodedRoleMapper(java.lang.String name, HardcodedRoleMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public HardcodedRoleMapper(java.lang.String name, HardcodedRoleMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private HardcodedRoleMapper(java.lang.String name, Output<java.lang.String> id, @Nullable HardcodedRoleMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/hardcodedRoleMapper:HardcodedRoleMapper", name, state, makeResourceOptions(options, id), false);
    }

    private static HardcodedRoleMapperArgs makeArgs(HardcodedRoleMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? HardcodedRoleMapperArgs.Empty : args;
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
    public static HardcodedRoleMapper get(java.lang.String name, Output<java.lang.String> id, @Nullable HardcodedRoleMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new HardcodedRoleMapper(name, id, state, options);
    }
}
