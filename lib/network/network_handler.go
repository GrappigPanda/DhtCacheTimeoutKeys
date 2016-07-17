package networkHandler

import (
	"github.com/GrappigPanda/Olivia/cache"
	"github.com/GrappigPanda/Olivia/lib/network/incoming"
	"github.com/GrappigPanda/Olivia/lib/network/message_handler"
)

// StartIncomingNetwork handles spinning up an incoming network router and
// doing any error checking (in the future) as well as ensuring that it
// continues running.
func StartIncomingNetwork(mh *message_handler.MessageHandler, cache *cache.Cache) {
	go incomingNetwork.StartNetworkRouter(mh, cache)
}
