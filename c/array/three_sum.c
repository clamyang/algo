// #15 三数之和
#include <stdio.h>

void main(void)
{
    // 重点：需要有序
    int array[] = {-4, -1, -1, 1, 2, 3};
    int len = sizeof(array) / sizeof(array[0]);

    for (int i = 0; i < len; i++)
    {
        if (array[i] > 0)
        {
            return;
        }

        if (i > 0 && array[i] == array[i - 1])
        {
            continue;
        }

        int left = i + 1;
        int right = len - 1;
        while (left < right)
        {
            if (array[i] + array[left] + array[right] > 0)
            {
                right--;
            }
            else if (array[i] + array[left] + array[right] < 0)
            {
                left++;
            }
            else
            {
                printf("%2d %2d %2d\n", array[i], array[left], array[right]);
                while (right > left && array[right] == array[right - 1])
                    right--;
                while (right > left && array[left] == array[left + 1])
                    left++;

                right--;
                left++;
            }
        }
    }
}
