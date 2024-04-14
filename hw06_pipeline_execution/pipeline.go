package hw06pipelineexecution

type (
	In    = <-chan interface{}
	Out   = In
	Bi    = chan interface{}
	Stage = func(in In) (out Out)
)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		in = Worker(done, stage(in))
	}
	return in
}

func Worker(done In, in In) Bi {
	out := make(Bi)
	go func() {
		defer close(out)
		for element := range in {
			select {
			case <-done:
				return
			default:
				out <- element
			}
		}
	}()
	return out
}
