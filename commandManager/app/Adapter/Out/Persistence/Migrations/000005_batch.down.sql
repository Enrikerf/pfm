

ALTER TABLE results
    DROP INDEX IDX_RESULTS_C_BATCH_ID,
    DROP FOREIGN KEY FK_RESULTS_C_BATCH_ID,
    ADD COLUMN task_id BIGINT      NOT NULL,
    ADD COLUMN task_uuid  VARCHAR(30) NOT NULL,
    DROP COLUMN batch_id,
    DROP COLUMN batch_uuid,
    ADD INDEX IDX_RESULTS_C_TASK_ID (task_id),
    ADD CONSTRAINT FK_RESULTS_C_TASK_ID FOREIGN KEY (task_id) REFERENCES tasks(id)
;

DROP TABLE batches;