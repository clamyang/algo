// #59 螺旋矩阵 II
#include <stdio.h>

void spiral(int num);

int main(int argc, char *argv[])
{
    spiral(10);
    return 0;
}

void spiral(int num)
{
    int ans[num][num];

    int loop = num / 2;
    int offset, count;
    int start_x, start_y;

    offset = count = 1;
    start_x = start_y = 0;

    while (loop)
    {
        int i = start_x;
        int j = start_y;

        // top
        for (; j < start_y + num - offset; j++)
        {
            ans[start_x][j] = count++;
        }

        // right
        for (; i < start_x + num - offset; i++)
        {
            ans[i][j] = count++;
        }

        // bottom
        for (; j > start_y; j--)
        {
            ans[i][j] = count++;
        }

        // left
        for (; i > start_x; i--)
        {
            ans[i][j] = count++;
        }

        offset += 2;
        start_x++;
        start_y++;
        loop--;
    }

    if (num % 2 == 1)
        ans[num / 2][num / 2] = num * num;

    for (int i = 0; i < num; i++)
    {
        for (int j = 0; j < num; j++)
        {
            printf("%3d ", ans[i][j]);
        }
        printf("\n");
    }
}
