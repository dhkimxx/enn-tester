package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8000")

	// cmd := exec.Command("adb", "devices")

	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	log.Fatalf("Error executing adb command: %s", err)
	// }

	// // fmt.Printf("Connected devices:\n%s\n", output)

	// // 출력 파싱 및 장치 목록 출력
	// lines := strings.Split(string(output), "\n")
	// for _, line := range lines {
	// 	if strings.Contains(line, "\tdevice") {
	// 		serial := fmt.Sprintf(strings.Fields(line)[0])
	// 		fmt.Printf("Serial: %s\n", serial)
	// 		cmd := exec.Command("adb", "-s", serial, "shell", "ls")
	// 		output, err := cmd.CombinedOutput()
	// 		if err != nil {
	// 			log.Fatalf("Error executing adb command: %s", err)
	// 		}
	// 		fmt.Printf("%s\n", output)
	// 	}
	// }
}
