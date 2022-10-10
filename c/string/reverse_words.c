// #151 反转字符串中的单词
// are you ok -> ok you are

#include <stdio.h>
#include <string.h>

void reverse_word(char statement[]);

void main(void)
{
    char statement[] = "are you ready to dance";
    reverse_word(statement);
}

// 垃圾方法：使用 split 分割，然后反转
// 不在此实现

// 厉害方法：空间复杂度 O(1)
void reverse_word(char statement[])
{
    for (int i = 0, j = strlen(statement) - 1; i < j; i++, j--)
    {
        char single = statement[j];
        statement[j] = statement[i];
        statement[i] = single;
    }

    for (size_t i = 0; i < strlen(statement); i++)
    {
        printf("%c", statement[i]);
    }

    printf("\n");

    // +1 means '\0'
    for (int i = 0, j = 0; j < strlen(statement) + 1; j++)
    {
        if (statement[j] != ' ' && statement[j] != '\0')
        {
            continue;
        }
        else
        {
            int temp = j - 1;
            while (i < temp)
            {
                char single = statement[temp];
                statement[temp--] = statement[i];
                statement[i++] = single;
            }
            i = j + 1;
        }
    }

    for (size_t i = 0; i < strlen(statement); i++)
    {
        printf("%c", statement[i]);
    }
}
