/*
 * 统计目录下所有文件单词出现的次数从大到小输出
 */
package main

import (
	"fmt"
	"os"
	"bufio"
	"io/ioutil"
	"strings"
	"unicode"
    "unicode/utf8"
	"sort"
)
type Pair struct {
	k string
	v int
}
type PairList []Pair
func (p PairList)Len() int{
	return len(p)
}
func (p PairList)Less(i, j int) bool{
	return p[i].v > p[j].v
}
func (p PairList)Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
// 提取单词
func SplitOnNonLetters(s string) []string {
    notALetter := func(char rune) bool { return !unicode.IsLetter(char) }
    return strings.FieldsFunc(s, notALetter)
}
func counter(path string, done chan map[string]int) {
	count := make(map[string]int)
	f, e := os.Open(path)
	defer f.Close()
	if e != nil {
		fmt.Println(e)
		return
	}
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
            if len(word) > utf8.UTFMax ||
                utf8.RuneCountInString(word) > 1 {
                count[strings.ToLower(word)] += 1
            }
        }
		if err != nil {
			break
		}
	}
	done <- count
}
func summary(done chan map[string]int, cnt int) map[string]int{
	count := make(map[string]int)
	for i := 0; i < cnt; i++ {
		m := <- done
		for k, v := range(m) {
			count[k] +=v
		}
	}
	return count
}
/* 出现次数从多到少输出 */
func printCountOrder(data map[string]int) {
	var pairs PairList
	for k, v := range(data) {
		pairs = append(pairs, Pair{k, v})
	}
	sort.Sort(pairs)
	for _, p := range(pairs) {
		fmt.Printf("%s : %d\n", p.k, p.v)
	}

}
func dfs(dir string, done chan map[string]int) int{
	cnt := 0
	files, e := ioutil.ReadDir(dir)
	if e != nil {
		fmt.Println(e)
		return 0
	}
	for _, file := range(files) {
		if file.IsDir() {
			if file.Name() == ".git" {
				continue
			}
			cnt += dfs(dir + "/" + file.Name(), done)
			continue
		}
		cnt++
		go counter(dir + "/" + file.Name(), done)
	}
	return cnt
}
func main() {
	var dir string
	fmt.Printf("Please enter file directory: ")
	fmt.Scanln(&dir)
	done := make(chan map[string]int)
	cnt := dfs(dir, done)
	data := summary(done, cnt)
	printCountOrder(data)
}