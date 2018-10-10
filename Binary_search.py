# 该函数返回的是 元素 在列表中的索引

def binary_search(list, item):

    # low and high 代表索引
    low = 0
    high = len(list)-1

    while low <= high:
        mid = int((low+high)/2)
        guess = list[mid]

        if guess == item:
            return mid

        if guess < item:

            # +1 是把 mid 排除出去 因为已经知道 mid 不在范围内
            low = mid + 1
        else:

            # -1 也是为了把 mid 排除出去
            high = mid - 1

    return None

my_list = [1, 3, 5, 7, 9,10,11,23]

print (binary_search(my_list, 11))
