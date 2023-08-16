package natscc

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/choria-io/fisk"
	"github.com/nats-io/nats.go"
)

type minimalNatsContext struct {
	URL   string
	Creds string
	NSC   string
}

// FlagConnect connects to a nats server similar to the nats cli and its context
// This uses the flag package and parses the flags. Make sure to add your flags before
// calling this function!
func FlagConnect(opts ...nats.Option) (*nats.Conn, error) {
	context := FlagContext()
	flag.Parse()
	return context.Connect(opts...)
}

type NccContext struct {
	context     string
	skipContext bool
	servers     string
	credsFile   *string
}

// FlagContext creates a context for connect to a nats server similar to the nats cli and its context
// This uses the flag package and prepares some flags to parse into the private fields.
func FlagContext() *NccContext {
	nccc := &NccContext{}
	flag.StringVar(&nccc.context, "context", "", "nats context")
	flag.BoolVar(&nccc.skipContext, "skip-context", false, "skipt the nats context evaluation")
	flag.StringVar(&nccc.servers, "s", "", "server")
	creds := ""
	nccc.credsFile = &creds
	flag.StringVar(nccc.credsFile, "c", "", "creds file")
	return nccc
}

// FiskContext creates a context based on fisk cli args
func FiskContext(cli *fisk.Application) *NccContext {
	nccc := &NccContext{}
	cli.Flag("context", "nats context").Short('C').StringVar(&nccc.context)
	cli.Flag("skip-context", "skipt the nats context evaluation").BoolVar(&nccc.skipContext)
	cli.Flag("server", "nats servers").Short('s').StringVar(&nccc.servers)
	nccc.credsFile = cli.Flag("creds", "creds file").Short('c').ExistingFile()
	return nccc
}

func (nccc *NccContext) Connect(opts ...nats.Option) (*nats.Conn, error) {
	servers, creds, err := ContextEval(nccc.servers, *nccc.credsFile, nccc.context, nccc.skipContext)
	if err != nil {
		return nil, err
	}
	return CredsConnect(servers, creds, opts...)
}

// NatsCredsConnect connects to a nats server using a creds file
func CredsConnect(servers string, credsFile string,
	opts ...nats.Option) (*nats.Conn, error) {
	// Connect and get the JetStream context.
	if credsFile != "" {
		opts = append(opts, nats.UserCredentials(credsFile))
	}
	nc, err := nats.Connect(servers, opts...)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

// ConnectInfo creates a string about who we connected to for showing to the user
func ConnectInfo(nc *nats.Conn) string {
	return fmt.Sprintf("Connected to %q (%q) @ %q (TLS: %t)",
		nc.ConnectedServerName(), nc.ConnectedClusterName(), nc.ConnectedUrlRedacted(), nc.TLSRequired())
}

// ContextEval returns the servers and the (optional) creds file
// that the nats cli would use for the same context situation
func ContextEval(natsServer string, credsFile string, natsContext string,
	skipContext bool) (server string, creds string, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", "", err
	}
	if natsContext == "" && !skipContext {
		// this may be highly internal, but it works well enough for us right now
		contextName, _ := os.ReadFile(homeDir + "/.config/nats/context.txt")
		if err == nil {
			natsContext = string(contextName)
		}
	}
	var cd minimalNatsContext
	if natsContext != "" {
		// this may be highly internal, but it works well enough for us right now
		buf, err := os.ReadFile(homeDir + "/.config/nats/context/" + natsContext + ".json")
		if err != nil {
			return "", "", err
		}
		err = json.Unmarshal(buf, &cd)
		if err != nil {
			return "", "", err
		}
	}
	if natsServer == "" {
		natsServer = cd.URL
	}

	if credsFile == "" {
		credsFile = cd.Creds
	}

	if credsFile == "" {
		if cd.NSC != "" {
			credsFile = homeDir + "/.local/share/nats/nsc/keys/creds/" + cd.NSC[6:] + ".creds"
		}
	}
	return natsServer, credsFile, err
}
