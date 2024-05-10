// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.ldap.RoleMapperArgs;
import com.pulumi.keycloak.ldap.inputs.RoleMapperState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for creating and managing role mappers for Keycloak users federated via LDAP.
 * 
 * The LDAP group mapper can be used to map an LDAP user&#39;s roles from some DN to Keycloak roles.
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
 * import com.pulumi.keycloak.ldap.RoleMapper;
 * import com.pulumi.keycloak.ldap.RoleMapperArgs;
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
 *         var ldapRoleMapper = new RoleMapper("ldapRoleMapper", RoleMapperArgs.builder()        
 *             .realmId(realm.id())
 *             .ldapUserFederationId(ldapUserFederation.id())
 *             .name("role-mapper")
 *             .ldapRolesDn("dc=example,dc=org")
 *             .roleNameLdapAttribute("cn")
 *             .roleObjectClasses("groupOfNames")
 *             .membershipAttributeType("DN")
 *             .membershipLdapAttribute("member")
 *             .membershipUserLdapAttribute("cn")
 *             .userRolesRetrieveStrategy("GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE")
 *             .memberofLdapAttribute("memberOf")
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
 * $ pulumi import keycloak:ldap/roleMapper:RoleMapper ldap_role_mapper my-realm/af2a6ca3-e4d7-49c3-b08b-1b3c70b4b860/3d923ece-1a91-4bf7-adaf-3b82f2a12b67
 * ```
 * 
 */
@ResourceType(type="keycloak:ldap/roleMapper:RoleMapper")
public class RoleMapper extends com.pulumi.resources.CustomResource {
    /**
     * When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `use_realm_roles_mapping` is `false`.
     * 
     */
    @Export(name="clientId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> clientId;

    /**
     * @return When specified, LDAP role mappings will be mapped to client role mappings tied to this client ID. Can only be set if `use_realm_roles_mapping` is `false`.
     * 
     */
    public Output<Optional<String>> clientId() {
        return Codegen.optional(this.clientId);
    }
    /**
     * The LDAP DN where roles can be found.
     * 
     */
    @Export(name="ldapRolesDn", refs={String.class}, tree="[0]")
    private Output<String> ldapRolesDn;

    /**
     * @return The LDAP DN where roles can be found.
     * 
     */
    public Output<String> ldapRolesDn() {
        return this.ldapRolesDn;
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
     * Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
     * 
     */
    @Export(name="memberofLdapAttribute", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> memberofLdapAttribute;

    /**
     * @return Specifies the name of the LDAP attribute on the LDAP user that contains the roles the user has. Defaults to `memberOf`. This is only used when
     * 
     */
    public Output<Optional<String>> memberofLdapAttribute() {
        return Codegen.optional(this.memberofLdapAttribute);
    }
    /**
     * Can be one of `DN` or `UID`. Defaults to `DN`.
     * 
     */
    @Export(name="membershipAttributeType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> membershipAttributeType;

    /**
     * @return Can be one of `DN` or `UID`. Defaults to `DN`.
     * 
     */
    public Output<Optional<String>> membershipAttributeType() {
        return Codegen.optional(this.membershipAttributeType);
    }
    /**
     * The name of the LDAP attribute that is used for membership mappings.
     * 
     */
    @Export(name="membershipLdapAttribute", refs={String.class}, tree="[0]")
    private Output<String> membershipLdapAttribute;

    /**
     * @return The name of the LDAP attribute that is used for membership mappings.
     * 
     */
    public Output<String> membershipLdapAttribute() {
        return this.membershipLdapAttribute;
    }
    /**
     * The name of the LDAP attribute on a user that is used for membership mappings.
     * 
     */
    @Export(name="membershipUserLdapAttribute", refs={String.class}, tree="[0]")
    private Output<String> membershipUserLdapAttribute;

    /**
     * @return The name of the LDAP attribute on a user that is used for membership mappings.
     * 
     */
    public Output<String> membershipUserLdapAttribute() {
        return this.membershipUserLdapAttribute;
    }
    /**
     * Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
     * 
     */
    @Export(name="mode", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> mode;

    /**
     * @return Can be one of `READ_ONLY`, `LDAP_ONLY` or `IMPORT`. Defaults to `READ_ONLY`.
     * 
     */
    public Output<Optional<String>> mode() {
        return Codegen.optional(this.mode);
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
     * The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
     * 
     */
    @Export(name="roleNameLdapAttribute", refs={String.class}, tree="[0]")
    private Output<String> roleNameLdapAttribute;

    /**
     * @return The name of the LDAP attribute that is used in role objects for the name and RDN of the role. Typically `cn`.
     * 
     */
    public Output<String> roleNameLdapAttribute() {
        return this.roleNameLdapAttribute;
    }
    /**
     * List of strings representing the object classes for the role. Must contain at least one.
     * 
     */
    @Export(name="roleObjectClasses", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> roleObjectClasses;

    /**
     * @return List of strings representing the object classes for the role. Must contain at least one.
     * 
     */
    public Output<List<String>> roleObjectClasses() {
        return this.roleObjectClasses;
    }
    /**
     * When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
     * 
     */
    @Export(name="rolesLdapFilter", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> rolesLdapFilter;

    /**
     * @return When specified, adds an additional custom filter to be used when querying for roles. Must start with `(` and end with `)`.
     * 
     */
    public Output<Optional<String>> rolesLdapFilter() {
        return Codegen.optional(this.rolesLdapFilter);
    }
    /**
     * When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
     * 
     */
    @Export(name="useRealmRolesMapping", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> useRealmRolesMapping;

    /**
     * @return When `true`, LDAP role mappings will be mapped to realm roles within Keycloak. Defaults to `true`.
     * 
     */
    public Output<Optional<Boolean>> useRealmRolesMapping() {
        return Codegen.optional(this.useRealmRolesMapping);
    }
    /**
     * Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
     * 
     */
    @Export(name="userRolesRetrieveStrategy", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> userRolesRetrieveStrategy;

    /**
     * @return Can be one of `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`, `GET_ROLES_FROM_USER_MEMBEROF_ATTRIBUTE`, or `LOAD_ROLES_BY_MEMBER_ATTRIBUTE_RECURSIVELY`. Defaults to `LOAD_ROLES_BY_MEMBER_ATTRIBUTE`.
     * 
     */
    public Output<Optional<String>> userRolesRetrieveStrategy() {
        return Codegen.optional(this.userRolesRetrieveStrategy);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RoleMapper(String name) {
        this(name, RoleMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RoleMapper(String name, RoleMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RoleMapper(String name, RoleMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/roleMapper:RoleMapper", name, args == null ? RoleMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RoleMapper(String name, Output<String> id, @Nullable RoleMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/roleMapper:RoleMapper", name, state, makeResourceOptions(options, id));
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
    public static RoleMapper get(String name, Output<String> id, @Nullable RoleMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RoleMapper(name, id, state, options);
    }
}
