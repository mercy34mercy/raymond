package lexer

// 置き換えた処理
/*
rOpenCommentDash = regexp.MustCompile(`^\{\{~?!--\s*`)
l.findRegexp(rOpenCommentDash);
*/
func checkROpenCommentDash(str string) (int, int) {
	if len(str) >= 5 && str[0:2] == "{{" {
		length := 2
		// Optionally skip the '~'
		if str[length] == '~' {
			length++
		}
		// Check for '!--'
		if len(str[length:]) >= 3 && str[length:length+3] == "!--" {
			length += 3
			// Skip any following whitespace characters
			for length < len(str) && (str[length] == ' ' || str[length] == '\t') {
				length++
			}
			return 0, length
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenComment = regexp.MustCompile(`^\{\{~?!\s*`)
findPattern(l.input[l.pos:], checkROpenComment)
*/
func checkROpenComment(str string) (int, int) {
	if len(str) >= 3 && str[0:2] == "{{" {
		length := 2
		// Optionally skip the '~'
		if str[length] == '~' {
			length++
		}
		// Check for '!'
		if length < len(str) && str[length] == '!' {
			length++
			// Skip any following whitespace characters
			for length < len(str) && (str[length] == ' ' || str[length] == '\t') {
				length++
			}
			return 0, length
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenInverse = regexp.MustCompile(`^\{\{~?\^`)
l.findRegexp(rOpenInverse)
*/
func checkROpenInverse(str string) (int, int) {
	if len(str) >= 3 && str[0:2] == "{{" {
		length := 2
		// Optionally skip the '~'
		if str[length] == '~' {
			length++
		}
		// Check for the following '^'
		if length < len(str) && str[length] == '^' {
			return 0, length + 1
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenPartial = regexp.MustCompile(`^\{\{~?>`)
l.findRegexp(rOpenPartial)
*/
func checkROpenPartial(str string) (int, int) {
	if len(str) >= 3 && str[0:2] == "{{" {
		length := 2
		// Optionally skip the '~'
		if str[length] == '~' {
			length++
		}
		// Check for the following '>'
		if length < len(str) && str[length] == '>' {
			return 0, length + 1
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenEndBlock = regexp.MustCompile(`^\{\{~?/`)
l.findRegexp(rOpenEndBlock)
*/
func checkROpenEndBlock(str string) (int, int) {
	if len(str) >= 3 && str[0:2] == "{{" {
		length := 2
		// Optionally skip the '~'
		if str[length] == '~' {
			length++
		}
		// Check for the following '/'
		if length < len(str) && str[length] == '/' {
			return 0, length + 1
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenBlock = regexp.MustCompile(`^\{\{~?#`)
l.findRegexp(rOpenBlock)
*/
func checkROpenBlock(str string) (int, int) {
	if len(str) >= 3 && str[0:2] == "{{" {
		length := 2
		// Optionally skip the '~'
		if str[length] == '~' {
			length++
		}
		// Check for the following '#'
		if length < len(str) && str[length] == '#' {
			return 0, length + 1
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenUnescaped = regexp.MustCompile(`^\{\{~?\{`)
l.findRegexp(rOpenUnescaped)
*/
func checkROpenUnescaped(str string) (int, int) {
	if len(str) >= 3 && str[0:2] == "{{" {
		length := 2
		// Optionally skip the first '~'
		if str[length] == '~' {
			length++
		}
		// Check for the following '{'
		if length < len(str) && str[length] == '{' {
			return 0, length + 1
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenEndRaw = regexp.MustCompile(`^\{\{\{\{/`)
l.findRegexp(rOpenEndRaw);
*/
func checkROpenEndRaw(str string) (int, int) {
	if len(str) >= 5 && str[0:5] == "{{{{/" {
		return 0, 5
	}
	return -1, 0
}

// 置き換えた処理
// この処理がかなり重く、置き換えることでかなり高速化される
/*
rOpenRaw = regexp.MustCompile(`^\{\{\{\{`)
l.findRegexp(rOpenRaw);
*/
func checkROpenRaw(str string) (int, int) {
	if len(str) >= 4 && str[0:4] == "{{{{" {
		return 0, 4
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpen = regexp.MustCompile(`^\{\{~?&?`)
l.findRegexp(rOpen)
*/
func checkROpen(str string) (int, int) {
	if len(str) >= 2 && str[0] == '{' && str[1] == '{' {
		length := 2
		// Count optional '~' and '&' characters
		for length < len(str) && (str[length] == '~' || str[length] == '&') {
			length++
		}
		return 0, length
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenInverseChain = regexp.MustCompile(`^\{\{~?\s*else`)
l.findRegexp(rOpenInverseChain)
*/
func checkROpenInverseChain(str string) (int, int) {
	if len(str) >= 4 && str[0] == '{' && str[1] == '{' {
		length := 2
		// Optionally skip the '~' character
		if length < len(str) && str[length] == '~' {
			length++
		}
		// Skip any whitespace characters
		for length < len(str) && (str[length] == ' ' || str[length] == '\t' || str[length] == '\n') {
			length++
		}
		// Check if "else" follows
		if length+4 <= len(str) && str[length:length+4] == "else" {
			return 0, length + 4
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rClose = regexp.MustCompile(`^~?\}\}`)
l.findRegexp(rClose)
*/
func checkRClose(str string) (int, int) {
	if len(str) >= 2 {
		length := 0
		// Optionally skip the '~'
		if str[length] == '~' {
			length++
		}
		// Check for the following '}}'
		if length < len(str)-1 && str[length] == '}' && str[length+1] == '}' {
			return 0, length + 2
		}
	}
	return -1, 0
}

// 置き換えた処理
/*
rCloseRaw = regexp.MustCompile(`^\}\}\}\}`)
l.findRegexp(rCloseRaw)
*/
func checkRCloseRaw(str string) (int, int) {
	if len(str) >= 4 && str[0:4] == "}}}}" {
		return 0, 4
	}
	return -1, 0
}

// 置き換えた処理
/*
rOpenBlockParams = regexp.MustCompile(`^as\s+\|`)
l.findRegexp(rOpenBlockParams)
*/
func checkROpenBlockParams(str string) (int, int) {
	if len(str) >= 4 && str[0:2] == "as" {
		length := 2
		// Skip any whitespace characters
		for length < len(str) && (str[length] == ' ' || str[length] == '\t') {
			length++
		}
		// Check for the following '|'
		if length < len(str) && str[length] == '|' {
			return 0, length + 1
		}
	}
	return -1, 0
}
