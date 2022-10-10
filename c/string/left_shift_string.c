// #剑指offer 58 左旋字符串
#include <stdio.h>
#include <string.h>

void left_shitf(char str[], int shift);
void left_shitf_hard(char str[], int shift);

void main(void)
{
    char str[] = "abccd";
    int shift = 2;
    // left_shitf(str, shift);
    left_shitf_hard(str, shift);
}

// 暴力方法：想象成字符的移动即可
void left_shitf(char str[], int shift)
{
    int len = strlen(str);
    int temp = len - 1;
    for (int i = shift - 1; i >= 0; i--)
    {
        char single = str[i];
        for (int j = i; j < len - 1; j++)
        {
            str[j] = str[j + 1];
        }
        str[temp--] = single;
    }

    for (size_t i = 0; i < len; i++)
    {
        printf("%c", str[i]);
    }
}

// 通过多次反转实现，降低时间复杂度
void left_shitf_hard(char str[], int shift)
{
}
