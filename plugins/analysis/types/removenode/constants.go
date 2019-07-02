package removenode

const (
	MARSHALED_PACKET_HEADER = 0x02

	MARSHALED_PACKET_HEADER_START = 0
	MARSHALED_PACKET_HEADER_SIZE  = 1
	MARSHALED_PACKET_HEADER_END   = MARSHALED_PACKET_HEADER_START + MARSHALED_PACKET_HEADER_SIZE

	MARSHALED_ID_START = MARSHALED_PACKET_HEADER_END
	MARSHALED_ID_SIZE  = 20
	MARSHALED_ID_END   = MARSHALED_ID_START + MARSHALED_ID_SIZE

	MARSHALED_TOTAL_SIZE = MARSHALED_ID_END
)
