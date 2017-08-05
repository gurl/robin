package web

// Group is used to categorize a group of actions
type Group struct {
	id     string `json:"id"`
	name   string `json:"name"`
	weight string `json:"weight"`
}

type Task struct {
	id     string `json:"id"`
	name   string `json:"name"`
	weight string `json:"weight"`
	group  string `json:"group"`
}

type Indicators struct {
	start string `json:"start"`
	end   string `json:"end"`
}

type Groups []Group
type Tasks []Task

type Registration struct {
	Tasks      Tasks      `json:"tasks"`
	Groups     Groups     `json:"groups"`
	Indicators Indicators `json:"indicators"`
}
