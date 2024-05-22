create table tokens (
  id uuid not null,
  token text not null,
  type varchar(20) not null,
  expire_at timestamp not null,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp default null,
  primary key (id)
);
create table auth_keys (
  id uuid not null,
  password varchar(72) not null,
  salt varchar(64) not null,
  nif varchar(9) not null,
  email varchar(62) not null,
  is_account_enabled boolean default false,
  user_type int default 0,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp default null,
  primary key (id),
  check (length(nif) = 9)
);
create table users (
  id uuid not null,
  auth_key_id uuid not null,
  profile_avatar text not null,
  user_name varchar(50) not null,
  first_name varchar(50) not null,
  last_name varchar(50) not null,
  phone_code varchar(4) not null,
  phone_number varchar(9) not null,
  address_street varchar(95) not null,
  address_number int not null,
  address_zip_code varchar(8) not null,
  address_city varchar(35) not null,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp default null,
  primary key (id),
  check (length(phone_number) = 9),
  check (length(address_zip_code) = 8),
  constraint auth_key_fk foreign key (auth_key_id) references auth_keys (id)
);
create table burn (
  id uuid not null,
  title varchar(50) not null,
  map_picture text not null,
  -- address_street varchar(95) default null,
  -- address_number int default null,
  -- address_zip_code varchar(8) default null,
  -- address_city varchar(35) default null,
  geo_location GEOMETRY(Point, 4326) DEFAULT NULL,
  has_aid_team boolean default false,
  reason int not null,
  type int not null,
  begin_at timestamp not null,
  completed_at timestamp default null,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp default null,
  primary key(id)
  -- check (length(address_zip_code) = 8)
);
create table autarchy (
  id uuid not null,
  manager_id uuid not null,
  title varchar(50) not null,
  geo_location GEOMETRY(Point, 4326) DEFAULT NULL,
  address_street varchar(95) not null,
  address_number int not null,
  address_zip_code varchar(8) not null,
  address_city varchar(35) not null,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp default null,
  primary key (id),
  constraint manager_id_fk foreign key (manager_id) references auth_keys(id)
);
create table autarchy_employee (
  id uuid not null,
  employee_id uuid not null,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp default null,
  primary key (id),
  constraint employee_id_fk foreign key (employee_id) references auth_keys(id)
);
create table burn_requests (
  auth_key_id uuid not null,
  -- autarchy_id uuid not null,
  burn_id uuid not null,
  initial_propose text not null,
  accepted boolean default false,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp default null,
  -- primary key (auth_key_id, autarchy_id, burn_id),
  primary key (auth_key_id, burn_id),
  constraint auth_key_id_fk foreign key (auth_key_id) references auth_keys(id),
  -- constraint autarchy_fk foreign key (autarchy_id) references autarchy(id),
  constraint burn_fk foreign key (burn_id) references burn(id)
);
create table burn_requests_states (
  auth_key_id uuid not null,
  -- autarchy_id uuid not null,
  burn_id uuid not null,
  state int default 0,
  observation text not null,
  created_at timestamp default now(),
  updated_at timestamp default now(),
  deleted_at timestamp default null,
  -- primary key (auth_key_id, autarchy_id, burn_id),
  primary key (auth_key_id, burn_id),
  -- constraint autarchy_fk foreign key (autarchy_id) references autarchy(id),
  FOREIGN KEY (auth_key_id, burn_id) REFERENCES burn_requests(auth_key_id, burn_id)
);
