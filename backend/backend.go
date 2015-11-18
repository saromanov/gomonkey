package backend

// Backend provides storing data to physics backend (file, db, etc...)
type Backend interface {

	// Write provides writing arguments to backend
	Write(string title, data)

	// Get returns arguments for testing func by title
	Get(string title)

	// String returns type of backend
	String()
}