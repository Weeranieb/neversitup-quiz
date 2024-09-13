package utils

func CountSmileyFace(faces []string) int {
	count := 0
	for _, face := range faces {
		if isSmileyFace(face) {
			count++
		}
	}
	return count
}

func isSmileyFace(face string) bool {

	if len(face) != 2 && len(face) != 3 {
		return false
	}

	eyes := face[0]
	mouth := face[len(face)-1]
	nose := face[1 : len(face)-1]

	// check eyes
	if eyes != ':' && eyes != ';' {
		return false
	}

	// check mouth
	if mouth != ')' && mouth != 'D' {
		return false
	}

	if len(nose) > 0 {
		if nose != "-" && nose != "~" {
			return false
		}
	}
	return true
}
