package providers

type RealtimeChatHubProvider interface {
	Get() interface{}
	Run()
	Stop()
}
