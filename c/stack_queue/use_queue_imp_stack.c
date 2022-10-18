// #225 使用队列实现栈的操作
#include <stdio.h>
#include <stdlib.h>

typedef struct
{
    int *arr;
    int head, tail;
} Stack;

void push(Stack *s, int x)
{
    s->arr[s->tail++] = x;
}

int pop(Stack *s)
{
    // 以下这一大段代码其实实现的就是
    // 除了最后一个元素，其余的先出队，再入队的操作
    for (int i = s->head; i < s->tail - 1; i++)
    {
        int head_val = s->arr[s->head];
        int head = s->head;
        int tail = s->tail;

        while (tail > 0)
        {
            // 1 2 3 4 5
            // 2 3 4 5 1
            s->arr[head] = s->arr[head + 1];
            head++;
            tail--;
        }

        s->arr[s->tail - 1] = head_val;
    }

    return s->arr[s->head++];
}

void main(void)
{
    Stack *s = (Stack *)malloc(sizeof(Stack));
    s->arr = (int *)malloc(sizeof(int) * 10);
    s->head = 0;
    s->tail = 0;

    push(s, 100);
    push(s, 101);
    push(s, 102);
    push(s, 103);
    push(s, 104);

    int val = pop(s);
    printf("%d %d\n", val, s->head);

    push(s, 105);
    val = pop(s);
    printf("%d %d\n", val, s->head);

    val = pop(s);
    printf("%d %d\n", val, s->head);
    val = pop(s);
    printf("%d %d\n", val, s->head);
    val = pop(s);
    printf("%d %d\n", val, s->head);
    val = pop(s);
    printf("%d %d\n", val, s->head);
}