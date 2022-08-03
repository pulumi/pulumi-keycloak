// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.keycloak.outputs.GetRealmKeysKey;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class GetRealmKeysResult {
    private final @Nullable List<String> algorithms;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return (Computed) A list of keys that match the filter criteria. Each key has the following attributes:
     * 
     */
    private final List<GetRealmKeysKey> keys;
    private final String realmId;
    /**
     * @return Key status (string)
     * 
     */
    private final @Nullable List<String> statuses;

    @CustomType.Constructor
    private GetRealmKeysResult(
        @CustomType.Parameter("algorithms") @Nullable List<String> algorithms,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("keys") List<GetRealmKeysKey> keys,
        @CustomType.Parameter("realmId") String realmId,
        @CustomType.Parameter("statuses") @Nullable List<String> statuses) {
        this.algorithms = algorithms;
        this.id = id;
        this.keys = keys;
        this.realmId = realmId;
        this.statuses = statuses;
    }

    public List<String> algorithms() {
        return this.algorithms == null ? List.of() : this.algorithms;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return (Computed) A list of keys that match the filter criteria. Each key has the following attributes:
     * 
     */
    public List<GetRealmKeysKey> keys() {
        return this.keys;
    }
    public String realmId() {
        return this.realmId;
    }
    /**
     * @return Key status (string)
     * 
     */
    public List<String> statuses() {
        return this.statuses == null ? List.of() : this.statuses;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetRealmKeysResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> algorithms;
        private String id;
        private List<GetRealmKeysKey> keys;
        private String realmId;
        private @Nullable List<String> statuses;

        public Builder() {
    	      // Empty
        }

        public Builder(GetRealmKeysResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.algorithms = defaults.algorithms;
    	      this.id = defaults.id;
    	      this.keys = defaults.keys;
    	      this.realmId = defaults.realmId;
    	      this.statuses = defaults.statuses;
        }

        public Builder algorithms(@Nullable List<String> algorithms) {
            this.algorithms = algorithms;
            return this;
        }
        public Builder algorithms(String... algorithms) {
            return algorithms(List.of(algorithms));
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder keys(List<GetRealmKeysKey> keys) {
            this.keys = Objects.requireNonNull(keys);
            return this;
        }
        public Builder keys(GetRealmKeysKey... keys) {
            return keys(List.of(keys));
        }
        public Builder realmId(String realmId) {
            this.realmId = Objects.requireNonNull(realmId);
            return this;
        }
        public Builder statuses(@Nullable List<String> statuses) {
            this.statuses = statuses;
            return this;
        }
        public Builder statuses(String... statuses) {
            return statuses(List.of(statuses));
        }        public GetRealmKeysResult build() {
            return new GetRealmKeysResult(algorithms, id, keys, realmId, statuses);
        }
    }
}
