package queryservice

type Session interface {
    Get(sessionId string) (string, error)
}
