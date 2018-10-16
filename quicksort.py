def quicksort(array):
    if len(array) < 2:
        return array
    else:
        pivot = array[0]

        less = [i for i in arr[1:] if i <= pivot]

        greater = [i for i in arr[1:] if i > pivot]
    # 递归调用
    return quicksort(less) + [pivot] + quicksort(greater)

