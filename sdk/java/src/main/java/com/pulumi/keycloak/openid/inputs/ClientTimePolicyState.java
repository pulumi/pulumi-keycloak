// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.keycloak.openid.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ClientTimePolicyState extends com.pulumi.resources.ResourceArgs {

    public static final ClientTimePolicyState Empty = new ClientTimePolicyState();

    @Import(name="dayMonth")
    private @Nullable Output<String> dayMonth;

    public Optional<Output<String>> dayMonth() {
        return Optional.ofNullable(this.dayMonth);
    }

    @Import(name="dayMonthEnd")
    private @Nullable Output<String> dayMonthEnd;

    public Optional<Output<String>> dayMonthEnd() {
        return Optional.ofNullable(this.dayMonthEnd);
    }

    @Import(name="decisionStrategy")
    private @Nullable Output<String> decisionStrategy;

    public Optional<Output<String>> decisionStrategy() {
        return Optional.ofNullable(this.decisionStrategy);
    }

    @Import(name="description")
    private @Nullable Output<String> description;

    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="hour")
    private @Nullable Output<String> hour;

    public Optional<Output<String>> hour() {
        return Optional.ofNullable(this.hour);
    }

    @Import(name="hourEnd")
    private @Nullable Output<String> hourEnd;

    public Optional<Output<String>> hourEnd() {
        return Optional.ofNullable(this.hourEnd);
    }

    @Import(name="logic")
    private @Nullable Output<String> logic;

    public Optional<Output<String>> logic() {
        return Optional.ofNullable(this.logic);
    }

    @Import(name="minute")
    private @Nullable Output<String> minute;

    public Optional<Output<String>> minute() {
        return Optional.ofNullable(this.minute);
    }

    @Import(name="minuteEnd")
    private @Nullable Output<String> minuteEnd;

    public Optional<Output<String>> minuteEnd() {
        return Optional.ofNullable(this.minuteEnd);
    }

    @Import(name="month")
    private @Nullable Output<String> month;

    public Optional<Output<String>> month() {
        return Optional.ofNullable(this.month);
    }

    @Import(name="monthEnd")
    private @Nullable Output<String> monthEnd;

    public Optional<Output<String>> monthEnd() {
        return Optional.ofNullable(this.monthEnd);
    }

    @Import(name="name")
    private @Nullable Output<String> name;

    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    @Import(name="notBefore")
    private @Nullable Output<String> notBefore;

    public Optional<Output<String>> notBefore() {
        return Optional.ofNullable(this.notBefore);
    }

    @Import(name="notOnOrAfter")
    private @Nullable Output<String> notOnOrAfter;

    public Optional<Output<String>> notOnOrAfter() {
        return Optional.ofNullable(this.notOnOrAfter);
    }

    @Import(name="realmId")
    private @Nullable Output<String> realmId;

    public Optional<Output<String>> realmId() {
        return Optional.ofNullable(this.realmId);
    }

    @Import(name="resourceServerId")
    private @Nullable Output<String> resourceServerId;

    public Optional<Output<String>> resourceServerId() {
        return Optional.ofNullable(this.resourceServerId);
    }

    @Import(name="year")
    private @Nullable Output<String> year;

    public Optional<Output<String>> year() {
        return Optional.ofNullable(this.year);
    }

    @Import(name="yearEnd")
    private @Nullable Output<String> yearEnd;

    public Optional<Output<String>> yearEnd() {
        return Optional.ofNullable(this.yearEnd);
    }

    private ClientTimePolicyState() {}

    private ClientTimePolicyState(ClientTimePolicyState $) {
        this.dayMonth = $.dayMonth;
        this.dayMonthEnd = $.dayMonthEnd;
        this.decisionStrategy = $.decisionStrategy;
        this.description = $.description;
        this.hour = $.hour;
        this.hourEnd = $.hourEnd;
        this.logic = $.logic;
        this.minute = $.minute;
        this.minuteEnd = $.minuteEnd;
        this.month = $.month;
        this.monthEnd = $.monthEnd;
        this.name = $.name;
        this.notBefore = $.notBefore;
        this.notOnOrAfter = $.notOnOrAfter;
        this.realmId = $.realmId;
        this.resourceServerId = $.resourceServerId;
        this.year = $.year;
        this.yearEnd = $.yearEnd;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ClientTimePolicyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ClientTimePolicyState $;

        public Builder() {
            $ = new ClientTimePolicyState();
        }

        public Builder(ClientTimePolicyState defaults) {
            $ = new ClientTimePolicyState(Objects.requireNonNull(defaults));
        }

        public Builder dayMonth(@Nullable Output<String> dayMonth) {
            $.dayMonth = dayMonth;
            return this;
        }

        public Builder dayMonth(String dayMonth) {
            return dayMonth(Output.of(dayMonth));
        }

        public Builder dayMonthEnd(@Nullable Output<String> dayMonthEnd) {
            $.dayMonthEnd = dayMonthEnd;
            return this;
        }

        public Builder dayMonthEnd(String dayMonthEnd) {
            return dayMonthEnd(Output.of(dayMonthEnd));
        }

        public Builder decisionStrategy(@Nullable Output<String> decisionStrategy) {
            $.decisionStrategy = decisionStrategy;
            return this;
        }

        public Builder decisionStrategy(String decisionStrategy) {
            return decisionStrategy(Output.of(decisionStrategy));
        }

        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder hour(@Nullable Output<String> hour) {
            $.hour = hour;
            return this;
        }

        public Builder hour(String hour) {
            return hour(Output.of(hour));
        }

        public Builder hourEnd(@Nullable Output<String> hourEnd) {
            $.hourEnd = hourEnd;
            return this;
        }

        public Builder hourEnd(String hourEnd) {
            return hourEnd(Output.of(hourEnd));
        }

        public Builder logic(@Nullable Output<String> logic) {
            $.logic = logic;
            return this;
        }

        public Builder logic(String logic) {
            return logic(Output.of(logic));
        }

        public Builder minute(@Nullable Output<String> minute) {
            $.minute = minute;
            return this;
        }

        public Builder minute(String minute) {
            return minute(Output.of(minute));
        }

        public Builder minuteEnd(@Nullable Output<String> minuteEnd) {
            $.minuteEnd = minuteEnd;
            return this;
        }

        public Builder minuteEnd(String minuteEnd) {
            return minuteEnd(Output.of(minuteEnd));
        }

        public Builder month(@Nullable Output<String> month) {
            $.month = month;
            return this;
        }

        public Builder month(String month) {
            return month(Output.of(month));
        }

        public Builder monthEnd(@Nullable Output<String> monthEnd) {
            $.monthEnd = monthEnd;
            return this;
        }

        public Builder monthEnd(String monthEnd) {
            return monthEnd(Output.of(monthEnd));
        }

        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        public Builder notBefore(@Nullable Output<String> notBefore) {
            $.notBefore = notBefore;
            return this;
        }

        public Builder notBefore(String notBefore) {
            return notBefore(Output.of(notBefore));
        }

        public Builder notOnOrAfter(@Nullable Output<String> notOnOrAfter) {
            $.notOnOrAfter = notOnOrAfter;
            return this;
        }

        public Builder notOnOrAfter(String notOnOrAfter) {
            return notOnOrAfter(Output.of(notOnOrAfter));
        }

        public Builder realmId(@Nullable Output<String> realmId) {
            $.realmId = realmId;
            return this;
        }

        public Builder realmId(String realmId) {
            return realmId(Output.of(realmId));
        }

        public Builder resourceServerId(@Nullable Output<String> resourceServerId) {
            $.resourceServerId = resourceServerId;
            return this;
        }

        public Builder resourceServerId(String resourceServerId) {
            return resourceServerId(Output.of(resourceServerId));
        }

        public Builder year(@Nullable Output<String> year) {
            $.year = year;
            return this;
        }

        public Builder year(String year) {
            return year(Output.of(year));
        }

        public Builder yearEnd(@Nullable Output<String> yearEnd) {
            $.yearEnd = yearEnd;
            return this;
        }

        public Builder yearEnd(String yearEnd) {
            return yearEnd(Output.of(yearEnd));
        }

        public ClientTimePolicyState build() {
            return $;
        }
    }

}
