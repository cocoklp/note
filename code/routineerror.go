// killContainersWithSyncResult kills all pod's containers with sync results.
func (m *kubeGenericRuntimeManager) killContainersWithSyncResult(pod *v1.Pod, runningPod kubecontainer.Pod, gracePeriodOverride *int64) (syncResults []*kubecontainer.SyncResult) {
    containerResults := make(chan *kubecontainer.SyncResult, len(runningPod.Containers))
    wg := sync.WaitGroup{}

    wg.Add(len(runningPod.Containers))
    for _, container := range runningPod.Containers {
        go func(container *kubecontainer.Container) {
            defer utilruntime.HandleCrash()
            defer wg.Done()

            killContainerResult := kubecontainer.NewSyncResult(kubecontainer.KillContainer, container.Name)
            if err := m.killContainer(pod, container.ID, container.Name, "Need to kill Pod", gracePeriodOverride); err != nil {
                killContainerResult.Fail(kubecontainer.ErrKillContainer, err.Error())
            }
            containerResults <- killContainerResult
        }(container)
    }
    wg.Wait()
    close(containerResults)

    for containerResult := range containerResults {
        syncResults = append(syncResults, containerResult)
    }
    return
}


// killPodWithSyncResult kills a runningPod and returns SyncResult.
// Note: The pod passed in could be *nil* when kubelet restarted.
func (m *kubeGenericRuntimeManager) killPodWithSyncResult(pod *v1.Pod, runningPod kubecontainer.Pod, gracePeriodOverride *int64) (result kubecontainer.PodSyncResult) {
    killContainerResults := m.killContainersWithSyncResult(pod, runningPod, gracePeriodOverride)
    for _, containerResult := range killContainerResults {
        result.AddSyncResult(containerResult)
    }
}

// PodSyncResult is the summary result of SyncPod() and KillPod()
type PodSyncResult struct {
    // Result of different sync actions
    SyncResults []*SyncResult
    // Error encountered in SyncPod() and KillPod() that is not already included in SyncResults
    SyncError error
}

// AddSyncResult adds multiple SyncResult to current PodSyncResult
func (p *PodSyncResult) AddSyncResult(result ...*SyncResult) {
    p.SyncResults = append(p.SyncResults, result...)
}

// Error returns an error summarizing all the errors in PodSyncResult
func (p *PodSyncResult) Error() error {
    errlist := []error{}
    if p.SyncError != nil {
        errlist = append(errlist, fmt.Errorf("failed to SyncPod: %v\n", p.SyncError))
    }
    for _, result := range p.SyncResults {
        if result.Error != nil {
            errlist = append(errlist, fmt.Errorf("failed to %q for %q with %v: %q\n", result.Action, result.Target,
                result.Error, result.Message))
        }
    }
    return utilerrors.NewAggregate(errlist)
}
// NewAggregate converts a slice of errors into an Aggregate interface, which
// is itself an implementation of the error interface.  If the slice is empty,
// this returns nil.
// It will check if any of the element of input error list is nil, to avoid
// nil pointer panic when call Error().
func NewAggregate(errlist []error) Aggregate {
    if len(errlist) == 0 {
        return nil
    }
    // In case of input error list contains nil
    var errs []error
    for _, e := range errlist {
        if e != nil {
            errs = append(errs, e)
        }
    }
    if len(errs) == 0 {
        return nil
    }
    return aggregate(errs)
}

// This helper implements the error and Errors interfaces.  Keeping it private
// prevents people from making an aggregate of 0 errors, which is not
// an error, but does satisfy the error interface.
type aggregate []error

// Error is part of the error interface.
func (agg aggregate) Error() string {
    if len(agg) == 0 {
        // This should never happen, really.
        return ""
    }
    if len(agg) == 1 {
        return agg[0].Error()
    }
    result := fmt.Sprintf("[%s", agg[0].Error())
    for i := 1; i < len(agg); i++ {
        result += fmt.Sprintf(", %s", agg[i].Error())
    }
    result += "]"
    return result
}
