CREATE TABLE tasks
(
    id         BIGINT NOT NULL AUTO_INCREMENT,
    uuid       VARCHAR(255) NOT NULL UNIQUE,
    host       VARCHAR(30)  NOT NULL,
    port       VARCHAR(30)  NOT NULL,
    command    VARCHAR(30)  NOT NULL,
    mode       VARCHAR(30)  NOT NULL,
    status     VARCHAR(30)  NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
)