package day21

import (
	"fmt"
	"github.com/nicolas-boisseau/AdventOfCode2022/common"
	"strconv"
	"strings"
)

type Monkey struct {
	left, right string
	name        string
	op          string
	yell        func() int
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
				yell: func() int {
					return number
				},
			})
		}
	}

	root := monkeys.Find("root")

	if !complex {
		return root.yell()
	} else {
		// change op for root (should return 0 when left and right are equals)
		root.yell = func() int {
			return monkeys.Find(root.left).yell() - monkeys.Find(root.right).yell()
		}

		mLeft := monkeys.Find(root.left)
		mRight := monkeys.Find(root.right)

		// Where is "humn" ?!
		humnSide := mLeft // default to left
		otherSide := mRight
		if mRight.IsInChilds(monkeys, "humn") {
			humnSide = mRight
			otherSide = mLeft
		}

		humn := monkeys.Find("humn")

		i := 0
		for humnSide != humn {
			i++
			mLeft = monkeys.Find(humnSide.left)
			mRight = monkeys.Find(humnSide.right)
			switch humnSide.op {
			case "+":

				if mRight.IsInChilds(monkeys, "humn") {
					humnSide = mRight
					newMonkey := &Monkey{
						name:  fmt.Sprintf("blabla%d", i),
						left:  otherSide.name,
						right: mLeft.name,
						op:    "-",
						yell:  performOp(monkeys, otherSide.name, mLeft.name, "-"),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				} else {
					humnSide = mLeft
					newMonkey := &Monkey{
						name:  fmt.Sprintf("blabla%d", i),
						left:  otherSide.name,
						right: mRight.name,
						op:    "-",
						yell:  performOp(monkeys, otherSide.name, mRight.name, "-"),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				}
			case "-":

				if mRight.IsInChilds(monkeys, "humn") {
					humnSide = mRight
					newMonkey := &Monkey{
						name:  fmt.Sprintf("blabla%d", i),
						left:  otherSide.name,
						right: mLeft.name,
						op:    "+",
						yell:  performOp(monkeys, otherSide.name, mLeft.name, "+"),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				} else {
					humnSide = mLeft
					newMonkey := &Monkey{
						name:  fmt.Sprintf("blabla%d", i),
						left:  otherSide.name,
						right: mRight.name,
						op:    "+",
						yell:  performOp(monkeys, otherSide.name, mRight.name, "+"),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				}
			case "/":

				if mRight.IsInChilds(monkeys, "humn") {
					humnSide = mRight
					newMonkey := &Monkey{
						name:  fmt.Sprintf("blabla%d", i),
						left:  otherSide.name,
						right: mLeft.name,
						op:    "*",
						yell:  performOp(monkeys, otherSide.name, mLeft.name, "*"),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				} else {
					humnSide = mLeft
					newMonkey := &Monkey{
						name:  fmt.Sprintf("blabla%d", i),
						left:  otherSide.name,
						right: mRight.name,
						op:    "*",
						yell:  performOp(monkeys, otherSide.name, mRight.name, "*"),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				}
			case "*":

				if mRight.IsInChilds(monkeys, "humn") {
					humnSide = mRight
					newMonkey := &Monkey{
						name:  fmt.Sprintf("blabla%d", i),
						left:  otherSide.name,
						right: mLeft.name,
						op:    "/",
						yell:  performOp(monkeys, otherSide.name, mLeft.name, "/"),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				} else {
					humnSide = mLeft
					newMonkey := &Monkey{
						name:  fmt.Sprintf("blabla%d", i),
						left:  otherSide.name,
						right: mRight.name,
						op:    "/",
						yell:  performOp(monkeys, otherSide.name, mRight.name, "/"),
					}
					otherSide = newMonkey
					monkeys.Add(newMonkey)
				}
			}
		}

		return otherSide.yell()
	}
}

func (m *Monkey) IsInChilds(monkeys *MonkeyList, name string) bool {
	if m.name == name {
		return true
	} else if m.left != "" && m.right != "" {
		return monkeys.Find(m.left).IsInChilds(monkeys, name) || monkeys.Find(m.right).IsInChilds(monkeys, name)
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

func performOp(monkeys *MonkeyList, left string, right string, op string) func() int {
	return func() int {
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
