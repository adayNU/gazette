package gazette

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"time"

	log "github.com/Sirupsen/logrus"
	etcd "github.com/coreos/etcd/client"
	"github.com/gorilla/mux"

	"github.com/pippio/gazette/consensus"
	"github.com/pippio/gazette/journal"
)

// API for creation of a new Journal. In particular, CreateAPI creates an Etcd
// item directory for the Journal under Gazette's consensus.Allocator root
// and responds to the client when the Journal is ready for transactions.
type CreateAPI struct {
	keysAPI          etcd.KeysAPI
	requiredReplicas int
}

func NewCreateAPI(keysAPI etcd.KeysAPI, requiredReplicas int) *CreateAPI {
	return &CreateAPI{keysAPI: keysAPI, requiredReplicas: requiredReplicas}
}

func (h *CreateAPI) Register(router *mux.Router) {
	router.NewRoute().Methods("POST").HandlerFunc(h.Create)
}

func (h *CreateAPI) Create(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[1:]
	path := path.Join(ServiceRoot, consensus.ItemsPrefix, url.QueryEscape(name))

	response, err := h.keysAPI.Set(context.Background(), path, "", &etcd.SetOptions{
		Dir:       true,
		PrevExist: etcd.PrevNoExist,
	})
	// Map a etcd NodeExist error into corresponding journal error.
	if etcdErr, _ := err.(etcd.Error); etcdErr.Code == etcd.ErrorCodeNodeExist {
		err = journal.ErrExists
	}

	if err != nil {
		http.Error(w, err.Error(), journal.StatusCodeForError(err))
		return
	}
	log.WithFields(log.Fields{"path": path, "name": name}).Info("created journal")

	// Briefly block until we see the required number of ready replicas under
	// the new item. If we returned immediately, the client will likely race
	// its next request against the consensus.Allocator (and often win!).
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var tree = response.Node
	var watcher = h.keysAPI.Watcher(path, &etcd.WatcherOptions{
		AfterIndex: response.Index,
		Recursive:  true,
	})

	for {
		var err error
		if response, err = watcher.Next(ctx); err != nil {
			http.Error(w, err.Error(), journal.StatusCodeForError(err))
			return
		} else if tree, err = consensus.PatchTree(tree, response); err != nil {
			http.Error(w, err.Error(), journal.StatusCodeForError(err))
			return
		}

		var readyCount int
		for _, node := range tree.Nodes {
			if node.Value == "ready" {
				readyCount += 1
			}
		}

		if readyCount > h.requiredReplicas {
			w.WriteHeader(http.StatusCreated)
			return
		}
	}
}
