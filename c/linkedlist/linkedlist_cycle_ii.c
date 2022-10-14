// #142 找到环形链表入口
#include <stdio.h>

struct Node
{
    struct Node *next;
    int val;
};

struct Node *entrance(struct Node *head)
{
    struct Node *slow = head, *fast = head;

    while (fast != NULL && fast->next != NULL)
    {
        slow = slow->next;
        fast = fast->next->next;
        if (slow == fast)
        {
            while (head != slow)
            {
                head = head->next;
                slow = slow->next;
            }

            return head;
        }
    }

    return NULL;
}