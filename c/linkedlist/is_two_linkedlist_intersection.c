// #160 判断两个链表是否相交
#include <stdio.h>
struct Node
{
    struct Node *next;
    int val;
};

void is_intersection(struct Node *h1, struct Node *h2)
{
    struct Node *cur1 = h1, *cur2 = h2;
    int len1 = 0, len2 = 0;

    while (cur1)
    {
        len1++;
        cur1 = cur1->next;
    }

    while (cur2)
    {
        len2++;
        cur2 = cur2->next;
    }

    int sub = 0;
    struct Node *slow, *fast;
    if (len1 > len2)
    {
        sub = len1 - len2;
        slow = h2;
        fast = h1;
    }
    else
    {
        sub = len2 - len1;
        slow = h1;
        fast = h2;
    }

    while (sub)
    {
        fast = fast->next;
        sub--;
    }

    while (fast != slow)
    {
        fast = fast->next;
        slow = slow->next;
    }
}