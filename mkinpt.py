import sys
import numpy as np

n = int(sys.argv[1])
xs = np.arange(n**2, dtype='uint32')
np.random.shuffle(xs)

sys.stdout.write(str(n) + "\n")
for i in range(n):
    sys.stdout.write(" ".join(map(str, xs[i*n:(i+1)*n])))
    sys.stdout.write("\n")


