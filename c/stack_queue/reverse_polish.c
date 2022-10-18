// #150 逆波兰表达式
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <ctype.h>

// 其实整体思路是对的，但是实现的时候，isdigit
// 没有像预期那样执行
long reverse_polish(char token[])
{
    long *stack = (long *)malloc(sizeof(long) * 10);
    long top = 0;

    stack[top] = token[0];

    for (int i = 1; i < strlen(token); i++)
    {
        int val = atoi(&token[i]);

        if (isdigit(val))
        {
            stack[++top] = val;
        }
        else
        {
            int target1 = stack[top - 1];
            int target2 = stack[top];

            switch (token[i])
            {
            case '+':
                stack[top - 1] = target1 + target2;
                break;
            case '-':
                stack[top - 1] = target1 - target2;
                break;
            case '*':
                stack[top - 1] = target1 * target2;
                break;
            case '/':
                stack[top - 1] = target1 / target2;
                break;
            default:
                break;
            }
            top--;
        }
    }
    return stack[0];
}

void main(void)
{
    char token[] = {'2', '1', '+', '3', '*'};
    for (int i = 0; i < strlen(token); i++)
    {
        printf("%c ", token[i]);
    }
    printf("%d\n", reverse_polish(token));
}