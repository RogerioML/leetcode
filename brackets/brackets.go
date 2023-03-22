package brackets

func IsValid(s string) bool {
	if len(s) <= 1 {
		return false
	}	
	chars := []byte(s)
	var pilha []byte
	for i := 0; i < len(chars); i++ {
		if chars[i] == '(' || chars[i] == '{' || chars[i] == '[' {
			pilha = append(pilha, chars[i])
		} else {
			if len(pilha) == 0 {
				return false
			}
			switch chars[i] {
			case ')':
				if pilha[len(pilha)-1] != '(' {
					return false
				}
			case '}':
				if pilha[len(pilha)-1] != '{' {
					return false
				}
			case ']':
				if pilha[len(pilha)-1] != '[' {
					return false
				}
			default:
				return false
			}
			pilha = pilha[:len(pilha)-1]
		}
	}
	return len(pilha) == 0
}
