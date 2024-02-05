package main
import "fmt"

func uniqueNames(a, b []string) []string {
	result := []string{}
	s := append(a,b...)
	keys := make(map[string]bool)
	for _, item := range(s) {
		if _, value := keys[item]; !value {
			keys[item] = true
			result = append(result,item)
		}
	}
	return result
}
    
func main() {
	// should print Ava, Emma, Olivia, Sophia
	fmt.Println(uniqueNames(
	[]string{"Ava", "Emma", "Olivia"}, 
	[]string{"Olivia", "Sophia", "Emma"}))
}