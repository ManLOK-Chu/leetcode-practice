package main

import (
	"fmt"
	"strings"
)

func wordBreak(s string, wordDict []string) []string {
	var n = len(s)
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	var results []string
	var paths []string
	var backtracking func(int)
	backtracking = func(index int) {
		if index == n {
			results = append(results, strings.Join(paths, " "))
			return
		}
		for i := index; i < n; i++ {
			str := s[index : i+1]
			if wordDictSet[str] {
				paths = append(paths, str)
				backtracking(i + 1)
				paths = paths[:len(paths)-1]
			}
		}
	}
	backtracking(0)
	return results
}

/*
public List<String> wordBreak(String s, List<String> wordDict) {
        Set<String> set = new HashSet<>(wordDict);
        List<String> ans = new ArrayList<>();
        dfs(s, ans, 0, new LinkedList<>(), set);
        return ans;
    }

    public void dfs(String s, List<String> ans, int idx, Deque<String> path, Set<String> set) {
        if (idx == s.length()) {
            ans.add(String.join(" ", path));
            return;
        }
        for (int i = idx; i < s.length(); i++) {
            String str = s.substring(idx, i + 1);
            if (set.contains(str)) {
                path.add(str);
                dfs(s, ans, i + 1, path, set);
                path.removeLast();
            }
        }
    }
*/
func main() {
	fmt.Println(wordBreak("catsanddog", []string{"cat", "cats", "and", "sand", "dog"}))
}
