// #1047  删除字符串中的所有相邻重复项
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void main(void)
{
    char adjacent[] = "abbaef";
    char *stack = (char *)malloc(sizeof(char) * 10);

    // 这里 top 初始为 0 也是可以的，不过我认为给 1 更好理解
    // 0 作为缓冲，-1 超出了索引下标的范围有些难理解
    // 另外，在这里进行初始化操作，就不需要在代码中判断 stack
    // 是否为空了，简化了代码。

    // 注意：即使要判断，也不要这样写条件 result.empty() || result.back() != s
    // 应该写成 result.back() != s || result.empty()
    // 减少了查询 stack 为空的次数
    int top = 1;
    stack[top] = adjacent[0];

    for (int i = 1; i < strlen(adjacent); i++)
    {
        if (stack[top] == adjacent[i])
        {
            top--;
            continue;
        }
        else
        {
            stack[++top] = adjacent[i];
        }
    }

    // 最后的入栈顺序为 e f
    // 所以严格定义来不能直接遍历 stack 输出 ef
    // 是需要 pop 的方式 输出 fe 然后再进行反转才可。

    // 投机取巧的话，可以直接倒序遍历。

    // 关键在于如何确定最终字符串的长度
}