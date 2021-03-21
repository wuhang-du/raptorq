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

