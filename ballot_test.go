package kshaka

import (
	"testing"
)

func Test_proposer_incBallot(t *testing.T) {

	tests := []struct {
		name string
		p    *proposerAcceptor
	}{
		{name: "increment ballot", p: newProposerAcceptor()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.p
			p.incBallot()
			p.incBallot()
			p.incBallot()

			if p.ballot.Counter != 3 {
				t.Errorf("\n p.incBallot() *3 \ngot = %#+v, \nwanted = %#+v", p.ballot.Counter, 3)
			}
		})
	}
}
