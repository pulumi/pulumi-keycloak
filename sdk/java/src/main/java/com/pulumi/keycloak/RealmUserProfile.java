// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.RealmUserProfileArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.RealmUserProfileState;
import com.pulumi.keycloak.outputs.RealmUserProfileAttribute;
import com.pulumi.keycloak.outputs.RealmUserProfileGroup;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Allows for managing Realm User Profiles within Keycloak.
 * 
 * A user profile defines a schema for representing user attributes and how they are managed within a realm.
 * This is a preview feature, hence not fully supported and disabled by default.
 * To enable it, start the server with one of the following flags:
 * - WildFly distribution: `-Dkeycloak.profile.feature.declarative_user_profile=enabled`
 * - Quarkus distribution: `--features=preview` or `--features=declarative-user-profile`
 * 
 * The realm linked to the `keycloak.RealmUserProfile` resource must have the user profile feature enabled.
 * It can be done via the administration UI, or by setting the `userProfileEnabled` realm attribute to `true`.
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
 * import com.pulumi.keycloak.RealmUserProfile;
 * import com.pulumi.keycloak.RealmUserProfileArgs;
 * import com.pulumi.keycloak.inputs.RealmUserProfileAttributeArgs;
 * import com.pulumi.keycloak.inputs.RealmUserProfileAttributePermissionsArgs;
 * import com.pulumi.keycloak.inputs.RealmUserProfileGroupArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *             .attributes(Map.of("userProfileEnabled", true))
 *             .build());
 * 
 *         var userprofile = new RealmUserProfile("userprofile", RealmUserProfileArgs.builder()
 *             .realmId(myRealm.id())
 *             .attributes(            
 *                 RealmUserProfileAttributeArgs.builder()
 *                     .name("field1")
 *                     .displayName("Field 1")
 *                     .group("group1")
 *                     .enabledWhenScopes("offline_access")
 *                     .requiredForRoles("user")
 *                     .requiredForScopes("offline_access")
 *                     .permissions(RealmUserProfileAttributePermissionsArgs.builder()
 *                         .views(                        
 *                             "admin",
 *                             "user")
 *                         .edits(                        
 *                             "admin",
 *                             "user")
 *                         .build())
 *                     .validators(                    
 *                         RealmUserProfileAttributeValidatorArgs.builder()
 *                             .name("person-name-prohibited-characters")
 *                             .build(),
 *                         RealmUserProfileAttributeValidatorArgs.builder()
 *                             .name("pattern")
 *                             .config(Map.ofEntries(
 *                                 Map.entry("pattern", "^[a-z]+$"),
 *                                 Map.entry("error-message", "Nope")
 *                             ))
 *                             .build())
 *                     .annotations(Map.of("foo", "bar"))
 *                     .build(),
 *                 RealmUserProfileAttributeArgs.builder()
 *                     .name("field2")
 *                     .validators(RealmUserProfileAttributeValidatorArgs.builder()
 *                         .name("options")
 *                         .config(Map.of("options", serializeJson(
 *                             jsonArray("opt1"))))
 *                         .build())
 *                     .annotations(Map.of("foo", serializeJson(
 *                         jsonObject(
 *                             jsonProperty("key", "val")
 *                         ))))
 *                     .build())
 *             .groups(            
 *                 RealmUserProfileGroupArgs.builder()
 *                     .name("group1")
 *                     .displayHeader("Group 1")
 *                     .displayDescription("A first group")
 *                     .annotations(Map.ofEntries(
 *                         Map.entry("foo", "bar"),
 *                         Map.entry("foo2", serializeJson(
 *                             jsonObject(
 *                                 jsonProperty("key", "val")
 *                             )))
 *                     ))
 *                     .build(),
 *                 RealmUserProfileGroupArgs.builder()
 *                     .name("group2")
 *                     .build())
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
 * This resource currently does not support importing.
 * 
 */
@ResourceType(type="keycloak:index/realmUserProfile:RealmUserProfile")
public class RealmUserProfile extends com.pulumi.resources.CustomResource {
    /**
     * An ordered list of attributes.
     * 
     */
    @Export(name="attributes", refs={List.class,RealmUserProfileAttribute.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RealmUserProfileAttribute>> attributes;

    /**
     * @return An ordered list of attributes.
     * 
     */
    public Output<Optional<List<RealmUserProfileAttribute>>> attributes() {
        return Codegen.optional(this.attributes);
    }
    /**
     * A list of groups.
     * 
     */
    @Export(name="groups", refs={List.class,RealmUserProfileGroup.class}, tree="[0,1]")
    private Output</* @Nullable */ List<RealmUserProfileGroup>> groups;

    /**
     * @return A list of groups.
     * 
     */
    public Output<Optional<List<RealmUserProfileGroup>>> groups() {
        return Codegen.optional(this.groups);
    }
    /**
     * The ID of the realm the user profile applies to.
     * 
     */
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    /**
     * @return The ID of the realm the user profile applies to.
     * 
     */
    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RealmUserProfile(java.lang.String name) {
        this(name, RealmUserProfileArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RealmUserProfile(java.lang.String name, RealmUserProfileArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RealmUserProfile(java.lang.String name, RealmUserProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmUserProfile:RealmUserProfile", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RealmUserProfile(java.lang.String name, Output<java.lang.String> id, @Nullable RealmUserProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmUserProfile:RealmUserProfile", name, state, makeResourceOptions(options, id), false);
    }

    private static RealmUserProfileArgs makeArgs(RealmUserProfileArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RealmUserProfileArgs.Empty : args;
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
    public static RealmUserProfile get(java.lang.String name, Output<java.lang.String> id, @Nullable RealmUserProfileState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RealmUserProfile(name, id, state, options);
    }
}
