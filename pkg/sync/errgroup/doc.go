// Package errgroup provides synchronization, error propagation, and Context
// The errgroup package provides goroutine synchronization and error cancellation for a set of subtask goroutines.
//
// Errgroup contains three common ways
//
//1、Direct use At this point, all tasks will not be canceled because of a task failure:
//		g := &errgroup.Group{}
//		g.Go(func(ctx context.Context) {
//			// NOTE: At this point ctx is context.Background()
//			// do something
//		})
//
//2、WithContext When using WithContext does not cause all tasks to be canceled because of a task failure:
//		g := errgroup.WithContext(ctx)
//		g.Go(func(ctx context.Context) {
//			// NOTE: At this point ctx is ctx passed by errgroup.WithContext
//			// do something
//		})
//
//3、WithCancel When using WithCancel, if a person fails, all *not-executed or in-progress* tasks will be canceled:
//		g := errgroup.WithCancel(ctx)
//		g.Go(func(ctx context.Context) {
//			// NOTE: At this point ctx is a ctx derived from ctx passed by errgroup.WithContext
//			// do something
//		})
//
//Set the maximum number of parallels GOMAXPROCS works for all three modes of use
//NOTE: Due to errgroup implementation issues, setting errgroup for GOMAXPROCS requires immediate calls to Wait() :
//
//		g := errgroup.WithCancel(ctx)
//		g.GOMAXPROCS(2)
//		// task1
//		g.Go(func(ctx context.Context) {
//			fmt.Println("task1")
//		})
//		// task2
//		g.Go(func(ctx context.Context) {
//			fmt.Println("task2")
//		})
//		// task3
//		g.Go(func(ctx context.Context) {
//			fmt.Println("task3")
//		})
//		// NOTE: The GOMAXPROCS set at this time is 2, and three tasks have been added. task1, task2, task3 At this time, task3 will not run!
//		// There is only a chance to run if Wait task3 is called.
//		g.Wait() // Task3 running
package errgroup
