package helpers
import(
	"fmt"
)

func CleanInterfaceData(data map[string]interface{}) map[string]interface{} {
	if data == nil {
		return nil
	}
	cleaned := make(map[string]interface{})
		for k, v := range data {
			if v != nil && v != "" && v != "00000000-0000-0000-0000-000000000000" { 
				cleaned[k] = v
				fmt.Println("Keeping key:", k, "with value:", v)
			}
	}
    return cleaned
}

