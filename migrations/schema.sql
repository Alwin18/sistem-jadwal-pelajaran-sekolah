CREATE TABLE "user" (
    id          BIGSERIAL PRIMARY KEY,
    username    VARCHAR NOT NULL,
    password    TEXT NOT NULL,
    role        VARCHAR,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP
);

CREATE TABLE teacher (
    id          BIGSERIAL PRIMARY KEY,
    user_id     INT,
    name        VARCHAR,
    nik         VARCHAR,
    created_at  TIMESTAMP,
    updated_at  TIMESTAMP,
    CONSTRAINT fk_teacher_user FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE schedules (
    id           BIGSERIAL PRIMARY KEY,
    teacher_id   INT NOT NULL,
    class_code   VARCHAR NOT NULL,
    class_name   VARCHAR NOT NULL,
    subject_code VARCHAR NOT NULL,
    date         DATE,
    jam_ke       INT,
    time_start   TIME,
    time_end     TIME,
    created_at   TIMESTAMP,
    updated_at   TIMESTAMP,
    CONSTRAINT fk_schedules_teacher FOREIGN KEY (teacher_id) REFERENCES teacher(id) ON UPDATE CASCADE ON DELETE CASCADE
);
