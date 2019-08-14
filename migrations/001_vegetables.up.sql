create table vegetables (
	id serial,
	name text, 
	calories integer,
	created_at timestamptz not null default now(),
	updated_at timestamptz not null default now()
);

create unique index vegetables_id_unique on vegetables (id);