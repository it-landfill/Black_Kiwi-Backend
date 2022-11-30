-- Mobile Query

-- Richiesta POI - IN: category, minRank, pos (order by dist)
SELECT poi.id, poi.name, poi.rank, cat.name, st_asgeojson(poi.geom) as coordinates, st_distance(poi.geom,st_geogfromtext('POINT(11.3428 44.4939)')) as meters
FROM "black-kiwi_data".poi_list as poi
JOIN "black-kiwi_data".categories as cat on poi.category = cat.id
WHERE cat.name = 'park' and poi.rank<0
ORDER BY meters;

-- Admin query
-- Visualizzare POI disponibili
SELECT poi.id, poi.name, poi.rank, cat.name, st_asgeojson(poi.geom) as coordinates
FROM "black-kiwi_data".poi_list as poi
JOIN "black-kiwi_data".categories as cat on poi.category = cat.id;

-- Add a new POI
INSERT INTO "black-kiwi_data".poi_list (id, geom, name, category, rank) 
VALUES (DEFAULT, ST_SetSRID(ST_MakePoint(-71.1043443253471, 42.3150676015829),4326), 'AjejeBa', 4, -1)

-- Delete a POI
DELETE FROM "black-kiwi_data".poi_list WHERE id = 32

-- Edit a POI
UPDATE "black-kiwi_data".poi_list SET name = 'fwe', category = 4, rank = 3 WHERE id = 29

-- Visualizzare posizione utenti
SELECT *
FROM "black-kiwi_data".requests;

-- Visualizzare quartieri colorati su base di poi in esso contenuti
SELECT json_build_object('type', 'FeatureCollection', 'features', json_agg(ST_AsGeoJSON(quartieri.*)::json))
FROM (SELECT quartieri.nomequart as name, quartieri.geom, count(poi.id) as value
      FROM "black-kiwi_data"."quartieri-bologna" as quartieri
               JOIN "black-kiwi_data".poi_list as poi on st_within(poi.geom, quartieri.geom)
      GROUP BY quartieri.nomequart, quartieri.geom) as quartieri;
-- Visualizzare quartieri colorati su base di checkin utenti
-- TODO


-- Clustering spaziale
SELECT req.id, req.timestamp, req.rank, req.category, req.geom, st_clusterkmeans(req.geom, 2) OVER() as cid
FROM "black-kiwi_data"."Requests" as req;