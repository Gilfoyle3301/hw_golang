package hw06pipelineexecution

type (
	In    = <-chan interface{}
	Out   = In
	Bi    = chan interface{}
	Stage = func(in In) (out Out)
)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for i := 0; i < len(stages); i++ {
		in = Worker(done, stages[i](in))
	}
	return in
}

func Worker(done In, out Out) Bi {
	localOut := make(Bi)

	go func() {
		defer close(localOut)
	loop:
		for {
			select {
			case value, ok := <-out:
				if ok {
					localOut <- value
				} else {
					break loop
				}
			case <-done:
				return
			}
		}
	}()

	return localOut
}
