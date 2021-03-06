package api

import (
	"common"
	"netConfig"
	"tcp"
)

var (
	g_cache_game_conn = make(map[int]*tcp.TCPConn)
)

func CallRpcGame(svrID int, rpc string, sendFun, recvFun func(*common.NetPack)) {
	GetGameConn(svrID).CallRpc(rpc, sendFun, recvFun)
}
func GetGameConn(svrId int) *tcp.TCPConn {
	conn, _ := g_cache_game_conn[svrId]
	if conn == nil || conn.IsClose() {
		conn = netConfig.GetTcpConn("game", svrId)
		g_cache_game_conn[svrId] = conn
	}
	return conn
}
