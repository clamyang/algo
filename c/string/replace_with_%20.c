// #剑指offer 05 替换字符串
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char *replace_string(char statement[]);
void replace_string_hard(char statement[]);

void main(void)
{
    char happy[] = "We are happy";
    // replace_string(happy);

    replace_string_hard(happy);
}

char *replace_string(char statement[])
{
    int length = strlen(statement);
    int space_count = 0;
    for (size_t i = 0; i < length; i++)
    {
        if (statement[i] == ' ')
        {
            space_count++;
        }
    }

    int new_length = length + 2 * space_count;
    char *new_str = (char *)malloc(sizeof(char) * new_length + 1);

    for (size_t i = 0, j = 0; i < length && j < new_length; i++)
    {
        if (statement[i] != ' ')
        {
            new_str[j++] = statement[i];
        }
        else
        {
            new_str[j++] = '%';
            new_str[j++] = '2';
            new_str[j++] = '0';
        }
    }

    new_str[new_length] = '\0';

    return new_str;
}

void replace_string_hard(char statement[])
{
    // 既然已经使用 malloc 分配了一个 O(n) 的字符串数组
    // 那直接同上面的方法相同就行了..
    // 以下内容是为了练习双指针的解题方法

    // 假设扩容好的字符数组如下
    char happy[] = "We are happy    ";
    int new_len = strlen(happy);
    int old_len = strlen(statement);

    for (int i = old_len, j = new_len; i > 0; i--, j--)
    {
        if (statement[i] != ' ')
        {
            happy[j] = statement[i];
        }
        else
        {
            happy[j--] = '0';
            happy[j--] = '2';
            happy[j] = '%';
        }
    }
}
