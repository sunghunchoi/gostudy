package main

import "fmt"

type PokemonPlay struct {
	PlayerName  string
	PokemonName string
}

func printGuideMessage() {
	fmt.Println("포켓몬스터를 생성하는 프로그램입니다. \n 닉네임과 포켓몬 이름을 차례대로 띄어 써주세요.")
	fmt.Print("닉네임과 포켓몬 이름을 입력해주세요 :")
}

func userInput() (string, string) {
	var nickName, pokemonName string
	fmt.Scanln(&nickName, &pokemonName)
	return nickName, pokemonName
}

func makePlayerInfo(nickName, pokemonName string) PokemonPlay {
	return PokemonPlay{
		PlayerName:  nickName,
		PokemonName: pokemonName,
	}
}

func printPlayerInfo(play PokemonPlay) {
	fmt.Printf("입력하신 닉네임은 %s, 포켓몬 이름은  %s 입니다. \n", play.PlayerName, play.PokemonName)
}

func main() {
	printGuideMessage()
	n, p := userInput()
	r := makePlayerInfo(n, p)
	printPlayerInfo(r)
}
