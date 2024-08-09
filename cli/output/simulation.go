package output

type Simulation []Charge

type Charge struct {
	Tax float32 `json:"tax"`
}
