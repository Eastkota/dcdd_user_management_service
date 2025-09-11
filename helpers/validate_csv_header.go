package helpers

func validateCSVHeader(header []string) bool {
	expected := []string{"Name", "MobileNo", "Email", "Gender", "Password"}
	if len(header) != len(expected) {
		return false
	}
	for i, h := range header {
		if h != expected[i] {
			return false
		}
	}
	return true
}