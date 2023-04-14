#!/bin/bash

mkdir tmp
cd tmp
curl https://archive.ics.uci.edu/ml/machine-learning-databases/housing/housing.data > housing.data
curl https://archive.ics.uci.edu/ml/machine-learning-databases/housing/housing.names > housing.names