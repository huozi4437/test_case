//求aSet集对所有bSet集的差集
package main

import (
	"fmt"
	"os"
	"strconv"
)

type PhonemeTask struct {
	BeginTime int
	EndTime   int
}

func DifferentSet(_req []int, _taskList [][]int) []*PhonemeTask {
	req := &PhonemeTask{BeginTime: _req[0], EndTime: _req[1]}
	taskList := make([]*PhonemeTask, 0)
	for _, _task := range _taskList {
		taskList = append(taskList, &PhonemeTask{BeginTime: _task[0], EndTime: _task[1]})
	}

	_phonemeTasks := make([]*PhonemeTask, 0)
	begin := req.BeginTime
	end := req.EndTime

	for _, task := range taskList {
		if begin >= task.BeginTime {
			if end <= task.EndTime {
				return _phonemeTasks
			}
			if begin < task.EndTime {
				begin = task.EndTime
			}
			continue
		} else {
			if end > task.EndTime {
				_task := &PhonemeTask{BeginTime: begin, EndTime: task.BeginTime}
				_phonemeTasks = append(_phonemeTasks, _task)
				begin = task.EndTime
			} else {
				if end > task.BeginTime {
					end = task.BeginTime
				}
				_task := &PhonemeTask{BeginTime: begin, EndTime: end}
				_phonemeTasks = append(_phonemeTasks, _task)

				return _phonemeTasks
			}
		}
	}
	_task := &PhonemeTask{BeginTime: begin, EndTime: end}
	_phonemeTasks = append(_phonemeTasks, _task)

	return _phonemeTasks
}

func main() {
	a, _ := strconv.Atoi(os.Args[1])
	b, _ := strconv.Atoi(os.Args[2])
	aSet := []int{a, b}
	bSet := [][]int{{5, 10}, {15, 20}, {25, 30}, {35, 40}}

	_phonemeTasks := DifferentSet(aSet, bSet)
	for _, subset := range _phonemeTasks {
		fmt.Printf("[%d, %d]\n", subset.BeginTime, subset.EndTime)
	}
}
