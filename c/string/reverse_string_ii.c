// #541 反转字符串 ii
#include <stdio.h>
#include <string.h>

void main(void)
{
    char alpha[] = {'a', 'b', 'c', 'd', 'e', 'f', 'g'};
    int length = strlen(alpha);
    int k = 2;
    // 2k -> every 4 alpha reverse 2 from head

    for (int i = 0; i < length; i += (2 * k))
    {
        int reverse = length - i < k ? length - i : k;

        int left = i;
        // NOTE: should sub 1
        int right = i + reverse -1;
        while (left < right)
        {
            char single = alpha[left];
            alpha[left] = alpha[right];
            alpha[right] = single;

            left++;
            right--;
        }
    }

    for (size_t i = 0; i < length; i++)
    {
        printf("%c ", alpha[i]);
    }
}