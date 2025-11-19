CREATE TABLE user (
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
    CONSTRAINT fk_teacher_user FOREIGN KEY (user_id) REFERENCES user(id) ON UPDATE CASCADE ON DELETE SET NULL
);

CREATE TABLE schedule (
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

-- INSERT
INSERT INTO user (id, username, "password", "role", created_at, updated_at) VALUES(1, 'admin', '$2a$10$6CVXK6Ftu9R8YENYcXQBUObrXJV2FrXvD/wy6XPsgwEIGI6Az/YqW', 'ADMIN', '2025-11-19 15:13:04.959', '2025-11-19 15:13:07.992');
INSERT INTO user (id, username, "password", "role", created_at, updated_at) VALUES(2, '2025001', '$2a$10$6CVXK6Ftu9R8YENYcXQBUObrXJV2FrXvD/wy6XPsgwEIGI6Az/YqW', 'TEACHER', '2025-11-19 15:13:04.959', '2025-11-19 15:13:07.992');
INSERT INTO user (id, username, "password", "role", created_at, updated_at) VALUES(3, '2025002', '$2a$10$6CVXK6Ftu9R8YENYcXQBUObrXJV2FrXvD/wy6XPsgwEIGI6Az/YqW', 'TEACHER', '2025-11-19 15:13:04.959', '2025-11-19 15:13:07.992');
INSERT INTO user (id, username, "password", "role", created_at, updated_at) VALUES(4, '2025003', '$2a$10$6CVXK6Ftu9R8YENYcXQBUObrXJV2FrXvD/wy6XPsgwEIGI6Az/YqW', 'TEACHER', '2025-11-19 15:13:04.959', '2025-11-19 15:13:07.992');
INSERT INTO user (id, username, "password", "role", created_at, updated_at) VALUES(5, '2025004', '$2a$10$6CVXK6Ftu9R8YENYcXQBUObrXJV2FrXvD/wy6XPsgwEIGI6Az/YqW', 'TEACHER', '2025-11-19 15:13:04.959', '2025-11-19 15:13:07.992');

INSERT INTO teacher (id, user_id, "name", nik, created_at, updated_at) VALUES(1, 2, 'TONO. S.Pd', '2025001', '2025-11-19 15:16:59.698', '2025-11-19 15:16:59.698');
INSERT INTO teacher (id, user_id, "name", nik, created_at, updated_at) VALUES(2, 3, 'WATI. S.Pd', '2025002', '2025-11-19 15:16:59.698', '2025-11-19 15:16:59.698');
INSERT INTO teacher (id, user_id, "name", nik, created_at, updated_at) VALUES(3, 4, 'SUWARNO. S.Pd', '2025003', '2025-11-19 15:16:59.698', '2025-11-19 15:16:59.698');
INSERT INTO teacher (id, user_id, "name", nik, created_at, updated_at) VALUES(4, 5, 'JAJANG. S.Pd', '2025004', '2025-11-19 15:16:59.698', '2025-11-19 15:16:59.698');

INSERT INTO schedule (id, teacher_id, class_code, class_name, subject_code, "date", jam_ke, time_start, time_end, created_at, updated_at) VALUES(1, 1, 'XA01', 'X-A', 'CHEM', '2025-02-10', 2, '08:40:00', '09:40:00', '2025-11-19 15:21:14.793', '2025-11-19 15:21:17.593');