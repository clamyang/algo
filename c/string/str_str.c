// #28 找出字符串中第一个匹配的下标
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void get_next(int *next, char sub_str[]);
int find_sub(int *next, char str[], char sub_str[]);
void get_next_normal(int *next, char sub_str[]);
int find_sub_normal(int *next, char str[], char sub_str[]);

void main(void)
{
    char str[] = "aabaabaafa";
    char sub_str[] = "aabaaf";

    int str_len = strlen(str);
    int sub_strlen = strlen(sub_str);

    int *next = (int *)malloc(sizeof(int) * sub_strlen);

    // get_next(next, sub_str);
    // printf("%d\n", find_sub(next, str, sub_str));

    get_next_normal(next, sub_str);
    printf("%d\n", find_sub_normal(next, str, sub_str));

    for (int i = 0; i < sub_strlen; i++)
    {
        printf("%2d", next[i]);
    }
}

// 获取子串的最长公共前后缀表
// 采用统一减一的方式实现
void get_next(int *next, char sub_str[])
{
    int sub_strlen = strlen(sub_str);
    int j = -1;
    next[0] = j; // 当子串只有一个字符时为 0;

    for (int i = 1; i < sub_strlen; i++)
    {
        while (j >= 0 && sub_str[i] != sub_str[j + 1])
        {
            j = next[j];
        }
        if (sub_str[i] == sub_str[j + 1])
        {
            j++;
        }
        next[i] = j;
    }
}

int find_sub(int *next, char str[], char sub_str[])
{
    int j = -1; // 指向 next 数组
    for (int i = 0; i < strlen(str); i++)
    {
        while (j >= 0 && str[i] != sub_str[j + 1])
        {
            j = next[j];
        }

        if (str[i] == sub_str[j + 1])
        {
            j++;
        }

        if (j + 1 == strlen(sub_str))
        {
            return i - strlen(sub_str) + 1;
        }
    }

    return -1;
}

void get_next_normal(int *next, char sub_str[])
{
    int j = 0;
    next[0] = j;

    for (int i = 1; i < strlen(sub_str); i++)
    {
        while (j > 0 && sub_str[i] != sub_str[j])
        {
            j = next[j - 1];
        }

        if (sub_str[i] == sub_str[j])
        {
            j++;
        }
        next[i] = j;
    }
}

int find_sub_normal(int *next, char str[], char sub_str[])
{
    int j = 0;
    for (int i = 0; i < strlen(str); i++)
    {
        while (j > 0 && str[i] != sub_str[j])
        {
            j = next[j - 1];
        }

        if (str[i] == sub_str[j])
        {
            j++;
        }

        if (j == strlen(sub_str))
        {
            return i - strlen(sub_str) + 1;
        }
    }
    return -1;
}
