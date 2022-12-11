import geopy.distance
import json
from genProbStat import generateUniformNoise
import numpy as np
import random
import pandas as pd
import geoJSON_utils

def load_df(path):
    return pd.read_csv(path, index_col=0)

def save_df(path, data):
    data.to_csv(path, index=True)

def save_json(path, data):
    with open(path, "w") as f:
        json.dump(data, f)

def load_json(path):
    with open(path, "r") as f:
        return json.load(f)

def get_distance(coords_1, coords_2):
    return geopy.distance.geodesic(coords_1, coords_2).m

def get_closest(coords, pois):
    min_dist = get_distance(coords["coordinates"], pois[0]["coordinates"])
    min_poi  = pois[0]

    for poi in pois:
        dist = get_distance(coords["coordinates"], poi["coordinates"])
        if dist < min_dist:
            min_dist = dist
            min_poi = poi

    return min_poi

def privacy_preservation_obj(real_pos_list, fake_pos_list):
    posLen = min(len(real_pos_list), len(fake_pos_list))
    tmpList = [get_distance(real_pos_list[x]["coordinates"], fake_pos_list[x]["coordinates"],) for x in range(posLen)]
    return tmpList

def privacy_preservation_arr(real_pos, fake_pos_list):
    tmpList = [get_distance(real_pos, fake_pos) for fake_pos in fake_pos_list]
    return tmpList

def quality_of_service(real_pos_list, fake_pos_list):
    posLen = min(len(real_pos_list), len(fake_pos_list))
    tmpList = [real_pos_list[x]["poiRank"]-fake_pos_list[x]["poiRank"] for x in range(posLen)]
    return np.square(tmpList)


if __name__ == "__main__":
    true_pos = load_json("data/true_pos.json")
    fake_pos = load_json("data/fake_uniform.json")
    #print(privacy_preservation(true_pos, fake_pos))
    print(np.square(quality_of_service(true_pos, fake_pos)).mean())


    
