#!/bin/bash

# Create Members Table
docker exec -it postgres psql -U docker acm -c "create table if not exists members (ID serial primary key not null, student_id varchar(15) not null, first_name varchar(255) not null, last_name varchar(255) not null, email varchar(255) not null, password varchar(255) not null, year varchar(30), github varchar(255), linkedin varchar(255), personal_website varchar(255), stack_overflow varchar(255), portfolium varchar(255), handshake varchar(255), slack varchar(50), discord varchar(50), thumbnail bytea null, active boolean, banned boolean, privilege varchar(50), created_at TIMESTAMPTZ default NOW());"

# Create Workshops Table
docker exec -it tuolumne php artisan migrate