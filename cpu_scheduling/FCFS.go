package main

import "fmt"

type process struct {
	arrival_time    int
	brust_time      int
	waiting_time    int
	turnaround_time int
}

func findWaitingTime(Process []process) {

	var serviceTime = make([]int, len(Process))

	serviceTime[0] = Process[0].arrival_time
	for i := 1; i < len(Process); i++ {
		prev := Process[i-1]
		serviceTime[i] = serviceTime[i-1] + prev.brust_time
		Process[i].waiting_time = serviceTime[i] - Process[i].arrival_time

		if Process[i].waiting_time < 0 {
			Process[i].waiting_time = 0
		}
	}
}

func findTurnAroundTime(Process []process) {
	for i := 0; i < len(Process); i++ {
		Process[i].turnaround_time = Process[i].brust_time + Process[i].waiting_time
	}
}

func findAvgTime(processes []process) {
	findWaitingTime(processes)
	findTurnAroundTime(processes)
	fmt.Print("Process ", "Burst ", " Arrival ", "Waiting ", "Turn-Around ", "Completion \n")

	total_turnaround, total_waiting_time := 0, 0
	for i := 0; i < len(processes); i++ {
		total_waiting_time += processes[i].waiting_time
		total_turnaround += processes[i].turnaround_time
		completion_time := processes[i].turnaround_time + processes[i].arrival_time

		fmt.Println(" ", i+1, "\t\t", processes[i].brust_time, "\t\t", processes[i].arrival_time, "\t\t", processes[i].waiting_time, "\t\t", processes[i].turnaround_time, "\t\t", completion_time)
	}

	fmt.Println("Average waiting time: ", float64(total_waiting_time)/float64(len(processes)))
	fmt.Println("Average turn around time: ", float64(total_turnaround)/float64(len(processes)))
}

func main() {
	processes := []process{
		{brust_time: 5, arrival_time: 0},
		{brust_time: 9, arrival_time: 3},
		{brust_time: 6, arrival_time: 6},
	}

	findAvgTime(processes)
}
