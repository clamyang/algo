// #459 重复的子字符串
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

void get_next(int *next, char str[]);

// 结论：除去最长相等前后缀的部分，如果原串对他取余为0
// 就说明该字符串是由它的子串组成
void main(void)
{
    char str[] = "abababc";

    int next_len = strlen(str);
    int *next = (int *)malloc(sizeof(int) * next_len);
    get_next(next, str);

    if ((next[next_len - 1] + 1) != -1 && (next_len % (next_len - (next[next_len - 1] + 1))))
    {
        printf("%d\n", 1);
        return;
    }

    printf("%d\n", 0);
}

void get_next(int *next, char str[])
{
    int j = -1;
    next[0] = j;

    for (int i = 1; i < strlen(str); i++)
    {
        while (j >= 0 && str[i] != str[j + 1])
        {
            j = next[j];
        }
        if (str[i] == str[j + 1])
        {
            j++;
        }
        next[i] = j;
    }
}
