-- Create or replace the burn_details_view
CREATE OR REPLACE VIEW burn_details_view AS
SELECT
    br.autarchy_id,
    br.auth_key_id AS author,
    b.id,
    b.title,
    b.map_picture,
    ST_X(geo_location)::float AS lat,
    ST_Y(geo_location)::float AS lon,
    b.has_aid_team,
    b.reason,
    b."type",
    b.begin_at,
    b.completed_at,
    brs.state,
    b.created_at,
    b.updated_at,
    b.deleted_at
FROM
    burn b
INNER JOIN burn_requests br 
    ON br.burn_id = b.id
INNER JOIN burn_requests_states brs
    ON brs.burn_id = b.id AND brs.auth_key_id = br.auth_key_id;

-- Create or replace the autarchy_details_view
CREATE OR REPLACE VIEW autarchy_details_view AS
SELECT
    ak.email,
    a.id,
    a.title,
    a.autarchy_avatar,
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
