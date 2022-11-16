#include <stdio.h>

/* 
    回溯可以解决的问题：
    1.组合问题
    2.切割问题，例如给一个字符串有多少种切割方式
    3.子集问题，
    4.排列问题
    5.棋盘问题
*/ 

void main()
{
    int n = 4, k = 2;
    // 1-4 中，所有 2 位数字的组合
    for (int i = 1; i <= n; i++)
    {
        for (int j = i + 1; j <= n; j++)
        {
            printf("%d-%d\n", i, j);
        }
    }
}
