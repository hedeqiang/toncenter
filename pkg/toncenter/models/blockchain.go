package models

// Block represents a blockchain block
type Block struct {
	AfterMerge             bool      `json:"after_merge"`
	AfterSplit             bool      `json:"after_split"`
	BeforeSplit            bool      `json:"before_split"`
	CreatedBy              string    `json:"created_by"`
	EndLt                  string    `json:"end_lt"`
	FileHash               string    `json:"file_hash"`
	Flags                  int       `json:"flags"`
	GenCatchainSeqno       int       `json:"gen_catchain_seqno"`
	GenUtime               string    `json:"gen_utime"`
	GlobalID               int       `json:"global_id"`
	KeyBlock               bool      `json:"key_block"`
	MasterRefSeqno         int       `json:"master_ref_seqno"`
	MasterchainBlockRef    BlockID   `json:"masterchain_block_ref"`
	MinRefMcSeqno          int       `json:"min_ref_mc_seqno"`
	PrevBlocks             []BlockID `json:"prev_blocks"`
	PrevKeyBlockSeqno      int       `json:"prev_key_block_seqno"`
	RandSeed               string    `json:"rand_seed"`
	RootHash               string    `json:"root_hash"`
	Seqno                  int       `json:"seqno"`
	Shard                  string    `json:"shard"`
	StartLt                string    `json:"start_lt"`
	TxCount                int       `json:"tx_count"`
	ValidatorListHashShort int       `json:"validator_list_hash_short"`
	Version                int       `json:"version"`
	VertSeqno              int       `json:"vert_seqno"`
	VertSeqnoIncr          bool      `json:"vert_seqno_incr"`
	WantMerge              bool      `json:"want_merge"`
	WantSplit              bool      `json:"want_split"`
	Workchain              int       `json:"workchain"`
}

// BlockID represents a block identifier
type BlockID struct {
	Seqno     int    `json:"seqno"`
	Shard     string `json:"shard"`
	Workchain int    `json:"workchain"`
}

// BlocksResponse represents response for blocks request
type BlocksResponse struct {
	Blocks []Block `json:"blocks"`
}
