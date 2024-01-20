-- migrate:up
CREATE EXTENSION "uuid-ossp";

CREATE TYPE task_priority AS ENUM ('p0', 'p1', 'p2', 'p3', 'p4', 'p5' );

CREATE TABLE tasks (
    id              UUID DEFAULT UUID_GENERATE_V4() PRIMARY KEY,
    title           TEXT,
    description     TEXT,
    priority        task_priority,
    created_at      TIMESTAMP WITH TIME ZONE,
    updated_at      TIMESTAMP WITH TIME ZONE,
    due_at          TIMESTAMP WITH TIME ZONE
);

-- migrate:down
DROP TABLE tasks;

DROP type task_priority;

DROP EXTENSION "uuid-ossp";
