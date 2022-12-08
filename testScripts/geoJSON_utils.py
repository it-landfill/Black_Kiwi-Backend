

def save_geoJSON(path, data):
    with open(path, "w") as f:
        geojson.dump(data, f)


def load_geoJSON(path):
    with open(path, "r") as f:
        return geojson.load(f)

def generate_random_geoJSON_point(center, radius, repeat):
    locations = generateUniformNoise(center, radius, repeat)

    features = {"type": "FeatureCollection", "features": []}

    for i, location in enumerate(locations):
        features["features"].append(
            {
                "type": "Feature",
                "geometry": {"type": "Point", "coordinates": location},
                "properties": {"id": i, "marker-color": "#00f"},
            }
        )

    return features

def update_nearest_poi(positions, pois):

    for coord in positions["features"]:
        _, closest, dist = get_closest(coord["geometry"]["coordinates"], pois)
        coord["properties"]["nearest_poi"] = closest
        coord["properties"]["distance"] = dist

    return positions

def set_point_color(features, color_str):
    for feature in features["features"]:
        feature["properties"]["marker-color"] = color_str

    return features

def join_geoJSON(data1, data2):
    return {"type": "FeatureCollection", "features": data1["features"] + data2["features"]}

def list_to_geoJSON(lst, cat=None):
    positions = {"type": "FeatureCollection", "features": []}

    for i, location in enumerate(lst):
        positions["features"].append(
            {
                "type": "Feature",
                "geometry": {"type": "Point", "coordinates": location},
                "properties": {"id": i, "type": cat},
            }
        )
    
    return positions

if __name__ == "__main__":
    pass
    #pois = geoJSON_utils.load_geoJSON("data/poi_list.geojson")
    #pois = set_point_color(pois, "#0f0")
    #print(pois)
    #save_geoJSON("data/poi_list.geojson",pois)
    #print(get_closest([11.356172561645508, 44.4977297671644], pois))

    #positions = generate_random_geoJSON_point([11.343083, 44.494332], 5, 50)
    #positions = load_geoJSON("data/positions.geojson")
    #positions = set_point_color(positions, "#00f")
    #print(positions)
    #positions = update_nearest_poi(positions, pois)
    #save_geoJSON("data/positions.geojson", positions)
    #save_geoJSON("data/locAndPois.geojson",join_geoJSON(positions, pois))
    #fake_positions = {"type": "FeatureCollection", "features": []}

    #for i, location in enumerate(positions["features"]):
    #    fake_pos = generateUniformNoise(location["geometry"]["coordinates"], 2, 1)
    #    fake_positions["features"].append(
    #        {
    #            "type": "Feature",
    #            "geometry": {"type": "Point", "coordinates": fake_pos[0]},
    #            "properties": {"id": i, "type": "fake"},
    #        }
    #    )
    #fake_positions = update_nearest_poi(fake_positions, pois)
    #fake_positions = set_point_color(fake_positions, "#f00")
    #save_geoJSON("data/testData/group1/truePositions.geojson", positions)
    #save_geoJSON("data/testData/group1/fakePositions.geojson", fake_positions)
    #save_geoJSON("data/trueFake.geojson", join_geoJSON(join_geoJSON(positions, pois), fake_positions))
    #print("PP\n")
    #print(privacy_preservation(positions["features"], fake_positions["features"]))
    #print("\nQOS\n")
    #print(quality_of_service(positions["features"], fake_positions["features"]))
    #clean_poi = [{"lat": x["geometry"]["coordinates"][0], "lon": x["geometry"]["coordinates"][1], "rank": x["properties"]["rank"]} for x in pois["features"]]
    #print(df.to_dict('index'))
    #save_json("data/poi_list.json", [{"id": ind, "coordinates": [val["lat"], val["lon"]], "rank":val["rank"]} for ind, val in df.to_dict('index').items()])
    #save_df("data/poi_list.csv", df)
    #print(get_closest({"lat": 11.356172561645508, "lon": 44.4977297671644}, df))