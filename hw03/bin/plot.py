#!/bin/python3
import numpy as np
import matplotlib.pyplot as plt

#DIR = "main_log/"
DIR = "new_log/"

y = []
for i in range(10):
    path = DIR + "log" + str(i+1) + ".txt" 
    f = open(path, "r")
    d = f.read()
    d = d.split('\n')
    d.pop()
    y.append(d)

f = []
for i in range(10):
    d = np.array(y[i])
    d = d.astype(np.float)
    f.append(d)

for i in f:
    plt.plot(i)

#plt.xlim(0,30)
plt.xlim(0,1600)
plt.show()
