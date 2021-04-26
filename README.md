# raptorq  

just for fun  


1. consumer: 解码
chunkManager()
setPiece(id,count,[]*piece)
AddReadyBlockChan(chan uint8)

2. 切片server:编码  
getRaqInfo()
registerPiece(uri,id,count) (chan []*piece,error)
missData(uri,id,count) ([]*piece,error)  

3. 微型服务器:转发
microRegisterPiece(uri,id,count) (chan []*piece,error)

4. tracker: 保留着切片与微型服务器。
SetRaqServer
SetMicroServer  

install:
sudo apt install libeigen3-dev
sudo ln -s /usr/include/eigen3/Eigen /usr/include/Eigen  

--什么是喷泉码（https://github.com/wuhang-du/go-raptorq）  
source:object --> source:block --> source:pieces  
中间传输的是pieces:
piece是由：  
- sbn(source-block-num)
- esi(encoding-piece-id)

需要同步的信息：
- Total source object size (in octets), 
- symbol size (chosen by the sender), 
- number of source blocks, 
- number of sub-blocks (an internal detail), 
- symbol alignment factor (another internal detail).  

New(input []byte, symbolSize uint16, minSubSymbolSize uint16,
		maxSubBlockSize uint32, alignment uint8) (Encoder, error)  

- input is the source object to encode
- symbolSize is the encoding symbol size, in octets
- minSubSymbolSize is the minimum encoding symbol size allowed, in octets.
- maxSubBlockSize is the maximum size block that is decodable in working
  memory, in octets.  Iff this is lower than the source object size, the
  source object will be split into more than one source blocks.  The
  maximum allowed value is 56403 * symbolSize.
- alignment is an internal alignment parameter, in bytes.  Typically this
}

// Decoder decodes encoding symbols and reconstructs one object from a series of
// symbols.
type Decoder interface {
	// Decoder needs to provide object information.
	ObjectInfo

	// Decode decodes a received encoding symbol.
	//
	// The result of decoding may not be available immediately after Decode
	// returns; IsSourceBlockReady or IsSourceObjectReady may not
	// immediately return true even if the symbol made the source block or
	// the source object available.
	// Use AddReadyBlockChan if immediate notification is needed.
	Decode(sbn uint8, esi uint32, symbol []byte)

	// IsSourceBlockReady returns whether the given source block has been fully
	// decoded and ready to be retrieved, or false if sbn is out of range.
	IsSourceBlockReady(sbn uint8) bool

	// IsSourceObjectReady returns whether the entire source object has been
	// fully decoded and ready to be retrieved.
	IsSourceObjectReady() bool

	// SourceBlock copies the given source block into the given buffer.  buf
	// should contain enough space to store the given source block (use
	// SourceBlockSize(sbn) to get the required size).
	SourceBlock(sbn uint8, buf []byte) (n int, err error)

	// SourceObject copies the source object into the given buffer.  buf should
	// contain enough space to store the entire source object (use
	// TransferLength() to get the required size).
	SourceObject(buf []byte) (n int, err error)

	// Free, on supported implementations, will free memory used for generating
	// encoding symbols for the given source block.  Once a source block has
	// been freed, calling Encode with its SBN may return an error.
	FreeSourceBlock(sbn uint8)

	// AddReadyBlockChan adds a channel through which the decoder shall avail
	// source blocks ready for retrieval as soon as they become available.
	//
	// If any block is currently available,
	// their source block number is immediately sent into the given channel.
	//
	// The added channel is closed when no more blocks are available,
	// e.g. when the decoder is closed or destroyed.
	//
	// If more than one channel is added,
	// newly available source block numbers are sent to all of them,
	// in no specific order.
	AddReadyBlockChan(chan<- uint8) (err error)

	// RemoveReadyBlockChan removes a channel previously added via
	// AddReadyBlockChan.  It does not close the removed channel.
	RemoveReadyBlockChan(chan<- uint8) (err error)

	// Close closes the Decoder.  After a Decoder is closed, all methods but
	// Close() will panic if called.
	Close() error
}