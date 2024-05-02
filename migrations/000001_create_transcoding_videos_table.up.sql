CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE transcode_videos
(
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    title         VARCHAR(100),
    is_downloaded BOOLEAN          DEFAULT false,
    is_transcoded BOOLEAN          DEFAULT false,
    is_uploaded   BOOLEAN          DEFAULT false,
    exception     TEXT,
    created_at    TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);