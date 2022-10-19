// #239 滑动窗口内的最大值
#include <stdio.h>
#include <stdlib.h>

// 暴力方法..？ 感觉不算是暴力方法吧..
// 我这个时间复杂度是 O(n) 并不是 O(n * k)

// 思考片刻后，发现这种方式实现存在问题：
// 如果滑动窗口特别大就不好处理了。需要定义的指针多了，需要比较的元素多了
// 时间复杂度自然会变成 O(n * k)
int *max_val(int arr[])
{
    int *result = (int *)malloc(sizeof(int) * 5);
    for (int i = 0, j = 1, k = 2; k < 5; i++, j++, k++)
    {
        int val = max(arr[i], arr[j], arr[k]);
        result[i] = val;
    }

    return result;
}

int max(int i, int j, int k)
{
    int max = 0;
    if (i > j)
    {
        max = i;
    }
    else
    {
        max = j;
    }

    if (k > max)
    {
        max = k;
    }

    return max;
}

void main(void)
{
    int array[] = {3, 4, 2, 5, 1};

    int *result = max_val(&array);
    for (int i = 0; i < 3; i++)
    {
        printf("%d\n", result[i]);
    }
}