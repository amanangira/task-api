-- migrate:up
CREATE EXTENSION "uuid-ossp";

CREATE TABLE tasks (
    id              UUID DEFAULT UUID_GENERATE_V4() PRIMARY KEY,
    title           TEXT,
    description     TEXT,
    priority        TEXT,
    created_at      TIMESTAMP WITH TIME ZONE,
    updated_at      TIMESTAMP WITH TIME ZONE,
    due_at          TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE tasks;
    DROP EXTENSION "uuid-ossp";
