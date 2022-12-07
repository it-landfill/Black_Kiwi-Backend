import geopy.distance
import geojson
from genProbStat import generateUniformNoise

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


if __name__ == "__main__":
    pois = load_geoJSON("data/poi_list.geojson")
    pois = set_point_color(pois, "#0f0")
    #print(pois)
    #save_geoJSON("data/poi_list.geojson",pois)
    #print(get_closest([11.356172561645508, 44.4977297671644], pois))

    data = generate_random_geoJSON_point([11.343083, 44.494332], 5, 100)
    #print(data)
    data = update_nearest_poi(data, pois)
    save_geoJSON("data/randwrefwrdom.geojson", data)

    save_geoJSON("data/locAndPois.geojson",join_geoJSON(data, pois))
    
