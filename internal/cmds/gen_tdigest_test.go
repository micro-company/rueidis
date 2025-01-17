// Code generated DO NOT EDIT

package cmds

import "testing"

func tdigest0(s Builder) {
	s.TdigestAdd().Key("1").Value(1).Value(1).Build()
	s.TdigestByrank().Key("1").Rank(1).Rank(1).Build()
	s.TdigestByrevrank().Key("1").ReverseRank(1).ReverseRank(1).Build()
	s.TdigestCdf().Key("1").Value(1).Value(1).Build()
	s.TdigestCreate().Key("1").Compression(1).Build()
	s.TdigestCreate().Key("1").Build()
	s.TdigestInfo().Key("1").Build()
	s.TdigestMax().Key("1").Build()
	s.TdigestMerge().DestinationKey("1").Numkeys(1).SourceKey("1").SourceKey("1").Compression(1).Override().Build()
	s.TdigestMerge().DestinationKey("1").Numkeys(1).SourceKey("1").SourceKey("1").Compression(1).Build()
	s.TdigestMerge().DestinationKey("1").Numkeys(1).SourceKey("1").SourceKey("1").Override().Build()
	s.TdigestMerge().DestinationKey("1").Numkeys(1).SourceKey("1").SourceKey("1").Build()
	s.TdigestMin().Key("1").Build()
	s.TdigestQuantile().Key("1").Quantile(1).Quantile(1).Build()
	s.TdigestRank().Key("1").Value(1).Value(1).Build()
	s.TdigestReset().Key("1").Build()
	s.TdigestRevrank().Key("1").Value(1).Value(1).Build()
	s.TdigestTrimmedMean().Key("1").LowCutQuantile(1).HighCutQuantile(1).Build()
}

func TestCommand_InitSlot_tdigest(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { tdigest0(s) })
}

func TestCommand_NoSlot_tdigest(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { tdigest0(s) })
}
