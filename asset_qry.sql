CREATE TABLE assets (
    id            INT AUTO_INCREMENT PRIMARY KEY,
    corporation   VARCHAR(255)    NOT NULL,
    department    VARCHAR(255)    NOT NULL,
    team          VARCHAR(255)    NOT NULL,
    name          VARCHAR(255)    NOT NULL,
    position      VARCHAR(255)    NOT NULL,
    location      VARCHAR(255)    NOT NULL,
    user_note     TEXT            NOT NULL,
    category      VARCHAR(255)    NOT NULL,
    item_type     VARCHAR(255)    NOT NULL,
    quantity      INT             NOT NULL,
    model         VARCHAR(255)    NOT NULL,
    manufacturer  VARCHAR(255)    NOT NULL,
    purchase_date DATE            NOT NULL,
    `usage`       VARCHAR(50)     NOT NULL,
    asset_note    TEXT            NOT NULL,
    asset_name    VARCHAR(255)    NOT NULL,
    asset_id      VARCHAR(255)    NOT NULL,

    -- 날짜 저장
    created_at        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at_unix   INT UNSIGNED NOT NULL DEFAULT (UNIX_TIMESTAMP())
);

CREATE TABLE system_info (
    id            INT AUTO_INCREMENT PRIMARY KEY,
    asset_id      INT NOT NULL,

    cpu           JSON       NOT NULL,
    motherboard   JSON       NOT NULL,
    ram           JSON       NOT NULL,
    gpu           JSON       NOT NULL,
    disk          JSON       NOT NULL,
    network       JSON       NOT NULL,

    -- 날짜 저장
    created_at        TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at_unix   INT UNSIGNED NOT NULL DEFAULT (UNIX_TIMESTAMP()),

    CONSTRAINT fk_sysinfo_asset
      FOREIGN KEY (asset_id)
      REFERENCES assets(id)
      ON DELETE CASCADE
);
CREATE TABLE windows_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    asset_id INT,
    caption VARCHAR(255),
    version VARCHAR(100),
    serial_number VARCHAR(255),
    product_key VARCHAR(255),
    partial_product_key VARCHAR(100),
    licensed INT,
    product_channel VARCHAR(255),
    kms_machine VARCHAR(255),
    grace_period_remaining INT,
    estimated_expire_date VARCHAR(50),
    cracked INT,
    FOREIGN KEY (asset_id) REFERENCES assets(id)
);
CREATE TABLE office_info (
    id INT AUTO_INCREMENT PRIMARY KEY,
    asset_id INT,
    name VARCHAR(255),
    version VARCHAR(100),
    partial_product_key VARCHAR(100),
    licensed INT,
    grace_period INT,
    estimated_expire_date VARCHAR(50),
    cracked INT,
    FOREIGN KEY (asset_id) REFERENCES assets(id)
);
