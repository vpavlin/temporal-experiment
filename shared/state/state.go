package state

var NewWorkflow = NewStatus("Workflow Started", "New")
var TransactionCreated = NewStatus("Transaction Created", "TxCreated")
var TransactionSent = NewStatus("Transaction sent to blockchain", "TxSent")
var EventReceived = NewStatus("Transaction Event received", "EventReceived")
var GasEstimated = NewStatus("Gas for Transaction estimated", "GasEstimated")
var PaymentCompleted = NewStatus("Payment completed successfully", "PaymentCompleted")

type Status struct {
	Message string
	Name    string
}

func NewStatus(message string, name string) Status {
	return Status{
		Message: message,
		Name:    name,
	}
}

type StatesCollection []Status

func (s *StatesCollection) Add(message string, name string) {
	s.Append(Status{
		Message: message,
		Name:    name,
	})
}

func (s *StatesCollection) Append(status Status) {
	*s = append(*s, status)
}
