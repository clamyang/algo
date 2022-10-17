#include <stdio.h>
#include <stdlib.h>

typedef struct
{
    int *in, *out;
    int in_top;
    int out_top;
} Queue;

// push
void push(Queue *q, int x)
{
    q->in[q->in_top++] = x;
}

int empty(Queue *q)
{
    if (q->in_top == 0 && q->out_top == 0)
    {
        return 1;
    }
    return 0;
}

// pop
int pop(Queue *q)
{
    if (empty(q))
    {
        return 0;
    }

    if (q->out_top == 0)
    {
        while (q->in_top > 0)
        {
            q->out[q->out_top++] = q->in[--q->in_top];
        }
    }
    return q->out[--q->out_top];
}

void main(void)
{

    Queue *q = (Queue *)malloc(sizeof(Queue));
    q->in = (int *)malloc(sizeof(int) * 6);
    q->out = (int *)malloc(sizeof(int) * 6);
    q->in_top = 0;
    q->out_top = 0;

    push(q, 1);
    push(q, 2);
    push(q, 3);
    push(q, 4);
    push(q, 5);

    // 采用两个栈实现如下
    // 第一个： 栈顶->栈底 5 4 3 2 1
    // 第二个： 栈顶->栈底 1 2 3 4 5

    // 很明显 第一个栈的第一个元素当作队列的尾部，
    // 第二个栈的第一个元素当作队列的头部，
    // 队列只有两个操作，入队、出队，并且遵循 FIFO 原则
    // 从队尾插入，从队头去除
    // 就对应从第一个栈 压入
    //       从第二个栈 弹出

    printf("%d\n", pop(q));
    printf("%d\n", pop(q));
    printf("%d\n", pop(q));
    printf("%d\n", pop(q));
    printf("%d\n", pop(q));

    return;
}