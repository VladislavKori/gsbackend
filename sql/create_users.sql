create table
    users (
        id serial primary key,
        email varchar(80),
        password text,
        avatar_url text default 'https://fiverr-res.cloudinary.com/images/q_auto,f_auto/gigs/213245707/original/66a67e36fe8227d15c8c310cc112b60e74af5d6f/design-avatar-cartoon-for-business-gaming-social-media.jpg',
        current_delivery_address_id integer default null,
        created_at timestamptz not null default now ()
    );

insert into users (email, password) values ('test@example.com', 'helloworld123') returning id;