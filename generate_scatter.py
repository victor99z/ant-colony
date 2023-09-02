import numpy as np
from matplotlib import pyplot as plt

DataOutput = open("output.csv")
DataInput = open("input.csv")

Array2d_input= np.loadtxt(DataInput, delimiter=",")
Array2d_result = np.loadtxt(DataOutput, delimiter=",")

fig = plt.figure()

ax1 = fig.add_subplot(121)
ax2 = fig.add_subplot(122)

# data_1=np.array(np.random.random((10,2)))*10
# data_2=np.array(np.random.random((10,2)))

ax1.pcolormesh(Array2d_input)
ax2.pcolormesh(Array2d_result)

ax1.set_title('Input')
ax1.set_xlabel('x')
ax1.set_ylabel('y')

ax2.set_title('Output')
ax2.set_xlabel('x')
ax2.set_ylabel('y')

plt.show()


# # Set the figure size
# plt.rcParams["figure.figsize"] = [7.50, 3.50]
# plt.rcParams["figure.autolayout"] = False

# # Random data of 100×3 dimension
# # data = np.array(np.random.random((100, 3)))

# # Scatter plot
# plt.pcolormesh(Array2d_input, cmap='magma')

# plt.pcolormesh(Array2d_result, cmap='magma')


# plt.title('Scatter plot')
# plt.xlabel('x-axis')
# plt.ylabel('y-axis')

# # Display the plot
# plt.show()