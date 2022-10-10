// #剑指offer 05 替换字符串
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char *replace_string(char statement[]);
char *replace_string_hard(char statement[]);

void main(void)
{
    char happy[] = "We are happy";
    replace_string(happy);
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

    int new_length = length + 3 * space_count - space_count;
    char *new_str = (char *)malloc(sizeof(char) * new_length);

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

    for (size_t i = 0; i < new_length; i++)
    {
        printf("%c ", new_str[i]);
    }

    return new_str;
}

char *replace_string_hard(char statement[])
{
}
