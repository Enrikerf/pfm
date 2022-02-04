package UpdateTask

type Command struct {
	Uuid          string
	Host          OptionalString
	Port          OptionalString
	Mode          OptionalString
	Status        OptionalString
	ExecutionMode OptionalString
}

type OptionalString struct {
	Change bool
	Value  string
}
