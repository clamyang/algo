#include <stdio.h>

struct Node
{
    int val;
    struct Node *Left;
    struct Node *Right;
};

// 前序
void pre_traverse(struct Node *node)
{
    if (node == NULL)
    {
        return;
    }
    printf("%d\n", node->val);
    pre_traverse(node->Left);
    pre_traverse(node->Right);
}

// 中序
void traverse(struct Node *node)
{
    if (node == NULL)
    {
        return;
    }

    traverse(node->Left);
    printf("%d\n", node->val);
    traverse(node->Right);
}

// 后序
void post_traverse(struct Node *node)
{
    if (node == NULL)
    {
        return;
    }

    traverse(node->Left);
    traverse(node->Right);
    printf("%d\n", node->val);
}

void main(void)
{
}
