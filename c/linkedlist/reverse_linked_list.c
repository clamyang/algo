// # 206 反转链表
#include <stdio.h>

struct Node
{
    struct Node *next;
    int val;
};

struct Node *reverse_linked_list(struct Node *head)
{
    struct Node *pre, *cur, *next;
    cur = head;
    pre = NULL;
    while (cur != NULL)
    {
        next = cur->next;
        cur->next = pre;
        pre = cur;

        cur = next;
    }

    return pre;
}
