create or replace view burn_details_view
as
select
		br.auth_key_id as author,
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
	inner join burn_requests br 
		on br.burn_id = b.id
	inner join burn_requests_states brs
		on brs.burn_id = b.id and brs.auth_key_id = br.auth_key_id 

create or replace view autarchy_details_view
as
select
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
	inner join auth_keys ak 
		on ak.id = a.auth_key_id