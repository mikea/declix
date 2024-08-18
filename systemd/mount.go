package systemd

import "mikea/declix/interfaces"

func (s *MountImpl) ExpectedState() (interfaces.State, error) {
	return s.State.(*UnitStateImpl), nil
}

func (s *MountImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	return determineUnitState(s, executor)
}

func (s *MountImpl) DetermineAction(current interfaces.State, expected interfaces.State) (interfaces.Action, error) {
	return determineUnitAction(current, expected)
}

func (s *MountImpl) RunAction(executor interfaces.CommandExecutor, action interfaces.Action, current interfaces.State, expectedState interfaces.State) error {
	return action.(*unitAction).Run(executor, s, current, expectedState)
}
