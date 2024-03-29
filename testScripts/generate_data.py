import utils
import genProbStat
import numpy as np
import pandas as pd
from tqdm import tqdm

def update_closest(pos, poi_list):
    closest = utils.get_closest({"coordinates": pos}, poi_list)
    return {"coordinates": pos, "poiRank": closest["rank"], "poiID": closest["id"]}

def generate_pos(n, center, radius, distrib_func, poi_list, dummies=False):
    pos = distrib_func(center, radius, n)

    if dummies:
        return pos

    pos_objs = [update_closest(x, poi_list) for x in pos]
    return pos_objs

def generate_true_pos(n):
    center = [11.343083, 44.494332]
    radius = 5
    file_path = "data/true_pos.json"

    true_pos = generate_pos(n, center, radius, genProbStat.generateUniformNoise, utils.load_json("data/poi_list.json"))
    for ind, _ in enumerate(true_pos):
        true_pos[ind]["id"] = ind

    utils.save_json(file_path, true_pos)

def generate_fake_pos(true_pos_path, radius, distrib_func, file_name):
    true_pos = utils.load_json(true_pos_path)
    poi_list = utils.load_json("data/poi_list.json")

    fake_pos = []
    for pos in true_pos:
        tmp = generate_pos(1 , pos["coordinates"], radius, distrib_func, poi_list)[0]
        tmp["id"] = pos["id"]
        fake_pos.append(tmp)

    utils.save_json(f"data/{file_name}", fake_pos)

def poisson(center, radius, n):
    return genProbStat.generatePoissonNoise(center, radius, 5, n)

def gaussian(center, radius, n):
    return genProbStat.generateGaussianNoise(center, radius, 1, 1, n)

def triangular(center, radius, n):
    return genProbStat.generateTriangularNoise(center, radius, 0, 2, 4, n)

def var_radius_data(n, min, max, divider):
    poi_list = utils.load_json("data/poi_list.json")

    center = [11.343083, 44.494332]
    closest = utils.get_closest({"coordinates": center}, poi_list)
    true_pos = [{"coordinates": center, "poiRank": closest["rank"]} for _ in range(n)]

    pp = []
    qos = []

    # Adjust the max radius to be a multiple of the radius divider
    max = max*divider

    for radius in range(min, max):
        uni = generate_pos(n, center, radius/divider, genProbStat.generateUniformNoise, poi_list)
        poi = generate_pos(n, center, radius/divider, poisson, poi_list)
        gau = generate_pos(n, center, radius/divider, gaussian, poi_list)
        tri = generate_pos(n, center, radius/divider, triangular, poi_list)

        uni_pp = utils.privacy_preservation_obj(true_pos, uni)
        poi_pp = utils.privacy_preservation_obj(true_pos, poi)
        gau_pp = utils.privacy_preservation_obj(true_pos, gau)
        tri_pp = utils.privacy_preservation_obj(true_pos, tri)

        uni_qos = utils.quality_of_service(true_pos, uni)
        poi_qos = utils.quality_of_service(true_pos, poi)
        gau_qos = utils.quality_of_service(true_pos, gau)
        tri_qos = utils.quality_of_service(true_pos, tri)

        for i in range(n):
            pp.append({"radius": radius/divider, "value": uni_pp[i], "type": "uniform"})
            pp.append({"radius": radius/divider, "value": poi_pp[i], "type": "poisson"})
            pp.append({"radius": radius/divider, "value": gau_pp[i], "type": "gaussian"})
            pp.append({"radius": radius/divider, "value": tri_pp[i], "type": "triangular"})

            qos.append({"radius": radius/divider, "value": uni_qos[i], "type": "uniform"})
            qos.append({"radius": radius/divider, "value": poi_qos[i], "type": "poisson"})
            qos.append({"radius": radius/divider, "value": gau_qos[i], "type": "gaussian"})
            qos.append({"radius": radius/divider, "value": tri_qos[i], "type": "triangular"})

    pp_df = pd.DataFrame(pp)
    utils.save_df("data/var_radius_pp.csv", pp_df)

    qos_df = pd.DataFrame(qos)
    utils.save_df("data/var_radius_qos.csv", qos_df)

def var_dummies_data(n, dummies, radius):
    center = [11.343083, 44.494332]

    pp = []

    pbar = tqdm(range(n))
    for _ in pbar:
        for dummies_num in range(2, dummies+1):
            pbar.set_description(f"Generating data for {dummies_num} dummies")
            uni = generate_pos(dummies_num, center, radius, genProbStat.generateUniformNoise, None, True)
            poi = generate_pos(dummies_num, center, radius, poisson, None, True)
            gau = generate_pos(dummies_num, center, radius, gaussian, None, True)
            tri = generate_pos(dummies_num, center, radius, triangular, None, True)

            uni_centroid = genProbStat.centroidCalculationAll(uni)
            poi_centroid = genProbStat.centroidCalculationAll(poi)
            gau_centroid = genProbStat.centroidCalculationAll(gau)
            tri_centroid = genProbStat.centroidCalculationAll(tri)

            uni_pp = utils.get_distance(center, uni_centroid)
            poi_pp = utils.get_distance(center, poi_centroid)
            gau_pp = utils.get_distance(center, gau_centroid)
            tri_pp = utils.get_distance(center, tri_centroid)

            pp.append({"dummies": dummies_num, "value": uni_pp, "type": "uniform"})
            pp.append({"dummies": dummies_num, "value": poi_pp, "type": "poisson"})
            pp.append({"dummies": dummies_num, "value": gau_pp, "type": "gaussian"})
            pp.append({"dummies": dummies_num, "value": tri_pp, "type": "triangular"})

    pp_df = pd.DataFrame(pp)
    utils.save_df("data/var_dummies_pp.csv", pp_df)


if __name__ == "__main__":
    #generate_true_pos(100)
    #generate_fake_pos("data/true_pos.json", 4, genProbStat.generateUniformNoise, "km4/fake_uniform.json")
    #generate_fake_pos("data/true_pos.json", 4, poisson, "km4/fake_poisson.json")
    #generate_fake_pos("data/true_pos.json", 4, gaussian, "km4/fake_gaussian.json")
    #generate_fake_pos("data/true_pos.json", 4, triangular, "km4/fake_triangular.json")
    #var_radius_data(100, 0, 4, 4)
    var_dummies_data(n=200, dummies=1000, radius=2)
