package types

type Url struct {
	Origin string `json:"origin"`
	Short  string `json:"short"`
}

type Error struct {
	Err string `json:"error"`
}
