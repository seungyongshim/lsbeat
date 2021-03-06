// Created by cgo -godefs - DO NOT EDIT
// cgo -godefs defs_solaris.go

package ipv4

const (
	sysIP_OPTIONS     = 0x1
	sysIP_HDRINCL     = 0x2
	sysIP_TOS         = 0x3
	sysIP_TTL         = 0x4
	sysIP_RECVOPTS    = 0x5
	sysIP_RECVRETOPTS = 0x6
	sysIP_RECVDSTADDR = 0x7
	sysIP_RETOPTS     = 0x8
	sysIP_RECVIF      = 0x9
	sysIP_RECVSLLA    = 0xa
	sysIP_RECVTTL     = 0xb

	sysIP_MULTICAST_IF           = 0x10
	sysIP_MULTICAST_TTL          = 0x11
	sysIP_MULTICAST_LOOP         = 0x12
	sysIP_ADD_MEMBERSHIP         = 0x13
	sysIP_DROP_MEMBERSHIP        = 0x14
	sysIP_BLOCK_SOURCE           = 0x15
	sysIP_UNBLOCK_SOURCE         = 0x16
	sysIP_ADD_SOURCE_MEMBERSHIP  = 0x17
	sysIP_DROP_SOURCE_MEMBERSHIP = 0x18
	sysIP_NEXTHOP                = 0x19

	sysIP_PKTINFO     = 0x1a
	sysIP_RECVPKTINFO = 0x1a
	sysIP_DONTFRAG    = 0x1b

	sysIP_BOUND_IF      = 0x41
	sysIP_UNSPEC_SRC    = 0x42
	sysIP_BROADCAST_TTL = 0x43
	sysIP_DHCPINIT_IF   = 0x45

	sysIP_REUSEADDR = 0x104
	sysIP_DONTROUTE = 0x105
	sysIP_BROADCAST = 0x106

	sysMCAST_JOIN_GROUP         = 0x29
	sysMCAST_LEAVE_GROUP        = 0x2a
	sysMCAST_BLOCK_SOURCE       = 0x2b
	sysMCAST_UNBLOCK_SOURCE     = 0x2c
	sysMCAST_JOIN_SOURCE_GROUP  = 0x2d
	sysMCAST_LEAVE_SOURCE_GROUP = 0x2e

	sizeofSockaddrStorage = 0x100
	sizeofSockaddrInet    = 0x10
	sizeofInetPktinfo     = 0xc

	sizeofIPMreq         = 0x8
	sizeofIPMreqSource   = 0xc
	sizeofGroupReq       = 0x104
	sizeofGroupSourceReq = 0x204
)

type sockaddrStorage struct {
	Family     uint16
	X_ss_pad1  [6]int8
	X_ss_align float64
	X_ss_pad2  [240]int8
}

type sockaddrInet struct {
	Family uint16
	Port   uint16
	Addr   [4]byte /* in_addr */
	Zero   [8]int8
}

type inetPktinfo struct {
	Ifindex  uint32
	Spec_dst [4]byte /* in_addr */
	Addr     [4]byte /* in_addr */
}

type ipMreq struct {
	Multiaddr [4]byte /* in_addr */
	Interface [4]byte /* in_addr */
}

type ipMreqSource struct {
	Multiaddr  [4]byte /* in_addr */
	Sourceaddr [4]byte /* in_addr */
	Interface  [4]byte /* in_addr */
}

type groupReq struct {
	Interface uint32
	Pad_cgo_0 [256]byte
}

type groupSourceReq struct {
	Interface uint32
	Pad_cgo_0 [256]byte
	Pad_cgo_1 [256]byte
}
