package domain

type ScenarioUseCase interface {
	AddUserScenario() error
	TestScenario()
}
