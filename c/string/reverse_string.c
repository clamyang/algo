// #344 反转字符串
#include <stdio.h>

void main(void)
{
    char hello[] = {'h', 'e', 'l', 'l', 'o'};

    int length = 5;
    for (int i = 0, j = length - 1; i < j; i++, j--)
    {
        char single = hello[i];

        hello[i] = hello[j];
        hello[j] = single;
    }

    for (size_t i = 0; i < length; i++)
    {
        printf("%c ", hello[i]);
    }
}