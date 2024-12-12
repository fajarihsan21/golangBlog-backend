CREATE EXTENSION "uuid-ossp";

create table public.tb_users (
	id uuid default uuid_generate_v4() not null,
	username varchar(50) not null,
	"password" varchar not null,
	is_active bool not null,
	created_at timestamp,
	updated_at timestamp,
	constraint tb_users_pkey primary key (id),
	constraint tb_users_username_key unique (username)
);

CREATE TYPE user_roles AS ENUM ('admin', 'user');
create table public.tb_users_profile (
	id uuid default uuid_generate_v4() not null,
	user_id uuid not null,
	"name" varchar not null,
	role user_roles not null,
	email varchar null,
	phone varchar null,
	is_active bool not null,
	created_at timestamp,
	updated_at timestamp,
	constraint tb_users_profile_pkey primary key (id),
	constraint tb_users_profile_unique_key unique (user_id),
	constraint tb_users_profile_user_id_fkey foreign key (user_id) references public.tb_users(id) on delete cascade on update cascade
);