# API Documentation

Демонстрационный сервис с простейшим интерфейсом, отображающий данные о заказе.

###  Клиент
*Метод*: GET   
*URL*: /

Возвращаяет HtmlHandler клиента

###  Получение заказа
*Метод*: GET   
*URL*: /order

Возвращает инфорацию о заказе

# Service Configuration

Файл конфигурации называется configuration.json и должен находиться в одной директории с исполняемым файлом

```json
{
  "port": 8080,
  "DataBase": {
    "user": "admin",
    "password": "admin",
    "dataBaseName": "some_DB_name",
    "sslmode": "disable"
  }
}
```

# Database Schema

Требуемые таблицы

```sql
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
```