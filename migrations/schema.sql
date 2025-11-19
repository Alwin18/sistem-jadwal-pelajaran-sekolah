table user {
  id bigserial [not null, primary key]
  username string [not null]
  password text [not null]
  role string
  created_at timestamp
  updated_at timestamp
}

table teacher {
  id bigserial [not null, primary key]
  user_id int4 [null, ref: > user.id]
  name string
  nik string
  created_at timestamp
  updated_at timestamp
}

table schedules {
  id bigserial [not null, primary key]
  teacher_id int4 [not null, ref: > teacher.id]
  class_code string [not null]
  class_name string [not null]
  subject_code string [not null]
  date date
  jam_ke int4
  time_start time
  time_end time
  created_at timestamp
  updated_at timestamp
}
