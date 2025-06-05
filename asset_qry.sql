CREATE TABLE assets (
    id            INT AUTO_INCREMENT PRIMARY KEY,
    corporation   VARCHAR(255)    NOT NULL,
    department    VARCHAR(255)    NOT NULL,
    team          VARCHAR(255)    NOT NULL,
    name          VARCHAR(255)    NOT NULL,
    position      VARCHAR(255)    NOT NULL,
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
    asset_id      VARCHAR(255)    NOT NULL
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
    
    CONSTRAINT fk_sysinfo_asset
      FOREIGN KEY (asset_id)
      REFERENCES assets(id)
      ON DELETE CASCADE
);
