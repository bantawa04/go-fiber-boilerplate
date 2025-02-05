CREATE TABLE IF NOT EXISTS todos 
(
    id          UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    title       VARCHAR(100) NOT NULL,
    status      VARCHAR(20)  NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMP    NOT NULL,
    updated_at  TIMESTAMP,
    deleted_at  TIMESTAMP,
    CONSTRAINT chk_todo_status CHECK (status IN ('pending', 'in_progress', 'completed'))
);