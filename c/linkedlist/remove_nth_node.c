// #19 删除链表倒数第 n 个节点
#include <stdio.h>

struct Node
{
    struct Node *next;
    int val;
};

// 1 2 3 4 5
// 1 2 3 5
// 就因为是倒数的，所以需要找到它前面的那个节点
// 如果是删除整数的第三个节点，同普通的删除一样

// 为什么快指针先移动 n 步，然后快慢同时移动
// 快指针指向末尾时，慢指针指向的就是要被删除节点
// 的前一个？
struct Node *remove_nth(struct Node *head, int nth)
{
    struct Node *slow, *fast;
    slow = head;
    fast = head;

    // 采用 -- 的方式，就不需要额外的变量用于计数
    while (nth)
    {
        fast = fast->next;
        nth--;
    }

    fast = fast->next;

    while (fast != NULL)
    {
        slow = slow->next;
        fast = fast->next;
    }

    slow->next = slow->next->next;
    return head;
}
