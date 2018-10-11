array = [1,2,85,4,8,9]

for i in range(len(array)):
  for j in range(i+1, len(array)):
    if array[j] > array[i]:
      array[j], array[i] = array[i], array[j]

print(array)