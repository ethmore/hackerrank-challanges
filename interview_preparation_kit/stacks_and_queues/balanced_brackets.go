package stacksandqueues

func isBalanced(s string) string {
	brackets := []rune(s)
	var openedBrackets []rune

	for _, bracket := range brackets {

		if bracket == '(' || bracket == '[' || bracket == '{' {
			openedBrackets = append(openedBrackets, bracket)
		} else if len(openedBrackets) > 0 {

			if bracket == ')' && openedBrackets[len(openedBrackets)-1] == '(' {
				openedBrackets = openedBrackets[:len(openedBrackets)-1]
			} else if bracket == ']' && openedBrackets[len(openedBrackets)-1] == '[' {
				openedBrackets = openedBrackets[:len(openedBrackets)-1]
			} else if bracket == '}' && openedBrackets[len(openedBrackets)-1] == '{' {
				openedBrackets = openedBrackets[:len(openedBrackets)-1]
			} else {
				return "NO"
			}
		} else {
			return "NO"
		}
	}
	if len(openedBrackets) == 0 {
		return "YES"
	}
	return "NO"
}
