
-- +migrate Up
create type alert_type_options as enum (
    'price_reaches',
    'price_rises_above',
    'price_drops_to',
    'change_is_over',
    'change_is_under'
);

create type alert_frequency_options as enum (
    'only_once',
    'once_a_day',
    'always'
);
create table if not exists user_token_price_alerts
(
    user_id       text,
    coin_gecko_id text,
    symbol        text,
    alert_type    alert_type_options,
    frequency     alert_frequency_options,
    price         float8
);

create index user_token_price_alerts_user_id_index
    on user_token_price_alerts (user_id);

-- +migrate Down
DROP ENUM IF EXISTS alert_type_options;
DROP ENUM IF EXISTS alert_frequency_options;
DROP TABLE IF EXISTS user_token_price_alerts;

