package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Ugh, off my one galore in this.
func TestTripletRangeSplit(t *testing.T) {
	base := triplet{20, 10, 5}
	//|------------|20--------24|--------|
	testCases := []struct {
		in                 rng
		out_overlapping    []rng
		out_notoverlapping []rng
	}{
		//|------------|20--------24|--------|
		{
			//|------------|20--------24|--------|
			//|-------------|21-------24|-------|
			in: rng{
				21, 4,
			},
			out_overlapping: []rng{
				{11, 4},
			},
			out_notoverlapping: nil,
		},
		{
			//|------------|20--------24|--------|
			//|--------------|21------24|-------|
			in: rng{
				22, 3,
			},
			out_overlapping: []rng{
				{12, 3},
			},
			out_notoverlapping: nil,
		},
		{
			//|------------|20--------24|--------|
			//|--------------|21-----23|---------|
			in: rng{
				22, 2,
			},
			out_overlapping: []rng{
				{12, 2},
			},
			out_notoverlapping: nil,
		},
		//|------------|20--------24|--------|
		//|-------------|21--------25|-------|
		{
			in: rng{
				21, 5,
			},
			out_overlapping: []rng{
				{11, 4},
			},
			out_notoverlapping: []rng{
				{25, 1},
			},
		},
		{
			//|------------|20--------24|--------|
			//|------------|20--------24|--------|
			in: rng{
				20, 5,
			},
			out_overlapping: []rng{
				{10, 5},
			},
			out_notoverlapping: nil,
		},
		//|------------|20--------24|--------|
		//|------------|20---------25|-------|
		{
			in: rng{
				20, 6,
			},
			out_overlapping: []rng{
				{10, 5},
			},
			out_notoverlapping: []rng{
				{25, 1},
			},
		},
		//|------------|20--------24|--------|
		//|-------------|21--------25|-------|
		{
			in: rng{
				21, 5,
			},
			out_overlapping: []rng{
				{11, 4},
			},
			out_notoverlapping: []rng{
				{25, 1},
			},
		},
		//|------------|20--------24|--------|
		//|-------------|21---------26|------|
		{
			in: rng{
				21, 6,
			},
			out_overlapping: []rng{
				{11, 4},
			},
			out_notoverlapping: []rng{
				{25, 2},
			},
		},
		//|------------|20--------24|--------|
		//|-|3-----18|-----------------------|
		{
			in: rng{
				3, 16,
			},
			out_overlapping: nil,
			out_notoverlapping: []rng{
				{3, 16},
			},
		},
		//|------------|20--------24|--------|
		//|--|3-----19|----------------------|
		{
			in: rng{
				3, 17,
			},
			out_overlapping: nil,
			out_notoverlapping: []rng{
				{3, 17},
			},
		},
		//|------------|20--------24|--------|
		//|--|3------20|---------------------|
		{
			in: rng{
				3, 18,
			},
			out_overlapping: []rng{
				{10, 1},
			},
			out_notoverlapping: []rng{
				{3, 17},
			},
		},
		//|------------|20--------24|--------|
		//|--|3-------21|--------------------|
		{
			in: rng{
				3, 19,
			},
			out_overlapping: []rng{
				{10, 2},
			},
			out_notoverlapping: []rng{
				{3, 17},
			},
		},
		//|---|20--------24|-----------------------|
		//|----------------|24------26|------------|
		{
			in: rng{
				24, 3,
			},
			out_overlapping: []rng{
				{14, 1},
			},
			out_notoverlapping: []rng{
				{25, 2},
			},
		},
		//|---|20--------24|-----------------------|
		//|------------------|25------27|----------|
		{
			in: rng{
				25, 3,
			},
			out_overlapping: nil,
			out_notoverlapping: []rng{
				{25, 3},
			},
		},
		//|---|20--------24|-----------------------|
		//|-12-------------------------------------|
		{
			in: rng{
				12, 1,
			},
			out_overlapping: nil,
			out_notoverlapping: []rng{
				{12, 1},
			},
		},
		//|---|20--------24|-----------------------|
		//|---|-20---------------------------------|
		{
			in: rng{
				20, 1,
			},
			out_overlapping: []rng{
				{10, 1},
			},
			out_notoverlapping: nil,
		},
		//|---|20--------24|-----------------------|
		//|-----|-22-------------------------------|
		{
			in: rng{
				22, 1,
			},
			out_overlapping: []rng{
				{12, 1},
			},
			out_notoverlapping: nil,
		},
		{
			in: rng{
				23, 1,
			},
			out_overlapping: []rng{
				{13, 1},
			},
			out_notoverlapping: nil,
		},
		{
			in: rng{
				24, 1,
			},
			out_overlapping: []rng{
				{14, 1},
			},
			out_notoverlapping: nil,
		},
		{
			in: rng{
				25, 1,
			},
			out_overlapping: nil,
			out_notoverlapping: []rng{
				{25, 1},
			},
		},
		{
			in: rng{
				26, 1,
			},
			out_overlapping: nil,
			out_notoverlapping: []rng{
				{26, 1},
			},
		},
	}

	for _, tc := range testCases {
		overlapping, not_overlapping := base.calcDestRanges([]rng{tc.in})
		assert.Equal(t, tc.out_overlapping, overlapping, tc)
		assert.Equal(t, tc.out_notoverlapping, not_overlapping, tc)
	}

}
