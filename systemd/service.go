package systemd

import (
	"mikea/declix/interfaces"
)

func (s *ServiceImpl) ExpectedState() (interfaces.State, error) {
	return s.State.(*UnitStateImpl), nil
}

func (s *ServiceImpl) DetermineState(executor interfaces.CommandExecutor) (interfaces.State, error) {
	return determineUnitState(s, executor)
}

func (s *ServiceImpl) DetermineAction(current interfaces.State, expected interfaces.State) (interfaces.Action, error) {
	return determineUnitAction(current, expected)
}

func (s *ServiceImpl) RunAction(executor interfaces.CommandExecutor, action interfaces.Action, current interfaces.State, expectedState interfaces.State) error {
	return action.(*unitAction).Run(executor, s, current, expectedState)
}
