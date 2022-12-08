import utils
import genProbStat
import seaborn as sns
import matplotlib.pyplot as plt
import pandas as pd

def generate_random_pos(center, radius):
	return genProbStat.generateUniformNoise(center, radius, 1)

def analyze_params():
	# Parameters
	folder = "data/km4"
	sns.set_theme(style="ticks", palette="pastel")

	# Load data
	true_pos = utils.load_json("data/true_pos.json")
	fake_files = ["fake_uniform", "fake_poisson", "fake_gaussian", "fake_triangular"]

	fake_pos = {}	
	qos = {}
	pp = {}

	for file in fake_files:
		fake_pos[file] = utils.load_json(f"{folder}/{file}.json")
		pp[file] = utils.privacy_preservation(true_pos, fake_pos[file])
		qos[file] = utils.quality_of_service(true_pos, fake_pos[file])

	
	fig, axes = plt.subplots(1, 2, figsize=(18, 10))

	sns.boxplot(data=pd.DataFrame(pp), ax=axes[0]).set(title='Privacy Preservation')
	sns.boxplot(data=pd.DataFrame(qos), ax=axes[1]).set(title='Quality of Service')
	sns.despine(offset=10, trim=True)
	plt.show(block=True)

def analyze_radius():
	pp = utils.load_df("data/var_radius_pp.csv")
	qos = utils.load_df("data/var_radius_qos.csv")

	fig, axes = plt.subplots(1, 2, figsize=(18, 10))
	sns.lineplot(data=pp, x="radius", y="value",ax=axes[0], hue="type").set(title='Privacy Preservation')
	sns.lineplot(data=qos, x="radius", y="value",ax=axes[1], hue="type").set(title='Quality of Service')
	sns.despine(offset=10, trim=True)
	plt.show(block=True)


if __name__ == "__main__":
	#analyze_params()
	analyze_radius()
	
