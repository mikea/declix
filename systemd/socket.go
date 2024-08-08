package systemd

import "mikea/declix/interfaces"

func (s *SocketImpl) ExpectedState() (interfaces.State, error) {
	return s.State.(*UnitStateImpl), nil
}

func (s *SocketImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	return determineUnitState(s, executor)
}

func (s *SocketImpl) DetermineAction(current interfaces.State, expected interfaces.State) (interfaces.Action, error) {
	return determineUnitAction(current, expected)
}

func (s *SocketImpl) RunAction(executor interfaces.CommandExecutor, action interfaces.Action, current interfaces.State, expectedState interfaces.State) error {
	return action.(*unitAction).Run(executor, s, current, expectedState)
}
