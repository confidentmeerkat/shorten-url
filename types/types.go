package types

type URL struct {
	Origin string `json:"origin"`
	Short  string `json:"short"`
}

type Error struct {
	Err string `json:"error"`
}
