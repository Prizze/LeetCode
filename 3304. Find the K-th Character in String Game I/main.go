package main

func kthCharacter(k int) byte {
	word := []byte{'a'} // Изначальное слово
	for len(word) < k {	
		generatedWord := make([]byte, len(word)) // Резервируем место той же длины
		var generatedChar byte				     // Для посимвольной генерации

		for i, char := range word {
			switch char {
			// Переход z -> a
			case 'z':
				generatedChar = 'a';
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
