package shardkv

import "shardmaster"
//
// Sharded key/value server.
// Lots of replica groups, each running op-at-a-time paxos.
// Shardmaster decides which group serves each shard.
// Shardmaster may change shard assignment from time to time.
//
// You will have to modify these definitions.
//

const (
	OK            = "OK"
	ErrNoKey      = "ErrNoKey"
	ErrWrongGroup = "ErrWrongGroup"
)

type Err string

// Put or Append
type PutAppendArgs struct {
	Me    int64
	MsgId int64
	Key   string
	Value string
	Op    string // "Put" or "Append"
	Shard int
}

type PutAppendReply struct {
	WrongLeader bool
	Err         Err
}

type GetArgs struct {
	Key string
	Shard int
}

type GetReply struct {
	WrongLeader bool
	Err         Err
	Value       string
}

type ReqShared struct {
	Shard int
	Config shardmaster.Config
}

type RespShared struct {
	Config  shardmaster.Config
	Shard  int
	Data    map[string]string
	MsgIDs  map[int64] int64
}

type ReqDeleteShared struct {
	Shard int
	Config shardmaster.Config
}

type RespDeleteShared struct {
	Shard int
	Config shardmaster.Config
}