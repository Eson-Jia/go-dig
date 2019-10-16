#include <iostream>
#include <set>
using namespace std;

class IntSetSTL
{
private:
    set<int> S;

public:
    IntSetSTL(int maxelements, int maxval) {}
    int size() { return S.size(); }
    void insert(int t) { S.insert(t); }
    void report(int *v)
    {
        int j = 0;
        for (auto i = S.begin(); i != S.end(); i++)
        {
            v[j++] = *i;
        }
    }
};

class IntSetBitVec
{
private:
    enum
    {
        BITSPERWORD = 32,
        SHIFT = 5,
        MASK = 0x1F
    };
    int n, hi, *x;
    void set(int i) { x[i >> SHIFT] |= (1 << (i & MASK)); }
    void clr(int i) { x[i >> SHIFT] &= ~(1 << (i & MASK)); }
    int test(int i) { return x[i >> SHIFT] & (1 << (i & MASK)); }

public:
    IntSetBitVec(int maxelements, int maxval)
    {
        hi = maxval;
        x = new int[1 + hi / BITSPERWORD];
        for (int i = 0; i < hi; i++)
            clr(i);
        n = 0;
    }
    int size() { return n; }
    void insert(int t)
    {
        if (test(t))
            return;
        set(t);
        n++;
    }
    void report(int *v)
    {
        int j = 0;
        for (int i = 0; i < hi; i++)
            if (test(i))
            {
                v[j++] = i;
                cout << "=>" << i << endl;
            }
    }
};

class IntSetBins
{
private:
    int n, bins, maxval;
    struct node
    {
        int val;
        node *next;
        node(int v, node *p)
        {
            val = v;
            next = p;
        }
    };
    node **bin, *sentinel;
    node *rinsert(node *p, int t)
    {
        if (p->val < t)
        {
            p->next = rinsert(p->next, t);
        }
        else if (p->val > t)
        {
            p = new node(t, p);
            n++;
        }
        return p;
    }

public:
    IntSetBins(int maxelements, int pmaxval)
    {
        bins = maxelements;
        maxval = pmaxval;
        bin = new node *[bins];
        sentinel = new node(maxval, 0);
        for (int i = 0; i < bins; i++)
            bin[i] = sentinel;
        n = 0;
    }
    int size() { return n; }
    void insert(int t)
    {
        int i = t / (1 + maxval / bins); // CHECK !
        bin[i] = rinsert(bin[i], t);
    }
    void report(int *v)
    {
        int j = 0;
        for (int i = 0; i < bins; i++)
            for (node *p = bin[i]; p != sentinel; p = p->next)
                v[j++] = p->val;
    }
};
int main()
{
    auto s = IntSetBitVec(100, 100);
    s.insert(12);
    s.insert(1);
    s.insert(34);
    auto buff = new int[s.size()];
    s.report(buff);
    return 0;
}