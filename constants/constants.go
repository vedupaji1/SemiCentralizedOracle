package temp11

const (
	PlatformName                       = "DataVerifier"
	PeersKey                           = "peers"
	PeersListKey                       = "peersList"
	MaxRetryForPeerListRequest         = 5
	TotalLeadingDuration               = 150 // 150 seconds
	EffectiveLeadingDuration           = 120 // 120 seconds
	LeadingPowerReleaseDuration        = 20  // 30 seconds
	BreakTime                          = 10  // 5 seconds
	WaitingTimeForPeersInEachRound     = 5   // 5 seconds
	RegistrationRequestMessage         = "Can You Add Me In Network"
	EthereumNodeSocketURL              = "wss://polygon-mumbai.g.alchemy.com/v2/barNHxwKcvdxJuDoKlbor5qx6mhT2C_O"
	EthereumNodeRPCURL                 = "https://polygon-mumbai.g.alchemy.com/v2/VUq8Nxd1_VYC2aA7R0ySc1kqL05cw0CF"
	RequestFulfillerContract           = "0xC072FA167EAc058c30701976689B8817B216270B"
	RequestReceiverContract            = "0x26226C0B001471910bDFcDC3B4A07e659dC7eD66"
	RequestReceivedEventTopicHash      = "0x77d170aa9240508fa5a2c20e6b50b4bdc9d4d03a132479fcfbe1593d96cfd340"
	MinETHBalanceForRequestFulfillment = 100000000000000000
	ChainID                            = 80001
	MaxGasForTransaction               = 5000000
)
