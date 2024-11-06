package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func cleanChannel(out Out) {
	for range out {
	}
}

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, stage := range stages {
		newInChann := make(Bi)
		out := stage(in)
		go func() {
			defer cleanChannel(out)
			defer close(newInChann)
			for {
				select {
				case <-done:
					return
				case val, ok := <-out:
					if !ok {
						return
					}
					newInChann <- val
				}
			}
		}()
		in = newInChann
	}
	return in
}
