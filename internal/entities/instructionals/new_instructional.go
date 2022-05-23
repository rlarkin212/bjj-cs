package instructionals

type NewInstructional struct {
	Title     string `json:"title"`
	Presenter string `json:"presenter"`
	Cover     string `json:"cover"`
	Part      int    `json:"part"`
	URL       string ` json:"url"`
}
