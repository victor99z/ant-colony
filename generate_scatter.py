import numpy as np
from matplotlib import pyplot as plt

CSVData = open("output.csv")
Array2d_result = np.loadtxt(CSVData, delimiter=",")

print(Array2d_result)

# Set the figure size
plt.rcParams["figure.figsize"] = [7.50, 3.50]
plt.rcParams["figure.autolayout"] = False

# Random data of 100×3 dimension
# data = np.array(np.random.random((100, 3)))

# Scatter plot
plt.pcolormesh(Array2d_result, cmap='magma')

plt.title('Scatter plot')
plt.xlabel('x-axis')
plt.ylabel('y-axis')

# Display the plot
plt.show()