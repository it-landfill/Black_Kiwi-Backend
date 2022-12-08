import geopy.distance
import geojson
from genProbStat import generateUniformNoise
import numpy as np
import random

def save_geoJSON(path, data):
    with open(path, "w") as f:
        geojson.dump(data, f)


def load_geoJSON(path):
    with open(path, "r") as f:
        return geojson.load(f)


def get_distance(coords_1, coords_2):
    return geopy.distance.geodesic(coords_1, coords_2).m


def get_closest(coords, data):
    min_dist = get_distance(coords, data["features"][0]["geometry"]["coordinates"])
    closest = 0

    for i, poi in enumerate(data["features"]):
        dist = get_distance(coords, poi["geometry"]["coordinates"])
        # print(f"{i} - {dist}")
        if dist < min_dist:
            min_dist = dist
            closest = i

    return closest, data[closest], min_dist

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

def privacy_preservation(real_pos_list, fake_pos_list):
    posLen = min(len(real_pos_list), len(fake_pos_list))
    tmpList = [get_distance(real_pos_list[x]["geometry"]["coordinates"], fake_pos_list[x]["geometry"]["coordinates"],) for x in range(posLen)]
    print(tmpList)
    return np.square(tmpList).mean()

def quality_of_service(real_pos_list, fake_pos_list):
    posLen = min(len(real_pos_list), len(fake_pos_list))
    tmpList = [real_pos_list[x]["properties"]["nearest_poi"]["properties"]["rank"]-fake_pos_list[x]["properties"]["nearest_poi"]["properties"]["rank"] for x in range(posLen)]
    print(tmpList)
    return np.square(tmpList).mean()

if __name__ == "__main__":
    pois = load_geoJSON("data/poi_list.geojson")
    pois = set_point_color(pois, "#0f0")
    #print(pois)
    #save_geoJSON("data/poi_list.geojson",pois)
    #print(get_closest([11.356172561645508, 44.4977297671644], pois))

    #positions = generate_random_geoJSON_point([11.343083, 44.494332], 5, 50)
    positions = load_geoJSON("data/positions.geojson")
    positions = set_point_color(positions, "#00f")
    #print(positions)
    #positions = update_nearest_poi(positions, pois)
    #save_geoJSON("data/positions.geojson", positions)
    #save_geoJSON("data/locAndPois.geojson",join_geoJSON(positions, pois))
    fake_positions = {"type": "FeatureCollection", "features": []}

    for i, location in enumerate(positions["features"]):
        fake_pos = generateUniformNoise(location["geometry"]["coordinates"], 2, 1)
        fake_positions["features"].append(
            {
                "type": "Feature",
                "geometry": {"type": "Point", "coordinates": fake_pos[0]},
                "properties": {"id": i, "type": "fake"},
            }
        )
    fake_positions = update_nearest_poi(fake_positions, pois)
    fake_positions = set_point_color(fake_positions, "#f00")

    save_geoJSON("data/testData/group1/truePositions.geojson", positions)
    save_geoJSON("data/testData/group1/fakePositions.geojson", fake_positions)
    #save_geoJSON("data/trueFake.geojson", join_geoJSON(join_geoJSON(positions, pois), fake_positions))
    #print("PP\n")
    #print(privacy_preservation(positions["features"], fake_positions["features"]))
    #print("\nQOS\n")
    #print(quality_of_service(positions["features"], fake_positions["features"]))
    
