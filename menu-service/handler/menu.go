package handler

import (
	"github.com/"
	"net/http"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)

	// response, err := json.Marshal(map[string]interface{}{
	// 	"success": true,
	// })
	// if err != nil {
	// 	fmt.Print("failed to generate response")
	// 	return
	// }
	// w.Write(response)

}
