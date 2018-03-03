#include <iostream>
#include <queue>

using namespace std;

typedef struct Node {
    int value;
    Node * left;
    Node * right;
};

void PrintNodeByLevel(Node *root) {
    queue<Node *> qu;
    qu.push(root);

    while (!qu.empty()) {
        Node *node = qu.front();
        // print node
        qu.pop();

        if (node.left != NULL) {
            qu.push(node.left);
        }

        if (node.right != NULL) {
            qu.push(node.right);
        }
    }
}


