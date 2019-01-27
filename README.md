# go-ipfs-plugin-i2p-swarm

Plugin for connecting an IPFS swarm over i2p

**WARNING:** This is *only* for forwarding the local swarm port over i2p, not
for connecting to i2p-hosted IPFS instances. It **will not route**
**communication between IPFS nodes over i2p(1)**. This means that it **doesn't**
**make your IPFS instance anonymous**, it just makes it *accssible to clients*
*anonymously(2)*. As such, it's probably not that useful to people in general
yet. It's also emphatically *not* a product of the i2p Project and doesn't carry
a guarantee from them. File an issue here, I'm happy to help.

see also: [go-ipfs-plugin-i2p-gateway](https://github.com/rtradeltd/go-ipfs-plugin-i2p-gateway)

### Notes:

(1) I'm going to do that too, but that's the hard part(the plan is to adapt
BiglyBT-style bridging, with a pure-clearnet peers, clearnet-to-i2p peers, and
pure-i2p peers. That may be less straightforward than the simple description
made it sound).

(2) Of course, that leaves the matter of i2p-compatible IPFS applications but
those are almost as simple as the gateway.
