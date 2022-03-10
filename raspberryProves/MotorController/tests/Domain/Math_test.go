package DomainTest

import (
	"github.com/Enrikerf/pfm/commands/MotorController/app/Domain/Math"
	"github.com/stretchr/testify/assert"
	"testing"
)

/*

factors

we got the following entries:
x
inMin 	- im
inMax 	- iM
outMin	- om
outMax 	- oM

that give to us the equivalence classes

relation between im,om
relation between im,oM
relation between iM,om
relation between iM,oM

because im<iM and om<oM always

values could be positives (p) or negatives (p) and relations > or <



TEST CASES
case    v       inInt   rel     outInt  -1000 -100 -20 -10 10 20 100 1000
1       im      nn      <       nn      (-1000,-100) 	(-20,-10)
2       iM      np      <       np      no sense
3       v       pp      <       pp      (10,20)			(100,1000)
4       im      np      CI      pp		(-100,20) 		(10,1000)
5       iM      nn      CI      nn		(-1000,-20) 	(-100,-10)
6       v       nn      CI      np		(-1000,-20) 	(-100,20)
7       im      pp      C       np		(10,100) 		(-100,1000)
8       iM      nn      C       pp		no sense
9       v       np      C       nn		no sense
10      iM      pp      CD      nn		no sense
11      im      nn      CD      np		no sense
12      v       np      CD      pp		no sense
13      im      np      >       nn		(-10,100) 		(-100,-20)
14      iM      pp      >       np		(20,100) 		(-100,10)
15      v       nn      >       pp		no sense
16      im      pp      CD      pp		(10,100) 		(20,1000)


TEST CASES
case    v       inInt   rel     outInt  -1000 -100 -20 -10 10 20 100 1000
1       im      nn      <       nn      (-1000,-100) 	(-20,-10)
2       v       pp      <       pp      (10,20)			(100,1000)
3       im      np      CI      pp		(-100,20) 		(10,1000)
4       iM      nn      CI      nn		(-1000,-20) 	(-100,-10)
5       v       nn      CI      np		(-1000,-20) 	(-100,20)
6       im      pp      C       np		(10,100) 		(-100,1000)
7       im      np      >       nn		(-10,100) 		(-100,-20)
8       iM      pp      >       np		(20,100) 		(-100,10)
9       im      pp      CD      pp		(10,100) 		(20,1000)
*/

const (
	neg1000 = -1000
	neg100  = -100
	neg20   = -20
	neg10   = -10
	pos10   = 10
	pos20   = 20
	pos100  = 100
	pos1000 = 1000
)

var tests1 = []struct {
	x,
	inMin,
	inMax,
	outMin,
	outMax,
	result float64
}{
	{
		x:      neg1000,
		inMin:  neg1000,
		inMax:  neg100,
		outMin: neg20,
		outMax: neg10,
		result: neg20,
	}, // 0
	{
		x:      pos10 + 5,
		inMin:  pos10,
		inMax:  pos20,
		outMin: pos100,
		outMax: pos1000,
		result: 550,
	}, // 1
	{
		x:      neg100,
		inMin:  neg100,
		inMax:  pos20,
		outMin: pos10,
		outMax: pos1000,
		result: pos10,
	}, // 2
	{
		x:      neg20,
		inMin:  neg1000,
		inMax:  neg20,
		outMin: neg100,
		outMax: neg10,
		result: neg10,
	}, // 3
	{
		x:      neg1000 + 140,
		inMin:  neg1000,
		inMax:  neg20,
		outMin: neg100,
		outMax: pos20,
		result: -82.85714285714286,
	}, // 4
	{
		x:      pos10,
		inMin:  pos10,
		inMax:  pos100,
		outMin: neg100,
		outMax: pos1000,
		result: neg100,
	}, // 5
	{
		x:      neg10,
		inMin:  neg10,
		inMax:  pos100,
		outMin: neg100,
		outMax: neg20,
		result: neg100,
	}, // 6
	{
		x:      pos100,
		inMin:  pos20,
		inMax:  pos100,
		outMin: neg100,
		outMax: pos10,
		result: pos10,
	}, // 7
	{
		x:      pos10,
		inMin:  pos10,
		inMax:  pos100,
		outMin: pos20,
		outMax: pos1000,
		result: pos20,
	}, // 8
}

func TestMapBetweenRanges(t *testing.T) {
	for index, test := range tests1 {
		assert.Equal(t, test.result, Math.MapBetweenRanges(test.x, test.inMin, test.inMax, test.outMin, test.outMax), "must be equals in index: %d ", index)
	}
}
