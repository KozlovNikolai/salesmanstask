#! python3

from ctypes import *
import numpy as np
import random
import matplotlib.pyplot as plt
import networkx as nx
from datetime import datetime
#------------------------------------------------------------------------------------
def distance(x1, y1, x2, y2):
    return ((x2 - x1) ** 2 + (y2 - y1) ** 2) ** 0.5
#------------------------------------------------------------------------------------
def tsp_branch(n, py_arr, lib):
    if n < 2:
        return {}
    flatten_arr =  list(np.concatenate(py_arr).flat)
    l = [-1] * (n * n - len(flatten_arr))
    flatten_arr = (flatten_arr + l)[:n * n:]
    int_arr =  (c_int  * (n * n))(*flatten_arr)      
    res = lib.tsp_branch(n, byref(int_arr))
    if res > 0:
        l = list(int_arr)[:res:]
        return {'len' : l.pop(0), 'steps' : l.pop(0), 'path' : l}
    else:
        return {}
#------------------------------------------------------------------------------------
lib_c = cdll.LoadLibrary(r"tsp_branch.dll")
lib_c.tsp_branch.argtypes = [c_int, c_void_p]
lib_c.tsp_branch.restype = c_int
#------------------------------------------------------------------------------------
INF = -1

#random.seed(1)

n = 26

v1 = []
points = {}
for i in range(n):
    points[i] = (random.randint(1, 10000), random.randint(1, 10000))

input_matrix = []
for i, vi in points.items(): 
    m1 = []
    for j, vj in points.items():
        if i == j:
             m1.append(INF)
        else:
            m1.append(int(distance(vi[0], vi[1], vj[0], vj[1])))
            v1.append([i, j, int(distance(vi[0], vi[1], vj[0], vj[1]))])
    input_matrix.append(m1.copy()) 

plt.figure(figsize=(6, 6))

graph = nx.Graph()
graph.add_nodes_from(points) 

# Добавляем дуги в граф
for i in v1:
    graph.add_edge(i[0], i[1], weight=i[2])
#-----------------------------------------------------------------
start_time = datetime.now()
res1 = tsp_branch(n, input_matrix, lib_c)
print(datetime.now() - start_time)

print('min_len =', res1['len'], ', steps =', res1['steps'], res1['path'])

if 'path' in res1:
    d1 = []
    for i, v in enumerate(res1['path']):
        d1.append([res1['path'][i-1], res1['path'][i]])
#-----------------------------------------------------------------
# Рисуем всё древо
nx.draw(graph, points, width=0.5, edge_color="#C0C0C0", with_labels=True)
nx.draw(graph, points, width=3, edge_color="red", edgelist=d1, style="-")