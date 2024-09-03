// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.keycloak.RealmEventsArgs;
import com.pulumi.keycloak.Utilities;
import com.pulumi.keycloak.inputs.RealmEventsState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * # keycloak.RealmEvents
 * 
 * Allows for managing Realm Events settings within Keycloak.
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
 * import com.pulumi.keycloak.RealmEvents;
 * import com.pulumi.keycloak.RealmEventsArgs;
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
 *             .build());
 * 
 *         var realmEvents = new RealmEvents("realmEvents", RealmEventsArgs.builder()
 *             .realmId(realm.id())
 *             .eventsEnabled(true)
 *             .eventsExpiration(3600)
 *             .adminEventsEnabled(true)
 *             .adminEventsDetailsEnabled(true)
 *             .enabledEventTypes(            
 *                 "LOGIN",
 *                 "LOGOUT")
 *             .eventsListeners("jboss-logging")
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
 * - `realm_id` - (Required) The name of the realm the event settings apply to.
 * - `admin_events_enabled` - (Optional) When true, admin events are saved to the database, making them available through the admin console. Defaults to `false`.
 * - `admin_events_details_enabled` - (Optional) When true, saved admin events will included detailed information for create/update requests. Defaults to `false`.
 * - `events_enabled` - (Optional) When true, events from `enabled_event_types` are saved to the database, making them available through the admin console. Defaults to `false`.
 * - `events_expiration` - (Optional) The amount of time in seconds events will be saved in the database. Defaults to `0` or never.
 * - `enabled_event_types` - (Optional) The event types that will be saved to the database. Omitting this field enables all event types. Defaults to `[]` or all event types.
 * - `events_listeners` - (Optional) The event listeners that events should be sent to. Defaults to `[]` or none. Note that new realms enable the `jboss-logging` listener by default, and this resource will remove that unless it is specified.
 * 
 */
@ResourceType(type="keycloak:index/realmEvents:RealmEvents")
public class RealmEvents extends com.pulumi.resources.CustomResource {
    @Export(name="adminEventsDetailsEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminEventsDetailsEnabled;

    public Output<Optional<Boolean>> adminEventsDetailsEnabled() {
        return Codegen.optional(this.adminEventsDetailsEnabled);
    }
    @Export(name="adminEventsEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> adminEventsEnabled;

    public Output<Optional<Boolean>> adminEventsEnabled() {
        return Codegen.optional(this.adminEventsEnabled);
    }
    @Export(name="enabledEventTypes", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> enabledEventTypes;

    public Output<Optional<List<String>>> enabledEventTypes() {
        return Codegen.optional(this.enabledEventTypes);
    }
    @Export(name="eventsEnabled", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> eventsEnabled;

    public Output<Optional<Boolean>> eventsEnabled() {
        return Codegen.optional(this.eventsEnabled);
    }
    @Export(name="eventsExpiration", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> eventsExpiration;

    public Output<Optional<Integer>> eventsExpiration() {
        return Codegen.optional(this.eventsExpiration);
    }
    @Export(name="eventsListeners", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> eventsListeners;

    public Output<Optional<List<String>>> eventsListeners() {
        return Codegen.optional(this.eventsListeners);
    }
    @Export(name="realmId", refs={String.class}, tree="[0]")
    private Output<String> realmId;

    public Output<String> realmId() {
        return this.realmId;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RealmEvents(java.lang.String name) {
        this(name, RealmEventsArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RealmEvents(java.lang.String name, RealmEventsArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RealmEvents(java.lang.String name, RealmEventsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmEvents:RealmEvents", name, makeArgs(args, options), makeResourceOptions(options, Codegen.empty()), false);
    }

    private RealmEvents(java.lang.String name, Output<java.lang.String> id, @Nullable RealmEventsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("keycloak:index/realmEvents:RealmEvents", name, state, makeResourceOptions(options, id), false);
    }

    private static RealmEventsArgs makeArgs(RealmEventsArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        if (options != null && options.getUrn().isPresent()) {
            return null;
        }
        return args == null ? RealmEventsArgs.Empty : args;
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
    public static RealmEvents get(java.lang.String name, Output<java.lang.String> id, @Nullable RealmEventsState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RealmEvents(name, id, state, options);
    }
}
