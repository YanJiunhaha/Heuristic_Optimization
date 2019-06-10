import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D

fig = plt.figure()
ax = Axes3D(fig)

dX = np.linspace(-5, 5, 100)
dY = np.linspace(-5, 5, 100)
X,Y = np.meshgrid(dX,dY)
Z = 4/((X-2)**2+(Y-2)**2+1)+3/((X-2)**2+(Y+2)**2+1)+2/((X+2)**2+(Y-2)**2+1)

max=0
x=0
y=0
for i in range(dX.size):
    for j in range(dY.size):
        if(max<Z[i][j]):
            max = Z[i][j]
            x=j
            y=i
    
print("x=",dX[x],",y=",dY[y],",z=",max)

ax.plot_surface(X,Y,Z, rstride=1, cstride=1, cmap=plt.get_cmap('rainbow'))


plt.show()
