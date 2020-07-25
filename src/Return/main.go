package main

import "fmt"

// 반환값에 이름을 붙여서 int형 count와 string슬라이스형 list를 반환
func getStudyMemberList(names ...string) (count int, list []string) {
	for _, n := range names {
		// list , count 는 이미 Named Return Parameter가 되었으므로 변수가 초기화 되어있는 상태.
		list = append(list, n)
		count++
	}
	//생략하면 큰일 난다.
	return
}

func main() {
	studyMemberList := make([]string, 0)
	studyMemberList = append(studyMemberList, "aaa")
	studyMemberList = append(studyMemberList, "bbb")
	studyMemberList = append(studyMemberList, "ccc")

	count, names := getStudyMemberList(studyMemberList...)
	fmt.Println(fmt.Sprintf("Count : %d , NameList: %v", count, names))

}
