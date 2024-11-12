CREATE TABLE files
(
    md5_id     CHAR(32) PRIMARY KEY DEFAULT gen_random_uuid(),
    filename   VARCHAR(255)         DEFAULT '<unnamed>',
    size       BIGINT      NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT current_timestamp
);

CREATE TABLE tags
(
    uuid       UUID PRIMARY KEY      DEFAULT gen_random_uuid(),
    tag_name   VARCHAR(255) NOT NULL,
    color      CHAR(7)      NOT NULL,
    created_at TIMESTAMPTZ  NOT NULL DEFAULT current_timestamp
);

CREATE TABLE file_tags
(
    file_id CHAR(32) NOT NULL,
    tag_id  UUID     NOT NULL,
    PRIMARY KEY (file_id, tag_id),
    -- CASCADE deleting all deleted files accordingly --
    FOREIGN KEY (file_id) REFERENCES files (md5_id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (uuid) ON DELETE CASCADE
);

-- Create two more indices --
CREATE INDEX idx_file_tags_file_id ON file_tags (file_id);
CREATE INDEX idx_file_tags_tag_id ON file_tags (tag_id);
