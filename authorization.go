// package main

// func (h *Handler) Run(addr string) {
// 	r := gin.Default()
// 	r.GET("/projects", h.List)
// 	r.Run(addr)
// }

// func (h *Handler) List(c *gin.Context) {
// 	c.JSON(200, []gin.H{
// 		{"id": 1, "title": "Дело"},
// 		{"id": 2, "title": "Витрина"}})
// }