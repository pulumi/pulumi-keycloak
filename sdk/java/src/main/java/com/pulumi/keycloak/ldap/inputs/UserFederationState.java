// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.ldap.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.keycloak.ldap.inputs.UserFederationCacheArgs;
import com.pulumi.keycloak.ldap.inputs.UserFederationKerberosArgs;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class UserFederationState extends com.pulumi.resources.ResourceArgs {

    public static final UserFederationState Empty = new UserFederationState();

    /**
     * The number of users to sync within a single transaction.
     * 
     */
    @Import(name="batchSizeForSync")
    private @Nullable Output<Integer> batchSizeForSync;

    /**
     * @return The number of users to sync within a single transaction.
     * 
     */
    public Optional<Output<Integer>> batchSizeForSync() {
        return Optional.ofNullable(this.batchSizeForSync);
    }

    /**
     * Password of LDAP admin.
     * 
     */
    @Import(name="bindCredential")
    private @Nullable Output<String> bindCredential;

    /**
     * @return Password of LDAP admin.
     * 
     */
    public Optional<Output<String>> bindCredential() {
        return Optional.ofNullable(this.bindCredential);
    }

    /**
     * DN of LDAP admin, which will be used by Keycloak to access LDAP server.
     * 
     */
    @Import(name="bindDn")
    private @Nullable Output<String> bindDn;

    /**
     * @return DN of LDAP admin, which will be used by Keycloak to access LDAP server.
     * 
     */
    public Optional<Output<String>> bindDn() {
        return Optional.ofNullable(this.bindDn);
    }

    /**
     * Settings regarding cache policy for this realm.
     * 
     */
    @Import(name="cache")
    private @Nullable Output<UserFederationCacheArgs> cache;

    /**
     * @return Settings regarding cache policy for this realm.
     * 
     */
    public Optional<Output<UserFederationCacheArgs>> cache() {
        return Optional.ofNullable(this.cache);
    }

    /**
     * How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
     * sync.
     * 
     */
    @Import(name="changedSyncPeriod")
    private @Nullable Output<Integer> changedSyncPeriod;

    /**
     * @return How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
     * sync.
     * 
     */
    public Optional<Output<Integer>> changedSyncPeriod() {
        return Optional.ofNullable(this.changedSyncPeriod);
    }

    /**
     * LDAP connection timeout (duration string)
     * 
     */
    @Import(name="connectionTimeout")
    private @Nullable Output<String> connectionTimeout;

    /**
     * @return LDAP connection timeout (duration string)
     * 
     */
    public Optional<Output<String>> connectionTimeout() {
        return Optional.ofNullable(this.connectionTimeout);
    }

    /**
     * Connection URL to the LDAP server.
     * 
     */
    @Import(name="connectionUrl")
    private @Nullable Output<String> connectionUrl;

    /**
     * @return Connection URL to the LDAP server.
     * 
     */
    public Optional<Output<String>> connectionUrl() {
        return Optional.ofNullable(this.connectionUrl);
    }

    /**
     * Additional LDAP filter for filtering searched users. Must begin with &#39;(&#39; and end with &#39;)&#39;.
     * 
     */
    @Import(name="customUserSearchFilter")
    private @Nullable Output<String> customUserSearchFilter;

    /**
     * @return Additional LDAP filter for filtering searched users. Must begin with &#39;(&#39; and end with &#39;)&#39;.
     * 
     */
    public Optional<Output<String>> customUserSearchFilter() {
        return Optional.ofNullable(this.customUserSearchFilter);
    }

    /**
     * When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP
     * user federation provider.
     * 
     */
    @Import(name="deleteDefaultMappers")
    private @Nullable Output<Boolean> deleteDefaultMappers;

    /**
     * @return When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP
     * user federation provider.
     * 
     */
    public Optional<Output<Boolean>> deleteDefaultMappers() {
        return Optional.ofNullable(this.deleteDefaultMappers);
    }

    /**
     * READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
     * 
     */
    @Import(name="editMode")
    private @Nullable Output<String> editMode;

    /**
     * @return READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
     * 
     */
    public Optional<Output<String>> editMode() {
        return Optional.ofNullable(this.editMode);
    }

    /**
     * When false, this provider will not be used when performing queries for users.
     * 
     */
    @Import(name="enabled")
    private @Nullable Output<Boolean> enabled;

    /**
     * @return When false, this provider will not be used when performing queries for users.
     * 
     */
    public Optional<Output<Boolean>> enabled() {
        return Optional.ofNullable(this.enabled);
    }

    /**
     * How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
     * 
     */
    @Import(name="fullSyncPeriod")
    private @Nullable Output<Integer> fullSyncPeriod;

    /**
     * @return How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
     * 
     */
    public Optional<Output<Integer>> fullSyncPeriod() {
        return Optional.ofNullable(this.fullSyncPeriod);
    }

    /**
     * When true, LDAP users will be imported into the Keycloak database.
     * 
     */
    @Import(name="importEnabled")
    private @Nullable Output<Boolean> importEnabled;

    /**
     * @return When true, LDAP users will be imported into the Keycloak database.
     * 
     */
    public Optional<Output<Boolean>> importEnabled() {
        return Optional.ofNullable(this.importEnabled);
    }

    /**
     * Settings regarding kerberos authentication for this realm.
     * 
     */
    @Import(name="kerberos")
    private @Nullable Output<UserFederationKerberosArgs> kerberos;

    /**
     * @return Settings regarding kerberos authentication for this realm.
     * 
     */
    public Optional<Output<UserFederationKerberosArgs>> kerberos() {
        return Optional.ofNullable(this.kerberos);
    }

    /**
     * Display name of the provider when displayed in the console.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Display name of the provider when displayed in the console.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * When true, Keycloak assumes the LDAP server supports pagination.
     * 
     */
    @Import(name="pagination")
    private @Nullable Output<Boolean> pagination;

    /**
     * @return When true, Keycloak assumes the LDAP server supports pagination.
     * 
     */
    public Optional<Output<Boolean>> pagination() {
        return Optional.ofNullable(this.pagination);
    }

    /**
     * Priority of this provider when looking up users. Lower values are first.
     * 
     */
    @Import(name="priority")
    private @Nullable Output<Integer> priority;

    /**
     * @return Priority of this provider when looking up users. Lower values are first.
     * 
     */
    public Optional<Output<Integer>> priority() {
        return Optional.ofNullable(this.priority);
    }

    /**
     * Name of the LDAP attribute to use as the relative distinguished name.
     * 
     */
    @Import(name="rdnLdapAttribute")
    private @Nullable Output<String> rdnLdapAttribute;

    /**
     * @return Name of the LDAP attribute to use as the relative distinguished name.
     * 
     */
    public Optional<Output<String>> rdnLdapAttribute() {
        return Optional.ofNullable(this.rdnLdapAttribute);
    }

    /**
     * LDAP read timeout (duration string)
     * 
     */
    @Import(name="readTimeout")
    private @Nullable Output<String> readTimeout;

    /**
     * @return LDAP read timeout (duration string)
     * 
     */
    public Optional<Output<String>> readTimeout() {
        return Optional.ofNullable(this.readTimeout);
    }

    /**
     * The realm this provider will provide user federation for.
     * 
     */
    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    /**
     * @return The realm this provider will provide user federation for.
     * 
     */
    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    /**
     * ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
     * 
     */
    @Import(name="searchScope")
    private @Nullable Output<String> searchScope;

    /**
     * @return ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
     * 
     */
    public Optional<Output<String>> searchScope() {
        return Optional.ofNullable(this.searchScope);
    }

    /**
     * When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
     * 
     */
    @Import(name="startTls")
    private @Nullable Output<Boolean> startTls;

    /**
     * @return When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
     * 
     */
    public Optional<Output<Boolean>> startTls() {
        return Optional.ofNullable(this.startTls);
    }

    /**
     * When true, newly created users will be synced back to LDAP.
     * 
     */
    @Import(name="syncRegistrations")
    private @Nullable Output<Boolean> syncRegistrations;

    /**
     * @return When true, newly created users will be synced back to LDAP.
     * 
     */
    public Optional<Output<Boolean>> syncRegistrations() {
        return Optional.ofNullable(this.syncRegistrations);
    }

    /**
     * If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
     * 
     */
    @Import(name="trustEmail")
    private @Nullable Output<Boolean> trustEmail;

    /**
     * @return If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
     * 
     */
    public Optional<Output<Boolean>> trustEmail() {
        return Optional.ofNullable(this.trustEmail);
    }

    /**
     * When `true`, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
     * 
     */
    @Import(name="usePasswordModifyExtendedOp")
    private @Nullable Output<Boolean> usePasswordModifyExtendedOp;

    /**
     * @return When `true`, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
     * 
     */
    public Optional<Output<Boolean>> usePasswordModifyExtendedOp() {
        return Optional.ofNullable(this.usePasswordModifyExtendedOp);
    }

    @Import(name="useTruststoreSpi")
    private @Nullable Output<String> useTruststoreSpi;

    public Optional<Output<String>> useTruststoreSpi() {
        return Optional.ofNullable(this.useTruststoreSpi);
    }

    /**
     * All values of LDAP objectClass attribute for users in LDAP.
     * 
     */
    @Import(name="userObjectClasses")
    private @Nullable Output<List<String>> userObjectClasses;

    /**
     * @return All values of LDAP objectClass attribute for users in LDAP.
     * 
     */
    public Optional<Output<List<String>>> userObjectClasses() {
        return Optional.ofNullable(this.userObjectClasses);
    }

    /**
     * Name of the LDAP attribute to use as the Keycloak username.
     * 
     */
    @Import(name="usernameLdapAttribute")
    private @Nullable Output<String> usernameLdapAttribute;

    /**
     * @return Name of the LDAP attribute to use as the Keycloak username.
     * 
     */
    public Optional<Output<String>> usernameLdapAttribute() {
        return Optional.ofNullable(this.usernameLdapAttribute);
    }

    /**
     * Full DN of LDAP tree where your users are.
     * 
     */
    @Import(name="usersDn")
    private @Nullable Output<String> usersDn;

    /**
     * @return Full DN of LDAP tree where your users are.
     * 
     */
    public Optional<Output<String>> usersDn() {
        return Optional.ofNullable(this.usersDn);
    }

    /**
     * Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
     * 
     */
    @Import(name="uuidLdapAttribute")
    private @Nullable Output<String> uuidLdapAttribute;

    /**
     * @return Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
     * 
     */
    public Optional<Output<String>> uuidLdapAttribute() {
        return Optional.ofNullable(this.uuidLdapAttribute);
    }

    /**
     * When true, Keycloak will validate passwords using the realm policy before updating it.
     * 
     */
    @Import(name="validatePasswordPolicy")
    private @Nullable Output<Boolean> validatePasswordPolicy;

    /**
     * @return When true, Keycloak will validate passwords using the realm policy before updating it.
     * 
     */
    public Optional<Output<Boolean>> validatePasswordPolicy() {
        return Optional.ofNullable(this.validatePasswordPolicy);
    }

    /**
     * LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
     * 
     */
    @Import(name="vendor")
    private @Nullable Output<String> vendor;

    /**
     * @return LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
     * 
     */
    public Optional<Output<String>> vendor() {
        return Optional.ofNullable(this.vendor);
    }

    private UserFederationState() {}

    private UserFederationState(UserFederationState $) {
        this.batchSizeForSync = $.batchSizeForSync;
        this.bindCredential = $.bindCredential;
        this.bindDn = $.bindDn;
        this.cache = $.cache;
        this.changedSyncPeriod = $.changedSyncPeriod;
        this.connectionTimeout = $.connectionTimeout;
        this.connectionUrl = $.connectionUrl;
        this.customUserSearchFilter = $.customUserSearchFilter;
        this.deleteDefaultMappers = $.deleteDefaultMappers;
        this.editMode = $.editMode;
        this.enabled = $.enabled;
        this.fullSyncPeriod = $.fullSyncPeriod;
        this.importEnabled = $.importEnabled;
        this.kerberos = $.kerberos;
        this.name = $.name;
        this.pagination = $.pagination;
        this.priority = $.priority;
        this.rdnLdapAttribute = $.rdnLdapAttribute;
        this.readTimeout = $.readTimeout;
        this.realmId = $.realmId;
        this.searchScope = $.searchScope;
        this.startTls = $.startTls;
        this.syncRegistrations = $.syncRegistrations;
        this.trustEmail = $.trustEmail;
        this.usePasswordModifyExtendedOp = $.usePasswordModifyExtendedOp;
        this.useTruststoreSpi = $.useTruststoreSpi;
        this.userObjectClasses = $.userObjectClasses;
        this.usernameLdapAttribute = $.usernameLdapAttribute;
        this.usersDn = $.usersDn;
        this.uuidLdapAttribute = $.uuidLdapAttribute;
        this.validatePasswordPolicy = $.validatePasswordPolicy;
        this.vendor = $.vendor;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(UserFederationState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private UserFederationState $;

        public Builder() {
            $ = new UserFederationState();
        }

        public Builder(UserFederationState defaults) {
            $ = new UserFederationState(Objects.requireNonNull(defaults));
        }

        /**
         * @param batchSizeForSync The number of users to sync within a single transaction.
         * 
         * @return builder
         * 
         */
        public Builder batchSizeForSync(@Nullable Output<Integer> batchSizeForSync) {
            $.batchSizeForSync = batchSizeForSync;
            return this;
        }

        /**
         * @param batchSizeForSync The number of users to sync within a single transaction.
         * 
         * @return builder
         * 
         */
        public Builder batchSizeForSync(Integer batchSizeForSync) {
            return batchSizeForSync(Output.of(batchSizeForSync));
        }

        /**
         * @param bindCredential Password of LDAP admin.
         * 
         * @return builder
         * 
         */
        public Builder bindCredential(@Nullable Output<String> bindCredential) {
            $.bindCredential = bindCredential;
            return this;
        }

        /**
         * @param bindCredential Password of LDAP admin.
         * 
         * @return builder
         * 
         */
        public Builder bindCredential(String bindCredential) {
            return bindCredential(Output.of(bindCredential));
        }

        /**
         * @param bindDn DN of LDAP admin, which will be used by Keycloak to access LDAP server.
         * 
         * @return builder
         * 
         */
        public Builder bindDn(@Nullable Output<String> bindDn) {
            $.bindDn = bindDn;
            return this;
        }

        /**
         * @param bindDn DN of LDAP admin, which will be used by Keycloak to access LDAP server.
         * 
         * @return builder
         * 
         */
        public Builder bindDn(String bindDn) {
            return bindDn(Output.of(bindDn));
        }

        /**
         * @param cache Settings regarding cache policy for this realm.
         * 
         * @return builder
         * 
         */
        public Builder cache(@Nullable Output<UserFederationCacheArgs> cache) {
            $.cache = cache;
            return this;
        }

        /**
         * @param cache Settings regarding cache policy for this realm.
         * 
         * @return builder
         * 
         */
        public Builder cache(UserFederationCacheArgs cache) {
            return cache(Output.of(cache));
        }

        /**
         * @param changedSyncPeriod How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
         * sync.
         * 
         * @return builder
         * 
         */
        public Builder changedSyncPeriod(@Nullable Output<Integer> changedSyncPeriod) {
            $.changedSyncPeriod = changedSyncPeriod;
            return this;
        }

        /**
         * @param changedSyncPeriod How frequently Keycloak should sync changed LDAP users, in seconds. Omit this property to disable periodic changed users
         * sync.
         * 
         * @return builder
         * 
         */
        public Builder changedSyncPeriod(Integer changedSyncPeriod) {
            return changedSyncPeriod(Output.of(changedSyncPeriod));
        }

        /**
         * @param connectionTimeout LDAP connection timeout (duration string)
         * 
         * @return builder
         * 
         */
        public Builder connectionTimeout(@Nullable Output<String> connectionTimeout) {
            $.connectionTimeout = connectionTimeout;
            return this;
        }

        /**
         * @param connectionTimeout LDAP connection timeout (duration string)
         * 
         * @return builder
         * 
         */
        public Builder connectionTimeout(String connectionTimeout) {
            return connectionTimeout(Output.of(connectionTimeout));
        }

        /**
         * @param connectionUrl Connection URL to the LDAP server.
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(@Nullable Output<String> connectionUrl) {
            $.connectionUrl = connectionUrl;
            return this;
        }

        /**
         * @param connectionUrl Connection URL to the LDAP server.
         * 
         * @return builder
         * 
         */
        public Builder connectionUrl(String connectionUrl) {
            return connectionUrl(Output.of(connectionUrl));
        }

        /**
         * @param customUserSearchFilter Additional LDAP filter for filtering searched users. Must begin with &#39;(&#39; and end with &#39;)&#39;.
         * 
         * @return builder
         * 
         */
        public Builder customUserSearchFilter(@Nullable Output<String> customUserSearchFilter) {
            $.customUserSearchFilter = customUserSearchFilter;
            return this;
        }

        /**
         * @param customUserSearchFilter Additional LDAP filter for filtering searched users. Must begin with &#39;(&#39; and end with &#39;)&#39;.
         * 
         * @return builder
         * 
         */
        public Builder customUserSearchFilter(String customUserSearchFilter) {
            return customUserSearchFilter(Output.of(customUserSearchFilter));
        }

        /**
         * @param deleteDefaultMappers When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP
         * user federation provider.
         * 
         * @return builder
         * 
         */
        public Builder deleteDefaultMappers(@Nullable Output<Boolean> deleteDefaultMappers) {
            $.deleteDefaultMappers = deleteDefaultMappers;
            return this;
        }

        /**
         * @param deleteDefaultMappers When true, the provider will delete the default mappers which are normally created by Keycloak when creating an LDAP
         * user federation provider.
         * 
         * @return builder
         * 
         */
        public Builder deleteDefaultMappers(Boolean deleteDefaultMappers) {
            return deleteDefaultMappers(Output.of(deleteDefaultMappers));
        }

        /**
         * @param editMode READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
         * 
         * @return builder
         * 
         */
        public Builder editMode(@Nullable Output<String> editMode) {
            $.editMode = editMode;
            return this;
        }

        /**
         * @param editMode READ_ONLY and WRITABLE are self-explanatory. UNSYNCED allows user data to be imported but not synced back to LDAP.
         * 
         * @return builder
         * 
         */
        public Builder editMode(String editMode) {
            return editMode(Output.of(editMode));
        }

        /**
         * @param enabled When false, this provider will not be used when performing queries for users.
         * 
         * @return builder
         * 
         */
        public Builder enabled(@Nullable Output<Boolean> enabled) {
            $.enabled = enabled;
            return this;
        }

        /**
         * @param enabled When false, this provider will not be used when performing queries for users.
         * 
         * @return builder
         * 
         */
        public Builder enabled(Boolean enabled) {
            return enabled(Output.of(enabled));
        }

        /**
         * @param fullSyncPeriod How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
         * 
         * @return builder
         * 
         */
        public Builder fullSyncPeriod(@Nullable Output<Integer> fullSyncPeriod) {
            $.fullSyncPeriod = fullSyncPeriod;
            return this;
        }

        /**
         * @param fullSyncPeriod How frequently Keycloak should sync all LDAP users, in seconds. Omit this property to disable periodic full sync.
         * 
         * @return builder
         * 
         */
        public Builder fullSyncPeriod(Integer fullSyncPeriod) {
            return fullSyncPeriod(Output.of(fullSyncPeriod));
        }

        /**
         * @param importEnabled When true, LDAP users will be imported into the Keycloak database.
         * 
         * @return builder
         * 
         */
        public Builder importEnabled(@Nullable Output<Boolean> importEnabled) {
            $.importEnabled = importEnabled;
            return this;
        }

        /**
         * @param importEnabled When true, LDAP users will be imported into the Keycloak database.
         * 
         * @return builder
         * 
         */
        public Builder importEnabled(Boolean importEnabled) {
            return importEnabled(Output.of(importEnabled));
        }

        /**
         * @param kerberos Settings regarding kerberos authentication for this realm.
         * 
         * @return builder
         * 
         */
        public Builder kerberos(@Nullable Output<UserFederationKerberosArgs> kerberos) {
            $.kerberos = kerberos;
            return this;
        }

        /**
         * @param kerberos Settings regarding kerberos authentication for this realm.
         * 
         * @return builder
         * 
         */
        public Builder kerberos(UserFederationKerberosArgs kerberos) {
            return kerberos(Output.of(kerberos));
        }

        /**
         * @param name Display name of the provider when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Display name of the provider when displayed in the console.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param pagination When true, Keycloak assumes the LDAP server supports pagination.
         * 
         * @return builder
         * 
         */
        public Builder pagination(@Nullable Output<Boolean> pagination) {
            $.pagination = pagination;
            return this;
        }

        /**
         * @param pagination When true, Keycloak assumes the LDAP server supports pagination.
         * 
         * @return builder
         * 
         */
        public Builder pagination(Boolean pagination) {
            return pagination(Output.of(pagination));
        }

        /**
         * @param priority Priority of this provider when looking up users. Lower values are first.
         * 
         * @return builder
         * 
         */
        public Builder priority(@Nullable Output<Integer> priority) {
            $.priority = priority;
            return this;
        }

        /**
         * @param priority Priority of this provider when looking up users. Lower values are first.
         * 
         * @return builder
         * 
         */
        public Builder priority(Integer priority) {
            return priority(Output.of(priority));
        }

        /**
         * @param rdnLdapAttribute Name of the LDAP attribute to use as the relative distinguished name.
         * 
         * @return builder
         * 
         */
        public Builder rdnLdapAttribute(@Nullable Output<String> rdnLdapAttribute) {
            $.rdnLdapAttribute = rdnLdapAttribute;
            return this;
        }

        /**
         * @param rdnLdapAttribute Name of the LDAP attribute to use as the relative distinguished name.
         * 
         * @return builder
         * 
         */
        public Builder rdnLdapAttribute(String rdnLdapAttribute) {
            return rdnLdapAttribute(Output.of(rdnLdapAttribute));
        }

        /**
         * @param readTimeout LDAP read timeout (duration string)
         * 
         * @return builder
         * 
         */
        public Builder readTimeout(@Nullable Output<String> readTimeout) {
            $.readTimeout = readTimeout;
            return this;
        }

        /**
         * @param readTimeout LDAP read timeout (duration string)
         * 
         * @return builder
         * 
         */
        public Builder readTimeout(String readTimeout) {
            return readTimeout(Output.of(readTimeout));
        }

        /**
         * @param realmId The realm this provider will provide user federation for.
         * 
         * @return builder
         * 
         */
        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        /**
         * @param realmId The realm this provider will provide user federation for.
         * 
         * @return builder
         * 
         */
        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        /**
         * @param searchScope ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
         * 
         * @return builder
         * 
         */
        public Builder searchScope(@Nullable Output<String> searchScope) {
            $.searchScope = searchScope;
            return this;
        }

        /**
         * @param searchScope ONE_LEVEL: only search for users in the DN specified by user_dn. SUBTREE: search entire LDAP subtree.
         * 
         * @return builder
         * 
         */
        public Builder searchScope(String searchScope) {
            return searchScope(Output.of(searchScope));
        }

        /**
         * @param startTls When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
         * 
         * @return builder
         * 
         */
        public Builder startTls(@Nullable Output<Boolean> startTls) {
            $.startTls = startTls;
            return this;
        }

        /**
         * @param startTls When true, Keycloak will encrypt the connection to LDAP using STARTTLS, which will disable connection pooling.
         * 
         * @return builder
         * 
         */
        public Builder startTls(Boolean startTls) {
            return startTls(Output.of(startTls));
        }

        /**
         * @param syncRegistrations When true, newly created users will be synced back to LDAP.
         * 
         * @return builder
         * 
         */
        public Builder syncRegistrations(@Nullable Output<Boolean> syncRegistrations) {
            $.syncRegistrations = syncRegistrations;
            return this;
        }

        /**
         * @param syncRegistrations When true, newly created users will be synced back to LDAP.
         * 
         * @return builder
         * 
         */
        public Builder syncRegistrations(Boolean syncRegistrations) {
            return syncRegistrations(Output.of(syncRegistrations));
        }

        /**
         * @param trustEmail If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
         * 
         * @return builder
         * 
         */
        public Builder trustEmail(@Nullable Output<Boolean> trustEmail) {
            $.trustEmail = trustEmail;
            return this;
        }

        /**
         * @param trustEmail If enabled, email provided by this provider is not verified even if verification is enabled for the realm.
         * 
         * @return builder
         * 
         */
        public Builder trustEmail(Boolean trustEmail) {
            return trustEmail(Output.of(trustEmail));
        }

        /**
         * @param usePasswordModifyExtendedOp When `true`, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
         * 
         * @return builder
         * 
         */
        public Builder usePasswordModifyExtendedOp(@Nullable Output<Boolean> usePasswordModifyExtendedOp) {
            $.usePasswordModifyExtendedOp = usePasswordModifyExtendedOp;
            return this;
        }

        /**
         * @param usePasswordModifyExtendedOp When `true`, use the LDAPv3 Password Modify Extended Operation (RFC-3062).
         * 
         * @return builder
         * 
         */
        public Builder usePasswordModifyExtendedOp(Boolean usePasswordModifyExtendedOp) {
            return usePasswordModifyExtendedOp(Output.of(usePasswordModifyExtendedOp));
        }

        public Builder useTruststoreSpi(@Nullable Output<String> useTruststoreSpi) {
            $.useTruststoreSpi = useTruststoreSpi;
            return this;
        }

        public Builder useTruststoreSpi(String useTruststoreSpi) {
            return useTruststoreSpi(Output.of(useTruststoreSpi));
        }

        /**
         * @param userObjectClasses All values of LDAP objectClass attribute for users in LDAP.
         * 
         * @return builder
         * 
         */
        public Builder userObjectClasses(@Nullable Output<List<String>> userObjectClasses) {
            $.userObjectClasses = userObjectClasses;
            return this;
        }

        /**
         * @param userObjectClasses All values of LDAP objectClass attribute for users in LDAP.
         * 
         * @return builder
         * 
         */
        public Builder userObjectClasses(List<String> userObjectClasses) {
            return userObjectClasses(Output.of(userObjectClasses));
        }

        /**
         * @param userObjectClasses All values of LDAP objectClass attribute for users in LDAP.
         * 
         * @return builder
         * 
         */
        public Builder userObjectClasses(String... userObjectClasses) {
            return userObjectClasses(List.of(userObjectClasses));
        }

        /**
         * @param usernameLdapAttribute Name of the LDAP attribute to use as the Keycloak username.
         * 
         * @return builder
         * 
         */
        public Builder usernameLdapAttribute(@Nullable Output<String> usernameLdapAttribute) {
            $.usernameLdapAttribute = usernameLdapAttribute;
            return this;
        }

        /**
         * @param usernameLdapAttribute Name of the LDAP attribute to use as the Keycloak username.
         * 
         * @return builder
         * 
         */
        public Builder usernameLdapAttribute(String usernameLdapAttribute) {
            return usernameLdapAttribute(Output.of(usernameLdapAttribute));
        }

        /**
         * @param usersDn Full DN of LDAP tree where your users are.
         * 
         * @return builder
         * 
         */
        public Builder usersDn(@Nullable Output<String> usersDn) {
            $.usersDn = usersDn;
            return this;
        }

        /**
         * @param usersDn Full DN of LDAP tree where your users are.
         * 
         * @return builder
         * 
         */
        public Builder usersDn(String usersDn) {
            return usersDn(Output.of(usersDn));
        }

        /**
         * @param uuidLdapAttribute Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
         * 
         * @return builder
         * 
         */
        public Builder uuidLdapAttribute(@Nullable Output<String> uuidLdapAttribute) {
            $.uuidLdapAttribute = uuidLdapAttribute;
            return this;
        }

        /**
         * @param uuidLdapAttribute Name of the LDAP attribute to use as a unique object identifier for objects in LDAP.
         * 
         * @return builder
         * 
         */
        public Builder uuidLdapAttribute(String uuidLdapAttribute) {
            return uuidLdapAttribute(Output.of(uuidLdapAttribute));
        }

        /**
         * @param validatePasswordPolicy When true, Keycloak will validate passwords using the realm policy before updating it.
         * 
         * @return builder
         * 
         */
        public Builder validatePasswordPolicy(@Nullable Output<Boolean> validatePasswordPolicy) {
            $.validatePasswordPolicy = validatePasswordPolicy;
            return this;
        }

        /**
         * @param validatePasswordPolicy When true, Keycloak will validate passwords using the realm policy before updating it.
         * 
         * @return builder
         * 
         */
        public Builder validatePasswordPolicy(Boolean validatePasswordPolicy) {
            return validatePasswordPolicy(Output.of(validatePasswordPolicy));
        }

        /**
         * @param vendor LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
         * 
         * @return builder
         * 
         */
        public Builder vendor(@Nullable Output<String> vendor) {
            $.vendor = vendor;
            return this;
        }

        /**
         * @param vendor LDAP vendor. I am almost certain this field does nothing, but the UI indicates that it is required.
         * 
         * @return builder
         * 
         */
        public Builder vendor(String vendor) {
            return vendor(Output.of(vendor));
        }

        public UserFederationState build() {
            return $;
        }
    }

}
