# go-ipfs-plugin-i2p-swarm

Plugin for connecting an IPFS swarm over i2p

**WARNING:** This is *only* for forwarding the local swarm port over i2p, not
for connecting to i2p-hosted IPFS instances. It **will not route**
**communication between IPFS nodes over i2p(1)**. This means that it **doesn't**
**make your IPFS instance anonymous**, it just makes it *accssible to clients*
*anonymously(2)*. As such, it's probably not that useful to people in general
yet. It's also emphatically *not* a product of the i2p Project and doesn't carry
a guarantee from them. File an issue here, I'm happy to help.

