package day21

import (
	"bytes"
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strconv"
	"strings"
)

type Monkey struct {
	left, right   string
	mLeft, mRight *Monkey
	name          string
	op            string
	yell          func() int64
	isHumn        bool
}

func (m *Monkey) String() string {
	buff := bytes.NewBufferString("")
	if m.mLeft != nil && m.mRight != nil {
		fmt.Fprint(buff, "(", m.mLeft.String(), m.op, m.mRight, ")")
	} else {
		if !m.isHumn {
			fmt.Fprintf(buff, "%d", m.yell())
		} else {
			fmt.Fprint(buff, "x")
		}

		//fmt.Fprintf(buff, "[%s]%d", m.name, m.yell())
	}
	return buff.String()
}

type MonkeyList struct {
	list []*Monkey
}

func newMonkeyList(capacity int) *MonkeyList {
	return &MonkeyList{
		list: make([]*Monkey, 0, capacity),
	}
}
func (ml *MonkeyList) Find(monkeyName string) *Monkey {
	for _, m := range ml.list {
		if m.name == monkeyName {
			return m
		}
	}
	return nil
}
func (ml *MonkeyList) Add(m *Monkey) {
	ml.list = append(ml.list, m)
}

func Process(fileName string, complex bool) int {
	lines := common.ReadLinesFromFile(fileName)

	// Create monkeys
	monkeys := newMonkeyList(len(lines))
	for _, line := range lines {

		splitted := strings.Split(line, ":")

		if strings.Contains(splitted[1], "+") {
			monkeys.Add(ExtractOpMonkey(splitted, monkeys, "+"))
		} else if strings.Contains(splitted[1], "-") {
			monkeys.Add(ExtractOpMonkey(splitted, monkeys, "-"))
		} else if strings.Contains(splitted[1], "/") {
			monkeys.Add(ExtractOpMonkey(splitted, monkeys, "/"))
		} else if strings.Contains(splitted[1], "*") {
			monkeys.Add(ExtractOpMonkey(splitted, monkeys, "*"))
		} else {
			number, err := strconv.Atoi(splitted[1][1:])
			if err != nil {
				panic("Cannot convert to number! " + err.Error())
			}
			monkeys.Add(&Monkey{
				name: splitted[0],
				yell: func() int64 {
					return int64(number)
				},
			})
		}
	}

	// evaluate monkeys
	for _, m := range monkeys.list {
		m.mLeft = monkeys.Find(m.left)
		m.mRight = monkeys.Find(m.right)
	}

	root := monkeys.Find("root")

	if !complex {
		return int(root.yell())
	} else {
		// change op for root (should return 0 when left and right are equals)
		root.yell = func() int64 {
			return root.mLeft.yell() - root.mRight.yell()
		}

		mLeft := root.mLeft
		mRight := root.mRight

		// Where is "humn" ?!
		humnSide := mLeft // default to left
		otherSide := mRight
		if mRight.IsInChilds(monkeys, "humn") {
			humnSide = mRight
			otherSide = mLeft
		}

		humn := monkeys.Find("humn")
		humn.isHumn = true

		i := 0
		for humnSide != humn {
			fmt.Println("==", i, "=======================")

			fmt.Println("HUMN SIDE:", humnSide)
			fmt.Println("OTHER SIDE:", otherSide)

			computedOtherSide := otherSide.yell()
			fmt.Println("Simplify other side to ", computedOtherSide)

			otherSide.mRight = nil
			otherSide.mLeft = nil
			otherSide.right = ""
			otherSide.left = ""
			otherSide.yell = func() int64 { return computedOtherSide }

			i++
			mLeft = humnSide.mLeft
			mRight = humnSide.mRight
			humnIsLeft := mLeft.IsInChilds(monkeys, "humn")

			if !humnIsLeft {
				computed := mLeft.yell()

				mLeft.mRight = nil
				mLeft.mLeft = nil
				mLeft.right = ""
				mLeft.left = ""
				mLeft.yell = func() int64 { return computed }
			} else {
				computed := mRight.yell()

				mRight.mRight = nil
				mRight.mLeft = nil
				mRight.right = ""
				mRight.left = ""
				mRight.yell = func() int64 { return computed }
			}

			switch humnSide.op {

			case "+":
				inverted := "-"
				humnSide, otherSide = CompensateOtherSide(mRight, monkeys, humnSide, i, otherSide, mLeft, inverted)
			case "-":
				inverted := "+"
				humnSide, otherSide = CompensateOtherSide(mRight, monkeys, humnSide, i, otherSide, mLeft, inverted)
			case "/":
				inverted := "*"
				//humnSide, otherSide = CompensateOtherSide(mRight, monkeys, humnSide, i, otherSide, mLeft, inverted)

				if mRight.IsInChilds(monkeys, "humn") {
					newMonkey := &Monkey{
						name:   fmt.Sprintf("blabla%d", i),
						left:   otherSide.name,
						right:  mRight.name,
						mLeft:  otherSide,
						mRight: mRight,
						op:     inverted,
						yell:   performOp(monkeys, otherSide.name, mRight.name, inverted),
					}
					humnSide = newMonkey
					otherSide = mLeft
					monkeys.Add(newMonkey)
				} else {
					humnSide = mLeft
					newMonkey := &Monkey{
						name:   fmt.Sprintf("blabla%d", i),
						left:   otherSide.name,
						right:  mRight.name,
						mLeft:  otherSide,
						mRight: mRight,
						op:     inverted,
						yell:   performOp(monkeys, otherSide.name, mRight.name, inverted),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				}
			case "*":
				inverted := "/"
				humnSide, otherSide = CompensateOtherSide(mRight, monkeys, humnSide, i, otherSide, mLeft, inverted)

			}
		}

		compensatingValue := otherSide.yell()

		humnSide.yell = func() int64 { return compensatingValue }
		fmt.Println("root.yell() (should be ZERO) = ", root.yell())

		return int(compensatingValue)
	}
}

func CompensateOtherSide(mRight *Monkey, monkeys *MonkeyList, humnSide *Monkey, i int, otherSide *Monkey, mLeft *Monkey, inverted string) (*Monkey, *Monkey) {
	humnIsRight := mRight.IsInChilds(monkeys, "humn")
	humnIsLeft := mLeft.IsInChilds(monkeys, "humn")
	fmt.Println("HumnIsLeft:", humnIsLeft, "HumnIsRight:", humnIsRight)
	if mRight.IsInChilds(monkeys, "humn") {
		humnSide = mRight
		newMonkey := &Monkey{
			name:   fmt.Sprintf("blabla%d", i),
			left:   otherSide.name,
			right:  mLeft.name,
			mLeft:  otherSide,
			mRight: mLeft,
			op:     inverted,
			yell:   performOp(monkeys, otherSide.name, mLeft.name, inverted),
		}
		otherSide = newMonkey
		monkeys.Add(newMonkey)
	} else {
		humnSide = mLeft
		newMonkey := &Monkey{
			name:   fmt.Sprintf("blabla%d", i),
			left:   otherSide.name,
			right:  mRight.name,
			mLeft:  otherSide,
			mRight: mRight,
			op:     inverted,
			yell:   performOp(monkeys, otherSide.name, mRight.name, inverted),
		}
		otherSide = newMonkey
		monkeys.Add(newMonkey)
	}
	return humnSide, otherSide
}

func (m *Monkey) IsInChilds(monkeys *MonkeyList, name string) bool {
	if m.name == name {
		return true
	} else if m.mLeft != nil && m.mRight != nil {
		return m.mLeft.IsInChilds(monkeys, name) || m.mRight.IsInChilds(monkeys, name)
	}
	return false
}

func ExtractOpMonkey(splitted []string, monkeys *MonkeyList, op string) *Monkey {
	var left, right string
	reader := strings.NewReader(splitted[1])
	format := "%s " + op + " %s"
	fmt.Fscanf(reader, format, &left, &right)

	return &Monkey{
		right: right,
		left:  left,
		name:  splitted[0],
		op:    op,
		yell:  performOp(monkeys, left, right, op),
	}
}

func performOp(monkeys *MonkeyList, left string, right string, op string) func() int64 {
	return func() int64 {
		mLeft := monkeys.Find(left)
		mRight := monkeys.Find(right)

		switch op {
		case "+":
			return mLeft.yell() + mRight.yell()
		case "-":
			return mLeft.yell() - mRight.yell()
		case "/":
			return mLeft.yell() / mRight.yell()
		case "*":
			return mLeft.yell() * mRight.yell()
		}
		panic("That should not happen!")
	}
}
