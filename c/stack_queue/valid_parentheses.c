// #20 有效的括号
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>

bool parentheses_v1(char list[])
{
    char *stack = (char *)malloc(sizeof(char) * 3);
    int left_count = 0;

    // {)}(  --> } false
    // {()}  --> })
    for (int i = 0; i < strlen(list); i++)
    {
        switch (list[i])
        {
        case '{':
            stack[left_count++] = '}';
            break;
        case '(':
            stack[left_count++] = ')';
            break;
        case '[':
            stack[left_count++] = ']';
            break;
        case '}':
        case ')':
        case ']':
            if (stack[--left_count] != list[i])
            {
                return false;
            }
            break;
        default:
            break;
        }
    }

    if (left_count != 0)
    {
        return false;
    }

    return true;
}

// 优化，能否在 for 循环结束前就判断出不满足
bool parentheses_v2(char list[])
{
    char *stack = (char *)malloc(sizeof(char) * 3);
    int left_count = 0;

    // {)}(  --> }
    // {()}  --> })
    // {{}   --> }}
    // {}}   --> }
    for (int i = 0; i < strlen(list); i++)
    {
        switch (list[i])
        {
        case '{':
            stack[left_count++] = '}';
            break;
        case '(':
            stack[left_count++] = ')';
            break;
        case '[':
            stack[left_count++] = ']';
            break;
        default:
            // 必须把 != 的条件放在后边，因为会出现索引为 -1 的情况；
            if (left_count == 0 || stack[left_count - 1] != list[i])
            {
                return false;
            }
            left_count--;
            break;
        }
    }

    return left_count == 0;
}

// 当存在 )( 情况，该算法无法正确执行
bool badness(char list[])
{
    int left1 = 0, left2 = 0, left3 = 0;
    int right1 = 0, right2 = 0, right3 = 0;

    for (int i = 0; i < strlen(list); i++)
    {
        switch (list[i])
        {
        case '{':
            left1++;
            break;
        case '(':
            left2++;
            break;
        case '[':
            left3++;
            break;
        case '}':
            right1++;
            break;
        case ')':
            right2++;
            break;
        case ']':
            right3++;
            break;
        default:
            break;
        }
    }

    if (left1 != right1 || left2 != right2 || left3 != right3)
    {
        return false;
    }

    return true;
}

void main(void)
{
    // true == 1
    char list[] = "{(}}";
    // printf("%d\n", badness(list));
    printf("%d\n", parentheses_v1(list));
    printf("%d\n", parentheses_v2(list));
}
