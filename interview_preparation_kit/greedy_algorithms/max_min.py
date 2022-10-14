def maxMin(k, arr):
    arr.sort()
    min = 9223372036854775807
    for idx in range(len(arr)-k+1):
        dif = arr[idx+k-1] - arr[idx]
        if min > dif:
            min = dif
    return min