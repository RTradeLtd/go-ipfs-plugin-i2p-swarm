package i2pswarm

import (
	"log"
	"os"
	"strings"

	"github.com/rtradeltd/go-ipfs-plugin-i2p-gateway/config"
	//TODO: Get a better understanding of gx.
	config "github.com/ipsn/go-ipfs/gxlibs/github.com/ipfs/go-ipfs-config"
	plugin "github.com/ipsn/go-ipfs/plugin"
	fsrepo "github.com/ipsn/go-ipfs/repo/fsrepo"
	peer "github.com/libp2p/go-libp2p-peer"
	"github.com/opentracing/opentracing-go"
)

type I2PSwarmPlugin struct {
	configPath    string
	config        *config.Config
	i2pconfigPath string
	i2pconfig     *i2pgateconfig.Config
	id            peer.ID

	forwardSwarm string
}

// I2PType will be used to identify this as the i2p gateway plugin to things
// that use it.
var I2PType = "i2pgate"

var _ plugin.Plugin = (*I2PSwarmPlugin)(nil)

// Name returns the plugin's name, satisfying the plugin.Plugin interface.
func (*I2PSwarmPlugin) Name() string {
	return "fwd-i2pgate"
}

// Version returns the plugin's version, satisfying the plugin.Plugin interface.
func (*I2PSwarmPlugin) Version() string {
	return "0.0.0"
}

// Init initializes plugin, satisfying the plugin.Plugin interface. Put any
// initialization logic here.
func (i *I2PSwarmPlugin) Init() error {
	var err error
	i.configPath, err = fsrepo.BestKnownPath()
	if err != nil {
		return err
	}
	err = os.Setenv("KEYS_PATH", i.configPath)
	if err != nil {
		return err
	}
	i.config, err = fsrepo.ConfigAt(i.configPath)
	if err != nil {
		return err
	}
	i.forwardSwarm = i.swarmString()

	i.i2pconfig, err = i2pgateconfig.ConfigAt(i.configPath)
	if err != nil {
		return err
	}

	err = i.configGateway()
	if err != nil {
		return err
	}

	i.i2pconfig, err = i.i2pconfig.Save(i.configPath)
	if err != nil {
		return err
	}
	go i.transportSwarm()
	return nil
}

func Setup() (*I2PSwarmPlugin, error) {
	var err error
	var i I2PSwarmPlugin
	i.configPath, err = fsrepo.BestKnownPath()
	if err != nil {
		return nil, err
	}
	err = os.Setenv("KEYS_PATH", i.configPath)
	if err != nil {
		return nil, err
	}
	i.config, err = fsrepo.ConfigAt(i.configPath)
	if err != nil {
		return nil, err
	}
	i.forwardSwarm = i.swarmString()
	log.Println("Prepared to forward:", i.forwardSwarm)
	i.i2pconfig, err = i2pgateconfig.ConfigAt(i.configPath)
	return &i, nil
}

func (i I2PSwarmPlugin) configGateway() error {
	err := i2pgateconfig.AddressSwarm(i.forwardSwarm, i.i2pconfig)
	if err != nil {
		return err
	}
	i.id, err = peer.IDFromString(i.idString())
	if err != nil {
		return err
	}
	log.Println(i.idString())
	i.i2pconfig, err = i.i2pconfig.Save(i.configPath)
	if err != nil {
		return err
	}
	return nil
}

func (i *I2PSwarmPlugin) swarmString() string {
	rpcaddress := ""
	for _, v := range i.config.Addresses.Swarm {
		rpcaddress += v
	}
	return unquote(string(rpcaddress))
}

func (i *I2PSwarmPlugin) idString() string {
	idbytes := i.config.Identity.PeerID
	return unquote(string(idbytes))
}

// I2PTypeName returns I2PType
func (*I2PSwarmPlugin) I2PTypeName() string {
	return I2PType
}

func unquote(s string) string {
	return strings.Replace(s, "\"", "", -1)
}

func (*I2PSwarmPlugin) InitTracer() (opentracing.Tracer, error) {
	return nil, nil
}
