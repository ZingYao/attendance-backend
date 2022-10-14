create user attendance;
ALTER USER attendance WITH PASSWORD 'localpass';
comment on role attendance is '考勤数据库管理员';

create database attendance
    with owner attendance;
comment on database attendance is '考勤数据库';
