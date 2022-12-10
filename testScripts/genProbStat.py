#!/usr/bin/python

import matplotlib.pyplot as plt
import numpy as np
from random import random
from math import sqrt, pi, cos, sin

def radiusCalibration(oldMaxRadius, newMaxRadius, actualRadius):
    return (oldMaxRadius * actualRadius) / newMaxRadius


# min:     min = 0 e max = 2 con passo 0.2
# max:     min = 0 e max = 4 con passo 0.2
# mode:    min = 0 e max = 4 con passo 0.2         
def generateTriangularNoise(location, radius, min, mode, max, repeat):
    locations = []
    # Repeat the process n times
    for i in range(repeat):
        # Generate a random radius and angle
        randomRadius = radiusCalibration(radius, max, np.random.triangular(min, mode, max))
        randomAngle = random() * 2 * pi
        # Convert polar to cartesian
        x = kilometersToDegreesLatitude(randomRadius * cos(randomAngle))
        y = kilometersToDegreesLongitude(randomRadius * sin(randomAngle), x)
        # Add the location
        x += location[0]
        y += location[1]
        # Add the point to the list
        locations.append((x, y))
    return locations


# Range radius 0.1 ... 2
# mean:                min = 0 e max = 2 con passo 0.1
# standard_deviation:  min = 0.1 e max = 2 con passo 0.1
def generateGaussianNoise(location, radius, mean, standardDeviation, repeat):
    locations = []
    # Repeat the process n times
    for i in range(repeat):
        # Generate a random radius and angle
        randomRadius = radiusCalibration(radius, mean + 3 * standardDeviation, abs(np.random.normal(mean, standardDeviation)))
        randomAngle = random() * 2 * pi
        # Convert polar to cartesian
        x = kilometersToDegreesLatitude(randomRadius * cos(randomAngle))
        y = kilometersToDegreesLongitude(randomRadius * sin(randomAngle), x)
        # Add the location
        x += location[0]
        y += location[1]
        # Add the point to the list
        locations.append((x, y))
    return locations


# Range radius 0.1 ... 2
# Range lambda in = 0.5 e max = 10 con passo 0.5
def generatePoissonNoise(location, radius, poissLambda, repeat):
    locations = []
    L = np.exp(-poissLambda)
    # Repeat the process n times
    for i in range(repeat):
        k = 0
        p = 1
        while p > L:
            k += 1
            p *= random()
        # Generate a random radius and angle
        randomRadius = radiusCalibration(radius, poissLambda + (3 * sqrt(poissLambda)), k - 1)
        randomAngle = random() * 2 * pi
        # Convert polar to cartesian
        x = kilometersToDegreesLatitude(randomRadius * cos(randomAngle))
        y = kilometersToDegreesLongitude(randomRadius * sin(randomAngle), x)
        # Add the location
        x += location[0]
        y += location[1]
        # Add the point to the list
        locations.append((x, y))
    return locations
   

# Range radius 0.1 ... 2
def generateUniformNoise(location, radius, repeat):
    locations = []
    # Repeat the process n times
    for i in range(repeat):
        # Generate a random radius and angle
        randomRadius = random() * radius
        randomAngle = random() * 2 * pi
        # Convert polar to cartesian
        x = kilometersToDegreesLatitude(randomRadius * cos(randomAngle))
        y = kilometersToDegreesLongitude(randomRadius * sin(randomAngle), x)
        # Add the location
        x += location[0]
        y += location[1]
        # Add the point to the list
        locations.append((x, y))
    return locations
 

def centroidCalculation(locations, firstNElements = 0):
    if firstNElements == 0:
        firstNElements = len(locations)
    x, y = zip(*locations)
    return (sum(x[:firstNElements]) / firstNElements, sum(y[:firstNElements]) / firstNElements)

def centroidCalculationAll(locations):
    x, y = zip(*locations)
    return (sum(x) / len(locations), sum(y) / len(locations))

def plot():
    # Number of points
    n = 1000
    rad = 4
    cordinates = [0, 0]
    
    # locations = generateUniformNoise((cordinates[0], cordinates[1]), rad, n)
    locations = generatePoissonNoise((cordinates[0], cordinates[1]), rad, 10, n)
    # locations = generateGaussianNoise((cordinates[0], cordinates[1]), rad, 2, 2, n)
    # locations = generateTriangularNoise((cordinates[0], cordinates[1]), rad, 1, 4, 4, n)
    # Plot the points
    x, y = zip(*locations)
    plt.scatter(x, y, s=1, c='red', alpha=0.5)
    # Calculate the centroid
    centroid = centroidCalculation(locations, 10)
    plt.scatter(centroid[0], centroid[1], s=30, c='blue', alpha=1)
    centroid = centroidCalculation(locations, 100)
    plt.scatter(centroid[0], centroid[1], s=30, c='purple', alpha=1)
    centroid = centroidCalculation(locations, 1000)
    plt.scatter(centroid[0], centroid[1], s=30, c='orange', alpha=1)

    plt.scatter(cordinates[0], cordinates[1], s=50, c='green', alpha=0.5)
    # plt.set_title('Plot uniform random points in a disk')
    plt.xlabel('Longitude')
    plt.ylabel('Latitude')
    plt.show()


# Funzione che permette di trasformare da kilometri a gradi (da applicare alla latitudine)
def kilometersToDegreesLatitude(km):
    return km / 111.111


# Funzione che permette di trasformare da kilometri a gradi (da applicare alla longitudine)
def kilometersToDegreesLongitude(km, latitude):
    return km / (111.111 * cos(latitude * pi / 180))


if __name__ == '__main__':
    plot()
