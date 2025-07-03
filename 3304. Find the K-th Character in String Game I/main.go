package main

func kthCharacter(k int) byte {
	word := []byte{'a'}
	for len(word) < k {
		generatedWord := make([]byte, len(word))
		var generatedChar byte

		for i, char := range word {
			switch char {
			case 'z':
				generatedChar = 'a';
				break;
			default:
				generatedChar = char + 1;
			}

			generatedWord[i] = generatedChar
		}

		word = append(word, generatedWord...)
	}

	return word[k-1]
}

func main() {
}
