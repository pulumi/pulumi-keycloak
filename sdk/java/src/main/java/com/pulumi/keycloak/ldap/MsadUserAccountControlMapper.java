// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.ldap.MsadUserAccountControlMapperArgs;
import com.pulumi.keycloak.ldap.inputs.MsadUserAccountControlMapperState;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## # keycloak.ldap.MsadUserAccountControlMapper
 * 
 * Allows for creating and managing MSAD user account control mappers for Keycloak
 * users federated via LDAP.
 * 
 * The MSAD (Microsoft Active Directory) user account control mapper is specific
 * to LDAP user federation providers that are pulling from AD, and it can propagate
 * AD user state to Keycloak in order to enforce settings like expired passwords
 * or disabled accounts.
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
 * import com.pulumi.keycloak.ldap.UserFederation;
 * import com.pulumi.keycloak.ldap.UserFederationArgs;
 * import com.pulumi.keycloak.ldap.MsadUserAccountControlMapper;
 * import com.pulumi.keycloak.ldap.MsadUserAccountControlMapperArgs;
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
 *         var ldapUserFederation = new UserFederation("ldapUserFederation", UserFederationArgs.builder()
 *             .name("ad")
 *             .realmId(realm.id())
 *             .usernameLdapAttribute("cn")
 *             .rdnLdapAttribute("cn")
 *             .uuidLdapAttribute("objectGUID")
 *             .userObjectClasses(            
 *                 "person",
 *                 "organizationalPerson",
 *                 "user")
 *             .connectionUrl("ldap://my-ad-server")
 *             .usersDn("dc=example,dc=org")
 *             .bindDn("cn=admin,dc=example,dc=org")
 *             .bindCredential("admin")
 *             .build());
 * 
 *         var msadUserAccountControlMapper = new MsadUserAccountControlMapper("msadUserAccountControlMapper", MsadUserAccountControlMapperArgs.builder()
 *             .realmId(realm.id())
 *             .ldapUserFederationId(ldapUserFederation.id())
 *             .name("msad-user-account-control-mapper")
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
 * - `realm_id` - (Required) The realm that this LDAP mapper will exist in.
 * - `ldap_user_federation_id` - (Required) The ID of the LDAP user federation provider to attach this mapper to.
 * - `name` - (Required) Display name of this mapper when displayed in the console.
 * - `ldap_password_policy_hints_enabled` - (Optional) When `true`, advanced password policies, such as password hints and previous password history will be used when writing new passwords to AD. Defaults to `false`.
 * 
 * ### Import
 * 
 * LDAP mappers can be imported using the format `{{realm_id}}/{{ldap_user_federation_id}}/{{ldap_mapper_id}}`.
 * The ID of the LDAP user federation provider and the mapper can be found within
 * the Keycloak GUI, and they are typically GUIDs:
 * 
 */
@ResourceType(type="keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper")
public class MsadUserAccountControlMapper extends com.pulumi.resources.CustomResource {
    @Export(name="ldapPasswordPolicyHintsEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> ldapPasswordPolicyHintsEnabled;

    public Output<Optional<Boolean>> ldapPasswordPolicyHintsEnabled() {
        return Codegen.optional(this.ldapPasswordPolicyHintsEnabled);
    }
    /**
     * The ldap user federation provider to attach this mapper to.
     * 
     */
    @Export(name="ldapUserFederationId", refs={String.class}, tree="[0]")
    private Output<String> ldapUserFederationId;

    /**
     * @return The ldap user federation provider to attach this mapper to.
     * 
     */
    public Output<String> ldapUserFederationId() {
        return this.ldapUserFederationId;
    }
    /**
     * Display name of the mapper when displayed in the console.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Display name of the mapper when displayed in the console.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The realm in which the ldap user federation provider exists.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The realm in which the ldap user federation provider exists.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MsadUserAccountControlMapper(String name) {
        this(name, MsadUserAccountControlMapperArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MsadUserAccountControlMapper(String name, MsadUserAccountControlMapperArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MsadUserAccountControlMapper(String name, MsadUserAccountControlMapperArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper", name, args == null ? MsadUserAccountControlMapperArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MsadUserAccountControlMapper(String name, Output<String> id, @Nullable MsadUserAccountControlMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:ldap/msadUserAccountControlMapper:MsadUserAccountControlMapper", name, state, makeResourceOptions(options, id));
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
    public static MsadUserAccountControlMapper get(String name, Output<String> id, @Nullable MsadUserAccountControlMapperState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MsadUserAccountControlMapper(name, id, state, options);
    }
}
