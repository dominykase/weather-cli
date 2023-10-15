package modes

type Mode interface {
    Handle() error
}
