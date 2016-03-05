package payloads

type JsonSerializer interface {
	Serialize() ([]byte, error)
}
