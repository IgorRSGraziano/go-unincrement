package main

type Scope struct {
	Start uint16
	End   uint16
}

var AlphaNumericScope = [3]Scope{
	{Start: 48, End: 57},  // 0-9
	{Start: 65, End: 90},  // A-Z
	{Start: 97, End: 122}, // a-z
}

// Get one char, return next char based on AlphaNumericScope
func NextChar(char byte) byte {
	var nextChar byte
	//if char its equal end of scope, return start of next scope, and if scope is the last, return start of first scope
	for i, scope := range AlphaNumericScope {
		if char >= byte(scope.Start) && char <= byte(scope.End) {
			if char == byte(scope.End) {
				if i == len(AlphaNumericScope)-1 {
					nextChar = byte(AlphaNumericScope[0].Start)
				} else {
					nextChar = byte(AlphaNumericScope[i+1].Start)
				}
			} else {
				nextChar = char + 1
			}
			break
		}
	}
	return nextChar
}

func main() {
	println(string(NextChar('9'))) // A
	println(string(NextChar('Z'))) // a
	println(string(NextChar('z'))) // 0
	println(string(NextChar('0'))) // 1
}
