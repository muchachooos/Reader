DROP TABLE new_order, item, delivery, payment;

CREATE TABLE new_order
(
    order_uid          VARCHAR(255) PRIMARY KEY,
    track_number       VARCHAR(255),
    entry              VARCHAR(255),
    locale             VARCHAR(255),
    internal_signature VARCHAR(255),
    customer_id        VARCHAR(255),
    delivery_service   VARCHAR(255),
    shardkey           VARCHAR(255),
    sm_id              INT,
    date_created       TIMESTAMP,
    oof_shard          VARCHAR(255)
);

INSERT INTO new_order (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ('b563feb7b2b84b6test2','WBILMTESTTRACK','WBIL', 'en', '', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1' )
;

CREATE TABLE item
(
    order_uid    VARCHAR(255),
    chrt_id      INT,
    track_number VARCHAR(255),
    price        INT,
    rid          VARCHAR(255),
    name         VARCHAR(255),
    sale         INT,
    size         VARCHAR(255),
    total_price  INT,
    nm_id        INT,
    brand        VARCHAR(255),
    status       INT,
    FOREIGN KEY (order_uid) REFERENCES new_order (order_uid)
);

INSERT INTO item (order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status)
VALUES ('b563feb7b2b84b6test',9999999,'WBILMTESTTRACK', 999,'ab4219087a764ae0btest','Mascaras', 30, 0,317,2389212,'Vivienne Sabo',202)
;

CREATE TABLE delivery
(
    order_uid VARCHAR(255),
    name      VARCHAR(255),
    phone     VARCHAR(255),
    zip       VARCHAR(255),
    city      VARCHAR(255),
    address   VARCHAR(255),
    region    VARCHAR(255),
    email     VARCHAR(255),
    FOREIGN KEY (order_uid) REFERENCES new_order (order_uid)
);

CREATE TABLE payment
(
    order_uid     VARCHAR(255),
    transaction   VARCHAR(255),
    request_id    VARCHAR(255),
    currency      VARCHAR(255),
    provider      VARCHAR(255),
    amount        INT,
    payment_dt    INT,
    bank          VARCHAR(255),
    delivery_cost INT,
    goods_total   INT,
    custom_fee    INT,
    FOREIGN KEY (order_uid) REFERENCES new_order (order_uid)
);
