// +build !libp2p

package i2pswarm

import (
	"github.com/eyedeekay/sam-forwarder"
	"github.com/rtradeltd/go-ipfs-plugin-i2p-gateway/config"
	"log"
)

func (i *I2PSwarmPlugin) transportSwarm() error {
	log.Println("Creating an i2p destination for the Swarm Server")
	host, err := i.i2pconfig.SwarmHost()
	if err != nil {
		return err
	}
	log.Println("host", host)
	port, err := i.i2pconfig.SwarmPort()
	if err != nil {
		return err
	}
	log.Println("port", port)
	GarlicForwarder, err := samforwarder.NewSAMForwarderFromOptions(
		samforwarder.SetSAMHost(i.i2pconfig.HostSAM()),
		samforwarder.SetSAMPort(i.i2pconfig.PortSAM()),
		samforwarder.SetHost(host),
		samforwarder.SetPort(port),
		samforwarder.SetType("server"),
		samforwarder.SetSaveFile(true),
		samforwarder.SetName("ipfs-gateway-swarm"),
		samforwarder.SetInLength(i.i2pconfig.InLength),
		samforwarder.SetOutLength(i.i2pconfig.OutLength),
		samforwarder.SetInVariance(i.i2pconfig.InVariance),
		samforwarder.SetOutVariance(i.i2pconfig.OutVariance),
		samforwarder.SetInQuantity(i.i2pconfig.InQuantity),
		samforwarder.SetOutQuantity(i.i2pconfig.OutQuantity),
		samforwarder.SetInBackups(i.i2pconfig.InBackupQuantity),
		samforwarder.SetOutBackups(i.i2pconfig.OutBackupQuantity),
		samforwarder.SetAllowZeroIn(i.i2pconfig.InAllowZeroHop),
		samforwarder.SetAllowZeroOut(i.i2pconfig.OutAllowZeroHop),
		samforwarder.SetCompress(i.i2pconfig.UseCompression),
		samforwarder.SetReduceIdle(i.i2pconfig.ReduceIdle),
		samforwarder.SetReduceIdleTimeMs(i.i2pconfig.ReduceIdleTime),
		samforwarder.SetReduceIdleQuantity(i.i2pconfig.ReduceIdleQuantity),
		samforwarder.SetAccessListType(i.i2pconfig.AccessListType),
		samforwarder.SetAccessList(i.i2pconfig.AccessList),
		samforwarder.SetEncrypt(i.i2pconfig.EncryptLeaseSet),
		samforwarder.SetLeaseSetKey(i.i2pconfig.EncryptedLeaseSetKey),
		samforwarder.SetLeaseSetPrivateKey(i.i2pconfig.EncryptedLeaseSetPrivateKey),
		samforwarder.SetLeaseSetPrivateSigningKey(i.i2pconfig.EncryptedLeaseSetPrivateSigningKey),
		samforwarder.SetMessageReliability(i.i2pconfig.MessageReliability),
	)
	log.Println("Generated Garlic Forwarder")
	if err != nil {
		return err
	}
	go GarlicForwarder.Serve()
	for {
		if len(GarlicForwarder.Base32()) > 51 {
			log.Println("base32: ", GarlicForwarder.Base32())
			break
		} else {
			log.Println("waiting for address")
		}
	}
	err = i2pgateconfig.ListenerBase32(GarlicForwarder.Base32(), i.i2pconfig)
	if err != nil {
		return err
	}
	err = i2pgateconfig.ListenerBase64(GarlicForwarder.Base64(), i.i2pconfig)
	if err != nil {
		return err
	}
	i.i2pconfig, err = i.i2pconfig.Save(i.configPath)
	if err != nil {
		return err
	}
	return nil
}
