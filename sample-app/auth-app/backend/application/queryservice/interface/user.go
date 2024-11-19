package queryservice

type Auth interface {
    Get(sessionId string) (string, error)
}
