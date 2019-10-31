#include <stdio.h>
#include <string.h>
#include <stdlib.h>

char buffer[100000];
char *word[1000];
int nword = 0;
int k = 2;
int wordremaining = 10000;

int wordCmp(char *p, char *q)
{
	int n = k;
	while (*p++ == *q++)
	{
		if (*p == 0 && n > 0)
			n--;
		if (n == 0)
			return 0;
	}
	return *p - *q;
}

char *skip(char *p, int n)
{
	while (n > 0)
	{
		if (*p++ == 0)
			n--;
	}
	return p;
}

int main()
{
	*word = buffer;
	while (scanf("%s", *word) != EOF)
	{
		word[nword + 1] = word[nword] + strlen(word[nword]) + 1;
		nword++;
	}
	qsort(word, nword, sizeof(word[0]), wordCmp);
	for (int i = 0; i < k; i++)
	{
		printf("%s\n", word[i]);
	}
	int l = -1, u = nword - 1, m;
	char *phrase;
	for (int i = 0; i < wordremaining; i++)
	{
		while (l != u - 1)
		{
			m = (l + u) / 2;
			if (word[m] < phrase)
				l = m;
			else
				u = m;
		}
		while (wordCmp())
		{
			/* code */
		}

		printf("%s\n", phrase);
	}
}