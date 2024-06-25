-- Create or replace the burn_details_view
CREATE OR REPLACE VIEW burn_details_view AS
with fetchStates as (
	select
		b.id,
		brs.state
	from burn_requests_states brs
	inner join burn b 
		on b.id = brs.burn_id
	order by brs.updated_at desc
)
select
	distinct on (b.id) 
	b.id,
    br.auth_key_id AS author,
    ak.email,
    ak.nif,
    u.user_name,
    u.profile_avatar,
    u.phone_code,
    u.phone_number,
    br.autarchy_id,
    b.title,
    b.map_picture,
    ST_X(geo_location)::float AS lat,
    ST_Y(geo_location)::float AS lon,
    b.has_aid_team,
    b.reason,
    b."type",
    b.address_street,
    b.address_number,
    b.address_zip_code,
    b.address_city,
    b.begin_at,
    b.completed_at,
    fs.state,
    b.created_at,
    b.updated_at,
    b.deleted_at
from
    burn b
INNER JOIN burn_requests br 
    ON br.burn_id = b.id
INNER JOIN fetchStates fs
	on fs.id = b.id
inner join auth_keys ak 
	on ak.id = br.auth_key_id 
inner join users u 
	on u.auth_key_id = ak.id 
;
-- Create or replace the autarchy_details_view
CREATE OR REPLACE VIEW autarchy_details_view AS
SELECT
    ak.email,
    ak.nif,
    a.id,
    a.title,
    a.profile_avatar,
    ST_X(geo_location)::float AS lat,
    ST_Y(geo_location)::float AS lon,
    a.phone_code,
    a.phone_number,
    a.address_street,
    a.address_number,
    a.address_zip_code,
    a.address_city,
    a.created_at,
    a.updated_at,
    a.deleted_at
FROM
    autarchy a
INNER JOIN auth_keys ak 
    ON ak.id = a.auth_key_id;